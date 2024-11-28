package mistplanmodifiers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// DefaultObject return a string plan modifier that sets the specified value if the planned value is Null.
func DefaultObject(t map[string]attr.Type, v map[string]attr.Value) planmodifier.Object {
	return defaultObject{
		typ: t,
		val: v,
	}
}

// defaultValue holds our default value and allows us to implement the `planmodifier.Object` interface
type defaultObject struct {
	typ map[string]attr.Type
	val map[string]attr.Value
}

// Description implements the `planmodifier.Object` interface
func (m defaultObject) Description(context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to %s", m.val)
}

// MarkdownDescription implements the `planmodifier.Object` interface
func (m defaultObject) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx) // reuse our plaintext Description
}

// PlanModifyObject implements the `planmodifier.Object` interface
func (m defaultObject) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	// If the attribute configuration is not null it is explicit; we should apply the planned value.
	if !req.ConfigValue.IsNull() {
		return
	}

	// Otherwise, the configuration is null, so apply the default value to the response.
	resp.PlanValue = types.ObjectValueMust(m.typ, m.val)
}
