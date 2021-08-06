// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkfirewall

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AWS Network Firewall Firewall Resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/networkfirewall"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networkfirewall.NewFirewall(ctx, "example", &networkfirewall.FirewallArgs{
// 			FirewallPolicyArn: pulumi.Any(aws_networkfirewall_firewall_policy.Example.Arn),
// 			VpcId:             pulumi.Any(aws_vpc.Example.Id),
// 			SubnetMappings: networkfirewall.FirewallSubnetMappingArray{
// 				&networkfirewall.FirewallSubnetMappingArgs{
// 					SubnetId: pulumi.Any(aws_subnet.Example.Id),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Tag1": pulumi.String("Value1"),
// 				"Tag2": pulumi.String("Value2"),
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
// Network Firewall Firewalls can be imported using their `ARN`.
//
// ```sh
//  $ pulumi import aws:networkfirewall/firewall:Firewall example arn:aws:network-firewall:us-west-1:123456789012:firewall/example
// ```
type Firewall struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) that identifies the firewall.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A boolean flag indicating whether it is possible to delete the firewall. Defaults to `false`.
	DeleteProtection pulumi.BoolPtrOutput `pulumi:"deleteProtection"`
	// A friendly description of the firewall.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the VPC Firewall policy.
	FirewallPolicyArn pulumi.StringOutput `pulumi:"firewallPolicyArn"`
	// A boolean flag indicating whether it is possible to change the associated firewall policy. Defaults to `false`.
	FirewallPolicyChangeProtection pulumi.BoolPtrOutput `pulumi:"firewallPolicyChangeProtection"`
	// Nested list of information about the current status of the firewall.
	FirewallStatuses FirewallFirewallStatusArrayOutput `pulumi:"firewallStatuses"`
	// A friendly name of the firewall.
	Name pulumi.StringOutput `pulumi:"name"`
	// A boolean flag indicating whether it is possible to change the associated subnet(s). Defaults to `false`.
	SubnetChangeProtection pulumi.BoolPtrOutput `pulumi:"subnetChangeProtection"`
	// Set of configuration blocks describing the public subnets. Each subnet must belong to a different Availability Zone in the VPC. AWS Network Firewall creates a firewall endpoint in each subnet. See Subnet Mapping below for details.
	SubnetMappings FirewallSubnetMappingArrayOutput `pulumi:"subnetMappings"`
	// Map of resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// A string token used when updating a firewall.
	UpdateToken pulumi.StringOutput `pulumi:"updateToken"`
	// The unique identifier of the VPC where AWS Network Firewall should create the firewall.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewFirewall registers a new resource with the given unique name, arguments, and options.
func NewFirewall(ctx *pulumi.Context,
	name string, args *FirewallArgs, opts ...pulumi.ResourceOption) (*Firewall, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FirewallPolicyArn == nil {
		return nil, errors.New("invalid value for required argument 'FirewallPolicyArn'")
	}
	if args.SubnetMappings == nil {
		return nil, errors.New("invalid value for required argument 'SubnetMappings'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource Firewall
	err := ctx.RegisterResource("aws:networkfirewall/firewall:Firewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewall gets an existing Firewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallState, opts ...pulumi.ResourceOption) (*Firewall, error) {
	var resource Firewall
	err := ctx.ReadResource("aws:networkfirewall/firewall:Firewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Firewall resources.
type firewallState struct {
	// The Amazon Resource Name (ARN) that identifies the firewall.
	Arn *string `pulumi:"arn"`
	// A boolean flag indicating whether it is possible to delete the firewall. Defaults to `false`.
	DeleteProtection *bool `pulumi:"deleteProtection"`
	// A friendly description of the firewall.
	Description *string `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the VPC Firewall policy.
	FirewallPolicyArn *string `pulumi:"firewallPolicyArn"`
	// A boolean flag indicating whether it is possible to change the associated firewall policy. Defaults to `false`.
	FirewallPolicyChangeProtection *bool `pulumi:"firewallPolicyChangeProtection"`
	// Nested list of information about the current status of the firewall.
	FirewallStatuses []FirewallFirewallStatus `pulumi:"firewallStatuses"`
	// A friendly name of the firewall.
	Name *string `pulumi:"name"`
	// A boolean flag indicating whether it is possible to change the associated subnet(s). Defaults to `false`.
	SubnetChangeProtection *bool `pulumi:"subnetChangeProtection"`
	// Set of configuration blocks describing the public subnets. Each subnet must belong to a different Availability Zone in the VPC. AWS Network Firewall creates a firewall endpoint in each subnet. See Subnet Mapping below for details.
	SubnetMappings []FirewallSubnetMapping `pulumi:"subnetMappings"`
	// Map of resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// A string token used when updating a firewall.
	UpdateToken *string `pulumi:"updateToken"`
	// The unique identifier of the VPC where AWS Network Firewall should create the firewall.
	VpcId *string `pulumi:"vpcId"`
}

type FirewallState struct {
	// The Amazon Resource Name (ARN) that identifies the firewall.
	Arn pulumi.StringPtrInput
	// A boolean flag indicating whether it is possible to delete the firewall. Defaults to `false`.
	DeleteProtection pulumi.BoolPtrInput
	// A friendly description of the firewall.
	Description pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the VPC Firewall policy.
	FirewallPolicyArn pulumi.StringPtrInput
	// A boolean flag indicating whether it is possible to change the associated firewall policy. Defaults to `false`.
	FirewallPolicyChangeProtection pulumi.BoolPtrInput
	// Nested list of information about the current status of the firewall.
	FirewallStatuses FirewallFirewallStatusArrayInput
	// A friendly name of the firewall.
	Name pulumi.StringPtrInput
	// A boolean flag indicating whether it is possible to change the associated subnet(s). Defaults to `false`.
	SubnetChangeProtection pulumi.BoolPtrInput
	// Set of configuration blocks describing the public subnets. Each subnet must belong to a different Availability Zone in the VPC. AWS Network Firewall creates a firewall endpoint in each subnet. See Subnet Mapping below for details.
	SubnetMappings FirewallSubnetMappingArrayInput
	// Map of resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// A string token used when updating a firewall.
	UpdateToken pulumi.StringPtrInput
	// The unique identifier of the VPC where AWS Network Firewall should create the firewall.
	VpcId pulumi.StringPtrInput
}

func (FirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallState)(nil)).Elem()
}

type firewallArgs struct {
	// A boolean flag indicating whether it is possible to delete the firewall. Defaults to `false`.
	DeleteProtection *bool `pulumi:"deleteProtection"`
	// A friendly description of the firewall.
	Description *string `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the VPC Firewall policy.
	FirewallPolicyArn string `pulumi:"firewallPolicyArn"`
	// A boolean flag indicating whether it is possible to change the associated firewall policy. Defaults to `false`.
	FirewallPolicyChangeProtection *bool `pulumi:"firewallPolicyChangeProtection"`
	// A friendly name of the firewall.
	Name *string `pulumi:"name"`
	// A boolean flag indicating whether it is possible to change the associated subnet(s). Defaults to `false`.
	SubnetChangeProtection *bool `pulumi:"subnetChangeProtection"`
	// Set of configuration blocks describing the public subnets. Each subnet must belong to a different Availability Zone in the VPC. AWS Network Firewall creates a firewall endpoint in each subnet. See Subnet Mapping below for details.
	SubnetMappings []FirewallSubnetMapping `pulumi:"subnetMappings"`
	// Map of resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The unique identifier of the VPC where AWS Network Firewall should create the firewall.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Firewall resource.
type FirewallArgs struct {
	// A boolean flag indicating whether it is possible to delete the firewall. Defaults to `false`.
	DeleteProtection pulumi.BoolPtrInput
	// A friendly description of the firewall.
	Description pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the VPC Firewall policy.
	FirewallPolicyArn pulumi.StringInput
	// A boolean flag indicating whether it is possible to change the associated firewall policy. Defaults to `false`.
	FirewallPolicyChangeProtection pulumi.BoolPtrInput
	// A friendly name of the firewall.
	Name pulumi.StringPtrInput
	// A boolean flag indicating whether it is possible to change the associated subnet(s). Defaults to `false`.
	SubnetChangeProtection pulumi.BoolPtrInput
	// Set of configuration blocks describing the public subnets. Each subnet must belong to a different Availability Zone in the VPC. AWS Network Firewall creates a firewall endpoint in each subnet. See Subnet Mapping below for details.
	SubnetMappings FirewallSubnetMappingArrayInput
	// Map of resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The unique identifier of the VPC where AWS Network Firewall should create the firewall.
	VpcId pulumi.StringInput
}

func (FirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallArgs)(nil)).Elem()
}

type FirewallInput interface {
	pulumi.Input

	ToFirewallOutput() FirewallOutput
	ToFirewallOutputWithContext(ctx context.Context) FirewallOutput
}

func (*Firewall) ElementType() reflect.Type {
	return reflect.TypeOf((*Firewall)(nil))
}

func (i *Firewall) ToFirewallOutput() FirewallOutput {
	return i.ToFirewallOutputWithContext(context.Background())
}

func (i *Firewall) ToFirewallOutputWithContext(ctx context.Context) FirewallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallOutput)
}

func (i *Firewall) ToFirewallPtrOutput() FirewallPtrOutput {
	return i.ToFirewallPtrOutputWithContext(context.Background())
}

func (i *Firewall) ToFirewallPtrOutputWithContext(ctx context.Context) FirewallPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPtrOutput)
}

type FirewallPtrInput interface {
	pulumi.Input

	ToFirewallPtrOutput() FirewallPtrOutput
	ToFirewallPtrOutputWithContext(ctx context.Context) FirewallPtrOutput
}

type firewallPtrType FirewallArgs

func (*firewallPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Firewall)(nil))
}

func (i *firewallPtrType) ToFirewallPtrOutput() FirewallPtrOutput {
	return i.ToFirewallPtrOutputWithContext(context.Background())
}

func (i *firewallPtrType) ToFirewallPtrOutputWithContext(ctx context.Context) FirewallPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPtrOutput)
}

// FirewallArrayInput is an input type that accepts FirewallArray and FirewallArrayOutput values.
// You can construct a concrete instance of `FirewallArrayInput` via:
//
//          FirewallArray{ FirewallArgs{...} }
type FirewallArrayInput interface {
	pulumi.Input

	ToFirewallArrayOutput() FirewallArrayOutput
	ToFirewallArrayOutputWithContext(context.Context) FirewallArrayOutput
}

type FirewallArray []FirewallInput

func (FirewallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Firewall)(nil)).Elem()
}

func (i FirewallArray) ToFirewallArrayOutput() FirewallArrayOutput {
	return i.ToFirewallArrayOutputWithContext(context.Background())
}

func (i FirewallArray) ToFirewallArrayOutputWithContext(ctx context.Context) FirewallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallArrayOutput)
}

// FirewallMapInput is an input type that accepts FirewallMap and FirewallMapOutput values.
// You can construct a concrete instance of `FirewallMapInput` via:
//
//          FirewallMap{ "key": FirewallArgs{...} }
type FirewallMapInput interface {
	pulumi.Input

	ToFirewallMapOutput() FirewallMapOutput
	ToFirewallMapOutputWithContext(context.Context) FirewallMapOutput
}

type FirewallMap map[string]FirewallInput

func (FirewallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Firewall)(nil)).Elem()
}

func (i FirewallMap) ToFirewallMapOutput() FirewallMapOutput {
	return i.ToFirewallMapOutputWithContext(context.Background())
}

func (i FirewallMap) ToFirewallMapOutputWithContext(ctx context.Context) FirewallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallMapOutput)
}

type FirewallOutput struct {
	*pulumi.OutputState
}

func (FirewallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Firewall)(nil))
}

func (o FirewallOutput) ToFirewallOutput() FirewallOutput {
	return o
}

func (o FirewallOutput) ToFirewallOutputWithContext(ctx context.Context) FirewallOutput {
	return o
}

func (o FirewallOutput) ToFirewallPtrOutput() FirewallPtrOutput {
	return o.ToFirewallPtrOutputWithContext(context.Background())
}

func (o FirewallOutput) ToFirewallPtrOutputWithContext(ctx context.Context) FirewallPtrOutput {
	return o.ApplyT(func(v Firewall) *Firewall {
		return &v
	}).(FirewallPtrOutput)
}

type FirewallPtrOutput struct {
	*pulumi.OutputState
}

func (FirewallPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Firewall)(nil))
}

func (o FirewallPtrOutput) ToFirewallPtrOutput() FirewallPtrOutput {
	return o
}

func (o FirewallPtrOutput) ToFirewallPtrOutputWithContext(ctx context.Context) FirewallPtrOutput {
	return o
}

type FirewallArrayOutput struct{ *pulumi.OutputState }

func (FirewallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Firewall)(nil))
}

func (o FirewallArrayOutput) ToFirewallArrayOutput() FirewallArrayOutput {
	return o
}

func (o FirewallArrayOutput) ToFirewallArrayOutputWithContext(ctx context.Context) FirewallArrayOutput {
	return o
}

func (o FirewallArrayOutput) Index(i pulumi.IntInput) FirewallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Firewall {
		return vs[0].([]Firewall)[vs[1].(int)]
	}).(FirewallOutput)
}

type FirewallMapOutput struct{ *pulumi.OutputState }

func (FirewallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Firewall)(nil))
}

func (o FirewallMapOutput) ToFirewallMapOutput() FirewallMapOutput {
	return o
}

func (o FirewallMapOutput) ToFirewallMapOutputWithContext(ctx context.Context) FirewallMapOutput {
	return o
}

func (o FirewallMapOutput) MapIndex(k pulumi.StringInput) FirewallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Firewall {
		return vs[0].(map[string]Firewall)[vs[1].(string)]
	}).(FirewallOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallOutput{})
	pulumi.RegisterOutputType(FirewallPtrOutput{})
	pulumi.RegisterOutputType(FirewallArrayOutput{})
	pulumi.RegisterOutputType(FirewallMapOutput{})
}
