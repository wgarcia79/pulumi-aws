// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Elastic File System (EFS) File System resource.
//
// ## Example Usage
// ### EFS File System w/ tags
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/efs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := efs.NewFileSystem(ctx, "foo", &efs.FileSystemArgs{
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("MyProduct"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using lifecycle policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/efs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := efs.NewFileSystem(ctx, "fooWithLifecylePolicy", &efs.FileSystemArgs{
// 			LifecyclePolicy: &efs.FileSystemLifecyclePolicyArgs{
// 				TransitionToIa: pulumi.String("AFTER_30_DAYS"),
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
// The EFS file systems can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:efs/fileSystem:FileSystem foo fs-6fa144c6
// ```
type FileSystem struct {
	pulumi.CustomResourceState

	// Amazon Resource Name of the file system.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A unique name (a maximum of 64 characters are allowed)
	// used as reference when creating the Elastic File System to ensure idempotent file
	// system creation. By default generated by this provider. See [Elastic File System]
	// (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
	CreationToken pulumi.StringOutput `pulumi:"creationToken"`
	// The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// If true, the disk will be encrypted.
	Encrypted pulumi.BoolOutput `pulumi:"encrypted"`
	// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object (documented below).
	LifecyclePolicy FileSystemLifecyclePolicyPtrOutput `pulumi:"lifecyclePolicy"`
	// The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
	PerformanceMode pulumi.StringOutput `pulumi:"performanceMode"`
	// The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughputMode` set to `provisioned`.
	ProvisionedThroughputInMibps pulumi.Float64PtrOutput `pulumi:"provisionedThroughputInMibps"`
	// A map of tags to assign to the file system.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisionedThroughputInMibps`.
	ThroughputMode pulumi.StringPtrOutput `pulumi:"throughputMode"`
}

// NewFileSystem registers a new resource with the given unique name, arguments, and options.
func NewFileSystem(ctx *pulumi.Context,
	name string, args *FileSystemArgs, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	if args == nil {
		args = &FileSystemArgs{}
	}

	var resource FileSystem
	err := ctx.RegisterResource("aws:efs/fileSystem:FileSystem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileSystem gets an existing FileSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileSystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileSystemState, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	var resource FileSystem
	err := ctx.ReadResource("aws:efs/fileSystem:FileSystem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileSystem resources.
type fileSystemState struct {
	// Amazon Resource Name of the file system.
	Arn *string `pulumi:"arn"`
	// A unique name (a maximum of 64 characters are allowed)
	// used as reference when creating the Elastic File System to ensure idempotent file
	// system creation. By default generated by this provider. See [Elastic File System]
	// (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
	CreationToken *string `pulumi:"creationToken"`
	// The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
	DnsName *string `pulumi:"dnsName"`
	// If true, the disk will be encrypted.
	Encrypted *bool `pulumi:"encrypted"`
	// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object (documented below).
	LifecyclePolicy *FileSystemLifecyclePolicy `pulumi:"lifecyclePolicy"`
	// The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
	PerformanceMode *string `pulumi:"performanceMode"`
	// The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughputMode` set to `provisioned`.
	ProvisionedThroughputInMibps *float64 `pulumi:"provisionedThroughputInMibps"`
	// A map of tags to assign to the file system.
	Tags map[string]string `pulumi:"tags"`
	// Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisionedThroughputInMibps`.
	ThroughputMode *string `pulumi:"throughputMode"`
}

type FileSystemState struct {
	// Amazon Resource Name of the file system.
	Arn pulumi.StringPtrInput
	// A unique name (a maximum of 64 characters are allowed)
	// used as reference when creating the Elastic File System to ensure idempotent file
	// system creation. By default generated by this provider. See [Elastic File System]
	// (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
	CreationToken pulumi.StringPtrInput
	// The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
	DnsName pulumi.StringPtrInput
	// If true, the disk will be encrypted.
	Encrypted pulumi.BoolPtrInput
	// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
	KmsKeyId pulumi.StringPtrInput
	// A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object (documented below).
	LifecyclePolicy FileSystemLifecyclePolicyPtrInput
	// The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
	PerformanceMode pulumi.StringPtrInput
	// The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughputMode` set to `provisioned`.
	ProvisionedThroughputInMibps pulumi.Float64PtrInput
	// A map of tags to assign to the file system.
	Tags pulumi.StringMapInput
	// Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisionedThroughputInMibps`.
	ThroughputMode pulumi.StringPtrInput
}

func (FileSystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemState)(nil)).Elem()
}

type fileSystemArgs struct {
	// A unique name (a maximum of 64 characters are allowed)
	// used as reference when creating the Elastic File System to ensure idempotent file
	// system creation. By default generated by this provider. See [Elastic File System]
	// (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
	CreationToken *string `pulumi:"creationToken"`
	// If true, the disk will be encrypted.
	Encrypted *bool `pulumi:"encrypted"`
	// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object (documented below).
	LifecyclePolicy *FileSystemLifecyclePolicy `pulumi:"lifecyclePolicy"`
	// The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
	PerformanceMode *string `pulumi:"performanceMode"`
	// The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughputMode` set to `provisioned`.
	ProvisionedThroughputInMibps *float64 `pulumi:"provisionedThroughputInMibps"`
	// A map of tags to assign to the file system.
	Tags map[string]string `pulumi:"tags"`
	// Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisionedThroughputInMibps`.
	ThroughputMode *string `pulumi:"throughputMode"`
}

// The set of arguments for constructing a FileSystem resource.
type FileSystemArgs struct {
	// A unique name (a maximum of 64 characters are allowed)
	// used as reference when creating the Elastic File System to ensure idempotent file
	// system creation. By default generated by this provider. See [Elastic File System]
	// (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
	CreationToken pulumi.StringPtrInput
	// If true, the disk will be encrypted.
	Encrypted pulumi.BoolPtrInput
	// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
	KmsKeyId pulumi.StringPtrInput
	// A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object (documented below).
	LifecyclePolicy FileSystemLifecyclePolicyPtrInput
	// The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
	PerformanceMode pulumi.StringPtrInput
	// The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughputMode` set to `provisioned`.
	ProvisionedThroughputInMibps pulumi.Float64PtrInput
	// A map of tags to assign to the file system.
	Tags pulumi.StringMapInput
	// Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisionedThroughputInMibps`.
	ThroughputMode pulumi.StringPtrInput
}

func (FileSystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemArgs)(nil)).Elem()
}

type FileSystemInput interface {
	pulumi.Input

	ToFileSystemOutput() FileSystemOutput
	ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput
}

func (*FileSystem) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystem)(nil))
}

func (i *FileSystem) ToFileSystemOutput() FileSystemOutput {
	return i.ToFileSystemOutputWithContext(context.Background())
}

func (i *FileSystem) ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemOutput)
}

func (i *FileSystem) ToFileSystemPtrOutput() FileSystemPtrOutput {
	return i.ToFileSystemPtrOutputWithContext(context.Background())
}

func (i *FileSystem) ToFileSystemPtrOutputWithContext(ctx context.Context) FileSystemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemPtrOutput)
}

type FileSystemPtrInput interface {
	pulumi.Input

	ToFileSystemPtrOutput() FileSystemPtrOutput
	ToFileSystemPtrOutputWithContext(ctx context.Context) FileSystemPtrOutput
}

type fileSystemPtrType FileSystemArgs

func (*fileSystemPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystem)(nil))
}

func (i *fileSystemPtrType) ToFileSystemPtrOutput() FileSystemPtrOutput {
	return i.ToFileSystemPtrOutputWithContext(context.Background())
}

func (i *fileSystemPtrType) ToFileSystemPtrOutputWithContext(ctx context.Context) FileSystemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemPtrOutput)
}

// FileSystemArrayInput is an input type that accepts FileSystemArray and FileSystemArrayOutput values.
// You can construct a concrete instance of `FileSystemArrayInput` via:
//
//          FileSystemArray{ FileSystemArgs{...} }
type FileSystemArrayInput interface {
	pulumi.Input

	ToFileSystemArrayOutput() FileSystemArrayOutput
	ToFileSystemArrayOutputWithContext(context.Context) FileSystemArrayOutput
}

type FileSystemArray []FileSystemInput

func (FileSystemArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FileSystem)(nil))
}

func (i FileSystemArray) ToFileSystemArrayOutput() FileSystemArrayOutput {
	return i.ToFileSystemArrayOutputWithContext(context.Background())
}

func (i FileSystemArray) ToFileSystemArrayOutputWithContext(ctx context.Context) FileSystemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemArrayOutput)
}

// FileSystemMapInput is an input type that accepts FileSystemMap and FileSystemMapOutput values.
// You can construct a concrete instance of `FileSystemMapInput` via:
//
//          FileSystemMap{ "key": FileSystemArgs{...} }
type FileSystemMapInput interface {
	pulumi.Input

	ToFileSystemMapOutput() FileSystemMapOutput
	ToFileSystemMapOutputWithContext(context.Context) FileSystemMapOutput
}

type FileSystemMap map[string]FileSystemInput

func (FileSystemMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FileSystem)(nil))
}

func (i FileSystemMap) ToFileSystemMapOutput() FileSystemMapOutput {
	return i.ToFileSystemMapOutputWithContext(context.Background())
}

func (i FileSystemMap) ToFileSystemMapOutputWithContext(ctx context.Context) FileSystemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemMapOutput)
}

type FileSystemOutput struct {
	*pulumi.OutputState
}

func (FileSystemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystem)(nil))
}

func (o FileSystemOutput) ToFileSystemOutput() FileSystemOutput {
	return o
}

func (o FileSystemOutput) ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput {
	return o
}

func (o FileSystemOutput) ToFileSystemPtrOutput() FileSystemPtrOutput {
	return o.ToFileSystemPtrOutputWithContext(context.Background())
}

func (o FileSystemOutput) ToFileSystemPtrOutputWithContext(ctx context.Context) FileSystemPtrOutput {
	return o.ApplyT(func(v FileSystem) *FileSystem {
		return &v
	}).(FileSystemPtrOutput)
}

type FileSystemPtrOutput struct {
	*pulumi.OutputState
}

func (FileSystemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystem)(nil))
}

func (o FileSystemPtrOutput) ToFileSystemPtrOutput() FileSystemPtrOutput {
	return o
}

func (o FileSystemPtrOutput) ToFileSystemPtrOutputWithContext(ctx context.Context) FileSystemPtrOutput {
	return o
}

type FileSystemArrayOutput struct{ *pulumi.OutputState }

func (FileSystemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FileSystem)(nil))
}

func (o FileSystemArrayOutput) ToFileSystemArrayOutput() FileSystemArrayOutput {
	return o
}

func (o FileSystemArrayOutput) ToFileSystemArrayOutputWithContext(ctx context.Context) FileSystemArrayOutput {
	return o
}

func (o FileSystemArrayOutput) Index(i pulumi.IntInput) FileSystemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FileSystem {
		return vs[0].([]FileSystem)[vs[1].(int)]
	}).(FileSystemOutput)
}

type FileSystemMapOutput struct{ *pulumi.OutputState }

func (FileSystemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FileSystem)(nil))
}

func (o FileSystemMapOutput) ToFileSystemMapOutput() FileSystemMapOutput {
	return o
}

func (o FileSystemMapOutput) ToFileSystemMapOutputWithContext(ctx context.Context) FileSystemMapOutput {
	return o
}

func (o FileSystemMapOutput) MapIndex(k pulumi.StringInput) FileSystemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FileSystem {
		return vs[0].(map[string]FileSystem)[vs[1].(string)]
	}).(FileSystemOutput)
}

func init() {
	pulumi.RegisterOutputType(FileSystemOutput{})
	pulumi.RegisterOutputType(FileSystemPtrOutput{})
	pulumi.RegisterOutputType(FileSystemArrayOutput{})
	pulumi.RegisterOutputType(FileSystemMapOutput{})
}
