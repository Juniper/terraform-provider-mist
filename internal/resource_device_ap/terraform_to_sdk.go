package resource_device_ap

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *DeviceApModel) (models.MistDevice, diag.Diagnostics) {
	data := models.DeviceAp{}
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	deviceprofile_id, e := uuid.Parse(plan.DeviceprofileId.ValueString())
	if e == nil {
		data.DeviceprofileId = models.NewOptional(&deviceprofile_id)
	} else {
		unset["deviceprofile_id"] = nil
	}
	map_id, e := uuid.Parse(plan.MapId.ValueString())
	if e == nil {
		data.MapId = &map_id
	} else {
		unset["map_id"] = nil
	}

	data.Name = plan.Name.ValueStringPointer()
	data.Notes = plan.Notes.ValueStringPointer()

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

	if !plan.Centrak.IsNull() && !plan.Centrak.IsUnknown() {
		data.Centrak = centrakTerraformToSdk(ctx, &diags, plan.Centrak)
	} else {
		unset["-centrak"] = ""
	}

	if !plan.ClientBridge.IsNull() && !plan.ClientBridge.IsUnknown() {
		data.ClientBridge = clientBridgeTerraformToSdk(ctx, &diags, plan.ClientBridge)
	} else {
		unset["-client_bridge"] = ""
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

	if !plan.Height.IsNull() && !plan.Height.IsUnknown() {
		data.Height = plan.Height.ValueFloat64Pointer()
	} else {
		unset["-height"] = ""
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

	if !plan.Locked.IsNull() && !plan.Locked.IsUnknown() {
		data.Locked = plan.Locked.ValueBoolPointer()
	} else {
		unset["-locked"] = ""
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

	if !plan.Orientation.IsNull() && !plan.Orientation.IsUnknown() {
		data.Orientation = models.ToPointer(int(plan.Orientation.ValueInt64()))
	} else {
		unset["-orientation"] = ""
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

	if !plan.X.IsNull() && !plan.X.IsUnknown() {
		data.X = plan.X.ValueFloat64Pointer()
	} else {
		unset["-x"] = ""
	}
	if !plan.Y.IsNull() && !plan.Y.IsUnknown() {
		data.Y = plan.Y.ValueFloat64Pointer()
	} else {
		unset["-y"] = ""
	}

	data.AdditionalProperties = unset

	mist_device := models.MistDeviceContainer.FromDeviceAp(data)
	return mist_device, diags
}
