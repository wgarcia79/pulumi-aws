// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling.Inputs
{

    public sealed class PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedLoadMetricSpecificationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes a scaling metric for a predictive scaling policy. Valid values are `ASGAverageCPUUtilization`, `ASGAverageNetworkIn`, `ASGAverageNetworkOut`, or `ALBRequestCountPerTarget`.
        /// </summary>
        [Input("predefinedMetricType", required: true)]
        public Input<string> PredefinedMetricType { get; set; } = null!;

        /// <summary>
        /// A label that uniquely identifies a specific Application Load Balancer target group from which to determine the request count served by your Auto Scaling group.
        /// </summary>
        [Input("resourceLabel", required: true)]
        public Input<string> ResourceLabel { get; set; } = null!;

        public PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedLoadMetricSpecificationGetArgs()
        {
        }
    }
}
