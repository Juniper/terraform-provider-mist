package datasource_device_ap_stats

import (
	"context"
	"encoding/json"
	"math/big"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func SdkToTerraform(ctx context.Context, l []models.ListOrgDevicesStatsResponse) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		ap_js, e := d.MarshalJSON()
		if e != nil {
			diags.AddError("Unable to Marshal AP Stats", e.Error())
		} else {
			ap := models.ApStats{}
			e := json.Unmarshal(ap_js, &ap)
			if e != nil {
				diags.AddError("Unable to unMarshal AP Stats", e.Error())
			}
			elem := deviceApStatSdkToTerraform(ctx, &diags, &ap)
			elements = append(elements, elem)
		}
	}

	dataSet, err := types.SetValue(DeviceApStatsValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func deviceApStatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApStats) DeviceApStatsValue {

	var auto_placement basetypes.ObjectValue = types.ObjectNull(AutoPlacementValue{}.AttributeTypes(ctx))
	var auto_upgrade_stat basetypes.ObjectValue = types.ObjectNull(AutoUpgradeStatValue{}.AttributeTypes(ctx))
	var ble_stat basetypes.ObjectValue = types.ObjectNull(BleStatValue{}.AttributeTypes(ctx))
	var cert_expiry basetypes.NumberValue
	var config_reverted basetypes.BoolValue
	var cpu_system basetypes.Int64Value
	var cpu_util basetypes.Int64Value
	var created_time basetypes.Int64Value
	var deviceprofile_id basetypes.StringValue
	var env_stat basetypes.ObjectValue = types.ObjectNull(EnvStatValue{}.AttributeTypes(ctx))
	var esl_stat basetypes.ObjectValue = types.ObjectNull(EslStatValue{}.AttributeTypes(ctx))
	var ext_ip basetypes.StringValue
	var fwupdate basetypes.ObjectValue = types.ObjectNull(FwupdateValue{}.AttributeTypes(ctx))
	var hw_rev basetypes.StringValue
	var id basetypes.StringValue
	var inactive_wired_vlans basetypes.ListValue = types.ListNull(types.Int64Type)
	var iot_stat basetypes.MapValue = types.MapNull(IotStatValue{}.Type(ctx))
	var ip basetypes.StringValue
	var ip_config basetypes.ObjectValue = types.ObjectNull(IpConfigValue{}.AttributeTypes(ctx))
	var ip_stat basetypes.ObjectValue = types.ObjectNull(IpStatValue{}.AttributeTypes(ctx))
	var l2tp_stat basetypes.MapValue = types.MapNull(L2tpStatValue{}.Type(ctx))
	var last_seen basetypes.NumberValue
	var last_trouble basetypes.ObjectValue = types.ObjectNull(LastTroubleValue{}.AttributeTypes(ctx))
	var led basetypes.ObjectValue = types.ObjectNull(LedValue{}.AttributeTypes(ctx))
	var lldp_stat basetypes.ObjectValue = types.ObjectNull(LldpStatValue{}.AttributeTypes(ctx))
	var locating basetypes.BoolValue
	var locked basetypes.BoolValue
	var mac basetypes.StringValue
	var map_id basetypes.StringValue
	var mem_used_kb basetypes.Int64Value
	var mesh_downlinks basetypes.MapValue = types.MapNull(MeshDownlinksValue{}.Type(ctx))
	var mesh_uplink basetypes.ObjectValue = types.ObjectNull(MeshUplinkValue{}.AttributeTypes(ctx))
	var model basetypes.StringValue
	var modified_time basetypes.Int64Value
	var mount basetypes.StringValue
	var name basetypes.StringValue
	var notes basetypes.StringValue
	var num_clients basetypes.Int64Value
	var org_id basetypes.StringValue
	var port_stat basetypes.MapValue = types.MapNull(PortStatValue{}.Type(ctx))
	var power_budget basetypes.Int64Value
	var power_constrained basetypes.BoolValue
	var power_opmode basetypes.StringValue
	var power_src basetypes.StringValue
	var radio_stat basetypes.ObjectValue = types.ObjectNull(RadioStatValue{}.AttributeTypes(ctx))
	var rx_bps basetypes.NumberValue
	var rx_bytes basetypes.Int64Value
	var rx_pkts basetypes.Int64Value
	var serial basetypes.StringValue
	var site_id basetypes.StringValue
	var status basetypes.StringValue
	var switch_redundancy basetypes.ObjectValue = types.ObjectNull(SwitchRedundancyValue{}.AttributeTypes(ctx))
	var tx_bps basetypes.NumberValue
	var tx_bytes basetypes.NumberValue
	var tx_pkts basetypes.NumberValue
	var uptime basetypes.NumberValue
	var usb_stat basetypes.ObjectValue = types.ObjectNull(UsbStatValue{}.AttributeTypes(ctx))
	var version basetypes.StringValue
	var x basetypes.Float64Value
	var y basetypes.Float64Value

	if d.AutoPlacement != nil {
		auto_placement = autoPlacementSdkToTerraform(ctx, diags, d.AutoPlacement)
	}
	if d.AutoUpgradeStat != nil {
		auto_upgrade_stat = autoUpgradeStatsSdkToTerraform(ctx, diags, d.AutoUpgradeStat)
	}
	if d.BleStat != nil {
		ble_stat = bleStatsSdkToTerraform(ctx, diags, d.BleStat)
	}
	if d.CertExpiry.Value() != nil {
		cert_expiry = types.NumberValue(big.NewFloat(*d.CertExpiry.Value()))
	}
	if d.ConfigReverted.Value() != nil {
		config_reverted = types.BoolValue(*d.ConfigReverted.Value())
	}
	if d.CpuSystem.Value() != nil {
		cpu_system = types.Int64Value(int64(*d.CpuSystem.Value()))
	}
	if d.CpuUtil.Value() != nil {
		cpu_util = types.Int64Value(int64(*d.CpuUtil.Value()))
	}
	if d.CreatedTime != nil {
		created_time = types.Int64Value(int64(*d.CreatedTime))
	}
	if d.DeviceprofileId.Value() != nil {
		deviceprofile_id = types.StringValue(d.DeviceprofileId.Value().String())
	}
	if d.EnvStat != nil {
		env_stat = envStatsSdkToTerraform(ctx, diags, d.EnvStat)
	}
	if d.EslStat.Value() != nil {
		esl_stat = eslStatsSdkToTerraform(ctx, diags, d.EslStat.Value())
	}
	if d.ExtIp.Value() != nil {
		ext_ip = types.StringValue(*d.ExtIp.Value())
	}
	if d.Fwupdate != nil {
		fwupdate = fwupdateSdkToTerraform(ctx, diags, d.Fwupdate)
	}
	if d.HwRev.Value() != nil {
		hw_rev = types.StringValue(*d.HwRev.Value())
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.InactiveWiredVlans != nil {
		inactive_wired_vlans = mist_transform.ListOfIntSdkToTerraform(ctx, d.InactiveWiredVlans)
	}
	if d.IotStat != nil && len(d.IotStat) > 0 {
		iot_stat = iotStatsSdkToTerraform(ctx, diags, d.IotStat)
	}
	if d.Ip.Value() != nil {
		ip = types.StringValue(*d.Ip.Value())
	}
	if d.IpConfig != nil {
		ip_config = ipConfigSdkToTerraform(ctx, diags, d.IpConfig)
	}
	if d.IpStat != nil {
		ip_stat = ipStatsSdkToTerraform(ctx, diags, d.IpStat)
	}
	if d.L2tpStat != nil && len(d.L2tpStat) > 0 {
		l2tp_stat = l2tpStatsSdkToTerraform(ctx, diags, d.L2tpStat)
	}
	if d.LastSeen.Value() != nil {
		last_seen = types.NumberValue(big.NewFloat(*d.LastSeen.Value()))
	}
	if d.LastTrouble != nil {
		last_trouble = lastTroubleSdkToTerraform(ctx, diags, d.LastTrouble)
	}
	if d.Led != nil {
		led = ledSdkToTerraform(ctx, diags, d.Led)
	}
	if d.LldpStat != nil {
		lldp_stat = lldpSdkToTerraform(ctx, diags, d.LldpStat)
	}
	if d.Locating.Value() != nil {
		locating = types.BoolValue(*d.Locating.Value())
	}
	if d.Locked.Value() != nil {
		locked = types.BoolValue(*d.Locked.Value())
	}
	if d.Mac != nil {
		mac = types.StringValue(*d.Mac)
	}
	if d.MapId.Value() != nil {
		map_id = types.StringValue(d.MapId.Value().String())
	}
	if d.MemUsedKb.Value() != nil {
		mem_used_kb = types.Int64Value(int64(*d.MemUsedKb.Value()))
	}
	if d.MeshDownlinks != nil && len(d.MeshDownlinks) > 0 {
		mesh_downlinks = meshDownlinksSdkToTerraform(ctx, diags, d.MeshDownlinks)
	}
	if d.MeshUplink != nil {
		mesh_uplink = meshUplinkSdkToTerraform(ctx, diags, d.MeshUplink)
	}
	if d.Model != nil {
		model = types.StringValue(*d.Model)
	}
	if d.ModifiedTime != nil {
		modified_time = types.Int64Value(int64(*d.ModifiedTime))
	}
	if d.Mount.Value() != nil {
		mount = types.StringValue(*d.Mount.Value())
	}
	if d.Name.Value() != nil {
		name = types.StringValue(*d.Name.Value())
	}
	if d.Notes.Value() != nil {
		notes = types.StringValue(*d.Notes.Value())
	}
	if d.NumClients.Value() != nil {
		num_clients = types.Int64Value(int64(*d.NumClients.Value()))
	}
	if d.OrgId.Value() != nil {
		org_id = types.StringValue(d.OrgId.Value().String())
	}
	if d.PortStat.Value() != nil && len(*d.PortStat.Value()) > 0 {
		port_stat = portStatdkToTerraform(ctx, diags, *d.PortStat.Value())
	}
	if d.PowerBudget.Value() != nil {
		power_budget = types.Int64Value(int64(*d.PowerBudget.Value()))
	}
	if d.PowerConstrained.Value() != nil {
		power_constrained = types.BoolValue(*d.PowerConstrained.Value())
	}
	if d.PowerOpmode.Value() != nil {
		power_opmode = types.StringValue(*d.PowerOpmode.Value())
	}
	if d.PowerSrc.Value() != nil {
		power_src = types.StringValue(*d.PowerSrc.Value())
	}
	if d.RadioStat != nil {
		radio_stat = radioStatSdkToTerraform(ctx, diags, d.RadioStat)
	}
	if d.RxBps.Value() != nil {
		rx_bps = types.NumberValue(big.NewFloat(*d.RxBps.Value()))
	}
	if d.RxBytes.Value() != nil {
		rx_bytes = types.Int64Value(int64(*d.RxBytes.Value()))
	}
	if d.RxPkts.Value() != nil {
		rx_pkts = types.Int64Value(int64(*d.RxPkts.Value()))
	}
	if d.Serial.Value() != nil {
		serial = types.StringValue(*d.Serial.Value())
	}
	if d.SiteId != nil {
		site_id = types.StringValue(d.SiteId.String())
	}
	if d.Status.Value() != nil {
		status = types.StringValue(*d.Status.Value())
	}
	if d.SwitchRedundancy != nil {
		switch_redundancy = SwitchRedundancySdkToTerraform(ctx, diags, d.SwitchRedundancy)
	}
	if d.TxBps.Value() != nil {
		tx_bps = types.NumberValue(big.NewFloat(*d.TxBps.Value()))
	}
	if d.TxBytes.Value() != nil {
		tx_bytes = types.NumberValue(big.NewFloat(*d.TxBytes.Value()))
	}
	if d.TxPkts.Value() != nil {
		tx_pkts = types.NumberValue(big.NewFloat(*d.TxPkts.Value()))
	}
	if d.Uptime.Value() != nil {
		uptime = types.NumberValue(big.NewFloat(*d.Uptime.Value()))
	}
	if d.UsbStat != nil {
		usb_stat = usbStatsSdkToTerraform(ctx, diags, d.UsbStat)
	}
	if d.Version.Value() != nil {
		version = types.StringValue(*d.Version.Value())
	}
	if d.X.Value() != nil {
		x = types.Float64Value(*d.X.Value())
	}
	if d.Y.Value() != nil {
		y = types.Float64Value(*d.Y.Value())
	}

	data_map_value := map[string]attr.Value{
		"auto_placement":       auto_placement,
		"auto_upgrade_stat":    auto_upgrade_stat,
		"ble_stat":             ble_stat,
		"cert_expiry":          cert_expiry,
		"config_reverted":      config_reverted,
		"cpu_system":           cpu_system,
		"cpu_util":             cpu_util,
		"created_time":         created_time,
		"deviceprofile_id":     deviceprofile_id,
		"env_stat":             env_stat,
		"esl_stat":             esl_stat,
		"ext_ip":               ext_ip,
		"fwupdate":             fwupdate,
		"hw_rev":               hw_rev,
		"id":                   id,
		"inactive_wired_vlans": inactive_wired_vlans,
		"iot_stat":             iot_stat,
		"ip":                   ip,
		"ip_config":            ip_config,
		"ip_stat":              ip_stat,
		"l2tp_stat":            l2tp_stat,
		"last_seen":            last_seen,
		"last_trouble":         last_trouble,
		"led":                  led,
		"lldp_stat":            lldp_stat,
		"locating":             locating,
		"locked":               locked,
		"mac":                  mac,
		"map_id":               map_id,
		"mem_used_kb":          mem_used_kb,
		"mesh_downlinks":       mesh_downlinks,
		"mesh_uplink":          mesh_uplink,
		"model":                model,
		"modified_time":        modified_time,
		"mount":                mount,
		"name":                 name,
		"notes":                notes,
		"num_clients":          num_clients,
		"org_id":               org_id,
		"port_stat":            port_stat,
		"power_budget":         power_budget,
		"power_constrained":    power_constrained,
		"power_opmode":         power_opmode,
		"power_src":            power_src,
		"radio_stat":           radio_stat,
		"rx_bps":               rx_bps,
		"rx_bytes":             rx_bytes,
		"rx_pkts":              rx_pkts,
		"serial":               serial,
		"site_id":              site_id,
		"status":               status,
		"switch_redundancy":    switch_redundancy,
		"tx_bps":               tx_bps,
		"tx_bytes":             tx_bytes,
		"tx_pkts":              tx_pkts,
		"uptime":               uptime,
		"usb_stat":             usb_stat,
		"version":              version,
		"x":                    x,
		"y":                    y,
	}
	data, e := NewDeviceApStatsValue(DeviceApStatsValue{}.AttributeTypes(ctx), data_map_value)
	diags.Append(e...)

	return data
}
