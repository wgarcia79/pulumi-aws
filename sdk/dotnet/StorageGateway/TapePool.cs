// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.StorageGateway
{
    /// <summary>
    /// Manages an AWS Storage Gateway Tape Pool.
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
    ///         var example = new Aws.StorageGateway.TapePool("example", new Aws.StorageGateway.TapePoolArgs
    ///         {
    ///             PoolName = "example",
    ///             StorageClass = "GLACIER",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class TapePool : Pulumi.CustomResource
    {
        /// <summary>
        /// Volume Amazon Resource Name (ARN), e.g. `aws_storagegateway_tape_pool.example arn:aws:storagegateway:us-east-1:123456789012:tapepool/pool-12345678`.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name of the new custom tape pool.
        /// </summary>
        [Output("poolName")]
        public Output<string> PoolName { get; private set; } = null!;

        /// <summary>
        /// Tape retention lock time is set in days. Tape retention lock can be enabled for up to 100 years (36,500 days). Default value is 0.
        /// </summary>
        [Output("retentionLockTimeInDays")]
        public Output<int?> RetentionLockTimeInDays { get; private set; } = null!;

        /// <summary>
        /// Tape retention lock can be configured in two modes. When configured in governance mode, AWS accounts with specific IAM permissions are authorized to remove the tape retention lock from archived virtual tapes. When configured in compliance mode, the tape retention lock cannot be removed by any user, including the root AWS account. Possible values are `COMPLIANCE`, `GOVERNANCE`, and `NONE`. Default value is `NONE`.
        /// </summary>
        [Output("retentionLockType")]
        public Output<string?> RetentionLockType { get; private set; } = null!;

        /// <summary>
        /// The storage class that is associated with the new custom pool. When you use your backup application to eject the tape, the tape is archived directly into the storage class that corresponds to the pool. Possible values are `DEEP_ARCHIVE` or `GLACIER`.
        /// </summary>
        [Output("storageClass")]
        public Output<string> StorageClass { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a TapePool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TapePool(string name, TapePoolArgs args, CustomResourceOptions? options = null)
            : base("aws:storagegateway/tapePool:TapePool", name, args ?? new TapePoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TapePool(string name, Input<string> id, TapePoolState? state = null, CustomResourceOptions? options = null)
            : base("aws:storagegateway/tapePool:TapePool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TapePool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TapePool Get(string name, Input<string> id, TapePoolState? state = null, CustomResourceOptions? options = null)
        {
            return new TapePool(name, id, state, options);
        }
    }

    public sealed class TapePoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the new custom tape pool.
        /// </summary>
        [Input("poolName", required: true)]
        public Input<string> PoolName { get; set; } = null!;

        /// <summary>
        /// Tape retention lock time is set in days. Tape retention lock can be enabled for up to 100 years (36,500 days). Default value is 0.
        /// </summary>
        [Input("retentionLockTimeInDays")]
        public Input<int>? RetentionLockTimeInDays { get; set; }

        /// <summary>
        /// Tape retention lock can be configured in two modes. When configured in governance mode, AWS accounts with specific IAM permissions are authorized to remove the tape retention lock from archived virtual tapes. When configured in compliance mode, the tape retention lock cannot be removed by any user, including the root AWS account. Possible values are `COMPLIANCE`, `GOVERNANCE`, and `NONE`. Default value is `NONE`.
        /// </summary>
        [Input("retentionLockType")]
        public Input<string>? RetentionLockType { get; set; }

        /// <summary>
        /// The storage class that is associated with the new custom pool. When you use your backup application to eject the tape, the tape is archived directly into the storage class that corresponds to the pool. Possible values are `DEEP_ARCHIVE` or `GLACIER`.
        /// </summary>
        [Input("storageClass", required: true)]
        public Input<string> StorageClass { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public TapePoolArgs()
        {
        }
    }

    public sealed class TapePoolState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Volume Amazon Resource Name (ARN), e.g. `aws_storagegateway_tape_pool.example arn:aws:storagegateway:us-east-1:123456789012:tapepool/pool-12345678`.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name of the new custom tape pool.
        /// </summary>
        [Input("poolName")]
        public Input<string>? PoolName { get; set; }

        /// <summary>
        /// Tape retention lock time is set in days. Tape retention lock can be enabled for up to 100 years (36,500 days). Default value is 0.
        /// </summary>
        [Input("retentionLockTimeInDays")]
        public Input<int>? RetentionLockTimeInDays { get; set; }

        /// <summary>
        /// Tape retention lock can be configured in two modes. When configured in governance mode, AWS accounts with specific IAM permissions are authorized to remove the tape retention lock from archived virtual tapes. When configured in compliance mode, the tape retention lock cannot be removed by any user, including the root AWS account. Possible values are `COMPLIANCE`, `GOVERNANCE`, and `NONE`. Default value is `NONE`.
        /// </summary>
        [Input("retentionLockType")]
        public Input<string>? RetentionLockType { get; set; }

        /// <summary>
        /// The storage class that is associated with the new custom pool. When you use your backup application to eject the tape, the tape is archived directly into the storage class that corresponds to the pool. Possible values are `DEEP_ARCHIVE` or `GLACIER`.
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public TapePoolState()
        {
        }
    }
}
