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
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var ComputeRegionDiskIamSchema = map[string]*schema.Schema{
	"project": {
		Type:             schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew:         true,
	},
	"region": {
		Type:             schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew:         true,
	},
	"name": {
		Type:             schema.TypeString,
		Required: true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}


type ComputeRegionDiskIamUpdater struct {
	project string
	region string
	name string
	d       TerraformResourceData
	Config  *Config
}

func ComputeRegionDiskIamUpdaterProducer(d TerraformResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	region, _ := getRegion(d, config)
	if region != "" {
		if err := d.Set("region", region); err != nil {
			return nil, fmt.Errorf("Error setting region: %s", err)
		}
	}
	values["region"] = region
	if v, ok := d.GetOk("name"); ok {
		values["name"] = v.(string)
	}


	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/disks/(?P<name>[^/]+)","(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)","(?P<region>[^/]+)/(?P<name>[^/]+)","(?P<name>[^/]+)"}, d, config, d.Get("name").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ComputeRegionDiskIamUpdater{
		project: values["project"],
		region: values["region"],
		name: values["name"],
		d:       d,
		Config:  config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("region", u.region); err != nil {
		return nil, fmt.Errorf("Error setting region: %s", err)
	}
	if err := d.Set("name", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}

	return u, nil
}

func ComputeRegionDiskIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project	}

	region, _ := getRegion(d, config)
	if region != "" {
		values["region"] = region	}


	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/disks/(?P<name>[^/]+)","(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)","(?P<region>[^/]+)/(?P<name>[^/]+)","(?P<name>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
    values[k] = v
	}

	u := &ComputeRegionDiskIamUpdater{
		project: values["project"],
		region: values["region"],
		name: values["name"],
		d:       d,
		Config:  config,
	}
	if err := d.Set("name", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *ComputeRegionDiskIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyRegionDiskUrl("getIamPolicy")
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

	policy, err := sendRequest(u.Config, "GET", project, url, userAgent, obj)
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

func (u *ComputeRegionDiskIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}


	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyRegionDiskUrl("setIamPolicy")
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

func (u *ComputeRegionDiskIamUpdater) qualifyRegionDiskUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{ComputeBasePath}}%s/%s", fmt.Sprintf("projects/%s/regions/%s/disks/%s", u.project, u.region, u.name), methodIdentifier)
  url, err := replaceVars(u.d, u.Config, urlTemplate)
  if err != nil {
      return "", err
  }
  return url, nil
}

func (u *ComputeRegionDiskIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/regions/%s/disks/%s", u.project, u.region, u.name)
}

func (u *ComputeRegionDiskIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-compute-regiondisk-%s", u.GetResourceId())
}

func (u *ComputeRegionDiskIamUpdater) DescribeResource() string {
	return fmt.Sprintf("compute regiondisk %q", u.GetResourceId())
}
