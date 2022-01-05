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

func TestAccNotebooksEnvironment_notebookEnvironmentBasicExample(t *testing.T) {
  t.Parallel()

	context := map[string]interface{} {
				"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
				Providers: testAccProviders,
								CheckDestroy: testAccCheckNotebooksEnvironmentDestroyProducer(t),
				Steps: []resource.TestStep{
			{
				Config: testAccNotebooksEnvironment_notebookEnvironmentBasicExample(context),
			},
					{
				ResourceName:      "google_notebooks_environment.environment",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
				},
	})
}

func testAccNotebooksEnvironment_notebookEnvironmentBasicExample(context map[string]interface{}) string {
  return Nprintf(`
resource "google_notebooks_environment" "environment" {
  name = "tf-test-notebooks-environment%{random_suffix}"
  location = "us-west1-a"  
  container_image {
    repository = "gcr.io/deeplearning-platform-release/base-cpu"
  } 
}
`, context)
}


func testAccCheckNotebooksEnvironmentDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_notebooks_environment" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

				config := googleProviderConfig(t)

		url, err := replaceVarsForTest(config, rs, "{{NotebooksBasePath}}projects/{{project}}/locations/{{location}}/environments/{{name}}")
		if err != nil {
			return err
		}

		billingProject := ""

		if config.BillingProject != "" {
			billingProject = config.BillingProject
		}

		_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
		if err == nil {
				return fmt.Errorf("NotebooksEnvironment still exists at %s", url)
			}
				}

		return nil
	}
}
