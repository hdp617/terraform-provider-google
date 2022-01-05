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
  "testing"

  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
  "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccMonitoringGroup_monitoringGroupBasicExample(t *testing.T) {
  t.Parallel()

	context := map[string]interface{} {
				"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
				Providers: testAccProviders,
								CheckDestroy: testAccCheckMonitoringGroupDestroyProducer(t),
				Steps: []resource.TestStep{
			{
				Config: testAccMonitoringGroup_monitoringGroupBasicExample(context),
			},
					{
				ResourceName:      "google_monitoring_group.basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
				},
	})
}

func testAccMonitoringGroup_monitoringGroupBasicExample(context map[string]interface{}) string {
  return Nprintf(`
resource "google_monitoring_group" "basic" {
  display_name = "tf-test MonitoringGroup%{random_suffix}"

  filter = "resource.metadata.region=\"europe-west2\""
}
`, context)
}

func TestAccMonitoringGroup_monitoringGroupSubgroupExample(t *testing.T) {
  t.Parallel()

	context := map[string]interface{} {
				"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
				Providers: testAccProviders,
								CheckDestroy: testAccCheckMonitoringGroupDestroyProducer(t),
				Steps: []resource.TestStep{
			{
				Config: testAccMonitoringGroup_monitoringGroupSubgroupExample(context),
			},
					{
				ResourceName:      "google_monitoring_group.subgroup",
				ImportState:       true,
				ImportStateVerify: true,
			},
				},
	})
}

func testAccMonitoringGroup_monitoringGroupSubgroupExample(context map[string]interface{}) string {
  return Nprintf(`
resource "google_monitoring_group" "parent" {
  display_name = "tf-test MonitoringParentGroup%{random_suffix}"
  filter       = "resource.metadata.region=\"europe-west2\""
}

resource "google_monitoring_group" "subgroup" {
  display_name = "tf-test MonitoringSubGroup%{random_suffix}"
  filter       = "resource.metadata.region=\"europe-west2\""
  parent_name  =  google_monitoring_group.parent.name
}
`, context)
}


func testAccCheckMonitoringGroupDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_monitoring_group" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

				config := googleProviderConfig(t)

		url, err := replaceVarsForTest(config, rs, "{{MonitoringBasePath}}v3/{{name}}")
		if err != nil {
			return err
		}

		billingProject := ""

		if config.BillingProject != "" {
			billingProject = config.BillingProject
		}

		_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil, isMonitoringConcurrentEditError)
		if err == nil {
				return fmt.Errorf("MonitoringGroup still exists at %s", url)
			}
				}

		return nil
	}
}
