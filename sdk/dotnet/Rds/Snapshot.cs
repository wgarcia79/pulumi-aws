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
    /// Manages an RDS database instance snapshot. For managing RDS database cluster snapshots, see the `aws.rds.ClusterSnapshot` resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bar = new Aws.Rds.Instance("bar", new Aws.Rds.InstanceArgs
    ///         {
    ///             AllocatedStorage = 10,
    ///             Engine = "mysql",
    ///             EngineVersion = "5.6.21",
    ///             InstanceClass = "db.t2.micro",
    ///             Name = "baz",
    ///             Password = "barbarbarbar",
    ///             Username = "foo",
    ///             MaintenanceWindow = "Fri:09:00-Fri:09:30",
    ///             BackupRetentionPeriod = 0,
    ///             ParameterGroupName = "default.mysql5.6",
    ///         });
    ///         var test = new Aws.Rds.Snapshot("test", new Aws.Rds.SnapshotArgs
    ///         {
    ///             DbInstanceIdentifier = bar.Id,
    ///             DbSnapshotIdentifier = "testsnapshot1234",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_db_snapshot` can be imported by using the snapshot identifier, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:rds/snapshot:Snapshot example my-snapshot
    /// ```
    /// </summary>
    [AwsResourceType("aws:rds/snapshot:Snapshot")]
    public partial class Snapshot : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the allocated storage size in gigabytes (GB).
        /// </summary>
        [Output("allocatedStorage")]
        public Output<int> AllocatedStorage { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// The DB Instance Identifier from which to take the snapshot.
        /// </summary>
        [Output("dbInstanceIdentifier")]
        public Output<string> DbInstanceIdentifier { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the DB snapshot.
        /// </summary>
        [Output("dbSnapshotArn")]
        public Output<string> DbSnapshotArn { get; private set; } = null!;

        /// <summary>
        /// The Identifier for the snapshot.
        /// </summary>
        [Output("dbSnapshotIdentifier")]
        public Output<string> DbSnapshotIdentifier { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the DB snapshot is encrypted.
        /// </summary>
        [Output("encrypted")]
        public Output<bool> Encrypted { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the database engine.
        /// </summary>
        [Output("engine")]
        public Output<string> Engine { get; private set; } = null!;

        /// <summary>
        /// Specifies the version of the database engine.
        /// </summary>
        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
        /// </summary>
        [Output("iops")]
        public Output<int> Iops { get; private set; } = null!;

        /// <summary>
        /// The ARN for the KMS encryption key.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// License model information for the restored DB instance.
        /// </summary>
        [Output("licenseModel")]
        public Output<string> LicenseModel { get; private set; } = null!;

        /// <summary>
        /// Provides the option group name for the DB snapshot.
        /// </summary>
        [Output("optionGroupName")]
        public Output<string> OptionGroupName { get; private set; } = null!;

        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        [Output("snapshotType")]
        public Output<string> SnapshotType { get; private set; } = null!;

        /// <summary>
        /// The DB snapshot Arn that the DB snapshot was copied from. It only has value in case of cross customer or cross region copy.
        /// </summary>
        [Output("sourceDbSnapshotIdentifier")]
        public Output<string> SourceDbSnapshotIdentifier { get; private set; } = null!;

        /// <summary>
        /// The region that the DB snapshot was created in or copied from.
        /// </summary>
        [Output("sourceRegion")]
        public Output<string> SourceRegion { get; private set; } = null!;

        /// <summary>
        /// Specifies the status of this DB snapshot.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the storage type associated with DB snapshot.
        /// </summary>
        [Output("storageType")]
        public Output<string> StorageType { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Provides the VPC ID associated with the DB snapshot.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Snapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Snapshot(string name, SnapshotArgs args, CustomResourceOptions? options = null)
            : base("aws:rds/snapshot:Snapshot", name, args ?? new SnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Snapshot(string name, Input<string> id, SnapshotState? state = null, CustomResourceOptions? options = null)
            : base("aws:rds/snapshot:Snapshot", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Snapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Snapshot Get(string name, Input<string> id, SnapshotState? state = null, CustomResourceOptions? options = null)
        {
            return new Snapshot(name, id, state, options);
        }
    }

    public sealed class SnapshotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DB Instance Identifier from which to take the snapshot.
        /// </summary>
        [Input("dbInstanceIdentifier", required: true)]
        public Input<string> DbInstanceIdentifier { get; set; } = null!;

        /// <summary>
        /// The Identifier for the snapshot.
        /// </summary>
        [Input("dbSnapshotIdentifier", required: true)]
        public Input<string> DbSnapshotIdentifier { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public SnapshotArgs()
        {
        }
    }

    public sealed class SnapshotState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the allocated storage size in gigabytes (GB).
        /// </summary>
        [Input("allocatedStorage")]
        public Input<int>? AllocatedStorage { get; set; }

        /// <summary>
        /// Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The DB Instance Identifier from which to take the snapshot.
        /// </summary>
        [Input("dbInstanceIdentifier")]
        public Input<string>? DbInstanceIdentifier { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) for the DB snapshot.
        /// </summary>
        [Input("dbSnapshotArn")]
        public Input<string>? DbSnapshotArn { get; set; }

        /// <summary>
        /// The Identifier for the snapshot.
        /// </summary>
        [Input("dbSnapshotIdentifier")]
        public Input<string>? DbSnapshotIdentifier { get; set; }

        /// <summary>
        /// Specifies whether the DB snapshot is encrypted.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// Specifies the name of the database engine.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// Specifies the version of the database engine.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
        /// </summary>
        [Input("iops")]
        public Input<int>? Iops { get; set; }

        /// <summary>
        /// The ARN for the KMS encryption key.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// License model information for the restored DB instance.
        /// </summary>
        [Input("licenseModel")]
        public Input<string>? LicenseModel { get; set; }

        /// <summary>
        /// Provides the option group name for the DB snapshot.
        /// </summary>
        [Input("optionGroupName")]
        public Input<string>? OptionGroupName { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("snapshotType")]
        public Input<string>? SnapshotType { get; set; }

        /// <summary>
        /// The DB snapshot Arn that the DB snapshot was copied from. It only has value in case of cross customer or cross region copy.
        /// </summary>
        [Input("sourceDbSnapshotIdentifier")]
        public Input<string>? SourceDbSnapshotIdentifier { get; set; }

        /// <summary>
        /// The region that the DB snapshot was created in or copied from.
        /// </summary>
        [Input("sourceRegion")]
        public Input<string>? SourceRegion { get; set; }

        /// <summary>
        /// Specifies the status of this DB snapshot.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the storage type associated with DB snapshot.
        /// </summary>
        [Input("storageType")]
        public Input<string>? StorageType { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// Provides the VPC ID associated with the DB snapshot.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public SnapshotState()
        {
        }
    }
}
