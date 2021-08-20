// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Chime
{
    /// <summary>
    /// Creates an Amazon Chime Voice Connector group under the administrator's AWS account. You can associate Amazon Chime Voice Connectors with the Amazon Chime Voice Connector group by including VoiceConnectorItems in the request.
    /// 
    /// You can include Amazon Chime Voice Connectors from different AWS Regions in your group. This creates a fault tolerant mechanism for fallback in case of availability events.
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
    ///         var vc1 = new Aws.Chime.VoiceConnector("vc1", new Aws.Chime.VoiceConnectorArgs
    ///         {
    ///             RequireEncryption = true,
    ///             AwsRegion = "us-east-1",
    ///         });
    ///         var vc2 = new Aws.Chime.VoiceConnector("vc2", new Aws.Chime.VoiceConnectorArgs
    ///         {
    ///             RequireEncryption = true,
    ///             AwsRegion = "us-west-2",
    ///         });
    ///         var @group = new Aws.Chime.VoiceConnectorGroup("group", new Aws.Chime.VoiceConnectorGroupArgs
    ///         {
    ///             Connectors = 
    ///             {
    ///                 new Aws.Chime.Inputs.VoiceConnectorGroupConnectorArgs
    ///                 {
    ///                     VoiceConnectorId = vc1.Id,
    ///                     Priority = 1,
    ///                 },
    ///                 new Aws.Chime.Inputs.VoiceConnectorGroupConnectorArgs
    ///                 {
    ///                     VoiceConnectorId = vc2.Id,
    ///                     Priority = 3,
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
    /// Configuration Recorder can be imported using the name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:chime/voiceConnectorGroup:VoiceConnectorGroup default example
    /// ```
    /// </summary>
    [AwsResourceType("aws:chime/voiceConnectorGroup:VoiceConnectorGroup")]
    public partial class VoiceConnectorGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Chime Voice Connectors to route inbound calls to.
        /// </summary>
        [Output("connectors")]
        public Output<ImmutableArray<Outputs.VoiceConnectorGroupConnector>> Connectors { get; private set; } = null!;

        /// <summary>
        /// The name of the Amazon Chime Voice Connector group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a VoiceConnectorGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VoiceConnectorGroup(string name, VoiceConnectorGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:chime/voiceConnectorGroup:VoiceConnectorGroup", name, args ?? new VoiceConnectorGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VoiceConnectorGroup(string name, Input<string> id, VoiceConnectorGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:chime/voiceConnectorGroup:VoiceConnectorGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VoiceConnectorGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VoiceConnectorGroup Get(string name, Input<string> id, VoiceConnectorGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new VoiceConnectorGroup(name, id, state, options);
        }
    }

    public sealed class VoiceConnectorGroupArgs : Pulumi.ResourceArgs
    {
        [Input("connectors")]
        private InputList<Inputs.VoiceConnectorGroupConnectorArgs>? _connectors;

        /// <summary>
        /// The Amazon Chime Voice Connectors to route inbound calls to.
        /// </summary>
        public InputList<Inputs.VoiceConnectorGroupConnectorArgs> Connectors
        {
            get => _connectors ?? (_connectors = new InputList<Inputs.VoiceConnectorGroupConnectorArgs>());
            set => _connectors = value;
        }

        /// <summary>
        /// The name of the Amazon Chime Voice Connector group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public VoiceConnectorGroupArgs()
        {
        }
    }

    public sealed class VoiceConnectorGroupState : Pulumi.ResourceArgs
    {
        [Input("connectors")]
        private InputList<Inputs.VoiceConnectorGroupConnectorGetArgs>? _connectors;

        /// <summary>
        /// The Amazon Chime Voice Connectors to route inbound calls to.
        /// </summary>
        public InputList<Inputs.VoiceConnectorGroupConnectorGetArgs> Connectors
        {
            get => _connectors ?? (_connectors = new InputList<Inputs.VoiceConnectorGroupConnectorGetArgs>());
            set => _connectors = value;
        }

        /// <summary>
        /// The name of the Amazon Chime Voice Connector group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public VoiceConnectorGroupState()
        {
        }
    }
}
