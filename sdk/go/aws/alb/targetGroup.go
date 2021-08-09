// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Target Group resource for use with Load Balancer resources.
//
// > **Note:** `alb.TargetGroup` is known as `lb.TargetGroup`. The functionality is identical.
//
// ## Example Usage
// ### Instance Target Group
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewTargetGroup(ctx, "test", &lb.TargetGroupArgs{
// 			Port:     pulumi.Int(80),
// 			Protocol: pulumi.String("HTTP"),
// 			VpcId:    main.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### IP Target Group
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewTargetGroup(ctx, "ip_example", &lb.TargetGroupArgs{
// 			Port:       pulumi.Int(80),
// 			Protocol:   pulumi.String("HTTP"),
// 			TargetType: pulumi.String("ip"),
// 			VpcId:      main.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Lambda Target Group
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
// 		_, err := lb.NewTargetGroup(ctx, "lambda_example", &lb.TargetGroupArgs{
// 			TargetType: pulumi.String("lambda"),
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
// Target Groups can be imported using their ARN, e.g.
//
// ```sh
//  $ pulumi import aws:alb/targetGroup:TargetGroup app_front_end arn:aws:elasticloadbalancing:us-west-2:187416307283:targetgroup/app-front-end/20cfe21448b66314
// ```
type TargetGroup struct {
	pulumi.CustomResourceState

	// ARN of the Target Group (matches `id`).
	Arn pulumi.StringOutput `pulumi:"arn"`
	// ARN suffix for use with CloudWatch Metrics.
	ArnSuffix pulumi.StringOutput `pulumi:"arnSuffix"`
	// Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
	DeregistrationDelay pulumi.IntPtrOutput `pulumi:"deregistrationDelay"`
	// Health Check configuration block. Detailed below.
	HealthCheck TargetGroupHealthCheckOutput `pulumi:"healthCheck"`
	// Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `targetType` is `lambda`. Default is `false`.
	LambdaMultiValueHeadersEnabled pulumi.BoolPtrOutput `pulumi:"lambdaMultiValueHeadersEnabled"`
	// Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is `roundRobin` or `leastOutstandingRequests`. The default is `roundRobin`.
	LoadBalancingAlgorithmType pulumi.StringOutput `pulumi:"loadBalancingAlgorithmType"`
	// Name of the Target Group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// Port to use to connect with the target. Valid values are either ports 1-65535, or `traffic-port`. Defaults to `traffic-port`.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// Whether client IP preservation is enabled. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#client-ip-preservation) for more information.
	PreserveClientIp pulumi.StringOutput `pulumi:"preserveClientIp"`
	// Protocol to use to connect with the target. Defaults to `HTTP`. Not applicable when `targetType` is `lambda`.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// Only applicable when `protocol` is `HTTP` or `HTTPS`. The protocol version. Specify GRPC to send requests to targets using gRPC. Specify HTTP2 to send requests to targets using HTTP/2. The default is HTTP1, which sends requests to targets using HTTP/1.1
	ProtocolVersion pulumi.StringOutput `pulumi:"protocolVersion"`
	// Whether to enable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information. Default is `false`.
	ProxyProtocolV2 pulumi.BoolPtrOutput `pulumi:"proxyProtocolV2"`
	// Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
	SlowStart pulumi.IntPtrOutput `pulumi:"slowStart"`
	// Stickiness configuration block. Detailed below.
	Stickiness TargetGroupStickinessOutput `pulumi:"stickiness"`
	// Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Type of target that you must specify when registering targets with this target group. The possible values are `instance` (targets are specified by instance ID) or `ip` (targets are specified by IP address) or `lambda` (targets are specified by lambda arn). The default is `instance`. Note that you can't specify targets for a target group using both instance IDs and IP addresses. If the target type is `ip`, specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group, the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10). You can't specify publicly routable IP addresses.
	TargetType pulumi.StringPtrOutput `pulumi:"targetType"`
	// Identifier of the VPC in which to create the target group. Required when `targetType` is `instance` or `ip`. Does not apply when `targetType` is `lambda`.
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
}

// NewTargetGroup registers a new resource with the given unique name, arguments, and options.
func NewTargetGroup(ctx *pulumi.Context,
	name string, args *TargetGroupArgs, opts ...pulumi.ResourceOption) (*TargetGroup, error) {
	if args == nil {
		args = &TargetGroupArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("aws:applicationloadbalancing/targetGroup:TargetGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource TargetGroup
	err := ctx.RegisterResource("aws:alb/targetGroup:TargetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetGroup gets an existing TargetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetGroupState, opts ...pulumi.ResourceOption) (*TargetGroup, error) {
	var resource TargetGroup
	err := ctx.ReadResource("aws:alb/targetGroup:TargetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetGroup resources.
type targetGroupState struct {
	// ARN of the Target Group (matches `id`).
	Arn *string `pulumi:"arn"`
	// ARN suffix for use with CloudWatch Metrics.
	ArnSuffix *string `pulumi:"arnSuffix"`
	// Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
	DeregistrationDelay *int `pulumi:"deregistrationDelay"`
	// Health Check configuration block. Detailed below.
	HealthCheck *TargetGroupHealthCheck `pulumi:"healthCheck"`
	// Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `targetType` is `lambda`. Default is `false`.
	LambdaMultiValueHeadersEnabled *bool `pulumi:"lambdaMultiValueHeadersEnabled"`
	// Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is `roundRobin` or `leastOutstandingRequests`. The default is `roundRobin`.
	LoadBalancingAlgorithmType *string `pulumi:"loadBalancingAlgorithmType"`
	// Name of the Target Group.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
	NamePrefix *string `pulumi:"namePrefix"`
	// Port to use to connect with the target. Valid values are either ports 1-65535, or `traffic-port`. Defaults to `traffic-port`.
	Port *int `pulumi:"port"`
	// Whether client IP preservation is enabled. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#client-ip-preservation) for more information.
	PreserveClientIp *string `pulumi:"preserveClientIp"`
	// Protocol to use to connect with the target. Defaults to `HTTP`. Not applicable when `targetType` is `lambda`.
	Protocol *string `pulumi:"protocol"`
	// Only applicable when `protocol` is `HTTP` or `HTTPS`. The protocol version. Specify GRPC to send requests to targets using gRPC. Specify HTTP2 to send requests to targets using HTTP/2. The default is HTTP1, which sends requests to targets using HTTP/1.1
	ProtocolVersion *string `pulumi:"protocolVersion"`
	// Whether to enable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information. Default is `false`.
	ProxyProtocolV2 *bool `pulumi:"proxyProtocolV2"`
	// Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
	SlowStart *int `pulumi:"slowStart"`
	// Stickiness configuration block. Detailed below.
	Stickiness *TargetGroupStickiness `pulumi:"stickiness"`
	// Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Type of target that you must specify when registering targets with this target group. The possible values are `instance` (targets are specified by instance ID) or `ip` (targets are specified by IP address) or `lambda` (targets are specified by lambda arn). The default is `instance`. Note that you can't specify targets for a target group using both instance IDs and IP addresses. If the target type is `ip`, specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group, the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10). You can't specify publicly routable IP addresses.
	TargetType *string `pulumi:"targetType"`
	// Identifier of the VPC in which to create the target group. Required when `targetType` is `instance` or `ip`. Does not apply when `targetType` is `lambda`.
	VpcId *string `pulumi:"vpcId"`
}

type TargetGroupState struct {
	// ARN of the Target Group (matches `id`).
	Arn pulumi.StringPtrInput
	// ARN suffix for use with CloudWatch Metrics.
	ArnSuffix pulumi.StringPtrInput
	// Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
	DeregistrationDelay pulumi.IntPtrInput
	// Health Check configuration block. Detailed below.
	HealthCheck TargetGroupHealthCheckPtrInput
	// Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `targetType` is `lambda`. Default is `false`.
	LambdaMultiValueHeadersEnabled pulumi.BoolPtrInput
	// Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is `roundRobin` or `leastOutstandingRequests`. The default is `roundRobin`.
	LoadBalancingAlgorithmType pulumi.StringPtrInput
	// Name of the Target Group.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
	NamePrefix pulumi.StringPtrInput
	// Port to use to connect with the target. Valid values are either ports 1-65535, or `traffic-port`. Defaults to `traffic-port`.
	Port pulumi.IntPtrInput
	// Whether client IP preservation is enabled. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#client-ip-preservation) for more information.
	PreserveClientIp pulumi.StringPtrInput
	// Protocol to use to connect with the target. Defaults to `HTTP`. Not applicable when `targetType` is `lambda`.
	Protocol pulumi.StringPtrInput
	// Only applicable when `protocol` is `HTTP` or `HTTPS`. The protocol version. Specify GRPC to send requests to targets using gRPC. Specify HTTP2 to send requests to targets using HTTP/2. The default is HTTP1, which sends requests to targets using HTTP/1.1
	ProtocolVersion pulumi.StringPtrInput
	// Whether to enable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information. Default is `false`.
	ProxyProtocolV2 pulumi.BoolPtrInput
	// Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
	SlowStart pulumi.IntPtrInput
	// Stickiness configuration block. Detailed below.
	Stickiness TargetGroupStickinessPtrInput
	// Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Type of target that you must specify when registering targets with this target group. The possible values are `instance` (targets are specified by instance ID) or `ip` (targets are specified by IP address) or `lambda` (targets are specified by lambda arn). The default is `instance`. Note that you can't specify targets for a target group using both instance IDs and IP addresses. If the target type is `ip`, specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group, the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10). You can't specify publicly routable IP addresses.
	TargetType pulumi.StringPtrInput
	// Identifier of the VPC in which to create the target group. Required when `targetType` is `instance` or `ip`. Does not apply when `targetType` is `lambda`.
	VpcId pulumi.StringPtrInput
}

func (TargetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGroupState)(nil)).Elem()
}

type targetGroupArgs struct {
	// Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
	DeregistrationDelay *int `pulumi:"deregistrationDelay"`
	// Health Check configuration block. Detailed below.
	HealthCheck *TargetGroupHealthCheck `pulumi:"healthCheck"`
	// Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `targetType` is `lambda`. Default is `false`.
	LambdaMultiValueHeadersEnabled *bool `pulumi:"lambdaMultiValueHeadersEnabled"`
	// Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is `roundRobin` or `leastOutstandingRequests`. The default is `roundRobin`.
	LoadBalancingAlgorithmType *string `pulumi:"loadBalancingAlgorithmType"`
	// Name of the Target Group.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
	NamePrefix *string `pulumi:"namePrefix"`
	// Port to use to connect with the target. Valid values are either ports 1-65535, or `traffic-port`. Defaults to `traffic-port`.
	Port *int `pulumi:"port"`
	// Whether client IP preservation is enabled. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#client-ip-preservation) for more information.
	PreserveClientIp *string `pulumi:"preserveClientIp"`
	// Protocol to use to connect with the target. Defaults to `HTTP`. Not applicable when `targetType` is `lambda`.
	Protocol *string `pulumi:"protocol"`
	// Only applicable when `protocol` is `HTTP` or `HTTPS`. The protocol version. Specify GRPC to send requests to targets using gRPC. Specify HTTP2 to send requests to targets using HTTP/2. The default is HTTP1, which sends requests to targets using HTTP/1.1
	ProtocolVersion *string `pulumi:"protocolVersion"`
	// Whether to enable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information. Default is `false`.
	ProxyProtocolV2 *bool `pulumi:"proxyProtocolV2"`
	// Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
	SlowStart *int `pulumi:"slowStart"`
	// Stickiness configuration block. Detailed below.
	Stickiness *TargetGroupStickiness `pulumi:"stickiness"`
	// Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Type of target that you must specify when registering targets with this target group. The possible values are `instance` (targets are specified by instance ID) or `ip` (targets are specified by IP address) or `lambda` (targets are specified by lambda arn). The default is `instance`. Note that you can't specify targets for a target group using both instance IDs and IP addresses. If the target type is `ip`, specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group, the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10). You can't specify publicly routable IP addresses.
	TargetType *string `pulumi:"targetType"`
	// Identifier of the VPC in which to create the target group. Required when `targetType` is `instance` or `ip`. Does not apply when `targetType` is `lambda`.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a TargetGroup resource.
type TargetGroupArgs struct {
	// Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
	DeregistrationDelay pulumi.IntPtrInput
	// Health Check configuration block. Detailed below.
	HealthCheck TargetGroupHealthCheckPtrInput
	// Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `targetType` is `lambda`. Default is `false`.
	LambdaMultiValueHeadersEnabled pulumi.BoolPtrInput
	// Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is `roundRobin` or `leastOutstandingRequests`. The default is `roundRobin`.
	LoadBalancingAlgorithmType pulumi.StringPtrInput
	// Name of the Target Group.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
	NamePrefix pulumi.StringPtrInput
	// Port to use to connect with the target. Valid values are either ports 1-65535, or `traffic-port`. Defaults to `traffic-port`.
	Port pulumi.IntPtrInput
	// Whether client IP preservation is enabled. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#client-ip-preservation) for more information.
	PreserveClientIp pulumi.StringPtrInput
	// Protocol to use to connect with the target. Defaults to `HTTP`. Not applicable when `targetType` is `lambda`.
	Protocol pulumi.StringPtrInput
	// Only applicable when `protocol` is `HTTP` or `HTTPS`. The protocol version. Specify GRPC to send requests to targets using gRPC. Specify HTTP2 to send requests to targets using HTTP/2. The default is HTTP1, which sends requests to targets using HTTP/1.1
	ProtocolVersion pulumi.StringPtrInput
	// Whether to enable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information. Default is `false`.
	ProxyProtocolV2 pulumi.BoolPtrInput
	// Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
	SlowStart pulumi.IntPtrInput
	// Stickiness configuration block. Detailed below.
	Stickiness TargetGroupStickinessPtrInput
	// Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Type of target that you must specify when registering targets with this target group. The possible values are `instance` (targets are specified by instance ID) or `ip` (targets are specified by IP address) or `lambda` (targets are specified by lambda arn). The default is `instance`. Note that you can't specify targets for a target group using both instance IDs and IP addresses. If the target type is `ip`, specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group, the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10). You can't specify publicly routable IP addresses.
	TargetType pulumi.StringPtrInput
	// Identifier of the VPC in which to create the target group. Required when `targetType` is `instance` or `ip`. Does not apply when `targetType` is `lambda`.
	VpcId pulumi.StringPtrInput
}

func (TargetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGroupArgs)(nil)).Elem()
}

type TargetGroupInput interface {
	pulumi.Input

	ToTargetGroupOutput() TargetGroupOutput
	ToTargetGroupOutputWithContext(ctx context.Context) TargetGroupOutput
}

func (*TargetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetGroup)(nil))
}

func (i *TargetGroup) ToTargetGroupOutput() TargetGroupOutput {
	return i.ToTargetGroupOutputWithContext(context.Background())
}

func (i *TargetGroup) ToTargetGroupOutputWithContext(ctx context.Context) TargetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetGroupOutput)
}

func (i *TargetGroup) ToTargetGroupPtrOutput() TargetGroupPtrOutput {
	return i.ToTargetGroupPtrOutputWithContext(context.Background())
}

func (i *TargetGroup) ToTargetGroupPtrOutputWithContext(ctx context.Context) TargetGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetGroupPtrOutput)
}

type TargetGroupPtrInput interface {
	pulumi.Input

	ToTargetGroupPtrOutput() TargetGroupPtrOutput
	ToTargetGroupPtrOutputWithContext(ctx context.Context) TargetGroupPtrOutput
}

type targetGroupPtrType TargetGroupArgs

func (*targetGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetGroup)(nil))
}

func (i *targetGroupPtrType) ToTargetGroupPtrOutput() TargetGroupPtrOutput {
	return i.ToTargetGroupPtrOutputWithContext(context.Background())
}

func (i *targetGroupPtrType) ToTargetGroupPtrOutputWithContext(ctx context.Context) TargetGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetGroupPtrOutput)
}

// TargetGroupArrayInput is an input type that accepts TargetGroupArray and TargetGroupArrayOutput values.
// You can construct a concrete instance of `TargetGroupArrayInput` via:
//
//          TargetGroupArray{ TargetGroupArgs{...} }
type TargetGroupArrayInput interface {
	pulumi.Input

	ToTargetGroupArrayOutput() TargetGroupArrayOutput
	ToTargetGroupArrayOutputWithContext(context.Context) TargetGroupArrayOutput
}

type TargetGroupArray []TargetGroupInput

func (TargetGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TargetGroup)(nil)).Elem()
}

func (i TargetGroupArray) ToTargetGroupArrayOutput() TargetGroupArrayOutput {
	return i.ToTargetGroupArrayOutputWithContext(context.Background())
}

func (i TargetGroupArray) ToTargetGroupArrayOutputWithContext(ctx context.Context) TargetGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetGroupArrayOutput)
}

// TargetGroupMapInput is an input type that accepts TargetGroupMap and TargetGroupMapOutput values.
// You can construct a concrete instance of `TargetGroupMapInput` via:
//
//          TargetGroupMap{ "key": TargetGroupArgs{...} }
type TargetGroupMapInput interface {
	pulumi.Input

	ToTargetGroupMapOutput() TargetGroupMapOutput
	ToTargetGroupMapOutputWithContext(context.Context) TargetGroupMapOutput
}

type TargetGroupMap map[string]TargetGroupInput

func (TargetGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TargetGroup)(nil)).Elem()
}

func (i TargetGroupMap) ToTargetGroupMapOutput() TargetGroupMapOutput {
	return i.ToTargetGroupMapOutputWithContext(context.Background())
}

func (i TargetGroupMap) ToTargetGroupMapOutputWithContext(ctx context.Context) TargetGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetGroupMapOutput)
}

type TargetGroupOutput struct{ *pulumi.OutputState }

func (TargetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetGroup)(nil))
}

func (o TargetGroupOutput) ToTargetGroupOutput() TargetGroupOutput {
	return o
}

func (o TargetGroupOutput) ToTargetGroupOutputWithContext(ctx context.Context) TargetGroupOutput {
	return o
}

func (o TargetGroupOutput) ToTargetGroupPtrOutput() TargetGroupPtrOutput {
	return o.ToTargetGroupPtrOutputWithContext(context.Background())
}

func (o TargetGroupOutput) ToTargetGroupPtrOutputWithContext(ctx context.Context) TargetGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TargetGroup) *TargetGroup {
		return &v
	}).(TargetGroupPtrOutput)
}

type TargetGroupPtrOutput struct{ *pulumi.OutputState }

func (TargetGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetGroup)(nil))
}

func (o TargetGroupPtrOutput) ToTargetGroupPtrOutput() TargetGroupPtrOutput {
	return o
}

func (o TargetGroupPtrOutput) ToTargetGroupPtrOutputWithContext(ctx context.Context) TargetGroupPtrOutput {
	return o
}

func (o TargetGroupPtrOutput) Elem() TargetGroupOutput {
	return o.ApplyT(func(v *TargetGroup) TargetGroup {
		if v != nil {
			return *v
		}
		var ret TargetGroup
		return ret
	}).(TargetGroupOutput)
}

type TargetGroupArrayOutput struct{ *pulumi.OutputState }

func (TargetGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetGroup)(nil))
}

func (o TargetGroupArrayOutput) ToTargetGroupArrayOutput() TargetGroupArrayOutput {
	return o
}

func (o TargetGroupArrayOutput) ToTargetGroupArrayOutputWithContext(ctx context.Context) TargetGroupArrayOutput {
	return o
}

func (o TargetGroupArrayOutput) Index(i pulumi.IntInput) TargetGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetGroup {
		return vs[0].([]TargetGroup)[vs[1].(int)]
	}).(TargetGroupOutput)
}

type TargetGroupMapOutput struct{ *pulumi.OutputState }

func (TargetGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TargetGroup)(nil))
}

func (o TargetGroupMapOutput) ToTargetGroupMapOutput() TargetGroupMapOutput {
	return o
}

func (o TargetGroupMapOutput) ToTargetGroupMapOutputWithContext(ctx context.Context) TargetGroupMapOutput {
	return o
}

func (o TargetGroupMapOutput) MapIndex(k pulumi.StringInput) TargetGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TargetGroup {
		return vs[0].(map[string]TargetGroup)[vs[1].(string)]
	}).(TargetGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(TargetGroupOutput{})
	pulumi.RegisterOutputType(TargetGroupPtrOutput{})
	pulumi.RegisterOutputType(TargetGroupArrayOutput{})
	pulumi.RegisterOutputType(TargetGroupMapOutput{})
}
