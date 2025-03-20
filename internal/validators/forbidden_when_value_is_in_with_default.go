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
	_ NineTypesValidator = ForbiddenWhenValueIsInWithDefaultValidator{}
)

type ForbiddenWhenValueIsInWithDefaultValidator struct {
	Expression   path.Expression
	Values       []attr.Value
	DefaultValue attr.Value
}

type ForbiddenWhenValueIsInWithDefaultRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
}

type ForbiddenWhenValueIsInWithDefaultResponse struct {
	Diagnostics diag.Diagnostics
}

func (o ForbiddenWhenValueIsInWithDefaultValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Ensures that no value or the default value is supplied when attribute at %q has value not in %s", o.Expression, o.Values)
}

func (o ForbiddenWhenValueIsInWithDefaultValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ForbiddenWhenValueIsInWithDefaultValidator) Validate(ctx context.Context, req ForbiddenWhenValueIsInWithDefaultRequest, resp *ForbiddenWhenValueIsInWithDefaultResponse) {
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
			forbidden := false
			for _, val := range o.Values {
				if !mpVal.IsUnknown() && val.Equal(mpVal) {
					forbidden = true
				}
			}
			if forbidden {
				resp.Diagnostics.Append(validatordiag.InvalidAttributeCombinationDiagnostic(
					req.Path,
					fmt.Sprintf("Attribute \"%s\" cannot be set when \"%s\" has value in%s, got: %s. Default value \"%s\" is allowed.", req.Path, mp, o.Values, mpVal.String(), o.DefaultValue.String()),
				))
			}
		}
	}
}

func (o ForbiddenWhenValueIsInWithDefaultValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	validateReq := ForbiddenWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenWhenValueIsInWithDefaultValidator) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := ForbiddenWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenWhenValueIsInWithDefaultValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := ForbiddenWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenWhenValueIsInWithDefaultValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := ForbiddenWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenWhenValueIsInWithDefaultValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := ForbiddenWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenWhenValueIsInWithDefaultValidator) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := ForbiddenWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenWhenValueIsInWithDefaultValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	validateReq := ForbiddenWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenWhenValueIsInWithDefaultValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := ForbiddenWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o ForbiddenWhenValueIsInWithDefaultValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := ForbiddenWhenValueIsInWithDefaultRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &ForbiddenWhenValueIsInWithDefaultResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func ForbiddenWhenValueIsInWithDefault(
	expression path.Expression,
	values []attr.Value,
	defaultValue attr.Value,
) ForbiddenWhenValueIsInWithDefaultValidator {
	return ForbiddenWhenValueIsInWithDefaultValidator{
		Expression:   expression,
		Values:       values,
		DefaultValue: defaultValue,
	}
}
