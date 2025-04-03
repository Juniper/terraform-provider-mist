package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func newTestProvider() provider.Provider {
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
			"host":        tftypes.NewValue(tftypes.String, "api.mistsys.com"),
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
	mdr := provider.MetadataResponse{}
	testProvider.Metadata(context.Background(), provider.MetadataRequest{}, &mdr)

	return testProvider
}

var (
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"mist": providerserver.NewProtocol6WithError(newTestProvider()),
	}
)
