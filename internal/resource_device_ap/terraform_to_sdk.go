package resource_device_ap

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *DeviceApModel) (models.MistDevice, diag.Diagnostics) {
	data := models.DeviceAp{}
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	if len(plan.MapId.ValueString()) > 0 {
		mapId, e := uuid.Parse(plan.MapId.ValueString())
		if e == nil {
			data.MapId = &mapId
		} else {
			diags.AddError("Bad value for map_id", e.Error())
		}
	} else {
		unset["-map_id"] = ""
	}

	data.Name = plan.Name.ValueStringPointer()

	if !plan.Notes.IsNull() && !plan.Notes.IsUnknown() {
		data.Notes = plan.Notes.ValueStringPointer()
	} else {
		unset["-notes"] = ""
	}

	if !plan.Aeroscout.IsNull() && !plan.Aeroscout.IsUnknown() {
		aeroscout := aeroscoutTerraformToSdk(plan.Aeroscout)
		data.Aeroscout = aeroscout
	} else {
		unset["-aeroscout"] = ""
	}

	if !plan.BleConfig.IsNull() && !plan.BleConfig.IsUnknown() {
		data.BleConfig = bleConfigTerraformToSdk(plan.BleConfig)
	} else {
		unset["-ble_config"] = ""
	}

	if !plan.Centrak.IsNull() && !plan.Centrak.IsUnknown() {
		data.Centrak = centrakTerraformToSdk(plan.Centrak)
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
		data.EslConfig = eslTerraformToSdk(plan.EslConfig)
	} else {
		unset["-esl_config"] = ""
	}
	if !plan.FlowControl.IsNull() && !plan.FlowControl.IsUnknown() {
		data.FlowControl = plan.FlowControl.ValueBoolPointer()
	} else {
		unset["-flow_control"] = ""
	}

	if !plan.Height.IsNull() && !plan.Height.IsUnknown() {
		data.Height = plan.Height.ValueFloat64Pointer()
	} else {
		unset["-height"] = ""
	}

	if !plan.IpConfig.IsNull() && !plan.IpConfig.IsUnknown() {
		ipConfig := ipConfigTerraformToSdk(plan.IpConfig)
		data.IpConfig = ipConfig
	} else {
		unset["-ip_config"] = ""
	}

	if !plan.Led.IsNull() && !plan.Led.IsUnknown() {
		led := ledTerraformToSdk(plan.Led)
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
		mesh := meshTerraformToSdk(plan.Mesh)
		data.Mesh = mesh
	} else {
		unset["-mesh"] = ""
	}

	if !plan.NtpServers.IsNull() && !plan.NtpServers.IsUnknown() {
		data.NtpServers = misttransform.ListOfStringTerraformToSdk(plan.NtpServers)
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

	mistDevice := models.MistDeviceContainer.FromDeviceAp(data)
	return mistDevice, diags
}
