package mistvalidator

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ validator.List = SwitchMatchingRuleNamesValidator{}

type SwitchMatchingRuleNamesValidator struct{}

func (o SwitchMatchingRuleNamesValidator) Description(_ context.Context) string {
	return "Ensures that value submitted by the user contains a Mist Variable"
}

func (o SwitchMatchingRuleNamesValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o SwitchMatchingRuleNamesValidator) ValidateList(_ context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	pattern := `^[a-z0-9_\-.]{1,32}$`
	re := regexp.MustCompile(pattern)
	var names []string

	rules := req.ConfigValue
	for idx, rule := range rules.Elements() {
		fieldName := reflect.ValueOf(rule).FieldByName("Name")

		if fieldName.IsValid() {
			var i interface{} = fieldName.Interface()
			s := i.(basetypes.StringValue)
			if idx != len(rules.Elements())-1 && strings.ToLower(s.ValueString()) == "default" {
				resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
					req.Path,
					fmt.Sprintf(
						"value can only use the rule name \"default\" for the last item in the list. Currently used by the rule %d/%d",
						idx+1,
						len(rules.Elements()),
					),
					s.ValueString(),
				))
			}
			if matched := re.MatchString(s.ValueString()); !matched {
				resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
					req.Path,
					fmt.Sprintf(
						"value can only have rule names containing a-z, 0-9, _, -, . and up to 32 characters; Please rename the rule %d/%d",
						idx+1,
						len(rules.Elements()),
					),
					s.ValueString(),
				))
			}
			if slices.Contains(names, s.ValueString()) {
				resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
					req.Path,
					fmt.Sprintf(
						"value cannot use the same rule name for multiple rules; Please rename the rule %d/%d",
						idx+1,
						len(rules.Elements()),
					),
					s.ValueString(),
				))
			}
			names = append(names, s.ValueString())
		}
	}
}

func SwitchMatchingRuleNames() validator.List {
	return SwitchMatchingRuleNamesValidator{}
}
