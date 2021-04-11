// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApplicationLoadBalancing.Outputs
{

    [OutputType]
    public sealed class ListenerDefaultActionForwardTargetGroup
    {
        /// <summary>
        /// ARN of the target group.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Weight. The range is 0 to 999.
        /// </summary>
        public readonly int? Weight;

        [OutputConstructor]
        private ListenerDefaultActionForwardTargetGroup(
            string arn,

            int? weight)
        {
            Arn = arn;
            Weight = weight;
        }
    }
}
