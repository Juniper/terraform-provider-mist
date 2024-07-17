package resource_org_gatewaytemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func servicePolicyAppQoESdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ServicePolicyAppqoe) basetypes.ObjectValue {
	var enabled basetypes.BoolValue = types.BoolValue(false)

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	r_attr_type := AppqoeValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"enabled": enabled,
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func servicePolicyEwfSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.ServicePolicyEwfRule) basetypes.ListValue {
	var data_list = []EwfValue{}
	for _, v := range d {
		var alert_only basetypes.BoolValue
		var block_message basetypes.StringValue
		var enabled basetypes.BoolValue = types.BoolValue(false)
		var profile basetypes.StringValue = types.StringValue("strict")

		if v.AlertOnly != nil {
			alert_only = types.BoolValue(*v.AlertOnly)
		}
		if v.BlockMessage != nil {
			block_message = types.StringValue(*v.BlockMessage)
		}
		if v.Enabled != nil {
			enabled = types.BoolValue(*v.Enabled)
		}
		if v.Profile != nil {
			profile = types.StringValue(string(*v.Profile))
		}

		data_map_attr_type := EwfValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"alert_only":    alert_only,
			"block_message": block_message,
			"enabled":       enabled,
			"profile":       profile,
		}
		data, e := NewEwfValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := EwfValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func servicePolicyIdpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.IdpConfig) basetypes.ObjectValue {
	var alert_only basetypes.BoolValue
	var enabled basetypes.BoolValue = types.BoolValue(false)
	var idpprofile_id basetypes.StringValue
	var profile basetypes.StringValue = types.StringValue("strict")

	if d != nil && d.AlertOnly != nil {
		alert_only = types.BoolValue(*d.AlertOnly)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.IdpprofileId != nil {
		idpprofile_id = types.StringValue(d.IdpprofileId.String())
	}
	if d != nil && d.Profile != nil {
		profile = types.StringValue(string(*d.Profile))
	}

	r_attr_type := IdpValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"alert_only":    alert_only,
		"enabled":       enabled,
		"idpprofile_id": idpprofile_id,
		"profile":       profile,
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func servicePoliciesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.ServicePolicy) basetypes.ListValue {
	var data_list = []ServicePoliciesValue{}

	for _, v := range d {

		var action basetypes.StringValue = types.StringValue("allow")
		var appqoe basetypes.ObjectValue = types.ObjectNull(AppqoeValue{}.AttributeTypes(ctx))
		var ewf basetypes.ListValue = types.ListNull(EwfValue{}.Type(ctx))
		var idp basetypes.ObjectValue = types.ObjectNull(IdpProfilesValue{}.AttributeTypes(ctx))
		var local_routing basetypes.BoolValue = types.BoolValue(false)
		var name basetypes.StringValue
		var path_preferences basetypes.StringValue
		var servicepolicy_id basetypes.StringValue
		var services basetypes.ListValue = mist_transform.ListOfStringSdkToTerraform(ctx, v.Services)
		var tenants basetypes.ListValue = mist_transform.ListOfStringSdkToTerraform(ctx, v.Tenants)

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
			local_routing = types.BoolValue(*v.LocalRouting)
		}
		if v.Name != nil {
			name = types.StringValue(*v.Name)
		}
		if v.PathPreferences != nil {
			path_preferences = types.StringValue(*v.PathPreferences)
		}
		if v.ServicepolicyId != nil {
			servicepolicy_id = types.StringValue(v.ServicepolicyId.String())
		}

		data_map_attr_type := ServicePoliciesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"action":           action,
			"appqoe":           appqoe,
			"ewf":              ewf,
			"idp":              idp,
			"local_routing":    local_routing,
			"name":             name,
			"path_preferences": path_preferences,
			"servicepolicy_id": servicepolicy_id,
			"services":         services,
			"tenants":          tenants,
		}

		data, e := NewServicePoliciesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := ServicePoliciesValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}
