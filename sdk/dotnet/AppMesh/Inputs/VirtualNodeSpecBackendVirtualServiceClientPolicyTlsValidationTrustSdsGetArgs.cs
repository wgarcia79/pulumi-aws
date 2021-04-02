// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh.Inputs
{

    public sealed class VirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustSdsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the secret for a virtual node's Transport Layer Security (TLS) Secret Discovery Service validation context trust.
        /// </summary>
        [Input("secretName", required: true)]
        public Input<string> SecretName { get; set; } = null!;

        public VirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustSdsGetArgs()
        {
        }
    }
}
