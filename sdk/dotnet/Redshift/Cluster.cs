// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.RedShift
{
    /// <summary>
    /// Provides a Redshift Cluster Resource.
    /// 
    /// &gt; **Note:** All arguments including the username and password will be stored in the raw state as plain-text.
    /// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/redshift_cluster.html.markdown.
    /// </summary>
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// If true , major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default is true
        /// </summary>
        [Output("allowVersionUpgrade")]
        public Output<bool?> AllowVersionUpgrade { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of cluster
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Even if automated snapshots are disabled, you can still create manual snapshots when you want with create-cluster-snapshot. Default is 1.
        /// </summary>
        [Output("automatedSnapshotRetentionPeriod")]
        public Output<int?> AutomatedSnapshotRetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. For example, if you have several EC2 instances running in a specific Availability Zone, then you might want the cluster to be provisioned in the same zone in order to decrease network latency.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// The Cluster Identifier. Must be a lower case
        /// string.
        /// </summary>
        [Output("clusterIdentifier")]
        public Output<string> ClusterIdentifier { get; private set; } = null!;

        /// <summary>
        /// The name of the parameter group to be associated with this cluster.
        /// </summary>
        [Output("clusterParameterGroupName")]
        public Output<string> ClusterParameterGroupName { get; private set; } = null!;

        /// <summary>
        /// The public key for the cluster
        /// </summary>
        [Output("clusterPublicKey")]
        public Output<string> ClusterPublicKey { get; private set; } = null!;

        /// <summary>
        /// The specific revision number of the database in the cluster
        /// </summary>
        [Output("clusterRevisionNumber")]
        public Output<string> ClusterRevisionNumber { get; private set; } = null!;

        /// <summary>
        /// A list of security groups to be associated with this cluster.
        /// </summary>
        [Output("clusterSecurityGroups")]
        public Output<ImmutableArray<string>> ClusterSecurityGroups { get; private set; } = null!;

        /// <summary>
        /// The name of a cluster subnet group to be associated with this cluster. If this parameter is not provided the resulting cluster will be deployed outside virtual private cloud (VPC).
        /// </summary>
        [Output("clusterSubnetGroupName")]
        public Output<string> ClusterSubnetGroupName { get; private set; } = null!;

        /// <summary>
        /// The cluster type to use. Either `single-node` or `multi-node`.
        /// </summary>
        [Output("clusterType")]
        public Output<string> ClusterType { get; private set; } = null!;

        /// <summary>
        /// The version of the Amazon Redshift engine software that you want to deploy on the cluster.
        /// The version selected runs on all the nodes in the cluster.
        /// </summary>
        [Output("clusterVersion")]
        public Output<string?> ClusterVersion { get; private set; } = null!;

        /// <summary>
        /// The name of the first database to be created when the cluster is created.
        /// If you do not provide a name, Amazon Redshift will create a default database called `dev`.
        /// </summary>
        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// The DNS name of the cluster
        /// </summary>
        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// The Elastic IP (EIP) address for the cluster.
        /// </summary>
        [Output("elasticIp")]
        public Output<string?> ElasticIp { get; private set; } = null!;

        /// <summary>
        /// If true , the data in the cluster is encrypted at rest.
        /// </summary>
        [Output("encrypted")]
        public Output<bool?> Encrypted { get; private set; } = null!;

        /// <summary>
        /// The connection endpoint
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// If true , enhanced VPC routing is enabled.
        /// </summary>
        [Output("enhancedVpcRouting")]
        public Output<bool> EnhancedVpcRouting { get; private set; } = null!;

        /// <summary>
        /// The identifier of the final snapshot that is to be created immediately before deleting the cluster. If this parameter is provided, `skip_final_snapshot` must be false.
        /// </summary>
        [Output("finalSnapshotIdentifier")]
        public Output<string?> FinalSnapshotIdentifier { get; private set; } = null!;

        /// <summary>
        /// A list of IAM Role ARNs to associate with the cluster. A Maximum of 10 can be associated to the cluster at any time.
        /// </summary>
        [Output("iamRoles")]
        public Output<ImmutableArray<string>> IamRoles { get; private set; } = null!;

        /// <summary>
        /// The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// Logging, documented below.
        /// </summary>
        [Output("logging")]
        public Output<Outputs.ClusterLogging?> Logging { get; private set; } = null!;

        /// <summary>
        /// Password for the master DB user.
        /// Note that this may show up in logs, and it will be stored in the state file. Password must contain at least 8 chars and
        /// contain at least one uppercase letter, one lowercase letter, and one number.
        /// </summary>
        [Output("masterPassword")]
        public Output<string?> MasterPassword { get; private set; } = null!;

        /// <summary>
        /// Username for the master DB user.
        /// </summary>
        [Output("masterUsername")]
        public Output<string?> MasterUsername { get; private set; } = null!;

        /// <summary>
        /// The node type to be provisioned for the cluster.
        /// </summary>
        [Output("nodeType")]
        public Output<string> NodeType { get; private set; } = null!;

        /// <summary>
        /// The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node. Default is 1.
        /// </summary>
        [Output("numberOfNodes")]
        public Output<int?> NumberOfNodes { get; private set; } = null!;

        /// <summary>
        /// The AWS customer account used to create or copy the snapshot. Required if you are restoring a snapshot you do not own, optional if you own the snapshot.
        /// </summary>
        [Output("ownerAccount")]
        public Output<string?> OwnerAccount { get; private set; } = null!;

        /// <summary>
        /// The port number on which the cluster accepts incoming connections.
        /// The cluster is accessible only via the JDBC and ODBC connection strings. Part of the connection string requires the port on which the cluster will listen for incoming connections. Default port is 5439.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// The weekly time range (in UTC) during which automated cluster maintenance can occur.
        /// Format: ddd:hh24:mi-ddd:hh24:mi
        /// </summary>
        [Output("preferredMaintenanceWindow")]
        public Output<string> PreferredMaintenanceWindow { get; private set; } = null!;

        /// <summary>
        /// If true, the cluster can be accessed from a public network. Default is `true`.
        /// </summary>
        [Output("publiclyAccessible")]
        public Output<bool?> PubliclyAccessible { get; private set; } = null!;

        /// <summary>
        /// Determines whether a final snapshot of the cluster is created before Amazon Redshift deletes the cluster. If true , a final cluster snapshot is not created. If false , a final cluster snapshot is created before the cluster is deleted. Default is false.
        /// </summary>
        [Output("skipFinalSnapshot")]
        public Output<bool?> SkipFinalSnapshot { get; private set; } = null!;

        /// <summary>
        /// The name of the cluster the source snapshot was created from.
        /// </summary>
        [Output("snapshotClusterIdentifier")]
        public Output<string?> SnapshotClusterIdentifier { get; private set; } = null!;

        /// <summary>
        /// Configuration of automatic copy of snapshots from one region to another. Documented below.
        /// </summary>
        [Output("snapshotCopy")]
        public Output<Outputs.ClusterSnapshotCopy?> SnapshotCopy { get; private set; } = null!;

        /// <summary>
        /// The name of the snapshot from which to create the new cluster.
        /// </summary>
        [Output("snapshotIdentifier")]
        public Output<string?> SnapshotIdentifier { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
        /// </summary>
        [Output("vpcSecurityGroupIds")]
        public Output<ImmutableArray<string>> VpcSecurityGroupIds { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("aws:redshift/cluster:Cluster", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("aws:redshift/cluster:Cluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true , major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default is true
        /// </summary>
        [Input("allowVersionUpgrade")]
        public Input<bool>? AllowVersionUpgrade { get; set; }

        /// <summary>
        /// The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Even if automated snapshots are disabled, you can still create manual snapshots when you want with create-cluster-snapshot. Default is 1.
        /// </summary>
        [Input("automatedSnapshotRetentionPeriod")]
        public Input<int>? AutomatedSnapshotRetentionPeriod { get; set; }

        /// <summary>
        /// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. For example, if you have several EC2 instances running in a specific Availability Zone, then you might want the cluster to be provisioned in the same zone in order to decrease network latency.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The Cluster Identifier. Must be a lower case
        /// string.
        /// </summary>
        [Input("clusterIdentifier", required: true)]
        public Input<string> ClusterIdentifier { get; set; } = null!;

        /// <summary>
        /// The name of the parameter group to be associated with this cluster.
        /// </summary>
        [Input("clusterParameterGroupName")]
        public Input<string>? ClusterParameterGroupName { get; set; }

        /// <summary>
        /// The public key for the cluster
        /// </summary>
        [Input("clusterPublicKey")]
        public Input<string>? ClusterPublicKey { get; set; }

        /// <summary>
        /// The specific revision number of the database in the cluster
        /// </summary>
        [Input("clusterRevisionNumber")]
        public Input<string>? ClusterRevisionNumber { get; set; }

        [Input("clusterSecurityGroups")]
        private InputList<string>? _clusterSecurityGroups;

        /// <summary>
        /// A list of security groups to be associated with this cluster.
        /// </summary>
        public InputList<string> ClusterSecurityGroups
        {
            get => _clusterSecurityGroups ?? (_clusterSecurityGroups = new InputList<string>());
            set => _clusterSecurityGroups = value;
        }

        /// <summary>
        /// The name of a cluster subnet group to be associated with this cluster. If this parameter is not provided the resulting cluster will be deployed outside virtual private cloud (VPC).
        /// </summary>
        [Input("clusterSubnetGroupName")]
        public Input<string>? ClusterSubnetGroupName { get; set; }

        /// <summary>
        /// The cluster type to use. Either `single-node` or `multi-node`.
        /// </summary>
        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        /// <summary>
        /// The version of the Amazon Redshift engine software that you want to deploy on the cluster.
        /// The version selected runs on all the nodes in the cluster.
        /// </summary>
        [Input("clusterVersion")]
        public Input<string>? ClusterVersion { get; set; }

        /// <summary>
        /// The name of the first database to be created when the cluster is created.
        /// If you do not provide a name, Amazon Redshift will create a default database called `dev`.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// The Elastic IP (EIP) address for the cluster.
        /// </summary>
        [Input("elasticIp")]
        public Input<string>? ElasticIp { get; set; }

        /// <summary>
        /// If true , the data in the cluster is encrypted at rest.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// The connection endpoint
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// If true , enhanced VPC routing is enabled.
        /// </summary>
        [Input("enhancedVpcRouting")]
        public Input<bool>? EnhancedVpcRouting { get; set; }

        /// <summary>
        /// The identifier of the final snapshot that is to be created immediately before deleting the cluster. If this parameter is provided, `skip_final_snapshot` must be false.
        /// </summary>
        [Input("finalSnapshotIdentifier")]
        public Input<string>? FinalSnapshotIdentifier { get; set; }

        [Input("iamRoles")]
        private InputList<string>? _iamRoles;

        /// <summary>
        /// A list of IAM Role ARNs to associate with the cluster. A Maximum of 10 can be associated to the cluster at any time.
        /// </summary>
        public InputList<string> IamRoles
        {
            get => _iamRoles ?? (_iamRoles = new InputList<string>());
            set => _iamRoles = value;
        }

        /// <summary>
        /// The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Logging, documented below.
        /// </summary>
        [Input("logging")]
        public Input<Inputs.ClusterLoggingArgs>? Logging { get; set; }

        /// <summary>
        /// Password for the master DB user.
        /// Note that this may show up in logs, and it will be stored in the state file. Password must contain at least 8 chars and
        /// contain at least one uppercase letter, one lowercase letter, and one number.
        /// </summary>
        [Input("masterPassword")]
        public Input<string>? MasterPassword { get; set; }

        /// <summary>
        /// Username for the master DB user.
        /// </summary>
        [Input("masterUsername")]
        public Input<string>? MasterUsername { get; set; }

        /// <summary>
        /// The node type to be provisioned for the cluster.
        /// </summary>
        [Input("nodeType", required: true)]
        public Input<string> NodeType { get; set; } = null!;

        /// <summary>
        /// The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node. Default is 1.
        /// </summary>
        [Input("numberOfNodes")]
        public Input<int>? NumberOfNodes { get; set; }

        /// <summary>
        /// The AWS customer account used to create or copy the snapshot. Required if you are restoring a snapshot you do not own, optional if you own the snapshot.
        /// </summary>
        [Input("ownerAccount")]
        public Input<string>? OwnerAccount { get; set; }

        /// <summary>
        /// The port number on which the cluster accepts incoming connections.
        /// The cluster is accessible only via the JDBC and ODBC connection strings. Part of the connection string requires the port on which the cluster will listen for incoming connections. Default port is 5439.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The weekly time range (in UTC) during which automated cluster maintenance can occur.
        /// Format: ddd:hh24:mi-ddd:hh24:mi
        /// </summary>
        [Input("preferredMaintenanceWindow")]
        public Input<string>? PreferredMaintenanceWindow { get; set; }

        /// <summary>
        /// If true, the cluster can be accessed from a public network. Default is `true`.
        /// </summary>
        [Input("publiclyAccessible")]
        public Input<bool>? PubliclyAccessible { get; set; }

        /// <summary>
        /// Determines whether a final snapshot of the cluster is created before Amazon Redshift deletes the cluster. If true , a final cluster snapshot is not created. If false , a final cluster snapshot is created before the cluster is deleted. Default is false.
        /// </summary>
        [Input("skipFinalSnapshot")]
        public Input<bool>? SkipFinalSnapshot { get; set; }

        /// <summary>
        /// The name of the cluster the source snapshot was created from.
        /// </summary>
        [Input("snapshotClusterIdentifier")]
        public Input<string>? SnapshotClusterIdentifier { get; set; }

        /// <summary>
        /// Configuration of automatic copy of snapshots from one region to another. Documented below.
        /// </summary>
        [Input("snapshotCopy")]
        public Input<Inputs.ClusterSnapshotCopyArgs>? SnapshotCopy { get; set; }

        /// <summary>
        /// The name of the snapshot from which to create the new cluster.
        /// </summary>
        [Input("snapshotIdentifier")]
        public Input<string>? SnapshotIdentifier { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vpcSecurityGroupIds")]
        private InputList<string>? _vpcSecurityGroupIds;

        /// <summary>
        /// A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
        /// </summary>
        public InputList<string> VpcSecurityGroupIds
        {
            get => _vpcSecurityGroupIds ?? (_vpcSecurityGroupIds = new InputList<string>());
            set => _vpcSecurityGroupIds = value;
        }

        public ClusterArgs()
        {
        }
    }

    public sealed class ClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true , major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default is true
        /// </summary>
        [Input("allowVersionUpgrade")]
        public Input<bool>? AllowVersionUpgrade { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of cluster
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Even if automated snapshots are disabled, you can still create manual snapshots when you want with create-cluster-snapshot. Default is 1.
        /// </summary>
        [Input("automatedSnapshotRetentionPeriod")]
        public Input<int>? AutomatedSnapshotRetentionPeriod { get; set; }

        /// <summary>
        /// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. For example, if you have several EC2 instances running in a specific Availability Zone, then you might want the cluster to be provisioned in the same zone in order to decrease network latency.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The Cluster Identifier. Must be a lower case
        /// string.
        /// </summary>
        [Input("clusterIdentifier")]
        public Input<string>? ClusterIdentifier { get; set; }

        /// <summary>
        /// The name of the parameter group to be associated with this cluster.
        /// </summary>
        [Input("clusterParameterGroupName")]
        public Input<string>? ClusterParameterGroupName { get; set; }

        /// <summary>
        /// The public key for the cluster
        /// </summary>
        [Input("clusterPublicKey")]
        public Input<string>? ClusterPublicKey { get; set; }

        /// <summary>
        /// The specific revision number of the database in the cluster
        /// </summary>
        [Input("clusterRevisionNumber")]
        public Input<string>? ClusterRevisionNumber { get; set; }

        [Input("clusterSecurityGroups")]
        private InputList<string>? _clusterSecurityGroups;

        /// <summary>
        /// A list of security groups to be associated with this cluster.
        /// </summary>
        public InputList<string> ClusterSecurityGroups
        {
            get => _clusterSecurityGroups ?? (_clusterSecurityGroups = new InputList<string>());
            set => _clusterSecurityGroups = value;
        }

        /// <summary>
        /// The name of a cluster subnet group to be associated with this cluster. If this parameter is not provided the resulting cluster will be deployed outside virtual private cloud (VPC).
        /// </summary>
        [Input("clusterSubnetGroupName")]
        public Input<string>? ClusterSubnetGroupName { get; set; }

        /// <summary>
        /// The cluster type to use. Either `single-node` or `multi-node`.
        /// </summary>
        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        /// <summary>
        /// The version of the Amazon Redshift engine software that you want to deploy on the cluster.
        /// The version selected runs on all the nodes in the cluster.
        /// </summary>
        [Input("clusterVersion")]
        public Input<string>? ClusterVersion { get; set; }

        /// <summary>
        /// The name of the first database to be created when the cluster is created.
        /// If you do not provide a name, Amazon Redshift will create a default database called `dev`.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// The DNS name of the cluster
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        /// <summary>
        /// The Elastic IP (EIP) address for the cluster.
        /// </summary>
        [Input("elasticIp")]
        public Input<string>? ElasticIp { get; set; }

        /// <summary>
        /// If true , the data in the cluster is encrypted at rest.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// The connection endpoint
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// If true , enhanced VPC routing is enabled.
        /// </summary>
        [Input("enhancedVpcRouting")]
        public Input<bool>? EnhancedVpcRouting { get; set; }

        /// <summary>
        /// The identifier of the final snapshot that is to be created immediately before deleting the cluster. If this parameter is provided, `skip_final_snapshot` must be false.
        /// </summary>
        [Input("finalSnapshotIdentifier")]
        public Input<string>? FinalSnapshotIdentifier { get; set; }

        [Input("iamRoles")]
        private InputList<string>? _iamRoles;

        /// <summary>
        /// A list of IAM Role ARNs to associate with the cluster. A Maximum of 10 can be associated to the cluster at any time.
        /// </summary>
        public InputList<string> IamRoles
        {
            get => _iamRoles ?? (_iamRoles = new InputList<string>());
            set => _iamRoles = value;
        }

        /// <summary>
        /// The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Logging, documented below.
        /// </summary>
        [Input("logging")]
        public Input<Inputs.ClusterLoggingGetArgs>? Logging { get; set; }

        /// <summary>
        /// Password for the master DB user.
        /// Note that this may show up in logs, and it will be stored in the state file. Password must contain at least 8 chars and
        /// contain at least one uppercase letter, one lowercase letter, and one number.
        /// </summary>
        [Input("masterPassword")]
        public Input<string>? MasterPassword { get; set; }

        /// <summary>
        /// Username for the master DB user.
        /// </summary>
        [Input("masterUsername")]
        public Input<string>? MasterUsername { get; set; }

        /// <summary>
        /// The node type to be provisioned for the cluster.
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        /// <summary>
        /// The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node. Default is 1.
        /// </summary>
        [Input("numberOfNodes")]
        public Input<int>? NumberOfNodes { get; set; }

        /// <summary>
        /// The AWS customer account used to create or copy the snapshot. Required if you are restoring a snapshot you do not own, optional if you own the snapshot.
        /// </summary>
        [Input("ownerAccount")]
        public Input<string>? OwnerAccount { get; set; }

        /// <summary>
        /// The port number on which the cluster accepts incoming connections.
        /// The cluster is accessible only via the JDBC and ODBC connection strings. Part of the connection string requires the port on which the cluster will listen for incoming connections. Default port is 5439.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The weekly time range (in UTC) during which automated cluster maintenance can occur.
        /// Format: ddd:hh24:mi-ddd:hh24:mi
        /// </summary>
        [Input("preferredMaintenanceWindow")]
        public Input<string>? PreferredMaintenanceWindow { get; set; }

        /// <summary>
        /// If true, the cluster can be accessed from a public network. Default is `true`.
        /// </summary>
        [Input("publiclyAccessible")]
        public Input<bool>? PubliclyAccessible { get; set; }

        /// <summary>
        /// Determines whether a final snapshot of the cluster is created before Amazon Redshift deletes the cluster. If true , a final cluster snapshot is not created. If false , a final cluster snapshot is created before the cluster is deleted. Default is false.
        /// </summary>
        [Input("skipFinalSnapshot")]
        public Input<bool>? SkipFinalSnapshot { get; set; }

        /// <summary>
        /// The name of the cluster the source snapshot was created from.
        /// </summary>
        [Input("snapshotClusterIdentifier")]
        public Input<string>? SnapshotClusterIdentifier { get; set; }

        /// <summary>
        /// Configuration of automatic copy of snapshots from one region to another. Documented below.
        /// </summary>
        [Input("snapshotCopy")]
        public Input<Inputs.ClusterSnapshotCopyGetArgs>? SnapshotCopy { get; set; }

        /// <summary>
        /// The name of the snapshot from which to create the new cluster.
        /// </summary>
        [Input("snapshotIdentifier")]
        public Input<string>? SnapshotIdentifier { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vpcSecurityGroupIds")]
        private InputList<string>? _vpcSecurityGroupIds;

        /// <summary>
        /// A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
        /// </summary>
        public InputList<string> VpcSecurityGroupIds
        {
            get => _vpcSecurityGroupIds ?? (_vpcSecurityGroupIds = new InputList<string>());
            set => _vpcSecurityGroupIds = value;
        }

        public ClusterState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ClusterLoggingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
        /// For more information on the permissions required for the bucket, please read the AWS [documentation](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        /// <summary>
        /// Enables logging information such as queries and connection attempts, for the specified Amazon Redshift cluster.
        /// </summary>
        [Input("enable", required: true)]
        public Input<bool> Enable { get; set; } = null!;

        /// <summary>
        /// The prefix applied to the log file names.
        /// </summary>
        [Input("s3KeyPrefix")]
        public Input<string>? S3KeyPrefix { get; set; }

        public ClusterLoggingArgs()
        {
        }
    }

    public sealed class ClusterLoggingGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
        /// For more information on the permissions required for the bucket, please read the AWS [documentation](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        /// <summary>
        /// Enables logging information such as queries and connection attempts, for the specified Amazon Redshift cluster.
        /// </summary>
        [Input("enable", required: true)]
        public Input<bool> Enable { get; set; } = null!;

        /// <summary>
        /// The prefix applied to the log file names.
        /// </summary>
        [Input("s3KeyPrefix")]
        public Input<string>? S3KeyPrefix { get; set; }

        public ClusterLoggingGetArgs()
        {
        }
    }

    public sealed class ClusterSnapshotCopyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination region that you want to copy snapshots to.
        /// </summary>
        [Input("destinationRegion", required: true)]
        public Input<string> DestinationRegion { get; set; } = null!;

        /// <summary>
        /// The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
        /// </summary>
        [Input("grantName")]
        public Input<string>? GrantName { get; set; }

        /// <summary>
        /// The number of days to retain automated snapshots in the destination region after they are copied from the source region. Defaults to `7`.
        /// </summary>
        [Input("retentionPeriod")]
        public Input<int>? RetentionPeriod { get; set; }

        public ClusterSnapshotCopyArgs()
        {
        }
    }

    public sealed class ClusterSnapshotCopyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination region that you want to copy snapshots to.
        /// </summary>
        [Input("destinationRegion", required: true)]
        public Input<string> DestinationRegion { get; set; } = null!;

        /// <summary>
        /// The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
        /// </summary>
        [Input("grantName")]
        public Input<string>? GrantName { get; set; }

        /// <summary>
        /// The number of days to retain automated snapshots in the destination region after they are copied from the source region. Defaults to `7`.
        /// </summary>
        [Input("retentionPeriod")]
        public Input<int>? RetentionPeriod { get; set; }

        public ClusterSnapshotCopyGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ClusterLogging
    {
        /// <summary>
        /// The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
        /// For more information on the permissions required for the bucket, please read the AWS [documentation](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
        /// </summary>
        public readonly string BucketName;
        /// <summary>
        /// Enables logging information such as queries and connection attempts, for the specified Amazon Redshift cluster.
        /// </summary>
        public readonly bool Enable;
        /// <summary>
        /// The prefix applied to the log file names.
        /// </summary>
        public readonly string S3KeyPrefix;

        [OutputConstructor]
        private ClusterLogging(
            string bucketName,
            bool enable,
            string s3KeyPrefix)
        {
            BucketName = bucketName;
            Enable = enable;
            S3KeyPrefix = s3KeyPrefix;
        }
    }

    [OutputType]
    public sealed class ClusterSnapshotCopy
    {
        /// <summary>
        /// The destination region that you want to copy snapshots to.
        /// </summary>
        public readonly string DestinationRegion;
        /// <summary>
        /// The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
        /// </summary>
        public readonly string? GrantName;
        /// <summary>
        /// The number of days to retain automated snapshots in the destination region after they are copied from the source region. Defaults to `7`.
        /// </summary>
        public readonly int? RetentionPeriod;

        [OutputConstructor]
        private ClusterSnapshotCopy(
            string destinationRegion,
            string? grantName,
            int? retentionPeriod)
        {
            DestinationRegion = destinationRegion;
            GrantName = grantName;
            RetentionPeriod = retentionPeriod;
        }
    }
    }
}
