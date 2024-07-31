// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_const_traffic_types

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

func ConstTrafficTypesDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"const_traffic_types": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"display": schema.StringAttribute{
							Computed: true,
						},
						"dscp": schema.Int64Attribute{
							Computed: true,
						},
						"failover_policy": schema.StringAttribute{
							Computed: true,
						},
						"max_jitter": schema.Int64Attribute{
							Computed: true,
						},
						"max_latency": schema.Int64Attribute{
							Computed: true,
						},
						"max_loss": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"traffic_class": schema.StringAttribute{
							Computed: true,
						},
					},
					CustomType: ConstTrafficTypesType{
						ObjectType: types.ObjectType{
							AttrTypes: ConstTrafficTypesValue{}.AttributeTypes(ctx),
						},
					},
				},
				Computed: true,
			},
		},
	}
}

type ConstTrafficTypesModel struct {
	ConstTrafficTypes types.Set `tfsdk:"const_traffic_types"`
}

var _ basetypes.ObjectTypable = ConstTrafficTypesType{}

type ConstTrafficTypesType struct {
	basetypes.ObjectType
}

func (t ConstTrafficTypesType) Equal(o attr.Type) bool {
	other, ok := o.(ConstTrafficTypesType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t ConstTrafficTypesType) String() string {
	return "ConstTrafficTypesType"
}

func (t ConstTrafficTypesType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
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

	dscpAttribute, ok := attributes["dscp"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`dscp is missing from object`)

		return nil, diags
	}

	dscpVal, ok := dscpAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`dscp expected to be basetypes.Int64Value, was: %T`, dscpAttribute))
	}

	failoverPolicyAttribute, ok := attributes["failover_policy"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`failover_policy is missing from object`)

		return nil, diags
	}

	failoverPolicyVal, ok := failoverPolicyAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`failover_policy expected to be basetypes.StringValue, was: %T`, failoverPolicyAttribute))
	}

	maxJitterAttribute, ok := attributes["max_jitter"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`max_jitter is missing from object`)

		return nil, diags
	}

	maxJitterVal, ok := maxJitterAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`max_jitter expected to be basetypes.Int64Value, was: %T`, maxJitterAttribute))
	}

	maxLatencyAttribute, ok := attributes["max_latency"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`max_latency is missing from object`)

		return nil, diags
	}

	maxLatencyVal, ok := maxLatencyAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`max_latency expected to be basetypes.Int64Value, was: %T`, maxLatencyAttribute))
	}

	maxLossAttribute, ok := attributes["max_loss"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`max_loss is missing from object`)

		return nil, diags
	}

	maxLossVal, ok := maxLossAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`max_loss expected to be basetypes.Int64Value, was: %T`, maxLossAttribute))
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

	trafficClassAttribute, ok := attributes["traffic_class"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`traffic_class is missing from object`)

		return nil, diags
	}

	trafficClassVal, ok := trafficClassAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`traffic_class expected to be basetypes.StringValue, was: %T`, trafficClassAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return ConstTrafficTypesValue{
		Display:        displayVal,
		Dscp:           dscpVal,
		FailoverPolicy: failoverPolicyVal,
		MaxJitter:      maxJitterVal,
		MaxLatency:     maxLatencyVal,
		MaxLoss:        maxLossVal,
		Name:           nameVal,
		TrafficClass:   trafficClassVal,
		state:          attr.ValueStateKnown,
	}, diags
}

func NewConstTrafficTypesValueNull() ConstTrafficTypesValue {
	return ConstTrafficTypesValue{
		state: attr.ValueStateNull,
	}
}

func NewConstTrafficTypesValueUnknown() ConstTrafficTypesValue {
	return ConstTrafficTypesValue{
		state: attr.ValueStateUnknown,
	}
}

func NewConstTrafficTypesValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (ConstTrafficTypesValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing ConstTrafficTypesValue Attribute Value",
				"While creating a ConstTrafficTypesValue value, a missing attribute value was detected. "+
					"A ConstTrafficTypesValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ConstTrafficTypesValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid ConstTrafficTypesValue Attribute Type",
				"While creating a ConstTrafficTypesValue value, an invalid attribute value was detected. "+
					"A ConstTrafficTypesValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ConstTrafficTypesValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("ConstTrafficTypesValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra ConstTrafficTypesValue Attribute Value",
				"While creating a ConstTrafficTypesValue value, an extra attribute value was detected. "+
					"A ConstTrafficTypesValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra ConstTrafficTypesValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewConstTrafficTypesValueUnknown(), diags
	}

	displayAttribute, ok := attributes["display"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`display is missing from object`)

		return NewConstTrafficTypesValueUnknown(), diags
	}

	displayVal, ok := displayAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`display expected to be basetypes.StringValue, was: %T`, displayAttribute))
	}

	dscpAttribute, ok := attributes["dscp"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`dscp is missing from object`)

		return NewConstTrafficTypesValueUnknown(), diags
	}

	dscpVal, ok := dscpAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`dscp expected to be basetypes.Int64Value, was: %T`, dscpAttribute))
	}

	failoverPolicyAttribute, ok := attributes["failover_policy"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`failover_policy is missing from object`)

		return NewConstTrafficTypesValueUnknown(), diags
	}

	failoverPolicyVal, ok := failoverPolicyAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`failover_policy expected to be basetypes.StringValue, was: %T`, failoverPolicyAttribute))
	}

	maxJitterAttribute, ok := attributes["max_jitter"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`max_jitter is missing from object`)

		return NewConstTrafficTypesValueUnknown(), diags
	}

	maxJitterVal, ok := maxJitterAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`max_jitter expected to be basetypes.Int64Value, was: %T`, maxJitterAttribute))
	}

	maxLatencyAttribute, ok := attributes["max_latency"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`max_latency is missing from object`)

		return NewConstTrafficTypesValueUnknown(), diags
	}

	maxLatencyVal, ok := maxLatencyAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`max_latency expected to be basetypes.Int64Value, was: %T`, maxLatencyAttribute))
	}

	maxLossAttribute, ok := attributes["max_loss"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`max_loss is missing from object`)

		return NewConstTrafficTypesValueUnknown(), diags
	}

	maxLossVal, ok := maxLossAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`max_loss expected to be basetypes.Int64Value, was: %T`, maxLossAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return NewConstTrafficTypesValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	trafficClassAttribute, ok := attributes["traffic_class"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`traffic_class is missing from object`)

		return NewConstTrafficTypesValueUnknown(), diags
	}

	trafficClassVal, ok := trafficClassAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`traffic_class expected to be basetypes.StringValue, was: %T`, trafficClassAttribute))
	}

	if diags.HasError() {
		return NewConstTrafficTypesValueUnknown(), diags
	}

	return ConstTrafficTypesValue{
		Display:        displayVal,
		Dscp:           dscpVal,
		FailoverPolicy: failoverPolicyVal,
		MaxJitter:      maxJitterVal,
		MaxLatency:     maxLatencyVal,
		MaxLoss:        maxLossVal,
		Name:           nameVal,
		TrafficClass:   trafficClassVal,
		state:          attr.ValueStateKnown,
	}, diags
}

func NewConstTrafficTypesValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) ConstTrafficTypesValue {
	object, diags := NewConstTrafficTypesValue(attributeTypes, attributes)

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

		panic("NewConstTrafficTypesValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t ConstTrafficTypesType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewConstTrafficTypesValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewConstTrafficTypesValueUnknown(), nil
	}

	if in.IsNull() {
		return NewConstTrafficTypesValueNull(), nil
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

	return NewConstTrafficTypesValueMust(ConstTrafficTypesValue{}.AttributeTypes(ctx), attributes), nil
}

func (t ConstTrafficTypesType) ValueType(ctx context.Context) attr.Value {
	return ConstTrafficTypesValue{}
}

var _ basetypes.ObjectValuable = ConstTrafficTypesValue{}

type ConstTrafficTypesValue struct {
	Display        basetypes.StringValue `tfsdk:"display"`
	Dscp           basetypes.Int64Value  `tfsdk:"dscp"`
	FailoverPolicy basetypes.StringValue `tfsdk:"failover_policy"`
	MaxJitter      basetypes.Int64Value  `tfsdk:"max_jitter"`
	MaxLatency     basetypes.Int64Value  `tfsdk:"max_latency"`
	MaxLoss        basetypes.Int64Value  `tfsdk:"max_loss"`
	Name           basetypes.StringValue `tfsdk:"name"`
	TrafficClass   basetypes.StringValue `tfsdk:"traffic_class"`
	state          attr.ValueState
}

func (v ConstTrafficTypesValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 8)

	var val tftypes.Value
	var err error

	attrTypes["display"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["dscp"] = basetypes.Int64Type{}.TerraformType(ctx)
	attrTypes["failover_policy"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["max_jitter"] = basetypes.Int64Type{}.TerraformType(ctx)
	attrTypes["max_latency"] = basetypes.Int64Type{}.TerraformType(ctx)
	attrTypes["max_loss"] = basetypes.Int64Type{}.TerraformType(ctx)
	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["traffic_class"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 8)

		val, err = v.Display.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["display"] = val

		val, err = v.Dscp.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["dscp"] = val

		val, err = v.FailoverPolicy.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["failover_policy"] = val

		val, err = v.MaxJitter.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["max_jitter"] = val

		val, err = v.MaxLatency.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["max_latency"] = val

		val, err = v.MaxLoss.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["max_loss"] = val

		val, err = v.Name.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["name"] = val

		val, err = v.TrafficClass.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["traffic_class"] = val

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

func (v ConstTrafficTypesValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v ConstTrafficTypesValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v ConstTrafficTypesValue) String() string {
	return "ConstTrafficTypesValue"
}

func (v ConstTrafficTypesValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"display":         basetypes.StringType{},
		"dscp":            basetypes.Int64Type{},
		"failover_policy": basetypes.StringType{},
		"max_jitter":      basetypes.Int64Type{},
		"max_latency":     basetypes.Int64Type{},
		"max_loss":        basetypes.Int64Type{},
		"name":            basetypes.StringType{},
		"traffic_class":   basetypes.StringType{},
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
			"display":         v.Display,
			"dscp":            v.Dscp,
			"failover_policy": v.FailoverPolicy,
			"max_jitter":      v.MaxJitter,
			"max_latency":     v.MaxLatency,
			"max_loss":        v.MaxLoss,
			"name":            v.Name,
			"traffic_class":   v.TrafficClass,
		})

	return objVal, diags
}

func (v ConstTrafficTypesValue) Equal(o attr.Value) bool {
	other, ok := o.(ConstTrafficTypesValue)

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

	if !v.Dscp.Equal(other.Dscp) {
		return false
	}

	if !v.FailoverPolicy.Equal(other.FailoverPolicy) {
		return false
	}

	if !v.MaxJitter.Equal(other.MaxJitter) {
		return false
	}

	if !v.MaxLatency.Equal(other.MaxLatency) {
		return false
	}

	if !v.MaxLoss.Equal(other.MaxLoss) {
		return false
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	if !v.TrafficClass.Equal(other.TrafficClass) {
		return false
	}

	return true
}

func (v ConstTrafficTypesValue) Type(ctx context.Context) attr.Type {
	return ConstTrafficTypesType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v ConstTrafficTypesValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"display":         basetypes.StringType{},
		"dscp":            basetypes.Int64Type{},
		"failover_policy": basetypes.StringType{},
		"max_jitter":      basetypes.Int64Type{},
		"max_latency":     basetypes.Int64Type{},
		"max_loss":        basetypes.Int64Type{},
		"name":            basetypes.StringType{},
		"traffic_class":   basetypes.StringType{},
	}
}