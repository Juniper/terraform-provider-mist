package mist_api

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func VlanAsString(vlanId models.VlanIdWithVariable) basetypes.StringValue {
	if v, ok := vlanId.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := vlanId.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}