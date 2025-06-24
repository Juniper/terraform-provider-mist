// WIP
package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgWlanPortalImageModel(t *testing.T) {
	// Create a test PNG file
	testImagePath := createTestPNGFile(t)

	type testStep struct {
		config OrgWlanPortalImageModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgWlanPortalImageModel{
						OrgId: GetTestOrgId(),
						File:  testImagePath,
					},
				},
			},
		},
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "org_wlan_portal_image"
			wlanTemplateName := "test_wlan_template"
			wlanName := "test_wlan"

			// Create single-step tests with combined config (WLAN template + WLAN + portal image)
			// Since portal images require a WLAN, and WLANs require a template,
			// we create all three in the same config but focus our checks on the portal image
			steps := make([]resource.TestStep, len(tCase.steps))

			for i, step := range tCase.steps {
				// Generate combined config: WLAN template + WLAN + portal image
				combinedConfig := generateOrgWlanPortalImageTestConfig(wlanTemplateName, wlanName, tName, step.config)

				// Focus checks on the portal image resource (WLAN template and WLAN are prerequisites)
				checks := newTestChecks(PrefixProviderName(resourceType) + "." + tName)
				checks.append(t, "TestCheckResourceAttr", "org_id", step.config.OrgId)
				checks.append(t, "TestCheckResourceAttrSet", "wlan_id")
				checks.append(t, "TestCheckResourceAttr", "file", step.config.File)

				steps[i] = resource.TestStep{
					Config: combinedConfig,
					Check:  resource.ComposeAggregateTestCheckFunc(checks.checks...),
				}

				// Log configuration for debugging
				t.Logf("\n// ------ Combined Config ------\n%s\n", combinedConfig)
			}

			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps:                    steps,
			})

		})
	}
}

// generateOrgWlanPortalImageTestConfig creates a combined configuration with
// WLAN template + WLAN + portal image since portal images require a WLAN ID
func generateOrgWlanPortalImageTestConfig(wlanTemplateName, wlanName, portalImageName string, portalImageConfig OrgWlanPortalImageModel) string {
	// Create the prerequisite WLAN template
	wlanTemplateConfig := OrgWlantemplateModel{
		Name:  "Test_WLAN_Template",
		OrgId: portalImageConfig.OrgId,
	}

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&wlanTemplateConfig, f.Body())
	wlanTemplateConfigStr := Render("org_wlantemplate", wlanTemplateName, string(f.Bytes()))

	// Create the WLAN that references the template
	templateRef := fmt.Sprintf("mist_org_wlantemplate.%s.id", wlanTemplateName)

	wlanConfig := OrgWlanModel{
		OrgId: portalImageConfig.OrgId,
		Ssid:  "TestSSID",
	}

	f = hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&wlanConfig, f.Body())
	// Add the template_id attribute to the body before rendering
	f.Body().SetAttributeRaw("template_id", hclwrite.TokensForIdentifier(templateRef))
	wlanConfigStr := Render("org_wlan", wlanName, string(f.Bytes()))

	// Create the portal image that references the WLAN
	wlanRef := fmt.Sprintf("mist_org_wlan.%s.id", wlanName)

	f = hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&portalImageConfig, f.Body())
	// Add the wlan_id attribute to the body before rendering
	f.Body().SetAttributeRaw("wlan_id", hclwrite.TokensForIdentifier(wlanRef))
	portalImageConfigStr := Render("org_wlan_portal_image", portalImageName, string(f.Bytes()))

	// Combine all three configs
	return wlanTemplateConfigStr + "\n\n" + wlanConfigStr + "\n\n" + portalImageConfigStr
}

// createTestPNGFile creates a minimal PNG file for testing purposes
func createTestPNGFile(t *testing.T) string {
	// Minimal PNG file data (1x1 transparent pixel)
	pngData := []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, // PNG signature
		0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52, // IHDR chunk header
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, // 1x1 dimensions
		0x08, 0x06, 0x00, 0x00, 0x00, 0x1F, 0x15, 0xC4, // bit depth, color type, etc.
		0x89, 0x00, 0x00, 0x00, 0x0A, 0x49, 0x44, 0x41, // IDAT chunk header
		0x54, 0x78, 0x9C, 0x63, 0x00, 0x01, 0x00, 0x00, // compressed image data
		0x05, 0x00, 0x01, 0x0D, 0x0A, 0x2D, 0xB4, 0x00, // (transparent pixel)
		0x00, 0x00, 0x00, 0x49, 0x45, 0x4E, 0x44, 0xAE, // IEND chunk
		0x42, 0x60, 0x82,
	}

	// Create temporary file
	tmpFile, err := os.CreateTemp("", "test-portal-image-*.png")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	// Write PNG data
	if _, err := tmpFile.Write(pngData); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		t.Fatalf("Failed to write PNG data: %v", err)
	}

	if err := tmpFile.Close(); err != nil {
		os.Remove(tmpFile.Name())
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Clean up on test completion
	t.Cleanup(func() {
		os.Remove(tmpFile.Name())
	})

	return tmpFile.Name()
}
