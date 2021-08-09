// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DMS (Data Migration Service) replication subnet group resource. DMS replication subnet groups can be created, updated, deleted, and imported.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/dms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dms.NewReplicationSubnetGroup(ctx, "test", &dms.ReplicationSubnetGroupArgs{
// 			ReplicationSubnetGroupDescription: pulumi.String("Test replication subnet group"),
// 			ReplicationSubnetGroupId:          pulumi.String("test-dms-replication-subnet-group-tf"),
// 			SubnetIds: pulumi.StringArray{
// 				pulumi.String("subnet-12345678"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("test"),
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
// Replication subnet groups can be imported using the `replication_subnet_group_id`, e.g.
//
// ```sh
//  $ pulumi import aws:dms/replicationSubnetGroup:ReplicationSubnetGroup test test-dms-replication-subnet-group-tf
// ```
type ReplicationSubnetGroup struct {
	pulumi.CustomResourceState

	ReplicationSubnetGroupArn pulumi.StringOutput `pulumi:"replicationSubnetGroupArn"`
	// The description for the subnet group.
	ReplicationSubnetGroupDescription pulumi.StringOutput `pulumi:"replicationSubnetGroupDescription"`
	// The name for the replication subnet group. This value is stored as a lowercase string.
	ReplicationSubnetGroupId pulumi.StringOutput `pulumi:"replicationSubnetGroupId"`
	// A list of the EC2 subnet IDs for the subnet group.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The ID of the VPC the subnet group is in.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewReplicationSubnetGroup registers a new resource with the given unique name, arguments, and options.
func NewReplicationSubnetGroup(ctx *pulumi.Context,
	name string, args *ReplicationSubnetGroupArgs, opts ...pulumi.ResourceOption) (*ReplicationSubnetGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ReplicationSubnetGroupDescription == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationSubnetGroupDescription'")
	}
	if args.ReplicationSubnetGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationSubnetGroupId'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	var resource ReplicationSubnetGroup
	err := ctx.RegisterResource("aws:dms/replicationSubnetGroup:ReplicationSubnetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationSubnetGroup gets an existing ReplicationSubnetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationSubnetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationSubnetGroupState, opts ...pulumi.ResourceOption) (*ReplicationSubnetGroup, error) {
	var resource ReplicationSubnetGroup
	err := ctx.ReadResource("aws:dms/replicationSubnetGroup:ReplicationSubnetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationSubnetGroup resources.
type replicationSubnetGroupState struct {
	ReplicationSubnetGroupArn *string `pulumi:"replicationSubnetGroupArn"`
	// The description for the subnet group.
	ReplicationSubnetGroupDescription *string `pulumi:"replicationSubnetGroupDescription"`
	// The name for the replication subnet group. This value is stored as a lowercase string.
	ReplicationSubnetGroupId *string `pulumi:"replicationSubnetGroupId"`
	// A list of the EC2 subnet IDs for the subnet group.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The ID of the VPC the subnet group is in.
	VpcId *string `pulumi:"vpcId"`
}

type ReplicationSubnetGroupState struct {
	ReplicationSubnetGroupArn pulumi.StringPtrInput
	// The description for the subnet group.
	ReplicationSubnetGroupDescription pulumi.StringPtrInput
	// The name for the replication subnet group. This value is stored as a lowercase string.
	ReplicationSubnetGroupId pulumi.StringPtrInput
	// A list of the EC2 subnet IDs for the subnet group.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The ID of the VPC the subnet group is in.
	VpcId pulumi.StringPtrInput
}

func (ReplicationSubnetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationSubnetGroupState)(nil)).Elem()
}

type replicationSubnetGroupArgs struct {
	// The description for the subnet group.
	ReplicationSubnetGroupDescription string `pulumi:"replicationSubnetGroupDescription"`
	// The name for the replication subnet group. This value is stored as a lowercase string.
	ReplicationSubnetGroupId string `pulumi:"replicationSubnetGroupId"`
	// A list of the EC2 subnet IDs for the subnet group.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a ReplicationSubnetGroup resource.
type ReplicationSubnetGroupArgs struct {
	// The description for the subnet group.
	ReplicationSubnetGroupDescription pulumi.StringInput
	// The name for the replication subnet group. This value is stored as a lowercase string.
	ReplicationSubnetGroupId pulumi.StringInput
	// A list of the EC2 subnet IDs for the subnet group.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (ReplicationSubnetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationSubnetGroupArgs)(nil)).Elem()
}

type ReplicationSubnetGroupInput interface {
	pulumi.Input

	ToReplicationSubnetGroupOutput() ReplicationSubnetGroupOutput
	ToReplicationSubnetGroupOutputWithContext(ctx context.Context) ReplicationSubnetGroupOutput
}

func (*ReplicationSubnetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationSubnetGroup)(nil))
}

func (i *ReplicationSubnetGroup) ToReplicationSubnetGroupOutput() ReplicationSubnetGroupOutput {
	return i.ToReplicationSubnetGroupOutputWithContext(context.Background())
}

func (i *ReplicationSubnetGroup) ToReplicationSubnetGroupOutputWithContext(ctx context.Context) ReplicationSubnetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationSubnetGroupOutput)
}

func (i *ReplicationSubnetGroup) ToReplicationSubnetGroupPtrOutput() ReplicationSubnetGroupPtrOutput {
	return i.ToReplicationSubnetGroupPtrOutputWithContext(context.Background())
}

func (i *ReplicationSubnetGroup) ToReplicationSubnetGroupPtrOutputWithContext(ctx context.Context) ReplicationSubnetGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationSubnetGroupPtrOutput)
}

type ReplicationSubnetGroupPtrInput interface {
	pulumi.Input

	ToReplicationSubnetGroupPtrOutput() ReplicationSubnetGroupPtrOutput
	ToReplicationSubnetGroupPtrOutputWithContext(ctx context.Context) ReplicationSubnetGroupPtrOutput
}

type replicationSubnetGroupPtrType ReplicationSubnetGroupArgs

func (*replicationSubnetGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationSubnetGroup)(nil))
}

func (i *replicationSubnetGroupPtrType) ToReplicationSubnetGroupPtrOutput() ReplicationSubnetGroupPtrOutput {
	return i.ToReplicationSubnetGroupPtrOutputWithContext(context.Background())
}

func (i *replicationSubnetGroupPtrType) ToReplicationSubnetGroupPtrOutputWithContext(ctx context.Context) ReplicationSubnetGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationSubnetGroupPtrOutput)
}

// ReplicationSubnetGroupArrayInput is an input type that accepts ReplicationSubnetGroupArray and ReplicationSubnetGroupArrayOutput values.
// You can construct a concrete instance of `ReplicationSubnetGroupArrayInput` via:
//
//          ReplicationSubnetGroupArray{ ReplicationSubnetGroupArgs{...} }
type ReplicationSubnetGroupArrayInput interface {
	pulumi.Input

	ToReplicationSubnetGroupArrayOutput() ReplicationSubnetGroupArrayOutput
	ToReplicationSubnetGroupArrayOutputWithContext(context.Context) ReplicationSubnetGroupArrayOutput
}

type ReplicationSubnetGroupArray []ReplicationSubnetGroupInput

func (ReplicationSubnetGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationSubnetGroup)(nil)).Elem()
}

func (i ReplicationSubnetGroupArray) ToReplicationSubnetGroupArrayOutput() ReplicationSubnetGroupArrayOutput {
	return i.ToReplicationSubnetGroupArrayOutputWithContext(context.Background())
}

func (i ReplicationSubnetGroupArray) ToReplicationSubnetGroupArrayOutputWithContext(ctx context.Context) ReplicationSubnetGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationSubnetGroupArrayOutput)
}

// ReplicationSubnetGroupMapInput is an input type that accepts ReplicationSubnetGroupMap and ReplicationSubnetGroupMapOutput values.
// You can construct a concrete instance of `ReplicationSubnetGroupMapInput` via:
//
//          ReplicationSubnetGroupMap{ "key": ReplicationSubnetGroupArgs{...} }
type ReplicationSubnetGroupMapInput interface {
	pulumi.Input

	ToReplicationSubnetGroupMapOutput() ReplicationSubnetGroupMapOutput
	ToReplicationSubnetGroupMapOutputWithContext(context.Context) ReplicationSubnetGroupMapOutput
}

type ReplicationSubnetGroupMap map[string]ReplicationSubnetGroupInput

func (ReplicationSubnetGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationSubnetGroup)(nil)).Elem()
}

func (i ReplicationSubnetGroupMap) ToReplicationSubnetGroupMapOutput() ReplicationSubnetGroupMapOutput {
	return i.ToReplicationSubnetGroupMapOutputWithContext(context.Background())
}

func (i ReplicationSubnetGroupMap) ToReplicationSubnetGroupMapOutputWithContext(ctx context.Context) ReplicationSubnetGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationSubnetGroupMapOutput)
}

type ReplicationSubnetGroupOutput struct{ *pulumi.OutputState }

func (ReplicationSubnetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationSubnetGroup)(nil))
}

func (o ReplicationSubnetGroupOutput) ToReplicationSubnetGroupOutput() ReplicationSubnetGroupOutput {
	return o
}

func (o ReplicationSubnetGroupOutput) ToReplicationSubnetGroupOutputWithContext(ctx context.Context) ReplicationSubnetGroupOutput {
	return o
}

func (o ReplicationSubnetGroupOutput) ToReplicationSubnetGroupPtrOutput() ReplicationSubnetGroupPtrOutput {
	return o.ToReplicationSubnetGroupPtrOutputWithContext(context.Background())
}

func (o ReplicationSubnetGroupOutput) ToReplicationSubnetGroupPtrOutputWithContext(ctx context.Context) ReplicationSubnetGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReplicationSubnetGroup) *ReplicationSubnetGroup {
		return &v
	}).(ReplicationSubnetGroupPtrOutput)
}

type ReplicationSubnetGroupPtrOutput struct{ *pulumi.OutputState }

func (ReplicationSubnetGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationSubnetGroup)(nil))
}

func (o ReplicationSubnetGroupPtrOutput) ToReplicationSubnetGroupPtrOutput() ReplicationSubnetGroupPtrOutput {
	return o
}

func (o ReplicationSubnetGroupPtrOutput) ToReplicationSubnetGroupPtrOutputWithContext(ctx context.Context) ReplicationSubnetGroupPtrOutput {
	return o
}

func (o ReplicationSubnetGroupPtrOutput) Elem() ReplicationSubnetGroupOutput {
	return o.ApplyT(func(v *ReplicationSubnetGroup) ReplicationSubnetGroup {
		if v != nil {
			return *v
		}
		var ret ReplicationSubnetGroup
		return ret
	}).(ReplicationSubnetGroupOutput)
}

type ReplicationSubnetGroupArrayOutput struct{ *pulumi.OutputState }

func (ReplicationSubnetGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReplicationSubnetGroup)(nil))
}

func (o ReplicationSubnetGroupArrayOutput) ToReplicationSubnetGroupArrayOutput() ReplicationSubnetGroupArrayOutput {
	return o
}

func (o ReplicationSubnetGroupArrayOutput) ToReplicationSubnetGroupArrayOutputWithContext(ctx context.Context) ReplicationSubnetGroupArrayOutput {
	return o
}

func (o ReplicationSubnetGroupArrayOutput) Index(i pulumi.IntInput) ReplicationSubnetGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReplicationSubnetGroup {
		return vs[0].([]ReplicationSubnetGroup)[vs[1].(int)]
	}).(ReplicationSubnetGroupOutput)
}

type ReplicationSubnetGroupMapOutput struct{ *pulumi.OutputState }

func (ReplicationSubnetGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReplicationSubnetGroup)(nil))
}

func (o ReplicationSubnetGroupMapOutput) ToReplicationSubnetGroupMapOutput() ReplicationSubnetGroupMapOutput {
	return o
}

func (o ReplicationSubnetGroupMapOutput) ToReplicationSubnetGroupMapOutputWithContext(ctx context.Context) ReplicationSubnetGroupMapOutput {
	return o
}

func (o ReplicationSubnetGroupMapOutput) MapIndex(k pulumi.StringInput) ReplicationSubnetGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReplicationSubnetGroup {
		return vs[0].(map[string]ReplicationSubnetGroup)[vs[1].(string)]
	}).(ReplicationSubnetGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicationSubnetGroupOutput{})
	pulumi.RegisterOutputType(ReplicationSubnetGroupPtrOutput{})
	pulumi.RegisterOutputType(ReplicationSubnetGroupArrayOutput{})
	pulumi.RegisterOutputType(ReplicationSubnetGroupMapOutput{})
}
