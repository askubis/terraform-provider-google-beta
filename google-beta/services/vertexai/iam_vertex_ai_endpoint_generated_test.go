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

package vertexai_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccVertexAIEndpointIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIEndpointIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_endpoint_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/endpoints/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("endpoint-name%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccVertexAIEndpointIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_endpoint_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/endpoints/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("endpoint-name%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVertexAIEndpointIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccVertexAIEndpointIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_endpoint_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/endpoints/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("endpoint-name%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVertexAIEndpointIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIEndpointIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_vertex_ai_endpoint_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_vertex_ai_endpoint_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("endpoint-name%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccVertexAIEndpointIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_vertex_ai_endpoint_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("endpoint-name%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccVertexAIEndpointIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_endpoint" "endpoint" {
  name         = "endpoint-name%{random_suffix}"
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  location     = "us-central1"
  region       = "us-central1"
  labels       = {
    label-one = "value-one"
  }
  private_service_connect_config {
    enable_private_service_connect = true
    project_allowlist = [
      "${data.google_project.project.project_id}"
    ]
    enable_secure_private_service_connect = false
  }
}

data "google_project" "project" {}

resource "google_vertex_ai_endpoint_iam_member" "foo" {
  project = google_vertex_ai_endpoint.endpoint.project
  location = google_vertex_ai_endpoint.endpoint.location
  endpoint = google_vertex_ai_endpoint.endpoint.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccVertexAIEndpointIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_endpoint" "endpoint" {
  name         = "endpoint-name%{random_suffix}"
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  location     = "us-central1"
  region       = "us-central1"
  labels       = {
    label-one = "value-one"
  }
  private_service_connect_config {
    enable_private_service_connect = true
    project_allowlist = [
      "${data.google_project.project.project_id}"
    ]
    enable_secure_private_service_connect = false
  }
}

data "google_project" "project" {}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_vertex_ai_endpoint_iam_policy" "foo" {
  project = google_vertex_ai_endpoint.endpoint.project
  location = google_vertex_ai_endpoint.endpoint.location
  endpoint = google_vertex_ai_endpoint.endpoint.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_vertex_ai_endpoint_iam_policy" "foo" {
  project = google_vertex_ai_endpoint.endpoint.project
  location = google_vertex_ai_endpoint.endpoint.location
  endpoint = google_vertex_ai_endpoint.endpoint.name
  depends_on = [
    google_vertex_ai_endpoint_iam_policy.foo
  ]
}
`, context)
}

func testAccVertexAIEndpointIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_endpoint" "endpoint" {
  name         = "endpoint-name%{random_suffix}"
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  location     = "us-central1"
  region       = "us-central1"
  labels       = {
    label-one = "value-one"
  }
  private_service_connect_config {
    enable_private_service_connect = true
    project_allowlist = [
      "${data.google_project.project.project_id}"
    ]
    enable_secure_private_service_connect = false
  }
}

data "google_project" "project" {}

data "google_iam_policy" "foo" {
}

resource "google_vertex_ai_endpoint_iam_policy" "foo" {
  project = google_vertex_ai_endpoint.endpoint.project
  location = google_vertex_ai_endpoint.endpoint.location
  endpoint = google_vertex_ai_endpoint.endpoint.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccVertexAIEndpointIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_endpoint" "endpoint" {
  name         = "endpoint-name%{random_suffix}"
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  location     = "us-central1"
  region       = "us-central1"
  labels       = {
    label-one = "value-one"
  }
  private_service_connect_config {
    enable_private_service_connect = true
    project_allowlist = [
      "${data.google_project.project.project_id}"
    ]
    enable_secure_private_service_connect = false
  }
}

data "google_project" "project" {}

resource "google_vertex_ai_endpoint_iam_binding" "foo" {
  project = google_vertex_ai_endpoint.endpoint.project
  location = google_vertex_ai_endpoint.endpoint.location
  endpoint = google_vertex_ai_endpoint.endpoint.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccVertexAIEndpointIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_endpoint" "endpoint" {
  name         = "endpoint-name%{random_suffix}"
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  location     = "us-central1"
  region       = "us-central1"
  labels       = {
    label-one = "value-one"
  }
  private_service_connect_config {
    enable_private_service_connect = true
    project_allowlist = [
      "${data.google_project.project.project_id}"
    ]
    enable_secure_private_service_connect = false
  }
}

data "google_project" "project" {}

resource "google_vertex_ai_endpoint_iam_binding" "foo" {
  project = google_vertex_ai_endpoint.endpoint.project
  location = google_vertex_ai_endpoint.endpoint.location
  endpoint = google_vertex_ai_endpoint.endpoint.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}