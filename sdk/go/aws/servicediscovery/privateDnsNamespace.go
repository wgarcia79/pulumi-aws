// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicediscovery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Service Discovery Private DNS Namespace resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/servicediscovery"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = servicediscovery.NewPrivateDnsNamespace(ctx, "examplePrivateDnsNamespace", &servicediscovery.PrivateDnsNamespaceArgs{
// 			Description: pulumi.String("example"),
// 			Vpc:         exampleVpc.ID(),
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
// Service Discovery Private DNS Namespace can be imported using the namespace ID and VPC ID, e.g.
//
// ```sh
//  $ pulumi import aws:servicediscovery/privateDnsNamespace:PrivateDnsNamespace example 0123456789:vpc-123345
// ```
type PrivateDnsNamespace struct {
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
	// The ID of VPC that you want to associate the namespace with.
	Vpc pulumi.StringOutput `pulumi:"vpc"`
}

// NewPrivateDnsNamespace registers a new resource with the given unique name, arguments, and options.
func NewPrivateDnsNamespace(ctx *pulumi.Context,
	name string, args *PrivateDnsNamespaceArgs, opts ...pulumi.ResourceOption) (*PrivateDnsNamespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Vpc == nil {
		return nil, errors.New("invalid value for required argument 'Vpc'")
	}
	var resource PrivateDnsNamespace
	err := ctx.RegisterResource("aws:servicediscovery/privateDnsNamespace:PrivateDnsNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateDnsNamespace gets an existing PrivateDnsNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateDnsNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateDnsNamespaceState, opts ...pulumi.ResourceOption) (*PrivateDnsNamespace, error) {
	var resource PrivateDnsNamespace
	err := ctx.ReadResource("aws:servicediscovery/privateDnsNamespace:PrivateDnsNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateDnsNamespace resources.
type privateDnsNamespaceState struct {
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
	// The ID of VPC that you want to associate the namespace with.
	Vpc *string `pulumi:"vpc"`
}

type PrivateDnsNamespaceState struct {
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
	// The ID of VPC that you want to associate the namespace with.
	Vpc pulumi.StringPtrInput
}

func (PrivateDnsNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDnsNamespaceState)(nil)).Elem()
}

type privateDnsNamespaceArgs struct {
	// The description that you specify for the namespace when you create it.
	Description *string `pulumi:"description"`
	// The name of the namespace.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the namespace. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The ID of VPC that you want to associate the namespace with.
	Vpc string `pulumi:"vpc"`
}

// The set of arguments for constructing a PrivateDnsNamespace resource.
type PrivateDnsNamespaceArgs struct {
	// The description that you specify for the namespace when you create it.
	Description pulumi.StringPtrInput
	// The name of the namespace.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the namespace. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The ID of VPC that you want to associate the namespace with.
	Vpc pulumi.StringInput
}

func (PrivateDnsNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDnsNamespaceArgs)(nil)).Elem()
}

type PrivateDnsNamespaceInput interface {
	pulumi.Input

	ToPrivateDnsNamespaceOutput() PrivateDnsNamespaceOutput
	ToPrivateDnsNamespaceOutputWithContext(ctx context.Context) PrivateDnsNamespaceOutput
}

func (*PrivateDnsNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateDnsNamespace)(nil))
}

func (i *PrivateDnsNamespace) ToPrivateDnsNamespaceOutput() PrivateDnsNamespaceOutput {
	return i.ToPrivateDnsNamespaceOutputWithContext(context.Background())
}

func (i *PrivateDnsNamespace) ToPrivateDnsNamespaceOutputWithContext(ctx context.Context) PrivateDnsNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDnsNamespaceOutput)
}

func (i *PrivateDnsNamespace) ToPrivateDnsNamespacePtrOutput() PrivateDnsNamespacePtrOutput {
	return i.ToPrivateDnsNamespacePtrOutputWithContext(context.Background())
}

func (i *PrivateDnsNamespace) ToPrivateDnsNamespacePtrOutputWithContext(ctx context.Context) PrivateDnsNamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDnsNamespacePtrOutput)
}

type PrivateDnsNamespacePtrInput interface {
	pulumi.Input

	ToPrivateDnsNamespacePtrOutput() PrivateDnsNamespacePtrOutput
	ToPrivateDnsNamespacePtrOutputWithContext(ctx context.Context) PrivateDnsNamespacePtrOutput
}

type privateDnsNamespacePtrType PrivateDnsNamespaceArgs

func (*privateDnsNamespacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateDnsNamespace)(nil))
}

func (i *privateDnsNamespacePtrType) ToPrivateDnsNamespacePtrOutput() PrivateDnsNamespacePtrOutput {
	return i.ToPrivateDnsNamespacePtrOutputWithContext(context.Background())
}

func (i *privateDnsNamespacePtrType) ToPrivateDnsNamespacePtrOutputWithContext(ctx context.Context) PrivateDnsNamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDnsNamespacePtrOutput)
}

// PrivateDnsNamespaceArrayInput is an input type that accepts PrivateDnsNamespaceArray and PrivateDnsNamespaceArrayOutput values.
// You can construct a concrete instance of `PrivateDnsNamespaceArrayInput` via:
//
//          PrivateDnsNamespaceArray{ PrivateDnsNamespaceArgs{...} }
type PrivateDnsNamespaceArrayInput interface {
	pulumi.Input

	ToPrivateDnsNamespaceArrayOutput() PrivateDnsNamespaceArrayOutput
	ToPrivateDnsNamespaceArrayOutputWithContext(context.Context) PrivateDnsNamespaceArrayOutput
}

type PrivateDnsNamespaceArray []PrivateDnsNamespaceInput

func (PrivateDnsNamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateDnsNamespace)(nil)).Elem()
}

func (i PrivateDnsNamespaceArray) ToPrivateDnsNamespaceArrayOutput() PrivateDnsNamespaceArrayOutput {
	return i.ToPrivateDnsNamespaceArrayOutputWithContext(context.Background())
}

func (i PrivateDnsNamespaceArray) ToPrivateDnsNamespaceArrayOutputWithContext(ctx context.Context) PrivateDnsNamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDnsNamespaceArrayOutput)
}

// PrivateDnsNamespaceMapInput is an input type that accepts PrivateDnsNamespaceMap and PrivateDnsNamespaceMapOutput values.
// You can construct a concrete instance of `PrivateDnsNamespaceMapInput` via:
//
//          PrivateDnsNamespaceMap{ "key": PrivateDnsNamespaceArgs{...} }
type PrivateDnsNamespaceMapInput interface {
	pulumi.Input

	ToPrivateDnsNamespaceMapOutput() PrivateDnsNamespaceMapOutput
	ToPrivateDnsNamespaceMapOutputWithContext(context.Context) PrivateDnsNamespaceMapOutput
}

type PrivateDnsNamespaceMap map[string]PrivateDnsNamespaceInput

func (PrivateDnsNamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateDnsNamespace)(nil)).Elem()
}

func (i PrivateDnsNamespaceMap) ToPrivateDnsNamespaceMapOutput() PrivateDnsNamespaceMapOutput {
	return i.ToPrivateDnsNamespaceMapOutputWithContext(context.Background())
}

func (i PrivateDnsNamespaceMap) ToPrivateDnsNamespaceMapOutputWithContext(ctx context.Context) PrivateDnsNamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDnsNamespaceMapOutput)
}

type PrivateDnsNamespaceOutput struct{ *pulumi.OutputState }

func (PrivateDnsNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateDnsNamespace)(nil))
}

func (o PrivateDnsNamespaceOutput) ToPrivateDnsNamespaceOutput() PrivateDnsNamespaceOutput {
	return o
}

func (o PrivateDnsNamespaceOutput) ToPrivateDnsNamespaceOutputWithContext(ctx context.Context) PrivateDnsNamespaceOutput {
	return o
}

func (o PrivateDnsNamespaceOutput) ToPrivateDnsNamespacePtrOutput() PrivateDnsNamespacePtrOutput {
	return o.ToPrivateDnsNamespacePtrOutputWithContext(context.Background())
}

func (o PrivateDnsNamespaceOutput) ToPrivateDnsNamespacePtrOutputWithContext(ctx context.Context) PrivateDnsNamespacePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateDnsNamespace) *PrivateDnsNamespace {
		return &v
	}).(PrivateDnsNamespacePtrOutput)
}

type PrivateDnsNamespacePtrOutput struct{ *pulumi.OutputState }

func (PrivateDnsNamespacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateDnsNamespace)(nil))
}

func (o PrivateDnsNamespacePtrOutput) ToPrivateDnsNamespacePtrOutput() PrivateDnsNamespacePtrOutput {
	return o
}

func (o PrivateDnsNamespacePtrOutput) ToPrivateDnsNamespacePtrOutputWithContext(ctx context.Context) PrivateDnsNamespacePtrOutput {
	return o
}

func (o PrivateDnsNamespacePtrOutput) Elem() PrivateDnsNamespaceOutput {
	return o.ApplyT(func(v *PrivateDnsNamespace) PrivateDnsNamespace {
		if v != nil {
			return *v
		}
		var ret PrivateDnsNamespace
		return ret
	}).(PrivateDnsNamespaceOutput)
}

type PrivateDnsNamespaceArrayOutput struct{ *pulumi.OutputState }

func (PrivateDnsNamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateDnsNamespace)(nil))
}

func (o PrivateDnsNamespaceArrayOutput) ToPrivateDnsNamespaceArrayOutput() PrivateDnsNamespaceArrayOutput {
	return o
}

func (o PrivateDnsNamespaceArrayOutput) ToPrivateDnsNamespaceArrayOutputWithContext(ctx context.Context) PrivateDnsNamespaceArrayOutput {
	return o
}

func (o PrivateDnsNamespaceArrayOutput) Index(i pulumi.IntInput) PrivateDnsNamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateDnsNamespace {
		return vs[0].([]PrivateDnsNamespace)[vs[1].(int)]
	}).(PrivateDnsNamespaceOutput)
}

type PrivateDnsNamespaceMapOutput struct{ *pulumi.OutputState }

func (PrivateDnsNamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PrivateDnsNamespace)(nil))
}

func (o PrivateDnsNamespaceMapOutput) ToPrivateDnsNamespaceMapOutput() PrivateDnsNamespaceMapOutput {
	return o
}

func (o PrivateDnsNamespaceMapOutput) ToPrivateDnsNamespaceMapOutputWithContext(ctx context.Context) PrivateDnsNamespaceMapOutput {
	return o
}

func (o PrivateDnsNamespaceMapOutput) MapIndex(k pulumi.StringInput) PrivateDnsNamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PrivateDnsNamespace {
		return vs[0].(map[string]PrivateDnsNamespace)[vs[1].(string)]
	}).(PrivateDnsNamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDnsNamespaceInput)(nil)).Elem(), &PrivateDnsNamespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDnsNamespacePtrInput)(nil)).Elem(), &PrivateDnsNamespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDnsNamespaceArrayInput)(nil)).Elem(), PrivateDnsNamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDnsNamespaceMapInput)(nil)).Elem(), PrivateDnsNamespaceMap{})
	pulumi.RegisterOutputType(PrivateDnsNamespaceOutput{})
	pulumi.RegisterOutputType(PrivateDnsNamespacePtrOutput{})
	pulumi.RegisterOutputType(PrivateDnsNamespaceArrayOutput{})
	pulumi.RegisterOutputType(PrivateDnsNamespaceMapOutput{})
}
