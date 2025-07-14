package provider

import (
	"fmt"
	"math"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
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
}

func (o *testChecks) append(t testing.TB, testCheckFuncName string, testCheckFuncArgs ...string) {
	t.Helper()

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

func (o *testChecks) appendSetNestedCheck(_ testing.TB, attrName string, m map[string]string) {
	o.checks = append(o.checks, resource.TestCheckTypeSetElemNestedAttrs(o.path, attrName, m))
	o.logLines.appendf("TestCheckTypeSetElemNestedAttrs(%s, %s, %s)", o.path, attrName, m)
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

func GetSiteWlanBaseConfig(org_ID string) (config string, siteRef string, wlanRef string) {
	siteConfig := SiteModel{
		Name:    "TestSite",
		OrgId:   org_ID,
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

func GetSiteBaseConfig(org_ID string) (config string, wlanRef string) {
	siteConfig := SiteModel{
		Name:    "TestSite",
		OrgId:   org_ID,
		Address: "TestAddress",
	}

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&siteConfig, f.Body())
	siteConfigStr := Render("site", siteConfig.Name, string(f.Bytes()))

	return siteConfigStr, fmt.Sprintf("mist_site.%s.id", siteConfig.Name)
	
// Helper function for creating string pointers
func stringPtr(s string) *string {
	return &s
}
