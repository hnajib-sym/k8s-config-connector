package google

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	compute "google.golang.org/api/compute/v0.beta"
)

func ResourceComputeProjectDefaultNetworkTier() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeProjectDefaultNetworkTierCreateOrUpdate,
		Read:   resourceComputeProjectDefaultNetworkTierRead,
		Update: resourceComputeProjectDefaultNetworkTierCreateOrUpdate,
		Delete: resourceComputeProjectDefaultNetworkTierDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
		},

		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"network_tier": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  `The default network tier to be configured for the project. This field can take the following values: PREMIUM or STANDARD.`,
				ValidateFunc: validation.StringInSlice([]string{"PREMIUM", "STANDARD"}, false),
			},

			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeProjectDefaultNetworkTierCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	projectID, err := getProject(d, config)
	if err != nil {
		return err
	}

	request := &compute.ProjectsSetDefaultNetworkTierRequest{
		NetworkTier: d.Get("network_tier").(string),
	}
	op, err := config.NewComputeClient(userAgent).Projects.SetDefaultNetworkTier(projectID, request).Do()
	if err != nil {
		return fmt.Errorf("SetDefaultNetworkTier failed: %s", err)
	}

	log.Printf("[DEBUG] SetDefaultNetworkTier: %d (%s)", op.Id, op.SelfLink)
	err = computeOperationWaitTime(config, op, projectID, "SetDefaultNetworkTier", userAgent, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("SetDefaultNetworkTier failed: %s", err)
	}

	d.SetId(projectID)

	return resourceComputeProjectDefaultNetworkTierRead(d, meta)
}

func resourceComputeProjectDefaultNetworkTierRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	projectId := d.Id()

	project, err := config.NewComputeClient(userAgent).Projects.Get(projectId).Do()
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("Project data for project %q", projectId))
	}

	err = d.Set("network_tier", project.DefaultNetworkTier)
	if err != nil {
		return fmt.Errorf("Error setting default network tier: %s", err)
	}

	if err := d.Set("project", projectId); err != nil {
		return fmt.Errorf("Error setting project: %s", err)
	}

	return nil
}

func resourceComputeProjectDefaultNetworkTierDelete(d *schema.ResourceData, meta interface{}) error {

	log.Printf("[WARNING] Default Network Tier will be only removed from Terraform state, but will be left intact on GCP.")

	return schema.RemoveFromState(d, meta)
}
