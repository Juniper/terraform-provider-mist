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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ NineTypesValidator = CannotBeTrueWhenValueIsValidator{}
)

type CannotBeTrueWhenValueIsValidator struct {
	Expression path.Expression
	Value      attr.Value
}

type CannotBeTrueWhenValueIsRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
}

type CannotBeTrueWhenValueIsResponse struct {
	Diagnostics diag.Diagnostics
}

func (o CannotBeTrueWhenValueIsValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Ensures that no value is supplied when attribute at %q has value different from %s", o.Expression, o.Value)
}

func (o CannotBeTrueWhenValueIsValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o CannotBeTrueWhenValueIsValidator) Validate(ctx context.Context, req CannotBeTrueWhenValueIsRequest, resp *CannotBeTrueWhenValueIsResponse) {
	// can't proceed while value is unknown
	if req.ConfigValue.IsUnknown() {
		return
	}

	// if we don't have a value, there's no need for further investigation
	if req.ConfigValue.IsNull() {
		return
	}

	// don't continue if the value is false
	var i interface{} = req.ConfigValue
	if !i.(types.Bool).ValueBool() {
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

			// raise error if the value matches
			if o.Value.Equal(mpVal) {
				resp.Diagnostics.Append(validatordiag.InvalidAttributeCombinationDiagnostic(
					req.Path,
					fmt.Sprintf("Attribute \"%s\" cannot be true when \"%s\" has value %s, got: %s", req.Path, mp, o.Value, mpVal.String()),
				))
			}
		}
	}
}

func (o CannotBeTrueWhenValueIsValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	validateReq := CannotBeTrueWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &CannotBeTrueWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o CannotBeTrueWhenValueIsValidator) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := CannotBeTrueWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &CannotBeTrueWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o CannotBeTrueWhenValueIsValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := CannotBeTrueWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &CannotBeTrueWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o CannotBeTrueWhenValueIsValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := CannotBeTrueWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &CannotBeTrueWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o CannotBeTrueWhenValueIsValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := CannotBeTrueWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &CannotBeTrueWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o CannotBeTrueWhenValueIsValidator) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := CannotBeTrueWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &CannotBeTrueWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o CannotBeTrueWhenValueIsValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	validateReq := CannotBeTrueWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &CannotBeTrueWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o CannotBeTrueWhenValueIsValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := CannotBeTrueWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &CannotBeTrueWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o CannotBeTrueWhenValueIsValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := CannotBeTrueWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &CannotBeTrueWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func CannotBeTrueWhenValueIs(expression path.Expression, value attr.Value) CannotBeTrueWhenValueIsValidator {
	return CannotBeTrueWhenValueIsValidator{
		Expression: expression,
		Value:      value,
	}
}
