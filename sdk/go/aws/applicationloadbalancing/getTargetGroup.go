// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package applicationloadbalancing

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// > **Note:** `alb.TargetGroup` is known as `lb.TargetGroup`. The functionality is identical.
//
// Provides information about a Load Balancer Target Group.
//
// This data source can prove useful when a module accepts an LB Target Group as an
// input variable and needs to know its attributes. It can also be used to get the ARN of
// an LB Target Group for use in other resources, given LB Target Group name.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		lbTgArn := ""
// 		if param := cfg.Get("lbTgArn"); param != "" {
// 			lbTgArn = param
// 		}
// 		lbTgName := ""
// 		if param := cfg.Get("lbTgName"); param != "" {
// 			lbTgName = param
// 		}
// 		opt0 := lbTgArn
// 		opt1 := lbTgName
// 		_, err := lb.LookupTargetGroup(ctx, &lb.LookupTargetGroupArgs{
// 			Arn:  &opt0,
// 			Name: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Deprecated: aws.applicationloadbalancing.getTargetGroup has been deprecated in favor of aws.alb.getTargetGroup
func LookupTargetGroup(ctx *pulumi.Context, args *LookupTargetGroupArgs, opts ...pulumi.InvokeOption) (*LookupTargetGroupResult, error) {
	var rv LookupTargetGroupResult
	err := ctx.Invoke("aws:applicationloadbalancing/getTargetGroup:getTargetGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTargetGroup.
type LookupTargetGroupArgs struct {
	// The full ARN of the target group.
	Arn *string `pulumi:"arn"`
	// The unique name of the target group.
	Name *string           `pulumi:"name"`
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getTargetGroup.
type LookupTargetGroupResult struct {
	Arn                 string                    `pulumi:"arn"`
	ArnSuffix           string                    `pulumi:"arnSuffix"`
	DeregistrationDelay int                       `pulumi:"deregistrationDelay"`
	HealthCheck         GetTargetGroupHealthCheck `pulumi:"healthCheck"`
	// The provider-assigned unique ID for this managed resource.
	Id                             string                   `pulumi:"id"`
	LambdaMultiValueHeadersEnabled bool                     `pulumi:"lambdaMultiValueHeadersEnabled"`
	LoadBalancingAlgorithmType     string                   `pulumi:"loadBalancingAlgorithmType"`
	Name                           string                   `pulumi:"name"`
	Port                           int                      `pulumi:"port"`
	PreserveClientIp               string                   `pulumi:"preserveClientIp"`
	Protocol                       string                   `pulumi:"protocol"`
	ProtocolVersion                string                   `pulumi:"protocolVersion"`
	ProxyProtocolV2                bool                     `pulumi:"proxyProtocolV2"`
	SlowStart                      int                      `pulumi:"slowStart"`
	Stickiness                     GetTargetGroupStickiness `pulumi:"stickiness"`
	Tags                           map[string]string        `pulumi:"tags"`
	TargetType                     string                   `pulumi:"targetType"`
	VpcId                          string                   `pulumi:"vpcId"`
}
