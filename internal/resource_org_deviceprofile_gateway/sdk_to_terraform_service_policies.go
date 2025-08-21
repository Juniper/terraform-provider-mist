package resource_org_deviceprofile_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func servicePolicyAppQoESdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ServicePolicyAppqoe) basetypes.ObjectValue {
	var enabled basetypes.BoolValue

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	rAttrValue := map[string]attr.Value{
		"enabled": enabled,
	}
	r, e := basetypes.NewObjectValue(AppqoeValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)
	return r
}

func servicePolicyEwfSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.ServicePolicyEwfRule) basetypes.ListValue {
	var dataList []EwfValue
	for _, v := range d {
		var alertOnly basetypes.BoolValue
		var blockMessage basetypes.StringValue
		var enabled basetypes.BoolValue
		var profile basetypes.StringValue

		if v.AlertOnly != nil {
			alertOnly = types.BoolValue(*v.AlertOnly)
		}
		if v.BlockMessage != nil {
			blockMessage = types.StringValue(*v.BlockMessage)
		}
		if v.Enabled != nil {
			enabled = types.BoolValue(*v.Enabled)
		}
		if v.Profile != nil {
			profile = types.StringValue(string(*v.Profile))
		}

		dataMapValue := map[string]attr.Value{
			"alert_only":    alertOnly,
			"block_message": blockMessage,
			"enabled":       enabled,
			"profile":       profile,
		}
		data, e := NewEwfValue(EwfValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, EwfValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func servicePolicyIdpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.IdpConfig) basetypes.ObjectValue {
	var alertOnly basetypes.BoolValue
	var enabled basetypes.BoolValue
	var idpprofileId basetypes.StringValue
	var profile basetypes.StringValue

	if d != nil && d.AlertOnly != nil {
		alertOnly = types.BoolValue(*d.AlertOnly)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.IdpprofileId != nil {
		idpprofileId = types.StringValue(d.IdpprofileId.String())
	}
	if d != nil && d.Profile != nil {
		profile = types.StringValue(*d.Profile)
	}

	rAttrValue := map[string]attr.Value{
		"alert_only":    alertOnly,
		"enabled":       enabled,
		"idpprofile_id": idpprofileId,
		"profile":       profile,
	}
	r, e := basetypes.NewObjectValue(IdpValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)
	return r
}

func avSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ServicePolicyAntivirus) basetypes.ObjectValue {

	var avprofileId basetypes.StringValue
	var enabled basetypes.BoolValue
	var profile basetypes.StringValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.AvprofileId != nil {
		avprofileId = types.StringValue(d.AvprofileId.String())
	}
	if d.Profile != nil && (d.AvprofileId == nil || *d.Profile != d.AvprofileId.String()) {
		profile = types.StringValue(*d.Profile)
	}

	rAttrValue := map[string]attr.Value{
		"avprofile_id": avprofileId,
		"enabled":      enabled,
		"profile":      profile,
	}
	r, e := basetypes.NewObjectValue(AntivirusValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)

	return r
}

func sslProxySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ServicePolicySslProxy) basetypes.ObjectValue {

	var CiphersCategory basetypes.StringValue
	var enabled basetypes.BoolValue

	if d.CiphersCategory != nil {
		CiphersCategory = types.StringValue(string(*d.CiphersCategory))
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	rAttrValue := map[string]attr.Value{
		"ciphers_category": CiphersCategory,
		"enabled":          enabled,
	}
	r, e := basetypes.NewObjectValue(SslProxyValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)

	return r
}

func servicePoliciesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.ServicePolicy) basetypes.ListValue {
	var dataList []ServicePoliciesValue

	for _, v := range d {

		var action basetypes.StringValue
		var antivirus = types.ObjectNull(AntivirusValue{}.AttributeTypes(ctx))
		var appqoe = types.ObjectNull(AppqoeValue{}.AttributeTypes(ctx))
		var ewf = types.ListNull(EwfValue{}.Type(ctx))
		var idp = types.ObjectNull(IdpValue{}.AttributeTypes(ctx))
		var localRouting basetypes.BoolValue
		var name basetypes.StringValue
		var pathPreference basetypes.StringValue
		var servicepolicyId basetypes.StringValue
		var services = types.ListNull(types.StringType)
		var sslProxy = types.ObjectNull(SslProxyValue{}.AttributeTypes(ctx))
		var tenants = types.ListNull(types.StringType)

		if v.Action != nil {
			action = types.StringValue(string(*v.Action))
		}
		if v.Antivirus != nil {
			antivirus = avSdkToTerraform(ctx, diags, v.Antivirus)
		}
		if v.Appqoe != nil {
			appqoe = servicePolicyAppQoESdkToTerraform(ctx, diags, v.Appqoe)
		}
		if v.Ewf != nil {
			ewf = servicePolicyEwfSdkToTerraform(ctx, diags, v.Ewf)
		}
		if v.Idp != nil {
			idp = servicePolicyIdpSdkToTerraform(ctx, diags, v.Idp)
		}
		if v.LocalRouting != nil {
			localRouting = types.BoolValue(*v.LocalRouting)
		}
		if v.Name != nil {
			name = types.StringValue(*v.Name)
		}
		if v.PathPreference != nil {
			pathPreference = types.StringValue(*v.PathPreference)
		}
		if v.Services != nil {
			services = mistutils.ListOfStringSdkToTerraform(v.Services)
		}
		if v.ServicepolicyId != nil {
			servicepolicyId = types.StringValue(v.ServicepolicyId.String())
		}
		if v.SslProxy != nil {
			sslProxy = sslProxySdkToTerraform(ctx, diags, v.SslProxy)
		}
		if v.Tenants != nil {
			tenants = mistutils.ListOfStringSdkToTerraform(v.Tenants)
		}

		dataMapValue := map[string]attr.Value{
			"action":           action,
			"antivirus":        antivirus,
			"appqoe":           appqoe,
			"ewf":              ewf,
			"idp":              idp,
			"local_routing":    localRouting,
			"name":             name,
			"path_preference":  pathPreference,
			"servicepolicy_id": servicepolicyId,
			"services":         services,
			"ssl_proxy":        sslProxy,
			"tenants":          tenants,
		}

		data, e := NewServicePoliciesValue(ServicePoliciesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, ServicePoliciesValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}
