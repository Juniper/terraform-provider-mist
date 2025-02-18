package resource_org_psk

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	mistapi "github.com/Juniper/terraform-provider-mist/internal/commons/api_response"
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(d *models.Psk) (OrgPskModel, diag.Diagnostics) {
	var state OrgPskModel
	var diags diag.Diagnostics

	var email types.String
	var expireTime = types.Int64Value(0)
	var expiryNotificationTime types.Int64
	var id types.String
	var mac types.String
	var macs = types.ListNull(types.StringType)
	var maxUsage types.Int64
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
		macs = misttransform.ListOfStringSdkToTerraform(d.Macs)
	}
	if d.MaxUsage != nil {
		maxUsage = types.Int64Value(int64(*d.MaxUsage))
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

	usage = types.StringValue(string(*d.Usage))

	if d.VlanId != nil {
		vlanId = mistapi.PskVlanAsString(*d.VlanId)
	}

	state.Email = email
	state.ExpireTime = expireTime
	state.ExpiryNotificationTime = expiryNotificationTime
	state.Id = id
	state.Mac = mac
	state.Macs = macs
	state.MaxUsage = maxUsage
	state.Name = name
	state.Note = note
	state.NotifyExpiry = notifyExpiry
	state.NotifyOnCreateOrEdit = notifyOnCreateOrEdit
	state.OldPassphrase = oldPassphrase
	state.Passphrase = passphrase
	state.OrgId = orgId
	state.Role = role
	state.Ssid = ssid
	state.Usage = usage
	state.VlanId = vlanId

	return state, diags

}
