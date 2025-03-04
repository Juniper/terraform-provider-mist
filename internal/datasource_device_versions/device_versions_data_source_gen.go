// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_device_versions

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DeviceVersionsDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device_versions": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"model": schema.StringAttribute{
							Computed:            true,
							Description:         "Device model (as seen in the device stats)",
							MarkdownDescription: "Device model (as seen in the device stats)",
						},
						"tag": schema.StringAttribute{
							Computed:            true,
							Description:         "Annotation, stable / beta / alpha. Or it can be empty or nothing which is likely a dev build",
							MarkdownDescription: "Annotation, stable / beta / alpha. Or it can be empty or nothing which is likely a dev build",
						},
						"version": schema.StringAttribute{
							Computed:            true,
							Description:         "Firmware version",
							MarkdownDescription: "Firmware version",
						},
					},
					CustomType: DeviceVersionsType{
						ObjectType: types.ObjectType{
							AttrTypes: DeviceVersionsValue{}.AttributeTypes(ctx),
						},
					},
				},
				Computed: true,
			},
			"model": schema.StringAttribute{
				Required:            true,
				Description:         "Fetch version for device model, use/combine with `type` as needed (for switch and gateway devices)",
				MarkdownDescription: "Fetch version for device model, use/combine with `type` as needed (for switch and gateway devices)",
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Required:            true,
				Description:         "enum: `ap`, `gateway`, `switch`",
				MarkdownDescription: "enum: `ap`, `gateway`, `switch`",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"",
						"ap",
						"gateway",
						"switch",
					),
				},
			},
		},
	}
}

type DeviceVersionsModel struct {
	DeviceVersions types.Set    `tfsdk:"device_versions"`
	Model          types.String `tfsdk:"model"`
	OrgId          types.String `tfsdk:"org_id"`
	Type           types.String `tfsdk:"type"`
}

var _ basetypes.ObjectTypable = DeviceVersionsType{}

type DeviceVersionsType struct {
	basetypes.ObjectType
}

func (t DeviceVersionsType) Equal(o attr.Type) bool {
	other, ok := o.(DeviceVersionsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t DeviceVersionsType) String() string {
	return "DeviceVersionsType"
}

func (t DeviceVersionsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	modelAttribute, ok := attributes["model"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`model is missing from object`)

		return nil, diags
	}

	modelVal, ok := modelAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`model expected to be basetypes.StringValue, was: %T`, modelAttribute))
	}

	tagAttribute, ok := attributes["tag"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`tag is missing from object`)

		return nil, diags
	}

	tagVal, ok := tagAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`tag expected to be basetypes.StringValue, was: %T`, tagAttribute))
	}

	versionAttribute, ok := attributes["version"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`version is missing from object`)

		return nil, diags
	}

	versionVal, ok := versionAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`version expected to be basetypes.StringValue, was: %T`, versionAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return DeviceVersionsValue{
		Model:   modelVal,
		Tag:     tagVal,
		Version: versionVal,
		state:   attr.ValueStateKnown,
	}, diags
}

func NewDeviceVersionsValueNull() DeviceVersionsValue {
	return DeviceVersionsValue{
		state: attr.ValueStateNull,
	}
}

func NewDeviceVersionsValueUnknown() DeviceVersionsValue {
	return DeviceVersionsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewDeviceVersionsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (DeviceVersionsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing DeviceVersionsValue Attribute Value",
				"While creating a DeviceVersionsValue value, a missing attribute value was detected. "+
					"A DeviceVersionsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("DeviceVersionsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid DeviceVersionsValue Attribute Type",
				"While creating a DeviceVersionsValue value, an invalid attribute value was detected. "+
					"A DeviceVersionsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("DeviceVersionsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("DeviceVersionsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra DeviceVersionsValue Attribute Value",
				"While creating a DeviceVersionsValue value, an extra attribute value was detected. "+
					"A DeviceVersionsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra DeviceVersionsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewDeviceVersionsValueUnknown(), diags
	}

	modelAttribute, ok := attributes["model"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`model is missing from object`)

		return NewDeviceVersionsValueUnknown(), diags
	}

	modelVal, ok := modelAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`model expected to be basetypes.StringValue, was: %T`, modelAttribute))
	}

	tagAttribute, ok := attributes["tag"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`tag is missing from object`)

		return NewDeviceVersionsValueUnknown(), diags
	}

	tagVal, ok := tagAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`tag expected to be basetypes.StringValue, was: %T`, tagAttribute))
	}

	versionAttribute, ok := attributes["version"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`version is missing from object`)

		return NewDeviceVersionsValueUnknown(), diags
	}

	versionVal, ok := versionAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`version expected to be basetypes.StringValue, was: %T`, versionAttribute))
	}

	if diags.HasError() {
		return NewDeviceVersionsValueUnknown(), diags
	}

	return DeviceVersionsValue{
		Model:   modelVal,
		Tag:     tagVal,
		Version: versionVal,
		state:   attr.ValueStateKnown,
	}, diags
}

func NewDeviceVersionsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) DeviceVersionsValue {
	object, diags := NewDeviceVersionsValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewDeviceVersionsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t DeviceVersionsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewDeviceVersionsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewDeviceVersionsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewDeviceVersionsValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewDeviceVersionsValueMust(DeviceVersionsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t DeviceVersionsType) ValueType(ctx context.Context) attr.Value {
	return DeviceVersionsValue{}
}

var _ basetypes.ObjectValuable = DeviceVersionsValue{}

type DeviceVersionsValue struct {
	Model   basetypes.StringValue `tfsdk:"model"`
	Tag     basetypes.StringValue `tfsdk:"tag"`
	Version basetypes.StringValue `tfsdk:"version"`
	state   attr.ValueState
}

func (v DeviceVersionsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 3)

	var val tftypes.Value
	var err error

	attrTypes["model"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["tag"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["version"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 3)

		val, err = v.Model.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["model"] = val

		val, err = v.Tag.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["tag"] = val

		val, err = v.Version.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["version"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v DeviceVersionsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v DeviceVersionsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v DeviceVersionsValue) String() string {
	return "DeviceVersionsValue"
}

func (v DeviceVersionsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"model":   basetypes.StringType{},
		"tag":     basetypes.StringType{},
		"version": basetypes.StringType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"model":   v.Model,
			"tag":     v.Tag,
			"version": v.Version,
		})

	return objVal, diags
}

func (v DeviceVersionsValue) Equal(o attr.Value) bool {
	other, ok := o.(DeviceVersionsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Model.Equal(other.Model) {
		return false
	}

	if !v.Tag.Equal(other.Tag) {
		return false
	}

	if !v.Version.Equal(other.Version) {
		return false
	}

	return true
}

func (v DeviceVersionsValue) Type(ctx context.Context) attr.Type {
	return DeviceVersionsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v DeviceVersionsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"model":   basetypes.StringType{},
		"tag":     basetypes.StringType{},
		"version": basetypes.StringType{},
	}
}
