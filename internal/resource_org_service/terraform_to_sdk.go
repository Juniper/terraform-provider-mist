package resource_org_service

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func specSpecsTerraformToSdk(d basetypes.ListValue) []models.ServiceSpec {
	var data []models.ServiceSpec
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		vState := vInterface.(SpecsValue)
		vData := models.ServiceSpec{}

		if vState.PortRange.ValueStringPointer() != nil {
			vData.PortRange = vState.PortRange.ValueStringPointer()
		}

		if vState.Protocol.ValueStringPointer() != nil {
			vData.Protocol = vState.Protocol.ValueStringPointer()
		}
		data = append(data, vData)
	}
	return data
}

func TerraformToSdk(plan *OrgServiceModel) (models.Service, diag.Diagnostics) {
	var diags diag.Diagnostics
	unset := make(map[string]interface{})
	data := models.Service{}

	data.Name = plan.Name.ValueStringPointer()
	data.Specs = specSpecsTerraformToSdk(plan.Specs)

	if !plan.Addresses.IsNull() && !plan.Addresses.IsUnknown() {
		data.Addresses = misttransform.ListOfStringTerraformToSdk(plan.Addresses)
	} else {
		unset["-addresses"] = ""
	}
	if !plan.AppCategories.IsNull() && !plan.AppCategories.IsUnknown() {
		data.AppCategories = misttransform.ListOfStringTerraformToSdk(plan.AppCategories)
	} else {
		unset["-app_categories"] = ""
	}
	if !plan.AppSubcategories.IsNull() && !plan.AppSubcategories.IsUnknown() {
		data.AppSubcategories = misttransform.ListOfStringTerraformToSdk(plan.AppSubcategories)
	} else {
		unset["-app_subcategories"] = ""
	}
	if !plan.Apps.IsNull() && !plan.Apps.IsUnknown() {
		data.Apps = misttransform.ListOfStringTerraformToSdk(plan.Apps)
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
		data.Hostnames = misttransform.ListOfStringTerraformToSdk(plan.Hostnames)
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
		data.Urls = misttransform.ListOfStringTerraformToSdk(plan.Urls)
	} else {
		unset["-urls"] = ""
	}

	data.AdditionalProperties = unset
	return data, diags
}
