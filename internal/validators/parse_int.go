package mistvalidator

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = ParseIntValidator{}

type ParseIntValidator struct {
	min int
	max int
}

func (o ParseIntValidator) Description(_ context.Context) string {
	return "Ensures that user submitted VLAN ID is either a valid VLAN ID string (1-4094) or contains a Variable"
}

func (o ParseIntValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ParseIntValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	str_value := req.ConfigValue.ValueString()
	int_value, e := strconv.Atoi(str_value)
	if e != nil || int_value < o.min || int_value > o.max {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			fmt.Sprintf("value must be an Integer between %s and %s", strconv.Itoa(o.min), strconv.Itoa(o.max)),
			str_value,
		))
		return
	}
}

func ParseInt(min int, max int) validator.String {
	return ParseIntValidator{
		min: min,
		max: max,
	}
}
