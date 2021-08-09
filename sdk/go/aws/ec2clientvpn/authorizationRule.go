// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2clientvpn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides authorization rules for AWS Client VPN endpoints. For more information on usage, please see the
// [AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2clientvpn"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2clientvpn.NewAuthorizationRule(ctx, "example", &ec2clientvpn.AuthorizationRuleArgs{
// 			ClientVpnEndpointId: pulumi.Any(aws_ec2_client_vpn_endpoint.Example.Id),
// 			TargetNetworkCidr:   pulumi.Any(aws_subnet.Example.Cidr_block),
// 			AuthorizeAllGroups:  pulumi.Bool(true),
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
// AWS Client VPN authorization rules can be imported using the endpoint ID and target network CIDR. If there is a specific group name that is included as well. All values are separated by a `,`.
//
// ```sh
//  $ pulumi import aws:ec2clientvpn/authorizationRule:AuthorizationRule example cvpn-endpoint-0ac3a1abbccddd666,10.1.0.0/24
// ```
//
// ```sh
//  $ pulumi import aws:ec2clientvpn/authorizationRule:AuthorizationRule example cvpn-endpoint-0ac3a1abbccddd666,10.1.0.0/24,team-a
// ```
type AuthorizationRule struct {
	pulumi.CustomResourceState

	// The ID of the group to which the authorization rule grants access. One of `accessGroupId` or `authorizeAllGroups` must be set.
	AccessGroupId pulumi.StringPtrOutput `pulumi:"accessGroupId"`
	// Indicates whether the authorization rule grants access to all clients. One of `accessGroupId` or `authorizeAllGroups` must be set.
	AuthorizeAllGroups pulumi.BoolPtrOutput `pulumi:"authorizeAllGroups"`
	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId pulumi.StringOutput `pulumi:"clientVpnEndpointId"`
	// A brief description of the authorization rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
	TargetNetworkCidr pulumi.StringOutput `pulumi:"targetNetworkCidr"`
}

// NewAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewAuthorizationRule(ctx *pulumi.Context,
	name string, args *AuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*AuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientVpnEndpointId == nil {
		return nil, errors.New("invalid value for required argument 'ClientVpnEndpointId'")
	}
	if args.TargetNetworkCidr == nil {
		return nil, errors.New("invalid value for required argument 'TargetNetworkCidr'")
	}
	var resource AuthorizationRule
	err := ctx.RegisterResource("aws:ec2clientvpn/authorizationRule:AuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorizationRule gets an existing AuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizationRuleState, opts ...pulumi.ResourceOption) (*AuthorizationRule, error) {
	var resource AuthorizationRule
	err := ctx.ReadResource("aws:ec2clientvpn/authorizationRule:AuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthorizationRule resources.
type authorizationRuleState struct {
	// The ID of the group to which the authorization rule grants access. One of `accessGroupId` or `authorizeAllGroups` must be set.
	AccessGroupId *string `pulumi:"accessGroupId"`
	// Indicates whether the authorization rule grants access to all clients. One of `accessGroupId` or `authorizeAllGroups` must be set.
	AuthorizeAllGroups *bool `pulumi:"authorizeAllGroups"`
	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId *string `pulumi:"clientVpnEndpointId"`
	// A brief description of the authorization rule.
	Description *string `pulumi:"description"`
	// The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
	TargetNetworkCidr *string `pulumi:"targetNetworkCidr"`
}

type AuthorizationRuleState struct {
	// The ID of the group to which the authorization rule grants access. One of `accessGroupId` or `authorizeAllGroups` must be set.
	AccessGroupId pulumi.StringPtrInput
	// Indicates whether the authorization rule grants access to all clients. One of `accessGroupId` or `authorizeAllGroups` must be set.
	AuthorizeAllGroups pulumi.BoolPtrInput
	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId pulumi.StringPtrInput
	// A brief description of the authorization rule.
	Description pulumi.StringPtrInput
	// The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
	TargetNetworkCidr pulumi.StringPtrInput
}

func (AuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationRuleState)(nil)).Elem()
}

type authorizationRuleArgs struct {
	// The ID of the group to which the authorization rule grants access. One of `accessGroupId` or `authorizeAllGroups` must be set.
	AccessGroupId *string `pulumi:"accessGroupId"`
	// Indicates whether the authorization rule grants access to all clients. One of `accessGroupId` or `authorizeAllGroups` must be set.
	AuthorizeAllGroups *bool `pulumi:"authorizeAllGroups"`
	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId string `pulumi:"clientVpnEndpointId"`
	// A brief description of the authorization rule.
	Description *string `pulumi:"description"`
	// The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
	TargetNetworkCidr string `pulumi:"targetNetworkCidr"`
}

// The set of arguments for constructing a AuthorizationRule resource.
type AuthorizationRuleArgs struct {
	// The ID of the group to which the authorization rule grants access. One of `accessGroupId` or `authorizeAllGroups` must be set.
	AccessGroupId pulumi.StringPtrInput
	// Indicates whether the authorization rule grants access to all clients. One of `accessGroupId` or `authorizeAllGroups` must be set.
	AuthorizeAllGroups pulumi.BoolPtrInput
	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId pulumi.StringInput
	// A brief description of the authorization rule.
	Description pulumi.StringPtrInput
	// The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
	TargetNetworkCidr pulumi.StringInput
}

func (AuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationRuleArgs)(nil)).Elem()
}

type AuthorizationRuleInput interface {
	pulumi.Input

	ToAuthorizationRuleOutput() AuthorizationRuleOutput
	ToAuthorizationRuleOutputWithContext(ctx context.Context) AuthorizationRuleOutput
}

func (*AuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationRule)(nil))
}

func (i *AuthorizationRule) ToAuthorizationRuleOutput() AuthorizationRuleOutput {
	return i.ToAuthorizationRuleOutputWithContext(context.Background())
}

func (i *AuthorizationRule) ToAuthorizationRuleOutputWithContext(ctx context.Context) AuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationRuleOutput)
}

func (i *AuthorizationRule) ToAuthorizationRulePtrOutput() AuthorizationRulePtrOutput {
	return i.ToAuthorizationRulePtrOutputWithContext(context.Background())
}

func (i *AuthorizationRule) ToAuthorizationRulePtrOutputWithContext(ctx context.Context) AuthorizationRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationRulePtrOutput)
}

type AuthorizationRulePtrInput interface {
	pulumi.Input

	ToAuthorizationRulePtrOutput() AuthorizationRulePtrOutput
	ToAuthorizationRulePtrOutputWithContext(ctx context.Context) AuthorizationRulePtrOutput
}

type authorizationRulePtrType AuthorizationRuleArgs

func (*authorizationRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationRule)(nil))
}

func (i *authorizationRulePtrType) ToAuthorizationRulePtrOutput() AuthorizationRulePtrOutput {
	return i.ToAuthorizationRulePtrOutputWithContext(context.Background())
}

func (i *authorizationRulePtrType) ToAuthorizationRulePtrOutputWithContext(ctx context.Context) AuthorizationRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationRulePtrOutput)
}

// AuthorizationRuleArrayInput is an input type that accepts AuthorizationRuleArray and AuthorizationRuleArrayOutput values.
// You can construct a concrete instance of `AuthorizationRuleArrayInput` via:
//
//          AuthorizationRuleArray{ AuthorizationRuleArgs{...} }
type AuthorizationRuleArrayInput interface {
	pulumi.Input

	ToAuthorizationRuleArrayOutput() AuthorizationRuleArrayOutput
	ToAuthorizationRuleArrayOutputWithContext(context.Context) AuthorizationRuleArrayOutput
}

type AuthorizationRuleArray []AuthorizationRuleInput

func (AuthorizationRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthorizationRule)(nil)).Elem()
}

func (i AuthorizationRuleArray) ToAuthorizationRuleArrayOutput() AuthorizationRuleArrayOutput {
	return i.ToAuthorizationRuleArrayOutputWithContext(context.Background())
}

func (i AuthorizationRuleArray) ToAuthorizationRuleArrayOutputWithContext(ctx context.Context) AuthorizationRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationRuleArrayOutput)
}

// AuthorizationRuleMapInput is an input type that accepts AuthorizationRuleMap and AuthorizationRuleMapOutput values.
// You can construct a concrete instance of `AuthorizationRuleMapInput` via:
//
//          AuthorizationRuleMap{ "key": AuthorizationRuleArgs{...} }
type AuthorizationRuleMapInput interface {
	pulumi.Input

	ToAuthorizationRuleMapOutput() AuthorizationRuleMapOutput
	ToAuthorizationRuleMapOutputWithContext(context.Context) AuthorizationRuleMapOutput
}

type AuthorizationRuleMap map[string]AuthorizationRuleInput

func (AuthorizationRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthorizationRule)(nil)).Elem()
}

func (i AuthorizationRuleMap) ToAuthorizationRuleMapOutput() AuthorizationRuleMapOutput {
	return i.ToAuthorizationRuleMapOutputWithContext(context.Background())
}

func (i AuthorizationRuleMap) ToAuthorizationRuleMapOutputWithContext(ctx context.Context) AuthorizationRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationRuleMapOutput)
}

type AuthorizationRuleOutput struct{ *pulumi.OutputState }

func (AuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationRule)(nil))
}

func (o AuthorizationRuleOutput) ToAuthorizationRuleOutput() AuthorizationRuleOutput {
	return o
}

func (o AuthorizationRuleOutput) ToAuthorizationRuleOutputWithContext(ctx context.Context) AuthorizationRuleOutput {
	return o
}

func (o AuthorizationRuleOutput) ToAuthorizationRulePtrOutput() AuthorizationRulePtrOutput {
	return o.ToAuthorizationRulePtrOutputWithContext(context.Background())
}

func (o AuthorizationRuleOutput) ToAuthorizationRulePtrOutputWithContext(ctx context.Context) AuthorizationRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthorizationRule) *AuthorizationRule {
		return &v
	}).(AuthorizationRulePtrOutput)
}

type AuthorizationRulePtrOutput struct{ *pulumi.OutputState }

func (AuthorizationRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationRule)(nil))
}

func (o AuthorizationRulePtrOutput) ToAuthorizationRulePtrOutput() AuthorizationRulePtrOutput {
	return o
}

func (o AuthorizationRulePtrOutput) ToAuthorizationRulePtrOutputWithContext(ctx context.Context) AuthorizationRulePtrOutput {
	return o
}

func (o AuthorizationRulePtrOutput) Elem() AuthorizationRuleOutput {
	return o.ApplyT(func(v *AuthorizationRule) AuthorizationRule {
		if v != nil {
			return *v
		}
		var ret AuthorizationRule
		return ret
	}).(AuthorizationRuleOutput)
}

type AuthorizationRuleArrayOutput struct{ *pulumi.OutputState }

func (AuthorizationRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthorizationRule)(nil))
}

func (o AuthorizationRuleArrayOutput) ToAuthorizationRuleArrayOutput() AuthorizationRuleArrayOutput {
	return o
}

func (o AuthorizationRuleArrayOutput) ToAuthorizationRuleArrayOutputWithContext(ctx context.Context) AuthorizationRuleArrayOutput {
	return o
}

func (o AuthorizationRuleArrayOutput) Index(i pulumi.IntInput) AuthorizationRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthorizationRule {
		return vs[0].([]AuthorizationRule)[vs[1].(int)]
	}).(AuthorizationRuleOutput)
}

type AuthorizationRuleMapOutput struct{ *pulumi.OutputState }

func (AuthorizationRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AuthorizationRule)(nil))
}

func (o AuthorizationRuleMapOutput) ToAuthorizationRuleMapOutput() AuthorizationRuleMapOutput {
	return o
}

func (o AuthorizationRuleMapOutput) ToAuthorizationRuleMapOutputWithContext(ctx context.Context) AuthorizationRuleMapOutput {
	return o
}

func (o AuthorizationRuleMapOutput) MapIndex(k pulumi.StringInput) AuthorizationRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AuthorizationRule {
		return vs[0].(map[string]AuthorizationRule)[vs[1].(string)]
	}).(AuthorizationRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizationRuleOutput{})
	pulumi.RegisterOutputType(AuthorizationRulePtrOutput{})
	pulumi.RegisterOutputType(AuthorizationRuleArrayOutput{})
	pulumi.RegisterOutputType(AuthorizationRuleMapOutput{})
}
