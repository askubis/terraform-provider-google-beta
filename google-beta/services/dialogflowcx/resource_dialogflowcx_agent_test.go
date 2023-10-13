// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package dialogflowcx_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDialogflowCXAgent_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDialogflowCXAgent_basic(context),
			},
			{
				ResourceName:      "google_dialogflow_cx_agent.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDialogflowCXAgent_full(context),
			},
			{
				ResourceName:      "google_dialogflow_cx_agent.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDialogflowCXAgent_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
	resource "google_dialogflow_cx_agent" "foobar" {
		display_name = "tf-test-%{random_suffix}"
		location = "global"
		default_language_code = "en"
		supported_language_codes = ["fr","de","es"]
		time_zone = "America/New_York"
		description = "Description 1."
		avatar_uri = "https://storage.cloud.google.com/dialogflow-test-host-image/cloud-logo.png"
	}
	`, context)
}

func testAccDialogflowCXAgent_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
	resource "google_dialogflow_cx_agent" "foobar" {
		display_name = "tf-test-%{random_suffix}update"
		location = "global"
		default_language_code = "en"
		supported_language_codes = ["no"]
		time_zone = "Europe/London"
		description = "Description 2!"
		avatar_uri = "https://storage.cloud.google.com/dialogflow-test-host-image/cloud-logo-2.png"
		enable_stackdriver_logging = true
        enable_spell_correction    = true
		speech_to_text_settings {
			enable_speech_adaptation = true
		}
	}
	  `, context)
}
