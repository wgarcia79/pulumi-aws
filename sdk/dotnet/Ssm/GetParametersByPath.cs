// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Aws.Ssm
{
    public static class GetParametersByPath
    {
        public static Task<GetParametersByPathResult> InvokeAsync(GetParametersByPathArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetParametersByPathResult>("aws:ssm/getParametersByPath:getParametersByPath", args ?? new GetParametersByPathArgs(), options.WithVersion());

        public static Output<GetParametersByPathResult> Invoke(GetParametersByPathInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetParametersByPathResult>("aws:ssm/getParametersByPath:getParametersByPath", args ?? new GetParametersByPathInvokeArgs(), options.WithVersion());
    }


    public sealed class GetParametersByPathArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The prefix path of the parameter.
        /// </summary>
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        /// <summary>
        /// Whether to return decrypted `SecureString` value. Defaults to `true`.
        /// </summary>
        [Input("withDecryption")]
        public bool? WithDecryption { get; set; }

        public GetParametersByPathArgs()
        {
        }
    }

    public sealed class GetParametersByPathInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The prefix path of the parameter.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Whether to return decrypted `SecureString` value. Defaults to `true`.
        /// </summary>
        [Input("withDecryption")]
        public Input<bool>? WithDecryption { get; set; }

        public GetParametersByPathInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetParametersByPathResult
    {
        public readonly ImmutableArray<string> Arns;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Names;
        public readonly string Path;
        public readonly ImmutableArray<string> Types;
        public readonly ImmutableArray<string> Values;
        public readonly bool? WithDecryption;

        [OutputConstructor]
        private GetParametersByPathResult(
            ImmutableArray<string> arns,

            string id,

            ImmutableArray<string> names,

            string path,

            ImmutableArray<string> types,

            ImmutableArray<string> values,

            bool? withDecryption)
        {
            Arns = arns;
            Id = id;
            Names = names;
            Path = path;
            Types = types;
            Values = values;
            WithDecryption = withDecryption;
        }
    }
}