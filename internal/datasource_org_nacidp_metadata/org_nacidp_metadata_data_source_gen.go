// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_org_nacidp_metadata

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func OrgNacidpMetadataDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"acs_url": schema.StringAttribute{
				Computed: true,
			},
			"entity_id": schema.StringAttribute{
				Computed: true,
			},
			"logout_url": schema.StringAttribute{
				Computed: true,
			},
			"metadata": schema.StringAttribute{
				Computed: true,
			},
			"nacidp_id": schema.StringAttribute{
				Required: true,
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

type OrgNacidpMetadataModel struct {
	AcsUrl    types.String `tfsdk:"acs_url"`
	EntityId  types.String `tfsdk:"entity_id"`
	LogoutUrl types.String `tfsdk:"logout_url"`
	Metadata  types.String `tfsdk:"metadata"`
	NacidpId  types.String `tfsdk:"nacidp_id"`
	OrgId     types.String `tfsdk:"org_id"`
}
