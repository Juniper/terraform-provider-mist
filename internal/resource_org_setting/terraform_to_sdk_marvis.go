package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func marvisSelfDrivingDomainTerraformToSdk(o basetypes.ObjectValue, newValue func(map[string]interface{}) interface{}) *models.MarvisSelfDrivingDomain {
	data := models.MarvisSelfDrivingDomain{}
	if o.IsNull() || o.IsUnknown() {
		return &data
	}
	attrs := o.Attributes()
	if v, ok := attrs["enabled"]; ok {
		if bv, ok := v.(basetypes.BoolValue); ok && !bv.IsNull() && !bv.IsUnknown() {
			data.Enabled = models.ToPointer(bv.ValueBool())
		}
	}
	return &data
}

func marvisSelfDrivingTerraformToSdk(ctx context.Context, o basetypes.ObjectValue) *models.MarvisSelfDriving {
	data := models.MarvisSelfDriving{}
	if o.IsNull() || o.IsUnknown() {
		return &data
	}
	d := NewSelfDrivingValueMust(o.AttributeTypes(ctx), o.Attributes())
	if !d.Wan.IsNull() && !d.Wan.IsUnknown() {
		data.Wan = marvisSelfDrivingDomainTerraformToSdk(d.Wan, nil)
	}
	if !d.Wired.IsNull() && !d.Wired.IsUnknown() {
		data.Wired = marvisSelfDrivingDomainTerraformToSdk(d.Wired, nil)
	}
	if !d.Wireless.IsNull() && !d.Wireless.IsUnknown() {
		data.Wireless = marvisSelfDrivingDomainTerraformToSdk(d.Wireless, nil)
	}
	return &data
}

func marvisTerraformToSdk(ctx context.Context, d MarvisValue) *models.OrgSettingMarvis {
	data := models.OrgSettingMarvis{}
	if !d.DisableProactiveMonitoring.IsNull() && !d.DisableProactiveMonitoring.IsUnknown() {
		data.DisableProactiveMonitoring = d.DisableProactiveMonitoring.ValueBoolPointer()
	}
	if !d.SelfDriving.IsNull() && !d.SelfDriving.IsUnknown() {
		data.SelfDriving = marvisSelfDrivingTerraformToSdk(ctx, d.SelfDriving)
	}
	return &data
}
