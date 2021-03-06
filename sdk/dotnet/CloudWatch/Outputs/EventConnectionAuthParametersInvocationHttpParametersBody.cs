// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch.Outputs
{

    [OutputType]
    public sealed class EventConnectionAuthParametersInvocationHttpParametersBody
    {
        /// <summary>
        /// Specified whether the value is secret.
        /// </summary>
        public readonly bool? IsValueSecret;
        /// <summary>
        /// Header Name.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// Header Value. Created and stored in AWS Secrets Manager.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private EventConnectionAuthParametersInvocationHttpParametersBody(
            bool? isValueSecret,

            string? key,

            string? value)
        {
            IsValueSecret = isValueSecret;
            Key = key;
            Value = value;
        }
    }
}
