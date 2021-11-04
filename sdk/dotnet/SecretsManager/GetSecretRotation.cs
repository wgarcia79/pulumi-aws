// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Aws.SecretsManager
{
    public static class GetSecretRotation
    {
        /// <summary>
        /// Retrieve information about a Secrets Manager secret rotation. To retrieve secret metadata, see the `aws.secretsmanager.Secret` data source. To retrieve a secret value, see the `aws.secretsmanager.SecretVersion` data source.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Retrieve Secret Rotation Configuration
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.SecretsManager.GetSecretRotation.InvokeAsync(new Aws.SecretsManager.GetSecretRotationArgs
        ///         {
        ///             SecretId = data.Aws_secretsmanager_secret.Example.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSecretRotationResult> InvokeAsync(GetSecretRotationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecretRotationResult>("aws:secretsmanager/getSecretRotation:getSecretRotation", args ?? new GetSecretRotationArgs(), options.WithVersion());

        /// <summary>
        /// Retrieve information about a Secrets Manager secret rotation. To retrieve secret metadata, see the `aws.secretsmanager.Secret` data source. To retrieve a secret value, see the `aws.secretsmanager.SecretVersion` data source.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Retrieve Secret Rotation Configuration
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.SecretsManager.GetSecretRotation.InvokeAsync(new Aws.SecretsManager.GetSecretRotationArgs
        ///         {
        ///             SecretId = data.Aws_secretsmanager_secret.Example.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSecretRotationResult> Invoke(GetSecretRotationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSecretRotationResult>("aws:secretsmanager/getSecretRotation:getSecretRotation", args ?? new GetSecretRotationInvokeArgs(), options.WithVersion());
    }


    public sealed class GetSecretRotationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the secret containing the version that you want to retrieve. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret.
        /// </summary>
        [Input("secretId", required: true)]
        public string SecretId { get; set; } = null!;

        public GetSecretRotationArgs()
        {
        }
    }

    public sealed class GetSecretRotationInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the secret containing the version that you want to retrieve. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret.
        /// </summary>
        [Input("secretId", required: true)]
        public Input<string> SecretId { get; set; } = null!;

        public GetSecretRotationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSecretRotationResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ARN of the secret.
        /// </summary>
        public readonly bool RotationEnabled;
        /// <summary>
        /// The decrypted part of the protected secret information that was originally provided as a string.
        /// </summary>
        public readonly string RotationLambdaArn;
        /// <summary>
        /// The decrypted part of the protected secret information that was originally provided as a binary. Base64 encoded.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSecretRotationRotationRuleResult> RotationRules;
        public readonly string SecretId;

        [OutputConstructor]
        private GetSecretRotationResult(
            string id,

            bool rotationEnabled,

            string rotationLambdaArn,

            ImmutableArray<Outputs.GetSecretRotationRotationRuleResult> rotationRules,

            string secretId)
        {
            Id = id;
            RotationEnabled = rotationEnabled;
            RotationLambdaArn = rotationLambdaArn;
            RotationRules = rotationRules;
            SecretId = secretId;
        }
    }
}
