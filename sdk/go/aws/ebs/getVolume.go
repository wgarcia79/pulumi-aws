// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about an EBS volume for use in other
// resources.
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
// 		opt0 := true
// 		_, err := ebs.LookupVolume(ctx, &ebs.LookupVolumeArgs{
// 			Filters: []ebs.GetVolumeFilter{
// 				ebs.GetVolumeFilter{
// 					Name: "volume-type",
// 					Values: []string{
// 						"gp2",
// 					},
// 				},
// 				ebs.GetVolumeFilter{
// 					Name: "tag:Name",
// 					Values: []string{
// 						"Example",
// 					},
// 				},
// 			},
// 			MostRecent: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("aws:ebs/getVolume:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVolume.
type LookupVolumeArgs struct {
	// One or more name/value pairs to filter off of. There are
	// several valid keys, for a full reference, check out
	// [describe-volumes in the AWS CLI reference][1].
	Filters []GetVolumeFilter `pulumi:"filters"`
	// If more than one result is returned, use the most
	// recent Volume.
	MostRecent *bool `pulumi:"mostRecent"`
	// A map of tags for the resource.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getVolume.
type LookupVolumeResult struct {
	// The volume ARN (e.g., arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
	Arn string `pulumi:"arn"`
	// The AZ where the EBS volume exists.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Whether the disk is encrypted.
	Encrypted bool              `pulumi:"encrypted"`
	Filters   []GetVolumeFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The amount of IOPS for the disk.
	Iops int `pulumi:"iops"`
	// The ARN for the KMS encryption key.
	KmsKeyId   string `pulumi:"kmsKeyId"`
	MostRecent *bool  `pulumi:"mostRecent"`
	// (Optional) Specifies whether Amazon EBS Multi-Attach is enabled.
	MultiAttachEnabled bool `pulumi:"multiAttachEnabled"`
	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn string `pulumi:"outpostArn"`
	// The size of the drive in GiBs.
	Size int `pulumi:"size"`
	// The snapshotId the EBS volume is based off.
	SnapshotId string `pulumi:"snapshotId"`
	// A map of tags for the resource.
	Tags map[string]string `pulumi:"tags"`
	// The throughput that the volume supports, in MiB/s.
	Throughput int `pulumi:"throughput"`
	// The volume ID (e.g., vol-59fcb34e).
	VolumeId string `pulumi:"volumeId"`
	// The type of EBS volume.
	VolumeType string `pulumi:"volumeType"`
}

func LookupVolumeOutput(ctx *pulumi.Context, args LookupVolumeOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeResult, error) {
			args := v.(LookupVolumeArgs)
			r, err := LookupVolume(ctx, &args, opts...)
			return *r, err
		}).(LookupVolumeResultOutput)
}

// A collection of arguments for invoking getVolume.
type LookupVolumeOutputArgs struct {
	// One or more name/value pairs to filter off of. There are
	// several valid keys, for a full reference, check out
	// [describe-volumes in the AWS CLI reference][1].
	Filters GetVolumeFilterArrayInput `pulumi:"filters"`
	// If more than one result is returned, use the most
	// recent Volume.
	MostRecent pulumi.BoolPtrInput `pulumi:"mostRecent"`
	// A map of tags for the resource.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupVolumeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeArgs)(nil)).Elem()
}

// A collection of values returned by getVolume.
type LookupVolumeResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeResult)(nil)).Elem()
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutput() LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutputWithContext(ctx context.Context) LookupVolumeResultOutput {
	return o
}

// The volume ARN (e.g., arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
func (o LookupVolumeResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The AZ where the EBS volume exists.
func (o LookupVolumeResultOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Whether the disk is encrypted.
func (o LookupVolumeResultOutput) Encrypted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVolumeResult) bool { return v.Encrypted }).(pulumi.BoolOutput)
}

func (o LookupVolumeResultOutput) Filters() GetVolumeFilterArrayOutput {
	return o.ApplyT(func(v LookupVolumeResult) []GetVolumeFilter { return v.Filters }).(GetVolumeFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVolumeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Id }).(pulumi.StringOutput)
}

// The amount of IOPS for the disk.
func (o LookupVolumeResultOutput) Iops() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVolumeResult) int { return v.Iops }).(pulumi.IntOutput)
}

// The ARN for the KMS encryption key.
func (o LookupVolumeResultOutput) KmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.KmsKeyId }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) MostRecent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.MostRecent }).(pulumi.BoolPtrOutput)
}

// (Optional) Specifies whether Amazon EBS Multi-Attach is enabled.
func (o LookupVolumeResultOutput) MultiAttachEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVolumeResult) bool { return v.MultiAttachEnabled }).(pulumi.BoolOutput)
}

// The Amazon Resource Name (ARN) of the Outpost.
func (o LookupVolumeResultOutput) OutpostArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.OutpostArn }).(pulumi.StringOutput)
}

// The size of the drive in GiBs.
func (o LookupVolumeResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVolumeResult) int { return v.Size }).(pulumi.IntOutput)
}

// The snapshotId the EBS volume is based off.
func (o LookupVolumeResultOutput) SnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.SnapshotId }).(pulumi.StringOutput)
}

// A map of tags for the resource.
func (o LookupVolumeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVolumeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The throughput that the volume supports, in MiB/s.
func (o LookupVolumeResultOutput) Throughput() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVolumeResult) int { return v.Throughput }).(pulumi.IntOutput)
}

// The volume ID (e.g., vol-59fcb34e).
func (o LookupVolumeResultOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.VolumeId }).(pulumi.StringOutput)
}

// The type of EBS volume.
func (o LookupVolumeResultOutput) VolumeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.VolumeType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeResultOutput{})
}
