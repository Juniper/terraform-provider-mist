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
	_ NineTypesValidator = MustDifferWhenValueIsValidator{}
)

type MustDifferWhenValueIsValidator struct {
	conditionExpression path.Expression // The field to check for the condition (e.g., "type")
	conditionValue      attr.Value      // The value that triggers the validation (e.g., "external")
	differExpression    path.Expression // The field that must differ (e.g., "local_as")
}

type MustDifferWhenValueIsRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
}

type MustDifferWhenValueIsResponse struct {
	Diagnostics diag.Diagnostics
}

func (o MustDifferWhenValueIsValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Ensures that the value differs from %q when %q has value %q", o.differExpression.String(), o.conditionExpression.String(), o.conditionValue)
}

func (o MustDifferWhenValueIsValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o MustDifferWhenValueIsValidator) Validate(ctx context.Context, req MustDifferWhenValueIsRequest, resp *MustDifferWhenValueIsResponse) {
	// can't proceed while value is unknown
	if req.ConfigValue.IsUnknown() {
		return
	}

	// if we don't have a value, nothing to validate
	if req.ConfigValue.IsNull() {
		return
	}

	// Check the condition field
	mergedConditionExpressions := req.PathExpression.MergeExpressions(o.conditionExpression)

	for _, expression := range mergedConditionExpressions {
		matchedPaths, diags := req.Config.PathMatches(ctx, expression)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		for _, mp := range matchedPaths {
			var conditionVal attr.Value
			resp.Diagnostics.Append(req.Config.GetAttribute(ctx, mp, &conditionVal)...)
			if resp.Diagnostics.HasError() {
				continue
			}

			// Unknown and Null attributes can't satisfy the condition
			if conditionVal.IsNull() || conditionVal.IsUnknown() {
				continue
			}

			// Check if condition is met
			if !conditionVal.Equal(o.conditionValue) {
				continue
			}

			// Condition is met, now check if current value differs from the target field
			mergedDifferExpressions := req.PathExpression.MergeExpressions(o.differExpression)

			for _, differExpr := range mergedDifferExpressions {
				differPaths, diags := req.Config.PathMatches(ctx, differExpr)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				for _, differPath := range differPaths {
					var differVal attr.Value
					resp.Diagnostics.Append(req.Config.GetAttribute(ctx, differPath, &differVal)...)
					if resp.Diagnostics.HasError() {
						continue
					}

					// If differ value is null or unknown, we can't validate
					if differVal.IsNull() || differVal.IsUnknown() {
						continue
					}

					// Check if values differ
					if req.ConfigValue.Equal(differVal) {
						resp.Diagnostics.AddAttributeError(
							req.Path,
							"Invalid Attribute Value Match",
							fmt.Sprintf("Attribute %s must differ from %s (value: %s) when %s is %s",
								req.Path,
								differPath.String(),
								differVal.String(),
								mp.String(),
								o.conditionValue.String(),
							),
						)
					}
				}
			}
		}
	}
}

func (o MustDifferWhenValueIsValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	validateReq := MustDifferWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustDifferWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o MustDifferWhenValueIsValidator) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := MustDifferWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustDifferWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o MustDifferWhenValueIsValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := MustDifferWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustDifferWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o MustDifferWhenValueIsValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := MustDifferWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustDifferWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o MustDifferWhenValueIsValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := MustDifferWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustDifferWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o MustDifferWhenValueIsValidator) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := MustDifferWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustDifferWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o MustDifferWhenValueIsValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	validateReq := MustDifferWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustDifferWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o MustDifferWhenValueIsValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := MustDifferWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustDifferWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o MustDifferWhenValueIsValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := MustDifferWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustDifferWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func MustDifferWhenValueIs(conditionExpression path.Expression, conditionValue attr.Value, differExpression path.Expression) MustDifferWhenValueIsValidator {
	return MustDifferWhenValueIsValidator{
		conditionExpression: conditionExpression,
		conditionValue:      conditionValue,
		differExpression:    differExpression,
	}
}
