package resource_device_ap

import (
	"context"

	mist_list "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *models.DeviceAp) (DeviceApModel, diag.Diagnostics) {
	var state DeviceApModel
	var diags diag.Diagnostics

	var aeroscout AeroscoutValue = NewAeroscoutValueNull()
	var ble_config BleConfigValue = NewBleConfigValueNull()
	var centrak CentrakValue = NewCentrakValueNull()
	var client_bridge ClientBridgeValue = NewClientBridgeValueNull()
	var deviceprofile_id types.String
	var disable_eth1 types.Bool
	var disable_eth2 types.Bool
	var disable_eth3 types.Bool
	var disable_module types.Bool
	var esl_config EslConfigValue = NewEslConfigValueNull()
	var height types.Float64
	var deviceId types.String
	var image1_url types.String
	var image2_url types.String
	var image3_url types.String
	var ip_config IpConfigValue = NewIpConfigValueNull()
	var led LedValue = NewLedValueNull()
	var locked types.Bool
	var map_id types.String
	var mesh MeshValue = NewMeshValueNull()
	var name types.String
	var notes types.String
	var ntp_servers types.List = types.ListNull(types.StringType)
	var org_id types.String
	var orientation types.Int64
	var poe_passthrough types.Bool
	var pwr_config PwrConfigValue = NewPwrConfigValueNull()
	var radio_config RadioConfigValue = NewRadioConfigValueNull()
	var site_id types.String
	var uplink_port_config UplinkPortConfigValue = NewUplinkPortConfigValueNull()
	var usb_config UsbConfigValue = NewUsbConfigValueNull()
	var vars types.Map = types.MapNull(types.StringType)
	var x types.Float64
	var y types.Float64

	var device_type types.String
	var serial types.String
	var mac types.String
	var model types.String

	if data.Aeroscout != nil {
		aeroscout = aeroscoutSdkToTerraform(ctx, &diags, data.Aeroscout)
	}
	if data.BleConfig != nil {
		ble_config = bleConfigsSdkToTerraform(ctx, &diags, data.BleConfig)
	}
	if data.Centrak != nil {
		centrak = centrakSdkToTerraform(ctx, &diags, data.Centrak)
	}
	if data.ClientBridge != nil {
		client_bridge = clientBridgeSdkToTerraform(ctx, &diags, data.ClientBridge)
	}
	if data.DeviceprofileId.Value() != nil {
		deviceprofile_id = types.StringValue(data.DeviceprofileId.Value().String())
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
	if data.Height != nil {
		height = types.Float64Value(float64(*data.Height))
	}
	if data.Id != nil {
		deviceId = types.StringValue(data.Id.String())
	}
	if data.Image1Url.Value() != nil {
		image1_url = types.StringValue(*data.Image1Url.Value())
	}
	if data.Image2Url.Value() != nil {
		image2_url = types.StringValue(*data.Image2Url.Value())
	}
	if data.Image3Url.Value() != nil {
		image3_url = types.StringValue(*data.Image3Url.Value())
	}
	if data.IpConfig != nil {
		ip_config = ipConfigSdkToTerraform(ctx, &diags, data.IpConfig)
	}
	if data.Led != nil {
		led = ledSdkToTerraform(ctx, &diags, data.Led)
	}
	if data.Locked != nil {
		locked = types.BoolValue(*data.Locked)
	}
	if data.MapId != nil {
		map_id = types.StringValue(data.MapId.String())
	}
	if data.Mesh != nil {
		mesh = meshSdkToTerraform(ctx, &diags, data.Mesh)
	}
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}
	if data.Notes != nil {
		notes = types.StringValue(*data.Notes)
	}
	if data.NtpServers != nil {
		ntp_servers = mist_list.ListOfStringSdkToTerraform(ctx, data.NtpServers)
	}
	if data.OrgId != nil {
		org_id = types.StringValue(data.OrgId.String())
	}
	if data.Orientation != nil {
		orientation = types.Int64Value(int64(*data.Orientation))
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
	if data.X != nil {
		x = types.Float64Value(float64(*data.X))
	}
	if data.Y != nil {
		y = types.Float64Value(float64(*data.Y))
	}

	if data.Type != nil {
		device_type = types.StringValue(string(*data.Type))
	}

	if data.Serial != nil {
		serial = types.StringValue(*data.Serial)
	}

	if data.Mac != nil {
		mac = types.StringValue(*data.Mac)
	}

	if data.Model != nil {
		model = types.StringValue(*data.Model)
	}
	state.Aeroscout = aeroscout
	state.BleConfig = ble_config
	state.Centrak = centrak
	state.ClientBridge = client_bridge
	state.DisableEth1 = disable_eth1
	state.DisableEth2 = disable_eth2
	state.DisableEth3 = disable_eth3
	state.DisableModule = disable_module
	state.DeviceprofileId = deviceprofile_id
	state.EslConfig = esl_config
	state.Height = height
	state.DeviceId = deviceId
	state.Image1Url = image1_url
	state.Image2Url = image2_url
	state.Image3Url = image3_url
	state.IpConfig = ip_config
	state.Led = led
	state.Locked = locked
	state.MapId = map_id
	state.Mesh = mesh
	state.Name = name
	state.NtpServers = ntp_servers
	state.Notes = notes
	state.Orientation = orientation
	state.OrgId = org_id
	state.PoePassthrough = poe_passthrough
	state.PwrConfig = pwr_config
	state.RadioConfig = radio_config
	state.SiteId = site_id
	state.UplinkPortConfig = uplink_port_config
	state.UsbConfig = usb_config
	state.Vars = vars
	state.X = x
	state.Y = y
	state.Type = device_type
	state.Serial = serial
	state.Mac = mac
	state.Model = model

	return state, diags
}
