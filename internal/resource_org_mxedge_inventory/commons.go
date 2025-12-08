package resource_org_mxedge_inventory

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func GenMxedgeMap(mxedges *basetypes.MapValue) (mxedgesMap map[string]*MxedgesValue, planMap map[string]string) {
	/*
		Generate a map[string]*MxedgesValue from the basetypes.MapValue

		parameters:
			mxedges: *basetypes.MapValue
				MapNested with each MxEdge in the Inventory of the Mist Org

		returns:
			map[string]*MxedgesValue
				key is the MxEdge Claim Code or MxEdge ID (UUID), value is the MxedgesValue
			map[string]string
				key is the uppercase version of the key, value is the original key
	*/
	mxedgesMap = make(map[string]*MxedgesValue)
	planMap = make(map[string]string)
	for key, v := range mxedges.Elements() {
		var msi interface{} = v
		var mxedge = msi.(MxedgesValue)
		mxedgesMap[strings.ToUpper(key)] = &mxedge
		planMap[strings.ToUpper(key)] = key
	}
	return mxedgesMap, planMap
}

func DetectMxedgeInfoType(diags *diag.Diagnostics, mxedgeInfo string) (isClaimcode bool, isUuid bool) {
	/*
		Function to detect the type of info (Claim Code or UUID)

		parameters
			diags : *diag.Diagnostics
			mxedgeInfo : string
				the string to test

		returns:
		bool
			true if it's a Claim Code
		bool
			true if it's a UUID

	*/
	reClaimcode1 := `^[0-9]{3}-[0-9]{3}-[0-9]{3}$`
	reClaimcode2 := `^[0-9A-Z]{15}$`
	reUuid := `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	if isValid, _ := regexp.MatchString(reClaimcode1, mxedgeInfo); isValid {
		return true, false
	} else if isValid, _ := regexp.MatchString(reClaimcode2, mxedgeInfo); isValid {
		return true, false
	} else if isValid, _ = regexp.MatchString(reUuid, mxedgeInfo); isValid {
		return false, true
	} else {
		diags.AddError(
			"Invalid MxEdge Key in \"org_mxedge_inventory\" resource",
			fmt.Sprintf("Unable to identify the type of key (claim code / mxedge id) for the MxEdge. got: \"%s\"", mxedgeInfo),
		)
	}
	return false, false
}
