package resource_org_setting

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func mistNacIdpsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.OrgSettingMistNacIdp) basetypes.ListValue {

	var data_list = []IdpsValue{}
	for _, d := range l {
		var exclude_realms basetypes.ListValue = types.ListNull(types.StringType)
		var id basetypes.StringValue
		var user_realms basetypes.ListValue = types.ListNull(types.StringType)

		if d.ExcludeRealms != nil {
			exclude_realms = mist_transform.ListOfStringSdkToTerraform(ctx, d.ExcludeRealms)
		}
		if d.Id != nil {
			id = types.StringValue(d.Id.String())
		}
		if d.UserRealms != nil {
			user_realms = mist_transform.ListOfStringSdkToTerraform(ctx, d.UserRealms)
		}

		data_map_attr_type := IdpsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"exclude_realms": exclude_realms,
			"id":             id,
			"user_realms":    user_realms,
		}
		data, e := NewIdpsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, IdpsValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

func mistNacServerCertSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingMistNacServerCert) basetypes.ObjectValue {
	var cert basetypes.StringValue
	var key basetypes.StringValue
	var password basetypes.StringValue

	if d.Cert != nil {
		cert = types.StringValue(*d.Cert)
	}
	if d.Key != nil {
		key = types.StringValue(*d.Key)
	}
	if d.Password != nil {
		password = types.StringValue(*d.Password)
	}

	data_map_attr_type := ServerCertValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"cert":     cert,
		"key":      key,
		"password": password,
	}
	r, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)
	return r
}

func mistNacSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingMistNac) MistNacValue {
	var cacerts basetypes.ListValue = types.ListNull(types.StringType)
	var default_idp_id basetypes.StringValue
	var eap_ssl_security_level basetypes.Int64Value
	var eu_only basetypes.BoolValue
	var idps basetypes.ListValue = types.ListNull(IdpsValue{}.Type(ctx))
	var server_cert basetypes.ObjectValue = types.ObjectNull(ServerCertValue{}.AttributeTypes(ctx))
	var use_ip_version basetypes.StringValue
	var use_ssl_port basetypes.BoolValue

	if d.Cacerts != nil {
		cacerts = mist_transform.ListOfStringSdkToTerraform(ctx, d.Cacerts)
	}
	if d.DefaultIdpId != nil {
		default_idp_id = types.StringValue(*d.DefaultIdpId)
	}
	if d.EapSslSecurityLevel != nil {
		eap_ssl_security_level = types.Int64Value(int64(*d.EapSslSecurityLevel))
	}
	if d.EuOnly != nil {
		eu_only = types.BoolValue(*d.EuOnly)
	}
	if d.Idps != nil {
		idps = mistNacIdpsSdkToTerraform(ctx, diags, d.Idps)
	}
	if d.ServerCert != nil {
		server_cert = mistNacServerCertSdkToTerraform(ctx, diags, d.ServerCert)
	}
	if d.UseIpVersion != nil {
		use_ip_version = types.StringValue(string(*d.UseIpVersion))
	}
	if d.UseSslPort != nil {
		use_ssl_port = types.BoolValue(*d.UseSslPort)
	}

	data_map_attr_type := MistNacValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"cacerts":                cacerts,
		"default_idp_id":         default_idp_id,
		"eap_ssl_security_level": eap_ssl_security_level,
		"eu_only":                eu_only,
		"idps":                   idps,
		"server_cert":            server_cert,
		"use_ip_version":         use_ip_version,
		"use_ssl_port":           use_ssl_port,
	}
	data, e := NewMistNacValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
