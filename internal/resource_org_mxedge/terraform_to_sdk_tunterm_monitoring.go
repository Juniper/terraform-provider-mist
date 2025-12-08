package resource_org_mxedge

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

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
