// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Outputs
{

    [OutputType]
    public sealed class FlowLogDestinationOptions
    {
        /// <summary>
        /// The format for the flow log. Default value: `plain-text`. Valid values: `plain-text`, `parquet`.
        /// </summary>
        public readonly string? FileFormat;
        /// <summary>
        /// Indicates whether to use Hive-compatible prefixes for flow logs stored in Amazon S3. Default value: `false`.
        /// </summary>
        public readonly bool? HiveCompatiblePartitions;
        /// <summary>
        /// Indicates whether to partition the flow log per hour. This reduces the cost and response time for queries. Default value: `false`.
        /// </summary>
        public readonly bool? PerHourPartition;

        [OutputConstructor]
        private FlowLogDestinationOptions(
            string? fileFormat,

            bool? hiveCompatiblePartitions,

            bool? perHourPartition)
        {
            FileFormat = fileFormat;
            HiveCompatiblePartitions = hiveCompatiblePartitions;
            PerHourPartition = perHourPartition;
        }
    }
}
