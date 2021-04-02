// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LB.Outputs
{

    [OutputType]
    public sealed class TargetGroupStickiness
    {
        /// <summary>
        /// Only used when the type is `lb_cookie`. The time period, in seconds, during which requests from a client should be routed to the same target. After this time period expires, the load balancer-generated cookie is considered stale. The range is 1 second to 1 week (604800 seconds). The default value is 1 day (86400 seconds).
        /// </summary>
        public readonly int? CookieDuration;
        /// <summary>
        /// Whether to enable `stickiness`. Default is `true`.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Type of sticky sessions. The only current possible values are `lb_cookie` for ALBs and `source_ip` for NLBs.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private TargetGroupStickiness(
            int? cookieDuration,

            bool? enabled,

            string type)
        {
            CookieDuration = cookieDuration;
            Enabled = enabled;
            Type = type;
        }
    }
}
