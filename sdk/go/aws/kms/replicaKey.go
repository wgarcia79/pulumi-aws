// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a KMS multi-Region replica key.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/kms"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/providers"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := providers.Newaws(ctx, "primary", &providers.awsArgs{
// 			Region: "us-east-1",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		primaryKey, err := kms.NewKey(ctx, "primaryKey", &kms.KeyArgs{
// 			Description:          pulumi.String("Multi-Region primary key"),
// 			DeletionWindowInDays: pulumi.Int(30),
// 			MultiRegion:          pulumi.Bool(true),
// 		}, pulumi.Provider(aws.Primary))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kms.NewReplicaKey(ctx, "replica", &kms.ReplicaKeyArgs{
// 			Description:          pulumi.String("Multi-Region replica key"),
// 			DeletionWindowInDays: pulumi.Int(7),
// 			PrimaryKeyArn:        primaryKey.Arn,
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
// KMS multi-Region replica keys can be imported using the `id`, e.g.,
//
// ```sh
//  $ pulumi import aws:kms/replicaKey:ReplicaKey example 1234abcd-12ab-34cd-56ef-1234567890ab
// ```
type ReplicaKey struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the replica key. The key ARNs of related multi-Region keys differ only in the Region value.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	// The default value is `false`.
	BypassPolicyLockoutSafetyCheck pulumi.BoolPtrOutput `pulumi:"bypassPolicyLockoutSafetyCheck"`
	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	DeletionWindowInDays pulumi.IntPtrOutput `pulumi:"deletionWindowInDays"`
	// A description of the KMS key.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. The default value is `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The key ID of the replica key. Related multi-Region keys have the same key ID.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// A Boolean value that specifies whether key rotation is enabled. This is a shared property of multi-Region keys.
	KeyRotationEnabled pulumi.BoolOutput `pulumi:"keyRotationEnabled"`
	// The type of key material in the KMS key. This is a shared property of multi-Region keys.
	KeySpec pulumi.StringOutput `pulumi:"keySpec"`
	// The [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. This is a shared property of multi-Region keys.
	KeyUsage pulumi.StringOutput `pulumi:"keyUsage"`
	Policy   pulumi.StringOutput `pulumi:"policy"`
	// The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
	PrimaryKeyArn pulumi.StringOutput    `pulumi:"primaryKeyArn"`
	Tags          pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewReplicaKey registers a new resource with the given unique name, arguments, and options.
func NewReplicaKey(ctx *pulumi.Context,
	name string, args *ReplicaKeyArgs, opts ...pulumi.ResourceOption) (*ReplicaKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrimaryKeyArn == nil {
		return nil, errors.New("invalid value for required argument 'PrimaryKeyArn'")
	}
	var resource ReplicaKey
	err := ctx.RegisterResource("aws:kms/replicaKey:ReplicaKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicaKey gets an existing ReplicaKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicaKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicaKeyState, opts ...pulumi.ResourceOption) (*ReplicaKey, error) {
	var resource ReplicaKey
	err := ctx.ReadResource("aws:kms/replicaKey:ReplicaKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicaKey resources.
type replicaKeyState struct {
	// The Amazon Resource Name (ARN) of the replica key. The key ARNs of related multi-Region keys differ only in the Region value.
	Arn *string `pulumi:"arn"`
	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	// The default value is `false`.
	BypassPolicyLockoutSafetyCheck *bool `pulumi:"bypassPolicyLockoutSafetyCheck"`
	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	DeletionWindowInDays *int `pulumi:"deletionWindowInDays"`
	// A description of the KMS key.
	Description *string `pulumi:"description"`
	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. The default value is `true`.
	Enabled *bool `pulumi:"enabled"`
	// The key ID of the replica key. Related multi-Region keys have the same key ID.
	KeyId *string `pulumi:"keyId"`
	// A Boolean value that specifies whether key rotation is enabled. This is a shared property of multi-Region keys.
	KeyRotationEnabled *bool `pulumi:"keyRotationEnabled"`
	// The type of key material in the KMS key. This is a shared property of multi-Region keys.
	KeySpec *string `pulumi:"keySpec"`
	// The [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. This is a shared property of multi-Region keys.
	KeyUsage *string `pulumi:"keyUsage"`
	Policy   *string `pulumi:"policy"`
	// The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
	PrimaryKeyArn *string           `pulumi:"primaryKeyArn"`
	Tags          map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ReplicaKeyState struct {
	// The Amazon Resource Name (ARN) of the replica key. The key ARNs of related multi-Region keys differ only in the Region value.
	Arn pulumi.StringPtrInput
	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	// The default value is `false`.
	BypassPolicyLockoutSafetyCheck pulumi.BoolPtrInput
	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	DeletionWindowInDays pulumi.IntPtrInput
	// A description of the KMS key.
	Description pulumi.StringPtrInput
	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. The default value is `true`.
	Enabled pulumi.BoolPtrInput
	// The key ID of the replica key. Related multi-Region keys have the same key ID.
	KeyId pulumi.StringPtrInput
	// A Boolean value that specifies whether key rotation is enabled. This is a shared property of multi-Region keys.
	KeyRotationEnabled pulumi.BoolPtrInput
	// The type of key material in the KMS key. This is a shared property of multi-Region keys.
	KeySpec pulumi.StringPtrInput
	// The [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. This is a shared property of multi-Region keys.
	KeyUsage pulumi.StringPtrInput
	Policy   pulumi.StringPtrInput
	// The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
	PrimaryKeyArn pulumi.StringPtrInput
	Tags          pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (ReplicaKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicaKeyState)(nil)).Elem()
}

type replicaKeyArgs struct {
	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	// The default value is `false`.
	BypassPolicyLockoutSafetyCheck *bool `pulumi:"bypassPolicyLockoutSafetyCheck"`
	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	DeletionWindowInDays *int `pulumi:"deletionWindowInDays"`
	// A description of the KMS key.
	Description *string `pulumi:"description"`
	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. The default value is `true`.
	Enabled *bool   `pulumi:"enabled"`
	Policy  *string `pulumi:"policy"`
	// The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
	PrimaryKeyArn string            `pulumi:"primaryKeyArn"`
	Tags          map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ReplicaKey resource.
type ReplicaKeyArgs struct {
	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	// The default value is `false`.
	BypassPolicyLockoutSafetyCheck pulumi.BoolPtrInput
	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	DeletionWindowInDays pulumi.IntPtrInput
	// A description of the KMS key.
	Description pulumi.StringPtrInput
	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. The default value is `true`.
	Enabled pulumi.BoolPtrInput
	Policy  pulumi.StringPtrInput
	// The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
	PrimaryKeyArn pulumi.StringInput
	Tags          pulumi.StringMapInput
}

func (ReplicaKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicaKeyArgs)(nil)).Elem()
}

type ReplicaKeyInput interface {
	pulumi.Input

	ToReplicaKeyOutput() ReplicaKeyOutput
	ToReplicaKeyOutputWithContext(ctx context.Context) ReplicaKeyOutput
}

func (*ReplicaKey) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicaKey)(nil))
}

func (i *ReplicaKey) ToReplicaKeyOutput() ReplicaKeyOutput {
	return i.ToReplicaKeyOutputWithContext(context.Background())
}

func (i *ReplicaKey) ToReplicaKeyOutputWithContext(ctx context.Context) ReplicaKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaKeyOutput)
}

func (i *ReplicaKey) ToReplicaKeyPtrOutput() ReplicaKeyPtrOutput {
	return i.ToReplicaKeyPtrOutputWithContext(context.Background())
}

func (i *ReplicaKey) ToReplicaKeyPtrOutputWithContext(ctx context.Context) ReplicaKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaKeyPtrOutput)
}

type ReplicaKeyPtrInput interface {
	pulumi.Input

	ToReplicaKeyPtrOutput() ReplicaKeyPtrOutput
	ToReplicaKeyPtrOutputWithContext(ctx context.Context) ReplicaKeyPtrOutput
}

type replicaKeyPtrType ReplicaKeyArgs

func (*replicaKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicaKey)(nil))
}

func (i *replicaKeyPtrType) ToReplicaKeyPtrOutput() ReplicaKeyPtrOutput {
	return i.ToReplicaKeyPtrOutputWithContext(context.Background())
}

func (i *replicaKeyPtrType) ToReplicaKeyPtrOutputWithContext(ctx context.Context) ReplicaKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaKeyPtrOutput)
}

// ReplicaKeyArrayInput is an input type that accepts ReplicaKeyArray and ReplicaKeyArrayOutput values.
// You can construct a concrete instance of `ReplicaKeyArrayInput` via:
//
//          ReplicaKeyArray{ ReplicaKeyArgs{...} }
type ReplicaKeyArrayInput interface {
	pulumi.Input

	ToReplicaKeyArrayOutput() ReplicaKeyArrayOutput
	ToReplicaKeyArrayOutputWithContext(context.Context) ReplicaKeyArrayOutput
}

type ReplicaKeyArray []ReplicaKeyInput

func (ReplicaKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicaKey)(nil)).Elem()
}

func (i ReplicaKeyArray) ToReplicaKeyArrayOutput() ReplicaKeyArrayOutput {
	return i.ToReplicaKeyArrayOutputWithContext(context.Background())
}

func (i ReplicaKeyArray) ToReplicaKeyArrayOutputWithContext(ctx context.Context) ReplicaKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaKeyArrayOutput)
}

// ReplicaKeyMapInput is an input type that accepts ReplicaKeyMap and ReplicaKeyMapOutput values.
// You can construct a concrete instance of `ReplicaKeyMapInput` via:
//
//          ReplicaKeyMap{ "key": ReplicaKeyArgs{...} }
type ReplicaKeyMapInput interface {
	pulumi.Input

	ToReplicaKeyMapOutput() ReplicaKeyMapOutput
	ToReplicaKeyMapOutputWithContext(context.Context) ReplicaKeyMapOutput
}

type ReplicaKeyMap map[string]ReplicaKeyInput

func (ReplicaKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicaKey)(nil)).Elem()
}

func (i ReplicaKeyMap) ToReplicaKeyMapOutput() ReplicaKeyMapOutput {
	return i.ToReplicaKeyMapOutputWithContext(context.Background())
}

func (i ReplicaKeyMap) ToReplicaKeyMapOutputWithContext(ctx context.Context) ReplicaKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaKeyMapOutput)
}

type ReplicaKeyOutput struct{ *pulumi.OutputState }

func (ReplicaKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicaKey)(nil))
}

func (o ReplicaKeyOutput) ToReplicaKeyOutput() ReplicaKeyOutput {
	return o
}

func (o ReplicaKeyOutput) ToReplicaKeyOutputWithContext(ctx context.Context) ReplicaKeyOutput {
	return o
}

func (o ReplicaKeyOutput) ToReplicaKeyPtrOutput() ReplicaKeyPtrOutput {
	return o.ToReplicaKeyPtrOutputWithContext(context.Background())
}

func (o ReplicaKeyOutput) ToReplicaKeyPtrOutputWithContext(ctx context.Context) ReplicaKeyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReplicaKey) *ReplicaKey {
		return &v
	}).(ReplicaKeyPtrOutput)
}

type ReplicaKeyPtrOutput struct{ *pulumi.OutputState }

func (ReplicaKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicaKey)(nil))
}

func (o ReplicaKeyPtrOutput) ToReplicaKeyPtrOutput() ReplicaKeyPtrOutput {
	return o
}

func (o ReplicaKeyPtrOutput) ToReplicaKeyPtrOutputWithContext(ctx context.Context) ReplicaKeyPtrOutput {
	return o
}

func (o ReplicaKeyPtrOutput) Elem() ReplicaKeyOutput {
	return o.ApplyT(func(v *ReplicaKey) ReplicaKey {
		if v != nil {
			return *v
		}
		var ret ReplicaKey
		return ret
	}).(ReplicaKeyOutput)
}

type ReplicaKeyArrayOutput struct{ *pulumi.OutputState }

func (ReplicaKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReplicaKey)(nil))
}

func (o ReplicaKeyArrayOutput) ToReplicaKeyArrayOutput() ReplicaKeyArrayOutput {
	return o
}

func (o ReplicaKeyArrayOutput) ToReplicaKeyArrayOutputWithContext(ctx context.Context) ReplicaKeyArrayOutput {
	return o
}

func (o ReplicaKeyArrayOutput) Index(i pulumi.IntInput) ReplicaKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReplicaKey {
		return vs[0].([]ReplicaKey)[vs[1].(int)]
	}).(ReplicaKeyOutput)
}

type ReplicaKeyMapOutput struct{ *pulumi.OutputState }

func (ReplicaKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReplicaKey)(nil))
}

func (o ReplicaKeyMapOutput) ToReplicaKeyMapOutput() ReplicaKeyMapOutput {
	return o
}

func (o ReplicaKeyMapOutput) ToReplicaKeyMapOutputWithContext(ctx context.Context) ReplicaKeyMapOutput {
	return o
}

func (o ReplicaKeyMapOutput) MapIndex(k pulumi.StringInput) ReplicaKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReplicaKey {
		return vs[0].(map[string]ReplicaKey)[vs[1].(string)]
	}).(ReplicaKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicaKeyInput)(nil)).Elem(), &ReplicaKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicaKeyPtrInput)(nil)).Elem(), &ReplicaKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicaKeyArrayInput)(nil)).Elem(), ReplicaKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicaKeyMapInput)(nil)).Elem(), ReplicaKeyMap{})
	pulumi.RegisterOutputType(ReplicaKeyOutput{})
	pulumi.RegisterOutputType(ReplicaKeyPtrOutput{})
	pulumi.RegisterOutputType(ReplicaKeyArrayOutput{})
	pulumi.RegisterOutputType(ReplicaKeyMapOutput{})
}
