// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Outputs
{

    [OutputType]
    public sealed class WebAclRuleOverrideAction
    {
        /// <summary>
        /// Override the rule action setting to count (i.e., only count matches). Configured as an empty block `{}`.
        /// </summary>
        public readonly Outputs.WebAclRuleOverrideActionCount? Count;
        /// <summary>
        /// Don't override the rule action setting. Configured as an empty block `{}`.
        /// </summary>
        public readonly Outputs.WebAclRuleOverrideActionNone? None;

        [OutputConstructor]
        private WebAclRuleOverrideAction(
            Outputs.WebAclRuleOverrideActionCount? count,

            Outputs.WebAclRuleOverrideActionNone? none)
        {
            Count = count;
            None = none;
        }
    }
}
