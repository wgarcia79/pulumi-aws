// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Aws.Efs
{
    public static class GetFileSystem
    {
        /// <summary>
        /// Provides information about an Elastic File System (EFS) File System.
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
        ///         var config = new Config();
        ///         var fileSystemId = config.Get("fileSystemId") ?? "";
        ///         var byId = Output.Create(Aws.Efs.GetFileSystem.InvokeAsync(new Aws.Efs.GetFileSystemArgs
        ///         {
        ///             FileSystemId = fileSystemId,
        ///         }));
        ///         var byTag = Output.Create(Aws.Efs.GetFileSystem.InvokeAsync(new Aws.Efs.GetFileSystemArgs
        ///         {
        ///             Tags = 
        ///             {
        ///                 { "Environment", "dev" },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFileSystemResult> InvokeAsync(GetFileSystemArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFileSystemResult>("aws:efs/getFileSystem:getFileSystem", args ?? new GetFileSystemArgs(), options.WithVersion());

        /// <summary>
        /// Provides information about an Elastic File System (EFS) File System.
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
        ///         var config = new Config();
        ///         var fileSystemId = config.Get("fileSystemId") ?? "";
        ///         var byId = Output.Create(Aws.Efs.GetFileSystem.InvokeAsync(new Aws.Efs.GetFileSystemArgs
        ///         {
        ///             FileSystemId = fileSystemId,
        ///         }));
        ///         var byTag = Output.Create(Aws.Efs.GetFileSystem.InvokeAsync(new Aws.Efs.GetFileSystemArgs
        ///         {
        ///             Tags = 
        ///             {
        ///                 { "Environment", "dev" },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFileSystemResult> Invoke(GetFileSystemInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFileSystemResult>("aws:efs/getFileSystem:getFileSystem", args ?? new GetFileSystemInvokeArgs(), options.WithVersion());
    }


    public sealed class GetFileSystemArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Restricts the list to the file system with this creation token.
        /// </summary>
        [Input("creationToken")]
        public string? CreationToken { get; set; }

        /// <summary>
        /// The ID that identifies the file system (e.g., fs-ccfc0d65).
        /// </summary>
        [Input("fileSystemId")]
        public string? FileSystemId { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Restricts the list to the file system with these tags.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetFileSystemArgs()
        {
        }
    }

    public sealed class GetFileSystemInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Restricts the list to the file system with this creation token.
        /// </summary>
        [Input("creationToken")]
        public Input<string>? CreationToken { get; set; }

        /// <summary>
        /// The ID that identifies the file system (e.g., fs-ccfc0d65).
        /// </summary>
        [Input("fileSystemId")]
        public Input<string>? FileSystemId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Restricts the list to the file system with these tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetFileSystemInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFileSystemResult
    {
        /// <summary>
        /// Amazon Resource Name of the file system.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The identifier of the Availability Zone in which the file system's One Zone storage classes exist.
        /// </summary>
        public readonly string AvailabilityZoneId;
        /// <summary>
        /// The Availability Zone name in which the file system's One Zone storage classes exist.
        /// </summary>
        public readonly string AvailabilityZoneName;
        public readonly string CreationToken;
        /// <summary>
        /// The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
        /// </summary>
        public readonly string DnsName;
        /// <summary>
        /// Whether EFS is encrypted.
        /// </summary>
        public readonly bool Encrypted;
        public readonly string FileSystemId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ARN for the KMS encryption key.
        /// </summary>
        public readonly string KmsKeyId;
        /// <summary>
        /// A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object.
        /// </summary>
        public readonly Outputs.GetFileSystemLifecyclePolicyResult LifecyclePolicy;
        /// <summary>
        /// The file system performance mode.
        /// </summary>
        public readonly string PerformanceMode;
        /// <summary>
        /// The throughput, measured in MiB/s, that you want to provision for the file system.
        /// * `tags` -A map of tags to assign to the file system.
        /// </summary>
        public readonly double ProvisionedThroughputInMibps;
        /// <summary>
        /// The current byte count used by the file system.
        /// </summary>
        public readonly int SizeInBytes;
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Throughput mode for the file system.
        /// </summary>
        public readonly string ThroughputMode;

        [OutputConstructor]
        private GetFileSystemResult(
            string arn,

            string availabilityZoneId,

            string availabilityZoneName,

            string creationToken,

            string dnsName,

            bool encrypted,

            string fileSystemId,

            string id,

            string kmsKeyId,

            Outputs.GetFileSystemLifecyclePolicyResult lifecyclePolicy,

            string performanceMode,

            double provisionedThroughputInMibps,

            int sizeInBytes,

            ImmutableDictionary<string, string> tags,

            string throughputMode)
        {
            Arn = arn;
            AvailabilityZoneId = availabilityZoneId;
            AvailabilityZoneName = availabilityZoneName;
            CreationToken = creationToken;
            DnsName = dnsName;
            Encrypted = encrypted;
            FileSystemId = fileSystemId;
            Id = id;
            KmsKeyId = kmsKeyId;
            LifecyclePolicy = lifecyclePolicy;
            PerformanceMode = performanceMode;
            ProvisionedThroughputInMibps = provisionedThroughputInMibps;
            SizeInBytes = sizeInBytes;
            Tags = tags;
            ThroughputMode = throughputMode;
        }
    }
}
