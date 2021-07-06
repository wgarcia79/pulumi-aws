// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Pinpoint
{
    /// <summary>
    /// Provides a Pinpoint App resource.
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
    ///         var example = new Aws.Pinpoint.App("example", new Aws.Pinpoint.AppArgs
    ///         {
    ///             Limits = new Aws.Pinpoint.Inputs.AppLimitsArgs
    ///             {
    ///                 MaximumDuration = 600,
    ///             },
    ///             QuietTime = new Aws.Pinpoint.Inputs.AppQuietTimeArgs
    ///             {
    ///                 End = "06:00",
    ///                 Start = "00:00",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Pinpoint App can be imported using the `application-id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:pinpoint/app:App name application-id
    /// ```
    /// </summary>
    [AwsResourceType("aws:pinpoint/app:App")]
    public partial class App : Pulumi.CustomResource
    {
        /// <summary>
        /// The Application ID of the Pinpoint App.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the PinPoint Application
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Specifies settings for invoking an AWS Lambda function that customizes a segment for a campaign
        /// </summary>
        [Output("campaignHook")]
        public Output<Outputs.AppCampaignHook?> CampaignHook { get; private set; } = null!;

        /// <summary>
        /// The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
        /// </summary>
        [Output("limits")]
        public Output<Outputs.AppLimits?> Limits { get; private set; } = null!;

        /// <summary>
        /// The application name. By default generated by this provider
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Pinpoint application. Conflicts with `name`
        /// </summary>
        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// The default quiet time for the app. Each campaign for this app sends no messages during this time unless the campaign overrides the default with a quiet time of its own
        /// </summary>
        [Output("quietTime")]
        public Output<Outputs.AppQuietTime?> QuietTime { get; private set; } = null!;

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
        /// Create a App resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public App(string name, AppArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:pinpoint/app:App", name, args ?? new AppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private App(string name, Input<string> id, AppState? state = null, CustomResourceOptions? options = null)
            : base("aws:pinpoint/app:App", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing App resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static App Get(string name, Input<string> id, AppState? state = null, CustomResourceOptions? options = null)
        {
            return new App(name, id, state, options);
        }
    }

    public sealed class AppArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies settings for invoking an AWS Lambda function that customizes a segment for a campaign
        /// </summary>
        [Input("campaignHook")]
        public Input<Inputs.AppCampaignHookArgs>? CampaignHook { get; set; }

        /// <summary>
        /// The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
        /// </summary>
        [Input("limits")]
        public Input<Inputs.AppLimitsArgs>? Limits { get; set; }

        /// <summary>
        /// The application name. By default generated by this provider
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Pinpoint application. Conflicts with `name`
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The default quiet time for the app. Each campaign for this app sends no messages during this time unless the campaign overrides the default with a quiet time of its own
        /// </summary>
        [Input("quietTime")]
        public Input<Inputs.AppQuietTimeArgs>? QuietTime { get; set; }

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

        public AppArgs()
        {
        }
    }

    public sealed class AppState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Application ID of the Pinpoint App.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the PinPoint Application
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Specifies settings for invoking an AWS Lambda function that customizes a segment for a campaign
        /// </summary>
        [Input("campaignHook")]
        public Input<Inputs.AppCampaignHookGetArgs>? CampaignHook { get; set; }

        /// <summary>
        /// The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
        /// </summary>
        [Input("limits")]
        public Input<Inputs.AppLimitsGetArgs>? Limits { get; set; }

        /// <summary>
        /// The application name. By default generated by this provider
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Pinpoint application. Conflicts with `name`
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The default quiet time for the app. Each campaign for this app sends no messages during this time unless the campaign overrides the default with a quiet time of its own
        /// </summary>
        [Input("quietTime")]
        public Input<Inputs.AppQuietTimeGetArgs>? QuietTime { get; set; }

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

        public AppState()
        {
        }
    }
}
