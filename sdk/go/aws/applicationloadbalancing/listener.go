// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package applicationloadbalancing

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Load Balancer Listener resource.
//
// > **Note:** `alb.Listener` is known as `lb.Listener`. The functionality is identical.
//
// ## Example Usage
// ### Forward Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		frontEndLoadBalancer, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		frontEndTargetGroup, err := lb.NewTargetGroup(ctx, "frontEndTargetGroup", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
// 			LoadBalancerArn: frontEndLoadBalancer.Arn,
// 			Port:            pulumi.Int(443),
// 			Protocol:        pulumi.String("HTTPS"),
// 			SslPolicy:       pulumi.String("ELBSecurityPolicy-2016-08"),
// 			CertificateArn:  pulumi.String("arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4"),
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					Type:           pulumi.String("forward"),
// 					TargetGroupArn: frontEndTargetGroup.Arn,
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
// To a NLB:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := lb.NewListener(ctx, "frontEnd", &lb.ListenerArgs{
// 			LoadBalancerArn: pulumi.Any(aws_lb.Front_end.Arn),
// 			Port:            pulumi.Int(443),
// 			Protocol:        pulumi.String("TLS"),
// 			CertificateArn:  pulumi.String("arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4"),
// 			AlpnPolicy:      pulumi.String("HTTP2Preferred"),
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					Type:           pulumi.String("forward"),
// 					TargetGroupArn: pulumi.Any(aws_lb_target_group.Front_end.Arn),
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
// ### Redirect Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		frontEndLoadBalancer, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
// 			LoadBalancerArn: frontEndLoadBalancer.Arn,
// 			Port:            pulumi.Int(80),
// 			Protocol:        pulumi.String("HTTP"),
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					Type: pulumi.String("redirect"),
// 					Redirect: &lb.ListenerDefaultActionRedirectArgs{
// 						Port:       pulumi.String("443"),
// 						Protocol:   pulumi.String("HTTPS"),
// 						StatusCode: pulumi.String("HTTP_301"),
// 					},
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
// ### Fixed-response Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		frontEndLoadBalancer, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
// 			LoadBalancerArn: frontEndLoadBalancer.Arn,
// 			Port:            pulumi.Int(80),
// 			Protocol:        pulumi.String("HTTP"),
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					Type: pulumi.String("fixed-response"),
// 					FixedResponse: &lb.ListenerDefaultActionFixedResponseArgs{
// 						ContentType: pulumi.String("text/plain"),
// 						MessageBody: pulumi.String("Fixed response content"),
// 						StatusCode:  pulumi.String("200"),
// 					},
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
// ### Authenticate-cognito Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cognito"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		frontEndLoadBalancer, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		frontEndTargetGroup, err := lb.NewTargetGroup(ctx, "frontEndTargetGroup", nil)
// 		if err != nil {
// 			return err
// 		}
// 		pool, err := cognito.NewUserPool(ctx, "pool", nil)
// 		if err != nil {
// 			return err
// 		}
// 		client, err := cognito.NewUserPoolClient(ctx, "client", nil)
// 		if err != nil {
// 			return err
// 		}
// 		domain, err := cognito.NewUserPoolDomain(ctx, "domain", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
// 			LoadBalancerArn: frontEndLoadBalancer.Arn,
// 			Port:            pulumi.Int(80),
// 			Protocol:        pulumi.String("HTTP"),
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					Type: pulumi.String("authenticate-cognito"),
// 					AuthenticateCognito: &lb.ListenerDefaultActionAuthenticateCognitoArgs{
// 						UserPoolArn:      pool.Arn,
// 						UserPoolClientId: client.ID(),
// 						UserPoolDomain:   domain.Domain,
// 					},
// 				},
// 				&lb.ListenerDefaultActionArgs{
// 					Type:           pulumi.String("forward"),
// 					TargetGroupArn: frontEndTargetGroup.Arn,
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
// ### Authenticate-OIDC Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		frontEndLoadBalancer, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		frontEndTargetGroup, err := lb.NewTargetGroup(ctx, "frontEndTargetGroup", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListener(ctx, "frontEndListener", &lb.ListenerArgs{
// 			LoadBalancerArn: frontEndLoadBalancer.Arn,
// 			Port:            pulumi.Int(80),
// 			Protocol:        pulumi.String("HTTP"),
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					Type: pulumi.String("authenticate-oidc"),
// 					AuthenticateOidc: &lb.ListenerDefaultActionAuthenticateOidcArgs{
// 						AuthorizationEndpoint: pulumi.String("https://example.com/authorization_endpoint"),
// 						ClientId:              pulumi.String("client_id"),
// 						ClientSecret:          pulumi.String("client_secret"),
// 						Issuer:                pulumi.String("https://example.com"),
// 						TokenEndpoint:         pulumi.String("https://example.com/token_endpoint"),
// 						UserInfoEndpoint:      pulumi.String("https://example.com/user_info_endpoint"),
// 					},
// 				},
// 				&lb.ListenerDefaultActionArgs{
// 					Type:           pulumi.String("forward"),
// 					TargetGroupArn: frontEndTargetGroup.Arn,
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
// ### Gateway Load Balancer Listener
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleLoadBalancer, err := lb.NewLoadBalancer(ctx, "exampleLoadBalancer", &lb.LoadBalancerArgs{
// 			LoadBalancerType: pulumi.String("gateway"),
// 			SubnetMappings: lb.LoadBalancerSubnetMappingArray{
// 				&lb.LoadBalancerSubnetMappingArgs{
// 					SubnetId: pulumi.Any(aws_subnet.Example.Id),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleTargetGroup, err := lb.NewTargetGroup(ctx, "exampleTargetGroup", &lb.TargetGroupArgs{
// 			Port:     pulumi.Int(6081),
// 			Protocol: pulumi.String("GENEVE"),
// 			VpcId:    pulumi.Any(aws_vpc.Example.Id),
// 			HealthCheck: &lb.TargetGroupHealthCheckArgs{
// 				Port:     pulumi.String("80"),
// 				Protocol: pulumi.String("HTTP"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListener(ctx, "exampleListener", &lb.ListenerArgs{
// 			LoadBalancerArn: exampleLoadBalancer.ID(),
// 			DefaultActions: lb.ListenerDefaultActionArray{
// 				&lb.ListenerDefaultActionArgs{
// 					TargetGroupArn: exampleTargetGroup.ID(),
// 					Type:           pulumi.String("forward"),
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
// Listeners can be imported using their ARN, e.g.
//
// ```sh
//  $ pulumi import aws:applicationloadbalancing/listener:Listener front_end arn:aws:elasticloadbalancing:us-west-2:187416307283:listener/app/front-end-alb/8e4497da625e2d8a/9ab28ade35828f96
// ```
//
// Deprecated: aws.applicationloadbalancing.Listener has been deprecated in favor of aws.alb.Listener
type Listener struct {
	pulumi.CustomResourceState

	// Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if `protocol` is `TLS`. Valid values are `HTTP1Only`, `HTTP2Only`, `HTTP2Optional`, `HTTP2Preferred`, and `None`.
	AlpnPolicy pulumi.StringPtrOutput `pulumi:"alpnPolicy"`
	// ARN of the target group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
	CertificateArn pulumi.StringPtrOutput `pulumi:"certificateArn"`
	// Configuration block for default actions. Detailed below.
	DefaultActions ListenerDefaultActionArrayOutput `pulumi:"defaultActions"`
	// ARN of the load balancer.
	LoadBalancerArn pulumi.StringOutput `pulumi:"loadBalancerArn"`
	// Port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// Protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy pulumi.StringOutput `pulumi:"sslPolicy"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewListener registers a new resource with the given unique name, arguments, and options.
func NewListener(ctx *pulumi.Context,
	name string, args *ListenerArgs, opts ...pulumi.ResourceOption) (*Listener, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultActions == nil {
		return nil, errors.New("invalid value for required argument 'DefaultActions'")
	}
	if args.LoadBalancerArn == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerArn'")
	}
	var resource Listener
	err := ctx.RegisterResource("aws:applicationloadbalancing/listener:Listener", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListener gets an existing Listener resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListener(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerState, opts ...pulumi.ResourceOption) (*Listener, error) {
	var resource Listener
	err := ctx.ReadResource("aws:applicationloadbalancing/listener:Listener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Listener resources.
type listenerState struct {
	// Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if `protocol` is `TLS`. Valid values are `HTTP1Only`, `HTTP2Only`, `HTTP2Optional`, `HTTP2Preferred`, and `None`.
	AlpnPolicy *string `pulumi:"alpnPolicy"`
	// ARN of the target group.
	Arn *string `pulumi:"arn"`
	// ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
	CertificateArn *string `pulumi:"certificateArn"`
	// Configuration block for default actions. Detailed below.
	DefaultActions []ListenerDefaultAction `pulumi:"defaultActions"`
	// ARN of the load balancer.
	LoadBalancerArn *string `pulumi:"loadBalancerArn"`
	// Port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
	Port *int `pulumi:"port"`
	// Protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
	Protocol *string `pulumi:"protocol"`
	// Name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy *string `pulumi:"sslPolicy"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ListenerState struct {
	// Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if `protocol` is `TLS`. Valid values are `HTTP1Only`, `HTTP2Only`, `HTTP2Optional`, `HTTP2Preferred`, and `None`.
	AlpnPolicy pulumi.StringPtrInput
	// ARN of the target group.
	Arn pulumi.StringPtrInput
	// ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
	CertificateArn pulumi.StringPtrInput
	// Configuration block for default actions. Detailed below.
	DefaultActions ListenerDefaultActionArrayInput
	// ARN of the load balancer.
	LoadBalancerArn pulumi.StringPtrInput
	// Port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
	Port pulumi.IntPtrInput
	// Protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
	Protocol pulumi.StringPtrInput
	// Name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (ListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerState)(nil)).Elem()
}

type listenerArgs struct {
	// Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if `protocol` is `TLS`. Valid values are `HTTP1Only`, `HTTP2Only`, `HTTP2Optional`, `HTTP2Preferred`, and `None`.
	AlpnPolicy *string `pulumi:"alpnPolicy"`
	// ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
	CertificateArn *string `pulumi:"certificateArn"`
	// Configuration block for default actions. Detailed below.
	DefaultActions []ListenerDefaultAction `pulumi:"defaultActions"`
	// ARN of the load balancer.
	LoadBalancerArn string `pulumi:"loadBalancerArn"`
	// Port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
	Port *int `pulumi:"port"`
	// Protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
	Protocol *string `pulumi:"protocol"`
	// Name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy *string `pulumi:"sslPolicy"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a Listener resource.
type ListenerArgs struct {
	// Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if `protocol` is `TLS`. Valid values are `HTTP1Only`, `HTTP2Only`, `HTTP2Optional`, `HTTP2Preferred`, and `None`.
	AlpnPolicy pulumi.StringPtrInput
	// ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
	CertificateArn pulumi.StringPtrInput
	// Configuration block for default actions. Detailed below.
	DefaultActions ListenerDefaultActionArrayInput
	// ARN of the load balancer.
	LoadBalancerArn pulumi.StringInput
	// Port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
	Port pulumi.IntPtrInput
	// Protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
	Protocol pulumi.StringPtrInput
	// Name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (ListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerArgs)(nil)).Elem()
}

type ListenerInput interface {
	pulumi.Input

	ToListenerOutput() ListenerOutput
	ToListenerOutputWithContext(ctx context.Context) ListenerOutput
}

func (*Listener) ElementType() reflect.Type {
	return reflect.TypeOf((*Listener)(nil))
}

func (i *Listener) ToListenerOutput() ListenerOutput {
	return i.ToListenerOutputWithContext(context.Background())
}

func (i *Listener) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerOutput)
}

func (i *Listener) ToListenerPtrOutput() ListenerPtrOutput {
	return i.ToListenerPtrOutputWithContext(context.Background())
}

func (i *Listener) ToListenerPtrOutputWithContext(ctx context.Context) ListenerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerPtrOutput)
}

type ListenerPtrInput interface {
	pulumi.Input

	ToListenerPtrOutput() ListenerPtrOutput
	ToListenerPtrOutputWithContext(ctx context.Context) ListenerPtrOutput
}

type listenerPtrType ListenerArgs

func (*listenerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil))
}

func (i *listenerPtrType) ToListenerPtrOutput() ListenerPtrOutput {
	return i.ToListenerPtrOutputWithContext(context.Background())
}

func (i *listenerPtrType) ToListenerPtrOutputWithContext(ctx context.Context) ListenerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerPtrOutput)
}

// ListenerArrayInput is an input type that accepts ListenerArray and ListenerArrayOutput values.
// You can construct a concrete instance of `ListenerArrayInput` via:
//
//          ListenerArray{ ListenerArgs{...} }
type ListenerArrayInput interface {
	pulumi.Input

	ToListenerArrayOutput() ListenerArrayOutput
	ToListenerArrayOutputWithContext(context.Context) ListenerArrayOutput
}

type ListenerArray []ListenerInput

func (ListenerArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Listener)(nil))
}

func (i ListenerArray) ToListenerArrayOutput() ListenerArrayOutput {
	return i.ToListenerArrayOutputWithContext(context.Background())
}

func (i ListenerArray) ToListenerArrayOutputWithContext(ctx context.Context) ListenerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerArrayOutput)
}

// ListenerMapInput is an input type that accepts ListenerMap and ListenerMapOutput values.
// You can construct a concrete instance of `ListenerMapInput` via:
//
//          ListenerMap{ "key": ListenerArgs{...} }
type ListenerMapInput interface {
	pulumi.Input

	ToListenerMapOutput() ListenerMapOutput
	ToListenerMapOutputWithContext(context.Context) ListenerMapOutput
}

type ListenerMap map[string]ListenerInput

func (ListenerMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Listener)(nil))
}

func (i ListenerMap) ToListenerMapOutput() ListenerMapOutput {
	return i.ToListenerMapOutputWithContext(context.Background())
}

func (i ListenerMap) ToListenerMapOutputWithContext(ctx context.Context) ListenerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerMapOutput)
}

type ListenerOutput struct {
	*pulumi.OutputState
}

func (ListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Listener)(nil))
}

func (o ListenerOutput) ToListenerOutput() ListenerOutput {
	return o
}

func (o ListenerOutput) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return o
}

func (o ListenerOutput) ToListenerPtrOutput() ListenerPtrOutput {
	return o.ToListenerPtrOutputWithContext(context.Background())
}

func (o ListenerOutput) ToListenerPtrOutputWithContext(ctx context.Context) ListenerPtrOutput {
	return o.ApplyT(func(v Listener) *Listener {
		return &v
	}).(ListenerPtrOutput)
}

type ListenerPtrOutput struct {
	*pulumi.OutputState
}

func (ListenerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil))
}

func (o ListenerPtrOutput) ToListenerPtrOutput() ListenerPtrOutput {
	return o
}

func (o ListenerPtrOutput) ToListenerPtrOutputWithContext(ctx context.Context) ListenerPtrOutput {
	return o
}

type ListenerArrayOutput struct{ *pulumi.OutputState }

func (ListenerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Listener)(nil))
}

func (o ListenerArrayOutput) ToListenerArrayOutput() ListenerArrayOutput {
	return o
}

func (o ListenerArrayOutput) ToListenerArrayOutputWithContext(ctx context.Context) ListenerArrayOutput {
	return o
}

func (o ListenerArrayOutput) Index(i pulumi.IntInput) ListenerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Listener {
		return vs[0].([]Listener)[vs[1].(int)]
	}).(ListenerOutput)
}

type ListenerMapOutput struct{ *pulumi.OutputState }

func (ListenerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Listener)(nil))
}

func (o ListenerMapOutput) ToListenerMapOutput() ListenerMapOutput {
	return o
}

func (o ListenerMapOutput) ToListenerMapOutputWithContext(ctx context.Context) ListenerMapOutput {
	return o
}

func (o ListenerMapOutput) MapIndex(k pulumi.StringInput) ListenerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Listener {
		return vs[0].(map[string]Listener)[vs[1].(string)]
	}).(ListenerOutput)
}

func init() {
	pulumi.RegisterOutputType(ListenerOutput{})
	pulumi.RegisterOutputType(ListenerPtrOutput{})
	pulumi.RegisterOutputType(ListenerArrayOutput{})
	pulumi.RegisterOutputType(ListenerMapOutput{})
}
