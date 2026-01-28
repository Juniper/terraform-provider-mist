package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_service"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgServiceModel(t *testing.T) {
	type testStep struct {
		config OrgServiceModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgServiceModel{
						OrgId: GetTestOrgId(),
						Name:  "test-service",
					},
				},
			},
		},
	}

	// Load fixture data following the org_wlan pattern
	b, err := os.ReadFile("fixtures/org_service_resource/org_service_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		fixtureOrgServiceModel := OrgServiceModel{}
		err = hcl.Decode(&fixtureOrgServiceModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgServiceModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgServiceModel,
				},
			},
		}
	}

	resourceType := "org_service"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_service.OrgServiceResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				// Generate Terraform configuration using automated HCL generation
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&step.config, f.Body())
				configStr := Render(resourceType, tName, string(f.Bytes()))

				checks := step.config.testChecks(t, resourceType, tName, tracker)
				chkLog := checks.string()
				stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

				// log config and checks here
				t.Logf("\n// ------ begin config for %s ------\n%s// -------- end config for %s ------\n\n", stepName, configStr, stepName)
				t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end checks for %s ------\n\n", stepName, chkLog, stepName)

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

func (s *OrgServiceModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.SetTracker(tracker)

	// Check fields in struct order
	// 1. Addresses array
	if len(s.Addresses) > 0 {
		checks.append(t, "TestCheckResourceAttr", "addresses.#", fmt.Sprintf("%d", len(s.Addresses)))
		for i, address := range s.Addresses {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("addresses.%d", i), address)
		}
	}

	// 2. AppCategories array
	if len(s.AppCategories) > 0 {
		checks.append(t, "TestCheckResourceAttr", "app_categories.#", fmt.Sprintf("%d", len(s.AppCategories)))
		for i, category := range s.AppCategories {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("app_categories.%d", i), category)
		}
	}

	// 3. AppSubcategories array
	if len(s.AppSubcategories) > 0 {
		checks.append(t, "TestCheckResourceAttr", "app_subcategories.#", fmt.Sprintf("%d", len(s.AppSubcategories)))
		for i, subcategory := range s.AppSubcategories {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("app_subcategories.%d", i), subcategory)
		}
	}

	// 4. Apps array
	if len(s.Apps) > 0 {
		checks.append(t, "TestCheckResourceAttr", "apps.#", fmt.Sprintf("%d", len(s.Apps)))
		for i, app := range s.Apps {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("apps.%d", i), app)
		}
	}

	// 5. ClientLimitDown
	if s.ClientLimitDown != nil {
		checks.append(t, "TestCheckResourceAttr", "client_limit_down", fmt.Sprintf("%d", *s.ClientLimitDown))
	}

	// 6. ClientLimitUp
	if s.ClientLimitUp != nil {
		checks.append(t, "TestCheckResourceAttr", "client_limit_up", fmt.Sprintf("%d", *s.ClientLimitUp))
	}

	// 7. Description
	if s.Description != nil {
		checks.append(t, "TestCheckResourceAttr", "description", *s.Description)
	}

	// 8. Dscp
	if s.Dscp != nil {
		checks.append(t, "TestCheckResourceAttr", "dscp", *s.Dscp)
	}

	// 9. FailoverPolicy
	if s.FailoverPolicy != nil {
		checks.append(t, "TestCheckResourceAttr", "failover_policy", *s.FailoverPolicy)
	}

	// 10. Hostnames array
	if len(s.Hostnames) > 0 {
		checks.append(t, "TestCheckResourceAttr", "hostnames.#", fmt.Sprintf("%d", len(s.Hostnames)))
		for i, hostname := range s.Hostnames {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("hostnames.%d", i), hostname)
		}
	}

	// 11. Id (computed-only)
	checks.append(t, "TestCheckResourceAttrSet", "id")

	// 12. MaxJitter
	if s.MaxJitter != nil {
		checks.append(t, "TestCheckResourceAttr", "max_jitter", *s.MaxJitter)
	}

	// 13. MaxLatency
	if s.MaxLatency != nil {
		checks.append(t, "TestCheckResourceAttr", "max_latency", *s.MaxLatency)
	}

	// 14. MaxLoss
	if s.MaxLoss != nil {
		checks.append(t, "TestCheckResourceAttr", "max_loss", *s.MaxLoss)
	}

	// 15. Name (required)
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)

	// 16. OrgId (required top-level field)
	checks.append(t, "TestCheckResourceAttr", "org_id", s.OrgId)

	// 17. ServiceLimitDown
	if s.ServiceLimitDown != nil {
		checks.append(t, "TestCheckResourceAttr", "service_limit_down", fmt.Sprintf("%d", *s.ServiceLimitDown))
	}

	// 18. ServiceLimitUp
	if s.ServiceLimitUp != nil {
		checks.append(t, "TestCheckResourceAttr", "service_limit_up", fmt.Sprintf("%d", *s.ServiceLimitUp))
	}

	// 19. SleEnabled
	if s.SleEnabled != nil {
		checks.append(t, "TestCheckResourceAttr", "sle_enabled", fmt.Sprintf("%t", *s.SleEnabled))
	}

	// 20. Specs array (nested objects)
	if len(s.Specs) > 0 {
		checks.append(t, "TestCheckResourceAttr", "specs.#", fmt.Sprintf("%d", len(s.Specs)))
		for i, spec := range s.Specs {
			if spec.Protocol != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("specs.%d.protocol", i), *spec.Protocol)
				// Only check port_range if protocol is tcp or udp (due to schema validation)
				if spec.PortRange != nil && (*spec.Protocol == "tcp" || *spec.Protocol == "udp") {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("specs.%d.port_range", i), *spec.PortRange)
				}
			}
		}
	}

	// 21. SsrRelaxedTcpStateEnforcement
	if s.SsrRelaxedTcpStateEnforcement != nil {
		checks.append(t, "TestCheckResourceAttr", "ssr_relaxed_tcp_state_enforcement", fmt.Sprintf("%t", *s.SsrRelaxedTcpStateEnforcement))
	}

	// 22. TrafficClass
	if s.TrafficClass != nil {
		checks.append(t, "TestCheckResourceAttr", "traffic_class", *s.TrafficClass)
	}

	// 23. TrafficType
	if s.TrafficType != nil {
		checks.append(t, "TestCheckResourceAttr", "traffic_type", *s.TrafficType)
	}

	// 24. Type
	if s.Type != nil {
		checks.append(t, "TestCheckResourceAttr", "type", *s.Type)
	}

	// 25. Urls array
	if len(s.Urls) > 0 {
		checks.append(t, "TestCheckResourceAttr", "urls.#", fmt.Sprintf("%d", len(s.Urls)))
		for i, url := range s.Urls {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("urls.%d", i), url)
		}
	}

	return checks
}
