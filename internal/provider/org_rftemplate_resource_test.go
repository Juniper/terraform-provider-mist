package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_rftemplate"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgRftemplateModel(t *testing.T) {
	type testStep struct {
		config OrgRftemplateModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgRftemplateModel{
						OrgId: GetTestOrgId(),
						Name:  "TestRfTemplate",
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_rftemplate_resource/org_rftemplate_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		fixtureOrgRftemplateModel := OrgRftemplateModel{}
		err = hcl.Decode(&fixtureOrgRftemplateModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgRftemplateModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgRftemplateModel,
				},
			},
		}
	}

	resourceType := "org_rftemplate"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_rftemplate.OrgRftemplateResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {

			// Create single-step tests with generated config
			steps := make([]resource.TestStep, len(tCase.steps))

			for i, step := range tCase.steps {
				// Generate Terraform configuration using automated HCL generation
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&step.config, f.Body())
				configStr := Render(resourceType, tName, string(f.Bytes()))

				// Generate checks for the rftemplate resource
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
}

func (o *OrgRftemplateModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.SetTracker(tracker)

	// Always check required fields
	checks.append(t, "TestCheckResourceAttrSet", "id")
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)

	// Check optional top-level fields
	if o.AntGain24 != nil {
		checks.append(t, "TestCheckResourceAttr", "ant_gain_24", fmt.Sprintf("%d", *o.AntGain24))
	}
	if o.AntGain5 != nil {
		checks.append(t, "TestCheckResourceAttr", "ant_gain_5", fmt.Sprintf("%d", *o.AntGain5))
	}
	if o.AntGain6 != nil {
		checks.append(t, "TestCheckResourceAttr", "ant_gain_6", fmt.Sprintf("%d", *o.AntGain6))
	}
	if o.Band24Usage != nil {
		checks.append(t, "TestCheckResourceAttr", "band_24_usage", *o.Band24Usage)
	}
	if o.CountryCode != nil {
		checks.append(t, "TestCheckResourceAttr", "country_code", *o.CountryCode)
	}
	if o.ScanningEnabled != nil {
		checks.append(t, "TestCheckResourceAttr", "scanning_enabled", fmt.Sprintf("%t", *o.ScanningEnabled))
	}

	// Check band_24 nested fields
	if o.Band24 != nil {
		band := o.Band24
		prefix := "band_24"

		if band.AllowRrmDisable != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".allow_rrm_disable", fmt.Sprintf("%t", *band.AllowRrmDisable))
		}
		if band.AntGain != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".ant_gain", fmt.Sprintf("%d", *band.AntGain))
		}
		if band.AntennaMode != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".antenna_mode", *band.AntennaMode)
		}
		if band.Bandwidth != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".bandwidth", fmt.Sprintf("%d", *band.Bandwidth))
		}
		if len(band.Channels) > 0 {
			checks.append(t, "TestCheckResourceAttr", prefix+".channels.#", fmt.Sprintf("%d", len(band.Channels)))
			for i, ch := range band.Channels {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.channels.%d", prefix, i), fmt.Sprintf("%d", ch))
			}
		}
		if band.Disabled != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".disabled", fmt.Sprintf("%t", *band.Disabled))
		}
		if band.Power != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".power", fmt.Sprintf("%d", *band.Power))
		}
		if band.PowerMax != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".power_max", fmt.Sprintf("%d", *band.PowerMax))
		}
		if band.PowerMin != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".power_min", fmt.Sprintf("%d", *band.PowerMin))
		}
		if band.Preamble != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".preamble", *band.Preamble)
		}
	}

	// Check band_5 nested fields
	if o.Band5 != nil {
		band := o.Band5
		prefix := "band_5"

		if band.AllowRrmDisable != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".allow_rrm_disable", fmt.Sprintf("%t", *band.AllowRrmDisable))
		}
		if band.AntGain != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".ant_gain", fmt.Sprintf("%d", *band.AntGain))
		}
		if band.AntennaMode != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".antenna_mode", *band.AntennaMode)
		}
		if band.Bandwidth != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".bandwidth", fmt.Sprintf("%d", *band.Bandwidth))
		}
		if len(band.Channels) > 0 {
			checks.append(t, "TestCheckResourceAttr", prefix+".channels.#", fmt.Sprintf("%d", len(band.Channels)))
			for i, ch := range band.Channels {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.channels.%d", prefix, i), fmt.Sprintf("%d", ch))
			}
		}
		if band.Disabled != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".disabled", fmt.Sprintf("%t", *band.Disabled))
		}
		if band.Power != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".power", fmt.Sprintf("%d", *band.Power))
		}
		if band.PowerMax != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".power_max", fmt.Sprintf("%d", *band.PowerMax))
		}
		if band.PowerMin != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".power_min", fmt.Sprintf("%d", *band.PowerMin))
		}
		if band.Preamble != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".preamble", *band.Preamble)
		}
	}

	// Check band_5_on_24_radio nested fields
	if o.Band5On24Radio != nil {
		band := o.Band5On24Radio
		prefix := "band_5_on_24_radio"

		if band.AllowRrmDisable != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".allow_rrm_disable", fmt.Sprintf("%t", *band.AllowRrmDisable))
		}
		if band.AntGain != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".ant_gain", fmt.Sprintf("%d", *band.AntGain))
		}
		if band.AntennaMode != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".antenna_mode", *band.AntennaMode)
		}
		if band.Bandwidth != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".bandwidth", fmt.Sprintf("%d", *band.Bandwidth))
		}
		if len(band.Channels) > 0 {
			checks.append(t, "TestCheckResourceAttr", prefix+".channels.#", fmt.Sprintf("%d", len(band.Channels)))
			for i, ch := range band.Channels {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.channels.%d", prefix, i), fmt.Sprintf("%d", ch))
			}
		}
		if band.Disabled != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".disabled", fmt.Sprintf("%t", *band.Disabled))
		}
		if band.Power != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".power", fmt.Sprintf("%d", *band.Power))
		}
		if band.PowerMax != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".power_max", fmt.Sprintf("%d", *band.PowerMax))
		}
		if band.PowerMin != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".power_min", fmt.Sprintf("%d", *band.PowerMin))
		}
		if band.Preamble != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".preamble", *band.Preamble)
		}
	}

	// Check band_6 nested fields
	if o.Band6 != nil {
		band := o.Band6
		prefix := "band_6"

		if band.AllowRrmDisable != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".allow_rrm_disable", fmt.Sprintf("%t", *band.AllowRrmDisable))
		}
		if band.AntGain != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".ant_gain", fmt.Sprintf("%d", *band.AntGain))
		}
		if band.AntennaMode != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".antenna_mode", *band.AntennaMode)
		}
		if band.Bandwidth != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".bandwidth", fmt.Sprintf("%d", *band.Bandwidth))
		}
		if len(band.Channels) > 0 {
			checks.append(t, "TestCheckResourceAttr", prefix+".channels.#", fmt.Sprintf("%d", len(band.Channels)))
			for i, ch := range band.Channels {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.channels.%d", prefix, i), fmt.Sprintf("%d", ch))
			}
		}
		if band.Disabled != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".disabled", fmt.Sprintf("%t", *band.Disabled))
		}
		if band.Power != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".power", fmt.Sprintf("%d", *band.Power))
		}
		if band.PowerMax != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".power_max", fmt.Sprintf("%d", *band.PowerMax))
		}
		if band.PowerMin != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".power_min", fmt.Sprintf("%d", *band.PowerMin))
		}
		if band.Preamble != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".preamble", *band.Preamble)
		}
		if band.StandardPower != nil {
			checks.append(t, "TestCheckResourceAttr", prefix+".standard_power", fmt.Sprintf("%t", *band.StandardPower))
		}
	}

	// Check model_specific map fields
	if len(o.ModelSpecific) > 0 {
		for modelKey, modelVal := range o.ModelSpecific {
			modelPrefix := fmt.Sprintf("model_specific.%s", modelKey)

			if modelVal.AntGain24 != nil {
				checks.append(t, "TestCheckResourceAttr", modelPrefix+".ant_gain_24", fmt.Sprintf("%d", *modelVal.AntGain24))
			}
			if modelVal.AntGain5 != nil {
				checks.append(t, "TestCheckResourceAttr", modelPrefix+".ant_gain_5", fmt.Sprintf("%d", *modelVal.AntGain5))
			}
			if modelVal.AntGain6 != nil {
				checks.append(t, "TestCheckResourceAttr", modelPrefix+".ant_gain_6", fmt.Sprintf("%d", *modelVal.AntGain6))
			}
			if modelVal.Band24Usage != nil {
				checks.append(t, "TestCheckResourceAttr", modelPrefix+".band_24_usage", *modelVal.Band24Usage)
			}

			// Check model_specific band_24
			if modelVal.Band24 != nil {
				band := modelVal.Band24
				bandPrefix := modelPrefix + ".band_24"

				if band.AllowRrmDisable != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".allow_rrm_disable", fmt.Sprintf("%t", *band.AllowRrmDisable))
				}
				if band.AntGain != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".ant_gain", fmt.Sprintf("%d", *band.AntGain))
				}
				if band.AntennaMode != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".antenna_mode", *band.AntennaMode)
				}
				if band.Bandwidth != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".bandwidth", fmt.Sprintf("%d", *band.Bandwidth))
				}
				if len(band.Channels) > 0 {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".channels.#", fmt.Sprintf("%d", len(band.Channels)))
					for i, ch := range band.Channels {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.channels.%d", bandPrefix, i), fmt.Sprintf("%d", ch))
					}
				}
				if band.Disabled != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".disabled", fmt.Sprintf("%t", *band.Disabled))
				}
				if band.Power != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".power", fmt.Sprintf("%d", *band.Power))
				}
				if band.PowerMax != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".power_max", fmt.Sprintf("%d", *band.PowerMax))
				}
				if band.PowerMin != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".power_min", fmt.Sprintf("%d", *band.PowerMin))
				}
				if band.Preamble != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".preamble", *band.Preamble)
				}
			}

			// Check model_specific band_5
			if modelVal.Band5 != nil {
				band := modelVal.Band5
				bandPrefix := modelPrefix + ".band_5"

				if band.AllowRrmDisable != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".allow_rrm_disable", fmt.Sprintf("%t", *band.AllowRrmDisable))
				}
				if band.AntGain != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".ant_gain", fmt.Sprintf("%d", *band.AntGain))
				}
				if band.AntennaMode != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".antenna_mode", *band.AntennaMode)
				}
				if band.Bandwidth != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".bandwidth", fmt.Sprintf("%d", *band.Bandwidth))
				}
				if len(band.Channels) > 0 {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".channels.#", fmt.Sprintf("%d", len(band.Channels)))
					for i, ch := range band.Channels {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.channels.%d", bandPrefix, i), fmt.Sprintf("%d", ch))
					}
				}
				if band.Disabled != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".disabled", fmt.Sprintf("%t", *band.Disabled))
				}
				if band.Power != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".power", fmt.Sprintf("%d", *band.Power))
				}
				if band.PowerMax != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".power_max", fmt.Sprintf("%d", *band.PowerMax))
				}
				if band.PowerMin != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".power_min", fmt.Sprintf("%d", *band.PowerMin))
				}
				if band.Preamble != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".preamble", *band.Preamble)
				}
			}

			// Check model_specific band_5_on_24_radio
			if modelVal.Band5On24Radio != nil {
				band := modelVal.Band5On24Radio
				bandPrefix := modelPrefix + ".band_5_on_24_radio"

				if band.AllowRrmDisable != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".allow_rrm_disable", fmt.Sprintf("%t", *band.AllowRrmDisable))
				}
				if band.AntGain != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".ant_gain", fmt.Sprintf("%d", *band.AntGain))
				}
				if band.AntennaMode != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".antenna_mode", *band.AntennaMode)
				}
				if band.Bandwidth != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".bandwidth", fmt.Sprintf("%d", *band.Bandwidth))
				}
				if len(band.Channels) > 0 {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".channels.#", fmt.Sprintf("%d", len(band.Channels)))
					for i, ch := range band.Channels {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.channels.%d", bandPrefix, i), fmt.Sprintf("%d", ch))
					}
				}
				if band.Disabled != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".disabled", fmt.Sprintf("%t", *band.Disabled))
				}
				if band.Power != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".power", fmt.Sprintf("%d", *band.Power))
				}
				if band.PowerMax != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".power_max", fmt.Sprintf("%d", *band.PowerMax))
				}
				if band.PowerMin != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".power_min", fmt.Sprintf("%d", *band.PowerMin))
				}
				if band.Preamble != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".preamble", *band.Preamble)
				}
			}

			// Check model_specific band_6
			if modelVal.Band6 != nil {
				band := modelVal.Band6
				bandPrefix := modelPrefix + ".band_6"

				if band.AllowRrmDisable != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".allow_rrm_disable", fmt.Sprintf("%t", *band.AllowRrmDisable))
				}
				if band.AntGain != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".ant_gain", fmt.Sprintf("%d", *band.AntGain))
				}
				if band.AntennaMode != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".antenna_mode", *band.AntennaMode)
				}
				if band.Bandwidth != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".bandwidth", fmt.Sprintf("%d", *band.Bandwidth))
				}
				if len(band.Channels) > 0 {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".channels.#", fmt.Sprintf("%d", len(band.Channels)))
					for i, ch := range band.Channels {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.channels.%d", bandPrefix, i), fmt.Sprintf("%d", ch))
					}
				}
				if band.Disabled != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".disabled", fmt.Sprintf("%t", *band.Disabled))
				}
				if band.Power != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".power", fmt.Sprintf("%d", *band.Power))
				}
				if band.PowerMax != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".power_max", fmt.Sprintf("%d", *band.PowerMax))
				}
				if band.PowerMin != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".power_min", fmt.Sprintf("%d", *band.PowerMin))
				}
				if band.Preamble != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".preamble", *band.Preamble)
				}
				if band.StandardPower != nil {
					checks.append(t, "TestCheckResourceAttr", bandPrefix+".standard_power", fmt.Sprintf("%t", *band.StandardPower))
				}
			}
		}
	}

	return checks
}
