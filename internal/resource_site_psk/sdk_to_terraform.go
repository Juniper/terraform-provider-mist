package resource_site_psk

import (
	"context"

	mist_api "github.com/Juniper/terraform-provider-mist/internal/commons/api_response"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, d *models.Psk) (SitePskModel, diag.Diagnostics) {
	var state SitePskModel
	var diags diag.Diagnostics

	var email types.String
	var expire_time types.Int64 = types.Int64Value(0)
	var expiry_notification_time types.Int64
	var id types.String
	var mac types.String
	var name types.String
	var note types.String
	var notify_expiry types.Bool
	var notify_on_create_or_edit types.Bool
	var old_passphrase types.String
	var org_id types.String
	var passphrase types.String
	var role types.String
	var ssid types.String
	var site_id types.String
	var usage types.String
	var vlan_id types.String

	if d.Email != nil {
		email = types.StringValue(*d.Email)
	}
	if d.ExpireTime.Value() != nil {
		expire_time = types.Int64Value(int64(*d.ExpireTime.Value()))
	}
	if d.ExpiryNotificationTime != nil {
		expiry_notification_time = types.Int64Value(int64(*d.ExpiryNotificationTime))
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.Mac != nil {
		mac = types.StringValue(*d.Mac)
	}

	name = types.StringValue(d.Name)

	if d.Note != nil {
		note = types.StringValue(*d.Note)
	}
	if d.NotifyExpiry != nil {
		notify_expiry = types.BoolValue(*d.NotifyExpiry)
	}
	if d.NotifyOnCreateOrEdit != nil {
		notify_on_create_or_edit = types.BoolValue(*d.NotifyOnCreateOrEdit)
	}
	if d.OldPassphrase != nil {
		old_passphrase = types.StringValue(*d.OldPassphrase)
	}
	if d.OrgId != nil {
		org_id = types.StringValue(d.OrgId.String())
	}

	passphrase = types.StringValue(d.Passphrase)

	if d.Role != nil {
		role = types.StringValue(*d.Role)
	}

	ssid = types.StringValue(d.Ssid)

	site_id = types.StringValue(string(d.SiteId.String()))

	usage = types.StringValue(string(*d.Usage))

	if d.VlanId != nil {
		vlan_id = mist_api.PskVlanAsString(*d.VlanId)
	}

	state.Email = email
	state.ExpireTime = expire_time
	state.ExpiryNotificationTime = expiry_notification_time
	state.Id = id
	state.Mac = mac
	state.Name = name
	state.Note = note
	state.NotifyExpiry = notify_expiry
	state.NotifyOnCreateOrEdit = notify_on_create_or_edit
	state.OldPassphrase = old_passphrase
	state.Passphrase = passphrase
	state.OrgId = org_id
	state.Role = role
	state.Ssid = ssid
	state.SiteId = site_id
	state.Usage = usage
	state.VlanId = vlan_id

	return state, diags

}
