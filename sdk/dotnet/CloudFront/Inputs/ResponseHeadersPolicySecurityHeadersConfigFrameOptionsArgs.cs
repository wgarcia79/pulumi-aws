// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront.Inputs
{

    public sealed class ResponseHeadersPolicySecurityHeadersConfigFrameOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value of the X-Frame-Options HTTP response header. Valid values: `DENY` | `SAMEORIGIN`
        /// </summary>
        [Input("frameOption", required: true)]
        public Input<string> FrameOption { get; set; } = null!;

        /// <summary>
        /// A Boolean value that determines whether CloudFront overrides the X-XSS-Protection HTTP response header received from the origin with the one specified in this response headers policy.
        /// </summary>
        [Input("override", required: true)]
        public Input<bool> Override { get; set; } = null!;

        public ResponseHeadersPolicySecurityHeadersConfigFrameOptionsArgs()
        {
        }
    }
}
