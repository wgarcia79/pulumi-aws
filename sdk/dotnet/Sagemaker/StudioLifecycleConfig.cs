// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker
{
    /// <summary>
    /// Provides a Sagemaker Studio Lifecycle Config resource.
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// Sagemaker Code Studio Lifecycle Configs can be imported using the `studio_lifecycle_config_name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:sagemaker/studioLifecycleConfig:StudioLifecycleConfig example example
    /// ```
    /// </summary>
    [AwsResourceType("aws:sagemaker/studioLifecycleConfig:StudioLifecycleConfig")]
    public partial class StudioLifecycleConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) assigned by AWS to this Studio Lifecycle Config.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The App type that the Lifecycle Configuration is attached to. Valid values are `JupyterServer` and `KernelGateway`.
        /// </summary>
        [Output("studioLifecycleConfigAppType")]
        public Output<string> StudioLifecycleConfigAppType { get; private set; } = null!;

        /// <summary>
        /// The content of your Studio Lifecycle Configuration script. This content must be base64 encoded.
        /// </summary>
        [Output("studioLifecycleConfigContent")]
        public Output<string> StudioLifecycleConfigContent { get; private set; } = null!;

        /// <summary>
        /// The name of the Studio Lifecycle Configuration to create.
        /// </summary>
        [Output("studioLifecycleConfigName")]
        public Output<string> StudioLifecycleConfigName { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a StudioLifecycleConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StudioLifecycleConfig(string name, StudioLifecycleConfigArgs args, CustomResourceOptions? options = null)
            : base("aws:sagemaker/studioLifecycleConfig:StudioLifecycleConfig", name, args ?? new StudioLifecycleConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StudioLifecycleConfig(string name, Input<string> id, StudioLifecycleConfigState? state = null, CustomResourceOptions? options = null)
            : base("aws:sagemaker/studioLifecycleConfig:StudioLifecycleConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StudioLifecycleConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StudioLifecycleConfig Get(string name, Input<string> id, StudioLifecycleConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new StudioLifecycleConfig(name, id, state, options);
        }
    }

    public sealed class StudioLifecycleConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The App type that the Lifecycle Configuration is attached to. Valid values are `JupyterServer` and `KernelGateway`.
        /// </summary>
        [Input("studioLifecycleConfigAppType", required: true)]
        public Input<string> StudioLifecycleConfigAppType { get; set; } = null!;

        /// <summary>
        /// The content of your Studio Lifecycle Configuration script. This content must be base64 encoded.
        /// </summary>
        [Input("studioLifecycleConfigContent", required: true)]
        public Input<string> StudioLifecycleConfigContent { get; set; } = null!;

        /// <summary>
        /// The name of the Studio Lifecycle Configuration to create.
        /// </summary>
        [Input("studioLifecycleConfigName", required: true)]
        public Input<string> StudioLifecycleConfigName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public StudioLifecycleConfigArgs()
        {
        }
    }

    public sealed class StudioLifecycleConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) assigned by AWS to this Studio Lifecycle Config.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The App type that the Lifecycle Configuration is attached to. Valid values are `JupyterServer` and `KernelGateway`.
        /// </summary>
        [Input("studioLifecycleConfigAppType")]
        public Input<string>? StudioLifecycleConfigAppType { get; set; }

        /// <summary>
        /// The content of your Studio Lifecycle Configuration script. This content must be base64 encoded.
        /// </summary>
        [Input("studioLifecycleConfigContent")]
        public Input<string>? StudioLifecycleConfigContent { get; set; }

        /// <summary>
        /// The name of the Studio Lifecycle Configuration to create.
        /// </summary>
        [Input("studioLifecycleConfigName")]
        public Input<string>? StudioLifecycleConfigName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public StudioLifecycleConfigState()
        {
        }
    }
}
