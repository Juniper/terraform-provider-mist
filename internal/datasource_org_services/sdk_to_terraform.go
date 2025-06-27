package datasource_org_services

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.Service, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := serviceSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func serviceSpecsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ServiceSpec) basetypes.ListValue {

	var dataList []SpecsValue

	for _, d := range l {
		var portRange basetypes.StringValue
		var protocol basetypes.StringValue

		if d.PortRange != nil {
			portRange = types.StringValue(*d.PortRange)
		}
		if d.Protocol != nil {
			protocol = types.StringValue(*d.Protocol)
		}

		dataMapValue := map[string]attr.Value{
			"port_range": portRange,
			"protocol":   protocol,
		}
		data, e := NewSpecsValue(SpecsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, SpecsValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func serviceSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Service) OrgServicesValue {

	var addresses = types.ListNull(types.StringType)
	var appCategories = types.ListNull(types.StringType)
	var appSubcategories = types.ListNull(types.StringType)
	var apps = types.ListNull(types.StringType)
	var clientLimitDown basetypes.Int64Value
	var clientLimitUp basetypes.Int64Value
	var createdTime basetypes.Float64Value
	var description basetypes.StringValue
	var dscp basetypes.StringValue
	var failoverPolicy basetypes.StringValue
	var hostnames = types.ListNull(types.StringType)
	var id basetypes.StringValue
	var maxJitter basetypes.StringValue
	var maxLatency basetypes.StringValue
	var maxLoss basetypes.StringValue
	var modifiedTime basetypes.Float64Value
	var name basetypes.StringValue
	var orgId basetypes.StringValue
	var serviceLimitDown types.Int64
	var serviceLimitUp types.Int64
	var sleEnabled types.Bool
	var specs = types.ListNull(SpecsValue{}.Type(ctx))
	var ssrRelaxedTcpStateEnforcement types.Bool
	var trafficClass types.String
	var trafficType types.String
	var serviceType types.String
	var urls = mistutils.ListOfStringSdkToTerraformEmpty()

	if d.Addresses != nil {
		addresses = mistutils.ListOfStringSdkToTerraform(d.Addresses)
	}
	if d.AppCategories != nil {
		appCategories = mistutils.ListOfStringSdkToTerraform(d.AppCategories)
	}
	if d.AppSubcategories != nil {
		appSubcategories = mistutils.ListOfStringSdkToTerraform(d.AppSubcategories)
	}
	if d.Apps != nil {
		addresses = mistutils.ListOfStringSdkToTerraform(d.Apps)
	}
	if d.ClientLimitDown != nil {
		clientLimitDown = types.Int64Value(int64(*d.ClientLimitDown))
	}
	if d.ClientLimitUp != nil {
		clientLimitUp = types.Int64Value(int64(*d.ClientLimitUp))
	}
	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.Description != nil {
		description = types.StringValue(*d.Description)
	}
	if d.Dscp != nil {
		dscp = types.StringValue(d.Dscp.String())
	}
	if d.FailoverPolicy != nil {
		failoverPolicy = types.StringValue(string(*d.FailoverPolicy))
	}
	if d.Hostnames != nil {
		hostnames = mistutils.ListOfStringSdkToTerraform(d.Hostnames)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.MaxJitter != nil {
		maxJitter = types.StringValue(d.MaxJitter.String())
	}
	if d.MaxLatency != nil {
		maxLatency = types.StringValue(d.MaxLatency.String())
	}
	if d.MaxLoss != nil {
		maxLoss = types.StringValue(d.MaxLoss.String())
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.ServiceLimitDown != nil {
		serviceLimitDown = types.Int64Value(int64(*d.ServiceLimitDown))
	}
	if d.ServiceLimitUp != nil {
		serviceLimitUp = types.Int64Value(int64(*d.ServiceLimitUp))
	}
	if d.SleEnabled != nil {
		sleEnabled = types.BoolValue(*d.SleEnabled)
	}
	if d.Specs != nil {
		specs = serviceSpecsSdkToTerraform(ctx, diags, d.Specs)
	}
	if d.SsrRelaxedTcpStateEnforcement != nil {
		ssrRelaxedTcpStateEnforcement = types.BoolValue(*d.SsrRelaxedTcpStateEnforcement)
	}
	if d.TrafficClass != nil {
		trafficClass = types.StringValue(string(*d.TrafficClass))
	}
	if d.TrafficType != nil {
		trafficType = types.StringValue(*d.TrafficType)
	}
	if d.Type != nil {
		serviceType = types.StringValue(string(*d.Type))
	}
	if d.Urls != nil {
		urls = mistutils.ListOfStringSdkToTerraform(d.Urls)
	}

	dataMapValue := map[string]attr.Value{
		"addresses":                         addresses,
		"app_categories":                    appCategories,
		"app_subcategories":                 appSubcategories,
		"apps":                              apps,
		"client_limit_down":                 clientLimitDown,
		"client_limit_up":                   clientLimitUp,
		"created_time":                      createdTime,
		"description":                       description,
		"dscp":                              dscp,
		"failover_policy":                   failoverPolicy,
		"hostnames":                         hostnames,
		"id":                                id,
		"max_jitter":                        maxJitter,
		"max_latency":                       maxLatency,
		"max_loss":                          maxLoss,
		"modified_time":                     modifiedTime,
		"name":                              name,
		"org_id":                            orgId,
		"service_limit_down":                serviceLimitDown,
		"service_limit_up":                  serviceLimitUp,
		"sle_enabled":                       sleEnabled,
		"specs":                             specs,
		"ssr_relaxed_tcp_state_enforcement": ssrRelaxedTcpStateEnforcement,
		"traffic_class":                     trafficClass,
		"traffic_type":                      trafficType,
		"type":                              serviceType,
		"urls":                              urls,
	}
	data, e := NewOrgServicesValue(OrgServicesValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
