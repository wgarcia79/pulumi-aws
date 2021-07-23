// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3.Inputs
{

    public sealed class BucketReplicationConfigurationRuleGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether delete markers are replicated. The only valid value is `Enabled`. To disable, omit this argument. This argument is only valid with V2 replication configurations (i.e., when `filter` is used).
        /// </summary>
        [Input("deleteMarkerReplicationStatus")]
        public Input<string>? DeleteMarkerReplicationStatus { get; set; }

        /// <summary>
        /// Specifies the destination for the rule (documented below).
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.BucketReplicationConfigurationRuleDestinationGetArgs> Destination { get; set; } = null!;

        /// <summary>
        /// Filter that identifies subset of objects to which the replication rule applies (documented below).
        /// </summary>
        [Input("filter")]
        public Input<Inputs.BucketReplicationConfigurationRuleFilterGetArgs>? Filter { get; set; }

        /// <summary>
        /// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Object keyname prefix identifying one or more objects to which the rule applies. Must be less than or equal to 1024 characters in length.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// The priority associated with the rule. Priority should only be set if `filter` is configured. If not provided, defaults to `0`. Priority must be unique between multiple rules.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Specifies special object selection criteria (documented below).
        /// </summary>
        [Input("sourceSelectionCriteria")]
        public Input<Inputs.BucketReplicationConfigurationRuleSourceSelectionCriteriaGetArgs>? SourceSelectionCriteria { get; set; }

        /// <summary>
        /// The status of the rule. Either `Enabled` or `Disabled`. The rule is ignored if status is not Enabled.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        public BucketReplicationConfigurationRuleGetArgs()
        {
        }
    }
}
