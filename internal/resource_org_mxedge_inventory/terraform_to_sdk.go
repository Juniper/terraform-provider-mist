package resource_org_mxedge_inventory

import (
	"fmt"
	"slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// processAction defines the required action for a specific MxEdge (assign/unassign/nothing)
func processAction(planSiteId *basetypes.StringValue, stateSiteId *basetypes.StringValue) (op string) {
	planSiteIdStr := ""
	if planSiteId != nil && !planSiteId.IsNull() {
		planSiteIdStr = planSiteId.ValueString()
	}

	stateSiteIdStr := ""
	if stateSiteId != nil && !stateSiteId.IsNull() {
		stateSiteIdStr = stateSiteId.ValueString()
	}

	// If both are the same (including both empty/null), no change needed
	if stateSiteIdStr == planSiteIdStr {
		return ""
	}

	// If plan is null/empty but state has a value, unassign
	if (planSiteId == nil || planSiteId.IsNull() || planSiteIdStr == "") && stateSiteIdStr != "" {
		return "unassign"
	}

	// If plan has a value, assign (or reassign if state was different)
	if planSiteIdStr != "" {
		return "assign"
	}

	return ""
}

// findMxedgeInState finds an MxEdge in the list coming from the Mist Inventory based on the Claim Code
// or the MxEdge ID
func findMxedgeInState(
	planMxedgeSiteId *basetypes.StringValue,
	stateMxedge *MxedgesValue,
) (op string, mxedgeId string, alreadyClaimed bool) {
	alreadyClaimed = false
	if stateMxedge != nil && !stateMxedge.IsNull() {
		// for already claimed MxEdges
		op = processAction(planMxedgeSiteId, &stateMxedge.SiteId)
		mxedgeId = stateMxedge.Id.ValueString()
		alreadyClaimed = true
	} else if !planMxedgeSiteId.IsNull() && planMxedgeSiteId.ValueString() != "" {
		// for MxEdges not claimed with the site_id set
		op = "assign"
	}

	return op, mxedgeId, alreadyClaimed
}

// processPlannedMxedges processes the planned MxEdges and detects which type of action should be applied. Depending
// on the required action, the MxEdge will be added to one of the required list
func processPlannedMxedges(
	diags *diag.Diagnostics,
	planMxedges *basetypes.MapValue,
	stateMxedgesMap *map[string]*MxedgesValue,
	claim *[]string,
	unassign *[]string,
	assignClaim *map[string]string,
	assign *map[string][]string,
) {
	for mxedgeInfo, m := range planMxedges.Elements() {
		var op, mxedgeId string
		var alreadyClaimed bool

		var mi interface{} = m
		var planMxedge = mi.(MxedgesValue)
		var mxedgeSiteId = planMxedge.SiteId
		stateMxedge := (*stateMxedgesMap)[strings.ToUpper(mxedgeInfo)]

		// mxedgeId will be empty if the MxEdge is not already in the state
		op, mxedgeId, alreadyClaimed = findMxedgeInState(&mxedgeSiteId, stateMxedge)
		isClaimCode, isUuid := DetectMxedgeInfoType(diags, mxedgeInfo)

		if !alreadyClaimed && isClaimCode {
			*claim = append(*claim, mxedgeInfo)
			if op == "assign" {
				(*assignClaim)[strings.ToUpper(mxedgeInfo)] = mxedgeSiteId.ValueString()
			}
		} else if alreadyClaimed || isUuid {
			if isUuid {
				mxedgeId = mxedgeInfo
				// When using UUID reference without state, determine action based on plan
				if op == "" {
					if mxedgeSiteId.IsNull() || mxedgeSiteId.ValueString() == "" {
						// Plan has no site_id, device should be unassigned
						op = "unassign"
					} else {
						// Plan has site_id, device should be assigned
						op = "assign"
					}
				}
			}
			switch op {
			case "assign":
				if !slices.Contains((*assign)[mxedgeSiteId.ValueString()], mxedgeId) {
					(*assign)[mxedgeSiteId.ValueString()] = append((*assign)[mxedgeSiteId.ValueString()], mxedgeId)
				}
			case "unassign":
				if !slices.Contains(*unassign, mxedgeId) {
					*unassign = append(*unassign, mxedgeId)
				}
			}
		} else if !isClaimCode && !isUuid {
			diags.AddError(
				"Unable to process an MxEdge in \"mist_org_mxedge_inventory\"",
				fmt.Sprintf("Invalid Claim Code / MxEdge ID format. Got: \"%s\"", mxedgeInfo),
			)
		}
	}
}

// processUnplannedMxedges processes the state MxEdges to detect which MxEdges must be deleted
func processUnplannedMxedges(
	planMxedgesMap *map[string]*MxedgesValue,
	stateMxedges *basetypes.MapValue,
	mxedgeIdsToDelete *[]string,
) {
	for mxedgeInfo, m := range stateMxedges.Elements() {
		var mi interface{} = m
		var mxedge = mi.(MxedgesValue)

		// If the MxEdge is not in the plan, it should be deleted
		if _, ok := (*planMxedgesMap)[mxedgeInfo]; !ok {
			if !mxedge.Id.IsNull() && mxedge.Id.ValueString() != "" {
				*mxedgeIdsToDelete = append(*mxedgeIdsToDelete, mxedge.Id.ValueString())
			}
		}
	}
}

// TerraformToSdk processes the Terraform plan and state to determine what operations need to be performed
// on MxEdges in the inventory
func TerraformToSdk(
	stateInventory *OrgMxedgeInventoryModel,
	planInventory *OrgMxedgeInventoryModel,
) (
	claim []string,
	unclaim []string,
	unassign []string,
	assignClaim map[string]string,
	assign map[string][]string,
	diags diag.Diagnostics,
) {
	assignClaim = make(map[string]string)
	assign = make(map[string][]string)

	// process MxEdges in the plan
	// check if MxEdges must be claimed/assigned/unassigned
	stateMxedgesMap, _ := GenMxedgeMap(&stateInventory.Mxedges)
	processPlannedMxedges(&diags, &planInventory.Mxedges, &stateMxedgesMap, &claim, &unassign, &assignClaim, &assign)

	// process MxEdges in the state
	// check if MxEdges must be deleted
	planMxedgesMap, _ := GenMxedgeMap(&planInventory.Mxedges)
	processUnplannedMxedges(&planMxedgesMap, &stateInventory.Mxedges, &unclaim)
	// Note: mxedgeIdsToDelete is not returned here as deletion is handled differently
	// MxEdges are deleted when removed from the map, not through TerraformToSdk

	return claim, unclaim, unassign, assignClaim, assign, diags
}

// DeleteOrgMxedgeInventory processes the state inventory during resource deletion to determine
// which MxEdges should be deleted
func DeleteOrgMxedgeInventory(
	stateInventory *OrgMxedgeInventoryModel,
) (mxedgeIdsToDelete []string, diags diag.Diagnostics) {
	planMxedgesMap := make(map[string]*MxedgesValue)
	processUnplannedMxedges(&planMxedgesMap, &stateInventory.Mxedges, &mxedgeIdsToDelete)

	return mxedgeIdsToDelete, diags
}
