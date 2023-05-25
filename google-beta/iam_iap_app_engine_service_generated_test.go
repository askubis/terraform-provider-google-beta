// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccIapAppEngineServiceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   RandString(t, 10),
		"role":            "roles/iap.httpsResourceAccessor",
		"project_id":      fmt.Sprintf("tf-test%s", RandString(t, 10)),
		"org_id":          acctest.GetTestOrgFromEnv(t),
		"billing_account": acctest.GetTestBillingAccountFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIapAppEngineServiceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_app_engine_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s roles/iap.httpsResourceAccessor", context["project_id"], context["project_id"], "default"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccIapAppEngineServiceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_iap_app_engine_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s roles/iap.httpsResourceAccessor", context["project_id"], context["project_id"], "default"),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapAppEngineServiceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   RandString(t, 10),
		"role":            "roles/iap.httpsResourceAccessor",
		"project_id":      fmt.Sprintf("tf-test%s", RandString(t, 10)),
		"org_id":          acctest.GetTestOrgFromEnv(t),
		"billing_account": acctest.GetTestBillingAccountFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccIapAppEngineServiceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_app_engine_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", context["project_id"], context["project_id"], "default"),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapAppEngineServiceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   RandString(t, 10),
		"role":            "roles/iap.httpsResourceAccessor",
		"project_id":      fmt.Sprintf("tf-test%s", RandString(t, 10)),
		"org_id":          acctest.GetTestOrgFromEnv(t),
		"billing_account": acctest.GetTestBillingAccountFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIapAppEngineServiceIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_iap_app_engine_service_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_iap_app_engine_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s", context["project_id"], context["project_id"], "default"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIapAppEngineServiceIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_iap_app_engine_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s", context["project_id"], context["project_id"], "default"),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapAppEngineServiceIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   RandString(t, 10),
		"role":            "roles/iap.httpsResourceAccessor",
		"project_id":      fmt.Sprintf("tf-test%s", RandString(t, 10)),
		"org_id":          acctest.GetTestOrgFromEnv(t),
		"billing_account": acctest.GetTestBillingAccountFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIapAppEngineServiceIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_app_engine_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s roles/iap.httpsResourceAccessor %s", context["project_id"], context["project_id"], "default", context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapAppEngineServiceIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   RandString(t, 10),
		"role":            "roles/iap.httpsResourceAccessor",
		"project_id":      fmt.Sprintf("tf-test%s", RandString(t, 10)),
		"org_id":          acctest.GetTestOrgFromEnv(t),
		"billing_account": acctest.GetTestBillingAccountFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIapAppEngineServiceIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_app_engine_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s roles/iap.httpsResourceAccessor", context["project_id"], context["project_id"], "default"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_app_engine_service_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s roles/iap.httpsResourceAccessor %s", context["project_id"], context["project_id"], "default", context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_app_engine_service_iam_binding.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s roles/iap.httpsResourceAccessor %s", context["project_id"], context["project_id"], "default", context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapAppEngineServiceIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   RandString(t, 10),
		"role":            "roles/iap.httpsResourceAccessor",
		"project_id":      fmt.Sprintf("tf-test%s", RandString(t, 10)),
		"org_id":          acctest.GetTestOrgFromEnv(t),
		"billing_account": acctest.GetTestBillingAccountFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIapAppEngineServiceIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_app_engine_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", context["project_id"], context["project_id"], "default", context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapAppEngineServiceIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   RandString(t, 10),
		"role":            "roles/iap.httpsResourceAccessor",
		"project_id":      fmt.Sprintf("tf-test%s", RandString(t, 10)),
		"org_id":          acctest.GetTestOrgFromEnv(t),
		"billing_account": acctest.GetTestBillingAccountFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIapAppEngineServiceIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_iap_app_engine_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", context["project_id"], context["project_id"], "default"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_app_engine_service_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", context["project_id"], context["project_id"], "default", context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_iap_app_engine_service_iam_member.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com %s", context["project_id"], context["project_id"], "default", context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapAppEngineServiceIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   RandString(t, 10),
		"role":            "roles/iap.httpsResourceAccessor",
		"project_id":      fmt.Sprintf("tf-test%s", RandString(t, 10)),
		"org_id":          acctest.GetTestOrgFromEnv(t),
		"billing_account": acctest.GetTestBillingAccountFromEnv(t),

		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	// Test should have 2 bindings: one with a description and one without. Any < chars are converted to a unicode character by the API.
	expectedPolicyData := Nprintf(`{"bindings":[{"condition":{"description":"%{condition_desc}","expression":"%{condition_expr}","title":"%{condition_title}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"},{"condition":{"expression":"%{condition_expr}","title":"%{condition_title}-no-description"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"}]}`, context)
	expectedPolicyData = strings.Replace(expectedPolicyData, "<", "\\u003c", -1)

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIapAppEngineServiceIamPolicy_withConditionGenerated(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO(SarahFrench) - uncomment once https://github.com/GoogleCloudPlatform/magic-modules/pull/6466 merged
					// resource.TestCheckResourceAttr("data.google_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttr("google_iap_app_engine_service_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", checkGoogleIamPolicy),
				),
			},
			{
				ResourceName:      "google_iap_app_engine_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s", context["project_id"], context["project_id"], "default"),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIapAppEngineServiceIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "my_project" {
  name            = "%{project_id}"
  project_id      = "%{project_id}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "project_service" {
  project = google_project.my_project.project_id
  service = "iap.googleapis.com"
}

resource "google_project_service" "cloudbuild_service" {
  project = google_project_service.project_service.project
  service = "cloudbuild.googleapis.com"
}

resource "google_app_engine_application" "app" {
  project     = google_project_service.cloudbuild_service.project
  location_id = "us-central"
}

resource "google_storage_bucket" "bucket" {
  project  = google_app_engine_application.app.project
  name     = "appengine-static-content-%{random_suffix}"
  location = "US"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "version" {
  project         = google_app_engine_application.app.project
  version_id      = "v2"
  service         = "default"
  runtime         = "nodejs10"
  noop_on_destroy = true
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.output_name}"
    }
  }
  env_variables = {
    port = "8080"
  }
}

resource "google_iap_app_engine_service_iam_member" "foo" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccIapAppEngineServiceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "my_project" {
  name            = "%{project_id}"
  project_id      = "%{project_id}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "project_service" {
  project = google_project.my_project.project_id
  service = "iap.googleapis.com"
}

resource "google_project_service" "cloudbuild_service" {
  project = google_project_service.project_service.project
  service = "cloudbuild.googleapis.com"
}

resource "google_app_engine_application" "app" {
  project     = google_project_service.cloudbuild_service.project
  location_id = "us-central"
}

resource "google_storage_bucket" "bucket" {
  project  = google_app_engine_application.app.project
  name     = "appengine-static-content-%{random_suffix}"
  location = "US"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "version" {
  project         = google_app_engine_application.app.project
  version_id      = "v2"
  service         = "default"
  runtime         = "nodejs10"
  noop_on_destroy = true
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.output_name}"
    }
  }
  env_variables = {
    port = "8080"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_iap_app_engine_service_iam_policy" "foo" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_iap_app_engine_service_iam_policy" "foo" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  depends_on = [
    google_iap_app_engine_service_iam_policy.foo
  ]
}
`, context)
}

func testAccIapAppEngineServiceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "my_project" {
  name            = "%{project_id}"
  project_id      = "%{project_id}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "project_service" {
  project = google_project.my_project.project_id
  service = "iap.googleapis.com"
}

resource "google_project_service" "cloudbuild_service" {
  project = google_project_service.project_service.project
  service = "cloudbuild.googleapis.com"
}

resource "google_app_engine_application" "app" {
  project     = google_project_service.cloudbuild_service.project
  location_id = "us-central"
}

resource "google_storage_bucket" "bucket" {
  project  = google_app_engine_application.app.project
  name     = "appengine-static-content-%{random_suffix}"
  location = "US"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "version" {
  project         = google_app_engine_application.app.project
  version_id      = "v2"
  service         = "default"
  runtime         = "nodejs10"
  noop_on_destroy = true
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.output_name}"
    }
  }
  env_variables = {
    port = "8080"
  }
}

data "google_iam_policy" "foo" {
}

resource "google_iap_app_engine_service_iam_policy" "foo" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccIapAppEngineServiceIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "my_project" {
  name            = "%{project_id}"
  project_id      = "%{project_id}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "project_service" {
  project = google_project.my_project.project_id
  service = "iap.googleapis.com"
}

resource "google_project_service" "cloudbuild_service" {
  project = google_project_service.project_service.project
  service = "cloudbuild.googleapis.com"
}

resource "google_app_engine_application" "app" {
  project     = google_project_service.cloudbuild_service.project
  location_id = "us-central"
}

resource "google_storage_bucket" "bucket" {
  project  = google_app_engine_application.app.project
  name     = "appengine-static-content-%{random_suffix}"
  location = "US"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "version" {
  project         = google_app_engine_application.app.project
  version_id      = "v2"
  service         = "default"
  runtime         = "nodejs10"
  noop_on_destroy = true
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.output_name}"
    }
  }
  env_variables = {
    port = "8080"
  }
}

resource "google_iap_app_engine_service_iam_binding" "foo" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccIapAppEngineServiceIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "my_project" {
  name            = "%{project_id}"
  project_id      = "%{project_id}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "project_service" {
  project = google_project.my_project.project_id
  service = "iap.googleapis.com"
}

resource "google_project_service" "cloudbuild_service" {
  project = google_project_service.project_service.project
  service = "cloudbuild.googleapis.com"
}

resource "google_app_engine_application" "app" {
  project     = google_project_service.cloudbuild_service.project
  location_id = "us-central"
}

resource "google_storage_bucket" "bucket" {
  project  = google_app_engine_application.app.project
  name     = "appengine-static-content-%{random_suffix}"
  location = "US"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "version" {
  project         = google_app_engine_application.app.project
  version_id      = "v2"
  service         = "default"
  runtime         = "nodejs10"
  noop_on_destroy = true
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.output_name}"
    }
  }
  env_variables = {
    port = "8080"
  }
}

resource "google_iap_app_engine_service_iam_binding" "foo" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccIapAppEngineServiceIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "my_project" {
  name            = "%{project_id}"
  project_id      = "%{project_id}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "project_service" {
  project = google_project.my_project.project_id
  service = "iap.googleapis.com"
}

resource "google_project_service" "cloudbuild_service" {
  project = google_project_service.project_service.project
  service = "cloudbuild.googleapis.com"
}

resource "google_app_engine_application" "app" {
  project     = google_project_service.cloudbuild_service.project
  location_id = "us-central"
}

resource "google_storage_bucket" "bucket" {
  project  = google_app_engine_application.app.project
  name     = "appengine-static-content-%{random_suffix}"
  location = "US"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "version" {
  project         = google_app_engine_application.app.project
  version_id      = "v2"
  service         = "default"
  runtime         = "nodejs10"
  noop_on_destroy = true
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.output_name}"
    }
  }
  env_variables = {
    port = "8080"
  }
}

resource "google_iap_app_engine_service_iam_binding" "foo" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccIapAppEngineServiceIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "my_project" {
  name            = "%{project_id}"
  project_id      = "%{project_id}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "project_service" {
  project = google_project.my_project.project_id
  service = "iap.googleapis.com"
}

resource "google_project_service" "cloudbuild_service" {
  project = google_project_service.project_service.project
  service = "cloudbuild.googleapis.com"
}

resource "google_app_engine_application" "app" {
  project     = google_project_service.cloudbuild_service.project
  location_id = "us-central"
}

resource "google_storage_bucket" "bucket" {
  project  = google_app_engine_application.app.project
  name     = "appengine-static-content-%{random_suffix}"
  location = "US"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "version" {
  project         = google_app_engine_application.app.project
  version_id      = "v2"
  service         = "default"
  runtime         = "nodejs10"
  noop_on_destroy = true
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.output_name}"
    }
  }
  env_variables = {
    port = "8080"
  }
}

resource "google_iap_app_engine_service_iam_binding" "foo" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_iap_app_engine_service_iam_binding" "foo2" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_iap_app_engine_service_iam_binding" "foo3" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccIapAppEngineServiceIamMember_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "my_project" {
  name            = "%{project_id}"
  project_id      = "%{project_id}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "project_service" {
  project = google_project.my_project.project_id
  service = "iap.googleapis.com"
}

resource "google_project_service" "cloudbuild_service" {
  project = google_project_service.project_service.project
  service = "cloudbuild.googleapis.com"
}

resource "google_app_engine_application" "app" {
  project     = google_project_service.cloudbuild_service.project
  location_id = "us-central"
}

resource "google_storage_bucket" "bucket" {
  project  = google_app_engine_application.app.project
  name     = "appengine-static-content-%{random_suffix}"
  location = "US"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "version" {
  project         = google_app_engine_application.app.project
  version_id      = "v2"
  service         = "default"
  runtime         = "nodejs10"
  noop_on_destroy = true
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.output_name}"
    }
  }
  env_variables = {
    port = "8080"
  }
}

resource "google_iap_app_engine_service_iam_member" "foo" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccIapAppEngineServiceIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "my_project" {
  name            = "%{project_id}"
  project_id      = "%{project_id}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "project_service" {
  project = google_project.my_project.project_id
  service = "iap.googleapis.com"
}

resource "google_project_service" "cloudbuild_service" {
  project = google_project_service.project_service.project
  service = "cloudbuild.googleapis.com"
}

resource "google_app_engine_application" "app" {
  project     = google_project_service.cloudbuild_service.project
  location_id = "us-central"
}

resource "google_storage_bucket" "bucket" {
  project  = google_app_engine_application.app.project
  name     = "appengine-static-content-%{random_suffix}"
  location = "US"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "version" {
  project         = google_app_engine_application.app.project
  version_id      = "v2"
  service         = "default"
  runtime         = "nodejs10"
  noop_on_destroy = true
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.output_name}"
    }
  }
  env_variables = {
    port = "8080"
  }
}

resource "google_iap_app_engine_service_iam_member" "foo" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_iap_app_engine_service_iam_member" "foo2" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_iap_app_engine_service_iam_member" "foo3" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccIapAppEngineServiceIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "my_project" {
  name            = "%{project_id}"
  project_id      = "%{project_id}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "project_service" {
  project = google_project.my_project.project_id
  service = "iap.googleapis.com"
}

resource "google_project_service" "cloudbuild_service" {
  project = google_project_service.project_service.project
  service = "cloudbuild.googleapis.com"
}

resource "google_app_engine_application" "app" {
  project     = google_project_service.cloudbuild_service.project
  location_id = "us-central"
}

resource "google_storage_bucket" "bucket" {
  project  = google_app_engine_application.app.project
  name     = "appengine-static-content-%{random_suffix}"
  location = "US"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "version" {
  project         = google_app_engine_application.app.project
  version_id      = "v2"
  service         = "default"
  runtime         = "nodejs10"
  noop_on_destroy = true
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.output_name}"
    }
  }
  env_variables = {
    port = "8080"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      # Check that lack of description doesn't cause any issues
      # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
      title       = "%{condition_title_no_desc}"
      expression  = "%{condition_expr_no_desc}"
    }
  }
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "%{condition_desc}"
      expression  = "%{condition_expr}"
    }
  }
}

resource "google_iap_app_engine_service_iam_policy" "foo" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
