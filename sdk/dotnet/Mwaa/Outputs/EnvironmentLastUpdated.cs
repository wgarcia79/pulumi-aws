// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Mwaa.Outputs
{

    [OutputType]
    public sealed class EnvironmentLastUpdated
    {
        /// <summary>
        /// The Created At date of the MWAA Environment
        /// * `logging_configuration.&lt;LOG_TYPE&gt;.cloud_watch_log_group_arn` - Provides the ARN for the CloudWatch group where the logs will be published
        /// </summary>
        public readonly string? CreatedAt;
        public readonly ImmutableArray<Outputs.EnvironmentLastUpdatedError> Errors;
        /// <summary>
        /// The status of the Amazon MWAA Environment
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private EnvironmentLastUpdated(
            string? createdAt,

            ImmutableArray<Outputs.EnvironmentLastUpdatedError> errors,

            string? status)
        {
            CreatedAt = createdAt;
            Errors = errors;
            Status = status;
        }
    }
}
