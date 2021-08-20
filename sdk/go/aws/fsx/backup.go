// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fsx

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a FSx Backup resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/fsx"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleLustreFileSystem, err := fsx.NewLustreFileSystem(ctx, "exampleLustreFileSystem", &fsx.LustreFileSystemArgs{
// 			StorageCapacity: pulumi.Int(1200),
// 			SubnetIds: pulumi.String{
// 				aws_subnet.Example.Id,
// 			},
// 			DeploymentType:           pulumi.String("PERSISTENT_1"),
// 			PerUnitStorageThroughput: pulumi.Int(50),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = fsx.NewBackup(ctx, "exampleBackup", &fsx.BackupArgs{
// 			FileSystemId: exampleLustreFileSystem.ID(),
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
// FSx Backups can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:fsx/backup:Backup example fs-543ab12b1ca672f33
// ```
type Backup struct {
	pulumi.CustomResourceState

	// Amazon Resource Name of the backup.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID of the file system to back up.
	FileSystemId pulumi.StringOutput `pulumi:"fileSystemId"`
	// The ID of the AWS Key Management Service (AWS KMS) key used to encrypt the backup of the Amazon FSx file system's data at rest.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// AWS account identifier that created the file system.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A map of tags to assign to the file system. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copyTagsToBackups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The type of the file system backup.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBackup registers a new resource with the given unique name, arguments, and options.
func NewBackup(ctx *pulumi.Context,
	name string, args *BackupArgs, opts ...pulumi.ResourceOption) (*Backup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileSystemId == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemId'")
	}
	var resource Backup
	err := ctx.RegisterResource("aws:fsx/backup:Backup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackup gets an existing Backup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupState, opts ...pulumi.ResourceOption) (*Backup, error) {
	var resource Backup
	err := ctx.ReadResource("aws:fsx/backup:Backup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Backup resources.
type backupState struct {
	// Amazon Resource Name of the backup.
	Arn *string `pulumi:"arn"`
	// The ID of the file system to back up.
	FileSystemId *string `pulumi:"fileSystemId"`
	// The ID of the AWS Key Management Service (AWS KMS) key used to encrypt the backup of the Amazon FSx file system's data at rest.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// AWS account identifier that created the file system.
	OwnerId *string `pulumi:"ownerId"`
	// A map of tags to assign to the file system. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copyTagsToBackups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The type of the file system backup.
	Type *string `pulumi:"type"`
}

type BackupState struct {
	// Amazon Resource Name of the backup.
	Arn pulumi.StringPtrInput
	// The ID of the file system to back up.
	FileSystemId pulumi.StringPtrInput
	// The ID of the AWS Key Management Service (AWS KMS) key used to encrypt the backup of the Amazon FSx file system's data at rest.
	KmsKeyId pulumi.StringPtrInput
	// AWS account identifier that created the file system.
	OwnerId pulumi.StringPtrInput
	// A map of tags to assign to the file system. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copyTagsToBackups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
	// The type of the file system backup.
	Type pulumi.StringPtrInput
}

func (BackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupState)(nil)).Elem()
}

type backupArgs struct {
	// The ID of the file system to back up.
	FileSystemId string `pulumi:"fileSystemId"`
	// A map of tags to assign to the file system. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copyTagsToBackups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a Backup resource.
type BackupArgs struct {
	// The ID of the file system to back up.
	FileSystemId pulumi.StringInput
	// A map of tags to assign to the file system. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copyTagsToBackups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
}

func (BackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupArgs)(nil)).Elem()
}

type BackupInput interface {
	pulumi.Input

	ToBackupOutput() BackupOutput
	ToBackupOutputWithContext(ctx context.Context) BackupOutput
}

func (*Backup) ElementType() reflect.Type {
	return reflect.TypeOf((*Backup)(nil))
}

func (i *Backup) ToBackupOutput() BackupOutput {
	return i.ToBackupOutputWithContext(context.Background())
}

func (i *Backup) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupOutput)
}

func (i *Backup) ToBackupPtrOutput() BackupPtrOutput {
	return i.ToBackupPtrOutputWithContext(context.Background())
}

func (i *Backup) ToBackupPtrOutputWithContext(ctx context.Context) BackupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPtrOutput)
}

type BackupPtrInput interface {
	pulumi.Input

	ToBackupPtrOutput() BackupPtrOutput
	ToBackupPtrOutputWithContext(ctx context.Context) BackupPtrOutput
}

type backupPtrType BackupArgs

func (*backupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil))
}

func (i *backupPtrType) ToBackupPtrOutput() BackupPtrOutput {
	return i.ToBackupPtrOutputWithContext(context.Background())
}

func (i *backupPtrType) ToBackupPtrOutputWithContext(ctx context.Context) BackupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPtrOutput)
}

// BackupArrayInput is an input type that accepts BackupArray and BackupArrayOutput values.
// You can construct a concrete instance of `BackupArrayInput` via:
//
//          BackupArray{ BackupArgs{...} }
type BackupArrayInput interface {
	pulumi.Input

	ToBackupArrayOutput() BackupArrayOutput
	ToBackupArrayOutputWithContext(context.Context) BackupArrayOutput
}

type BackupArray []BackupInput

func (BackupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Backup)(nil)).Elem()
}

func (i BackupArray) ToBackupArrayOutput() BackupArrayOutput {
	return i.ToBackupArrayOutputWithContext(context.Background())
}

func (i BackupArray) ToBackupArrayOutputWithContext(ctx context.Context) BackupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupArrayOutput)
}

// BackupMapInput is an input type that accepts BackupMap and BackupMapOutput values.
// You can construct a concrete instance of `BackupMapInput` via:
//
//          BackupMap{ "key": BackupArgs{...} }
type BackupMapInput interface {
	pulumi.Input

	ToBackupMapOutput() BackupMapOutput
	ToBackupMapOutputWithContext(context.Context) BackupMapOutput
}

type BackupMap map[string]BackupInput

func (BackupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Backup)(nil)).Elem()
}

func (i BackupMap) ToBackupMapOutput() BackupMapOutput {
	return i.ToBackupMapOutputWithContext(context.Background())
}

func (i BackupMap) ToBackupMapOutputWithContext(ctx context.Context) BackupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupMapOutput)
}

type BackupOutput struct{ *pulumi.OutputState }

func (BackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Backup)(nil))
}

func (o BackupOutput) ToBackupOutput() BackupOutput {
	return o
}

func (o BackupOutput) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return o
}

func (o BackupOutput) ToBackupPtrOutput() BackupPtrOutput {
	return o.ToBackupPtrOutputWithContext(context.Background())
}

func (o BackupOutput) ToBackupPtrOutputWithContext(ctx context.Context) BackupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Backup) *Backup {
		return &v
	}).(BackupPtrOutput)
}

type BackupPtrOutput struct{ *pulumi.OutputState }

func (BackupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil))
}

func (o BackupPtrOutput) ToBackupPtrOutput() BackupPtrOutput {
	return o
}

func (o BackupPtrOutput) ToBackupPtrOutputWithContext(ctx context.Context) BackupPtrOutput {
	return o
}

func (o BackupPtrOutput) Elem() BackupOutput {
	return o.ApplyT(func(v *Backup) Backup {
		if v != nil {
			return *v
		}
		var ret Backup
		return ret
	}).(BackupOutput)
}

type BackupArrayOutput struct{ *pulumi.OutputState }

func (BackupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Backup)(nil))
}

func (o BackupArrayOutput) ToBackupArrayOutput() BackupArrayOutput {
	return o
}

func (o BackupArrayOutput) ToBackupArrayOutputWithContext(ctx context.Context) BackupArrayOutput {
	return o
}

func (o BackupArrayOutput) Index(i pulumi.IntInput) BackupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Backup {
		return vs[0].([]Backup)[vs[1].(int)]
	}).(BackupOutput)
}

type BackupMapOutput struct{ *pulumi.OutputState }

func (BackupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Backup)(nil))
}

func (o BackupMapOutput) ToBackupMapOutput() BackupMapOutput {
	return o
}

func (o BackupMapOutput) ToBackupMapOutputWithContext(ctx context.Context) BackupMapOutput {
	return o
}

func (o BackupMapOutput) MapIndex(k pulumi.StringInput) BackupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Backup {
		return vs[0].(map[string]Backup)[vs[1].(string)]
	}).(BackupOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupOutput{})
	pulumi.RegisterOutputType(BackupPtrOutput{})
	pulumi.RegisterOutputType(BackupArrayOutput{})
	pulumi.RegisterOutputType(BackupMapOutput{})
}
