// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Waf.Outputs
{

    [OutputType]
    public sealed class SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch
    {
        /// <summary>
        /// When `type` is `HEADER`, enter the name of the header that you want to search, e.g., `User-Agent` or `Referer`.
        /// If `type` is any other value, omit this field.
        /// </summary>
        public readonly string? Data;
        /// <summary>
        /// The part of the web request that you want AWS WAF to search for a specified string.
        /// e.g., `HEADER`, `METHOD` or `BODY`.
        /// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
        /// for all supported values.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch(
            string? data,

            string type)
        {
            Data = data;
            Type = type;
        }
    }
}
