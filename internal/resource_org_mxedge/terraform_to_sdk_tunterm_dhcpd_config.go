package resource_org_mxedge

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tuntermDhcpdConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) *models.MxedgeTuntermDhcpdConfig {
	data := models.MxedgeTuntermDhcpdConfig{}

	if d.IsNull() || d.IsUnknown() {
		return &data
	}

	// Convert the map to a map of TuntermDhcpdConfigValue
	var tfMap map[string]TuntermDhcpdConfigValue
	d.ElementsAs(ctx, &tfMap, false)

	// Create the AdditionalProperties map
	additionalProps := make(map[string]models.MxedgeTuntermDhcpdConfigProperty)

	for key, value := range tfMap {
		prop := models.MxedgeTuntermDhcpdConfigProperty{}

		if !value.Enabled.IsNull() && !value.Enabled.IsUnknown() {
			prop.Enabled = value.Enabled.ValueBoolPointer()
		}

		if !value.Servers.IsNull() && !value.Servers.IsUnknown() {
			prop.Servers = mistutils.ListOfStringTerraformToSdk(value.Servers)
		}

		if !value.TuntermDhcpdConfigType.IsNull() && !value.TuntermDhcpdConfigType.IsUnknown() {
			prop.Type = (*models.MxedgeTuntermDhcpdConfigTypeEnum)(value.TuntermDhcpdConfigType.ValueStringPointer())
		}

		additionalProps[key] = prop
	}

	data.AdditionalProperties = additionalProps

	return &data
}
