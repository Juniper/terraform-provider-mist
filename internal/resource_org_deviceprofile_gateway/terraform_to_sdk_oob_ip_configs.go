package resource_org_deviceprofile_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func oobIpConfigNode1TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.GatewayOobIpConfigNode1 {
	data := models.GatewayOobIpConfigNode1{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewNode1Value(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.Ip.ValueStringPointer() != nil {
				data.Ip = plan.Ip.ValueStringPointer()
			}
			if plan.Netmask.ValueStringPointer() != nil {
				data.Netmask = plan.Netmask.ValueStringPointer()
			}
			if plan.Network.ValueStringPointer() != nil {
				data.Network = plan.Network.ValueStringPointer()
			}

			if plan.Node1Type.ValueStringPointer() != nil {
				data.Type = (*models.IpTypeEnum)(plan.Node1Type.ValueStringPointer())
			}
			if plan.UseMgmtVrf.ValueBoolPointer() != nil {
				data.UseMgmtVrf = plan.UseMgmtVrf.ValueBoolPointer()
			}
			if plan.UseMgmtVrfForHostOut.ValueBoolPointer() != nil {
				data.UseMgmtVrfForHostOut = plan.UseMgmtVrfForHostOut.ValueBoolPointer()
			}
		}
	}
	return &data
}
func oobIpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d OobIpConfigValue) *models.GatewayOobIpConfig {

	data := models.GatewayOobIpConfig{}

	if d.Ip.ValueStringPointer() != nil {
		data.Ip = d.Ip.ValueStringPointer()
	}
	if d.Netmask.ValueStringPointer() != nil {
		data.Netmask = d.Netmask.ValueStringPointer()
	}
	if d.Network.ValueStringPointer() != nil {
		data.Network = d.Network.ValueStringPointer()
	}
	if !d.Node1.IsNull() && !d.Node1.IsUnknown() {
		data.Node1 = oobIpConfigNode1TerraformToSdk(ctx, diags, d.Node1)
	}
	if d.OobIpConfigType.ValueStringPointer() != nil {
		data.Type = (*models.IpTypeEnum)(d.OobIpConfigType.ValueStringPointer())
	}
	if d.UseMgmtVrf.ValueBoolPointer() != nil {
		data.UseMgmtVrf = d.UseMgmtVrf.ValueBoolPointer()
	}
	if d.UseMgmtVrfForHostOut.ValueBoolPointer() != nil {
		data.UseMgmtVrfForHostOut = d.UseMgmtVrfForHostOut.ValueBoolPointer()
	}

	return &data
}
