package datasource_org_wlans

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func hotspot20SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanHotspot20) basetypes.ObjectValue {
	var domain_name basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var enabled basetypes.BoolValue
	var nai_realms basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var operators basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var rcoi basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var venue_name basetypes.StringValue

	if d != nil && d.DomainName != nil {
		domain_name = mist_transform.ListOfStringSdkToTerraform(ctx, d.DomainName)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.NaiRealms != nil {
		nai_realms = mist_transform.ListOfStringSdkToTerraform(ctx, d.NaiRealms)
	}
	if d != nil && d.Operators != nil {
		var operators_list []attr.Value
		for _, v := range d.Operators {
			operators_list = append(operators_list, types.StringValue(string(v)))
		}
		operators = types.ListValueMust(basetypes.StringType{}, operators_list)
	}
	if d != nil && d.Rcoi != nil {
		rcoi = mist_transform.ListOfStringSdkToTerraform(ctx, d.Rcoi)
	}
	if d != nil && d.VenueName != nil {
		venue_name = types.StringValue(*d.VenueName)
	}

	data_map_attr_type := Hotspot20Value{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"domain_name": domain_name,
		"enabled":     enabled,
		"nai_realms":  nai_realms,
		"operators":   operators,
		"rcoi":        rcoi,
		"venue_name":  venue_name,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
