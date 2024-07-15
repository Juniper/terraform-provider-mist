package resource_org_wxtag

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data models.WxlanTag) (OrgWxtagModel, diag.Diagnostics) {
	var state OrgWxtagModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.Id.String())
	state.OrgId = types.StringValue(data.OrgId.String())
	state.SiteId = types.StringValue(data.SiteId.String())

	state.Name = types.StringValue(data.Name)
	state.LastIps = mist_transform.ListOfStringSdkToTerraform(ctx, data.LastIps)
	state.Mac = types.StringValue(*data.Mac.Value())
	state.Match = types.StringValue(string(*data.Match))
	state.Op = types.StringValue(string(*data.Op))
	state.ResourceMac = types.StringValue(*data.ResourceMac.Value())
	state.Services = mist_transform.ListOfStringSdkToTerraform(ctx, data.Services)

	specs := specsSdkToTerraform(ctx, &diags, data.Specs)
	state.Specs = specs

	state.Subnet = types.StringValue(*data.Subnet)
	state.Type = types.StringValue(string(data.Type))
	state.Values = mist_transform.ListOfStringSdkToTerraform(ctx, data.Values)
	state.VlanId = types.Int64Value(int64(*data.VlanId))

	return state, diags
}
