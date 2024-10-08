// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_const_alarms

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ConstAlarmsDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"const_alarms": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"display": schema.StringAttribute{
							Computed:            true,
							Description:         "Description of the alarm type",
							MarkdownDescription: "Description of the alarm type",
						},
						"group": schema.StringAttribute{
							Computed:            true,
							Description:         "Group to which the alarm belongs",
							MarkdownDescription: "Group to which the alarm belongs",
						},
						"key": schema.StringAttribute{
							Computed:            true,
							Description:         "Key name of the alarm type",
							MarkdownDescription: "Key name of the alarm type",
						},
						"severity": schema.StringAttribute{
							Computed:            true,
							Description:         "Severity of the alarm",
							MarkdownDescription: "Severity of the alarm",
						},
					},
					CustomType: ConstAlarmsType{
						ObjectType: types.ObjectType{
							AttrTypes: ConstAlarmsValue{}.AttributeTypes(ctx),
						},
					},
				},
				Computed: true,
			},
		},
	}
}

type ConstAlarmsModel struct {
	ConstAlarms types.Set `tfsdk:"const_alarms"`
}

var _ basetypes.ObjectTypable = ConstAlarmsType{}

type ConstAlarmsType struct {
	basetypes.ObjectType
}

func (t ConstAlarmsType) Equal(o attr.Type) bool {
	other, ok := o.(ConstAlarmsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t ConstAlarmsType) String() string {
	return "ConstAlarmsType"
}

func (t ConstAlarmsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	displayAttribute, ok := attributes["display"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`display is missing from object`)

		return nil, diags
	}

	displayVal, ok := displayAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`display expected to be basetypes.StringValue, was: %T`, displayAttribute))
	}

	groupAttribute, ok := attributes["group"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`group is missing from object`)

		return nil, diags
	}

	groupVal, ok := groupAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`group expected to be basetypes.StringValue, was: %T`, groupAttribute))
	}

	keyAttribute, ok := attributes["key"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`key is missing from object`)

		return nil, diags
	}

	keyVal, ok := keyAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`key expected to be basetypes.StringValue, was: %T`, keyAttribute))
	}

	severityAttribute, ok := attributes["severity"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`severity is missing from object`)

		return nil, diags
	}

	severityVal, ok := severityAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`severity expected to be basetypes.StringValue, was: %T`, severityAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return ConstAlarmsValue{
		Display:  displayVal,
		Group:    groupVal,
		Key:      keyVal,
		Severity: severityVal,
		state:    attr.ValueStateKnown,
	}, diags
}

func NewConstAlarmsValueNull() ConstAlarmsValue {
	return ConstAlarmsValue{
		state: attr.ValueStateNull,
	}
}

func NewConstAlarmsValueUnknown() ConstAlarmsValue {
	return ConstAlarmsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewConstAlarmsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (ConstAlarmsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing ConstAlarmsValue Attribute Value",
				"While creating a ConstAlarmsValue value, a missing attribute value was detected. "+
					"A ConstAlarmsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ConstAlarmsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid ConstAlarmsValue Attribute Type",
				"While creating a ConstAlarmsValue value, an invalid attribute value was detected. "+
					"A ConstAlarmsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ConstAlarmsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("ConstAlarmsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra ConstAlarmsValue Attribute Value",
				"While creating a ConstAlarmsValue value, an extra attribute value was detected. "+
					"A ConstAlarmsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra ConstAlarmsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewConstAlarmsValueUnknown(), diags
	}

	displayAttribute, ok := attributes["display"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`display is missing from object`)

		return NewConstAlarmsValueUnknown(), diags
	}

	displayVal, ok := displayAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`display expected to be basetypes.StringValue, was: %T`, displayAttribute))
	}

	groupAttribute, ok := attributes["group"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`group is missing from object`)

		return NewConstAlarmsValueUnknown(), diags
	}

	groupVal, ok := groupAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`group expected to be basetypes.StringValue, was: %T`, groupAttribute))
	}

	keyAttribute, ok := attributes["key"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`key is missing from object`)

		return NewConstAlarmsValueUnknown(), diags
	}

	keyVal, ok := keyAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`key expected to be basetypes.StringValue, was: %T`, keyAttribute))
	}

	severityAttribute, ok := attributes["severity"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`severity is missing from object`)

		return NewConstAlarmsValueUnknown(), diags
	}

	severityVal, ok := severityAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`severity expected to be basetypes.StringValue, was: %T`, severityAttribute))
	}

	if diags.HasError() {
		return NewConstAlarmsValueUnknown(), diags
	}

	return ConstAlarmsValue{
		Display:  displayVal,
		Group:    groupVal,
		Key:      keyVal,
		Severity: severityVal,
		state:    attr.ValueStateKnown,
	}, diags
}

func NewConstAlarmsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) ConstAlarmsValue {
	object, diags := NewConstAlarmsValue(attributeTypes, attributes)

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

		panic("NewConstAlarmsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t ConstAlarmsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewConstAlarmsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewConstAlarmsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewConstAlarmsValueNull(), nil
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

	return NewConstAlarmsValueMust(ConstAlarmsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t ConstAlarmsType) ValueType(ctx context.Context) attr.Value {
	return ConstAlarmsValue{}
}

var _ basetypes.ObjectValuable = ConstAlarmsValue{}

type ConstAlarmsValue struct {
	Display  basetypes.StringValue `tfsdk:"display"`
	Group    basetypes.StringValue `tfsdk:"group"`
	Key      basetypes.StringValue `tfsdk:"key"`
	Severity basetypes.StringValue `tfsdk:"severity"`
	state    attr.ValueState
}

func (v ConstAlarmsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 4)

	var val tftypes.Value
	var err error

	attrTypes["display"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["group"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["key"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["severity"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 4)

		val, err = v.Display.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["display"] = val

		val, err = v.Group.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["group"] = val

		val, err = v.Key.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["key"] = val

		val, err = v.Severity.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["severity"] = val

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

func (v ConstAlarmsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v ConstAlarmsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v ConstAlarmsValue) String() string {
	return "ConstAlarmsValue"
}

func (v ConstAlarmsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"display":  basetypes.StringType{},
		"group":    basetypes.StringType{},
		"key":      basetypes.StringType{},
		"severity": basetypes.StringType{},
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
			"display":  v.Display,
			"group":    v.Group,
			"key":      v.Key,
			"severity": v.Severity,
		})

	return objVal, diags
}

func (v ConstAlarmsValue) Equal(o attr.Value) bool {
	other, ok := o.(ConstAlarmsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Display.Equal(other.Display) {
		return false
	}

	if !v.Group.Equal(other.Group) {
		return false
	}

	if !v.Key.Equal(other.Key) {
		return false
	}

	if !v.Severity.Equal(other.Severity) {
		return false
	}

	return true
}

func (v ConstAlarmsValue) Type(ctx context.Context) attr.Type {
	return ConstAlarmsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v ConstAlarmsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"display":  basetypes.StringType{},
		"group":    basetypes.StringType{},
		"key":      basetypes.StringType{},
		"severity": basetypes.StringType{},
	}
}
