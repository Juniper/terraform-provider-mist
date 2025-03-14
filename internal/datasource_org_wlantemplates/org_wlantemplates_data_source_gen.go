// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_org_wlantemplates

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

func OrgWlantemplatesDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"org_wlantemplates": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_time": schema.Float64Attribute{
							Computed:            true,
							Description:         "When the object has been created, in epoch",
							MarkdownDescription: "When the object has been created, in epoch",
						},
						"id": schema.StringAttribute{
							Computed:            true,
							Description:         "Unique ID of the object instance in the Mist Organization",
							MarkdownDescription: "Unique ID of the object instance in the Mist Organization",
						},
						"modified_time": schema.Float64Attribute{
							Computed:            true,
							Description:         "When the object has been modified for the last time, in epoch",
							MarkdownDescription: "When the object has been modified for the last time, in epoch",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"org_id": schema.StringAttribute{
							Computed: true,
						},
					},
					CustomType: OrgWlantemplatesType{
						ObjectType: types.ObjectType{
							AttrTypes: OrgWlantemplatesValue{}.AttributeTypes(ctx),
						},
					},
				},
				Computed: true,
			},
		},
	}
}

type OrgWlantemplatesModel struct {
	OrgId            types.String `tfsdk:"org_id"`
	OrgWlantemplates types.Set    `tfsdk:"org_wlantemplates"`
}

var _ basetypes.ObjectTypable = OrgWlantemplatesType{}

type OrgWlantemplatesType struct {
	basetypes.ObjectType
}

func (t OrgWlantemplatesType) Equal(o attr.Type) bool {
	other, ok := o.(OrgWlantemplatesType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t OrgWlantemplatesType) String() string {
	return "OrgWlantemplatesType"
}

func (t OrgWlantemplatesType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	createdTimeAttribute, ok := attributes["created_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`created_time is missing from object`)

		return nil, diags
	}

	createdTimeVal, ok := createdTimeAttribute.(basetypes.Float64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`created_time expected to be basetypes.Float64Value, was: %T`, createdTimeAttribute))
	}

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return nil, diags
	}

	idVal, ok := idAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be basetypes.StringValue, was: %T`, idAttribute))
	}

	modifiedTimeAttribute, ok := attributes["modified_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`modified_time is missing from object`)

		return nil, diags
	}

	modifiedTimeVal, ok := modifiedTimeAttribute.(basetypes.Float64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`modified_time expected to be basetypes.Float64Value, was: %T`, modifiedTimeAttribute))
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

	orgIdAttribute, ok := attributes["org_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`org_id is missing from object`)

		return nil, diags
	}

	orgIdVal, ok := orgIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`org_id expected to be basetypes.StringValue, was: %T`, orgIdAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return OrgWlantemplatesValue{
		CreatedTime:  createdTimeVal,
		Id:           idVal,
		ModifiedTime: modifiedTimeVal,
		Name:         nameVal,
		OrgId:        orgIdVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewOrgWlantemplatesValueNull() OrgWlantemplatesValue {
	return OrgWlantemplatesValue{
		state: attr.ValueStateNull,
	}
}

func NewOrgWlantemplatesValueUnknown() OrgWlantemplatesValue {
	return OrgWlantemplatesValue{
		state: attr.ValueStateUnknown,
	}
}

func NewOrgWlantemplatesValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (OrgWlantemplatesValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing OrgWlantemplatesValue Attribute Value",
				"While creating a OrgWlantemplatesValue value, a missing attribute value was detected. "+
					"A OrgWlantemplatesValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("OrgWlantemplatesValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid OrgWlantemplatesValue Attribute Type",
				"While creating a OrgWlantemplatesValue value, an invalid attribute value was detected. "+
					"A OrgWlantemplatesValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("OrgWlantemplatesValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("OrgWlantemplatesValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra OrgWlantemplatesValue Attribute Value",
				"While creating a OrgWlantemplatesValue value, an extra attribute value was detected. "+
					"A OrgWlantemplatesValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra OrgWlantemplatesValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewOrgWlantemplatesValueUnknown(), diags
	}

	createdTimeAttribute, ok := attributes["created_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`created_time is missing from object`)

		return NewOrgWlantemplatesValueUnknown(), diags
	}

	createdTimeVal, ok := createdTimeAttribute.(basetypes.Float64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`created_time expected to be basetypes.Float64Value, was: %T`, createdTimeAttribute))
	}

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return NewOrgWlantemplatesValueUnknown(), diags
	}

	idVal, ok := idAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be basetypes.StringValue, was: %T`, idAttribute))
	}

	modifiedTimeAttribute, ok := attributes["modified_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`modified_time is missing from object`)

		return NewOrgWlantemplatesValueUnknown(), diags
	}

	modifiedTimeVal, ok := modifiedTimeAttribute.(basetypes.Float64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`modified_time expected to be basetypes.Float64Value, was: %T`, modifiedTimeAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return NewOrgWlantemplatesValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	orgIdAttribute, ok := attributes["org_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`org_id is missing from object`)

		return NewOrgWlantemplatesValueUnknown(), diags
	}

	orgIdVal, ok := orgIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`org_id expected to be basetypes.StringValue, was: %T`, orgIdAttribute))
	}

	if diags.HasError() {
		return NewOrgWlantemplatesValueUnknown(), diags
	}

	return OrgWlantemplatesValue{
		CreatedTime:  createdTimeVal,
		Id:           idVal,
		ModifiedTime: modifiedTimeVal,
		Name:         nameVal,
		OrgId:        orgIdVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewOrgWlantemplatesValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) OrgWlantemplatesValue {
	object, diags := NewOrgWlantemplatesValue(attributeTypes, attributes)

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

		panic("NewOrgWlantemplatesValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t OrgWlantemplatesType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewOrgWlantemplatesValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewOrgWlantemplatesValueUnknown(), nil
	}

	if in.IsNull() {
		return NewOrgWlantemplatesValueNull(), nil
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

	return NewOrgWlantemplatesValueMust(OrgWlantemplatesValue{}.AttributeTypes(ctx), attributes), nil
}

func (t OrgWlantemplatesType) ValueType(ctx context.Context) attr.Value {
	return OrgWlantemplatesValue{}
}

var _ basetypes.ObjectValuable = OrgWlantemplatesValue{}

type OrgWlantemplatesValue struct {
	CreatedTime  basetypes.Float64Value `tfsdk:"created_time"`
	Id           basetypes.StringValue  `tfsdk:"id"`
	ModifiedTime basetypes.Float64Value `tfsdk:"modified_time"`
	Name         basetypes.StringValue  `tfsdk:"name"`
	OrgId        basetypes.StringValue  `tfsdk:"org_id"`
	state        attr.ValueState
}

func (v OrgWlantemplatesValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 5)

	var val tftypes.Value
	var err error

	attrTypes["created_time"] = basetypes.Float64Type{}.TerraformType(ctx)
	attrTypes["id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["modified_time"] = basetypes.Float64Type{}.TerraformType(ctx)
	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["org_id"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 5)

		val, err = v.CreatedTime.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["created_time"] = val

		val, err = v.Id.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["id"] = val

		val, err = v.ModifiedTime.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["modified_time"] = val

		val, err = v.Name.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["name"] = val

		val, err = v.OrgId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["org_id"] = val

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

func (v OrgWlantemplatesValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v OrgWlantemplatesValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v OrgWlantemplatesValue) String() string {
	return "OrgWlantemplatesValue"
}

func (v OrgWlantemplatesValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"created_time":  basetypes.Float64Type{},
		"id":            basetypes.StringType{},
		"modified_time": basetypes.Float64Type{},
		"name":          basetypes.StringType{},
		"org_id":        basetypes.StringType{},
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
			"created_time":  v.CreatedTime,
			"id":            v.Id,
			"modified_time": v.ModifiedTime,
			"name":          v.Name,
			"org_id":        v.OrgId,
		})

	return objVal, diags
}

func (v OrgWlantemplatesValue) Equal(o attr.Value) bool {
	other, ok := o.(OrgWlantemplatesValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.CreatedTime.Equal(other.CreatedTime) {
		return false
	}

	if !v.Id.Equal(other.Id) {
		return false
	}

	if !v.ModifiedTime.Equal(other.ModifiedTime) {
		return false
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	if !v.OrgId.Equal(other.OrgId) {
		return false
	}

	return true
}

func (v OrgWlantemplatesValue) Type(ctx context.Context) attr.Type {
	return OrgWlantemplatesType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v OrgWlantemplatesValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"created_time":  basetypes.Float64Type{},
		"id":            basetypes.StringType{},
		"modified_time": basetypes.Float64Type{},
		"name":          basetypes.StringType{},
		"org_id":        basetypes.StringType{},
	}
}
