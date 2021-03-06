// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling.Outputs
{

    [OutputType]
    public sealed class GetAmiIdsFilterResult
    {
        /// <summary>
        /// The name of the filter. The valid values are: `auto-scaling-group`, `key`, `value`, and `propagate-at-launch`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of the filter.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetAmiIdsFilterResult(
            string name,

            ImmutableArray<string> values)
        {
            Name = name;
            Values = values;
        }
    }
}
