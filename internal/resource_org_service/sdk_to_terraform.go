package resource_org_service

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

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

func SdkToTerraform(ctx context.Context, d *models.Service) (OrgServiceModel, diag.Diagnostics) {
	var state OrgServiceModel
	var diags diag.Diagnostics

	var addresses = misttransform.ListOfStringSdkToTerraformEmpty()
	var appCategories = misttransform.ListOfStringSdkToTerraformEmpty()
	var appSubcategories = misttransform.ListOfStringSdkToTerraformEmpty()
	var apps = misttransform.ListOfStringSdkToTerraformEmpty()
	var clientLimitDown types.Int64
	var clientLimitUp types.Int64
	var description types.String
	var dscp types.String
	var failoverPolicy types.String
	var hostnames = misttransform.ListOfStringSdkToTerraformEmpty()
	var id types.String
	var maxJitter types.String
	var maxLatency types.String
	var maxLoss types.String
	var name types.String
	var orgId types.String
	var serviceLimitDown types.Int64
	var serviceLimitUp types.Int64
	var sleEnabled types.Bool
	var specs = types.ListNull(SpecsValue{}.Type(ctx))
	var ssrRelaxedTcpStateEnforcement types.Bool
	var trafficClass types.String
	var trafficType types.String
	var serviceType types.String
	var urls = misttransform.ListOfStringSdkToTerraformEmpty()

	if d.Addresses != nil {
		addresses = misttransform.ListOfStringSdkToTerraform(d.Addresses)
	}
	if d.AppCategories != nil {
		appCategories = misttransform.ListOfStringSdkToTerraform(d.AppCategories)
	}
	if d.AppSubcategories != nil {
		appSubcategories = misttransform.ListOfStringSdkToTerraform(d.AppSubcategories)
	}
	if d.Apps != nil {
		apps = misttransform.ListOfStringSdkToTerraform(d.Apps)
	}
	if d.ClientLimitDown != nil {
		clientLimitDown = types.Int64Value(int64(*d.ClientLimitDown))
	}
	if d.ClientLimitUp != nil {
		clientLimitUp = types.Int64Value(int64(*d.ClientLimitUp))
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
		hostnames = misttransform.ListOfStringSdkToTerraform(d.Hostnames)
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
		specs = serviceSpecsSdkToTerraform(ctx, &diags, d.Specs)
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
		urls = misttransform.ListOfStringSdkToTerraform(d.Urls)
	}

	state.Addresses = addresses
	state.AppCategories = appCategories
	state.AppSubcategories = appSubcategories
	state.Apps = apps
	state.ClientLimitDown = clientLimitDown
	state.ClientLimitUp = clientLimitUp
	state.Description = description
	state.Dscp = dscp
	state.FailoverPolicy = failoverPolicy
	state.Hostnames = hostnames
	state.Id = id
	state.MaxJitter = maxJitter
	state.MaxLatency = maxLatency
	state.MaxLoss = maxLoss
	state.Name = name
	state.OrgId = orgId
	state.ServiceLimitDown = serviceLimitDown
	state.ServiceLimitUp = serviceLimitUp
	state.SleEnabled = sleEnabled
	state.Specs = specs
	state.SsrRelaxedTcpStateEnforcement = ssrRelaxedTcpStateEnforcement
	state.TrafficClass = trafficClass
	state.TrafficType = trafficType
	state.Type = serviceType
	state.Urls = urls

	return state, diags

}
