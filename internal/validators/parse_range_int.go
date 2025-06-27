package mistvalidator

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = ParseRangeOfIntValidator{}

type ParseRangeOfIntValidator struct {
	min          int
	max          int
	equalAllowed bool
}

func (o ParseRangeOfIntValidator) Description(_ context.Context) string {
	return "Ensures that user submitted value is a range of Integers between the min and max values"
}

func (o ParseRangeOfIntValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ParseRangeOfIntValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	values := strings.Split(req.ConfigValue.ValueString(), "-")

	if len(values) != 2 {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			fmt.Sprintf("must be a range of Integers between %s and %s (e.g. %s-%s)", strconv.Itoa(o.min), strconv.Itoa(o.max), strconv.Itoa(o.min), strconv.Itoa(o.max)),
			req.ConfigValue.ValueString(),
		))
		return
	}

	for _, strValue := range values {
		intValue, e := strconv.Atoi(strValue)
		if e != nil || intValue < o.min || intValue > o.max {
			resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
				req.Path,
				fmt.Sprintf("must be a range of Integers between %s and %s", strconv.Itoa(o.min), strconv.Itoa(o.max)),
				strValue,
			))
			return
		}
	}

	valueOne, _ := strconv.Atoi(values[0])
	valueTwo, _ := strconv.Atoi(values[1])
	if (!o.equalAllowed && valueOne >= valueTwo) || (o.equalAllowed && valueOne > valueTwo) {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			fmt.Sprintf("must be a range of Integers between %s and %s, meaning the first value must be lower than the second value", strconv.Itoa(o.min), strconv.Itoa(o.max)),
			req.ConfigValue.ValueString(),
		))
		return
	}
}

func ParseRangeOfInt(min int, max int, equalAllowed bool) validator.String {
	return ParseRangeOfIntValidator{
		min:          min,
		max:          max,
		equalAllowed: equalAllowed,
	}
}
