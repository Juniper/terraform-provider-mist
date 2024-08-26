package resource_site_psk

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *SitePskModel) (models.Psk, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.Psk{}
	unset := make(map[string]interface{})

	if !plan.Email.IsNull() && !plan.Email.IsUnknown() {
		data.Email = plan.Email.ValueStringPointer()
	} else {
		unset["-email"] = ""
	}

	if !plan.ExpireTime.IsNull() && !plan.ExpireTime.IsUnknown() {
		data.ExpireTime = models.NewOptional(models.ToPointer(int(plan.ExpireTime.ValueInt64())))
	} else {
		unset["-expire_time"] = ""
	}

	if !plan.ExpiryNotificationTime.IsNull() && !plan.ExpiryNotificationTime.IsUnknown() {
		data.ExpiryNotificationTime = models.ToPointer(int(plan.ExpiryNotificationTime.ValueInt64()))
	} else {
		unset["-expiry_notification_time"] = ""
	}

	if !plan.Mac.IsNull() && !plan.Mac.IsUnknown() {
		data.Mac = plan.Mac.ValueStringPointer()
	} else {
		unset["-mac"] = ""
	}

	if !plan.Macs.IsNull() && !plan.Macs.IsUnknown() {
		data.Macs = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Macs)
	} else {
		unset["-macs"] = ""
	}

	data.Name = plan.Name.String()

	if !plan.Note.IsNull() && !plan.Note.IsUnknown() {
		data.Note = plan.Note.ValueStringPointer()
	} else {
		unset["-note"] = ""
	}

	if !plan.NotifyExpiry.IsNull() && !plan.NotifyExpiry.IsUnknown() {
		data.NotifyExpiry = plan.NotifyExpiry.ValueBoolPointer()
	} else {
		unset["-notify_expiry"] = ""
	}

	if !plan.NotifyOnCreateOrEdit.IsNull() && !plan.NotifyOnCreateOrEdit.IsUnknown() {
		data.NotifyOnCreateOrEdit = plan.NotifyOnCreateOrEdit.ValueBoolPointer()
	} else {
		unset["-notify_on_create_or_edit"] = ""
	}

	if !plan.OldPassphrase.IsNull() && !plan.OldPassphrase.IsUnknown() {
		data.OldPassphrase = plan.OldPassphrase.ValueStringPointer()
	} else {
		unset["-old_passphrase"] = ""
	}

	if !plan.Passphrase.IsNull() && !plan.Passphrase.IsUnknown() {
		data.Passphrase = plan.Passphrase.ValueString()
	} else {
		unset["-passphrase"] = ""
	}

	if !plan.Role.IsNull() && !plan.Role.IsUnknown() {
		data.Role = plan.Role.ValueStringPointer()
	} else {
		unset["-role"] = ""
	}

	if !plan.Ssid.IsNull() && !plan.Ssid.IsUnknown() {
		data.Ssid = plan.Ssid.ValueString()
	} else {
		unset["-ssid"] = ""
	}

	if !plan.Usage.IsNull() && !plan.Usage.IsUnknown() {
		data.Usage = models.ToPointer(models.PskUsageEnum(plan.Usage.ValueString()))
	} else {
		unset["-usage"] = ""
	}

	if !plan.VlanId.IsNull() && !plan.VlanId.IsUnknown() {
		data.VlanId = models.ToPointer(models.PskVlanIdContainer.FromString(plan.VlanId.ValueString()))
	} else {
		unset["-vlan_id"] = ""
	}

	data.AdditionalProperties = unset
	return data, diags
}
