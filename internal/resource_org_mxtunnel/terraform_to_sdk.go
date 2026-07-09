package resource_org_mxtunnel

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgMxtunnelModel) (*models.Mxtunnel, diag.Diagnostics) {
	var data models.Mxtunnel
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	if !plan.AnchorMxtunnelIds.IsNull() && !plan.AnchorMxtunnelIds.IsUnknown() {
		strs := mistutils.ListOfStringTerraformToSdk(plan.AnchorMxtunnelIds)
		var uuids []uuid.UUID
		for _, s := range strs {
			u, e := uuid.Parse(s)
			if e != nil {
				diags.AddError("Invalid anchor_mxtunnel_ids value", e.Error())
				continue
			}
			uuids = append(uuids, u)
		}
		data.AnchorMxtunnelIds = uuids
	} else {
		unset["-anchor_mxtunnel_ids"] = ""
	}

	if !plan.AutoPreemption.IsNull() && !plan.AutoPreemption.IsUnknown() {
		data.AutoPreemption = autoPreemptionTerraformToSdk(ctx, &diags, plan.AutoPreemption)
	} else {
		unset["-auto_preemption"] = ""
	}

	if !plan.HelloInterval.IsNull() && !plan.HelloInterval.IsUnknown() {
		data.HelloInterval = models.NewOptional(models.ToPointer(int(plan.HelloInterval.ValueInt64())))
	} else {
		data.HelloInterval = models.NewOptional[int](nil)
	}

	if !plan.HelloRetries.IsNull() && !plan.HelloRetries.IsUnknown() {
		data.HelloRetries = models.NewOptional(models.ToPointer(int(plan.HelloRetries.ValueInt64())))
	} else {
		data.HelloRetries = models.NewOptional[int](nil)
	}

	if !plan.Ipsec.IsNull() && !plan.Ipsec.IsUnknown() {
		data.Ipsec = ipsecTerraformToSdk(ctx, &diags, plan.Ipsec)
	} else {
		unset["-ipsec"] = ""
	}

	if !plan.Mtu.IsNull() && !plan.Mtu.IsUnknown() {
		data.Mtu = models.ToPointer(int(plan.Mtu.ValueInt64()))
	} else {
		unset["-mtu"] = ""
	}

	if !plan.MxclusterIds.IsNull() && !plan.MxclusterIds.IsUnknown() {
		strs := mistutils.ListOfStringTerraformToSdk(plan.MxclusterIds)
		var uuids []uuid.UUID
		for _, s := range strs {
			u, e := uuid.Parse(s)
			if e != nil {
				diags.AddError("Invalid mxcluster_ids value", e.Error())
				continue
			}
			uuids = append(uuids, u)
		}
		data.MxclusterIds = uuids
	} else {
		unset["-mxcluster_ids"] = ""
	}

	data.Name = models.NewOptional(plan.Name.ValueStringPointer())

	// org_id is required, always parse it
	orgId, e := uuid.Parse(plan.OrgId.ValueString())
	if e == nil {
		data.OrgId = &orgId
	} else {
		diags.AddError("Invalid value for org_id", e.Error())
	}

	if !plan.Protocol.IsNull() && !plan.Protocol.IsUnknown() && plan.Protocol.ValueString() != "" {
		data.Protocol = (*models.MxtunnelProtocolEnum)(plan.Protocol.ValueStringPointer())
	} else {
		unset["-protocol"] = ""
	}

	if !plan.VlanIds.IsNull() && !plan.VlanIds.IsUnknown() {
		data.VlanIds = mistutils.ListOfIntTerraformToSdk(plan.VlanIds)
	} else {
		unset["-vlan_ids"] = ""
	}

	data.AdditionalProperties = unset

	return &data, diags
}
