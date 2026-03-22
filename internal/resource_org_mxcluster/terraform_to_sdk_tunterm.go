package resource_org_mxcluster

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tuntermDhcpdConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) *models.TuntermDhcpdConfig {
	data := models.TuntermDhcpdConfig{}

	if d.IsNull() || d.IsUnknown() {
		return &data
	}

	// Convert the map to a map of TuntermDhcpdConfigValue
	var tfMap map[string]TuntermDhcpdConfigValue
	d.ElementsAs(ctx, &tfMap, false)

	// Create the AdditionalProperties map
	additionalProps := make(map[string]models.TuntermDhcpdConfigProperty)

	for key, value := range tfMap {
		prop := models.TuntermDhcpdConfigProperty{}

		if !value.Enabled.IsNull() && !value.Enabled.IsUnknown() {
			prop.Enabled = value.Enabled.ValueBoolPointer()
		}

		if !value.Servers.IsNull() && !value.Servers.IsUnknown() {
			prop.Servers = mistutils.ListOfStringTerraformToSdk(value.Servers)
		}

		if !value.TuntermDhcpdConfigType.IsNull() && !value.TuntermDhcpdConfigType.IsUnknown() {
			prop.Type = (*models.TuntermDhcpdTypeEnum)(value.TuntermDhcpdConfigType.ValueStringPointer())
		}

		additionalProps[key] = prop
	}

	data.AdditionalProperties = additionalProps

	return &data
}

func tuntermExtraRoutesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.MxclusterTuntermExtraRoute {
	data := make(map[string]models.MxclusterTuntermExtraRoute)

	if d.IsNull() || d.IsUnknown() {
		return data
	}

	// Convert the map to a map of TuntermExtraRoutesValue
	var tfMap map[string]TuntermExtraRoutesValue
	d.ElementsAs(ctx, &tfMap, false)

	for key, value := range tfMap {
		route := models.MxclusterTuntermExtraRoute{}

		if !value.Via.IsNull() && !value.Via.IsUnknown() {
			route.Via = value.Via.ValueStringPointer()
		}

		data[key] = route
	}

	return data
}

func tuntermMonitoringTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) [][]models.TuntermMonitoringItem {
	var data [][]models.TuntermMonitoringItem

	if d.IsNull() || d.IsUnknown() {
		return data
	}

	// Get the outer list of lists
	var outerList []basetypes.ListValue
	d.ElementsAs(ctx, &outerList, false)

	for _, innerListValue := range outerList {
		var innerList []models.TuntermMonitoringItem

		// Get each inner list of objects
		var objectsList []basetypes.ObjectValue
		innerListValue.ElementsAs(ctx, &objectsList, false)

		for _, objValue := range objectsList {
			item := models.TuntermMonitoringItem{}

			// Extract attributes from the object
			attrs := objValue.Attributes()

			if host, ok := attrs["host"].(basetypes.StringValue); ok && !host.IsNull() && !host.IsUnknown() {
				item.Host = host.ValueStringPointer()
			}

			if port, ok := attrs["port"].(basetypes.Int64Value); ok && !port.IsNull() && !port.IsUnknown() {
				item.Port = models.ToPointer(int(port.ValueInt64()))
			}

			if protocol, ok := attrs["protocol"].(basetypes.StringValue); ok && !protocol.IsNull() && !protocol.IsUnknown() {
				item.Protocol = (*models.TuntermMonitoringProtocolEnum)(protocol.ValueStringPointer())
			}

			if srcVlanId, ok := attrs["src_vlan_id"].(basetypes.Int64Value); ok && !srcVlanId.IsNull() && !srcVlanId.IsUnknown() {
				item.SrcVlanId = models.ToPointer(int(srcVlanId.ValueInt64()))
			}

			if timeout, ok := attrs["timeout"].(basetypes.Int64Value); ok && !timeout.IsNull() && !timeout.IsUnknown() {
				item.Timeout = models.ToPointer(int(timeout.ValueInt64()))
			}

			innerList = append(innerList, item)
		}

		data = append(data, innerList)
	}

	return data
}
