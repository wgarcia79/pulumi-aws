// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Snapshot of an EBS Volume.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ebs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := ebs.NewVolume(ctx, "example", &ebs.VolumeArgs{
// 			AvailabilityZone: pulumi.String("us-west-2a"),
// 			Size:             pulumi.Int(40),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("HelloWorld"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ebs.NewSnapshot(ctx, "exampleSnapshot", &ebs.SnapshotArgs{
// 			VolumeId: example.ID(),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("HelloWorld_snap"),
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
// EBS Snapshot can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:ebs/snapshot:Snapshot id snap-049df61146c4d7901
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the EBS Snapshot.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The data encryption key identifier for the snapshot.
	DataEncryptionKeyId pulumi.StringOutput `pulumi:"dataEncryptionKeyId"`
	// A description of what the snapshot is.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the snapshot is encrypted.
	Encrypted pulumi.BoolOutput `pulumi:"encrypted"`
	// The ARN for the KMS encryption key.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
	OwnerAlias pulumi.StringOutput `pulumi:"ownerAlias"`
	// The AWS account ID of the EBS snapshot owner.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A map of tags for the snapshot.
	Tags    pulumi.StringMapOutput `pulumi:"tags"`
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The Volume ID of which to make a snapshot.
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
	// The size of the drive in GiBs.
	VolumeSize pulumi.IntOutput `pulumi:"volumeSize"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VolumeId == nil {
		return nil, errors.New("invalid value for required argument 'VolumeId'")
	}
	var resource Snapshot
	err := ctx.RegisterResource("aws:ebs/snapshot:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("aws:ebs/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// Amazon Resource Name (ARN) of the EBS Snapshot.
	Arn *string `pulumi:"arn"`
	// The data encryption key identifier for the snapshot.
	DataEncryptionKeyId *string `pulumi:"dataEncryptionKeyId"`
	// A description of what the snapshot is.
	Description *string `pulumi:"description"`
	// Whether the snapshot is encrypted.
	Encrypted *bool `pulumi:"encrypted"`
	// The ARN for the KMS encryption key.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
	OwnerAlias *string `pulumi:"ownerAlias"`
	// The AWS account ID of the EBS snapshot owner.
	OwnerId *string `pulumi:"ownerId"`
	// A map of tags for the snapshot.
	Tags    map[string]string `pulumi:"tags"`
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The Volume ID of which to make a snapshot.
	VolumeId *string `pulumi:"volumeId"`
	// The size of the drive in GiBs.
	VolumeSize *int `pulumi:"volumeSize"`
}

type SnapshotState struct {
	// Amazon Resource Name (ARN) of the EBS Snapshot.
	Arn pulumi.StringPtrInput
	// The data encryption key identifier for the snapshot.
	DataEncryptionKeyId pulumi.StringPtrInput
	// A description of what the snapshot is.
	Description pulumi.StringPtrInput
	// Whether the snapshot is encrypted.
	Encrypted pulumi.BoolPtrInput
	// The ARN for the KMS encryption key.
	KmsKeyId pulumi.StringPtrInput
	// Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
	OwnerAlias pulumi.StringPtrInput
	// The AWS account ID of the EBS snapshot owner.
	OwnerId pulumi.StringPtrInput
	// A map of tags for the snapshot.
	Tags    pulumi.StringMapInput
	TagsAll pulumi.StringMapInput
	// The Volume ID of which to make a snapshot.
	VolumeId pulumi.StringPtrInput
	// The size of the drive in GiBs.
	VolumeSize pulumi.IntPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// A description of what the snapshot is.
	Description *string `pulumi:"description"`
	// A map of tags for the snapshot.
	Tags    map[string]string `pulumi:"tags"`
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The Volume ID of which to make a snapshot.
	VolumeId string `pulumi:"volumeId"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// A description of what the snapshot is.
	Description pulumi.StringPtrInput
	// A map of tags for the snapshot.
	Tags    pulumi.StringMapInput
	TagsAll pulumi.StringMapInput
	// The Volume ID of which to make a snapshot.
	VolumeId pulumi.StringInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((*Snapshot)(nil))
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

func (i *Snapshot) ToSnapshotPtrOutput() SnapshotPtrOutput {
	return i.ToSnapshotPtrOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotPtrOutputWithContext(ctx context.Context) SnapshotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotPtrOutput)
}

type SnapshotPtrInput interface {
	pulumi.Input

	ToSnapshotPtrOutput() SnapshotPtrOutput
	ToSnapshotPtrOutputWithContext(ctx context.Context) SnapshotPtrOutput
}

type snapshotPtrType SnapshotArgs

func (*snapshotPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil))
}

func (i *snapshotPtrType) ToSnapshotPtrOutput() SnapshotPtrOutput {
	return i.ToSnapshotPtrOutputWithContext(context.Background())
}

func (i *snapshotPtrType) ToSnapshotPtrOutputWithContext(ctx context.Context) SnapshotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotPtrOutput)
}

// SnapshotArrayInput is an input type that accepts SnapshotArray and SnapshotArrayOutput values.
// You can construct a concrete instance of `SnapshotArrayInput` via:
//
//          SnapshotArray{ SnapshotArgs{...} }
type SnapshotArrayInput interface {
	pulumi.Input

	ToSnapshotArrayOutput() SnapshotArrayOutput
	ToSnapshotArrayOutputWithContext(context.Context) SnapshotArrayOutput
}

type SnapshotArray []SnapshotInput

func (SnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (i SnapshotArray) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return i.ToSnapshotArrayOutputWithContext(context.Background())
}

func (i SnapshotArray) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotArrayOutput)
}

// SnapshotMapInput is an input type that accepts SnapshotMap and SnapshotMapOutput values.
// You can construct a concrete instance of `SnapshotMapInput` via:
//
//          SnapshotMap{ "key": SnapshotArgs{...} }
type SnapshotMapInput interface {
	pulumi.Input

	ToSnapshotMapOutput() SnapshotMapOutput
	ToSnapshotMapOutputWithContext(context.Context) SnapshotMapOutput
}

type SnapshotMap map[string]SnapshotInput

func (SnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (i SnapshotMap) ToSnapshotMapOutput() SnapshotMapOutput {
	return i.ToSnapshotMapOutputWithContext(context.Background())
}

func (i SnapshotMap) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotMapOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Snapshot)(nil))
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotPtrOutput() SnapshotPtrOutput {
	return o.ToSnapshotPtrOutputWithContext(context.Background())
}

func (o SnapshotOutput) ToSnapshotPtrOutputWithContext(ctx context.Context) SnapshotPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Snapshot) *Snapshot {
		return &v
	}).(SnapshotPtrOutput)
}

type SnapshotPtrOutput struct{ *pulumi.OutputState }

func (SnapshotPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil))
}

func (o SnapshotPtrOutput) ToSnapshotPtrOutput() SnapshotPtrOutput {
	return o
}

func (o SnapshotPtrOutput) ToSnapshotPtrOutputWithContext(ctx context.Context) SnapshotPtrOutput {
	return o
}

func (o SnapshotPtrOutput) Elem() SnapshotOutput {
	return o.ApplyT(func(v *Snapshot) Snapshot {
		if v != nil {
			return *v
		}
		var ret Snapshot
		return ret
	}).(SnapshotOutput)
}

type SnapshotArrayOutput struct{ *pulumi.OutputState }

func (SnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Snapshot)(nil))
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) Index(i pulumi.IntInput) SnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Snapshot {
		return vs[0].([]Snapshot)[vs[1].(int)]
	}).(SnapshotOutput)
}

type SnapshotMapOutput struct{ *pulumi.OutputState }

func (SnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Snapshot)(nil))
}

func (o SnapshotMapOutput) ToSnapshotMapOutput() SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) MapIndex(k pulumi.StringInput) SnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Snapshot {
		return vs[0].(map[string]Snapshot)[vs[1].(string)]
	}).(SnapshotOutput)
}

func init() {
	pulumi.RegisterOutputType(SnapshotOutput{})
	pulumi.RegisterOutputType(SnapshotPtrOutput{})
	pulumi.RegisterOutputType(SnapshotArrayOutput{})
	pulumi.RegisterOutputType(SnapshotMapOutput{})
}
