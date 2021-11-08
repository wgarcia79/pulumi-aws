// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodePipeline.Inputs
{

    public sealed class WebhookFilterGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [JSON path](https://github.com/json-path/JsonPath) to filter on.
        /// </summary>
        [Input("jsonPath", required: true)]
        public Input<string> JsonPath { get; set; } = null!;

        /// <summary>
        /// The value to match on (e.g., `refs/heads/{Branch}`). See [AWS docs](https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_WebhookFilterRule.html) for details.
        /// </summary>
        [Input("matchEquals", required: true)]
        public Input<string> MatchEquals { get; set; } = null!;

        public WebhookFilterGetArgs()
        {
        }
    }
}
