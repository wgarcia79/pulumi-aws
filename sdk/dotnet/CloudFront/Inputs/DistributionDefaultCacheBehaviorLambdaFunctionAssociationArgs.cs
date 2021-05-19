// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront.Inputs
{

    public sealed class DistributionDefaultCacheBehaviorLambdaFunctionAssociationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The specific event to trigger this function.
        /// Valid values: `viewer-request` or `viewer-response`
        /// </summary>
        [Input("eventType", required: true)]
        public Input<string> EventType { get; set; } = null!;

        /// <summary>
        /// When set to true it exposes the request body to the lambda function. Defaults to false. Valid values: `true`, `false`.
        /// </summary>
        [Input("includeBody")]
        public Input<bool>? IncludeBody { get; set; }

        /// <summary>
        /// ARN of the Lambda function.
        /// </summary>
        [Input("lambdaArn", required: true)]
        public Input<string> LambdaArn { get; set; } = null!;

        public DistributionDefaultCacheBehaviorLambdaFunctionAssociationArgs()
        {
        }
    }
}
