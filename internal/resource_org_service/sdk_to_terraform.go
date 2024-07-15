package resource_org_service

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func serviceSpecsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ServiceSpec) basetypes.ListValue {

	var data_list = []SpecsValue{}

	for _, d := range l {
		var port_range basetypes.StringValue
		var protocol basetypes.StringValue

		if d.PortRange != nil {
			port_range = types.StringValue(*d.PortRange)
		}
		if d.Protocol != nil {
			protocol = types.StringValue(*d.Protocol)
		}

		data_map_attr_type := SpecsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"port_range": port_range,
			"protocol":   protocol,
		}
		data, e := NewSpecsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := SpecsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func SdkToTerraform(ctx context.Context, d *models.Service) (OrgServiceModel, diag.Diagnostics) {
	var state OrgServiceModel
	var diags diag.Diagnostics

	var addresses types.List = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var app_categories types.List = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var app_subcategories types.List = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var apps types.List = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var description types.String
	var dscp types.Int64
	var failover_policy types.String
	var hostnames types.List = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var id types.String
	var max_jitter types.Int64
	var max_latency types.Int64
	var max_loss types.Int64
	var name types.String
	var org_id types.String
	var sle_enabled types.Bool
	var specs types.List = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var ssr_relaxed_tcp_state_enforcement types.Bool
	var traffic_class types.String
	var traffic_type types.String
	var service_type types.String
	var urls types.List = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)

	if d.Addresses != nil {
		addresses = mist_transform.ListOfStringSdkToTerraform(ctx, d.Addresses)
	}
	if d.AppCategories != nil {
		app_categories = mist_transform.ListOfStringSdkToTerraform(ctx, d.AppCategories)
	}
	if d.AppSubcategories != nil {
		app_subcategories = mist_transform.ListOfStringSdkToTerraform(ctx, d.AppSubcategories)
	}
	if d.Apps != nil {
		apps = mist_transform.ListOfStringSdkToTerraform(ctx, d.Apps)
	}
	if d.Description != nil {
		description = types.StringValue(string(*d.Description))
	}
	if d.Dscp != nil {
		dscp = types.Int64Value(int64(*d.Dscp))
	}
	if d.FailoverPolicy != nil {
		failover_policy = types.StringValue(string(*d.FailoverPolicy))
	}
	if d.Hostnames != nil {
		hostnames = mist_transform.ListOfStringSdkToTerraform(ctx, d.Hostnames)
	}
	if d.Id != nil {
		id = types.StringValue(string(d.Id.String()))
	}
	if d.MaxJitter != nil {
		max_jitter = types.Int64Value(int64(*d.MaxJitter))
	}
	if d.MaxLatency != nil {
		max_latency = types.Int64Value(int64(*d.MaxLatency))
	}
	if d.MaxLoss != nil {
		max_loss = types.Int64Value(int64(*d.MaxLoss))
	}
	if d.Name != nil {
		name = types.StringValue(string(*d.Name))
	}
	if d.OrgId != nil {
		org_id = types.StringValue(d.OrgId.String())
	}
	if d.SleEnabled != nil {
		sle_enabled = types.BoolValue(*d.SleEnabled)
	}
	if d.Specs != nil {
		specs = serviceSpecsSdkToTerraform(ctx, &diags, d.Specs)
	}
	if d.SsrRelaxedTcpStateEnforcement != nil {
		ssr_relaxed_tcp_state_enforcement = types.BoolValue(*d.SsrRelaxedTcpStateEnforcement)
	}
	if d.TrafficClass != nil {
		traffic_class = types.StringValue(string(*d.TrafficClass))
	}
	if d.TrafficType != nil {
		traffic_type = types.StringValue(string(*d.TrafficType))
	}
	if d.Type != nil {
		service_type = types.StringValue(string(*d.Type))
	}
	if d.Urls != nil {
		urls = mist_transform.ListOfStringSdkToTerraform(ctx, d.Urls)
	}

	state.Addresses = addresses
	state.AppCategories = app_categories
	state.AppSubcategories = app_subcategories
	state.Apps = apps
	state.Description = description
	state.Dscp = dscp
	state.FailoverPolicy = failover_policy
	state.Hostnames = hostnames
	state.Id = id
	state.MaxJitter = max_jitter
	state.MaxLatency = max_latency
	state.MaxLoss = max_loss
	state.Name = name
	state.OrgId = org_id
	state.SleEnabled = sle_enabled
	state.Specs = specs
	state.SsrRelaxedTcpStateEnforcement = ssr_relaxed_tcp_state_enforcement
	state.TrafficClass = traffic_class
	state.TrafficType = traffic_type
	state.Type = service_type
	state.Urls = urls

	return state, diags

}
