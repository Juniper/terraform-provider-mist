package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"os"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	defaultTag        = "v0.0.0"
	defaultApiTimeout = 10
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
	Host       types.String  `tfsdk:"host"`
	Apitoken   types.String  `tfsdk:"apitoken"`
	Username   types.String  `tfsdk:"username"`
	Password   types.String  `tfsdk:"password"`
	ApiTimeout types.Float64 `tfsdk:"api_timeout"`
}

func (p *mistProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {

	resp.Schema = schema.Schema{
		MarkdownDescription: "The Mist Provider allows Terraform to manage Juniper Mist Organizations.\n\n" +
			"It is mainly focusing on day 0 and day 1 operations (provisionning and delpyment) but will be " +
			"completed over time.\n\nUse the navigation tree to the left to read about the available resources " +
			"and data sources.\n\nIt is possible to use API Token or Username/Password authentication (without 2FA)" +
			", but only one method should be configured.\n\nThis version is supporting the following Mist Clouds:\n" +
			"* Global 01 (api.mist.com)\n" +
			"* Global 02 (api.gc1.mist.com)\n" +
			"* Global 03 (api.ac2.mist.com)\n" +
			"* Global 04 (api.gc2.mist.com)\n" +
			"* EMEA 01 (api.eu.mist.com)\n" +
			"* EMEA 02 (api.gc3.mist.com)\n" +
			"* EMEA 03 (api.ac6.mist.com)\n" +
			"* APAC 01 (api.ac5.mist.com)",
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				MarkdownDescription: "URL of the Mist Cloud, e.g. `api.mist.com`.\n" +
					"It is also possible to pass the Mist Cloud host with the environment variable `" + envHost + "`.",
				Required: true,
			},
			"apitoken": schema.StringAttribute{
				MarkdownDescription: "For API Token authentication, the Mist API Token.\n" +
					"The preferred approach is to pass the API Token as environment variables `" + envApitoken + "`.",
				Optional:  true,
				Sensitive: true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "For username/password authentication, the Mist Account username.\n" +
					"The preferred approach is to pass the API Token as environment variables `" + envUsername + "`.",
				Optional: true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "For username/password authentication, the Mist Account password.\n" +
					"The preferred approach is to pass the API Token as environment variables `" + envPassword + "`.",
				Optional:  true,
				Sensitive: true,
			},
			"api_timeout": schema.Float64Attribute{
				MarkdownDescription: fmt.Sprintf("Timeout in seconds for completing API transactions "+
					"with the Mist Cloud. Omit for default value of %d seconds. Value of 0 results in "+
					"infinite timeout.",
					defaultApiTimeout),
				Optional:   true,
				Validators: []validator.Float64{float64validator.AtLeast(0)},
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

	host := os.Getenv(envHost)
	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	apitoken := os.Getenv(envApitoken)
	if !config.Apitoken.IsNull() {
		apitoken = config.Apitoken.ValueString()
	}

	username := os.Getenv(envUsername)
	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	password := os.Getenv(envPassword)
	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	var api_timeout float64 = 10
	if s, ok := os.LookupEnv(envApiTimeout); ok && config.ApiTimeout.IsNull() {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			diags.AddError(fmt.Sprintf("error parsing environment variable %q", envApiTimeout), err.Error())
		}
		if v < 0 {
			diags.AddError(fmt.Sprintf("invalid value in environment variable %q", envApiTimeout),
				fmt.Sprintf("minimum permitted value is 0, got %d", int64(v)))
		}
		config.ApiTimeout = types.Float64Value(v)
	} else if !config.ApiTimeout.IsNull() {
		api_timeout = config.ApiTimeout.ValueFloat64()
	}

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing MIST API Host",
			"The provider cannot create the MIST API client because there is a missing or empty value for the MIST API host. "+
				"Set the host value in the configuration or use the `"+envHost+"` environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
	if apitoken == "" && (username == "" && password == "") {
		resp.Diagnostics.AddError(
			"Missing MIST API Authentication",
			"The provider cannot create the MIST API client because the authentication configuration is missing. "+
				"Set the Authentication values in the configuration or in the environment variables: "+
				" * apitoken (environment variable `"+envApitoken+"`)"+
				" * username and password (environment variables`"+envUsername+"` and `"+envPassword+"`)"+
				"If either is already set, ensure the value is not empty.",
		)
	} else if apitoken == "" && (username != "" && password == "") {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing MIST API Password",
			"The provider cannot create the MIST API client because there is a  a missing or empty value for the MIST Username whereas the MIST Password is configured. "+
				"Set the host value in the configuration or use the `"+envUsername+"` environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	} else if apitoken == "" && (username == "" && password != "") {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing MIST API Password",
			"The provider cannot create the MIST API client because there is a  a missing or empty value for the MIST Password whereas the MIST Username is configured. "+
				"Set the host value in the configuration or use the `"+envPassword+"` environment variable. "+
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
			mistapi.WithHttpConfiguration(
				mistapi.CreateHttpConfiguration(
					mistapi.WithTimeout(api_timeout),
				),
			),
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
				mistapi.WithHttpConfiguration(
					mistapi.CreateHttpConfiguration(
						mistapi.WithTimeout(api_timeout),
					),
				),
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
			return
		} else if apiResponse.Response.StatusCode != 200 {
			resp.Diagnostics.AddError("Authentication Failed", "Incorrect login/password")
			return
		}

		// Process the Response Headers to extract the CSRF Token
		csrfTokenSet := false
		for hNAme, hVal := range apiResponse.Response.Header {
			if hNAme == "Set-Cookie" {
				for _, cooky := range hVal {
					for _, cVal := range strings.Split(cooky, ";") {
						if strings.HasPrefix(cVal, "csrftoken") {
							csrfToken_string := strings.Split(cVal, "=")[1]
							csrfToken := mistapi.NewCsrfTokenCredentials(string(csrfToken_string))
							client_config = mistapi.CreateConfiguration(
								mistapi.WithHttpConfiguration(
									mistapi.CreateHttpConfiguration(
										mistapi.WithTimeout(api_timeout),
									),
								),
								mistapi.WithEnvironment(mist_cloud),
								mistapi.WithBasicAuthCredentials(
									mistapi.NewBasicAuthCredentials(username, password),
								),
								mistapi.WithCsrfTokenCredentials(csrfToken),
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
			return
		}
	}

	// Use the  configuration to create the client and test the credentials
	client = mistapi.NewClient(client_config)
	_, err := client.SelfAccount().GetSelf(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Authentication Error", err.Error())
		return
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
		NewConstAppCategoriesDataSource,
		NewConstAppSubCategoriesDataSource,
		NewConstCountriesDataSource,
		NewConstTrafficTypesDataSource,
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
		NewOrgServicepoliciesDataSource,
		NewOrgIdpprofilesDataSource,
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
		NewOrgSettingResource,
		NewOrgServicepolicyResource,
		NewOrgIdpprofileResource,
	}
}
