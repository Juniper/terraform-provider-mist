package mistvalidator

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

var (
	_ NineTypesValidator = ForbiddenWhenValueIsWithDefaultValidator{}
)

type ForbiddenWhenValueIsWithDefaultValidator struct {
	Expression   path.Expression
	Value        attr.Value
	DefaultValue attr.Value
}

type ForbiddenWhenValueIsWithDefaultRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
}

type ForbiddenWhenValueIsWithDefaultResponse struct {
	Diagnostics diag.Diagnostics
}

func (o ForbiddenWhenValueIsWithDefaultValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Ensures that no value or the default value is supplied when attribute at %q has value %s", o.Expression, o.Value)
}

func (o ForbiddenWhenValueIsWithDefaultValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ForbiddenWhenValueIsWithDefaultValidator) Validate(ctx context.Context, req ForbiddenWhenValueIsWithDefaultRequest, resp *ForbiddenWhenValueIsWithDefaultResponse) {
	// can't proceed while value is unknown
	if req.ConfigValue.IsUnknown() {
		return
	}

	// if we don't have a value there's no need for further investigation
	if req.ConfigValue.IsNull() {
		return
	}

	// if the value matches the allowed default value there's no need for further investigation
	if req.ConfigValue.Equal(o.DefaultValue) {
		return
	}

	mergedExpressions := req.PathExpression.MergeExpressions(o.Expression)

	for _, expression := range mergedExpressions {
		matchedPaths, diags := req.Config.PathMatches(ctx, expression)
		resp.Diagnostics.Append(diags...)
		if diags.HasError() {
			return
		}

		for _, mp := range matchedPaths {
			// If the user specifies the same attribute this mist_validator is applied to,
			// also as part of the input, skip it
			if mp.Equal(req.Path) {
				continue
			}

			// get the attribute we'll be checking against
			var mpVal attr.Value
			resp.Diagnostics.Append(req.Config.GetAttribute(ctx, mp, &mpVal)...)
			if diags.HasError() {
				continue // Collect all errors
			}

			// Unknown attributes can't satisfy the valueIs condition
			if mpVal.IsUnknown() {
				return
			}

			// is the forbidden value found in the matched path?
			if o.Value.Equal(mpVal) {
				resp.Diagnostics.Append(validatordiag.InvalidAttributeCombinationDiagnostic(
					req.Path,
					fmt.Sprintf("Attribute \"%s\" cannot be set when \"%s\" has value %s, got: %s. Default value \"%s\" is allowed.", req.Path, mp, mpVal.String(), req.ConfigValue.String(), o.DefaultValue.String()),
				))
			}
		}
	}
}

func (o ForbiddenWhenValueIsWithDefaultValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	validateReq := ForbiddenWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenWhenValueIsWithDefaultValidator) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := ForbiddenWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenWhenValueIsWithDefaultValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := ForbiddenWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenWhenValueIsWithDefaultValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := ForbiddenWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenWhenValueIsWithDefaultValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := ForbiddenWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenWhenValueIsWithDefaultValidator) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := ForbiddenWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenWhenValueIsWithDefaultValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	validateReq := ForbiddenWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenWhenValueIsWithDefaultValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := ForbiddenWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenWhenValueIsWithDefaultValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := ForbiddenWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func ForbiddenWhenValueIsWithDefault(
	expression path.Expression,
	value attr.Value,
	defaultValue attr.Value,
) ForbiddenWhenValueIsWithDefaultValidator {
	return ForbiddenWhenValueIsWithDefaultValidator{
		Expression:   expression,
		Value:        value,
		DefaultValue: defaultValue,
	}
}
