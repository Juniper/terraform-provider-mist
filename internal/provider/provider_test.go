package provider

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/tmunzer/mistapi-go/mistapi"
)

func TestDebugEnvVariable_True(t *testing.T) {
	t.Setenv("MIST_API_DEBUG", "true")

	var p mistProviderModel
	var diags diag.Diagnostics
	p.fromEnv(context.Background(), &diags)
	if !p.ApiDebug.ValueBool() {
		t.Errorf("ApiDebug not set correctly from environment variable")
	}
}
func TestDebugEnvVariable_False(t *testing.T) {
	t.Setenv("MIST_API_DEBUG", "false")

	var p mistProviderModel
	var diags diag.Diagnostics
	p.fromEnv(context.Background(), &diags)
	if p.ApiDebug.ValueBool() {
		t.Errorf("ApiDebug not set correctly from environment variable")
	}
}
func TestDebugEnvVariable_Null(t *testing.T) {
	var p mistProviderModel
	var diags diag.Diagnostics
	p.fromEnv(context.Background(), &diags)
	if p.ApiDebug.ValueBool() {
		t.Errorf("ApiDebug not set correctly from environment variable")
	}
}

func TestFromEnv_ApiDebug(t *testing.T) {
	tests := []struct {
		name          string
		envValue      string
		expectedValue bool
		expectError   bool
		envIsUnset    bool
	}{
		{
			name:          "ApiDebug true from env variable",
			envValue:      "true",
			expectedValue: true,
		},
		{
			name:          "ApiDebug false from env variable",
			envValue:      "false",
			expectedValue: false,
		},
		{
			name:        "Invalid ApiDebug value",
			envValue:    "invalid",
			expectError: true,
		},
		{
			name:       "ApiDebug unset",
			envIsUnset: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.envIsUnset {
				t.Setenv("MIST_API_DEBUG", tt.envValue)
			} else {
				os.Unsetenv("MIST_API_DEBUG")
			}

			var p = mistProviderModel{
				ApiDebug: types.BoolNull(),
			}

			var diags diag.Diagnostics

			p.fromEnv(context.Background(), &diags)

			if tt.expectError {
				if !diags.HasError() {
					t.Errorf("expected an error but got none")
				}
				return
			} else if diags.HasError() {
				t.Fatalf("unexpected diagnostics: %v", diags)
			}

			if !p.ApiDebug.IsNull() {
				if p.ApiDebug.ValueBool() != tt.expectedValue {
					t.Errorf("expected ApiDebug to be %v, got %v", tt.expectedValue, p.ApiDebug.ValueBool())
				}
			} else if !tt.envIsUnset {
				t.Errorf("expected ApiDebug to be set, but it was null")
			}
		})
	}
}

func TestIntegration_Configure_Invalid_Configuration(t *testing.T) {

	testProvider := New()()
	schemaResponse := new(provider.SchemaResponse)

	testProvider.Schema(context.Background(), provider.SchemaRequest{}, schemaResponse)

	testConfig := tfsdk.Config{
		Raw: tftypes.NewValue(tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"host":        tftypes.String,
				"apitoken":    tftypes.String,
				"username":    tftypes.String,
				"password":    tftypes.String,
				"api_timeout": tftypes.Number,
				"api_debug":   tftypes.Bool,
				"proxy":       tftypes.String,
			},
		}, map[string]tftypes.Value{
			"host":        tftypes.NewValue(tftypes.String, "api.mist.com"),
			"apitoken":    tftypes.NewValue(tftypes.String, "fake-token"),
			"username":    tftypes.NewValue(tftypes.String, nil),
			"password":    tftypes.NewValue(tftypes.String, nil),
			"api_timeout": tftypes.NewValue(tftypes.Number, nil),
			"api_debug":   tftypes.NewValue(tftypes.Bool, nil),
			"proxy":       tftypes.NewValue(tftypes.String, nil),
		}),
		Schema: schemaResponse.Schema,
	}

	req := provider.ConfigureRequest{
		Config: testConfig,
	}
	resp := provider.ConfigureResponse{}

	testProvider.Configure(context.Background(), req, &resp)

	if resp.Diagnostics.HasError() {
		assert.Equal(t, diag.Diagnostics{diag.NewErrorDiagnostic("Authentication Error", "ResponseHttp401Error occurred: Unauthorized")}, resp.Diagnostics)
	}
}
func TestIntegration_Configure_API_Token(t *testing.T) {
	ValidateEnvVars(t, []string{"TEST_MIST_API_TOKEN"})

	testProvider := New()()
	schemaResponse := new(provider.SchemaResponse)

	testProvider.Schema(context.Background(), provider.SchemaRequest{}, schemaResponse)

	testConfig := tfsdk.Config{
		Raw: tftypes.NewValue(tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"host":        tftypes.String,
				"apitoken":    tftypes.String,
				"username":    tftypes.String,
				"password":    tftypes.String,
				"api_timeout": tftypes.Number,
				"api_debug":   tftypes.Bool,
				"proxy":       tftypes.String,
			},
		}, map[string]tftypes.Value{
			"host":        tftypes.NewValue(tftypes.String, "api.mist.com"),
			"apitoken":    tftypes.NewValue(tftypes.String, os.Getenv("TEST_MIST_API_TOKEN")),
			"username":    tftypes.NewValue(tftypes.String, nil),
			"password":    tftypes.NewValue(tftypes.String, nil),
			"api_timeout": tftypes.NewValue(tftypes.Number, nil),
			"api_debug":   tftypes.NewValue(tftypes.Bool, nil),
			"proxy":       tftypes.NewValue(tftypes.String, nil),
		}),
		Schema: schemaResponse.Schema,
	}

	req := provider.ConfigureRequest{
		Config: testConfig,
	}
	resp := provider.ConfigureResponse{}

	testProvider.Configure(context.Background(), req, &resp)
	assert.NotEmpty(t, resp.ResourceData.(mistapi.ClientInterface).Configuration().ApiTokenCredentials())
	assert.Empty(t, resp.ResourceData.(mistapi.ClientInterface).Configuration().BasicAuthCredentials())
	assert.Empty(t, resp.ResourceData.(mistapi.ClientInterface).Configuration().CsrfTokenCredentials())
}

func TestIntegration_Configure_Basic_Auth(t *testing.T) {
	ValidateEnvVars(t, []string{"TEST_MIST_USERNAME", "TEST_MIST_PASSWORD"})

	testProvider := New()()
	schemaResponse := new(provider.SchemaResponse)

	testProvider.Schema(context.Background(), provider.SchemaRequest{}, schemaResponse)

	testConfig := tfsdk.Config{
		Raw: tftypes.NewValue(tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"host":        tftypes.String,
				"apitoken":    tftypes.String,
				"username":    tftypes.String,
				"password":    tftypes.String,
				"api_timeout": tftypes.Number,
				"api_debug":   tftypes.Bool,
				"proxy":       tftypes.String,
			},
		}, map[string]tftypes.Value{
			"host":        tftypes.NewValue(tftypes.String, "api.mist.com"),
			"apitoken":    tftypes.NewValue(tftypes.String, nil),
			"username":    tftypes.NewValue(tftypes.String, os.Getenv("TEST_MIST_USERNAME")),
			"password":    tftypes.NewValue(tftypes.String, os.Getenv("TEST_MIST_PASSWORD")),
			"api_timeout": tftypes.NewValue(tftypes.Number, nil),
			"api_debug":   tftypes.NewValue(tftypes.Bool, nil),
			"proxy":       tftypes.NewValue(tftypes.String, nil),
		}),
		Schema: schemaResponse.Schema,
	}

	req := provider.ConfigureRequest{
		Config: testConfig,
	}
	resp := provider.ConfigureResponse{}

	testProvider.Configure(context.Background(), req, &resp)
	assert.Empty(t, resp.ResourceData.(mistapi.ClientInterface).Configuration().ApiTokenCredentials())
	assert.NotEmpty(t, resp.ResourceData.(mistapi.ClientInterface).Configuration().BasicAuthCredentials())
	assert.NotEmpty(t, resp.ResourceData.(mistapi.ClientInterface).Configuration().CsrfTokenCredentials())
}

func TestIntegration_Configure_API_DEBUG(t *testing.T) {
	ValidateEnvVars(t, []string{"TEST_MIST_API_TOKEN"})

	testProvider := New()()
	schemaResponse := new(provider.SchemaResponse)

	testProvider.Schema(context.Background(), provider.SchemaRequest{}, schemaResponse)

	testConfig := tfsdk.Config{
		Raw: tftypes.NewValue(tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"host":        tftypes.String,
				"apitoken":    tftypes.String,
				"username":    tftypes.String,
				"password":    tftypes.String,
				"api_timeout": tftypes.Number,
				"api_debug":   tftypes.Bool,
				"proxy":       tftypes.String,
			},
		}, map[string]tftypes.Value{
			"host":        tftypes.NewValue(tftypes.String, "api.mist.com"),
			"apitoken":    tftypes.NewValue(tftypes.String, os.Getenv("TEST_MIST_API_TOKEN")),
			"username":    tftypes.NewValue(tftypes.String, nil),
			"password":    tftypes.NewValue(tftypes.String, nil),
			"api_timeout": tftypes.NewValue(tftypes.Number, nil),
			"api_debug":   tftypes.NewValue(tftypes.Bool, true),
			"proxy":       tftypes.NewValue(tftypes.String, nil),
		}),
		Schema: schemaResponse.Schema,
	}

	req := provider.ConfigureRequest{
		Config: testConfig,
	}
	resp := provider.ConfigureResponse{}

	testProvider.Configure(context.Background(), req, &resp)
	assert.False(t, resp.Diagnostics.HasError(), "Unexpected error in diagnostics: %v", resp.Diagnostics)
}

func ValidateEnvVars(t *testing.T, requiredVars []string) {
	t.Helper()

	var missingVars []string
	for _, envVar := range requiredVars {
		if _, exists := os.LookupEnv(envVar); !exists {
			missingVars = append(missingVars, envVar)
		}
	}

	if len(missingVars) > 0 {
		t.Errorf("Test missing required environment variables: %v", missingVars)
		t.FailNow()
	}
}
