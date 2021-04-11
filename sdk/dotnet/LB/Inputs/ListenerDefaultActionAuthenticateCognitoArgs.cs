// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LB.Inputs
{

    public sealed class ListenerDefaultActionAuthenticateCognitoArgs : Pulumi.ResourceArgs
    {
        [Input("authenticationRequestExtraParams")]
        private InputMap<string>? _authenticationRequestExtraParams;

        /// <summary>
        /// Query parameters to include in the redirect request to the authorization endpoint. Max: 10.
        /// </summary>
        public InputMap<string> AuthenticationRequestExtraParams
        {
            get => _authenticationRequestExtraParams ?? (_authenticationRequestExtraParams = new InputMap<string>());
            set => _authenticationRequestExtraParams = value;
        }

        /// <summary>
        /// Behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
        /// </summary>
        [Input("onUnauthenticatedRequest")]
        public Input<string>? OnUnauthenticatedRequest { get; set; }

        /// <summary>
        /// Set of user claims to be requested from the IdP.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// Name of the cookie used to maintain session information.
        /// </summary>
        [Input("sessionCookieName")]
        public Input<string>? SessionCookieName { get; set; }

        /// <summary>
        /// Maximum duration of the authentication session, in seconds.
        /// </summary>
        [Input("sessionTimeout")]
        public Input<int>? SessionTimeout { get; set; }

        /// <summary>
        /// ARN of the Cognito user pool.
        /// </summary>
        [Input("userPoolArn", required: true)]
        public Input<string> UserPoolArn { get; set; } = null!;

        /// <summary>
        /// ID of the Cognito user pool client.
        /// </summary>
        [Input("userPoolClientId", required: true)]
        public Input<string> UserPoolClientId { get; set; } = null!;

        /// <summary>
        /// Domain prefix or fully-qualified domain name of the Cognito user pool.
        /// </summary>
        [Input("userPoolDomain", required: true)]
        public Input<string> UserPoolDomain { get; set; } = null!;

        public ListenerDefaultActionAuthenticateCognitoArgs()
        {
        }
    }
}
