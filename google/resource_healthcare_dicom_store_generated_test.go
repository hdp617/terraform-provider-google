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

func TestAccHealthcareDicomStore_healthcareDicomStoreBasicExample(t *testing.T) {
  t.Parallel()

	context := map[string]interface{} {
				"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
				Providers: testAccProviders,
								CheckDestroy: testAccCheckHealthcareDicomStoreDestroyProducer(t),
				Steps: []resource.TestStep{
			{
				Config: testAccHealthcareDicomStore_healthcareDicomStoreBasicExample(context),
			},
					{
				ResourceName:      "google_healthcare_dicom_store.default",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"self_link", "dataset"},
			},
				},
	})
}

func testAccHealthcareDicomStore_healthcareDicomStoreBasicExample(context map[string]interface{}) string {
  return Nprintf(`
resource "google_healthcare_dicom_store" "default" {
  name    = "tf-test-example-dicom-store%{random_suffix}"
  dataset = google_healthcare_dataset.dataset.id

  notification_config {
    pubsub_topic = google_pubsub_topic.topic.id
  }

  labels = {
    label1 = "labelvalue1"
  }
}

resource "google_pubsub_topic" "topic" {
  name     = "tf-test-dicom-notifications%{random_suffix}"
}

resource "google_healthcare_dataset" "dataset" {
  name     = "tf-test-example-dataset%{random_suffix}"
  location = "us-central1"
}
`, context)
}


func testAccCheckHealthcareDicomStoreDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_healthcare_dicom_store" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

				config := googleProviderConfig(t)

		url, err := replaceVarsForTest(config, rs, "{{HealthcareBasePath}}{{dataset}}/dicomStores/{{name}}")
		if err != nil {
			return err
		}

		billingProject := ""

		if config.BillingProject != "" {
			billingProject = config.BillingProject
		}

		_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
		if err == nil {
				return fmt.Errorf("HealthcareDicomStore still exists at %s", url)
			}
				}

		return nil
	}
}
