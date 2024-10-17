package provider

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	// "github.com/Juniper/terraform-provider-mist/internal/provider"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_inventory"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func SearchInventoryByMacFunction_Valid(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"example": providerserver.NewProtocol6WithError(provider.New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::example::echo("test-value")
                }`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test", knownvalue.StringExact("test-value")),
				},
			},
		},
	})
}

// The example implementation does not return any errors, however
// this acceptance test verifies how the function should behave if it did.
func SearchInventoryByMacFunction_Invalid(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"example": providerserver.NewProtocol6WithError(provider.New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::example::echo("invalid")
                }`,
				ExpectError: regexp.MustCompile(`error summary`),
			},
		},
	})
}

// The example implementation does not enable AllowNullValue, however this
// acceptance test shows how to verify the behavior.
func SearchInventoryByMacFunction_Null(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"example": providerserver.NewProtocol6WithError(provider.New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::example::echo(null)
                }`,
				ExpectError: regexp.MustCompile(`Invalid Function Call`),
			},
		},
	})
}

// The example implementation does not enable AllowUnknownValues, however this
// acceptance test shows how to verify the behavior.
func SearchInventoryByMacFunction_Unknown(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"example": providerserver.NewProtocol6WithError(provider.New()),
		},
		Steps: []resource.TestStep{
			{
				Config: `
                terraform_data "test" {
                    input = "test-value"
                }
                
                output "test" {
                    value = provider::example::echo(terraform_data.test.output)
                }`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownOutputValue("test"),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test", knownvalue.StringExact("test-value")),
				},
			},
		},
	})
}
func TestSearchInventoryFunctionRun(t *testing.T) {
	var ctx context.Context
	var inventory_av []attr.Value
	device := resource_org_inventory.NewDevicesValueNull()
	device.DeviceprofileId = types.StringNull()
	device.Magic = types.StringValue("0123ABCD456")
	device.Hostname = types.StringValue("hostname")
	device.Id = types.StringValue("00000000-0000-0000-1000-c0ffee000000")
	device.Mac = types.StringValue("c0ffee000000")
	device.Model = types.StringValue("AP47")
	device.OrgId = types.StringValue("992bf4b9-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
	device.Serial = types.StringValue("1440XXXXXXXX")
	device.SiteId = types.StringValue("d7c8364e-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
	device.DevicesType = types.StringValue("ap")
	device.UnclaimWhenDestroyed = types.BoolValue(false)
	device.VcMac = types.StringValue("")

	inventory_av = append(inventory_av, device)
	inventory := types.SetValueMust(resource_org_inventory.DevicesValue{}.Type(ctx), inventory_av)

	var inventory_null_av []attr.Value
	inventory_null := types.SetValueMust(resource_org_inventory.DevicesValue{}.Type(ctx), inventory_null_av)
	t.Parallel()

	var mac_valid types.String = types.StringValue("c0ffee000000")
	var mac_invalid types.String = types.StringValue("c0ffee000001")
	var mac_null types.String = types.StringNull()

	testCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		// The example implementation uses the Go built-in string type, however
		// if AllowNullValue was enabled and *string or types.String was used,
		// this test case shows how the function would be expected to behave.
		"no-inventry": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData(
					[]attr.Value{
						inventory_null,
						mac_valid,
					},
				),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "The inventory provided is emtpy"),
				Result: function.NewResultData(resource_org_inventory.NewDevicesValueNull()),
			},
		},
		"no-mac": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData(
					[]attr.Value{
						inventory,
						mac_null,
					},
				),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(resource_org_inventory.NewDevicesValueNull()),
			},
		},
		// The example implementation uses the Go built-in string type, however
		// if AllowUnknownValues was enabled and types.String was used,
		// this test case shows how the function would be expected to behave.
		"unknown": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData(
					[]attr.Value{
						types.SetUnknown(resource_org_inventory.DevicesValue{}.Type(ctx)),
						mac_valid,
					},
				),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(resource_org_inventory.NewDevicesValueNull()),
			},
		},
		"value-valid": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData(
					[]attr.Value{
						inventory,
						mac_valid,
					},
				),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(device),
			},
		},
		// The example implementation does not return an error, however
		// this test case shows how the function would be expected to behave if
		// it did.
		"value-invalid": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData(
					[]attr.Value{
						inventory,
						mac_invalid,
					},
				),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, fmt.Sprintf("Unable to find a device with MAC Address \"%s\" in the provided inventory", mac_invalid)),
				Result: function.NewResultData(resource_org_inventory.NewDevicesValueNull()),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := function.RunResponse{
				Result: function.NewResultData(types.StringUnknown()),
			}

			provider.SearchInventoryByMacFunction{}.Run(context.Background(), testCase.request, &got)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
