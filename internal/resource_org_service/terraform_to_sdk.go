package resource_org_service

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func specSpecsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.ServiceSpec {
	var data []models.ServiceSpec
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		v_state := v_interface.(SpecsValue)
		v_data := models.ServiceSpec{}

		if v_state.PortRange.ValueStringPointer() != nil {
			v_data.PortRange = v_state.PortRange.ValueStringPointer()
		}

		if v_state.Protocol.ValueStringPointer() != nil {
			v_data.Protocol = v_state.Protocol.ValueStringPointer()
		}
		data = append(data, v_data)
	}
	return data
}

func TerraformToSdk(ctx context.Context, plan *OrgServiceModel) (models.Service, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.Service{}

	data.Name = plan.Name.ValueStringPointer()
	data.Specs = specSpecsTerraformToSdk(ctx, &diags, plan.Specs)

	if !plan.Addresses.IsNull() && !plan.Addresses.IsUnknown() {
		data.Addresses = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Addresses)
	}
	if !plan.AppCategories.IsNull() && !plan.AppCategories.IsUnknown() {
		data.AppCategories = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AppCategories)
	}
	if !plan.AppSubcategories.IsNull() && !plan.AppSubcategories.IsUnknown() {
		data.AppSubcategories = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AppSubcategories)
	}
	if !plan.Apps.IsNull() && !plan.Apps.IsUnknown() {
		data.Apps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Apps)
	}
	if plan.Description.ValueStringPointer() != nil {
		data.Description = plan.Description.ValueStringPointer()
	}
	if plan.Dscp.ValueStringPointer() != nil {
		data.Dscp = models.ToPointer(models.ServiceDscpContainer.FromString(plan.Dscp.ValueString()))
	}
	if plan.FailoverPolicy.ValueStringPointer() != nil {
		data.FailoverPolicy = models.ToPointer(models.ServiceFailoverPolicyEnum(plan.FailoverPolicy.ValueString()))
	}
	if !plan.Hostnames.IsNull() && !plan.Hostnames.IsUnknown() {
		data.Hostnames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Hostnames)
	}
	if plan.MaxJitter.ValueStringPointer() != nil {
		data.MaxJitter = models.ToPointer(models.ServiceMaxJitterContainer.FromString(plan.MaxJitter.ValueString()))
	}
	if plan.MaxLatency.ValueStringPointer() != nil {
		data.MaxLatency = models.ToPointer(models.ServiceMaxLatencyContainer.FromString(plan.MaxLatency.ValueString()))
	}
	if plan.MaxLoss.ValueStringPointer() != nil {
		data.MaxLoss = models.ToPointer(models.ServiceMaxLossContainer.FromString(plan.MaxLoss.ValueString()))
	}
	if plan.SleEnabled.ValueBoolPointer() != nil {
		data.SleEnabled = plan.SleEnabled.ValueBoolPointer()
	}
	if plan.SsrRelaxedTcpStateEnforcement.ValueBoolPointer() != nil {
		data.SsrRelaxedTcpStateEnforcement = plan.SsrRelaxedTcpStateEnforcement.ValueBoolPointer()
	}
	if plan.TrafficClass.ValueStringPointer() != nil {
		data.TrafficClass = models.ToPointer(models.ServiceTrafficClassEnum(plan.TrafficClass.ValueString()))
	}
	if plan.TrafficType.ValueStringPointer() != nil {
		data.TrafficType = plan.TrafficType.ValueStringPointer()
	}
	if plan.Type.ValueStringPointer() != nil {
		data.Type = (*models.ServiceTypeEnum)(plan.Type.ValueStringPointer())
	}
	if !plan.Urls.IsNull() && !plan.Urls.IsUnknown() {
		data.Urls = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Urls)
	}

	return data, diags
}
