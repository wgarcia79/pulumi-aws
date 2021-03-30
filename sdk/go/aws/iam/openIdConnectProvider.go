// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an IAM OpenID Connect provider.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iam.NewOpenIdConnectProvider(ctx, "_default", &iam.OpenIdConnectProviderArgs{
// 			ClientIdLists: pulumi.StringArray{
// 				pulumi.String("266362248691-342342xasdasdasda-apps.googleusercontent.com"),
// 			},
// 			ThumbprintLists: []interface{}{},
// 			Url:             pulumi.String("https://accounts.google.com"),
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
// IAM OpenID Connect Providers can be imported using the `arn`, e.g.
//
// ```sh
//  $ pulumi import aws:iam/openIdConnectProvider:OpenIdConnectProvider default arn:aws:iam::123456789012:oidc-provider/accounts.google.com
// ```
type OpenIdConnectProvider struct {
	pulumi.CustomResourceState

	// The ARN assigned by AWS for this provider.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
	ClientIdLists pulumi.StringArrayOutput `pulumi:"clientIdLists"`
	// Map of resource tags for the IAM OIDC provider.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
	ThumbprintLists pulumi.StringArrayOutput `pulumi:"thumbprintLists"`
	// The URL of the identity provider. Corresponds to the _iss_ claim.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewOpenIdConnectProvider registers a new resource with the given unique name, arguments, and options.
func NewOpenIdConnectProvider(ctx *pulumi.Context,
	name string, args *OpenIdConnectProviderArgs, opts ...pulumi.ResourceOption) (*OpenIdConnectProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientIdLists == nil {
		return nil, errors.New("invalid value for required argument 'ClientIdLists'")
	}
	if args.ThumbprintLists == nil {
		return nil, errors.New("invalid value for required argument 'ThumbprintLists'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	var resource OpenIdConnectProvider
	err := ctx.RegisterResource("aws:iam/openIdConnectProvider:OpenIdConnectProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOpenIdConnectProvider gets an existing OpenIdConnectProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOpenIdConnectProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OpenIdConnectProviderState, opts ...pulumi.ResourceOption) (*OpenIdConnectProvider, error) {
	var resource OpenIdConnectProvider
	err := ctx.ReadResource("aws:iam/openIdConnectProvider:OpenIdConnectProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OpenIdConnectProvider resources.
type openIdConnectProviderState struct {
	// The ARN assigned by AWS for this provider.
	Arn *string `pulumi:"arn"`
	// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
	ClientIdLists []string `pulumi:"clientIdLists"`
	// Map of resource tags for the IAM OIDC provider.
	Tags map[string]string `pulumi:"tags"`
	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
	ThumbprintLists []string `pulumi:"thumbprintLists"`
	// The URL of the identity provider. Corresponds to the _iss_ claim.
	Url *string `pulumi:"url"`
}

type OpenIdConnectProviderState struct {
	// The ARN assigned by AWS for this provider.
	Arn pulumi.StringPtrInput
	// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
	ClientIdLists pulumi.StringArrayInput
	// Map of resource tags for the IAM OIDC provider.
	Tags pulumi.StringMapInput
	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
	ThumbprintLists pulumi.StringArrayInput
	// The URL of the identity provider. Corresponds to the _iss_ claim.
	Url pulumi.StringPtrInput
}

func (OpenIdConnectProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*openIdConnectProviderState)(nil)).Elem()
}

type openIdConnectProviderArgs struct {
	// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
	ClientIdLists []string `pulumi:"clientIdLists"`
	// Map of resource tags for the IAM OIDC provider.
	Tags map[string]string `pulumi:"tags"`
	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
	ThumbprintLists []string `pulumi:"thumbprintLists"`
	// The URL of the identity provider. Corresponds to the _iss_ claim.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a OpenIdConnectProvider resource.
type OpenIdConnectProviderArgs struct {
	// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
	ClientIdLists pulumi.StringArrayInput
	// Map of resource tags for the IAM OIDC provider.
	Tags pulumi.StringMapInput
	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
	ThumbprintLists pulumi.StringArrayInput
	// The URL of the identity provider. Corresponds to the _iss_ claim.
	Url pulumi.StringInput
}

func (OpenIdConnectProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*openIdConnectProviderArgs)(nil)).Elem()
}

type OpenIdConnectProviderInput interface {
	pulumi.Input

	ToOpenIdConnectProviderOutput() OpenIdConnectProviderOutput
	ToOpenIdConnectProviderOutputWithContext(ctx context.Context) OpenIdConnectProviderOutput
}

func (*OpenIdConnectProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectProvider)(nil))
}

func (i *OpenIdConnectProvider) ToOpenIdConnectProviderOutput() OpenIdConnectProviderOutput {
	return i.ToOpenIdConnectProviderOutputWithContext(context.Background())
}

func (i *OpenIdConnectProvider) ToOpenIdConnectProviderOutputWithContext(ctx context.Context) OpenIdConnectProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectProviderOutput)
}

func (i *OpenIdConnectProvider) ToOpenIdConnectProviderPtrOutput() OpenIdConnectProviderPtrOutput {
	return i.ToOpenIdConnectProviderPtrOutputWithContext(context.Background())
}

func (i *OpenIdConnectProvider) ToOpenIdConnectProviderPtrOutputWithContext(ctx context.Context) OpenIdConnectProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectProviderPtrOutput)
}

type OpenIdConnectProviderPtrInput interface {
	pulumi.Input

	ToOpenIdConnectProviderPtrOutput() OpenIdConnectProviderPtrOutput
	ToOpenIdConnectProviderPtrOutputWithContext(ctx context.Context) OpenIdConnectProviderPtrOutput
}

type openIdConnectProviderPtrType OpenIdConnectProviderArgs

func (*openIdConnectProviderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectProvider)(nil))
}

func (i *openIdConnectProviderPtrType) ToOpenIdConnectProviderPtrOutput() OpenIdConnectProviderPtrOutput {
	return i.ToOpenIdConnectProviderPtrOutputWithContext(context.Background())
}

func (i *openIdConnectProviderPtrType) ToOpenIdConnectProviderPtrOutputWithContext(ctx context.Context) OpenIdConnectProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectProviderPtrOutput)
}

// OpenIdConnectProviderArrayInput is an input type that accepts OpenIdConnectProviderArray and OpenIdConnectProviderArrayOutput values.
// You can construct a concrete instance of `OpenIdConnectProviderArrayInput` via:
//
//          OpenIdConnectProviderArray{ OpenIdConnectProviderArgs{...} }
type OpenIdConnectProviderArrayInput interface {
	pulumi.Input

	ToOpenIdConnectProviderArrayOutput() OpenIdConnectProviderArrayOutput
	ToOpenIdConnectProviderArrayOutputWithContext(context.Context) OpenIdConnectProviderArrayOutput
}

type OpenIdConnectProviderArray []OpenIdConnectProviderInput

func (OpenIdConnectProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*OpenIdConnectProvider)(nil))
}

func (i OpenIdConnectProviderArray) ToOpenIdConnectProviderArrayOutput() OpenIdConnectProviderArrayOutput {
	return i.ToOpenIdConnectProviderArrayOutputWithContext(context.Background())
}

func (i OpenIdConnectProviderArray) ToOpenIdConnectProviderArrayOutputWithContext(ctx context.Context) OpenIdConnectProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectProviderArrayOutput)
}

// OpenIdConnectProviderMapInput is an input type that accepts OpenIdConnectProviderMap and OpenIdConnectProviderMapOutput values.
// You can construct a concrete instance of `OpenIdConnectProviderMapInput` via:
//
//          OpenIdConnectProviderMap{ "key": OpenIdConnectProviderArgs{...} }
type OpenIdConnectProviderMapInput interface {
	pulumi.Input

	ToOpenIdConnectProviderMapOutput() OpenIdConnectProviderMapOutput
	ToOpenIdConnectProviderMapOutputWithContext(context.Context) OpenIdConnectProviderMapOutput
}

type OpenIdConnectProviderMap map[string]OpenIdConnectProviderInput

func (OpenIdConnectProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*OpenIdConnectProvider)(nil))
}

func (i OpenIdConnectProviderMap) ToOpenIdConnectProviderMapOutput() OpenIdConnectProviderMapOutput {
	return i.ToOpenIdConnectProviderMapOutputWithContext(context.Background())
}

func (i OpenIdConnectProviderMap) ToOpenIdConnectProviderMapOutputWithContext(ctx context.Context) OpenIdConnectProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectProviderMapOutput)
}

type OpenIdConnectProviderOutput struct {
	*pulumi.OutputState
}

func (OpenIdConnectProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectProvider)(nil))
}

func (o OpenIdConnectProviderOutput) ToOpenIdConnectProviderOutput() OpenIdConnectProviderOutput {
	return o
}

func (o OpenIdConnectProviderOutput) ToOpenIdConnectProviderOutputWithContext(ctx context.Context) OpenIdConnectProviderOutput {
	return o
}

func (o OpenIdConnectProviderOutput) ToOpenIdConnectProviderPtrOutput() OpenIdConnectProviderPtrOutput {
	return o.ToOpenIdConnectProviderPtrOutputWithContext(context.Background())
}

func (o OpenIdConnectProviderOutput) ToOpenIdConnectProviderPtrOutputWithContext(ctx context.Context) OpenIdConnectProviderPtrOutput {
	return o.ApplyT(func(v OpenIdConnectProvider) *OpenIdConnectProvider {
		return &v
	}).(OpenIdConnectProviderPtrOutput)
}

type OpenIdConnectProviderPtrOutput struct {
	*pulumi.OutputState
}

func (OpenIdConnectProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectProvider)(nil))
}

func (o OpenIdConnectProviderPtrOutput) ToOpenIdConnectProviderPtrOutput() OpenIdConnectProviderPtrOutput {
	return o
}

func (o OpenIdConnectProviderPtrOutput) ToOpenIdConnectProviderPtrOutputWithContext(ctx context.Context) OpenIdConnectProviderPtrOutput {
	return o
}

type OpenIdConnectProviderArrayOutput struct{ *pulumi.OutputState }

func (OpenIdConnectProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OpenIdConnectProvider)(nil))
}

func (o OpenIdConnectProviderArrayOutput) ToOpenIdConnectProviderArrayOutput() OpenIdConnectProviderArrayOutput {
	return o
}

func (o OpenIdConnectProviderArrayOutput) ToOpenIdConnectProviderArrayOutputWithContext(ctx context.Context) OpenIdConnectProviderArrayOutput {
	return o
}

func (o OpenIdConnectProviderArrayOutput) Index(i pulumi.IntInput) OpenIdConnectProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OpenIdConnectProvider {
		return vs[0].([]OpenIdConnectProvider)[vs[1].(int)]
	}).(OpenIdConnectProviderOutput)
}

type OpenIdConnectProviderMapOutput struct{ *pulumi.OutputState }

func (OpenIdConnectProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OpenIdConnectProvider)(nil))
}

func (o OpenIdConnectProviderMapOutput) ToOpenIdConnectProviderMapOutput() OpenIdConnectProviderMapOutput {
	return o
}

func (o OpenIdConnectProviderMapOutput) ToOpenIdConnectProviderMapOutputWithContext(ctx context.Context) OpenIdConnectProviderMapOutput {
	return o
}

func (o OpenIdConnectProviderMapOutput) MapIndex(k pulumi.StringInput) OpenIdConnectProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OpenIdConnectProvider {
		return vs[0].(map[string]OpenIdConnectProvider)[vs[1].(string)]
	}).(OpenIdConnectProviderOutput)
}

func init() {
	pulumi.RegisterOutputType(OpenIdConnectProviderOutput{})
	pulumi.RegisterOutputType(OpenIdConnectProviderPtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectProviderArrayOutput{})
	pulumi.RegisterOutputType(OpenIdConnectProviderMapOutput{})
}
