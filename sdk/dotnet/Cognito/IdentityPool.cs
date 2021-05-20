// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito
{
    /// <summary>
    /// Provides an AWS Cognito Identity Pool.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.IO;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = new Aws.Iam.SamlProvider("default", new Aws.Iam.SamlProviderArgs
    ///         {
    ///             SamlMetadataDocument = File.ReadAllText("saml-metadata.xml"),
    ///         });
    ///         var main = new Aws.Cognito.IdentityPool("main", new Aws.Cognito.IdentityPoolArgs
    ///         {
    ///             IdentityPoolName = "identity pool",
    ///             AllowUnauthenticatedIdentities = false,
    ///             AllowClassicFlow = false,
    ///             CognitoIdentityProviders = 
    ///             {
    ///                 new Aws.Cognito.Inputs.IdentityPoolCognitoIdentityProviderArgs
    ///                 {
    ///                     ClientId = "6lhlkkfbfb4q5kpp90urffae",
    ///                     ProviderName = "cognito-idp.us-east-1.amazonaws.com/us-east-1_Tv0493apJ",
    ///                     ServerSideTokenCheck = false,
    ///                 },
    ///                 new Aws.Cognito.Inputs.IdentityPoolCognitoIdentityProviderArgs
    ///                 {
    ///                     ClientId = "7kodkvfqfb4qfkp39eurffae",
    ///                     ProviderName = "cognito-idp.us-east-1.amazonaws.com/eu-west-1_Zr231apJu",
    ///                     ServerSideTokenCheck = false,
    ///                 },
    ///             },
    ///             SupportedLoginProviders = 
    ///             {
    ///                 { "graph.facebook.com", "7346241598935552" },
    ///                 { "accounts.google.com", "123456789012.apps.googleusercontent.com" },
    ///             },
    ///             SamlProviderArns = 
    ///             {
    ///                 @default.Arn,
    ///             },
    ///             OpenidConnectProviderArns = 
    ///             {
    ///                 "arn:aws:iam::123456789012:oidc-provider/id.example.com",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cognito Identity Pool can be imported using the name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:cognito/identityPool:IdentityPool mypool &lt;identity-pool-id&gt;
    /// ```
    /// </summary>
    [AwsResourceType("aws:cognito/identityPool:IdentityPool")]
    public partial class IdentityPool : Pulumi.CustomResource
    {
        /// <summary>
        /// Enables or disables the classic / basic authentication flow. Default is `false`.
        /// </summary>
        [Output("allowClassicFlow")]
        public Output<bool?> AllowClassicFlow { get; private set; } = null!;

        /// <summary>
        /// Whether the identity pool supports unauthenticated logins or not.
        /// </summary>
        [Output("allowUnauthenticatedIdentities")]
        public Output<bool?> AllowUnauthenticatedIdentities { get; private set; } = null!;

        /// <summary>
        /// The ARN of the identity pool.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// An array of Amazon Cognito Identity user pools and their client IDs.
        /// </summary>
        [Output("cognitoIdentityProviders")]
        public Output<ImmutableArray<Outputs.IdentityPoolCognitoIdentityProvider>> CognitoIdentityProviders { get; private set; } = null!;

        /// <summary>
        /// The "domain" by which Cognito will refer to your users. This name acts as a placeholder that allows your
        /// backend and the Cognito service to communicate about the developer provider.
        /// </summary>
        [Output("developerProviderName")]
        public Output<string?> DeveloperProviderName { get; private set; } = null!;

        /// <summary>
        /// The Cognito Identity Pool name.
        /// </summary>
        [Output("identityPoolName")]
        public Output<string> IdentityPoolName { get; private set; } = null!;

        /// <summary>
        /// Set of OpendID Connect provider ARNs.
        /// </summary>
        [Output("openidConnectProviderArns")]
        public Output<ImmutableArray<string>> OpenidConnectProviderArns { get; private set; } = null!;

        /// <summary>
        /// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity.
        /// </summary>
        [Output("samlProviderArns")]
        public Output<ImmutableArray<string>> SamlProviderArns { get; private set; } = null!;

        /// <summary>
        /// Key-Value pairs mapping provider names to provider app IDs.
        /// </summary>
        [Output("supportedLoginProviders")]
        public Output<ImmutableDictionary<string, string>?> SupportedLoginProviders { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the Identity Pool. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityPool(string name, IdentityPoolArgs args, CustomResourceOptions? options = null)
            : base("aws:cognito/identityPool:IdentityPool", name, args ?? new IdentityPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityPool(string name, Input<string> id, IdentityPoolState? state = null, CustomResourceOptions? options = null)
            : base("aws:cognito/identityPool:IdentityPool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IdentityPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityPool Get(string name, Input<string> id, IdentityPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityPool(name, id, state, options);
        }
    }

    public sealed class IdentityPoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables or disables the classic / basic authentication flow. Default is `false`.
        /// </summary>
        [Input("allowClassicFlow")]
        public Input<bool>? AllowClassicFlow { get; set; }

        /// <summary>
        /// Whether the identity pool supports unauthenticated logins or not.
        /// </summary>
        [Input("allowUnauthenticatedIdentities")]
        public Input<bool>? AllowUnauthenticatedIdentities { get; set; }

        [Input("cognitoIdentityProviders")]
        private InputList<Inputs.IdentityPoolCognitoIdentityProviderArgs>? _cognitoIdentityProviders;

        /// <summary>
        /// An array of Amazon Cognito Identity user pools and their client IDs.
        /// </summary>
        public InputList<Inputs.IdentityPoolCognitoIdentityProviderArgs> CognitoIdentityProviders
        {
            get => _cognitoIdentityProviders ?? (_cognitoIdentityProviders = new InputList<Inputs.IdentityPoolCognitoIdentityProviderArgs>());
            set => _cognitoIdentityProviders = value;
        }

        /// <summary>
        /// The "domain" by which Cognito will refer to your users. This name acts as a placeholder that allows your
        /// backend and the Cognito service to communicate about the developer provider.
        /// </summary>
        [Input("developerProviderName")]
        public Input<string>? DeveloperProviderName { get; set; }

        /// <summary>
        /// The Cognito Identity Pool name.
        /// </summary>
        [Input("identityPoolName", required: true)]
        public Input<string> IdentityPoolName { get; set; } = null!;

        [Input("openidConnectProviderArns")]
        private InputList<string>? _openidConnectProviderArns;

        /// <summary>
        /// Set of OpendID Connect provider ARNs.
        /// </summary>
        public InputList<string> OpenidConnectProviderArns
        {
            get => _openidConnectProviderArns ?? (_openidConnectProviderArns = new InputList<string>());
            set => _openidConnectProviderArns = value;
        }

        [Input("samlProviderArns")]
        private InputList<string>? _samlProviderArns;

        /// <summary>
        /// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity.
        /// </summary>
        public InputList<string> SamlProviderArns
        {
            get => _samlProviderArns ?? (_samlProviderArns = new InputList<string>());
            set => _samlProviderArns = value;
        }

        [Input("supportedLoginProviders")]
        private InputMap<string>? _supportedLoginProviders;

        /// <summary>
        /// Key-Value pairs mapping provider names to provider app IDs.
        /// </summary>
        public InputMap<string> SupportedLoginProviders
        {
            get => _supportedLoginProviders ?? (_supportedLoginProviders = new InputMap<string>());
            set => _supportedLoginProviders = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the Identity Pool. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public IdentityPoolArgs()
        {
        }
    }

    public sealed class IdentityPoolState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables or disables the classic / basic authentication flow. Default is `false`.
        /// </summary>
        [Input("allowClassicFlow")]
        public Input<bool>? AllowClassicFlow { get; set; }

        /// <summary>
        /// Whether the identity pool supports unauthenticated logins or not.
        /// </summary>
        [Input("allowUnauthenticatedIdentities")]
        public Input<bool>? AllowUnauthenticatedIdentities { get; set; }

        /// <summary>
        /// The ARN of the identity pool.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("cognitoIdentityProviders")]
        private InputList<Inputs.IdentityPoolCognitoIdentityProviderGetArgs>? _cognitoIdentityProviders;

        /// <summary>
        /// An array of Amazon Cognito Identity user pools and their client IDs.
        /// </summary>
        public InputList<Inputs.IdentityPoolCognitoIdentityProviderGetArgs> CognitoIdentityProviders
        {
            get => _cognitoIdentityProviders ?? (_cognitoIdentityProviders = new InputList<Inputs.IdentityPoolCognitoIdentityProviderGetArgs>());
            set => _cognitoIdentityProviders = value;
        }

        /// <summary>
        /// The "domain" by which Cognito will refer to your users. This name acts as a placeholder that allows your
        /// backend and the Cognito service to communicate about the developer provider.
        /// </summary>
        [Input("developerProviderName")]
        public Input<string>? DeveloperProviderName { get; set; }

        /// <summary>
        /// The Cognito Identity Pool name.
        /// </summary>
        [Input("identityPoolName")]
        public Input<string>? IdentityPoolName { get; set; }

        [Input("openidConnectProviderArns")]
        private InputList<string>? _openidConnectProviderArns;

        /// <summary>
        /// Set of OpendID Connect provider ARNs.
        /// </summary>
        public InputList<string> OpenidConnectProviderArns
        {
            get => _openidConnectProviderArns ?? (_openidConnectProviderArns = new InputList<string>());
            set => _openidConnectProviderArns = value;
        }

        [Input("samlProviderArns")]
        private InputList<string>? _samlProviderArns;

        /// <summary>
        /// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity.
        /// </summary>
        public InputList<string> SamlProviderArns
        {
            get => _samlProviderArns ?? (_samlProviderArns = new InputList<string>());
            set => _samlProviderArns = value;
        }

        [Input("supportedLoginProviders")]
        private InputMap<string>? _supportedLoginProviders;

        /// <summary>
        /// Key-Value pairs mapping provider names to provider app IDs.
        /// </summary>
        public InputMap<string> SupportedLoginProviders
        {
            get => _supportedLoginProviders ?? (_supportedLoginProviders = new InputMap<string>());
            set => _supportedLoginProviders = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the Identity Pool. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public IdentityPoolState()
        {
        }
    }
}
