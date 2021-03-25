// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `route53.ResolverRule` provides details about a specific Route53 Resolver rule.
//
// ## Example Usage
//
// The following example shows how to get a Route53 Resolver rule based on its associated domain name and rule type.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "subdomain.example.com"
// 		opt1 := "SYSTEM"
// 		_, err := route53.LookupResolverRule(ctx, &route53.LookupResolverRuleArgs{
// 			DomainName: &opt0,
// 			RuleType:   &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupResolverRule(ctx *pulumi.Context, args *LookupResolverRuleArgs, opts ...pulumi.InvokeOption) (*LookupResolverRuleResult, error) {
	var rv LookupResolverRuleResult
	err := ctx.Invoke("aws:route53/getResolverRule:getResolverRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResolverRule.
type LookupResolverRuleArgs struct {
	// The domain name the desired resolver rule forwards DNS queries for. Conflicts with `resolverRuleId`.
	DomainName *string `pulumi:"domainName"`
	// The friendly name of the desired resolver rule. Conflicts with `resolverRuleId`.
	Name *string `pulumi:"name"`
	// The ID of the outbound resolver endpoint of the desired resolver rule. Conflicts with `resolverRuleId`.
	ResolverEndpointId *string `pulumi:"resolverEndpointId"`
	// The ID of the desired resolver rule. Conflicts with `domainName`, `name`, `resolverEndpointId` and `ruleType`.
	ResolverRuleId *string `pulumi:"resolverRuleId"`
	// The rule type of the desired resolver rule. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`. Conflicts with `resolverRuleId`.
	RuleType *string `pulumi:"ruleType"`
	// A map of tags assigned to the resolver rule.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getResolverRule.
type LookupResolverRuleResult struct {
	// The ARN (Amazon Resource Name) for the resolver rule.
	Arn        string `pulumi:"arn"`
	DomainName string `pulumi:"domainName"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.
	OwnerId            string `pulumi:"ownerId"`
	ResolverEndpointId string `pulumi:"resolverEndpointId"`
	ResolverRuleId     string `pulumi:"resolverRuleId"`
	RuleType           string `pulumi:"ruleType"`
	// Whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account.
	// Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
	ShareStatus string `pulumi:"shareStatus"`
	// A map of tags assigned to the resolver rule.
	Tags map[string]string `pulumi:"tags"`
}
