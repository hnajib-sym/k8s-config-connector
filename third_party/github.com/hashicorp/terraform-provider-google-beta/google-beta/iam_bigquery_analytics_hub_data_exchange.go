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

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var BigqueryAnalyticsHubDataExchangeIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"location": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"data_exchange_id": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type BigqueryAnalyticsHubDataExchangeIamUpdater struct {
	project        string
	location       string
	dataExchangeId string
	d              TerraformResourceData
	Config         *Config
}

func BigqueryAnalyticsHubDataExchangeIamUpdaterProducer(d TerraformResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	location, _ := getLocation(d, config)
	if location != "" {
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	values["location"] = location
	if v, ok := d.GetOk("data_exchange_id"); ok {
		values["data_exchange_id"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/dataExchanges/(?P<data_exchange_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<data_exchange_id>[^/]+)", "(?P<location>[^/]+)/(?P<data_exchange_id>[^/]+)", "(?P<data_exchange_id>[^/]+)"}, d, config, d.Get("data_exchange_id").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &BigqueryAnalyticsHubDataExchangeIamUpdater{
		project:        values["project"],
		location:       values["location"],
		dataExchangeId: values["data_exchange_id"],
		d:              d,
		Config:         config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("location", u.location); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("data_exchange_id", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting data_exchange_id: %s", err)
	}

	return u, nil
}

func BigqueryAnalyticsHubDataExchangeIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	location, _ := getLocation(d, config)
	if location != "" {
		values["location"] = location
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/dataExchanges/(?P<data_exchange_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<data_exchange_id>[^/]+)", "(?P<location>[^/]+)/(?P<data_exchange_id>[^/]+)", "(?P<data_exchange_id>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &BigqueryAnalyticsHubDataExchangeIamUpdater{
		project:        values["project"],
		location:       values["location"],
		dataExchangeId: values["data_exchange_id"],
		d:              d,
		Config:         config,
	}
	if err := d.Set("data_exchange_id", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting data_exchange_id: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *BigqueryAnalyticsHubDataExchangeIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyDataExchangeUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := generateUserAgentString(u.d, u.Config.userAgent)
	if err != nil {
		return nil, err
	}

	policy, err := sendRequest(u.Config, "POST", project, url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *BigqueryAnalyticsHubDataExchangeIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyDataExchangeUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.userAgent)
	if err != nil {
		return err
	}

	_, err = sendRequestWithTimeout(u.Config, "POST", project, url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *BigqueryAnalyticsHubDataExchangeIamUpdater) qualifyDataExchangeUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{BigqueryAnalyticsHubBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", u.project, u.location, u.dataExchangeId), methodIdentifier)
	url, err := replaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *BigqueryAnalyticsHubDataExchangeIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", u.project, u.location, u.dataExchangeId)
}

func (u *BigqueryAnalyticsHubDataExchangeIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-bigqueryanalyticshub-dataexchange-%s", u.GetResourceId())
}

func (u *BigqueryAnalyticsHubDataExchangeIamUpdater) DescribeResource() string {
	return fmt.Sprintf("bigqueryanalyticshub dataexchange %q", u.GetResourceId())
}
