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

func TestAccHealthcareDataset_healthcareDatasetBasicExample(t *testing.T) {
  t.Parallel()

	context := map[string]interface{} {
				"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
				Providers: testAccProviders,
								CheckDestroy: testAccCheckHealthcareDatasetDestroyProducer(t),
				Steps: []resource.TestStep{
			{
				Config: testAccHealthcareDataset_healthcareDatasetBasicExample(context),
			},
					{
				ResourceName:      "google_healthcare_dataset.default",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"self_link", "location"},
			},
				},
	})
}

func testAccHealthcareDataset_healthcareDatasetBasicExample(context map[string]interface{}) string {
  return Nprintf(`
resource "google_healthcare_dataset" "default" {
  name      = "tf-test-example-dataset%{random_suffix}"
  location  = "us-central1"
  time_zone = "UTC"
}
`, context)
}


func testAccCheckHealthcareDatasetDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_healthcare_dataset" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

				config := googleProviderConfig(t)

		url, err := replaceVarsForTest(config, rs, "{{HealthcareBasePath}}projects/{{project}}/locations/{{location}}/datasets/{{name}}")
		if err != nil {
			return err
		}

		billingProject := ""

		if config.BillingProject != "" {
			billingProject = config.BillingProject
		}

		_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil, healthcareDatasetNotInitialized)
		if err == nil {
				return fmt.Errorf("HealthcareDataset still exists at %s", url)
			}
				}

		return nil
	}
}
