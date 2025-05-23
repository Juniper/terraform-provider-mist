// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_org

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func OrgResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"alarmtemplate_id": schema.StringAttribute{
				Optional: true,
			},
			"allow_mist": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Unique ID of the object instance in the Mist Organization",
				MarkdownDescription: "Unique ID of the object instance in the Mist Organization",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"msp_id": schema.StringAttribute{
				Computed: true,
			},
			"msp_logo_url": schema.StringAttribute{
				Computed:            true,
				Description:         "logo uploaded by the MSP with advanced tier, only present if provided",
				MarkdownDescription: "logo uploaded by the MSP with advanced tier, only present if provided",
			},
			"msp_name": schema.StringAttribute{
				Computed:            true,
				Description:         "Name of the msp the org belongs to",
				MarkdownDescription: "Name of the msp the org belongs to",
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"orggroup_ids": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"session_expiry": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Validators: []validator.Int64{
					int64validator.Between(10, 20160),
				},
				Default: int64default.StaticInt64(1440),
			},
		},
	}
}

type OrgModel struct {
	AlarmtemplateId types.String `tfsdk:"alarmtemplate_id"`
	AllowMist       types.Bool   `tfsdk:"allow_mist"`
	Id              types.String `tfsdk:"id"`
	MspId           types.String `tfsdk:"msp_id"`
	MspLogoUrl      types.String `tfsdk:"msp_logo_url"`
	MspName         types.String `tfsdk:"msp_name"`
	Name            types.String `tfsdk:"name"`
	OrggroupIds     types.List   `tfsdk:"orggroup_ids"`
	SessionExpiry   types.Int64  `tfsdk:"session_expiry"`
}
