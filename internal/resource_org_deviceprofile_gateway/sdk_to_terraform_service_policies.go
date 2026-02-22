package resource_org_deviceprofile_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func servicePolicyAppQoESdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ServicePolicyAppqoe) basetypes.ObjectValue {
	var enabled basetypes.BoolValue

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	rAttrValue := map[string]attr.Value{
		"enabled": enabled,
	}
	r, e := basetypes.NewObjectValue(AppqoeValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)
	return r
}

func servicePolicyEwfSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.ServicePolicyEwfRule) basetypes.ListValue {
	var dataList []EwfValue
	for _, v := range d {
		var alertOnly basetypes.BoolValue
		var blockMessage basetypes.StringValue
		var enabled basetypes.BoolValue
		var profile basetypes.StringValue

		if v.AlertOnly != nil {
			alertOnly = types.BoolValue(*v.AlertOnly)
		}
		if v.BlockMessage != nil {
			blockMessage = types.StringValue(*v.BlockMessage)
		}
		if v.Enabled != nil {
			enabled = types.BoolValue(*v.Enabled)
		}
		if v.Profile != nil {
			profile = types.StringValue(string(*v.Profile))
		}

		dataMapValue := map[string]attr.Value{
			"alert_only":    alertOnly,
			"block_message": blockMessage,
			"enabled":       enabled,
			"profile":       profile,
		}
		data, e := NewEwfValue(EwfValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, EwfValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func servicePolicyIdpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.IdpConfig) basetypes.ObjectValue {
	var alertOnly basetypes.BoolValue
	var enabled basetypes.BoolValue
	var idpprofileId basetypes.StringValue
	var profile basetypes.StringValue

	if d != nil && d.AlertOnly != nil {
		alertOnly = types.BoolValue(*d.AlertOnly)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.IdpprofileId != nil {
		idpprofileId = types.StringValue(d.IdpprofileId.String())
	}
	if d != nil && d.Profile != nil {
		profile = types.StringValue(*d.Profile)
	}

	rAttrValue := map[string]attr.Value{
		"alert_only":    alertOnly,
		"enabled":       enabled,
		"idpprofile_id": idpprofileId,
		"profile":       profile,
	}
	r, e := basetypes.NewObjectValue(IdpValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)
	return r
}

func avSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ServicePolicyAntivirus) basetypes.ObjectValue {

	var avprofileId basetypes.StringValue
	var enabled basetypes.BoolValue
	var profile basetypes.StringValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.AvprofileId != nil {
		avprofileId = types.StringValue(d.AvprofileId.String())
	}
	if d.Profile != nil && (d.AvprofileId == nil || *d.Profile != d.AvprofileId.String()) {
		profile = types.StringValue(*d.Profile)
	}

	rAttrValue := map[string]attr.Value{
		"avprofile_id": avprofileId,
		"enabled":      enabled,
		"profile":      profile,
	}
	r, e := basetypes.NewObjectValue(AntivirusValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)

	return r
}

func skyatpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ServicePolicySkyatp) basetypes.ObjectValue {

	var dnsDgaDetection = types.ObjectNull(DnsDgaDetectionValue{}.AttributeTypes(ctx))
	var dnsTunnelDetection = types.ObjectNull(DnsTunnelDetectionValue{}.AttributeTypes(ctx))
	var httpInspection = types.ObjectNull(HttpInspectionValue{}.AttributeTypes(ctx))
	var iotDevicePolicy = types.ObjectNull(IotDevicePolicyValue{}.AttributeTypes(ctx))

	if d.DnsDgaDetection != nil {
		var enabled basetypes.BoolValue
		var profile basetypes.StringValue

		if d.DnsDgaDetection.Enabled != nil {
			enabled = types.BoolValue(*d.DnsDgaDetection.Enabled)
		}
		if d.DnsDgaDetection.Profile != nil {
			profile = types.StringValue(string(*d.DnsDgaDetection.Profile))
		}

		dnsDgaAttr := map[string]attr.Value{
			"enabled": enabled,
			"profile": profile,
		}
		dnsDgaDetection, _ = types.ObjectValue(DnsDgaDetectionValue{}.AttributeTypes(ctx), dnsDgaAttr)
	}
	if d.DnsTunnelDetection != nil {
		var enabled basetypes.BoolValue
		var profile basetypes.StringValue

		if d.DnsTunnelDetection.Enabled != nil {
			enabled = types.BoolValue(*d.DnsTunnelDetection.Enabled)
		}
		if d.DnsTunnelDetection.Profile != nil {
			profile = types.StringValue(string(*d.DnsTunnelDetection.Profile))
		}

		dnsTunnelAttr := map[string]attr.Value{
			"enabled": enabled,
			"profile": profile,
		}
		dnsTunnelDetection, _ = types.ObjectValue(DnsTunnelDetectionValue{}.AttributeTypes(ctx), dnsTunnelAttr)
	}
	if d.HttpInspection != nil {
		var enabled basetypes.BoolValue
		var profile basetypes.StringValue

		if d.HttpInspection.Enabled != nil {
			enabled = types.BoolValue(*d.HttpInspection.Enabled)
		}
		if d.HttpInspection.Profile != nil {
			profile = types.StringValue(string(*d.HttpInspection.Profile))
		}

		httpInspAttr := map[string]attr.Value{
			"enabled": enabled,
			"profile": profile,
		}
		httpInspection, _ = types.ObjectValue(HttpInspectionValue{}.AttributeTypes(ctx), httpInspAttr)
	}
	if d.IotDevicePolicy != nil {
		var enabled basetypes.BoolValue

		if d.IotDevicePolicy.Enabled != nil {
			enabled = types.BoolValue(*d.IotDevicePolicy.Enabled)
		}

		iotPolicyAttr := map[string]attr.Value{
			"enabled": enabled,
		}
		iotDevicePolicy, _ = types.ObjectValue(IotDevicePolicyValue{}.AttributeTypes(ctx), iotPolicyAttr)
	}

	rAttrValue := map[string]attr.Value{
		"dns_dga_detection":    dnsDgaDetection,
		"dns_tunnel_detection": dnsTunnelDetection,
		"http_inspection":      httpInspection,
		"iot_device_policy":    iotDevicePolicy,
	}
	r, e := basetypes.NewObjectValue(SkyatpValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)

	return r
}

func syslogSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ServicePolicySyslog) basetypes.ObjectValue {

	var enabled basetypes.BoolValue
	var serverNames = types.ListNull(types.StringType)

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.ServerNames != nil {
		serverNames = mistutils.ListOfStringSdkToTerraform(d.ServerNames)
	}

	rAttrValue := map[string]attr.Value{
		"enabled":      enabled,
		"server_names": serverNames,
	}
	r, e := basetypes.NewObjectValue(SyslogValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)

	return r
}

func sslProxySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ServicePolicySslProxy) basetypes.ObjectValue {

	var CiphersCategory basetypes.StringValue
	var enabled basetypes.BoolValue

	if d.CiphersCategory != nil {
		CiphersCategory = types.StringValue(string(*d.CiphersCategory))
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	rAttrValue := map[string]attr.Value{
		"ciphers_category": CiphersCategory,
		"enabled":          enabled,
	}
	r, e := basetypes.NewObjectValue(SslProxyValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)

	return r
}

func servicePoliciesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.ServicePolicy) basetypes.ListValue {
	var dataList []ServicePoliciesValue

	for _, v := range d {

		var action basetypes.StringValue
		var antivirus = types.ObjectNull(AntivirusValue{}.AttributeTypes(ctx))
		var appqoe = types.ObjectNull(AppqoeValue{}.AttributeTypes(ctx))
		var ewf = types.ListNull(EwfValue{}.Type(ctx))
		var idp = types.ObjectNull(IdpValue{}.AttributeTypes(ctx))
		var localRouting basetypes.BoolValue
		var name basetypes.StringValue
		var pathPreference basetypes.StringValue
		var servicepolicyId basetypes.StringValue
		var services = types.ListNull(types.StringType)
		var skyatp = types.ObjectNull(SkyatpValue{}.AttributeTypes(ctx))
		var sslProxy = types.ObjectNull(SslProxyValue{}.AttributeTypes(ctx))
		var syslog = types.ObjectNull(SyslogValue{}.AttributeTypes(ctx))
		var tenants = types.ListNull(types.StringType)

		if v.Action != nil {
			action = types.StringValue(string(*v.Action))
		}
		if v.Antivirus != nil {
			antivirus = avSdkToTerraform(ctx, diags, v.Antivirus)
		}
		if v.Appqoe != nil {
			appqoe = servicePolicyAppQoESdkToTerraform(ctx, diags, v.Appqoe)
		}
		if v.Ewf != nil {
			ewf = servicePolicyEwfSdkToTerraform(ctx, diags, v.Ewf)
		}
		if v.Idp != nil {
			idp = servicePolicyIdpSdkToTerraform(ctx, diags, v.Idp)
		}
		if v.LocalRouting != nil {
			localRouting = types.BoolValue(*v.LocalRouting)
		}
		if v.Name != nil {
			name = types.StringValue(*v.Name)
		}
		if v.PathPreference != nil {
			pathPreference = types.StringValue(*v.PathPreference)
		}
		if v.Services != nil {
			services = mistutils.ListOfStringSdkToTerraform(v.Services)
		}
		if v.ServicepolicyId != nil {
			servicepolicyId = types.StringValue(v.ServicepolicyId.String())
		}
		if v.Skyatp != nil {
			skyatp = skyatpSdkToTerraform(ctx, diags, v.Skyatp)
		}
		if v.SslProxy != nil {
			sslProxy = sslProxySdkToTerraform(ctx, diags, v.SslProxy)
		}
		if v.Syslog != nil {
			syslog = syslogSdkToTerraform(ctx, diags, v.Syslog)
		}
		if v.Tenants != nil {
			tenants = mistutils.ListOfStringSdkToTerraform(v.Tenants)
		}

		dataMapValue := map[string]attr.Value{
			"action":           action,
			"antivirus":        antivirus,
			"appqoe":           appqoe,
			"ewf":              ewf,
			"idp":              idp,
			"local_routing":    localRouting,
			"name":             name,
			"path_preference":  pathPreference,
			"servicepolicy_id": servicepolicyId,
			"services":         services,
			"skyatp":           skyatp,
			"ssl_proxy":        sslProxy,
			"syslog":           syslog,
			"tenants":          tenants,
		}

		data, e := NewServicePoliciesValue(ServicePoliciesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, ServicePoliciesValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}
