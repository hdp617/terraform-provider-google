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

func TestAccFirestoreDocument_firestoreDocumentBasicExample(t *testing.T) {
  t.Parallel()

	context := map[string]interface{} {
    			"project_id": getTestFirestoreProjectFromEnv(t),
				"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
				Providers: testAccProviders,
								CheckDestroy: testAccCheckFirestoreDocumentDestroyProducer(t),
				Steps: []resource.TestStep{
			{
				Config: testAccFirestoreDocument_firestoreDocumentBasicExample(context),
			},
					{
				ResourceName:      "google_firestore_document.mydoc",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"database", "collection", "document_id"},
			},
				},
	})
}

func testAccFirestoreDocument_firestoreDocumentBasicExample(context map[string]interface{}) string {
  return Nprintf(`
resource "google_firestore_document" "mydoc" {
  project     = "%{project_id}"
  collection  = "somenewcollection"
  document_id = "my-doc-%{random_suffix}"
  fields      = "{\"something\":{\"mapValue\":{\"fields\":{\"akey\":{\"stringValue\":\"avalue\"}}}}}"
}
`, context)
}

func TestAccFirestoreDocument_firestoreDocumentNestedDocumentExample(t *testing.T) {
  t.Parallel()

	context := map[string]interface{} {
    			"project_id": getTestFirestoreProjectFromEnv(t),
				"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
				Providers: testAccProviders,
								CheckDestroy: testAccCheckFirestoreDocumentDestroyProducer(t),
				Steps: []resource.TestStep{
			{
				Config: testAccFirestoreDocument_firestoreDocumentNestedDocumentExample(context),
			},
					{
				ResourceName:      "google_firestore_document.mydoc",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"database", "collection", "document_id"},
			},
				},
	})
}

func testAccFirestoreDocument_firestoreDocumentNestedDocumentExample(context map[string]interface{}) string {
  return Nprintf(`
resource "google_firestore_document" "mydoc" {
  project     = "%{project_id}"
  collection  = "somenewcollection"
  document_id = "my-doc-%{random_suffix}"
  fields      = "{\"something\":{\"mapValue\":{\"fields\":{\"akey\":{\"stringValue\":\"avalue\"}}}}}"
}

resource "google_firestore_document" "sub_document" {
  project     = "%{project_id}"
  collection  = "${google_firestore_document.mydoc.path}/subdocs"
  document_id = "bitcoinkey"
  fields      = "{\"something\":{\"mapValue\":{\"fields\":{\"ayo\":{\"stringValue\":\"val2\"}}}}}"
}

resource "google_firestore_document" "sub_sub_document" {
  project     = "%{project_id}"
  collection  = "${google_firestore_document.sub_document.path}/subsubdocs"
  document_id = "asecret"
  fields      = "{\"something\":{\"mapValue\":{\"fields\":{\"secret\":{\"stringValue\":\"hithere\"}}}}}"
}
`, context)
}


func testAccCheckFirestoreDocumentDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_firestore_document" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

				config := googleProviderConfig(t)

		url, err := replaceVarsForTest(config, rs, "{{FirestoreBasePath}}{{name}}")
		if err != nil {
			return err
		}

		billingProject := ""

		if config.BillingProject != "" {
			billingProject = config.BillingProject
		}

		_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
		if err == nil {
				return fmt.Errorf("FirestoreDocument still exists at %s", url)
			}
				}

		return nil
	}
}
