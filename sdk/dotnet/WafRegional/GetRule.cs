// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Aws.WafRegional
{
    public static class GetRule
    {
        /// <summary>
        /// `aws.wafregional.Rule` Retrieves a WAF Regional Rule Resource Id.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.WafRegional.GetRule.InvokeAsync(new Aws.WafRegional.GetRuleArgs
        ///         {
        ///             Name = "tfWAFRegionalRule",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRuleResult> InvokeAsync(GetRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRuleResult>("aws:wafregional/getRule:getRule", args ?? new GetRuleArgs(), options.WithVersion());

        /// <summary>
        /// `aws.wafregional.Rule` Retrieves a WAF Regional Rule Resource Id.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.WafRegional.GetRule.InvokeAsync(new Aws.WafRegional.GetRuleArgs
        ///         {
        ///             Name = "tfWAFRegionalRule",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRuleResult> Invoke(GetRuleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRuleResult>("aws:wafregional/getRule:getRule", args ?? new GetRuleInvokeArgs(), options.WithVersion());
    }


    public sealed class GetRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the WAF Regional rule.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetRuleArgs()
        {
        }
    }

    public sealed class GetRuleInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the WAF Regional rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetRuleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRuleResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetRuleResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
