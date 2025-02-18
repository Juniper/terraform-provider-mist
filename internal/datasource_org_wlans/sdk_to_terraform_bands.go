package datasource_org_wlans

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bandsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.Dot11BandEnum) basetypes.ListValue {

	var dataList []attr.Value
	for _, d := range l {
		dataList = append(dataList, types.StringValue(string(d)))
	}
	r, e := types.ListValueFrom(ctx, basetypes.StringType{}, dataList)
	diags.Append(e...)

	return r

}
