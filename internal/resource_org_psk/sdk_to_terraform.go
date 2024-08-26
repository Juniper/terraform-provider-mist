package resource_org_psk

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, d *models.Psk) (OrgPskModel, diag.Diagnostics) {
	var state OrgPskModel
	var diags diag.Diagnostics

	var email types.String
	var expire_time types.Int64 = types.Int64Value(0)
	var expiry_notification_time types.Int64
	var id types.String
	var mac types.String
	var macs types.List = types.ListNull(types.StringType)
	var max_usage types.Int64
	var name types.String
	var note types.String
	var notify_expiry types.Bool
	var notify_on_create_or_edit types.Bool
	var old_passphrase types.String
	var org_id types.String
	var passphrase types.String
	var role types.String
	var ssid types.String
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
	if d.Macs != nil {
		macs = mist_transform.ListOfStringSdkToTerraform(ctx, d.Macs)
	}
	if d.MaxUsage != nil {
		max_usage = types.Int64Value(int64(*d.MaxUsage))
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

	usage = types.StringValue(string(*d.Usage))

	if d.VlanId != nil {
		vlan_id = types.StringValue(d.VlanId.String())
	}

	state.Email = email
	state.ExpireTime = expire_time
	state.ExpiryNotificationTime = expiry_notification_time
	state.Id = id
	state.Mac = mac
	state.Macs = macs
	state.MaxUsage = max_usage
	state.Name = name
	state.Note = note
	state.NotifyExpiry = notify_expiry
	state.NotifyOnCreateOrEdit = notify_on_create_or_edit
	state.OldPassphrase = old_passphrase
	state.Passphrase = passphrase
	state.OrgId = org_id
	state.Role = role
	state.Ssid = ssid
	state.Usage = usage
	state.VlanId = vlan_id

	return state, diags

}
