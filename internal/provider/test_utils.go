package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/tools/validators"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func GetTestOrgId() string {
	return os.Getenv("MIST_TEST_ORG_ID")
}

func PrefixProviderName(name string) string {
	if strings.HasPrefix(name, "mist_") {
		return name
	}
	return "mist_" + name
}

func Render(resourceType, resourceName, config string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
%s
}`, PrefixProviderName(resourceType), resourceName, config)
}

func newTestChecks(path string) testChecks {
	return testChecks{path: path}
}

type testChecks struct {
	path     string
	logLines lineNumberer
	checks   []resource.TestCheckFunc
	tracker  *validators.FieldCoverageTracker // Optional field coverage tracker
}

// SetTracker enables field coverage tracking for this test check
func (o *testChecks) SetTracker(tracker *validators.FieldCoverageTracker) {
	o.tracker = tracker
}

func (o *testChecks) append(t testing.TB, testCheckFuncName string, testCheckFuncArgs ...string) {
	t.Helper()

	// Track field coverage if enabled via environment variable and tracker is set
	if os.Getenv("MIST_TRACK_FIELD_COVERAGE") != "" && o.tracker != nil && len(testCheckFuncArgs) > 0 {
		o.tracker.MarkFieldAsTested(testCheckFuncArgs[0])
	}

	switch testCheckFuncName {
	case "TestCheckResourceAttrSet":
		if len(testCheckFuncArgs) != 1 {
			t.Fatalf("%s requires 1 args, got %d", testCheckFuncName, len(testCheckFuncArgs))
		}
		o.checks = append(o.checks, resource.TestCheckResourceAttrSet(o.path, testCheckFuncArgs[0]))
		o.logLines.appendf("TestCheckResourceAttrSet(%s, %q)", o.path, testCheckFuncArgs[0])
	case "TestCheckNoResourceAttr":
		if len(testCheckFuncArgs) != 1 {
			t.Fatalf("%s requires 1 args, got %d", testCheckFuncName, len(testCheckFuncArgs))
		}
		o.checks = append(o.checks, resource.TestCheckNoResourceAttr(o.path, testCheckFuncArgs[0]))
		o.logLines.appendf("TestCheckNoResourceAttr(%s, %q)", o.path, testCheckFuncArgs[0])
	case "TestCheckResourceAttr":
		if len(testCheckFuncArgs) != 2 {
			t.Fatalf("%s requires 2 args, got %d", testCheckFuncName, len(testCheckFuncArgs))
		}
		o.checks = append(o.checks, resource.TestCheckResourceAttr(o.path, testCheckFuncArgs[0], testCheckFuncArgs[1]))
		o.logLines.appendf("TestCheckResourceAttr(%s, %q, %q)", o.path, testCheckFuncArgs[0], testCheckFuncArgs[1])
	case "TestCheckTypeSetElemAttr":
		if len(testCheckFuncArgs) != 2 {
			t.Fatalf("%s requires 2 args, got %d", testCheckFuncName, len(testCheckFuncArgs))
		}
		o.checks = append(o.checks, resource.TestCheckTypeSetElemAttr(o.path, testCheckFuncArgs[0], testCheckFuncArgs[1]))
	case "TestCheckResourceAttrPair":
		if len(testCheckFuncArgs) != 2 {
			t.Fatalf("%s requires 2 args, got %d", testCheckFuncName, len(testCheckFuncArgs))
		}
		o.checks = append(o.checks, resource.TestCheckResourceAttrPair(o.path, testCheckFuncArgs[0], o.path, testCheckFuncArgs[1]))
	default:
		t.Fatalf("unknown test check function: %s", testCheckFuncName)
	}
}

func (o *testChecks) string() string {
	return o.logLines.string()
}

type lineNumberer struct {
	lines []string
	base  int
}

func (o *lineNumberer) setBase(base int) error {
	switch base {
	case 2:
	case 8:
	case 10:
	case 16:
	default:
		return fmt.Errorf("base %d not supported", base)
	}

	o.base = base
	return nil
}

func (o *lineNumberer) append(l string) {
	o.lines = append(o.lines, l)
}

func (o *lineNumberer) appendf(format string, a ...any) {
	o.append(fmt.Sprintf(format, a...))
}

func (o *lineNumberer) string() string {
	count := len(o.lines)
	if count == 0 {
		return ""
	}

	base := o.base
	if base == 0 {
		base = 10
	}

	formatStr, _ := padFormatStr(count, base) // err ignored because only valid base can exist here

	sb := new(strings.Builder)
	for i, line := range o.lines {
		sb.WriteString(fmt.Sprintf(formatStr, i+1) + " " + line + "\n")
	}

	return sb.String()
}

func padFormatStr(n, base int) (string, error) {
	var baseChar string
	switch base {
	case 2:
		baseChar = "b"
	case 8:
		baseChar = "o"
	case 10:
		baseChar = "d"
	case 16:
		baseChar = "x"
	default:
		return "", fmt.Errorf("base %d not supported", base)
	}

	c := int(math.Floor(math.Log(float64(n))/math.Log(float64(base)))) + 1
	return fmt.Sprintf("%%%d%s", c, baseChar), nil
}

func TestLineNumbererString(t *testing.T) {
	type testCase struct {
		lines    []string
		base     int
		expected string
	}

	testCases := []testCase{
		{
			lines:    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"},
			base:     2,
			expected: "    1 a\n   10 b\n   11 c\n  100 d\n  101 e\n  110 f\n  111 g\n 1000 h\n 1001 i\n 1010 j\n 1011 k\n 1100 l\n 1101 m\n 1110 n\n 1111 o\n10000 p\n",
		},
		{
			lines:    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"},
			base:     8,
			expected: " 1 a\n 2 b\n 3 c\n 4 d\n 5 e\n 6 f\n 7 g\n10 h\n11 i\n12 j\n13 k\n14 l\n15 m\n16 n\n17 o\n20 p\n",
		},
		{
			lines:    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"},
			base:     10,
			expected: " 1 a\n 2 b\n 3 c\n 4 d\n 5 e\n 6 f\n 7 g\n 8 h\n 9 i\n10 j\n11 k\n12 l\n13 m\n14 n\n15 o\n16 p\n",
		},
		{
			lines:    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"},
			base:     16,
			expected: " 1 a\n 2 b\n 3 c\n 4 d\n 5 e\n 6 f\n 7 g\n 8 h\n 9 i\n a j\n b k\n c l\n d m\n e n\n f o\n10 p\n",
		},
	}

	for _, tCase := range testCases {
		ln := new(lineNumberer)
		require.NoError(t, ln.setBase(tCase.base))
		for _, line := range tCase.lines {
			ln.append(line)
		}
		result := ln.string()

		assert.Equal(t, tCase.expected, result)
	}
}

// CreateTestPNGFile creates a minimal PNG file for testing purposes
func CreateTestPNGFile(t *testing.T) string {
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
	_, err = tmpFile.Write(pngData)
	if err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		t.Fatalf("Failed to write PNG data: %v", err)
	}

	err = tmpFile.Close()
	if err != nil {
		os.Remove(tmpFile.Name())
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Clean up on test completion
	t.Cleanup(func() {
		os.Remove(tmpFile.Name())
	})

	return tmpFile.Name()
}

func GetOrgWlanBaseConfig(orgID string) (config string, wlanRef string) {
	// Create the prerequisite WLAN template
	wlanTemplateConfig := OrgWlantemplateModel{
		Name:  "Test_WLAN_Template",
		OrgId: orgID,
	}

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&wlanTemplateConfig, f.Body())
	wlanTemplateConfigStr := Render("org_wlantemplate", "test_wlan_template", string(f.Bytes()))

	// Create the WLAN that references the template
	templateRef := fmt.Sprintf("mist_org_wlantemplate.%s.id", "test_wlan_template")

	wlanConfig := OrgWlanModel{
		OrgId: orgID,
		Ssid:  "TestSSID",
	}

	f = hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&wlanConfig, f.Body())
	// Add the template_id attribute to the body before rendering
	f.Body().SetAttributeRaw("template_id", hclwrite.TokensForIdentifier(templateRef))
	wlanConfigStr := Render("org_wlan", "wlanName", string(f.Bytes()))

	return wlanTemplateConfigStr + "\n\n" + wlanConfigStr, "mist_org_wlan.wlanName.id"
}

func GetSiteWlanBaseConfig(orgID string) (config string, siteRef string, wlanRef string) {
	siteConfig := SiteModel{
		Name:    "TestSite",
		OrgId:   orgID,
		Address: "TestAddress",
	}

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&siteConfig, f.Body())
	siteConfigStr := Render("site", siteConfig.Name, string(f.Bytes()))

	siteRef = fmt.Sprintf("mist_site.%s.id", siteConfig.Name)

	wlanConfig := SiteWlanModel{
		Ssid: "TestSSID",
	}

	f = hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&wlanConfig, f.Body())
	f.Body().SetAttributeRaw("site_id", hclwrite.TokensForIdentifier(siteRef))
	wlanConfigStr := Render("site_wlan", "wlanName", string(f.Bytes()))

	return siteConfigStr + "\n\n" + wlanConfigStr, fmt.Sprintf("mist_site.%s.id", siteConfig.Name), "mist_site_wlan.wlanName.id"
}

func GetSiteBaseConfig(orgID string) (config string, siteRef string) {
	siteConfig := SiteModel{
		Name:    "TestSite",
		OrgId:   orgID,
		Address: "TestAddress",
	}

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&siteConfig, f.Body())
	siteConfigStr := Render("site", siteConfig.Name, string(f.Bytes()))

	return siteConfigStr, fmt.Sprintf("mist_site.%s.id", siteConfig.Name)
}

func GetSitegroupBaseConfig(orgID string) (config string, sitegroupRef string) {
	sitegroupConfig := OrgSitegroupModel{
		Name:  "TestSitegroup",
		OrgId: orgID,
	}

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&sitegroupConfig, f.Body())
	sitegroupConfigStr := Render("org_sitegroup", sitegroupConfig.Name, string(f.Bytes()))

	return sitegroupConfigStr, fmt.Sprintf("mist_org_sitegroup.%s.id", sitegroupConfig.Name)
}

// Helper function for creating string pointers
func stringPtr(s string) *string {
	return &s
}

// TrackFieldCoverage enables field coverage tracking for test checks
// when the MIST_TRACK_FIELD_COVERAGE environment variable is set.
// It extracts schema fields and configures tracking for field coverage reporting.
func TrackFieldCoverage(t testing.TB, checks *testChecks, resourceName string, schemaFunc func(context.Context) schema.Schema) {
	if os.Getenv("MIST_TRACK_FIELD_COVERAGE") == "" {
		return
	}

	ctx := context.Background()
	schemaObj := schemaFunc(ctx)
	tracker := validators.ExtractAllSchemaFields(resourceName, schemaObj.Attributes)
	checks.SetTracker(tracker)
}

// FieldCoverageReport writes the current state of the FieldCoverageTracker to a JSON file.
func FieldCoverageReport(t testing.TB, checks *testChecks) {
	t.Helper()

	if checks.tracker == nil {
		return
	}

	type CoverageReport struct {
		ResourceName                string   `json:"resource_name"`
		TestedFieldsCnt             int      `json:"tested_fields_count"`
		UntestedFieldsCnt           int      `json:"untested_fields_count"`
		UntestedFields              []string `json:"untested_fields"`
		UnknownFieldsCnt            int      `json:"unknown_fields_count"`
		UnknownFields               []string `json:"unknown_fields"`
		SchemaExtractionFailuresCnt int      `json:"schema_extraction_failures_count"`
		SchemaExtractionFailures    []string `json:"schema_extraction_failures"`
	}

	// Build report
	untestedFields := make([]string, 0)
	for path, field := range checks.tracker.SchemaFields {
		if !field.Computed && !field.IsTested && !isContainerType(field.SchemaAttr) {
			untestedFields = append(untestedFields, path)
		}
	}

	// Convert unknown fields map to sorted slice
	unknownFields := make([]string, 0, len(checks.tracker.UnknownFields))
	for path := range checks.tracker.UnknownFields {
		unknownFields = append(unknownFields, path)
	}

	sort.Strings(unknownFields)
	sort.Strings(untestedFields)
	report := CoverageReport{
		ResourceName:                checks.tracker.ResourceName,
		TestedFieldsCnt:             len(checks.tracker.NormalizedFields),
		UntestedFieldsCnt:           len(untestedFields),
		UntestedFields:              untestedFields,
		UnknownFieldsCnt:            len(unknownFields),
		UnknownFields:               unknownFields,
		SchemaExtractionFailuresCnt: len(checks.tracker.SchemaExtractionFailures),
		SchemaExtractionFailures:    checks.tracker.SchemaExtractionFailures,
	}

	// Write JSON files to tools/reports directory
	err := writeJSONReport(fmt.Sprintf("../../tools/reports/%s_report.json", checks.tracker.ResourceName), report)
	if err != nil {
		t.Errorf("failed to write field coverage report: %v", err)
	}
}

// isContainerType checks if an attribute is a container type
// Container types cannot be tested by themselves and are thus excluded from test coverage counts
func isContainerType(attr schema.Attribute) bool {
	_, isSingleNested := attr.(schema.SingleNestedAttribute)
	_, isMapNested := attr.(schema.MapNestedAttribute)
	return isSingleNested || isMapNested
}

// writeJSONReport writes data as indented JSON to the specified file
func writeJSONReport(filename string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
