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

func TestAccNotebooksInstance_notebookInstanceBasicExample(t *testing.T) {
  t.Parallel()

	context := map[string]interface{} {
				"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
				Providers: testAccProviders,
								CheckDestroy: testAccCheckNotebooksInstanceDestroyProducer(t),
				Steps: []resource.TestStep{
			{
				Config: testAccNotebooksInstance_notebookInstanceBasicExample(context),
			},
					{
				ResourceName:      "google_notebooks_instance.instance",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"name", "instance_owners", "boot_disk_type", "boot_disk_size_gb", "data_disk_size_gb", "no_remove_data_disk", "metadata", "vm_image", "container_image", "location"},
			},
				},
	})
}

func testAccNotebooksInstance_notebookInstanceBasicExample(context map[string]interface{}) string {
  return Nprintf(`
resource "google_notebooks_instance" "instance" {
  name = "tf-test-notebooks-instance%{random_suffix}"
  location = "us-west1-a"
  machine_type = "e2-medium"
  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-cpu"
  }
}
`, context)
}

func TestAccNotebooksInstance_notebookInstanceBasicContainerExample(t *testing.T) {
  t.Parallel()

	context := map[string]interface{} {
				"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
				Providers: testAccProviders,
								CheckDestroy: testAccCheckNotebooksInstanceDestroyProducer(t),
				Steps: []resource.TestStep{
			{
				Config: testAccNotebooksInstance_notebookInstanceBasicContainerExample(context),
			},
					{
				ResourceName:      "google_notebooks_instance.instance",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"name", "instance_owners", "boot_disk_type", "boot_disk_size_gb", "data_disk_size_gb", "no_remove_data_disk", "metadata", "vm_image", "container_image", "location"},
			},
				},
	})
}

func testAccNotebooksInstance_notebookInstanceBasicContainerExample(context map[string]interface{}) string {
  return Nprintf(`
resource "google_notebooks_instance" "instance" {
  name = "tf-test-notebooks-instance%{random_suffix}"
  location = "us-west1-a"
  machine_type = "e2-medium"
  metadata = {
    proxy-mode = "service_account"
    terraform  = "true"
  }
  container_image {
    repository = "gcr.io/deeplearning-platform-release/base-cpu"
    tag = "latest"
  }
}
`, context)
}

func TestAccNotebooksInstance_notebookInstanceBasicGpuExample(t *testing.T) {
  t.Parallel()

	context := map[string]interface{} {
				"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
				Providers: testAccProviders,
								CheckDestroy: testAccCheckNotebooksInstanceDestroyProducer(t),
				Steps: []resource.TestStep{
			{
				Config: testAccNotebooksInstance_notebookInstanceBasicGpuExample(context),
			},
					{
				ResourceName:      "google_notebooks_instance.instance",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"name", "instance_owners", "boot_disk_type", "boot_disk_size_gb", "data_disk_size_gb", "no_remove_data_disk", "metadata", "vm_image", "container_image", "location"},
			},
				},
	})
}

func testAccNotebooksInstance_notebookInstanceBasicGpuExample(context map[string]interface{}) string {
  return Nprintf(`
resource "google_notebooks_instance" "instance" {
  name = "tf-test-notebooks-instance%{random_suffix}"
  location = "us-west1-a"
  machine_type = "n1-standard-1" // can't be e2 because of accelerator

  install_gpu_driver = true
  accelerator_config {
    type         = "NVIDIA_TESLA_T4"
    core_count   = 1
  }
  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-gpu"
  }
}
`, context)
}

func TestAccNotebooksInstance_notebookInstanceFullExample(t *testing.T) {
  t.Parallel()

	context := map[string]interface{} {
    			"service_account": getTestServiceAccountFromEnv(t),
				"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
				Providers: testAccProviders,
								CheckDestroy: testAccCheckNotebooksInstanceDestroyProducer(t),
				Steps: []resource.TestStep{
			{
				Config: testAccNotebooksInstance_notebookInstanceFullExample(context),
			},
					{
				ResourceName:      "google_notebooks_instance.instance",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"name", "instance_owners", "boot_disk_type", "boot_disk_size_gb", "data_disk_size_gb", "no_remove_data_disk", "metadata", "vm_image", "container_image", "location"},
			},
				},
	})
}

func testAccNotebooksInstance_notebookInstanceFullExample(context map[string]interface{}) string {
  return Nprintf(`
resource "google_notebooks_instance" "instance" {
  name = "tf-test-notebooks-instance%{random_suffix}"
  location = "us-central1-a"
  machine_type = "e2-medium"

  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-cpu"
  }

  instance_owners = ["admin@hashicorptest.com"]
  service_account = "%{service_account}"

  install_gpu_driver = true
  boot_disk_type = "PD_SSD"
  boot_disk_size_gb = 110

  no_public_ip = true
  no_proxy_access = true

  network = data.google_compute_network.my_network.id
  subnet = data.google_compute_subnetwork.my_subnetwork.id

  labels = {
    k = "val"
  }

  metadata = {
    terraform = "true"
  }
}

data "google_compute_network" "my_network" {
  name = "default"
}

data "google_compute_subnetwork" "my_subnetwork" {
  name   = "default"
  region = "us-central1"
}
`, context)
}


func testAccCheckNotebooksInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_notebooks_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

				config := googleProviderConfig(t)

		url, err := replaceVarsForTest(config, rs, "{{NotebooksBasePath}}projects/{{project}}/locations/{{location}}/instances/{{name}}")
		if err != nil {
			return err
		}

		billingProject := ""

		if config.BillingProject != "" {
			billingProject = config.BillingProject
		}

		_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
		if err == nil {
				return fmt.Errorf("NotebooksInstance still exists at %s", url)
			}
				}

		return nil
	}
}
