// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a domain resource for the specified domain (e.g., example.com).
// You cannot register a new domain name using Lightsail. You must register
// a domain name using Amazon Route 53 or another domain name registrar.
// If you have already registered your domain, you can enter its name in
// this parameter to manage the DNS records for that domain.
//
// > **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lightsail"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := lightsail.NewDomain(ctx, "domainTest", &lightsail.DomainArgs{
// 			DomainName: pulumi.String("mydomain.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Domain struct {
	pulumi.CustomResourceState

	// The ARN of the Lightsail domain
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the Lightsail domain to manage
	DomainName pulumi.StringOutput `pulumi:"domainName"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	var resource Domain
	err := ctx.RegisterResource("aws:lightsail/domain:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("aws:lightsail/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// The ARN of the Lightsail domain
	Arn *string `pulumi:"arn"`
	// The name of the Lightsail domain to manage
	DomainName *string `pulumi:"domainName"`
}

type DomainState struct {
	// The ARN of the Lightsail domain
	Arn pulumi.StringPtrInput
	// The name of the Lightsail domain to manage
	DomainName pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// The name of the Lightsail domain to manage
	DomainName string `pulumi:"domainName"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// The name of the Lightsail domain to manage
	DomainName pulumi.StringInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil))
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

func (i *Domain) ToDomainPtrOutput() DomainPtrOutput {
	return i.ToDomainPtrOutputWithContext(context.Background())
}

func (i *Domain) ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPtrOutput)
}

type DomainPtrInput interface {
	pulumi.Input

	ToDomainPtrOutput() DomainPtrOutput
	ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput
}

type domainPtrType DomainArgs

func (*domainPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil))
}

func (i *domainPtrType) ToDomainPtrOutput() DomainPtrOutput {
	return i.ToDomainPtrOutputWithContext(context.Background())
}

func (i *domainPtrType) ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPtrOutput)
}

// DomainArrayInput is an input type that accepts DomainArray and DomainArrayOutput values.
// You can construct a concrete instance of `DomainArrayInput` via:
//
//          DomainArray{ DomainArgs{...} }
type DomainArrayInput interface {
	pulumi.Input

	ToDomainArrayOutput() DomainArrayOutput
	ToDomainArrayOutputWithContext(context.Context) DomainArrayOutput
}

type DomainArray []DomainInput

func (DomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Domain)(nil)).Elem()
}

func (i DomainArray) ToDomainArrayOutput() DomainArrayOutput {
	return i.ToDomainArrayOutputWithContext(context.Background())
}

func (i DomainArray) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainArrayOutput)
}

// DomainMapInput is an input type that accepts DomainMap and DomainMapOutput values.
// You can construct a concrete instance of `DomainMapInput` via:
//
//          DomainMap{ "key": DomainArgs{...} }
type DomainMapInput interface {
	pulumi.Input

	ToDomainMapOutput() DomainMapOutput
	ToDomainMapOutputWithContext(context.Context) DomainMapOutput
}

type DomainMap map[string]DomainInput

func (DomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Domain)(nil)).Elem()
}

func (i DomainMap) ToDomainMapOutput() DomainMapOutput {
	return i.ToDomainMapOutputWithContext(context.Background())
}

func (i DomainMap) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMapOutput)
}

type DomainOutput struct {
	*pulumi.OutputState
}

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil))
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

func (o DomainOutput) ToDomainPtrOutput() DomainPtrOutput {
	return o.ToDomainPtrOutputWithContext(context.Background())
}

func (o DomainOutput) ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput {
	return o.ApplyT(func(v Domain) *Domain {
		return &v
	}).(DomainPtrOutput)
}

type DomainPtrOutput struct {
	*pulumi.OutputState
}

func (DomainPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil))
}

func (o DomainPtrOutput) ToDomainPtrOutput() DomainPtrOutput {
	return o
}

func (o DomainPtrOutput) ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput {
	return o
}

type DomainArrayOutput struct{ *pulumi.OutputState }

func (DomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Domain)(nil))
}

func (o DomainArrayOutput) ToDomainArrayOutput() DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) Index(i pulumi.IntInput) DomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Domain {
		return vs[0].([]Domain)[vs[1].(int)]
	}).(DomainOutput)
}

type DomainMapOutput struct{ *pulumi.OutputState }

func (DomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Domain)(nil))
}

func (o DomainMapOutput) ToDomainMapOutput() DomainMapOutput {
	return o
}

func (o DomainMapOutput) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return o
}

func (o DomainMapOutput) MapIndex(k pulumi.StringInput) DomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Domain {
		return vs[0].(map[string]Domain)[vs[1].(string)]
	}).(DomainOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
	pulumi.RegisterOutputType(DomainPtrOutput{})
	pulumi.RegisterOutputType(DomainArrayOutput{})
	pulumi.RegisterOutputType(DomainMapOutput{})
}
