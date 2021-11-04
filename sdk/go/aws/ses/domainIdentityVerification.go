// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a successful verification of an SES domain identity.
//
// Most commonly, this resource is used together with `route53.Record` and
// `ses.DomainIdentity` to request an SES domain identity,
// deploy the required DNS verification records, and wait for verification to complete.
//
// > **WARNING:** This resource implements a part of the verification workflow. It does not represent a real-world entity in AWS, therefore changing or deleting this resource on its own has no immediate effect.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := ses.NewDomainIdentity(ctx, "example", &ses.DomainIdentityArgs{
// 			Domain: pulumi.String("example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAmazonsesVerificationRecord, err := route53.NewRecord(ctx, "exampleAmazonsesVerificationRecord", &route53.RecordArgs{
// 			ZoneId: pulumi.Any(aws_route53_zone.Example.Zone_id),
// 			Name: example.ID().ApplyT(func(id string) (string, error) {
// 				return fmt.Sprintf("%v%v", "_amazonses.", id), nil
// 			}).(pulumi.StringOutput),
// 			Type: pulumi.String("TXT"),
// 			Ttl:  pulumi.Int(600),
// 			Records: pulumi.StringArray{
// 				example.VerificationToken,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ses.NewDomainIdentityVerification(ctx, "exampleVerification", &ses.DomainIdentityVerificationArgs{
// 			Domain: example.ID(),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleAmazonsesVerificationRecord,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DomainIdentityVerification struct {
	pulumi.CustomResourceState

	// The ARN of the domain identity.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The domain name of the SES domain identity to verify.
	Domain pulumi.StringOutput `pulumi:"domain"`
}

// NewDomainIdentityVerification registers a new resource with the given unique name, arguments, and options.
func NewDomainIdentityVerification(ctx *pulumi.Context,
	name string, args *DomainIdentityVerificationArgs, opts ...pulumi.ResourceOption) (*DomainIdentityVerification, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	var resource DomainIdentityVerification
	err := ctx.RegisterResource("aws:ses/domainIdentityVerification:DomainIdentityVerification", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainIdentityVerification gets an existing DomainIdentityVerification resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainIdentityVerification(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainIdentityVerificationState, opts ...pulumi.ResourceOption) (*DomainIdentityVerification, error) {
	var resource DomainIdentityVerification
	err := ctx.ReadResource("aws:ses/domainIdentityVerification:DomainIdentityVerification", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainIdentityVerification resources.
type domainIdentityVerificationState struct {
	// The ARN of the domain identity.
	Arn *string `pulumi:"arn"`
	// The domain name of the SES domain identity to verify.
	Domain *string `pulumi:"domain"`
}

type DomainIdentityVerificationState struct {
	// The ARN of the domain identity.
	Arn pulumi.StringPtrInput
	// The domain name of the SES domain identity to verify.
	Domain pulumi.StringPtrInput
}

func (DomainIdentityVerificationState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainIdentityVerificationState)(nil)).Elem()
}

type domainIdentityVerificationArgs struct {
	// The domain name of the SES domain identity to verify.
	Domain string `pulumi:"domain"`
}

// The set of arguments for constructing a DomainIdentityVerification resource.
type DomainIdentityVerificationArgs struct {
	// The domain name of the SES domain identity to verify.
	Domain pulumi.StringInput
}

func (DomainIdentityVerificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainIdentityVerificationArgs)(nil)).Elem()
}

type DomainIdentityVerificationInput interface {
	pulumi.Input

	ToDomainIdentityVerificationOutput() DomainIdentityVerificationOutput
	ToDomainIdentityVerificationOutputWithContext(ctx context.Context) DomainIdentityVerificationOutput
}

func (*DomainIdentityVerification) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainIdentityVerification)(nil))
}

func (i *DomainIdentityVerification) ToDomainIdentityVerificationOutput() DomainIdentityVerificationOutput {
	return i.ToDomainIdentityVerificationOutputWithContext(context.Background())
}

func (i *DomainIdentityVerification) ToDomainIdentityVerificationOutputWithContext(ctx context.Context) DomainIdentityVerificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainIdentityVerificationOutput)
}

func (i *DomainIdentityVerification) ToDomainIdentityVerificationPtrOutput() DomainIdentityVerificationPtrOutput {
	return i.ToDomainIdentityVerificationPtrOutputWithContext(context.Background())
}

func (i *DomainIdentityVerification) ToDomainIdentityVerificationPtrOutputWithContext(ctx context.Context) DomainIdentityVerificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainIdentityVerificationPtrOutput)
}

type DomainIdentityVerificationPtrInput interface {
	pulumi.Input

	ToDomainIdentityVerificationPtrOutput() DomainIdentityVerificationPtrOutput
	ToDomainIdentityVerificationPtrOutputWithContext(ctx context.Context) DomainIdentityVerificationPtrOutput
}

type domainIdentityVerificationPtrType DomainIdentityVerificationArgs

func (*domainIdentityVerificationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainIdentityVerification)(nil))
}

func (i *domainIdentityVerificationPtrType) ToDomainIdentityVerificationPtrOutput() DomainIdentityVerificationPtrOutput {
	return i.ToDomainIdentityVerificationPtrOutputWithContext(context.Background())
}

func (i *domainIdentityVerificationPtrType) ToDomainIdentityVerificationPtrOutputWithContext(ctx context.Context) DomainIdentityVerificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainIdentityVerificationPtrOutput)
}

// DomainIdentityVerificationArrayInput is an input type that accepts DomainIdentityVerificationArray and DomainIdentityVerificationArrayOutput values.
// You can construct a concrete instance of `DomainIdentityVerificationArrayInput` via:
//
//          DomainIdentityVerificationArray{ DomainIdentityVerificationArgs{...} }
type DomainIdentityVerificationArrayInput interface {
	pulumi.Input

	ToDomainIdentityVerificationArrayOutput() DomainIdentityVerificationArrayOutput
	ToDomainIdentityVerificationArrayOutputWithContext(context.Context) DomainIdentityVerificationArrayOutput
}

type DomainIdentityVerificationArray []DomainIdentityVerificationInput

func (DomainIdentityVerificationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainIdentityVerification)(nil)).Elem()
}

func (i DomainIdentityVerificationArray) ToDomainIdentityVerificationArrayOutput() DomainIdentityVerificationArrayOutput {
	return i.ToDomainIdentityVerificationArrayOutputWithContext(context.Background())
}

func (i DomainIdentityVerificationArray) ToDomainIdentityVerificationArrayOutputWithContext(ctx context.Context) DomainIdentityVerificationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainIdentityVerificationArrayOutput)
}

// DomainIdentityVerificationMapInput is an input type that accepts DomainIdentityVerificationMap and DomainIdentityVerificationMapOutput values.
// You can construct a concrete instance of `DomainIdentityVerificationMapInput` via:
//
//          DomainIdentityVerificationMap{ "key": DomainIdentityVerificationArgs{...} }
type DomainIdentityVerificationMapInput interface {
	pulumi.Input

	ToDomainIdentityVerificationMapOutput() DomainIdentityVerificationMapOutput
	ToDomainIdentityVerificationMapOutputWithContext(context.Context) DomainIdentityVerificationMapOutput
}

type DomainIdentityVerificationMap map[string]DomainIdentityVerificationInput

func (DomainIdentityVerificationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainIdentityVerification)(nil)).Elem()
}

func (i DomainIdentityVerificationMap) ToDomainIdentityVerificationMapOutput() DomainIdentityVerificationMapOutput {
	return i.ToDomainIdentityVerificationMapOutputWithContext(context.Background())
}

func (i DomainIdentityVerificationMap) ToDomainIdentityVerificationMapOutputWithContext(ctx context.Context) DomainIdentityVerificationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainIdentityVerificationMapOutput)
}

type DomainIdentityVerificationOutput struct{ *pulumi.OutputState }

func (DomainIdentityVerificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainIdentityVerification)(nil))
}

func (o DomainIdentityVerificationOutput) ToDomainIdentityVerificationOutput() DomainIdentityVerificationOutput {
	return o
}

func (o DomainIdentityVerificationOutput) ToDomainIdentityVerificationOutputWithContext(ctx context.Context) DomainIdentityVerificationOutput {
	return o
}

func (o DomainIdentityVerificationOutput) ToDomainIdentityVerificationPtrOutput() DomainIdentityVerificationPtrOutput {
	return o.ToDomainIdentityVerificationPtrOutputWithContext(context.Background())
}

func (o DomainIdentityVerificationOutput) ToDomainIdentityVerificationPtrOutputWithContext(ctx context.Context) DomainIdentityVerificationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DomainIdentityVerification) *DomainIdentityVerification {
		return &v
	}).(DomainIdentityVerificationPtrOutput)
}

type DomainIdentityVerificationPtrOutput struct{ *pulumi.OutputState }

func (DomainIdentityVerificationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainIdentityVerification)(nil))
}

func (o DomainIdentityVerificationPtrOutput) ToDomainIdentityVerificationPtrOutput() DomainIdentityVerificationPtrOutput {
	return o
}

func (o DomainIdentityVerificationPtrOutput) ToDomainIdentityVerificationPtrOutputWithContext(ctx context.Context) DomainIdentityVerificationPtrOutput {
	return o
}

func (o DomainIdentityVerificationPtrOutput) Elem() DomainIdentityVerificationOutput {
	return o.ApplyT(func(v *DomainIdentityVerification) DomainIdentityVerification {
		if v != nil {
			return *v
		}
		var ret DomainIdentityVerification
		return ret
	}).(DomainIdentityVerificationOutput)
}

type DomainIdentityVerificationArrayOutput struct{ *pulumi.OutputState }

func (DomainIdentityVerificationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainIdentityVerification)(nil))
}

func (o DomainIdentityVerificationArrayOutput) ToDomainIdentityVerificationArrayOutput() DomainIdentityVerificationArrayOutput {
	return o
}

func (o DomainIdentityVerificationArrayOutput) ToDomainIdentityVerificationArrayOutputWithContext(ctx context.Context) DomainIdentityVerificationArrayOutput {
	return o
}

func (o DomainIdentityVerificationArrayOutput) Index(i pulumi.IntInput) DomainIdentityVerificationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainIdentityVerification {
		return vs[0].([]DomainIdentityVerification)[vs[1].(int)]
	}).(DomainIdentityVerificationOutput)
}

type DomainIdentityVerificationMapOutput struct{ *pulumi.OutputState }

func (DomainIdentityVerificationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DomainIdentityVerification)(nil))
}

func (o DomainIdentityVerificationMapOutput) ToDomainIdentityVerificationMapOutput() DomainIdentityVerificationMapOutput {
	return o
}

func (o DomainIdentityVerificationMapOutput) ToDomainIdentityVerificationMapOutputWithContext(ctx context.Context) DomainIdentityVerificationMapOutput {
	return o
}

func (o DomainIdentityVerificationMapOutput) MapIndex(k pulumi.StringInput) DomainIdentityVerificationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DomainIdentityVerification {
		return vs[0].(map[string]DomainIdentityVerification)[vs[1].(string)]
	}).(DomainIdentityVerificationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainIdentityVerificationInput)(nil)).Elem(), &DomainIdentityVerification{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainIdentityVerificationPtrInput)(nil)).Elem(), &DomainIdentityVerification{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainIdentityVerificationArrayInput)(nil)).Elem(), DomainIdentityVerificationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainIdentityVerificationMapInput)(nil)).Elem(), DomainIdentityVerificationMap{})
	pulumi.RegisterOutputType(DomainIdentityVerificationOutput{})
	pulumi.RegisterOutputType(DomainIdentityVerificationPtrOutput{})
	pulumi.RegisterOutputType(DomainIdentityVerificationArrayOutput{})
	pulumi.RegisterOutputType(DomainIdentityVerificationMapOutput{})
}
