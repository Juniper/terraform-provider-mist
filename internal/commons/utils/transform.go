package mist_transform

import (
	"context"
	"math/big"
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ListOfStringTerraformToSdk(ctx context.Context, list basetypes.ListValue) []string {
	var items []string
	for _, item := range list.Elements() {
		var s_interface interface{} = item
		s := s_interface.(basetypes.StringValue)
		items = append(items, s.ValueString())
	}
	return items
}

func ListOfStringSdkToTerraform(ctx context.Context, data []string) basetypes.ListValue {
	var items []attr.Value
	var items_type attr.Type = basetypes.StringType{}
	for _, item := range data {
		items = append(items, types.StringValue(item))
	}
	list, _ := types.ListValue(items_type, items)
	return list
}

func ListOfStringSdkToTerraformEmpty(ctx context.Context) basetypes.ListValue {
	var items []attr.Value
	var items_type attr.Type = basetypes.StringType{}
	list, _ := types.ListValue(items_type, items)
	return list
}

func ListOfIntTerraformToSdk(ctx context.Context, list basetypes.ListValue) []int {
	var items []int
	for _, item := range list.Elements() {
		var item_interface interface{} = item
		i := item_interface.(int64)
		items = append(items, int(i))
	}
	return items
}

func ListOfFloat64SdkToTerraform(ctx context.Context, data []float64) basetypes.ListValue {
	var items []attr.Value
	var items_type attr.Type = basetypes.Float64Type{}
	for _, item := range data {
		items = append(items, types.Float64Value(float64(item)))
	}
	list, _ := types.ListValue(items_type, items)
	return list
}

func ListOfNumberSdkToTerraform(ctx context.Context, data []float64) basetypes.ListValue {
	var items []attr.Value
	var items_type attr.Type = basetypes.NumberType{}
	for _, item := range data {
		items = append(items, types.NumberValue(big.NewFloat(item)))
	}
	list, _ := types.ListValue(items_type, items)
	return list
}

func ListOfIntSdkToTerraform(ctx context.Context, data []int) basetypes.ListValue {
	var items []attr.Value
	var items_type attr.Type = basetypes.Int64Type{}
	for _, item := range data {
		items = append(items, types.Int64Value(int64(item)))
	}
	list, _ := types.ListValue(items_type, items)
	return list
}

func ListOfIntSdkToTerraformEmpty(ctx context.Context) basetypes.ListValue {
	var items []attr.Value
	var items_type attr.Type = basetypes.Int64Type{}
	list, _ := types.ListValue(items_type, items)
	return list
}

func ListOfUuidTerraformToSdk(ctx context.Context, list basetypes.ListValue) []uuid.UUID {
	var items []uuid.UUID
	for _, item := range list.Elements() {
		items = append(items, uuid.MustParse(strings.ReplaceAll(item.String(), "\"", "")))
	}
	return items
}

func ListOfUuidSdkToTerraform(ctx context.Context, data []uuid.UUID) basetypes.ListValue {
	var items []attr.Value
	var items_type attr.Type = basetypes.StringType{}
	for _, item := range data {
		items = append(items, types.StringValue(item.String()))
	}
	list, _ := types.ListValue(items_type, items)
	return list
}
func ListOfUuidSdkToTerraformEmpty(ctx context.Context) basetypes.ListValue {
	var items []attr.Value
	var items_type attr.Type = basetypes.StringType{}
	list, _ := types.ListValue(items_type, items)
	return list
}
