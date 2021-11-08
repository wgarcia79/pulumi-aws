// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an Amazon CloudFront origin access identity.
//
// For information about CloudFront distributions, see the
// [Amazon CloudFront Developer Guide](http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Introduction.html). For more information on generating
// origin access identities, see
// [Using an Origin Access Identity to Restrict Access to Your Amazon S3 Content][2].
//
// ## Example Usage
//
// The following example below creates a CloudFront origin access identity.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudfront"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudfront.NewOriginAccessIdentity(ctx, "example", &cloudfront.OriginAccessIdentityArgs{
// 			Comment: pulumi.String("Some comment"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Using With CloudFront
//
// Normally, when referencing an origin access identity in CloudFront, you need to
// prefix the ID with the `origin-access-identity/cloudfront/` special path.
// The `cloudfrontAccessIdentityPath` allows this to be circumvented.
// The below snippet demonstrates use with the `s3OriginConfig` structure for the
// `cloudfront.Distribution` resource:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudfront"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudfront.NewDistribution(ctx, "example", &cloudfront.DistributionArgs{
// 			Origins: cloudfront.DistributionOriginArray{
// 				&cloudfront.DistributionOriginArgs{
// 					S3OriginConfig: &cloudfront.DistributionOriginS3OriginConfigArgs{
// 						OriginAccessIdentity: pulumi.Any(aws_cloudfront_origin_access_identity.Example.Cloudfront_access_identity_path),
// 					},
// 				},
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
// Cloudfront Origin Access Identities can be imported using the `id`, e.g.,
//
// ```sh
//  $ pulumi import aws:cloudfront/originAccessIdentity:OriginAccessIdentity origin_access E74FTE3AEXAMPLE
// ```
type OriginAccessIdentity struct {
	pulumi.CustomResourceState

	// Internal value used by CloudFront to allow future
	// updates to the origin access identity.
	CallerReference pulumi.StringOutput `pulumi:"callerReference"`
	// A shortcut to the full path for the
	// origin access identity to use in CloudFront, see below.
	CloudfrontAccessIdentityPath pulumi.StringOutput `pulumi:"cloudfrontAccessIdentityPath"`
	// An optional comment for the origin access identity.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// The current version of the origin access identity's information.
	// For example: `E2QWRUHAPOMQZL`.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A pre-generated ARN for use in S3 bucket policies (see below).
	// Example: `arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity
	// E2QWRUHAPOMQZL`.
	IamArn pulumi.StringOutput `pulumi:"iamArn"`
	// The Amazon S3 canonical user ID for the origin
	// access identity, which you use when giving the origin access identity read
	// permission to an object in Amazon S3.
	S3CanonicalUserId pulumi.StringOutput `pulumi:"s3CanonicalUserId"`
}

// NewOriginAccessIdentity registers a new resource with the given unique name, arguments, and options.
func NewOriginAccessIdentity(ctx *pulumi.Context,
	name string, args *OriginAccessIdentityArgs, opts ...pulumi.ResourceOption) (*OriginAccessIdentity, error) {
	if args == nil {
		args = &OriginAccessIdentityArgs{}
	}

	var resource OriginAccessIdentity
	err := ctx.RegisterResource("aws:cloudfront/originAccessIdentity:OriginAccessIdentity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOriginAccessIdentity gets an existing OriginAccessIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOriginAccessIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OriginAccessIdentityState, opts ...pulumi.ResourceOption) (*OriginAccessIdentity, error) {
	var resource OriginAccessIdentity
	err := ctx.ReadResource("aws:cloudfront/originAccessIdentity:OriginAccessIdentity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OriginAccessIdentity resources.
type originAccessIdentityState struct {
	// Internal value used by CloudFront to allow future
	// updates to the origin access identity.
	CallerReference *string `pulumi:"callerReference"`
	// A shortcut to the full path for the
	// origin access identity to use in CloudFront, see below.
	CloudfrontAccessIdentityPath *string `pulumi:"cloudfrontAccessIdentityPath"`
	// An optional comment for the origin access identity.
	Comment *string `pulumi:"comment"`
	// The current version of the origin access identity's information.
	// For example: `E2QWRUHAPOMQZL`.
	Etag *string `pulumi:"etag"`
	// A pre-generated ARN for use in S3 bucket policies (see below).
	// Example: `arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity
	// E2QWRUHAPOMQZL`.
	IamArn *string `pulumi:"iamArn"`
	// The Amazon S3 canonical user ID for the origin
	// access identity, which you use when giving the origin access identity read
	// permission to an object in Amazon S3.
	S3CanonicalUserId *string `pulumi:"s3CanonicalUserId"`
}

type OriginAccessIdentityState struct {
	// Internal value used by CloudFront to allow future
	// updates to the origin access identity.
	CallerReference pulumi.StringPtrInput
	// A shortcut to the full path for the
	// origin access identity to use in CloudFront, see below.
	CloudfrontAccessIdentityPath pulumi.StringPtrInput
	// An optional comment for the origin access identity.
	Comment pulumi.StringPtrInput
	// The current version of the origin access identity's information.
	// For example: `E2QWRUHAPOMQZL`.
	Etag pulumi.StringPtrInput
	// A pre-generated ARN for use in S3 bucket policies (see below).
	// Example: `arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity
	// E2QWRUHAPOMQZL`.
	IamArn pulumi.StringPtrInput
	// The Amazon S3 canonical user ID for the origin
	// access identity, which you use when giving the origin access identity read
	// permission to an object in Amazon S3.
	S3CanonicalUserId pulumi.StringPtrInput
}

func (OriginAccessIdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*originAccessIdentityState)(nil)).Elem()
}

type originAccessIdentityArgs struct {
	// An optional comment for the origin access identity.
	Comment *string `pulumi:"comment"`
}

// The set of arguments for constructing a OriginAccessIdentity resource.
type OriginAccessIdentityArgs struct {
	// An optional comment for the origin access identity.
	Comment pulumi.StringPtrInput
}

func (OriginAccessIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*originAccessIdentityArgs)(nil)).Elem()
}

type OriginAccessIdentityInput interface {
	pulumi.Input

	ToOriginAccessIdentityOutput() OriginAccessIdentityOutput
	ToOriginAccessIdentityOutputWithContext(ctx context.Context) OriginAccessIdentityOutput
}

func (*OriginAccessIdentity) ElementType() reflect.Type {
	return reflect.TypeOf((*OriginAccessIdentity)(nil))
}

func (i *OriginAccessIdentity) ToOriginAccessIdentityOutput() OriginAccessIdentityOutput {
	return i.ToOriginAccessIdentityOutputWithContext(context.Background())
}

func (i *OriginAccessIdentity) ToOriginAccessIdentityOutputWithContext(ctx context.Context) OriginAccessIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginAccessIdentityOutput)
}

func (i *OriginAccessIdentity) ToOriginAccessIdentityPtrOutput() OriginAccessIdentityPtrOutput {
	return i.ToOriginAccessIdentityPtrOutputWithContext(context.Background())
}

func (i *OriginAccessIdentity) ToOriginAccessIdentityPtrOutputWithContext(ctx context.Context) OriginAccessIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginAccessIdentityPtrOutput)
}

type OriginAccessIdentityPtrInput interface {
	pulumi.Input

	ToOriginAccessIdentityPtrOutput() OriginAccessIdentityPtrOutput
	ToOriginAccessIdentityPtrOutputWithContext(ctx context.Context) OriginAccessIdentityPtrOutput
}

type originAccessIdentityPtrType OriginAccessIdentityArgs

func (*originAccessIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OriginAccessIdentity)(nil))
}

func (i *originAccessIdentityPtrType) ToOriginAccessIdentityPtrOutput() OriginAccessIdentityPtrOutput {
	return i.ToOriginAccessIdentityPtrOutputWithContext(context.Background())
}

func (i *originAccessIdentityPtrType) ToOriginAccessIdentityPtrOutputWithContext(ctx context.Context) OriginAccessIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginAccessIdentityPtrOutput)
}

// OriginAccessIdentityArrayInput is an input type that accepts OriginAccessIdentityArray and OriginAccessIdentityArrayOutput values.
// You can construct a concrete instance of `OriginAccessIdentityArrayInput` via:
//
//          OriginAccessIdentityArray{ OriginAccessIdentityArgs{...} }
type OriginAccessIdentityArrayInput interface {
	pulumi.Input

	ToOriginAccessIdentityArrayOutput() OriginAccessIdentityArrayOutput
	ToOriginAccessIdentityArrayOutputWithContext(context.Context) OriginAccessIdentityArrayOutput
}

type OriginAccessIdentityArray []OriginAccessIdentityInput

func (OriginAccessIdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OriginAccessIdentity)(nil)).Elem()
}

func (i OriginAccessIdentityArray) ToOriginAccessIdentityArrayOutput() OriginAccessIdentityArrayOutput {
	return i.ToOriginAccessIdentityArrayOutputWithContext(context.Background())
}

func (i OriginAccessIdentityArray) ToOriginAccessIdentityArrayOutputWithContext(ctx context.Context) OriginAccessIdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginAccessIdentityArrayOutput)
}

// OriginAccessIdentityMapInput is an input type that accepts OriginAccessIdentityMap and OriginAccessIdentityMapOutput values.
// You can construct a concrete instance of `OriginAccessIdentityMapInput` via:
//
//          OriginAccessIdentityMap{ "key": OriginAccessIdentityArgs{...} }
type OriginAccessIdentityMapInput interface {
	pulumi.Input

	ToOriginAccessIdentityMapOutput() OriginAccessIdentityMapOutput
	ToOriginAccessIdentityMapOutputWithContext(context.Context) OriginAccessIdentityMapOutput
}

type OriginAccessIdentityMap map[string]OriginAccessIdentityInput

func (OriginAccessIdentityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OriginAccessIdentity)(nil)).Elem()
}

func (i OriginAccessIdentityMap) ToOriginAccessIdentityMapOutput() OriginAccessIdentityMapOutput {
	return i.ToOriginAccessIdentityMapOutputWithContext(context.Background())
}

func (i OriginAccessIdentityMap) ToOriginAccessIdentityMapOutputWithContext(ctx context.Context) OriginAccessIdentityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginAccessIdentityMapOutput)
}

type OriginAccessIdentityOutput struct{ *pulumi.OutputState }

func (OriginAccessIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OriginAccessIdentity)(nil))
}

func (o OriginAccessIdentityOutput) ToOriginAccessIdentityOutput() OriginAccessIdentityOutput {
	return o
}

func (o OriginAccessIdentityOutput) ToOriginAccessIdentityOutputWithContext(ctx context.Context) OriginAccessIdentityOutput {
	return o
}

func (o OriginAccessIdentityOutput) ToOriginAccessIdentityPtrOutput() OriginAccessIdentityPtrOutput {
	return o.ToOriginAccessIdentityPtrOutputWithContext(context.Background())
}

func (o OriginAccessIdentityOutput) ToOriginAccessIdentityPtrOutputWithContext(ctx context.Context) OriginAccessIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OriginAccessIdentity) *OriginAccessIdentity {
		return &v
	}).(OriginAccessIdentityPtrOutput)
}

type OriginAccessIdentityPtrOutput struct{ *pulumi.OutputState }

func (OriginAccessIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OriginAccessIdentity)(nil))
}

func (o OriginAccessIdentityPtrOutput) ToOriginAccessIdentityPtrOutput() OriginAccessIdentityPtrOutput {
	return o
}

func (o OriginAccessIdentityPtrOutput) ToOriginAccessIdentityPtrOutputWithContext(ctx context.Context) OriginAccessIdentityPtrOutput {
	return o
}

func (o OriginAccessIdentityPtrOutput) Elem() OriginAccessIdentityOutput {
	return o.ApplyT(func(v *OriginAccessIdentity) OriginAccessIdentity {
		if v != nil {
			return *v
		}
		var ret OriginAccessIdentity
		return ret
	}).(OriginAccessIdentityOutput)
}

type OriginAccessIdentityArrayOutput struct{ *pulumi.OutputState }

func (OriginAccessIdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OriginAccessIdentity)(nil))
}

func (o OriginAccessIdentityArrayOutput) ToOriginAccessIdentityArrayOutput() OriginAccessIdentityArrayOutput {
	return o
}

func (o OriginAccessIdentityArrayOutput) ToOriginAccessIdentityArrayOutputWithContext(ctx context.Context) OriginAccessIdentityArrayOutput {
	return o
}

func (o OriginAccessIdentityArrayOutput) Index(i pulumi.IntInput) OriginAccessIdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OriginAccessIdentity {
		return vs[0].([]OriginAccessIdentity)[vs[1].(int)]
	}).(OriginAccessIdentityOutput)
}

type OriginAccessIdentityMapOutput struct{ *pulumi.OutputState }

func (OriginAccessIdentityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OriginAccessIdentity)(nil))
}

func (o OriginAccessIdentityMapOutput) ToOriginAccessIdentityMapOutput() OriginAccessIdentityMapOutput {
	return o
}

func (o OriginAccessIdentityMapOutput) ToOriginAccessIdentityMapOutputWithContext(ctx context.Context) OriginAccessIdentityMapOutput {
	return o
}

func (o OriginAccessIdentityMapOutput) MapIndex(k pulumi.StringInput) OriginAccessIdentityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OriginAccessIdentity {
		return vs[0].(map[string]OriginAccessIdentity)[vs[1].(string)]
	}).(OriginAccessIdentityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OriginAccessIdentityInput)(nil)).Elem(), &OriginAccessIdentity{})
	pulumi.RegisterInputType(reflect.TypeOf((*OriginAccessIdentityPtrInput)(nil)).Elem(), &OriginAccessIdentity{})
	pulumi.RegisterInputType(reflect.TypeOf((*OriginAccessIdentityArrayInput)(nil)).Elem(), OriginAccessIdentityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OriginAccessIdentityMapInput)(nil)).Elem(), OriginAccessIdentityMap{})
	pulumi.RegisterOutputType(OriginAccessIdentityOutput{})
	pulumi.RegisterOutputType(OriginAccessIdentityPtrOutput{})
	pulumi.RegisterOutputType(OriginAccessIdentityArrayOutput{})
	pulumi.RegisterOutputType(OriginAccessIdentityMapOutput{})
}
