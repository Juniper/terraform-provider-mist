package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func switchMgmtSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingSwitchMgmt) SwitchMgmtValue {

	var ap_affinity_threshold basetypes.Int64Value

	if d.ApAffinityThreshold != nil {
		ap_affinity_threshold = types.Int64Value(int64(*d.ApAffinityThreshold))
	}

	data_map_attr_type := SwitchMgmtValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"ap_affinity_threshold": ap_affinity_threshold,
	}
	data, e := NewSwitchMgmtValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data

}
