package datasource_device_gateway_stats

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func clusterConfigControlLinkInfoSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsClusterConfigControlLinkInfo) basetypes.ObjectValue {
	var name basetypes.StringValue
	var status basetypes.StringValue

	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.Status != nil {
		status = types.StringValue(*d.Status)
	}

	dataMapValue := map[string]attr.Value{
		"name":   name,
		"status": status,
	}

	data, e := types.ObjectValue(ControlLinkInfoValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func clusterConfigStatsEthernetConnectionSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.StatsClusterConfigEthernetConnectionItem) basetypes.ListValue {
	var dataList []EthernetConnectionValue
	for _, d := range l {

		var name basetypes.StringValue
		var status basetypes.StringValue

		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.Status != nil {
			status = types.StringValue(*d.Status)
		}

		dataMapValue := map[string]attr.Value{
			"name":   name,
			"status": status,
		}
		data, e := NewEthernetConnectionValue(EthernetConnectionValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, EthernetConnectionValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func clusterConfigFabricLinkInfoSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsClusterConfigFabricLinkInfo) basetypes.ObjectValue {
	var dataPlaneNotifiedStatus basetypes.StringValue
	var interfaceLink = types.ListNull(types.StringType)
	var internalStatus basetypes.StringValue
	var state basetypes.StringValue
	var status basetypes.StringValue

	if d.DataPlaneNotifiedStatus != nil {
		dataPlaneNotifiedStatus = types.StringValue(*d.DataPlaneNotifiedStatus)
	}
	if d.Interface != nil {
		interfaceLink = mistutils.ListOfStringSdkToTerraform(d.Interface)
	}
	if d.InternalStatus != nil {
		internalStatus = types.StringValue(*d.InternalStatus)
	}
	if d.State != nil {
		state = types.StringValue(*d.State)
	}
	if d.Status != nil {
		status = types.StringValue(*d.Status)
	}

	dataMapValue := map[string]attr.Value{
		"data_plane_notified_status": dataPlaneNotifiedStatus,
		"interface":                  interfaceLink,
		"internal_status":            internalStatus,
		"state":                      state,
		"status":                     status,
	}
	data, e := types.ObjectValue(FabricLinkInfoValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func clusterConfigStatsRedundancySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.StatsClusterConfigRedundancyGroupInfoItem) basetypes.ListValue {
	var dataList []RedundancyGroupInformationValue
	for _, d := range l {

		var id basetypes.Int64Value
		var monitoringFailure basetypes.StringValue
		var threshold basetypes.Int64Value

		if d.Id != nil {
			id = types.Int64Value(int64(*d.Id))
		}
		if d.MonitoringFailure != nil {
			monitoringFailure = types.StringValue(*d.MonitoringFailure)
		}
		if d.Threshold != nil {
			threshold = types.Int64Value(int64(*d.Threshold))
		}

		dataMapValue := map[string]attr.Value{
			"id":                 id,
			"monitoring_failure": monitoringFailure,
			"threshold":          threshold,
		}
		data, e := NewRedundancyGroupInformationValue(RedundancyGroupInformationValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, RedundancyGroupInformationValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func clusterConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsClusterConfig) basetypes.ObjectValue {
	var configuration basetypes.StringValue
	var controlLinkInfo = types.ObjectNull(ControlLinkInfoValue{}.AttributeTypes(ctx))
	var ethernetConnection = types.ListNull(EthernetConnectionValue{}.Type(ctx))
	var fabricLinkInfo = types.ObjectNull(FabricLinkInfoValue{}.AttributeTypes(ctx))
	var lastStatusChangeReason basetypes.StringValue
	var operational basetypes.StringValue
	var primaryNodeHealth basetypes.StringValue
	var redundancyGroupInformation = types.ListNull(RedundancyGroupInformationValue{}.Type(ctx))
	var secondaryNodeHealth basetypes.StringValue
	var status basetypes.StringValue

	if d.Configuration != nil {
		configuration = types.StringValue(*d.Configuration)
	}
	if d.ControlLinkInfo != nil {
		controlLinkInfo = clusterConfigControlLinkInfoSdkToTerraform(ctx, diags, d.ControlLinkInfo)
	}
	if d.EthernetConnection != nil {
		ethernetConnection = clusterConfigStatsEthernetConnectionSdkToTerraform(ctx, diags, d.EthernetConnection)
	}
	if d.FabricLinkInfo != nil {
		fabricLinkInfo = clusterConfigFabricLinkInfoSdkToTerraform(ctx, diags, d.FabricLinkInfo)
	}
	if d.LastStatusChangeReason != nil {
		lastStatusChangeReason = types.StringValue(*d.LastStatusChangeReason)
	}
	if d.Operational != nil {
		operational = types.StringValue(*d.Operational)
	}
	if d.PrimaryNodeHealth != nil {
		primaryNodeHealth = types.StringValue(*d.PrimaryNodeHealth)
	}
	if d.RedundancyGroupInformation != nil {
		redundancyGroupInformation = clusterConfigStatsRedundancySdkToTerraform(ctx, diags, d.RedundancyGroupInformation)
	}
	if d.SecondaryNodeHealth != nil {
		secondaryNodeHealth = types.StringValue(*d.SecondaryNodeHealth)
	}
	if d.Status != nil {
		status = types.StringValue(*d.Status)
	}

	dataMapValue := map[string]attr.Value{
		"configuration":                configuration,
		"control_link_info":            controlLinkInfo,
		"ethernet_connection":          ethernetConnection,
		"fabric_link_info":             fabricLinkInfo,
		"last_status_change_reason":    lastStatusChangeReason,
		"operational":                  operational,
		"primary_node_health":          primaryNodeHealth,
		"redundancy_group_information": redundancyGroupInformation,
		"secondary_node_health":        secondaryNodeHealth,
		"status":                       status,
	}

	data, e := types.ObjectValue(ClusterConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
