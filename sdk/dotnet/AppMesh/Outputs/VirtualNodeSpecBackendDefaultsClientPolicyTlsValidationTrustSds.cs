// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh.Outputs
{

    [OutputType]
    public sealed class VirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustSds
    {
        /// <summary>
        /// The name of the secret for a virtual node's Transport Layer Security (TLS) Secret Discovery Service validation context trust.
        /// </summary>
        public readonly string SecretName;

        [OutputConstructor]
        private VirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustSds(string secretName)
        {
            SecretName = secretName;
        }
    }
}
