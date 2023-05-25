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

func ResourceIdentityPlatformTenantInboundSamlConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceIdentityPlatformTenantInboundSamlConfigCreate,
		Read:   resourceIdentityPlatformTenantInboundSamlConfigRead,
		Update: resourceIdentityPlatformTenantInboundSamlConfigUpdate,
		Delete: resourceIdentityPlatformTenantInboundSamlConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIdentityPlatformTenantInboundSamlConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Human friendly display name.`,
			},
			"idp_config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `SAML IdP configuration when the project acts as the relying party`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"idp_certificates": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"x509_certificate": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The x509 certificate`,
									},
								},
							},
						},
						"idp_entity_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Unique identifier for all SAML entities`,
						},
						"sso_url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `URL to send Authentication request to.`,
						},
						"sign_request": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Indicates if outbounding SAMLRequest should be signed.`,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
alphanumeric character, and have at least 2 characters.`,
			},
			"sp_config": {
				Type:     schema.TypeList,
				Required: true,
				Description: `SAML SP (Service Provider) configuration when the project acts as the relying party to receive
and accept an authentication assertion issued by a SAML identity provider.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"callback_uri": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Callback URI where responses from IDP are handled. Must start with 'https://'.`,
						},
						"sp_entity_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Unique identifier for all SAML entities.`,
						},
						"sp_certificates": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"x509_certificate": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `The x509 certificate`,
									},
								},
							},
						},
					},
				},
			},
			"tenant": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the tenant where this inbound SAML config resource exists`,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `If this config allows users to sign in with the provider.`,
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

func resourceIdentityPlatformTenantInboundSamlConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandIdentityPlatformTenantInboundSamlConfigName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	displayNameProp, err := expandIdentityPlatformTenantInboundSamlConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enabledProp, err := expandIdentityPlatformTenantInboundSamlConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}
	idpConfigProp, err := expandIdentityPlatformTenantInboundSamlConfigIdpConfig(d.Get("idp_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("idp_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(idpConfigProp)) && (ok || !reflect.DeepEqual(v, idpConfigProp)) {
		obj["idpConfig"] = idpConfigProp
	}
	spConfigProp, err := expandIdentityPlatformTenantInboundSamlConfigSpConfig(d.Get("sp_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("sp_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(spConfigProp)) && (ok || !reflect.DeepEqual(v, spConfigProp)) {
		obj["spConfig"] = spConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{tenant}}/inboundSamlConfigs?inboundSamlConfigId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TenantInboundSamlConfig: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TenantInboundSamlConfig: %s", err)
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
		return fmt.Errorf("Error creating TenantInboundSamlConfig: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/tenants/{{tenant}}/inboundSamlConfigs/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating TenantInboundSamlConfig %q: %#v", d.Id(), res)

	return resourceIdentityPlatformTenantInboundSamlConfigRead(d, meta)
}

func resourceIdentityPlatformTenantInboundSamlConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{tenant}}/inboundSamlConfigs/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TenantInboundSamlConfig: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IdentityPlatformTenantInboundSamlConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading TenantInboundSamlConfig: %s", err)
	}

	if err := d.Set("name", flattenIdentityPlatformTenantInboundSamlConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading TenantInboundSamlConfig: %s", err)
	}
	if err := d.Set("display_name", flattenIdentityPlatformTenantInboundSamlConfigDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading TenantInboundSamlConfig: %s", err)
	}
	if err := d.Set("enabled", flattenIdentityPlatformTenantInboundSamlConfigEnabled(res["enabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading TenantInboundSamlConfig: %s", err)
	}
	if err := d.Set("idp_config", flattenIdentityPlatformTenantInboundSamlConfigIdpConfig(res["idpConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading TenantInboundSamlConfig: %s", err)
	}
	if err := d.Set("sp_config", flattenIdentityPlatformTenantInboundSamlConfigSpConfig(res["spConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading TenantInboundSamlConfig: %s", err)
	}

	return nil
}

func resourceIdentityPlatformTenantInboundSamlConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TenantInboundSamlConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandIdentityPlatformTenantInboundSamlConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enabledProp, err := expandIdentityPlatformTenantInboundSamlConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}
	idpConfigProp, err := expandIdentityPlatformTenantInboundSamlConfigIdpConfig(d.Get("idp_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("idp_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, idpConfigProp)) {
		obj["idpConfig"] = idpConfigProp
	}
	spConfigProp, err := expandIdentityPlatformTenantInboundSamlConfigSpConfig(d.Get("sp_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("sp_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, spConfigProp)) {
		obj["spConfig"] = spConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{tenant}}/inboundSamlConfigs/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating TenantInboundSamlConfig %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("enabled") {
		updateMask = append(updateMask, "enabled")
	}

	if d.HasChange("idp_config") {
		updateMask = append(updateMask, "idpConfig")
	}

	if d.HasChange("sp_config") {
		updateMask = append(updateMask, "spConfig")
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
		return fmt.Errorf("Error updating TenantInboundSamlConfig %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating TenantInboundSamlConfig %q: %#v", d.Id(), res)
	}

	return resourceIdentityPlatformTenantInboundSamlConfigRead(d, meta)
}

func resourceIdentityPlatformTenantInboundSamlConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TenantInboundSamlConfig: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{tenant}}/inboundSamlConfigs/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting TenantInboundSamlConfig %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "TenantInboundSamlConfig")
	}

	log.Printf("[DEBUG] Finished deleting TenantInboundSamlConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceIdentityPlatformTenantInboundSamlConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/tenants/(?P<tenant>[^/]+)/inboundSamlConfigs/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<tenant>[^/]+)/(?P<name>[^/]+)",
		"(?P<tenant>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/tenants/{{tenant}}/inboundSamlConfigs/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIdentityPlatformTenantInboundSamlConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenIdentityPlatformTenantInboundSamlConfigDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantInboundSamlConfigEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantInboundSamlConfigIdpConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["idp_entity_id"] =
		flattenIdentityPlatformTenantInboundSamlConfigIdpConfigIdpEntityId(original["idpEntityId"], d, config)
	transformed["sso_url"] =
		flattenIdentityPlatformTenantInboundSamlConfigIdpConfigSsoUrl(original["ssoUrl"], d, config)
	transformed["sign_request"] =
		flattenIdentityPlatformTenantInboundSamlConfigIdpConfigSignRequest(original["signRequest"], d, config)
	transformed["idp_certificates"] =
		flattenIdentityPlatformTenantInboundSamlConfigIdpConfigIdpCertificates(original["idpCertificates"], d, config)
	return []interface{}{transformed}
}
func flattenIdentityPlatformTenantInboundSamlConfigIdpConfigIdpEntityId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantInboundSamlConfigIdpConfigSsoUrl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantInboundSamlConfigIdpConfigSignRequest(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantInboundSamlConfigIdpConfigIdpCertificates(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"x509_certificate": flattenIdentityPlatformTenantInboundSamlConfigIdpConfigIdpCertificatesX509Certificate(original["x509Certificate"], d, config),
		})
	}
	return transformed
}
func flattenIdentityPlatformTenantInboundSamlConfigIdpConfigIdpCertificatesX509Certificate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantInboundSamlConfigSpConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["sp_entity_id"] =
		flattenIdentityPlatformTenantInboundSamlConfigSpConfigSpEntityId(original["spEntityId"], d, config)
	transformed["callback_uri"] =
		flattenIdentityPlatformTenantInboundSamlConfigSpConfigCallbackUri(original["callbackUri"], d, config)
	transformed["sp_certificates"] =
		flattenIdentityPlatformTenantInboundSamlConfigSpConfigSpCertificates(original["spCertificates"], d, config)
	return []interface{}{transformed}
}
func flattenIdentityPlatformTenantInboundSamlConfigSpConfigSpEntityId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantInboundSamlConfigSpConfigCallbackUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantInboundSamlConfigSpConfigSpCertificates(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"x509_certificate": flattenIdentityPlatformTenantInboundSamlConfigSpConfigSpCertificatesX509Certificate(original["x509Certificate"], d, config),
		})
	}
	return transformed
}
func flattenIdentityPlatformTenantInboundSamlConfigSpConfigSpCertificatesX509Certificate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandIdentityPlatformTenantInboundSamlConfigName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigIdpConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIdpEntityId, err := expandIdentityPlatformTenantInboundSamlConfigIdpConfigIdpEntityId(original["idp_entity_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdpEntityId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["idpEntityId"] = transformedIdpEntityId
	}

	transformedSsoUrl, err := expandIdentityPlatformTenantInboundSamlConfigIdpConfigSsoUrl(original["sso_url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSsoUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ssoUrl"] = transformedSsoUrl
	}

	transformedSignRequest, err := expandIdentityPlatformTenantInboundSamlConfigIdpConfigSignRequest(original["sign_request"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSignRequest); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["signRequest"] = transformedSignRequest
	}

	transformedIdpCertificates, err := expandIdentityPlatformTenantInboundSamlConfigIdpConfigIdpCertificates(original["idp_certificates"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdpCertificates); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["idpCertificates"] = transformedIdpCertificates
	}

	return transformed, nil
}

func expandIdentityPlatformTenantInboundSamlConfigIdpConfigIdpEntityId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigIdpConfigSsoUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigIdpConfigSignRequest(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigIdpConfigIdpCertificates(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedX509Certificate, err := expandIdentityPlatformTenantInboundSamlConfigIdpConfigIdpCertificatesX509Certificate(original["x509_certificate"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedX509Certificate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["x509Certificate"] = transformedX509Certificate
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandIdentityPlatformTenantInboundSamlConfigIdpConfigIdpCertificatesX509Certificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigSpConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSpEntityId, err := expandIdentityPlatformTenantInboundSamlConfigSpConfigSpEntityId(original["sp_entity_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSpEntityId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["spEntityId"] = transformedSpEntityId
	}

	transformedCallbackUri, err := expandIdentityPlatformTenantInboundSamlConfigSpConfigCallbackUri(original["callback_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCallbackUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["callbackUri"] = transformedCallbackUri
	}

	transformedSpCertificates, err := expandIdentityPlatformTenantInboundSamlConfigSpConfigSpCertificates(original["sp_certificates"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSpCertificates); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["spCertificates"] = transformedSpCertificates
	}

	return transformed, nil
}

func expandIdentityPlatformTenantInboundSamlConfigSpConfigSpEntityId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigSpConfigCallbackUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigSpConfigSpCertificates(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedX509Certificate, err := expandIdentityPlatformTenantInboundSamlConfigSpConfigSpCertificatesX509Certificate(original["x509_certificate"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedX509Certificate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["x509Certificate"] = transformedX509Certificate
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandIdentityPlatformTenantInboundSamlConfigSpConfigSpCertificatesX509Certificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
