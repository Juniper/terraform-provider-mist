package mistvalidator

import (
	"context"
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

	value := req.ConfigValue.ValueString()
	if vlan, e := strconv.Atoi(value); e == nil {
		if vlan < o.min || vlan > o.max {
			resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
				req.Path,
				"value must be an Integer between 1 and 4094 or must contain a variable (\"{{...}}\")",
				value,
			))
			return
		}
	}
}

func ParseInt(min int, max int) validator.String {
	return ParseIntValidator{
		min: min,
		max: max,
	}
}
