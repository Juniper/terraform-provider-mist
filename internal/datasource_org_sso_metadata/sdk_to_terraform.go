package datasource_org_sso_metadata

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func SdkToTerraform(data *models.SamlMetadata) OrgSsoMetadataModel {
	var ds OrgSsoMetadataModel

	ds.AcsUrl = types.StringValue(*data.AcsUrl)
	ds.EntityId = types.StringValue(*data.EntityId)
	ds.LogoutUrl = types.StringValue(*data.LogoutUrl)
	ds.Metadata = types.StringValue(*data.Metadata)

	return ds
}
