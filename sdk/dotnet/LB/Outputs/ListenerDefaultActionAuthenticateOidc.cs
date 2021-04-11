// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LB.Outputs
{

    [OutputType]
    public sealed class ListenerDefaultActionAuthenticateOidc
    {
        /// <summary>
        /// Query parameters to include in the redirect request to the authorization endpoint. Max: 10.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? AuthenticationRequestExtraParams;
        /// <summary>
        /// Authorization endpoint of the IdP.
        /// </summary>
        public readonly string AuthorizationEndpoint;
        /// <summary>
        /// OAuth 2.0 client identifier.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// OAuth 2.0 client secret.
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// OIDC issuer identifier of the IdP.
        /// </summary>
        public readonly string Issuer;
        /// <summary>
        /// Behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
        /// </summary>
        public readonly string? OnUnauthenticatedRequest;
        /// <summary>
        /// Set of user claims to be requested from the IdP.
        /// </summary>
        public readonly string? Scope;
        /// <summary>
        /// Name of the cookie used to maintain session information.
        /// </summary>
        public readonly string? SessionCookieName;
        /// <summary>
        /// Maximum duration of the authentication session, in seconds.
        /// </summary>
        public readonly int? SessionTimeout;
        /// <summary>
        /// Token endpoint of the IdP.
        /// </summary>
        public readonly string TokenEndpoint;
        /// <summary>
        /// User info endpoint of the IdP.
        /// </summary>
        public readonly string UserInfoEndpoint;

        [OutputConstructor]
        private ListenerDefaultActionAuthenticateOidc(
            ImmutableDictionary<string, string>? authenticationRequestExtraParams,

            string authorizationEndpoint,

            string clientId,

            string clientSecret,

            string issuer,

            string? onUnauthenticatedRequest,

            string? scope,

            string? sessionCookieName,

            int? sessionTimeout,

            string tokenEndpoint,

            string userInfoEndpoint)
        {
            AuthenticationRequestExtraParams = authenticationRequestExtraParams;
            AuthorizationEndpoint = authorizationEndpoint;
            ClientId = clientId;
            ClientSecret = clientSecret;
            Issuer = issuer;
            OnUnauthenticatedRequest = onUnauthenticatedRequest;
            Scope = scope;
            SessionCookieName = sessionCookieName;
            SessionTimeout = sessionTimeout;
            TokenEndpoint = tokenEndpoint;
            UserInfoEndpoint = userInfoEndpoint;
        }
    }
}
