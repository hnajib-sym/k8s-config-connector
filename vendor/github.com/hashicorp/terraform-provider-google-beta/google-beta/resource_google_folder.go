package google

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	resourceManagerV3 "google.golang.org/api/cloudresourcemanager/v3"
)

var activeFolderNotFoundError = errors.New("active folder not found")

func ResourceGoogleFolder() *schema.Resource {
	return &schema.Resource{
		Create: resourceGoogleFolderCreate,
		Read:   resourceGoogleFolderRead,
		Update: resourceGoogleFolderUpdate,
		Delete: resourceGoogleFolderDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGoogleFolderImportState,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Read:   schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"parent_org_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The organization id of the parent Organization. Exactly one of parent_org_id or parent_folder_id must be specified.`,
			},
			"parent_folder_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The folder id of the parent Folder. Exactly one of parent_org_id or parent_folder_id must be specified.`,
			},
			// Must be unique amongst its siblings.
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The folder's display name. A folder's display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.`,
			},
			"folder_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The folder id from the name "folders/{folder_id}"`,
			},
			// Format is 'folders/{folder_id}.
			// The terraform id holds the same value.
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the Folder. Its format is folders/{folder_id}.`,
			},
			"lifecycle_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The lifecycle state of the folder such as ACTIVE or DELETE_REQUESTED.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Timestamp when the Folder was created. Assigned by the server. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceGoogleFolderCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	displayName := d.Get("display_name").(string)
	parent, err := getParentID(d)
	if err != nil {
		return fmt.Errorf("Error getting parent for folder '%s': %s", displayName, err)
	}

	// Check if there's an ACTIVE folder with the given display_name in the
	// given parent first before trying to create a new folder. This allows
	// users to acquire existing folders by specifying the folder's
	// display_name and parent.
	folder, err := getActiveFolderByDisplayName(displayName, parent, userAgent, config)
	if err != nil && !errors.Is(err, activeFolderNotFoundError) {
		return fmt.Errorf("Error checking if folder '%s' in '%s' exists: %s", displayName, parent, err)
	} else if err == nil {
		// An ACTIVE folder with the given display_name in the given parent is found.
		d.SetId(folder.Name)
		return resourceGoogleFolderRead(d, meta)
	}

	var op *resourceManagerV3.Operation
	err = retryTimeDuration(func() error {
		var reqErr error
		op, reqErr = config.NewResourceManagerV3Client(userAgent).Folders.Create(&resourceManagerV3.Folder{
			DisplayName: displayName,
			Parent:      parent,
		}).Do()
		return reqErr
	}, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating folder '%s' in '%s': %s", displayName, parent, err)
	}

	opAsMap, err := ConvertToMap(op)
	if err != nil {
		return err
	}

	err = resourceManagerOperationWaitTime(config, opAsMap, "creating folder", userAgent, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating folder '%s' in '%s': %s", displayName, parent, err)
	}

	// Since we waited above, the operation is guaranteed to have been successful by this point.
	waitOp, err := config.NewResourceManagerClient(userAgent).Operations.Get(op.Name).Do()
	if err != nil {
		return fmt.Errorf("The folder '%s' has been created but we could not retrieve its id. Delete the folder manually and retry or use 'terraform import': %s", displayName, err)
	}

	// Requires 3 successive checks for safety. Nested IFs are used to avoid 3 error statement with the same message.
	var responseMap map[string]interface{}
	if err := json.Unmarshal(waitOp.Response, &responseMap); err == nil {
		if val, ok := responseMap["name"]; ok {
			if name, ok := val.(string); ok {
				d.SetId(name)
				return resourceGoogleFolderRead(d, meta)
			}
		}
	}
	return fmt.Errorf("The folder '%s' has been created but we could not retrieve its id. Delete the folder manually and retry or use 'terraform import'", displayName)
}

func resourceGoogleFolderRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	folder, err := getGoogleFolder(d.Id(), userAgent, d, config)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("Folder Not Found : %s", d.Id()))
	}

	// If the folder has been deleted from outside Terraform, remove it from state file.
	if folder.State != "ACTIVE" {
		log.Printf("[WARN] Removing folder '%s' because its state is '%s' (requires 'ACTIVE').", d.Id(), folder.State)
		d.SetId("")
		return nil
	}

	if err := d.Set("name", folder.Name); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	folderId := strings.TrimPrefix(folder.Name, "folders/")
	if err := d.Set("folder_id", folderId); err != nil {
		return fmt.Errorf("Error setting folder_id: %s", err)
	}
	if err := d.Set("display_name", folder.DisplayName); err != nil {
		return fmt.Errorf("Error setting display_name: %s", err)
	}
	if err := d.Set("lifecycle_state", folder.State); err != nil {
		return fmt.Errorf("Error setting lifecycle_state: %s", err)
	}
	if err := d.Set("create_time", folder.CreateTime); err != nil {
		return fmt.Errorf("Error setting create_time: %s", err)
	}

	if strings.HasPrefix(folder.Parent, "organizations/") {
		orgId := strings.TrimPrefix(folder.Parent, "organizations/")
		if err := d.Set("parent_org_id", orgId); err != nil {
			return fmt.Errorf("Error setting parent_org_id: %s", err)
		}
	} else if strings.HasPrefix(folder.Parent, "folders/") {
		folderId := strings.TrimPrefix(folder.Parent, "folders/")
		if err := d.Set("parent_folder_id", folderId); err != nil {
			return fmt.Errorf("Error setting parent_folder_id: %s", err)
		}
	} else {
		return fmt.Errorf("Error reading folder '%s' since its parent '%s' has an unrecognizable format.", folder.DisplayName, folder.Parent)
	}

	return nil
}

func resourceGoogleFolderUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	displayName := d.Get("display_name").(string)

	d.Partial(true)
	if d.HasChange("display_name") {
		err := retry(func() error {
			_, reqErr := config.NewResourceManagerV3Client(userAgent).Folders.Patch(d.Id(), &resourceManagerV3.Folder{
				DisplayName: displayName,
			}).Do()
			return reqErr
		})
		if err != nil {
			return fmt.Errorf("Error updating display_name to '%s': %s", displayName, err)
		}
	}

	if d.HasChange("parent_org_id") || d.HasChange("parent_folder_id") {
		newParent, err := getParentID(d)
		if err != nil {
			return fmt.Errorf("Error getting parent for folder '%s': %s", displayName, err)
		}

		var op *resourceManagerV3.Operation
		err = retry(func() error {
			var reqErr error
			op, reqErr = config.NewResourceManagerV3Client(userAgent).Folders.Move(d.Id(), &resourceManagerV3.MoveFolderRequest{
				DestinationParent: newParent,
			}).Do()
			return reqErr
		})
		if err != nil {
			return fmt.Errorf("Error moving folder '%s' to '%s': %s", displayName, newParent, err)
		}

		opAsMap, err := ConvertToMap(op)
		if err != nil {
			return err
		}

		err = resourceManagerOperationWaitTime(config, opAsMap, "move folder", userAgent, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error moving folder '%s' to '%s': %s", displayName, newParent, err)
		}
	}

	d.Partial(false)

	return nil
}

func resourceGoogleFolderDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	displayName := d.Get("display_name").(string)

	var op *resourceManagerV3.Operation
	err = retryTimeDuration(func() error {
		var reqErr error
		op, reqErr = config.NewResourceManagerV3Client(userAgent).Folders.Delete(d.Id()).Do()
		return reqErr
	}, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return fmt.Errorf("Error deleting folder '%s': %s", displayName, err)
	}

	opAsMap, err := ConvertToMap(op)
	if err != nil {
		return err
	}

	err = resourceManagerOperationWaitTime(config, opAsMap, "deleting folder", userAgent, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return fmt.Errorf("Error deleting folder '%s': %s", displayName, err)
	}

	return nil
}

func resourceGoogleFolderImportState(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	id := d.Id()

	if !strings.HasPrefix(d.Id(), "folders/") {
		id = fmt.Sprintf("folders/%s", id)
	}

	d.SetId(id)
	d.Set("name", id)

	return []*schema.ResourceData{d}, nil
}

// Util to get a Folder resource from API. Note that folder described by name is not necessarily the
// ResourceData resource.
func getGoogleFolder(folderName, userAgent string, d *schema.ResourceData, config *Config) (*resourceManagerV3.Folder, error) {
	var folder *resourceManagerV3.Folder
	err := retryTimeDuration(func() error {
		var reqErr error
		folder, reqErr = config.NewResourceManagerV3Client(userAgent).Folders.Get(folderName).Do()
		return reqErr
	}, d.Timeout(schema.TimeoutRead))
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func getActiveFolderByDisplayName(displayName, parent, userAgent string, config *Config) (*resourceManagerV3.Folder, error) {
	pageToken := ""
	for ok := true; ok; ok = pageToken != "" {
		query := fmt.Sprintf("state=ACTIVE AND parent=%s AND displayName=\"%s\"", parent, displayName)
		searchResponse, err := config.NewResourceManagerV3Client(userAgent).Folders.Search().Query(query).PageToken(pageToken).Do()
		if err != nil {
			if isGoogleApiErrorWithCode(err, 404) {
				return nil, activeFolderNotFoundError
			}
			return nil, fmt.Errorf("error searching for folders with query '%v': %v", query, err)
		}
		for _, folder := range searchResponse.Folders {
			if folder.DisplayName == displayName {
				return folder, nil
			}
		}
		pageToken = searchResponse.NextPageToken
	}
	return nil, activeFolderNotFoundError
}

func getParentID(d *schema.ResourceData) (string, error) {
	orgId := d.Get("parent_org_id").(string)
	folderId := d.Get("parent_folder_id").(string)

	if orgId != "" && folderId != "" {
		return "", fmt.Errorf("'parent_org_id' and 'parent_folder_id' cannot be both set.")
	}
	if orgId != "" {
		return "organizations/" + orgId, nil
	}
	if folderId != "" {
		return "folders/" + folderId, nil
	}
	return "", fmt.Errorf("exactly one of 'parent_org_id' or 'parent_folder_id' must be specified.")
}
