// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LB.Inputs
{

    public sealed class ListenerDefaultActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration block for using Amazon Cognito to authenticate users. Specify only when `type` is `authenticate-cognito`. Detailed below.
        /// </summary>
        [Input("authenticateCognito")]
        public Input<Inputs.ListenerDefaultActionAuthenticateCognitoArgs>? AuthenticateCognito { get; set; }

        /// <summary>
        /// Configuration block for an identity provider that is compliant with OpenID Connect (OIDC). Specify only when `type` is `authenticate-oidc`. Detailed below.
        /// </summary>
        [Input("authenticateOidc")]
        public Input<Inputs.ListenerDefaultActionAuthenticateOidcArgs>? AuthenticateOidc { get; set; }

        /// <summary>
        /// Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
        /// </summary>
        [Input("fixedResponse")]
        public Input<Inputs.ListenerDefaultActionFixedResponseArgs>? FixedResponse { get; set; }

        /// <summary>
        /// Configuration block for creating an action that distributes requests among one or more target groups. Specify only if `type` is `forward`. If you specify both `forward` block and `target_group_arn` attribute, you can specify only one target group using `forward` and it must be the same target group specified in `target_group_arn`. Detailed below.
        /// </summary>
        [Input("forward")]
        public Input<Inputs.ListenerDefaultActionForwardArgs>? Forward { get; set; }

        /// <summary>
        /// Order for the action. This value is required for rules with multiple actions. The action with the lowest value for order is performed first. Valid values are between `1` and `50000`.
        /// </summary>
        [Input("order")]
        public Input<int>? Order { get; set; }

        /// <summary>
        /// Configuration block for creating a redirect action. Required if `type` is `redirect`. Detailed below.
        /// </summary>
        [Input("redirect")]
        public Input<Inputs.ListenerDefaultActionRedirectArgs>? Redirect { get; set; }

        /// <summary>
        /// ARN of the Target Group to which to route traffic. Specify only if `type` is `forward` and you want to route to a single target group. To route to one or more target groups, use a `forward` block instead.
        /// </summary>
        [Input("targetGroupArn")]
        public Input<string>? TargetGroupArn { get; set; }

        /// <summary>
        /// Type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ListenerDefaultActionArgs()
        {
        }
    }
}
