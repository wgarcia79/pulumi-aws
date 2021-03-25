// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attaches a policy to an S3 bucket resource.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
// 	"encoding/json"
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bucket, err := s3.NewBucket(ctx, "bucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucketPolicy(ctx, "bucketPolicy", &s3.BucketPolicyArgs{
// 			Bucket: bucket.ID(),
// 			Policy: pulumi.All(bucket.Arn, bucket.Arn).ApplyT(func(_args []interface{}) (pulumi.String, error) {
// 				bucketArn := _args[0].(string)
// 				bucketArn1 := _args[1].(string)
// 				var _zero pulumi.String
// 				tmpJSON0, err := json.Marshal(map[string]interface{}{
// 					"Version": "2012-10-17",
// 					"Id":      "MYBUCKETPOLICY",
// 					"Statement": []map[string]interface{}{
// 						map[string]interface{}{
// 							"Sid":       "IPAllow",
// 							"Effect":    "Deny",
// 							"Principal": "*",
// 							"Action":    "s3:*",
// 							"Resource": []string{
// 								bucketArn,
// 								fmt.Sprintf("%v%v", bucketArn1, "/*"),
// 							},
// 							"Condition": map[string]interface{}{
// 								"IpAddress": map[string]interface{}{
// 									"aws:SourceIp": "8.8.8.8/32",
// 								},
// 							},
// 						},
// 					},
// 				})
// 				if err != nil {
// 					return _zero, err
// 				}
// 				json0 := string(tmpJSON0)
// 				return pulumi.String(json0), nil
// 			}).(pulumi.StringOutput),
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
// S3 bucket policies can be imported using the bucket name, e.g.
//
// ```sh
//  $ pulumi import aws:s3/bucketPolicy:BucketPolicy example my-bucket-name
// ```
type BucketPolicy struct {
	pulumi.CustomResourceState

	// The name of the bucket to which to apply the policy.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The text of the policy. Note: Bucket policies are limited to 20 KB in size.
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewBucketPolicy registers a new resource with the given unique name, arguments, and options.
func NewBucketPolicy(ctx *pulumi.Context,
	name string, args *BucketPolicyArgs, opts ...pulumi.ResourceOption) (*BucketPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	var resource BucketPolicy
	err := ctx.RegisterResource("aws:s3/bucketPolicy:BucketPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketPolicy gets an existing BucketPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketPolicyState, opts ...pulumi.ResourceOption) (*BucketPolicy, error) {
	var resource BucketPolicy
	err := ctx.ReadResource("aws:s3/bucketPolicy:BucketPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketPolicy resources.
type bucketPolicyState struct {
	// The name of the bucket to which to apply the policy.
	Bucket *string `pulumi:"bucket"`
	// The text of the policy. Note: Bucket policies are limited to 20 KB in size.
	Policy *string `pulumi:"policy"`
}

type BucketPolicyState struct {
	// The name of the bucket to which to apply the policy.
	Bucket pulumi.StringPtrInput
	// The text of the policy. Note: Bucket policies are limited to 20 KB in size.
	Policy pulumi.StringPtrInput
}

func (BucketPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketPolicyState)(nil)).Elem()
}

type bucketPolicyArgs struct {
	// The name of the bucket to which to apply the policy.
	Bucket string `pulumi:"bucket"`
	// The text of the policy. Note: Bucket policies are limited to 20 KB in size.
	Policy interface{} `pulumi:"policy"`
}

// The set of arguments for constructing a BucketPolicy resource.
type BucketPolicyArgs struct {
	// The name of the bucket to which to apply the policy.
	Bucket pulumi.StringInput
	// The text of the policy. Note: Bucket policies are limited to 20 KB in size.
	Policy pulumi.Input
}

func (BucketPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketPolicyArgs)(nil)).Elem()
}

type BucketPolicyInput interface {
	pulumi.Input

	ToBucketPolicyOutput() BucketPolicyOutput
	ToBucketPolicyOutputWithContext(ctx context.Context) BucketPolicyOutput
}

func (*BucketPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketPolicy)(nil))
}

func (i *BucketPolicy) ToBucketPolicyOutput() BucketPolicyOutput {
	return i.ToBucketPolicyOutputWithContext(context.Background())
}

func (i *BucketPolicy) ToBucketPolicyOutputWithContext(ctx context.Context) BucketPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPolicyOutput)
}

func (i *BucketPolicy) ToBucketPolicyPtrOutput() BucketPolicyPtrOutput {
	return i.ToBucketPolicyPtrOutputWithContext(context.Background())
}

func (i *BucketPolicy) ToBucketPolicyPtrOutputWithContext(ctx context.Context) BucketPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPolicyPtrOutput)
}

type BucketPolicyPtrInput interface {
	pulumi.Input

	ToBucketPolicyPtrOutput() BucketPolicyPtrOutput
	ToBucketPolicyPtrOutputWithContext(ctx context.Context) BucketPolicyPtrOutput
}

type bucketPolicyPtrType BucketPolicyArgs

func (*bucketPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketPolicy)(nil))
}

func (i *bucketPolicyPtrType) ToBucketPolicyPtrOutput() BucketPolicyPtrOutput {
	return i.ToBucketPolicyPtrOutputWithContext(context.Background())
}

func (i *bucketPolicyPtrType) ToBucketPolicyPtrOutputWithContext(ctx context.Context) BucketPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPolicyPtrOutput)
}

// BucketPolicyArrayInput is an input type that accepts BucketPolicyArray and BucketPolicyArrayOutput values.
// You can construct a concrete instance of `BucketPolicyArrayInput` via:
//
//          BucketPolicyArray{ BucketPolicyArgs{...} }
type BucketPolicyArrayInput interface {
	pulumi.Input

	ToBucketPolicyArrayOutput() BucketPolicyArrayOutput
	ToBucketPolicyArrayOutputWithContext(context.Context) BucketPolicyArrayOutput
}

type BucketPolicyArray []BucketPolicyInput

func (BucketPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*BucketPolicy)(nil))
}

func (i BucketPolicyArray) ToBucketPolicyArrayOutput() BucketPolicyArrayOutput {
	return i.ToBucketPolicyArrayOutputWithContext(context.Background())
}

func (i BucketPolicyArray) ToBucketPolicyArrayOutputWithContext(ctx context.Context) BucketPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPolicyArrayOutput)
}

// BucketPolicyMapInput is an input type that accepts BucketPolicyMap and BucketPolicyMapOutput values.
// You can construct a concrete instance of `BucketPolicyMapInput` via:
//
//          BucketPolicyMap{ "key": BucketPolicyArgs{...} }
type BucketPolicyMapInput interface {
	pulumi.Input

	ToBucketPolicyMapOutput() BucketPolicyMapOutput
	ToBucketPolicyMapOutputWithContext(context.Context) BucketPolicyMapOutput
}

type BucketPolicyMap map[string]BucketPolicyInput

func (BucketPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*BucketPolicy)(nil))
}

func (i BucketPolicyMap) ToBucketPolicyMapOutput() BucketPolicyMapOutput {
	return i.ToBucketPolicyMapOutputWithContext(context.Background())
}

func (i BucketPolicyMap) ToBucketPolicyMapOutputWithContext(ctx context.Context) BucketPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPolicyMapOutput)
}

type BucketPolicyOutput struct {
	*pulumi.OutputState
}

func (BucketPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketPolicy)(nil))
}

func (o BucketPolicyOutput) ToBucketPolicyOutput() BucketPolicyOutput {
	return o
}

func (o BucketPolicyOutput) ToBucketPolicyOutputWithContext(ctx context.Context) BucketPolicyOutput {
	return o
}

func (o BucketPolicyOutput) ToBucketPolicyPtrOutput() BucketPolicyPtrOutput {
	return o.ToBucketPolicyPtrOutputWithContext(context.Background())
}

func (o BucketPolicyOutput) ToBucketPolicyPtrOutputWithContext(ctx context.Context) BucketPolicyPtrOutput {
	return o.ApplyT(func(v BucketPolicy) *BucketPolicy {
		return &v
	}).(BucketPolicyPtrOutput)
}

type BucketPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (BucketPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketPolicy)(nil))
}

func (o BucketPolicyPtrOutput) ToBucketPolicyPtrOutput() BucketPolicyPtrOutput {
	return o
}

func (o BucketPolicyPtrOutput) ToBucketPolicyPtrOutputWithContext(ctx context.Context) BucketPolicyPtrOutput {
	return o
}

type BucketPolicyArrayOutput struct{ *pulumi.OutputState }

func (BucketPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BucketPolicy)(nil))
}

func (o BucketPolicyArrayOutput) ToBucketPolicyArrayOutput() BucketPolicyArrayOutput {
	return o
}

func (o BucketPolicyArrayOutput) ToBucketPolicyArrayOutputWithContext(ctx context.Context) BucketPolicyArrayOutput {
	return o
}

func (o BucketPolicyArrayOutput) Index(i pulumi.IntInput) BucketPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BucketPolicy {
		return vs[0].([]BucketPolicy)[vs[1].(int)]
	}).(BucketPolicyOutput)
}

type BucketPolicyMapOutput struct{ *pulumi.OutputState }

func (BucketPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BucketPolicy)(nil))
}

func (o BucketPolicyMapOutput) ToBucketPolicyMapOutput() BucketPolicyMapOutput {
	return o
}

func (o BucketPolicyMapOutput) ToBucketPolicyMapOutputWithContext(ctx context.Context) BucketPolicyMapOutput {
	return o
}

func (o BucketPolicyMapOutput) MapIndex(k pulumi.StringInput) BucketPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BucketPolicy {
		return vs[0].(map[string]BucketPolicy)[vs[1].(string)]
	}).(BucketPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(BucketPolicyOutput{})
	pulumi.RegisterOutputType(BucketPolicyPtrOutput{})
	pulumi.RegisterOutputType(BucketPolicyArrayOutput{})
	pulumi.RegisterOutputType(BucketPolicyMapOutput{})
}
