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
    public sealed class SqlInjectionMatchSetSqlInjectionMatchTuple
    {
        /// <summary>
        /// Specifies where in a web request to look for snippets of malicious SQL code.
        /// </summary>
        public readonly Outputs.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch FieldToMatch;
        /// <summary>
        /// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
        /// If you specify a transformation, AWS WAF performs the transformation on `field_to_match` before inspecting a request for a match.
        /// e.g., `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
        /// See [docs](https://docs.aws.amazon.com/waf/latest/APIReference/API_regional_SqlInjectionMatchTuple.html#WAF-Type-regional_SqlInjectionMatchTuple-TextTransformation)
        /// for all supported values.
        /// </summary>
        public readonly string TextTransformation;

        [OutputConstructor]
        private SqlInjectionMatchSetSqlInjectionMatchTuple(
            Outputs.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch fieldToMatch,

            string textTransformation)
        {
            FieldToMatch = fieldToMatch;
            TextTransformation = textTransformation;
        }
    }
}
