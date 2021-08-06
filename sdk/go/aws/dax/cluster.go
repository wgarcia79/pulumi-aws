// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dax

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DAX Cluster resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/dax"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dax.NewCluster(ctx, "bar", &dax.ClusterArgs{
// 			ClusterName:       pulumi.String("cluster-example"),
// 			IamRoleArn:        pulumi.Any(data.Aws_iam_role.Example.Arn),
// 			NodeType:          pulumi.String("dax.r4.large"),
// 			ReplicationFactor: pulumi.Int(1),
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
// DAX Clusters can be imported using the `cluster_name`, e.g.
//
// ```sh
//  $ pulumi import aws:dax/cluster:Cluster my_cluster my_cluster
// ```
//
//  [1]http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.concepts.cluster.html#DAX.concepts.nodes
type Cluster struct {
	pulumi.CustomResourceState

	// The ARN of the DAX cluster
	Arn pulumi.StringOutput `pulumi:"arn"`
	// List of Availability Zones in which the
	// nodes will be created
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// The DNS name of the DAX cluster without the port appended
	ClusterAddress pulumi.StringOutput `pulumi:"clusterAddress"`
	// Group identifier. DAX converts this name to
	// lowercase
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// The configuration endpoint for this DAX cluster,
	// consisting of a DNS name and a port number
	ConfigurationEndpoint pulumi.StringOutput `pulumi:"configurationEndpoint"`
	// Description for the cluster
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A valid Amazon Resource Name (ARN) that identifies
	// an IAM role. At runtime, DAX will assume this role and use the role's
	// permissions to access DynamoDB on your behalf
	IamRoleArn pulumi.StringOutput `pulumi:"iamRoleArn"`
	// Specifies the weekly time range for when
	// maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
	// `sun:05:00-sun:09:00`
	MaintenanceWindow pulumi.StringOutput `pulumi:"maintenanceWindow"`
	// The compute and memory capacity of the nodes. See
	// [Nodes](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.concepts.cluster.html#DAX.concepts.nodes) for supported node types
	NodeType pulumi.StringOutput `pulumi:"nodeType"`
	// List of node objects including `id`, `address`, `port` and
	// `availabilityZone`. Referenceable e.g. as
	// `${aws_dax_cluster.test.nodes.0.address}`
	Nodes ClusterNodeArrayOutput `pulumi:"nodes"`
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send DAX notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn pulumi.StringPtrOutput `pulumi:"notificationTopicArn"`
	// Name of the parameter group to associate
	// with this DAX cluster
	ParameterGroupName pulumi.StringOutput `pulumi:"parameterGroupName"`
	// The port used by the configuration endpoint
	Port pulumi.IntOutput `pulumi:"port"`
	// The number of nodes in the DAX cluster. A
	// replication factor of 1 will create a single-node cluster, without any read
	// replicas
	ReplicationFactor pulumi.IntOutput `pulumi:"replicationFactor"`
	// One or more VPC security groups associated
	// with the cluster
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// Encrypt at rest options
	ServerSideEncryption ClusterServerSideEncryptionPtrOutput `pulumi:"serverSideEncryption"`
	// Name of the subnet group to be used for the
	// cluster
	SubnetGroupName pulumi.StringOutput `pulumi:"subnetGroupName"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.IamRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'IamRoleArn'")
	}
	if args.NodeType == nil {
		return nil, errors.New("invalid value for required argument 'NodeType'")
	}
	if args.ReplicationFactor == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationFactor'")
	}
	var resource Cluster
	err := ctx.RegisterResource("aws:dax/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("aws:dax/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The ARN of the DAX cluster
	Arn *string `pulumi:"arn"`
	// List of Availability Zones in which the
	// nodes will be created
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The DNS name of the DAX cluster without the port appended
	ClusterAddress *string `pulumi:"clusterAddress"`
	// Group identifier. DAX converts this name to
	// lowercase
	ClusterName *string `pulumi:"clusterName"`
	// The configuration endpoint for this DAX cluster,
	// consisting of a DNS name and a port number
	ConfigurationEndpoint *string `pulumi:"configurationEndpoint"`
	// Description for the cluster
	Description *string `pulumi:"description"`
	// A valid Amazon Resource Name (ARN) that identifies
	// an IAM role. At runtime, DAX will assume this role and use the role's
	// permissions to access DynamoDB on your behalf
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// Specifies the weekly time range for when
	// maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
	// `sun:05:00-sun:09:00`
	MaintenanceWindow *string `pulumi:"maintenanceWindow"`
	// The compute and memory capacity of the nodes. See
	// [Nodes](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.concepts.cluster.html#DAX.concepts.nodes) for supported node types
	NodeType *string `pulumi:"nodeType"`
	// List of node objects including `id`, `address`, `port` and
	// `availabilityZone`. Referenceable e.g. as
	// `${aws_dax_cluster.test.nodes.0.address}`
	Nodes []ClusterNode `pulumi:"nodes"`
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send DAX notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn *string `pulumi:"notificationTopicArn"`
	// Name of the parameter group to associate
	// with this DAX cluster
	ParameterGroupName *string `pulumi:"parameterGroupName"`
	// The port used by the configuration endpoint
	Port *int `pulumi:"port"`
	// The number of nodes in the DAX cluster. A
	// replication factor of 1 will create a single-node cluster, without any read
	// replicas
	ReplicationFactor *int `pulumi:"replicationFactor"`
	// One or more VPC security groups associated
	// with the cluster
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Encrypt at rest options
	ServerSideEncryption *ClusterServerSideEncryption `pulumi:"serverSideEncryption"`
	// Name of the subnet group to be used for the
	// cluster
	SubnetGroupName *string `pulumi:"subnetGroupName"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ClusterState struct {
	// The ARN of the DAX cluster
	Arn pulumi.StringPtrInput
	// List of Availability Zones in which the
	// nodes will be created
	AvailabilityZones pulumi.StringArrayInput
	// The DNS name of the DAX cluster without the port appended
	ClusterAddress pulumi.StringPtrInput
	// Group identifier. DAX converts this name to
	// lowercase
	ClusterName pulumi.StringPtrInput
	// The configuration endpoint for this DAX cluster,
	// consisting of a DNS name and a port number
	ConfigurationEndpoint pulumi.StringPtrInput
	// Description for the cluster
	Description pulumi.StringPtrInput
	// A valid Amazon Resource Name (ARN) that identifies
	// an IAM role. At runtime, DAX will assume this role and use the role's
	// permissions to access DynamoDB on your behalf
	IamRoleArn pulumi.StringPtrInput
	// Specifies the weekly time range for when
	// maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
	// `sun:05:00-sun:09:00`
	MaintenanceWindow pulumi.StringPtrInput
	// The compute and memory capacity of the nodes. See
	// [Nodes](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.concepts.cluster.html#DAX.concepts.nodes) for supported node types
	NodeType pulumi.StringPtrInput
	// List of node objects including `id`, `address`, `port` and
	// `availabilityZone`. Referenceable e.g. as
	// `${aws_dax_cluster.test.nodes.0.address}`
	Nodes ClusterNodeArrayInput
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send DAX notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn pulumi.StringPtrInput
	// Name of the parameter group to associate
	// with this DAX cluster
	ParameterGroupName pulumi.StringPtrInput
	// The port used by the configuration endpoint
	Port pulumi.IntPtrInput
	// The number of nodes in the DAX cluster. A
	// replication factor of 1 will create a single-node cluster, without any read
	// replicas
	ReplicationFactor pulumi.IntPtrInput
	// One or more VPC security groups associated
	// with the cluster
	SecurityGroupIds pulumi.StringArrayInput
	// Encrypt at rest options
	ServerSideEncryption ClusterServerSideEncryptionPtrInput
	// Name of the subnet group to be used for the
	// cluster
	SubnetGroupName pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// List of Availability Zones in which the
	// nodes will be created
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Group identifier. DAX converts this name to
	// lowercase
	ClusterName string `pulumi:"clusterName"`
	// Description for the cluster
	Description *string `pulumi:"description"`
	// A valid Amazon Resource Name (ARN) that identifies
	// an IAM role. At runtime, DAX will assume this role and use the role's
	// permissions to access DynamoDB on your behalf
	IamRoleArn string `pulumi:"iamRoleArn"`
	// Specifies the weekly time range for when
	// maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
	// `sun:05:00-sun:09:00`
	MaintenanceWindow *string `pulumi:"maintenanceWindow"`
	// The compute and memory capacity of the nodes. See
	// [Nodes](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.concepts.cluster.html#DAX.concepts.nodes) for supported node types
	NodeType string `pulumi:"nodeType"`
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send DAX notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn *string `pulumi:"notificationTopicArn"`
	// Name of the parameter group to associate
	// with this DAX cluster
	ParameterGroupName *string `pulumi:"parameterGroupName"`
	// The number of nodes in the DAX cluster. A
	// replication factor of 1 will create a single-node cluster, without any read
	// replicas
	ReplicationFactor int `pulumi:"replicationFactor"`
	// One or more VPC security groups associated
	// with the cluster
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Encrypt at rest options
	ServerSideEncryption *ClusterServerSideEncryption `pulumi:"serverSideEncryption"`
	// Name of the subnet group to be used for the
	// cluster
	SubnetGroupName *string `pulumi:"subnetGroupName"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// List of Availability Zones in which the
	// nodes will be created
	AvailabilityZones pulumi.StringArrayInput
	// Group identifier. DAX converts this name to
	// lowercase
	ClusterName pulumi.StringInput
	// Description for the cluster
	Description pulumi.StringPtrInput
	// A valid Amazon Resource Name (ARN) that identifies
	// an IAM role. At runtime, DAX will assume this role and use the role's
	// permissions to access DynamoDB on your behalf
	IamRoleArn pulumi.StringInput
	// Specifies the weekly time range for when
	// maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
	// `sun:05:00-sun:09:00`
	MaintenanceWindow pulumi.StringPtrInput
	// The compute and memory capacity of the nodes. See
	// [Nodes](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.concepts.cluster.html#DAX.concepts.nodes) for supported node types
	NodeType pulumi.StringInput
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send DAX notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn pulumi.StringPtrInput
	// Name of the parameter group to associate
	// with this DAX cluster
	ParameterGroupName pulumi.StringPtrInput
	// The number of nodes in the DAX cluster. A
	// replication factor of 1 will create a single-node cluster, without any read
	// replicas
	ReplicationFactor pulumi.IntInput
	// One or more VPC security groups associated
	// with the cluster
	SecurityGroupIds pulumi.StringArrayInput
	// Encrypt at rest options
	ServerSideEncryption ClusterServerSideEncryptionPtrInput
	// Name of the subnet group to be used for the
	// cluster
	SubnetGroupName pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil))
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

func (i *Cluster) ToClusterPtrOutput() ClusterPtrOutput {
	return i.ToClusterPtrOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPtrOutput)
}

type ClusterPtrInput interface {
	pulumi.Input

	ToClusterPtrOutput() ClusterPtrOutput
	ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput
}

type clusterPtrType ClusterArgs

func (*clusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil))
}

func (i *clusterPtrType) ToClusterPtrOutput() ClusterPtrOutput {
	return i.ToClusterPtrOutputWithContext(context.Background())
}

func (i *clusterPtrType) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPtrOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//          ClusterArray{ ClusterArgs{...} }
type ClusterArrayInput interface {
	pulumi.Input

	ToClusterArrayOutput() ClusterArrayOutput
	ToClusterArrayOutputWithContext(context.Context) ClusterArrayOutput
}

type ClusterArray []ClusterInput

func (ClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (i ClusterArray) ToClusterArrayOutput() ClusterArrayOutput {
	return i.ToClusterArrayOutputWithContext(context.Background())
}

func (i ClusterArray) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterArrayOutput)
}

// ClusterMapInput is an input type that accepts ClusterMap and ClusterMapOutput values.
// You can construct a concrete instance of `ClusterMapInput` via:
//
//          ClusterMap{ "key": ClusterArgs{...} }
type ClusterMapInput interface {
	pulumi.Input

	ToClusterMapOutput() ClusterMapOutput
	ToClusterMapOutputWithContext(context.Context) ClusterMapOutput
}

type ClusterMap map[string]ClusterInput

func (ClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (i ClusterMap) ToClusterMapOutput() ClusterMapOutput {
	return i.ToClusterMapOutputWithContext(context.Background())
}

func (i ClusterMap) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMapOutput)
}

type ClusterOutput struct {
	*pulumi.OutputState
}

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil))
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterPtrOutput() ClusterPtrOutput {
	return o.ToClusterPtrOutputWithContext(context.Background())
}

func (o ClusterOutput) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return o.ApplyT(func(v Cluster) *Cluster {
		return &v
	}).(ClusterPtrOutput)
}

type ClusterPtrOutput struct {
	*pulumi.OutputState
}

func (ClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil))
}

func (o ClusterPtrOutput) ToClusterPtrOutput() ClusterPtrOutput {
	return o
}

func (o ClusterPtrOutput) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return o
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Cluster)(nil))
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Cluster {
		return vs[0].([]Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Cluster)(nil))
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Cluster {
		return vs[0].(map[string]Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterPtrOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}
