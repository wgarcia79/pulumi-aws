// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an ElastiCache Security Group to control access to one or more cache
// clusters.
//
// > **NOTE:** ElastiCache Security Groups are for use only when working with an
// ElastiCache cluster **outside** of a VPC. If you are using a VPC, see the
// ElastiCache Subnet Group resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/elasticache"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		barSecurityGroup, err := ec2.NewSecurityGroup(ctx, "barSecurityGroup", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elasticache.NewSecurityGroup(ctx, "barElasticache_securityGroupSecurityGroup", &elasticache.SecurityGroupArgs{
// 			SecurityGroupNames: pulumi.StringArray{
// 				barSecurityGroup.Name,
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
// ElastiCache Security Groups can be imported by name, e.g.
//
// ```sh
//  $ pulumi import aws:elasticache/securityGroup:SecurityGroup my_ec_security_group ec-security-group-1
// ```
type SecurityGroup struct {
	pulumi.CustomResourceState

	// description for the cache security group. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`
	// Name for the cache security group. This value is stored as a lowercase string.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of EC2 security group names to be
	// authorized for ingress to the cache security group
	SecurityGroupNames pulumi.StringArrayOutput `pulumi:"securityGroupNames"`
}

// NewSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroup(ctx *pulumi.Context,
	name string, args *SecurityGroupArgs, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SecurityGroupNames == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupNames'")
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource SecurityGroup
	err := ctx.RegisterResource("aws:elasticache/securityGroup:SecurityGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:elasticache/securityGroup:SecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroup resources.
type securityGroupState struct {
	// description for the cache security group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// Name for the cache security group. This value is stored as a lowercase string.
	Name *string `pulumi:"name"`
	// List of EC2 security group names to be
	// authorized for ingress to the cache security group
	SecurityGroupNames []string `pulumi:"securityGroupNames"`
}

type SecurityGroupState struct {
	// description for the cache security group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// Name for the cache security group. This value is stored as a lowercase string.
	Name pulumi.StringPtrInput
	// List of EC2 security group names to be
	// authorized for ingress to the cache security group
	SecurityGroupNames pulumi.StringArrayInput
}

func (SecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupState)(nil)).Elem()
}

type securityGroupArgs struct {
	// description for the cache security group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// Name for the cache security group. This value is stored as a lowercase string.
	Name *string `pulumi:"name"`
	// List of EC2 security group names to be
	// authorized for ingress to the cache security group
	SecurityGroupNames []string `pulumi:"securityGroupNames"`
}

// The set of arguments for constructing a SecurityGroup resource.
type SecurityGroupArgs struct {
	// description for the cache security group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// Name for the cache security group. This value is stored as a lowercase string.
	Name pulumi.StringPtrInput
	// List of EC2 security group names to be
	// authorized for ingress to the cache security group
	SecurityGroupNames pulumi.StringArrayInput
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
	return reflect.TypeOf(([]*SecurityGroup)(nil))
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
	return reflect.TypeOf((map[string]*SecurityGroup)(nil))
}

func (i SecurityGroupMap) ToSecurityGroupMapOutput() SecurityGroupMapOutput {
	return i.ToSecurityGroupMapOutputWithContext(context.Background())
}

func (i SecurityGroupMap) ToSecurityGroupMapOutputWithContext(ctx context.Context) SecurityGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupMapOutput)
}

type SecurityGroupOutput struct {
	*pulumi.OutputState
}

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
	return o.ApplyT(func(v SecurityGroup) *SecurityGroup {
		return &v
	}).(SecurityGroupPtrOutput)
}

type SecurityGroupPtrOutput struct {
	*pulumi.OutputState
}

func (SecurityGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroup)(nil))
}

func (o SecurityGroupPtrOutput) ToSecurityGroupPtrOutput() SecurityGroupPtrOutput {
	return o
}

func (o SecurityGroupPtrOutput) ToSecurityGroupPtrOutputWithContext(ctx context.Context) SecurityGroupPtrOutput {
	return o
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
	pulumi.RegisterOutputType(SecurityGroupOutput{})
	pulumi.RegisterOutputType(SecurityGroupPtrOutput{})
	pulumi.RegisterOutputType(SecurityGroupArrayOutput{})
	pulumi.RegisterOutputType(SecurityGroupMapOutput{})
}
