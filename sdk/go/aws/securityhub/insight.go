// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package securityhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Security Hub custom insight resource. See the [Managing custom insights section](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-custom-insights.html) of the AWS User Guide for more information.
//
// ## Example Usage
// ### Filter by AWS account ID
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/securityhub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = securityhub.NewInsight(ctx, "exampleInsight", &securityhub.InsightArgs{
// 			Filters: &securityhub.InsightFiltersArgs{
// 				AwsAccountIds: securityhub.InsightFiltersAwsAccountIdArray{
// 					&securityhub.InsightFiltersAwsAccountIdArgs{
// 						Comparison: pulumi.String("EQUALS"),
// 						Value:      pulumi.String("1234567890"),
// 					},
// 					&securityhub.InsightFiltersAwsAccountIdArgs{
// 						Comparison: pulumi.String("EQUALS"),
// 						Value:      pulumi.String("09876543210"),
// 					},
// 				},
// 			},
// 			GroupByAttribute: pulumi.String("AwsAccountId"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleAccount,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Filter by date range
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/securityhub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = securityhub.NewInsight(ctx, "exampleInsight", &securityhub.InsightArgs{
// 			Filters: &securityhub.InsightFiltersArgs{
// 				CreatedAts: securityhub.InsightFiltersCreatedAtArray{
// 					&securityhub.InsightFiltersCreatedAtArgs{
// 						DateRange: &securityhub.InsightFiltersCreatedAtDateRangeArgs{
// 							Unit:  pulumi.String("DAYS"),
// 							Value: pulumi.Int(5),
// 						},
// 					},
// 				},
// 			},
// 			GroupByAttribute: pulumi.String("CreatedAt"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleAccount,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Filter by destination IPv4 address
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/securityhub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = securityhub.NewInsight(ctx, "exampleInsight", &securityhub.InsightArgs{
// 			Filters: &securityhub.InsightFiltersArgs{
// 				NetworkDestinationIpv4s: securityhub.InsightFiltersNetworkDestinationIpv4Array{
// 					&securityhub.InsightFiltersNetworkDestinationIpv4Args{
// 						Cidr: pulumi.String("10.0.0.0/16"),
// 					},
// 				},
// 			},
// 			GroupByAttribute: pulumi.String("NetworkDestinationIpV4"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleAccount,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Filter by finding's confidence
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/securityhub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = securityhub.NewInsight(ctx, "exampleInsight", &securityhub.InsightArgs{
// 			Filters: &securityhub.InsightFiltersArgs{
// 				Confidences: securityhub.InsightFiltersConfidenceArray{
// 					&securityhub.InsightFiltersConfidenceArgs{
// 						Gte: pulumi.String("80"),
// 					},
// 				},
// 			},
// 			GroupByAttribute: pulumi.String("Confidence"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleAccount,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Filter by resource tags
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/securityhub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = securityhub.NewInsight(ctx, "exampleInsight", &securityhub.InsightArgs{
// 			Filters: &securityhub.InsightFiltersArgs{
// 				ResourceTags: securityhub.InsightFiltersResourceTagArray{
// 					&securityhub.InsightFiltersResourceTagArgs{
// 						Comparison: pulumi.String("EQUALS"),
// 						Key:        pulumi.String("Environment"),
// 						Value:      pulumi.String("Production"),
// 					},
// 				},
// 			},
// 			GroupByAttribute: pulumi.String("ResourceTags"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleAccount,
// 		}))
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
// Security Hub insights can be imported using the ARN, e.g.
//
// ```sh
//  $ pulumi import aws:securityhub/insight:Insight example arn:aws:securityhub:us-west-2:1234567890:insight/1234567890/custom/91299ed7-abd0-4e44-a858-d0b15e37141a
// ```
type Insight struct {
	pulumi.CustomResourceState

	// ARN of the insight.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A configuration block including one or more (up to 10 distinct) attributes used to filter the findings included in the insight. The insight only includes findings that match criteria defined in the filters. See filters below for more details.
	Filters InsightFiltersOutput `pulumi:"filters"`
	// The attribute used to group the findings for the insight e.g. if an insight is grouped by `ResourceId`, then the insight produces a list of resource identifiers.
	GroupByAttribute pulumi.StringOutput `pulumi:"groupByAttribute"`
	// The name of the custom insight.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewInsight registers a new resource with the given unique name, arguments, and options.
func NewInsight(ctx *pulumi.Context,
	name string, args *InsightArgs, opts ...pulumi.ResourceOption) (*Insight, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Filters == nil {
		return nil, errors.New("invalid value for required argument 'Filters'")
	}
	if args.GroupByAttribute == nil {
		return nil, errors.New("invalid value for required argument 'GroupByAttribute'")
	}
	var resource Insight
	err := ctx.RegisterResource("aws:securityhub/insight:Insight", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInsight gets an existing Insight resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInsight(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InsightState, opts ...pulumi.ResourceOption) (*Insight, error) {
	var resource Insight
	err := ctx.ReadResource("aws:securityhub/insight:Insight", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Insight resources.
type insightState struct {
	// ARN of the insight.
	Arn *string `pulumi:"arn"`
	// A configuration block including one or more (up to 10 distinct) attributes used to filter the findings included in the insight. The insight only includes findings that match criteria defined in the filters. See filters below for more details.
	Filters *InsightFilters `pulumi:"filters"`
	// The attribute used to group the findings for the insight e.g. if an insight is grouped by `ResourceId`, then the insight produces a list of resource identifiers.
	GroupByAttribute *string `pulumi:"groupByAttribute"`
	// The name of the custom insight.
	Name *string `pulumi:"name"`
}

type InsightState struct {
	// ARN of the insight.
	Arn pulumi.StringPtrInput
	// A configuration block including one or more (up to 10 distinct) attributes used to filter the findings included in the insight. The insight only includes findings that match criteria defined in the filters. See filters below for more details.
	Filters InsightFiltersPtrInput
	// The attribute used to group the findings for the insight e.g. if an insight is grouped by `ResourceId`, then the insight produces a list of resource identifiers.
	GroupByAttribute pulumi.StringPtrInput
	// The name of the custom insight.
	Name pulumi.StringPtrInput
}

func (InsightState) ElementType() reflect.Type {
	return reflect.TypeOf((*insightState)(nil)).Elem()
}

type insightArgs struct {
	// A configuration block including one or more (up to 10 distinct) attributes used to filter the findings included in the insight. The insight only includes findings that match criteria defined in the filters. See filters below for more details.
	Filters InsightFilters `pulumi:"filters"`
	// The attribute used to group the findings for the insight e.g. if an insight is grouped by `ResourceId`, then the insight produces a list of resource identifiers.
	GroupByAttribute string `pulumi:"groupByAttribute"`
	// The name of the custom insight.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Insight resource.
type InsightArgs struct {
	// A configuration block including one or more (up to 10 distinct) attributes used to filter the findings included in the insight. The insight only includes findings that match criteria defined in the filters. See filters below for more details.
	Filters InsightFiltersInput
	// The attribute used to group the findings for the insight e.g. if an insight is grouped by `ResourceId`, then the insight produces a list of resource identifiers.
	GroupByAttribute pulumi.StringInput
	// The name of the custom insight.
	Name pulumi.StringPtrInput
}

func (InsightArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*insightArgs)(nil)).Elem()
}

type InsightInput interface {
	pulumi.Input

	ToInsightOutput() InsightOutput
	ToInsightOutputWithContext(ctx context.Context) InsightOutput
}

func (*Insight) ElementType() reflect.Type {
	return reflect.TypeOf((*Insight)(nil))
}

func (i *Insight) ToInsightOutput() InsightOutput {
	return i.ToInsightOutputWithContext(context.Background())
}

func (i *Insight) ToInsightOutputWithContext(ctx context.Context) InsightOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InsightOutput)
}

func (i *Insight) ToInsightPtrOutput() InsightPtrOutput {
	return i.ToInsightPtrOutputWithContext(context.Background())
}

func (i *Insight) ToInsightPtrOutputWithContext(ctx context.Context) InsightPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InsightPtrOutput)
}

type InsightPtrInput interface {
	pulumi.Input

	ToInsightPtrOutput() InsightPtrOutput
	ToInsightPtrOutputWithContext(ctx context.Context) InsightPtrOutput
}

type insightPtrType InsightArgs

func (*insightPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Insight)(nil))
}

func (i *insightPtrType) ToInsightPtrOutput() InsightPtrOutput {
	return i.ToInsightPtrOutputWithContext(context.Background())
}

func (i *insightPtrType) ToInsightPtrOutputWithContext(ctx context.Context) InsightPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InsightPtrOutput)
}

// InsightArrayInput is an input type that accepts InsightArray and InsightArrayOutput values.
// You can construct a concrete instance of `InsightArrayInput` via:
//
//          InsightArray{ InsightArgs{...} }
type InsightArrayInput interface {
	pulumi.Input

	ToInsightArrayOutput() InsightArrayOutput
	ToInsightArrayOutputWithContext(context.Context) InsightArrayOutput
}

type InsightArray []InsightInput

func (InsightArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Insight)(nil)).Elem()
}

func (i InsightArray) ToInsightArrayOutput() InsightArrayOutput {
	return i.ToInsightArrayOutputWithContext(context.Background())
}

func (i InsightArray) ToInsightArrayOutputWithContext(ctx context.Context) InsightArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InsightArrayOutput)
}

// InsightMapInput is an input type that accepts InsightMap and InsightMapOutput values.
// You can construct a concrete instance of `InsightMapInput` via:
//
//          InsightMap{ "key": InsightArgs{...} }
type InsightMapInput interface {
	pulumi.Input

	ToInsightMapOutput() InsightMapOutput
	ToInsightMapOutputWithContext(context.Context) InsightMapOutput
}

type InsightMap map[string]InsightInput

func (InsightMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Insight)(nil)).Elem()
}

func (i InsightMap) ToInsightMapOutput() InsightMapOutput {
	return i.ToInsightMapOutputWithContext(context.Background())
}

func (i InsightMap) ToInsightMapOutputWithContext(ctx context.Context) InsightMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InsightMapOutput)
}

type InsightOutput struct{ *pulumi.OutputState }

func (InsightOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Insight)(nil))
}

func (o InsightOutput) ToInsightOutput() InsightOutput {
	return o
}

func (o InsightOutput) ToInsightOutputWithContext(ctx context.Context) InsightOutput {
	return o
}

func (o InsightOutput) ToInsightPtrOutput() InsightPtrOutput {
	return o.ToInsightPtrOutputWithContext(context.Background())
}

func (o InsightOutput) ToInsightPtrOutputWithContext(ctx context.Context) InsightPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Insight) *Insight {
		return &v
	}).(InsightPtrOutput)
}

type InsightPtrOutput struct{ *pulumi.OutputState }

func (InsightPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Insight)(nil))
}

func (o InsightPtrOutput) ToInsightPtrOutput() InsightPtrOutput {
	return o
}

func (o InsightPtrOutput) ToInsightPtrOutputWithContext(ctx context.Context) InsightPtrOutput {
	return o
}

func (o InsightPtrOutput) Elem() InsightOutput {
	return o.ApplyT(func(v *Insight) Insight {
		if v != nil {
			return *v
		}
		var ret Insight
		return ret
	}).(InsightOutput)
}

type InsightArrayOutput struct{ *pulumi.OutputState }

func (InsightArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Insight)(nil))
}

func (o InsightArrayOutput) ToInsightArrayOutput() InsightArrayOutput {
	return o
}

func (o InsightArrayOutput) ToInsightArrayOutputWithContext(ctx context.Context) InsightArrayOutput {
	return o
}

func (o InsightArrayOutput) Index(i pulumi.IntInput) InsightOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Insight {
		return vs[0].([]Insight)[vs[1].(int)]
	}).(InsightOutput)
}

type InsightMapOutput struct{ *pulumi.OutputState }

func (InsightMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Insight)(nil))
}

func (o InsightMapOutput) ToInsightMapOutput() InsightMapOutput {
	return o
}

func (o InsightMapOutput) ToInsightMapOutputWithContext(ctx context.Context) InsightMapOutput {
	return o
}

func (o InsightMapOutput) MapIndex(k pulumi.StringInput) InsightOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Insight {
		return vs[0].(map[string]Insight)[vs[1].(string)]
	}).(InsightOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InsightInput)(nil)).Elem(), &Insight{})
	pulumi.RegisterInputType(reflect.TypeOf((*InsightPtrInput)(nil)).Elem(), &Insight{})
	pulumi.RegisterInputType(reflect.TypeOf((*InsightArrayInput)(nil)).Elem(), InsightArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InsightMapInput)(nil)).Elem(), InsightMap{})
	pulumi.RegisterOutputType(InsightOutput{})
	pulumi.RegisterOutputType(InsightPtrOutput{})
	pulumi.RegisterOutputType(InsightArrayOutput{})
	pulumi.RegisterOutputType(InsightMapOutput{})
}
