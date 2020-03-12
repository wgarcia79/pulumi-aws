// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ebs

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a Snapshot of a snapshot.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ebs_snapshot_copy.html.markdown.
type SnapshotCopy struct {
	pulumi.CustomResourceState

	// The data encryption key identifier for the snapshot.
	// * `sourceSnapshotId` The ARN of the copied snapshot.
	// * `sourceRegion` The region of the source snapshot.
	DataEncryptionKeyId pulumi.StringOutput `pulumi:"dataEncryptionKeyId"`
	// A description of what the snapshot is.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the snapshot is encrypted.
	Encrypted pulumi.BoolPtrOutput `pulumi:"encrypted"`
	// The ARN for the KMS encryption key.
	// * `sourceSnapshotId` The ARN for the snapshot to be copied.
	// * `sourceRegion` The region of the source snapshot.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
	OwnerAlias pulumi.StringOutput `pulumi:"ownerAlias"`
	// The AWS account ID of the snapshot owner.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	SourceRegion pulumi.StringOutput `pulumi:"sourceRegion"`
	SourceSnapshotId pulumi.StringOutput `pulumi:"sourceSnapshotId"`
	// A mapping of tags for the snapshot.
	Tags pulumi.MapOutput `pulumi:"tags"`
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
	// The size of the drive in GiBs.
	VolumeSize pulumi.IntOutput `pulumi:"volumeSize"`
}

// NewSnapshotCopy registers a new resource with the given unique name, arguments, and options.
func NewSnapshotCopy(ctx *pulumi.Context,
	name string, args *SnapshotCopyArgs, opts ...pulumi.ResourceOption) (*SnapshotCopy, error) {
	if args == nil || args.SourceRegion == nil {
		return nil, errors.New("missing required argument 'SourceRegion'")
	}
	if args == nil || args.SourceSnapshotId == nil {
		return nil, errors.New("missing required argument 'SourceSnapshotId'")
	}
	if args == nil {
		args = &SnapshotCopyArgs{}
	}
	var resource SnapshotCopy
	err := ctx.RegisterResource("aws:ebs/snapshotCopy:SnapshotCopy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotCopy gets an existing SnapshotCopy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotCopy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotCopyState, opts ...pulumi.ResourceOption) (*SnapshotCopy, error) {
	var resource SnapshotCopy
	err := ctx.ReadResource("aws:ebs/snapshotCopy:SnapshotCopy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotCopy resources.
type snapshotCopyState struct {
	// The data encryption key identifier for the snapshot.
	// * `sourceSnapshotId` The ARN of the copied snapshot.
	// * `sourceRegion` The region of the source snapshot.
	DataEncryptionKeyId *string `pulumi:"dataEncryptionKeyId"`
	// A description of what the snapshot is.
	Description *string `pulumi:"description"`
	// Whether the snapshot is encrypted.
	Encrypted *bool `pulumi:"encrypted"`
	// The ARN for the KMS encryption key.
	// * `sourceSnapshotId` The ARN for the snapshot to be copied.
	// * `sourceRegion` The region of the source snapshot.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
	OwnerAlias *string `pulumi:"ownerAlias"`
	// The AWS account ID of the snapshot owner.
	OwnerId *string `pulumi:"ownerId"`
	SourceRegion *string `pulumi:"sourceRegion"`
	SourceSnapshotId *string `pulumi:"sourceSnapshotId"`
	// A mapping of tags for the snapshot.
	Tags map[string]interface{} `pulumi:"tags"`
	VolumeId *string `pulumi:"volumeId"`
	// The size of the drive in GiBs.
	VolumeSize *int `pulumi:"volumeSize"`
}

type SnapshotCopyState struct {
	// The data encryption key identifier for the snapshot.
	// * `sourceSnapshotId` The ARN of the copied snapshot.
	// * `sourceRegion` The region of the source snapshot.
	DataEncryptionKeyId pulumi.StringPtrInput
	// A description of what the snapshot is.
	Description pulumi.StringPtrInput
	// Whether the snapshot is encrypted.
	Encrypted pulumi.BoolPtrInput
	// The ARN for the KMS encryption key.
	// * `sourceSnapshotId` The ARN for the snapshot to be copied.
	// * `sourceRegion` The region of the source snapshot.
	KmsKeyId pulumi.StringPtrInput
	// Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
	OwnerAlias pulumi.StringPtrInput
	// The AWS account ID of the snapshot owner.
	OwnerId pulumi.StringPtrInput
	SourceRegion pulumi.StringPtrInput
	SourceSnapshotId pulumi.StringPtrInput
	// A mapping of tags for the snapshot.
	Tags pulumi.MapInput
	VolumeId pulumi.StringPtrInput
	// The size of the drive in GiBs.
	VolumeSize pulumi.IntPtrInput
}

func (SnapshotCopyState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotCopyState)(nil)).Elem()
}

type snapshotCopyArgs struct {
	// A description of what the snapshot is.
	Description *string `pulumi:"description"`
	// Whether the snapshot is encrypted.
	Encrypted *bool `pulumi:"encrypted"`
	// The ARN for the KMS encryption key.
	// * `sourceSnapshotId` The ARN for the snapshot to be copied.
	// * `sourceRegion` The region of the source snapshot.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	SourceRegion string `pulumi:"sourceRegion"`
	SourceSnapshotId string `pulumi:"sourceSnapshotId"`
	// A mapping of tags for the snapshot.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a SnapshotCopy resource.
type SnapshotCopyArgs struct {
	// A description of what the snapshot is.
	Description pulumi.StringPtrInput
	// Whether the snapshot is encrypted.
	Encrypted pulumi.BoolPtrInput
	// The ARN for the KMS encryption key.
	// * `sourceSnapshotId` The ARN for the snapshot to be copied.
	// * `sourceRegion` The region of the source snapshot.
	KmsKeyId pulumi.StringPtrInput
	SourceRegion pulumi.StringInput
	SourceSnapshotId pulumi.StringInput
	// A mapping of tags for the snapshot.
	Tags pulumi.MapInput
}

func (SnapshotCopyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotCopyArgs)(nil)).Elem()
}

