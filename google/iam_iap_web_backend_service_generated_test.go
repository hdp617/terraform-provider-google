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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

																				func TestAccIapWebBackendServiceIamBindingGenerated(t *testing.T) {
	t.Parallel()

context := map[string]interface{}{
	"random_suffix": randString(t, 10),
	"role":          "roles/iap.httpsResourceAccessor",
}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
				Steps: []resource.TestStep{
			{
				Config: testAccIapWebBackendServiceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor", getTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccIapWebBackendServiceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor", getTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebBackendServiceIamMemberGenerated(t *testing.T) {
	t.Parallel()

context := map[string]interface{}{
	"random_suffix": randString(t, 10),
	"role":          "roles/iap.httpsResourceAccessor",
}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
				Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccIapWebBackendServiceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", getTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapWebBackendServiceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

context := map[string]interface{}{
	"random_suffix": randString(t, 10),
	"role":          "roles/iap.httpsResourceAccessor",
}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
				Steps: []resource.TestStep{
			{
				Config: testAccIapWebBackendServiceIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s", getTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIapWebBackendServiceIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_iap_web_backend_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/compute/services/%s", getTestProjectFromEnv(), fmt.Sprintf("tf-test-backend-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}


func testAccIapWebBackendServiceIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_member" "foo" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccIapWebBackendServiceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_iap_web_backend_service_iam_policy" "foo" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccIapWebBackendServiceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

data "google_iam_policy" "foo" {
}

resource "google_iap_web_backend_service_iam_policy" "foo" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccIapWebBackendServiceIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_binding" "foo" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccIapWebBackendServiceIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_iap_web_backend_service_iam_binding" "foo" {
  project = google_compute_backend_service.default.project
  web_backend_service = google_compute_backend_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
