// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_org_usermacs

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

func OrgUsermacsDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"labels": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "Optional, array of strings of labels",
				MarkdownDescription: "Optional, array of strings of labels",
			},
			"mac": schema.StringAttribute{
				Optional:            true,
				Description:         "Partial/full MAC address",
				MarkdownDescription: "Partial/full MAC address",
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"org_usermacs": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:            true,
							Description:         "Unique ID of the object instance in the Mist Organization",
							MarkdownDescription: "Unique ID of the object instance in the Mist Organization",
						},
						"labels": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
						"mac": schema.StringAttribute{
							Computed:            true,
							Description:         "Only non-local-admin MAC is accepted",
							MarkdownDescription: "Only non-local-admin MAC is accepted",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"notes": schema.StringAttribute{
							Computed: true,
						},
						"radius_group": schema.StringAttribute{
							Computed: true,
						},
						"vlan": schema.StringAttribute{
							Computed: true,
						},
					},
					CustomType: OrgUsermacsType{
						ObjectType: types.ObjectType{
							AttrTypes: OrgUsermacsValue{}.AttributeTypes(ctx),
						},
					},
				},
				Computed: true,
			},
		},
	}
}

type OrgUsermacsModel struct {
	Labels      types.List   `tfsdk:"labels"`
	Mac         types.String `tfsdk:"mac"`
	OrgId       types.String `tfsdk:"org_id"`
	OrgUsermacs types.Set    `tfsdk:"org_usermacs"`
}

var _ basetypes.ObjectTypable = OrgUsermacsType{}

type OrgUsermacsType struct {
	basetypes.ObjectType
}

func (t OrgUsermacsType) Equal(o attr.Type) bool {
	other, ok := o.(OrgUsermacsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t OrgUsermacsType) String() string {
	return "OrgUsermacsType"
}

func (t OrgUsermacsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

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

	labelsAttribute, ok := attributes["labels"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`labels is missing from object`)

		return nil, diags
	}

	labelsVal, ok := labelsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`labels expected to be basetypes.ListValue, was: %T`, labelsAttribute))
	}

	macAttribute, ok := attributes["mac"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`mac is missing from object`)

		return nil, diags
	}

	macVal, ok := macAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`mac expected to be basetypes.StringValue, was: %T`, macAttribute))
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

	notesAttribute, ok := attributes["notes"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`notes is missing from object`)

		return nil, diags
	}

	notesVal, ok := notesAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`notes expected to be basetypes.StringValue, was: %T`, notesAttribute))
	}

	radiusGroupAttribute, ok := attributes["radius_group"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`radius_group is missing from object`)

		return nil, diags
	}

	radiusGroupVal, ok := radiusGroupAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`radius_group expected to be basetypes.StringValue, was: %T`, radiusGroupAttribute))
	}

	vlanAttribute, ok := attributes["vlan"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`vlan is missing from object`)

		return nil, diags
	}

	vlanVal, ok := vlanAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`vlan expected to be basetypes.StringValue, was: %T`, vlanAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return OrgUsermacsValue{
		Id:          idVal,
		Labels:      labelsVal,
		Mac:         macVal,
		Name:        nameVal,
		Notes:       notesVal,
		RadiusGroup: radiusGroupVal,
		Vlan:        vlanVal,
		state:       attr.ValueStateKnown,
	}, diags
}

func NewOrgUsermacsValueNull() OrgUsermacsValue {
	return OrgUsermacsValue{
		state: attr.ValueStateNull,
	}
}

func NewOrgUsermacsValueUnknown() OrgUsermacsValue {
	return OrgUsermacsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewOrgUsermacsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (OrgUsermacsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing OrgUsermacsValue Attribute Value",
				"While creating a OrgUsermacsValue value, a missing attribute value was detected. "+
					"A OrgUsermacsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("OrgUsermacsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid OrgUsermacsValue Attribute Type",
				"While creating a OrgUsermacsValue value, an invalid attribute value was detected. "+
					"A OrgUsermacsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("OrgUsermacsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("OrgUsermacsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra OrgUsermacsValue Attribute Value",
				"While creating a OrgUsermacsValue value, an extra attribute value was detected. "+
					"A OrgUsermacsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra OrgUsermacsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewOrgUsermacsValueUnknown(), diags
	}

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return NewOrgUsermacsValueUnknown(), diags
	}

	idVal, ok := idAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be basetypes.StringValue, was: %T`, idAttribute))
	}

	labelsAttribute, ok := attributes["labels"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`labels is missing from object`)

		return NewOrgUsermacsValueUnknown(), diags
	}

	labelsVal, ok := labelsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`labels expected to be basetypes.ListValue, was: %T`, labelsAttribute))
	}

	macAttribute, ok := attributes["mac"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`mac is missing from object`)

		return NewOrgUsermacsValueUnknown(), diags
	}

	macVal, ok := macAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`mac expected to be basetypes.StringValue, was: %T`, macAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return NewOrgUsermacsValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	notesAttribute, ok := attributes["notes"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`notes is missing from object`)

		return NewOrgUsermacsValueUnknown(), diags
	}

	notesVal, ok := notesAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`notes expected to be basetypes.StringValue, was: %T`, notesAttribute))
	}

	radiusGroupAttribute, ok := attributes["radius_group"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`radius_group is missing from object`)

		return NewOrgUsermacsValueUnknown(), diags
	}

	radiusGroupVal, ok := radiusGroupAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`radius_group expected to be basetypes.StringValue, was: %T`, radiusGroupAttribute))
	}

	vlanAttribute, ok := attributes["vlan"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`vlan is missing from object`)

		return NewOrgUsermacsValueUnknown(), diags
	}

	vlanVal, ok := vlanAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`vlan expected to be basetypes.StringValue, was: %T`, vlanAttribute))
	}

	if diags.HasError() {
		return NewOrgUsermacsValueUnknown(), diags
	}

	return OrgUsermacsValue{
		Id:          idVal,
		Labels:      labelsVal,
		Mac:         macVal,
		Name:        nameVal,
		Notes:       notesVal,
		RadiusGroup: radiusGroupVal,
		Vlan:        vlanVal,
		state:       attr.ValueStateKnown,
	}, diags
}

func NewOrgUsermacsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) OrgUsermacsValue {
	object, diags := NewOrgUsermacsValue(attributeTypes, attributes)

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

		panic("NewOrgUsermacsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t OrgUsermacsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewOrgUsermacsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewOrgUsermacsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewOrgUsermacsValueNull(), nil
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

	return NewOrgUsermacsValueMust(OrgUsermacsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t OrgUsermacsType) ValueType(ctx context.Context) attr.Value {
	return OrgUsermacsValue{}
}

var _ basetypes.ObjectValuable = OrgUsermacsValue{}

type OrgUsermacsValue struct {
	Id          basetypes.StringValue `tfsdk:"id"`
	Labels      basetypes.ListValue   `tfsdk:"labels"`
	Mac         basetypes.StringValue `tfsdk:"mac"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Notes       basetypes.StringValue `tfsdk:"notes"`
	RadiusGroup basetypes.StringValue `tfsdk:"radius_group"`
	Vlan        basetypes.StringValue `tfsdk:"vlan"`
	state       attr.ValueState
}

func (v OrgUsermacsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 7)

	var val tftypes.Value
	var err error

	attrTypes["id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["labels"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)
	attrTypes["mac"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["notes"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["radius_group"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["vlan"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 7)

		val, err = v.Id.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["id"] = val

		val, err = v.Labels.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["labels"] = val

		val, err = v.Mac.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["mac"] = val

		val, err = v.Name.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["name"] = val

		val, err = v.Notes.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["notes"] = val

		val, err = v.RadiusGroup.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["radius_group"] = val

		val, err = v.Vlan.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["vlan"] = val

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

func (v OrgUsermacsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v OrgUsermacsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v OrgUsermacsValue) String() string {
	return "OrgUsermacsValue"
}

func (v OrgUsermacsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	labelsVal, d := types.ListValue(types.StringType, v.Labels.Elements())

	diags.Append(d...)

	if d.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"id": basetypes.StringType{},
			"labels": basetypes.ListType{
				ElemType: types.StringType,
			},
			"mac":          basetypes.StringType{},
			"name":         basetypes.StringType{},
			"notes":        basetypes.StringType{},
			"radius_group": basetypes.StringType{},
			"vlan":         basetypes.StringType{},
		}), diags
	}

	attributeTypes := map[string]attr.Type{
		"id": basetypes.StringType{},
		"labels": basetypes.ListType{
			ElemType: types.StringType,
		},
		"mac":          basetypes.StringType{},
		"name":         basetypes.StringType{},
		"notes":        basetypes.StringType{},
		"radius_group": basetypes.StringType{},
		"vlan":         basetypes.StringType{},
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
			"id":           v.Id,
			"labels":       labelsVal,
			"mac":          v.Mac,
			"name":         v.Name,
			"notes":        v.Notes,
			"radius_group": v.RadiusGroup,
			"vlan":         v.Vlan,
		})

	return objVal, diags
}

func (v OrgUsermacsValue) Equal(o attr.Value) bool {
	other, ok := o.(OrgUsermacsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Id.Equal(other.Id) {
		return false
	}

	if !v.Labels.Equal(other.Labels) {
		return false
	}

	if !v.Mac.Equal(other.Mac) {
		return false
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	if !v.Notes.Equal(other.Notes) {
		return false
	}

	if !v.RadiusGroup.Equal(other.RadiusGroup) {
		return false
	}

	if !v.Vlan.Equal(other.Vlan) {
		return false
	}

	return true
}

func (v OrgUsermacsValue) Type(ctx context.Context) attr.Type {
	return OrgUsermacsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v OrgUsermacsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"id": basetypes.StringType{},
		"labels": basetypes.ListType{
			ElemType: types.StringType,
		},
		"mac":          basetypes.StringType{},
		"name":         basetypes.StringType{},
		"notes":        basetypes.StringType{},
		"radius_group": basetypes.StringType{},
		"vlan":         basetypes.StringType{},
	}
}
