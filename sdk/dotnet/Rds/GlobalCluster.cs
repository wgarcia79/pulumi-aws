// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds
{
    /// <summary>
    /// Manages an RDS Global Cluster, which is an Aurora global database spread across multiple regions. The global database contains a single primary cluster with read-write capability, and a read-only secondary cluster that receives data from the primary cluster through high-speed replication performed by the Aurora storage subsystem.
    /// 
    /// More information about Aurora global databases can be found in the [Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database-creating).
    /// 
    /// ## Example Usage
    /// ### New MySQL Global Cluster
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Rds.GlobalCluster("example", new Aws.Rds.GlobalClusterArgs
    ///         {
    ///             GlobalClusterIdentifier = "global-test",
    ///             Engine = "aurora",
    ///             EngineVersion = "5.6.mysql_aurora.1.22.2",
    ///             DatabaseName = "example_db",
    ///         });
    ///         var primaryCluster = new Aws.Rds.Cluster("primaryCluster", new Aws.Rds.ClusterArgs
    ///         {
    ///             Engine = example.Engine,
    ///             EngineVersion = example.EngineVersion,
    ///             ClusterIdentifier = "test-primary-cluster",
    ///             MasterUsername = "username",
    ///             MasterPassword = "somepass123",
    ///             DatabaseName = "example_db",
    ///             GlobalClusterIdentifier = example.Id,
    ///             DbSubnetGroupName = "default",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = aws.Primary,
    ///         });
    ///         var primaryClusterInstance = new Aws.Rds.ClusterInstance("primaryClusterInstance", new Aws.Rds.ClusterInstanceArgs
    ///         {
    ///             Engine = example.Engine,
    ///             EngineVersion = example.EngineVersion,
    ///             Identifier = "test-primary-cluster-instance",
    ///             ClusterIdentifier = primaryCluster.Id,
    ///             InstanceClass = "db.r4.large",
    ///             DbSubnetGroupName = "default",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = aws.Primary,
    ///         });
    ///         var secondaryCluster = new Aws.Rds.Cluster("secondaryCluster", new Aws.Rds.ClusterArgs
    ///         {
    ///             Engine = example.Engine,
    ///             EngineVersion = example.EngineVersion,
    ///             ClusterIdentifier = "test-secondary-cluster",
    ///             GlobalClusterIdentifier = example.Id,
    ///             DbSubnetGroupName = "default",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = aws.Secondary,
    ///         });
    ///         var secondaryClusterInstance = new Aws.Rds.ClusterInstance("secondaryClusterInstance", new Aws.Rds.ClusterInstanceArgs
    ///         {
    ///             Engine = example.Engine,
    ///             EngineVersion = example.EngineVersion,
    ///             Identifier = "test-secondary-cluster-instance",
    ///             ClusterIdentifier = secondaryCluster.Id,
    ///             InstanceClass = "db.r4.large",
    ///             DbSubnetGroupName = "default",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = aws.Secondary,
    ///             DependsOn = 
    ///             {
    ///                 primaryClusterInstance,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### New PostgreSQL Global Cluster
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var primary = new Aws.Provider("primary", new Aws.ProviderArgs
    ///         {
    ///             Region = "us-east-2",
    ///         });
    ///         var secondary = new Aws.Provider("secondary", new Aws.ProviderArgs
    ///         {
    ///             Region = "us-east-1",
    ///         });
    ///         var example = new Aws.Rds.GlobalCluster("example", new Aws.Rds.GlobalClusterArgs
    ///         {
    ///             GlobalClusterIdentifier = "global-test",
    ///             Engine = "aurora-postgresql",
    ///             EngineVersion = "11.9",
    ///             DatabaseName = "example_db",
    ///         });
    ///         var primaryCluster = new Aws.Rds.Cluster("primaryCluster", new Aws.Rds.ClusterArgs
    ///         {
    ///             Engine = example.Engine,
    ///             EngineVersion = example.EngineVersion,
    ///             ClusterIdentifier = "test-primary-cluster",
    ///             MasterUsername = "username",
    ///             MasterPassword = "somepass123",
    ///             DatabaseName = "example_db",
    ///             GlobalClusterIdentifier = example.Id,
    ///             DbSubnetGroupName = "default",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = aws.Primary,
    ///         });
    ///         var primaryClusterInstance = new Aws.Rds.ClusterInstance("primaryClusterInstance", new Aws.Rds.ClusterInstanceArgs
    ///         {
    ///             Engine = example.Engine,
    ///             EngineVersion = example.EngineVersion,
    ///             Identifier = "test-primary-cluster-instance",
    ///             ClusterIdentifier = primaryCluster.Id,
    ///             InstanceClass = "db.r4.large",
    ///             DbSubnetGroupName = "default",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = aws.Primary,
    ///         });
    ///         var secondaryCluster = new Aws.Rds.Cluster("secondaryCluster", new Aws.Rds.ClusterArgs
    ///         {
    ///             Engine = example.Engine,
    ///             EngineVersion = example.EngineVersion,
    ///             ClusterIdentifier = "test-secondary-cluster",
    ///             GlobalClusterIdentifier = example.Id,
    ///             SkipFinalSnapshot = true,
    ///             DbSubnetGroupName = "default",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = aws.Secondary,
    ///             DependsOn = 
    ///             {
    ///                 primaryClusterInstance,
    ///             },
    ///         });
    ///         var secondaryClusterInstance = new Aws.Rds.ClusterInstance("secondaryClusterInstance", new Aws.Rds.ClusterInstanceArgs
    ///         {
    ///             Engine = example.Engine,
    ///             EngineVersion = example.EngineVersion,
    ///             Identifier = "test-secondary-cluster-instance",
    ///             ClusterIdentifier = secondaryCluster.Id,
    ///             InstanceClass = "db.r4.large",
    ///             DbSubnetGroupName = "default",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = aws.Secondary,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### New Global Cluster From Existing DB Cluster
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // ... other configuration ...
    ///         var exampleCluster = new Aws.Rds.Cluster("exampleCluster", new Aws.Rds.ClusterArgs
    ///         {
    ///         });
    ///         var exampleGlobalCluster = new Aws.Rds.GlobalCluster("exampleGlobalCluster", new Aws.Rds.GlobalClusterArgs
    ///         {
    ///             ForceDestroy = true,
    ///             GlobalClusterIdentifier = "example",
    ///             SourceDbClusterIdentifier = exampleCluster.Arn,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_rds_global_cluster` can be imported by using the RDS Global Cluster identifier, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:rds/globalCluster:GlobalCluster example example
    /// ```
    /// 
    ///  Certain resource arguments, like `source_db_cluster_identifier`, do not have an API method for reading the information after creation. If the argument is set in the provider configuration on an imported resource, the provider will always show a difference. To workaround this behavior, either omit the argument from the the provider configuration or use `ignore_changes` to hide the difference, e.g. terraform resource "aws_rds_global_cluster" "example" {
    /// 
    /// # ... other configuration ...
    /// 
    /// # There is no API for reading source_db_cluster_identifier
    /// 
    ///  lifecycle {
    /// 
    ///  ignore_changes = [source_db_cluster_identifier]
    /// 
    ///  } }
    /// </summary>
    [AwsResourceType("aws:rds/globalCluster:GlobalCluster")]
    public partial class GlobalCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// RDS Global Cluster Amazon Resource Name (ARN)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Name for an automatically created database on cluster creation.
        /// </summary>
        [Output("databaseName")]
        public Output<string?> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool?> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `source_db_cluster_identifier`.
        /// </summary>
        [Output("engine")]
        public Output<string> Engine { get; private set; } = null!;

        /// <summary>
        /// Engine version of the Aurora global database.
        /// * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
        /// </summary>
        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// Enable to remove DB Cluster members from Global Cluster on destroy. Required with `source_db_cluster_identifier`.
        /// </summary>
        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        /// <summary>
        /// The global cluster identifier.
        /// </summary>
        [Output("globalClusterIdentifier")]
        public Output<string> GlobalClusterIdentifier { get; private set; } = null!;

        /// <summary>
        /// Set of objects containing Global Cluster members.
        /// </summary>
        [Output("globalClusterMembers")]
        public Output<ImmutableArray<Outputs.GlobalClusterGlobalClusterMember>> GlobalClusterMembers { get; private set; } = null!;

        /// <summary>
        /// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
        /// </summary>
        [Output("globalClusterResourceId")]
        public Output<string> GlobalClusterResourceId { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
        /// </summary>
        [Output("sourceDbClusterIdentifier")]
        public Output<string> SourceDbClusterIdentifier { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the DB cluster is encrypted. The default is `false` unless `source_db_cluster_identifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Output("storageEncrypted")]
        public Output<bool> StorageEncrypted { get; private set; } = null!;


        /// <summary>
        /// Create a GlobalCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GlobalCluster(string name, GlobalClusterArgs args, CustomResourceOptions? options = null)
            : base("aws:rds/globalCluster:GlobalCluster", name, args ?? new GlobalClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GlobalCluster(string name, Input<string> id, GlobalClusterState? state = null, CustomResourceOptions? options = null)
            : base("aws:rds/globalCluster:GlobalCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GlobalCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GlobalCluster Get(string name, Input<string> id, GlobalClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new GlobalCluster(name, id, state, options);
        }
    }

    public sealed class GlobalClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name for an automatically created database on cluster creation.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `source_db_cluster_identifier`.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// Engine version of the Aurora global database.
        /// * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// Enable to remove DB Cluster members from Global Cluster on destroy. Required with `source_db_cluster_identifier`.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// The global cluster identifier.
        /// </summary>
        [Input("globalClusterIdentifier", required: true)]
        public Input<string> GlobalClusterIdentifier { get; set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
        /// </summary>
        [Input("sourceDbClusterIdentifier")]
        public Input<string>? SourceDbClusterIdentifier { get; set; }

        /// <summary>
        /// Specifies whether the DB cluster is encrypted. The default is `false` unless `source_db_cluster_identifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Input("storageEncrypted")]
        public Input<bool>? StorageEncrypted { get; set; }

        public GlobalClusterArgs()
        {
        }
    }

    public sealed class GlobalClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// RDS Global Cluster Amazon Resource Name (ARN)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Name for an automatically created database on cluster creation.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `source_db_cluster_identifier`.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// Engine version of the Aurora global database.
        /// * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// Enable to remove DB Cluster members from Global Cluster on destroy. Required with `source_db_cluster_identifier`.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// The global cluster identifier.
        /// </summary>
        [Input("globalClusterIdentifier")]
        public Input<string>? GlobalClusterIdentifier { get; set; }

        [Input("globalClusterMembers")]
        private InputList<Inputs.GlobalClusterGlobalClusterMemberGetArgs>? _globalClusterMembers;

        /// <summary>
        /// Set of objects containing Global Cluster members.
        /// </summary>
        public InputList<Inputs.GlobalClusterGlobalClusterMemberGetArgs> GlobalClusterMembers
        {
            get => _globalClusterMembers ?? (_globalClusterMembers = new InputList<Inputs.GlobalClusterGlobalClusterMemberGetArgs>());
            set => _globalClusterMembers = value;
        }

        /// <summary>
        /// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
        /// </summary>
        [Input("globalClusterResourceId")]
        public Input<string>? GlobalClusterResourceId { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
        /// </summary>
        [Input("sourceDbClusterIdentifier")]
        public Input<string>? SourceDbClusterIdentifier { get; set; }

        /// <summary>
        /// Specifies whether the DB cluster is encrypted. The default is `false` unless `source_db_cluster_identifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Input("storageEncrypted")]
        public Input<bool>? StorageEncrypted { get; set; }

        public GlobalClusterState()
        {
        }
    }
}
