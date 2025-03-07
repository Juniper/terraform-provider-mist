package resource_site_wxtag

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data models.WxlanTag) (SiteWxtagModel, diag.Diagnostics) {
	var state SiteWxtagModel
	var diags diag.Diagnostics

	var mac types.String
	var match types.String
	var op types.String
	var specs = types.ListNull(SpecsValue{}.Type(ctx))
	var values = types.ListNull(types.StringType)
	var vlanId types.String

	if data.Mac.Value() != nil {
		mac = types.StringValue(*data.Mac.Value())
	}
	if data.Match != nil {
		match = types.StringValue(string(*data.Match))
	}
	if data.Op != nil {
		op = types.StringValue(string(*data.Op))
	}
	if data.Specs != nil {
		specs = specsSdkToTerraform(ctx, &diags, data.Specs)
	}
	if data.Values != nil {
		values = mistutils.ListOfStringSdkToTerraform(data.Values)
	}
	if data.VlanId != nil {
		vlanId = types.StringValue(data.VlanId.String())
	}

	state.Id = types.StringValue(data.Id.String())
	state.SiteId = types.StringValue(data.SiteId.String())

	state.Name = types.StringValue(data.Name)

	state.Mac = mac
	state.Match = match
	state.Op = op

	state.Specs = specs

	state.Type = types.StringValue(string(data.Type))
	state.Values = values
	state.VlanId = vlanId

	return state, diags
}
