package resource_org_nactag

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data models.NacTag) (OrgNactagModel, diag.Diagnostics) {
	var state OrgNactagModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.Id.String())
	state.OrgId = types.StringValue(data.OrgId.String())
	state.Name = types.StringValue(data.Name)

	state.AllowUsermacOverride = types.BoolValue(*data.AllowUsermacOverride)
	state.EgressVlanNames = mist_transform.ListOfStringSdkToTerraform(ctx, data.EgressVlanNames)
	state.GbpTag = types.Int64Value(int64(*data.GbpTag))
	state.Match = types.StringValue(string(*data.Match))
	state.MatchAll = types.BoolValue(*data.MatchAll)
	state.RadiusAttrs = mist_transform.ListOfStringSdkToTerraform(ctx, data.RadiusAttrs)
	state.RadiusGroup = types.StringValue(*data.RadiusGroup)
	state.RadiusVendorAttrs = mist_transform.ListOfStringSdkToTerraform(ctx, data.RadiusVendorAttrs)
	state.SessionTimeout = types.Int64Value(int64(*data.SessionTimeout))
	state.Type = types.StringValue(string(data.Type))
	state.Values = mist_transform.ListOfStringSdkToTerraform(ctx, data.Values)
	state.Vlan = types.StringValue(*data.Vlan)

	return state, diags
}
