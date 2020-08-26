// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
)

func resourceIdentityPlatformInboundSamlConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceIdentityPlatformInboundSamlConfigCreate,
		Read:   resourceIdentityPlatformInboundSamlConfigRead,
		Update: resourceIdentityPlatformInboundSamlConfigUpdate,
		Delete: resourceIdentityPlatformInboundSamlConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIdentityPlatformInboundSamlConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
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
							Description: `The IdP's certificate data to verify the signature in the SAMLResponse issued by the IDP.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"x509_certificate": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The IdP's x509 certificate.`,
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
							Optional:    true,
							Description: `Callback URI where responses from IDP are handled. Must start with 'https://'.`,
						},
						"sp_entity_id": {
							Type:        schema.TypeString,
							Optional:    true,
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
	}
}

func resourceIdentityPlatformInboundSamlConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandIdentityPlatformInboundSamlConfigName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	displayNameProp, err := expandIdentityPlatformInboundSamlConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enabledProp, err := expandIdentityPlatformInboundSamlConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !isEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}
	idpConfigProp, err := expandIdentityPlatformInboundSamlConfigIdpConfig(d.Get("idp_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("idp_config"); !isEmptyValue(reflect.ValueOf(idpConfigProp)) && (ok || !reflect.DeepEqual(v, idpConfigProp)) {
		obj["idpConfig"] = idpConfigProp
	}
	spConfigProp, err := expandIdentityPlatformInboundSamlConfigSpConfig(d.Get("sp_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("sp_config"); !isEmptyValue(reflect.ValueOf(spConfigProp)) && (ok || !reflect.DeepEqual(v, spConfigProp)) {
		obj["spConfig"] = spConfigProp
	}

	url, err := replaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/inboundSamlConfigs?inboundSamlConfigId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new InboundSamlConfig: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating InboundSamlConfig: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/inboundSamlConfigs/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating InboundSamlConfig %q: %#v", d.Id(), res)

	return resourceIdentityPlatformInboundSamlConfigRead(d, meta)
}

func resourceIdentityPlatformInboundSamlConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/inboundSamlConfigs/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("IdentityPlatformInboundSamlConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading InboundSamlConfig: %s", err)
	}

	if err := d.Set("name", flattenIdentityPlatformInboundSamlConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading InboundSamlConfig: %s", err)
	}
	if err := d.Set("display_name", flattenIdentityPlatformInboundSamlConfigDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading InboundSamlConfig: %s", err)
	}
	if err := d.Set("enabled", flattenIdentityPlatformInboundSamlConfigEnabled(res["enabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading InboundSamlConfig: %s", err)
	}
	if err := d.Set("idp_config", flattenIdentityPlatformInboundSamlConfigIdpConfig(res["idpConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading InboundSamlConfig: %s", err)
	}
	if err := d.Set("sp_config", flattenIdentityPlatformInboundSamlConfigSpConfig(res["spConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading InboundSamlConfig: %s", err)
	}

	return nil
}

func resourceIdentityPlatformInboundSamlConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandIdentityPlatformInboundSamlConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enabledProp, err := expandIdentityPlatformInboundSamlConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}
	idpConfigProp, err := expandIdentityPlatformInboundSamlConfigIdpConfig(d.Get("idp_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("idp_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, idpConfigProp)) {
		obj["idpConfig"] = idpConfigProp
	}
	spConfigProp, err := expandIdentityPlatformInboundSamlConfigSpConfig(d.Get("sp_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("sp_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, spConfigProp)) {
		obj["spConfig"] = spConfigProp
	}

	url, err := replaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/inboundSamlConfigs/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating InboundSamlConfig %q: %#v", d.Id(), obj)
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
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating InboundSamlConfig %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating InboundSamlConfig %q: %#v", d.Id(), res)
	}

	return resourceIdentityPlatformInboundSamlConfigRead(d, meta)
}

func resourceIdentityPlatformInboundSamlConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/inboundSamlConfigs/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting InboundSamlConfig %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "InboundSamlConfig")
	}

	log.Printf("[DEBUG] Finished deleting InboundSamlConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceIdentityPlatformInboundSamlConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/inboundSamlConfigs/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/inboundSamlConfigs/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIdentityPlatformInboundSamlConfigName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenIdentityPlatformInboundSamlConfigDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIdentityPlatformInboundSamlConfigEnabled(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIdentityPlatformInboundSamlConfigIdpConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["idp_entity_id"] =
		flattenIdentityPlatformInboundSamlConfigIdpConfigIdpEntityId(original["idpEntityId"], d, config)
	transformed["sso_url"] =
		flattenIdentityPlatformInboundSamlConfigIdpConfigSsoUrl(original["ssoUrl"], d, config)
	transformed["sign_request"] =
		flattenIdentityPlatformInboundSamlConfigIdpConfigSignRequest(original["signRequest"], d, config)
	transformed["idp_certificates"] =
		flattenIdentityPlatformInboundSamlConfigIdpConfigIdpCertificates(original["idpCertificates"], d, config)
	return []interface{}{transformed}
}
func flattenIdentityPlatformInboundSamlConfigIdpConfigIdpEntityId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIdentityPlatformInboundSamlConfigIdpConfigSsoUrl(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIdentityPlatformInboundSamlConfigIdpConfigSignRequest(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIdentityPlatformInboundSamlConfigIdpConfigIdpCertificates(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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
			"x509_certificate": flattenIdentityPlatformInboundSamlConfigIdpConfigIdpCertificatesX509Certificate(original["x509Certificate"], d, config),
		})
	}
	return transformed
}
func flattenIdentityPlatformInboundSamlConfigIdpConfigIdpCertificatesX509Certificate(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIdentityPlatformInboundSamlConfigSpConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["sp_entity_id"] =
		flattenIdentityPlatformInboundSamlConfigSpConfigSpEntityId(original["spEntityId"], d, config)
	transformed["callback_uri"] =
		flattenIdentityPlatformInboundSamlConfigSpConfigCallbackUri(original["callbackUri"], d, config)
	transformed["sp_certificates"] =
		flattenIdentityPlatformInboundSamlConfigSpConfigSpCertificates(original["spCertificates"], d, config)
	return []interface{}{transformed}
}
func flattenIdentityPlatformInboundSamlConfigSpConfigSpEntityId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIdentityPlatformInboundSamlConfigSpConfigCallbackUri(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIdentityPlatformInboundSamlConfigSpConfigSpCertificates(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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
			"x509_certificate": flattenIdentityPlatformInboundSamlConfigSpConfigSpCertificatesX509Certificate(original["x509Certificate"], d, config),
		})
	}
	return transformed
}
func flattenIdentityPlatformInboundSamlConfigSpConfigSpCertificatesX509Certificate(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandIdentityPlatformInboundSamlConfigName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformInboundSamlConfigDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformInboundSamlConfigEnabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformInboundSamlConfigIdpConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIdpEntityId, err := expandIdentityPlatformInboundSamlConfigIdpConfigIdpEntityId(original["idp_entity_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdpEntityId); val.IsValid() && !isEmptyValue(val) {
		transformed["idpEntityId"] = transformedIdpEntityId
	}

	transformedSsoUrl, err := expandIdentityPlatformInboundSamlConfigIdpConfigSsoUrl(original["sso_url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSsoUrl); val.IsValid() && !isEmptyValue(val) {
		transformed["ssoUrl"] = transformedSsoUrl
	}

	transformedSignRequest, err := expandIdentityPlatformInboundSamlConfigIdpConfigSignRequest(original["sign_request"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSignRequest); val.IsValid() && !isEmptyValue(val) {
		transformed["signRequest"] = transformedSignRequest
	}

	transformedIdpCertificates, err := expandIdentityPlatformInboundSamlConfigIdpConfigIdpCertificates(original["idp_certificates"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdpCertificates); val.IsValid() && !isEmptyValue(val) {
		transformed["idpCertificates"] = transformedIdpCertificates
	}

	return transformed, nil
}

func expandIdentityPlatformInboundSamlConfigIdpConfigIdpEntityId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformInboundSamlConfigIdpConfigSsoUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformInboundSamlConfigIdpConfigSignRequest(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformInboundSamlConfigIdpConfigIdpCertificates(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedX509Certificate, err := expandIdentityPlatformInboundSamlConfigIdpConfigIdpCertificatesX509Certificate(original["x509_certificate"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedX509Certificate); val.IsValid() && !isEmptyValue(val) {
			transformed["x509Certificate"] = transformedX509Certificate
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandIdentityPlatformInboundSamlConfigIdpConfigIdpCertificatesX509Certificate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformInboundSamlConfigSpConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSpEntityId, err := expandIdentityPlatformInboundSamlConfigSpConfigSpEntityId(original["sp_entity_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSpEntityId); val.IsValid() && !isEmptyValue(val) {
		transformed["spEntityId"] = transformedSpEntityId
	}

	transformedCallbackUri, err := expandIdentityPlatformInboundSamlConfigSpConfigCallbackUri(original["callback_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCallbackUri); val.IsValid() && !isEmptyValue(val) {
		transformed["callbackUri"] = transformedCallbackUri
	}

	transformedSpCertificates, err := expandIdentityPlatformInboundSamlConfigSpConfigSpCertificates(original["sp_certificates"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSpCertificates); val.IsValid() && !isEmptyValue(val) {
		transformed["spCertificates"] = transformedSpCertificates
	}

	return transformed, nil
}

func expandIdentityPlatformInboundSamlConfigSpConfigSpEntityId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformInboundSamlConfigSpConfigCallbackUri(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformInboundSamlConfigSpConfigSpCertificates(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedX509Certificate, err := expandIdentityPlatformInboundSamlConfigSpConfigSpCertificatesX509Certificate(original["x509_certificate"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedX509Certificate); val.IsValid() && !isEmptyValue(val) {
			transformed["x509Certificate"] = transformedX509Certificate
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandIdentityPlatformInboundSamlConfigSpConfigSpCertificatesX509Certificate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
