package google

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/serviceusage/v1"
)

// These services can only be enabled as a side-effect of enabling other services,
// so don't bother storing them in the config or using them for diffing.
var ignoredProjectServices = []string{"dataproc-control.googleapis.com", "source.googleapis.com", "stackdriverprovisioning.googleapis.com"}
var ignoredProjectServicesSet = golangSetFromStringSlice(ignoredProjectServices)

// Services that can't be user-specified but are otherwise valid. Renamed
// services should be added to this set during major releases.
var bannedProjectServices = []string{"bigquery-json.googleapis.com"}

// Service Renames
// we expect when a service is renamed:
// - both service names will continue to be able to be set
// - setting one will effectively enable the other as a dependent
// - GET will return whichever service name is requested
// - LIST responses will not contain the old service name
// renames may be reverted, though, so we should canonicalise both ways until
// the old service is fully removed from the provider
//
// We handle service renames in the provider by pretending that we've read both
// the old and new service names from the API if we see either, and only setting
// the one(s) that existed in prior state in config (if any). If neither exists,
// we'll set the old service name in state.
// Additionally, in case of service rename rollbacks or unexpected early
// removals of services, if we fail to create or delete a service that's been
// renamed we'll retry using an alternate name.
// We try creation by the user-specified value followed by the other value.
// We try deletion by the old value followed by the new value.

// map from old -> new names of services that have been renamed
// these should be removed during major provider versions. comment here with
// "DEPRECATED FOR {{version}} next to entries slated for removal in {{version}}
// upon removal, we should disallow the old name from being used even if it's
// not gone from the underlying API yet
var renamedServices = map[string]string{}

// renamedServices in reverse (new -> old)
var renamedServicesByNewServiceNames = reverseStringMap(renamedServices)

// renamedServices expressed as both old -> new and new -> old
var renamedServicesByOldAndNewServiceNames = mergeStringMaps(renamedServices, renamedServicesByNewServiceNames)

const maxServiceUsageBatchSize = 20

func validateProjectServiceService(val interface{}, key string) (warns []string, errs []error) {
	bannedServicesFunc := StringNotInSlice(append(ignoredProjectServices, bannedProjectServices...), false)
	warns, errs = bannedServicesFunc(val, key)
	if len(errs) > 0 {
		return
	}

	// StringNotInSlice already validates that this is a string
	v, _ := val.(string)
	if !strings.Contains(v, ".") {
		errs = append(errs, fmt.Errorf("expected %s to be a domain like serviceusage.googleapis.com", v))
	}
	return
}

func ResourceGoogleProjectService() *schema.Resource {
	return &schema.Resource{
		Create: resourceGoogleProjectServiceCreate,
		Read:   resourceGoogleProjectServiceRead,
		Delete: resourceGoogleProjectServiceDelete,
		Update: resourceGoogleProjectServiceUpdate,

		Importer: &schema.ResourceImporter{
			State: resourceGoogleProjectServiceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateProjectServiceService,
			},
			"project": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareResourceNames,
			},

			"disable_dependent_services": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"disable_on_destroy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceGoogleProjectServiceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/services/(?P<service>[^/]+)",
		"(?P<project>[^/]+)/(?P<service>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/services/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func resourceGoogleProjectServiceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	project = GetResourceNameFromSelfLink(project)

	srv := d.Get("service").(string)
	id := project + "/" + srv

	// Check if the service has already been enabled
	servicesRaw, err := BatchRequestReadServices(project, d, config)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("Project Service %s", d.Id()))
	}
	servicesList := servicesRaw.(map[string]struct{})
	if _, ok := servicesList[srv]; ok {
		log.Printf("[DEBUG] service %s was already found to be enabled in project %s", srv, project)
		d.SetId(id)
		if err := d.Set("project", project); err != nil {
			return fmt.Errorf("Error setting project: %s", err)
		}
		if err := d.Set("service", srv); err != nil {
			return fmt.Errorf("Error setting service: %s", err)
		}
		return nil
	}

	err = BatchRequestEnableService(srv, project, d, config)
	if err != nil {
		return err
	}
	d.SetId(id)
	return resourceGoogleProjectServiceRead(d, meta)
}

func resourceGoogleProjectServiceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	project = GetResourceNameFromSelfLink(project)

	servicesRaw, err := BatchRequestReadServices(project, d, config)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("Project Service %s", d.Id()))
	}
	servicesList := servicesRaw.(map[string]struct{})

	srv := d.Get("service").(string)
	if _, ok := servicesList[srv]; ok {
		if err := d.Set("project", project); err != nil {
			return fmt.Errorf("Error setting project: %s", err)
		}
		if err := d.Set("service", srv); err != nil {
			return fmt.Errorf("Error setting service: %s", err)
		}
		return nil
	}

	log.Printf("[DEBUG] service %s not in enabled services for project %s, removing from state", srv, project)
	d.SetId("")
	return nil
}

func resourceGoogleProjectServiceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	if disable := d.Get("disable_on_destroy"); !(disable.(bool)) {
		log.Printf("[WARN] Project service %q disable_on_destroy is false, skip disabling service", d.Id())
		d.SetId("")
		return nil
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	project = GetResourceNameFromSelfLink(project)

	service := d.Get("service").(string)
	disableDependencies := d.Get("disable_dependent_services").(bool)
	if err = disableServiceUsageProjectService(service, project, d, config, disableDependencies); err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("Project Service %s", d.Id()))
	}

	d.SetId("")
	return nil
}

func resourceGoogleProjectServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	// This update method is no-op because the only updatable fields
	// are state/config-only, i.e. they aren't sent in requests to the API.
	return nil
}

// Disables a project service.
func disableServiceUsageProjectService(service, project string, d *schema.ResourceData, config *Config, disableDependentServices bool) error {
	err := retryTimeDuration(func() error {
		billingProject := project
		userAgent, err := generateUserAgentString(d, config.userAgent)
		if err != nil {
			return err
		}
		name := fmt.Sprintf("projects/%s/services/%s", project, service)
		servicesDisableCall := config.NewServiceUsageClient(userAgent).Services.Disable(name, &serviceusage.DisableServiceRequest{
			DisableDependentServices: disableDependentServices,
		})
		if config.UserProjectOverride {
			// err == nil indicates that the billing_project value was found
			if bp, err := getBillingProject(d, config); err == nil {
				billingProject = bp
			}
			servicesDisableCall.Header().Add("X-Goog-User-Project", billingProject)
		}
		sop, err := servicesDisableCall.Do()
		if err != nil {
			return err
		}
		// Wait for the operation to complete
		waitErr := serviceUsageOperationWait(config, sop, billingProject, "api to disable", userAgent, d.Timeout(schema.TimeoutDelete))
		if waitErr != nil {
			return waitErr
		}
		return nil
	}, d.Timeout(schema.TimeoutDelete), serviceUsageServiceBeingActivated)
	if err != nil {
		return fmt.Errorf("Error disabling service %q for project %q: %v", service, project, err)
	}
	return nil
}
