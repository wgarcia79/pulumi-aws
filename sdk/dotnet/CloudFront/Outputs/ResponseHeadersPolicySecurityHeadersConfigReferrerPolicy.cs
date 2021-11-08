// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront.Outputs
{

    [OutputType]
    public sealed class ResponseHeadersPolicySecurityHeadersConfigReferrerPolicy
    {
        /// <summary>
        /// A Boolean value that determines whether CloudFront overrides the X-XSS-Protection HTTP response header received from the origin with the one specified in this response headers policy.
        /// </summary>
        public readonly bool Override;
        /// <summary>
        /// The value of the Referrer-Policy HTTP response header. Valid Values: `no-referrer` | `no-referrer-when-downgrade` | `origin` | `origin-when-cross-origin` | `same-origin` | `strict-origin` | `strict-origin-when-cross-origin` | `unsafe-url`
        /// </summary>
        public readonly string ReferrerPolicy;

        [OutputConstructor]
        private ResponseHeadersPolicySecurityHeadersConfigReferrerPolicy(
            bool @override,

            string referrerPolicy)
        {
            Override = @override;
            ReferrerPolicy = referrerPolicy;
        }
    }
}
