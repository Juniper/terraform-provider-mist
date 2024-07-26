package datasource_org_servicepolicies

import (
	"context"
	"math/big"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l []models.OrgServicePolicy) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		elem := servicepolicieSdkToTerraform(ctx, &diags, d)
		elements = append(elements, elem)
	}

	dataSet, err := types.SetValue(OrgServicepoliciesValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func servicepolicieSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.OrgServicePolicy) OrgServicepoliciesValue {

	var action types.String
	var appqoe basetypes.ObjectValue = types.ObjectNull(AppqoeValue{}.AttributeTypes(ctx))
	var created_time basetypes.NumberValue
	var ewf types.List = types.ListNull(EwfValue{}.Type(ctx))
	var id types.String
	var idp basetypes.ObjectValue = types.ObjectNull(IdpValue{}.AttributeTypes(ctx))
	var local_routing types.Bool
	var modified_time basetypes.NumberValue
	var name types.String
	var org_id types.String
	var path_preference types.String
	var services types.List = types.ListNull(types.StringType)
	var tenants types.List = types.ListNull(types.StringType)

	if d.Action != nil {
		action = types.StringValue(string(*d.Action))
	}
	if d.Appqoe != nil {
		appqoe = appQoeToTerraform(ctx, diags, d.Appqoe)
	}
	if d.CreatedTime != nil {
		created_time = types.NumberValue(big.NewFloat(*d.CreatedTime))
	}
	if d.Ewf != nil {
		ewf = ewfSdkToTerraform(ctx, diags, d.Ewf)
	}
	id = types.StringValue(d.Id.String())

	if d.Idp != nil {
		idp = idpSdkToTerraform(ctx, diags, d.Idp)
	}
	if d.ModifiedTime != nil {
		modified_time = types.NumberValue(big.NewFloat(*d.ModifiedTime))
	}
	if d.LocalRouting != nil {
		local_routing = types.BoolValue(*d.LocalRouting)
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	org_id = types.StringValue(d.OrgId.String())

	if d.PathPreference != nil {
		path_preference = types.StringValue(*d.PathPreference)
	}
	if d.Services != nil {
		services = mist_transform.ListOfStringSdkToTerraform(ctx, d.Services)
	}
	if d.Tenants != nil {
		tenants = mist_transform.ListOfStringSdkToTerraform(ctx, d.Tenants)
	}

	data_map_attr_type := OrgServicepoliciesValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"created_time":    created_time,
		"action":          action,
		"appqoe":          appqoe,
		"ewf":             ewf,
		"id":              id,
		"idp":             idp,
		"local_routing":   local_routing,
		"modified_time":   modified_time,
		"name":            name,
		"org_id":          org_id,
		"path_preference": path_preference,
		"services":        services,
		"tenants":         tenants,
	}
	data, e := NewOrgServicepoliciesValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
