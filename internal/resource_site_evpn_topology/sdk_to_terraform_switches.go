package resource_site_evpn_topology

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func switchesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.EvpnTopologySwitch) basetypes.MapValue {
	data_map := make(map[string]SwitchesValue)
	for _, d := range l {
		var deviceprofile_id basetypes.StringValue
		// var downlink_ips basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		// var downlinks basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		// var esilaglinks basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		var evpn_id basetypes.Int64Value
		var mac basetypes.StringValue
		var model basetypes.StringValue
		var pod basetypes.Int64Value = types.Int64Value(1)
		var pods basetypes.ListValue = mist_transform.ListOfIntSdkToTerraformEmpty(ctx)
		var role basetypes.StringValue
		var router_id basetypes.StringValue
		var site_id basetypes.StringValue
		// var suggested_downlinks basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		// var suggested_esilaglinks basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		// var suggested_uplinks basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		// var uplinks basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)

		if d.DeviceprofileId != nil {
			deviceprofile_id = types.StringValue(d.DeviceprofileId.String())
		}
		// if d.DownlinkIps != nil {
		// 	downlink_ips = mist_transform.ListOfStringSdkToTerraform(ctx, d.DownlinkIps)
		// }
		// if d.Downlinks != nil {
		// 	downlinks = mist_transform.ListOfStringSdkToTerraform(ctx, d.Downlinks)
		// }
		// if d.Esilaglinks != nil {
		// 	esilaglinks = mist_transform.ListOfStringSdkToTerraform(ctx, d.Esilaglinks)
		// }
		if d.EvpnId != nil {
			evpn_id = types.Int64Value(int64(*d.EvpnId))
		}

		mac = types.StringValue(d.Mac)

		if d.Model != nil {
			model = types.StringValue(*d.Model)
		}
		if d.Pod != nil {
			pod = types.Int64Value(int64(*d.Pod))
		}
		if d.Pods != nil {
			pods = mist_transform.ListOfIntSdkToTerraform(ctx, d.Pods)
		}

		role = types.StringValue(string(d.Role))

		if d.RouterId != nil {
			router_id = types.StringValue(*d.RouterId)
		}
		if d.SiteId != nil {
			site_id = types.StringValue(d.SiteId.String())
		}
		// if d.SuggestedDownlinks != nil {
		// 	suggested_downlinks = mist_transform.ListOfStringSdkToTerraform(ctx, d.SuggestedDownlinks)
		// }
		// if d.SuggestedEsilaglinks != nil {
		// 	suggested_esilaglinks = mist_transform.ListOfStringSdkToTerraform(ctx, d.SuggestedEsilaglinks)
		// }
		// if d.SuggestedUplinks != nil {
		// 	suggested_uplinks = mist_transform.ListOfStringSdkToTerraform(ctx, d.SuggestedUplinks)
		// }
		// if d.Uplinks != nil {
		// 	uplinks = mist_transform.ListOfStringSdkToTerraform(ctx, d.Uplinks)
		// }

		data_map_attr_type := SwitchesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"deviceprofile_id": deviceprofile_id,
			// "downlink_ips":     downlink_ips,
			// "downlinks":             downlinks,
			// "esilaglinks":           esilaglinks,
			"evpn_id":   evpn_id,
			"mac":       mac,
			"model":     model,
			"pod":       pod,
			"pods":      pods,
			"role":      role,
			"router_id": router_id,
			"site_id":   site_id,
			// "suggested_downlinks":   suggested_downlinks,
			// "suggested_esilaglinks": suggested_esilaglinks,
			// "suggested_uplinks":     suggested_uplinks,
			// "uplinks":               uplinks,
		}
		data, e := NewSwitchesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_map[mac.ValueString()] = data
	}
	data_list_type := SwitchesValue{}.Type(ctx)
	r, e := types.MapValueFrom(ctx, data_list_type, data_map)
	diags.Append(e...)
	return r
}
