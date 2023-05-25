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
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceCertificateManagerCertificateMapEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertificateManagerCertificateMapEntryCreate,
		Read:   resourceCertificateManagerCertificateMapEntryRead,
		Update: resourceCertificateManagerCertificateMapEntryUpdate,
		Delete: resourceCertificateManagerCertificateMapEntryDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCertificateManagerCertificateMapEntryImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"certificates": {
				Type:             schema.TypeList,
				Required:         true,
				DiffSuppressFunc: tpgresource.ProjectNumberDiffSuppress,
				Description: `A set of Certificates defines for the given hostname.
There can be defined up to fifteen certificates in each Certificate Map Entry.
Each certificate must match pattern projects/*/locations/*/certificates/*.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"map": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `A map entry that is inputted into the cetrificate map`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `A user-defined name of the Certificate Map Entry. Certificate Map Entry
names must be unique globally and match pattern
'projects/*/locations/*/certificateMaps/*/certificateMapEntries/*'`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A human-readable description of the resource.`,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
selecting a proper certificate.`,
				ExactlyOneOf: []string{"hostname", "matcher"},
			},
			"labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Optional: true,
				Description: `Set of labels associated with a Certificate Map Entry.
An object containing a list of "key": value pairs.
Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"matcher": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  `A predefined matcher for particular cases, other than SNI selection`,
				ExactlyOneOf: []string{"hostname", "matcher"},
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Creation timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `A serving state of this Certificate Map Entry.`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Update timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceCertificateManagerCertificateMapEntryCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandCertificateManagerCertificateMapEntryDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandCertificateManagerCertificateMapEntryLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	certificatesProp, err := expandCertificateManagerCertificateMapEntryCertificates(d.Get("certificates"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("certificates"); !tpgresource.IsEmptyValue(reflect.ValueOf(certificatesProp)) && (ok || !reflect.DeepEqual(v, certificatesProp)) {
		obj["certificates"] = certificatesProp
	}
	hostnameProp, err := expandCertificateManagerCertificateMapEntryHostname(d.Get("hostname"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("hostname"); !tpgresource.IsEmptyValue(reflect.ValueOf(hostnameProp)) && (ok || !reflect.DeepEqual(v, hostnameProp)) {
		obj["hostname"] = hostnameProp
	}
	matcherProp, err := expandCertificateManagerCertificateMapEntryMatcher(d.Get("matcher"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("matcher"); !tpgresource.IsEmptyValue(reflect.ValueOf(matcherProp)) && (ok || !reflect.DeepEqual(v, matcherProp)) {
		obj["matcher"] = matcherProp
	}
	nameProp, err := expandCertificateManagerCertificateMapEntryName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CertificateManagerBasePath}}projects/{{project}}/locations/global/certificateMaps/{{map}}/certificateMapEntries?certificateMapEntryId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new CertificateMapEntry: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for CertificateMapEntry: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating CertificateMapEntry: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/certificateMaps/{{map}}/certificateMapEntries/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = CertificateManagerOperationWaitTime(
		config, res, project, "Creating CertificateMapEntry", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create CertificateMapEntry: %s", err)
	}

	log.Printf("[DEBUG] Finished creating CertificateMapEntry %q: %#v", d.Id(), res)

	return resourceCertificateManagerCertificateMapEntryRead(d, meta)
}

func resourceCertificateManagerCertificateMapEntryRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CertificateManagerBasePath}}projects/{{project}}/locations/global/certificateMaps/{{map}}/certificateMapEntries/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for CertificateMapEntry: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("CertificateManagerCertificateMapEntry %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading CertificateMapEntry: %s", err)
	}

	if err := d.Set("description", flattenCertificateManagerCertificateMapEntryDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading CertificateMapEntry: %s", err)
	}
	if err := d.Set("create_time", flattenCertificateManagerCertificateMapEntryCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading CertificateMapEntry: %s", err)
	}
	if err := d.Set("update_time", flattenCertificateManagerCertificateMapEntryUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading CertificateMapEntry: %s", err)
	}
	if err := d.Set("labels", flattenCertificateManagerCertificateMapEntryLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading CertificateMapEntry: %s", err)
	}
	if err := d.Set("certificates", flattenCertificateManagerCertificateMapEntryCertificates(res["certificates"], d, config)); err != nil {
		return fmt.Errorf("Error reading CertificateMapEntry: %s", err)
	}
	if err := d.Set("state", flattenCertificateManagerCertificateMapEntryState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading CertificateMapEntry: %s", err)
	}
	if err := d.Set("hostname", flattenCertificateManagerCertificateMapEntryHostname(res["hostname"], d, config)); err != nil {
		return fmt.Errorf("Error reading CertificateMapEntry: %s", err)
	}
	if err := d.Set("matcher", flattenCertificateManagerCertificateMapEntryMatcher(res["matcher"], d, config)); err != nil {
		return fmt.Errorf("Error reading CertificateMapEntry: %s", err)
	}
	if err := d.Set("name", flattenCertificateManagerCertificateMapEntryName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading CertificateMapEntry: %s", err)
	}

	return nil
}

func resourceCertificateManagerCertificateMapEntryUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for CertificateMapEntry: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandCertificateManagerCertificateMapEntryDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandCertificateManagerCertificateMapEntryLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	certificatesProp, err := expandCertificateManagerCertificateMapEntryCertificates(d.Get("certificates"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("certificates"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, certificatesProp)) {
		obj["certificates"] = certificatesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CertificateManagerBasePath}}projects/{{project}}/locations/global/certificateMaps/{{map}}/certificateMapEntries/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating CertificateMapEntry %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("certificates") {
		updateMask = append(updateMask, "certificates")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating CertificateMapEntry %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating CertificateMapEntry %q: %#v", d.Id(), res)
	}

	err = CertificateManagerOperationWaitTime(
		config, res, project, "Updating CertificateMapEntry", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceCertificateManagerCertificateMapEntryRead(d, meta)
}

func resourceCertificateManagerCertificateMapEntryDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for CertificateMapEntry: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{CertificateManagerBasePath}}projects/{{project}}/locations/global/certificateMaps/{{map}}/certificateMapEntries/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting CertificateMapEntry %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "CertificateMapEntry")
	}

	err = CertificateManagerOperationWaitTime(
		config, res, project, "Deleting CertificateMapEntry", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting CertificateMapEntry %q: %#v", d.Id(), res)
	return nil
}

func resourceCertificateManagerCertificateMapEntryImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/global/certificateMaps/(?P<map>[^/]+)/certificateMapEntries/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<map>[^/]+)/(?P<name>[^/]+)",
		"(?P<map>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/certificateMaps/{{map}}/certificateMapEntries/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenCertificateManagerCertificateMapEntryDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateMapEntryCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateMapEntryUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateMapEntryLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateMapEntryCertificates(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateMapEntryState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateMapEntryHostname(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateMapEntryMatcher(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateMapEntryName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func expandCertificateManagerCertificateMapEntryDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateMapEntryLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCertificateManagerCertificateMapEntryCertificates(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateMapEntryHostname(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateMapEntryMatcher(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateMapEntryName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.GetResourceNameFromSelfLink(v.(string)), nil
}
