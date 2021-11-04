// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Amazon Redshift security group. You use security groups to control access to non-VPC clusters
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/redshift"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := redshift.NewSecurityGroup(ctx, "_default", &redshift.SecurityGroupArgs{
// 			Ingress: redshift.SecurityGroupIngressArray{
// 				&redshift.SecurityGroupIngressArgs{
// 					Cidr: pulumi.String("10.0.0.0/24"),
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
// Redshift security groups can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:redshift/securityGroup:SecurityGroup testgroup1 redshift_test_group
// ```
type SecurityGroup struct {
	pulumi.CustomResourceState

	// The description of the Redshift security group. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`
	// A list of ingress rules.
	Ingress SecurityGroupIngressArrayOutput `pulumi:"ingress"`
	// The name of the Redshift security group.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroup(ctx *pulumi.Context,
	name string, args *SecurityGroupArgs, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ingress == nil {
		return nil, errors.New("invalid value for required argument 'Ingress'")
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource SecurityGroup
	err := ctx.RegisterResource("aws:redshift/securityGroup:SecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityGroup gets an existing SecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityGroupState, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	var resource SecurityGroup
	err := ctx.ReadResource("aws:redshift/securityGroup:SecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroup resources.
type securityGroupState struct {
	// The description of the Redshift security group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// A list of ingress rules.
	Ingress []SecurityGroupIngress `pulumi:"ingress"`
	// The name of the Redshift security group.
	Name *string `pulumi:"name"`
}

type SecurityGroupState struct {
	// The description of the Redshift security group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// A list of ingress rules.
	Ingress SecurityGroupIngressArrayInput
	// The name of the Redshift security group.
	Name pulumi.StringPtrInput
}

func (SecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupState)(nil)).Elem()
}

type securityGroupArgs struct {
	// The description of the Redshift security group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// A list of ingress rules.
	Ingress []SecurityGroupIngress `pulumi:"ingress"`
	// The name of the Redshift security group.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a SecurityGroup resource.
type SecurityGroupArgs struct {
	// The description of the Redshift security group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// A list of ingress rules.
	Ingress SecurityGroupIngressArrayInput
	// The name of the Redshift security group.
	Name pulumi.StringPtrInput
}

func (SecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupArgs)(nil)).Elem()
}

type SecurityGroupInput interface {
	pulumi.Input

	ToSecurityGroupOutput() SecurityGroupOutput
	ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput
}

func (*SecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroup)(nil))
}

func (i *SecurityGroup) ToSecurityGroupOutput() SecurityGroupOutput {
	return i.ToSecurityGroupOutputWithContext(context.Background())
}

func (i *SecurityGroup) ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupOutput)
}

func (i *SecurityGroup) ToSecurityGroupPtrOutput() SecurityGroupPtrOutput {
	return i.ToSecurityGroupPtrOutputWithContext(context.Background())
}

func (i *SecurityGroup) ToSecurityGroupPtrOutputWithContext(ctx context.Context) SecurityGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupPtrOutput)
}

type SecurityGroupPtrInput interface {
	pulumi.Input

	ToSecurityGroupPtrOutput() SecurityGroupPtrOutput
	ToSecurityGroupPtrOutputWithContext(ctx context.Context) SecurityGroupPtrOutput
}

type securityGroupPtrType SecurityGroupArgs

func (*securityGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroup)(nil))
}

func (i *securityGroupPtrType) ToSecurityGroupPtrOutput() SecurityGroupPtrOutput {
	return i.ToSecurityGroupPtrOutputWithContext(context.Background())
}

func (i *securityGroupPtrType) ToSecurityGroupPtrOutputWithContext(ctx context.Context) SecurityGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupPtrOutput)
}

// SecurityGroupArrayInput is an input type that accepts SecurityGroupArray and SecurityGroupArrayOutput values.
// You can construct a concrete instance of `SecurityGroupArrayInput` via:
//
//          SecurityGroupArray{ SecurityGroupArgs{...} }
type SecurityGroupArrayInput interface {
	pulumi.Input

	ToSecurityGroupArrayOutput() SecurityGroupArrayOutput
	ToSecurityGroupArrayOutputWithContext(context.Context) SecurityGroupArrayOutput
}

type SecurityGroupArray []SecurityGroupInput

func (SecurityGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityGroup)(nil)).Elem()
}

func (i SecurityGroupArray) ToSecurityGroupArrayOutput() SecurityGroupArrayOutput {
	return i.ToSecurityGroupArrayOutputWithContext(context.Background())
}

func (i SecurityGroupArray) ToSecurityGroupArrayOutputWithContext(ctx context.Context) SecurityGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupArrayOutput)
}

// SecurityGroupMapInput is an input type that accepts SecurityGroupMap and SecurityGroupMapOutput values.
// You can construct a concrete instance of `SecurityGroupMapInput` via:
//
//          SecurityGroupMap{ "key": SecurityGroupArgs{...} }
type SecurityGroupMapInput interface {
	pulumi.Input

	ToSecurityGroupMapOutput() SecurityGroupMapOutput
	ToSecurityGroupMapOutputWithContext(context.Context) SecurityGroupMapOutput
}

type SecurityGroupMap map[string]SecurityGroupInput

func (SecurityGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityGroup)(nil)).Elem()
}

func (i SecurityGroupMap) ToSecurityGroupMapOutput() SecurityGroupMapOutput {
	return i.ToSecurityGroupMapOutputWithContext(context.Background())
}

func (i SecurityGroupMap) ToSecurityGroupMapOutputWithContext(ctx context.Context) SecurityGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupMapOutput)
}

type SecurityGroupOutput struct{ *pulumi.OutputState }

func (SecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroup)(nil))
}

func (o SecurityGroupOutput) ToSecurityGroupOutput() SecurityGroupOutput {
	return o
}

func (o SecurityGroupOutput) ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput {
	return o
}

func (o SecurityGroupOutput) ToSecurityGroupPtrOutput() SecurityGroupPtrOutput {
	return o.ToSecurityGroupPtrOutputWithContext(context.Background())
}

func (o SecurityGroupOutput) ToSecurityGroupPtrOutputWithContext(ctx context.Context) SecurityGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityGroup) *SecurityGroup {
		return &v
	}).(SecurityGroupPtrOutput)
}

type SecurityGroupPtrOutput struct{ *pulumi.OutputState }

func (SecurityGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroup)(nil))
}

func (o SecurityGroupPtrOutput) ToSecurityGroupPtrOutput() SecurityGroupPtrOutput {
	return o
}

func (o SecurityGroupPtrOutput) ToSecurityGroupPtrOutputWithContext(ctx context.Context) SecurityGroupPtrOutput {
	return o
}

func (o SecurityGroupPtrOutput) Elem() SecurityGroupOutput {
	return o.ApplyT(func(v *SecurityGroup) SecurityGroup {
		if v != nil {
			return *v
		}
		var ret SecurityGroup
		return ret
	}).(SecurityGroupOutput)
}

type SecurityGroupArrayOutput struct{ *pulumi.OutputState }

func (SecurityGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityGroup)(nil))
}

func (o SecurityGroupArrayOutput) ToSecurityGroupArrayOutput() SecurityGroupArrayOutput {
	return o
}

func (o SecurityGroupArrayOutput) ToSecurityGroupArrayOutputWithContext(ctx context.Context) SecurityGroupArrayOutput {
	return o
}

func (o SecurityGroupArrayOutput) Index(i pulumi.IntInput) SecurityGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityGroup {
		return vs[0].([]SecurityGroup)[vs[1].(int)]
	}).(SecurityGroupOutput)
}

type SecurityGroupMapOutput struct{ *pulumi.OutputState }

func (SecurityGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SecurityGroup)(nil))
}

func (o SecurityGroupMapOutput) ToSecurityGroupMapOutput() SecurityGroupMapOutput {
	return o
}

func (o SecurityGroupMapOutput) ToSecurityGroupMapOutputWithContext(ctx context.Context) SecurityGroupMapOutput {
	return o
}

func (o SecurityGroupMapOutput) MapIndex(k pulumi.StringInput) SecurityGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SecurityGroup {
		return vs[0].(map[string]SecurityGroup)[vs[1].(string)]
	}).(SecurityGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupInput)(nil)).Elem(), &SecurityGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupPtrInput)(nil)).Elem(), &SecurityGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupArrayInput)(nil)).Elem(), SecurityGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupMapInput)(nil)).Elem(), SecurityGroupMap{})
	pulumi.RegisterOutputType(SecurityGroupOutput{})
	pulumi.RegisterOutputType(SecurityGroupPtrOutput{})
	pulumi.RegisterOutputType(SecurityGroupArrayOutput{})
	pulumi.RegisterOutputType(SecurityGroupMapOutput{})
}
