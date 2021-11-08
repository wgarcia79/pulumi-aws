// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafRegional.Outputs
{

    [OutputType]
    public sealed class XssMatchSetXssMatchTupleFieldToMatch
    {
        /// <summary>
        /// When the value of `type` is `HEADER`, enter the name of the header that you want the WAF to search, for example, `User-Agent` or `Referer`. If the value of `type` is any other value, omit `data`.
        /// </summary>
        public readonly string? Data;
        /// <summary>
        /// The part of the web request that you want AWS WAF to search for a specified stringE.g., `HEADER` or `METHOD`
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private XssMatchSetXssMatchTupleFieldToMatch(
            string? data,

            string type)
        {
            Data = data;
            Type = type;
        }
    }
}
