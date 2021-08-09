// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DirectConnect
{
    /// <summary>
    /// Provides a Direct Connect LAG. Connections can be added to the LAG via the `aws.directconnect.Connection` and `aws.directconnect.ConnectionAssociation` resources.
    /// 
    /// &gt; *NOTE:* When creating a LAG, Direct Connect requires creating a Connection. This provider will remove this unmanaged connection during resource creation.
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
    ///         var hoge = new Aws.DirectConnect.LinkAggregationGroup("hoge", new Aws.DirectConnect.LinkAggregationGroupArgs
    ///         {
    ///             ConnectionsBandwidth = "1Gbps",
    ///             ForceDestroy = true,
    ///             Location = "EqDC2",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Direct Connect LAGs can be imported using the `lag id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:directconnect/linkAggregationGroup:LinkAggregationGroup test_lag dxlag-fgnsp5rq
    /// ```
    /// </summary>
    [AwsResourceType("aws:directconnect/linkAggregationGroup:LinkAggregationGroup")]
    public partial class LinkAggregationGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the LAG.
        /// * `jumbo_frame_capable` -Indicates whether jumbo frames (9001 MTU) are supported.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
        /// </summary>
        [Output("connectionsBandwidth")]
        public Output<string> ConnectionsBandwidth { get; private set; } = null!;

        /// <summary>
        /// A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
        /// </summary>
        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).
        /// </summary>
        [Output("hasLogicalRedundancy")]
        public Output<string> HasLogicalRedundancy { get; private set; } = null!;

        [Output("jumboFrameCapable")]
        public Output<bool> JumboFrameCapable { get; private set; } = null!;

        /// <summary>
        /// The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the LAG.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a LinkAggregationGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LinkAggregationGroup(string name, LinkAggregationGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:directconnect/linkAggregationGroup:LinkAggregationGroup", name, args ?? new LinkAggregationGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LinkAggregationGroup(string name, Input<string> id, LinkAggregationGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:directconnect/linkAggregationGroup:LinkAggregationGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LinkAggregationGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LinkAggregationGroup Get(string name, Input<string> id, LinkAggregationGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new LinkAggregationGroup(name, id, state, options);
        }
    }

    public sealed class LinkAggregationGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
        /// </summary>
        [Input("connectionsBandwidth", required: true)]
        public Input<string> ConnectionsBandwidth { get; set; } = null!;

        /// <summary>
        /// A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the LAG.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public LinkAggregationGroupArgs()
        {
        }
    }

    public sealed class LinkAggregationGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the LAG.
        /// * `jumbo_frame_capable` -Indicates whether jumbo frames (9001 MTU) are supported.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
        /// </summary>
        [Input("connectionsBandwidth")]
        public Input<string>? ConnectionsBandwidth { get; set; }

        /// <summary>
        /// A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).
        /// </summary>
        [Input("hasLogicalRedundancy")]
        public Input<string>? HasLogicalRedundancy { get; set; }

        [Input("jumboFrameCapable")]
        public Input<bool>? JumboFrameCapable { get; set; }

        /// <summary>
        /// The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the LAG.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public LinkAggregationGroupState()
        {
        }
    }
}
