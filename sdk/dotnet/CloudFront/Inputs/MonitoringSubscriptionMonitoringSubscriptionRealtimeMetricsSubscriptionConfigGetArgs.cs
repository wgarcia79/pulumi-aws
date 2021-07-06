// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront.Inputs
{

    public sealed class MonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A flag that indicates whether additional CloudWatch metrics are enabled for a given CloudFront distribution. Valid values are `Enabled` and `Disabled`. See below.
        /// </summary>
        [Input("realtimeMetricsSubscriptionStatus", required: true)]
        public Input<string> RealtimeMetricsSubscriptionStatus { get; set; } = null!;

        public MonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigGetArgs()
        {
        }
    }
}
