// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage an S3 Access Point.
//
// ## Example Usage
// ### AWS Partition Bucket
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleBucket, err := s3.NewBucket(ctx, "exampleBucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewAccessPoint(ctx, "exampleAccessPoint", &s3.AccessPointArgs{
// 			Bucket: exampleBucket.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### S3 on Outposts Bucket
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3control"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleBucket, err := s3control.NewBucket(ctx, "exampleBucket", &s3control.BucketArgs{
// 			Bucket: pulumi.String("example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewAccessPoint(ctx, "exampleAccessPoint", &s3.AccessPointArgs{
// 			Bucket: exampleBucket.Arn,
// 			VpcConfiguration: &s3.AccessPointVpcConfigurationArgs{
// 				VpcId: exampleVpc.ID(),
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
// For Access Points associated with an AWS Partition S3 Bucket, this resource can be imported using the `account_id` and `name` separated by a colon (`:`), e.g.
//
// ```sh
//  $ pulumi import aws:s3/accessPoint:AccessPoint example 123456789012:example
// ```
//
//  For Access Points associated with an S3 on Outposts Bucket, this resource can be imported using the Amazon Resource Name (ARN), e.g.
//
// ```sh
//  $ pulumi import aws:s3/accessPoint:AccessPoint example arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-1234567890123456/accesspoint/example
// ```
type AccessPoint struct {
	pulumi.CustomResourceState

	// The AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the provider.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// Amazon Resource Name (ARN) of the S3 Access Point.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of an AWS Partition S3 Bucket or the Amazon Resource Name (ARN) of S3 on Outposts Bucket that you want to associate this access point with.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The DNS domain name of the S3 Access Point in the format _`name`_-_`accountId`_.s3-accesspoint._region_.amazonaws.com.
	// Note: S3 access points only support secure access by HTTPS. HTTP isn't supported.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Indicates whether this access point currently has a policy that allows public access.
	HasPublicAccessPolicy pulumi.BoolOutput `pulumi:"hasPublicAccessPolicy"`
	// The name you want to assign to this access point.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates whether this access point allows access from the public Internet. Values are `VPC` (the access point doesn't allow access from the public Internet) and `Internet` (the access point allows access from the public Internet, subject to the access point and bucket access policies).
	NetworkOrigin pulumi.StringOutput `pulumi:"networkOrigin"`
	// A valid JSON document that specifies the policy that you want to apply to this access point.
	Policy pulumi.StringPtrOutput `pulumi:"policy"`
	// Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
	PublicAccessBlockConfiguration AccessPointPublicAccessBlockConfigurationPtrOutput `pulumi:"publicAccessBlockConfiguration"`
	// Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
	VpcConfiguration AccessPointVpcConfigurationPtrOutput `pulumi:"vpcConfiguration"`
}

// NewAccessPoint registers a new resource with the given unique name, arguments, and options.
func NewAccessPoint(ctx *pulumi.Context,
	name string, args *AccessPointArgs, opts ...pulumi.ResourceOption) (*AccessPoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	var resource AccessPoint
	err := ctx.RegisterResource("aws:s3/accessPoint:AccessPoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPoint gets an existing AccessPoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPointState, opts ...pulumi.ResourceOption) (*AccessPoint, error) {
	var resource AccessPoint
	err := ctx.ReadResource("aws:s3/accessPoint:AccessPoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPoint resources.
type accessPointState struct {
	// The AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the provider.
	AccountId *string `pulumi:"accountId"`
	// Amazon Resource Name (ARN) of the S3 Access Point.
	Arn *string `pulumi:"arn"`
	// The name of an AWS Partition S3 Bucket or the Amazon Resource Name (ARN) of S3 on Outposts Bucket that you want to associate this access point with.
	Bucket *string `pulumi:"bucket"`
	// The DNS domain name of the S3 Access Point in the format _`name`_-_`accountId`_.s3-accesspoint._region_.amazonaws.com.
	// Note: S3 access points only support secure access by HTTPS. HTTP isn't supported.
	DomainName *string `pulumi:"domainName"`
	// Indicates whether this access point currently has a policy that allows public access.
	HasPublicAccessPolicy *bool `pulumi:"hasPublicAccessPolicy"`
	// The name you want to assign to this access point.
	Name *string `pulumi:"name"`
	// Indicates whether this access point allows access from the public Internet. Values are `VPC` (the access point doesn't allow access from the public Internet) and `Internet` (the access point allows access from the public Internet, subject to the access point and bucket access policies).
	NetworkOrigin *string `pulumi:"networkOrigin"`
	// A valid JSON document that specifies the policy that you want to apply to this access point.
	Policy *string `pulumi:"policy"`
	// Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
	PublicAccessBlockConfiguration *AccessPointPublicAccessBlockConfiguration `pulumi:"publicAccessBlockConfiguration"`
	// Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
	VpcConfiguration *AccessPointVpcConfiguration `pulumi:"vpcConfiguration"`
}

type AccessPointState struct {
	// The AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the provider.
	AccountId pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the S3 Access Point.
	Arn pulumi.StringPtrInput
	// The name of an AWS Partition S3 Bucket or the Amazon Resource Name (ARN) of S3 on Outposts Bucket that you want to associate this access point with.
	Bucket pulumi.StringPtrInput
	// The DNS domain name of the S3 Access Point in the format _`name`_-_`accountId`_.s3-accesspoint._region_.amazonaws.com.
	// Note: S3 access points only support secure access by HTTPS. HTTP isn't supported.
	DomainName pulumi.StringPtrInput
	// Indicates whether this access point currently has a policy that allows public access.
	HasPublicAccessPolicy pulumi.BoolPtrInput
	// The name you want to assign to this access point.
	Name pulumi.StringPtrInput
	// Indicates whether this access point allows access from the public Internet. Values are `VPC` (the access point doesn't allow access from the public Internet) and `Internet` (the access point allows access from the public Internet, subject to the access point and bucket access policies).
	NetworkOrigin pulumi.StringPtrInput
	// A valid JSON document that specifies the policy that you want to apply to this access point.
	Policy pulumi.StringPtrInput
	// Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
	PublicAccessBlockConfiguration AccessPointPublicAccessBlockConfigurationPtrInput
	// Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
	VpcConfiguration AccessPointVpcConfigurationPtrInput
}

func (AccessPointState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPointState)(nil)).Elem()
}

type accessPointArgs struct {
	// The AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the provider.
	AccountId *string `pulumi:"accountId"`
	// The name of an AWS Partition S3 Bucket or the Amazon Resource Name (ARN) of S3 on Outposts Bucket that you want to associate this access point with.
	Bucket string `pulumi:"bucket"`
	// The name you want to assign to this access point.
	Name *string `pulumi:"name"`
	// A valid JSON document that specifies the policy that you want to apply to this access point.
	Policy *string `pulumi:"policy"`
	// Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
	PublicAccessBlockConfiguration *AccessPointPublicAccessBlockConfiguration `pulumi:"publicAccessBlockConfiguration"`
	// Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
	VpcConfiguration *AccessPointVpcConfiguration `pulumi:"vpcConfiguration"`
}

// The set of arguments for constructing a AccessPoint resource.
type AccessPointArgs struct {
	// The AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the provider.
	AccountId pulumi.StringPtrInput
	// The name of an AWS Partition S3 Bucket or the Amazon Resource Name (ARN) of S3 on Outposts Bucket that you want to associate this access point with.
	Bucket pulumi.StringInput
	// The name you want to assign to this access point.
	Name pulumi.StringPtrInput
	// A valid JSON document that specifies the policy that you want to apply to this access point.
	Policy pulumi.StringPtrInput
	// Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
	PublicAccessBlockConfiguration AccessPointPublicAccessBlockConfigurationPtrInput
	// Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
	VpcConfiguration AccessPointVpcConfigurationPtrInput
}

func (AccessPointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPointArgs)(nil)).Elem()
}

type AccessPointInput interface {
	pulumi.Input

	ToAccessPointOutput() AccessPointOutput
	ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput
}

func (*AccessPoint) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPoint)(nil))
}

func (i *AccessPoint) ToAccessPointOutput() AccessPointOutput {
	return i.ToAccessPointOutputWithContext(context.Background())
}

func (i *AccessPoint) ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPointOutput)
}

func (i *AccessPoint) ToAccessPointPtrOutput() AccessPointPtrOutput {
	return i.ToAccessPointPtrOutputWithContext(context.Background())
}

func (i *AccessPoint) ToAccessPointPtrOutputWithContext(ctx context.Context) AccessPointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPointPtrOutput)
}

type AccessPointPtrInput interface {
	pulumi.Input

	ToAccessPointPtrOutput() AccessPointPtrOutput
	ToAccessPointPtrOutputWithContext(ctx context.Context) AccessPointPtrOutput
}

type accessPointPtrType AccessPointArgs

func (*accessPointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPoint)(nil))
}

func (i *accessPointPtrType) ToAccessPointPtrOutput() AccessPointPtrOutput {
	return i.ToAccessPointPtrOutputWithContext(context.Background())
}

func (i *accessPointPtrType) ToAccessPointPtrOutputWithContext(ctx context.Context) AccessPointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPointPtrOutput)
}

// AccessPointArrayInput is an input type that accepts AccessPointArray and AccessPointArrayOutput values.
// You can construct a concrete instance of `AccessPointArrayInput` via:
//
//          AccessPointArray{ AccessPointArgs{...} }
type AccessPointArrayInput interface {
	pulumi.Input

	ToAccessPointArrayOutput() AccessPointArrayOutput
	ToAccessPointArrayOutputWithContext(context.Context) AccessPointArrayOutput
}

type AccessPointArray []AccessPointInput

func (AccessPointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessPoint)(nil)).Elem()
}

func (i AccessPointArray) ToAccessPointArrayOutput() AccessPointArrayOutput {
	return i.ToAccessPointArrayOutputWithContext(context.Background())
}

func (i AccessPointArray) ToAccessPointArrayOutputWithContext(ctx context.Context) AccessPointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPointArrayOutput)
}

// AccessPointMapInput is an input type that accepts AccessPointMap and AccessPointMapOutput values.
// You can construct a concrete instance of `AccessPointMapInput` via:
//
//          AccessPointMap{ "key": AccessPointArgs{...} }
type AccessPointMapInput interface {
	pulumi.Input

	ToAccessPointMapOutput() AccessPointMapOutput
	ToAccessPointMapOutputWithContext(context.Context) AccessPointMapOutput
}

type AccessPointMap map[string]AccessPointInput

func (AccessPointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessPoint)(nil)).Elem()
}

func (i AccessPointMap) ToAccessPointMapOutput() AccessPointMapOutput {
	return i.ToAccessPointMapOutputWithContext(context.Background())
}

func (i AccessPointMap) ToAccessPointMapOutputWithContext(ctx context.Context) AccessPointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPointMapOutput)
}

type AccessPointOutput struct {
	*pulumi.OutputState
}

func (AccessPointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPoint)(nil))
}

func (o AccessPointOutput) ToAccessPointOutput() AccessPointOutput {
	return o
}

func (o AccessPointOutput) ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput {
	return o
}

func (o AccessPointOutput) ToAccessPointPtrOutput() AccessPointPtrOutput {
	return o.ToAccessPointPtrOutputWithContext(context.Background())
}

func (o AccessPointOutput) ToAccessPointPtrOutputWithContext(ctx context.Context) AccessPointPtrOutput {
	return o.ApplyT(func(v AccessPoint) *AccessPoint {
		return &v
	}).(AccessPointPtrOutput)
}

type AccessPointPtrOutput struct {
	*pulumi.OutputState
}

func (AccessPointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPoint)(nil))
}

func (o AccessPointPtrOutput) ToAccessPointPtrOutput() AccessPointPtrOutput {
	return o
}

func (o AccessPointPtrOutput) ToAccessPointPtrOutputWithContext(ctx context.Context) AccessPointPtrOutput {
	return o
}

type AccessPointArrayOutput struct{ *pulumi.OutputState }

func (AccessPointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessPoint)(nil))
}

func (o AccessPointArrayOutput) ToAccessPointArrayOutput() AccessPointArrayOutput {
	return o
}

func (o AccessPointArrayOutput) ToAccessPointArrayOutputWithContext(ctx context.Context) AccessPointArrayOutput {
	return o
}

func (o AccessPointArrayOutput) Index(i pulumi.IntInput) AccessPointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessPoint {
		return vs[0].([]AccessPoint)[vs[1].(int)]
	}).(AccessPointOutput)
}

type AccessPointMapOutput struct{ *pulumi.OutputState }

func (AccessPointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AccessPoint)(nil))
}

func (o AccessPointMapOutput) ToAccessPointMapOutput() AccessPointMapOutput {
	return o
}

func (o AccessPointMapOutput) ToAccessPointMapOutputWithContext(ctx context.Context) AccessPointMapOutput {
	return o
}

func (o AccessPointMapOutput) MapIndex(k pulumi.StringInput) AccessPointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AccessPoint {
		return vs[0].(map[string]AccessPoint)[vs[1].(string)]
	}).(AccessPointOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessPointOutput{})
	pulumi.RegisterOutputType(AccessPointPtrOutput{})
	pulumi.RegisterOutputType(AccessPointArrayOutput{})
	pulumi.RegisterOutputType(AccessPointMapOutput{})
}
