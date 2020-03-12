// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package apigateway

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the name and value of a pre-existing API Key, for
// example to supply credentials for a dependency microservice.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/api_gateway_api_key.html.markdown.
func GetKey(ctx *pulumi.Context, args *GetKeyArgs, opts ...pulumi.InvokeOption) (*GetKeyResult, error) {
	var rv GetKeyResult
	err := ctx.Invoke("aws:apigateway/getKey:getKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKey.
type GetKeyArgs struct {
	// The ID of the API Key to look up.
	Id string `pulumi:"id"`
	Tags map[string]interface{} `pulumi:"tags"`
}


// A collection of values returned by getKey.
type GetKeyResult struct {
	// The date and time when the API Key was created.
	CreatedDate string `pulumi:"createdDate"`
	// The description of the API Key.
	Description string `pulumi:"description"`
	// Specifies whether the API Key is enabled.
	Enabled bool `pulumi:"enabled"`
	// Set to the ID of the API Key.
	Id string `pulumi:"id"`
	// The date and time when the API Key was last updated.
	LastUpdatedDate string `pulumi:"lastUpdatedDate"`
	// Set to the name of the API Key.
	Name string `pulumi:"name"`
	// A mapping of tags for the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// Set to the value of the API Key.
	Value string `pulumi:"value"`
}

