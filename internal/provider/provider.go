package provider

import (
	"context"
	"strings"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"os"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	envHost     = "MIST_HOST"
	envApitoken = "MIST_API_TOKEN"
	envUsername = "MIST_USERNAME"
	envPassword = "MIST_PASSWORD"
)

var _ provider.Provider = (*mistProvider)(nil)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &mistProvider{}
	}
}

type mistProvider struct {
	version string
}
type mistProviderData struct {
	client mistapi.ClientInterface
}

type mistProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Apitoken types.String `tfsdk:"apitoken"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func (p *mistProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {

	resp.Schema = schema.Schema{
		MarkdownDescription: "The Mist Provider allows Terraform to manage Juniper Mist.",
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				MarkdownDescription: "URL of the Mist Cloud, e.g. `api.mist.com`\n." +
					"The preferred approach is to pass the credentials as environment variables `",
				Optional: true,
			},
			"apitoken": schema.StringAttribute{
				MarkdownDescription: "For Api Token authentication, the Mist API Token",
				Optional:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "For username/password authentication, the Mist Account username",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "For username/password authentication, the Mist Account password",
				Optional:            true,
			},
		},
	}
}

func (p *mistProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config mistProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Mist API Host",
			"The provider cannot create the Mist API client as there is an unknown configuration value for the Mist API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the HOST environment variable.",
		)
	}
	if config.Apitoken.IsUnknown() && (config.Username.IsUnknown() && config.Password.IsUnknown()) {
		resp.Diagnostics.AddError(
			"Unknown Mist API Authentication",
			"The provider cannot create the Mist API client as there is an unknown authentication configuration. "+
				"Either the API Token or the Username/Password must be statically set in the configuration or as environment variables.",
		)
	} else if config.Apitoken.IsUnknown() && (!config.Username.IsUnknown() && config.Password.IsUnknown()) {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown Mist API Password",
			"The provider cannot create the Mist API client as there is an unknown configuration value for the Mist Username whereas the MIST Username is configured. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the MIST_USERNAME environment variable.",
		)
	} else if config.Apitoken.IsUnknown() && (config.Username.IsUnknown() && !config.Password.IsUnknown()) {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown Mist API Password",
			"The provider cannot create the Mist API client as there is an unknown configuration value for the Mist Password whereas the MIST Username is configured. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the MIST_PASSWORD environment variable.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	host := os.Getenv("MIST_HOST")
	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	apitoken := os.Getenv("MIST_APITOKEN")
	if !config.Apitoken.IsNull() {
		apitoken = config.Apitoken.ValueString()
	}

	username := os.Getenv("MIST_USERNAME")
	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	password := os.Getenv("MIST_PASSWORD")
	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing MIST API Host",
			"The provider cannot create the MIST API client as there is a missing or empty value for the MIST API host. "+
				"Set the host value in the configuration or use the MIST_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
	if apitoken == "" && (username == "" && password == "") {
		resp.Diagnostics.AddError(
			"Missing MIST API Authentication",
			"The provider cannot create the MIST API client as there the authentication configuration is missing. "+
				"Set the Authentication values in the configuration or in the environment variables: "+
				" * apitoken (environment variable \"APITOKEN\")"+
				" * username and password (environment variables \"USERNAME\" and \"PASSWORD\")"+
				"If either is already set, ensure the value is not empty.",
		)
	} else if apitoken == "" && (username != "" && password == "") {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing MIST API Password",
			"The provider cannot create the MIST API client as there is a  a missing or empty value for the MIST Username whereas the MIST Password is configured. "+
				"Set the host value in the configuration or use the MIST_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	} else if apitoken == "" && (username == "" && password != "") {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing MIST API Password",
			"The provider cannot create the MIST API client as there is a  a missing or empty value for the MIST Password whereas the MIST Username is configured. "+
				"Set the host value in the configuration or use the MIST_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	var mist_cloud mistapi.Environment
	switch host {
	case "api.mist.com":
		mist_cloud = mistapi.MIST_GLOBAL_01
	case "api.gc1.mist.com":
		mist_cloud = mistapi.MIST_GLOBAL_02
	case "api.ac2.mist.com":
		mist_cloud = mistapi.MIST_GLOBAL_03
	case "api.gc2.mist.com":
		mist_cloud = mistapi.MIST_GLOBAL_04
	case "api.eu.mist.com":
		mist_cloud = mistapi.MIST_EMEA_01
	case "api.gc3.mist.com":
		mist_cloud = mistapi.MIST_EMEA_02
	case "api.ac6.mist.com":
		mist_cloud = mistapi.MIST_EMEA_03
	case "api.ac5.mist.com":
		mist_cloud = mistapi.MIST_APAC_01
	case "api.mistsys.com":
		mist_cloud = mistapi.AWS_STAGING
	default:
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Wrong Mist Host",
			"The configured host \""+host+"\" is not a valid Mist Host. Please refer to the documentation to get the possible values",
		)
		return
	}

	var client mistapi.ClientInterface
	var client_config mistapi.Configuration

	// configure the client for API Token Auth
	if apitoken != "" {
		client_config = mistapi.CreateConfiguration(
			mistapi.WithEnvironment(mist_cloud),
			mistapi.WithApiTokenCredentials(
				mistapi.NewApiTokenCredentials("Token "+apitoken),
			),
		)
		// configure the client for Basic Auth + CSRF
	} else {
		// Initiate the login API Call
		tmp_client := mistapi.NewClient(
			mistapi.CreateConfiguration(
				mistapi.WithEnvironment(mist_cloud),
				mistapi.WithBasicAuthCredentials(
					mistapi.NewBasicAuthCredentials(username, password),
				),
			),
		)
		body := models.Login{}
		body.Email = username
		body.Password = password
		apiResponse, err := tmp_client.AdminsLogin().Login(ctx, &body)
		if err != nil {
			resp.Diagnostics.AddError("Authentication Error", err.Error())
			if resp.Diagnostics.HasError() {
				return
			}
			// Process the Response Headers to extract the CSRF Token
		} else {
			csrfTokenSet := false
			for hNAme, hVal := range apiResponse.Response.Header {
				if hNAme == "Set-Cookie" {
					for _, cooky := range hVal {
						for _, cVal := range strings.Split(cooky, ";") {
							if strings.HasPrefix(cVal, "csrftoken") {
								csrfToken_string := strings.Split(cVal, "=")[1]
								test := mistapi.NewCsrfTokenCredentials(string(csrfToken_string))
								client_config = mistapi.CreateConfiguration(
									mistapi.WithEnvironment(mist_cloud),
									mistapi.WithBasicAuthCredentials(
										mistapi.NewBasicAuthCredentials(username, password),
									),
									mistapi.WithCsrfTokenCredentials(test),
								)
								csrfTokenSet = true
							}
						}
					}
				}
			}
			// IF CSRF Token not set, raise an error and exit
			if !csrfTokenSet {
				resp.Diagnostics.AddError("Authentication Error", "Unable to extract the CSRF Token from the Authentication response")
				if resp.Diagnostics.HasError() {
					return
				}
			}
		}
	}

	// Use the  configuration to create the client and test the credentials
	client = mistapi.NewClient(client_config)
	_, err := client.SelfAccount().GetSelf(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Authentication Error", err.Error())
		if resp.Diagnostics.HasError() {
			return
		}
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *mistProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "mist"
	resp.Version = p.version
}

func (p *mistProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewCountriesDataSource,
		NewDeviceApStatsDataSource,
		NewDeviceSwitchStatsDataSource,
		NewDeviceGatewayStatsDataSource,
		NewOrgGatewaytemplatesDataSource,
		NewOrgInventoryDataSource,
		NewOrgNacrulesDataSource,
		NewOrgNactagsDataSource,
		NewOrgNetworksDataSource,
		NewOrgNetworktemplatesDataSource,
		NewOrgRftemplatesDataSource,
		NewOrgServicesDataSource,
		NewOrgSitegroupsDataSource,
		NewOrgVpnsDataSource,
		NewOrgWlantemplatesDataSource,
		NewOrgWxtagsDataSource,
		NewSitesDataSource,
		NewOrgDeviceprofilesApDataSource,
		NewOrgDeviceprofilesGatewayDataSource,
	}
}

func (p *mistProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewOrgResource,
		NewOrgSiteGroupResource,
		NewOrgNetworkTemplate,
		NewOrgServiceResource,
		NewOrgNetworkResource,
		NewOrgGatewayTemplate,
		NewOrgInventory,
		NewOrgNacTag,
		NewOrgNacRule,
		NewOrgRfTemplate,
		NewOrgVpn,
		NewOrgWlanTemplate,
		NewOrgWlan,
		NewOrgWxTag,
		NewOrgWxRule,
		NewSiteResource,
		NewSiteSettingResource,
		NewSiteNetworkTemplate,
		NewSiteWlan,
		NewSiteWxRule,
		NewSiteWxTag,
		NewDeviceApResource,
		NewDeviceSwitchResource,
		NewDeviceGatewayResource,
		NewDeviceGatewayClusterResource,
		NewOrgDeviceprofileAp,
		NewOrgDeviceprofileAssign,
		NewOrgDeviceprofileGateway,
	}
}
