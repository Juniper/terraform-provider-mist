package mist_transform

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"math/big"
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// /////// STRING
func ListOfStringTerraformToSdk(list basetypes.ListValue) []string {
	var items []string
	for _, item := range list.Elements() {
		var sInterface interface{} = item
		s := sInterface.(basetypes.StringValue)
		items = append(items, s.ValueString())
	}
	return items
}

func ListOfStringSdkToTerraform(data []string) basetypes.ListValue {
	var items []attr.Value
	var itemsType attr.Type = basetypes.StringType{}
	for _, item := range data {
		items = append(items, types.StringValue(item))
	}
	list, _ := types.ListValue(itemsType, items)
	return list
}

func ListOfStringSdkToTerraformEmpty() basetypes.ListValue {
	var items []attr.Value
	var itemsType attr.Type = basetypes.StringType{}
	list, _ := types.ListValue(itemsType, items)
	return list
}

// ///////
func ListOfFloat64SdkToTerraform(data []float64) basetypes.ListValue {
	var items []attr.Value
	var itemsType attr.Type = basetypes.Float64Type{}
	for _, item := range data {
		items = append(items, types.Float64Value(item))
	}
	list, _ := types.ListValue(itemsType, items)
	return list
}

func ListOfNumberSdkToTerraform(data []float64) basetypes.ListValue {
	var items []attr.Value
	var itemsType attr.Type = basetypes.NumberType{}
	for _, item := range data {
		items = append(items, types.NumberValue(big.NewFloat(item)))
	}
	list, _ := types.ListValue(itemsType, items)
	return list
}

// /////// INT
func ListOfIntTerraformToSdk(list basetypes.ListValue) []int {
	var items []int
	for _, item := range list.Elements() {
		var itemInterface interface{} = item
		i := itemInterface.(basetypes.Int64Value)
		items = append(items, int(i.ValueInt64()))
	}
	return items
}

func ListOfIntSdkToTerraform(data []int) basetypes.ListValue {
	var items []attr.Value
	var itemsType attr.Type = basetypes.Int64Type{}
	for _, item := range data {
		items = append(items, types.Int64Value(int64(item)))
	}
	list, _ := types.ListValue(itemsType, items)
	return list
}

func ListOfIntSdkToTerraformEmpty() basetypes.ListValue {
	var items []attr.Value
	var itemsType attr.Type = basetypes.Int64Type{}
	list, _ := types.ListValue(itemsType, items)
	return list
}

// /////// UUID
func ListOfUuidTerraformToSdk(list basetypes.ListValue) []uuid.UUID {
	var items []uuid.UUID
	for _, item := range list.Elements() {
		items = append(items, uuid.MustParse(strings.ReplaceAll(item.String(), "\"", "")))
	}
	return items
}

func ListOfUuidSdkToTerraform(data []uuid.UUID) basetypes.ListValue {
	var items []attr.Value
	var itemsType attr.Type = basetypes.StringType{}
	for _, item := range data {
		items = append(items, types.StringValue(item.String()))
	}
	list, _ := types.ListValue(itemsType, items)
	return list
}
func ListOfUuidSdkToTerraformEmpty() basetypes.ListValue {
	var items []attr.Value
	var itemsType attr.Type = basetypes.StringType{}
	list, _ := types.ListValue(itemsType, items)
	return list
}

// /////// DOT11
func ListODot11TerraformToSdk(list basetypes.ListValue) []models.Dot11BandEnum {
	var items []models.Dot11BandEnum
	for _, item := range list.Elements() {
		var sInterface interface{} = item
		s := sInterface.(basetypes.StringValue)
		items = append(items, (models.Dot11BandEnum)(s.ValueString()))
	}
	return items
}

func ListOfDot11SdkToTerraform(data []models.Dot11BandEnum) basetypes.ListValue {
	var items []attr.Value
	var itemsType attr.Type = basetypes.StringType{}
	for _, item := range data {
		items = append(items, types.StringValue(string(item)))
	}
	list, _ := types.ListValue(itemsType, items)
	return list
}

func ListOfDot11SdkToTerraformEmpty() basetypes.ListValue {
	var items []attr.Value
	var itemsType attr.Type = basetypes.StringType{}
	list, _ := types.ListValue(itemsType, items)
	return list
}
