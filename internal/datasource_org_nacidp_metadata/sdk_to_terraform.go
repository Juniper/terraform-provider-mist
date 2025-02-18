package datasource_org_nacidp_metadata

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func SdkToTerraform(data *models.SamlMetadata) OrgNacidpMetadataModel {
	var ds OrgNacidpMetadataModel

	if data.AcsUrl != nil {
		ds.AcsUrl = types.StringValue(*data.AcsUrl)
	}
	if data.EntityId != nil {
		ds.EntityId = types.StringValue(*data.EntityId)
	}
	if data.LogoutUrl != nil {
		ds.LogoutUrl = types.StringValue(*data.LogoutUrl)
	}
	if data.Metadata != nil {
		ds.Metadata = types.StringValue(*data.Metadata)
	}

	return ds
}
