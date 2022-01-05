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

func TestAccComputeTargetHttpProxy_targetHttpProxyBasicExample(t *testing.T) {
  t.Parallel()

	context := map[string]interface{} {
				"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
				Providers: testAccProviders,
								CheckDestroy: testAccCheckComputeTargetHttpProxyDestroyProducer(t),
				Steps: []resource.TestStep{
			{
				Config: testAccComputeTargetHttpProxy_targetHttpProxyBasicExample(context),
			},
					{
				ResourceName:      "google_compute_target_http_proxy.default",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"url_map"},
			},
				},
	})
}

func testAccComputeTargetHttpProxy_targetHttpProxyBasicExample(context map[string]interface{}) string {
  return Nprintf(`
resource "google_compute_target_http_proxy" "default" {
  name    = "tf-test-test-proxy%{random_suffix}"
  url_map = google_compute_url_map.default.id
}

resource "google_compute_url_map" "default" {
  name            = "tf-test-url-map%{random_suffix}"
  default_service = google_compute_backend_service.default.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.id
    }
  }
}

resource "google_compute_backend_service" "default" {
  name        = "tf-test-backend-service%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-http-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, context)
}

func TestAccComputeTargetHttpProxy_targetHttpProxyHttpsRedirectExample(t *testing.T) {
  t.Parallel()

	context := map[string]interface{} {
				"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
				Providers: testAccProviders,
								CheckDestroy: testAccCheckComputeTargetHttpProxyDestroyProducer(t),
				Steps: []resource.TestStep{
			{
				Config: testAccComputeTargetHttpProxy_targetHttpProxyHttpsRedirectExample(context),
			},
					{
				ResourceName:      "google_compute_target_http_proxy.default",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"url_map"},
			},
				},
	})
}

func testAccComputeTargetHttpProxy_targetHttpProxyHttpsRedirectExample(context map[string]interface{}) string {
  return Nprintf(`
resource "google_compute_target_http_proxy" "default" {
  name    = "tf-test-test-https-redirect-proxy%{random_suffix}"
  url_map = google_compute_url_map.default.id
}

resource "google_compute_url_map" "default" {
  name            = "tf-test-url-map%{random_suffix}"
  default_url_redirect {
    https_redirect = true
    strip_query    = false
  }
}
`, context)
}


func testAccCheckComputeTargetHttpProxyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_target_http_proxy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

				config := googleProviderConfig(t)

		url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/targetHttpProxies/{{name}}")
		if err != nil {
			return err
		}

		billingProject := ""

		if config.BillingProject != "" {
			billingProject = config.BillingProject
		}

		_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
		if err == nil {
				return fmt.Errorf("ComputeTargetHttpProxy still exists at %s", url)
			}
				}

		return nil
	}
}
