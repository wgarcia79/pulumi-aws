// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront.Inputs
{

    public sealed class ResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// TThe policy directives and their values that CloudFront includes as values for the Content-Security-Policy HTTP response header.
        /// </summary>
        [Input("contentSecurityPolicy", required: true)]
        public Input<string> ContentSecurityPolicy { get; set; } = null!;

        /// <summary>
        /// A Boolean value that determines whether CloudFront overrides the X-XSS-Protection HTTP response header received from the origin with the one specified in this response headers policy.
        /// </summary>
        [Input("override", required: true)]
        public Input<bool> Override { get; set; } = null!;

        public ResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyArgs()
        {
        }
    }
}
