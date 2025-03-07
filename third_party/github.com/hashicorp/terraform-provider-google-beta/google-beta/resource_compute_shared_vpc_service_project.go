package google

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/googleapi"

	compute "google.golang.org/api/compute/v0.beta"
)

func ResourceComputeSharedVpcServiceProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeSharedVpcServiceProjectCreate,
		Read:   resourceComputeSharedVpcServiceProjectRead,
		Delete: resourceComputeSharedVpcServiceProjectDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"host_project": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of a host project to associate.`,
			},
			"service_project": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the project that will serve as a Shared VPC service project.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeSharedVpcServiceProjectCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	hostProject := d.Get("host_project").(string)
	serviceProject := d.Get("service_project").(string)

	req := &compute.ProjectsEnableXpnResourceRequest{
		XpnResource: &compute.XpnResourceId{
			Id:   serviceProject,
			Type: "PROJECT",
		},
	}
	op, err := config.NewComputeClient(userAgent).Projects.EnableXpnResource(hostProject, req).Do()
	if err != nil {
		return err
	}
	err = computeOperationWaitTime(config, op, hostProject, "Enabling Shared VPC Resource", userAgent, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", hostProject, serviceProject))

	return nil
}

func resourceComputeSharedVpcServiceProjectRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	split := strings.Split(d.Id(), "/")
	if len(split) != 2 {
		return fmt.Errorf("Error parsing resource ID %s", d.Id())
	}
	hostProject := split[0]
	serviceProject := split[1]

	associatedHostProject, err := config.NewComputeClient(userAgent).Projects.GetXpnHost(serviceProject).Do()
	if err != nil {
		log.Printf("[WARN] Removing shared VPC service. The service project is not associated with any host")

		d.SetId("")
		return nil
	}

	if hostProject != associatedHostProject.Name {
		log.Printf("[WARN] Removing shared VPC service. Expected associated host project to be '%s', got '%s'", hostProject, associatedHostProject.Name)
		d.SetId("")
		return nil
	}

	if err := d.Set("host_project", hostProject); err != nil {
		return fmt.Errorf("Error setting host_project: %s", err)
	}
	if err := d.Set("service_project", serviceProject); err != nil {
		return fmt.Errorf("Error setting service_project: %s", err)
	}

	return nil
}

func resourceComputeSharedVpcServiceProjectDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	hostProject := d.Get("host_project").(string)
	serviceProject := d.Get("service_project").(string)

	if err := disableXpnResource(d, config, hostProject, serviceProject); err != nil {
		// Don't fail if the service project is already disabled.
		if !isDisabledXpnResourceError(err) {
			return fmt.Errorf("Error disabling Shared VPC Resource %q: %s", serviceProject, err)
		}
	}

	return nil
}

func disableXpnResource(d *schema.ResourceData, config *Config, hostProject, project string) error {
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	req := &compute.ProjectsDisableXpnResourceRequest{
		XpnResource: &compute.XpnResourceId{
			Id:   project,
			Type: "PROJECT",
		},
	}
	op, err := config.NewComputeClient(userAgent).Projects.DisableXpnResource(hostProject, req).Do()
	if err != nil {
		return err
	}
	err = computeOperationWaitTime(config, op, hostProject, "Disabling Shared VPC Resource", userAgent, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	return nil
}

func isDisabledXpnResourceError(err error) bool {
	if gerr, ok := err.(*googleapi.Error); ok {
		if gerr.Code == 400 && len(gerr.Errors) > 0 && gerr.Errors[0].Reason == "invalidResourceUsage" {
			return true
		}
	}
	return false
}
