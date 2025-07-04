// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_org_idpprofile

import (
	"context"
	"fmt"
	"github.com/Juniper/terraform-provider-mist/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func OrgIdpprofileResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_profile": schema.StringAttribute{
				Required:            true,
				Description:         "enum: `critical`, `standard`, `strict`",
				MarkdownDescription: "enum: `critical`, `standard`, `strict`",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"",
						"critical",
						"standard",
						"strict",
					),
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Unique ID of the object instance in the Mist Organization",
				MarkdownDescription: "Unique ID of the object instance in the Mist Organization",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.All(stringvalidator.LengthBetween(2, 32), mistvalidator.ParseName()),
				},
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"overwrites": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "enum:\n  * alert (default)\n  * drop: silently dropping packets\n  * close: notify client/server to close connection",
							MarkdownDescription: "enum:\n  * alert (default)\n  * drop: silently dropping packets\n  * close: notify client/server to close connection",
							Validators: []validator.String{
								stringvalidator.OneOf(
									"",
									"alert",
									"close",
									"drop",
								),
							},
							Default: stringdefault.StaticString("alert"),
						},
						"matching": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"attack_name": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"dst_subnet": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"severity": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
							},
							CustomType: MatchingType{
								ObjectType: types.ObjectType{
									AttrTypes: MatchingValue{}.AttributeTypes(ctx),
								},
							},
							Optional: true,
						},
						"name": schema.StringAttribute{
							Required: true,
						},
					},
					CustomType: OverwritesType{
						ObjectType: types.ObjectType{
							AttrTypes: OverwritesValue{}.AttributeTypes(ctx),
						},
					},
				},
				Optional: true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
		},
	}
}

type OrgIdpprofileModel struct {
	BaseProfile types.String `tfsdk:"base_profile"`
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	OrgId       types.String `tfsdk:"org_id"`
	Overwrites  types.List   `tfsdk:"overwrites"`
}

var _ basetypes.ObjectTypable = OverwritesType{}

type OverwritesType struct {
	basetypes.ObjectType
}

func (t OverwritesType) Equal(o attr.Type) bool {
	other, ok := o.(OverwritesType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t OverwritesType) String() string {
	return "OverwritesType"
}

func (t OverwritesType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	actionAttribute, ok := attributes["action"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`action is missing from object`)

		return nil, diags
	}

	actionVal, ok := actionAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`action expected to be basetypes.StringValue, was: %T`, actionAttribute))
	}

	matchingAttribute, ok := attributes["matching"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`matching is missing from object`)

		return nil, diags
	}

	matchingVal, ok := matchingAttribute.(basetypes.ObjectValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`matching expected to be basetypes.ObjectValue, was: %T`, matchingAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return nil, diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return OverwritesValue{
		Action:   actionVal,
		Matching: matchingVal,
		Name:     nameVal,
		state:    attr.ValueStateKnown,
	}, diags
}

func NewOverwritesValueNull() OverwritesValue {
	return OverwritesValue{
		state: attr.ValueStateNull,
	}
}

func NewOverwritesValueUnknown() OverwritesValue {
	return OverwritesValue{
		state: attr.ValueStateUnknown,
	}
}

func NewOverwritesValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (OverwritesValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing OverwritesValue Attribute Value",
				"While creating a OverwritesValue value, a missing attribute value was detected. "+
					"A OverwritesValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("OverwritesValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid OverwritesValue Attribute Type",
				"While creating a OverwritesValue value, an invalid attribute value was detected. "+
					"A OverwritesValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("OverwritesValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("OverwritesValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra OverwritesValue Attribute Value",
				"While creating a OverwritesValue value, an extra attribute value was detected. "+
					"A OverwritesValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra OverwritesValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewOverwritesValueUnknown(), diags
	}

	actionAttribute, ok := attributes["action"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`action is missing from object`)

		return NewOverwritesValueUnknown(), diags
	}

	actionVal, ok := actionAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`action expected to be basetypes.StringValue, was: %T`, actionAttribute))
	}

	matchingAttribute, ok := attributes["matching"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`matching is missing from object`)

		return NewOverwritesValueUnknown(), diags
	}

	matchingVal, ok := matchingAttribute.(basetypes.ObjectValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`matching expected to be basetypes.ObjectValue, was: %T`, matchingAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return NewOverwritesValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	if diags.HasError() {
		return NewOverwritesValueUnknown(), diags
	}

	return OverwritesValue{
		Action:   actionVal,
		Matching: matchingVal,
		Name:     nameVal,
		state:    attr.ValueStateKnown,
	}, diags
}

func NewOverwritesValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) OverwritesValue {
	object, diags := NewOverwritesValue(attributeTypes, attributes)

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

		panic("NewOverwritesValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t OverwritesType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewOverwritesValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewOverwritesValueUnknown(), nil
	}

	if in.IsNull() {
		return NewOverwritesValueNull(), nil
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

	return NewOverwritesValueMust(OverwritesValue{}.AttributeTypes(ctx), attributes), nil
}

func (t OverwritesType) ValueType(ctx context.Context) attr.Value {
	return OverwritesValue{}
}

var _ basetypes.ObjectValuable = OverwritesValue{}

type OverwritesValue struct {
	Action   basetypes.StringValue `tfsdk:"action"`
	Matching basetypes.ObjectValue `tfsdk:"matching"`
	Name     basetypes.StringValue `tfsdk:"name"`
	state    attr.ValueState
}

func (v OverwritesValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 3)

	var val tftypes.Value
	var err error

	attrTypes["action"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["matching"] = basetypes.ObjectType{
		AttrTypes: MatchingValue{}.AttributeTypes(ctx),
	}.TerraformType(ctx)
	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 3)

		val, err = v.Action.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["action"] = val

		val, err = v.Matching.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["matching"] = val

		val, err = v.Name.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["name"] = val

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

func (v OverwritesValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v OverwritesValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v OverwritesValue) String() string {
	return "OverwritesValue"
}

func (v OverwritesValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var matching basetypes.ObjectValue

	if v.Matching.IsNull() {
		matching = types.ObjectNull(
			MatchingValue{}.AttributeTypes(ctx),
		)
	}

	if v.Matching.IsUnknown() {
		matching = types.ObjectUnknown(
			MatchingValue{}.AttributeTypes(ctx),
		)
	}

	if !v.Matching.IsNull() && !v.Matching.IsUnknown() {
		matching = types.ObjectValueMust(
			MatchingValue{}.AttributeTypes(ctx),
			v.Matching.Attributes(),
		)
	}

	attributeTypes := map[string]attr.Type{
		"action": basetypes.StringType{},
		"matching": basetypes.ObjectType{
			AttrTypes: MatchingValue{}.AttributeTypes(ctx),
		},
		"name": basetypes.StringType{},
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
			"action":   v.Action,
			"matching": matching,
			"name":     v.Name,
		})

	return objVal, diags
}

func (v OverwritesValue) Equal(o attr.Value) bool {
	other, ok := o.(OverwritesValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Action.Equal(other.Action) {
		return false
	}

	if !v.Matching.Equal(other.Matching) {
		return false
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	return true
}

func (v OverwritesValue) Type(ctx context.Context) attr.Type {
	return OverwritesType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v OverwritesValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"action": basetypes.StringType{},
		"matching": basetypes.ObjectType{
			AttrTypes: MatchingValue{}.AttributeTypes(ctx),
		},
		"name": basetypes.StringType{},
	}
}

var _ basetypes.ObjectTypable = MatchingType{}

type MatchingType struct {
	basetypes.ObjectType
}

func (t MatchingType) Equal(o attr.Type) bool {
	other, ok := o.(MatchingType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t MatchingType) String() string {
	return "MatchingType"
}

func (t MatchingType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	attackNameAttribute, ok := attributes["attack_name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`attack_name is missing from object`)

		return nil, diags
	}

	attackNameVal, ok := attackNameAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`attack_name expected to be basetypes.ListValue, was: %T`, attackNameAttribute))
	}

	dstSubnetAttribute, ok := attributes["dst_subnet"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`dst_subnet is missing from object`)

		return nil, diags
	}

	dstSubnetVal, ok := dstSubnetAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`dst_subnet expected to be basetypes.ListValue, was: %T`, dstSubnetAttribute))
	}

	severityAttribute, ok := attributes["severity"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`severity is missing from object`)

		return nil, diags
	}

	severityVal, ok := severityAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`severity expected to be basetypes.ListValue, was: %T`, severityAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return MatchingValue{
		AttackName: attackNameVal,
		DstSubnet:  dstSubnetVal,
		Severity:   severityVal,
		state:      attr.ValueStateKnown,
	}, diags
}

func NewMatchingValueNull() MatchingValue {
	return MatchingValue{
		state: attr.ValueStateNull,
	}
}

func NewMatchingValueUnknown() MatchingValue {
	return MatchingValue{
		state: attr.ValueStateUnknown,
	}
}

func NewMatchingValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (MatchingValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing MatchingValue Attribute Value",
				"While creating a MatchingValue value, a missing attribute value was detected. "+
					"A MatchingValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("MatchingValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid MatchingValue Attribute Type",
				"While creating a MatchingValue value, an invalid attribute value was detected. "+
					"A MatchingValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("MatchingValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("MatchingValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra MatchingValue Attribute Value",
				"While creating a MatchingValue value, an extra attribute value was detected. "+
					"A MatchingValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra MatchingValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewMatchingValueUnknown(), diags
	}

	attackNameAttribute, ok := attributes["attack_name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`attack_name is missing from object`)

		return NewMatchingValueUnknown(), diags
	}

	attackNameVal, ok := attackNameAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`attack_name expected to be basetypes.ListValue, was: %T`, attackNameAttribute))
	}

	dstSubnetAttribute, ok := attributes["dst_subnet"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`dst_subnet is missing from object`)

		return NewMatchingValueUnknown(), diags
	}

	dstSubnetVal, ok := dstSubnetAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`dst_subnet expected to be basetypes.ListValue, was: %T`, dstSubnetAttribute))
	}

	severityAttribute, ok := attributes["severity"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`severity is missing from object`)

		return NewMatchingValueUnknown(), diags
	}

	severityVal, ok := severityAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`severity expected to be basetypes.ListValue, was: %T`, severityAttribute))
	}

	if diags.HasError() {
		return NewMatchingValueUnknown(), diags
	}

	return MatchingValue{
		AttackName: attackNameVal,
		DstSubnet:  dstSubnetVal,
		Severity:   severityVal,
		state:      attr.ValueStateKnown,
	}, diags
}

func NewMatchingValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) MatchingValue {
	object, diags := NewMatchingValue(attributeTypes, attributes)

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

		panic("NewMatchingValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t MatchingType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewMatchingValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewMatchingValueUnknown(), nil
	}

	if in.IsNull() {
		return NewMatchingValueNull(), nil
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

	return NewMatchingValueMust(MatchingValue{}.AttributeTypes(ctx), attributes), nil
}

func (t MatchingType) ValueType(ctx context.Context) attr.Value {
	return MatchingValue{}
}

var _ basetypes.ObjectValuable = MatchingValue{}

type MatchingValue struct {
	AttackName basetypes.ListValue `tfsdk:"attack_name"`
	DstSubnet  basetypes.ListValue `tfsdk:"dst_subnet"`
	Severity   basetypes.ListValue `tfsdk:"severity"`
	state      attr.ValueState
}

func (v MatchingValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 3)

	var val tftypes.Value
	var err error

	attrTypes["attack_name"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)
	attrTypes["dst_subnet"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)
	attrTypes["severity"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 3)

		val, err = v.AttackName.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["attack_name"] = val

		val, err = v.DstSubnet.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["dst_subnet"] = val

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

func (v MatchingValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v MatchingValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v MatchingValue) String() string {
	return "MatchingValue"
}

func (v MatchingValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var attackNameVal basetypes.ListValue
	switch {
	case v.AttackName.IsUnknown():
		attackNameVal = types.ListUnknown(types.StringType)
	case v.AttackName.IsNull():
		attackNameVal = types.ListNull(types.StringType)
	default:
		var d diag.Diagnostics
		attackNameVal, d = types.ListValue(types.StringType, v.AttackName.Elements())
		diags.Append(d...)
	}

	if diags.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"attack_name": basetypes.ListType{
				ElemType: types.StringType,
			},
			"dst_subnet": basetypes.ListType{
				ElemType: types.StringType,
			},
			"severity": basetypes.ListType{
				ElemType: types.StringType,
			},
		}), diags
	}

	var dstSubnetVal basetypes.ListValue
	switch {
	case v.DstSubnet.IsUnknown():
		dstSubnetVal = types.ListUnknown(types.StringType)
	case v.DstSubnet.IsNull():
		dstSubnetVal = types.ListNull(types.StringType)
	default:
		var d diag.Diagnostics
		dstSubnetVal, d = types.ListValue(types.StringType, v.DstSubnet.Elements())
		diags.Append(d...)
	}

	if diags.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"attack_name": basetypes.ListType{
				ElemType: types.StringType,
			},
			"dst_subnet": basetypes.ListType{
				ElemType: types.StringType,
			},
			"severity": basetypes.ListType{
				ElemType: types.StringType,
			},
		}), diags
	}

	var severityVal basetypes.ListValue
	switch {
	case v.Severity.IsUnknown():
		severityVal = types.ListUnknown(types.StringType)
	case v.Severity.IsNull():
		severityVal = types.ListNull(types.StringType)
	default:
		var d diag.Diagnostics
		severityVal, d = types.ListValue(types.StringType, v.Severity.Elements())
		diags.Append(d...)
	}

	if diags.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"attack_name": basetypes.ListType{
				ElemType: types.StringType,
			},
			"dst_subnet": basetypes.ListType{
				ElemType: types.StringType,
			},
			"severity": basetypes.ListType{
				ElemType: types.StringType,
			},
		}), diags
	}

	attributeTypes := map[string]attr.Type{
		"attack_name": basetypes.ListType{
			ElemType: types.StringType,
		},
		"dst_subnet": basetypes.ListType{
			ElemType: types.StringType,
		},
		"severity": basetypes.ListType{
			ElemType: types.StringType,
		},
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
			"attack_name": attackNameVal,
			"dst_subnet":  dstSubnetVal,
			"severity":    severityVal,
		})

	return objVal, diags
}

func (v MatchingValue) Equal(o attr.Value) bool {
	other, ok := o.(MatchingValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.AttackName.Equal(other.AttackName) {
		return false
	}

	if !v.DstSubnet.Equal(other.DstSubnet) {
		return false
	}

	if !v.Severity.Equal(other.Severity) {
		return false
	}

	return true
}

func (v MatchingValue) Type(ctx context.Context) attr.Type {
	return MatchingType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v MatchingValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"attack_name": basetypes.ListType{
			ElemType: types.StringType,
		},
		"dst_subnet": basetypes.ListType{
			ElemType: types.StringType,
		},
		"severity": basetypes.ListType{
			ElemType: types.StringType,
		},
	}
}
