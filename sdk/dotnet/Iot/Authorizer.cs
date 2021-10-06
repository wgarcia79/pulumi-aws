// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iot
{
    /// <summary>
    /// Creates and manages an AWS IoT Authorizer.
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
    ///         var example = new Aws.Iot.Authorizer("example", new Aws.Iot.AuthorizerArgs
    ///         {
    ///             AuthorizerFunctionArn = aws_lambda_function.Example.Arn,
    ///             SigningDisabled = false,
    ///             Status = "ACTIVE",
    ///             TokenKeyName = "Token-Header",
    ///             TokenSigningPublicKeys = 
    ///             {
    ///                 { "Key1", File.ReadAllText("test-fixtures/iot-authorizer-signing-key.pem") },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// IOT Authorizers can be imported using the name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:iot/authorizer:Authorizer example example
    /// ```
    /// </summary>
    [AwsResourceType("aws:iot/authorizer:Authorizer")]
    public partial class Authorizer : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the authorizer.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ARN of the authorizer's Lambda function.
        /// </summary>
        [Output("authorizerFunctionArn")]
        public Output<string> AuthorizerFunctionArn { get; private set; } = null!;

        /// <summary>
        /// The name of the authorizer.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies whether AWS IoT validates the token signature in an authorization request. Default: `false`.
        /// </summary>
        [Output("signingDisabled")]
        public Output<bool?> SigningDisabled { get; private set; } = null!;

        /// <summary>
        /// The status of Authorizer request at creation. Valid values: `ACTIVE`, `INACTIVE`. Default: `ACTIVE`.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// The name of the token key used to extract the token from the HTTP headers. This value is required if signing is enabled in your authorizer.
        /// </summary>
        [Output("tokenKeyName")]
        public Output<string?> TokenKeyName { get; private set; } = null!;

        /// <summary>
        /// The public keys used to verify the digital signature returned by your custom authentication service. This value is required if signing is enabled in your authorizer.
        /// </summary>
        [Output("tokenSigningPublicKeys")]
        public Output<ImmutableDictionary<string, string>?> TokenSigningPublicKeys { get; private set; } = null!;


        /// <summary>
        /// Create a Authorizer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Authorizer(string name, AuthorizerArgs args, CustomResourceOptions? options = null)
            : base("aws:iot/authorizer:Authorizer", name, args ?? new AuthorizerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Authorizer(string name, Input<string> id, AuthorizerState? state = null, CustomResourceOptions? options = null)
            : base("aws:iot/authorizer:Authorizer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Authorizer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Authorizer Get(string name, Input<string> id, AuthorizerState? state = null, CustomResourceOptions? options = null)
        {
            return new Authorizer(name, id, state, options);
        }
    }

    public sealed class AuthorizerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the authorizer's Lambda function.
        /// </summary>
        [Input("authorizerFunctionArn", required: true)]
        public Input<string> AuthorizerFunctionArn { get; set; } = null!;

        /// <summary>
        /// The name of the authorizer.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies whether AWS IoT validates the token signature in an authorization request. Default: `false`.
        /// </summary>
        [Input("signingDisabled")]
        public Input<bool>? SigningDisabled { get; set; }

        /// <summary>
        /// The status of Authorizer request at creation. Valid values: `ACTIVE`, `INACTIVE`. Default: `ACTIVE`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The name of the token key used to extract the token from the HTTP headers. This value is required if signing is enabled in your authorizer.
        /// </summary>
        [Input("tokenKeyName")]
        public Input<string>? TokenKeyName { get; set; }

        [Input("tokenSigningPublicKeys")]
        private InputMap<string>? _tokenSigningPublicKeys;

        /// <summary>
        /// The public keys used to verify the digital signature returned by your custom authentication service. This value is required if signing is enabled in your authorizer.
        /// </summary>
        public InputMap<string> TokenSigningPublicKeys
        {
            get => _tokenSigningPublicKeys ?? (_tokenSigningPublicKeys = new InputMap<string>());
            set => _tokenSigningPublicKeys = value;
        }

        public AuthorizerArgs()
        {
        }
    }

    public sealed class AuthorizerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the authorizer.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The ARN of the authorizer's Lambda function.
        /// </summary>
        [Input("authorizerFunctionArn")]
        public Input<string>? AuthorizerFunctionArn { get; set; }

        /// <summary>
        /// The name of the authorizer.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies whether AWS IoT validates the token signature in an authorization request. Default: `false`.
        /// </summary>
        [Input("signingDisabled")]
        public Input<bool>? SigningDisabled { get; set; }

        /// <summary>
        /// The status of Authorizer request at creation. Valid values: `ACTIVE`, `INACTIVE`. Default: `ACTIVE`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The name of the token key used to extract the token from the HTTP headers. This value is required if signing is enabled in your authorizer.
        /// </summary>
        [Input("tokenKeyName")]
        public Input<string>? TokenKeyName { get; set; }

        [Input("tokenSigningPublicKeys")]
        private InputMap<string>? _tokenSigningPublicKeys;

        /// <summary>
        /// The public keys used to verify the digital signature returned by your custom authentication service. This value is required if signing is enabled in your authorizer.
        /// </summary>
        public InputMap<string> TokenSigningPublicKeys
        {
            get => _tokenSigningPublicKeys ?? (_tokenSigningPublicKeys = new InputMap<string>());
            set => _tokenSigningPublicKeys = value;
        }

        public AuthorizerState()
        {
        }
    }
}
