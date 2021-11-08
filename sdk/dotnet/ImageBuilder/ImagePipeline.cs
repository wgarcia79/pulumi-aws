// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ImageBuilder
{
    /// <summary>
    /// Manages an Image Builder Image Pipeline.
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
    ///         var example = new Aws.ImageBuilder.ImagePipeline("example", new Aws.ImageBuilder.ImagePipelineArgs
    ///         {
    ///             ImageRecipeArn = aws_imagebuilder_image_recipe.Example.Arn,
    ///             InfrastructureConfigurationArn = aws_imagebuilder_infrastructure_configuration.Example.Arn,
    ///             Schedule = new Aws.ImageBuilder.Inputs.ImagePipelineScheduleArgs
    ///             {
    ///                 ScheduleExpression = "cron(0 0 * * ? *)",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_imagebuilder_image_pipeline` resources can be imported using the Amazon Resource Name (ARN), e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:imagebuilder/imagePipeline:ImagePipeline example arn:aws:imagebuilder:us-east-1:123456789012:image-pipeline/example
    /// ```
    /// </summary>
    [AwsResourceType("aws:imagebuilder/imagePipeline:ImagePipeline")]
    public partial class ImagePipeline : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the image pipeline.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Date the image pipeline was created.
        /// </summary>
        [Output("dateCreated")]
        public Output<string> DateCreated { get; private set; } = null!;

        /// <summary>
        /// Date the image pipeline was last run.
        /// </summary>
        [Output("dateLastRun")]
        public Output<string> DateLastRun { get; private set; } = null!;

        /// <summary>
        /// Date the image pipeline will run next.
        /// </summary>
        [Output("dateNextRun")]
        public Output<string> DateNextRun { get; private set; } = null!;

        /// <summary>
        /// Date the image pipeline was updated.
        /// </summary>
        [Output("dateUpdated")]
        public Output<string> DateUpdated { get; private set; } = null!;

        /// <summary>
        /// Description of the image pipeline.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
        /// </summary>
        [Output("distributionConfigurationArn")]
        public Output<string?> DistributionConfigurationArn { get; private set; } = null!;

        /// <summary>
        /// Whether additional information about the image being created is collected. Defaults to `true`.
        /// </summary>
        [Output("enhancedImageMetadataEnabled")]
        public Output<bool?> EnhancedImageMetadataEnabled { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
        /// </summary>
        [Output("imageRecipeArn")]
        public Output<string> ImageRecipeArn { get; private set; } = null!;

        /// <summary>
        /// Configuration block with image tests configuration. Detailed below.
        /// </summary>
        [Output("imageTestsConfiguration")]
        public Output<Outputs.ImagePipelineImageTestsConfiguration> ImageTestsConfiguration { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
        /// </summary>
        [Output("infrastructureConfigurationArn")]
        public Output<string> InfrastructureConfigurationArn { get; private set; } = null!;

        /// <summary>
        /// Name of the image pipeline.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Platform of the image pipeline.
        /// </summary>
        [Output("platform")]
        public Output<string> Platform { get; private set; } = null!;

        /// <summary>
        /// Configuration block with schedule settings. Detailed below.
        /// </summary>
        [Output("schedule")]
        public Output<Outputs.ImagePipelineSchedule?> Schedule { get; private set; } = null!;

        /// <summary>
        /// Status of the image pipeline. Valid values are `DISABLED` and `ENABLED`. Defaults to `ENABLED`.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags for the image pipeline. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a ImagePipeline resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ImagePipeline(string name, ImagePipelineArgs args, CustomResourceOptions? options = null)
            : base("aws:imagebuilder/imagePipeline:ImagePipeline", name, args ?? new ImagePipelineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ImagePipeline(string name, Input<string> id, ImagePipelineState? state = null, CustomResourceOptions? options = null)
            : base("aws:imagebuilder/imagePipeline:ImagePipeline", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ImagePipeline resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ImagePipeline Get(string name, Input<string> id, ImagePipelineState? state = null, CustomResourceOptions? options = null)
        {
            return new ImagePipeline(name, id, state, options);
        }
    }

    public sealed class ImagePipelineArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the image pipeline.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
        /// </summary>
        [Input("distributionConfigurationArn")]
        public Input<string>? DistributionConfigurationArn { get; set; }

        /// <summary>
        /// Whether additional information about the image being created is collected. Defaults to `true`.
        /// </summary>
        [Input("enhancedImageMetadataEnabled")]
        public Input<bool>? EnhancedImageMetadataEnabled { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
        /// </summary>
        [Input("imageRecipeArn", required: true)]
        public Input<string> ImageRecipeArn { get; set; } = null!;

        /// <summary>
        /// Configuration block with image tests configuration. Detailed below.
        /// </summary>
        [Input("imageTestsConfiguration")]
        public Input<Inputs.ImagePipelineImageTestsConfigurationArgs>? ImageTestsConfiguration { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
        /// </summary>
        [Input("infrastructureConfigurationArn", required: true)]
        public Input<string> InfrastructureConfigurationArn { get; set; } = null!;

        /// <summary>
        /// Name of the image pipeline.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configuration block with schedule settings. Detailed below.
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.ImagePipelineScheduleArgs>? Schedule { get; set; }

        /// <summary>
        /// Status of the image pipeline. Valid values are `DISABLED` and `ENABLED`. Defaults to `ENABLED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags for the image pipeline. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ImagePipelineArgs()
        {
        }
    }

    public sealed class ImagePipelineState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the image pipeline.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Date the image pipeline was created.
        /// </summary>
        [Input("dateCreated")]
        public Input<string>? DateCreated { get; set; }

        /// <summary>
        /// Date the image pipeline was last run.
        /// </summary>
        [Input("dateLastRun")]
        public Input<string>? DateLastRun { get; set; }

        /// <summary>
        /// Date the image pipeline will run next.
        /// </summary>
        [Input("dateNextRun")]
        public Input<string>? DateNextRun { get; set; }

        /// <summary>
        /// Date the image pipeline was updated.
        /// </summary>
        [Input("dateUpdated")]
        public Input<string>? DateUpdated { get; set; }

        /// <summary>
        /// Description of the image pipeline.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
        /// </summary>
        [Input("distributionConfigurationArn")]
        public Input<string>? DistributionConfigurationArn { get; set; }

        /// <summary>
        /// Whether additional information about the image being created is collected. Defaults to `true`.
        /// </summary>
        [Input("enhancedImageMetadataEnabled")]
        public Input<bool>? EnhancedImageMetadataEnabled { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
        /// </summary>
        [Input("imageRecipeArn")]
        public Input<string>? ImageRecipeArn { get; set; }

        /// <summary>
        /// Configuration block with image tests configuration. Detailed below.
        /// </summary>
        [Input("imageTestsConfiguration")]
        public Input<Inputs.ImagePipelineImageTestsConfigurationGetArgs>? ImageTestsConfiguration { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
        /// </summary>
        [Input("infrastructureConfigurationArn")]
        public Input<string>? InfrastructureConfigurationArn { get; set; }

        /// <summary>
        /// Name of the image pipeline.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Platform of the image pipeline.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// Configuration block with schedule settings. Detailed below.
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.ImagePipelineScheduleGetArgs>? Schedule { get; set; }

        /// <summary>
        /// Status of the image pipeline. Valid values are `DISABLED` and `ENABLED`. Defaults to `ENABLED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags for the image pipeline. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public ImagePipelineState()
        {
        }
    }
}
