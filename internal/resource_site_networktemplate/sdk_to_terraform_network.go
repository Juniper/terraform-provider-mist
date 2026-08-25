package resource_site_networktemplate

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func NetworksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.SwitchNetwork) basetypes.MapValue {

	stateValueMapValue := make(map[string]attr.Value)
	for k, d := range m {

		var isolation basetypes.BoolValue
		var isolationVlanId basetypes.StringValue
		var multicast = types.ObjectNull(MulticastValue{}.AttributeTypes(ctx))
		var gateway basetypes.StringValue
		var gateway6 basetypes.StringValue
		var subnet basetypes.StringValue
		var subnet6 basetypes.StringValue
		var vlanId basetypes.StringValue

		if d.Isolation != nil {
			isolation = types.BoolValue(*d.Isolation)
		}
		if d.IsolationVlanId != nil {
			isolationVlanId = types.StringValue(*d.IsolationVlanId)
		}
		if d.Multicast != nil {
			var enabled basetypes.BoolValue
			var igmpVersion basetypes.StringValue
			if d.Multicast.Enabled != nil {
				enabled = types.BoolValue(*d.Multicast.Enabled)
			}
			if d.Multicast.IgmpVersion != nil {
				igmpVersion = types.StringValue(string(*d.Multicast.IgmpVersion))
			}
			mv, e := NewMulticastValue(MulticastValue{}.AttributeTypes(ctx), map[string]attr.Value{
				"enabled":      enabled,
				"igmp_version": igmpVersion,
			})
			diags.Append(e...)
			o, e2 := mv.ToObjectValue(ctx)
			diags.Append(e2...)
			multicast = o
		}
		if d.Gateway != nil {
			gateway = types.StringValue(*d.Gateway)
		}
		if d.Gateway6 != nil {
			gateway6 = types.StringValue(*d.Gateway6)
		}
		if d.Subnet != nil {
			subnet = types.StringValue(*d.Subnet)
		}
		if d.Subnet6 != nil {
			subnet6 = types.StringValue(*d.Subnet6)
		}
		vlanId = mistutils.VlanAsString(d.VlanId)

		dataMapValue := map[string]attr.Value{
			"isolation":         isolation,
			"isolation_vlan_id": isolationVlanId,
			"multicast":         multicast,
			"gateway":           gateway,
			"gateway6":          gateway6,
			"subnet":            subnet,
			"subnet6":           subnet6,
			"vlan_id":           vlanId,
		}
		data, e := NewNetworksValue(NetworksValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMapValue[k] = data
	}
	stateResultMapType := NetworksValue{}.Type(ctx)
	stateResultMap, e := types.MapValueFrom(ctx, stateResultMapType, stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}
