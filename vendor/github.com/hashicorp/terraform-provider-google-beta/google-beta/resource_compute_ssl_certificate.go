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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceComputeSslCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeSslCertificateCreate,
		Read:   resourceComputeSslCertificateRead,
		Delete: resourceComputeSslCertificateDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeSslCertificateImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"certificate": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The certificate in PEM format.
The certificate chain must be no greater than 5 certs long.
The chain must include at least one intermediate cert.`,
				Sensitive: true,
			},
			"private_key": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: sha256DiffSuppress,
				Description:      `The write-only private key in PEM format.`,
				Sensitive:        true,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"name": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateGCEName,
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.


These are in the same namespace as the managed SSL certificates.`,
			},
			"certificate_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The unique identifier for the resource.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"expire_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Expire time of the certificate in RFC3339 text format.`,
			},
			"name_prefix": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"name"},
				Description:   "Creates a unique name beginning with the specified prefix. Conflicts with name.",
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					// https://cloud.google.com/compute/docs/reference/latest/sslCertificates#resource
					// uuid is 26 characters, limit the prefix to 37.
					value := v.(string)
					if len(value) > 37 {
						errors = append(errors, fmt.Errorf(
							"%q cannot be longer than 37 characters, name is limited to 63", k))
					}
					return
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeSslCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	certificateProp, err := expandComputeSslCertificateCertificate(d.Get("certificate"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("certificate"); !isEmptyValue(reflect.ValueOf(certificateProp)) && (ok || !reflect.DeepEqual(v, certificateProp)) {
		obj["certificate"] = certificateProp
	}
	descriptionProp, err := expandComputeSslCertificateDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeSslCertificateName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	privateKeyProp, err := expandComputeSslCertificatePrivateKey(d.Get("private_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("private_key"); !isEmptyValue(reflect.ValueOf(privateKeyProp)) && (ok || !reflect.DeepEqual(v, privateKeyProp)) {
		obj["privateKey"] = privateKeyProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/sslCertificates")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new SslCertificate: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SslCertificate: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating SslCertificate: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/global/sslCertificates/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating SslCertificate", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create SslCertificate: %s", err)
	}

	log.Printf("[DEBUG] Finished creating SslCertificate %q: %#v", d.Id(), res)

	return resourceComputeSslCertificateRead(d, meta)
}

func resourceComputeSslCertificateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/sslCertificates/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SslCertificate: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeSslCertificate %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading SslCertificate: %s", err)
	}

	if err := d.Set("certificate", flattenComputeSslCertificateCertificate(res["certificate"], d, config)); err != nil {
		return fmt.Errorf("Error reading SslCertificate: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeSslCertificateCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading SslCertificate: %s", err)
	}
	if err := d.Set("description", flattenComputeSslCertificateDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading SslCertificate: %s", err)
	}
	if err := d.Set("expire_time", flattenComputeSslCertificateExpireTime(res["expireTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading SslCertificate: %s", err)
	}
	if err := d.Set("certificate_id", flattenComputeSslCertificateCertificateId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading SslCertificate: %s", err)
	}
	if err := d.Set("name", flattenComputeSslCertificateName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading SslCertificate: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading SslCertificate: %s", err)
	}

	return nil
}

func resourceComputeSslCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SslCertificate: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/sslCertificates/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting SslCertificate %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "SslCertificate")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting SslCertificate", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting SslCertificate %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeSslCertificateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/sslCertificates/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/global/sslCertificates/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeSslCertificateCertificate(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSslCertificateCreationTimestamp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSslCertificateDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSslCertificateExpireTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSslCertificateCertificateId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := stringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeSslCertificateName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandComputeSslCertificateCertificate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSslCertificateDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSslCertificateName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	var certName string
	if v, ok := d.GetOk("name"); ok {
		certName = v.(string)
	} else if v, ok := d.GetOk("name_prefix"); ok {
		certName = resource.PrefixedUniqueId(v.(string))
	} else {
		certName = resource.UniqueId()
	}

	// We need to get the {{name}} into schema to set the ID using ReplaceVars
	if err := d.Set("name", certName); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}

	return certName, nil
}

func expandComputeSslCertificatePrivateKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
