package resource_org_setting

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func mgmtSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingMgmt) MgmtValue {
	var mxtunnel_ids basetypes.ListValue = mist_transform.ListOfUuidSdkToTerraformEmpty(ctx)
	var use_mxtunnel basetypes.BoolValue
	var use_wxtunnel basetypes.BoolValue

	if d.MxtunnelIds != nil {
		mxtunnel_ids = mist_transform.ListOfUuidSdkToTerraform(ctx, d.MxtunnelIds)
	}
	if d.UseMxtunnel != nil {
		use_mxtunnel = types.BoolValue(*d.UseMxtunnel)
	}
	if d.UseWxtunnel != nil {
		use_wxtunnel = types.BoolValue(*d.UseWxtunnel)
	}

	data_map_attr_type := MgmtValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"mxtunnel_ids": mxtunnel_ids,
		"use_mxtunnel": use_mxtunnel,
		"use_wxtunnel": use_wxtunnel,
	}
	data, e := NewMgmtValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
