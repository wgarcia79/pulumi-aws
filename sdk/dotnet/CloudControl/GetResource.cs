// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudControl
{
    public static class GetResource
    {
        /// <summary>
        /// Provides details for a Cloud Control API Resource. The reading of these resources is proxied through Cloud Control API handlers to the backend service.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.CloudControl.GetResource.InvokeAsync(new Aws.CloudControl.GetResourceArgs
        ///         {
        ///             Identifier = "example",
        ///             TypeName = "AWS::ECS::Cluster",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetResourceResult> InvokeAsync(GetResourceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourceResult>("aws:cloudcontrol/getResource:getResource", args ?? new GetResourceArgs(), options.WithVersion());
    }


    public sealed class GetResourceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Identifier of the CloudFormation resource type. For example, `vpc-12345678`.
        /// </summary>
        [Input("identifier", required: true)]
        public string Identifier { get; set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the IAM Role to assume for operations.
        /// </summary>
        [Input("roleArn")]
        public string? RoleArn { get; set; }

        /// <summary>
        /// CloudFormation resource type name. For example, `AWS::EC2::VPC`.
        /// </summary>
        [Input("typeName", required: true)]
        public string TypeName { get; set; } = null!;

        /// <summary>
        /// Identifier of the CloudFormation resource type version.
        /// </summary>
        [Input("typeVersionId")]
        public string? TypeVersionId { get; set; }

        public GetResourceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResourceResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Identifier;
        public readonly string Properties;
        public readonly string? RoleArn;
        public readonly string TypeName;
        public readonly string? TypeVersionId;

        [OutputConstructor]
        private GetResourceResult(
            string id,

            string identifier,

            string properties,

            string? roleArn,

            string typeName,

            string? typeVersionId)
        {
            Id = id;
            Identifier = identifier;
            Properties = properties;
            RoleArn = roleArn;
            TypeName = typeName;
            TypeVersionId = typeVersionId;
        }
    }
}
