// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3Outposts.Inputs
{

    public sealed class EndpointNetworkInterfaceGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of the Elastic Network Interface (ENI).
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        public EndpointNetworkInterfaceGetArgs()
        {
        }
    }
}
