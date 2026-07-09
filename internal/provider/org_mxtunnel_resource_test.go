package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_mxtunnel"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgMxtunnelModel(t *testing.T) {
	type testStep struct {
		config OrgMxtunnelModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgMxtunnelModel{
						Name:  "Test Org Mxtunnel",
						OrgId: GetTestOrgId(),
					},
				},
			},
		},
	}

	fixtures, err := os.ReadFile("fixtures/org_mxtunnel_resource/org_mxtunnel_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	for i, fixture := range strings.Split(string(fixtures), "␞") {
		fixtureOrgMxtunnelModel := OrgMxtunnelModel{}
		err = hcl.Decode(&fixtureOrgMxtunnelModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgMxtunnelModel.OrgId = GetTestOrgId()
		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgMxtunnelModel,
				},
			},
		}
	}

	resourceType := "org_mxtunnel"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_mxtunnel.OrgMxtunnelResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())

				combinedConfig := Render(resourceType, tName, string(f.Bytes()))

				checks := config.testChecks(t, resourceType, tName, tracker)
				chkLog := checks.string()
				stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

				t.Logf("\n// ------ begin config for %s ------\n%s// -------- end config for %s ------\n\n", stepName, combinedConfig, stepName)
				t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end checks for %s ------\n\n", stepName, chkLog, stepName)

				steps[i] = resource.TestStep{
					Config: combinedConfig,
					Check:  resource.ComposeAggregateTestCheckFunc(checks.checks...),
				}
			}

			resource.Test(t, resource.TestCase{
				IsUnitTest:               true,
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps:                    steps,
			})
		})
	}
	if tracker != nil {
		tracker.FieldCoverageReport(t)
	}
}

func (o *OrgMxtunnelModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	// Required fields
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)

	// Optional basic fields
	if o.HelloInterval != nil {
		checks.append(t, "TestCheckResourceAttr", "hello_interval", fmt.Sprintf("%d", *o.HelloInterval))
	}
	if o.HelloRetries != nil {
		checks.append(t, "TestCheckResourceAttr", "hello_retries", fmt.Sprintf("%d", *o.HelloRetries))
	}
	if o.Mtu != nil {
		checks.append(t, "TestCheckResourceAttr", "mtu", fmt.Sprintf("%d", *o.Mtu))
	}
	if o.Protocol != nil {
		checks.append(t, "TestCheckResourceAttr", "protocol", *o.Protocol)
	}
	if len(o.VlanIds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "vlan_ids.#", fmt.Sprintf("%d", len(o.VlanIds)))
	}
	if len(o.AnchorMxtunnelIds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "anchor_mxtunnel_ids.#", fmt.Sprintf("%d", len(o.AnchorMxtunnelIds)))
	}
	if len(o.MxclusterIds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "mxcluster_ids.#", fmt.Sprintf("%d", len(o.MxclusterIds)))
	}

	// auto_preemption
	if o.AutoPreemption != nil {
		if o.AutoPreemption.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "auto_preemption.enabled", fmt.Sprintf("%t", *o.AutoPreemption.Enabled))
		}
		if o.AutoPreemption.DayOfWeek != nil {
			checks.append(t, "TestCheckResourceAttr", "auto_preemption.day_of_week", *o.AutoPreemption.DayOfWeek)
		}
		if o.AutoPreemption.TimeOfDay != nil {
			checks.append(t, "TestCheckResourceAttr", "auto_preemption.time_of_day", *o.AutoPreemption.TimeOfDay)
		}
	}

	// ipsec
	if o.Ipsec != nil {
		if o.Ipsec.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "ipsec.enabled", fmt.Sprintf("%t", *o.Ipsec.Enabled))
		}
		if o.Ipsec.SplitTunnel != nil {
			checks.append(t, "TestCheckResourceAttr", "ipsec.split_tunnel", fmt.Sprintf("%t", *o.Ipsec.SplitTunnel))
		}
		if o.Ipsec.UseMxedge != nil {
			checks.append(t, "TestCheckResourceAttr", "ipsec.use_mxedge", fmt.Sprintf("%t", *o.Ipsec.UseMxedge))
		}
		if len(o.Ipsec.DnsServers) > 0 {
			checks.append(t, "TestCheckResourceAttr", "ipsec.dns_servers.#", fmt.Sprintf("%d", len(o.Ipsec.DnsServers)))
		}
		if len(o.Ipsec.DnsSuffix) > 0 {
			checks.append(t, "TestCheckResourceAttr", "ipsec.dns_suffix.#", fmt.Sprintf("%d", len(o.Ipsec.DnsSuffix)))
		}
		if len(o.Ipsec.ExtraRoutes) > 0 {
			checks.append(t, "TestCheckResourceAttr", "ipsec.extra_routes.#", fmt.Sprintf("%d", len(o.Ipsec.ExtraRoutes)))
			for i, route := range o.Ipsec.ExtraRoutes {
				if route.Dest != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ipsec.extra_routes.%d.dest", i), *route.Dest)
				}
				if route.NextHop != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ipsec.extra_routes.%d.next_hop", i), *route.NextHop)
				}
			}
		}
	}

	return checks
}
