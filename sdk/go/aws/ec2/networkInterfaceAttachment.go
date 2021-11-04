// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attach an Elastic network interface (ENI) resource with EC2 instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.NewNetworkInterfaceAttachment(ctx, "test", &ec2.NetworkInterfaceAttachmentArgs{
// 			InstanceId:         pulumi.Any(aws_instance.Test.Id),
// 			NetworkInterfaceId: pulumi.Any(aws_network_interface.Test.Id),
// 			DeviceIndex:        pulumi.Int(0),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type NetworkInterfaceAttachment struct {
	pulumi.CustomResourceState

	// The ENI Attachment ID.
	AttachmentId pulumi.StringOutput `pulumi:"attachmentId"`
	// Network interface index (int).
	DeviceIndex pulumi.IntOutput `pulumi:"deviceIndex"`
	// Instance ID to attach.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// ENI ID to attach.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The status of the Network Interface Attachment.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewNetworkInterfaceAttachment registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterfaceAttachment(ctx *pulumi.Context,
	name string, args *NetworkInterfaceAttachmentArgs, opts ...pulumi.ResourceOption) (*NetworkInterfaceAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceIndex == nil {
		return nil, errors.New("invalid value for required argument 'DeviceIndex'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.NetworkInterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInterfaceId'")
	}
	var resource NetworkInterfaceAttachment
	err := ctx.RegisterResource("aws:ec2/networkInterfaceAttachment:NetworkInterfaceAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterfaceAttachment gets an existing NetworkInterfaceAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterfaceAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceAttachmentState, opts ...pulumi.ResourceOption) (*NetworkInterfaceAttachment, error) {
	var resource NetworkInterfaceAttachment
	err := ctx.ReadResource("aws:ec2/networkInterfaceAttachment:NetworkInterfaceAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterfaceAttachment resources.
type networkInterfaceAttachmentState struct {
	// The ENI Attachment ID.
	AttachmentId *string `pulumi:"attachmentId"`
	// Network interface index (int).
	DeviceIndex *int `pulumi:"deviceIndex"`
	// Instance ID to attach.
	InstanceId *string `pulumi:"instanceId"`
	// ENI ID to attach.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The status of the Network Interface Attachment.
	Status *string `pulumi:"status"`
}

type NetworkInterfaceAttachmentState struct {
	// The ENI Attachment ID.
	AttachmentId pulumi.StringPtrInput
	// Network interface index (int).
	DeviceIndex pulumi.IntPtrInput
	// Instance ID to attach.
	InstanceId pulumi.StringPtrInput
	// ENI ID to attach.
	NetworkInterfaceId pulumi.StringPtrInput
	// The status of the Network Interface Attachment.
	Status pulumi.StringPtrInput
}

func (NetworkInterfaceAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceAttachmentState)(nil)).Elem()
}

type networkInterfaceAttachmentArgs struct {
	// Network interface index (int).
	DeviceIndex int `pulumi:"deviceIndex"`
	// Instance ID to attach.
	InstanceId string `pulumi:"instanceId"`
	// ENI ID to attach.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
}

// The set of arguments for constructing a NetworkInterfaceAttachment resource.
type NetworkInterfaceAttachmentArgs struct {
	// Network interface index (int).
	DeviceIndex pulumi.IntInput
	// Instance ID to attach.
	InstanceId pulumi.StringInput
	// ENI ID to attach.
	NetworkInterfaceId pulumi.StringInput
}

func (NetworkInterfaceAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceAttachmentArgs)(nil)).Elem()
}

type NetworkInterfaceAttachmentInput interface {
	pulumi.Input

	ToNetworkInterfaceAttachmentOutput() NetworkInterfaceAttachmentOutput
	ToNetworkInterfaceAttachmentOutputWithContext(ctx context.Context) NetworkInterfaceAttachmentOutput
}

func (*NetworkInterfaceAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceAttachment)(nil))
}

func (i *NetworkInterfaceAttachment) ToNetworkInterfaceAttachmentOutput() NetworkInterfaceAttachmentOutput {
	return i.ToNetworkInterfaceAttachmentOutputWithContext(context.Background())
}

func (i *NetworkInterfaceAttachment) ToNetworkInterfaceAttachmentOutputWithContext(ctx context.Context) NetworkInterfaceAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceAttachmentOutput)
}

func (i *NetworkInterfaceAttachment) ToNetworkInterfaceAttachmentPtrOutput() NetworkInterfaceAttachmentPtrOutput {
	return i.ToNetworkInterfaceAttachmentPtrOutputWithContext(context.Background())
}

func (i *NetworkInterfaceAttachment) ToNetworkInterfaceAttachmentPtrOutputWithContext(ctx context.Context) NetworkInterfaceAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceAttachmentPtrOutput)
}

type NetworkInterfaceAttachmentPtrInput interface {
	pulumi.Input

	ToNetworkInterfaceAttachmentPtrOutput() NetworkInterfaceAttachmentPtrOutput
	ToNetworkInterfaceAttachmentPtrOutputWithContext(ctx context.Context) NetworkInterfaceAttachmentPtrOutput
}

type networkInterfaceAttachmentPtrType NetworkInterfaceAttachmentArgs

func (*networkInterfaceAttachmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceAttachment)(nil))
}

func (i *networkInterfaceAttachmentPtrType) ToNetworkInterfaceAttachmentPtrOutput() NetworkInterfaceAttachmentPtrOutput {
	return i.ToNetworkInterfaceAttachmentPtrOutputWithContext(context.Background())
}

func (i *networkInterfaceAttachmentPtrType) ToNetworkInterfaceAttachmentPtrOutputWithContext(ctx context.Context) NetworkInterfaceAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceAttachmentPtrOutput)
}

// NetworkInterfaceAttachmentArrayInput is an input type that accepts NetworkInterfaceAttachmentArray and NetworkInterfaceAttachmentArrayOutput values.
// You can construct a concrete instance of `NetworkInterfaceAttachmentArrayInput` via:
//
//          NetworkInterfaceAttachmentArray{ NetworkInterfaceAttachmentArgs{...} }
type NetworkInterfaceAttachmentArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceAttachmentArrayOutput() NetworkInterfaceAttachmentArrayOutput
	ToNetworkInterfaceAttachmentArrayOutputWithContext(context.Context) NetworkInterfaceAttachmentArrayOutput
}

type NetworkInterfaceAttachmentArray []NetworkInterfaceAttachmentInput

func (NetworkInterfaceAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkInterfaceAttachment)(nil)).Elem()
}

func (i NetworkInterfaceAttachmentArray) ToNetworkInterfaceAttachmentArrayOutput() NetworkInterfaceAttachmentArrayOutput {
	return i.ToNetworkInterfaceAttachmentArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceAttachmentArray) ToNetworkInterfaceAttachmentArrayOutputWithContext(ctx context.Context) NetworkInterfaceAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceAttachmentArrayOutput)
}

// NetworkInterfaceAttachmentMapInput is an input type that accepts NetworkInterfaceAttachmentMap and NetworkInterfaceAttachmentMapOutput values.
// You can construct a concrete instance of `NetworkInterfaceAttachmentMapInput` via:
//
//          NetworkInterfaceAttachmentMap{ "key": NetworkInterfaceAttachmentArgs{...} }
type NetworkInterfaceAttachmentMapInput interface {
	pulumi.Input

	ToNetworkInterfaceAttachmentMapOutput() NetworkInterfaceAttachmentMapOutput
	ToNetworkInterfaceAttachmentMapOutputWithContext(context.Context) NetworkInterfaceAttachmentMapOutput
}

type NetworkInterfaceAttachmentMap map[string]NetworkInterfaceAttachmentInput

func (NetworkInterfaceAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkInterfaceAttachment)(nil)).Elem()
}

func (i NetworkInterfaceAttachmentMap) ToNetworkInterfaceAttachmentMapOutput() NetworkInterfaceAttachmentMapOutput {
	return i.ToNetworkInterfaceAttachmentMapOutputWithContext(context.Background())
}

func (i NetworkInterfaceAttachmentMap) ToNetworkInterfaceAttachmentMapOutputWithContext(ctx context.Context) NetworkInterfaceAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceAttachmentMapOutput)
}

type NetworkInterfaceAttachmentOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceAttachment)(nil))
}

func (o NetworkInterfaceAttachmentOutput) ToNetworkInterfaceAttachmentOutput() NetworkInterfaceAttachmentOutput {
	return o
}

func (o NetworkInterfaceAttachmentOutput) ToNetworkInterfaceAttachmentOutputWithContext(ctx context.Context) NetworkInterfaceAttachmentOutput {
	return o
}

func (o NetworkInterfaceAttachmentOutput) ToNetworkInterfaceAttachmentPtrOutput() NetworkInterfaceAttachmentPtrOutput {
	return o.ToNetworkInterfaceAttachmentPtrOutputWithContext(context.Background())
}

func (o NetworkInterfaceAttachmentOutput) ToNetworkInterfaceAttachmentPtrOutputWithContext(ctx context.Context) NetworkInterfaceAttachmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkInterfaceAttachment) *NetworkInterfaceAttachment {
		return &v
	}).(NetworkInterfaceAttachmentPtrOutput)
}

type NetworkInterfaceAttachmentPtrOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceAttachment)(nil))
}

func (o NetworkInterfaceAttachmentPtrOutput) ToNetworkInterfaceAttachmentPtrOutput() NetworkInterfaceAttachmentPtrOutput {
	return o
}

func (o NetworkInterfaceAttachmentPtrOutput) ToNetworkInterfaceAttachmentPtrOutputWithContext(ctx context.Context) NetworkInterfaceAttachmentPtrOutput {
	return o
}

func (o NetworkInterfaceAttachmentPtrOutput) Elem() NetworkInterfaceAttachmentOutput {
	return o.ApplyT(func(v *NetworkInterfaceAttachment) NetworkInterfaceAttachment {
		if v != nil {
			return *v
		}
		var ret NetworkInterfaceAttachment
		return ret
	}).(NetworkInterfaceAttachmentOutput)
}

type NetworkInterfaceAttachmentArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceAttachment)(nil))
}

func (o NetworkInterfaceAttachmentArrayOutput) ToNetworkInterfaceAttachmentArrayOutput() NetworkInterfaceAttachmentArrayOutput {
	return o
}

func (o NetworkInterfaceAttachmentArrayOutput) ToNetworkInterfaceAttachmentArrayOutputWithContext(ctx context.Context) NetworkInterfaceAttachmentArrayOutput {
	return o
}

func (o NetworkInterfaceAttachmentArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceAttachment {
		return vs[0].([]NetworkInterfaceAttachment)[vs[1].(int)]
	}).(NetworkInterfaceAttachmentOutput)
}

type NetworkInterfaceAttachmentMapOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkInterfaceAttachment)(nil))
}

func (o NetworkInterfaceAttachmentMapOutput) ToNetworkInterfaceAttachmentMapOutput() NetworkInterfaceAttachmentMapOutput {
	return o
}

func (o NetworkInterfaceAttachmentMapOutput) ToNetworkInterfaceAttachmentMapOutputWithContext(ctx context.Context) NetworkInterfaceAttachmentMapOutput {
	return o
}

func (o NetworkInterfaceAttachmentMapOutput) MapIndex(k pulumi.StringInput) NetworkInterfaceAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkInterfaceAttachment {
		return vs[0].(map[string]NetworkInterfaceAttachment)[vs[1].(string)]
	}).(NetworkInterfaceAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceAttachmentInput)(nil)).Elem(), &NetworkInterfaceAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceAttachmentPtrInput)(nil)).Elem(), &NetworkInterfaceAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceAttachmentArrayInput)(nil)).Elem(), NetworkInterfaceAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceAttachmentMapInput)(nil)).Elem(), NetworkInterfaceAttachmentMap{})
	pulumi.RegisterOutputType(NetworkInterfaceAttachmentOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceAttachmentPtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceAttachmentArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceAttachmentMapOutput{})
}
