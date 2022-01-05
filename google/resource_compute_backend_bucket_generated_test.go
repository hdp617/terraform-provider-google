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

func TestAccComputeBackendBucket_backendBucketBasicExample(t *testing.T) {
  t.Parallel()

	context := map[string]interface{} {
				"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
				Providers: testAccProviders,
								CheckDestroy: testAccCheckComputeBackendBucketDestroyProducer(t),
				Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendBucket_backendBucketBasicExample(context),
			},
					{
				ResourceName:      "google_compute_backend_bucket.image_backend",
				ImportState:       true,
				ImportStateVerify: true,
			},
				},
	})
}

func testAccComputeBackendBucket_backendBucketBasicExample(context map[string]interface{}) string {
  return Nprintf(`
resource "google_compute_backend_bucket" "image_backend" {
  name        = "tf-test-image-backend-bucket%{random_suffix}"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.image_bucket.name
  enable_cdn  = true
}

resource "google_storage_bucket" "image_bucket" {
  name     = "tf-test-image-store-bucket%{random_suffix}"
  location = "EU"
}
`, context)
}

func TestAccComputeBackendBucket_backendBucketFullExample(t *testing.T) {
  t.Parallel()

	context := map[string]interface{} {
				"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
				Providers: testAccProviders,
								CheckDestroy: testAccCheckComputeBackendBucketDestroyProducer(t),
				Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendBucket_backendBucketFullExample(context),
			},
					{
				ResourceName:      "google_compute_backend_bucket.image_backend_full",
				ImportState:       true,
				ImportStateVerify: true,
			},
				},
	})
}

func testAccComputeBackendBucket_backendBucketFullExample(context map[string]interface{}) string {
  return Nprintf(`
resource "google_compute_backend_bucket" "image_backend_full" {
  name        = "tf-test-image-backend-bucket-full%{random_suffix}"
  description = "Contains beautiful beta mages"
  bucket_name = google_storage_bucket.image_backend_full.name
  enable_cdn  = true
  cdn_policy {
    cache_mode = "CACHE_ALL_STATIC"
    default_ttl = 3600
    client_ttl  = 7200
    max_ttl     = 10800
    negative_caching = true
  }
  custom_response_headers = [
    "X-Client-Geo-Location:{client_region},{client_city}",
    "X-Tested-By:Magic-Modules"
  ]
}

resource "google_storage_bucket" "image_backend_full" {
  name     = "tf-test-image-store-bucket-full%{random_suffix}"
  location = "EU"
}
`, context)
}


func testAccCheckComputeBackendBucketDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_backend_bucket" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

				config := googleProviderConfig(t)

		url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/backendBuckets/{{name}}")
		if err != nil {
			return err
		}

		billingProject := ""

		if config.BillingProject != "" {
			billingProject = config.BillingProject
		}

		_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
		if err == nil {
				return fmt.Errorf("ComputeBackendBucket still exists at %s", url)
			}
				}

		return nil
	}
}
