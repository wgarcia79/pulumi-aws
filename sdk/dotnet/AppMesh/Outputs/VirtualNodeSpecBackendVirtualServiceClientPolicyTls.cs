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
    public sealed class VirtualNodeSpecBackendVirtualServiceClientPolicyTls
    {
        /// <summary>
        /// The listener's TLS certificate.
        /// </summary>
        public readonly Outputs.VirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificate? Certificate;
        /// <summary>
        /// Whether the policy is enforced. Default is `true`.
        /// </summary>
        public readonly bool? Enforce;
        /// <summary>
        /// One or more ports that the policy is enforced for.
        /// </summary>
        public readonly ImmutableArray<int> Ports;
        /// <summary>
        /// The listener's Transport Layer Security (TLS) validation context.
        /// </summary>
        public readonly Outputs.VirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidation Validation;

        [OutputConstructor]
        private VirtualNodeSpecBackendVirtualServiceClientPolicyTls(
            Outputs.VirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificate? certificate,

            bool? enforce,

            ImmutableArray<int> ports,

            Outputs.VirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidation validation)
        {
            Certificate = certificate;
            Enforce = enforce;
            Ports = ports;
            Validation = validation;
        }
    }
}
