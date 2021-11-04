// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Aws.ServerlessRepository
{
    public static class GetApplication
    {
        /// <summary>
        /// Use this data source to get information about an AWS Serverless Application Repository application. For example, this can be used to determine the required `capabilities` for an application.
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
        ///         var exampleApplication = Output.Create(Aws.ServerlessRepository.GetApplication.InvokeAsync(new Aws.ServerlessRepository.GetApplicationArgs
        ///         {
        ///             ApplicationId = "arn:aws:serverlessrepo:us-east-1:123456789012:applications/ExampleApplication",
        ///         }));
        ///         var exampleCloudFormationStack = new Aws.ServerlessRepository.CloudFormationStack("exampleCloudFormationStack", new Aws.ServerlessRepository.CloudFormationStackArgs
        ///         {
        ///             ApplicationId = exampleApplication.Apply(exampleApplication =&gt; exampleApplication.ApplicationId),
        ///             SemanticVersion = exampleApplication.Apply(exampleApplication =&gt; exampleApplication.SemanticVersion),
        ///             Capabilities = exampleApplication.Apply(exampleApplication =&gt; exampleApplication.RequiredCapabilities),
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetApplicationResult> InvokeAsync(GetApplicationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationResult>("aws:serverlessrepository/getApplication:getApplication", args ?? new GetApplicationArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to get information about an AWS Serverless Application Repository application. For example, this can be used to determine the required `capabilities` for an application.
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
        ///         var exampleApplication = Output.Create(Aws.ServerlessRepository.GetApplication.InvokeAsync(new Aws.ServerlessRepository.GetApplicationArgs
        ///         {
        ///             ApplicationId = "arn:aws:serverlessrepo:us-east-1:123456789012:applications/ExampleApplication",
        ///         }));
        ///         var exampleCloudFormationStack = new Aws.ServerlessRepository.CloudFormationStack("exampleCloudFormationStack", new Aws.ServerlessRepository.CloudFormationStackArgs
        ///         {
        ///             ApplicationId = exampleApplication.Apply(exampleApplication =&gt; exampleApplication.ApplicationId),
        ///             SemanticVersion = exampleApplication.Apply(exampleApplication =&gt; exampleApplication.SemanticVersion),
        ///             Capabilities = exampleApplication.Apply(exampleApplication =&gt; exampleApplication.RequiredCapabilities),
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetApplicationResult> Invoke(GetApplicationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetApplicationResult>("aws:serverlessrepository/getApplication:getApplication", args ?? new GetApplicationInvokeArgs(), options.WithVersion());
    }


    public sealed class GetApplicationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the application.
        /// </summary>
        [Input("applicationId", required: true)]
        public string ApplicationId { get; set; } = null!;

        /// <summary>
        /// The requested version of the application. By default, retrieves the latest version.
        /// </summary>
        [Input("semanticVersion")]
        public string? SemanticVersion { get; set; }

        public GetApplicationArgs()
        {
        }
    }

    public sealed class GetApplicationInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the application.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// The requested version of the application. By default, retrieves the latest version.
        /// </summary>
        [Input("semanticVersion")]
        public Input<string>? SemanticVersion { get; set; }

        public GetApplicationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApplicationResult
    {
        /// <summary>
        /// The ARN of the application.
        /// </summary>
        public readonly string ApplicationId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the application.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A list of capabilities describing the permissions needed to deploy the application.
        /// </summary>
        public readonly ImmutableArray<string> RequiredCapabilities;
        public readonly string SemanticVersion;
        /// <summary>
        /// A URL pointing to the source code of the application version.
        /// </summary>
        public readonly string SourceCodeUrl;
        /// <summary>
        /// A URL pointing to the Cloud Formation template for the application version.
        /// </summary>
        public readonly string TemplateUrl;

        [OutputConstructor]
        private GetApplicationResult(
            string applicationId,

            string id,

            string name,

            ImmutableArray<string> requiredCapabilities,

            string semanticVersion,

            string sourceCodeUrl,

            string templateUrl)
        {
            ApplicationId = applicationId;
            Id = id;
            Name = name;
            RequiredCapabilities = requiredCapabilities;
            SemanticVersion = semanticVersion;
            SourceCodeUrl = sourceCodeUrl;
            TemplateUrl = templateUrl;
        }
    }
}
