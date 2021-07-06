// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iam
{
    public static class GetSessionContext
    {
        public static Task<GetSessionContextResult> InvokeAsync(GetSessionContextArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSessionContextResult>("aws:iam/getSessionContext:getSessionContext", args ?? new GetSessionContextArgs(), options.WithVersion());
    }


    public sealed class GetSessionContextArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ARN for an assumed role.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetSessionContextArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSessionContextResult
    {
        public readonly string Arn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IAM source role ARN if `arn` corresponds to an STS assumed role. Otherwise, `issuer_arn` is equal to `arn`.
        /// </summary>
        public readonly string IssuerArn;
        /// <summary>
        /// Unique identifier of the IAM role that issues the STS assumed role.
        /// </summary>
        public readonly string IssuerId;
        /// <summary>
        /// Name of the source role. Only available if `arn` corresponds to an STS assumed role.
        /// </summary>
        public readonly string IssuerName;
        /// <summary>
        /// Name of the STS session. Only available if `arn` corresponds to an STS assumed role.
        /// </summary>
        public readonly string SessionName;

        [OutputConstructor]
        private GetSessionContextResult(
            string arn,

            string id,

            string issuerArn,

            string issuerId,

            string issuerName,

            string sessionName)
        {
            Arn = arn;
            Id = id;
            IssuerArn = issuerArn;
            IssuerId = issuerId;
            IssuerName = issuerName;
            SessionName = sessionName;
        }
    }
}
