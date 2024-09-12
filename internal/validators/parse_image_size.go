package mistvalidator

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = ParseImageSizeValidator{}

type ParseImageSizeValidator struct {
	maxSize int
}

func (o ParseImageSizeValidator) Description(_ context.Context) string {
	return "Ensures that value submitted by the user contains is a path to a valid Image"
}

func (o ParseImageSizeValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ParseImageSizeValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()
	file, err := os.Open(value)
	if err != nil {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"value is not a valid Image file path; Unable to open image file",
			value,
		))
		return
	}

	defer file.Close()

	state, err := file.Stat()
	if err != nil {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"value is not a valid Image file path; Unable to read image stats",
			value,
		))
		return
	} else if state.Size() > int64(o.maxSize) {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			fmt.Sprintf("value is not a valid Image file; Image file must be less than %vB", o.maxSize),
			fmt.Sprintf("%vB", state.Size()),
		))
	}
}

func ParseImageSize(maxSize int) validator.String {
	return ParseImageSizeValidator{
		maxSize: maxSize,
	}
}
