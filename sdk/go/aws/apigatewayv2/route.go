// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Amazon API Gateway Version 2 route.
// More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).
//
// ## Example Usage
// ### Basic
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
// 		_, err := apigatewayv2.NewRoute(ctx, "example", &apigatewayv2.RouteArgs{
// 			ApiId:    pulumi.Any(aws_apigatewayv2_api.Example.Id),
// 			RouteKey: pulumi.String(fmt.Sprintf("%v%v", "$", "default")),
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
// `aws_apigatewayv2_route` can be imported by using the API identifier and route identifier, e.g.
//
// ```sh
//  $ pulumi import aws:apigatewayv2/route:Route example aabbccddee/1122334
// ```
type Route struct {
	pulumi.CustomResourceState

	// The API identifier.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// Boolean whether an API key is required for the route. Defaults to `false`.
	ApiKeyRequired pulumi.BoolPtrOutput `pulumi:"apiKeyRequired"`
	// The authorization scopes supported by this route. The scopes are used with a JWT authorizer to authorize the method invocation.
	AuthorizationScopes pulumi.StringArrayOutput `pulumi:"authorizationScopes"`
	// The authorization type for the route.
	// For WebSocket APIs, valid values are `NONE` for open access, `AWS_IAM` for using AWS IAM permissions, and `CUSTOM` for using a Lambda authorizer.
	// For HTTP APIs, valid values are `NONE` for open access, or `JWT` for using JSON Web Tokens.
	// Defaults to `NONE`.
	AuthorizationType pulumi.StringPtrOutput `pulumi:"authorizationType"`
	// The identifier of the `apigatewayv2.Authorizer` resource to be associated with this route, if the authorizationType is `CUSTOM`.
	AuthorizerId pulumi.StringPtrOutput `pulumi:"authorizerId"`
	// The [model selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) for the route.
	ModelSelectionExpression pulumi.StringPtrOutput `pulumi:"modelSelectionExpression"`
	// The operation name for the route. Must be between 1 and 64 characters in length.
	OperationName pulumi.StringPtrOutput `pulumi:"operationName"`
	// The request models for the route.
	RequestModels pulumi.StringMapOutput `pulumi:"requestModels"`
	// The route key for the route. For HTTP APIs, the route key can be either `$default`, or a combination of an HTTP method and resource path, for example, `GET /pets`.
	RouteKey pulumi.StringOutput `pulumi:"routeKey"`
	// The [route response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-response-selection-expressions) for the route.
	RouteResponseSelectionExpression pulumi.StringPtrOutput `pulumi:"routeResponseSelectionExpression"`
	// The target for the route. Must be between 1 and 128 characters in length.
	Target pulumi.StringPtrOutput `pulumi:"target"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.RouteKey == nil {
		return nil, errors.New("invalid value for required argument 'RouteKey'")
	}
	var resource Route
	err := ctx.RegisterResource("aws:apigatewayv2/route:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("aws:apigatewayv2/route:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
	// The API identifier.
	ApiId *string `pulumi:"apiId"`
	// Boolean whether an API key is required for the route. Defaults to `false`.
	ApiKeyRequired *bool `pulumi:"apiKeyRequired"`
	// The authorization scopes supported by this route. The scopes are used with a JWT authorizer to authorize the method invocation.
	AuthorizationScopes []string `pulumi:"authorizationScopes"`
	// The authorization type for the route.
	// For WebSocket APIs, valid values are `NONE` for open access, `AWS_IAM` for using AWS IAM permissions, and `CUSTOM` for using a Lambda authorizer.
	// For HTTP APIs, valid values are `NONE` for open access, or `JWT` for using JSON Web Tokens.
	// Defaults to `NONE`.
	AuthorizationType *string `pulumi:"authorizationType"`
	// The identifier of the `apigatewayv2.Authorizer` resource to be associated with this route, if the authorizationType is `CUSTOM`.
	AuthorizerId *string `pulumi:"authorizerId"`
	// The [model selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) for the route.
	ModelSelectionExpression *string `pulumi:"modelSelectionExpression"`
	// The operation name for the route. Must be between 1 and 64 characters in length.
	OperationName *string `pulumi:"operationName"`
	// The request models for the route.
	RequestModels map[string]string `pulumi:"requestModels"`
	// The route key for the route. For HTTP APIs, the route key can be either `$default`, or a combination of an HTTP method and resource path, for example, `GET /pets`.
	RouteKey *string `pulumi:"routeKey"`
	// The [route response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-response-selection-expressions) for the route.
	RouteResponseSelectionExpression *string `pulumi:"routeResponseSelectionExpression"`
	// The target for the route. Must be between 1 and 128 characters in length.
	Target *string `pulumi:"target"`
}

type RouteState struct {
	// The API identifier.
	ApiId pulumi.StringPtrInput
	// Boolean whether an API key is required for the route. Defaults to `false`.
	ApiKeyRequired pulumi.BoolPtrInput
	// The authorization scopes supported by this route. The scopes are used with a JWT authorizer to authorize the method invocation.
	AuthorizationScopes pulumi.StringArrayInput
	// The authorization type for the route.
	// For WebSocket APIs, valid values are `NONE` for open access, `AWS_IAM` for using AWS IAM permissions, and `CUSTOM` for using a Lambda authorizer.
	// For HTTP APIs, valid values are `NONE` for open access, or `JWT` for using JSON Web Tokens.
	// Defaults to `NONE`.
	AuthorizationType pulumi.StringPtrInput
	// The identifier of the `apigatewayv2.Authorizer` resource to be associated with this route, if the authorizationType is `CUSTOM`.
	AuthorizerId pulumi.StringPtrInput
	// The [model selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) for the route.
	ModelSelectionExpression pulumi.StringPtrInput
	// The operation name for the route. Must be between 1 and 64 characters in length.
	OperationName pulumi.StringPtrInput
	// The request models for the route.
	RequestModels pulumi.StringMapInput
	// The route key for the route. For HTTP APIs, the route key can be either `$default`, or a combination of an HTTP method and resource path, for example, `GET /pets`.
	RouteKey pulumi.StringPtrInput
	// The [route response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-response-selection-expressions) for the route.
	RouteResponseSelectionExpression pulumi.StringPtrInput
	// The target for the route. Must be between 1 and 128 characters in length.
	Target pulumi.StringPtrInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// The API identifier.
	ApiId string `pulumi:"apiId"`
	// Boolean whether an API key is required for the route. Defaults to `false`.
	ApiKeyRequired *bool `pulumi:"apiKeyRequired"`
	// The authorization scopes supported by this route. The scopes are used with a JWT authorizer to authorize the method invocation.
	AuthorizationScopes []string `pulumi:"authorizationScopes"`
	// The authorization type for the route.
	// For WebSocket APIs, valid values are `NONE` for open access, `AWS_IAM` for using AWS IAM permissions, and `CUSTOM` for using a Lambda authorizer.
	// For HTTP APIs, valid values are `NONE` for open access, or `JWT` for using JSON Web Tokens.
	// Defaults to `NONE`.
	AuthorizationType *string `pulumi:"authorizationType"`
	// The identifier of the `apigatewayv2.Authorizer` resource to be associated with this route, if the authorizationType is `CUSTOM`.
	AuthorizerId *string `pulumi:"authorizerId"`
	// The [model selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) for the route.
	ModelSelectionExpression *string `pulumi:"modelSelectionExpression"`
	// The operation name for the route. Must be between 1 and 64 characters in length.
	OperationName *string `pulumi:"operationName"`
	// The request models for the route.
	RequestModels map[string]string `pulumi:"requestModels"`
	// The route key for the route. For HTTP APIs, the route key can be either `$default`, or a combination of an HTTP method and resource path, for example, `GET /pets`.
	RouteKey string `pulumi:"routeKey"`
	// The [route response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-response-selection-expressions) for the route.
	RouteResponseSelectionExpression *string `pulumi:"routeResponseSelectionExpression"`
	// The target for the route. Must be between 1 and 128 characters in length.
	Target *string `pulumi:"target"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// The API identifier.
	ApiId pulumi.StringInput
	// Boolean whether an API key is required for the route. Defaults to `false`.
	ApiKeyRequired pulumi.BoolPtrInput
	// The authorization scopes supported by this route. The scopes are used with a JWT authorizer to authorize the method invocation.
	AuthorizationScopes pulumi.StringArrayInput
	// The authorization type for the route.
	// For WebSocket APIs, valid values are `NONE` for open access, `AWS_IAM` for using AWS IAM permissions, and `CUSTOM` for using a Lambda authorizer.
	// For HTTP APIs, valid values are `NONE` for open access, or `JWT` for using JSON Web Tokens.
	// Defaults to `NONE`.
	AuthorizationType pulumi.StringPtrInput
	// The identifier of the `apigatewayv2.Authorizer` resource to be associated with this route, if the authorizationType is `CUSTOM`.
	AuthorizerId pulumi.StringPtrInput
	// The [model selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) for the route.
	ModelSelectionExpression pulumi.StringPtrInput
	// The operation name for the route. Must be between 1 and 64 characters in length.
	OperationName pulumi.StringPtrInput
	// The request models for the route.
	RequestModels pulumi.StringMapInput
	// The route key for the route. For HTTP APIs, the route key can be either `$default`, or a combination of an HTTP method and resource path, for example, `GET /pets`.
	RouteKey pulumi.StringInput
	// The [route response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-response-selection-expressions) for the route.
	RouteResponseSelectionExpression pulumi.StringPtrInput
	// The target for the route. Must be between 1 and 128 characters in length.
	Target pulumi.StringPtrInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil))
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

func (i *Route) ToRoutePtrOutput() RoutePtrOutput {
	return i.ToRoutePtrOutputWithContext(context.Background())
}

func (i *Route) ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutePtrOutput)
}

type RoutePtrInput interface {
	pulumi.Input

	ToRoutePtrOutput() RoutePtrOutput
	ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput
}

type routePtrType RouteArgs

func (*routePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil))
}

func (i *routePtrType) ToRoutePtrOutput() RoutePtrOutput {
	return i.ToRoutePtrOutputWithContext(context.Background())
}

func (i *routePtrType) ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutePtrOutput)
}

// RouteArrayInput is an input type that accepts RouteArray and RouteArrayOutput values.
// You can construct a concrete instance of `RouteArrayInput` via:
//
//          RouteArray{ RouteArgs{...} }
type RouteArrayInput interface {
	pulumi.Input

	ToRouteArrayOutput() RouteArrayOutput
	ToRouteArrayOutputWithContext(context.Context) RouteArrayOutput
}

type RouteArray []RouteInput

func (RouteArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Route)(nil))
}

func (i RouteArray) ToRouteArrayOutput() RouteArrayOutput {
	return i.ToRouteArrayOutputWithContext(context.Background())
}

func (i RouteArray) ToRouteArrayOutputWithContext(ctx context.Context) RouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteArrayOutput)
}

// RouteMapInput is an input type that accepts RouteMap and RouteMapOutput values.
// You can construct a concrete instance of `RouteMapInput` via:
//
//          RouteMap{ "key": RouteArgs{...} }
type RouteMapInput interface {
	pulumi.Input

	ToRouteMapOutput() RouteMapOutput
	ToRouteMapOutputWithContext(context.Context) RouteMapOutput
}

type RouteMap map[string]RouteInput

func (RouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Route)(nil))
}

func (i RouteMap) ToRouteMapOutput() RouteMapOutput {
	return i.ToRouteMapOutputWithContext(context.Background())
}

func (i RouteMap) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteMapOutput)
}

type RouteOutput struct {
	*pulumi.OutputState
}

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil))
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func (o RouteOutput) ToRoutePtrOutput() RoutePtrOutput {
	return o.ToRoutePtrOutputWithContext(context.Background())
}

func (o RouteOutput) ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput {
	return o.ApplyT(func(v Route) *Route {
		return &v
	}).(RoutePtrOutput)
}

type RoutePtrOutput struct {
	*pulumi.OutputState
}

func (RoutePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil))
}

func (o RoutePtrOutput) ToRoutePtrOutput() RoutePtrOutput {
	return o
}

func (o RoutePtrOutput) ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput {
	return o
}

type RouteArrayOutput struct{ *pulumi.OutputState }

func (RouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Route)(nil))
}

func (o RouteArrayOutput) ToRouteArrayOutput() RouteArrayOutput {
	return o
}

func (o RouteArrayOutput) ToRouteArrayOutputWithContext(ctx context.Context) RouteArrayOutput {
	return o
}

func (o RouteArrayOutput) Index(i pulumi.IntInput) RouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Route {
		return vs[0].([]Route)[vs[1].(int)]
	}).(RouteOutput)
}

type RouteMapOutput struct{ *pulumi.OutputState }

func (RouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Route)(nil))
}

func (o RouteMapOutput) ToRouteMapOutput() RouteMapOutput {
	return o
}

func (o RouteMapOutput) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return o
}

func (o RouteMapOutput) MapIndex(k pulumi.StringInput) RouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Route {
		return vs[0].(map[string]Route)[vs[1].(string)]
	}).(RouteOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteOutput{})
	pulumi.RegisterOutputType(RoutePtrOutput{})
	pulumi.RegisterOutputType(RouteArrayOutput{})
	pulumi.RegisterOutputType(RouteMapOutput{})
}
