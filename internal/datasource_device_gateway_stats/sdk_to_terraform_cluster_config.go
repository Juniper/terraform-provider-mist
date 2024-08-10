package datasource_device_gateway_stats

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

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

	data_map_attr_type := ControlLinkInfoValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"name":   name,
		"status": status,
	}

	data, e := types.ObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func clusterConfigStatsEthernetConnectionSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.StatsClusterConfigEthernetConnectionItem) basetypes.ListValue {
	var data_list = []EthernetConnectionValue{}
	for _, d := range l {

		var name basetypes.StringValue
		var status basetypes.StringValue

		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.Status != nil {
			status = types.StringValue(*d.Status)
		}

		data_map_attr_type := EthernetConnectionValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"name":   name,
			"status": status,
		}
		data, e := NewEthernetConnectionValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, EthernetConnectionValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

func clusterConfigFabricLinkInfoSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsClusterConfigFabricLinkInfo) basetypes.ObjectValue {
	var data_plane_notified_status basetypes.StringValue
	var interface_link basetypes.ListValue = types.ListNull(types.StringType)
	var internal_status basetypes.StringValue
	var state basetypes.StringValue
	var status basetypes.StringValue

	if d.DataPlaneNotifiedStatus != nil {
		data_plane_notified_status = types.StringValue(*d.DataPlaneNotifiedStatus)
	}
	if d.Interface != nil {
		interface_link = mist_transform.ListOfStringSdkToTerraform(ctx, d.Interface)
	}
	if d.InternalStatus != nil {
		internal_status = types.StringValue(*d.InternalStatus)
	}
	if d.State != nil {
		state = types.StringValue(*d.State)
	}
	if d.Status != nil {
		status = types.StringValue(*d.Status)
	}

	data_map_attr_type := FabricLinkInfoValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"data_plane_notified_status": data_plane_notified_status,
		"interface":                  interface_link,
		"internal_status":            internal_status,
		"state":                      state,
		"status":                     status,
	}
	data, e := types.ObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func clusterConfigStatsRedundancySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.StatsClusterConfigRedundancyGroupInfoItem) basetypes.ListValue {
	var data_list = []RedundancyGroupInformationValue{}
	for _, d := range l {

		var id basetypes.Int64Value
		var monitoring_failure basetypes.StringValue
		var threshold basetypes.Int64Value

		if d.Id != nil {
			id = types.Int64Value(int64(*d.Id))
		}
		if d.MonitoringFailure != nil {
			monitoring_failure = types.StringValue(*d.MonitoringFailure)
		}
		if d.Threshold != nil {
			threshold = types.Int64Value(int64(*d.Threshold))
		}

		data_map_attr_type := RedundancyGroupInformationValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"id":                 id,
			"monitoring_failure": monitoring_failure,
			"threshold":          threshold,
		}
		data, e := NewRedundancyGroupInformationValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, RedundancyGroupInformationValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

func clusterConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsClusterConfig) basetypes.ObjectValue {
	var configuration basetypes.StringValue
	var control_link_info basetypes.ObjectValue = types.ObjectNull(ControlLinkInfoValue{}.AttributeTypes(ctx))
	var ethernet_connection basetypes.ListValue = types.ListNull(EthernetConnectionValue{}.Type(ctx))
	var fabric_link_info basetypes.ObjectValue = types.ObjectNull(FabricLinkInfoValue{}.AttributeTypes(ctx))
	var last_status_change_reason basetypes.StringValue
	var operational basetypes.StringValue
	var primary_node_health basetypes.StringValue
	var redundancy_group_information basetypes.ListValue = types.ListNull(RedundancyGroupInformationValue{}.Type(ctx))
	var secondary_node_health basetypes.StringValue
	var status basetypes.StringValue

	if d.Configuration != nil {
		configuration = types.StringValue(*d.Configuration)
	}
	if d.ControlLinkInfo != nil {
		control_link_info = clusterConfigControlLinkInfoSdkToTerraform(ctx, diags, d.ControlLinkInfo)
	}
	if d.EthernetConnection != nil {
		ethernet_connection = clusterConfigStatsEthernetConnectionSdkToTerraform(ctx, diags, d.EthernetConnection)
	}
	if d.FabricLinkInfo != nil {
		fabric_link_info = clusterConfigFabricLinkInfoSdkToTerraform(ctx, diags, d.FabricLinkInfo)
	}
	if d.LastStatusChangeReason != nil {
		last_status_change_reason = types.StringValue(*d.LastStatusChangeReason)
	}
	if d.Operational != nil {
		operational = types.StringValue(*d.Operational)
	}
	if d.PrimaryNodeHealth != nil {
		primary_node_health = types.StringValue(*d.PrimaryNodeHealth)
	}
	if d.RedundancyGroupInformation != nil {
		redundancy_group_information = clusterConfigStatsRedundancySdkToTerraform(ctx, diags, d.RedundancyGroupInformation)
	}
	if d.SecondaryNodeHealth != nil {
		secondary_node_health = types.StringValue(*d.SecondaryNodeHealth)
	}
	if d.Status != nil {
		status = types.StringValue(*d.Status)
	}

	data_map_attr_type := ClusterConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"configuration":                configuration,
		"control_link_info":            control_link_info,
		"ethernet_connection":          ethernet_connection,
		"fabric_link_info":             fabric_link_info,
		"last_status_change_reason":    last_status_change_reason,
		"operational":                  operational,
		"primary_node_health":          primary_node_health,
		"redundancy_group_information": redundancy_group_information,
		"secondary_node_health":        secondary_node_health,
		"status":                       status,
	}

	data, e := types.ObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
