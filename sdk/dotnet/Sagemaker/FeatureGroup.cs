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
    /// Provides a SageMaker Feature Group resource.
    /// 
    /// ## Example Usage
    /// 
    /// Basic usage:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Sagemaker.FeatureGroup("example", new Aws.Sagemaker.FeatureGroupArgs
    ///         {
    ///             FeatureGroupName = "example",
    ///             RecordIdentifierFeatureName = "example",
    ///             EventTimeFeatureName = "example",
    ///             RoleArn = aws_iam_role.Test.Arn,
    ///             FeatureDefinitions = 
    ///             {
    ///                 new Aws.Sagemaker.Inputs.FeatureGroupFeatureDefinitionArgs
    ///                 {
    ///                     FeatureName = "example",
    ///                     FeatureType = "String",
    ///                 },
    ///             },
    ///             OnlineStoreConfig = new Aws.Sagemaker.Inputs.FeatureGroupOnlineStoreConfigArgs
    ///             {
    ///                 EnableOnlineStore = true,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Feature Groups can be imported using the `name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:sagemaker/featureGroup:FeatureGroup test_feature_group feature_group-foo
    /// ```
    /// </summary>
    [AwsResourceType("aws:sagemaker/featureGroup:FeatureGroup")]
    public partial class FeatureGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) assigned by AWS to this feature_group.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A free-form description of a Feature Group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the feature that stores the EventTime of a Record in a Feature Group.
        /// </summary>
        [Output("eventTimeFeatureName")]
        public Output<string> EventTimeFeatureName { get; private set; } = null!;

        /// <summary>
        /// A list of Feature names and types. See Feature Definition Below.
        /// </summary>
        [Output("featureDefinitions")]
        public Output<ImmutableArray<Outputs.FeatureGroupFeatureDefinition>> FeatureDefinitions { get; private set; } = null!;

        /// <summary>
        /// The name of the Feature Group. The name must be unique within an AWS Region in an AWS account.
        /// </summary>
        [Output("featureGroupName")]
        public Output<string> FeatureGroupName { get; private set; } = null!;

        /// <summary>
        /// The Offline Feature Store Configuration. See Offline Store Config Below.
        /// </summary>
        [Output("offlineStoreConfig")]
        public Output<Outputs.FeatureGroupOfflineStoreConfig?> OfflineStoreConfig { get; private set; } = null!;

        /// <summary>
        /// The Online Feature Store Configuration. See Online Store Config Below.
        /// </summary>
        [Output("onlineStoreConfig")]
        public Output<Outputs.FeatureGroupOnlineStoreConfig?> OnlineStoreConfig { get; private set; } = null!;

        /// <summary>
        /// The name of the Feature whose value uniquely identifies a Record defined in the Feature Store. Only the latest record per identifier value will be stored in the Online Store.
        /// </summary>
        [Output("recordIdentifierFeatureName")]
        public Output<string> RecordIdentifierFeatureName { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM execution role used to persist data into the Offline Store if an `offline_store_config` is provided.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// Map of resource tags for the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a FeatureGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FeatureGroup(string name, FeatureGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:sagemaker/featureGroup:FeatureGroup", name, args ?? new FeatureGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FeatureGroup(string name, Input<string> id, FeatureGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:sagemaker/featureGroup:FeatureGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FeatureGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FeatureGroup Get(string name, Input<string> id, FeatureGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new FeatureGroup(name, id, state, options);
        }
    }

    public sealed class FeatureGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A free-form description of a Feature Group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the feature that stores the EventTime of a Record in a Feature Group.
        /// </summary>
        [Input("eventTimeFeatureName", required: true)]
        public Input<string> EventTimeFeatureName { get; set; } = null!;

        [Input("featureDefinitions", required: true)]
        private InputList<Inputs.FeatureGroupFeatureDefinitionArgs>? _featureDefinitions;

        /// <summary>
        /// A list of Feature names and types. See Feature Definition Below.
        /// </summary>
        public InputList<Inputs.FeatureGroupFeatureDefinitionArgs> FeatureDefinitions
        {
            get => _featureDefinitions ?? (_featureDefinitions = new InputList<Inputs.FeatureGroupFeatureDefinitionArgs>());
            set => _featureDefinitions = value;
        }

        /// <summary>
        /// The name of the Feature Group. The name must be unique within an AWS Region in an AWS account.
        /// </summary>
        [Input("featureGroupName", required: true)]
        public Input<string> FeatureGroupName { get; set; } = null!;

        /// <summary>
        /// The Offline Feature Store Configuration. See Offline Store Config Below.
        /// </summary>
        [Input("offlineStoreConfig")]
        public Input<Inputs.FeatureGroupOfflineStoreConfigArgs>? OfflineStoreConfig { get; set; }

        /// <summary>
        /// The Online Feature Store Configuration. See Online Store Config Below.
        /// </summary>
        [Input("onlineStoreConfig")]
        public Input<Inputs.FeatureGroupOnlineStoreConfigArgs>? OnlineStoreConfig { get; set; }

        /// <summary>
        /// The name of the Feature whose value uniquely identifies a Record defined in the Feature Store. Only the latest record per identifier value will be stored in the Online Store.
        /// </summary>
        [Input("recordIdentifierFeatureName", required: true)]
        public Input<string> RecordIdentifierFeatureName { get; set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM execution role used to persist data into the Offline Store if an `offline_store_config` is provided.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of resource tags for the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public FeatureGroupArgs()
        {
        }
    }

    public sealed class FeatureGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) assigned by AWS to this feature_group.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// A free-form description of a Feature Group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the feature that stores the EventTime of a Record in a Feature Group.
        /// </summary>
        [Input("eventTimeFeatureName")]
        public Input<string>? EventTimeFeatureName { get; set; }

        [Input("featureDefinitions")]
        private InputList<Inputs.FeatureGroupFeatureDefinitionGetArgs>? _featureDefinitions;

        /// <summary>
        /// A list of Feature names and types. See Feature Definition Below.
        /// </summary>
        public InputList<Inputs.FeatureGroupFeatureDefinitionGetArgs> FeatureDefinitions
        {
            get => _featureDefinitions ?? (_featureDefinitions = new InputList<Inputs.FeatureGroupFeatureDefinitionGetArgs>());
            set => _featureDefinitions = value;
        }

        /// <summary>
        /// The name of the Feature Group. The name must be unique within an AWS Region in an AWS account.
        /// </summary>
        [Input("featureGroupName")]
        public Input<string>? FeatureGroupName { get; set; }

        /// <summary>
        /// The Offline Feature Store Configuration. See Offline Store Config Below.
        /// </summary>
        [Input("offlineStoreConfig")]
        public Input<Inputs.FeatureGroupOfflineStoreConfigGetArgs>? OfflineStoreConfig { get; set; }

        /// <summary>
        /// The Online Feature Store Configuration. See Online Store Config Below.
        /// </summary>
        [Input("onlineStoreConfig")]
        public Input<Inputs.FeatureGroupOnlineStoreConfigGetArgs>? OnlineStoreConfig { get; set; }

        /// <summary>
        /// The name of the Feature whose value uniquely identifies a Record defined in the Feature Store. Only the latest record per identifier value will be stored in the Online Store.
        /// </summary>
        [Input("recordIdentifierFeatureName")]
        public Input<string>? RecordIdentifierFeatureName { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM execution role used to persist data into the Offline Store if an `offline_store_config` is provided.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of resource tags for the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public FeatureGroupState()
        {
        }
    }
}
