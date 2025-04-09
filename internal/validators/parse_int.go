package mistvalidator

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = ParseIntValidator{}

type ParseIntValidator struct {
	min int64
	max int64
}

func (o ParseIntValidator) Description(_ context.Context) string {
	return "Ensures that user submitted value is is between the min and max values"
}

func (o ParseIntValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ParseIntValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	strValue := strings.TrimSpace(req.ConfigValue.ValueString())
	intValue, e := strconv.ParseInt(strValue, 10, 64)
	if e != nil || intValue < o.min || intValue > o.max {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			fmt.Sprintf("must be an Integer between %s and %s", strconv.FormatInt(o.min, 10), strconv.FormatInt(o.max, 10)),
			strValue,
		))
		return
	}
}

func ParseInt(min int64, max int64) validator.String {
	return ParseIntValidator{
		min: min,
		max: max,
	}
}
