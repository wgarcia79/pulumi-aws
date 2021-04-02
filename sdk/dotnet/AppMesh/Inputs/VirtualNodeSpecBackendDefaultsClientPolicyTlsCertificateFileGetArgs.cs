// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh.Inputs
{

    public sealed class VirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFileGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate trust chain for a certificate stored on the file system of the mesh endpoint that the proxy is running on. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("certificateChain", required: true)]
        public Input<string> CertificateChain { get; set; } = null!;

        /// <summary>
        /// The private key for a certificate stored on the file system of the virtual node that the proxy is running on. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("privateKey", required: true)]
        public Input<string> PrivateKey { get; set; } = null!;

        public VirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFileGetArgs()
        {
        }
    }
}
