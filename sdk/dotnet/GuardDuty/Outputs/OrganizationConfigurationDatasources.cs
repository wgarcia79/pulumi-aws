// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GuardDuty.Outputs
{

    [OutputType]
    public sealed class OrganizationConfigurationDatasources
    {
        /// <summary>
        /// Configuration for the builds to store logs to S3.
        /// </summary>
        public readonly Outputs.OrganizationConfigurationDatasourcesS3Logs? S3Logs;

        [OutputConstructor]
        private OrganizationConfigurationDatasources(Outputs.OrganizationConfigurationDatasourcesS3Logs? s3Logs)
        {
            S3Logs = s3Logs;
        }
    }
}
