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
	unset := make(map[string]interface{})
	data := models.Service{}

	data.Name = plan.Name.ValueStringPointer()
	data.Specs = specSpecsTerraformToSdk(ctx, &diags, plan.Specs)

	if !plan.Addresses.IsNull() && !plan.Addresses.IsUnknown() {
		data.Addresses = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Addresses)
	} else {
		unset["-addresses"] = ""
	}
	if !plan.AppCategories.IsNull() && !plan.AppCategories.IsUnknown() {
		data.AppCategories = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AppCategories)
	} else {
		unset["-app_categories"] = ""
	}
	if !plan.AppSubcategories.IsNull() && !plan.AppSubcategories.IsUnknown() {
		data.AppSubcategories = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AppSubcategories)
	} else {
		unset["-app_subcategories"] = ""
	}
	if !plan.Apps.IsNull() && !plan.Apps.IsUnknown() {
		data.Apps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Apps)
	} else {
		unset["-apps"] = ""
	}
	if plan.Description.ValueStringPointer() != nil {
		data.Description = plan.Description.ValueStringPointer()
	} else {
		unset["-descritpion"] = ""
	}
	if plan.Dscp.ValueStringPointer() != nil {
		data.Dscp = models.ToPointer(models.ServiceDscpContainer.FromString(plan.Dscp.ValueString()))
	} else {
		unset["-dscp"] = ""
	}
	if plan.FailoverPolicy.ValueStringPointer() != nil {
		data.FailoverPolicy = models.ToPointer(models.ServiceFailoverPolicyEnum(plan.FailoverPolicy.ValueString()))
	} else {
		unset["-failover_policy"] = ""
	}
	if !plan.Hostnames.IsNull() && !plan.Hostnames.IsUnknown() {
		data.Hostnames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Hostnames)
	} else {
		unset["-hostnames"] = ""
	}
	if plan.MaxJitter.ValueStringPointer() != nil {
		data.MaxJitter = models.ToPointer(models.ServiceMaxJitterContainer.FromString(plan.MaxJitter.ValueString()))
	} else {
		unset["-max_jitter"] = ""
	}
	if plan.MaxLatency.ValueStringPointer() != nil {
		data.MaxLatency = models.ToPointer(models.ServiceMaxLatencyContainer.FromString(plan.MaxLatency.ValueString()))
	} else {
		unset["-max_latency"] = ""
	}
	if plan.MaxLoss.ValueStringPointer() != nil {
		data.MaxLoss = models.ToPointer(models.ServiceMaxLossContainer.FromString(plan.MaxLoss.ValueString()))
	} else {
		unset["-max_loss"] = ""
	}
	if plan.SleEnabled.ValueBoolPointer() != nil {
		data.SleEnabled = plan.SleEnabled.ValueBoolPointer()
	} else {
		unset["-sle_enables"] = ""
	}
	if plan.SsrRelaxedTcpStateEnforcement.ValueBoolPointer() != nil {
		data.SsrRelaxedTcpStateEnforcement = plan.SsrRelaxedTcpStateEnforcement.ValueBoolPointer()
	} else {
		unset["-ssr_relaxed_tcp_state_enforcement"] = ""
	}
	if plan.TrafficClass.ValueStringPointer() != nil {
		data.TrafficClass = models.ToPointer(models.ServiceTrafficClassEnum(plan.TrafficClass.ValueString()))
	} else {
		unset["-traffic_class"] = ""
	}
	if plan.TrafficType.ValueStringPointer() != nil {
		data.TrafficType = plan.TrafficType.ValueStringPointer()
	} else {
		unset["-traffic_type"] = ""
	}
	if plan.Type.ValueStringPointer() != nil {
		data.Type = (*models.ServiceTypeEnum)(plan.Type.ValueStringPointer())
	} else {
		unset["-type"] = ""
	}
	if !plan.Urls.IsNull() && !plan.Urls.IsUnknown() {
		data.Urls = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Urls)
	} else {
		unset["-urls"] = ""
	}

	data.AdditionalProperties = unset
	return data, diags
}
