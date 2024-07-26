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

// This type of validator must satisfy all types.
var (
	_ NineTypesValidator = AtLeastNOfValidator{}
)

// AtLeastNOfValidator is the underlying struct implementing AtLeastNOf.
type AtLeastNOfValidator struct {
	N               int
	PathExpressions path.Expressions
}

type AtLeastNOfValidatorRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	N              int
	Path           path.Path
	PathExpression path.Expression
}

type AtLeastNOfValidatorResponse struct {
	Diagnostics diag.Diagnostics
}

func (o AtLeastNOfValidator) Description(ctx context.Context) string {
	return o.MarkdownDescription(ctx)
}

func (o AtLeastNOfValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("Ensure that at least %d attributes from this collection is set: %s", o.N, o.PathExpressions)
}

func (o AtLeastNOfValidator) Validate(ctx context.Context, req AtLeastNOfValidatorRequest, resp *AtLeastNOfValidatorResponse) {
	expressions := req.PathExpression.MergeExpressions(o.PathExpressions...)

	var notNullPaths []path.Path
	for _, expression := range expressions {
		matchedPaths, diags := req.Config.PathMatches(ctx, expression)

		resp.Diagnostics.Append(diags...)

		// Collect all errors
		if diags.HasError() {
			continue
		}

		for _, mp := range matchedPaths {
			var mpVal attr.Value
			diags := req.Config.GetAttribute(ctx, mp, &mpVal)
			resp.Diagnostics.Append(diags...)

			// Collect all errors
			if diags.HasError() {
				continue
			}

			// Delay validation until all involved attribute have a known value
			if mpVal.IsUnknown() {
				return
			}

			if !mpVal.IsNull() {
				notNullPaths = append(notNullPaths, mp)
			}
		}
	}

	if len(notNullPaths) >= req.N {
		return // this is the desired outcome: more non-null paths than the limit.
	}

	resp.Diagnostics.Append(validatordiag.InvalidAttributeCombinationDiagnostic(
		req.Path,
		fmt.Sprintf("At least %d attributes out of %s should be specified, but %d non-null attributes were found",
			req.N, expressions, len(notNullPaths)),
	))
}

func (o AtLeastNOfValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	validateReq := AtLeastNOfValidatorRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		N:              o.N,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &AtLeastNOfValidatorResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AtLeastNOfValidator) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := AtLeastNOfValidatorRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		N:              o.N,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &AtLeastNOfValidatorResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AtLeastNOfValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := AtLeastNOfValidatorRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		N:              o.N,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &AtLeastNOfValidatorResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AtLeastNOfValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := AtLeastNOfValidatorRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		N:              o.N,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &AtLeastNOfValidatorResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AtLeastNOfValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := AtLeastNOfValidatorRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		N:              o.N,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &AtLeastNOfValidatorResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AtLeastNOfValidator) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := AtLeastNOfValidatorRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		N:              o.N,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &AtLeastNOfValidatorResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AtLeastNOfValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	validateReq := AtLeastNOfValidatorRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		N:              o.N,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &AtLeastNOfValidatorResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AtLeastNOfValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := AtLeastNOfValidatorRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		N:              o.N,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &AtLeastNOfValidatorResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o AtLeastNOfValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := AtLeastNOfValidatorRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		N:              o.N,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &AtLeastNOfValidatorResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

// AtLeastNOf checks that of a set of path.Expression,
// including the attribute this validator is applied to,
// at least 'n' have a non-null value.
//
// Any relative path.Expression will be resolved using the attribute being
// validated.
func AtLeastNOf(n int, expressions ...path.Expression) NineTypesValidator {
	return AtLeastNOfValidator{
		N:               n,
		PathExpressions: expressions,
	}
}
