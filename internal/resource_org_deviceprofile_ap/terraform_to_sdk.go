package resource_org_deviceprofile_ap

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgDeviceprofileApModel) (models.Deviceprofile, diag.Diagnostics) {
	var data models.DeviceprofileAp
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	data.Name = plan.Name.ValueStringPointer()

	if !plan.Aeroscout.IsNull() && !plan.Aeroscout.IsUnknown() {
		data.Aeroscout = aeroscoutTerraformToSdk(plan.Aeroscout)
	} else {
		unset["-aeroscout"] = ""
	}

	if !plan.BleConfig.IsNull() && !plan.BleConfig.IsUnknown() {
		data.BleConfig = bleConfigTerraformToSdk(plan.BleConfig)
	} else {
		unset["-ble_config"] = ""
	}

	if !plan.DisableEth1.IsNull() && !plan.DisableEth1.IsUnknown() {
		data.DisableEth1 = plan.DisableEth1.ValueBoolPointer()
	} else {
		unset["-disable_eth1"] = ""
	}
	if !plan.DisableEth2.IsNull() && !plan.DisableEth2.IsUnknown() {
		data.DisableEth2 = plan.DisableEth2.ValueBoolPointer()
	} else {
		unset["-disable_eth2"] = ""
	}
	if !plan.DisableEth3.IsNull() && !plan.DisableEth3.IsUnknown() {
		data.DisableEth3 = plan.DisableEth3.ValueBoolPointer()
	} else {
		unset["-disable_eth3"] = ""
	}
	if !plan.DisableModule.IsNull() && !plan.DisableModule.IsUnknown() {
		data.DisableModule = plan.DisableModule.ValueBoolPointer()
	} else {
		unset["-disable_eth3"] = ""
	}
	if !plan.EslConfig.IsNull() && !plan.EslConfig.IsUnknown() {
		data.EslConfig = eslTerraformToSdk(plan.EslConfig)
	} else {
		unset["-esl_config"] = ""
	}

	if !plan.IpConfig.IsNull() && !plan.IpConfig.IsUnknown() {
		data.IpConfig = ipConfigTerraformToSdk(plan.IpConfig)
	} else {
		unset["-ip_config"] = ""
	}

	if !plan.LacpConfig.IsNull() && !plan.LacpConfig.IsUnknown() {
		data.LacpConfig = lacpConfigTerraformToSdk(plan.LacpConfig)
	} else {
		unset["-lacp_config"] = ""
	}

	if !plan.Led.IsNull() && !plan.Led.IsUnknown() {
		data.Led = ledTerraformToSdk(plan.Led)
	} else {
		unset["-led"] = ""
	}

	if !plan.Mesh.IsNull() && !plan.Mesh.IsUnknown() {
		data.Mesh = meshTerraformToSdk(plan.Mesh)
	} else {
		unset["-mesh"] = ""
	}

	if !plan.NtpServers.IsNull() && !plan.NtpServers.IsUnknown() {
		data.NtpServers = misttransform.ListOfStringTerraformToSdk(plan.NtpServers)
	} else {
		unset["-ntp_servers"] = ""
	}

	if !plan.PoePassthrough.IsNull() && !plan.PoePassthrough.IsUnknown() {
		data.PoePassthrough = plan.PoePassthrough.ValueBoolPointer()
	} else {
		unset["-poe_passthrough"] = ""
	}

	if !plan.PwrConfig.IsNull() && !plan.PwrConfig.IsUnknown() {
		data.PwrConfig = pwrConfigTerraformToSdk(plan.PwrConfig)
	} else {
		unset["-pwr_config"] = ""
	}

	if !plan.RadioConfig.IsNull() && !plan.RadioConfig.IsUnknown() {
		data.RadioConfig = radioConfigTerraformToSdk(ctx, &diags, plan.RadioConfig)
	} else {
		unset["-radio_config"] = ""
	}

	if !plan.UplinkPortConfig.IsNull() && !plan.UplinkPortConfig.IsUnknown() {
		data.UplinkPortConfig = uplinkPortConfigTerraformToSdk(plan.UplinkPortConfig)
	} else {
		unset["-uplink_port_config"] = ""
	}

	if !plan.UsbConfig.IsNull() && !plan.UsbConfig.IsUnknown() {
		data.UsbConfig = usbConfigTerraformToSdk(plan.UsbConfig)
	} else {
		unset["-usb_config"] = ""
	}

	if !plan.Vars.IsNull() && !plan.Vars.IsUnknown() {
		data.Vars = varsTerraformToSdk(plan.Vars)
	} else {
		unset["-vars"] = ""
	}

	data.Type = string(models.ConstDeviceTypeApEnum_AP)

	data.AdditionalProperties = unset

	deviceprofile := models.DeviceprofileContainer.FromDeviceprofileAp(data)
	return deviceprofile, diags
}
