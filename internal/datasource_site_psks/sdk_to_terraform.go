package datasource_site_psks

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l []models.Psk) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		elem := pskSdkToTerraform(ctx, &diags, d)
		elements = append(elements, elem)
	}

	dataSet, err := types.SetValue(SitePsksValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func pskSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.Psk) SitePsksValue {
	var state SitePsksValue

	var email types.String
	var expire_time types.Int64
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

	if d.SiteId != nil {
		site_id = types.StringValue(d.SiteId.String())
	}

	ssid = types.StringValue(d.Ssid)

	usage = types.StringValue(string(*d.Usage))

	if d.VlanId != nil {
		vlan_id = types.StringValue(d.VlanId.String())
	}

	data_map_attr_type := SitePsksValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"email":                    email,
		"expire_time":              expire_time,
		"expiry_notification_time": expiry_notification_time,
		"id":                       id,
		"mac":                      mac,
		"name":                     name,
		"note":                     note,
		"notify_expiry":            notify_expiry,
		"notify_on_create_or_edit": notify_on_create_or_edit,
		"old_passphrase":           old_passphrase,
		"org_id":                   org_id,
		"passphrase":               passphrase,
		"role":                     role,
		"site_id":                  site_id,
		"ssid":                     ssid,
		"usage":                    usage,
		"vlan_id":                  vlan_id,
	}

	state, e := NewSitePsksValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return state

}
