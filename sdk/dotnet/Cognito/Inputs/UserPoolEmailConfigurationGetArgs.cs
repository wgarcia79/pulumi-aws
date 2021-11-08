// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito.Inputs
{

    public sealed class UserPoolEmailConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Email configuration set name from SES.
        /// </summary>
        [Input("configurationSet")]
        public Input<string>? ConfigurationSet { get; set; }

        /// <summary>
        /// Email delivery method to use. `COGNITO_DEFAULT` for the default email functionality built into Cognito or `DEVELOPER` to use your Amazon SES configuration.
        /// </summary>
        [Input("emailSendingAccount")]
        public Input<string>? EmailSendingAccount { get; set; }

        /// <summary>
        /// Sender’s email address or sender’s display name with their email address (e.g., `john@example.com`, `John Smith &lt;john@example.com&gt;` or `\"John Smith Ph.D.\" &lt;john@example.com&gt;`). Escaped double quotes are required around display names that contain certain characters as specified in [RFC 5322](https://tools.ietf.org/html/rfc5322).
        /// </summary>
        [Input("fromEmailAddress")]
        public Input<string>? FromEmailAddress { get; set; }

        /// <summary>
        /// REPLY-TO email address.
        /// </summary>
        [Input("replyToEmailAddress")]
        public Input<string>? ReplyToEmailAddress { get; set; }

        /// <summary>
        /// ARN of the SES verified email identity to to use. Required if `email_sending_account` is set to `DEVELOPER`.
        /// </summary>
        [Input("sourceArn")]
        public Input<string>? SourceArn { get; set; }

        public UserPoolEmailConfigurationGetArgs()
        {
        }
    }
}
