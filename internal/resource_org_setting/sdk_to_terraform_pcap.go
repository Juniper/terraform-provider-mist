package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func pcapSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingPcap) PcapValue {

	var bucket basetypes.StringValue
	var max_pkt_len basetypes.Int64Value

	if d.Bucket != nil {
		bucket = types.StringValue(*d.Bucket)
	}
	if d.MaxPktLen != nil {
		max_pkt_len = types.Int64Value(int64(*d.MaxPktLen))
	}

	data_map_attr_type := PcapValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"bucket":      bucket,
		"max_pkt_len": max_pkt_len,
	}
	data, e := NewPcapValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data

}
