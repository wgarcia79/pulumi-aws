// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppStream
{
    /// <summary>
    /// Provides an AppStream stack.
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
    ///         var example = new Aws.AppStream.Stack("example", new Aws.AppStream.StackArgs
    ///         {
    ///             ApplicationSettings = new Aws.AppStream.Inputs.StackApplicationSettingsArgs
    ///             {
    ///                 Enabled = true,
    ///                 SettingsGroup = "SettingsGroup",
    ///             },
    ///             Description = "stack description",
    ///             DisplayName = "stack display name",
    ///             FeedbackUrl = "http://your-domain/feedback",
    ///             RedirectUrl = "http://your-domain/redirect",
    ///             StorageConnectors = 
    ///             {
    ///                 new Aws.AppStream.Inputs.StackStorageConnectorArgs
    ///                 {
    ///                     ConnectorType = "HOMEFOLDERS",
    ///                 },
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "TagName", "TagValue" },
    ///             },
    ///             UserSettings = 
    ///             {
    ///                 new Aws.AppStream.Inputs.StackUserSettingArgs
    ///                 {
    ///                     Action = "CLIPBOARD_COPY_FROM_LOCAL_DEVICE",
    ///                     Permission = "ENABLED",
    ///                 },
    ///                 new Aws.AppStream.Inputs.StackUserSettingArgs
    ///                 {
    ///                     Action = "CLIPBOARD_COPY_TO_LOCAL_DEVICE",
    ///                     Permission = "ENABLED",
    ///                 },
    ///                 new Aws.AppStream.Inputs.StackUserSettingArgs
    ///                 {
    ///                     Action = "FILE_UPLOAD",
    ///                     Permission = "ENABLED",
    ///                 },
    ///                 new Aws.AppStream.Inputs.StackUserSettingArgs
    ///                 {
    ///                     Action = "FILE_DOWNLOAD",
    ///                     Permission = "ENABLED",
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
    /// `aws_appstream_stack` can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:appstream/stack:Stack example stackID
    /// ```
    /// </summary>
    [AwsResourceType("aws:appstream/stack:Stack")]
    public partial class Stack : Pulumi.CustomResource
    {
        [Output("accessEndpoints")]
        public Output<ImmutableArray<Outputs.StackAccessEndpoint>> AccessEndpoints { get; private set; } = null!;

        /// <summary>
        /// Settings for application settings persistence.
        /// </summary>
        [Output("applicationSettings")]
        public Output<Outputs.StackApplicationSettings> ApplicationSettings { get; private set; } = null!;

        /// <summary>
        /// ARN of the appstream stack.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Date and time, in UTC and extended RFC 3339 format, when the stack was created.
        /// </summary>
        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

        /// <summary>
        /// Description for the AppStream stack.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Stack name to display.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
        /// </summary>
        [Output("embedHostDomains")]
        public Output<ImmutableArray<string>> EmbedHostDomains { get; private set; } = null!;

        /// <summary>
        /// URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
        /// </summary>
        [Output("feedbackUrl")]
        public Output<string> FeedbackUrl { get; private set; } = null!;

        /// <summary>
        /// Unique name for the AppStream stack.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// URL that users are redirected to after their streaming session ends.
        /// </summary>
        [Output("redirectUrl")]
        public Output<string> RedirectUrl { get; private set; } = null!;

        /// <summary>
        /// Configuration block for the storage connectors to enable. See below.
        /// </summary>
        [Output("storageConnectors")]
        public Output<ImmutableArray<Outputs.StackStorageConnector>> StorageConnectors { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Configuration block for the actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled. See below.
        /// </summary>
        [Output("userSettings")]
        public Output<ImmutableArray<Outputs.StackUserSetting>> UserSettings { get; private set; } = null!;


        /// <summary>
        /// Create a Stack resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Stack(string name, StackArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:appstream/stack:Stack", name, args ?? new StackArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Stack(string name, Input<string> id, StackState? state = null, CustomResourceOptions? options = null)
            : base("aws:appstream/stack:Stack", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Stack resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Stack Get(string name, Input<string> id, StackState? state = null, CustomResourceOptions? options = null)
        {
            return new Stack(name, id, state, options);
        }
    }

    public sealed class StackArgs : Pulumi.ResourceArgs
    {
        [Input("accessEndpoints")]
        private InputList<Inputs.StackAccessEndpointArgs>? _accessEndpoints;
        public InputList<Inputs.StackAccessEndpointArgs> AccessEndpoints
        {
            get => _accessEndpoints ?? (_accessEndpoints = new InputList<Inputs.StackAccessEndpointArgs>());
            set => _accessEndpoints = value;
        }

        /// <summary>
        /// Settings for application settings persistence.
        /// </summary>
        [Input("applicationSettings")]
        public Input<Inputs.StackApplicationSettingsArgs>? ApplicationSettings { get; set; }

        /// <summary>
        /// Description for the AppStream stack.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Stack name to display.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("embedHostDomains")]
        private InputList<string>? _embedHostDomains;

        /// <summary>
        /// Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
        /// </summary>
        public InputList<string> EmbedHostDomains
        {
            get => _embedHostDomains ?? (_embedHostDomains = new InputList<string>());
            set => _embedHostDomains = value;
        }

        /// <summary>
        /// URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
        /// </summary>
        [Input("feedbackUrl")]
        public Input<string>? FeedbackUrl { get; set; }

        /// <summary>
        /// Unique name for the AppStream stack.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// URL that users are redirected to after their streaming session ends.
        /// </summary>
        [Input("redirectUrl")]
        public Input<string>? RedirectUrl { get; set; }

        [Input("storageConnectors")]
        private InputList<Inputs.StackStorageConnectorArgs>? _storageConnectors;

        /// <summary>
        /// Configuration block for the storage connectors to enable. See below.
        /// </summary>
        public InputList<Inputs.StackStorageConnectorArgs> StorageConnectors
        {
            get => _storageConnectors ?? (_storageConnectors = new InputList<Inputs.StackStorageConnectorArgs>());
            set => _storageConnectors = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        [Input("userSettings")]
        private InputList<Inputs.StackUserSettingArgs>? _userSettings;

        /// <summary>
        /// Configuration block for the actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled. See below.
        /// </summary>
        public InputList<Inputs.StackUserSettingArgs> UserSettings
        {
            get => _userSettings ?? (_userSettings = new InputList<Inputs.StackUserSettingArgs>());
            set => _userSettings = value;
        }

        public StackArgs()
        {
        }
    }

    public sealed class StackState : Pulumi.ResourceArgs
    {
        [Input("accessEndpoints")]
        private InputList<Inputs.StackAccessEndpointGetArgs>? _accessEndpoints;
        public InputList<Inputs.StackAccessEndpointGetArgs> AccessEndpoints
        {
            get => _accessEndpoints ?? (_accessEndpoints = new InputList<Inputs.StackAccessEndpointGetArgs>());
            set => _accessEndpoints = value;
        }

        /// <summary>
        /// Settings for application settings persistence.
        /// </summary>
        [Input("applicationSettings")]
        public Input<Inputs.StackApplicationSettingsGetArgs>? ApplicationSettings { get; set; }

        /// <summary>
        /// ARN of the appstream stack.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Date and time, in UTC and extended RFC 3339 format, when the stack was created.
        /// </summary>
        [Input("createdTime")]
        public Input<string>? CreatedTime { get; set; }

        /// <summary>
        /// Description for the AppStream stack.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Stack name to display.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("embedHostDomains")]
        private InputList<string>? _embedHostDomains;

        /// <summary>
        /// Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
        /// </summary>
        public InputList<string> EmbedHostDomains
        {
            get => _embedHostDomains ?? (_embedHostDomains = new InputList<string>());
            set => _embedHostDomains = value;
        }

        /// <summary>
        /// URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
        /// </summary>
        [Input("feedbackUrl")]
        public Input<string>? FeedbackUrl { get; set; }

        /// <summary>
        /// Unique name for the AppStream stack.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// URL that users are redirected to after their streaming session ends.
        /// </summary>
        [Input("redirectUrl")]
        public Input<string>? RedirectUrl { get; set; }

        [Input("storageConnectors")]
        private InputList<Inputs.StackStorageConnectorGetArgs>? _storageConnectors;

        /// <summary>
        /// Configuration block for the storage connectors to enable. See below.
        /// </summary>
        public InputList<Inputs.StackStorageConnectorGetArgs> StorageConnectors
        {
            get => _storageConnectors ?? (_storageConnectors = new InputList<Inputs.StackStorageConnectorGetArgs>());
            set => _storageConnectors = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        [Input("userSettings")]
        private InputList<Inputs.StackUserSettingGetArgs>? _userSettings;

        /// <summary>
        /// Configuration block for the actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled. See below.
        /// </summary>
        public InputList<Inputs.StackUserSettingGetArgs> UserSettings
        {
            get => _userSettings ?? (_userSettings = new InputList<Inputs.StackUserSettingGetArgs>());
            set => _userSettings = value;
        }

        public StackState()
        {
        }
    }
}
