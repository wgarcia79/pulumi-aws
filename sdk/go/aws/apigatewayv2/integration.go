// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Amazon API Gateway Version 2 integration.
// More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigatewayv2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigatewayv2.NewIntegration(ctx, "example", &apigatewayv2.IntegrationArgs{
// 			ApiId:           pulumi.Any(aws_apigatewayv2_api.Example.Id),
// 			IntegrationType: pulumi.String("MOCK"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### AWS Service Integration
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigatewayv2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigatewayv2.NewIntegration(ctx, "example", &apigatewayv2.IntegrationArgs{
// 			ApiId:              pulumi.Any(aws_apigatewayv2_api.Example.Id),
// 			CredentialsArn:     pulumi.Any(aws_iam_role.Example.Arn),
// 			Description:        pulumi.String("SQS example"),
// 			IntegrationType:    pulumi.String("AWS_PROXY"),
// 			IntegrationSubtype: pulumi.String("SQS-SendMessage"),
// 			RequestParameters: pulumi.StringMap{
// 				"QueueUrl":    pulumi.String(fmt.Sprintf("%v%v", "$", "request.header.queueUrl")),
// 				"MessageBody": pulumi.String(fmt.Sprintf("%v%v", "$", "request.body.message")),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// `aws_apigatewayv2_integration` can be imported by using the API identifier and integration identifier, e.g.
//
// ```sh
//  $ pulumi import aws:apigatewayv2/integration:Integration example aabbccddee/1122334
// ```
type Integration struct {
	pulumi.CustomResourceState

	// The API identifier.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// The ID of the VPC link for a private integration. Supported only for HTTP APIs. Must be between 1 and 1024 characters in length.
	ConnectionId pulumi.StringPtrOutput `pulumi:"connectionId"`
	// The type of the network connection to the integration endpoint. Valid values: `INTERNET`, `VPC_LINK`. Default is `INTERNET`.
	ConnectionType pulumi.StringPtrOutput `pulumi:"connectionType"`
	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`. Supported only for WebSocket APIs.
	ContentHandlingStrategy pulumi.StringPtrOutput `pulumi:"contentHandlingStrategy"`
	// The credentials required for the integration, if any.
	CredentialsArn pulumi.StringPtrOutput `pulumi:"credentialsArn"`
	// The description of the integration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration's HTTP method. Must be specified if `integrationType` is not `MOCK`.
	IntegrationMethod pulumi.StringPtrOutput `pulumi:"integrationMethod"`
	// The [integration response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-integration-response-selection-expressions) for the integration.
	IntegrationResponseSelectionExpression pulumi.StringOutput `pulumi:"integrationResponseSelectionExpression"`
	// Specifies the AWS service action to invoke. Supported only for HTTP APIs when `integrationType` is `AWS_PROXY`. See the [AWS service integration reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html) documentation for supported values. Must be between 1 and 128 characters in length.
	IntegrationSubtype pulumi.StringPtrOutput `pulumi:"integrationSubtype"`
	// The integration type of an integration.
	// Valid values: `AWS`, `AWS_PROXY`, `HTTP`, `HTTP_PROXY`, `MOCK`.
	IntegrationType pulumi.StringOutput `pulumi:"integrationType"`
	// The URI of the Lambda function for a Lambda proxy integration, when `integrationType` is `AWS_PROXY`.
	// For an `HTTP` integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
	IntegrationUri pulumi.StringPtrOutput `pulumi:"integrationUri"`
	// The pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the `requestTemplates` attribute.
	// Valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`. Default is `WHEN_NO_MATCH`. Supported only for WebSocket APIs.
	PassthroughBehavior pulumi.StringPtrOutput `pulumi:"passthroughBehavior"`
	// The [format of the payload](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html#http-api-develop-integrations-lambda.proxy-format) sent to an integration. Valid values: `1.0`, `2.0`. Default is `1.0`.
	PayloadFormatVersion pulumi.StringPtrOutput `pulumi:"payloadFormatVersion"`
	// For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend.
	// For HTTP APIs with a specified `integrationSubtype`, a key-value map specifying parameters that are passed to `AWS_PROXY` integrations.
	// For HTTP APIs without a specified `integrationSubtype`, a key-value map specifying how to transform HTTP requests before sending them to the backend.
	// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html) for details.
	RequestParameters pulumi.StringMapOutput `pulumi:"requestParameters"`
	// A map of [Velocity](https://velocity.apache.org/) templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
	RequestTemplates pulumi.StringMapOutput `pulumi:"requestTemplates"`
	// Mappings to transform the HTTP response from a backend integration before returning the response to clients. Supported only for HTTP APIs.
	ResponseParameters IntegrationResponseParameterArrayOutput `pulumi:"responseParameters"`
	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration.
	TemplateSelectionExpression pulumi.StringPtrOutput `pulumi:"templateSelectionExpression"`
	TimeoutMilliseconds         pulumi.IntOutput       `pulumi:"timeoutMilliseconds"`
	// The TLS configuration for a private integration. Supported only for HTTP APIs.
	TlsConfig IntegrationTlsConfigPtrOutput `pulumi:"tlsConfig"`
}

// NewIntegration registers a new resource with the given unique name, arguments, and options.
func NewIntegration(ctx *pulumi.Context,
	name string, args *IntegrationArgs, opts ...pulumi.ResourceOption) (*Integration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.IntegrationType == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationType'")
	}
	var resource Integration
	err := ctx.RegisterResource("aws:apigatewayv2/integration:Integration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegration gets an existing Integration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationState, opts ...pulumi.ResourceOption) (*Integration, error) {
	var resource Integration
	err := ctx.ReadResource("aws:apigatewayv2/integration:Integration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Integration resources.
type integrationState struct {
	// The API identifier.
	ApiId *string `pulumi:"apiId"`
	// The ID of the VPC link for a private integration. Supported only for HTTP APIs. Must be between 1 and 1024 characters in length.
	ConnectionId *string `pulumi:"connectionId"`
	// The type of the network connection to the integration endpoint. Valid values: `INTERNET`, `VPC_LINK`. Default is `INTERNET`.
	ConnectionType *string `pulumi:"connectionType"`
	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`. Supported only for WebSocket APIs.
	ContentHandlingStrategy *string `pulumi:"contentHandlingStrategy"`
	// The credentials required for the integration, if any.
	CredentialsArn *string `pulumi:"credentialsArn"`
	// The description of the integration.
	Description *string `pulumi:"description"`
	// The integration's HTTP method. Must be specified if `integrationType` is not `MOCK`.
	IntegrationMethod *string `pulumi:"integrationMethod"`
	// The [integration response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-integration-response-selection-expressions) for the integration.
	IntegrationResponseSelectionExpression *string `pulumi:"integrationResponseSelectionExpression"`
	// Specifies the AWS service action to invoke. Supported only for HTTP APIs when `integrationType` is `AWS_PROXY`. See the [AWS service integration reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html) documentation for supported values. Must be between 1 and 128 characters in length.
	IntegrationSubtype *string `pulumi:"integrationSubtype"`
	// The integration type of an integration.
	// Valid values: `AWS`, `AWS_PROXY`, `HTTP`, `HTTP_PROXY`, `MOCK`.
	IntegrationType *string `pulumi:"integrationType"`
	// The URI of the Lambda function for a Lambda proxy integration, when `integrationType` is `AWS_PROXY`.
	// For an `HTTP` integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
	IntegrationUri *string `pulumi:"integrationUri"`
	// The pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the `requestTemplates` attribute.
	// Valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`. Default is `WHEN_NO_MATCH`. Supported only for WebSocket APIs.
	PassthroughBehavior *string `pulumi:"passthroughBehavior"`
	// The [format of the payload](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html#http-api-develop-integrations-lambda.proxy-format) sent to an integration. Valid values: `1.0`, `2.0`. Default is `1.0`.
	PayloadFormatVersion *string `pulumi:"payloadFormatVersion"`
	// For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend.
	// For HTTP APIs with a specified `integrationSubtype`, a key-value map specifying parameters that are passed to `AWS_PROXY` integrations.
	// For HTTP APIs without a specified `integrationSubtype`, a key-value map specifying how to transform HTTP requests before sending them to the backend.
	// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html) for details.
	RequestParameters map[string]string `pulumi:"requestParameters"`
	// A map of [Velocity](https://velocity.apache.org/) templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
	RequestTemplates map[string]string `pulumi:"requestTemplates"`
	// Mappings to transform the HTTP response from a backend integration before returning the response to clients. Supported only for HTTP APIs.
	ResponseParameters []IntegrationResponseParameter `pulumi:"responseParameters"`
	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration.
	TemplateSelectionExpression *string `pulumi:"templateSelectionExpression"`
	TimeoutMilliseconds         *int    `pulumi:"timeoutMilliseconds"`
	// The TLS configuration for a private integration. Supported only for HTTP APIs.
	TlsConfig *IntegrationTlsConfig `pulumi:"tlsConfig"`
}

type IntegrationState struct {
	// The API identifier.
	ApiId pulumi.StringPtrInput
	// The ID of the VPC link for a private integration. Supported only for HTTP APIs. Must be between 1 and 1024 characters in length.
	ConnectionId pulumi.StringPtrInput
	// The type of the network connection to the integration endpoint. Valid values: `INTERNET`, `VPC_LINK`. Default is `INTERNET`.
	ConnectionType pulumi.StringPtrInput
	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`. Supported only for WebSocket APIs.
	ContentHandlingStrategy pulumi.StringPtrInput
	// The credentials required for the integration, if any.
	CredentialsArn pulumi.StringPtrInput
	// The description of the integration.
	Description pulumi.StringPtrInput
	// The integration's HTTP method. Must be specified if `integrationType` is not `MOCK`.
	IntegrationMethod pulumi.StringPtrInput
	// The [integration response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-integration-response-selection-expressions) for the integration.
	IntegrationResponseSelectionExpression pulumi.StringPtrInput
	// Specifies the AWS service action to invoke. Supported only for HTTP APIs when `integrationType` is `AWS_PROXY`. See the [AWS service integration reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html) documentation for supported values. Must be between 1 and 128 characters in length.
	IntegrationSubtype pulumi.StringPtrInput
	// The integration type of an integration.
	// Valid values: `AWS`, `AWS_PROXY`, `HTTP`, `HTTP_PROXY`, `MOCK`.
	IntegrationType pulumi.StringPtrInput
	// The URI of the Lambda function for a Lambda proxy integration, when `integrationType` is `AWS_PROXY`.
	// For an `HTTP` integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
	IntegrationUri pulumi.StringPtrInput
	// The pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the `requestTemplates` attribute.
	// Valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`. Default is `WHEN_NO_MATCH`. Supported only for WebSocket APIs.
	PassthroughBehavior pulumi.StringPtrInput
	// The [format of the payload](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html#http-api-develop-integrations-lambda.proxy-format) sent to an integration. Valid values: `1.0`, `2.0`. Default is `1.0`.
	PayloadFormatVersion pulumi.StringPtrInput
	// For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend.
	// For HTTP APIs with a specified `integrationSubtype`, a key-value map specifying parameters that are passed to `AWS_PROXY` integrations.
	// For HTTP APIs without a specified `integrationSubtype`, a key-value map specifying how to transform HTTP requests before sending them to the backend.
	// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html) for details.
	RequestParameters pulumi.StringMapInput
	// A map of [Velocity](https://velocity.apache.org/) templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
	RequestTemplates pulumi.StringMapInput
	// Mappings to transform the HTTP response from a backend integration before returning the response to clients. Supported only for HTTP APIs.
	ResponseParameters IntegrationResponseParameterArrayInput
	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration.
	TemplateSelectionExpression pulumi.StringPtrInput
	TimeoutMilliseconds         pulumi.IntPtrInput
	// The TLS configuration for a private integration. Supported only for HTTP APIs.
	TlsConfig IntegrationTlsConfigPtrInput
}

func (IntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationState)(nil)).Elem()
}

type integrationArgs struct {
	// The API identifier.
	ApiId string `pulumi:"apiId"`
	// The ID of the VPC link for a private integration. Supported only for HTTP APIs. Must be between 1 and 1024 characters in length.
	ConnectionId *string `pulumi:"connectionId"`
	// The type of the network connection to the integration endpoint. Valid values: `INTERNET`, `VPC_LINK`. Default is `INTERNET`.
	ConnectionType *string `pulumi:"connectionType"`
	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`. Supported only for WebSocket APIs.
	ContentHandlingStrategy *string `pulumi:"contentHandlingStrategy"`
	// The credentials required for the integration, if any.
	CredentialsArn *string `pulumi:"credentialsArn"`
	// The description of the integration.
	Description *string `pulumi:"description"`
	// The integration's HTTP method. Must be specified if `integrationType` is not `MOCK`.
	IntegrationMethod *string `pulumi:"integrationMethod"`
	// Specifies the AWS service action to invoke. Supported only for HTTP APIs when `integrationType` is `AWS_PROXY`. See the [AWS service integration reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html) documentation for supported values. Must be between 1 and 128 characters in length.
	IntegrationSubtype *string `pulumi:"integrationSubtype"`
	// The integration type of an integration.
	// Valid values: `AWS`, `AWS_PROXY`, `HTTP`, `HTTP_PROXY`, `MOCK`.
	IntegrationType string `pulumi:"integrationType"`
	// The URI of the Lambda function for a Lambda proxy integration, when `integrationType` is `AWS_PROXY`.
	// For an `HTTP` integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
	IntegrationUri *string `pulumi:"integrationUri"`
	// The pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the `requestTemplates` attribute.
	// Valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`. Default is `WHEN_NO_MATCH`. Supported only for WebSocket APIs.
	PassthroughBehavior *string `pulumi:"passthroughBehavior"`
	// The [format of the payload](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html#http-api-develop-integrations-lambda.proxy-format) sent to an integration. Valid values: `1.0`, `2.0`. Default is `1.0`.
	PayloadFormatVersion *string `pulumi:"payloadFormatVersion"`
	// For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend.
	// For HTTP APIs with a specified `integrationSubtype`, a key-value map specifying parameters that are passed to `AWS_PROXY` integrations.
	// For HTTP APIs without a specified `integrationSubtype`, a key-value map specifying how to transform HTTP requests before sending them to the backend.
	// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html) for details.
	RequestParameters map[string]string `pulumi:"requestParameters"`
	// A map of [Velocity](https://velocity.apache.org/) templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
	RequestTemplates map[string]string `pulumi:"requestTemplates"`
	// Mappings to transform the HTTP response from a backend integration before returning the response to clients. Supported only for HTTP APIs.
	ResponseParameters []IntegrationResponseParameter `pulumi:"responseParameters"`
	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration.
	TemplateSelectionExpression *string `pulumi:"templateSelectionExpression"`
	TimeoutMilliseconds         *int    `pulumi:"timeoutMilliseconds"`
	// The TLS configuration for a private integration. Supported only for HTTP APIs.
	TlsConfig *IntegrationTlsConfig `pulumi:"tlsConfig"`
}

// The set of arguments for constructing a Integration resource.
type IntegrationArgs struct {
	// The API identifier.
	ApiId pulumi.StringInput
	// The ID of the VPC link for a private integration. Supported only for HTTP APIs. Must be between 1 and 1024 characters in length.
	ConnectionId pulumi.StringPtrInput
	// The type of the network connection to the integration endpoint. Valid values: `INTERNET`, `VPC_LINK`. Default is `INTERNET`.
	ConnectionType pulumi.StringPtrInput
	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`. Supported only for WebSocket APIs.
	ContentHandlingStrategy pulumi.StringPtrInput
	// The credentials required for the integration, if any.
	CredentialsArn pulumi.StringPtrInput
	// The description of the integration.
	Description pulumi.StringPtrInput
	// The integration's HTTP method. Must be specified if `integrationType` is not `MOCK`.
	IntegrationMethod pulumi.StringPtrInput
	// Specifies the AWS service action to invoke. Supported only for HTTP APIs when `integrationType` is `AWS_PROXY`. See the [AWS service integration reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html) documentation for supported values. Must be between 1 and 128 characters in length.
	IntegrationSubtype pulumi.StringPtrInput
	// The integration type of an integration.
	// Valid values: `AWS`, `AWS_PROXY`, `HTTP`, `HTTP_PROXY`, `MOCK`.
	IntegrationType pulumi.StringInput
	// The URI of the Lambda function for a Lambda proxy integration, when `integrationType` is `AWS_PROXY`.
	// For an `HTTP` integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
	IntegrationUri pulumi.StringPtrInput
	// The pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the `requestTemplates` attribute.
	// Valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`. Default is `WHEN_NO_MATCH`. Supported only for WebSocket APIs.
	PassthroughBehavior pulumi.StringPtrInput
	// The [format of the payload](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html#http-api-develop-integrations-lambda.proxy-format) sent to an integration. Valid values: `1.0`, `2.0`. Default is `1.0`.
	PayloadFormatVersion pulumi.StringPtrInput
	// For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend.
	// For HTTP APIs with a specified `integrationSubtype`, a key-value map specifying parameters that are passed to `AWS_PROXY` integrations.
	// For HTTP APIs without a specified `integrationSubtype`, a key-value map specifying how to transform HTTP requests before sending them to the backend.
	// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html) for details.
	RequestParameters pulumi.StringMapInput
	// A map of [Velocity](https://velocity.apache.org/) templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
	RequestTemplates pulumi.StringMapInput
	// Mappings to transform the HTTP response from a backend integration before returning the response to clients. Supported only for HTTP APIs.
	ResponseParameters IntegrationResponseParameterArrayInput
	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration.
	TemplateSelectionExpression pulumi.StringPtrInput
	TimeoutMilliseconds         pulumi.IntPtrInput
	// The TLS configuration for a private integration. Supported only for HTTP APIs.
	TlsConfig IntegrationTlsConfigPtrInput
}

func (IntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationArgs)(nil)).Elem()
}

type IntegrationInput interface {
	pulumi.Input

	ToIntegrationOutput() IntegrationOutput
	ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput
}

func (*Integration) ElementType() reflect.Type {
	return reflect.TypeOf((*Integration)(nil))
}

func (i *Integration) ToIntegrationOutput() IntegrationOutput {
	return i.ToIntegrationOutputWithContext(context.Background())
}

func (i *Integration) ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationOutput)
}

func (i *Integration) ToIntegrationPtrOutput() IntegrationPtrOutput {
	return i.ToIntegrationPtrOutputWithContext(context.Background())
}

func (i *Integration) ToIntegrationPtrOutputWithContext(ctx context.Context) IntegrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationPtrOutput)
}

type IntegrationPtrInput interface {
	pulumi.Input

	ToIntegrationPtrOutput() IntegrationPtrOutput
	ToIntegrationPtrOutputWithContext(ctx context.Context) IntegrationPtrOutput
}

type integrationPtrType IntegrationArgs

func (*integrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Integration)(nil))
}

func (i *integrationPtrType) ToIntegrationPtrOutput() IntegrationPtrOutput {
	return i.ToIntegrationPtrOutputWithContext(context.Background())
}

func (i *integrationPtrType) ToIntegrationPtrOutputWithContext(ctx context.Context) IntegrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationPtrOutput)
}

// IntegrationArrayInput is an input type that accepts IntegrationArray and IntegrationArrayOutput values.
// You can construct a concrete instance of `IntegrationArrayInput` via:
//
//          IntegrationArray{ IntegrationArgs{...} }
type IntegrationArrayInput interface {
	pulumi.Input

	ToIntegrationArrayOutput() IntegrationArrayOutput
	ToIntegrationArrayOutputWithContext(context.Context) IntegrationArrayOutput
}

type IntegrationArray []IntegrationInput

func (IntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Integration)(nil))
}

func (i IntegrationArray) ToIntegrationArrayOutput() IntegrationArrayOutput {
	return i.ToIntegrationArrayOutputWithContext(context.Background())
}

func (i IntegrationArray) ToIntegrationArrayOutputWithContext(ctx context.Context) IntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationArrayOutput)
}

// IntegrationMapInput is an input type that accepts IntegrationMap and IntegrationMapOutput values.
// You can construct a concrete instance of `IntegrationMapInput` via:
//
//          IntegrationMap{ "key": IntegrationArgs{...} }
type IntegrationMapInput interface {
	pulumi.Input

	ToIntegrationMapOutput() IntegrationMapOutput
	ToIntegrationMapOutputWithContext(context.Context) IntegrationMapOutput
}

type IntegrationMap map[string]IntegrationInput

func (IntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Integration)(nil))
}

func (i IntegrationMap) ToIntegrationMapOutput() IntegrationMapOutput {
	return i.ToIntegrationMapOutputWithContext(context.Background())
}

func (i IntegrationMap) ToIntegrationMapOutputWithContext(ctx context.Context) IntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationMapOutput)
}

type IntegrationOutput struct {
	*pulumi.OutputState
}

func (IntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Integration)(nil))
}

func (o IntegrationOutput) ToIntegrationOutput() IntegrationOutput {
	return o
}

func (o IntegrationOutput) ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput {
	return o
}

func (o IntegrationOutput) ToIntegrationPtrOutput() IntegrationPtrOutput {
	return o.ToIntegrationPtrOutputWithContext(context.Background())
}

func (o IntegrationOutput) ToIntegrationPtrOutputWithContext(ctx context.Context) IntegrationPtrOutput {
	return o.ApplyT(func(v Integration) *Integration {
		return &v
	}).(IntegrationPtrOutput)
}

type IntegrationPtrOutput struct {
	*pulumi.OutputState
}

func (IntegrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Integration)(nil))
}

func (o IntegrationPtrOutput) ToIntegrationPtrOutput() IntegrationPtrOutput {
	return o
}

func (o IntegrationPtrOutput) ToIntegrationPtrOutputWithContext(ctx context.Context) IntegrationPtrOutput {
	return o
}

type IntegrationArrayOutput struct{ *pulumi.OutputState }

func (IntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Integration)(nil))
}

func (o IntegrationArrayOutput) ToIntegrationArrayOutput() IntegrationArrayOutput {
	return o
}

func (o IntegrationArrayOutput) ToIntegrationArrayOutputWithContext(ctx context.Context) IntegrationArrayOutput {
	return o
}

func (o IntegrationArrayOutput) Index(i pulumi.IntInput) IntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Integration {
		return vs[0].([]Integration)[vs[1].(int)]
	}).(IntegrationOutput)
}

type IntegrationMapOutput struct{ *pulumi.OutputState }

func (IntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Integration)(nil))
}

func (o IntegrationMapOutput) ToIntegrationMapOutput() IntegrationMapOutput {
	return o
}

func (o IntegrationMapOutput) ToIntegrationMapOutputWithContext(ctx context.Context) IntegrationMapOutput {
	return o
}

func (o IntegrationMapOutput) MapIndex(k pulumi.StringInput) IntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Integration {
		return vs[0].(map[string]Integration)[vs[1].(string)]
	}).(IntegrationOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationOutput{})
	pulumi.RegisterOutputType(IntegrationPtrOutput{})
	pulumi.RegisterOutputType(IntegrationArrayOutput{})
	pulumi.RegisterOutputType(IntegrationMapOutput{})
}
