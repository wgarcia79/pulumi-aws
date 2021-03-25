// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Signer Signing Profile Permission. That is, a cross-account permission for a signing profile.
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
// 		prodSp, err := signer.NewSigningProfile(ctx, "prodSp", &signer.SigningProfileArgs{
// 			PlatformId: pulumi.String("AWSLambda-SHA384-ECDSA"),
// 			NamePrefix: pulumi.String("prod_sp_"),
// 			SignatureValidityPeriod: &signer.SigningProfileSignatureValidityPeriodArgs{
// 				Value: pulumi.Int(5),
// 				Type:  pulumi.String("YEARS"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"tag1": pulumi.String("value1"),
// 				"tag2": pulumi.String("value2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = signer.NewSigningProfilePermission(ctx, "spPermission1", &signer.SigningProfilePermissionArgs{
// 			ProfileName: prodSp.Name,
// 			Action:      pulumi.String("signer:StartSigningJob"),
// 			Principal:   pulumi.Any(_var.Aws_account),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = signer.NewSigningProfilePermission(ctx, "spPermission2", &signer.SigningProfilePermissionArgs{
// 			ProfileName: prodSp.Name,
// 			Action:      pulumi.String("signer:GetSigningProfile"),
// 			Principal:   pulumi.Any(_var.Aws_team_role_arn),
// 			StatementId: pulumi.String("ProdAccountStartSigningJob_StatementId"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = signer.NewSigningProfilePermission(ctx, "spPermission3", &signer.SigningProfilePermissionArgs{
// 			ProfileName:       prodSp.Name,
// 			Action:            pulumi.String("signer:RevokeSignature"),
// 			Principal:         pulumi.String("123456789012"),
// 			ProfileVersion:    prodSp.Version,
// 			StatementIdPrefix: pulumi.String("version-permission-"),
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
// Signer signing profile permission statements can be imported using profile_name/statement_id, e.g.
//
// ```sh
//  $ pulumi import aws:signer/signingProfilePermission:SigningProfilePermission test_signer_signing_profile_permission prod_profile_DdW3Mk1foYL88fajut4mTVFGpuwfd4ACO6ANL0D1uIj7lrn8adK/ProdAccountStartSigningJobStatementId
// ```
type SigningProfilePermission struct {
	pulumi.CustomResourceState

	// An AWS Signer action permitted as part of cross-account permissions. Valid values: `signer:StartSigningJob`, `signer:GetSigningProfile`, or `signer:RevokeSignature`.
	Action pulumi.StringOutput `pulumi:"action"`
	// The AWS principal to be granted a cross-account permission.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// Name of the signing profile to add the cross-account permissions.
	ProfileName pulumi.StringOutput `pulumi:"profileName"`
	// The signing profile version that a permission applies to.
	ProfileVersion pulumi.StringOutput `pulumi:"profileVersion"`
	// A unique statement identifier. By default generated by the provider.
	StatementId pulumi.StringOutput `pulumi:"statementId"`
	// A statement identifier prefix. The provider will generate a unique suffix. Conflicts with `statementId`.
	StatementIdPrefix pulumi.StringPtrOutput `pulumi:"statementIdPrefix"`
}

// NewSigningProfilePermission registers a new resource with the given unique name, arguments, and options.
func NewSigningProfilePermission(ctx *pulumi.Context,
	name string, args *SigningProfilePermissionArgs, opts ...pulumi.ResourceOption) (*SigningProfilePermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	var resource SigningProfilePermission
	err := ctx.RegisterResource("aws:signer/signingProfilePermission:SigningProfilePermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSigningProfilePermission gets an existing SigningProfilePermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSigningProfilePermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SigningProfilePermissionState, opts ...pulumi.ResourceOption) (*SigningProfilePermission, error) {
	var resource SigningProfilePermission
	err := ctx.ReadResource("aws:signer/signingProfilePermission:SigningProfilePermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SigningProfilePermission resources.
type signingProfilePermissionState struct {
	// An AWS Signer action permitted as part of cross-account permissions. Valid values: `signer:StartSigningJob`, `signer:GetSigningProfile`, or `signer:RevokeSignature`.
	Action *string `pulumi:"action"`
	// The AWS principal to be granted a cross-account permission.
	Principal *string `pulumi:"principal"`
	// Name of the signing profile to add the cross-account permissions.
	ProfileName *string `pulumi:"profileName"`
	// The signing profile version that a permission applies to.
	ProfileVersion *string `pulumi:"profileVersion"`
	// A unique statement identifier. By default generated by the provider.
	StatementId *string `pulumi:"statementId"`
	// A statement identifier prefix. The provider will generate a unique suffix. Conflicts with `statementId`.
	StatementIdPrefix *string `pulumi:"statementIdPrefix"`
}

type SigningProfilePermissionState struct {
	// An AWS Signer action permitted as part of cross-account permissions. Valid values: `signer:StartSigningJob`, `signer:GetSigningProfile`, or `signer:RevokeSignature`.
	Action pulumi.StringPtrInput
	// The AWS principal to be granted a cross-account permission.
	Principal pulumi.StringPtrInput
	// Name of the signing profile to add the cross-account permissions.
	ProfileName pulumi.StringPtrInput
	// The signing profile version that a permission applies to.
	ProfileVersion pulumi.StringPtrInput
	// A unique statement identifier. By default generated by the provider.
	StatementId pulumi.StringPtrInput
	// A statement identifier prefix. The provider will generate a unique suffix. Conflicts with `statementId`.
	StatementIdPrefix pulumi.StringPtrInput
}

func (SigningProfilePermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*signingProfilePermissionState)(nil)).Elem()
}

type signingProfilePermissionArgs struct {
	// An AWS Signer action permitted as part of cross-account permissions. Valid values: `signer:StartSigningJob`, `signer:GetSigningProfile`, or `signer:RevokeSignature`.
	Action string `pulumi:"action"`
	// The AWS principal to be granted a cross-account permission.
	Principal string `pulumi:"principal"`
	// Name of the signing profile to add the cross-account permissions.
	ProfileName string `pulumi:"profileName"`
	// The signing profile version that a permission applies to.
	ProfileVersion *string `pulumi:"profileVersion"`
	// A unique statement identifier. By default generated by the provider.
	StatementId *string `pulumi:"statementId"`
	// A statement identifier prefix. The provider will generate a unique suffix. Conflicts with `statementId`.
	StatementIdPrefix *string `pulumi:"statementIdPrefix"`
}

// The set of arguments for constructing a SigningProfilePermission resource.
type SigningProfilePermissionArgs struct {
	// An AWS Signer action permitted as part of cross-account permissions. Valid values: `signer:StartSigningJob`, `signer:GetSigningProfile`, or `signer:RevokeSignature`.
	Action pulumi.StringInput
	// The AWS principal to be granted a cross-account permission.
	Principal pulumi.StringInput
	// Name of the signing profile to add the cross-account permissions.
	ProfileName pulumi.StringInput
	// The signing profile version that a permission applies to.
	ProfileVersion pulumi.StringPtrInput
	// A unique statement identifier. By default generated by the provider.
	StatementId pulumi.StringPtrInput
	// A statement identifier prefix. The provider will generate a unique suffix. Conflicts with `statementId`.
	StatementIdPrefix pulumi.StringPtrInput
}

func (SigningProfilePermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signingProfilePermissionArgs)(nil)).Elem()
}

type SigningProfilePermissionInput interface {
	pulumi.Input

	ToSigningProfilePermissionOutput() SigningProfilePermissionOutput
	ToSigningProfilePermissionOutputWithContext(ctx context.Context) SigningProfilePermissionOutput
}

func (*SigningProfilePermission) ElementType() reflect.Type {
	return reflect.TypeOf((*SigningProfilePermission)(nil))
}

func (i *SigningProfilePermission) ToSigningProfilePermissionOutput() SigningProfilePermissionOutput {
	return i.ToSigningProfilePermissionOutputWithContext(context.Background())
}

func (i *SigningProfilePermission) ToSigningProfilePermissionOutputWithContext(ctx context.Context) SigningProfilePermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SigningProfilePermissionOutput)
}

func (i *SigningProfilePermission) ToSigningProfilePermissionPtrOutput() SigningProfilePermissionPtrOutput {
	return i.ToSigningProfilePermissionPtrOutputWithContext(context.Background())
}

func (i *SigningProfilePermission) ToSigningProfilePermissionPtrOutputWithContext(ctx context.Context) SigningProfilePermissionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SigningProfilePermissionPtrOutput)
}

type SigningProfilePermissionPtrInput interface {
	pulumi.Input

	ToSigningProfilePermissionPtrOutput() SigningProfilePermissionPtrOutput
	ToSigningProfilePermissionPtrOutputWithContext(ctx context.Context) SigningProfilePermissionPtrOutput
}

type signingProfilePermissionPtrType SigningProfilePermissionArgs

func (*signingProfilePermissionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SigningProfilePermission)(nil))
}

func (i *signingProfilePermissionPtrType) ToSigningProfilePermissionPtrOutput() SigningProfilePermissionPtrOutput {
	return i.ToSigningProfilePermissionPtrOutputWithContext(context.Background())
}

func (i *signingProfilePermissionPtrType) ToSigningProfilePermissionPtrOutputWithContext(ctx context.Context) SigningProfilePermissionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SigningProfilePermissionPtrOutput)
}

// SigningProfilePermissionArrayInput is an input type that accepts SigningProfilePermissionArray and SigningProfilePermissionArrayOutput values.
// You can construct a concrete instance of `SigningProfilePermissionArrayInput` via:
//
//          SigningProfilePermissionArray{ SigningProfilePermissionArgs{...} }
type SigningProfilePermissionArrayInput interface {
	pulumi.Input

	ToSigningProfilePermissionArrayOutput() SigningProfilePermissionArrayOutput
	ToSigningProfilePermissionArrayOutputWithContext(context.Context) SigningProfilePermissionArrayOutput
}

type SigningProfilePermissionArray []SigningProfilePermissionInput

func (SigningProfilePermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SigningProfilePermission)(nil))
}

func (i SigningProfilePermissionArray) ToSigningProfilePermissionArrayOutput() SigningProfilePermissionArrayOutput {
	return i.ToSigningProfilePermissionArrayOutputWithContext(context.Background())
}

func (i SigningProfilePermissionArray) ToSigningProfilePermissionArrayOutputWithContext(ctx context.Context) SigningProfilePermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SigningProfilePermissionArrayOutput)
}

// SigningProfilePermissionMapInput is an input type that accepts SigningProfilePermissionMap and SigningProfilePermissionMapOutput values.
// You can construct a concrete instance of `SigningProfilePermissionMapInput` via:
//
//          SigningProfilePermissionMap{ "key": SigningProfilePermissionArgs{...} }
type SigningProfilePermissionMapInput interface {
	pulumi.Input

	ToSigningProfilePermissionMapOutput() SigningProfilePermissionMapOutput
	ToSigningProfilePermissionMapOutputWithContext(context.Context) SigningProfilePermissionMapOutput
}

type SigningProfilePermissionMap map[string]SigningProfilePermissionInput

func (SigningProfilePermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SigningProfilePermission)(nil))
}

func (i SigningProfilePermissionMap) ToSigningProfilePermissionMapOutput() SigningProfilePermissionMapOutput {
	return i.ToSigningProfilePermissionMapOutputWithContext(context.Background())
}

func (i SigningProfilePermissionMap) ToSigningProfilePermissionMapOutputWithContext(ctx context.Context) SigningProfilePermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SigningProfilePermissionMapOutput)
}

type SigningProfilePermissionOutput struct {
	*pulumi.OutputState
}

func (SigningProfilePermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SigningProfilePermission)(nil))
}

func (o SigningProfilePermissionOutput) ToSigningProfilePermissionOutput() SigningProfilePermissionOutput {
	return o
}

func (o SigningProfilePermissionOutput) ToSigningProfilePermissionOutputWithContext(ctx context.Context) SigningProfilePermissionOutput {
	return o
}

func (o SigningProfilePermissionOutput) ToSigningProfilePermissionPtrOutput() SigningProfilePermissionPtrOutput {
	return o.ToSigningProfilePermissionPtrOutputWithContext(context.Background())
}

func (o SigningProfilePermissionOutput) ToSigningProfilePermissionPtrOutputWithContext(ctx context.Context) SigningProfilePermissionPtrOutput {
	return o.ApplyT(func(v SigningProfilePermission) *SigningProfilePermission {
		return &v
	}).(SigningProfilePermissionPtrOutput)
}

type SigningProfilePermissionPtrOutput struct {
	*pulumi.OutputState
}

func (SigningProfilePermissionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SigningProfilePermission)(nil))
}

func (o SigningProfilePermissionPtrOutput) ToSigningProfilePermissionPtrOutput() SigningProfilePermissionPtrOutput {
	return o
}

func (o SigningProfilePermissionPtrOutput) ToSigningProfilePermissionPtrOutputWithContext(ctx context.Context) SigningProfilePermissionPtrOutput {
	return o
}

type SigningProfilePermissionArrayOutput struct{ *pulumi.OutputState }

func (SigningProfilePermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SigningProfilePermission)(nil))
}

func (o SigningProfilePermissionArrayOutput) ToSigningProfilePermissionArrayOutput() SigningProfilePermissionArrayOutput {
	return o
}

func (o SigningProfilePermissionArrayOutput) ToSigningProfilePermissionArrayOutputWithContext(ctx context.Context) SigningProfilePermissionArrayOutput {
	return o
}

func (o SigningProfilePermissionArrayOutput) Index(i pulumi.IntInput) SigningProfilePermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SigningProfilePermission {
		return vs[0].([]SigningProfilePermission)[vs[1].(int)]
	}).(SigningProfilePermissionOutput)
}

type SigningProfilePermissionMapOutput struct{ *pulumi.OutputState }

func (SigningProfilePermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SigningProfilePermission)(nil))
}

func (o SigningProfilePermissionMapOutput) ToSigningProfilePermissionMapOutput() SigningProfilePermissionMapOutput {
	return o
}

func (o SigningProfilePermissionMapOutput) ToSigningProfilePermissionMapOutputWithContext(ctx context.Context) SigningProfilePermissionMapOutput {
	return o
}

func (o SigningProfilePermissionMapOutput) MapIndex(k pulumi.StringInput) SigningProfilePermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SigningProfilePermission {
		return vs[0].(map[string]SigningProfilePermission)[vs[1].(string)]
	}).(SigningProfilePermissionOutput)
}

func init() {
	pulumi.RegisterOutputType(SigningProfilePermissionOutput{})
	pulumi.RegisterOutputType(SigningProfilePermissionPtrOutput{})
	pulumi.RegisterOutputType(SigningProfilePermissionArrayOutput{})
	pulumi.RegisterOutputType(SigningProfilePermissionMapOutput{})
}
