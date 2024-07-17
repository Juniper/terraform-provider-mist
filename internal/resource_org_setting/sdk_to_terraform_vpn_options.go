package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func vpnOptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingVpnOptions) VpnOptionsValue {

	var as_base basetypes.Int64Value
	var st_subnet basetypes.StringValue

	if d.AsBase != nil {
		as_base = types.Int64Value(int64(*d.AsBase))
	}
	if d.StSubnet != nil {
		st_subnet = types.StringValue(*d.StSubnet)
	}

	data_map_attr_type := VpnOptionsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"as_base":   as_base,
		"st_subnet": st_subnet,
	}
	data, e := NewVpnOptionsValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data

}
