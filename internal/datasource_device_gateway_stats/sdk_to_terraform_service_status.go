package datasource_device_gateway_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func serviceStatusSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsGatewayServiceStatus) basetypes.ObjectValue {

	var appidInstallResult basetypes.StringValue
	var appidInstallTimestamp basetypes.StringValue
	var appidStatus basetypes.StringValue
	var appidVersion basetypes.Int64Value
	var ewfStatus basetypes.StringValue
	var idpInstallResult basetypes.StringValue
	var idpInstallTimestamp basetypes.StringValue
	var idpPolicy basetypes.StringValue
	var idpStatus basetypes.StringValue
	var idpUpdateTimestamp basetypes.StringValue

	if d.AppidInstallResult != nil {
		appidInstallResult = types.StringValue(*d.AppidInstallResult)
	}
	if d.AppidInstallTimestamp != nil {
		appidInstallTimestamp = types.StringValue(*d.AppidInstallTimestamp)
	}
	if d.AppidStatus != nil {
		appidStatus = types.StringValue(*d.AppidStatus)
	}
	if d.AppidVersion != nil {
		appidVersion = types.Int64Value(int64(*d.AppidVersion))
	}
	if d.EwfStatus != nil {
		ewfStatus = types.StringValue(*d.EwfStatus)
	}
	if d.IdpInstallResult != nil {
		idpInstallResult = types.StringValue(*d.IdpInstallResult)
	}
	if d.IdpInstallTimestamp != nil {
		idpInstallTimestamp = types.StringValue(*d.IdpInstallTimestamp)
	}
	if d.IdpPolicy != nil {
		idpPolicy = types.StringValue(*d.IdpPolicy)
	}
	if d.IdpStatus != nil {
		idpStatus = types.StringValue(*d.IdpStatus)
	}
	if d.IdpUpdateTimestamp != nil {
		idpUpdateTimestamp = types.StringValue(*d.IdpUpdateTimestamp)
	}

	dataMapValue := map[string]attr.Value{
		"appid_install_result":    appidInstallResult,
		"appid_install_timestamp": appidInstallTimestamp,
		"appid_status":            appidStatus,
		"appid_version":           appidVersion,
		"ewf_status":              ewfStatus,
		"idp_install_result":      idpInstallResult,
		"idp_install_timestamp":   idpInstallTimestamp,
		"idp_policy":              idpPolicy,
		"idp_status":              idpStatus,
		"idp_update_timestamp":    idpUpdateTimestamp,
	}
	data, e := basetypes.NewObjectValue(ServiceStatusValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
