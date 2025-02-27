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
	_ NineTypesValidator = ForbiddenForbiddenWhenValueIsNotNullValidator{}
)

type ForbiddenForbiddenWhenValueIsNotNullValidator struct {
	expression path.Expression
}

type ForbiddenForbiddenWhenValueIsNotNullRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
}

type ForbiddenForbiddenWhenValueIsNotNullResponse struct {
	Diagnostics diag.Diagnostics
}

func (o ForbiddenForbiddenWhenValueIsNotNullValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Ensures that a value is supplied when attribute %q is not defined", o.expression.String())
}

func (o ForbiddenForbiddenWhenValueIsNotNullValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ForbiddenForbiddenWhenValueIsNotNullValidator) Validate(ctx context.Context, req ForbiddenForbiddenWhenValueIsNotNullRequest, resp *ForbiddenForbiddenWhenValueIsNotNullResponse) {
	// can't proceed while value is unknown
	if req.ConfigValue.IsUnknown() {
		return
	}

	// if we have a value there's no need for further investigation
	if req.ConfigValue.IsNull() {
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

			// Unknown attributes can't satisfy the valueIs condition
			if !mpVal.IsUnknown() && !mpVal.IsNull() {
				resp.Diagnostics.Append(validatordiag.InvalidAttributeCombinationDiagnostic(
					req.Path,
					fmt.Sprintf("attribute %s is forbidden when %s is set, got: %s", req.Path, mp, mpVal.String()),
				))
			}
		}
	}
}

func (o ForbiddenForbiddenWhenValueIsNotNullValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	validateReq := ForbiddenForbiddenWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenForbiddenWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenForbiddenWhenValueIsNotNullValidator) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := ForbiddenForbiddenWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenForbiddenWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenForbiddenWhenValueIsNotNullValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := ForbiddenForbiddenWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenForbiddenWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenForbiddenWhenValueIsNotNullValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := ForbiddenForbiddenWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenForbiddenWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenForbiddenWhenValueIsNotNullValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := ForbiddenForbiddenWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenForbiddenWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenForbiddenWhenValueIsNotNullValidator) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := ForbiddenForbiddenWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenForbiddenWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenForbiddenWhenValueIsNotNullValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	validateReq := ForbiddenForbiddenWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenForbiddenWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenForbiddenWhenValueIsNotNullValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := ForbiddenForbiddenWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenForbiddenWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenForbiddenWhenValueIsNotNullValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := ForbiddenForbiddenWhenValueIsNotNullRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenForbiddenWhenValueIsNotNullResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func ForbiddenForbiddenWhenValueIsNotNull(expression path.Expression) ForbiddenForbiddenWhenValueIsNotNullValidator {
	return ForbiddenForbiddenWhenValueIsNotNullValidator{
		expression: expression,
	}
}
