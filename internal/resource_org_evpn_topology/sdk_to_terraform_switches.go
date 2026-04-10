package resource_org_evpn_topology

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func switchesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.EvpnTopologySwitch) basetypes.MapValue {
	dataMap := make(map[string]SwitchesValue)
	for _, d := range l {
		var deviceprofileId basetypes.StringValue
		var downlinkIps = types.ListNull(types.StringType)
		var downlinks = types.ListNull(types.StringType)
		var esilaglinks = types.ListNull(types.StringType)
		var evpnId basetypes.Int64Value
		var mac basetypes.StringValue
		var model basetypes.StringValue
		var pod basetypes.Int64Value
		var pods = types.ListNull(types.Int64Type)
		var role basetypes.StringValue
		var routerId basetypes.StringValue
		var siteId basetypes.StringValue
		var suggestedDownlinks = types.ListNull(types.StringType)
		var suggestedEsilaglinks = types.ListNull(types.StringType)
		var suggestedUplinks = types.ListNull(types.StringType)
		var uplinks basetypes.ListValue

		if d.DeviceprofileId != nil {
			deviceprofileId = types.StringValue(d.DeviceprofileId.String())
		}
		if len(d.DownlinkIps) > 0 {
			downlinkIps = mistutils.ListOfStringSdkToTerraform(d.DownlinkIps)
		}
		if len(d.Downlinks) > 0 {
			downlinks = mistutils.ListOfStringSdkToTerraform(d.Downlinks)
		}
		if len(d.Esilaglinks) > 0 {
			esilaglinks = mistutils.ListOfStringSdkToTerraform(d.Esilaglinks)
		}
		if d.EvpnId != nil {
			evpnId = types.Int64Value(int64(*d.EvpnId))
		}

		mac = types.StringValue(d.Mac)

		if d.Model != nil {
			model = types.StringValue(*d.Model)
		}
		if d.Pod != nil {
			pod = types.Int64Value(int64(*d.Pod))
		}
		if len(d.Pods) > 0 {
			pods = mistutils.ListOfIntSdkToTerraform(d.Pods)
		}

		role = types.StringValue(string(d.Role))

		if d.RouterId != nil {
			routerId = types.StringValue(*d.RouterId)
		}
		if d.SiteId != nil {
			siteId = types.StringValue(d.SiteId.String())
		}
		if len(d.SuggestedDownlinks) > 0 {
			suggestedDownlinks = mistutils.ListOfStringSdkToTerraform(d.SuggestedDownlinks)
		}
		if len(d.SuggestedEsilaglinks) > 0 {
			suggestedEsilaglinks = mistutils.ListOfStringSdkToTerraform(d.SuggestedEsilaglinks)
		}
		if len(d.SuggestedUplinks) > 0 {
			suggestedUplinks = mistutils.ListOfStringSdkToTerraform(d.SuggestedUplinks)
		}
		if len(d.Uplinks) > 0 {
			uplinks = mistutils.ListOfStringSdkToTerraform(d.Uplinks)
		}

		dataMapValue := map[string]attr.Value{
			"deviceprofile_id":      deviceprofileId,
			"downlink_ips":          downlinkIps,
			"downlinks":             downlinks,
			"esilaglinks":           esilaglinks,
			"evpn_id":               evpnId,
			"mac":                   mac,
			"model":                 model,
			"pod":                   pod,
			"pods":                  pods,
			"role":                  role,
			"router_id":             routerId,
			"site_id":               siteId,
			"suggested_downlinks":   suggestedDownlinks,
			"suggested_esilaglinks": suggestedEsilaglinks,
			"suggested_uplinks":     suggestedUplinks,
			"uplinks":               uplinks,
		}
		data, e := NewSwitchesValue(SwitchesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataMap[data.Mac.ValueString()] = data
	}
	datalistType := SwitchesValue{}.Type(ctx)
	r, e := types.MapValueFrom(ctx, datalistType, dataMap)
	diags.Append(e...)
	return r
}
