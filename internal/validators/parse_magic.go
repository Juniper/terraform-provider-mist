package mistvalidator

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = ParseMagicValidator{}

type ParseMagicValidator struct{}

func (o ParseMagicValidator) Description(_ context.Context) string {
	return "Ensures that value submitted by the user contains a Mist Maciable"
}

func (o ParseMagicValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ParseMagicValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	re_variable := `^[0-9A-Z]{15}$`

	value := req.ConfigValue.ValueString()
	if is_valid, err := regexp.MatchString(re_variable, value); !is_valid || err != nil {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"value is not a valid Claim Code; format is \"[0-9A-Z]{15}\" (e.g \"01234ABCDE56789\")",
			value,
		))
		return
	}
}

func ParseMagic() validator.String {
	return ParseMagicValidator{}
}
