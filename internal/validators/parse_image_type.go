package mistvalidator

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = ParseImageTypeValidator{}

type ParseImageTypeValidator struct {
	allowPng bool
	allowJpg bool
}

func (o ParseImageTypeValidator) Description(_ context.Context) string {
	return "Ensures that value submitted by the user contains is a path to a valid Image"
}

func (o ParseImageTypeValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ParseImageTypeValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
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
	fileData, err := io.ReadAll(file)
	if err != nil {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"value is not a valid Image file path; Unable to read image file",
			value,
		))
		return
	}

	contentType := http.DetectContentType(fileData)
	if o.allowPng && contentType == "image/png" {
		return
	} else if o.allowJpg && contentType == "image/jpeg" {
		return
	} else {
		var allowed []string
		if o.allowJpg {
			allowed = append(allowed, "`image/jpeg`")
		}
		if o.allowPng {
			allowed = append(allowed, "`image/png`")
		}
		tmp := strings.Join(allowed, ", ")
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			fmt.Sprintf("value %s is not a valid Image file; Image file must be %s", value, tmp),
			contentType,
		))
		return
	}
}

func ParseImageType(allowPng bool, allowJpg bool) validator.String {
	return ParseImageTypeValidator{
		allowPng: allowPng,
		allowJpg: allowJpg,
	}
}
