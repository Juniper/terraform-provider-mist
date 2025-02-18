package mistplanmodifiers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// useStateForUnknownModifier implements the plan modifier.
type OnlyRefreshIfValueHasChangedModifier struct {
	Expression path.Expression
}

// Description returns a human-readable description of the plan modifier.
func (m OnlyRefreshIfValueHasChangedModifier) Description(_ context.Context) string {
	return fmt.Sprintf("Only refresh the value if when attribute at %q is updated.", m.Expression)
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m OnlyRefreshIfValueHasChangedModifier) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("Only refresh the value if when attribute at %q is updated.", m.Expression)
}

func compareList(ctx context.Context, req planmodifier.StringRequest, mp path.Path) bool {
	var needRefresh = false
	var s basetypes.ListValue
	var p basetypes.ListValue
	req.State.GetAttribute(ctx, mp, &s)
	req.Plan.GetAttribute(ctx, mp, &p)

	var sl []string
	var pl []string
	for _, v := range s.Elements() {
		sl = append(sl, v.String())
	}
	for _, v := range p.Elements() {
		pl = append(pl, v.String())
	}
	if len(sl) != len(pl) {
		needRefresh = true
	} else {
		for i := range sl {
			if sl[i] != pl[i] {
				needRefresh = true
			}
		}
	}
	return needRefresh
}
func compareString(ctx context.Context, req planmodifier.StringRequest, mp path.Path) bool {
	var needRefresh = false
	var s basetypes.StringValue
	var p basetypes.StringValue
	req.State.GetAttribute(ctx, mp, &s)
	req.Plan.GetAttribute(ctx, mp, &p)

	if s != p {
		needRefresh = true
	}
	return needRefresh
}
func compareBool(ctx context.Context, req planmodifier.StringRequest, mp path.Path) bool {
	var needRefresh = false
	var s basetypes.BoolValue
	var p basetypes.BoolValue
	req.State.GetAttribute(ctx, mp, &s)
	req.Plan.GetAttribute(ctx, mp, &p)

	if s != p {
		needRefresh = true
	}
	return needRefresh
}
func compareFloat64(ctx context.Context, req planmodifier.StringRequest, mp path.Path) bool {
	var needRefresh = false
	var s basetypes.Float64Value
	var p basetypes.Float64Value
	req.State.GetAttribute(ctx, mp, &s)
	req.Plan.GetAttribute(ctx, mp, &p)

	if s != p {
		needRefresh = true
	}
	return needRefresh
}
func compareInt64(ctx context.Context, req planmodifier.StringRequest, mp path.Path) bool {
	var needRefresh = false
	var s basetypes.Int64Value
	var p basetypes.Int64Value
	req.State.GetAttribute(ctx, mp, &s)
	req.Plan.GetAttribute(ctx, mp, &p)

	if s != p {
		needRefresh = true
	}
	return needRefresh
}

// PlanModifyString PlanModifyBool implements the plan modification logic.
// It will reuse the state value if the provided attribute is not changed
func (m OnlyRefreshIfValueHasChangedModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	needRefresh := true
	mergedExpressions := req.PathExpression.MergeExpressions(m.Expression)

	for _, expression := range mergedExpressions {
		matchedPaths, e := req.State.PathMatches(ctx, expression)
		if e != nil {
			return
		}
		for _, mp := range matchedPaths {
			attr, e := req.Plan.Schema.AttributeAtPath(ctx, mp)
			if e != nil {
				continue
			}
			attrType := attr.GetType().String()

			switch attrType {
			case "types.ListType[basetypes.StringType]":
				needRefresh = compareList(ctx, req, mp)
			case "types.BoolType":
				needRefresh = compareBool(ctx, req, mp)
			case "types.Float64Type":
				needRefresh = compareFloat64(ctx, req, mp)
			case "types.Int64Type":
				needRefresh = compareInt64(ctx, req, mp)
			case "types.StringType":
				needRefresh = compareString(ctx, req, mp)
			default:
				continue
			}
		}
	}
	if !needRefresh {
		resp.PlanValue = req.StateValue
	}
}

func OnlyRefreshIfValueHasChanged(expression path.Expression) planmodifier.String {
	return OnlyRefreshIfValueHasChangedModifier{
		Expression: expression,
	}
}
