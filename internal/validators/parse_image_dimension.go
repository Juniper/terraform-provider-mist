package mistvalidator

import (
	"context"
	"fmt"
	"image"
	"io"
	"os"

	_ "image/jpeg"
	_ "image/png"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = ParseImageDimensionValidator{}

type ParseImageDimensionValidator struct {
	x int
	y int
}

func (o ParseImageDimensionValidator) Description(_ context.Context) string {
	return "Ensures that value submitted by the user contains is a path to a valid Image"
}

func (o ParseImageDimensionValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ParseImageDimensionValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
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

	config, _, err := image.DecodeConfig((io.Reader)(file))
	if err != nil {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"value is not a valid Image file path; Unable to decode image file",
			value,
		))
	}
	if config.Width > o.x {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			fmt.Sprintf("value is not a valid Image file; Image width must be less than %vpx", o.x),
			fmt.Sprintf("%vpx", config.Width),
		))
	}
	if config.Height > o.y {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			fmt.Sprintf("value is not a valid Image file; Image height must be less than %vpx", o.y),
			fmt.Sprintf("%vpx", config.Height),
		))
	}
}

func ParseImageDimension(x int, y int) validator.String {
	return ParseImageDimensionValidator{
		x: x,
		y: y,
	}
}
