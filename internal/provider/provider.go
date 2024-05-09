// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure devopsApiProvider satisfies various provider interfaces.
var _ provider.Provider = &devopsApiProvider{}
var _ provider.ProviderWithFunctions = &devopsApiProvider{}

// devopsApiProvider defines the provider implementation.
type devopsApiProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

func (p *devopsApiProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "devopsApi"
	resp.Version = p.version
}

// devopsApiProviderModel maps provider schema data to a Go type.
type devopsApiProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

// user defines the endpoint value when declaring this provider in the TF configuration
func (p *devopsApiProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "API endpoint",
				Optional:            true,
			},
		},
	}
}

func (p *devopsApiProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data devopsApiProvider

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	client := http.DefaultClient
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *devopsApiProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewExampleResource,
	}
}

func (p *devopsApiProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewExampleDataSource,
	}
}

func (p *devopsApiProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewExampleFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &devopsApiProvider{
			version: version,
		}
	}
}
