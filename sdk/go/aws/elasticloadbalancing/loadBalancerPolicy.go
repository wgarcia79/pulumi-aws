// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancing

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a load balancer policy, which can be attached to an ELB listener or backend server.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/elb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := elb.NewLoadBalancer(ctx, "wu_tang", &elb.LoadBalancerArgs{
// 			AvailabilityZones: pulumi.StringArray{
// 				pulumi.String("us-east-1a"),
// 			},
// 			Listeners: elb.LoadBalancerListenerArray{
// 				&elb.LoadBalancerListenerArgs{
// 					InstancePort:     pulumi.Int(443),
// 					InstanceProtocol: pulumi.String("http"),
// 					LbPort:           pulumi.Int(443),
// 					LbProtocol:       pulumi.String("https"),
// 					SslCertificateId: pulumi.String("arn:aws:iam::000000000000:server-certificate/wu-tang.net"),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("wu-tang"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elb.NewLoadBalancerPolicy(ctx, "wu_tang_ca_pubkey_policy", &elb.LoadBalancerPolicyArgs{
// 			LoadBalancerName: wu_tang.Name,
// 			PolicyName:       pulumi.String("wu-tang-ca-pubkey-policy"),
// 			PolicyTypeName:   pulumi.String("PublicKeyPolicyType"),
// 			PolicyAttributes: elb.LoadBalancerPolicyPolicyAttributeArray{
// 				&elb.LoadBalancerPolicyPolicyAttributeArgs{
// 					Name:  pulumi.String("PublicKey"),
// 					Value: readFileOrPanic("wu-tang-pubkey"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elb.NewLoadBalancerPolicy(ctx, "wu_tang_root_ca_backend_auth_policy", &elb.LoadBalancerPolicyArgs{
// 			LoadBalancerName: wu_tang.Name,
// 			PolicyName:       pulumi.String("wu-tang-root-ca-backend-auth-policy"),
// 			PolicyTypeName:   pulumi.String("BackendServerAuthenticationPolicyType"),
// 			PolicyAttributes: elb.LoadBalancerPolicyPolicyAttributeArray{
// 				&elb.LoadBalancerPolicyPolicyAttributeArgs{
// 					Name:  pulumi.String("PublicKeyPolicyName"),
// 					Value: pulumi.Any(aws_load_balancer_policy.Wu - tang - root - ca - pubkey - policy.Policy_name),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elb.NewLoadBalancerPolicy(ctx, "wu_tang_ssl", &elb.LoadBalancerPolicyArgs{
// 			LoadBalancerName: wu_tang.Name,
// 			PolicyName:       pulumi.String("wu-tang-ssl"),
// 			PolicyTypeName:   pulumi.String("SSLNegotiationPolicyType"),
// 			PolicyAttributes: elb.LoadBalancerPolicyPolicyAttributeArray{
// 				&elb.LoadBalancerPolicyPolicyAttributeArgs{
// 					Name:  pulumi.String("ECDHE-ECDSA-AES128-GCM-SHA256"),
// 					Value: pulumi.String("true"),
// 				},
// 				&elb.LoadBalancerPolicyPolicyAttributeArgs{
// 					Name:  pulumi.String("Protocol-TLSv1.2"),
// 					Value: pulumi.String("true"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elb.NewLoadBalancerPolicy(ctx, "wu_tang_ssl_tls_1_1", &elb.LoadBalancerPolicyArgs{
// 			LoadBalancerName: wu_tang.Name,
// 			PolicyName:       pulumi.String("wu-tang-ssl"),
// 			PolicyTypeName:   pulumi.String("SSLNegotiationPolicyType"),
// 			PolicyAttributes: elb.LoadBalancerPolicyPolicyAttributeArray{
// 				&elb.LoadBalancerPolicyPolicyAttributeArgs{
// 					Name:  pulumi.String("Reference-Security-Policy"),
// 					Value: pulumi.String("ELBSecurityPolicy-TLS-1-1-2017-01"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elb.NewLoadBalancerBackendServerPolicy(ctx, "wu_tang_backend_auth_policies_443", &elb.LoadBalancerBackendServerPolicyArgs{
// 			LoadBalancerName: wu_tang.Name,
// 			InstancePort:     pulumi.Int(443),
// 			PolicyNames: pulumi.StringArray{
// 				wu_tang_root_ca_backend_auth_policy.PolicyName,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elb.NewListenerPolicy(ctx, "wu_tang_listener_policies_443", &elb.ListenerPolicyArgs{
// 			LoadBalancerName: wu_tang.Name,
// 			LoadBalancerPort: pulumi.Int(443),
// 			PolicyNames: pulumi.StringArray{
// 				wu_tang_ssl.PolicyName,
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
// Deprecated: aws.elasticloadbalancing.LoadBalancerPolicy has been deprecated in favor of aws.elb.LoadBalancerPolicy
type LoadBalancerPolicy struct {
	pulumi.CustomResourceState

	// The load balancer on which the policy is defined.
	LoadBalancerName pulumi.StringOutput `pulumi:"loadBalancerName"`
	// Policy attribute to apply to the policy.
	PolicyAttributes LoadBalancerPolicyPolicyAttributeArrayOutput `pulumi:"policyAttributes"`
	// The name of the load balancer policy.
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
	// The policy type.
	PolicyTypeName pulumi.StringOutput `pulumi:"policyTypeName"`
}

// NewLoadBalancerPolicy registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancerPolicy(ctx *pulumi.Context,
	name string, args *LoadBalancerPolicyArgs, opts ...pulumi.ResourceOption) (*LoadBalancerPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadBalancerName == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerName'")
	}
	if args.PolicyName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyName'")
	}
	if args.PolicyTypeName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyTypeName'")
	}
	var resource LoadBalancerPolicy
	err := ctx.RegisterResource("aws:elasticloadbalancing/loadBalancerPolicy:LoadBalancerPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancerPolicy gets an existing LoadBalancerPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancerPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerPolicyState, opts ...pulumi.ResourceOption) (*LoadBalancerPolicy, error) {
	var resource LoadBalancerPolicy
	err := ctx.ReadResource("aws:elasticloadbalancing/loadBalancerPolicy:LoadBalancerPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancerPolicy resources.
type loadBalancerPolicyState struct {
	// The load balancer on which the policy is defined.
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// Policy attribute to apply to the policy.
	PolicyAttributes []LoadBalancerPolicyPolicyAttribute `pulumi:"policyAttributes"`
	// The name of the load balancer policy.
	PolicyName *string `pulumi:"policyName"`
	// The policy type.
	PolicyTypeName *string `pulumi:"policyTypeName"`
}

type LoadBalancerPolicyState struct {
	// The load balancer on which the policy is defined.
	LoadBalancerName pulumi.StringPtrInput
	// Policy attribute to apply to the policy.
	PolicyAttributes LoadBalancerPolicyPolicyAttributeArrayInput
	// The name of the load balancer policy.
	PolicyName pulumi.StringPtrInput
	// The policy type.
	PolicyTypeName pulumi.StringPtrInput
}

func (LoadBalancerPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerPolicyState)(nil)).Elem()
}

type loadBalancerPolicyArgs struct {
	// The load balancer on which the policy is defined.
	LoadBalancerName string `pulumi:"loadBalancerName"`
	// Policy attribute to apply to the policy.
	PolicyAttributes []LoadBalancerPolicyPolicyAttribute `pulumi:"policyAttributes"`
	// The name of the load balancer policy.
	PolicyName string `pulumi:"policyName"`
	// The policy type.
	PolicyTypeName string `pulumi:"policyTypeName"`
}

// The set of arguments for constructing a LoadBalancerPolicy resource.
type LoadBalancerPolicyArgs struct {
	// The load balancer on which the policy is defined.
	LoadBalancerName pulumi.StringInput
	// Policy attribute to apply to the policy.
	PolicyAttributes LoadBalancerPolicyPolicyAttributeArrayInput
	// The name of the load balancer policy.
	PolicyName pulumi.StringInput
	// The policy type.
	PolicyTypeName pulumi.StringInput
}

func (LoadBalancerPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerPolicyArgs)(nil)).Elem()
}

type LoadBalancerPolicyInput interface {
	pulumi.Input

	ToLoadBalancerPolicyOutput() LoadBalancerPolicyOutput
	ToLoadBalancerPolicyOutputWithContext(ctx context.Context) LoadBalancerPolicyOutput
}

func (*LoadBalancerPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerPolicy)(nil))
}

func (i *LoadBalancerPolicy) ToLoadBalancerPolicyOutput() LoadBalancerPolicyOutput {
	return i.ToLoadBalancerPolicyOutputWithContext(context.Background())
}

func (i *LoadBalancerPolicy) ToLoadBalancerPolicyOutputWithContext(ctx context.Context) LoadBalancerPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerPolicyOutput)
}

func (i *LoadBalancerPolicy) ToLoadBalancerPolicyPtrOutput() LoadBalancerPolicyPtrOutput {
	return i.ToLoadBalancerPolicyPtrOutputWithContext(context.Background())
}

func (i *LoadBalancerPolicy) ToLoadBalancerPolicyPtrOutputWithContext(ctx context.Context) LoadBalancerPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerPolicyPtrOutput)
}

type LoadBalancerPolicyPtrInput interface {
	pulumi.Input

	ToLoadBalancerPolicyPtrOutput() LoadBalancerPolicyPtrOutput
	ToLoadBalancerPolicyPtrOutputWithContext(ctx context.Context) LoadBalancerPolicyPtrOutput
}

type loadBalancerPolicyPtrType LoadBalancerPolicyArgs

func (*loadBalancerPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerPolicy)(nil))
}

func (i *loadBalancerPolicyPtrType) ToLoadBalancerPolicyPtrOutput() LoadBalancerPolicyPtrOutput {
	return i.ToLoadBalancerPolicyPtrOutputWithContext(context.Background())
}

func (i *loadBalancerPolicyPtrType) ToLoadBalancerPolicyPtrOutputWithContext(ctx context.Context) LoadBalancerPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerPolicyPtrOutput)
}

// LoadBalancerPolicyArrayInput is an input type that accepts LoadBalancerPolicyArray and LoadBalancerPolicyArrayOutput values.
// You can construct a concrete instance of `LoadBalancerPolicyArrayInput` via:
//
//          LoadBalancerPolicyArray{ LoadBalancerPolicyArgs{...} }
type LoadBalancerPolicyArrayInput interface {
	pulumi.Input

	ToLoadBalancerPolicyArrayOutput() LoadBalancerPolicyArrayOutput
	ToLoadBalancerPolicyArrayOutputWithContext(context.Context) LoadBalancerPolicyArrayOutput
}

type LoadBalancerPolicyArray []LoadBalancerPolicyInput

func (LoadBalancerPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadBalancerPolicy)(nil)).Elem()
}

func (i LoadBalancerPolicyArray) ToLoadBalancerPolicyArrayOutput() LoadBalancerPolicyArrayOutput {
	return i.ToLoadBalancerPolicyArrayOutputWithContext(context.Background())
}

func (i LoadBalancerPolicyArray) ToLoadBalancerPolicyArrayOutputWithContext(ctx context.Context) LoadBalancerPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerPolicyArrayOutput)
}

// LoadBalancerPolicyMapInput is an input type that accepts LoadBalancerPolicyMap and LoadBalancerPolicyMapOutput values.
// You can construct a concrete instance of `LoadBalancerPolicyMapInput` via:
//
//          LoadBalancerPolicyMap{ "key": LoadBalancerPolicyArgs{...} }
type LoadBalancerPolicyMapInput interface {
	pulumi.Input

	ToLoadBalancerPolicyMapOutput() LoadBalancerPolicyMapOutput
	ToLoadBalancerPolicyMapOutputWithContext(context.Context) LoadBalancerPolicyMapOutput
}

type LoadBalancerPolicyMap map[string]LoadBalancerPolicyInput

func (LoadBalancerPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadBalancerPolicy)(nil)).Elem()
}

func (i LoadBalancerPolicyMap) ToLoadBalancerPolicyMapOutput() LoadBalancerPolicyMapOutput {
	return i.ToLoadBalancerPolicyMapOutputWithContext(context.Background())
}

func (i LoadBalancerPolicyMap) ToLoadBalancerPolicyMapOutputWithContext(ctx context.Context) LoadBalancerPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerPolicyMapOutput)
}

type LoadBalancerPolicyOutput struct{ *pulumi.OutputState }

func (LoadBalancerPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerPolicy)(nil))
}

func (o LoadBalancerPolicyOutput) ToLoadBalancerPolicyOutput() LoadBalancerPolicyOutput {
	return o
}

func (o LoadBalancerPolicyOutput) ToLoadBalancerPolicyOutputWithContext(ctx context.Context) LoadBalancerPolicyOutput {
	return o
}

func (o LoadBalancerPolicyOutput) ToLoadBalancerPolicyPtrOutput() LoadBalancerPolicyPtrOutput {
	return o.ToLoadBalancerPolicyPtrOutputWithContext(context.Background())
}

func (o LoadBalancerPolicyOutput) ToLoadBalancerPolicyPtrOutputWithContext(ctx context.Context) LoadBalancerPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoadBalancerPolicy) *LoadBalancerPolicy {
		return &v
	}).(LoadBalancerPolicyPtrOutput)
}

type LoadBalancerPolicyPtrOutput struct{ *pulumi.OutputState }

func (LoadBalancerPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerPolicy)(nil))
}

func (o LoadBalancerPolicyPtrOutput) ToLoadBalancerPolicyPtrOutput() LoadBalancerPolicyPtrOutput {
	return o
}

func (o LoadBalancerPolicyPtrOutput) ToLoadBalancerPolicyPtrOutputWithContext(ctx context.Context) LoadBalancerPolicyPtrOutput {
	return o
}

func (o LoadBalancerPolicyPtrOutput) Elem() LoadBalancerPolicyOutput {
	return o.ApplyT(func(v *LoadBalancerPolicy) LoadBalancerPolicy {
		if v != nil {
			return *v
		}
		var ret LoadBalancerPolicy
		return ret
	}).(LoadBalancerPolicyOutput)
}

type LoadBalancerPolicyArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerPolicy)(nil))
}

func (o LoadBalancerPolicyArrayOutput) ToLoadBalancerPolicyArrayOutput() LoadBalancerPolicyArrayOutput {
	return o
}

func (o LoadBalancerPolicyArrayOutput) ToLoadBalancerPolicyArrayOutputWithContext(ctx context.Context) LoadBalancerPolicyArrayOutput {
	return o
}

func (o LoadBalancerPolicyArrayOutput) Index(i pulumi.IntInput) LoadBalancerPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancerPolicy {
		return vs[0].([]LoadBalancerPolicy)[vs[1].(int)]
	}).(LoadBalancerPolicyOutput)
}

type LoadBalancerPolicyMapOutput struct{ *pulumi.OutputState }

func (LoadBalancerPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LoadBalancerPolicy)(nil))
}

func (o LoadBalancerPolicyMapOutput) ToLoadBalancerPolicyMapOutput() LoadBalancerPolicyMapOutput {
	return o
}

func (o LoadBalancerPolicyMapOutput) ToLoadBalancerPolicyMapOutputWithContext(ctx context.Context) LoadBalancerPolicyMapOutput {
	return o
}

func (o LoadBalancerPolicyMapOutput) MapIndex(k pulumi.StringInput) LoadBalancerPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LoadBalancerPolicy {
		return vs[0].(map[string]LoadBalancerPolicy)[vs[1].(string)]
	}).(LoadBalancerPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerPolicyInput)(nil)).Elem(), &LoadBalancerPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerPolicyPtrInput)(nil)).Elem(), &LoadBalancerPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerPolicyArrayInput)(nil)).Elem(), LoadBalancerPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerPolicyMapInput)(nil)).Elem(), LoadBalancerPolicyMap{})
	pulumi.RegisterOutputType(LoadBalancerPolicyOutput{})
	pulumi.RegisterOutputType(LoadBalancerPolicyPtrOutput{})
	pulumi.RegisterOutputType(LoadBalancerPolicyArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerPolicyMapOutput{})
}
