// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAccessApprovalProjectSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessApprovalProjectSettingsCreate,
		Read:   resourceAccessApprovalProjectSettingsRead,
		Update: resourceAccessApprovalProjectSettingsUpdate,
		Delete: resourceAccessApprovalProjectSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessApprovalProjectSettingsImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"enrolled_services": {
				Type:     schema.TypeSet,
				Required: true,
				Description: `A list of Google Cloud Services for which the given resource has Access Approval enrolled.
Access requests for the resource given by name against any of these services contained here will be required
to have explicit approval. Enrollment can only be done on an all or nothing basis.

A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.`,
				Elem: accessapprovalProjectSettingsEnrolledServicesSchema(),
				Set:  accessApprovalEnrolledServicesHash,
			},
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `ID of the project of the access approval settings.`,
			},
			"active_key_version": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The asymmetric crypto key version to use for signing approval requests.
Empty active_key_version indicates that a Google-managed key should be used for signing.
This property will be ignored if set by an ancestor of the resource, and new non-empty values may not be set.`,
			},
			"notification_emails": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				Description: `A list of email addresses to which notifications relating to approval requests should be sent.
Notifications relating to a resource will be sent to all emails in the settings of ancestor
resources of that resource. A maximum of 50 email addresses are allowed.`,
				MaxItems: 50,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Deprecated:  "Deprecated in favor of `project_id`",
				Description: `Deprecated in favor of 'project_id'`,
			},
			"ancestor_has_active_key_version": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `If the field is true, that indicates that an ancestor of this Project has set active_key_version.`,
			},
			"enrolled_ancestor": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors of the Project.`,
			},
			"invalid_key_version": {
				Type:     schema.TypeBool,
				Computed: true,
				Description: `If the field is true, that indicates that there is some configuration issue with the active_key_version
configured on this Project (e.g. it doesn't exist or the Access Approval service account doesn't have the
correct permissions on it, etc.) This key version is not necessarily the effective key version at this level,
as key versions are inherited top-down.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the settings. Format is "projects/{project_id}/accessApprovalSettings"`,
			},
		},
		UseJSONNumber: true,
	}
}

func accessapprovalProjectSettingsEnrolledServicesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_product": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The product for which Access Approval will be enrolled. Allowed values are listed (case-sensitive):
  all
  appengine.googleapis.com
  bigquery.googleapis.com
  bigtable.googleapis.com
  cloudkms.googleapis.com
  compute.googleapis.com
  dataflow.googleapis.com
  iam.googleapis.com
  pubsub.googleapis.com
  storage.googleapis.com`,
			},
			"enrollment_level": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateEnum([]string{"BLOCK_ALL", ""}),
				Description:  `The enrollment level of the service. Default value: "BLOCK_ALL" Possible values: ["BLOCK_ALL"]`,
				Default:      "BLOCK_ALL",
			},
		},
	}
}

func resourceAccessApprovalProjectSettingsCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	notificationEmailsProp, err := expandAccessApprovalProjectSettingsNotificationEmails(d.Get("notification_emails"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_emails"); !isEmptyValue(reflect.ValueOf(notificationEmailsProp)) && (ok || !reflect.DeepEqual(v, notificationEmailsProp)) {
		obj["notificationEmails"] = notificationEmailsProp
	}
	enrolledServicesProp, err := expandAccessApprovalProjectSettingsEnrolledServices(d.Get("enrolled_services"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enrolled_services"); !isEmptyValue(reflect.ValueOf(enrolledServicesProp)) && (ok || !reflect.DeepEqual(v, enrolledServicesProp)) {
		obj["enrolledServices"] = enrolledServicesProp
	}
	activeKeyVersionProp, err := expandAccessApprovalProjectSettingsActiveKeyVersion(d.Get("active_key_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("active_key_version"); !isEmptyValue(reflect.ValueOf(activeKeyVersionProp)) && (ok || !reflect.DeepEqual(v, activeKeyVersionProp)) {
		obj["activeKeyVersion"] = activeKeyVersionProp
	}
	projectProp, err := expandAccessApprovalProjectSettingsProject(d.Get("project"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("project"); !isEmptyValue(reflect.ValueOf(projectProp)) && (ok || !reflect.DeepEqual(v, projectProp)) {
		obj["project"] = projectProp
	}

	url, err := replaceVars(d, config, "{{AccessApprovalBasePath}}projects/{{project_id}}/accessApprovalSettings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ProjectSettings: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	updateMask := []string{}

	if d.HasChange("notification_emails") {
		updateMask = append(updateMask, "notificationEmails")
	}

	if d.HasChange("enrolled_services") {
		updateMask = append(updateMask, "enrolledServices")
	}

	if d.HasChange("active_key_version") {
		updateMask = append(updateMask, "activeKeyVersion")
	}

	if d.HasChange("project") {
		updateMask = append(updateMask, "project")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ProjectSettings: %s", err)
	}
	if err := d.Set("name", flattenAccessApprovalProjectSettingsName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project_id}}/accessApprovalSettings")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ProjectSettings %q: %#v", d.Id(), res)

	return resourceAccessApprovalProjectSettingsRead(d, meta)
}

func resourceAccessApprovalProjectSettingsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{AccessApprovalBasePath}}projects/{{project_id}}/accessApprovalSettings")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AccessApprovalProjectSettings %q", d.Id()))
	}

	if err := d.Set("name", flattenAccessApprovalProjectSettingsName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectSettings: %s", err)
	}
	if err := d.Set("notification_emails", flattenAccessApprovalProjectSettingsNotificationEmails(res["notificationEmails"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectSettings: %s", err)
	}
	if err := d.Set("enrolled_services", flattenAccessApprovalProjectSettingsEnrolledServices(res["enrolledServices"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectSettings: %s", err)
	}
	if err := d.Set("enrolled_ancestor", flattenAccessApprovalProjectSettingsEnrolledAncestor(res["enrolledAncestor"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectSettings: %s", err)
	}
	if err := d.Set("active_key_version", flattenAccessApprovalProjectSettingsActiveKeyVersion(res["activeKeyVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectSettings: %s", err)
	}
	if err := d.Set("ancestor_has_active_key_version", flattenAccessApprovalProjectSettingsAncestorHasActiveKeyVersion(res["ancestorHasActiveKeyVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectSettings: %s", err)
	}
	if err := d.Set("invalid_key_version", flattenAccessApprovalProjectSettingsInvalidKeyVersion(res["invalidKeyVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectSettings: %s", err)
	}
	if err := d.Set("project", flattenAccessApprovalProjectSettingsProject(res["project"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectSettings: %s", err)
	}

	return nil
}

func resourceAccessApprovalProjectSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	notificationEmailsProp, err := expandAccessApprovalProjectSettingsNotificationEmails(d.Get("notification_emails"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_emails"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, notificationEmailsProp)) {
		obj["notificationEmails"] = notificationEmailsProp
	}
	enrolledServicesProp, err := expandAccessApprovalProjectSettingsEnrolledServices(d.Get("enrolled_services"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enrolled_services"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enrolledServicesProp)) {
		obj["enrolledServices"] = enrolledServicesProp
	}
	activeKeyVersionProp, err := expandAccessApprovalProjectSettingsActiveKeyVersion(d.Get("active_key_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("active_key_version"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, activeKeyVersionProp)) {
		obj["activeKeyVersion"] = activeKeyVersionProp
	}
	projectProp, err := expandAccessApprovalProjectSettingsProject(d.Get("project"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("project"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, projectProp)) {
		obj["project"] = projectProp
	}

	url, err := replaceVars(d, config, "{{AccessApprovalBasePath}}projects/{{project_id}}/accessApprovalSettings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ProjectSettings %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("notification_emails") {
		updateMask = append(updateMask, "notificationEmails")
	}

	if d.HasChange("enrolled_services") {
		updateMask = append(updateMask, "enrolledServices")
	}

	if d.HasChange("active_key_version") {
		updateMask = append(updateMask, "activeKeyVersion")
	}

	if d.HasChange("project") {
		updateMask = append(updateMask, "project")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating ProjectSettings %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ProjectSettings %q: %#v", d.Id(), res)
	}

	return resourceAccessApprovalProjectSettingsRead(d, meta)
}

func resourceAccessApprovalProjectSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["notificationEmails"] = []string{}
	obj["enrolledServices"] = []string{}
	obj["activeKeyVersion"] = ""

	url, err := replaceVars(d, config, "{{AccessApprovalBasePath}}projects/{{project_id}}/accessApprovalSettings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Emptying ProjectSettings %q: %#v", d.Id(), obj)
	updateMask := []string{}

	updateMask = append(updateMask, "notificationEmails")
	updateMask = append(updateMask, "enrolledServices")
	updateMask = append(updateMask, "activeKeyVersion")

	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	res, err := sendRequestWithTimeout(config, "PATCH", "", url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error emptying ProjectSettings %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished emptying ProjectSettings %q: %#v", d.Id(), res)
	}

	return nil
}

func resourceAccessApprovalProjectSettingsImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project_id>[^/]+)/accessApprovalSettings",
		"(?P<project_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project_id}}/accessApprovalSettings")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAccessApprovalProjectSettingsName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsNotificationEmails(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenAccessApprovalProjectSettingsEnrolledServices(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(accessApprovalEnrolledServicesHash, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"cloud_product":    flattenAccessApprovalProjectSettingsEnrolledServicesCloudProduct(original["cloudProduct"], d, config),
			"enrollment_level": flattenAccessApprovalProjectSettingsEnrolledServicesEnrollmentLevel(original["enrollmentLevel"], d, config),
		})
	}
	return transformed
}
func flattenAccessApprovalProjectSettingsEnrolledServicesCloudProduct(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsEnrolledServicesEnrollmentLevel(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsEnrolledAncestor(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsActiveKeyVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsAncestorHasActiveKeyVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsInvalidKeyVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsProject(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandAccessApprovalProjectSettingsNotificationEmails(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandAccessApprovalProjectSettingsEnrolledServices(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedCloudProduct, err := expandAccessApprovalProjectSettingsEnrolledServicesCloudProduct(original["cloud_product"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCloudProduct); val.IsValid() && !isEmptyValue(val) {
			transformed["cloudProduct"] = transformedCloudProduct
		}

		transformedEnrollmentLevel, err := expandAccessApprovalProjectSettingsEnrolledServicesEnrollmentLevel(original["enrollment_level"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEnrollmentLevel); val.IsValid() && !isEmptyValue(val) {
			transformed["enrollmentLevel"] = transformedEnrollmentLevel
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAccessApprovalProjectSettingsEnrolledServicesCloudProduct(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessApprovalProjectSettingsEnrolledServicesEnrollmentLevel(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessApprovalProjectSettingsActiveKeyVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessApprovalProjectSettingsProject(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
