// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an AWS Storage Gateway upload buffer.
//
// > **NOTE:** The Storage Gateway API provides no method to remove an upload buffer disk. Destroying this resource does not perform any Storage Gateway actions.
//
// ## Example Usage
// ### Cached and VTL Gateway Type
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := aws_volume_attachment.Test.Device_name
// 		testLocalDisk, err := storagegateway.GetLocalDisk(ctx, &storagegateway.GetLocalDiskArgs{
// 			DiskNode:   &opt0,
// 			GatewayArn: aws_storagegateway_gateway.Test.Arn,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storagegateway.NewUploadBuffer(ctx, "testUploadBuffer", &storagegateway.UploadBufferArgs{
// 			DiskPath:   pulumi.String(testLocalDisk.DiskPath),
// 			GatewayArn: pulumi.Any(aws_storagegateway_gateway.Test.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Stored Gateway Type
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := aws_volume_attachment.Test.Device_name
// 		_, err := storagegateway.GetLocalDisk(ctx, &storagegateway.GetLocalDiskArgs{
// 			DiskNode:   &opt0,
// 			GatewayArn: aws_storagegateway_gateway.Test.Arn,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storagegateway.NewUploadBuffer(ctx, "example", &storagegateway.UploadBufferArgs{
// 			DiskId:     pulumi.Any(data.Aws_storagegateway_local_disk.Example.Id),
// 			GatewayArn: pulumi.Any(aws_storagegateway_gateway.Example.Arn),
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
// `aws_storagegateway_upload_buffer` can be imported by using the gateway Amazon Resource Name (ARN) and local disk identifier separated with a colon (`:`), e.g.,
//
// ```sh
//  $ pulumi import aws:storagegateway/uploadBuffer:UploadBuffer example arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678:pci-0000:03:00.0-scsi-0:0:0:0
// ```
type UploadBuffer struct {
	pulumi.CustomResourceState

	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId pulumi.StringOutput `pulumi:"diskId"`
	// Local disk path. For example, `/dev/nvme1n1`.
	DiskPath pulumi.StringOutput `pulumi:"diskPath"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringOutput `pulumi:"gatewayArn"`
}

// NewUploadBuffer registers a new resource with the given unique name, arguments, and options.
func NewUploadBuffer(ctx *pulumi.Context,
	name string, args *UploadBufferArgs, opts ...pulumi.ResourceOption) (*UploadBuffer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayArn == nil {
		return nil, errors.New("invalid value for required argument 'GatewayArn'")
	}
	var resource UploadBuffer
	err := ctx.RegisterResource("aws:storagegateway/uploadBuffer:UploadBuffer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUploadBuffer gets an existing UploadBuffer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUploadBuffer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UploadBufferState, opts ...pulumi.ResourceOption) (*UploadBuffer, error) {
	var resource UploadBuffer
	err := ctx.ReadResource("aws:storagegateway/uploadBuffer:UploadBuffer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UploadBuffer resources.
type uploadBufferState struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId *string `pulumi:"diskId"`
	// Local disk path. For example, `/dev/nvme1n1`.
	DiskPath *string `pulumi:"diskPath"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn *string `pulumi:"gatewayArn"`
}

type UploadBufferState struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId pulumi.StringPtrInput
	// Local disk path. For example, `/dev/nvme1n1`.
	DiskPath pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringPtrInput
}

func (UploadBufferState) ElementType() reflect.Type {
	return reflect.TypeOf((*uploadBufferState)(nil)).Elem()
}

type uploadBufferArgs struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId *string `pulumi:"diskId"`
	// Local disk path. For example, `/dev/nvme1n1`.
	DiskPath *string `pulumi:"diskPath"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn string `pulumi:"gatewayArn"`
}

// The set of arguments for constructing a UploadBuffer resource.
type UploadBufferArgs struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId pulumi.StringPtrInput
	// Local disk path. For example, `/dev/nvme1n1`.
	DiskPath pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringInput
}

func (UploadBufferArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*uploadBufferArgs)(nil)).Elem()
}

type UploadBufferInput interface {
	pulumi.Input

	ToUploadBufferOutput() UploadBufferOutput
	ToUploadBufferOutputWithContext(ctx context.Context) UploadBufferOutput
}

func (*UploadBuffer) ElementType() reflect.Type {
	return reflect.TypeOf((*UploadBuffer)(nil))
}

func (i *UploadBuffer) ToUploadBufferOutput() UploadBufferOutput {
	return i.ToUploadBufferOutputWithContext(context.Background())
}

func (i *UploadBuffer) ToUploadBufferOutputWithContext(ctx context.Context) UploadBufferOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadBufferOutput)
}

func (i *UploadBuffer) ToUploadBufferPtrOutput() UploadBufferPtrOutput {
	return i.ToUploadBufferPtrOutputWithContext(context.Background())
}

func (i *UploadBuffer) ToUploadBufferPtrOutputWithContext(ctx context.Context) UploadBufferPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadBufferPtrOutput)
}

type UploadBufferPtrInput interface {
	pulumi.Input

	ToUploadBufferPtrOutput() UploadBufferPtrOutput
	ToUploadBufferPtrOutputWithContext(ctx context.Context) UploadBufferPtrOutput
}

type uploadBufferPtrType UploadBufferArgs

func (*uploadBufferPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UploadBuffer)(nil))
}

func (i *uploadBufferPtrType) ToUploadBufferPtrOutput() UploadBufferPtrOutput {
	return i.ToUploadBufferPtrOutputWithContext(context.Background())
}

func (i *uploadBufferPtrType) ToUploadBufferPtrOutputWithContext(ctx context.Context) UploadBufferPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadBufferPtrOutput)
}

// UploadBufferArrayInput is an input type that accepts UploadBufferArray and UploadBufferArrayOutput values.
// You can construct a concrete instance of `UploadBufferArrayInput` via:
//
//          UploadBufferArray{ UploadBufferArgs{...} }
type UploadBufferArrayInput interface {
	pulumi.Input

	ToUploadBufferArrayOutput() UploadBufferArrayOutput
	ToUploadBufferArrayOutputWithContext(context.Context) UploadBufferArrayOutput
}

type UploadBufferArray []UploadBufferInput

func (UploadBufferArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UploadBuffer)(nil)).Elem()
}

func (i UploadBufferArray) ToUploadBufferArrayOutput() UploadBufferArrayOutput {
	return i.ToUploadBufferArrayOutputWithContext(context.Background())
}

func (i UploadBufferArray) ToUploadBufferArrayOutputWithContext(ctx context.Context) UploadBufferArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadBufferArrayOutput)
}

// UploadBufferMapInput is an input type that accepts UploadBufferMap and UploadBufferMapOutput values.
// You can construct a concrete instance of `UploadBufferMapInput` via:
//
//          UploadBufferMap{ "key": UploadBufferArgs{...} }
type UploadBufferMapInput interface {
	pulumi.Input

	ToUploadBufferMapOutput() UploadBufferMapOutput
	ToUploadBufferMapOutputWithContext(context.Context) UploadBufferMapOutput
}

type UploadBufferMap map[string]UploadBufferInput

func (UploadBufferMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UploadBuffer)(nil)).Elem()
}

func (i UploadBufferMap) ToUploadBufferMapOutput() UploadBufferMapOutput {
	return i.ToUploadBufferMapOutputWithContext(context.Background())
}

func (i UploadBufferMap) ToUploadBufferMapOutputWithContext(ctx context.Context) UploadBufferMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadBufferMapOutput)
}

type UploadBufferOutput struct{ *pulumi.OutputState }

func (UploadBufferOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UploadBuffer)(nil))
}

func (o UploadBufferOutput) ToUploadBufferOutput() UploadBufferOutput {
	return o
}

func (o UploadBufferOutput) ToUploadBufferOutputWithContext(ctx context.Context) UploadBufferOutput {
	return o
}

func (o UploadBufferOutput) ToUploadBufferPtrOutput() UploadBufferPtrOutput {
	return o.ToUploadBufferPtrOutputWithContext(context.Background())
}

func (o UploadBufferOutput) ToUploadBufferPtrOutputWithContext(ctx context.Context) UploadBufferPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UploadBuffer) *UploadBuffer {
		return &v
	}).(UploadBufferPtrOutput)
}

type UploadBufferPtrOutput struct{ *pulumi.OutputState }

func (UploadBufferPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UploadBuffer)(nil))
}

func (o UploadBufferPtrOutput) ToUploadBufferPtrOutput() UploadBufferPtrOutput {
	return o
}

func (o UploadBufferPtrOutput) ToUploadBufferPtrOutputWithContext(ctx context.Context) UploadBufferPtrOutput {
	return o
}

func (o UploadBufferPtrOutput) Elem() UploadBufferOutput {
	return o.ApplyT(func(v *UploadBuffer) UploadBuffer {
		if v != nil {
			return *v
		}
		var ret UploadBuffer
		return ret
	}).(UploadBufferOutput)
}

type UploadBufferArrayOutput struct{ *pulumi.OutputState }

func (UploadBufferArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UploadBuffer)(nil))
}

func (o UploadBufferArrayOutput) ToUploadBufferArrayOutput() UploadBufferArrayOutput {
	return o
}

func (o UploadBufferArrayOutput) ToUploadBufferArrayOutputWithContext(ctx context.Context) UploadBufferArrayOutput {
	return o
}

func (o UploadBufferArrayOutput) Index(i pulumi.IntInput) UploadBufferOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UploadBuffer {
		return vs[0].([]UploadBuffer)[vs[1].(int)]
	}).(UploadBufferOutput)
}

type UploadBufferMapOutput struct{ *pulumi.OutputState }

func (UploadBufferMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UploadBuffer)(nil))
}

func (o UploadBufferMapOutput) ToUploadBufferMapOutput() UploadBufferMapOutput {
	return o
}

func (o UploadBufferMapOutput) ToUploadBufferMapOutputWithContext(ctx context.Context) UploadBufferMapOutput {
	return o
}

func (o UploadBufferMapOutput) MapIndex(k pulumi.StringInput) UploadBufferOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UploadBuffer {
		return vs[0].(map[string]UploadBuffer)[vs[1].(string)]
	}).(UploadBufferOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UploadBufferInput)(nil)).Elem(), &UploadBuffer{})
	pulumi.RegisterInputType(reflect.TypeOf((*UploadBufferPtrInput)(nil)).Elem(), &UploadBuffer{})
	pulumi.RegisterInputType(reflect.TypeOf((*UploadBufferArrayInput)(nil)).Elem(), UploadBufferArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UploadBufferMapInput)(nil)).Elem(), UploadBufferMap{})
	pulumi.RegisterOutputType(UploadBufferOutput{})
	pulumi.RegisterOutputType(UploadBufferPtrOutput{})
	pulumi.RegisterOutputType(UploadBufferArrayOutput{})
	pulumi.RegisterOutputType(UploadBufferMapOutput{})
}
