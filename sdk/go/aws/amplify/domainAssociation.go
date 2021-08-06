// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package amplify

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Amplify Domain Association resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/amplify"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleApp, err := amplify.NewApp(ctx, "exampleApp", &amplify.AppArgs{
// 			CustomRules: amplify.AppCustomRuleArray{
// 				&amplify.AppCustomRuleArgs{
// 					Source: pulumi.String("https://example.com"),
// 					Status: pulumi.String("302"),
// 					Target: pulumi.String("https://www.example.com"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		master, err := amplify.NewBranch(ctx, "master", &amplify.BranchArgs{
// 			AppId:      exampleApp.ID(),
// 			BranchName: pulumi.String("master"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = amplify.NewDomainAssociation(ctx, "exampleDomainAssociation", &amplify.DomainAssociationArgs{
// 			AppId:      exampleApp.ID(),
// 			DomainName: pulumi.String("example.com"),
// 			SubDomains: amplify.DomainAssociationSubDomainArray{
// 				&amplify.DomainAssociationSubDomainArgs{
// 					BranchName: master.BranchName,
// 					Prefix:     pulumi.String(""),
// 				},
// 				&amplify.DomainAssociationSubDomainArgs{
// 					BranchName: master.BranchName,
// 					Prefix:     pulumi.String("www"),
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
// Amplify domain association can be imported using `app_id` and `domain_name`, e.g.
//
// ```sh
//  $ pulumi import aws:amplify/domainAssociation:DomainAssociation app d2ypk4k47z8u6/example.com
// ```
type DomainAssociation struct {
	pulumi.CustomResourceState

	// The unique ID for an Amplify app.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The Amazon Resource Name (ARN) for the domain association.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The DNS record for certificate verification.
	CertificateVerificationDnsRecord pulumi.StringOutput `pulumi:"certificateVerificationDnsRecord"`
	// The domain name for the domain association.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The setting for the subdomain. Documented below.
	SubDomains DomainAssociationSubDomainArrayOutput `pulumi:"subDomains"`
	// If enabled, the resource will wait for the domain association status to change to `PENDING_DEPLOYMENT` or `AVAILABLE`. Setting this to `false` will skip the process. Default: `true`.
	WaitForVerification pulumi.BoolPtrOutput `pulumi:"waitForVerification"`
}

// NewDomainAssociation registers a new resource with the given unique name, arguments, and options.
func NewDomainAssociation(ctx *pulumi.Context,
	name string, args *DomainAssociationArgs, opts ...pulumi.ResourceOption) (*DomainAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.SubDomains == nil {
		return nil, errors.New("invalid value for required argument 'SubDomains'")
	}
	var resource DomainAssociation
	err := ctx.RegisterResource("aws:amplify/domainAssociation:DomainAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainAssociation gets an existing DomainAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainAssociationState, opts ...pulumi.ResourceOption) (*DomainAssociation, error) {
	var resource DomainAssociation
	err := ctx.ReadResource("aws:amplify/domainAssociation:DomainAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainAssociation resources.
type domainAssociationState struct {
	// The unique ID for an Amplify app.
	AppId *string `pulumi:"appId"`
	// The Amazon Resource Name (ARN) for the domain association.
	Arn *string `pulumi:"arn"`
	// The DNS record for certificate verification.
	CertificateVerificationDnsRecord *string `pulumi:"certificateVerificationDnsRecord"`
	// The domain name for the domain association.
	DomainName *string `pulumi:"domainName"`
	// The setting for the subdomain. Documented below.
	SubDomains []DomainAssociationSubDomain `pulumi:"subDomains"`
	// If enabled, the resource will wait for the domain association status to change to `PENDING_DEPLOYMENT` or `AVAILABLE`. Setting this to `false` will skip the process. Default: `true`.
	WaitForVerification *bool `pulumi:"waitForVerification"`
}

type DomainAssociationState struct {
	// The unique ID for an Amplify app.
	AppId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for the domain association.
	Arn pulumi.StringPtrInput
	// The DNS record for certificate verification.
	CertificateVerificationDnsRecord pulumi.StringPtrInput
	// The domain name for the domain association.
	DomainName pulumi.StringPtrInput
	// The setting for the subdomain. Documented below.
	SubDomains DomainAssociationSubDomainArrayInput
	// If enabled, the resource will wait for the domain association status to change to `PENDING_DEPLOYMENT` or `AVAILABLE`. Setting this to `false` will skip the process. Default: `true`.
	WaitForVerification pulumi.BoolPtrInput
}

func (DomainAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainAssociationState)(nil)).Elem()
}

type domainAssociationArgs struct {
	// The unique ID for an Amplify app.
	AppId string `pulumi:"appId"`
	// The domain name for the domain association.
	DomainName string `pulumi:"domainName"`
	// The setting for the subdomain. Documented below.
	SubDomains []DomainAssociationSubDomain `pulumi:"subDomains"`
	// If enabled, the resource will wait for the domain association status to change to `PENDING_DEPLOYMENT` or `AVAILABLE`. Setting this to `false` will skip the process. Default: `true`.
	WaitForVerification *bool `pulumi:"waitForVerification"`
}

// The set of arguments for constructing a DomainAssociation resource.
type DomainAssociationArgs struct {
	// The unique ID for an Amplify app.
	AppId pulumi.StringInput
	// The domain name for the domain association.
	DomainName pulumi.StringInput
	// The setting for the subdomain. Documented below.
	SubDomains DomainAssociationSubDomainArrayInput
	// If enabled, the resource will wait for the domain association status to change to `PENDING_DEPLOYMENT` or `AVAILABLE`. Setting this to `false` will skip the process. Default: `true`.
	WaitForVerification pulumi.BoolPtrInput
}

func (DomainAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainAssociationArgs)(nil)).Elem()
}

type DomainAssociationInput interface {
	pulumi.Input

	ToDomainAssociationOutput() DomainAssociationOutput
	ToDomainAssociationOutputWithContext(ctx context.Context) DomainAssociationOutput
}

func (*DomainAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainAssociation)(nil))
}

func (i *DomainAssociation) ToDomainAssociationOutput() DomainAssociationOutput {
	return i.ToDomainAssociationOutputWithContext(context.Background())
}

func (i *DomainAssociation) ToDomainAssociationOutputWithContext(ctx context.Context) DomainAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAssociationOutput)
}

func (i *DomainAssociation) ToDomainAssociationPtrOutput() DomainAssociationPtrOutput {
	return i.ToDomainAssociationPtrOutputWithContext(context.Background())
}

func (i *DomainAssociation) ToDomainAssociationPtrOutputWithContext(ctx context.Context) DomainAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAssociationPtrOutput)
}

type DomainAssociationPtrInput interface {
	pulumi.Input

	ToDomainAssociationPtrOutput() DomainAssociationPtrOutput
	ToDomainAssociationPtrOutputWithContext(ctx context.Context) DomainAssociationPtrOutput
}

type domainAssociationPtrType DomainAssociationArgs

func (*domainAssociationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainAssociation)(nil))
}

func (i *domainAssociationPtrType) ToDomainAssociationPtrOutput() DomainAssociationPtrOutput {
	return i.ToDomainAssociationPtrOutputWithContext(context.Background())
}

func (i *domainAssociationPtrType) ToDomainAssociationPtrOutputWithContext(ctx context.Context) DomainAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAssociationPtrOutput)
}

// DomainAssociationArrayInput is an input type that accepts DomainAssociationArray and DomainAssociationArrayOutput values.
// You can construct a concrete instance of `DomainAssociationArrayInput` via:
//
//          DomainAssociationArray{ DomainAssociationArgs{...} }
type DomainAssociationArrayInput interface {
	pulumi.Input

	ToDomainAssociationArrayOutput() DomainAssociationArrayOutput
	ToDomainAssociationArrayOutputWithContext(context.Context) DomainAssociationArrayOutput
}

type DomainAssociationArray []DomainAssociationInput

func (DomainAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainAssociation)(nil)).Elem()
}

func (i DomainAssociationArray) ToDomainAssociationArrayOutput() DomainAssociationArrayOutput {
	return i.ToDomainAssociationArrayOutputWithContext(context.Background())
}

func (i DomainAssociationArray) ToDomainAssociationArrayOutputWithContext(ctx context.Context) DomainAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAssociationArrayOutput)
}

// DomainAssociationMapInput is an input type that accepts DomainAssociationMap and DomainAssociationMapOutput values.
// You can construct a concrete instance of `DomainAssociationMapInput` via:
//
//          DomainAssociationMap{ "key": DomainAssociationArgs{...} }
type DomainAssociationMapInput interface {
	pulumi.Input

	ToDomainAssociationMapOutput() DomainAssociationMapOutput
	ToDomainAssociationMapOutputWithContext(context.Context) DomainAssociationMapOutput
}

type DomainAssociationMap map[string]DomainAssociationInput

func (DomainAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainAssociation)(nil)).Elem()
}

func (i DomainAssociationMap) ToDomainAssociationMapOutput() DomainAssociationMapOutput {
	return i.ToDomainAssociationMapOutputWithContext(context.Background())
}

func (i DomainAssociationMap) ToDomainAssociationMapOutputWithContext(ctx context.Context) DomainAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAssociationMapOutput)
}

type DomainAssociationOutput struct {
	*pulumi.OutputState
}

func (DomainAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainAssociation)(nil))
}

func (o DomainAssociationOutput) ToDomainAssociationOutput() DomainAssociationOutput {
	return o
}

func (o DomainAssociationOutput) ToDomainAssociationOutputWithContext(ctx context.Context) DomainAssociationOutput {
	return o
}

func (o DomainAssociationOutput) ToDomainAssociationPtrOutput() DomainAssociationPtrOutput {
	return o.ToDomainAssociationPtrOutputWithContext(context.Background())
}

func (o DomainAssociationOutput) ToDomainAssociationPtrOutputWithContext(ctx context.Context) DomainAssociationPtrOutput {
	return o.ApplyT(func(v DomainAssociation) *DomainAssociation {
		return &v
	}).(DomainAssociationPtrOutput)
}

type DomainAssociationPtrOutput struct {
	*pulumi.OutputState
}

func (DomainAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainAssociation)(nil))
}

func (o DomainAssociationPtrOutput) ToDomainAssociationPtrOutput() DomainAssociationPtrOutput {
	return o
}

func (o DomainAssociationPtrOutput) ToDomainAssociationPtrOutputWithContext(ctx context.Context) DomainAssociationPtrOutput {
	return o
}

type DomainAssociationArrayOutput struct{ *pulumi.OutputState }

func (DomainAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainAssociation)(nil))
}

func (o DomainAssociationArrayOutput) ToDomainAssociationArrayOutput() DomainAssociationArrayOutput {
	return o
}

func (o DomainAssociationArrayOutput) ToDomainAssociationArrayOutputWithContext(ctx context.Context) DomainAssociationArrayOutput {
	return o
}

func (o DomainAssociationArrayOutput) Index(i pulumi.IntInput) DomainAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainAssociation {
		return vs[0].([]DomainAssociation)[vs[1].(int)]
	}).(DomainAssociationOutput)
}

type DomainAssociationMapOutput struct{ *pulumi.OutputState }

func (DomainAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DomainAssociation)(nil))
}

func (o DomainAssociationMapOutput) ToDomainAssociationMapOutput() DomainAssociationMapOutput {
	return o
}

func (o DomainAssociationMapOutput) ToDomainAssociationMapOutputWithContext(ctx context.Context) DomainAssociationMapOutput {
	return o
}

func (o DomainAssociationMapOutput) MapIndex(k pulumi.StringInput) DomainAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DomainAssociation {
		return vs[0].(map[string]DomainAssociation)[vs[1].(string)]
	}).(DomainAssociationOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainAssociationOutput{})
	pulumi.RegisterOutputType(DomainAssociationPtrOutput{})
	pulumi.RegisterOutputType(DomainAssociationArrayOutput{})
	pulumi.RegisterOutputType(DomainAssociationMapOutput{})
}
