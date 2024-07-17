package resource_site_wxtag

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data models.WxlanTag) (SiteWxtagModel, diag.Diagnostics) {
	var state SiteWxtagModel
	var diags diag.Diagnostics

	var last_ips types.List = types.ListNull(types.StringType)
	var mac types.String
	var match types.String
	var op types.String
	var resource_mac types.String
	var services types.List = types.ListNull(types.StringType)
	var specs types.List = types.ListNull(SpecsValue{}.Type(ctx))
	var subnet types.String
	var values types.List = types.ListNull(types.StringType)
	var vlan_id types.Int64

	if data.LastIps != nil {
		last_ips = mist_transform.ListOfStringSdkToTerraform(ctx, data.LastIps)
	}
	if data.Mac.Value() != nil {
		mac = types.StringValue(*data.Mac.Value())
	}
	if data.Match != nil {
		match = types.StringValue(string(*data.Match))
	}
	if data.Op != nil {
		op = types.StringValue(string(*data.Op))
	}
	if data.ResourceMac.Value() != nil {
		resource_mac = types.StringValue(*data.ResourceMac.Value())
	}
	if data.Services != nil {
		services = mist_transform.ListOfStringSdkToTerraform(ctx, data.Services)
	}
	if data.Specs != nil {
		specs = specsSdkToTerraform(ctx, &diags, data.Specs)
	}
	if data.Subnet != nil {
		subnet = types.StringValue(*data.Subnet)
	}
	if data.Values != nil {
		values = mist_transform.ListOfStringSdkToTerraform(ctx, data.Values)
	}
	if data.VlanId != nil {
		vlan_id = types.Int64Value(int64(*data.VlanId))
	}

	state.Id = types.StringValue(data.Id.String())
	state.OrgId = types.StringValue(data.OrgId.String())
	state.SiteId = types.StringValue(data.SiteId.String())

	state.Name = types.StringValue(data.Name)

	state.LastIps = last_ips
	state.Mac = mac
	state.Match = match
	state.Op = op
	state.ResourceMac = resource_mac
	state.Services = services

	state.Specs = specs

	state.Subnet = subnet
	state.Type = types.StringValue(string(data.Type))
	state.Values = values
	state.VlanId = vlan_id

	return state, diags
}
