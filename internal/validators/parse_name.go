package mistvalidator

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = ParseNameValidator{}

type ParseNameValidator struct{}

func (o ParseNameValidator) Description(_ context.Context) string {
	return "Ensures that value submitted by the user contains a Mist Variable"
}

func (o ParseNameValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ParseNameValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	re_variable := `^[a-zA-Z0-9][a-zA-Z0-9_-]*[a-zA-Z0-9]$`

	value := req.ConfigValue.ValueString()
	if is_name, err := regexp.MatchString(re_variable, value); !is_name || err != nil {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"value must only use alphanumerics, underscores, or dashes; start and end with an alphanumeric; include at least one alphabetic character; have no dashes or underscores before the first alphabetic character",
			value,
		))
		return
	}
}

func ParseName() validator.String {
	return ParseNameValidator{}
}
