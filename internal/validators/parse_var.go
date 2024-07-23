package mistvalidator

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = ParseVarValidator{}

type ParseVarValidator struct{}

func (o ParseVarValidator) Description(_ context.Context) string {
	return "Ensures that value submitted by the user contains a Mist Variable"
}

func (o ParseVarValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ParseVarValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	re_variable := `\{\{\w*\}\}`

	value := req.ConfigValue.ValueString()
	if has_var, err := regexp.MatchString(re_variable, value); !has_var || err != nil {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"value does not contain a Mist variable (\"{{...}}\")",
			value,
		))
		return
	}
}

func ParseVar() validator.String {
	return ParseVarValidator{}
}
