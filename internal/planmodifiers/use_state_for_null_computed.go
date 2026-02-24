package mistplanmodifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// UseStateForNullComputedBool returns a plan modifier that copies the prior state value
// to the planned value when the computed value is null. This is specifically designed
// for backwards compatibility when the API stops returning certain fields.
//
// This modifier differs from UseStateForUnknown in that it:
// - Only activates when the computed value is null (not unknown)
// - Preserves the prior state value instead of replacing with null
// - Does not activate when there's no prior state (new resource)
//
// Use case: API fields that are being deprecated but still need to maintain state
// to prevent "inconsistent result after apply" errors during API transitions.
func UseStateForNullComputedBool() planmodifier.Bool {
	return useStateForNullComputedBoolModifier{}
}

type useStateForNullComputedBoolModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m useStateForNullComputedBoolModifier) Description(_ context.Context) string {
	return "Preserves the prior state value when the computed value is null, for backwards compatibility with API changes."
}

// MarkdownDescription returns a markdown formatted description of the plan modifier.
func (m useStateForNullComputedBoolModifier) MarkdownDescription(_ context.Context) string {
	return "Preserves the prior state value when the computed value is null, for backwards compatibility with API changes."
}

// PlanModifyBool implements the plan modifier logic.
func (m useStateForNullComputedBoolModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	// If there is no prior state (new resource), do nothing
	if req.State.Raw.IsNull() {
		return
	}

	// If the prior state value is null or unknown, do nothing
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}

	// If the plan value is not null, the API is still returning the value
	// In this case, respect the API response
	if !req.PlanValue.IsNull() {
		return
	}

	// At this point:
	// - We have a prior state with a known value
	// - The API didn't return a value (plan is null)
	// - We should preserve the prior state value
	resp.PlanValue = req.StateValue
}
