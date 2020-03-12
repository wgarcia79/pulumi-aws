// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticLoadBalancingV2
{
    /// <summary>
    /// Provides a Load Balancer Listener resource.
    /// 
    /// &gt; **Note:** `aws.alb.Listener` is known as `aws.lb.Listener`. The functionality is identical.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_listener.html.markdown.
    /// </summary>
    public partial class Listener : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the listener (matches `id`)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`aws.lb.ListenerCertificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
        /// </summary>
        [Output("certificateArn")]
        public Output<string?> CertificateArn { get; private set; } = null!;

        /// <summary>
        /// An Action block. Action blocks are documented below.
        /// </summary>
        [Output("defaultActions")]
        public Output<ImmutableArray<Outputs.ListenerDefaultActions>> DefaultActions { get; private set; } = null!;

        /// <summary>
        /// The ARN of the load balancer.
        /// </summary>
        [Output("loadBalancerArn")]
        public Output<string> LoadBalancerArn { get; private set; } = null!;

        /// <summary>
        /// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;

        /// <summary>
        /// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
        /// </summary>
        [Output("sslPolicy")]
        public Output<string> SslPolicy { get; private set; } = null!;


        /// <summary>
        /// Create a Listener resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Listener(string name, ListenerArgs args, CustomResourceOptions? options = null)
            : base("aws:elasticloadbalancingv2/listener:Listener", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Listener(string name, Input<string> id, ListenerState? state = null, CustomResourceOptions? options = null)
            : base("aws:elasticloadbalancingv2/listener:Listener", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Listener resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Listener Get(string name, Input<string> id, ListenerState? state = null, CustomResourceOptions? options = null)
        {
            return new Listener(name, id, state, options);
        }
    }

    public sealed class ListenerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`aws.lb.ListenerCertificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
        /// </summary>
        [Input("certificateArn")]
        public Input<string>? CertificateArn { get; set; }

        [Input("defaultActions", required: true)]
        private InputList<Inputs.ListenerDefaultActionsArgs>? _defaultActions;

        /// <summary>
        /// An Action block. Action blocks are documented below.
        /// </summary>
        public InputList<Inputs.ListenerDefaultActionsArgs> DefaultActions
        {
            get => _defaultActions ?? (_defaultActions = new InputList<Inputs.ListenerDefaultActionsArgs>());
            set => _defaultActions = value;
        }

        /// <summary>
        /// The ARN of the load balancer.
        /// </summary>
        [Input("loadBalancerArn", required: true)]
        public Input<string> LoadBalancerArn { get; set; } = null!;

        /// <summary>
        /// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
        /// </summary>
        [Input("sslPolicy")]
        public Input<string>? SslPolicy { get; set; }

        public ListenerArgs()
        {
        }
    }

    public sealed class ListenerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the listener (matches `id`)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`aws.lb.ListenerCertificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
        /// </summary>
        [Input("certificateArn")]
        public Input<string>? CertificateArn { get; set; }

        [Input("defaultActions")]
        private InputList<Inputs.ListenerDefaultActionsGetArgs>? _defaultActions;

        /// <summary>
        /// An Action block. Action blocks are documented below.
        /// </summary>
        public InputList<Inputs.ListenerDefaultActionsGetArgs> DefaultActions
        {
            get => _defaultActions ?? (_defaultActions = new InputList<Inputs.ListenerDefaultActionsGetArgs>());
            set => _defaultActions = value;
        }

        /// <summary>
        /// The ARN of the load balancer.
        /// </summary>
        [Input("loadBalancerArn")]
        public Input<string>? LoadBalancerArn { get; set; }

        /// <summary>
        /// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
        /// </summary>
        [Input("sslPolicy")]
        public Input<string>? SslPolicy { get; set; }

        public ListenerState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ListenerDefaultActionsArgs : Pulumi.ResourceArgs
    {
        [Input("authenticateCognito")]
        public Input<ListenerDefaultActionsAuthenticateCognitoArgs>? AuthenticateCognito { get; set; }

        [Input("authenticateOidc")]
        public Input<ListenerDefaultActionsAuthenticateOidcArgs>? AuthenticateOidc { get; set; }

        /// <summary>
        /// Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
        /// </summary>
        [Input("fixedResponse")]
        public Input<ListenerDefaultActionsFixedResponseArgs>? FixedResponse { get; set; }

        [Input("order")]
        public Input<int>? Order { get; set; }

        /// <summary>
        /// Information for creating a redirect action. Required if `type` is `redirect`.
        /// </summary>
        [Input("redirect")]
        public Input<ListenerDefaultActionsRedirectArgs>? Redirect { get; set; }

        /// <summary>
        /// The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
        /// </summary>
        [Input("targetGroupArn")]
        public Input<string>? TargetGroupArn { get; set; }

        /// <summary>
        /// The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ListenerDefaultActionsArgs()
        {
        }
    }

    public sealed class ListenerDefaultActionsAuthenticateCognitoArgs : Pulumi.ResourceArgs
    {
        [Input("authenticationRequestExtraParams")]
        private InputMap<object>? _authenticationRequestExtraParams;

        /// <summary>
        /// The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
        /// </summary>
        public InputMap<object> AuthenticationRequestExtraParams
        {
            get => _authenticationRequestExtraParams ?? (_authenticationRequestExtraParams = new InputMap<object>());
            set => _authenticationRequestExtraParams = value;
        }

        /// <summary>
        /// The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
        /// </summary>
        [Input("onUnauthenticatedRequest")]
        public Input<string>? OnUnauthenticatedRequest { get; set; }

        /// <summary>
        /// The set of user claims to be requested from the IdP.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// The name of the cookie used to maintain session information.
        /// </summary>
        [Input("sessionCookieName")]
        public Input<string>? SessionCookieName { get; set; }

        /// <summary>
        /// The maximum duration of the authentication session, in seconds.
        /// </summary>
        [Input("sessionTimeout")]
        public Input<int>? SessionTimeout { get; set; }

        /// <summary>
        /// The ARN of the Cognito user pool.
        /// </summary>
        [Input("userPoolArn", required: true)]
        public Input<string> UserPoolArn { get; set; } = null!;

        /// <summary>
        /// The ID of the Cognito user pool client.
        /// </summary>
        [Input("userPoolClientId", required: true)]
        public Input<string> UserPoolClientId { get; set; } = null!;

        /// <summary>
        /// The domain prefix or fully-qualified domain name of the Cognito user pool.
        /// </summary>
        [Input("userPoolDomain", required: true)]
        public Input<string> UserPoolDomain { get; set; } = null!;

        public ListenerDefaultActionsAuthenticateCognitoArgs()
        {
        }
    }

    public sealed class ListenerDefaultActionsAuthenticateCognitoGetArgs : Pulumi.ResourceArgs
    {
        [Input("authenticationRequestExtraParams")]
        private InputMap<object>? _authenticationRequestExtraParams;

        /// <summary>
        /// The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
        /// </summary>
        public InputMap<object> AuthenticationRequestExtraParams
        {
            get => _authenticationRequestExtraParams ?? (_authenticationRequestExtraParams = new InputMap<object>());
            set => _authenticationRequestExtraParams = value;
        }

        /// <summary>
        /// The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
        /// </summary>
        [Input("onUnauthenticatedRequest")]
        public Input<string>? OnUnauthenticatedRequest { get; set; }

        /// <summary>
        /// The set of user claims to be requested from the IdP.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// The name of the cookie used to maintain session information.
        /// </summary>
        [Input("sessionCookieName")]
        public Input<string>? SessionCookieName { get; set; }

        /// <summary>
        /// The maximum duration of the authentication session, in seconds.
        /// </summary>
        [Input("sessionTimeout")]
        public Input<int>? SessionTimeout { get; set; }

        /// <summary>
        /// The ARN of the Cognito user pool.
        /// </summary>
        [Input("userPoolArn", required: true)]
        public Input<string> UserPoolArn { get; set; } = null!;

        /// <summary>
        /// The ID of the Cognito user pool client.
        /// </summary>
        [Input("userPoolClientId", required: true)]
        public Input<string> UserPoolClientId { get; set; } = null!;

        /// <summary>
        /// The domain prefix or fully-qualified domain name of the Cognito user pool.
        /// </summary>
        [Input("userPoolDomain", required: true)]
        public Input<string> UserPoolDomain { get; set; } = null!;

        public ListenerDefaultActionsAuthenticateCognitoGetArgs()
        {
        }
    }

    public sealed class ListenerDefaultActionsAuthenticateOidcArgs : Pulumi.ResourceArgs
    {
        [Input("authenticationRequestExtraParams")]
        private InputMap<object>? _authenticationRequestExtraParams;

        /// <summary>
        /// The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
        /// </summary>
        public InputMap<object> AuthenticationRequestExtraParams
        {
            get => _authenticationRequestExtraParams ?? (_authenticationRequestExtraParams = new InputMap<object>());
            set => _authenticationRequestExtraParams = value;
        }

        /// <summary>
        /// The authorization endpoint of the IdP.
        /// </summary>
        [Input("authorizationEndpoint", required: true)]
        public Input<string> AuthorizationEndpoint { get; set; } = null!;

        /// <summary>
        /// The OAuth 2.0 client identifier.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The OAuth 2.0 client secret.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        /// <summary>
        /// The OIDC issuer identifier of the IdP.
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        /// <summary>
        /// The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
        /// </summary>
        [Input("onUnauthenticatedRequest")]
        public Input<string>? OnUnauthenticatedRequest { get; set; }

        /// <summary>
        /// The set of user claims to be requested from the IdP.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// The name of the cookie used to maintain session information.
        /// </summary>
        [Input("sessionCookieName")]
        public Input<string>? SessionCookieName { get; set; }

        /// <summary>
        /// The maximum duration of the authentication session, in seconds.
        /// </summary>
        [Input("sessionTimeout")]
        public Input<int>? SessionTimeout { get; set; }

        /// <summary>
        /// The token endpoint of the IdP.
        /// </summary>
        [Input("tokenEndpoint", required: true)]
        public Input<string> TokenEndpoint { get; set; } = null!;

        /// <summary>
        /// The user info endpoint of the IdP.
        /// </summary>
        [Input("userInfoEndpoint", required: true)]
        public Input<string> UserInfoEndpoint { get; set; } = null!;

        public ListenerDefaultActionsAuthenticateOidcArgs()
        {
        }
    }

    public sealed class ListenerDefaultActionsAuthenticateOidcGetArgs : Pulumi.ResourceArgs
    {
        [Input("authenticationRequestExtraParams")]
        private InputMap<object>? _authenticationRequestExtraParams;

        /// <summary>
        /// The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
        /// </summary>
        public InputMap<object> AuthenticationRequestExtraParams
        {
            get => _authenticationRequestExtraParams ?? (_authenticationRequestExtraParams = new InputMap<object>());
            set => _authenticationRequestExtraParams = value;
        }

        /// <summary>
        /// The authorization endpoint of the IdP.
        /// </summary>
        [Input("authorizationEndpoint", required: true)]
        public Input<string> AuthorizationEndpoint { get; set; } = null!;

        /// <summary>
        /// The OAuth 2.0 client identifier.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The OAuth 2.0 client secret.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        /// <summary>
        /// The OIDC issuer identifier of the IdP.
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        /// <summary>
        /// The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
        /// </summary>
        [Input("onUnauthenticatedRequest")]
        public Input<string>? OnUnauthenticatedRequest { get; set; }

        /// <summary>
        /// The set of user claims to be requested from the IdP.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// The name of the cookie used to maintain session information.
        /// </summary>
        [Input("sessionCookieName")]
        public Input<string>? SessionCookieName { get; set; }

        /// <summary>
        /// The maximum duration of the authentication session, in seconds.
        /// </summary>
        [Input("sessionTimeout")]
        public Input<int>? SessionTimeout { get; set; }

        /// <summary>
        /// The token endpoint of the IdP.
        /// </summary>
        [Input("tokenEndpoint", required: true)]
        public Input<string> TokenEndpoint { get; set; } = null!;

        /// <summary>
        /// The user info endpoint of the IdP.
        /// </summary>
        [Input("userInfoEndpoint", required: true)]
        public Input<string> UserInfoEndpoint { get; set; } = null!;

        public ListenerDefaultActionsAuthenticateOidcGetArgs()
        {
        }
    }

    public sealed class ListenerDefaultActionsFixedResponseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.
        /// </summary>
        [Input("contentType", required: true)]
        public Input<string> ContentType { get; set; } = null!;

        /// <summary>
        /// The message body.
        /// </summary>
        [Input("messageBody")]
        public Input<string>? MessageBody { get; set; }

        /// <summary>
        /// The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
        /// </summary>
        [Input("statusCode")]
        public Input<string>? StatusCode { get; set; }

        public ListenerDefaultActionsFixedResponseArgs()
        {
        }
    }

    public sealed class ListenerDefaultActionsFixedResponseGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.
        /// </summary>
        [Input("contentType", required: true)]
        public Input<string> ContentType { get; set; } = null!;

        /// <summary>
        /// The message body.
        /// </summary>
        [Input("messageBody")]
        public Input<string>? MessageBody { get; set; }

        /// <summary>
        /// The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
        /// </summary>
        [Input("statusCode")]
        public Input<string>? StatusCode { get; set; }

        public ListenerDefaultActionsFixedResponseGetArgs()
        {
        }
    }

    public sealed class ListenerDefaultActionsGetArgs : Pulumi.ResourceArgs
    {
        [Input("authenticateCognito")]
        public Input<ListenerDefaultActionsAuthenticateCognitoGetArgs>? AuthenticateCognito { get; set; }

        [Input("authenticateOidc")]
        public Input<ListenerDefaultActionsAuthenticateOidcGetArgs>? AuthenticateOidc { get; set; }

        /// <summary>
        /// Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
        /// </summary>
        [Input("fixedResponse")]
        public Input<ListenerDefaultActionsFixedResponseGetArgs>? FixedResponse { get; set; }

        [Input("order")]
        public Input<int>? Order { get; set; }

        /// <summary>
        /// Information for creating a redirect action. Required if `type` is `redirect`.
        /// </summary>
        [Input("redirect")]
        public Input<ListenerDefaultActionsRedirectGetArgs>? Redirect { get; set; }

        /// <summary>
        /// The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
        /// </summary>
        [Input("targetGroupArn")]
        public Input<string>? TargetGroupArn { get; set; }

        /// <summary>
        /// The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ListenerDefaultActionsGetArgs()
        {
        }
    }

    public sealed class ListenerDefaultActionsRedirectArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The hostname. This component is not percent-encoded. The hostname can contain `#{host}`. Defaults to `#{host}`.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to `/#{path}`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to `#{query}`.
        /// </summary>
        [Input("query")]
        public Input<string>? Query { get; set; }

        /// <summary>
        /// The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
        /// </summary>
        [Input("statusCode", required: true)]
        public Input<string> StatusCode { get; set; } = null!;

        public ListenerDefaultActionsRedirectArgs()
        {
        }
    }

    public sealed class ListenerDefaultActionsRedirectGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The hostname. This component is not percent-encoded. The hostname can contain `#{host}`. Defaults to `#{host}`.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to `/#{path}`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to `#{query}`.
        /// </summary>
        [Input("query")]
        public Input<string>? Query { get; set; }

        /// <summary>
        /// The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
        /// </summary>
        [Input("statusCode", required: true)]
        public Input<string> StatusCode { get; set; } = null!;

        public ListenerDefaultActionsRedirectGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ListenerDefaultActions
    {
        public readonly ListenerDefaultActionsAuthenticateCognito? AuthenticateCognito;
        public readonly ListenerDefaultActionsAuthenticateOidc? AuthenticateOidc;
        /// <summary>
        /// Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
        /// </summary>
        public readonly ListenerDefaultActionsFixedResponse? FixedResponse;
        public readonly int Order;
        /// <summary>
        /// Information for creating a redirect action. Required if `type` is `redirect`.
        /// </summary>
        public readonly ListenerDefaultActionsRedirect? Redirect;
        /// <summary>
        /// The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
        /// </summary>
        public readonly string? TargetGroupArn;
        /// <summary>
        /// The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ListenerDefaultActions(
            ListenerDefaultActionsAuthenticateCognito? authenticateCognito,
            ListenerDefaultActionsAuthenticateOidc? authenticateOidc,
            ListenerDefaultActionsFixedResponse? fixedResponse,
            int order,
            ListenerDefaultActionsRedirect? redirect,
            string? targetGroupArn,
            string type)
        {
            AuthenticateCognito = authenticateCognito;
            AuthenticateOidc = authenticateOidc;
            FixedResponse = fixedResponse;
            Order = order;
            Redirect = redirect;
            TargetGroupArn = targetGroupArn;
            Type = type;
        }
    }

    [OutputType]
    public sealed class ListenerDefaultActionsAuthenticateCognito
    {
        /// <summary>
        /// The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? AuthenticationRequestExtraParams;
        /// <summary>
        /// The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
        /// </summary>
        public readonly string OnUnauthenticatedRequest;
        /// <summary>
        /// The set of user claims to be requested from the IdP.
        /// </summary>
        public readonly string Scope;
        /// <summary>
        /// The name of the cookie used to maintain session information.
        /// </summary>
        public readonly string SessionCookieName;
        /// <summary>
        /// The maximum duration of the authentication session, in seconds.
        /// </summary>
        public readonly int SessionTimeout;
        /// <summary>
        /// The ARN of the Cognito user pool.
        /// </summary>
        public readonly string UserPoolArn;
        /// <summary>
        /// The ID of the Cognito user pool client.
        /// </summary>
        public readonly string UserPoolClientId;
        /// <summary>
        /// The domain prefix or fully-qualified domain name of the Cognito user pool.
        /// </summary>
        public readonly string UserPoolDomain;

        [OutputConstructor]
        private ListenerDefaultActionsAuthenticateCognito(
            ImmutableDictionary<string, object>? authenticationRequestExtraParams,
            string onUnauthenticatedRequest,
            string scope,
            string sessionCookieName,
            int sessionTimeout,
            string userPoolArn,
            string userPoolClientId,
            string userPoolDomain)
        {
            AuthenticationRequestExtraParams = authenticationRequestExtraParams;
            OnUnauthenticatedRequest = onUnauthenticatedRequest;
            Scope = scope;
            SessionCookieName = sessionCookieName;
            SessionTimeout = sessionTimeout;
            UserPoolArn = userPoolArn;
            UserPoolClientId = userPoolClientId;
            UserPoolDomain = userPoolDomain;
        }
    }

    [OutputType]
    public sealed class ListenerDefaultActionsAuthenticateOidc
    {
        /// <summary>
        /// The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? AuthenticationRequestExtraParams;
        /// <summary>
        /// The authorization endpoint of the IdP.
        /// </summary>
        public readonly string AuthorizationEndpoint;
        /// <summary>
        /// The OAuth 2.0 client identifier.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The OAuth 2.0 client secret.
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// The OIDC issuer identifier of the IdP.
        /// </summary>
        public readonly string Issuer;
        /// <summary>
        /// The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
        /// </summary>
        public readonly string OnUnauthenticatedRequest;
        /// <summary>
        /// The set of user claims to be requested from the IdP.
        /// </summary>
        public readonly string Scope;
        /// <summary>
        /// The name of the cookie used to maintain session information.
        /// </summary>
        public readonly string SessionCookieName;
        /// <summary>
        /// The maximum duration of the authentication session, in seconds.
        /// </summary>
        public readonly int SessionTimeout;
        /// <summary>
        /// The token endpoint of the IdP.
        /// </summary>
        public readonly string TokenEndpoint;
        /// <summary>
        /// The user info endpoint of the IdP.
        /// </summary>
        public readonly string UserInfoEndpoint;

        [OutputConstructor]
        private ListenerDefaultActionsAuthenticateOidc(
            ImmutableDictionary<string, object>? authenticationRequestExtraParams,
            string authorizationEndpoint,
            string clientId,
            string clientSecret,
            string issuer,
            string onUnauthenticatedRequest,
            string scope,
            string sessionCookieName,
            int sessionTimeout,
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

    [OutputType]
    public sealed class ListenerDefaultActionsFixedResponse
    {
        /// <summary>
        /// The content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.
        /// </summary>
        public readonly string ContentType;
        /// <summary>
        /// The message body.
        /// </summary>
        public readonly string? MessageBody;
        /// <summary>
        /// The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
        /// </summary>
        public readonly string StatusCode;

        [OutputConstructor]
        private ListenerDefaultActionsFixedResponse(
            string contentType,
            string? messageBody,
            string statusCode)
        {
            ContentType = contentType;
            MessageBody = messageBody;
            StatusCode = statusCode;
        }
    }

    [OutputType]
    public sealed class ListenerDefaultActionsRedirect
    {
        /// <summary>
        /// The hostname. This component is not percent-encoded. The hostname can contain `#{host}`. Defaults to `#{host}`.
        /// </summary>
        public readonly string? Host;
        /// <summary>
        /// The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to `/#{path}`.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
        /// </summary>
        public readonly string? Port;
        /// <summary>
        /// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to `#{query}`.
        /// </summary>
        public readonly string? Query;
        /// <summary>
        /// The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
        /// </summary>
        public readonly string StatusCode;

        [OutputConstructor]
        private ListenerDefaultActionsRedirect(
            string? host,
            string? path,
            string? port,
            string? protocol,
            string? query,
            string statusCode)
        {
            Host = host;
            Path = path;
            Port = port;
            Protocol = protocol;
            Query = query;
            StatusCode = statusCode;
        }
    }
    }
}
