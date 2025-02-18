package resource_org_networktemplate

import (
	"context"

	mistapi "github.com/Juniper/terraform-provider-mist/internal/commons/api_response"
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
		vlanId = mistapi.VlanAsString(d.VlanId)

		dataMapValue := map[string]attr.Value{
			"isolation":         isolation,
			"isolation_vlan_id": isolationVlanId,
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
