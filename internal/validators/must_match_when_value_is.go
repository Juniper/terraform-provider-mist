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
	_ NineTypesValidator = MustMatchWhenValueIsValidator{}
)

type MustMatchWhenValueIsValidator struct {
	conditionExpression path.Expression // The field to check for the condition (e.g., "type")
	conditionValue      attr.Value      // The value that triggers the validation (e.g., "internal")
	matchExpression     path.Expression // The field that must match (e.g., "local_as")
}

type MustMatchWhenValueIsRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
}

type MustMatchWhenValueIsResponse struct {
	Diagnostics diag.Diagnostics
}

func (o MustMatchWhenValueIsValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Ensures that the value matches %q when %q has value %q", o.matchExpression.String(), o.conditionExpression.String(), o.conditionValue)
}

func (o MustMatchWhenValueIsValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o MustMatchWhenValueIsValidator) Validate(ctx context.Context, req MustMatchWhenValueIsRequest, resp *MustMatchWhenValueIsResponse) {
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

			// Condition is met, now check if current value matches the target field
			mergedMatchExpressions := req.PathExpression.MergeExpressions(o.matchExpression)

			for _, matchExpr := range mergedMatchExpressions {
				matchPaths, diags := req.Config.PathMatches(ctx, matchExpr)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				for _, matchPath := range matchPaths {
					var matchVal attr.Value
					resp.Diagnostics.Append(req.Config.GetAttribute(ctx, matchPath, &matchVal)...)
					if resp.Diagnostics.HasError() {
						continue
					}

					// If match value is null or unknown, we can't validate
					if matchVal.IsNull() || matchVal.IsUnknown() {
						continue
					}

					// Check if values match
					if !req.ConfigValue.Equal(matchVal) {
						resp.Diagnostics.AddAttributeError(
							req.Path,
							"Invalid Attribute Value Match",
							fmt.Sprintf("Attribute %s must match %s (value: %s) when %s is %s, got: %s",
								req.Path,
								matchPath.String(),
								matchVal.String(),
								mp.String(),
								o.conditionValue.String(),
								req.ConfigValue.String(),
							),
						)
					}
				}
			}
		}
	}
}

func (o MustMatchWhenValueIsValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	validateReq := MustMatchWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustMatchWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o MustMatchWhenValueIsValidator) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := MustMatchWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustMatchWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o MustMatchWhenValueIsValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := MustMatchWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustMatchWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o MustMatchWhenValueIsValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := MustMatchWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustMatchWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o MustMatchWhenValueIsValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := MustMatchWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustMatchWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o MustMatchWhenValueIsValidator) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := MustMatchWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustMatchWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o MustMatchWhenValueIsValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	validateReq := MustMatchWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustMatchWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o MustMatchWhenValueIsValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := MustMatchWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustMatchWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (o MustMatchWhenValueIsValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := MustMatchWhenValueIsRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}

	validateResp := &MustMatchWhenValueIsResponse{}

	o.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func MustMatchWhenValueIs(conditionExpression path.Expression, conditionValue attr.Value, matchExpression path.Expression) MustMatchWhenValueIsValidator {
	return MustMatchWhenValueIsValidator{
		conditionExpression: conditionExpression,
		conditionValue:      conditionValue,
		matchExpression:     matchExpression,
	}
}
