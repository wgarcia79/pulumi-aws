// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GuardDuty.Inputs
{

    public sealed class OrganizationConfigurationDatasourcesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration for the builds to store logs to S3.
        /// </summary>
        [Input("s3Logs")]
        public Input<Inputs.OrganizationConfigurationDatasourcesS3LogsGetArgs>? S3Logs { get; set; }

        public OrganizationConfigurationDatasourcesGetArgs()
        {
        }
    }
}
