// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GuardDuty
{
    /// <summary>
    /// Provides a resource to manage a GuardDuty filter.
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
    ///         var myFilter = new Aws.GuardDuty.Filter("myFilter", new Aws.GuardDuty.FilterArgs
    ///         {
    ///             Action = "ARCHIVE",
    ///             DetectorId = aws_guardduty_detector.Example.Id,
    ///             Rank = 1,
    ///             FindingCriteria = new Aws.GuardDuty.Inputs.FilterFindingCriteriaArgs
    ///             {
    ///                 Criterions = 
    ///                 {
    ///                     new Aws.GuardDuty.Inputs.FilterFindingCriteriaCriterionArgs
    ///                     {
    ///                         Field = "region",
    ///                         Equals = 
    ///                         {
    ///                             "eu-west-1",
    ///                         },
    ///                     },
    ///                     new Aws.GuardDuty.Inputs.FilterFindingCriteriaCriterionArgs
    ///                     {
    ///                         Field = "service.additionalInfo.threatListName",
    ///                         NotEquals = 
    ///                         {
    ///                             "some-threat",
    ///                             "another-threat",
    ///                         },
    ///                     },
    ///                     new Aws.GuardDuty.Inputs.FilterFindingCriteriaCriterionArgs
    ///                     {
    ///                         Field = "updatedAt",
    ///                         GreaterThan = "2020-01-01T00:00:00Z",
    ///                         LessThan = "2020-02-01T00:00:00Z",
    ///                     },
    ///                     new Aws.GuardDuty.Inputs.FilterFindingCriteriaCriterionArgs
    ///                     {
    ///                         Field = "severity",
    ///                         GreaterThanOrEqual = "4",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Filter : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// The ARN of the GuardDuty filter.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Description of the filter.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// ID of a GuardDuty detector, attached to your account.
        /// </summary>
        [Output("detectorId")]
        public Output<string> DetectorId { get; private set; } = null!;

        /// <summary>
        /// Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
        /// </summary>
        [Output("findingCriteria")]
        public Output<Outputs.FilterFindingCriteria> FindingCriteria { get; private set; } = null!;

        /// <summary>
        /// The name of your filter.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
        /// </summary>
        [Output("rank")]
        public Output<int> Rank { get; private set; } = null!;

        /// <summary>
        /// The tags that you want to add to the Filter resource. A tag consists of a key and a value.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Filter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Filter(string name, FilterArgs args, CustomResourceOptions? options = null)
            : base("aws:guardduty/filter:Filter", name, args ?? new FilterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Filter(string name, Input<string> id, FilterState? state = null, CustomResourceOptions? options = null)
            : base("aws:guardduty/filter:Filter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Filter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Filter Get(string name, Input<string> id, FilterState? state = null, CustomResourceOptions? options = null)
        {
            return new Filter(name, id, state, options);
        }
    }

    public sealed class FilterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// Description of the filter.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of a GuardDuty detector, attached to your account.
        /// </summary>
        [Input("detectorId", required: true)]
        public Input<string> DetectorId { get; set; } = null!;

        /// <summary>
        /// Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
        /// </summary>
        [Input("findingCriteria", required: true)]
        public Input<Inputs.FilterFindingCriteriaArgs> FindingCriteria { get; set; } = null!;

        /// <summary>
        /// The name of your filter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
        /// </summary>
        [Input("rank", required: true)]
        public Input<int> Rank { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags that you want to add to the Filter resource. A tag consists of a key and a value.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public FilterArgs()
        {
        }
    }

    public sealed class FilterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// The ARN of the GuardDuty filter.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Description of the filter.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of a GuardDuty detector, attached to your account.
        /// </summary>
        [Input("detectorId")]
        public Input<string>? DetectorId { get; set; }

        /// <summary>
        /// Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
        /// </summary>
        [Input("findingCriteria")]
        public Input<Inputs.FilterFindingCriteriaGetArgs>? FindingCriteria { get; set; }

        /// <summary>
        /// The name of your filter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
        /// </summary>
        [Input("rank")]
        public Input<int>? Rank { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags that you want to add to the Filter resource. A tag consists of a key and a value.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public FilterState()
        {
        }
    }
}
