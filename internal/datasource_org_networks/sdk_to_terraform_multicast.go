package datasource_org_networks

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func groupMulticastSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.NetworkMulticastGroup) basetypes.MapValue {
	stateValueMapValue := make(map[string]attr.Value)
	for k, v := range d {
		var rpIp basetypes.StringValue

		if v.RpIp != nil {
			rpIp = types.StringValue(*v.RpIp)
		}

		dataMapValue := map[string]attr.Value{
			"rp_ip": rpIp,
		}
		n, e := NewGroupsValue(GroupsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMapValue[k] = n
	}
	stateResultMap, e := types.MapValueFrom(ctx, GroupsValue{}.Type(ctx), stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}

func MulticastSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.NetworkMulticast) basetypes.ObjectValue {
	var disableIgmp basetypes.BoolValue
	var enabled basetypes.BoolValue
	var groups = types.MapNull(GroupsValue{}.Type(ctx))

	if d.DisableIgmp != nil {
		disableIgmp = types.BoolValue(*d.DisableIgmp)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Groups != nil {
		groups = groupMulticastSdkToTerraform(ctx, diags, d.Groups)
	}

	dataMapValue := map[string]attr.Value{
		"disable_igmp": disableIgmp,
		"enabled":      enabled,
		"groups":       groups,
	}
	data, e := basetypes.NewObjectValueFrom(ctx, MulticastValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
