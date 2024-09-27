package mistvalidator

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

var (
	_ NineTypesValidator = RequiredWhenValueIsNotNullValidator{}
)

type RequiredWhenValueIsNotNullValidator struct {
	expression path.Expression
}

type RequiredWhenValueIsNotNullRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
}

type RequiredWhenValueIsNotNullResponse struct {
	Diagnostics diag.Diagnostics
}

func (o RequiredWhenValueIsNotNullValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Ensures that a value is supplied when attribute %q is not defined", o.expression.String())
}

func (o RequiredWhenValueIsNotNullValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o RequiredWhenValueIsNotNullValidator) Validate(ctx context.Context, req RequiredWhenValueIsNotNullRequest, resp *RequiredWhenValueIsNotNullResponse) {
	// can't proceed while value is unknown
	if req.ConfigValue.IsUnknown() {
		return
	}

	// if we have a value there's no need for further investigation
	if !req.ConfigValue.IsNull() {
		return
	}

	mergedExpressions := req.PathExpression.MergeExpressions(o.expression)

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

			var mpVal attr.Value
			resp.Diagnostics.Append(req.Config.GetAttribute(ctx, mp, &mpVal)...)
			if resp.Diagnostics.HasError() {
				continue // Collect all errors
			}

			// Unknown and Null attributes can't satisfy the valueIs condition
			if !mpVal.IsNull() {
				resp.Diagnostics.AddAttributeError(
					req.Path,
					"Missing required attribute",
					fmt.Sprintf("Attribute %s required when %s is defined.", req.Path, mp.String()),
				)
			}
		}
	}
}

func (o RequiredWhenValueIsNotNullValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	validateReq := RequiredWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &RequiredWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o RequiredWhenValueIsNotNullValidator) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := RequiredWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &RequiredWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o RequiredWhenValueIsNotNullValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := RequiredWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &RequiredWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o RequiredWhenValueIsNotNullValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := RequiredWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &RequiredWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o RequiredWhenValueIsNotNullValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := RequiredWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &RequiredWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o RequiredWhenValueIsNotNullValidator) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := RequiredWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &RequiredWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o RequiredWhenValueIsNotNullValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	validateReq := RequiredWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &RequiredWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o RequiredWhenValueIsNotNullValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := RequiredWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &RequiredWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o RequiredWhenValueIsNotNullValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := RequiredWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &RequiredWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func RequiredWhenValueIsNotNull(expression path.Expression) RequiredWhenValueIsNotNullValidator {
	return RequiredWhenValueIsNotNullValidator{
		expression: expression,
	}
}
