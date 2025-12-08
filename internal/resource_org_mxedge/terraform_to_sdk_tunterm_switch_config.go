package resource_org_mxedge

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tuntermSwitchConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) *models.MxedgeTuntermSwitchConfigs {
	data := models.MxedgeTuntermSwitchConfigs{}

	if d.IsNull() || d.IsUnknown() {
		return &data
	}

	// Convert the map to a map of TuntermSwitchConfigValue
	var tfMap map[string]TuntermSwitchConfigValue
	d.ElementsAs(ctx, &tfMap, false)

	// Create the AdditionalProperties map
	additionalProps := make(map[string]models.MxedgeTuntermSwitchConfig)

	for key, value := range tfMap {
		config := models.MxedgeTuntermSwitchConfig{}

		if !value.PortVlanId.IsNull() && !value.PortVlanId.IsUnknown() {
			config.PortVlanId = models.ToPointer(int(value.PortVlanId.ValueInt64()))
		}

		if !value.VlanIds.IsNull() && !value.VlanIds.IsUnknown() {
			vlanIdStrings := mistutils.ListOfStringTerraformToSdk(value.VlanIds)
			vlanIds := make([]models.VlanIdWithVariable, len(vlanIdStrings))
			for i, vlanId := range vlanIdStrings {
				vlanIds[i] = models.VlanIdWithVariableContainer.FromString(vlanId)
			}
			config.VlanIds = vlanIds
		}

		additionalProps[key] = config
	}

	data.AdditionalProperties = additionalProps

	return &data
}
