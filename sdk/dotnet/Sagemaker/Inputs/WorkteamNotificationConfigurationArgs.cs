// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Inputs
{

    public sealed class WorkteamNotificationConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN for the SNS topic to which notifications should be published.
        /// </summary>
        [Input("notificationTopicArn")]
        public Input<string>? NotificationTopicArn { get; set; }

        public WorkteamNotificationConfigurationArgs()
        {
        }
    }
}
