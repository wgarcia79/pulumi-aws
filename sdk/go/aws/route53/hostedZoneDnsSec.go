// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages Route 53 Hosted Zone Domain Name System Security Extensions (DNSSEC). For more information about managing DNSSEC in Route 53, see the [Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"encoding/json"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/kms"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"Statement": []interface{}{
// 				map[string]interface{}{
// 					"Action": []string{
// 						"kms:DescribeKey",
// 						"kms:GetPublicKey",
// 						"kms:Sign",
// 					},
// 					"Effect": "Allow",
// 					"Principal": map[string]interface{}{
// 						"Service": "api-service.dnssec.route53.aws.internal",
// 					},
// 					"Sid": "Route 53 DNSSEC Permissions",
// 				},
// 				map[string]interface{}{
// 					"Action": "kms:*",
// 					"Effect": "Allow",
// 					"Principal": map[string]interface{}{
// 						"AWS": "*",
// 					},
// 					"Resource": "*",
// 					"Sid":      "IAM User Permissions",
// 				},
// 			},
// 			"Version": "2012-10-17",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		_, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
// 			CustomerMasterKeySpec: pulumi.String("ECC_NIST_P256"),
// 			DeletionWindowInDays:  pulumi.Int(7),
// 			KeyUsage:              pulumi.String("SIGN_VERIFY"),
// 			Policy:                pulumi.String(json0),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewZone(ctx, "exampleZone", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleKeySigningKey, err := route53.NewKeySigningKey(ctx, "exampleKeySigningKey", &route53.KeySigningKeyArgs{
// 			HostedZoneId:            pulumi.Any(aws_route53_zone.Test.Id),
// 			KeyManagementServiceArn: pulumi.Any(aws_kms_key.Test.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewHostedZoneDnsSec(ctx, "exampleHostedZoneDnsSec", &route53.HostedZoneDnsSecArgs{
// 			HostedZoneId: exampleKeySigningKey.HostedZoneId,
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
// `aws_route53_hosted_zone_dnssec` resources can be imported by using the Route 53 Hosted Zone identifier, e.g.
//
// ```sh
//  $ pulumi import aws:route53/hostedZoneDnsSec:HostedZoneDnsSec example Z1D633PJN98FT9
// ```
type HostedZoneDnsSec struct {
	pulumi.CustomResourceState

	// Identifier of the Route 53 Hosted Zone.
	HostedZoneId pulumi.StringOutput `pulumi:"hostedZoneId"`
	// Hosted Zone signing status. Valid values: `SIGNING`, `NOT_SIGNING`. Defaults to `SIGNING`.
	SigningStatus pulumi.StringPtrOutput `pulumi:"signingStatus"`
}

// NewHostedZoneDnsSec registers a new resource with the given unique name, arguments, and options.
func NewHostedZoneDnsSec(ctx *pulumi.Context,
	name string, args *HostedZoneDnsSecArgs, opts ...pulumi.ResourceOption) (*HostedZoneDnsSec, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostedZoneId == nil {
		return nil, errors.New("invalid value for required argument 'HostedZoneId'")
	}
	var resource HostedZoneDnsSec
	err := ctx.RegisterResource("aws:route53/hostedZoneDnsSec:HostedZoneDnsSec", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostedZoneDnsSec gets an existing HostedZoneDnsSec resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostedZoneDnsSec(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostedZoneDnsSecState, opts ...pulumi.ResourceOption) (*HostedZoneDnsSec, error) {
	var resource HostedZoneDnsSec
	err := ctx.ReadResource("aws:route53/hostedZoneDnsSec:HostedZoneDnsSec", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostedZoneDnsSec resources.
type hostedZoneDnsSecState struct {
	// Identifier of the Route 53 Hosted Zone.
	HostedZoneId *string `pulumi:"hostedZoneId"`
	// Hosted Zone signing status. Valid values: `SIGNING`, `NOT_SIGNING`. Defaults to `SIGNING`.
	SigningStatus *string `pulumi:"signingStatus"`
}

type HostedZoneDnsSecState struct {
	// Identifier of the Route 53 Hosted Zone.
	HostedZoneId pulumi.StringPtrInput
	// Hosted Zone signing status. Valid values: `SIGNING`, `NOT_SIGNING`. Defaults to `SIGNING`.
	SigningStatus pulumi.StringPtrInput
}

func (HostedZoneDnsSecState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedZoneDnsSecState)(nil)).Elem()
}

type hostedZoneDnsSecArgs struct {
	// Identifier of the Route 53 Hosted Zone.
	HostedZoneId string `pulumi:"hostedZoneId"`
	// Hosted Zone signing status. Valid values: `SIGNING`, `NOT_SIGNING`. Defaults to `SIGNING`.
	SigningStatus *string `pulumi:"signingStatus"`
}

// The set of arguments for constructing a HostedZoneDnsSec resource.
type HostedZoneDnsSecArgs struct {
	// Identifier of the Route 53 Hosted Zone.
	HostedZoneId pulumi.StringInput
	// Hosted Zone signing status. Valid values: `SIGNING`, `NOT_SIGNING`. Defaults to `SIGNING`.
	SigningStatus pulumi.StringPtrInput
}

func (HostedZoneDnsSecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedZoneDnsSecArgs)(nil)).Elem()
}

type HostedZoneDnsSecInput interface {
	pulumi.Input

	ToHostedZoneDnsSecOutput() HostedZoneDnsSecOutput
	ToHostedZoneDnsSecOutputWithContext(ctx context.Context) HostedZoneDnsSecOutput
}

func (*HostedZoneDnsSec) ElementType() reflect.Type {
	return reflect.TypeOf((*HostedZoneDnsSec)(nil))
}

func (i *HostedZoneDnsSec) ToHostedZoneDnsSecOutput() HostedZoneDnsSecOutput {
	return i.ToHostedZoneDnsSecOutputWithContext(context.Background())
}

func (i *HostedZoneDnsSec) ToHostedZoneDnsSecOutputWithContext(ctx context.Context) HostedZoneDnsSecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedZoneDnsSecOutput)
}

func (i *HostedZoneDnsSec) ToHostedZoneDnsSecPtrOutput() HostedZoneDnsSecPtrOutput {
	return i.ToHostedZoneDnsSecPtrOutputWithContext(context.Background())
}

func (i *HostedZoneDnsSec) ToHostedZoneDnsSecPtrOutputWithContext(ctx context.Context) HostedZoneDnsSecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedZoneDnsSecPtrOutput)
}

type HostedZoneDnsSecPtrInput interface {
	pulumi.Input

	ToHostedZoneDnsSecPtrOutput() HostedZoneDnsSecPtrOutput
	ToHostedZoneDnsSecPtrOutputWithContext(ctx context.Context) HostedZoneDnsSecPtrOutput
}

type hostedZoneDnsSecPtrType HostedZoneDnsSecArgs

func (*hostedZoneDnsSecPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HostedZoneDnsSec)(nil))
}

func (i *hostedZoneDnsSecPtrType) ToHostedZoneDnsSecPtrOutput() HostedZoneDnsSecPtrOutput {
	return i.ToHostedZoneDnsSecPtrOutputWithContext(context.Background())
}

func (i *hostedZoneDnsSecPtrType) ToHostedZoneDnsSecPtrOutputWithContext(ctx context.Context) HostedZoneDnsSecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedZoneDnsSecPtrOutput)
}

// HostedZoneDnsSecArrayInput is an input type that accepts HostedZoneDnsSecArray and HostedZoneDnsSecArrayOutput values.
// You can construct a concrete instance of `HostedZoneDnsSecArrayInput` via:
//
//          HostedZoneDnsSecArray{ HostedZoneDnsSecArgs{...} }
type HostedZoneDnsSecArrayInput interface {
	pulumi.Input

	ToHostedZoneDnsSecArrayOutput() HostedZoneDnsSecArrayOutput
	ToHostedZoneDnsSecArrayOutputWithContext(context.Context) HostedZoneDnsSecArrayOutput
}

type HostedZoneDnsSecArray []HostedZoneDnsSecInput

func (HostedZoneDnsSecArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*HostedZoneDnsSec)(nil))
}

func (i HostedZoneDnsSecArray) ToHostedZoneDnsSecArrayOutput() HostedZoneDnsSecArrayOutput {
	return i.ToHostedZoneDnsSecArrayOutputWithContext(context.Background())
}

func (i HostedZoneDnsSecArray) ToHostedZoneDnsSecArrayOutputWithContext(ctx context.Context) HostedZoneDnsSecArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedZoneDnsSecArrayOutput)
}

// HostedZoneDnsSecMapInput is an input type that accepts HostedZoneDnsSecMap and HostedZoneDnsSecMapOutput values.
// You can construct a concrete instance of `HostedZoneDnsSecMapInput` via:
//
//          HostedZoneDnsSecMap{ "key": HostedZoneDnsSecArgs{...} }
type HostedZoneDnsSecMapInput interface {
	pulumi.Input

	ToHostedZoneDnsSecMapOutput() HostedZoneDnsSecMapOutput
	ToHostedZoneDnsSecMapOutputWithContext(context.Context) HostedZoneDnsSecMapOutput
}

type HostedZoneDnsSecMap map[string]HostedZoneDnsSecInput

func (HostedZoneDnsSecMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*HostedZoneDnsSec)(nil))
}

func (i HostedZoneDnsSecMap) ToHostedZoneDnsSecMapOutput() HostedZoneDnsSecMapOutput {
	return i.ToHostedZoneDnsSecMapOutputWithContext(context.Background())
}

func (i HostedZoneDnsSecMap) ToHostedZoneDnsSecMapOutputWithContext(ctx context.Context) HostedZoneDnsSecMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedZoneDnsSecMapOutput)
}

type HostedZoneDnsSecOutput struct {
	*pulumi.OutputState
}

func (HostedZoneDnsSecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostedZoneDnsSec)(nil))
}

func (o HostedZoneDnsSecOutput) ToHostedZoneDnsSecOutput() HostedZoneDnsSecOutput {
	return o
}

func (o HostedZoneDnsSecOutput) ToHostedZoneDnsSecOutputWithContext(ctx context.Context) HostedZoneDnsSecOutput {
	return o
}

func (o HostedZoneDnsSecOutput) ToHostedZoneDnsSecPtrOutput() HostedZoneDnsSecPtrOutput {
	return o.ToHostedZoneDnsSecPtrOutputWithContext(context.Background())
}

func (o HostedZoneDnsSecOutput) ToHostedZoneDnsSecPtrOutputWithContext(ctx context.Context) HostedZoneDnsSecPtrOutput {
	return o.ApplyT(func(v HostedZoneDnsSec) *HostedZoneDnsSec {
		return &v
	}).(HostedZoneDnsSecPtrOutput)
}

type HostedZoneDnsSecPtrOutput struct {
	*pulumi.OutputState
}

func (HostedZoneDnsSecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostedZoneDnsSec)(nil))
}

func (o HostedZoneDnsSecPtrOutput) ToHostedZoneDnsSecPtrOutput() HostedZoneDnsSecPtrOutput {
	return o
}

func (o HostedZoneDnsSecPtrOutput) ToHostedZoneDnsSecPtrOutputWithContext(ctx context.Context) HostedZoneDnsSecPtrOutput {
	return o
}

type HostedZoneDnsSecArrayOutput struct{ *pulumi.OutputState }

func (HostedZoneDnsSecArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostedZoneDnsSec)(nil))
}

func (o HostedZoneDnsSecArrayOutput) ToHostedZoneDnsSecArrayOutput() HostedZoneDnsSecArrayOutput {
	return o
}

func (o HostedZoneDnsSecArrayOutput) ToHostedZoneDnsSecArrayOutputWithContext(ctx context.Context) HostedZoneDnsSecArrayOutput {
	return o
}

func (o HostedZoneDnsSecArrayOutput) Index(i pulumi.IntInput) HostedZoneDnsSecOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostedZoneDnsSec {
		return vs[0].([]HostedZoneDnsSec)[vs[1].(int)]
	}).(HostedZoneDnsSecOutput)
}

type HostedZoneDnsSecMapOutput struct{ *pulumi.OutputState }

func (HostedZoneDnsSecMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]HostedZoneDnsSec)(nil))
}

func (o HostedZoneDnsSecMapOutput) ToHostedZoneDnsSecMapOutput() HostedZoneDnsSecMapOutput {
	return o
}

func (o HostedZoneDnsSecMapOutput) ToHostedZoneDnsSecMapOutputWithContext(ctx context.Context) HostedZoneDnsSecMapOutput {
	return o
}

func (o HostedZoneDnsSecMapOutput) MapIndex(k pulumi.StringInput) HostedZoneDnsSecOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) HostedZoneDnsSec {
		return vs[0].(map[string]HostedZoneDnsSec)[vs[1].(string)]
	}).(HostedZoneDnsSecOutput)
}

func init() {
	pulumi.RegisterOutputType(HostedZoneDnsSecOutput{})
	pulumi.RegisterOutputType(HostedZoneDnsSecPtrOutput{})
	pulumi.RegisterOutputType(HostedZoneDnsSecArrayOutput{})
	pulumi.RegisterOutputType(HostedZoneDnsSecMapOutput{})
}
