package mistvalidator

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = ParseSubnetValidator{}

type ParseSubnetValidator struct {
	cidrFormat    bool
	decimalFormat bool
}

func (o ParseSubnetValidator) Description(_ context.Context) string {
	switch {
	case o.cidrFormat && o.decimalFormat:
		return "Ensures that user submitted IPv4 Subnet has a valid CIDR format (e.g. \"/24\") and Decimal format (e.g. \"255.255.255.0\") - this usage is likely a mistake in the provider code"
	case o.cidrFormat:
		return "Ensures that user submitted IPv4 Subnet has a valid CIDR format (e.g. \"/24\")"
	case o.decimalFormat:
		return "Ensures that user submitted IPv4 Subnet has a valid Decimal format (e.g. \"255.255.255.0\")"
	default:
		return "Ensures that user submitted IPv4 Subnet has a valid CIDR format (e.g. \"/24\") or Decimal format (e.g. \"255.255.255.0\")"
	}
}

func (o ParseSubnetValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ParseSubnetValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	re_subnet := `^(255)\.(0|128|192|224|240|248|252|254|255)\.(0|128|192|224|240|248|252|254|255)\.(0|128|192|224|240|248|252|254|255)`

	value := req.ConfigValue.ValueString()

	if strings.HasPrefix(value, "/") {
		if mask, e := strconv.Atoi(strings.Replace(value, "/", "", 1)); e != nil || mask < 0 || mask > 32 {
			resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
				req.Path,
				"value is not a valid Netmask",
				value,
			))
			return
		}
	} else if is_netmask, err := regexp.MatchString(re_subnet, value); !is_netmask || err != nil {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"value is not a valid Netmask",
			value,
		))
		return
	}
}

func ParseSubnet(cidrFormat bool, decimalFormat bool) validator.String {
	return ParseSubnetValidator{
		cidrFormat:    cidrFormat,
		decimalFormat: decimalFormat,
	}
}