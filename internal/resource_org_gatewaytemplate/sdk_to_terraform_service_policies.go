package resource_org_gatewaytemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func servicePolicyAppQoESdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ServicePolicyAppqoe) basetypes.ObjectValue {
	var enabled = types.BoolValue(false)

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	rAttrType := AppqoeValue{}.AttributeTypes(ctx)
	rAttrValue := map[string]attr.Value{
		"enabled": enabled,
	}
	r, e := basetypes.NewObjectValue(rAttrType, rAttrValue)
	diags.Append(e...)
	return r
}

func servicePolicyEwfSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.ServicePolicyEwfRule) basetypes.ListValue {
	var dataList []EwfValue
	for _, v := range d {
		var alertOnly basetypes.BoolValue
		var blockMessage basetypes.StringValue
		var enabled = types.BoolValue(false)
		var profile = types.StringValue("strict")

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
	datalistType := EwfValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)
	return r
}

func servicePolicyIdpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.IdpConfig) basetypes.ObjectValue {
	var alertOnly basetypes.BoolValue
	var enabled = types.BoolValue(false)
	var idpprofileId basetypes.StringValue
	var profile = types.StringValue("strict")

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

	rAttrType := IdpValue{}.AttributeTypes(ctx)
	rAttrValue := map[string]attr.Value{
		"alert_only":    alertOnly,
		"enabled":       enabled,
		"idpprofile_id": idpprofileId,
		"profile":       profile,
	}
	r, e := basetypes.NewObjectValue(rAttrType, rAttrValue)
	diags.Append(e...)
	return r
}

func servicePoliciesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.ServicePolicy) basetypes.ListValue {
	var dataList []ServicePoliciesValue

	for _, v := range d {

		var action basetypes.StringValue
		var appqoe = types.ObjectNull(AppqoeValue{}.AttributeTypes(ctx))
		var ewf = types.ListNull(EwfValue{}.Type(ctx))
		var idp = types.ObjectNull(IdpValue{}.AttributeTypes(ctx))
		var localRouting basetypes.BoolValue
		var name basetypes.StringValue
		var pathPreference basetypes.StringValue
		var servicepolicyId basetypes.StringValue
		var services = misttransform.ListOfStringSdkToTerraform(v.Services)
		var tenants = misttransform.ListOfStringSdkToTerraform(v.Tenants)

		if v.Action != nil {
			action = types.StringValue(string(*v.Action))
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
		if v.ServicepolicyId != nil {
			servicepolicyId = types.StringValue(v.ServicepolicyId.String())
		}

		dataMapValue := map[string]attr.Value{
			"action":           action,
			"appqoe":           appqoe,
			"ewf":              ewf,
			"idp":              idp,
			"local_routing":    localRouting,
			"name":             name,
			"path_preference":  pathPreference,
			"servicepolicy_id": servicepolicyId,
			"services":         services,
			"tenants":          tenants,
		}

		data, e := NewServicePoliciesValue(ServicePoliciesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		dataList = append(dataList, data)
	}
	datalistType := ServicePoliciesValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)
	return r
}
