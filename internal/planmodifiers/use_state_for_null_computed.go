package mistplanmodifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// UseStateForNullComputedBool returns a plan modifier that copies the prior state value
// to the planned value when the planned value is null. This is specifically designed
// for backwards compatibility scenarios involving Optional+Computed fields.
//
// This modifier differs from UseStateForUnknown in that it:
// - Only activates when the plan value is null (not unknown)
// - Preserves the prior state value instead of allowing it to become null
// - Does not activate when there's no prior state (new resource)
//
// Use case: Deprecated Optional+Computed fields where the user omits the attribute from
// their config but the field still has a value in state (from a previous API response).
// Without this modifier, the plan would show the field changing from its state value to
// null, causing "inconsistent result after apply" errors.
//
// Note: This modifier requires that the field NOT have a Default value, as defaults
// prevent the plan value from being null when the user omits the attribute.
func UseStateForNullComputedBool() planmodifier.Bool {
	return useStateForNullComputedBoolModifier{}
}

type useStateForNullComputedBoolModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m useStateForNullComputedBoolModifier) Description(_ context.Context) string {
	return "When the user omits an Optional+Computed attribute from their config, preserves the value from state instead of planning it to null."
}

// MarkdownDescription returns a markdown formatted description of the plan modifier.
func (m useStateForNullComputedBoolModifier) MarkdownDescription(_ context.Context) string {
	return "When the user omits an Optional+Computed attribute from their config, preserves the value from state instead of planning it to null."
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

	// If the plan value is not null, it means either:
	// - The user explicitly set the value in their config, OR
	// - A schema Default populated the value
	// In either case, respect that value instead of preserving state
	if !req.PlanValue.IsNull() {
		return
	}

	// At this point:
	// - We have a prior state with a known value
	// - The plan value is null (user omitted the attribute and no default exists)
	// - We should preserve the prior state value to prevent unnecessary drift
	resp.PlanValue = req.StateValue
}
