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
	_ NineTypesValidator = AllowedWhenValueIsWithDefaultValidator{}
)

type AllowedWhenValueIsWithDefaultValidator struct {
	Expression   path.Expression
	Value        attr.Value
	DefaultValue attr.Value
}

type AllowedWhenValueIsWithDefaultRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
}

type AllowedWhenValueIsWithDefaultResponse struct {
	Diagnostics diag.Diagnostics
}

func (o AllowedWhenValueIsWithDefaultValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Ensures that no value or the default value is supplied when attribute at %q has value different from %s", o.Expression, o.Value)
}

func (o AllowedWhenValueIsWithDefaultValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o AllowedWhenValueIsWithDefaultValidator) Validate(ctx context.Context, req AllowedWhenValueIsWithDefaultRequest, resp *AllowedWhenValueIsWithDefaultResponse) {
	// can't proceed while value is unknown
	if req.ConfigValue.IsUnknown() {
		return
	}

	// if we don't have a value there's no need for further investigation
	if req.ConfigValue.IsNull() {
		return
	}

	if req.ConfigValue.Equal(o.DefaultValue) {
		return
	}

	mergedExpressions := req.PathExpression.MergeExpressions(o.Expression)

	for _, expression := range mergedExpressions {
		matchedPaths, diags := req.Config.PathMatches(ctx, expression)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
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
			if resp.Diagnostics.HasError() {
				continue // Collect all errors
			}

			// Unknown attributes can't satisfy the valueIs condition
			if mpVal.IsUnknown() || !o.Value.Equal(mpVal) {
				resp.Diagnostics.Append(validatordiag.InvalidAttributeCombinationDiagnostic(
					req.Path,
					fmt.Sprintf("attribute %s is only allowed when %s has value in %s, got: %s. Default value %s is allowed.", req.Path, mp, o.Value, mpVal.String(), o.DefaultValue.String()),
				))
			}
		}
	}
}

func (o AllowedWhenValueIsWithDefaultValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	validateReq := AllowedWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AllowedWhenValueIsWithDefaultValidator) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := AllowedWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AllowedWhenValueIsWithDefaultValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := AllowedWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AllowedWhenValueIsWithDefaultValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := AllowedWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AllowedWhenValueIsWithDefaultValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := AllowedWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AllowedWhenValueIsWithDefaultValidator) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := AllowedWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AllowedWhenValueIsWithDefaultValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	validateReq := AllowedWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AllowedWhenValueIsWithDefaultValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := AllowedWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AllowedWhenValueIsWithDefaultValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := AllowedWhenValueIsWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func AllowedWhenValueIsWithDefault(
	expression path.Expression,
	value attr.Value,
	default_value attr.Value,
) AllowedWhenValueIsWithDefaultValidator {
	return AllowedWhenValueIsWithDefaultValidator{
		Expression:   expression,
		Value:        value,
		DefaultValue: default_value,
	}
}
