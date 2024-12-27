package mistvalidator

import (
	"context"
	"net"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = ParseIpPortValidator{}

type ParseIpPortValidator struct {
	requireIp bool
	allowVar  bool
}

func (o ParseIpPortValidator) Description(_ context.Context) string {
	switch {
	case o.requireIp && o.allowVar:
		return "Ensures that the supplied value can be parsed as an \"IPv4:Port\". Both IPv4 and Port can be a variable (i.e. \"{{myvar}}\")"
	case o.requireIp:
		return "Ensures that the supplied value can be parsed as an \"IPv4:Port\""
	case o.allowVar:
		return "Ensures that the supplied can be parsed as either an \"IPv4:Port\" or \":Port\". Both IPv4 and Port can be a variable (i.e. \"{{myvar}}\")"
	default:
		return "Ensures that the supplied can be parsed as either an \"IPv4:Port\" or \":Port\""
	}
}

func (o ParseIpPortValidator) MarkdownDescription(ctx context.Context) string {
	return o.Description(ctx)
}

func (o ParseIpPortValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	if split := strings.Split(value, ":"); len(split) != 2 {
		switch {
		case o.requireIp:
			resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
				req.Path, "value must be an IPv4:Port (i.e. \"10.10.10.10:443\")", value))
		default:
			resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
				req.Path, "value must be an IPv4:Port (i.e. \"10.10.10.10:443\") or a Port (i.e. \":443\")", value))
		}
		return

	} else {
		// Validate IP Address
		if split[0] != "" || o.requireIp {
			if isVar := checkIsVar(split[0]); !o.allowVar || !isVar {

				ip := net.ParseIP(split[0])
				if ip == nil {
					resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
						req.Path, "value must be an IP:Port (i.e. \"10.10.10.10:443\")", value))
					return
				}

				if len(ip.To4()) != net.IPv4len {
					resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
						req.Path, "value must be an IPv4:Port", value))
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

func ParseIpPort(
	requireIp bool,
	allowVar bool,
) validator.String {
	return ParseIpPortValidator{
		requireIp: requireIp,
		allowVar:  allowVar,
	}
}
