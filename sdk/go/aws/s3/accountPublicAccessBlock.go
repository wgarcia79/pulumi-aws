// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages S3 account-level Public Access Block configuration. For more information about these settings, see the [AWS S3 Block Public Access documentation](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html).
//
// > **NOTE:** Each AWS account may only have one S3 Public Access Block configuration. Multiple configurations of the resource against the same AWS account will cause a perpetual difference.
//
// > Advanced usage: To use a custom API endpoint for this resource, use the `s3control` endpoint provider configuration, not the `s3` endpoint provider configuration.
//
// ## Example Usage
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
// 		_, err := s3.NewAccountPublicAccessBlock(ctx, "example", &s3.AccountPublicAccessBlockArgs{
// 			BlockPublicAcls:   pulumi.Bool(true),
// 			BlockPublicPolicy: pulumi.Bool(true),
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
// `aws_s3_account_public_access_block` can be imported by using the AWS account ID, e.g.
//
// ```sh
//  $ pulumi import aws:s3/accountPublicAccessBlock:AccountPublicAccessBlock example 123456789012
// ```
type AccountPublicAccessBlock struct {
	pulumi.CustomResourceState

	// AWS account ID to configure. Defaults to automatically determined account ID of the this provider AWS provider.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
	// * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
	// * PUT Object calls will fail if the request includes an object ACL.
	BlockPublicAcls pulumi.BoolPtrOutput `pulumi:"blockPublicAcls"`
	// Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
	// * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	BlockPublicPolicy pulumi.BoolPtrOutput `pulumi:"blockPublicPolicy"`
	// Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
	// * Ignore all public ACLs on buckets in this account and any objects that they contain.
	IgnorePublicAcls pulumi.BoolPtrOutput `pulumi:"ignorePublicAcls"`
	// Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
	// * Only the bucket owner and AWS Services can access buckets with public policies.
	RestrictPublicBuckets pulumi.BoolPtrOutput `pulumi:"restrictPublicBuckets"`
}

// NewAccountPublicAccessBlock registers a new resource with the given unique name, arguments, and options.
func NewAccountPublicAccessBlock(ctx *pulumi.Context,
	name string, args *AccountPublicAccessBlockArgs, opts ...pulumi.ResourceOption) (*AccountPublicAccessBlock, error) {
	if args == nil {
		args = &AccountPublicAccessBlockArgs{}
	}

	var resource AccountPublicAccessBlock
	err := ctx.RegisterResource("aws:s3/accountPublicAccessBlock:AccountPublicAccessBlock", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountPublicAccessBlock gets an existing AccountPublicAccessBlock resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountPublicAccessBlock(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountPublicAccessBlockState, opts ...pulumi.ResourceOption) (*AccountPublicAccessBlock, error) {
	var resource AccountPublicAccessBlock
	err := ctx.ReadResource("aws:s3/accountPublicAccessBlock:AccountPublicAccessBlock", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountPublicAccessBlock resources.
type accountPublicAccessBlockState struct {
	// AWS account ID to configure. Defaults to automatically determined account ID of the this provider AWS provider.
	AccountId *string `pulumi:"accountId"`
	// Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
	// * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
	// * PUT Object calls will fail if the request includes an object ACL.
	BlockPublicAcls *bool `pulumi:"blockPublicAcls"`
	// Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
	// * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	BlockPublicPolicy *bool `pulumi:"blockPublicPolicy"`
	// Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
	// * Ignore all public ACLs on buckets in this account and any objects that they contain.
	IgnorePublicAcls *bool `pulumi:"ignorePublicAcls"`
	// Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
	// * Only the bucket owner and AWS Services can access buckets with public policies.
	RestrictPublicBuckets *bool `pulumi:"restrictPublicBuckets"`
}

type AccountPublicAccessBlockState struct {
	// AWS account ID to configure. Defaults to automatically determined account ID of the this provider AWS provider.
	AccountId pulumi.StringPtrInput
	// Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
	// * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
	// * PUT Object calls will fail if the request includes an object ACL.
	BlockPublicAcls pulumi.BoolPtrInput
	// Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
	// * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	BlockPublicPolicy pulumi.BoolPtrInput
	// Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
	// * Ignore all public ACLs on buckets in this account and any objects that they contain.
	IgnorePublicAcls pulumi.BoolPtrInput
	// Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
	// * Only the bucket owner and AWS Services can access buckets with public policies.
	RestrictPublicBuckets pulumi.BoolPtrInput
}

func (AccountPublicAccessBlockState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountPublicAccessBlockState)(nil)).Elem()
}

type accountPublicAccessBlockArgs struct {
	// AWS account ID to configure. Defaults to automatically determined account ID of the this provider AWS provider.
	AccountId *string `pulumi:"accountId"`
	// Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
	// * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
	// * PUT Object calls will fail if the request includes an object ACL.
	BlockPublicAcls *bool `pulumi:"blockPublicAcls"`
	// Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
	// * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	BlockPublicPolicy *bool `pulumi:"blockPublicPolicy"`
	// Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
	// * Ignore all public ACLs on buckets in this account and any objects that they contain.
	IgnorePublicAcls *bool `pulumi:"ignorePublicAcls"`
	// Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
	// * Only the bucket owner and AWS Services can access buckets with public policies.
	RestrictPublicBuckets *bool `pulumi:"restrictPublicBuckets"`
}

// The set of arguments for constructing a AccountPublicAccessBlock resource.
type AccountPublicAccessBlockArgs struct {
	// AWS account ID to configure. Defaults to automatically determined account ID of the this provider AWS provider.
	AccountId pulumi.StringPtrInput
	// Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
	// * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
	// * PUT Object calls will fail if the request includes an object ACL.
	BlockPublicAcls pulumi.BoolPtrInput
	// Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
	// * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	BlockPublicPolicy pulumi.BoolPtrInput
	// Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
	// * Ignore all public ACLs on buckets in this account and any objects that they contain.
	IgnorePublicAcls pulumi.BoolPtrInput
	// Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
	// * Only the bucket owner and AWS Services can access buckets with public policies.
	RestrictPublicBuckets pulumi.BoolPtrInput
}

func (AccountPublicAccessBlockArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountPublicAccessBlockArgs)(nil)).Elem()
}

type AccountPublicAccessBlockInput interface {
	pulumi.Input

	ToAccountPublicAccessBlockOutput() AccountPublicAccessBlockOutput
	ToAccountPublicAccessBlockOutputWithContext(ctx context.Context) AccountPublicAccessBlockOutput
}

func (*AccountPublicAccessBlock) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountPublicAccessBlock)(nil))
}

func (i *AccountPublicAccessBlock) ToAccountPublicAccessBlockOutput() AccountPublicAccessBlockOutput {
	return i.ToAccountPublicAccessBlockOutputWithContext(context.Background())
}

func (i *AccountPublicAccessBlock) ToAccountPublicAccessBlockOutputWithContext(ctx context.Context) AccountPublicAccessBlockOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPublicAccessBlockOutput)
}

func (i *AccountPublicAccessBlock) ToAccountPublicAccessBlockPtrOutput() AccountPublicAccessBlockPtrOutput {
	return i.ToAccountPublicAccessBlockPtrOutputWithContext(context.Background())
}

func (i *AccountPublicAccessBlock) ToAccountPublicAccessBlockPtrOutputWithContext(ctx context.Context) AccountPublicAccessBlockPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPublicAccessBlockPtrOutput)
}

type AccountPublicAccessBlockPtrInput interface {
	pulumi.Input

	ToAccountPublicAccessBlockPtrOutput() AccountPublicAccessBlockPtrOutput
	ToAccountPublicAccessBlockPtrOutputWithContext(ctx context.Context) AccountPublicAccessBlockPtrOutput
}

type accountPublicAccessBlockPtrType AccountPublicAccessBlockArgs

func (*accountPublicAccessBlockPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountPublicAccessBlock)(nil))
}

func (i *accountPublicAccessBlockPtrType) ToAccountPublicAccessBlockPtrOutput() AccountPublicAccessBlockPtrOutput {
	return i.ToAccountPublicAccessBlockPtrOutputWithContext(context.Background())
}

func (i *accountPublicAccessBlockPtrType) ToAccountPublicAccessBlockPtrOutputWithContext(ctx context.Context) AccountPublicAccessBlockPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPublicAccessBlockPtrOutput)
}

// AccountPublicAccessBlockArrayInput is an input type that accepts AccountPublicAccessBlockArray and AccountPublicAccessBlockArrayOutput values.
// You can construct a concrete instance of `AccountPublicAccessBlockArrayInput` via:
//
//          AccountPublicAccessBlockArray{ AccountPublicAccessBlockArgs{...} }
type AccountPublicAccessBlockArrayInput interface {
	pulumi.Input

	ToAccountPublicAccessBlockArrayOutput() AccountPublicAccessBlockArrayOutput
	ToAccountPublicAccessBlockArrayOutputWithContext(context.Context) AccountPublicAccessBlockArrayOutput
}

type AccountPublicAccessBlockArray []AccountPublicAccessBlockInput

func (AccountPublicAccessBlockArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountPublicAccessBlock)(nil)).Elem()
}

func (i AccountPublicAccessBlockArray) ToAccountPublicAccessBlockArrayOutput() AccountPublicAccessBlockArrayOutput {
	return i.ToAccountPublicAccessBlockArrayOutputWithContext(context.Background())
}

func (i AccountPublicAccessBlockArray) ToAccountPublicAccessBlockArrayOutputWithContext(ctx context.Context) AccountPublicAccessBlockArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPublicAccessBlockArrayOutput)
}

// AccountPublicAccessBlockMapInput is an input type that accepts AccountPublicAccessBlockMap and AccountPublicAccessBlockMapOutput values.
// You can construct a concrete instance of `AccountPublicAccessBlockMapInput` via:
//
//          AccountPublicAccessBlockMap{ "key": AccountPublicAccessBlockArgs{...} }
type AccountPublicAccessBlockMapInput interface {
	pulumi.Input

	ToAccountPublicAccessBlockMapOutput() AccountPublicAccessBlockMapOutput
	ToAccountPublicAccessBlockMapOutputWithContext(context.Context) AccountPublicAccessBlockMapOutput
}

type AccountPublicAccessBlockMap map[string]AccountPublicAccessBlockInput

func (AccountPublicAccessBlockMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountPublicAccessBlock)(nil)).Elem()
}

func (i AccountPublicAccessBlockMap) ToAccountPublicAccessBlockMapOutput() AccountPublicAccessBlockMapOutput {
	return i.ToAccountPublicAccessBlockMapOutputWithContext(context.Background())
}

func (i AccountPublicAccessBlockMap) ToAccountPublicAccessBlockMapOutputWithContext(ctx context.Context) AccountPublicAccessBlockMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPublicAccessBlockMapOutput)
}

type AccountPublicAccessBlockOutput struct{ *pulumi.OutputState }

func (AccountPublicAccessBlockOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountPublicAccessBlock)(nil))
}

func (o AccountPublicAccessBlockOutput) ToAccountPublicAccessBlockOutput() AccountPublicAccessBlockOutput {
	return o
}

func (o AccountPublicAccessBlockOutput) ToAccountPublicAccessBlockOutputWithContext(ctx context.Context) AccountPublicAccessBlockOutput {
	return o
}

func (o AccountPublicAccessBlockOutput) ToAccountPublicAccessBlockPtrOutput() AccountPublicAccessBlockPtrOutput {
	return o.ToAccountPublicAccessBlockPtrOutputWithContext(context.Background())
}

func (o AccountPublicAccessBlockOutput) ToAccountPublicAccessBlockPtrOutputWithContext(ctx context.Context) AccountPublicAccessBlockPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountPublicAccessBlock) *AccountPublicAccessBlock {
		return &v
	}).(AccountPublicAccessBlockPtrOutput)
}

type AccountPublicAccessBlockPtrOutput struct{ *pulumi.OutputState }

func (AccountPublicAccessBlockPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountPublicAccessBlock)(nil))
}

func (o AccountPublicAccessBlockPtrOutput) ToAccountPublicAccessBlockPtrOutput() AccountPublicAccessBlockPtrOutput {
	return o
}

func (o AccountPublicAccessBlockPtrOutput) ToAccountPublicAccessBlockPtrOutputWithContext(ctx context.Context) AccountPublicAccessBlockPtrOutput {
	return o
}

func (o AccountPublicAccessBlockPtrOutput) Elem() AccountPublicAccessBlockOutput {
	return o.ApplyT(func(v *AccountPublicAccessBlock) AccountPublicAccessBlock {
		if v != nil {
			return *v
		}
		var ret AccountPublicAccessBlock
		return ret
	}).(AccountPublicAccessBlockOutput)
}

type AccountPublicAccessBlockArrayOutput struct{ *pulumi.OutputState }

func (AccountPublicAccessBlockArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccountPublicAccessBlock)(nil))
}

func (o AccountPublicAccessBlockArrayOutput) ToAccountPublicAccessBlockArrayOutput() AccountPublicAccessBlockArrayOutput {
	return o
}

func (o AccountPublicAccessBlockArrayOutput) ToAccountPublicAccessBlockArrayOutputWithContext(ctx context.Context) AccountPublicAccessBlockArrayOutput {
	return o
}

func (o AccountPublicAccessBlockArrayOutput) Index(i pulumi.IntInput) AccountPublicAccessBlockOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccountPublicAccessBlock {
		return vs[0].([]AccountPublicAccessBlock)[vs[1].(int)]
	}).(AccountPublicAccessBlockOutput)
}

type AccountPublicAccessBlockMapOutput struct{ *pulumi.OutputState }

func (AccountPublicAccessBlockMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AccountPublicAccessBlock)(nil))
}

func (o AccountPublicAccessBlockMapOutput) ToAccountPublicAccessBlockMapOutput() AccountPublicAccessBlockMapOutput {
	return o
}

func (o AccountPublicAccessBlockMapOutput) ToAccountPublicAccessBlockMapOutputWithContext(ctx context.Context) AccountPublicAccessBlockMapOutput {
	return o
}

func (o AccountPublicAccessBlockMapOutput) MapIndex(k pulumi.StringInput) AccountPublicAccessBlockOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AccountPublicAccessBlock {
		return vs[0].(map[string]AccountPublicAccessBlock)[vs[1].(string)]
	}).(AccountPublicAccessBlockOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPublicAccessBlockInput)(nil)).Elem(), &AccountPublicAccessBlock{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPublicAccessBlockPtrInput)(nil)).Elem(), &AccountPublicAccessBlock{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPublicAccessBlockArrayInput)(nil)).Elem(), AccountPublicAccessBlockArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPublicAccessBlockMapInput)(nil)).Elem(), AccountPublicAccessBlockMap{})
	pulumi.RegisterOutputType(AccountPublicAccessBlockOutput{})
	pulumi.RegisterOutputType(AccountPublicAccessBlockPtrOutput{})
	pulumi.RegisterOutputType(AccountPublicAccessBlockArrayOutput{})
	pulumi.RegisterOutputType(AccountPublicAccessBlockMapOutput{})
}
