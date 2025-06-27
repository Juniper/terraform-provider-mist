package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func junosShellAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingJunosShellAccess) JunosShellAccessValue {

	var admin basetypes.StringValue
	var helpdesk basetypes.StringValue
	var read basetypes.StringValue
	var write basetypes.StringValue

	if d.Admin != nil {
		admin = types.StringValue(string(*d.Admin))
	}
	if d.Helpdesk != nil {
		helpdesk = types.StringValue(string(*d.Helpdesk))
	}
	if d.Read != nil {
		read = types.StringValue(string(*d.Read))
	}
	if d.Write != nil {
		write = types.StringValue(string(*d.Write))
	}

	dataMapValue := map[string]attr.Value{
		"admin":    admin,
		"helpdesk": helpdesk,
		"read":     read,
		"write":    write,
	}
	data, e := NewJunosShellAccessValue(JunosShellAccessValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
