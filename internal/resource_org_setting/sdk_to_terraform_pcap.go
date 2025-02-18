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
	var maxPktLen basetypes.Int64Value

	if d.Bucket != nil {
		bucket = types.StringValue(*d.Bucket)
	}
	if d.MaxPktLen != nil {
		maxPktLen = types.Int64Value(int64(*d.MaxPktLen))
	}

	dataMapValue := map[string]attr.Value{
		"bucket":      bucket,
		"max_pkt_len": maxPktLen,
	}
	data, e := NewPcapValue(PcapValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data

}
