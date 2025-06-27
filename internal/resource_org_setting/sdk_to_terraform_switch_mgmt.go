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

	var apAffinityThreshold basetypes.Int64Value
	var removeExistingConfigs basetypes.BoolValue

	if d.ApAffinityThreshold != nil {
		apAffinityThreshold = types.Int64Value(int64(*d.ApAffinityThreshold))
	}

	dataMapValue := map[string]attr.Value{
		"ap_affinity_threshold":   apAffinityThreshold,
		"remove_existing_configs": removeExistingConfigs,
	}
	data, e := NewSwitchMgmtValue(SwitchMgmtValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data

}
