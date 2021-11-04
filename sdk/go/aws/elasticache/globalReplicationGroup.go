// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an ElastiCache Global Replication Group resource, which manages replication between two or more Replication Groups in different regions. For more information, see the [ElastiCache User Guide](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Redis-Global-Datastore.html).
//
// ## Example Usage
// ### Global replication group with one secondary replication group
//
// The global replication group depends on the primary group existing. Secondary replication groups depend on the global replication group. the provider dependency management will handle this transparently using resource value references.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/elasticache"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		primary, err := elasticache.NewReplicationGroup(ctx, "primary", &elasticache.ReplicationGroupArgs{
// 			ReplicationGroupDescription: pulumi.String("primary replication group"),
// 			Engine:                      pulumi.String("redis"),
// 			EngineVersion:               pulumi.String("5.0.6"),
// 			NodeType:                    pulumi.String("cache.m5.large"),
// 			NumberCacheClusters:         pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		example, err := elasticache.NewGlobalReplicationGroup(ctx, "example", &elasticache.GlobalReplicationGroupArgs{
// 			GlobalReplicationGroupIdSuffix: pulumi.String("example"),
// 			PrimaryReplicationGroupId:      primary.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elasticache.NewReplicationGroup(ctx, "secondary", &elasticache.ReplicationGroupArgs{
// 			ReplicationGroupDescription: pulumi.String("secondary replication group"),
// 			GlobalReplicationGroupId:    example.GlobalReplicationGroupId,
// 			NumberCacheClusters:         pulumi.Int(1),
// 		}, pulumi.Provider(aws.Other_region))
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
// ElastiCache Global Replication Groups can be imported using the `global_replication_group_id`, e.g.
//
// ```sh
//  $ pulumi import aws:elasticache/globalReplicationGroup:GlobalReplicationGroup my_global_replication_group okuqm-global-replication-group-1
// ```
type GlobalReplicationGroup struct {
	pulumi.CustomResourceState

	// (**DEPRECATED** use `engineVersionActual` instead) The full version number of the cache engine running on the members of this global replication group.
	//
	// Deprecated: Use engine_version_actual instead
	ActualEngineVersion pulumi.StringOutput `pulumi:"actualEngineVersion"`
	// The ARN of the ElastiCache Global Replication Group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A flag that indicate whether the encryption at rest is enabled.
	AtRestEncryptionEnabled pulumi.BoolOutput `pulumi:"atRestEncryptionEnabled"`
	// A flag that indicate whether AuthToken (password) is enabled.
	AuthTokenEnabled pulumi.BoolOutput `pulumi:"authTokenEnabled"`
	// The instance class used. See AWS documentation for information on [supported node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html) and [guidance on selecting node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/nodes-select-size.html).
	CacheNodeType pulumi.StringOutput `pulumi:"cacheNodeType"`
	// Indicates whether the Global Datastore is cluster enabled.
	ClusterEnabled pulumi.BoolOutput `pulumi:"clusterEnabled"`
	// The name of the cache engine to be used for the clusters in this global replication group.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// The full version number of the cache engine running on the members of this global replication group.
	EngineVersionActual pulumi.StringOutput `pulumi:"engineVersionActual"`
	// A user-created description for the global replication group.
	GlobalReplicationGroupDescription pulumi.StringPtrOutput `pulumi:"globalReplicationGroupDescription"`
	// The full ID of the global replication group.
	GlobalReplicationGroupId pulumi.StringOutput `pulumi:"globalReplicationGroupId"`
	// The suffix name of a Global Datastore. If `globalReplicationGroupIdSuffix` is changed, creates a new resource.
	GlobalReplicationGroupIdSuffix pulumi.StringOutput `pulumi:"globalReplicationGroupIdSuffix"`
	// The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If `primaryReplicationGroupId` is changed, creates a new resource.
	PrimaryReplicationGroupId pulumi.StringOutput `pulumi:"primaryReplicationGroupId"`
	// A flag that indicates whether the encryption in transit is enabled.
	TransitEncryptionEnabled pulumi.BoolOutput `pulumi:"transitEncryptionEnabled"`
}

// NewGlobalReplicationGroup registers a new resource with the given unique name, arguments, and options.
func NewGlobalReplicationGroup(ctx *pulumi.Context,
	name string, args *GlobalReplicationGroupArgs, opts ...pulumi.ResourceOption) (*GlobalReplicationGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GlobalReplicationGroupIdSuffix == nil {
		return nil, errors.New("invalid value for required argument 'GlobalReplicationGroupIdSuffix'")
	}
	if args.PrimaryReplicationGroupId == nil {
		return nil, errors.New("invalid value for required argument 'PrimaryReplicationGroupId'")
	}
	var resource GlobalReplicationGroup
	err := ctx.RegisterResource("aws:elasticache/globalReplicationGroup:GlobalReplicationGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalReplicationGroup gets an existing GlobalReplicationGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalReplicationGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalReplicationGroupState, opts ...pulumi.ResourceOption) (*GlobalReplicationGroup, error) {
	var resource GlobalReplicationGroup
	err := ctx.ReadResource("aws:elasticache/globalReplicationGroup:GlobalReplicationGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalReplicationGroup resources.
type globalReplicationGroupState struct {
	// (**DEPRECATED** use `engineVersionActual` instead) The full version number of the cache engine running on the members of this global replication group.
	//
	// Deprecated: Use engine_version_actual instead
	ActualEngineVersion *string `pulumi:"actualEngineVersion"`
	// The ARN of the ElastiCache Global Replication Group.
	Arn *string `pulumi:"arn"`
	// A flag that indicate whether the encryption at rest is enabled.
	AtRestEncryptionEnabled *bool `pulumi:"atRestEncryptionEnabled"`
	// A flag that indicate whether AuthToken (password) is enabled.
	AuthTokenEnabled *bool `pulumi:"authTokenEnabled"`
	// The instance class used. See AWS documentation for information on [supported node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html) and [guidance on selecting node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/nodes-select-size.html).
	CacheNodeType *string `pulumi:"cacheNodeType"`
	// Indicates whether the Global Datastore is cluster enabled.
	ClusterEnabled *bool `pulumi:"clusterEnabled"`
	// The name of the cache engine to be used for the clusters in this global replication group.
	Engine *string `pulumi:"engine"`
	// The full version number of the cache engine running on the members of this global replication group.
	EngineVersionActual *string `pulumi:"engineVersionActual"`
	// A user-created description for the global replication group.
	GlobalReplicationGroupDescription *string `pulumi:"globalReplicationGroupDescription"`
	// The full ID of the global replication group.
	GlobalReplicationGroupId *string `pulumi:"globalReplicationGroupId"`
	// The suffix name of a Global Datastore. If `globalReplicationGroupIdSuffix` is changed, creates a new resource.
	GlobalReplicationGroupIdSuffix *string `pulumi:"globalReplicationGroupIdSuffix"`
	// The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If `primaryReplicationGroupId` is changed, creates a new resource.
	PrimaryReplicationGroupId *string `pulumi:"primaryReplicationGroupId"`
	// A flag that indicates whether the encryption in transit is enabled.
	TransitEncryptionEnabled *bool `pulumi:"transitEncryptionEnabled"`
}

type GlobalReplicationGroupState struct {
	// (**DEPRECATED** use `engineVersionActual` instead) The full version number of the cache engine running on the members of this global replication group.
	//
	// Deprecated: Use engine_version_actual instead
	ActualEngineVersion pulumi.StringPtrInput
	// The ARN of the ElastiCache Global Replication Group.
	Arn pulumi.StringPtrInput
	// A flag that indicate whether the encryption at rest is enabled.
	AtRestEncryptionEnabled pulumi.BoolPtrInput
	// A flag that indicate whether AuthToken (password) is enabled.
	AuthTokenEnabled pulumi.BoolPtrInput
	// The instance class used. See AWS documentation for information on [supported node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html) and [guidance on selecting node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/nodes-select-size.html).
	CacheNodeType pulumi.StringPtrInput
	// Indicates whether the Global Datastore is cluster enabled.
	ClusterEnabled pulumi.BoolPtrInput
	// The name of the cache engine to be used for the clusters in this global replication group.
	Engine pulumi.StringPtrInput
	// The full version number of the cache engine running on the members of this global replication group.
	EngineVersionActual pulumi.StringPtrInput
	// A user-created description for the global replication group.
	GlobalReplicationGroupDescription pulumi.StringPtrInput
	// The full ID of the global replication group.
	GlobalReplicationGroupId pulumi.StringPtrInput
	// The suffix name of a Global Datastore. If `globalReplicationGroupIdSuffix` is changed, creates a new resource.
	GlobalReplicationGroupIdSuffix pulumi.StringPtrInput
	// The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If `primaryReplicationGroupId` is changed, creates a new resource.
	PrimaryReplicationGroupId pulumi.StringPtrInput
	// A flag that indicates whether the encryption in transit is enabled.
	TransitEncryptionEnabled pulumi.BoolPtrInput
}

func (GlobalReplicationGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalReplicationGroupState)(nil)).Elem()
}

type globalReplicationGroupArgs struct {
	// A user-created description for the global replication group.
	GlobalReplicationGroupDescription *string `pulumi:"globalReplicationGroupDescription"`
	// The suffix name of a Global Datastore. If `globalReplicationGroupIdSuffix` is changed, creates a new resource.
	GlobalReplicationGroupIdSuffix string `pulumi:"globalReplicationGroupIdSuffix"`
	// The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If `primaryReplicationGroupId` is changed, creates a new resource.
	PrimaryReplicationGroupId string `pulumi:"primaryReplicationGroupId"`
}

// The set of arguments for constructing a GlobalReplicationGroup resource.
type GlobalReplicationGroupArgs struct {
	// A user-created description for the global replication group.
	GlobalReplicationGroupDescription pulumi.StringPtrInput
	// The suffix name of a Global Datastore. If `globalReplicationGroupIdSuffix` is changed, creates a new resource.
	GlobalReplicationGroupIdSuffix pulumi.StringInput
	// The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If `primaryReplicationGroupId` is changed, creates a new resource.
	PrimaryReplicationGroupId pulumi.StringInput
}

func (GlobalReplicationGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalReplicationGroupArgs)(nil)).Elem()
}

type GlobalReplicationGroupInput interface {
	pulumi.Input

	ToGlobalReplicationGroupOutput() GlobalReplicationGroupOutput
	ToGlobalReplicationGroupOutputWithContext(ctx context.Context) GlobalReplicationGroupOutput
}

func (*GlobalReplicationGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalReplicationGroup)(nil))
}

func (i *GlobalReplicationGroup) ToGlobalReplicationGroupOutput() GlobalReplicationGroupOutput {
	return i.ToGlobalReplicationGroupOutputWithContext(context.Background())
}

func (i *GlobalReplicationGroup) ToGlobalReplicationGroupOutputWithContext(ctx context.Context) GlobalReplicationGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalReplicationGroupOutput)
}

func (i *GlobalReplicationGroup) ToGlobalReplicationGroupPtrOutput() GlobalReplicationGroupPtrOutput {
	return i.ToGlobalReplicationGroupPtrOutputWithContext(context.Background())
}

func (i *GlobalReplicationGroup) ToGlobalReplicationGroupPtrOutputWithContext(ctx context.Context) GlobalReplicationGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalReplicationGroupPtrOutput)
}

type GlobalReplicationGroupPtrInput interface {
	pulumi.Input

	ToGlobalReplicationGroupPtrOutput() GlobalReplicationGroupPtrOutput
	ToGlobalReplicationGroupPtrOutputWithContext(ctx context.Context) GlobalReplicationGroupPtrOutput
}

type globalReplicationGroupPtrType GlobalReplicationGroupArgs

func (*globalReplicationGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalReplicationGroup)(nil))
}

func (i *globalReplicationGroupPtrType) ToGlobalReplicationGroupPtrOutput() GlobalReplicationGroupPtrOutput {
	return i.ToGlobalReplicationGroupPtrOutputWithContext(context.Background())
}

func (i *globalReplicationGroupPtrType) ToGlobalReplicationGroupPtrOutputWithContext(ctx context.Context) GlobalReplicationGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalReplicationGroupPtrOutput)
}

// GlobalReplicationGroupArrayInput is an input type that accepts GlobalReplicationGroupArray and GlobalReplicationGroupArrayOutput values.
// You can construct a concrete instance of `GlobalReplicationGroupArrayInput` via:
//
//          GlobalReplicationGroupArray{ GlobalReplicationGroupArgs{...} }
type GlobalReplicationGroupArrayInput interface {
	pulumi.Input

	ToGlobalReplicationGroupArrayOutput() GlobalReplicationGroupArrayOutput
	ToGlobalReplicationGroupArrayOutputWithContext(context.Context) GlobalReplicationGroupArrayOutput
}

type GlobalReplicationGroupArray []GlobalReplicationGroupInput

func (GlobalReplicationGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GlobalReplicationGroup)(nil)).Elem()
}

func (i GlobalReplicationGroupArray) ToGlobalReplicationGroupArrayOutput() GlobalReplicationGroupArrayOutput {
	return i.ToGlobalReplicationGroupArrayOutputWithContext(context.Background())
}

func (i GlobalReplicationGroupArray) ToGlobalReplicationGroupArrayOutputWithContext(ctx context.Context) GlobalReplicationGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalReplicationGroupArrayOutput)
}

// GlobalReplicationGroupMapInput is an input type that accepts GlobalReplicationGroupMap and GlobalReplicationGroupMapOutput values.
// You can construct a concrete instance of `GlobalReplicationGroupMapInput` via:
//
//          GlobalReplicationGroupMap{ "key": GlobalReplicationGroupArgs{...} }
type GlobalReplicationGroupMapInput interface {
	pulumi.Input

	ToGlobalReplicationGroupMapOutput() GlobalReplicationGroupMapOutput
	ToGlobalReplicationGroupMapOutputWithContext(context.Context) GlobalReplicationGroupMapOutput
}

type GlobalReplicationGroupMap map[string]GlobalReplicationGroupInput

func (GlobalReplicationGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GlobalReplicationGroup)(nil)).Elem()
}

func (i GlobalReplicationGroupMap) ToGlobalReplicationGroupMapOutput() GlobalReplicationGroupMapOutput {
	return i.ToGlobalReplicationGroupMapOutputWithContext(context.Background())
}

func (i GlobalReplicationGroupMap) ToGlobalReplicationGroupMapOutputWithContext(ctx context.Context) GlobalReplicationGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalReplicationGroupMapOutput)
}

type GlobalReplicationGroupOutput struct{ *pulumi.OutputState }

func (GlobalReplicationGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalReplicationGroup)(nil))
}

func (o GlobalReplicationGroupOutput) ToGlobalReplicationGroupOutput() GlobalReplicationGroupOutput {
	return o
}

func (o GlobalReplicationGroupOutput) ToGlobalReplicationGroupOutputWithContext(ctx context.Context) GlobalReplicationGroupOutput {
	return o
}

func (o GlobalReplicationGroupOutput) ToGlobalReplicationGroupPtrOutput() GlobalReplicationGroupPtrOutput {
	return o.ToGlobalReplicationGroupPtrOutputWithContext(context.Background())
}

func (o GlobalReplicationGroupOutput) ToGlobalReplicationGroupPtrOutputWithContext(ctx context.Context) GlobalReplicationGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GlobalReplicationGroup) *GlobalReplicationGroup {
		return &v
	}).(GlobalReplicationGroupPtrOutput)
}

type GlobalReplicationGroupPtrOutput struct{ *pulumi.OutputState }

func (GlobalReplicationGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalReplicationGroup)(nil))
}

func (o GlobalReplicationGroupPtrOutput) ToGlobalReplicationGroupPtrOutput() GlobalReplicationGroupPtrOutput {
	return o
}

func (o GlobalReplicationGroupPtrOutput) ToGlobalReplicationGroupPtrOutputWithContext(ctx context.Context) GlobalReplicationGroupPtrOutput {
	return o
}

func (o GlobalReplicationGroupPtrOutput) Elem() GlobalReplicationGroupOutput {
	return o.ApplyT(func(v *GlobalReplicationGroup) GlobalReplicationGroup {
		if v != nil {
			return *v
		}
		var ret GlobalReplicationGroup
		return ret
	}).(GlobalReplicationGroupOutput)
}

type GlobalReplicationGroupArrayOutput struct{ *pulumi.OutputState }

func (GlobalReplicationGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GlobalReplicationGroup)(nil))
}

func (o GlobalReplicationGroupArrayOutput) ToGlobalReplicationGroupArrayOutput() GlobalReplicationGroupArrayOutput {
	return o
}

func (o GlobalReplicationGroupArrayOutput) ToGlobalReplicationGroupArrayOutputWithContext(ctx context.Context) GlobalReplicationGroupArrayOutput {
	return o
}

func (o GlobalReplicationGroupArrayOutput) Index(i pulumi.IntInput) GlobalReplicationGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GlobalReplicationGroup {
		return vs[0].([]GlobalReplicationGroup)[vs[1].(int)]
	}).(GlobalReplicationGroupOutput)
}

type GlobalReplicationGroupMapOutput struct{ *pulumi.OutputState }

func (GlobalReplicationGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GlobalReplicationGroup)(nil))
}

func (o GlobalReplicationGroupMapOutput) ToGlobalReplicationGroupMapOutput() GlobalReplicationGroupMapOutput {
	return o
}

func (o GlobalReplicationGroupMapOutput) ToGlobalReplicationGroupMapOutputWithContext(ctx context.Context) GlobalReplicationGroupMapOutput {
	return o
}

func (o GlobalReplicationGroupMapOutput) MapIndex(k pulumi.StringInput) GlobalReplicationGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GlobalReplicationGroup {
		return vs[0].(map[string]GlobalReplicationGroup)[vs[1].(string)]
	}).(GlobalReplicationGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalReplicationGroupInput)(nil)).Elem(), &GlobalReplicationGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalReplicationGroupPtrInput)(nil)).Elem(), &GlobalReplicationGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalReplicationGroupArrayInput)(nil)).Elem(), GlobalReplicationGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalReplicationGroupMapInput)(nil)).Elem(), GlobalReplicationGroupMap{})
	pulumi.RegisterOutputType(GlobalReplicationGroupOutput{})
	pulumi.RegisterOutputType(GlobalReplicationGroupPtrOutput{})
	pulumi.RegisterOutputType(GlobalReplicationGroupArrayOutput{})
	pulumi.RegisterOutputType(GlobalReplicationGroupMapOutput{})
}
