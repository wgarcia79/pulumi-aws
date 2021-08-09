// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicediscovery

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Service Discovery Public DNS Namespace resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/servicediscovery"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicediscovery.NewPublicDnsNamespace(ctx, "example", &servicediscovery.PublicDnsNamespaceArgs{
// 			Description: pulumi.String("example"),
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
// Service Discovery Public DNS Namespace can be imported using the namespace ID, e.g.
//
// ```sh
//  $ pulumi import aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace example 0123456789
// ```
type PublicDnsNamespace struct {
	pulumi.CustomResourceState

	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description that you specify for the namespace when you create it.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
	HostedZone pulumi.StringOutput `pulumi:"hostedZone"`
	// The name of the namespace.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of tags to assign to the namespace. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewPublicDnsNamespace registers a new resource with the given unique name, arguments, and options.
func NewPublicDnsNamespace(ctx *pulumi.Context,
	name string, args *PublicDnsNamespaceArgs, opts ...pulumi.ResourceOption) (*PublicDnsNamespace, error) {
	if args == nil {
		args = &PublicDnsNamespaceArgs{}
	}

	var resource PublicDnsNamespace
	err := ctx.RegisterResource("aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicDnsNamespace gets an existing PublicDnsNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicDnsNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicDnsNamespaceState, opts ...pulumi.ResourceOption) (*PublicDnsNamespace, error) {
	var resource PublicDnsNamespace
	err := ctx.ReadResource("aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicDnsNamespace resources.
type publicDnsNamespaceState struct {
	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn *string `pulumi:"arn"`
	// The description that you specify for the namespace when you create it.
	Description *string `pulumi:"description"`
	// The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
	HostedZone *string `pulumi:"hostedZone"`
	// The name of the namespace.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the namespace. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type PublicDnsNamespaceState struct {
	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn pulumi.StringPtrInput
	// The description that you specify for the namespace when you create it.
	Description pulumi.StringPtrInput
	// The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
	HostedZone pulumi.StringPtrInput
	// The name of the namespace.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the namespace. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (PublicDnsNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicDnsNamespaceState)(nil)).Elem()
}

type publicDnsNamespaceArgs struct {
	// The description that you specify for the namespace when you create it.
	Description *string `pulumi:"description"`
	// The name of the namespace.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the namespace. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a PublicDnsNamespace resource.
type PublicDnsNamespaceArgs struct {
	// The description that you specify for the namespace when you create it.
	Description pulumi.StringPtrInput
	// The name of the namespace.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the namespace. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (PublicDnsNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicDnsNamespaceArgs)(nil)).Elem()
}

type PublicDnsNamespaceInput interface {
	pulumi.Input

	ToPublicDnsNamespaceOutput() PublicDnsNamespaceOutput
	ToPublicDnsNamespaceOutputWithContext(ctx context.Context) PublicDnsNamespaceOutput
}

func (*PublicDnsNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicDnsNamespace)(nil))
}

func (i *PublicDnsNamespace) ToPublicDnsNamespaceOutput() PublicDnsNamespaceOutput {
	return i.ToPublicDnsNamespaceOutputWithContext(context.Background())
}

func (i *PublicDnsNamespace) ToPublicDnsNamespaceOutputWithContext(ctx context.Context) PublicDnsNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicDnsNamespaceOutput)
}

func (i *PublicDnsNamespace) ToPublicDnsNamespacePtrOutput() PublicDnsNamespacePtrOutput {
	return i.ToPublicDnsNamespacePtrOutputWithContext(context.Background())
}

func (i *PublicDnsNamespace) ToPublicDnsNamespacePtrOutputWithContext(ctx context.Context) PublicDnsNamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicDnsNamespacePtrOutput)
}

type PublicDnsNamespacePtrInput interface {
	pulumi.Input

	ToPublicDnsNamespacePtrOutput() PublicDnsNamespacePtrOutput
	ToPublicDnsNamespacePtrOutputWithContext(ctx context.Context) PublicDnsNamespacePtrOutput
}

type publicDnsNamespacePtrType PublicDnsNamespaceArgs

func (*publicDnsNamespacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicDnsNamespace)(nil))
}

func (i *publicDnsNamespacePtrType) ToPublicDnsNamespacePtrOutput() PublicDnsNamespacePtrOutput {
	return i.ToPublicDnsNamespacePtrOutputWithContext(context.Background())
}

func (i *publicDnsNamespacePtrType) ToPublicDnsNamespacePtrOutputWithContext(ctx context.Context) PublicDnsNamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicDnsNamespacePtrOutput)
}

// PublicDnsNamespaceArrayInput is an input type that accepts PublicDnsNamespaceArray and PublicDnsNamespaceArrayOutput values.
// You can construct a concrete instance of `PublicDnsNamespaceArrayInput` via:
//
//          PublicDnsNamespaceArray{ PublicDnsNamespaceArgs{...} }
type PublicDnsNamespaceArrayInput interface {
	pulumi.Input

	ToPublicDnsNamespaceArrayOutput() PublicDnsNamespaceArrayOutput
	ToPublicDnsNamespaceArrayOutputWithContext(context.Context) PublicDnsNamespaceArrayOutput
}

type PublicDnsNamespaceArray []PublicDnsNamespaceInput

func (PublicDnsNamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicDnsNamespace)(nil)).Elem()
}

func (i PublicDnsNamespaceArray) ToPublicDnsNamespaceArrayOutput() PublicDnsNamespaceArrayOutput {
	return i.ToPublicDnsNamespaceArrayOutputWithContext(context.Background())
}

func (i PublicDnsNamespaceArray) ToPublicDnsNamespaceArrayOutputWithContext(ctx context.Context) PublicDnsNamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicDnsNamespaceArrayOutput)
}

// PublicDnsNamespaceMapInput is an input type that accepts PublicDnsNamespaceMap and PublicDnsNamespaceMapOutput values.
// You can construct a concrete instance of `PublicDnsNamespaceMapInput` via:
//
//          PublicDnsNamespaceMap{ "key": PublicDnsNamespaceArgs{...} }
type PublicDnsNamespaceMapInput interface {
	pulumi.Input

	ToPublicDnsNamespaceMapOutput() PublicDnsNamespaceMapOutput
	ToPublicDnsNamespaceMapOutputWithContext(context.Context) PublicDnsNamespaceMapOutput
}

type PublicDnsNamespaceMap map[string]PublicDnsNamespaceInput

func (PublicDnsNamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicDnsNamespace)(nil)).Elem()
}

func (i PublicDnsNamespaceMap) ToPublicDnsNamespaceMapOutput() PublicDnsNamespaceMapOutput {
	return i.ToPublicDnsNamespaceMapOutputWithContext(context.Background())
}

func (i PublicDnsNamespaceMap) ToPublicDnsNamespaceMapOutputWithContext(ctx context.Context) PublicDnsNamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicDnsNamespaceMapOutput)
}

type PublicDnsNamespaceOutput struct{ *pulumi.OutputState }

func (PublicDnsNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicDnsNamespace)(nil))
}

func (o PublicDnsNamespaceOutput) ToPublicDnsNamespaceOutput() PublicDnsNamespaceOutput {
	return o
}

func (o PublicDnsNamespaceOutput) ToPublicDnsNamespaceOutputWithContext(ctx context.Context) PublicDnsNamespaceOutput {
	return o
}

func (o PublicDnsNamespaceOutput) ToPublicDnsNamespacePtrOutput() PublicDnsNamespacePtrOutput {
	return o.ToPublicDnsNamespacePtrOutputWithContext(context.Background())
}

func (o PublicDnsNamespaceOutput) ToPublicDnsNamespacePtrOutputWithContext(ctx context.Context) PublicDnsNamespacePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicDnsNamespace) *PublicDnsNamespace {
		return &v
	}).(PublicDnsNamespacePtrOutput)
}

type PublicDnsNamespacePtrOutput struct{ *pulumi.OutputState }

func (PublicDnsNamespacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicDnsNamespace)(nil))
}

func (o PublicDnsNamespacePtrOutput) ToPublicDnsNamespacePtrOutput() PublicDnsNamespacePtrOutput {
	return o
}

func (o PublicDnsNamespacePtrOutput) ToPublicDnsNamespacePtrOutputWithContext(ctx context.Context) PublicDnsNamespacePtrOutput {
	return o
}

func (o PublicDnsNamespacePtrOutput) Elem() PublicDnsNamespaceOutput {
	return o.ApplyT(func(v *PublicDnsNamespace) PublicDnsNamespace {
		if v != nil {
			return *v
		}
		var ret PublicDnsNamespace
		return ret
	}).(PublicDnsNamespaceOutput)
}

type PublicDnsNamespaceArrayOutput struct{ *pulumi.OutputState }

func (PublicDnsNamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PublicDnsNamespace)(nil))
}

func (o PublicDnsNamespaceArrayOutput) ToPublicDnsNamespaceArrayOutput() PublicDnsNamespaceArrayOutput {
	return o
}

func (o PublicDnsNamespaceArrayOutput) ToPublicDnsNamespaceArrayOutputWithContext(ctx context.Context) PublicDnsNamespaceArrayOutput {
	return o
}

func (o PublicDnsNamespaceArrayOutput) Index(i pulumi.IntInput) PublicDnsNamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PublicDnsNamespace {
		return vs[0].([]PublicDnsNamespace)[vs[1].(int)]
	}).(PublicDnsNamespaceOutput)
}

type PublicDnsNamespaceMapOutput struct{ *pulumi.OutputState }

func (PublicDnsNamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PublicDnsNamespace)(nil))
}

func (o PublicDnsNamespaceMapOutput) ToPublicDnsNamespaceMapOutput() PublicDnsNamespaceMapOutput {
	return o
}

func (o PublicDnsNamespaceMapOutput) ToPublicDnsNamespaceMapOutputWithContext(ctx context.Context) PublicDnsNamespaceMapOutput {
	return o
}

func (o PublicDnsNamespaceMapOutput) MapIndex(k pulumi.StringInput) PublicDnsNamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PublicDnsNamespace {
		return vs[0].(map[string]PublicDnsNamespace)[vs[1].(string)]
	}).(PublicDnsNamespaceOutput)
}

func init() {
	pulumi.RegisterOutputType(PublicDnsNamespaceOutput{})
	pulumi.RegisterOutputType(PublicDnsNamespacePtrOutput{})
	pulumi.RegisterOutputType(PublicDnsNamespaceArrayOutput{})
	pulumi.RegisterOutputType(PublicDnsNamespaceMapOutput{})
}
