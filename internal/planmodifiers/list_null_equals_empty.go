package mistplanmodifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// ListNullEqualsEmpty returns a plan modifier that treats null and empty lists
// as semantically equivalent. This prevents drift when the API does not
// distinguish between an absent field (null) and an empty list, while the
// Terraform configuration may use either form.
func ListNullEqualsEmpty() planmodifier.List {
	return listNullEqualsEmptyModifier{}
}

type listNullEqualsEmptyModifier struct{}

func (m listNullEqualsEmptyModifier) Description(_ context.Context) string {
	return "Treats null and empty list as semantically equivalent to prevent configuration drift."
}

func (m listNullEqualsEmptyModifier) MarkdownDescription(_ context.Context) string {
	return "Treats null and empty list as semantically equivalent to prevent configuration drift."
}

func (m listNullEqualsEmptyModifier) PlanModifyList(_ context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	// Do nothing for new resources (no prior state)
	if req.State.Raw.IsNull() {
		return
	}

	planEmpty := req.PlanValue.IsNull() || len(req.PlanValue.Elements()) == 0
	stateEmpty := req.StateValue.IsNull() || len(req.StateValue.Elements()) == 0

	// If both are semantically empty, align the plan with the state to suppress drift
	if planEmpty && stateEmpty {
		resp.PlanValue = req.StateValue
	}
}
