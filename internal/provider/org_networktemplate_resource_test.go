package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_networktemplate"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgNetworktemplateModel(t *testing.T) {
	type testStep struct {
		config OrgNetworktemplateModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgNetworktemplateModel{
						OrgId: GetTestOrgId(),
						Name:  "test-networktemplate",
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_networktemplate_resource/org_networktemplate_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		FixtureOrgNetworktemplateModel := OrgNetworktemplateModel{}

		err = hcl.Decode(&FixtureOrgNetworktemplateModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		FixtureOrgNetworktemplateModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgNetworktemplateModel,
				},
			},
		}
	}

	resourceType := "org_networktemplate"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_networktemplate.OrgNetworktemplateResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				configStr := Render(resourceType, tName, string(f.Bytes()))

				checks := config.testChecks(t, resourceType, tName, tracker)
				chkLog := checks.string()
				stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

				t.Logf("\n// ------ begin config for %s ------\n%s// -------- end config for %s ------\n\n", stepName, configStr, stepName)
				t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end config for %s ------\n\n", stepName, chkLog, stepName)

				steps[i] = resource.TestStep{
					Config: configStr,
					Check:  resource.ComposeAggregateTestCheckFunc(checks.checks...),
				}
			}

			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps:                    steps,
			})
		})
	}
	if tracker != nil {
		tracker.FieldCoverageReport(t)
	}
}

func (o *OrgNetworktemplateModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.SetTracker(tracker)

	// Check required fields
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)

	// Optional list attributes
	if len(o.AclPolicies) > 0 {
		checks.append(t, "TestCheckResourceAttr", "acl_policies.#", fmt.Sprintf("%d", len(o.AclPolicies)))
		for i, policy := range o.AclPolicies {
			basePath := fmt.Sprintf("acl_policies.%d", i)
			if policy.Name != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".name", *policy.Name)
			}
			if len(policy.SrcTags) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".src_tags.#", fmt.Sprintf("%d", len(policy.SrcTags)))
				for j, srcTag := range policy.SrcTags {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.src_tags.%d", basePath, j), srcTag)
				}
			}
			if len(policy.Actions) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".actions.#", fmt.Sprintf("%d", len(policy.Actions)))
				for j, action := range policy.Actions {
					actionPath := fmt.Sprintf("%s.actions.%d", basePath, j)
					if action.Action != nil {
						checks.append(t, "TestCheckResourceAttr", actionPath+".action", *action.Action)
					}
					checks.append(t, "TestCheckResourceAttr", actionPath+".dst_tag", action.DstTag)
				}
			}
		}
	}

	if len(o.AdditionalConfigCmds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "additional_config_cmds.#", fmt.Sprintf("%d", len(o.AdditionalConfigCmds)))
		for i, cmd := range o.AdditionalConfigCmds {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("additional_config_cmds.%d", i), cmd)
		}
	}

	// BgpConfig map
	if len(o.BgpConfig) > 0 {
		for key, bgp := range o.BgpConfig {
			basePath := fmt.Sprintf("bgp_config.%s", key)
			if bgp.AuthKey != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".auth_key", *bgp.AuthKey)
			}
			if bgp.BfdMinimumInterval != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".bfd_minimum_interval", fmt.Sprintf("%d", *bgp.BfdMinimumInterval))
			}
			if bgp.ExportPolicy != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".export_policy", *bgp.ExportPolicy)
			}
			if bgp.HoldTime != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".hold_time", fmt.Sprintf("%d", *bgp.HoldTime))
			}

			if bgp.ImportPolicy != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".import_policy", *bgp.ImportPolicy)
			}
			checks.append(t, "TestCheckResourceAttr", basePath+".local_as", bgp.LocalAs)
			if len(bgp.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".networks.#", fmt.Sprintf("%d", len(bgp.Networks)))
				for i, network := range bgp.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.networks.%d", basePath, i), network)
				}
			}
			checks.append(t, "TestCheckResourceAttr", basePath+".type", bgp.BgpConfigType)
			// BgpConfig neighbors
			if len(bgp.Neighbors) > 0 {
				for neighborKey, neighbor := range bgp.Neighbors {
					neighborPath := fmt.Sprintf("%s.neighbors.%s", basePath, neighborKey)
					if neighbor.ExportPolicy != nil {
						checks.append(t, "TestCheckResourceAttr", neighborPath+".export_policy", *neighbor.ExportPolicy)
					}
					if neighbor.HoldTime != nil {
						checks.append(t, "TestCheckResourceAttr", neighborPath+".hold_time", fmt.Sprintf("%d", *neighbor.HoldTime))
					}
					if neighbor.ImportPolicy != nil {
						checks.append(t, "TestCheckResourceAttr", neighborPath+".import_policy", *neighbor.ImportPolicy)
					}
					if neighbor.MultihopTtl != nil {
						checks.append(t, "TestCheckResourceAttr", neighborPath+".multihop_ttl", fmt.Sprintf("%d", *neighbor.MultihopTtl))
					}
					checks.append(t, "TestCheckResourceAttr", neighborPath+".neighbor_as", neighbor.NeighborAs)
				}
			}
		}
	}

	if len(o.DnsServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "dns_servers.#", fmt.Sprintf("%d", len(o.DnsServers)))
		for i, server := range o.DnsServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dns_servers.%d", i), server)
		}
	}

	if len(o.DnsSuffix) > 0 {
		checks.append(t, "TestCheckResourceAttr", "dns_suffix.#", fmt.Sprintf("%d", len(o.DnsSuffix)))
		for i, suffix := range o.DnsSuffix {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dns_suffix.%d", i), suffix)
		}
	}

	if len(o.NtpServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ntp_servers.#", fmt.Sprintf("%d", len(o.NtpServers)))
		for i, server := range o.NtpServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ntp_servers.%d", i), server)
		}
	}

	// Optional boolean attributes
	if o.RemoveExistingConfigs != nil {
		checks.append(t, "TestCheckResourceAttr", "remove_existing_configs", fmt.Sprintf("%t", *o.RemoveExistingConfigs))
	}

	// Optional complex object attributes
	if o.DhcpSnooping != nil {
		if o.DhcpSnooping.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "dhcp_snooping.enabled", fmt.Sprintf("%t", *o.DhcpSnooping.Enabled))
		}
		if o.DhcpSnooping.AllNetworks != nil {
			checks.append(t, "TestCheckResourceAttr", "dhcp_snooping.all_networks", fmt.Sprintf("%t", *o.DhcpSnooping.AllNetworks))
		}
		if o.DhcpSnooping.EnableArpSpoofCheck != nil {
			checks.append(t, "TestCheckResourceAttr", "dhcp_snooping.enable_arp_spoof_check", fmt.Sprintf("%t", *o.DhcpSnooping.EnableArpSpoofCheck))
		}
		if o.DhcpSnooping.EnableIpSourceGuard != nil {
			checks.append(t, "TestCheckResourceAttr", "dhcp_snooping.enable_ip_source_guard", fmt.Sprintf("%t", *o.DhcpSnooping.EnableIpSourceGuard))
		}
		if len(o.DhcpSnooping.Networks) > 0 {
			checks.append(t, "TestCheckResourceAttr", "dhcp_snooping.networks.#", fmt.Sprintf("%d", len(o.DhcpSnooping.Networks)))
			for i, network := range o.DhcpSnooping.Networks {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcp_snooping.networks.%d", i), network)
			}
		}
	}

	if o.MistNac != nil {
		if o.MistNac.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.enabled", fmt.Sprintf("%t", *o.MistNac.Enabled))
		}
		if o.MistNac.Network != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.network", *o.MistNac.Network)
		}
	}

	if o.RadiusConfig != nil {
		// Optional boolean fields
		if o.RadiusConfig.AcctImmediateUpdate != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.acct_immediate_update", fmt.Sprintf("%t", *o.RadiusConfig.AcctImmediateUpdate))
		}
		if o.RadiusConfig.CoaEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.coa_enabled", fmt.Sprintf("%t", *o.RadiusConfig.CoaEnabled))
		}
		if o.RadiusConfig.FastDot1xTimers != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.fast_dot1x_timers", fmt.Sprintf("%t", *o.RadiusConfig.FastDot1xTimers))
		}

		// Optional integer fields
		if o.RadiusConfig.AcctInterimInterval != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.acct_interim_interval", fmt.Sprintf("%d", int(*o.RadiusConfig.AcctInterimInterval)))
		}
		if o.RadiusConfig.AuthServersRetries != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.auth_servers_retries", fmt.Sprintf("%d", int(*o.RadiusConfig.AuthServersRetries)))
		}
		if o.RadiusConfig.AuthServersTimeout != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.auth_servers_timeout", fmt.Sprintf("%d", int(*o.RadiusConfig.AuthServersTimeout)))
		}

		// Optional string fields
		if o.RadiusConfig.AuthServerSelection != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.auth_server_selection", *o.RadiusConfig.AuthServerSelection)
		}
		if o.RadiusConfig.CoaPort != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.coa_port", *o.RadiusConfig.CoaPort)
		}
		if o.RadiusConfig.Network != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.network", *o.RadiusConfig.Network)
		}
		if o.RadiusConfig.SourceIp != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.source_ip", *o.RadiusConfig.SourceIp)
		}

		// AcctServers list
		if len(o.RadiusConfig.AcctServers) > 0 {
			checks.append(t, "TestCheckResourceAttr", "radius_config.acct_servers.#", fmt.Sprintf("%d", len(o.RadiusConfig.AcctServers)))
			for i, server := range o.RadiusConfig.AcctServers {
				serverPath := fmt.Sprintf("radius_config.acct_servers.%d", i)

				// Required fields
				checks.append(t, "TestCheckResourceAttr", serverPath+".host", server.Host)
				checks.append(t, "TestCheckResourceAttr", serverPath+".secret", server.Secret)

				// Optional fields
				if server.KeywrapEnabled != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".keywrap_enabled", fmt.Sprintf("%t", *server.KeywrapEnabled))
				}
				if server.KeywrapFormat != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".keywrap_format", *server.KeywrapFormat)
				}
				if server.KeywrapKek != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".keywrap_kek", *server.KeywrapKek)
				}
				if server.KeywrapMack != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".keywrap_mack", *server.KeywrapMack)
				}
				if server.Port != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".port", *server.Port)
				}
			}
		}

		// AuthServers list
		if len(o.RadiusConfig.AuthServers) > 0 {
			checks.append(t, "TestCheckResourceAttr", "radius_config.auth_servers.#", fmt.Sprintf("%d", len(o.RadiusConfig.AuthServers)))
			for i, server := range o.RadiusConfig.AuthServers {
				serverPath := fmt.Sprintf("radius_config.auth_servers.%d", i)

				// Required fields
				checks.append(t, "TestCheckResourceAttr", serverPath+".host", server.Host)
				checks.append(t, "TestCheckResourceAttr", serverPath+".secret", server.Secret)

				// Optional fields
				if server.KeywrapEnabled != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".keywrap_enabled", fmt.Sprintf("%t", *server.KeywrapEnabled))
				}
				if server.KeywrapFormat != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".keywrap_format", *server.KeywrapFormat)
				}
				if server.KeywrapKek != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".keywrap_kek", *server.KeywrapKek)
				}
				if server.KeywrapMack != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".keywrap_mack", *server.KeywrapMack)
				}
				if server.Port != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".port", *server.Port)
				}
				if server.RequireMessageAuthenticator != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".require_message_authenticator", fmt.Sprintf("%t", *server.RequireMessageAuthenticator))
				}
			}
		}
	}

	if o.RemoteSyslog != nil {
		// Optional boolean fields
		if o.RemoteSyslog.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.enabled", fmt.Sprintf("%t", *o.RemoteSyslog.Enabled))
		}
		if o.RemoteSyslog.SendToAllServers != nil {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.send_to_all_servers", fmt.Sprintf("%t", *o.RemoteSyslog.SendToAllServers))
		}

		// Optional string fields
		if o.RemoteSyslog.Network != nil {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.network", *o.RemoteSyslog.Network)
		}
		if o.RemoteSyslog.TimeFormat != nil {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.time_format", *o.RemoteSyslog.TimeFormat)
		}

		// Cacerts list
		if len(o.RemoteSyslog.Cacerts) > 0 {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.cacerts.#", fmt.Sprintf("%d", len(o.RemoteSyslog.Cacerts)))
			for i, cacert := range o.RemoteSyslog.Cacerts {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.cacerts.%d", i), cacert)
			}
		}

		// Archive configuration
		if o.RemoteSyslog.Archive != nil {
			if o.RemoteSyslog.Archive.Files != nil {
				checks.append(t, "TestCheckResourceAttr", "remote_syslog.archive.files", *o.RemoteSyslog.Archive.Files)
			}
			if o.RemoteSyslog.Archive.Size != nil {
				checks.append(t, "TestCheckResourceAttr", "remote_syslog.archive.size", *o.RemoteSyslog.Archive.Size)
			}
		}

		// Console configuration
		if o.RemoteSyslog.Console != nil {
			if len(o.RemoteSyslog.Console.Contents) > 0 {
				checks.append(t, "TestCheckResourceAttr", "remote_syslog.console.contents.#", fmt.Sprintf("%d", len(o.RemoteSyslog.Console.Contents)))
				for i, content := range o.RemoteSyslog.Console.Contents {
					contentPath := fmt.Sprintf("remote_syslog.console.contents.%d", i)
					if content.Facility != nil {
						checks.append(t, "TestCheckResourceAttr", contentPath+".facility", *content.Facility)
					}
					if content.Severity != nil {
						checks.append(t, "TestCheckResourceAttr", contentPath+".severity", *content.Severity)
					}
				}
			}
		}

		// Files configuration
		if len(o.RemoteSyslog.Files) > 0 {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.files.#", fmt.Sprintf("%d", len(o.RemoteSyslog.Files)))
			for i, file := range o.RemoteSyslog.Files {
				filePath := fmt.Sprintf("remote_syslog.files.%d", i)

				// Optional boolean fields
				if file.EnableTls != nil {
					checks.append(t, "TestCheckResourceAttr", filePath+".enable_tls", fmt.Sprintf("%t", *file.EnableTls))
				}
				if file.ExplicitPriority != nil {
					checks.append(t, "TestCheckResourceAttr", filePath+".explicit_priority", fmt.Sprintf("%t", *file.ExplicitPriority))
				}
				if file.StructuredData != nil {
					checks.append(t, "TestCheckResourceAttr", filePath+".structured_data", fmt.Sprintf("%t", *file.StructuredData))
				}

				// Optional string fields
				if file.File != nil {
					checks.append(t, "TestCheckResourceAttr", filePath+".file", *file.File)
				}
				if file.Match != nil {
					checks.append(t, "TestCheckResourceAttr", filePath+".match", *file.Match)
				}

				// File archive
				if file.Archive != nil {
					if file.Archive.Files != nil {
						checks.append(t, "TestCheckResourceAttr", filePath+".archive.files", *file.Archive.Files)
					}
					if file.Archive.Size != nil {
						checks.append(t, "TestCheckResourceAttr", filePath+".archive.size", *file.Archive.Size)
					}
				}

				// File contents
				if len(file.Contents) > 0 {
					checks.append(t, "TestCheckResourceAttr", filePath+".contents.#", fmt.Sprintf("%d", len(file.Contents)))
					for j, content := range file.Contents {
						contentPath := fmt.Sprintf("%s.contents.%d", filePath, j)
						if content.Facility != nil {
							checks.append(t, "TestCheckResourceAttr", contentPath+".facility", *content.Facility)
						}
						if content.Severity != nil {
							checks.append(t, "TestCheckResourceAttr", contentPath+".severity", *content.Severity)
						}
					}
				}
			}
		}

		// Servers configuration
		if len(o.RemoteSyslog.Servers) > 0 {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.servers.#", fmt.Sprintf("%d", len(o.RemoteSyslog.Servers)))
			for i, server := range o.RemoteSyslog.Servers {
				serverPath := fmt.Sprintf("remote_syslog.servers.%d", i)

				// Optional boolean fields
				if server.ExplicitPriority != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".explicit_priority", fmt.Sprintf("%t", *server.ExplicitPriority))
				}
				if server.StructuredData != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".structured_data", fmt.Sprintf("%t", *server.StructuredData))
				}

				// Optional string fields
				if server.Facility != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".facility", *server.Facility)
				}
				if server.Host != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".host", *server.Host)
				}
				if server.Match != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".match", *server.Match)
				}
				if server.Port != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".port", *server.Port)
				}
				if server.Protocol != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".protocol", *server.Protocol)
				}
				if server.RoutingInstance != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".routing_instance", *server.RoutingInstance)
				}
				if server.ServerName != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".server_name", *server.ServerName)
				}
				if server.Severity != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".severity", *server.Severity)
				}
				if server.SourceAddress != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".source_address", *server.SourceAddress)
				}
				if server.Tag != nil {
					checks.append(t, "TestCheckResourceAttr", serverPath+".tag", *server.Tag)
				}

				// Server contents
				if len(server.Contents) > 0 {
					checks.append(t, "TestCheckResourceAttr", serverPath+".contents.#", fmt.Sprintf("%d", len(server.Contents)))
					for j, content := range server.Contents {
						contentPath := fmt.Sprintf("%s.contents.%d", serverPath, j)
						if content.Facility != nil {
							checks.append(t, "TestCheckResourceAttr", contentPath+".facility", *content.Facility)
						}
						if content.Severity != nil {
							checks.append(t, "TestCheckResourceAttr", contentPath+".severity", *content.Severity)
						}
					}
				}
			}
		}

		// Users configuration
		if len(o.RemoteSyslog.Users) > 0 {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.users.#", fmt.Sprintf("%d", len(o.RemoteSyslog.Users)))
			for i, user := range o.RemoteSyslog.Users {
				userPath := fmt.Sprintf("remote_syslog.users.%d", i)

				// Optional string fields
				if user.Match != nil {
					checks.append(t, "TestCheckResourceAttr", userPath+".match", *user.Match)
				}
				if user.User != nil {
					checks.append(t, "TestCheckResourceAttr", userPath+".user", *user.User)
				}

				// User contents
				if len(user.Contents) > 0 {
					checks.append(t, "TestCheckResourceAttr", userPath+".contents.#", fmt.Sprintf("%d", len(user.Contents)))
					for j, content := range user.Contents {
						contentPath := fmt.Sprintf("%s.contents.%d", userPath, j)
						if content.Facility != nil {
							checks.append(t, "TestCheckResourceAttr", contentPath+".facility", *content.Facility)
						}
						if content.Severity != nil {
							checks.append(t, "TestCheckResourceAttr", contentPath+".severity", *content.Severity)
						}
					}
				}
			}
		}
	}

	// Check routing_policies if present
	if len(o.RoutingPolicies) > 0 {
		checks.append(t, "TestCheckResourceAttr", "routing_policies.%", fmt.Sprintf("%d", len(o.RoutingPolicies)))
		for k, v := range o.RoutingPolicies {
			if len(v.Terms) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.#", k), fmt.Sprintf("%d", len(v.Terms)))
				// Check for terms by name using TestCheckTypeSetElemNestedAttrs to handle ordering
				for _, term := range v.Terms {
					termChecks := make(map[string]string)
					termChecks["name"] = term.Name

					if term.RoutingPolicyTermActions != nil {
						if term.RoutingPolicyTermActions.Accept != nil {
							termChecks["actions.accept"] = fmt.Sprintf("%t", *term.RoutingPolicyTermActions.Accept)
						}
						if term.RoutingPolicyTermActions.LocalPreference != nil {
							termChecks["actions.local_preference"] = *term.RoutingPolicyTermActions.LocalPreference
						}
						if len(term.RoutingPolicyTermActions.Community) > 0 {
							termChecks["actions.community.#"] = fmt.Sprintf("%d", len(term.RoutingPolicyTermActions.Community))
							for j, community := range term.RoutingPolicyTermActions.Community {
								termChecks[fmt.Sprintf("actions.community.%d", j)] = community
							}
						}
						if len(term.RoutingPolicyTermActions.PrependAsPath) > 0 {
							termChecks["actions.prepend_as_path.#"] = fmt.Sprintf("%d", len(term.RoutingPolicyTermActions.PrependAsPath))
							for j, prependAsPath := range term.RoutingPolicyTermActions.PrependAsPath {
								termChecks[fmt.Sprintf("actions.prepend_as_path.%d", j)] = prependAsPath
							}
						}
					}
					if term.Matching != nil {
						if len(term.Matching.AsPath) > 0 {
							termChecks["matching.as_path.#"] = fmt.Sprintf("%d", len(term.Matching.AsPath))
							for j, asPath := range term.Matching.AsPath {
								termChecks[fmt.Sprintf("matching.as_path.%d", j)] = asPath
							}
						}
						if len(term.Matching.Community) > 0 {
							termChecks["matching.community.#"] = fmt.Sprintf("%d", len(term.Matching.Community))
							for j, community := range term.Matching.Community {
								termChecks[fmt.Sprintf("matching.community.%d", j)] = community
							}
						}
						if len(term.Matching.Prefix) > 0 {
							termChecks["matching.prefix.#"] = fmt.Sprintf("%d", len(term.Matching.Prefix))
							for j, prefix := range term.Matching.Prefix {
								termChecks[fmt.Sprintf("matching.prefix.%d", j)] = prefix
							}
						}
						if len(term.Matching.Protocol) > 0 {
							termChecks["matching.protocol.#"] = fmt.Sprintf("%d", len(term.Matching.Protocol))
							for j, protocol := range term.Matching.Protocol {
								termChecks[fmt.Sprintf("matching.protocol.%d", j)] = protocol
							}
						}
					}
					checks.appendSetNestedCheck(t, fmt.Sprintf("routing_policies.%s.terms.*", k), termChecks)
				}
			}
		}
	}

	if o.SnmpConfig != nil {
		// Optional boolean field
		if o.SnmpConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.enabled", fmt.Sprintf("%t", *o.SnmpConfig.Enabled))
		}

		// Optional string fields
		if o.SnmpConfig.Contact != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.contact", *o.SnmpConfig.Contact)
		}
		if o.SnmpConfig.Description != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.description", *o.SnmpConfig.Description)
		}
		if o.SnmpConfig.EngineId != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.engine_id", *o.SnmpConfig.EngineId)
		}
		if o.SnmpConfig.EngineIdType != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.engine_id_type", *o.SnmpConfig.EngineIdType)
		}
		if o.SnmpConfig.Location != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.location", *o.SnmpConfig.Location)
		}
		if o.SnmpConfig.Name != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.name", *o.SnmpConfig.Name)
		}
		if o.SnmpConfig.Network != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.network", *o.SnmpConfig.Network)
		}

		// ClientList configuration
		if len(o.SnmpConfig.ClientList) > 0 {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.client_list.#", fmt.Sprintf("%d", len(o.SnmpConfig.ClientList)))
			for i, clientList := range o.SnmpConfig.ClientList {
				clientPath := fmt.Sprintf("snmp_config.client_list.%d", i)

				if clientList.ClientListName != nil {
					checks.append(t, "TestCheckResourceAttr", clientPath+".client_list_name", *clientList.ClientListName)
				}

				if len(clientList.Clients) > 0 {
					checks.append(t, "TestCheckResourceAttr", clientPath+".clients.#", fmt.Sprintf("%d", len(clientList.Clients)))
					for j, client := range clientList.Clients {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.clients.%d", clientPath, j), client)
					}
				}
			}
		}

		// TrapGroups configuration
		if len(o.SnmpConfig.TrapGroups) > 0 {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.trap_groups.#", fmt.Sprintf("%d", len(o.SnmpConfig.TrapGroups)))
			for i, trapGroup := range o.SnmpConfig.TrapGroups {
				trapPath := fmt.Sprintf("snmp_config.trap_groups.%d", i)

				if trapGroup.GroupName != nil {
					checks.append(t, "TestCheckResourceAttr", trapPath+".group_name", *trapGroup.GroupName)
				}
				if trapGroup.Version != nil {
					checks.append(t, "TestCheckResourceAttr", trapPath+".version", *trapGroup.Version)
				}

				if len(trapGroup.Categories) > 0 {
					checks.append(t, "TestCheckResourceAttr", trapPath+".categories.#", fmt.Sprintf("%d", len(trapGroup.Categories)))
					for j, category := range trapGroup.Categories {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.categories.%d", trapPath, j), category)
					}
				}

				if len(trapGroup.Targets) > 0 {
					checks.append(t, "TestCheckResourceAttr", trapPath+".targets.#", fmt.Sprintf("%d", len(trapGroup.Targets)))
					for j, target := range trapGroup.Targets {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.targets.%d", trapPath, j), target)
					}
				}
			}
		}

		// V2cConfig configuration
		if len(o.SnmpConfig.V2cConfig) > 0 {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.v2c_config.#", fmt.Sprintf("%d", len(o.SnmpConfig.V2cConfig)))
			for i, v2cConfig := range o.SnmpConfig.V2cConfig {
				v2cPath := fmt.Sprintf("snmp_config.v2c_config.%d", i)

				if v2cConfig.Authorization != nil {
					checks.append(t, "TestCheckResourceAttr", v2cPath+".authorization", *v2cConfig.Authorization)
				}
				if v2cConfig.ClientListName != nil {
					checks.append(t, "TestCheckResourceAttr", v2cPath+".client_list_name", *v2cConfig.ClientListName)
				}
				if v2cConfig.CommunityName != nil {
					checks.append(t, "TestCheckResourceAttr", v2cPath+".community_name", *v2cConfig.CommunityName)
				}
				if v2cConfig.View != nil {
					checks.append(t, "TestCheckResourceAttr", v2cPath+".view", *v2cConfig.View)
				}
			}
		}

		// V3Config configuration (basic structure - can be expanded with nested structs)
		if o.SnmpConfig.V3Config != nil {
			// Notify configuration
			if len(o.SnmpConfig.V3Config.Notify) > 0 {
				checks.append(t, "TestCheckResourceAttr", "snmp_config.v3_config.notify.#", fmt.Sprintf("%d", len(o.SnmpConfig.V3Config.Notify)))
				for i, notify := range o.SnmpConfig.V3Config.Notify {
					notifyPath := fmt.Sprintf("snmp_config.v3_config.notify.%d", i)
					checks.append(t, "TestCheckResourceAttr", notifyPath+".name", notify.Name)
					checks.append(t, "TestCheckResourceAttr", notifyPath+".tag", notify.Tag)
					checks.append(t, "TestCheckResourceAttr", notifyPath+".type", notify.NotifyType)
				}
			}

			// NotifyFilter configuration
			if len(o.SnmpConfig.V3Config.NotifyFilter) > 0 {
				checks.append(t, "TestCheckResourceAttr", "snmp_config.v3_config.notify_filter.#", fmt.Sprintf("%d", len(o.SnmpConfig.V3Config.NotifyFilter)))
				for i, notifyFilter := range o.SnmpConfig.V3Config.NotifyFilter {
					filterPath := fmt.Sprintf("snmp_config.v3_config.notify_filter.%d", i)
					if notifyFilter.ProfileName != nil {
						checks.append(t, "TestCheckResourceAttr", filterPath+".profile_name", *notifyFilter.ProfileName)
					}
					if len(notifyFilter.Snmpv3Contents) > 0 {
						checks.append(t, "TestCheckResourceAttr", filterPath+".contents.#", fmt.Sprintf("%d", len(notifyFilter.Snmpv3Contents)))
						for j, content := range notifyFilter.Snmpv3Contents {
							contentPath := fmt.Sprintf("%s.contents.%d", filterPath, j)
							checks.append(t, "TestCheckResourceAttr", contentPath+".oid", content.Oid)
							if content.Include != nil {
								checks.append(t, "TestCheckResourceAttr", contentPath+".include", fmt.Sprintf("%t", *content.Include))
							}
						}
					}
				}
			}

			// TargetAddress configuration
			if len(o.SnmpConfig.V3Config.TargetAddress) > 0 {
				checks.append(t, "TestCheckResourceAttr", "snmp_config.v3_config.target_address.#", fmt.Sprintf("%d", len(o.SnmpConfig.V3Config.TargetAddress)))
				for i, targetAddr := range o.SnmpConfig.V3Config.TargetAddress {
					addrPath := fmt.Sprintf("snmp_config.v3_config.target_address.%d", i)
					checks.append(t, "TestCheckResourceAttr", addrPath+".address", targetAddr.Address)
					checks.append(t, "TestCheckResourceAttr", addrPath+".address_mask", targetAddr.AddressMask)
					checks.append(t, "TestCheckResourceAttr", addrPath+".target_address_name", targetAddr.TargetAddressName)
					if targetAddr.Port != nil {
						checks.append(t, "TestCheckResourceAttr", addrPath+".port", *targetAddr.Port)
					}
					if targetAddr.TagList != nil {
						checks.append(t, "TestCheckResourceAttr", addrPath+".tag_list", *targetAddr.TagList)
					}
					if targetAddr.TargetParameters != nil {
						checks.append(t, "TestCheckResourceAttr", addrPath+".target_parameters", *targetAddr.TargetParameters)
					}
				}
			}

			// TargetParameters configuration
			if len(o.SnmpConfig.V3Config.TargetParameters) > 0 {
				checks.append(t, "TestCheckResourceAttr", "snmp_config.v3_config.target_parameters.#", fmt.Sprintf("%d", len(o.SnmpConfig.V3Config.TargetParameters)))
				for i, targetParam := range o.SnmpConfig.V3Config.TargetParameters {
					paramPath := fmt.Sprintf("snmp_config.v3_config.target_parameters.%d", i)
					checks.append(t, "TestCheckResourceAttr", paramPath+".message_processing_model", targetParam.MessageProcessingModel)
					checks.append(t, "TestCheckResourceAttr", paramPath+".name", targetParam.Name)
					if targetParam.NotifyFilter != nil {
						checks.append(t, "TestCheckResourceAttr", paramPath+".notify_filter", *targetParam.NotifyFilter)
					}
					if targetParam.SecurityLevel != nil {
						checks.append(t, "TestCheckResourceAttr", paramPath+".security_level", *targetParam.SecurityLevel)
					}
					if targetParam.SecurityModel != nil {
						checks.append(t, "TestCheckResourceAttr", paramPath+".security_model", *targetParam.SecurityModel)
					}
					if targetParam.SecurityName != nil {
						checks.append(t, "TestCheckResourceAttr", paramPath+".security_name", *targetParam.SecurityName)
					}
				}
			}

			// Usm configuration
			if len(o.SnmpConfig.V3Config.Usm) > 0 {
				checks.append(t, "TestCheckResourceAttr", "snmp_config.v3_config.usm.#", fmt.Sprintf("%d", len(o.SnmpConfig.V3Config.Usm)))
				for i, usm := range o.SnmpConfig.V3Config.Usm {
					usmPath := fmt.Sprintf("snmp_config.v3_config.usm.%d", i)
					checks.append(t, "TestCheckResourceAttr", usmPath+".engine_type", usm.EngineType)
					if usm.RemoteEngineId != nil {
						checks.append(t, "TestCheckResourceAttr", usmPath+".remote_engine_id", *usm.RemoteEngineId)
					}
					if len(usm.Snmpv3Users) > 0 {
						checks.append(t, "TestCheckResourceAttr", usmPath+".users.#", fmt.Sprintf("%d", len(usm.Snmpv3Users)))
						for j, user := range usm.Snmpv3Users {
							userPath := fmt.Sprintf("%s.users.%d", usmPath, j)
							if user.Name != nil {
								checks.append(t, "TestCheckResourceAttr", userPath+".name", *user.Name)
							}
							if user.AuthenticationType != nil {
								checks.append(t, "TestCheckResourceAttr", userPath+".authentication_type", *user.AuthenticationType)
							}
							if user.AuthenticationPassword != nil {
								checks.append(t, "TestCheckResourceAttr", userPath+".authentication_password", *user.AuthenticationPassword)
							}
							if user.EncryptionType != nil {
								checks.append(t, "TestCheckResourceAttr", userPath+".encryption_type", *user.EncryptionType)
							}
							if user.EncryptionPassword != nil {
								checks.append(t, "TestCheckResourceAttr", userPath+".encryption_password", *user.EncryptionPassword)
							}
						}
					}
				}
			}

			// Vacm configuration
			if o.SnmpConfig.V3Config.Vacm != nil {
				if len(o.SnmpConfig.V3Config.Vacm.Access) > 0 {
					checks.append(t, "TestCheckResourceAttr", "snmp_config.v3_config.vacm.access.#", fmt.Sprintf("%d", len(o.SnmpConfig.V3Config.Vacm.Access)))
					for i, access := range o.SnmpConfig.V3Config.Vacm.Access {
						accessPath := fmt.Sprintf("snmp_config.v3_config.vacm.access.%d", i)
						if access.GroupName != nil {
							checks.append(t, "TestCheckResourceAttr", accessPath+".group_name", *access.GroupName)
						}
						if len(access.PrefixList) > 0 {
							checks.append(t, "TestCheckResourceAttr", accessPath+".prefix_list.#", fmt.Sprintf("%d", len(access.PrefixList)))
							for j, prefix := range access.PrefixList {
								prefixPath := fmt.Sprintf("%s.prefix_list.%d", accessPath, j)
								if prefix.ContextPrefix != nil {
									checks.append(t, "TestCheckResourceAttr", prefixPath+".context_prefix", *prefix.ContextPrefix)
								}
								if prefix.NotifyView != nil {
									checks.append(t, "TestCheckResourceAttr", prefixPath+".notify_view", *prefix.NotifyView)
								}
								if prefix.ReadView != nil {
									checks.append(t, "TestCheckResourceAttr", prefixPath+".read_view", *prefix.ReadView)
								}
								if prefix.SecurityLevel != nil {
									checks.append(t, "TestCheckResourceAttr", prefixPath+".security_level", *prefix.SecurityLevel)
								}
								if prefix.SecurityModel != nil {
									checks.append(t, "TestCheckResourceAttr", prefixPath+".security_model", *prefix.SecurityModel)
								}
								if prefix.PrefixListType != nil {
									checks.append(t, "TestCheckResourceAttr", prefixPath+".type", *prefix.PrefixListType)
								}
								if prefix.WriteView != nil {
									checks.append(t, "TestCheckResourceAttr", prefixPath+".write_view", *prefix.WriteView)
								}
							}
						}
					}
				}

				if o.SnmpConfig.V3Config.Vacm.SecurityToGroup != nil {
					if o.SnmpConfig.V3Config.Vacm.SecurityToGroup.SecurityModel != nil {
						checks.append(t, "TestCheckResourceAttr", "snmp_config.v3_config.vacm.security_to_group.security_model", *o.SnmpConfig.V3Config.Vacm.SecurityToGroup.SecurityModel)
					}
					if len(o.SnmpConfig.V3Config.Vacm.SecurityToGroup.Snmpv3VacmContent) > 0 {
						checks.append(t, "TestCheckResourceAttr", "snmp_config.v3_config.vacm.security_to_group.content.#", fmt.Sprintf("%d", len(o.SnmpConfig.V3Config.Vacm.SecurityToGroup.Snmpv3VacmContent)))
						for i, content := range o.SnmpConfig.V3Config.Vacm.SecurityToGroup.Snmpv3VacmContent {
							contentPath := fmt.Sprintf("snmp_config.v3_config.vacm.security_to_group.content.%d", i)
							if content.Group != nil {
								checks.append(t, "TestCheckResourceAttr", contentPath+".group", *content.Group)
							}
							if content.SecurityName != nil {
								checks.append(t, "TestCheckResourceAttr", contentPath+".security_name", *content.SecurityName)
							}
						}
					}
				}
			}
		}

		// Views configuration
		if len(o.SnmpConfig.Views) > 0 {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.views.#", fmt.Sprintf("%d", len(o.SnmpConfig.Views)))
			for i, view := range o.SnmpConfig.Views {
				viewPath := fmt.Sprintf("snmp_config.views.%d", i)
				if view.ViewName != nil {
					checks.append(t, "TestCheckResourceAttr", viewPath+".view_name", *view.ViewName)
				}
				if view.Oid != nil {
					checks.append(t, "TestCheckResourceAttr", viewPath+".oid", *view.Oid)
				}
				if view.Include != nil {
					checks.append(t, "TestCheckResourceAttr", viewPath+".include", fmt.Sprintf("%t", *view.Include))
				}
			}
		}
	}

	if o.SwitchMatching != nil {
		if o.SwitchMatching.Enable != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_matching.enable", fmt.Sprintf("%t", *o.SwitchMatching.Enable))
		}
		if len(o.SwitchMatching.MatchingRules) > 0 {
			checks.append(t, "TestCheckResourceAttr", "switch_matching.rules.#", fmt.Sprintf("%d", len(o.SwitchMatching.MatchingRules)))
			for i, rule := range o.SwitchMatching.MatchingRules {
				basePath := fmt.Sprintf("switch_matching.rules.%d", i)
				if rule.Name != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".name", *rule.Name)
				}
				if rule.MatchModel != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".match_model", *rule.MatchModel)
				}
				if rule.MatchName != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".match_name", *rule.MatchName)
				}
				if rule.MatchNameOffset != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".match_name_offset", fmt.Sprintf("%d", int(*rule.MatchNameOffset)))
				}
				if rule.MatchRole != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".match_role", *rule.MatchRole)
				}
				if len(rule.AdditionalConfigCmds) > 0 {
					checks.append(t, "TestCheckResourceAttr", basePath+".additional_config_cmds.#", fmt.Sprintf("%d", len(rule.AdditionalConfigCmds)))
					for j, cmd := range rule.AdditionalConfigCmds {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.additional_config_cmds.%d", basePath, j), cmd)
					}
				}
				if rule.IpConfig != nil {
					if rule.IpConfig.Network != nil {
						checks.append(t, "TestCheckResourceAttr", basePath+".ip_config.network", *rule.IpConfig.Network)
					}
					if rule.IpConfig.IpConfigType != nil {
						checks.append(t, "TestCheckResourceAttr", basePath+".ip_config.type", *rule.IpConfig.IpConfigType)
					}
				}
				if rule.OobIpConfig != nil {
					if rule.OobIpConfig.OobIpConfigType != nil {
						checks.append(t, "TestCheckResourceAttr", basePath+".oob_ip_config.type", *rule.OobIpConfig.OobIpConfigType)
					}
					if rule.OobIpConfig.UseMgmtVrf != nil {
						checks.append(t, "TestCheckResourceAttr", basePath+".oob_ip_config.use_mgmt_vrf", fmt.Sprintf("%t", *rule.OobIpConfig.UseMgmtVrf))
					}
					if rule.OobIpConfig.UseMgmtVrfForHostOut != nil {
						checks.append(t, "TestCheckResourceAttr", basePath+".oob_ip_config.use_mgmt_vrf_for_host_out", fmt.Sprintf("%t", *rule.OobIpConfig.UseMgmtVrfForHostOut))
					}
				}
				if rule.StpConfig != nil {
					if rule.StpConfig.BridgePriority != nil {
						checks.append(t, "TestCheckResourceAttr", basePath+".stp_config.bridge_priority", *rule.StpConfig.BridgePriority)
					}
				}
				// PortConfig map
				if len(rule.PortConfig) > 0 {
					for portKey, portCfg := range rule.PortConfig {
						portPath := fmt.Sprintf("%s.port_config.%s", basePath, portKey)
						checks.append(t, "TestCheckResourceAttr", portPath+".usage", portCfg.Usage)
						if portCfg.Speed != nil {
							checks.append(t, "TestCheckResourceAttr", portPath+".speed", string(*portCfg.Speed))
						}
						if len(portCfg.Networks) > 0 {
							checks.append(t, "TestCheckResourceAttr", portPath+".networks.#", fmt.Sprintf("%d", len(portCfg.Networks)))
							for j, network := range portCfg.Networks {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.networks.%d", portPath, j), network)
							}
						}
					}
				}
			}
		}
	}

	if o.SwitchMgmt != nil {
		// Optional integer fields
		if o.SwitchMgmt.ApAffinityThreshold != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.ap_affinity_threshold", fmt.Sprintf("%d", int(*o.SwitchMgmt.ApAffinityThreshold)))
		}
		if o.SwitchMgmt.CliIdleTimeout != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.cli_idle_timeout", fmt.Sprintf("%d", int(*o.SwitchMgmt.CliIdleTimeout)))
		}
		if o.SwitchMgmt.ConfigRevertTimer != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.config_revert_timer", fmt.Sprintf("%d", int(*o.SwitchMgmt.ConfigRevertTimer)))
		}

		// Optional boolean fields
		if o.SwitchMgmt.DhcpOptionFqdn != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.dhcp_option_fqdn", fmt.Sprintf("%t", *o.SwitchMgmt.DhcpOptionFqdn))
		}
		if o.SwitchMgmt.DisableOobDownAlarm != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.disable_oob_down_alarm", fmt.Sprintf("%t", *o.SwitchMgmt.DisableOobDownAlarm))
		}
		if o.SwitchMgmt.FipsEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.fips_enabled", fmt.Sprintf("%t", *o.SwitchMgmt.FipsEnabled))
		}
		if o.SwitchMgmt.RemoveExistingConfigs != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.remove_existing_configs", fmt.Sprintf("%t", *o.SwitchMgmt.RemoveExistingConfigs))
		}
		if o.SwitchMgmt.UseMxedgeProxy != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.use_mxedge_proxy", fmt.Sprintf("%t", *o.SwitchMgmt.UseMxedgeProxy))
		}

		// Optional string fields
		if o.SwitchMgmt.CliBanner != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.cli_banner", *o.SwitchMgmt.CliBanner)
		}
		if o.SwitchMgmt.MxedgeProxyHost != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.mxedge_proxy_host", *o.SwitchMgmt.MxedgeProxyHost)
		}
		if o.SwitchMgmt.MxedgeProxyPort != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.mxedge_proxy_port", *o.SwitchMgmt.MxedgeProxyPort)
		}
		if o.SwitchMgmt.RootPassword != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.root_password", *o.SwitchMgmt.RootPassword)
		}

		// LocalAccounts map
		if len(o.SwitchMgmt.LocalAccounts) > 0 {
			for key, account := range o.SwitchMgmt.LocalAccounts {
				accountPath := fmt.Sprintf("switch_mgmt.local_accounts.%s", key)
				if account.Password != nil {
					checks.append(t, "TestCheckResourceAttr", accountPath+".password", *account.Password)
				}
				if account.Role != nil {
					checks.append(t, "TestCheckResourceAttr", accountPath+".role", *account.Role)
				}
			}
		}

		// ProtectRe configuration
		if o.SwitchMgmt.ProtectRe != nil {
			if o.SwitchMgmt.ProtectRe.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.protect_re.enabled", fmt.Sprintf("%t", *o.SwitchMgmt.ProtectRe.Enabled))
			}
			if o.SwitchMgmt.ProtectRe.HitCount != nil {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.protect_re.hit_count", fmt.Sprintf("%t", *o.SwitchMgmt.ProtectRe.HitCount))
			}

			if len(o.SwitchMgmt.ProtectRe.AllowedServices) > 0 {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.protect_re.allowed_services.#", fmt.Sprintf("%d", len(o.SwitchMgmt.ProtectRe.AllowedServices)))
				for i, service := range o.SwitchMgmt.ProtectRe.AllowedServices {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.protect_re.allowed_services.%d", i), service)
				}
			}

			if len(o.SwitchMgmt.ProtectRe.TrustedHosts) > 0 {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.protect_re.trusted_hosts.#", fmt.Sprintf("%d", len(o.SwitchMgmt.ProtectRe.TrustedHosts)))
				for i, host := range o.SwitchMgmt.ProtectRe.TrustedHosts {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.protect_re.trusted_hosts.%d", i), host)
				}
			}

			// Custom protect_re rules
			if len(o.SwitchMgmt.ProtectRe.Custom) > 0 {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.protect_re.custom.#", fmt.Sprintf("%d", len(o.SwitchMgmt.ProtectRe.Custom)))
				for i, custom := range o.SwitchMgmt.ProtectRe.Custom {
					customPath := fmt.Sprintf("switch_mgmt.protect_re.custom.%d", i)

					// Optional string fields
					if custom.PortRange != nil {
						checks.append(t, "TestCheckResourceAttr", customPath+".port_range", *custom.PortRange)
					}
					if custom.Protocol != nil {
						checks.append(t, "TestCheckResourceAttr", customPath+".protocol", *custom.Protocol)
					}

					// Subnets list
					if len(custom.Subnets) > 0 {
						checks.append(t, "TestCheckResourceAttr", customPath+".subnets.#", fmt.Sprintf("%d", len(custom.Subnets)))
						for j, subnet := range custom.Subnets {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.subnets.%d", customPath, j), subnet)
						}
					}
				}
			}
		}

		// Tacacs configuration
		if o.SwitchMgmt.Tacacs != nil {
			// Optional boolean field
			if o.SwitchMgmt.Tacacs.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.tacacs.enabled", fmt.Sprintf("%t", *o.SwitchMgmt.Tacacs.Enabled))
			}

			// Optional string fields
			if o.SwitchMgmt.Tacacs.DefaultRole != nil {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.tacacs.default_role", *o.SwitchMgmt.Tacacs.DefaultRole)
			}
			if o.SwitchMgmt.Tacacs.Network != nil {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.tacacs.network", *o.SwitchMgmt.Tacacs.Network)
			}

			// TacacctServers (accounting servers) list
			if len(o.SwitchMgmt.Tacacs.TacacctServers) > 0 {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.tacacs.acct_servers.#", fmt.Sprintf("%d", len(o.SwitchMgmt.Tacacs.TacacctServers)))
				for i, server := range o.SwitchMgmt.Tacacs.TacacctServers {
					serverPath := fmt.Sprintf("switch_mgmt.tacacs.acct_servers.%d", i)

					// Optional string fields
					if server.Host != nil {
						checks.append(t, "TestCheckResourceAttr", serverPath+".host", *server.Host)
					}
					if server.Port != nil {
						checks.append(t, "TestCheckResourceAttr", serverPath+".port", *server.Port)
					}
					if server.Secret != nil {
						checks.append(t, "TestCheckResourceAttr", serverPath+".secret", *server.Secret)
					}

					// Optional integer field
					if server.Timeout != nil {
						checks.append(t, "TestCheckResourceAttr", serverPath+".timeout", fmt.Sprintf("%d", int(*server.Timeout)))
					}
				}
			}

			// TacplusServers (authentication servers) list
			if len(o.SwitchMgmt.Tacacs.TacplusServers) > 0 {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.tacacs.tacplus_servers.#", fmt.Sprintf("%d", len(o.SwitchMgmt.Tacacs.TacplusServers)))
				for i, server := range o.SwitchMgmt.Tacacs.TacplusServers {
					serverPath := fmt.Sprintf("switch_mgmt.tacacs.tacplus_servers.%d", i)

					// Optional string fields
					if server.Host != nil {
						checks.append(t, "TestCheckResourceAttr", serverPath+".host", *server.Host)
					}
					if server.Port != nil {
						checks.append(t, "TestCheckResourceAttr", serverPath+".port", *server.Port)
					}
					if server.Secret != nil {
						checks.append(t, "TestCheckResourceAttr", serverPath+".secret", *server.Secret)
					}

					// Optional integer field
					if server.Timeout != nil {
						checks.append(t, "TestCheckResourceAttr", serverPath+".timeout", fmt.Sprintf("%d", int(*server.Timeout)))
					}
				}
			}
		}
	}

	if o.VrfConfig != nil {
		if o.VrfConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "vrf_config.enabled", fmt.Sprintf("%t", *o.VrfConfig.Enabled))
		}
	}

	// Map attributes
	if len(o.AclTags) > 0 {
		for key, tag := range o.AclTags {
			basePath := fmt.Sprintf("acl_tags.%s", key)

			// Check AclTagsType (required field)
			checks.append(t, "TestCheckResourceAttr", basePath+".type", tag.AclTagsType)

			// Optional list fields
			if len(tag.EtherTypes) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".ether_types.#", fmt.Sprintf("%d", len(tag.EtherTypes)))
				for i, etherType := range tag.EtherTypes {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.ether_types.%d", basePath, i), etherType)
				}
			}
			if len(tag.Macs) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".macs.#", fmt.Sprintf("%d", len(tag.Macs)))
				for i, mac := range tag.Macs {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.macs.%d", basePath, i), mac)
				}
			}
			if len(tag.Subnets) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".subnets.#", fmt.Sprintf("%d", len(tag.Subnets)))
				for i, subnet := range tag.Subnets {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.subnets.%d", basePath, i), subnet)
				}
			}

			// Optional pointer fields
			if tag.GbpTag != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".gbp_tag", fmt.Sprintf("%d", int(*tag.GbpTag)))
			}
			if tag.Network != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".network", *tag.Network)
			}
			if tag.PortUsage != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".port_usage", *tag.PortUsage)
			}
			if tag.RadiusGroup != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".radius_group", *tag.RadiusGroup)
			}

			// Check Specs list
			if len(tag.Specs) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".specs.#", fmt.Sprintf("%d", len(tag.Specs)))
				for i, spec := range tag.Specs {
					specPath := fmt.Sprintf("%s.specs.%d", basePath, i)
					if spec.PortRange != nil {
						checks.append(t, "TestCheckResourceAttr", specPath+".port_range", *spec.PortRange)
					}
					if spec.Protocol != nil {
						checks.append(t, "TestCheckResourceAttr", specPath+".protocol", *spec.Protocol)
					}
				}
			}
		}
	}

	if len(o.ExtraRoutes) > 0 {
		for key, route := range o.ExtraRoutes {
			basePath := fmt.Sprintf("extra_routes.%s", key)

			// Check required via field
			checks.append(t, "TestCheckResourceAttr", basePath+".via", route.Via)

			// Optional boolean fields
			if route.Discard != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".discard", fmt.Sprintf("%t", *route.Discard))
			}
			if route.NoResolve != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".no_resolve", fmt.Sprintf("%t", *route.NoResolve))
			}

			// Optional integer fields
			if route.Metric != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".metric", fmt.Sprintf("%d", int(*route.Metric)))
			}
			if route.Preference != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".preference", fmt.Sprintf("%d", int(*route.Preference)))
			}

			// NextQualified map
			if len(route.NextQualified) > 0 {
				for nqKey, nq := range route.NextQualified {
					nqPath := fmt.Sprintf("%s.next_qualified.%s", basePath, nqKey)
					if nq.Metric != nil {
						checks.append(t, "TestCheckResourceAttr", nqPath+".metric", fmt.Sprintf("%d", int(*nq.Metric)))
					}
					if nq.Preference != nil {
						checks.append(t, "TestCheckResourceAttr", nqPath+".preference", fmt.Sprintf("%d", int(*nq.Preference)))
					}
				}
			}
		}
	}

	if len(o.ExtraRoutes6) > 0 {
		for key, route6 := range o.ExtraRoutes6 {
			basePath := fmt.Sprintf("extra_routes6.%s", key)

			// Check required via field
			checks.append(t, "TestCheckResourceAttr", basePath+".via", route6.Via)

			// Optional boolean fields
			if route6.Discard != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".discard", fmt.Sprintf("%t", *route6.Discard))
			}
			if route6.NoResolve != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".no_resolve", fmt.Sprintf("%t", *route6.NoResolve))
			}

			// Optional integer fields
			if route6.Metric != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".metric", fmt.Sprintf("%d", int(*route6.Metric)))
			}
			if route6.Preference != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".preference", fmt.Sprintf("%d", int(*route6.Preference)))
			}

			// NextQualified map
			if len(route6.NextQualified) > 0 {
				for nqKey, nq := range route6.NextQualified {
					nqPath := fmt.Sprintf("%s.next_qualified.%s", basePath, nqKey)
					if nq.Metric != nil {
						checks.append(t, "TestCheckResourceAttr", nqPath+".metric", fmt.Sprintf("%d", int(*nq.Metric)))
					}
					if nq.Preference != nil {
						checks.append(t, "TestCheckResourceAttr", nqPath+".preference", fmt.Sprintf("%d", int(*nq.Preference)))
					}
				}
			}
		}
	}

	if len(o.Networks) > 0 {
		for key, network := range o.Networks {
			basePath := fmt.Sprintf("networks.%s", key)

			// Check required vlan_id field
			checks.append(t, "TestCheckResourceAttr", basePath+".vlan_id", network.VlanId)

			// Optional string fields
			if network.Gateway != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".gateway", *network.Gateway)
			}
			if network.Gateway6 != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".gateway6", *network.Gateway6)
			}
			if network.IsolationVlanId != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".isolation_vlan_id", *network.IsolationVlanId)
			}
			if network.Subnet != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".subnet", *network.Subnet)
			}
			if network.Subnet6 != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".subnet6", *network.Subnet6)
			}

			// Optional boolean field
			if network.Isolation != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".isolation", fmt.Sprintf("%t", *network.Isolation))
			}
		}
	}

	if len(o.OspfAreas) > 0 {
		for key, area := range o.OspfAreas {
			basePath := fmt.Sprintf("ospf_areas.%s", key)

			// Optional boolean field
			if area.IncludeLoopback != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".include_loopback", fmt.Sprintf("%t", *area.IncludeLoopback))
			}

			// Optional string field
			if area.OspfAreasType != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".type", *area.OspfAreasType)
			}

			// OspfNetworks map
			if len(area.OspfNetworks) > 0 {
				for netKey, network := range area.OspfNetworks {
					netPath := fmt.Sprintf("%s.networks.%s", basePath, netKey)

					// Optional string fields
					if network.AuthPassword != nil {
						checks.append(t, "TestCheckResourceAttr", netPath+".auth_password", *network.AuthPassword)
					}
					if network.AuthType != nil {
						checks.append(t, "TestCheckResourceAttr", netPath+".auth_type", *network.AuthType)
					}
					if network.ExportPolicy != nil {
						checks.append(t, "TestCheckResourceAttr", netPath+".export_policy", *network.ExportPolicy)
					}
					if network.ImportPolicy != nil {
						checks.append(t, "TestCheckResourceAttr", netPath+".import_policy", *network.ImportPolicy)
					}
					if network.InterfaceType != nil {
						checks.append(t, "TestCheckResourceAttr", netPath+".interface_type", *network.InterfaceType)
					}

					// Optional integer fields
					if network.BfdMinimumInterval != nil {
						checks.append(t, "TestCheckResourceAttr", netPath+".bfd_minimum_interval", fmt.Sprintf("%d", int(*network.BfdMinimumInterval)))
					}
					if network.DeadInterval != nil {
						checks.append(t, "TestCheckResourceAttr", netPath+".dead_interval", fmt.Sprintf("%d", int(*network.DeadInterval)))
					}
					if network.HelloInterval != nil {
						checks.append(t, "TestCheckResourceAttr", netPath+".hello_interval", fmt.Sprintf("%d", int(*network.HelloInterval)))
					}
					if network.Metric != nil {
						checks.append(t, "TestCheckResourceAttr", netPath+".metric", fmt.Sprintf("%d", int(*network.Metric)))
					}

					// Optional boolean fields
					if network.NoReadvertiseToOverlay != nil {
						checks.append(t, "TestCheckResourceAttr", netPath+".no_readvertise_to_overlay", fmt.Sprintf("%t", *network.NoReadvertiseToOverlay))
					}
					if network.Passive != nil {
						checks.append(t, "TestCheckResourceAttr", netPath+".passive", fmt.Sprintf("%t", *network.Passive))
					}

					// AuthKeys map
					if len(network.AuthKeys) > 0 {
						for authKey, authValue := range network.AuthKeys {
							authPath := fmt.Sprintf("%s.auth_keys.%s", netPath, authKey)
							checks.append(t, "TestCheckResourceAttr", authPath, authValue)
						}
					}
				}
			}
		}
	}

	if len(o.PortMirroring) > 0 {
		for key, mirror := range o.PortMirroring {
			basePath := fmt.Sprintf("port_mirroring.%s", key)

			// Optional string fields
			if mirror.OutputIpAddress != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".output_ip_address", *mirror.OutputIpAddress)
			}
			if mirror.OutputNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".output_network", *mirror.OutputNetwork)
			}
			if mirror.OutputPortId != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".output_port_id", *mirror.OutputPortId)
			}

			// Optional list fields
			if len(mirror.InputNetworksIngress) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".input_networks_ingress.#", fmt.Sprintf("%d", len(mirror.InputNetworksIngress)))
				for i, network := range mirror.InputNetworksIngress {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.input_networks_ingress.%d", basePath, i), network)
				}
			}

			if len(mirror.InputPortIdsEgress) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".input_port_ids_egress.#", fmt.Sprintf("%d", len(mirror.InputPortIdsEgress)))
				for i, portId := range mirror.InputPortIdsEgress {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.input_port_ids_egress.%d", basePath, i), portId)
				}
			}

			if len(mirror.InputPortIdsIngress) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".input_port_ids_ingress.#", fmt.Sprintf("%d", len(mirror.InputPortIdsIngress)))
				for i, portId := range mirror.InputPortIdsIngress {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.input_port_ids_ingress.%d", basePath, i), portId)
				}
			}
		}
	}

	if len(o.PortUsages) > 0 {
		for key, usage := range o.PortUsages {
			basePath := fmt.Sprintf("port_usages.%s", key)

			// Add comprehensive port usage checks
			if usage.AllNetworks != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".all_networks", fmt.Sprintf("%t", *usage.AllNetworks))
			}
			if usage.AllowDhcpd != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".allow_dhcpd", fmt.Sprintf("%t", *usage.AllowDhcpd))
			}
			if usage.AllowMultipleSupplicants != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".allow_multiple_supplicants", fmt.Sprintf("%t", *usage.AllowMultipleSupplicants))
			}
			if usage.BypassAuthWhenServerDown != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".bypass_auth_when_server_down", fmt.Sprintf("%t", *usage.BypassAuthWhenServerDown))
			}
			if usage.BypassAuthWhenServerDownForUnknownClient != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".bypass_auth_when_server_down_for_unknown_client", fmt.Sprintf("%t", *usage.BypassAuthWhenServerDownForUnknownClient))
			}
			if usage.BypassAuthWhenServerDownForVoip != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".bypass_auth_when_server_down_for_voip", fmt.Sprintf("%t", *usage.BypassAuthWhenServerDownForVoip))
			}
			if usage.Description != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".description", *usage.Description)
			}
			if usage.DisableAutoneg != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".disable_autoneg", fmt.Sprintf("%t", *usage.DisableAutoneg))
			}
			if usage.Disabled != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".disabled", fmt.Sprintf("%t", *usage.Disabled))
			}
			if usage.Duplex != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".duplex", *usage.Duplex)
			}
			if len(usage.DynamicVlanNetworks) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".dynamic_vlan_networks.#", fmt.Sprintf("%d", len(usage.DynamicVlanNetworks)))
				for i, network := range usage.DynamicVlanNetworks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.dynamic_vlan_networks.%d", basePath, i), network)
				}
			}
			if usage.EnableMacAuth != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".enable_mac_auth", fmt.Sprintf("%t", *usage.EnableMacAuth))
			}
			if usage.EnableQos != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".enable_qos", fmt.Sprintf("%t", *usage.EnableQos))
			}
			if usage.GuestNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".guest_network", *usage.GuestNetwork)
			}
			if usage.InterSwitchLink != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".inter_switch_link", fmt.Sprintf("%t", *usage.InterSwitchLink))
			}
			if usage.MacAuthOnly != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".mac_auth_only", fmt.Sprintf("%t", *usage.MacAuthOnly))
			}
			if usage.MacAuthPreferred != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".mac_auth_preferred", fmt.Sprintf("%t", *usage.MacAuthPreferred))
			}
			if usage.MacAuthProtocol != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".mac_auth_protocol", *usage.MacAuthProtocol)
			}
			if usage.MacLimit != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".mac_limit", *usage.MacLimit)
			}
			if usage.Mode != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".mode", *usage.Mode)
			}
			if usage.Mtu != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".mtu", *usage.Mtu)
			}
			if len(usage.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".networks.#", fmt.Sprintf("%d", len(usage.Networks)))
				for i, network := range usage.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.networks.%d", basePath, i), network)
				}
			}
			if usage.PersistMac != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".persist_mac", fmt.Sprintf("%t", *usage.PersistMac))
			}
			if usage.PoeDisabled != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".poe_disabled", fmt.Sprintf("%t", *usage.PoeDisabled))
			}
			if usage.PoePriority != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".poe_priority", *usage.PoePriority)
			}
			if usage.PortAuth != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".port_auth", *usage.PortAuth)
			}
			if usage.PortNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".port_network", *usage.PortNetwork)
			}
			if usage.ReauthInterval != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".reauth_interval", *usage.ReauthInterval)
			}
			if usage.ResetDefaultWhen != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".reset_default_when", *usage.ResetDefaultWhen)
			}
			if len(usage.Rules) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".rules.#", fmt.Sprintf("%d", len(usage.Rules)))
				for i, rule := range usage.Rules {
					rulePath := fmt.Sprintf("%s.rules.%d", basePath, i)

					// Check required src field
					checks.append(t, "TestCheckResourceAttr", rulePath+".src", rule.Src)

					// Optional string fields
					if rule.Equals != nil {
						checks.append(t, "TestCheckResourceAttr", rulePath+".equals", *rule.Equals)
					}
					if rule.Expression != nil {
						checks.append(t, "TestCheckResourceAttr", rulePath+".expression", *rule.Expression)
					}
					if rule.Usage != nil {
						checks.append(t, "TestCheckResourceAttr", rulePath+".usage", *rule.Usage)
					}

					// EqualsAny list
					if len(rule.EqualsAny) > 0 {
						checks.append(t, "TestCheckResourceAttr", rulePath+".equals_any.#", fmt.Sprintf("%d", len(rule.EqualsAny)))
						for j, equalsAny := range rule.EqualsAny {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.equals_any.%d", rulePath, j), equalsAny)
						}
					}
				}
			}
			if usage.ServerFailNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".server_fail_network", *usage.ServerFailNetwork)
			}
			if usage.ServerRejectNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".server_reject_network", *usage.ServerRejectNetwork)
			}
			if usage.Speed != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".speed", *usage.Speed)
			}
			if usage.StormControl != nil {
				if usage.StormControl.DisablePort != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".storm_control.disable_port", fmt.Sprintf("%t", *usage.StormControl.DisablePort))
				}
				if usage.StormControl.NoBroadcast != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".storm_control.no_broadcast", fmt.Sprintf("%t", *usage.StormControl.NoBroadcast))
				}
				if usage.StormControl.NoMulticast != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".storm_control.no_multicast", fmt.Sprintf("%t", *usage.StormControl.NoMulticast))
				}
				if usage.StormControl.NoRegisteredMulticast != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".storm_control.no_registered_multicast", fmt.Sprintf("%t", *usage.StormControl.NoRegisteredMulticast))
				}
				if usage.StormControl.NoUnknownUnicast != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".storm_control.no_unknown_unicast", fmt.Sprintf("%t", *usage.StormControl.NoUnknownUnicast))
				}
				if usage.StormControl.Percentage != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".storm_control.percentage", fmt.Sprintf("%d", int(*usage.StormControl.Percentage)))
				}
			}
			if usage.StpDisable != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".stp_disable", fmt.Sprintf("%t", *usage.StpDisable))
			}
			if usage.StpEdge != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".stp_edge", fmt.Sprintf("%t", *usage.StpEdge))
			}
			if usage.StpNoRootPort != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".stp_no_root_port", fmt.Sprintf("%t", *usage.StpNoRootPort))
			}
			if usage.StpP2p != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".stp_p2p", fmt.Sprintf("%t", *usage.StpP2p))
			}
			if usage.StpRequired != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".stp_required", fmt.Sprintf("%t", *usage.StpRequired))
			}
			if usage.UiEvpntopoId != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".ui_evpntopo_id", *usage.UiEvpntopoId)
			}
			if usage.UseVstp != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".use_vstp", fmt.Sprintf("%t", *usage.UseVstp))
			}
			if usage.VoipNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".voip_network", *usage.VoipNetwork)
			}
		}
	}

	if len(o.VrfInstances) > 0 {
		for key, instance := range o.VrfInstances {
			basePath := fmt.Sprintf("vrf_instances.%s", key)

			if len(instance.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".networks.#", fmt.Sprintf("%d", len(instance.Networks)))
				for i, network := range instance.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.networks.%d", basePath, i), network)
				}
			}

			if len(instance.VrfExtraRoutes) > 0 {
				for routeKey, route := range instance.VrfExtraRoutes {
					routePath := fmt.Sprintf("%s.extra_routes.%s", basePath, routeKey)
					// Check required via field
					checks.append(t, "TestCheckResourceAttr", routePath+".via", route.Via)
				}
			}
		}
	}

	return checks
}
