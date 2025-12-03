package resource_org_setting

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func ssrProxyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.SsrProxy {
	data := models.SsrProxy{}
	if !d.IsNull() && !d.IsUnknown() {
		item, e := NewProxyValue(ProxyValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		if e != nil {
			diags.Append(e...)
			return data
		} else {
			if item.Disabled.ValueBoolPointer() != nil {
				data.Disabled = item.Disabled.ValueBoolPointer()
			}
			if item.Url.ValueStringPointer() != nil {
				data.Url = item.Url.ValueStringPointer()
			}
		}
	}
	return data
}

func ssrAutoUpgradeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.SettingSsrAutoUpgrade {
	data := models.SettingSsrAutoUpgrade{}
	if !d.IsNull() && !d.IsUnknown() {
		item, e := NewSsrAutoUpgradeValue(SsrAutoUpgradeValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		if e != nil {
			diags.Append(e...)
			return data
		} else {
			if item.Channel.ValueStringPointer() != nil {
				channel := models.SsrUpgradeChannelEnum(*item.Channel.ValueStringPointer())
				data.Channel = &channel
			}
			if !item.CustomVersions.IsNull() && !item.CustomVersions.IsUnknown() {
				rMap := make(map[string]string)
				for k, v := range item.CustomVersions.Elements() {
					var vInterface interface{} = v
					vString := vInterface.(basetypes.StringValue)
					if vString.ValueStringPointer() != nil {
						rMap[k] = vString.ValueString()
					}
				}
				data.CustomVersions = rMap
			}
			if item.Enabled.ValueBoolPointer() != nil {
				data.Enabled = models.ToPointer(item.Enabled.ValueBool())
			}
		}
	}
	return data
}

func ssrTerraformToSdk(d SsrValue) *models.SettingSsr {
	data := models.SettingSsr{}

	if !d.ConductorHosts.IsNull() && !d.ConductorHosts.IsUnknown() {
		data.ConductorHosts = mistutils.ListOfStringTerraformToSdk(d.ConductorHosts)
	}
	if d.ConductorToken.ValueStringPointer() != nil {
		data.ConductorToken = d.ConductorToken.ValueStringPointer()
	}
	if d.DisableStats.ValueBoolPointer() != nil {
		data.DisableStats = d.DisableStats.ValueBoolPointer()
	}
	if !d.Proxy.IsNull() && !d.Proxy.IsUnknown() {
		data.Proxy = models.ToPointer(ssrProxyTerraformToSdk(context.Background(), &diag.Diagnostics{}, d.Proxy))
	}
	if !d.SsrAutoUpgrade.IsNull() && !d.SsrAutoUpgrade.IsUnknown() {
		data.AutoUpgrade = models.ToPointer(ssrAutoUpgradeTerraformToSdk(context.Background(), &diag.Diagnostics{}, d.SsrAutoUpgrade))
	}

	return &data
}
