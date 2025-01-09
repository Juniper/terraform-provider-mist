package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/apimatic/go-core-runtime/logger"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"os"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/function"
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

var (
	_ provider.Provider              = &mistProvider{}
	_ provider.ProviderWithFunctions = &mistProvider{}
)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &mistProvider{}
	}
}

type mistProvider struct {
	version string
}

type mistProviderModel struct {
	Host       types.String  `tfsdk:"host"`
	Apitoken   types.String  `tfsdk:"apitoken"`
	Username   types.String  `tfsdk:"username"`
	Password   types.String  `tfsdk:"password"`
	ApiTimeout types.Float64 `tfsdk:"api_timeout"`
	ApiDebug   types.Bool    `tfsdk:"api_debug"`
	Proxy      types.String  `tfsdk:"proxy"`
}

func (p *mistProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {

	resp.Schema = schema.Schema{
		MarkdownDescription: "The Mist Provider allows Terraform to manage Juniper Mist Organizations.\n\n" +
			"It is mainly focusing on day 0 and day 1 operations (provisionning and delpyment) but will be " +
			"completed over time.\n\nUse the navigation tree to the left to read about the available resources " +
			"and data sources.\n\nIt is possible to use API Token or Username/Password authentication (without 2FA)" +
			", but only one method should be configured.\n\n## Supported Mist Clouds\n\nThis provider can be used with " +
			"the following Mist Clouds:\n" +
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
				MarkdownDescription: "URL of the Mist Cloud, e.g. `api.mist.com`.",
				Optional:            true,
			},
			"apitoken": schema.StringAttribute{
				MarkdownDescription: "For API Token authentication, the Mist API Token.",
				Optional:            true,
				Sensitive:           true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "For username/password authentication, the Mist Account username.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "For username/password authentication, the Mist Account password.",
				Optional:            true,
				Sensitive:           true,
			},
			"api_debug": schema.BoolAttribute{
				MarkdownDescription: "Flag to enable debugging API calls. Default is false.",
				Optional:            true,
			},
			"api_timeout": schema.Float64Attribute{
				MarkdownDescription: fmt.Sprintf("Timeout in seconds for completing API transactions "+
					"with the Mist Cloud. Omit for default value of %d seconds. Value of 0 results in "+
					"infinite timeout.",
					defaultApiTimeout),
				Optional:   true,
				Validators: []validator.Float64{float64validator.AtLeast(0)},
			},
			"proxy": schema.StringAttribute{
				MarkdownDescription: "Requests use the configured proxy to reach the Mist Cloud.\n" +
					"The value may be either a complete URL or a `[username:password@]host[:port]`, in which case the `http` scheme is assumed. " +
					"The schemes `http`, `https`, and `socks5` are supported.",
				Optional: true,
			},
		},
	}
}

func (p *mistProviderModel) fromEnv(_ context.Context, diags *diag.Diagnostics) {
	if s, ok := os.LookupEnv(envHost); ok && p.Host.IsNull() {
		if !strings.HasPrefix(s, "api.") {
			diags.AddError(fmt.Sprintf("error parsing environment variable %q", envHost),
				fmt.Sprintf("The configured Mist Host does not match the supported Clouds; got %q", s))
		}
		p.Host = types.StringValue(s)
	}

	if s, ok := os.LookupEnv(envApitoken); ok && p.Apitoken.IsNull() {
		if len(s) < 1 {
			diags.AddError(fmt.Sprintf("error parsing environment variable %q", envApitoken),
				fmt.Sprintf("minimum string length 1; got %q", s))
		}
		p.Apitoken = types.StringValue(s)
	}

	if s, ok := os.LookupEnv(envUsername); ok && p.Username.IsNull() {
		if len(s) < 1 {
			diags.AddError(fmt.Sprintf("error parsing environment variable %q", envUsername),
				fmt.Sprintf("minimum string length 1; got %q", s))
		}
		p.Username = types.StringValue(s)
	}

	if s, ok := os.LookupEnv(envPassword); ok && p.Password.IsNull() {
		if len(s) < 1 {
			diags.AddError(fmt.Sprintf("error parsing environment variable %q", envPassword),
				fmt.Sprintf("minimum string length 1; got %q", s))
		}
		p.Password = types.StringValue(s)
	}

	if s, ok := os.LookupEnv(envApiTimeout); ok && p.ApiTimeout.IsNull() {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			diags.AddError(fmt.Sprintf("error parsing environment variable %q", envApiTimeout), err.Error())
		}
		if v < 0 {
			diags.AddError(fmt.Sprintf("invalid value in environment variable %q", envApiTimeout),
				fmt.Sprintf("minimum permitted value is 0, got %d", int64(v)))
		}
		p.ApiTimeout = types.Float64Value(v)
	}

	if s, ok := os.LookupEnv(envProxy); ok && p.Proxy.IsNull() {
		if len(s) < 1 {
			diags.AddError(fmt.Sprintf("error parsing environment variable %q", envProxy),
				fmt.Sprintf("minimum string length 1; got %q", s))
		}
		p.Proxy = types.StringValue(s)
	}

	if s, ok := os.LookupEnv(envDebug); ok && p.ApiDebug.IsNull() {
		v, err := strconv.ParseBool(s)
		if err != nil {
			diags.AddError(fmt.Sprintf("error parsing MIST_API_DEBUG environment variable %q", envDebug), err.Error())
		}
		p.ApiDebug = types.BoolValue(v)
	}
}

func (p *mistProviderModel) validateConfig(_ context.Context, diags *diag.Diagnostics) {
	if p.Host.ValueString() == "" {
		diags.AddAttributeError(
			path.Root("host"),
			"Missing MIST API Host",
			"The provider cannot create the MIST API client because there is a missing or empty value for the MIST API host. "+
				"Set the host value in the configuration or use the `"+envHost+"` environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if p.Apitoken.ValueString() == "" && (p.Username.ValueString() == "" && p.Password.ValueString() == "") {
		diags.AddError(
			"Missing MIST API Authentication",
			"The provider cannot create the MIST API client because the authentication configuration is missing. "+
				"Set the Authentication values in the configuration or in the environment variables: "+
				" * apitoken (environment variable `"+envApitoken+"`)"+
				" * username and password (environment variables`"+envUsername+"` and `"+envPassword+"`)"+
				"If either is already set, ensure the value is not empty.",
		)
	} else if p.Apitoken.ValueString() == "" && (p.Username.ValueString() != "" && p.Password.ValueString() == "") {
		diags.AddAttributeError(
			path.Root("username"),
			"Missing MIST API Password",
			"The provider cannot create the MIST API client because there is a  a missing or empty value for the MIST Username whereas the MIST Password is configured. "+
				"Set the host value in the configuration or use the `"+envUsername+"` environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	} else if p.Apitoken.ValueString() == "" && (p.Username.ValueString() == "" && p.Password.ValueString() != "") {
		diags.AddAttributeError(
			path.Root("password"),
			"Missing MIST API Password",
			"The provider cannot create the MIST API client because there is a  a missing or empty value for the MIST Password whereas the MIST Username is configured. "+
				"Set the host value in the configuration or use the `"+envPassword+"` environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
}

func (p *mistProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config mistProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieve missing config elements from environment
	config.fromEnv(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	config.validateConfig(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.ApiTimeout.IsNull() {
		config.ApiTimeout = types.Float64Value(defaultApiTimeout)
	}

	var proxy_url *url.URL
	if !config.Proxy.IsNull() {
		proxy_string := config.Proxy.ValueString()
		if !strings.HasPrefix(proxy_string, "http://") &&
			!strings.HasPrefix(proxy_string, "https://") &&
			!strings.HasPrefix(proxy_string, "socks5://") {
			proxy_string = "http://" + proxy_string
		}
		u, err := url.Parse(proxy_string)
		if err != nil {
			resp.Diagnostics.AddError("Unable to parse proxy configuration", err.Error())
			return
		}
		proxy_url = u
	}

	var mist_cloud mistapi.Environment
	switch config.Host.ValueString() {
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
	case "api.us.mist-federal.com":
		mist_cloud = mistapi.GOV_CLOUD
	default:
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Wrong Mist Host",
			"The configured host \""+config.Host.ValueString()+"\" is not a valid Mist Host. Please refer to the documentation to get the possible values",
		)
		return
	}

	var DefaultTransport http.RoundTripper = http.DefaultTransport
	if proxy_url != nil {
		DefaultTransport = &http.Transport{
			Proxy: http.ProxyURL(proxy_url),
		}
	}

	var client_config mistapi.Configuration

	var configOptions []mistapi.ConfigurationOptions
	var tfLogger = NewTFlogger(ctx)
	if config.ApiDebug.ValueBool() {
		tfLogger.Warn("API Debugging enabled")
		resp.Diagnostics.AddWarning("API Debugging enabled", "API debugging is enabled. Request and response bodies, headers, or other sensitive data may be logged. Use with caution.")
		loggerConfig := mistapi.WithLoggerConfiguration(
			mistapi.WithLogger(tfLogger),
			mistapi.WithLevel(logger.Level_DEBUG),
			mistapi.WithRequestConfiguration(
				mistapi.WithRequestBody(true),
				mistapi.WithRequestHeaders(true),
			),
			mistapi.WithResponseConfiguration(
				mistapi.WithResponseHeaders(true),
				mistapi.WithResponseBody(true),
			),
		)
		configOptions = append(configOptions, loggerConfig)
	}
	configOptions = append(configOptions, mistapi.WithEnvironment(mist_cloud))

	var httpConfig mistapi.ConfigurationOptions = mistapi.WithHttpConfiguration(
		mistapi.CreateHttpConfiguration(
			mistapi.WithTimeout(config.ApiTimeout.ValueFloat64()),
			mistapi.WithTransport(DefaultTransport),
		),
	)
	configOptions = append(configOptions, httpConfig)

	// configure the client for API Token Auth
	if config.Apitoken.ValueString() != "" {

		var apiTokenConfig = mistapi.WithApiTokenCredentials(
			mistapi.NewApiTokenCredentials("Token " + config.Apitoken.ValueString()),
		)
		configOptions = append(configOptions, apiTokenConfig)

		client_config = mistapi.CreateConfiguration(configOptions...)

		// configure the client for Basic Auth + CSRF
	} else {
		// Initiate the login API Call
		var basic_auth_config = mistapi.WithBasicAuthCredentials(
			mistapi.NewBasicAuthCredentials(config.Username.ValueString(), config.Password.ValueString()),
		)
		configOptions = append(configOptions, basic_auth_config)
		tmp_client := mistapi.NewClient(
			mistapi.CreateConfiguration(configOptions...),
		)

		body := models.Login{}
		body.Email = config.Username.ValueString()
		body.Password = config.Password.ValueString()
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
							csrfTokenConfig := mistapi.WithCsrfTokenCredentials(
								mistapi.NewCsrfTokenCredentials(csrfToken_string),
							)
							configOptions = append(configOptions, csrfTokenConfig)
							client_config = mistapi.CreateConfiguration(configOptions...)
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
	var client mistapi.ClientInterface = mistapi.NewClient(client_config)
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
		NewConstApplicationsDataSource,
		NewConstAppSubCategoriesDataSource,
		NewConstCountriesDataSource,
		NewConstTrafficTypesDataSource,
		NewDeviceApStatsDataSource,
		NewDeviceSwitchStatsDataSource,
		NewDeviceGatewayStatsDataSource,
		NewOrgGatewaytemplatesDataSource,
		NewOrgInventoryDataSource,
		NewOrgAlarmtemplatesDataSource,
		NewOrgNacrulesDataSource,
		NewOrgNactagsDataSource,
		NewOrgNacEndpointsDataSource,
		NewOrgNetworksDataSource,
		NewOrgNetworktemplatesDataSource,
		NewOrgRftemplatesDataSource,
		NewOrgServicesDataSource,
		NewOrgSitegroupsDataSource,
		NewOrgVpnsDataSource,
		NewOrgSsoRolesDataSource,
		NewOrgWlantemplatesDataSource,
		NewOrgWxtagsDataSource,
		NewSitesDataSource,
		NewOrgDeviceprofilesApDataSource,
		NewOrgDeviceprofilesGatewayDataSource,
		NewOrgServicepoliciesDataSource,
		NewOrgIdpprofilesDataSource,
		NewOrgPsksDataSource,
		NewSitePsksDataSource,
		NewOrgWebhooksDataSource,
		NewSiteWebhooksDataSource,
		NewConstAlarmsDataSource,
		NewOrgNacidpMetadataDataSource,
		NewOrgSsoMetadataDataSource,
		NewConstWebhooksDataSource,
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
		NewOrgNacIdp,
		NewOrgNacEndpointResource,
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
		NewDeviceImage,
		NewOrgDeviceprofileAp,
		NewOrgDeviceprofileAssign,
		NewOrgDeviceprofileGateway,
		NewOrgSettingResource,
		NewOrgServicepolicyResource,
		NewOrgIdpprofileResource,
		NewOrgPsk,
		NewSitePsk,
		NewOrgWebhook,
		NewSiteWebhook,
		NewOrgWlanPortalImage,
		NewOrgWlanPortalTemplate,
		NewSiteWlanPortalImage,
		NewSiteWlanPortalTemplate,
		NewOrgApitoken,
		NewOrgSso,
		NewOrgSsoRole,
		NewOrgAlarmtemplateResource,
		NewOrgEvpnTopologyResource,
		NewSiteEvpnTopologyResource,
	}
}

func (p *mistProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{
		NewSearchInventoryByClaimcodeFunction,
		NewSearchInventoryByMacFunction,
		NewSearchInventoryBySerialFunction,
	}
}
