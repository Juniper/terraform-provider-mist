package datasource_org_psks

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.Psk, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := pskSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func pskSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Psk) OrgPsksValue {
	var state OrgPsksValue

	var adminSsoId types.String
	var createdTime basetypes.Float64Value
	var email types.String
	var expireTime types.Int64
	var expiryNotificationTime types.Int64
	var id types.String
	var mac types.String
	var macs = types.ListNull(types.StringType)
	var maxUsage types.Int64
	var modifiedTime basetypes.Float64Value
	var name types.String
	var note types.String
	var notifyExpiry types.Bool
	var notifyOnCreateOrEdit types.Bool
	var oldPassphrase types.String
	var orgId types.String
	var passphrase types.String
	var role types.String
	var ssid types.String
	var usage types.String
	var vlanId types.String

	if d.AdminSsoId != nil {
		adminSsoId = types.StringValue(*d.AdminSsoId)
	}
	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.Email != nil {
		email = types.StringValue(*d.Email)
	}
	if d.ExpireTime.Value() != nil {
		expireTime = types.Int64Value(int64(*d.ExpireTime.Value()))
	}
	if d.ExpiryNotificationTime != nil {
		expiryNotificationTime = types.Int64Value(int64(*d.ExpiryNotificationTime))
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.Mac != nil {
		mac = types.StringValue(*d.Mac)
	}
	if d.Macs != nil {
		macs = mistutils.ListOfStringSdkToTerraform(d.Macs)
	}
	if d.MaxUsage != nil {
		maxUsage = types.Int64Value(int64(*d.MaxUsage))
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}

	name = types.StringValue(d.Name)

	if d.Note != nil {
		note = types.StringValue(*d.Note)
	}
	if d.NotifyExpiry != nil {
		notifyExpiry = types.BoolValue(*d.NotifyExpiry)
	}
	if d.NotifyOnCreateOrEdit != nil {
		notifyOnCreateOrEdit = types.BoolValue(*d.NotifyOnCreateOrEdit)
	}
	if d.OldPassphrase != nil {
		oldPassphrase = types.StringValue(*d.OldPassphrase)
	}
	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}

	passphrase = types.StringValue(d.Passphrase)

	if d.Role != nil {
		role = types.StringValue(*d.Role)
	}

	ssid = types.StringValue(d.Ssid)

	if d.Usage != nil {
		usage = types.StringValue(string(*d.Usage))
	}

	if d.VlanId != nil {
		vlanId = types.StringValue(d.VlanId.String())
	}

	dataMapValue := map[string]attr.Value{
		"admin_sso_id":             adminSsoId,
		"created_time":             createdTime,
		"email":                    email,
		"expire_time":              expireTime,
		"expiry_notification_time": expiryNotificationTime,
		"id":                       id,
		"mac":                      mac,
		"macs":                     macs,
		"max_usage":                maxUsage,
		"modified_time":            modifiedTime,
		"name":                     name,
		"note":                     note,
		"notify_expiry":            notifyExpiry,
		"notify_on_create_or_edit": notifyOnCreateOrEdit,
		"old_passphrase":           oldPassphrase,
		"org_id":                   orgId,
		"passphrase":               passphrase,
		"role":                     role,
		"ssid":                     ssid,
		"usage":                    usage,
		"vlan_id":                  vlanId,
	}

	state, e := NewOrgPsksValue(OrgPsksValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return state

}
