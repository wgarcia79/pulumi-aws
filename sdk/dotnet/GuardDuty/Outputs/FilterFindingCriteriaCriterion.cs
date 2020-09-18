// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GuardDuty.Outputs
{

    [OutputType]
    public sealed class FilterFindingCriteriaCriterion
    {
        /// <summary>
        /// List of string values to be evaluated.
        /// </summary>
        public readonly ImmutableArray<string> Equals;
        /// <summary>
        /// The name of the field to be evaluated. The full list of field names can be found in [AWS documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_filter-findings.html#filter_criteria).
        /// </summary>
        public readonly string Field;
        /// <summary>
        /// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
        /// </summary>
        public readonly string? GreaterThan;
        /// <summary>
        /// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
        /// </summary>
        public readonly string? GreaterThanOrEqual;
        /// <summary>
        /// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
        /// </summary>
        public readonly string? LessThan;
        /// <summary>
        /// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
        /// </summary>
        public readonly string? LessThanOrEqual;
        /// <summary>
        /// List of string values to be evaluated.
        /// </summary>
        public readonly ImmutableArray<string> NotEquals;

        [OutputConstructor]
        private FilterFindingCriteriaCriterion(
            ImmutableArray<string> equals,

            string field,

            string? greaterThan,

            string? greaterThanOrEqual,

            string? lessThan,

            string? lessThanOrEqual,

            ImmutableArray<string> notEquals)
        {
            Equals = equals;
            Field = field;
            GreaterThan = greaterThan;
            GreaterThanOrEqual = greaterThanOrEqual;
            LessThan = lessThan;
            LessThanOrEqual = lessThanOrEqual;
            NotEquals = notEquals;
        }
    }
}
