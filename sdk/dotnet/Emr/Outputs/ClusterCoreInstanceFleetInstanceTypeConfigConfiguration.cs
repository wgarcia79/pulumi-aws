// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Emr.Outputs
{

    [OutputType]
    public sealed class ClusterCoreInstanceFleetInstanceTypeConfigConfiguration
    {
        /// <summary>
        /// Classification within a configuration.
        /// </summary>
        public readonly string? Classification;
        /// <summary>
        /// Key-Value map of Java properties that are set when the step runs. You can use these properties to pass key value pairs to your main function.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Properties;

        [OutputConstructor]
        private ClusterCoreInstanceFleetInstanceTypeConfigConfiguration(
            string? classification,

            ImmutableDictionary<string, object>? properties)
        {
            Classification = classification;
            Properties = properties;
        }
    }
}
