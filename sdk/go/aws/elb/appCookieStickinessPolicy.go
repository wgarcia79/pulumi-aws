// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an application cookie stickiness policy, which allows an ELB to wed its sticky cookie's expiration to a cookie generated by your application.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/elb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		lb, err := elb.NewLoadBalancer(ctx, "lb", &elb.LoadBalancerArgs{
// 			AvailabilityZones: pulumi.StringArray{
// 				pulumi.String("us-east-1a"),
// 			},
// 			Listeners: elb.LoadBalancerListenerArray{
// 				&elb.LoadBalancerListenerArgs{
// 					InstancePort:     pulumi.Int(8000),
// 					InstanceProtocol: pulumi.String("http"),
// 					LbPort:           pulumi.Int(80),
// 					LbProtocol:       pulumi.String("http"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elb.NewAppCookieStickinessPolicy(ctx, "foo", &elb.AppCookieStickinessPolicyArgs{
// 			LoadBalancer: lb.Name,
// 			LbPort:       pulumi.Int(80),
// 			CookieName:   pulumi.String("MyAppCookie"),
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
// Application cookie stickiness policies can be imported using the ELB name, port, and policy name separated by colons (`:`), e.g.
//
// ```sh
//  $ pulumi import aws:elb/appCookieStickinessPolicy:AppCookieStickinessPolicy example my-elb:80:my-policy
// ```
type AppCookieStickinessPolicy struct {
	pulumi.CustomResourceState

	// The application cookie whose lifetime the ELB's cookie should follow.
	CookieName pulumi.StringOutput `pulumi:"cookieName"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntOutput `pulumi:"lbPort"`
	// The name of load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringOutput `pulumi:"loadBalancer"`
	// The name of the stickiness policy.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewAppCookieStickinessPolicy registers a new resource with the given unique name, arguments, and options.
func NewAppCookieStickinessPolicy(ctx *pulumi.Context,
	name string, args *AppCookieStickinessPolicyArgs, opts ...pulumi.ResourceOption) (*AppCookieStickinessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CookieName == nil {
		return nil, errors.New("invalid value for required argument 'CookieName'")
	}
	if args.LbPort == nil {
		return nil, errors.New("invalid value for required argument 'LbPort'")
	}
	if args.LoadBalancer == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancer'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("aws:elasticloadbalancing/appCookieStickinessPolicy:AppCookieStickinessPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource AppCookieStickinessPolicy
	err := ctx.RegisterResource("aws:elb/appCookieStickinessPolicy:AppCookieStickinessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppCookieStickinessPolicy gets an existing AppCookieStickinessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppCookieStickinessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppCookieStickinessPolicyState, opts ...pulumi.ResourceOption) (*AppCookieStickinessPolicy, error) {
	var resource AppCookieStickinessPolicy
	err := ctx.ReadResource("aws:elb/appCookieStickinessPolicy:AppCookieStickinessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppCookieStickinessPolicy resources.
type appCookieStickinessPolicyState struct {
	// The application cookie whose lifetime the ELB's cookie should follow.
	CookieName *string `pulumi:"cookieName"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort *int `pulumi:"lbPort"`
	// The name of load balancer to which the policy
	// should be attached.
	LoadBalancer *string `pulumi:"loadBalancer"`
	// The name of the stickiness policy.
	Name *string `pulumi:"name"`
}

type AppCookieStickinessPolicyState struct {
	// The application cookie whose lifetime the ELB's cookie should follow.
	CookieName pulumi.StringPtrInput
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntPtrInput
	// The name of load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringPtrInput
	// The name of the stickiness policy.
	Name pulumi.StringPtrInput
}

func (AppCookieStickinessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*appCookieStickinessPolicyState)(nil)).Elem()
}

type appCookieStickinessPolicyArgs struct {
	// The application cookie whose lifetime the ELB's cookie should follow.
	CookieName string `pulumi:"cookieName"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort int `pulumi:"lbPort"`
	// The name of load balancer to which the policy
	// should be attached.
	LoadBalancer string `pulumi:"loadBalancer"`
	// The name of the stickiness policy.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a AppCookieStickinessPolicy resource.
type AppCookieStickinessPolicyArgs struct {
	// The application cookie whose lifetime the ELB's cookie should follow.
	CookieName pulumi.StringInput
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntInput
	// The name of load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringInput
	// The name of the stickiness policy.
	Name pulumi.StringPtrInput
}

func (AppCookieStickinessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appCookieStickinessPolicyArgs)(nil)).Elem()
}

type AppCookieStickinessPolicyInput interface {
	pulumi.Input

	ToAppCookieStickinessPolicyOutput() AppCookieStickinessPolicyOutput
	ToAppCookieStickinessPolicyOutputWithContext(ctx context.Context) AppCookieStickinessPolicyOutput
}

func (*AppCookieStickinessPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*AppCookieStickinessPolicy)(nil))
}

func (i *AppCookieStickinessPolicy) ToAppCookieStickinessPolicyOutput() AppCookieStickinessPolicyOutput {
	return i.ToAppCookieStickinessPolicyOutputWithContext(context.Background())
}

func (i *AppCookieStickinessPolicy) ToAppCookieStickinessPolicyOutputWithContext(ctx context.Context) AppCookieStickinessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppCookieStickinessPolicyOutput)
}

func (i *AppCookieStickinessPolicy) ToAppCookieStickinessPolicyPtrOutput() AppCookieStickinessPolicyPtrOutput {
	return i.ToAppCookieStickinessPolicyPtrOutputWithContext(context.Background())
}

func (i *AppCookieStickinessPolicy) ToAppCookieStickinessPolicyPtrOutputWithContext(ctx context.Context) AppCookieStickinessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppCookieStickinessPolicyPtrOutput)
}

type AppCookieStickinessPolicyPtrInput interface {
	pulumi.Input

	ToAppCookieStickinessPolicyPtrOutput() AppCookieStickinessPolicyPtrOutput
	ToAppCookieStickinessPolicyPtrOutputWithContext(ctx context.Context) AppCookieStickinessPolicyPtrOutput
}

type appCookieStickinessPolicyPtrType AppCookieStickinessPolicyArgs

func (*appCookieStickinessPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppCookieStickinessPolicy)(nil))
}

func (i *appCookieStickinessPolicyPtrType) ToAppCookieStickinessPolicyPtrOutput() AppCookieStickinessPolicyPtrOutput {
	return i.ToAppCookieStickinessPolicyPtrOutputWithContext(context.Background())
}

func (i *appCookieStickinessPolicyPtrType) ToAppCookieStickinessPolicyPtrOutputWithContext(ctx context.Context) AppCookieStickinessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppCookieStickinessPolicyPtrOutput)
}

// AppCookieStickinessPolicyArrayInput is an input type that accepts AppCookieStickinessPolicyArray and AppCookieStickinessPolicyArrayOutput values.
// You can construct a concrete instance of `AppCookieStickinessPolicyArrayInput` via:
//
//          AppCookieStickinessPolicyArray{ AppCookieStickinessPolicyArgs{...} }
type AppCookieStickinessPolicyArrayInput interface {
	pulumi.Input

	ToAppCookieStickinessPolicyArrayOutput() AppCookieStickinessPolicyArrayOutput
	ToAppCookieStickinessPolicyArrayOutputWithContext(context.Context) AppCookieStickinessPolicyArrayOutput
}

type AppCookieStickinessPolicyArray []AppCookieStickinessPolicyInput

func (AppCookieStickinessPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AppCookieStickinessPolicy)(nil))
}

func (i AppCookieStickinessPolicyArray) ToAppCookieStickinessPolicyArrayOutput() AppCookieStickinessPolicyArrayOutput {
	return i.ToAppCookieStickinessPolicyArrayOutputWithContext(context.Background())
}

func (i AppCookieStickinessPolicyArray) ToAppCookieStickinessPolicyArrayOutputWithContext(ctx context.Context) AppCookieStickinessPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppCookieStickinessPolicyArrayOutput)
}

// AppCookieStickinessPolicyMapInput is an input type that accepts AppCookieStickinessPolicyMap and AppCookieStickinessPolicyMapOutput values.
// You can construct a concrete instance of `AppCookieStickinessPolicyMapInput` via:
//
//          AppCookieStickinessPolicyMap{ "key": AppCookieStickinessPolicyArgs{...} }
type AppCookieStickinessPolicyMapInput interface {
	pulumi.Input

	ToAppCookieStickinessPolicyMapOutput() AppCookieStickinessPolicyMapOutput
	ToAppCookieStickinessPolicyMapOutputWithContext(context.Context) AppCookieStickinessPolicyMapOutput
}

type AppCookieStickinessPolicyMap map[string]AppCookieStickinessPolicyInput

func (AppCookieStickinessPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AppCookieStickinessPolicy)(nil))
}

func (i AppCookieStickinessPolicyMap) ToAppCookieStickinessPolicyMapOutput() AppCookieStickinessPolicyMapOutput {
	return i.ToAppCookieStickinessPolicyMapOutputWithContext(context.Background())
}

func (i AppCookieStickinessPolicyMap) ToAppCookieStickinessPolicyMapOutputWithContext(ctx context.Context) AppCookieStickinessPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppCookieStickinessPolicyMapOutput)
}

type AppCookieStickinessPolicyOutput struct {
	*pulumi.OutputState
}

func (AppCookieStickinessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppCookieStickinessPolicy)(nil))
}

func (o AppCookieStickinessPolicyOutput) ToAppCookieStickinessPolicyOutput() AppCookieStickinessPolicyOutput {
	return o
}

func (o AppCookieStickinessPolicyOutput) ToAppCookieStickinessPolicyOutputWithContext(ctx context.Context) AppCookieStickinessPolicyOutput {
	return o
}

func (o AppCookieStickinessPolicyOutput) ToAppCookieStickinessPolicyPtrOutput() AppCookieStickinessPolicyPtrOutput {
	return o.ToAppCookieStickinessPolicyPtrOutputWithContext(context.Background())
}

func (o AppCookieStickinessPolicyOutput) ToAppCookieStickinessPolicyPtrOutputWithContext(ctx context.Context) AppCookieStickinessPolicyPtrOutput {
	return o.ApplyT(func(v AppCookieStickinessPolicy) *AppCookieStickinessPolicy {
		return &v
	}).(AppCookieStickinessPolicyPtrOutput)
}

type AppCookieStickinessPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (AppCookieStickinessPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppCookieStickinessPolicy)(nil))
}

func (o AppCookieStickinessPolicyPtrOutput) ToAppCookieStickinessPolicyPtrOutput() AppCookieStickinessPolicyPtrOutput {
	return o
}

func (o AppCookieStickinessPolicyPtrOutput) ToAppCookieStickinessPolicyPtrOutputWithContext(ctx context.Context) AppCookieStickinessPolicyPtrOutput {
	return o
}

type AppCookieStickinessPolicyArrayOutput struct{ *pulumi.OutputState }

func (AppCookieStickinessPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppCookieStickinessPolicy)(nil))
}

func (o AppCookieStickinessPolicyArrayOutput) ToAppCookieStickinessPolicyArrayOutput() AppCookieStickinessPolicyArrayOutput {
	return o
}

func (o AppCookieStickinessPolicyArrayOutput) ToAppCookieStickinessPolicyArrayOutputWithContext(ctx context.Context) AppCookieStickinessPolicyArrayOutput {
	return o
}

func (o AppCookieStickinessPolicyArrayOutput) Index(i pulumi.IntInput) AppCookieStickinessPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppCookieStickinessPolicy {
		return vs[0].([]AppCookieStickinessPolicy)[vs[1].(int)]
	}).(AppCookieStickinessPolicyOutput)
}

type AppCookieStickinessPolicyMapOutput struct{ *pulumi.OutputState }

func (AppCookieStickinessPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppCookieStickinessPolicy)(nil))
}

func (o AppCookieStickinessPolicyMapOutput) ToAppCookieStickinessPolicyMapOutput() AppCookieStickinessPolicyMapOutput {
	return o
}

func (o AppCookieStickinessPolicyMapOutput) ToAppCookieStickinessPolicyMapOutputWithContext(ctx context.Context) AppCookieStickinessPolicyMapOutput {
	return o
}

func (o AppCookieStickinessPolicyMapOutput) MapIndex(k pulumi.StringInput) AppCookieStickinessPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppCookieStickinessPolicy {
		return vs[0].(map[string]AppCookieStickinessPolicy)[vs[1].(string)]
	}).(AppCookieStickinessPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(AppCookieStickinessPolicyOutput{})
	pulumi.RegisterOutputType(AppCookieStickinessPolicyPtrOutput{})
	pulumi.RegisterOutputType(AppCookieStickinessPolicyArrayOutput{})
	pulumi.RegisterOutputType(AppCookieStickinessPolicyMapOutput{})
}
