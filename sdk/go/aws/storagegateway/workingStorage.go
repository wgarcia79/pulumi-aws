// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an AWS Storage Gateway working storage.
//
// > **NOTE:** The Storage Gateway API provides no method to remove a working storage disk. Destroying this resource does not perform any Storage Gateway actions.
//
// ## Example Usage
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
// 		_, err := storagegateway.NewWorkingStorage(ctx, "example", &storagegateway.WorkingStorageArgs{
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
// `aws_storagegateway_working_storage` can be imported by using the gateway Amazon Resource Name (ARN) and local disk identifier separated with a colon (`:`), e.g.
//
// ```sh
//  $ pulumi import aws:storagegateway/workingStorage:WorkingStorage example arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678:pci-0000:03:00.0-scsi-0:0:0:0
// ```
type WorkingStorage struct {
	pulumi.CustomResourceState

	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId pulumi.StringOutput `pulumi:"diskId"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringOutput `pulumi:"gatewayArn"`
}

// NewWorkingStorage registers a new resource with the given unique name, arguments, and options.
func NewWorkingStorage(ctx *pulumi.Context,
	name string, args *WorkingStorageArgs, opts ...pulumi.ResourceOption) (*WorkingStorage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskId == nil {
		return nil, errors.New("invalid value for required argument 'DiskId'")
	}
	if args.GatewayArn == nil {
		return nil, errors.New("invalid value for required argument 'GatewayArn'")
	}
	var resource WorkingStorage
	err := ctx.RegisterResource("aws:storagegateway/workingStorage:WorkingStorage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkingStorage gets an existing WorkingStorage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkingStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkingStorageState, opts ...pulumi.ResourceOption) (*WorkingStorage, error) {
	var resource WorkingStorage
	err := ctx.ReadResource("aws:storagegateway/workingStorage:WorkingStorage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkingStorage resources.
type workingStorageState struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId *string `pulumi:"diskId"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn *string `pulumi:"gatewayArn"`
}

type WorkingStorageState struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringPtrInput
}

func (WorkingStorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*workingStorageState)(nil)).Elem()
}

type workingStorageArgs struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId string `pulumi:"diskId"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn string `pulumi:"gatewayArn"`
}

// The set of arguments for constructing a WorkingStorage resource.
type WorkingStorageArgs struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId pulumi.StringInput
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringInput
}

func (WorkingStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workingStorageArgs)(nil)).Elem()
}

type WorkingStorageInput interface {
	pulumi.Input

	ToWorkingStorageOutput() WorkingStorageOutput
	ToWorkingStorageOutputWithContext(ctx context.Context) WorkingStorageOutput
}

func (*WorkingStorage) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkingStorage)(nil))
}

func (i *WorkingStorage) ToWorkingStorageOutput() WorkingStorageOutput {
	return i.ToWorkingStorageOutputWithContext(context.Background())
}

func (i *WorkingStorage) ToWorkingStorageOutputWithContext(ctx context.Context) WorkingStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkingStorageOutput)
}

func (i *WorkingStorage) ToWorkingStoragePtrOutput() WorkingStoragePtrOutput {
	return i.ToWorkingStoragePtrOutputWithContext(context.Background())
}

func (i *WorkingStorage) ToWorkingStoragePtrOutputWithContext(ctx context.Context) WorkingStoragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkingStoragePtrOutput)
}

type WorkingStoragePtrInput interface {
	pulumi.Input

	ToWorkingStoragePtrOutput() WorkingStoragePtrOutput
	ToWorkingStoragePtrOutputWithContext(ctx context.Context) WorkingStoragePtrOutput
}

type workingStoragePtrType WorkingStorageArgs

func (*workingStoragePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkingStorage)(nil))
}

func (i *workingStoragePtrType) ToWorkingStoragePtrOutput() WorkingStoragePtrOutput {
	return i.ToWorkingStoragePtrOutputWithContext(context.Background())
}

func (i *workingStoragePtrType) ToWorkingStoragePtrOutputWithContext(ctx context.Context) WorkingStoragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkingStoragePtrOutput)
}

// WorkingStorageArrayInput is an input type that accepts WorkingStorageArray and WorkingStorageArrayOutput values.
// You can construct a concrete instance of `WorkingStorageArrayInput` via:
//
//          WorkingStorageArray{ WorkingStorageArgs{...} }
type WorkingStorageArrayInput interface {
	pulumi.Input

	ToWorkingStorageArrayOutput() WorkingStorageArrayOutput
	ToWorkingStorageArrayOutputWithContext(context.Context) WorkingStorageArrayOutput
}

type WorkingStorageArray []WorkingStorageInput

func (WorkingStorageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkingStorage)(nil)).Elem()
}

func (i WorkingStorageArray) ToWorkingStorageArrayOutput() WorkingStorageArrayOutput {
	return i.ToWorkingStorageArrayOutputWithContext(context.Background())
}

func (i WorkingStorageArray) ToWorkingStorageArrayOutputWithContext(ctx context.Context) WorkingStorageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkingStorageArrayOutput)
}

// WorkingStorageMapInput is an input type that accepts WorkingStorageMap and WorkingStorageMapOutput values.
// You can construct a concrete instance of `WorkingStorageMapInput` via:
//
//          WorkingStorageMap{ "key": WorkingStorageArgs{...} }
type WorkingStorageMapInput interface {
	pulumi.Input

	ToWorkingStorageMapOutput() WorkingStorageMapOutput
	ToWorkingStorageMapOutputWithContext(context.Context) WorkingStorageMapOutput
}

type WorkingStorageMap map[string]WorkingStorageInput

func (WorkingStorageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkingStorage)(nil)).Elem()
}

func (i WorkingStorageMap) ToWorkingStorageMapOutput() WorkingStorageMapOutput {
	return i.ToWorkingStorageMapOutputWithContext(context.Background())
}

func (i WorkingStorageMap) ToWorkingStorageMapOutputWithContext(ctx context.Context) WorkingStorageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkingStorageMapOutput)
}

type WorkingStorageOutput struct{ *pulumi.OutputState }

func (WorkingStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkingStorage)(nil))
}

func (o WorkingStorageOutput) ToWorkingStorageOutput() WorkingStorageOutput {
	return o
}

func (o WorkingStorageOutput) ToWorkingStorageOutputWithContext(ctx context.Context) WorkingStorageOutput {
	return o
}

func (o WorkingStorageOutput) ToWorkingStoragePtrOutput() WorkingStoragePtrOutput {
	return o.ToWorkingStoragePtrOutputWithContext(context.Background())
}

func (o WorkingStorageOutput) ToWorkingStoragePtrOutputWithContext(ctx context.Context) WorkingStoragePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkingStorage) *WorkingStorage {
		return &v
	}).(WorkingStoragePtrOutput)
}

type WorkingStoragePtrOutput struct{ *pulumi.OutputState }

func (WorkingStoragePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkingStorage)(nil))
}

func (o WorkingStoragePtrOutput) ToWorkingStoragePtrOutput() WorkingStoragePtrOutput {
	return o
}

func (o WorkingStoragePtrOutput) ToWorkingStoragePtrOutputWithContext(ctx context.Context) WorkingStoragePtrOutput {
	return o
}

func (o WorkingStoragePtrOutput) Elem() WorkingStorageOutput {
	return o.ApplyT(func(v *WorkingStorage) WorkingStorage {
		if v != nil {
			return *v
		}
		var ret WorkingStorage
		return ret
	}).(WorkingStorageOutput)
}

type WorkingStorageArrayOutput struct{ *pulumi.OutputState }

func (WorkingStorageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkingStorage)(nil))
}

func (o WorkingStorageArrayOutput) ToWorkingStorageArrayOutput() WorkingStorageArrayOutput {
	return o
}

func (o WorkingStorageArrayOutput) ToWorkingStorageArrayOutputWithContext(ctx context.Context) WorkingStorageArrayOutput {
	return o
}

func (o WorkingStorageArrayOutput) Index(i pulumi.IntInput) WorkingStorageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkingStorage {
		return vs[0].([]WorkingStorage)[vs[1].(int)]
	}).(WorkingStorageOutput)
}

type WorkingStorageMapOutput struct{ *pulumi.OutputState }

func (WorkingStorageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkingStorage)(nil))
}

func (o WorkingStorageMapOutput) ToWorkingStorageMapOutput() WorkingStorageMapOutput {
	return o
}

func (o WorkingStorageMapOutput) ToWorkingStorageMapOutputWithContext(ctx context.Context) WorkingStorageMapOutput {
	return o
}

func (o WorkingStorageMapOutput) MapIndex(k pulumi.StringInput) WorkingStorageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WorkingStorage {
		return vs[0].(map[string]WorkingStorage)[vs[1].(string)]
	}).(WorkingStorageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkingStorageInput)(nil)).Elem(), &WorkingStorage{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkingStoragePtrInput)(nil)).Elem(), &WorkingStorage{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkingStorageArrayInput)(nil)).Elem(), WorkingStorageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkingStorageMapInput)(nil)).Elem(), WorkingStorageMap{})
	pulumi.RegisterOutputType(WorkingStorageOutput{})
	pulumi.RegisterOutputType(WorkingStoragePtrOutput{})
	pulumi.RegisterOutputType(WorkingStorageArrayOutput{})
	pulumi.RegisterOutputType(WorkingStorageMapOutput{})
}
