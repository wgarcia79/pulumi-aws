// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to manage the accepter's side of a Direct Connect hosted public virtual interface.
// This resource accepts ownership of a public virtual interface created by another AWS account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/directconnect"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/providers"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := providers.Newaws(ctx, "accepter", nil)
// 		if err != nil {
// 			return err
// 		}
// 		accepterCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		creator, err := directconnect.NewHostedPublicVirtualInterface(ctx, "creator", &directconnect.HostedPublicVirtualInterfaceArgs{
// 			ConnectionId:    pulumi.String("dxcon-zzzzzzzz"),
// 			OwnerAccountId:  pulumi.String(accepterCallerIdentity.AccountId),
// 			Vlan:            pulumi.Int(4094),
// 			AddressFamily:   pulumi.String("ipv4"),
// 			BgpAsn:          pulumi.Int(65352),
// 			CustomerAddress: pulumi.String("175.45.176.1/30"),
// 			AmazonAddress:   pulumi.String("175.45.176.2/30"),
// 			RouteFilterPrefixes: pulumi.StringArray{
// 				pulumi.String("210.52.109.0/24"),
// 				pulumi.String("175.45.176.0/22"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = directconnect.NewHostedPublicVirtualInterfaceAccepter(ctx, "accepterHostedPublicVirtualInterfaceAccepter", &directconnect.HostedPublicVirtualInterfaceAccepterArgs{
// 			VirtualInterfaceId: creator.ID(),
// 			Tags: pulumi.StringMap{
// 				"Side": pulumi.String("Accepter"),
// 			},
// 		}, pulumi.Provider(aws.Accepter))
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
// Direct Connect hosted public virtual interfaces can be imported using the `vif id`, e.g.
//
// ```sh
//  $ pulumi import aws:directconnect/hostedPublicVirtualInterfaceAccepter:HostedPublicVirtualInterfaceAccepter test dxvif-33cc44dd
// ```
type HostedPublicVirtualInterfaceAccepter struct {
	pulumi.CustomResourceState

	// The ARN of the virtual interface.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringOutput `pulumi:"virtualInterfaceId"`
}

// NewHostedPublicVirtualInterfaceAccepter registers a new resource with the given unique name, arguments, and options.
func NewHostedPublicVirtualInterfaceAccepter(ctx *pulumi.Context,
	name string, args *HostedPublicVirtualInterfaceAccepterArgs, opts ...pulumi.ResourceOption) (*HostedPublicVirtualInterfaceAccepter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VirtualInterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualInterfaceId'")
	}
	var resource HostedPublicVirtualInterfaceAccepter
	err := ctx.RegisterResource("aws:directconnect/hostedPublicVirtualInterfaceAccepter:HostedPublicVirtualInterfaceAccepter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostedPublicVirtualInterfaceAccepter gets an existing HostedPublicVirtualInterfaceAccepter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostedPublicVirtualInterfaceAccepter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostedPublicVirtualInterfaceAccepterState, opts ...pulumi.ResourceOption) (*HostedPublicVirtualInterfaceAccepter, error) {
	var resource HostedPublicVirtualInterfaceAccepter
	err := ctx.ReadResource("aws:directconnect/hostedPublicVirtualInterfaceAccepter:HostedPublicVirtualInterfaceAccepter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostedPublicVirtualInterfaceAccepter resources.
type hostedPublicVirtualInterfaceAccepterState struct {
	// The ARN of the virtual interface.
	Arn *string `pulumi:"arn"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId *string `pulumi:"virtualInterfaceId"`
}

type HostedPublicVirtualInterfaceAccepterState struct {
	// The ARN of the virtual interface.
	Arn pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringPtrInput
}

func (HostedPublicVirtualInterfaceAccepterState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedPublicVirtualInterfaceAccepterState)(nil)).Elem()
}

type hostedPublicVirtualInterfaceAccepterArgs struct {
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId string `pulumi:"virtualInterfaceId"`
}

// The set of arguments for constructing a HostedPublicVirtualInterfaceAccepter resource.
type HostedPublicVirtualInterfaceAccepterArgs struct {
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringInput
}

func (HostedPublicVirtualInterfaceAccepterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedPublicVirtualInterfaceAccepterArgs)(nil)).Elem()
}

type HostedPublicVirtualInterfaceAccepterInput interface {
	pulumi.Input

	ToHostedPublicVirtualInterfaceAccepterOutput() HostedPublicVirtualInterfaceAccepterOutput
	ToHostedPublicVirtualInterfaceAccepterOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceAccepterOutput
}

func (*HostedPublicVirtualInterfaceAccepter) ElementType() reflect.Type {
	return reflect.TypeOf((*HostedPublicVirtualInterfaceAccepter)(nil))
}

func (i *HostedPublicVirtualInterfaceAccepter) ToHostedPublicVirtualInterfaceAccepterOutput() HostedPublicVirtualInterfaceAccepterOutput {
	return i.ToHostedPublicVirtualInterfaceAccepterOutputWithContext(context.Background())
}

func (i *HostedPublicVirtualInterfaceAccepter) ToHostedPublicVirtualInterfaceAccepterOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceAccepterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedPublicVirtualInterfaceAccepterOutput)
}

func (i *HostedPublicVirtualInterfaceAccepter) ToHostedPublicVirtualInterfaceAccepterPtrOutput() HostedPublicVirtualInterfaceAccepterPtrOutput {
	return i.ToHostedPublicVirtualInterfaceAccepterPtrOutputWithContext(context.Background())
}

func (i *HostedPublicVirtualInterfaceAccepter) ToHostedPublicVirtualInterfaceAccepterPtrOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceAccepterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedPublicVirtualInterfaceAccepterPtrOutput)
}

type HostedPublicVirtualInterfaceAccepterPtrInput interface {
	pulumi.Input

	ToHostedPublicVirtualInterfaceAccepterPtrOutput() HostedPublicVirtualInterfaceAccepterPtrOutput
	ToHostedPublicVirtualInterfaceAccepterPtrOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceAccepterPtrOutput
}

type hostedPublicVirtualInterfaceAccepterPtrType HostedPublicVirtualInterfaceAccepterArgs

func (*hostedPublicVirtualInterfaceAccepterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HostedPublicVirtualInterfaceAccepter)(nil))
}

func (i *hostedPublicVirtualInterfaceAccepterPtrType) ToHostedPublicVirtualInterfaceAccepterPtrOutput() HostedPublicVirtualInterfaceAccepterPtrOutput {
	return i.ToHostedPublicVirtualInterfaceAccepterPtrOutputWithContext(context.Background())
}

func (i *hostedPublicVirtualInterfaceAccepterPtrType) ToHostedPublicVirtualInterfaceAccepterPtrOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceAccepterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedPublicVirtualInterfaceAccepterPtrOutput)
}

// HostedPublicVirtualInterfaceAccepterArrayInput is an input type that accepts HostedPublicVirtualInterfaceAccepterArray and HostedPublicVirtualInterfaceAccepterArrayOutput values.
// You can construct a concrete instance of `HostedPublicVirtualInterfaceAccepterArrayInput` via:
//
//          HostedPublicVirtualInterfaceAccepterArray{ HostedPublicVirtualInterfaceAccepterArgs{...} }
type HostedPublicVirtualInterfaceAccepterArrayInput interface {
	pulumi.Input

	ToHostedPublicVirtualInterfaceAccepterArrayOutput() HostedPublicVirtualInterfaceAccepterArrayOutput
	ToHostedPublicVirtualInterfaceAccepterArrayOutputWithContext(context.Context) HostedPublicVirtualInterfaceAccepterArrayOutput
}

type HostedPublicVirtualInterfaceAccepterArray []HostedPublicVirtualInterfaceAccepterInput

func (HostedPublicVirtualInterfaceAccepterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*HostedPublicVirtualInterfaceAccepter)(nil))
}

func (i HostedPublicVirtualInterfaceAccepterArray) ToHostedPublicVirtualInterfaceAccepterArrayOutput() HostedPublicVirtualInterfaceAccepterArrayOutput {
	return i.ToHostedPublicVirtualInterfaceAccepterArrayOutputWithContext(context.Background())
}

func (i HostedPublicVirtualInterfaceAccepterArray) ToHostedPublicVirtualInterfaceAccepterArrayOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceAccepterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedPublicVirtualInterfaceAccepterArrayOutput)
}

// HostedPublicVirtualInterfaceAccepterMapInput is an input type that accepts HostedPublicVirtualInterfaceAccepterMap and HostedPublicVirtualInterfaceAccepterMapOutput values.
// You can construct a concrete instance of `HostedPublicVirtualInterfaceAccepterMapInput` via:
//
//          HostedPublicVirtualInterfaceAccepterMap{ "key": HostedPublicVirtualInterfaceAccepterArgs{...} }
type HostedPublicVirtualInterfaceAccepterMapInput interface {
	pulumi.Input

	ToHostedPublicVirtualInterfaceAccepterMapOutput() HostedPublicVirtualInterfaceAccepterMapOutput
	ToHostedPublicVirtualInterfaceAccepterMapOutputWithContext(context.Context) HostedPublicVirtualInterfaceAccepterMapOutput
}

type HostedPublicVirtualInterfaceAccepterMap map[string]HostedPublicVirtualInterfaceAccepterInput

func (HostedPublicVirtualInterfaceAccepterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*HostedPublicVirtualInterfaceAccepter)(nil))
}

func (i HostedPublicVirtualInterfaceAccepterMap) ToHostedPublicVirtualInterfaceAccepterMapOutput() HostedPublicVirtualInterfaceAccepterMapOutput {
	return i.ToHostedPublicVirtualInterfaceAccepterMapOutputWithContext(context.Background())
}

func (i HostedPublicVirtualInterfaceAccepterMap) ToHostedPublicVirtualInterfaceAccepterMapOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceAccepterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedPublicVirtualInterfaceAccepterMapOutput)
}

type HostedPublicVirtualInterfaceAccepterOutput struct {
	*pulumi.OutputState
}

func (HostedPublicVirtualInterfaceAccepterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostedPublicVirtualInterfaceAccepter)(nil))
}

func (o HostedPublicVirtualInterfaceAccepterOutput) ToHostedPublicVirtualInterfaceAccepterOutput() HostedPublicVirtualInterfaceAccepterOutput {
	return o
}

func (o HostedPublicVirtualInterfaceAccepterOutput) ToHostedPublicVirtualInterfaceAccepterOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceAccepterOutput {
	return o
}

func (o HostedPublicVirtualInterfaceAccepterOutput) ToHostedPublicVirtualInterfaceAccepterPtrOutput() HostedPublicVirtualInterfaceAccepterPtrOutput {
	return o.ToHostedPublicVirtualInterfaceAccepterPtrOutputWithContext(context.Background())
}

func (o HostedPublicVirtualInterfaceAccepterOutput) ToHostedPublicVirtualInterfaceAccepterPtrOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceAccepterPtrOutput {
	return o.ApplyT(func(v HostedPublicVirtualInterfaceAccepter) *HostedPublicVirtualInterfaceAccepter {
		return &v
	}).(HostedPublicVirtualInterfaceAccepterPtrOutput)
}

type HostedPublicVirtualInterfaceAccepterPtrOutput struct {
	*pulumi.OutputState
}

func (HostedPublicVirtualInterfaceAccepterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostedPublicVirtualInterfaceAccepter)(nil))
}

func (o HostedPublicVirtualInterfaceAccepterPtrOutput) ToHostedPublicVirtualInterfaceAccepterPtrOutput() HostedPublicVirtualInterfaceAccepterPtrOutput {
	return o
}

func (o HostedPublicVirtualInterfaceAccepterPtrOutput) ToHostedPublicVirtualInterfaceAccepterPtrOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceAccepterPtrOutput {
	return o
}

type HostedPublicVirtualInterfaceAccepterArrayOutput struct{ *pulumi.OutputState }

func (HostedPublicVirtualInterfaceAccepterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostedPublicVirtualInterfaceAccepter)(nil))
}

func (o HostedPublicVirtualInterfaceAccepterArrayOutput) ToHostedPublicVirtualInterfaceAccepterArrayOutput() HostedPublicVirtualInterfaceAccepterArrayOutput {
	return o
}

func (o HostedPublicVirtualInterfaceAccepterArrayOutput) ToHostedPublicVirtualInterfaceAccepterArrayOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceAccepterArrayOutput {
	return o
}

func (o HostedPublicVirtualInterfaceAccepterArrayOutput) Index(i pulumi.IntInput) HostedPublicVirtualInterfaceAccepterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostedPublicVirtualInterfaceAccepter {
		return vs[0].([]HostedPublicVirtualInterfaceAccepter)[vs[1].(int)]
	}).(HostedPublicVirtualInterfaceAccepterOutput)
}

type HostedPublicVirtualInterfaceAccepterMapOutput struct{ *pulumi.OutputState }

func (HostedPublicVirtualInterfaceAccepterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]HostedPublicVirtualInterfaceAccepter)(nil))
}

func (o HostedPublicVirtualInterfaceAccepterMapOutput) ToHostedPublicVirtualInterfaceAccepterMapOutput() HostedPublicVirtualInterfaceAccepterMapOutput {
	return o
}

func (o HostedPublicVirtualInterfaceAccepterMapOutput) ToHostedPublicVirtualInterfaceAccepterMapOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceAccepterMapOutput {
	return o
}

func (o HostedPublicVirtualInterfaceAccepterMapOutput) MapIndex(k pulumi.StringInput) HostedPublicVirtualInterfaceAccepterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) HostedPublicVirtualInterfaceAccepter {
		return vs[0].(map[string]HostedPublicVirtualInterfaceAccepter)[vs[1].(string)]
	}).(HostedPublicVirtualInterfaceAccepterOutput)
}

func init() {
	pulumi.RegisterOutputType(HostedPublicVirtualInterfaceAccepterOutput{})
	pulumi.RegisterOutputType(HostedPublicVirtualInterfaceAccepterPtrOutput{})
	pulumi.RegisterOutputType(HostedPublicVirtualInterfaceAccepterArrayOutput{})
	pulumi.RegisterOutputType(HostedPublicVirtualInterfaceAccepterMapOutput{})
}
