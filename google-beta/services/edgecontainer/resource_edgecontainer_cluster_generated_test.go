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

package edgecontainer_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccEdgecontainerCluster_edgecontainerClusterExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckEdgecontainerClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccEdgecontainerCluster_edgecontainerClusterExample(context),
			},
			{
				ResourceName:            "google_edgecontainer_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "terraform_labels"},
			},
		},
	})
}

func testAccEdgecontainerCluster_edgecontainerClusterExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_edgecontainer_cluster" "default" {
  name = "tf-test-basic-cluster%{random_suffix}"
  location = "us-central1"

  authorization {
    admin_users {
      username = "admin@hashicorptest.com"
    }
  }

  networking {
    cluster_ipv4_cidr_blocks = ["10.0.0.0/16"]
    services_ipv4_cidr_blocks = ["10.1.0.0/16"]
  }

  fleet {
    project = "projects/${data.google_project.project.number}"
  }

  labels = {
    my_key    = "my_val"
    other_key = "other_val"
  }
}

data "google_project" "project" {}
`, context)
}

func TestAccEdgecontainerCluster_edgecontainerClusterWithMaintenanceWindowExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckEdgecontainerClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccEdgecontainerCluster_edgecontainerClusterWithMaintenanceWindowExample(context),
			},
			{
				ResourceName:            "google_edgecontainer_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "terraform_labels"},
			},
		},
	})
}

func testAccEdgecontainerCluster_edgecontainerClusterWithMaintenanceWindowExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_edgecontainer_cluster" "default" {
  name = "tf-test-cluster-with-maintenance%{random_suffix}"
  location = "us-central1"

  authorization {
    admin_users {
      username = "admin@hashicorptest.com"
    }
  }

  networking {
    cluster_ipv4_cidr_blocks = ["10.0.0.0/16"]
    services_ipv4_cidr_blocks = ["10.1.0.0/16"]
  }

  fleet {
    project = "projects/${data.google_project.project.number}"
  }

  maintenance_policy {
    window {
      recurring_window {
        window {
          start_time = "2023-01-01T08:00:00Z"
          end_time = "2023-01-01T17:00:00Z"
        }

        recurrence = "FREQ=WEEKLY;BYDAY=SA"
      }
    }
  }
}

data "google_project" "project" {}
`, context)
}

func testAccCheckEdgecontainerClusterDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_edgecontainer_cluster" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{EdgecontainerBasePath}}projects/{{project}}/locations/{{location}}/clusters/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("EdgecontainerCluster still exists at %s", url)
			}
		}

		return nil
	}
}
