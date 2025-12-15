package resource_org_mxedge_inventory

import (
	"context"
	"fmt"
	"strings"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func processMistInventory(
	ctx context.Context,
	diags *diag.Diagnostics,
	data *[]models.Mxedge,
	mistMxedgesByClaimCode *map[string]*MxedgesValue,
	mistMxedgesById *map[string]*MxedgesValue,
	idToMagic map[string]string,
) {
	/*
		Function to process the MxEdge Inventory list retrieved from Mist. This returns the map with all the MxEdges
		in the inventory and generates the maps used in the other functions:
		- mistMxedgesByClaimCode: map to find an MxEdge based on its claim code (if any)
		- mistMxedgesById: map to find an MxEdge based on its ID

		parameters:
			ctx : context.Context
			diags :  *diag.Diagnostics
			data : *[]models.Mxedge
				MxEdge Inventory list retrieved from Mist
			mistMxedgesByClaimCode : *map[string]*MxedgesValue
				map to find an MxEdge based on its claim code (if any). The Key is the MxEdge Claim Code, the value
				is the MxEdge Data
			mistMxedgesById : *map[string]*MxedgesValue
				map to find an MxEdge based on its ID. The Key is the MxEdge ID (UUID), the value
				is the MxEdge Data
	*/

	for _, m := range *data {
		var claimCode = types.StringValue("")
		var forSite basetypes.BoolValue
		var id basetypes.StringValue
		var model basetypes.StringValue
		var name basetypes.StringValue
		var orgId basetypes.StringValue
		var siteId basetypes.StringValue

		if m.Magic != nil {
			claimCode = types.StringValue(*m.Magic)
		}
		if m.ForSite != nil {
			forSite = types.BoolValue(*m.ForSite)
		}
		if m.Id != nil {
			id = types.StringValue(m.Id.String())
		}
		model = types.StringValue(m.Model)
		name = types.StringValue(m.Name)
		if m.OrgId != nil {
			orgId = types.StringValue(m.OrgId.String())
		}
		if m.SiteId != nil {
			siteId = types.StringValue(m.SiteId.String())
		}

		var mac basetypes.StringValue
		if m.Mac != nil {
			mac = types.StringValue(*m.Mac)
		}

		dataMapValue := map[string]attr.Value{
			"for_site":   forSite,
			"id":         id,
			"mac":        mac,
			"claim_code": claimCode,
			"model":      model,
			"name":       name,
			"org_id":     orgId,
			"site_id":    siteId,
		}
		newMxedge, e := NewMxedgesValue(MxedgesValue{}.AttributeTypes(ctx), dataMapValue)
		if e != nil {
			diags.Append(e...)
		} else {
			var nMagic = strings.ToUpper(newMxedge.Magic.ValueString())
			var nId = strings.ToUpper(newMxedge.Id.ValueString())
			(*mistMxedgesById)[nId] = &newMxedge
			if idToMagic != nil {
				if _, ok := idToMagic[nId]; ok {
					fmt.Printf("KDJ Mapping ID %s to Claim Code %s\n", nId, idToMagic[nId])
					nMagic = strings.ToUpper(idToMagic[nId])
				}
			}
			if nMagic != "" {
				(*mistMxedgesByClaimCode)[nMagic] = &newMxedge
			}
		}
	}
}

func processImport(
	ctx context.Context,
	diags *diag.Diagnostics,
	mistMxedgesById *map[string]*MxedgesValue,
) basetypes.MapValue {
	/*
		Function used when using a TF import

		parameters:
			ctx : context.Context
			diags :  *diag.Diagnostics
			mistMxedgesById : *map[string]*MxedgesValue
				map of the MxEdges retrieved in the Mist Inventory

		returns:
			basetypes.MapValue
				map of MxedgesValue to save. Key is the MxEdge Claim Code or the MxEdge ID, Value is a MxedgesValue
				Nested Object with the SiteId and the Read only attributes from Mist
	*/
	newStateMxedgesMap := make(map[string]attr.Value)

	for _, mxedge := range *mistMxedgesById {
		if mxedge.Magic.ValueStringPointer() == nil || len(mxedge.Magic.ValueString()) == 0 {
			newStateMxedgesMap[mxedge.Id.ValueString()] = mxedge
		} else {
			newStateMxedgesMap[mxedge.Magic.ValueString()] = mxedge
		}
	}
	newStateMxedges, e := types.MapValueFrom(ctx, MxedgesValue{}.Type(ctx), newStateMxedgesMap)
	diags.Append(e...)
	return newStateMxedges
}

func processSync(
	ctx context.Context,
	diags *diag.Diagnostics,
	refInventoryMxedges *map[string]*MxedgesValue,
	refPlanMap *map[string]string,
	mistMxedgesByClaimCode *map[string]*MxedgesValue,
	mistMxedgesById *map[string]*MxedgesValue,
) (newStateMxedgesMap basetypes.MapValue) {
	/*
		Function used when syncing state with Mist inventory. Generates the MxEdges map

		parameters:
			ctx : context.Context
			diags :  *diag.Diagnostics
			refInventoryMxedges : *map[string]*MxedgesValue
				map of MxEdges from the reference inventory (plan/state)
			refPlanMap : *map[string]string
				map to get the original key format from the standardized key
			mistMxedgesByClaimCode : *map[string]*MxedgesValue
				map of the MxEdges retrieved from Mist indexed by claim code
			mistMxedgesById : *map[string]*MxedgesValue
				map of the MxEdges retrieved from Mist indexed by ID

		returns:
			basetypes.MapValue
				map of MxedgesValue to save. Key is the MxEdge Claim Code or ID, Value is a MxedgesValue
				Nested Object with the SiteId and attributes from Mist
	*/
	newStateMxedges := make(map[string]attr.Value)

	for mxedgeInfoStandardized := range *refInventoryMxedges {
		mxedgeInfo := (*refPlanMap)[mxedgeInfoStandardized]
		isClaimCode, isUuid := DetectMxedgeInfoType(diags, strings.ToUpper(mxedgeInfo))

		if isClaimCode {
			if mxedgeFromMist, ok := (*mistMxedgesByClaimCode)[strings.ToUpper(mxedgeInfo)]; ok {
				newStateMxedges[mxedgeInfo] = mxedgeFromMist
			}
		} else if isUuid {
			if mxedgeFromMist, ok := (*mistMxedgesById)[strings.ToUpper(mxedgeInfo)]; ok {
				newStateMxedges[mxedgeInfo] = mxedgeFromMist
			}
		} else {
			diags.AddError(
				"MxEdge not found",
				fmt.Sprintf("Invalid Claim Code / MxEdge ID format. Got: \"%s\"", mxedgeInfo),
			)
		}
	}
	newStateMxedgesMap, e := types.MapValueFrom(ctx, MxedgesValue{}.Type(ctx), newStateMxedges)
	diags.Append(e...)
	return newStateMxedgesMap
}

func mapSdkToTerraform(
	ctx context.Context,
	orgId string,
	data *[]models.Mxedge,
	refInventory *OrgMxedgeInventoryModel,
	idToMagic map[string]string,
) (state OrgMxedgeInventoryModel, diags diag.Diagnostics) {
	mistMxedgesByClaimCode := make(map[string]*MxedgesValue)
	mistMxedgesById := make(map[string]*MxedgesValue)

	processMistInventory(ctx, &diags, data, &mistMxedgesByClaimCode, &mistMxedgesById, idToMagic)

	for k, v := range mistMxedgesByClaimCode {
		fmt.Printf("KDJ MxEdge by ClaimCode: %s => %+v\n", k, v)
	}
	for k, v := range mistMxedgesById {
		fmt.Printf("KDJ MxEdge by ID: %s => %+v\n", k, v)
	}

	/*
		If it's for an Import (no refInventory.OrgId), then generate the inventory with:
		- basetypes.StringValue OrgId with the import orgId
		- MapNested Mxedges with the list of MxEdges

		If it's for a Sync (refInventory.OrgId set):
		- basetypes.StringValue OrgId with the refInventory.OrgId
		- MapNested Mxedges with the list of MxEdges in the refInventory and in the Mist Inventory
	*/
	if refInventory.OrgId.ValueStringPointer() == nil {
		fmt.Println("KDJ Processing Import...")
		state.OrgId = types.StringValue(orgId)
		state.Mxedges = processImport(ctx, &diags, &mistMxedgesById)
	} else {
		fmt.Println("KDJ Processing Sync...")
		state.OrgId = refInventory.OrgId
		refInventoryMxedgesMap, refPlanMap := GenMxedgeMap(&refInventory.Mxedges)
		state.Mxedges = processSync(ctx, &diags, &refInventoryMxedgesMap, &refPlanMap, &mistMxedgesByClaimCode, &mistMxedgesById)
	}

	return state, diags
}

func SdkToTerraform(
	ctx context.Context,
	orgId string,
	data *[]models.Mxedge,
	refInventory *OrgMxedgeInventoryModel,
	idToMagic map[string]string,
) (state OrgMxedgeInventoryModel, diags diag.Diagnostics) {
	state, diags = mapSdkToTerraform(ctx, orgId, data, refInventory, idToMagic)
	return state, diags
}
