package mistvalidator

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = ParseCidrPortValidator{}

type ParseCidrPortValidator struct {
	requireIp bool
	allowVar  bool
}

func (o ParseCidrPortValidator) Description(_ context.Context) string {
	switch {
	case o.requireIp && o.allowVar:
		return "Ensures that the supplied value can be parsed as an \"IPv4Prefix:Port\". Both IPv4Prefix and Port can be a variable (i.e. \"{{myvar}}\")"
	case o.requireIp:
		return "Ensures that the supplied value can be parsed as an \"IPv4Prefix:Port\""
	case o.allowVar:
		return "Ensures that the supplied value can be parsed as an \"IPv4Prefix:Port\" or \":Port\". Both IPv4Prefix and Port can be a variable (i.e. \"{{myvar}}\")"
	default:
		return "Ensures that the supplied value can be parsed as an \"IPv4Prefix:Port\" or \":Port\""
	}
}

func (o ParseCidrPortValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ParseCidrPortValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	if split := strings.Split(value, ":"); len(split) != 2 {
		switch {
		case o.requireIp:
			resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
				req.Path, "value must be an IPv4CIDR:Port (i.e. \"10.10.10.10/24:443\")", value))
		default:
			resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
				req.Path, "value must be an IPv4CIDR:Port (i.e. \"10.10.10.10/24:443\") or a Port (i.e. \":443\")", value))
		}
		return

	} else {
		// Validate IP CIDR
		if split[0] != "" || o.requireIp {

			if isVar := checkIsVar(split[0]); !o.allowVar || !isVar {

				ip, ipNet, err := net.ParseCIDR(split[0])
				if err != nil || ip == nil || ipNet == nil {
					resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
						req.Path,
						"value is not a valid CIDR notation prefix",
						split[0]))
					return
				}

				if !ipNet.IP.Equal(ip) {
					resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
						req.Path,
						fmt.Sprintf("value is not a valid CIDR base address (did you mean %q?)", ipNet.String()),
						split[0],
					))
				}

				if len(ip.To4()) != net.IPv4len {
					resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
						req.Path, "value is not an IPv4 CIDR prefix", split[0]))
					return
				}
			}
		}

		// Validate Port
		if isVar := checkIsVar(split[1]); !o.allowVar || !isVar {
			port, e := strconv.Atoi(split[1])
			if e != nil || port < 0 || port > 65535 {
				resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
					req.Path, "value doesn't have a valid Port Number", value))
				return
			}
		}
	}
}

func ParseCidrPort(
	requireIp bool,
	allowVar bool,
) validator.String {
	return ParseCidrPortValidator{
		requireIp: requireIp,
		allowVar:  allowVar,
	}
}
