// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage a GuardDuty filter.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/guardduty"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := guardduty.NewFilter(ctx, "myFilter", &guardduty.FilterArgs{
// 			Action:     pulumi.String("ARCHIVE"),
// 			DetectorId: pulumi.Any(aws_guardduty_detector.Example.Id),
// 			Rank:       pulumi.Int(1),
// 			FindingCriteria: &guardduty.FilterFindingCriteriaArgs{
// 				Criterions: guardduty.FilterFindingCriteriaCriterionArray{
// 					&guardduty.FilterFindingCriteriaCriterionArgs{
// 						Field: pulumi.String("region"),
// 						Equals: pulumi.StringArray{
// 							pulumi.String("eu-west-1"),
// 						},
// 					},
// 					&guardduty.FilterFindingCriteriaCriterionArgs{
// 						Field: pulumi.String("service.additionalInfo.threatListName"),
// 						NotEquals: pulumi.StringArray{
// 							pulumi.String("some-threat"),
// 							pulumi.String("another-threat"),
// 						},
// 					},
// 					&guardduty.FilterFindingCriteriaCriterionArgs{
// 						Field:       pulumi.String("updatedAt"),
// 						GreaterThan: pulumi.String("2020-01-01T00:00:00Z"),
// 						LessThan:    pulumi.String("2020-02-01T00:00:00Z"),
// 					},
// 					&guardduty.FilterFindingCriteriaCriterionArgs{
// 						Field:              pulumi.String("severity"),
// 						GreaterThanOrEqual: pulumi.String("4"),
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
//
// ## Import
//
// GuardDuty filters can be imported using the detector ID and filter's name separated by a colon, e.g.
//
// ```sh
//  $ pulumi import aws:guardduty/filter:Filter MyFilter 00b00fd5aecc0ab60a708659477e9617:MyFilter
// ```
type Filter struct {
	pulumi.CustomResourceState

	// Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
	Action pulumi.StringOutput `pulumi:"action"`
	// The ARN of the GuardDuty filter.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the filter.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// ID of a GuardDuty detector, attached to your account.
	DetectorId pulumi.StringOutput `pulumi:"detectorId"`
	// Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
	FindingCriteria FilterFindingCriteriaOutput `pulumi:"findingCriteria"`
	// The name of your filter.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
	Rank pulumi.IntOutput `pulumi:"rank"`
	// The tags that you want to add to the Filter resource. A tag consists of a key and a value. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewFilter registers a new resource with the given unique name, arguments, and options.
func NewFilter(ctx *pulumi.Context,
	name string, args *FilterArgs, opts ...pulumi.ResourceOption) (*Filter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.DetectorId == nil {
		return nil, errors.New("invalid value for required argument 'DetectorId'")
	}
	if args.FindingCriteria == nil {
		return nil, errors.New("invalid value for required argument 'FindingCriteria'")
	}
	if args.Rank == nil {
		return nil, errors.New("invalid value for required argument 'Rank'")
	}
	var resource Filter
	err := ctx.RegisterResource("aws:guardduty/filter:Filter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFilter gets an existing Filter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FilterState, opts ...pulumi.ResourceOption) (*Filter, error) {
	var resource Filter
	err := ctx.ReadResource("aws:guardduty/filter:Filter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Filter resources.
type filterState struct {
	// Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
	Action *string `pulumi:"action"`
	// The ARN of the GuardDuty filter.
	Arn *string `pulumi:"arn"`
	// Description of the filter.
	Description *string `pulumi:"description"`
	// ID of a GuardDuty detector, attached to your account.
	DetectorId *string `pulumi:"detectorId"`
	// Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
	FindingCriteria *FilterFindingCriteria `pulumi:"findingCriteria"`
	// The name of your filter.
	Name *string `pulumi:"name"`
	// Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
	Rank *int `pulumi:"rank"`
	// The tags that you want to add to the Filter resource. A tag consists of a key and a value. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type FilterState struct {
	// Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
	Action pulumi.StringPtrInput
	// The ARN of the GuardDuty filter.
	Arn pulumi.StringPtrInput
	// Description of the filter.
	Description pulumi.StringPtrInput
	// ID of a GuardDuty detector, attached to your account.
	DetectorId pulumi.StringPtrInput
	// Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
	FindingCriteria FilterFindingCriteriaPtrInput
	// The name of your filter.
	Name pulumi.StringPtrInput
	// Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
	Rank pulumi.IntPtrInput
	// The tags that you want to add to the Filter resource. A tag consists of a key and a value. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (FilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*filterState)(nil)).Elem()
}

type filterArgs struct {
	// Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
	Action string `pulumi:"action"`
	// Description of the filter.
	Description *string `pulumi:"description"`
	// ID of a GuardDuty detector, attached to your account.
	DetectorId string `pulumi:"detectorId"`
	// Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
	FindingCriteria FilterFindingCriteria `pulumi:"findingCriteria"`
	// The name of your filter.
	Name *string `pulumi:"name"`
	// Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
	Rank int `pulumi:"rank"`
	// The tags that you want to add to the Filter resource. A tag consists of a key and a value. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a Filter resource.
type FilterArgs struct {
	// Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
	Action pulumi.StringInput
	// Description of the filter.
	Description pulumi.StringPtrInput
	// ID of a GuardDuty detector, attached to your account.
	DetectorId pulumi.StringInput
	// Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
	FindingCriteria FilterFindingCriteriaInput
	// The name of your filter.
	Name pulumi.StringPtrInput
	// Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
	Rank pulumi.IntInput
	// The tags that you want to add to the Filter resource. A tag consists of a key and a value. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (FilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*filterArgs)(nil)).Elem()
}

type FilterInput interface {
	pulumi.Input

	ToFilterOutput() FilterOutput
	ToFilterOutputWithContext(ctx context.Context) FilterOutput
}

func (*Filter) ElementType() reflect.Type {
	return reflect.TypeOf((*Filter)(nil))
}

func (i *Filter) ToFilterOutput() FilterOutput {
	return i.ToFilterOutputWithContext(context.Background())
}

func (i *Filter) ToFilterOutputWithContext(ctx context.Context) FilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterOutput)
}

func (i *Filter) ToFilterPtrOutput() FilterPtrOutput {
	return i.ToFilterPtrOutputWithContext(context.Background())
}

func (i *Filter) ToFilterPtrOutputWithContext(ctx context.Context) FilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterPtrOutput)
}

type FilterPtrInput interface {
	pulumi.Input

	ToFilterPtrOutput() FilterPtrOutput
	ToFilterPtrOutputWithContext(ctx context.Context) FilterPtrOutput
}

type filterPtrType FilterArgs

func (*filterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Filter)(nil))
}

func (i *filterPtrType) ToFilterPtrOutput() FilterPtrOutput {
	return i.ToFilterPtrOutputWithContext(context.Background())
}

func (i *filterPtrType) ToFilterPtrOutputWithContext(ctx context.Context) FilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterPtrOutput)
}

// FilterArrayInput is an input type that accepts FilterArray and FilterArrayOutput values.
// You can construct a concrete instance of `FilterArrayInput` via:
//
//          FilterArray{ FilterArgs{...} }
type FilterArrayInput interface {
	pulumi.Input

	ToFilterArrayOutput() FilterArrayOutput
	ToFilterArrayOutputWithContext(context.Context) FilterArrayOutput
}

type FilterArray []FilterInput

func (FilterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Filter)(nil))
}

func (i FilterArray) ToFilterArrayOutput() FilterArrayOutput {
	return i.ToFilterArrayOutputWithContext(context.Background())
}

func (i FilterArray) ToFilterArrayOutputWithContext(ctx context.Context) FilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterArrayOutput)
}

// FilterMapInput is an input type that accepts FilterMap and FilterMapOutput values.
// You can construct a concrete instance of `FilterMapInput` via:
//
//          FilterMap{ "key": FilterArgs{...} }
type FilterMapInput interface {
	pulumi.Input

	ToFilterMapOutput() FilterMapOutput
	ToFilterMapOutputWithContext(context.Context) FilterMapOutput
}

type FilterMap map[string]FilterInput

func (FilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Filter)(nil))
}

func (i FilterMap) ToFilterMapOutput() FilterMapOutput {
	return i.ToFilterMapOutputWithContext(context.Background())
}

func (i FilterMap) ToFilterMapOutputWithContext(ctx context.Context) FilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterMapOutput)
}

type FilterOutput struct {
	*pulumi.OutputState
}

func (FilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Filter)(nil))
}

func (o FilterOutput) ToFilterOutput() FilterOutput {
	return o
}

func (o FilterOutput) ToFilterOutputWithContext(ctx context.Context) FilterOutput {
	return o
}

func (o FilterOutput) ToFilterPtrOutput() FilterPtrOutput {
	return o.ToFilterPtrOutputWithContext(context.Background())
}

func (o FilterOutput) ToFilterPtrOutputWithContext(ctx context.Context) FilterPtrOutput {
	return o.ApplyT(func(v Filter) *Filter {
		return &v
	}).(FilterPtrOutput)
}

type FilterPtrOutput struct {
	*pulumi.OutputState
}

func (FilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Filter)(nil))
}

func (o FilterPtrOutput) ToFilterPtrOutput() FilterPtrOutput {
	return o
}

func (o FilterPtrOutput) ToFilterPtrOutputWithContext(ctx context.Context) FilterPtrOutput {
	return o
}

type FilterArrayOutput struct{ *pulumi.OutputState }

func (FilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Filter)(nil))
}

func (o FilterArrayOutput) ToFilterArrayOutput() FilterArrayOutput {
	return o
}

func (o FilterArrayOutput) ToFilterArrayOutputWithContext(ctx context.Context) FilterArrayOutput {
	return o
}

func (o FilterArrayOutput) Index(i pulumi.IntInput) FilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Filter {
		return vs[0].([]Filter)[vs[1].(int)]
	}).(FilterOutput)
}

type FilterMapOutput struct{ *pulumi.OutputState }

func (FilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Filter)(nil))
}

func (o FilterMapOutput) ToFilterMapOutput() FilterMapOutput {
	return o
}

func (o FilterMapOutput) ToFilterMapOutputWithContext(ctx context.Context) FilterMapOutput {
	return o
}

func (o FilterMapOutput) MapIndex(k pulumi.StringInput) FilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Filter {
		return vs[0].(map[string]Filter)[vs[1].(string)]
	}).(FilterOutput)
}

func init() {
	pulumi.RegisterOutputType(FilterOutput{})
	pulumi.RegisterOutputType(FilterPtrOutput{})
	pulumi.RegisterOutputType(FilterArrayOutput{})
	pulumi.RegisterOutputType(FilterMapOutput{})
}
