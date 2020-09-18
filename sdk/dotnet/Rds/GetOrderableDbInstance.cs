// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds
{
    public static class GetOrderableDbInstance
    {
        /// <summary>
        /// Information about RDS orderable DB instances.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var test = Output.Create(Aws.Rds.GetOrderableDbInstance.InvokeAsync(new Aws.Rds.GetOrderableDbInstanceArgs
        ///         {
        ///             Engine = "mysql",
        ///             EngineVersion = "5.7.22",
        ///             LicenseModel = "general-public-license",
        ///             PreferredInstanceClasses = 
        ///             {
        ///                 "db.r6.xlarge",
        ///                 "db.m4.large",
        ///                 "db.t3.small",
        ///             },
        ///             StorageType = "standard",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOrderableDbInstanceResult> InvokeAsync(GetOrderableDbInstanceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrderableDbInstanceResult>("aws:rds/getOrderableDbInstance:getOrderableDbInstance", args ?? new GetOrderableDbInstanceArgs(), options.WithVersion());
    }


    public sealed class GetOrderableDbInstanceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Availability zone group.
        /// </summary>
        [Input("availabilityZoneGroup")]
        public string? AvailabilityZoneGroup { get; set; }

        /// <summary>
        /// DB engine. Engine values include `aurora`, `aurora-mysql`, `aurora-postgresql`, `docdb`, `mariadb`, `mysql`, `neptune`, `oracle-ee`, `oracle-se`, `oracle-se1`, `oracle-se2`, `postgres`, `sqlserver-ee`, `sqlserver-ex`, `sqlserver-se`, and `sqlserver-web`.
        /// </summary>
        [Input("engine", required: true)]
        public string Engine { get; set; } = null!;

        /// <summary>
        /// Version of the DB engine.
        /// </summary>
        [Input("engineVersion")]
        public string? EngineVersion { get; set; }

        /// <summary>
        /// DB instance class. Examples of classes are `db.m3.2xlarge`, `db.t2.small`, and `db.m3.medium`.
        /// </summary>
        [Input("instanceClass")]
        public string? InstanceClass { get; set; }

        /// <summary>
        /// License model. Examples of license models are `general-public-license`, `bring-your-own-license`, and `amazon-license`.
        /// </summary>
        [Input("licenseModel")]
        public string? LicenseModel { get; set; }

        [Input("preferredEngineVersions")]
        private List<string>? _preferredEngineVersions;
        public List<string> PreferredEngineVersions
        {
            get => _preferredEngineVersions ?? (_preferredEngineVersions = new List<string>());
            set => _preferredEngineVersions = value;
        }

        [Input("preferredInstanceClasses")]
        private List<string>? _preferredInstanceClasses;

        /// <summary>
        /// Ordered list of preferred RDS DB instance classes. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned.
        /// </summary>
        public List<string> PreferredInstanceClasses
        {
            get => _preferredInstanceClasses ?? (_preferredInstanceClasses = new List<string>());
            set => _preferredInstanceClasses = value;
        }

        /// <summary>
        /// Storage types. Examples of storage types are `standard`, `io1`, `gp2`, and `aurora`.
        /// </summary>
        [Input("storageType")]
        public string? StorageType { get; set; }

        /// <summary>
        /// Enable this to ensure a DB instance supports Enhanced Monitoring at intervals from 1 to 60 seconds.
        /// </summary>
        [Input("supportsEnhancedMonitoring")]
        public bool? SupportsEnhancedMonitoring { get; set; }

        /// <summary>
        /// Enable this to ensure a DB instance supports Aurora global databases with a specific combination of other DB engine attributes.
        /// </summary>
        [Input("supportsGlobalDatabases")]
        public bool? SupportsGlobalDatabases { get; set; }

        /// <summary>
        /// Enable this to ensure a DB instance supports IAM database authentication.
        /// </summary>
        [Input("supportsIamDatabaseAuthentication")]
        public bool? SupportsIamDatabaseAuthentication { get; set; }

        /// <summary>
        /// Enable this to ensure a DB instance supports provisioned IOPS.
        /// </summary>
        [Input("supportsIops")]
        public bool? SupportsIops { get; set; }

        /// <summary>
        /// Enable this to ensure a DB instance supports Kerberos Authentication.
        /// </summary>
        [Input("supportsKerberosAuthentication")]
        public bool? SupportsKerberosAuthentication { get; set; }

        /// <summary>
        /// Enable this to ensure a DB instance supports Performance Insights.
        /// </summary>
        [Input("supportsPerformanceInsights")]
        public bool? SupportsPerformanceInsights { get; set; }

        /// <summary>
        /// Enable this to ensure Amazon RDS can automatically scale storage for DB instances that use the specified DB instance class.
        /// </summary>
        [Input("supportsStorageAutoscaling")]
        public bool? SupportsStorageAutoscaling { get; set; }

        /// <summary>
        /// Enable this to ensure a DB instance supports encrypted storage.
        /// </summary>
        [Input("supportsStorageEncryption")]
        public bool? SupportsStorageEncryption { get; set; }

        /// <summary>
        /// Boolean that indicates whether to show only VPC or non-VPC offerings.
        /// </summary>
        [Input("vpc")]
        public bool? Vpc { get; set; }

        public GetOrderableDbInstanceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOrderableDbInstanceResult
    {
        public readonly string AvailabilityZoneGroup;
        /// <summary>
        /// Availability zones where the instance is available.
        /// </summary>
        public readonly ImmutableArray<string> AvailabilityZones;
        public readonly string Engine;
        public readonly string EngineVersion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceClass;
        public readonly string LicenseModel;
        /// <summary>
        /// Maximum total provisioned IOPS for a DB instance.
        /// </summary>
        public readonly int MaxIopsPerDbInstance;
        /// <summary>
        /// Maximum provisioned IOPS per GiB for a DB instance.
        /// </summary>
        public readonly double MaxIopsPerGib;
        /// <summary>
        /// Maximum storage size for a DB instance.
        /// </summary>
        public readonly int MaxStorageSize;
        /// <summary>
        /// Minimum total provisioned IOPS for a DB instance.
        /// </summary>
        public readonly int MinIopsPerDbInstance;
        /// <summary>
        /// Minimum provisioned IOPS per GiB for a DB instance.
        /// </summary>
        public readonly double MinIopsPerGib;
        /// <summary>
        /// Minimum storage size for a DB instance.
        /// </summary>
        public readonly int MinStorageSize;
        /// <summary>
        /// Whether a DB instance is Multi-AZ capable.
        /// </summary>
        public readonly bool MultiAzCapable;
        /// <summary>
        /// Whether a DB instance supports RDS on Outposts.
        /// </summary>
        public readonly bool OutpostCapable;
        public readonly ImmutableArray<string> PreferredEngineVersions;
        public readonly ImmutableArray<string> PreferredInstanceClasses;
        /// <summary>
        /// Whether a DB instance can have a read replica.
        /// </summary>
        public readonly bool ReadReplicaCapable;
        public readonly string StorageType;
        /// <summary>
        /// A list of the supported DB engine modes.
        /// </summary>
        public readonly ImmutableArray<string> SupportedEngineModes;
        public readonly bool SupportsEnhancedMonitoring;
        public readonly bool SupportsGlobalDatabases;
        public readonly bool SupportsIamDatabaseAuthentication;
        public readonly bool SupportsIops;
        public readonly bool SupportsKerberosAuthentication;
        public readonly bool SupportsPerformanceInsights;
        public readonly bool SupportsStorageAutoscaling;
        public readonly bool SupportsStorageEncryption;
        public readonly bool Vpc;

        [OutputConstructor]
        private GetOrderableDbInstanceResult(
            string availabilityZoneGroup,

            ImmutableArray<string> availabilityZones,

            string engine,

            string engineVersion,

            string id,

            string instanceClass,

            string licenseModel,

            int maxIopsPerDbInstance,

            double maxIopsPerGib,

            int maxStorageSize,

            int minIopsPerDbInstance,

            double minIopsPerGib,

            int minStorageSize,

            bool multiAzCapable,

            bool outpostCapable,

            ImmutableArray<string> preferredEngineVersions,

            ImmutableArray<string> preferredInstanceClasses,

            bool readReplicaCapable,

            string storageType,

            ImmutableArray<string> supportedEngineModes,

            bool supportsEnhancedMonitoring,

            bool supportsGlobalDatabases,

            bool supportsIamDatabaseAuthentication,

            bool supportsIops,

            bool supportsKerberosAuthentication,

            bool supportsPerformanceInsights,

            bool supportsStorageAutoscaling,

            bool supportsStorageEncryption,

            bool vpc)
        {
            AvailabilityZoneGroup = availabilityZoneGroup;
            AvailabilityZones = availabilityZones;
            Engine = engine;
            EngineVersion = engineVersion;
            Id = id;
            InstanceClass = instanceClass;
            LicenseModel = licenseModel;
            MaxIopsPerDbInstance = maxIopsPerDbInstance;
            MaxIopsPerGib = maxIopsPerGib;
            MaxStorageSize = maxStorageSize;
            MinIopsPerDbInstance = minIopsPerDbInstance;
            MinIopsPerGib = minIopsPerGib;
            MinStorageSize = minStorageSize;
            MultiAzCapable = multiAzCapable;
            OutpostCapable = outpostCapable;
            PreferredEngineVersions = preferredEngineVersions;
            PreferredInstanceClasses = preferredInstanceClasses;
            ReadReplicaCapable = readReplicaCapable;
            StorageType = storageType;
            SupportedEngineModes = supportedEngineModes;
            SupportsEnhancedMonitoring = supportsEnhancedMonitoring;
            SupportsGlobalDatabases = supportsGlobalDatabases;
            SupportsIamDatabaseAuthentication = supportsIamDatabaseAuthentication;
            SupportsIops = supportsIops;
            SupportsKerberosAuthentication = supportsKerberosAuthentication;
            SupportsPerformanceInsights = supportsPerformanceInsights;
            SupportsStorageAutoscaling = supportsStorageAutoscaling;
            SupportsStorageEncryption = supportsStorageEncryption;
            Vpc = vpc;
        }
    }
}
