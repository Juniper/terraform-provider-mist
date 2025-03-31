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
	_ NineTypesValidator = AllowedWhenValueIsInWithDefaultValidator{}
)

type AllowedWhenValueIsInWithDefaultValidator struct {
	Expression   path.Expression
	Values       []attr.Value
	DefaultValue attr.Value
}

type AllowedWhenValueIsInWithDefaultRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
}

type AllowedWhenValueIsInWithDefaultResponse struct {
	Diagnostics diag.Diagnostics
}

func (o AllowedWhenValueIsInWithDefaultValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Ensures that no value or the default value is supplied when attribute at %q has value not in %s", o.Expression, o.Values)
}

func (o AllowedWhenValueIsInWithDefaultValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o AllowedWhenValueIsInWithDefaultValidator) Validate(ctx context.Context, req AllowedWhenValueIsInWithDefaultRequest, resp *AllowedWhenValueIsInWithDefaultResponse) {
	// can't proceed while value is unknown
	if req.ConfigValue.IsUnknown() {
		return
	}

	// if we don't have a value, there's no need for further investigation
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

			// Unknown attributes can't satisfy the ValueIsIn condition
			allowed := false
			for _, val := range o.Values {
				if !mpVal.IsUnknown() && val.Equal(mpVal) {
					allowed = true
				}
			}
			if !allowed {
				resp.Diagnostics.Append(validatordiag.InvalidAttributeCombinationDiagnostic(
					req.Path,
					fmt.Sprintf("Attribute \"%s\" is only allowed when \"%s\" has value in %s, got: %s. Default value \"%s\" is allowed.", req.Path, mp, o.Values, mpVal.String(), o.DefaultValue.String()),
				))
			}
		}
	}
}

func (o AllowedWhenValueIsInWithDefaultValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	validateReq := AllowedWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AllowedWhenValueIsInWithDefaultValidator) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := AllowedWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AllowedWhenValueIsInWithDefaultValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := AllowedWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AllowedWhenValueIsInWithDefaultValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := AllowedWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AllowedWhenValueIsInWithDefaultValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := AllowedWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AllowedWhenValueIsInWithDefaultValidator) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := AllowedWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AllowedWhenValueIsInWithDefaultValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	validateReq := AllowedWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AllowedWhenValueIsInWithDefaultValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := AllowedWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AllowedWhenValueIsInWithDefaultValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := AllowedWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &AllowedWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func AllowedWhenValueIsInWithDefault(
	expression path.Expression,
	values []attr.Value,
	defaultValue attr.Value,
) AllowedWhenValueIsInWithDefaultValidator {
	return AllowedWhenValueIsInWithDefaultValidator{
		Expression:   expression,
		Values:       values,
		DefaultValue: defaultValue,
	}
}
