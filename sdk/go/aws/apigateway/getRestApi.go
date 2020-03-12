// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package apigateway

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the id and rootResourceId of a REST API in
// API Gateway. To fetch the REST API you must provide a name to match against. 
// As there is no unique name constraint on REST APIs this data source will 
// error if there is more than one match.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/api_gateway_rest_api.html.markdown.
func LookupRestApi(ctx *pulumi.Context, args *LookupRestApiArgs, opts ...pulumi.InvokeOption) (*LookupRestApiResult, error) {
	var rv LookupRestApiResult
	err := ctx.Invoke("aws:apigateway/getRestApi:getRestApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRestApi.
type LookupRestApiArgs struct {
	// The name of the REST API to look up. If no REST API is found with this name, an error will be returned. If multiple REST APIs are found with this name, an error will be returned.
	Name string `pulumi:"name"`
	Tags map[string]interface{} `pulumi:"tags"`
}


// A collection of values returned by getRestApi.
type LookupRestApiResult struct {
	// The source of the API key for requests.
	ApiKeySource string `pulumi:"apiKeySource"`
	// The ARN of the REST API.
	Arn string `pulumi:"arn"`
	// The list of binary media types supported by the REST API.
	BinaryMediaTypes []string `pulumi:"binaryMediaTypes"`
	// The description of the REST API.
	Description string `pulumi:"description"`
	// The endpoint configuration of this RestApi showing the endpoint types of the API.
	EndpointConfigurations []GetRestApiEndpointConfiguration `pulumi:"endpointConfigurations"`
	// The execution ARN part to be used in [`lambdaPermission`](https://www.terraform.io/docs/providers/aws/r/lambda_permission.html)'s `sourceArn` when allowing API Gateway to invoke a Lambda function, e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j`, which can be concatenated with allowed stage, method and resource path.
	ExecutionArn string `pulumi:"executionArn"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Minimum response size to compress for the REST API.
	MinimumCompressionSize int `pulumi:"minimumCompressionSize"`
	Name string `pulumi:"name"`
	// JSON formatted policy document that controls access to the API Gateway.
	Policy string `pulumi:"policy"`
	// Set to the ID of the API Gateway Resource on the found REST API where the route matches '/'.
	RootResourceId string `pulumi:"rootResourceId"`
	// Key-value mapping of resource tags.
	Tags map[string]interface{} `pulumi:"tags"`
}

