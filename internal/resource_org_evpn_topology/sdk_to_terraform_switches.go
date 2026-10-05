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
		var pod basetypes.Int64Value
		var pods = types.ListNull(types.Int64Type)
		var role basetypes.StringValue

		if d.Pod != nil {
			pod = types.Int64Value(int64(*d.Pod))
		}
		if d.Pods != nil {
			pods = mistutils.ListOfIntSdkToTerraform(d.Pods)
		}

		role = types.StringValue(string(d.Role))

		dataMapValue := map[string]attr.Value{
			"pod":  pod,
			"pods": pods,
			"role": role,
		}
		data, e := NewSwitchesValue(SwitchesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataMap[d.Mac] = data
	}
	datalistType := SwitchesValue{}.Type(ctx)
	r, e := types.MapValueFrom(ctx, datalistType, dataMap)
	diags.Append(e...)
	return r
}
