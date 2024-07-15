package resource_org_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bandsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.Dot11BandEnum) basetypes.ListValue {

	var data_list []attr.Value
	for _, d := range l {
		data_list = append(data_list, types.StringValue(string(d)))
	}
	r, e := types.ListValueFrom(ctx, basetypes.StringType{}, data_list)
	diags.Append(e...)

	return r

}
