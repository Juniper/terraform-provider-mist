package datasource_device_gateway_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func serviceStatusSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.GatewayStatsServiceStatus) basetypes.ObjectValue {

	var appid_install_result basetypes.StringValue
	var appid_install_timestamp basetypes.StringValue
	var appid_status basetypes.StringValue
	var appid_version basetypes.Int64Value
	var ewf_status basetypes.StringValue
	var idp_install_result basetypes.StringValue
	var idp_install_timestamp basetypes.StringValue
	var idp_policy basetypes.StringValue
	var idp_status basetypes.StringValue
	var idp_update_timestamp basetypes.StringValue

	if d.AppidInstallResult != nil {
		appid_install_result = types.StringValue(*d.AppidInstallResult)
	}
	if d.AppidInstallTimestamp != nil {
		appid_install_timestamp = types.StringValue(*d.AppidInstallTimestamp)
	}
	if d.AppidStatus != nil {
		appid_status = types.StringValue(*d.AppidStatus)
	}
	if d.AppidVersion != nil {
		appid_version = types.Int64Value(int64(*d.AppidVersion))
	}
	if d.EwfStatus != nil {
		ewf_status = types.StringValue(*d.EwfStatus)
	}
	if d.IdpInstallResult != nil {
		idp_install_result = types.StringValue(*d.IdpInstallResult)
	}
	if d.IdpInstallTimestamp != nil {
		idp_install_timestamp = types.StringValue(*d.IdpInstallTimestamp)
	}
	if d.IdpPolicy != nil {
		idp_policy = types.StringValue(*d.IdpPolicy)
	}
	if d.IdpStatus != nil {
		idp_status = types.StringValue(*d.IdpStatus)
	}
	if d.IdpUpdateTimestamp != nil {
		idp_update_timestamp = types.StringValue(*d.IdpUpdateTimestamp)
	}

	data_map_attr_type := ServiceStatusValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"appid_install_result":    appid_install_result,
		"appid_install_timestamp": appid_install_timestamp,
		"appid_status":            appid_status,
		"appid_version":           appid_version,
		"ewf_status":              ewf_status,
		"idp_install_result":      idp_install_result,
		"idp_install_timestamp":   idp_install_timestamp,
		"idp_policy":              idp_policy,
		"idp_status":              idp_status,
		"idp_update_timestamp":    idp_update_timestamp,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
