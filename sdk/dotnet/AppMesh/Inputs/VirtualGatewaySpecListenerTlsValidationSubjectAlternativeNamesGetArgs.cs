// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh.Inputs
{

    public sealed class VirtualGatewaySpecListenerTlsValidationSubjectAlternativeNamesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The criteria for determining a SAN's match.
        /// </summary>
        [Input("match", required: true)]
        public Input<Inputs.VirtualGatewaySpecListenerTlsValidationSubjectAlternativeNamesMatchGetArgs> Match { get; set; } = null!;

        public VirtualGatewaySpecListenerTlsValidationSubjectAlternativeNamesGetArgs()
        {
        }
    }
}
