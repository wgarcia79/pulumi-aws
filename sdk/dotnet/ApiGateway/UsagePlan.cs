// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    /// <summary>
    /// Provides an API Gateway Usage Plan.
    /// 
    /// ## Import
    /// 
    /// AWS API Gateway Usage Plan can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:apigateway/usagePlan:UsagePlan myusageplan &lt;usage_plan_id&gt;
    /// ```
    /// </summary>
    [AwsResourceType("aws:apigateway/usagePlan:UsagePlan")]
    public partial class UsagePlan : Pulumi.CustomResource
    {
        /// <summary>
        /// The associated API stages of the usage plan.
        /// </summary>
        [Output("apiStages")]
        public Output<ImmutableArray<Outputs.UsagePlanApiStage>> ApiStages { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The description of a usage plan.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the usage plan.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
        /// </summary>
        [Output("productCode")]
        public Output<string?> ProductCode { get; private set; } = null!;

        /// <summary>
        /// The quota settings of the usage plan.
        /// </summary>
        [Output("quotaSettings")]
        public Output<Outputs.UsagePlanQuotaSettings?> QuotaSettings { get; private set; } = null!;

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
        /// The throttling limits of the usage plan.
        /// </summary>
        [Output("throttleSettings")]
        public Output<Outputs.UsagePlanThrottleSettings?> ThrottleSettings { get; private set; } = null!;


        /// <summary>
        /// Create a UsagePlan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UsagePlan(string name, UsagePlanArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/usagePlan:UsagePlan", name, args ?? new UsagePlanArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UsagePlan(string name, Input<string> id, UsagePlanState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/usagePlan:UsagePlan", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UsagePlan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UsagePlan Get(string name, Input<string> id, UsagePlanState? state = null, CustomResourceOptions? options = null)
        {
            return new UsagePlan(name, id, state, options);
        }
    }

    public sealed class UsagePlanArgs : Pulumi.ResourceArgs
    {
        [Input("apiStages")]
        private InputList<Inputs.UsagePlanApiStageArgs>? _apiStages;

        /// <summary>
        /// The associated API stages of the usage plan.
        /// </summary>
        public InputList<Inputs.UsagePlanApiStageArgs> ApiStages
        {
            get => _apiStages ?? (_apiStages = new InputList<Inputs.UsagePlanApiStageArgs>());
            set => _apiStages = value;
        }

        /// <summary>
        /// The description of a usage plan.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the usage plan.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
        /// </summary>
        [Input("productCode")]
        public Input<string>? ProductCode { get; set; }

        /// <summary>
        /// The quota settings of the usage plan.
        /// </summary>
        [Input("quotaSettings")]
        public Input<Inputs.UsagePlanQuotaSettingsArgs>? QuotaSettings { get; set; }

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
        /// The throttling limits of the usage plan.
        /// </summary>
        [Input("throttleSettings")]
        public Input<Inputs.UsagePlanThrottleSettingsArgs>? ThrottleSettings { get; set; }

        public UsagePlanArgs()
        {
        }
    }

    public sealed class UsagePlanState : Pulumi.ResourceArgs
    {
        [Input("apiStages")]
        private InputList<Inputs.UsagePlanApiStageGetArgs>? _apiStages;

        /// <summary>
        /// The associated API stages of the usage plan.
        /// </summary>
        public InputList<Inputs.UsagePlanApiStageGetArgs> ApiStages
        {
            get => _apiStages ?? (_apiStages = new InputList<Inputs.UsagePlanApiStageGetArgs>());
            set => _apiStages = value;
        }

        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The description of a usage plan.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the usage plan.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
        /// </summary>
        [Input("productCode")]
        public Input<string>? ProductCode { get; set; }

        /// <summary>
        /// The quota settings of the usage plan.
        /// </summary>
        [Input("quotaSettings")]
        public Input<Inputs.UsagePlanQuotaSettingsGetArgs>? QuotaSettings { get; set; }

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
        /// The throttling limits of the usage plan.
        /// </summary>
        [Input("throttleSettings")]
        public Input<Inputs.UsagePlanThrottleSettingsGetArgs>? ThrottleSettings { get; set; }

        public UsagePlanState()
        {
        }
    }
}
