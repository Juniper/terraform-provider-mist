package resource_org_setting

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func mgmtSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingMgmt) MgmtValue {
	var mxtunnelIds = mistutils.ListOfUuidSdkToTerraformEmpty()
	var useMxtunnel basetypes.BoolValue
	var useWxtunnel basetypes.BoolValue

	if d.MxtunnelIds != nil {
		mxtunnelIds = mistutils.ListOfUuidSdkToTerraform(d.MxtunnelIds)
	}
	if d.UseMxtunnel != nil {
		useMxtunnel = types.BoolValue(*d.UseMxtunnel)
	}
	if d.UseWxtunnel != nil {
		useWxtunnel = types.BoolValue(*d.UseWxtunnel)
	}

	dataMapValue := map[string]attr.Value{
		"mxtunnel_ids": mxtunnelIds,
		"use_mxtunnel": useMxtunnel,
		"use_wxtunnel": useWxtunnel,
	}
	data, e := NewMgmtValue(MgmtValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
