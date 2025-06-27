package resource_site_evpn_topology

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
		var evpnId basetypes.Int64Value
		var mac basetypes.StringValue
		var model basetypes.StringValue
		var pod = types.Int64Value(1)
		var pods = mistutils.ListOfIntSdkToTerraformEmpty()
		var role basetypes.StringValue
		var routerId basetypes.StringValue
		var siteId basetypes.StringValue

		if d.DeviceprofileId != nil {
			deviceprofileId = types.StringValue(d.DeviceprofileId.String())
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
		if d.Pods != nil {
			pods = mistutils.ListOfIntSdkToTerraform(d.Pods)
		}

		role = types.StringValue(string(d.Role))

		if d.RouterId != nil {
			routerId = types.StringValue(*d.RouterId)
		}
		if d.SiteId != nil {
			siteId = types.StringValue(d.SiteId.String())
		}

		dataMapValue := map[string]attr.Value{
			"deviceprofile_id": deviceprofileId,
			"evpn_id":          evpnId,
			"mac":              mac,
			"model":            model,
			"pod":              pod,
			"pods":             pods,
			"role":             role,
			"router_id":        routerId,
			"site_id":          siteId,
		}
		data, e := NewSwitchesValue(SwitchesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataMap[mac.ValueString()] = data
	}
	datalistType := SwitchesValue{}.Type(ctx)
	r, e := types.MapValueFrom(ctx, datalistType, dataMap)
	diags.Append(e...)
	return r
}
