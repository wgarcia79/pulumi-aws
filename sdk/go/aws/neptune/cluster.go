// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Neptune Cluster Resource. A Cluster Resource defines attributes that are
// applied to the entire cluster of Neptune Cluster Instances.
//
// Changes to a Neptune Cluster can occur when you manually change a
// parameter, such as `backupRetentionPeriod`, and are reflected in the next maintenance
// window. Because of this, this provider may report a difference in its planning
// phase because a modification has not yet taken place. You can use the
// `applyImmediately` flag to instruct the service to apply the change immediately
// (see documentation below).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/neptune"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := neptune.NewCluster(ctx, "_default", &neptune.ClusterArgs{
// 			ApplyImmediately:                 pulumi.Bool(true),
// 			BackupRetentionPeriod:            pulumi.Int(5),
// 			ClusterIdentifier:                pulumi.String("neptune-cluster-demo"),
// 			Engine:                           pulumi.String("neptune"),
// 			IamDatabaseAuthenticationEnabled: pulumi.Bool(true),
// 			PreferredBackupWindow:            pulumi.String("07:00-09:00"),
// 			SkipFinalSnapshot:                pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// > **Note:** AWS Neptune does not support user name/password–based access control.
// See the AWS [Docs](https://docs.aws.amazon.com/neptune/latest/userguide/limits.html) for more information.
//
// ## Import
//
// `aws_neptune_cluster` can be imported by using the cluster identifier, e.g.,
//
// ```sh
//  $ pulumi import aws:neptune/cluster:Cluster example my-cluster
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately pulumi.BoolOutput `pulumi:"applyImmediately"`
	// The Neptune Cluster Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A list of EC2 Availability Zones that instances in the Neptune cluster can be created in.
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// The days to retain backups for. Default `1`
	BackupRetentionPeriod pulumi.IntPtrOutput `pulumi:"backupRetentionPeriod"`
	// The cluster identifier. If omitted, this provider will assign a random, unique identifier.
	ClusterIdentifier pulumi.StringOutput `pulumi:"clusterIdentifier"`
	// Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `clusterIdentifier`.
	ClusterIdentifierPrefix pulumi.StringOutput `pulumi:"clusterIdentifierPrefix"`
	// List of Neptune Instances that are a part of this cluster
	ClusterMembers pulumi.StringArrayOutput `pulumi:"clusterMembers"`
	// The Neptune Cluster Resource ID
	ClusterResourceId pulumi.StringOutput `pulumi:"clusterResourceId"`
	// If set to true, tags are copied to any snapshot of the DB cluster that is created.
	CopyTagsToSnapshot pulumi.BoolPtrOutput `pulumi:"copyTagsToSnapshot"`
	// A value that indicates whether the DB cluster has deletion protection enabled.The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	DeletionProtection pulumi.BoolPtrOutput `pulumi:"deletionProtection"`
	// A list of the log types this DB cluster is configured to export to Cloudwatch Logs. Currently only supports `audit`.
	EnableCloudwatchLogsExports pulumi.StringArrayOutput `pulumi:"enableCloudwatchLogsExports"`
	// The DNS address of the Neptune instance
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The name of the database engine to be used for this Neptune cluster. Defaults to `neptune`.
	Engine pulumi.StringPtrOutput `pulumi:"engine"`
	// The database engine version.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// The name of your final Neptune snapshot when this Neptune cluster is deleted. If omitted, no final snapshot will be made.
	FinalSnapshotIdentifier pulumi.StringPtrOutput `pulumi:"finalSnapshotIdentifier"`
	// The Route53 Hosted Zone ID of the endpoint
	HostedZoneId pulumi.StringOutput `pulumi:"hostedZoneId"`
	// Specifies whether or not mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.
	IamDatabaseAuthenticationEnabled pulumi.BoolPtrOutput `pulumi:"iamDatabaseAuthenticationEnabled"`
	// A List of ARNs for the IAM roles to associate to the Neptune Cluster.
	IamRoles pulumi.StringArrayOutput `pulumi:"iamRoles"`
	// The ARN for the KMS encryption key. When specifying `kmsKeyArn`, `storageEncrypted` needs to be set to true.
	KmsKeyArn pulumi.StringOutput `pulumi:"kmsKeyArn"`
	// A cluster parameter group to associate with the cluster.
	NeptuneClusterParameterGroupName pulumi.StringPtrOutput `pulumi:"neptuneClusterParameterGroupName"`
	// A Neptune subnet group to associate with this Neptune instance.
	NeptuneSubnetGroupName pulumi.StringOutput `pulumi:"neptuneSubnetGroupName"`
	// The port on which the Neptune accepts connections. Default is `8182`.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter. Time in UTC. Default: A 30-minute window selected at random from an 8-hour block of time per regionE.g., 04:00-09:00
	PreferredBackupWindow pulumi.StringOutput `pulumi:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur, in (UTC) e.g., wed:04:00-wed:04:30
	PreferredMaintenanceWindow pulumi.StringOutput `pulumi:"preferredMaintenanceWindow"`
	// A read-only endpoint for the Neptune cluster, automatically load-balanced across replicas
	ReaderEndpoint pulumi.StringOutput `pulumi:"readerEndpoint"`
	// ARN of a source Neptune cluster or Neptune instance if this Neptune cluster is to be created as a Read Replica.
	ReplicationSourceIdentifier pulumi.StringPtrOutput `pulumi:"replicationSourceIdentifier"`
	// Determines whether a final Neptune snapshot is created before the Neptune cluster is deleted. If true is specified, no Neptune snapshot is created. If false is specified, a Neptune snapshot is created before the Neptune cluster is deleted, using the value from `finalSnapshotIdentifier`. Default is `false`.
	SkipFinalSnapshot pulumi.BoolPtrOutput `pulumi:"skipFinalSnapshot"`
	// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a Neptune cluster snapshot, or the ARN when specifying a Neptune snapshot.
	SnapshotIdentifier pulumi.StringPtrOutput `pulumi:"snapshotIdentifier"`
	// Specifies whether the Neptune cluster is encrypted. The default is `false` if not specified.
	StorageEncrypted pulumi.BoolPtrOutput `pulumi:"storageEncrypted"`
	// A map of tags to assign to the Neptune cluster. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// List of VPC security groups to associate with the Cluster
	VpcSecurityGroupIds pulumi.StringArrayOutput `pulumi:"vpcSecurityGroupIds"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		args = &ClusterArgs{}
	}

	var resource Cluster
	err := ctx.RegisterResource("aws:neptune/cluster:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:neptune/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately *bool `pulumi:"applyImmediately"`
	// The Neptune Cluster Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// A list of EC2 Availability Zones that instances in the Neptune cluster can be created in.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The days to retain backups for. Default `1`
	BackupRetentionPeriod *int `pulumi:"backupRetentionPeriod"`
	// The cluster identifier. If omitted, this provider will assign a random, unique identifier.
	ClusterIdentifier *string `pulumi:"clusterIdentifier"`
	// Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `clusterIdentifier`.
	ClusterIdentifierPrefix *string `pulumi:"clusterIdentifierPrefix"`
	// List of Neptune Instances that are a part of this cluster
	ClusterMembers []string `pulumi:"clusterMembers"`
	// The Neptune Cluster Resource ID
	ClusterResourceId *string `pulumi:"clusterResourceId"`
	// If set to true, tags are copied to any snapshot of the DB cluster that is created.
	CopyTagsToSnapshot *bool `pulumi:"copyTagsToSnapshot"`
	// A value that indicates whether the DB cluster has deletion protection enabled.The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// A list of the log types this DB cluster is configured to export to Cloudwatch Logs. Currently only supports `audit`.
	EnableCloudwatchLogsExports []string `pulumi:"enableCloudwatchLogsExports"`
	// The DNS address of the Neptune instance
	Endpoint *string `pulumi:"endpoint"`
	// The name of the database engine to be used for this Neptune cluster. Defaults to `neptune`.
	Engine *string `pulumi:"engine"`
	// The database engine version.
	EngineVersion *string `pulumi:"engineVersion"`
	// The name of your final Neptune snapshot when this Neptune cluster is deleted. If omitted, no final snapshot will be made.
	FinalSnapshotIdentifier *string `pulumi:"finalSnapshotIdentifier"`
	// The Route53 Hosted Zone ID of the endpoint
	HostedZoneId *string `pulumi:"hostedZoneId"`
	// Specifies whether or not mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.
	IamDatabaseAuthenticationEnabled *bool `pulumi:"iamDatabaseAuthenticationEnabled"`
	// A List of ARNs for the IAM roles to associate to the Neptune Cluster.
	IamRoles []string `pulumi:"iamRoles"`
	// The ARN for the KMS encryption key. When specifying `kmsKeyArn`, `storageEncrypted` needs to be set to true.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// A cluster parameter group to associate with the cluster.
	NeptuneClusterParameterGroupName *string `pulumi:"neptuneClusterParameterGroupName"`
	// A Neptune subnet group to associate with this Neptune instance.
	NeptuneSubnetGroupName *string `pulumi:"neptuneSubnetGroupName"`
	// The port on which the Neptune accepts connections. Default is `8182`.
	Port *int `pulumi:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter. Time in UTC. Default: A 30-minute window selected at random from an 8-hour block of time per regionE.g., 04:00-09:00
	PreferredBackupWindow *string `pulumi:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur, in (UTC) e.g., wed:04:00-wed:04:30
	PreferredMaintenanceWindow *string `pulumi:"preferredMaintenanceWindow"`
	// A read-only endpoint for the Neptune cluster, automatically load-balanced across replicas
	ReaderEndpoint *string `pulumi:"readerEndpoint"`
	// ARN of a source Neptune cluster or Neptune instance if this Neptune cluster is to be created as a Read Replica.
	ReplicationSourceIdentifier *string `pulumi:"replicationSourceIdentifier"`
	// Determines whether a final Neptune snapshot is created before the Neptune cluster is deleted. If true is specified, no Neptune snapshot is created. If false is specified, a Neptune snapshot is created before the Neptune cluster is deleted, using the value from `finalSnapshotIdentifier`. Default is `false`.
	SkipFinalSnapshot *bool `pulumi:"skipFinalSnapshot"`
	// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a Neptune cluster snapshot, or the ARN when specifying a Neptune snapshot.
	SnapshotIdentifier *string `pulumi:"snapshotIdentifier"`
	// Specifies whether the Neptune cluster is encrypted. The default is `false` if not specified.
	StorageEncrypted *bool `pulumi:"storageEncrypted"`
	// A map of tags to assign to the Neptune cluster. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// List of VPC security groups to associate with the Cluster
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
}

type ClusterState struct {
	// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately pulumi.BoolPtrInput
	// The Neptune Cluster Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// A list of EC2 Availability Zones that instances in the Neptune cluster can be created in.
	AvailabilityZones pulumi.StringArrayInput
	// The days to retain backups for. Default `1`
	BackupRetentionPeriod pulumi.IntPtrInput
	// The cluster identifier. If omitted, this provider will assign a random, unique identifier.
	ClusterIdentifier pulumi.StringPtrInput
	// Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `clusterIdentifier`.
	ClusterIdentifierPrefix pulumi.StringPtrInput
	// List of Neptune Instances that are a part of this cluster
	ClusterMembers pulumi.StringArrayInput
	// The Neptune Cluster Resource ID
	ClusterResourceId pulumi.StringPtrInput
	// If set to true, tags are copied to any snapshot of the DB cluster that is created.
	CopyTagsToSnapshot pulumi.BoolPtrInput
	// A value that indicates whether the DB cluster has deletion protection enabled.The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	DeletionProtection pulumi.BoolPtrInput
	// A list of the log types this DB cluster is configured to export to Cloudwatch Logs. Currently only supports `audit`.
	EnableCloudwatchLogsExports pulumi.StringArrayInput
	// The DNS address of the Neptune instance
	Endpoint pulumi.StringPtrInput
	// The name of the database engine to be used for this Neptune cluster. Defaults to `neptune`.
	Engine pulumi.StringPtrInput
	// The database engine version.
	EngineVersion pulumi.StringPtrInput
	// The name of your final Neptune snapshot when this Neptune cluster is deleted. If omitted, no final snapshot will be made.
	FinalSnapshotIdentifier pulumi.StringPtrInput
	// The Route53 Hosted Zone ID of the endpoint
	HostedZoneId pulumi.StringPtrInput
	// Specifies whether or not mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.
	IamDatabaseAuthenticationEnabled pulumi.BoolPtrInput
	// A List of ARNs for the IAM roles to associate to the Neptune Cluster.
	IamRoles pulumi.StringArrayInput
	// The ARN for the KMS encryption key. When specifying `kmsKeyArn`, `storageEncrypted` needs to be set to true.
	KmsKeyArn pulumi.StringPtrInput
	// A cluster parameter group to associate with the cluster.
	NeptuneClusterParameterGroupName pulumi.StringPtrInput
	// A Neptune subnet group to associate with this Neptune instance.
	NeptuneSubnetGroupName pulumi.StringPtrInput
	// The port on which the Neptune accepts connections. Default is `8182`.
	Port pulumi.IntPtrInput
	// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter. Time in UTC. Default: A 30-minute window selected at random from an 8-hour block of time per regionE.g., 04:00-09:00
	PreferredBackupWindow pulumi.StringPtrInput
	// The weekly time range during which system maintenance can occur, in (UTC) e.g., wed:04:00-wed:04:30
	PreferredMaintenanceWindow pulumi.StringPtrInput
	// A read-only endpoint for the Neptune cluster, automatically load-balanced across replicas
	ReaderEndpoint pulumi.StringPtrInput
	// ARN of a source Neptune cluster or Neptune instance if this Neptune cluster is to be created as a Read Replica.
	ReplicationSourceIdentifier pulumi.StringPtrInput
	// Determines whether a final Neptune snapshot is created before the Neptune cluster is deleted. If true is specified, no Neptune snapshot is created. If false is specified, a Neptune snapshot is created before the Neptune cluster is deleted, using the value from `finalSnapshotIdentifier`. Default is `false`.
	SkipFinalSnapshot pulumi.BoolPtrInput
	// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a Neptune cluster snapshot, or the ARN when specifying a Neptune snapshot.
	SnapshotIdentifier pulumi.StringPtrInput
	// Specifies whether the Neptune cluster is encrypted. The default is `false` if not specified.
	StorageEncrypted pulumi.BoolPtrInput
	// A map of tags to assign to the Neptune cluster. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// List of VPC security groups to associate with the Cluster
	VpcSecurityGroupIds pulumi.StringArrayInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately *bool `pulumi:"applyImmediately"`
	// A list of EC2 Availability Zones that instances in the Neptune cluster can be created in.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The days to retain backups for. Default `1`
	BackupRetentionPeriod *int `pulumi:"backupRetentionPeriod"`
	// The cluster identifier. If omitted, this provider will assign a random, unique identifier.
	ClusterIdentifier *string `pulumi:"clusterIdentifier"`
	// Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `clusterIdentifier`.
	ClusterIdentifierPrefix *string `pulumi:"clusterIdentifierPrefix"`
	// If set to true, tags are copied to any snapshot of the DB cluster that is created.
	CopyTagsToSnapshot *bool `pulumi:"copyTagsToSnapshot"`
	// A value that indicates whether the DB cluster has deletion protection enabled.The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// A list of the log types this DB cluster is configured to export to Cloudwatch Logs. Currently only supports `audit`.
	EnableCloudwatchLogsExports []string `pulumi:"enableCloudwatchLogsExports"`
	// The name of the database engine to be used for this Neptune cluster. Defaults to `neptune`.
	Engine *string `pulumi:"engine"`
	// The database engine version.
	EngineVersion *string `pulumi:"engineVersion"`
	// The name of your final Neptune snapshot when this Neptune cluster is deleted. If omitted, no final snapshot will be made.
	FinalSnapshotIdentifier *string `pulumi:"finalSnapshotIdentifier"`
	// Specifies whether or not mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.
	IamDatabaseAuthenticationEnabled *bool `pulumi:"iamDatabaseAuthenticationEnabled"`
	// A List of ARNs for the IAM roles to associate to the Neptune Cluster.
	IamRoles []string `pulumi:"iamRoles"`
	// The ARN for the KMS encryption key. When specifying `kmsKeyArn`, `storageEncrypted` needs to be set to true.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// A cluster parameter group to associate with the cluster.
	NeptuneClusterParameterGroupName *string `pulumi:"neptuneClusterParameterGroupName"`
	// A Neptune subnet group to associate with this Neptune instance.
	NeptuneSubnetGroupName *string `pulumi:"neptuneSubnetGroupName"`
	// The port on which the Neptune accepts connections. Default is `8182`.
	Port *int `pulumi:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter. Time in UTC. Default: A 30-minute window selected at random from an 8-hour block of time per regionE.g., 04:00-09:00
	PreferredBackupWindow *string `pulumi:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur, in (UTC) e.g., wed:04:00-wed:04:30
	PreferredMaintenanceWindow *string `pulumi:"preferredMaintenanceWindow"`
	// ARN of a source Neptune cluster or Neptune instance if this Neptune cluster is to be created as a Read Replica.
	ReplicationSourceIdentifier *string `pulumi:"replicationSourceIdentifier"`
	// Determines whether a final Neptune snapshot is created before the Neptune cluster is deleted. If true is specified, no Neptune snapshot is created. If false is specified, a Neptune snapshot is created before the Neptune cluster is deleted, using the value from `finalSnapshotIdentifier`. Default is `false`.
	SkipFinalSnapshot *bool `pulumi:"skipFinalSnapshot"`
	// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a Neptune cluster snapshot, or the ARN when specifying a Neptune snapshot.
	SnapshotIdentifier *string `pulumi:"snapshotIdentifier"`
	// Specifies whether the Neptune cluster is encrypted. The default is `false` if not specified.
	StorageEncrypted *bool `pulumi:"storageEncrypted"`
	// A map of tags to assign to the Neptune cluster. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// List of VPC security groups to associate with the Cluster
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately pulumi.BoolPtrInput
	// A list of EC2 Availability Zones that instances in the Neptune cluster can be created in.
	AvailabilityZones pulumi.StringArrayInput
	// The days to retain backups for. Default `1`
	BackupRetentionPeriod pulumi.IntPtrInput
	// The cluster identifier. If omitted, this provider will assign a random, unique identifier.
	ClusterIdentifier pulumi.StringPtrInput
	// Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `clusterIdentifier`.
	ClusterIdentifierPrefix pulumi.StringPtrInput
	// If set to true, tags are copied to any snapshot of the DB cluster that is created.
	CopyTagsToSnapshot pulumi.BoolPtrInput
	// A value that indicates whether the DB cluster has deletion protection enabled.The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	DeletionProtection pulumi.BoolPtrInput
	// A list of the log types this DB cluster is configured to export to Cloudwatch Logs. Currently only supports `audit`.
	EnableCloudwatchLogsExports pulumi.StringArrayInput
	// The name of the database engine to be used for this Neptune cluster. Defaults to `neptune`.
	Engine pulumi.StringPtrInput
	// The database engine version.
	EngineVersion pulumi.StringPtrInput
	// The name of your final Neptune snapshot when this Neptune cluster is deleted. If omitted, no final snapshot will be made.
	FinalSnapshotIdentifier pulumi.StringPtrInput
	// Specifies whether or not mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.
	IamDatabaseAuthenticationEnabled pulumi.BoolPtrInput
	// A List of ARNs for the IAM roles to associate to the Neptune Cluster.
	IamRoles pulumi.StringArrayInput
	// The ARN for the KMS encryption key. When specifying `kmsKeyArn`, `storageEncrypted` needs to be set to true.
	KmsKeyArn pulumi.StringPtrInput
	// A cluster parameter group to associate with the cluster.
	NeptuneClusterParameterGroupName pulumi.StringPtrInput
	// A Neptune subnet group to associate with this Neptune instance.
	NeptuneSubnetGroupName pulumi.StringPtrInput
	// The port on which the Neptune accepts connections. Default is `8182`.
	Port pulumi.IntPtrInput
	// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter. Time in UTC. Default: A 30-minute window selected at random from an 8-hour block of time per regionE.g., 04:00-09:00
	PreferredBackupWindow pulumi.StringPtrInput
	// The weekly time range during which system maintenance can occur, in (UTC) e.g., wed:04:00-wed:04:30
	PreferredMaintenanceWindow pulumi.StringPtrInput
	// ARN of a source Neptune cluster or Neptune instance if this Neptune cluster is to be created as a Read Replica.
	ReplicationSourceIdentifier pulumi.StringPtrInput
	// Determines whether a final Neptune snapshot is created before the Neptune cluster is deleted. If true is specified, no Neptune snapshot is created. If false is specified, a Neptune snapshot is created before the Neptune cluster is deleted, using the value from `finalSnapshotIdentifier`. Default is `false`.
	SkipFinalSnapshot pulumi.BoolPtrInput
	// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a Neptune cluster snapshot, or the ARN when specifying a Neptune snapshot.
	SnapshotIdentifier pulumi.StringPtrInput
	// Specifies whether the Neptune cluster is encrypted. The default is `false` if not specified.
	StorageEncrypted pulumi.BoolPtrInput
	// A map of tags to assign to the Neptune cluster. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// List of VPC security groups to associate with the Cluster
	VpcSecurityGroupIds pulumi.StringArrayInput
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

type ClusterOutput struct{ *pulumi.OutputState }

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
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Cluster) *Cluster {
		return &v
	}).(ClusterPtrOutput)
}

type ClusterPtrOutput struct{ *pulumi.OutputState }

func (ClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil))
}

func (o ClusterPtrOutput) ToClusterPtrOutput() ClusterPtrOutput {
	return o
}

func (o ClusterPtrOutput) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return o
}

func (o ClusterPtrOutput) Elem() ClusterOutput {
	return o.ApplyT(func(v *Cluster) Cluster {
		if v != nil {
			return *v
		}
		var ret Cluster
		return ret
	}).(ClusterOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterPtrInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterArrayInput)(nil)).Elem(), ClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMapInput)(nil)).Elem(), ClusterMap{})
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterPtrOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}
