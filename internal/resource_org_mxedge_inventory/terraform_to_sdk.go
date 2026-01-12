package resource_org_mxedge_inventory

import (
	"fmt"
	"slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

/*
processAction defines the required action for a specific MxEdge (assign/unassign/nothing)

parameters:

	planSiteId : *basetypes.StringValue
		planned siteId for the MxEdge
	stateSiteId : *basetypes.StringValue
		state siteId for the MxEdge

returns:

	string
		the op to apply to the MxEdge (assign/unassign/nothing)
*/
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

/*
findMxedgeInState finds an MxEdge in the list coming from the Mist Inventory based on the Claim Code
or the MxEdge ID

parameters:

	planMxedgeSiteId : basetypes.StringValue
		the planned MxEdge Site ID
	stateMxedge : *MxedgesValue
		MxEdge from the state

returns:

	string
		the op to apply to the MxEdge (assign/unassign/nothing)
	string
		the MxEdge ID (required for assign/unassign ops)
	bool
		if the MxEdge is already claimed (only used when planMxedgeInfo is a claim code)
*/
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

/*
processPlannedMxedges processes the planned MxEdges and detects which type of action should be applied. Depending
on the required action, the MxEdge will be added to one of the required list

parameters:

	diags: *diag.Diagnostics
	planMxedges : *basetypes.MapValue
		map of MxEdges in the plan. Key is the MxEdge Claim Code or MxEdge ID, Value is a MxedgesValue Nested
		Object with the SiteId and the information retrieved from the Mist Inventory
	stateMxedgesMap : *map[string]*MxedgesValue
		map of MxEdges in the state. Key is the MxEdge Claim Code or MxEdge ID, Value is a MxedgesValue Nested
		Object with the SiteId and the information retrieved from the Mist Inventory
	claim : *[]string
		list of claim codes (string) that must be claimed to the Mist Org
	unassign : *[]string
		list of MxEdge IDs (string) that must be unassigned from Mist Sites
	assignClaim : *map[string]string
		map of ClaimCodes / SiteId of the MxEdges that must be claimed then assigned to a site. This is required
		because we don't have the MxEdge ID at this time (we only have the claim code, the MxEdge ID
		which is required for the "assign" op will be known after the claim)
		the key is the MxEdge Claim Code
		the value is the site id where the MxEdge must be assigned to after the claim
	assign : *map[string][]string
		map of siteId / list of MxEdge IDs (string) that must be assigned to a site
		the key is the siteId where the MxEdge(s) must be assigned to
		the value is a list of MxEdge IDs that must be assigned to the site
*/
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

/*
TerraformToSdk processes the Terraform plan and state to determine what operations need to be performed
on MxEdges in the inventory

returns:

	claim : []string
		list of claim codes that need to be claimed
	unassign : []string
		list of MxEdge IDs that need to be unassigned from sites
	assignClaim : map[string]string
		map of claim codes to site IDs for MxEdges that need to be claimed and assigned
	assign : map[string][]string
		map of site IDs to lists of MxEdge IDs that need to be assigned
	diags : diag.Diagnostics
		any diagnostics/errors encountered during processing
*/
func TerraformToSdk(
	stateInventory *OrgMxedgeInventoryModel,
	planInventory *OrgMxedgeInventoryModel,
) (
	claim []string,
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

	return claim, unassign, assignClaim, assign, diags
}
