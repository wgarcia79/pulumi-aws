// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppRunner.Inputs
{

    public sealed class ServiceSourceConfigurationAuthenticationConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the IAM role that grants the App Runner service access to a source repository. Required for ECR image repositories (but not for ECR Public)
        /// </summary>
        [Input("accessRoleArn")]
        public Input<string>? AccessRoleArn { get; set; }

        /// <summary>
        /// ARN of the App Runner connection that enables the App Runner service to connect to a source repository. Required for GitHub code repositories.
        /// </summary>
        [Input("connectionArn")]
        public Input<string>? ConnectionArn { get; set; }

        public ServiceSourceConfigurationAuthenticationConfigurationGetArgs()
        {
        }
    }
}
