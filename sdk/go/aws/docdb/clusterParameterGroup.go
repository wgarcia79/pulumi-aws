// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a DocumentDB Cluster Parameter Group
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/docdb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := docdb.NewClusterParameterGroup(ctx, "example", &docdb.ClusterParameterGroupArgs{
// 			Description: pulumi.String("docdb cluster parameter group"),
// 			Family:      pulumi.String("docdb3.6"),
// 			Parameters: docdb.ClusterParameterGroupParameterArray{
// 				&docdb.ClusterParameterGroupParameterArgs{
// 					Name:  pulumi.String("tls"),
// 					Value: pulumi.String("enabled"),
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
// DocumentDB Cluster Parameter Groups can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:docdb/clusterParameterGroup:ClusterParameterGroup cluster_pg production-pg-1
// ```
type ClusterParameterGroup struct {
	pulumi.CustomResourceState

	// The ARN of the documentDB cluster parameter group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the documentDB cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The family of the documentDB cluster parameter group.
	Family pulumi.StringOutput `pulumi:"family"`
	// The name of the documentDB parameter.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// A list of documentDB parameters to apply. Setting parameters to system default values may show a difference on imported resources.
	Parameters ClusterParameterGroupParameterArrayOutput `pulumi:"parameters"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewClusterParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewClusterParameterGroup(ctx *pulumi.Context,
	name string, args *ClusterParameterGroupArgs, opts ...pulumi.ResourceOption) (*ClusterParameterGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Family == nil {
		return nil, errors.New("invalid value for required argument 'Family'")
	}
	var resource ClusterParameterGroup
	err := ctx.RegisterResource("aws:docdb/clusterParameterGroup:ClusterParameterGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterParameterGroup gets an existing ClusterParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterParameterGroupState, opts ...pulumi.ResourceOption) (*ClusterParameterGroup, error) {
	var resource ClusterParameterGroup
	err := ctx.ReadResource("aws:docdb/clusterParameterGroup:ClusterParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterParameterGroup resources.
type clusterParameterGroupState struct {
	// The ARN of the documentDB cluster parameter group.
	Arn *string `pulumi:"arn"`
	// The description of the documentDB cluster parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the documentDB cluster parameter group.
	Family *string `pulumi:"family"`
	// The name of the documentDB parameter.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of documentDB parameters to apply. Setting parameters to system default values may show a difference on imported resources.
	Parameters []ClusterParameterGroupParameter `pulumi:"parameters"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ClusterParameterGroupState struct {
	// The ARN of the documentDB cluster parameter group.
	Arn pulumi.StringPtrInput
	// The description of the documentDB cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the documentDB cluster parameter group.
	Family pulumi.StringPtrInput
	// The name of the documentDB parameter.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// A list of documentDB parameters to apply. Setting parameters to system default values may show a difference on imported resources.
	Parameters ClusterParameterGroupParameterArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ClusterParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterParameterGroupState)(nil)).Elem()
}

type clusterParameterGroupArgs struct {
	// The description of the documentDB cluster parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the documentDB cluster parameter group.
	Family string `pulumi:"family"`
	// The name of the documentDB parameter.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of documentDB parameters to apply. Setting parameters to system default values may show a difference on imported resources.
	Parameters []ClusterParameterGroupParameter `pulumi:"parameters"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterParameterGroup resource.
type ClusterParameterGroupArgs struct {
	// The description of the documentDB cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the documentDB cluster parameter group.
	Family pulumi.StringInput
	// The name of the documentDB parameter.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// A list of documentDB parameters to apply. Setting parameters to system default values may show a difference on imported resources.
	Parameters ClusterParameterGroupParameterArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ClusterParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterParameterGroupArgs)(nil)).Elem()
}

type ClusterParameterGroupInput interface {
	pulumi.Input

	ToClusterParameterGroupOutput() ClusterParameterGroupOutput
	ToClusterParameterGroupOutputWithContext(ctx context.Context) ClusterParameterGroupOutput
}

func (*ClusterParameterGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterParameterGroup)(nil))
}

func (i *ClusterParameterGroup) ToClusterParameterGroupOutput() ClusterParameterGroupOutput {
	return i.ToClusterParameterGroupOutputWithContext(context.Background())
}

func (i *ClusterParameterGroup) ToClusterParameterGroupOutputWithContext(ctx context.Context) ClusterParameterGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupOutput)
}

func (i *ClusterParameterGroup) ToClusterParameterGroupPtrOutput() ClusterParameterGroupPtrOutput {
	return i.ToClusterParameterGroupPtrOutputWithContext(context.Background())
}

func (i *ClusterParameterGroup) ToClusterParameterGroupPtrOutputWithContext(ctx context.Context) ClusterParameterGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupPtrOutput)
}

type ClusterParameterGroupPtrInput interface {
	pulumi.Input

	ToClusterParameterGroupPtrOutput() ClusterParameterGroupPtrOutput
	ToClusterParameterGroupPtrOutputWithContext(ctx context.Context) ClusterParameterGroupPtrOutput
}

type clusterParameterGroupPtrType ClusterParameterGroupArgs

func (*clusterParameterGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterParameterGroup)(nil))
}

func (i *clusterParameterGroupPtrType) ToClusterParameterGroupPtrOutput() ClusterParameterGroupPtrOutput {
	return i.ToClusterParameterGroupPtrOutputWithContext(context.Background())
}

func (i *clusterParameterGroupPtrType) ToClusterParameterGroupPtrOutputWithContext(ctx context.Context) ClusterParameterGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupPtrOutput)
}

// ClusterParameterGroupArrayInput is an input type that accepts ClusterParameterGroupArray and ClusterParameterGroupArrayOutput values.
// You can construct a concrete instance of `ClusterParameterGroupArrayInput` via:
//
//          ClusterParameterGroupArray{ ClusterParameterGroupArgs{...} }
type ClusterParameterGroupArrayInput interface {
	pulumi.Input

	ToClusterParameterGroupArrayOutput() ClusterParameterGroupArrayOutput
	ToClusterParameterGroupArrayOutputWithContext(context.Context) ClusterParameterGroupArrayOutput
}

type ClusterParameterGroupArray []ClusterParameterGroupInput

func (ClusterParameterGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ClusterParameterGroup)(nil))
}

func (i ClusterParameterGroupArray) ToClusterParameterGroupArrayOutput() ClusterParameterGroupArrayOutput {
	return i.ToClusterParameterGroupArrayOutputWithContext(context.Background())
}

func (i ClusterParameterGroupArray) ToClusterParameterGroupArrayOutputWithContext(ctx context.Context) ClusterParameterGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupArrayOutput)
}

// ClusterParameterGroupMapInput is an input type that accepts ClusterParameterGroupMap and ClusterParameterGroupMapOutput values.
// You can construct a concrete instance of `ClusterParameterGroupMapInput` via:
//
//          ClusterParameterGroupMap{ "key": ClusterParameterGroupArgs{...} }
type ClusterParameterGroupMapInput interface {
	pulumi.Input

	ToClusterParameterGroupMapOutput() ClusterParameterGroupMapOutput
	ToClusterParameterGroupMapOutputWithContext(context.Context) ClusterParameterGroupMapOutput
}

type ClusterParameterGroupMap map[string]ClusterParameterGroupInput

func (ClusterParameterGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ClusterParameterGroup)(nil))
}

func (i ClusterParameterGroupMap) ToClusterParameterGroupMapOutput() ClusterParameterGroupMapOutput {
	return i.ToClusterParameterGroupMapOutputWithContext(context.Background())
}

func (i ClusterParameterGroupMap) ToClusterParameterGroupMapOutputWithContext(ctx context.Context) ClusterParameterGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupMapOutput)
}

type ClusterParameterGroupOutput struct {
	*pulumi.OutputState
}

func (ClusterParameterGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterParameterGroup)(nil))
}

func (o ClusterParameterGroupOutput) ToClusterParameterGroupOutput() ClusterParameterGroupOutput {
	return o
}

func (o ClusterParameterGroupOutput) ToClusterParameterGroupOutputWithContext(ctx context.Context) ClusterParameterGroupOutput {
	return o
}

func (o ClusterParameterGroupOutput) ToClusterParameterGroupPtrOutput() ClusterParameterGroupPtrOutput {
	return o.ToClusterParameterGroupPtrOutputWithContext(context.Background())
}

func (o ClusterParameterGroupOutput) ToClusterParameterGroupPtrOutputWithContext(ctx context.Context) ClusterParameterGroupPtrOutput {
	return o.ApplyT(func(v ClusterParameterGroup) *ClusterParameterGroup {
		return &v
	}).(ClusterParameterGroupPtrOutput)
}

type ClusterParameterGroupPtrOutput struct {
	*pulumi.OutputState
}

func (ClusterParameterGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterParameterGroup)(nil))
}

func (o ClusterParameterGroupPtrOutput) ToClusterParameterGroupPtrOutput() ClusterParameterGroupPtrOutput {
	return o
}

func (o ClusterParameterGroupPtrOutput) ToClusterParameterGroupPtrOutputWithContext(ctx context.Context) ClusterParameterGroupPtrOutput {
	return o
}

type ClusterParameterGroupArrayOutput struct{ *pulumi.OutputState }

func (ClusterParameterGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterParameterGroup)(nil))
}

func (o ClusterParameterGroupArrayOutput) ToClusterParameterGroupArrayOutput() ClusterParameterGroupArrayOutput {
	return o
}

func (o ClusterParameterGroupArrayOutput) ToClusterParameterGroupArrayOutputWithContext(ctx context.Context) ClusterParameterGroupArrayOutput {
	return o
}

func (o ClusterParameterGroupArrayOutput) Index(i pulumi.IntInput) ClusterParameterGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterParameterGroup {
		return vs[0].([]ClusterParameterGroup)[vs[1].(int)]
	}).(ClusterParameterGroupOutput)
}

type ClusterParameterGroupMapOutput struct{ *pulumi.OutputState }

func (ClusterParameterGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ClusterParameterGroup)(nil))
}

func (o ClusterParameterGroupMapOutput) ToClusterParameterGroupMapOutput() ClusterParameterGroupMapOutput {
	return o
}

func (o ClusterParameterGroupMapOutput) ToClusterParameterGroupMapOutputWithContext(ctx context.Context) ClusterParameterGroupMapOutput {
	return o
}

func (o ClusterParameterGroupMapOutput) MapIndex(k pulumi.StringInput) ClusterParameterGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ClusterParameterGroup {
		return vs[0].(map[string]ClusterParameterGroup)[vs[1].(string)]
	}).(ClusterParameterGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterParameterGroupOutput{})
	pulumi.RegisterOutputType(ClusterParameterGroupPtrOutput{})
	pulumi.RegisterOutputType(ClusterParameterGroupArrayOutput{})
	pulumi.RegisterOutputType(ClusterParameterGroupMapOutput{})
}
