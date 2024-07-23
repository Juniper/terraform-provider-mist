package mistvalidator

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = ParseMacValidator{}

type ParseMacValidator struct{}

func (o ParseMacValidator) Description(_ context.Context) string {
	return "Ensures that value submitted by the user contains a Mist Maciable"
}

func (o ParseMacValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ParseMacValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	re_variable := `[0-9a-f]{12}`

	value := req.ConfigValue.ValueString()
	if is_valid, err := regexp.MatchString(re_variable, value); !is_valid || err != nil {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"value is not a valid MAC Address; format is \"5684dae9ac8b\"",
			value,
		))
		return
	}
}

func ParseMac() validator.String {
	return ParseMacValidator{}
}
