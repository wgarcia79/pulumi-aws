// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecr
{
    /// <summary>
    /// Provides an Elastic Container Registry Replication Configuration.
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
    ///         var current = Output.Create(Aws.GetCallerIdentity.InvokeAsync());
    ///         var exampleRegions = Output.Create(Aws.GetRegions.InvokeAsync());
    ///         var exampleReplicationConfiguration = new Aws.Ecr.ReplicationConfiguration("exampleReplicationConfiguration", new Aws.Ecr.ReplicationConfigurationArgs
    ///         {
    ///             ReplicationConfiguration = new Aws.Ecr.Inputs.ReplicationConfigurationReplicationConfigurationArgs
    ///             {
    ///                 Rule = new Aws.Ecr.Inputs.ReplicationConfigurationReplicationConfigurationRuleArgs
    ///                 {
    ///                     Destinations = 
    ///                     {
    ///                         new Aws.Ecr.Inputs.ReplicationConfigurationReplicationConfigurationRuleDestinationArgs
    ///                         {
    ///                             Region = exampleRegions.Apply(exampleRegions =&gt; exampleRegions.Names[0]),
    ///                             RegistryId = current.Apply(current =&gt; current.AccountId),
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECR Replication Configuration can be imported using the `registry_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:ecr/replicationConfiguration:ReplicationConfiguration service 012345678912
    /// ```
    /// </summary>
    [AwsResourceType("aws:ecr/replicationConfiguration:ReplicationConfiguration")]
    public partial class ReplicationConfiguration : Pulumi.CustomResource
    {
        /// <summary>
        /// The account ID of the destination registry to replicate to.
        /// </summary>
        [Output("registryId")]
        public Output<string> RegistryId { get; private set; } = null!;

        /// <summary>
        /// Replication configuration for a registry. See Replication Configuration.
        /// </summary>
        [Output("replicationConfiguration")]
        public Output<Outputs.ReplicationConfigurationReplicationConfiguration?> ReplicationConfigurationDetails { get; private set; } = null!;


        /// <summary>
        /// Create a ReplicationConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplicationConfiguration(string name, ReplicationConfigurationArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:ecr/replicationConfiguration:ReplicationConfiguration", name, args ?? new ReplicationConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReplicationConfiguration(string name, Input<string> id, ReplicationConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("aws:ecr/replicationConfiguration:ReplicationConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReplicationConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplicationConfiguration Get(string name, Input<string> id, ReplicationConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new ReplicationConfiguration(name, id, state, options);
        }
    }

    public sealed class ReplicationConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Replication configuration for a registry. See Replication Configuration.
        /// </summary>
        [Input("replicationConfiguration")]
        public Input<Inputs.ReplicationConfigurationReplicationConfigurationArgs>? ReplicationConfigurationDetails { get; set; }

        public ReplicationConfigurationArgs()
        {
        }
    }

    public sealed class ReplicationConfigurationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account ID of the destination registry to replicate to.
        /// </summary>
        [Input("registryId")]
        public Input<string>? RegistryId { get; set; }

        /// <summary>
        /// Replication configuration for a registry. See Replication Configuration.
        /// </summary>
        [Input("replicationConfiguration")]
        public Input<Inputs.ReplicationConfigurationReplicationConfigurationGetArgs>? ReplicationConfigurationDetails { get; set; }

        public ReplicationConfigurationState()
        {
        }
    }
}
