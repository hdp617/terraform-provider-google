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

																															func TestAccPrivatecaCaPoolIamBindingGenerated(t *testing.T) {
	t.Parallel()

context := map[string]interface{}{
	"random_suffix": randString(t, 10),
	"role":          "roles/privateca.certificateManager",
}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
				Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCaPoolIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s roles/privateca.certificateManager", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccPrivatecaCaPoolIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s roles/privateca.certificateManager", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPrivatecaCaPoolIamMemberGenerated(t *testing.T) {
	t.Parallel()

context := map[string]interface{}{
	"random_suffix": randString(t, 10),
	"role":          "roles/privateca.certificateManager",
}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
				Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccPrivatecaCaPoolIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s roles/privateca.certificateManager user:admin@hashicorptest.com", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPrivatecaCaPoolIamPolicyGenerated(t *testing.T) {
	t.Parallel()

context := map[string]interface{}{
	"random_suffix": randString(t, 10),
	"role":          "roles/privateca.certificateManager",
}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
				Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCaPoolIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccPrivatecaCaPoolIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_privateca_ca_pool_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/caPools/%s", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-my-pool%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}


func testAccPrivatecaCaPoolIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
}

resource "google_privateca_ca_pool_iam_member" "foo" {
  ca_pool = google_privateca_ca_pool.default.id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccPrivatecaCaPoolIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_privateca_ca_pool_iam_policy" "foo" {
  ca_pool = google_privateca_ca_pool.default.id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccPrivatecaCaPoolIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
}

data "google_iam_policy" "foo" {
}

resource "google_privateca_ca_pool_iam_policy" "foo" {
  ca_pool = google_privateca_ca_pool.default.id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccPrivatecaCaPoolIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
}

resource "google_privateca_ca_pool_iam_binding" "foo" {
  ca_pool = google_privateca_ca_pool.default.id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccPrivatecaCaPoolIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_ca_pool" "default" {
  name = "tf-test-my-pool%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
}

resource "google_privateca_ca_pool_iam_binding" "foo" {
  ca_pool = google_privateca_ca_pool.default.id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
