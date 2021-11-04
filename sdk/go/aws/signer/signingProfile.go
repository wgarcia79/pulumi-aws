// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Signer Signing Profile. A signing profile contains information about the code signing configuration parameters that can be used by a given code signing user.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/signer"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := signer.NewSigningProfile(ctx, "testSp", &signer.SigningProfileArgs{
// 			PlatformId: pulumi.String("AWSLambda-SHA384-ECDSA"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = signer.NewSigningProfile(ctx, "prodSp", &signer.SigningProfileArgs{
// 			NamePrefix: pulumi.String("prod_sp_"),
// 			PlatformId: pulumi.String("AWSLambda-SHA384-ECDSA"),
// 			SignatureValidityPeriod: &signer.SigningProfileSignatureValidityPeriodArgs{
// 				Type:  pulumi.String("YEARS"),
// 				Value: pulumi.Int(5),
// 			},
// 			Tags: pulumi.StringMap{
// 				"tag1": pulumi.String("value1"),
// 				"tag2": pulumi.String("value2"),
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
// Signer signing profiles can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:signer/signingProfile:SigningProfile test_signer_signing_profile test_sp_DdW3Mk1foYL88fajut4mTVFGpuwfd4ACO6ANL0D1uIj7lrn8adK
// ```
type SigningProfile struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) for the signing profile.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A unique signing profile name. By default generated by the provider. Signing profile names are immutable and cannot be reused after canceled.
	Name pulumi.StringOutput `pulumi:"name"`
	// A signing profile name prefix. The provider will generate a unique suffix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// A human-readable name for the signing platform associated with the signing profile.
	PlatformDisplayName pulumi.StringOutput `pulumi:"platformDisplayName"`
	// The ID of the platform that is used by the target signing profile.
	PlatformId pulumi.StringOutput `pulumi:"platformId"`
	// Revocation information for a signing profile.
	RevocationRecords SigningProfileRevocationRecordArrayOutput `pulumi:"revocationRecords"`
	// The validity period for a signing job.
	SignatureValidityPeriod SigningProfileSignatureValidityPeriodOutput `pulumi:"signatureValidityPeriod"`
	// The status of the target signing profile.
	Status pulumi.StringOutput `pulumi:"status"`
	// A list of tags associated with the signing profile. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The current version of the signing profile.
	Version pulumi.StringOutput `pulumi:"version"`
	// The signing profile ARN, including the profile version.
	VersionArn pulumi.StringOutput `pulumi:"versionArn"`
}

// NewSigningProfile registers a new resource with the given unique name, arguments, and options.
func NewSigningProfile(ctx *pulumi.Context,
	name string, args *SigningProfileArgs, opts ...pulumi.ResourceOption) (*SigningProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PlatformId == nil {
		return nil, errors.New("invalid value for required argument 'PlatformId'")
	}
	var resource SigningProfile
	err := ctx.RegisterResource("aws:signer/signingProfile:SigningProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSigningProfile gets an existing SigningProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSigningProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SigningProfileState, opts ...pulumi.ResourceOption) (*SigningProfile, error) {
	var resource SigningProfile
	err := ctx.ReadResource("aws:signer/signingProfile:SigningProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SigningProfile resources.
type signingProfileState struct {
	// The Amazon Resource Name (ARN) for the signing profile.
	Arn *string `pulumi:"arn"`
	// A unique signing profile name. By default generated by the provider. Signing profile names are immutable and cannot be reused after canceled.
	Name *string `pulumi:"name"`
	// A signing profile name prefix. The provider will generate a unique suffix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// A human-readable name for the signing platform associated with the signing profile.
	PlatformDisplayName *string `pulumi:"platformDisplayName"`
	// The ID of the platform that is used by the target signing profile.
	PlatformId *string `pulumi:"platformId"`
	// Revocation information for a signing profile.
	RevocationRecords []SigningProfileRevocationRecord `pulumi:"revocationRecords"`
	// The validity period for a signing job.
	SignatureValidityPeriod *SigningProfileSignatureValidityPeriod `pulumi:"signatureValidityPeriod"`
	// The status of the target signing profile.
	Status *string `pulumi:"status"`
	// A list of tags associated with the signing profile. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The current version of the signing profile.
	Version *string `pulumi:"version"`
	// The signing profile ARN, including the profile version.
	VersionArn *string `pulumi:"versionArn"`
}

type SigningProfileState struct {
	// The Amazon Resource Name (ARN) for the signing profile.
	Arn pulumi.StringPtrInput
	// A unique signing profile name. By default generated by the provider. Signing profile names are immutable and cannot be reused after canceled.
	Name pulumi.StringPtrInput
	// A signing profile name prefix. The provider will generate a unique suffix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// A human-readable name for the signing platform associated with the signing profile.
	PlatformDisplayName pulumi.StringPtrInput
	// The ID of the platform that is used by the target signing profile.
	PlatformId pulumi.StringPtrInput
	// Revocation information for a signing profile.
	RevocationRecords SigningProfileRevocationRecordArrayInput
	// The validity period for a signing job.
	SignatureValidityPeriod SigningProfileSignatureValidityPeriodPtrInput
	// The status of the target signing profile.
	Status pulumi.StringPtrInput
	// A list of tags associated with the signing profile. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The current version of the signing profile.
	Version pulumi.StringPtrInput
	// The signing profile ARN, including the profile version.
	VersionArn pulumi.StringPtrInput
}

func (SigningProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*signingProfileState)(nil)).Elem()
}

type signingProfileArgs struct {
	// A unique signing profile name. By default generated by the provider. Signing profile names are immutable and cannot be reused after canceled.
	Name *string `pulumi:"name"`
	// A signing profile name prefix. The provider will generate a unique suffix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The ID of the platform that is used by the target signing profile.
	PlatformId string `pulumi:"platformId"`
	// The validity period for a signing job.
	SignatureValidityPeriod *SigningProfileSignatureValidityPeriod `pulumi:"signatureValidityPeriod"`
	// A list of tags associated with the signing profile. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SigningProfile resource.
type SigningProfileArgs struct {
	// A unique signing profile name. By default generated by the provider. Signing profile names are immutable and cannot be reused after canceled.
	Name pulumi.StringPtrInput
	// A signing profile name prefix. The provider will generate a unique suffix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The ID of the platform that is used by the target signing profile.
	PlatformId pulumi.StringInput
	// The validity period for a signing job.
	SignatureValidityPeriod SigningProfileSignatureValidityPeriodPtrInput
	// A list of tags associated with the signing profile. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (SigningProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signingProfileArgs)(nil)).Elem()
}

type SigningProfileInput interface {
	pulumi.Input

	ToSigningProfileOutput() SigningProfileOutput
	ToSigningProfileOutputWithContext(ctx context.Context) SigningProfileOutput
}

func (*SigningProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*SigningProfile)(nil))
}

func (i *SigningProfile) ToSigningProfileOutput() SigningProfileOutput {
	return i.ToSigningProfileOutputWithContext(context.Background())
}

func (i *SigningProfile) ToSigningProfileOutputWithContext(ctx context.Context) SigningProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SigningProfileOutput)
}

func (i *SigningProfile) ToSigningProfilePtrOutput() SigningProfilePtrOutput {
	return i.ToSigningProfilePtrOutputWithContext(context.Background())
}

func (i *SigningProfile) ToSigningProfilePtrOutputWithContext(ctx context.Context) SigningProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SigningProfilePtrOutput)
}

type SigningProfilePtrInput interface {
	pulumi.Input

	ToSigningProfilePtrOutput() SigningProfilePtrOutput
	ToSigningProfilePtrOutputWithContext(ctx context.Context) SigningProfilePtrOutput
}

type signingProfilePtrType SigningProfileArgs

func (*signingProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SigningProfile)(nil))
}

func (i *signingProfilePtrType) ToSigningProfilePtrOutput() SigningProfilePtrOutput {
	return i.ToSigningProfilePtrOutputWithContext(context.Background())
}

func (i *signingProfilePtrType) ToSigningProfilePtrOutputWithContext(ctx context.Context) SigningProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SigningProfilePtrOutput)
}

// SigningProfileArrayInput is an input type that accepts SigningProfileArray and SigningProfileArrayOutput values.
// You can construct a concrete instance of `SigningProfileArrayInput` via:
//
//          SigningProfileArray{ SigningProfileArgs{...} }
type SigningProfileArrayInput interface {
	pulumi.Input

	ToSigningProfileArrayOutput() SigningProfileArrayOutput
	ToSigningProfileArrayOutputWithContext(context.Context) SigningProfileArrayOutput
}

type SigningProfileArray []SigningProfileInput

func (SigningProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SigningProfile)(nil)).Elem()
}

func (i SigningProfileArray) ToSigningProfileArrayOutput() SigningProfileArrayOutput {
	return i.ToSigningProfileArrayOutputWithContext(context.Background())
}

func (i SigningProfileArray) ToSigningProfileArrayOutputWithContext(ctx context.Context) SigningProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SigningProfileArrayOutput)
}

// SigningProfileMapInput is an input type that accepts SigningProfileMap and SigningProfileMapOutput values.
// You can construct a concrete instance of `SigningProfileMapInput` via:
//
//          SigningProfileMap{ "key": SigningProfileArgs{...} }
type SigningProfileMapInput interface {
	pulumi.Input

	ToSigningProfileMapOutput() SigningProfileMapOutput
	ToSigningProfileMapOutputWithContext(context.Context) SigningProfileMapOutput
}

type SigningProfileMap map[string]SigningProfileInput

func (SigningProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SigningProfile)(nil)).Elem()
}

func (i SigningProfileMap) ToSigningProfileMapOutput() SigningProfileMapOutput {
	return i.ToSigningProfileMapOutputWithContext(context.Background())
}

func (i SigningProfileMap) ToSigningProfileMapOutputWithContext(ctx context.Context) SigningProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SigningProfileMapOutput)
}

type SigningProfileOutput struct{ *pulumi.OutputState }

func (SigningProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SigningProfile)(nil))
}

func (o SigningProfileOutput) ToSigningProfileOutput() SigningProfileOutput {
	return o
}

func (o SigningProfileOutput) ToSigningProfileOutputWithContext(ctx context.Context) SigningProfileOutput {
	return o
}

func (o SigningProfileOutput) ToSigningProfilePtrOutput() SigningProfilePtrOutput {
	return o.ToSigningProfilePtrOutputWithContext(context.Background())
}

func (o SigningProfileOutput) ToSigningProfilePtrOutputWithContext(ctx context.Context) SigningProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SigningProfile) *SigningProfile {
		return &v
	}).(SigningProfilePtrOutput)
}

type SigningProfilePtrOutput struct{ *pulumi.OutputState }

func (SigningProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SigningProfile)(nil))
}

func (o SigningProfilePtrOutput) ToSigningProfilePtrOutput() SigningProfilePtrOutput {
	return o
}

func (o SigningProfilePtrOutput) ToSigningProfilePtrOutputWithContext(ctx context.Context) SigningProfilePtrOutput {
	return o
}

func (o SigningProfilePtrOutput) Elem() SigningProfileOutput {
	return o.ApplyT(func(v *SigningProfile) SigningProfile {
		if v != nil {
			return *v
		}
		var ret SigningProfile
		return ret
	}).(SigningProfileOutput)
}

type SigningProfileArrayOutput struct{ *pulumi.OutputState }

func (SigningProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SigningProfile)(nil))
}

func (o SigningProfileArrayOutput) ToSigningProfileArrayOutput() SigningProfileArrayOutput {
	return o
}

func (o SigningProfileArrayOutput) ToSigningProfileArrayOutputWithContext(ctx context.Context) SigningProfileArrayOutput {
	return o
}

func (o SigningProfileArrayOutput) Index(i pulumi.IntInput) SigningProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SigningProfile {
		return vs[0].([]SigningProfile)[vs[1].(int)]
	}).(SigningProfileOutput)
}

type SigningProfileMapOutput struct{ *pulumi.OutputState }

func (SigningProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SigningProfile)(nil))
}

func (o SigningProfileMapOutput) ToSigningProfileMapOutput() SigningProfileMapOutput {
	return o
}

func (o SigningProfileMapOutput) ToSigningProfileMapOutputWithContext(ctx context.Context) SigningProfileMapOutput {
	return o
}

func (o SigningProfileMapOutput) MapIndex(k pulumi.StringInput) SigningProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SigningProfile {
		return vs[0].(map[string]SigningProfile)[vs[1].(string)]
	}).(SigningProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SigningProfileInput)(nil)).Elem(), &SigningProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*SigningProfilePtrInput)(nil)).Elem(), &SigningProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*SigningProfileArrayInput)(nil)).Elem(), SigningProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SigningProfileMapInput)(nil)).Elem(), SigningProfileMap{})
	pulumi.RegisterOutputType(SigningProfileOutput{})
	pulumi.RegisterOutputType(SigningProfilePtrOutput{})
	pulumi.RegisterOutputType(SigningProfileArrayOutput{})
	pulumi.RegisterOutputType(SigningProfileMapOutput{})
}
