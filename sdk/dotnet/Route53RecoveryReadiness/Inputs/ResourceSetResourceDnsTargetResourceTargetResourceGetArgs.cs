// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53RecoveryReadiness.Inputs
{

    public sealed class ResourceSetResourceDnsTargetResourceTargetResourceGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// NLB resource a DNS Target Resource points to. Required if `r53_resource` is not set.
        /// </summary>
        [Input("nlbResource")]
        public Input<Inputs.ResourceSetResourceDnsTargetResourceTargetResourceNlbResourceGetArgs>? NlbResource { get; set; }

        /// <summary>
        /// Route53 resource a DNS Target Resource record points to.
        /// </summary>
        [Input("r53Resource")]
        public Input<Inputs.ResourceSetResourceDnsTargetResourceTargetResourceR53ResourceGetArgs>? R53Resource { get; set; }

        public ResourceSetResourceDnsTargetResourceTargetResourceGetArgs()
        {
        }
    }
}
