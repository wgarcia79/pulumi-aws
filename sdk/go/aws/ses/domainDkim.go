// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an SES domain DKIM generation resource.
//
// Domain ownership needs to be confirmed first using `ses.DomainIdentity` resource.
//
// ## Import
//
// DKIM tokens can be imported using the `domain` attribute, e.g.
//
// ```sh
//  $ pulumi import aws:ses/domainDkim:DomainDkim example example.com
// ```
type DomainDkim struct {
	pulumi.CustomResourceState

	// DKIM tokens generated by SES.
	// These tokens should be used to create CNAME records used to verify SES Easy DKIM.
	// See below for an example of how this might be achieved
	// when the domain is hosted in Route 53 and managed by this provider.
	// Find out more about verifying domains in Amazon SES
	// in the [AWS SES docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
	DkimTokens pulumi.StringArrayOutput `pulumi:"dkimTokens"`
	// Verified domain name to generate DKIM tokens for.
	Domain pulumi.StringOutput `pulumi:"domain"`
}

// NewDomainDkim registers a new resource with the given unique name, arguments, and options.
func NewDomainDkim(ctx *pulumi.Context,
	name string, args *DomainDkimArgs, opts ...pulumi.ResourceOption) (*DomainDkim, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	var resource DomainDkim
	err := ctx.RegisterResource("aws:ses/domainDkim:DomainDkim", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainDkim gets an existing DomainDkim resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainDkim(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainDkimState, opts ...pulumi.ResourceOption) (*DomainDkim, error) {
	var resource DomainDkim
	err := ctx.ReadResource("aws:ses/domainDkim:DomainDkim", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainDkim resources.
type domainDkimState struct {
	// DKIM tokens generated by SES.
	// These tokens should be used to create CNAME records used to verify SES Easy DKIM.
	// See below for an example of how this might be achieved
	// when the domain is hosted in Route 53 and managed by this provider.
	// Find out more about verifying domains in Amazon SES
	// in the [AWS SES docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
	DkimTokens []string `pulumi:"dkimTokens"`
	// Verified domain name to generate DKIM tokens for.
	Domain *string `pulumi:"domain"`
}

type DomainDkimState struct {
	// DKIM tokens generated by SES.
	// These tokens should be used to create CNAME records used to verify SES Easy DKIM.
	// See below for an example of how this might be achieved
	// when the domain is hosted in Route 53 and managed by this provider.
	// Find out more about verifying domains in Amazon SES
	// in the [AWS SES docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
	DkimTokens pulumi.StringArrayInput
	// Verified domain name to generate DKIM tokens for.
	Domain pulumi.StringPtrInput
}

func (DomainDkimState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainDkimState)(nil)).Elem()
}

type domainDkimArgs struct {
	// Verified domain name to generate DKIM tokens for.
	Domain string `pulumi:"domain"`
}

// The set of arguments for constructing a DomainDkim resource.
type DomainDkimArgs struct {
	// Verified domain name to generate DKIM tokens for.
	Domain pulumi.StringInput
}

func (DomainDkimArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainDkimArgs)(nil)).Elem()
}

type DomainDkimInput interface {
	pulumi.Input

	ToDomainDkimOutput() DomainDkimOutput
	ToDomainDkimOutputWithContext(ctx context.Context) DomainDkimOutput
}

func (*DomainDkim) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainDkim)(nil))
}

func (i *DomainDkim) ToDomainDkimOutput() DomainDkimOutput {
	return i.ToDomainDkimOutputWithContext(context.Background())
}

func (i *DomainDkim) ToDomainDkimOutputWithContext(ctx context.Context) DomainDkimOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainDkimOutput)
}

func (i *DomainDkim) ToDomainDkimPtrOutput() DomainDkimPtrOutput {
	return i.ToDomainDkimPtrOutputWithContext(context.Background())
}

func (i *DomainDkim) ToDomainDkimPtrOutputWithContext(ctx context.Context) DomainDkimPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainDkimPtrOutput)
}

type DomainDkimPtrInput interface {
	pulumi.Input

	ToDomainDkimPtrOutput() DomainDkimPtrOutput
	ToDomainDkimPtrOutputWithContext(ctx context.Context) DomainDkimPtrOutput
}

type domainDkimPtrType DomainDkimArgs

func (*domainDkimPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainDkim)(nil))
}

func (i *domainDkimPtrType) ToDomainDkimPtrOutput() DomainDkimPtrOutput {
	return i.ToDomainDkimPtrOutputWithContext(context.Background())
}

func (i *domainDkimPtrType) ToDomainDkimPtrOutputWithContext(ctx context.Context) DomainDkimPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainDkimPtrOutput)
}

// DomainDkimArrayInput is an input type that accepts DomainDkimArray and DomainDkimArrayOutput values.
// You can construct a concrete instance of `DomainDkimArrayInput` via:
//
//          DomainDkimArray{ DomainDkimArgs{...} }
type DomainDkimArrayInput interface {
	pulumi.Input

	ToDomainDkimArrayOutput() DomainDkimArrayOutput
	ToDomainDkimArrayOutputWithContext(context.Context) DomainDkimArrayOutput
}

type DomainDkimArray []DomainDkimInput

func (DomainDkimArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainDkim)(nil)).Elem()
}

func (i DomainDkimArray) ToDomainDkimArrayOutput() DomainDkimArrayOutput {
	return i.ToDomainDkimArrayOutputWithContext(context.Background())
}

func (i DomainDkimArray) ToDomainDkimArrayOutputWithContext(ctx context.Context) DomainDkimArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainDkimArrayOutput)
}

// DomainDkimMapInput is an input type that accepts DomainDkimMap and DomainDkimMapOutput values.
// You can construct a concrete instance of `DomainDkimMapInput` via:
//
//          DomainDkimMap{ "key": DomainDkimArgs{...} }
type DomainDkimMapInput interface {
	pulumi.Input

	ToDomainDkimMapOutput() DomainDkimMapOutput
	ToDomainDkimMapOutputWithContext(context.Context) DomainDkimMapOutput
}

type DomainDkimMap map[string]DomainDkimInput

func (DomainDkimMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainDkim)(nil)).Elem()
}

func (i DomainDkimMap) ToDomainDkimMapOutput() DomainDkimMapOutput {
	return i.ToDomainDkimMapOutputWithContext(context.Background())
}

func (i DomainDkimMap) ToDomainDkimMapOutputWithContext(ctx context.Context) DomainDkimMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainDkimMapOutput)
}

type DomainDkimOutput struct{ *pulumi.OutputState }

func (DomainDkimOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainDkim)(nil))
}

func (o DomainDkimOutput) ToDomainDkimOutput() DomainDkimOutput {
	return o
}

func (o DomainDkimOutput) ToDomainDkimOutputWithContext(ctx context.Context) DomainDkimOutput {
	return o
}

func (o DomainDkimOutput) ToDomainDkimPtrOutput() DomainDkimPtrOutput {
	return o.ToDomainDkimPtrOutputWithContext(context.Background())
}

func (o DomainDkimOutput) ToDomainDkimPtrOutputWithContext(ctx context.Context) DomainDkimPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DomainDkim) *DomainDkim {
		return &v
	}).(DomainDkimPtrOutput)
}

type DomainDkimPtrOutput struct{ *pulumi.OutputState }

func (DomainDkimPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainDkim)(nil))
}

func (o DomainDkimPtrOutput) ToDomainDkimPtrOutput() DomainDkimPtrOutput {
	return o
}

func (o DomainDkimPtrOutput) ToDomainDkimPtrOutputWithContext(ctx context.Context) DomainDkimPtrOutput {
	return o
}

func (o DomainDkimPtrOutput) Elem() DomainDkimOutput {
	return o.ApplyT(func(v *DomainDkim) DomainDkim {
		if v != nil {
			return *v
		}
		var ret DomainDkim
		return ret
	}).(DomainDkimOutput)
}

type DomainDkimArrayOutput struct{ *pulumi.OutputState }

func (DomainDkimArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainDkim)(nil))
}

func (o DomainDkimArrayOutput) ToDomainDkimArrayOutput() DomainDkimArrayOutput {
	return o
}

func (o DomainDkimArrayOutput) ToDomainDkimArrayOutputWithContext(ctx context.Context) DomainDkimArrayOutput {
	return o
}

func (o DomainDkimArrayOutput) Index(i pulumi.IntInput) DomainDkimOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainDkim {
		return vs[0].([]DomainDkim)[vs[1].(int)]
	}).(DomainDkimOutput)
}

type DomainDkimMapOutput struct{ *pulumi.OutputState }

func (DomainDkimMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DomainDkim)(nil))
}

func (o DomainDkimMapOutput) ToDomainDkimMapOutput() DomainDkimMapOutput {
	return o
}

func (o DomainDkimMapOutput) ToDomainDkimMapOutputWithContext(ctx context.Context) DomainDkimMapOutput {
	return o
}

func (o DomainDkimMapOutput) MapIndex(k pulumi.StringInput) DomainDkimOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DomainDkim {
		return vs[0].(map[string]DomainDkim)[vs[1].(string)]
	}).(DomainDkimOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainDkimOutput{})
	pulumi.RegisterOutputType(DomainDkimPtrOutput{})
	pulumi.RegisterOutputType(DomainDkimArrayOutput{})
	pulumi.RegisterOutputType(DomainDkimMapOutput{})
}
