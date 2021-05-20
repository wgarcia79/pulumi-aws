// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an API Gateway API Key.
//
// > **NOTE:** Since the API Gateway usage plans feature was launched on August 11, 2016, usage plans are now **required** to associate an API key with an API stage.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigateway.NewApiKey(ctx, "myDemoApiKey", nil)
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
// API Gateway Keys can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:apigateway/apiKey:ApiKey my_demo_key 8bklk8bl1k3sB38D9B3l0enyWT8c09B30lkq0blk
// ```
type ApiKey struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The creation date of the API key
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The API key description. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`
	// Specifies whether the API key can be used by callers. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The last update date of the API key
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// The name of the API key
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The value of the API key. If not specified, it will be automatically generated by AWS on creation.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewApiKey registers a new resource with the given unique name, arguments, and options.
func NewApiKey(ctx *pulumi.Context,
	name string, args *ApiKeyArgs, opts ...pulumi.ResourceOption) (*ApiKey, error) {
	if args == nil {
		args = &ApiKeyArgs{}
	}

	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource ApiKey
	err := ctx.RegisterResource("aws:apigateway/apiKey:ApiKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiKey gets an existing ApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiKeyState, opts ...pulumi.ResourceOption) (*ApiKey, error) {
	var resource ApiKey
	err := ctx.ReadResource("aws:apigateway/apiKey:ApiKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiKey resources.
type apiKeyState struct {
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// The creation date of the API key
	CreatedDate *string `pulumi:"createdDate"`
	// The API key description. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// Specifies whether the API key can be used by callers. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The last update date of the API key
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// The name of the API key
	Name *string `pulumi:"name"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The value of the API key. If not specified, it will be automatically generated by AWS on creation.
	Value *string `pulumi:"value"`
}

type ApiKeyState struct {
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// The creation date of the API key
	CreatedDate pulumi.StringPtrInput
	// The API key description. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// Specifies whether the API key can be used by callers. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The last update date of the API key
	LastUpdatedDate pulumi.StringPtrInput
	// The name of the API key
	Name pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The value of the API key. If not specified, it will be automatically generated by AWS on creation.
	Value pulumi.StringPtrInput
}

func (ApiKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiKeyState)(nil)).Elem()
}

type apiKeyArgs struct {
	// The API key description. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// Specifies whether the API key can be used by callers. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The name of the API key
	Name *string `pulumi:"name"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The value of the API key. If not specified, it will be automatically generated by AWS on creation.
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a ApiKey resource.
type ApiKeyArgs struct {
	// The API key description. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// Specifies whether the API key can be used by callers. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The name of the API key
	Name pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The value of the API key. If not specified, it will be automatically generated by AWS on creation.
	Value pulumi.StringPtrInput
}

func (ApiKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiKeyArgs)(nil)).Elem()
}

type ApiKeyInput interface {
	pulumi.Input

	ToApiKeyOutput() ApiKeyOutput
	ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput
}

func (*ApiKey) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiKey)(nil))
}

func (i *ApiKey) ToApiKeyOutput() ApiKeyOutput {
	return i.ToApiKeyOutputWithContext(context.Background())
}

func (i *ApiKey) ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyOutput)
}

func (i *ApiKey) ToApiKeyPtrOutput() ApiKeyPtrOutput {
	return i.ToApiKeyPtrOutputWithContext(context.Background())
}

func (i *ApiKey) ToApiKeyPtrOutputWithContext(ctx context.Context) ApiKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyPtrOutput)
}

type ApiKeyPtrInput interface {
	pulumi.Input

	ToApiKeyPtrOutput() ApiKeyPtrOutput
	ToApiKeyPtrOutputWithContext(ctx context.Context) ApiKeyPtrOutput
}

type apiKeyPtrType ApiKeyArgs

func (*apiKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiKey)(nil))
}

func (i *apiKeyPtrType) ToApiKeyPtrOutput() ApiKeyPtrOutput {
	return i.ToApiKeyPtrOutputWithContext(context.Background())
}

func (i *apiKeyPtrType) ToApiKeyPtrOutputWithContext(ctx context.Context) ApiKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyPtrOutput)
}

// ApiKeyArrayInput is an input type that accepts ApiKeyArray and ApiKeyArrayOutput values.
// You can construct a concrete instance of `ApiKeyArrayInput` via:
//
//          ApiKeyArray{ ApiKeyArgs{...} }
type ApiKeyArrayInput interface {
	pulumi.Input

	ToApiKeyArrayOutput() ApiKeyArrayOutput
	ToApiKeyArrayOutputWithContext(context.Context) ApiKeyArrayOutput
}

type ApiKeyArray []ApiKeyInput

func (ApiKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ApiKey)(nil))
}

func (i ApiKeyArray) ToApiKeyArrayOutput() ApiKeyArrayOutput {
	return i.ToApiKeyArrayOutputWithContext(context.Background())
}

func (i ApiKeyArray) ToApiKeyArrayOutputWithContext(ctx context.Context) ApiKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyArrayOutput)
}

// ApiKeyMapInput is an input type that accepts ApiKeyMap and ApiKeyMapOutput values.
// You can construct a concrete instance of `ApiKeyMapInput` via:
//
//          ApiKeyMap{ "key": ApiKeyArgs{...} }
type ApiKeyMapInput interface {
	pulumi.Input

	ToApiKeyMapOutput() ApiKeyMapOutput
	ToApiKeyMapOutputWithContext(context.Context) ApiKeyMapOutput
}

type ApiKeyMap map[string]ApiKeyInput

func (ApiKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ApiKey)(nil))
}

func (i ApiKeyMap) ToApiKeyMapOutput() ApiKeyMapOutput {
	return i.ToApiKeyMapOutputWithContext(context.Background())
}

func (i ApiKeyMap) ToApiKeyMapOutputWithContext(ctx context.Context) ApiKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyMapOutput)
}

type ApiKeyOutput struct {
	*pulumi.OutputState
}

func (ApiKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiKey)(nil))
}

func (o ApiKeyOutput) ToApiKeyOutput() ApiKeyOutput {
	return o
}

func (o ApiKeyOutput) ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput {
	return o
}

func (o ApiKeyOutput) ToApiKeyPtrOutput() ApiKeyPtrOutput {
	return o.ToApiKeyPtrOutputWithContext(context.Background())
}

func (o ApiKeyOutput) ToApiKeyPtrOutputWithContext(ctx context.Context) ApiKeyPtrOutput {
	return o.ApplyT(func(v ApiKey) *ApiKey {
		return &v
	}).(ApiKeyPtrOutput)
}

type ApiKeyPtrOutput struct {
	*pulumi.OutputState
}

func (ApiKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiKey)(nil))
}

func (o ApiKeyPtrOutput) ToApiKeyPtrOutput() ApiKeyPtrOutput {
	return o
}

func (o ApiKeyPtrOutput) ToApiKeyPtrOutputWithContext(ctx context.Context) ApiKeyPtrOutput {
	return o
}

type ApiKeyArrayOutput struct{ *pulumi.OutputState }

func (ApiKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiKey)(nil))
}

func (o ApiKeyArrayOutput) ToApiKeyArrayOutput() ApiKeyArrayOutput {
	return o
}

func (o ApiKeyArrayOutput) ToApiKeyArrayOutputWithContext(ctx context.Context) ApiKeyArrayOutput {
	return o
}

func (o ApiKeyArrayOutput) Index(i pulumi.IntInput) ApiKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiKey {
		return vs[0].([]ApiKey)[vs[1].(int)]
	}).(ApiKeyOutput)
}

type ApiKeyMapOutput struct{ *pulumi.OutputState }

func (ApiKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiKey)(nil))
}

func (o ApiKeyMapOutput) ToApiKeyMapOutput() ApiKeyMapOutput {
	return o
}

func (o ApiKeyMapOutput) ToApiKeyMapOutputWithContext(ctx context.Context) ApiKeyMapOutput {
	return o
}

func (o ApiKeyMapOutput) MapIndex(k pulumi.StringInput) ApiKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApiKey {
		return vs[0].(map[string]ApiKey)[vs[1].(string)]
	}).(ApiKeyOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiKeyOutput{})
	pulumi.RegisterOutputType(ApiKeyPtrOutput{})
	pulumi.RegisterOutputType(ApiKeyArrayOutput{})
	pulumi.RegisterOutputType(ApiKeyMapOutput{})
}
