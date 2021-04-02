// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh.Inputs
{

    public sealed class VirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The TLS validation context trust for an AWS Certificate Manager (ACM) certificate.
        /// </summary>
        [Input("acm")]
        public Input<Inputs.VirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustAcmGetArgs>? Acm { get; set; }

        /// <summary>
        /// The TLS validation context trust for a local file certificate.
        /// </summary>
        [Input("file")]
        public Input<Inputs.VirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustFileGetArgs>? File { get; set; }

        /// <summary>
        /// The TLS validation context trust for a [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.
        /// </summary>
        [Input("sds")]
        public Input<Inputs.VirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustSdsGetArgs>? Sds { get; set; }

        public VirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustGetArgs()
        {
        }
    }
}
