package resource_org_mxtunnel

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *models.Mxtunnel) (OrgMxtunnelModel, diag.Diagnostics) {
	var state OrgMxtunnelModel
	var diags diag.Diagnostics

	var anchorMxtunnelIds = types.ListNull(types.StringType)
	var autoPreemption = NewAutoPreemptionValueNull()
	var helloInterval = types.Int64Null()
	var helloRetries = types.Int64Null()
	var id = types.StringNull()
	var ipsec = NewIpsecValueNull()
	var mtu = types.Int64Null()
	var mxclusterIds = types.ListNull(types.StringType)
	var name = types.StringNull()
	var orgId = types.StringNull()
	var protocol = types.StringNull()
	var vlanIds = types.ListNull(types.Int64Type)

	if data.AnchorMxtunnelIds != nil {
		var items []string
		for _, v := range data.AnchorMxtunnelIds {
			items = append(items, v.String())
		}
		anchorMxtunnelIds = mistutils.ListOfStringSdkToTerraform(items)
	}

	if data.AutoPreemption != nil {
		autoPreemption = autoPreemptionSdkToTerraform(ctx, &diags, data.AutoPreemption)
	}

	if data.HelloInterval.Value() != nil {
		helloInterval = types.Int64Value(int64(*data.HelloInterval.Value()))
	}

	if data.HelloRetries.Value() != nil {
		helloRetries = types.Int64Value(int64(*data.HelloRetries.Value()))
	}

	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}

	if data.Ipsec != nil {
		ipsec = ipsecSdkToTerraform(ctx, &diags, data.Ipsec)
	}

	if data.Mtu != nil {
		mtu = types.Int64Value(int64(*data.Mtu))
	}

	if data.MxclusterIds != nil {
		var items []string
		for _, v := range data.MxclusterIds {
			items = append(items, v.String())
		}
		mxclusterIds = mistutils.ListOfStringSdkToTerraform(items)
	}

	if data.Name.Value() != nil {
		name = types.StringValue(*data.Name.Value())
	}

	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}

	if data.Protocol != nil {
		protocol = types.StringValue(string(*data.Protocol))
	}

	if data.VlanIds != nil {
		vlanIds = mistutils.ListOfIntSdkToTerraform(data.VlanIds)
	}

	state.AnchorMxtunnelIds = anchorMxtunnelIds
	state.AutoPreemption = autoPreemption
	state.HelloInterval = helloInterval
	state.HelloRetries = helloRetries
	state.Id = id
	state.Ipsec = ipsec
	state.Mtu = mtu
	state.MxclusterIds = mxclusterIds
	state.Name = name
	state.OrgId = orgId
	state.Protocol = protocol
	state.VlanIds = vlanIds

	return state, diags
}
