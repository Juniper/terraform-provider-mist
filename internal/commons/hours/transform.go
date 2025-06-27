package mist_hours

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func HoursSdkToTerraform(diags *diag.Diagnostics, d *models.Hours) basetypes.ObjectValue {

	var set = false
	var mon types.String
	var tue types.String
	var wed types.String
	var thu types.String
	var fri types.String
	var sat types.String
	var sun types.String

	if d.Mon != nil {
		set = true
		mon = types.StringValue(*d.Mon)
	}
	if d.Tue != nil {
		set = true
		tue = types.StringValue(*d.Tue)
	}
	if d.Wed != nil {
		set = true
		wed = types.StringValue(*d.Wed)
	}
	if d.Thu != nil {
		set = true
		thu = types.StringValue(*d.Thu)
	}
	if d.Fri != nil {
		set = true
		fri = types.StringValue(*d.Fri)
	}
	if d.Sat != nil {
		set = true
		sat = types.StringValue(*d.Sat)
	}
	if d.Sun != nil {
		set = true
		sun = types.StringValue(*d.Sun)
	}

	if set {
		rAttrValue := map[string]attr.Value{
			"mon": mon,
			"tue": tue,
			"wed": wed,
			"thu": thu,
			"fri": fri,
			"sat": sat,
			"sun": sun,
		}
		r, e := basetypes.NewObjectValue(HoursValue{}.AttributeTypes(), rAttrValue)
		diags.Append(e...)
		return r
	} else {
		return types.ObjectNull(HoursValue{}.AttributeTypes())
	}
}

func HoursTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.Hours {
	data := models.Hours{}
	if !d.IsNull() && !d.IsUnknown() {
		v, e := NewHoursValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			data = models.Hours{
				Mon: v.Mon.ValueStringPointer(),
				Tue: v.Tue.ValueStringPointer(),
				Wed: v.Wed.ValueStringPointer(),
				Thu: v.Thu.ValueStringPointer(),
				Fri: v.Fri.ValueStringPointer(),
				Sat: v.Sat.ValueStringPointer(),
				Sun: v.Sun.ValueStringPointer(),
			}
		}
	}
	return &data
}

/************************************************************************

DEFINITION

**************************************************************************/

var _ basetypes.ObjectTypable = HoursType{}

type HoursType struct {
	basetypes.ObjectType
}

func (t HoursType) Equal(o attr.Type) bool {
	other, ok := o.(HoursType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t HoursType) String() string {
	return "HoursType"
}

func (t HoursType) ValueFromObject(_ context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	friAttribute, ok := attributes["fri"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`fri is missing from object`)

		return nil, diags
	}

	friVal, ok := friAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`fri expected to be basetypes.StringValue, was: %T`, friAttribute))
	}

	monAttribute, ok := attributes["mon"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`mon is missing from object`)

		return nil, diags
	}

	monVal, ok := monAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`mon expected to be basetypes.StringValue, was: %T`, monAttribute))
	}

	satAttribute, ok := attributes["sat"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`sat is missing from object`)

		return nil, diags
	}

	satVal, ok := satAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`sat expected to be basetypes.StringValue, was: %T`, satAttribute))
	}

	sunAttribute, ok := attributes["sun"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`sun is missing from object`)

		return nil, diags
	}

	sunVal, ok := sunAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`sun expected to be basetypes.StringValue, was: %T`, sunAttribute))
	}

	thuAttribute, ok := attributes["thu"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`thu is missing from object`)

		return nil, diags
	}

	thuVal, ok := thuAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`thu expected to be basetypes.StringValue, was: %T`, thuAttribute))
	}

	tueAttribute, ok := attributes["tue"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`tue is missing from object`)

		return nil, diags
	}

	tueVal, ok := tueAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`tue expected to be basetypes.StringValue, was: %T`, tueAttribute))
	}

	wedAttribute, ok := attributes["wed"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`wed is missing from object`)

		return nil, diags
	}

	wedVal, ok := wedAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`wed expected to be basetypes.StringValue, was: %T`, wedAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return HoursValue{
		Fri:   friVal,
		Mon:   monVal,
		Sat:   satVal,
		Sun:   sunVal,
		Thu:   thuVal,
		Tue:   tueVal,
		Wed:   wedVal,
		state: attr.ValueStateKnown,
	}, diags
}

func NewHoursValueNull() HoursValue {
	return HoursValue{
		state: attr.ValueStateNull,
	}
}

func NewHoursValueUnknown() HoursValue {
	return HoursValue{
		state: attr.ValueStateUnknown,
	}
}

func NewHoursValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (HoursValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing HoursValue Attribute Value",
				"While creating a HoursValue value, a missing attribute value was detected. "+
					"A HoursValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("HoursValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid HoursValue Attribute Type",
				"While creating a HoursValue value, an invalid attribute value was detected. "+
					"A HoursValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("HoursValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("HoursValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra HoursValue Attribute Value",
				"While creating a HoursValue value, an extra attribute value was detected. "+
					"A HoursValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra HoursValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewHoursValueUnknown(), diags
	}

	friAttribute, ok := attributes["fri"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`fri is missing from object`)

		return NewHoursValueUnknown(), diags
	}

	friVal, ok := friAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`fri expected to be basetypes.StringValue, was: %T`, friAttribute))
	}

	monAttribute, ok := attributes["mon"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`mon is missing from object`)

		return NewHoursValueUnknown(), diags
	}

	monVal, ok := monAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`mon expected to be basetypes.StringValue, was: %T`, monAttribute))
	}

	satAttribute, ok := attributes["sat"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`sat is missing from object`)

		return NewHoursValueUnknown(), diags
	}

	satVal, ok := satAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`sat expected to be basetypes.StringValue, was: %T`, satAttribute))
	}

	sunAttribute, ok := attributes["sun"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`sun is missing from object`)

		return NewHoursValueUnknown(), diags
	}

	sunVal, ok := sunAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`sun expected to be basetypes.StringValue, was: %T`, sunAttribute))
	}

	thuAttribute, ok := attributes["thu"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`thu is missing from object`)

		return NewHoursValueUnknown(), diags
	}

	thuVal, ok := thuAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`thu expected to be basetypes.StringValue, was: %T`, thuAttribute))
	}

	tueAttribute, ok := attributes["tue"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`tue is missing from object`)

		return NewHoursValueUnknown(), diags
	}

	tueVal, ok := tueAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`tue expected to be basetypes.StringValue, was: %T`, tueAttribute))
	}

	wedAttribute, ok := attributes["wed"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`wed is missing from object`)

		return NewHoursValueUnknown(), diags
	}

	wedVal, ok := wedAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`wed expected to be basetypes.StringValue, was: %T`, wedAttribute))
	}

	if diags.HasError() {
		return NewHoursValueUnknown(), diags
	}

	return HoursValue{
		Fri:   friVal,
		Mon:   monVal,
		Sat:   satVal,
		Sun:   sunVal,
		Thu:   thuVal,
		Tue:   tueVal,
		Wed:   wedVal,
		state: attr.ValueStateKnown,
	}, diags
}

func NewHoursValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) HoursValue {
	object, diags := NewHoursValue(attributeTypes, attributes)

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

		panic("NewHoursValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t HoursType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewHoursValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewHoursValueUnknown(), nil
	}

	if in.IsNull() {
		return NewHoursValueNull(), nil
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

	return NewHoursValueMust(HoursValue{}.AttributeTypes(), attributes), nil
}

func (t HoursType) ValueType(_ context.Context) attr.Value {
	return HoursValue{}
}

var _ basetypes.ObjectValuable = HoursValue{}

type HoursValue struct {
	Fri   basetypes.StringValue `tfsdk:"fri"`
	Mon   basetypes.StringValue `tfsdk:"mon"`
	Sat   basetypes.StringValue `tfsdk:"sat"`
	Sun   basetypes.StringValue `tfsdk:"sun"`
	Thu   basetypes.StringValue `tfsdk:"thu"`
	Tue   basetypes.StringValue `tfsdk:"tue"`
	Wed   basetypes.StringValue `tfsdk:"wed"`
	state attr.ValueState
}

func (v HoursValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 7)

	var val tftypes.Value
	var err error

	attrTypes["fri"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["mon"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["sat"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["sun"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["thu"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["tue"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["wed"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 7)

		val, err = v.Fri.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["fri"] = val

		val, err = v.Mon.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["mon"] = val

		val, err = v.Sat.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["sat"] = val

		val, err = v.Sun.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["sun"] = val

		val, err = v.Thu.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["thu"] = val

		val, err = v.Tue.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["tue"] = val

		val, err = v.Wed.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["wed"] = val

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

func (v HoursValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v HoursValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v HoursValue) String() string {
	return "HoursValue"
}

func (v HoursValue) ToObjectValue(_ context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"fri": basetypes.StringType{},
		"mon": basetypes.StringType{},
		"sat": basetypes.StringType{},
		"sun": basetypes.StringType{},
		"thu": basetypes.StringType{},
		"tue": basetypes.StringType{},
		"wed": basetypes.StringType{},
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
			"fri": v.Fri,
			"mon": v.Mon,
			"sat": v.Sat,
			"sun": v.Sun,
			"thu": v.Thu,
			"tue": v.Tue,
			"wed": v.Wed,
		})

	return objVal, diags
}

func (v HoursValue) Equal(o attr.Value) bool {
	other, ok := o.(HoursValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Fri.Equal(other.Fri) {
		return false
	}

	if !v.Mon.Equal(other.Mon) {
		return false
	}

	if !v.Sat.Equal(other.Sat) {
		return false
	}

	if !v.Sun.Equal(other.Sun) {
		return false
	}

	if !v.Thu.Equal(other.Thu) {
		return false
	}

	if !v.Tue.Equal(other.Tue) {
		return false
	}

	if !v.Wed.Equal(other.Wed) {
		return false
	}

	return true
}

func (v HoursValue) Type(context.Context) attr.Type {
	return HoursType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(),
		},
	}
}

func (v HoursValue) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"fri": basetypes.StringType{},
		"mon": basetypes.StringType{},
		"sat": basetypes.StringType{},
		"sun": basetypes.StringType{},
		"thu": basetypes.StringType{},
		"tue": basetypes.StringType{},
		"wed": basetypes.StringType{},
	}
}
