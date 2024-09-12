package resource_org_deviceprofile_ap

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgDeviceprofileApModel) (models.Deviceprofile, diag.Diagnostics) {
	data := models.DeviceprofileAp{}
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	data.Name = plan.Name.ValueStringPointer()

	if !plan.Aeroscout.IsNull() && !plan.Aeroscout.IsUnknown() {
		aeroscout := aeroscoutTerraformToSdk(ctx, &diags, plan.Aeroscout)
		data.Aeroscout = aeroscout
	} else {
		unset["-aeroscout"] = ""
	}

	if !plan.BleConfig.IsNull() && !plan.BleConfig.IsUnknown() {
		data.BleConfig = bleConfigTerraformToSdk(ctx, &diags, plan.BleConfig)
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
		data.EslConfig = eslTerraformToSdk(ctx, &diags, plan.EslConfig)
	} else {
		unset["-esl_config"] = ""
	}

	if !plan.IpConfig.IsNull() && !plan.IpConfig.IsUnknown() {
		ip_config := ipConfigTerraformToSdk(ctx, &diags, plan.IpConfig)
		data.IpConfig = ip_config
	} else {
		unset["-ip_config"] = ""
	}

	if !plan.Led.IsNull() && !plan.Led.IsUnknown() {
		led := ledTerraformToSdk(ctx, &diags, plan.Led)
		data.Led = led
	} else {
		unset["-led"] = ""
	}

	if !plan.Mesh.IsNull() && !plan.Mesh.IsUnknown() {
		mesh := meshTerraformToSdk(ctx, &diags, plan.Mesh)
		data.Mesh = mesh
	} else {
		unset["-mesh"] = ""
	}

	if !plan.NtpServers.IsNull() && !plan.NtpServers.IsUnknown() {
		data.NtpServers = mist_transform.ListOfStringTerraformToSdk(ctx, plan.NtpServers)
	} else {
		unset["-ntp_servers"] = ""
	}

	if !plan.PoePassthrough.IsNull() && !plan.PoePassthrough.IsUnknown() {
		data.PoePassthrough = plan.PoePassthrough.ValueBoolPointer()
	} else {
		unset["-poe_passthrough"] = ""
	}

	if !plan.PwrConfig.IsNull() && !plan.PwrConfig.IsUnknown() {
		data.PwrConfig = pwrConfigTerraformToSdk(ctx, &diags, plan.PwrConfig)
	} else {
		unset["-pwr_config"] = ""
	}

	if !plan.RadioConfig.IsNull() && !plan.RadioConfig.IsUnknown() {
		data.RadioConfig = radioConfigTerraformToSdk(ctx, &diags, plan.RadioConfig)
	} else {
		unset["-radio_config"] = ""
	}

	if !plan.UplinkPortConfig.IsNull() && !plan.UplinkPortConfig.IsUnknown() {
		data.UplinkPortConfig = uplinkPortConfigTerraformToSdk(ctx, &diags, plan.UplinkPortConfig)
	} else {
		unset["-uplink_port_config"] = ""
	}

	if !plan.UsbConfig.IsNull() && !plan.UsbConfig.IsUnknown() {
		data.UsbConfig = usbConfigTerraformToSdk(ctx, &diags, plan.UsbConfig)
	} else {
		unset["-usb_config"] = ""
	}

	if !plan.Vars.IsNull() && !plan.Vars.IsUnknown() {
		data.Vars = varsTerraformToSdk(ctx, &diags, plan.Vars)
	} else {
		unset["-vars"] = ""
	}

	data.Type = string(models.ConstDeviceTypeApEnum_AP)

	data.AdditionalProperties = unset

	deviceprofile := models.DeviceprofileContainer.FromDeviceprofileAp(data)
	return deviceprofile, diags
}
