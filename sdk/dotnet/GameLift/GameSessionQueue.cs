// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GameLift
{
    /// <summary>
    /// Provides an Gamelift Game Session Queue resource.
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
    ///         var test = new Aws.GameLift.GameSessionQueue("test", new Aws.GameLift.GameSessionQueueArgs
    ///         {
    ///             Destinations = 
    ///             {
    ///                 aws_gamelift_fleet.Us_west_2_fleet.Arn,
    ///                 aws_gamelift_fleet.Eu_central_1_fleet.Arn,
    ///             },
    ///             PlayerLatencyPolicies = 
    ///             {
    ///                 new Aws.GameLift.Inputs.GameSessionQueuePlayerLatencyPolicyArgs
    ///                 {
    ///                     MaximumIndividualPlayerLatencyMilliseconds = 100,
    ///                     PolicyDurationSeconds = 5,
    ///                 },
    ///                 new Aws.GameLift.Inputs.GameSessionQueuePlayerLatencyPolicyArgs
    ///                 {
    ///                     MaximumIndividualPlayerLatencyMilliseconds = 200,
    ///                 },
    ///             },
    ///             TimeoutInSeconds = 60,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Gamelift Game Session Queues can be imported by their `name`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:gamelift/gameSessionQueue:GameSessionQueue example example
    /// ```
    /// </summary>
    [AwsResourceType("aws:gamelift/gameSessionQueue:GameSessionQueue")]
    public partial class GameSessionQueue : Pulumi.CustomResource
    {
        /// <summary>
        /// Game Session Queue ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// List of fleet/alias ARNs used by session queue for placing game sessions.
        /// </summary>
        [Output("destinations")]
        public Output<ImmutableArray<string>> Destinations { get; private set; } = null!;

        /// <summary>
        /// Name of the session queue.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// One or more policies used to choose fleet based on player latency. See below.
        /// </summary>
        [Output("playerLatencyPolicies")]
        public Output<ImmutableArray<Outputs.GameSessionQueuePlayerLatencyPolicy>> PlayerLatencyPolicies { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Maximum time a game session request can remain in the queue.
        /// </summary>
        [Output("timeoutInSeconds")]
        public Output<int?> TimeoutInSeconds { get; private set; } = null!;


        /// <summary>
        /// Create a GameSessionQueue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GameSessionQueue(string name, GameSessionQueueArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:gamelift/gameSessionQueue:GameSessionQueue", name, args ?? new GameSessionQueueArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GameSessionQueue(string name, Input<string> id, GameSessionQueueState? state = null, CustomResourceOptions? options = null)
            : base("aws:gamelift/gameSessionQueue:GameSessionQueue", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GameSessionQueue resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GameSessionQueue Get(string name, Input<string> id, GameSessionQueueState? state = null, CustomResourceOptions? options = null)
        {
            return new GameSessionQueue(name, id, state, options);
        }
    }

    public sealed class GameSessionQueueArgs : Pulumi.ResourceArgs
    {
        [Input("destinations")]
        private InputList<string>? _destinations;

        /// <summary>
        /// List of fleet/alias ARNs used by session queue for placing game sessions.
        /// </summary>
        public InputList<string> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<string>());
            set => _destinations = value;
        }

        /// <summary>
        /// Name of the session queue.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("playerLatencyPolicies")]
        private InputList<Inputs.GameSessionQueuePlayerLatencyPolicyArgs>? _playerLatencyPolicies;

        /// <summary>
        /// One or more policies used to choose fleet based on player latency. See below.
        /// </summary>
        public InputList<Inputs.GameSessionQueuePlayerLatencyPolicyArgs> PlayerLatencyPolicies
        {
            get => _playerLatencyPolicies ?? (_playerLatencyPolicies = new InputList<Inputs.GameSessionQueuePlayerLatencyPolicyArgs>());
            set => _playerLatencyPolicies = value;
        }

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

        /// <summary>
        /// Maximum time a game session request can remain in the queue.
        /// </summary>
        [Input("timeoutInSeconds")]
        public Input<int>? TimeoutInSeconds { get; set; }

        public GameSessionQueueArgs()
        {
        }
    }

    public sealed class GameSessionQueueState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Game Session Queue ARN.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("destinations")]
        private InputList<string>? _destinations;

        /// <summary>
        /// List of fleet/alias ARNs used by session queue for placing game sessions.
        /// </summary>
        public InputList<string> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<string>());
            set => _destinations = value;
        }

        /// <summary>
        /// Name of the session queue.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("playerLatencyPolicies")]
        private InputList<Inputs.GameSessionQueuePlayerLatencyPolicyGetArgs>? _playerLatencyPolicies;

        /// <summary>
        /// One or more policies used to choose fleet based on player latency. See below.
        /// </summary>
        public InputList<Inputs.GameSessionQueuePlayerLatencyPolicyGetArgs> PlayerLatencyPolicies
        {
            get => _playerLatencyPolicies ?? (_playerLatencyPolicies = new InputList<Inputs.GameSessionQueuePlayerLatencyPolicyGetArgs>());
            set => _playerLatencyPolicies = value;
        }

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
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// Maximum time a game session request can remain in the queue.
        /// </summary>
        [Input("timeoutInSeconds")]
        public Input<int>? TimeoutInSeconds { get; set; }

        public GameSessionQueueState()
        {
        }
    }
}
