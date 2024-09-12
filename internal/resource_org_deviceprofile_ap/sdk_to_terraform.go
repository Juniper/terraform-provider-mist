package resource_org_deviceprofile_ap

import (
	"context"

	mist_list "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *models.DeviceprofileAp) (OrgDeviceprofileApModel, diag.Diagnostics) {
	var state OrgDeviceprofileApModel
	var diags diag.Diagnostics

	var aeroscout AeroscoutValue = NewAeroscoutValueNull()
	var ble_config BleConfigValue = NewBleConfigValueNull()
	var disable_eth1 types.Bool
	var disable_eth2 types.Bool
	var disable_eth3 types.Bool
	var disable_module types.Bool
	var esl_config EslConfigValue = NewEslConfigValueNull()
	var ip_config IpConfigValue = NewIpConfigValueNull()
	var led LedValue = NewLedValueNull()
	var mesh MeshValue = NewMeshValueNull()
	var name types.String
	var ntp_servers types.List = types.ListNull(types.StringType)
	var org_id types.String
	var poe_passthrough types.Bool
	var profile_id types.String
	var pwr_config PwrConfigValue = NewPwrConfigValueNull()
	var radio_config RadioConfigValue = NewRadioConfigValueNull()
	var site_id types.String
	var uplink_port_config UplinkPortConfigValue = NewUplinkPortConfigValueNull()
	var usb_config UsbConfigValue = NewUsbConfigValueNull()
	var vars types.Map = types.MapNull(types.StringType)

	var profile_type types.String

	if data.Aeroscout != nil {
		aeroscout = aeroscoutSdkToTerraform(ctx, &diags, data.Aeroscout)
	}
	if data.BleConfig != nil {
		ble_config = bleConfigsSdkToTerraform(ctx, &diags, data.BleConfig)
	}
	if data.DisableEth1 != nil {
		disable_eth1 = types.BoolValue(*data.DisableEth1)
	}
	if data.DisableEth2 != nil {
		disable_eth2 = types.BoolValue(*data.DisableEth2)
	}
	if data.DisableEth3 != nil {
		disable_eth3 = types.BoolValue(*data.DisableEth3)
	}
	if data.DisableModule != nil {
		disable_module = types.BoolValue(*data.DisableModule)
	}
	if data.EslConfig != nil {
		esl_config = eslSdkToTerraform(ctx, &diags, data.EslConfig)
	}
	if data.Id != nil {
		profile_id = types.StringValue(data.Id.String())
	}
	if data.IpConfig != nil {
		ip_config = ipConfigSdkToTerraform(ctx, &diags, data.IpConfig)
	}
	if data.Led != nil {
		led = ledSdkToTerraform(ctx, &diags, data.Led)
	}
	if data.Mesh != nil {
		mesh = meshSdkToTerraform(ctx, &diags, data.Mesh)
	}

	name = types.StringValue(*data.Name)

	if data.NtpServers != nil {
		ntp_servers = mist_list.ListOfStringSdkToTerraform(ctx, data.NtpServers)
	}
	if data.OrgId != nil {
		org_id = types.StringValue(data.OrgId.String())
	}
	if data.PoePassthrough != nil {
		poe_passthrough = types.BoolValue(*data.PoePassthrough)
	}
	if data.PwrConfig != nil {
		pwr_config = pwrConfigSdkToTerraform(ctx, &diags, data.PwrConfig)
	}
	if data.RadioConfig != nil {
		radio_config = radioConfigSdkToTerraform(ctx, &diags, data.RadioConfig)
	}
	if data.SiteId != nil {
		site_id = types.StringValue(data.SiteId.String())
	}
	if data.UplinkPortConfig != nil {
		uplink_port_config = uplinkPortConfigSdkToTerraform(ctx, &diags, data.UplinkPortConfig)
	}
	if data.UsbConfig != nil {
		usb_config = usbConfigSdkToTerraform(ctx, &diags, data.UsbConfig)
	}
	if data.Vars != nil {
		vars = varsSdkToTerraform(ctx, &diags, data.Vars)
	}

	profile_type = types.StringValue(string(data.Type))

	state.Aeroscout = aeroscout
	state.BleConfig = ble_config
	state.DisableEth1 = disable_eth1
	state.DisableEth2 = disable_eth2
	state.DisableEth3 = disable_eth3
	state.DisableModule = disable_module
	state.EslConfig = esl_config
	state.Id = profile_id
	state.IpConfig = ip_config
	state.Led = led
	state.Mesh = mesh
	state.Name = name
	state.NtpServers = ntp_servers
	state.OrgId = org_id
	state.PoePassthrough = poe_passthrough
	state.PwrConfig = pwr_config
	state.RadioConfig = radio_config
	state.SiteId = site_id
	state.UplinkPortConfig = uplink_port_config
	state.UsbConfig = usb_config
	state.Vars = vars
	state.Type = profile_type

	return state, diags
}
