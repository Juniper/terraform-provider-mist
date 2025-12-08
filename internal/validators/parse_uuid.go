package mistvalidator

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = ParseUuidValidator{}

type ParseUuidValidator struct{}

func (o ParseUuidValidator) Description(_ context.Context) string {
	return "Ensures that value submitted by the user is a valid UUID"
}

func (o ParseUuidValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ParseUuidValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()
	if _, err := uuid.Parse(value); err != nil {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"value is not a valid UUID; format is \"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx\" (e.g \"550e8400-e29b-41d4-a716-446655440000\")",
			value,
		))
		return
	}
}

func ParseUuid() validator.String {
	return ParseUuidValidator{}
}
