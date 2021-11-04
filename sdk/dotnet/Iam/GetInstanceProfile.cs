// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Aws.Iam
{
    public static class GetInstanceProfile
    {
        /// <summary>
        /// This data source can be used to fetch information about a specific
        /// IAM instance profile. By using this data source, you can reference IAM
        /// instance profile properties without having to hard code ARNs as input.
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
        ///         var example = Output.Create(Aws.Iam.GetInstanceProfile.InvokeAsync(new Aws.Iam.GetInstanceProfileArgs
        ///         {
        ///             Name = "an_example_instance_profile_name",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceProfileResult> InvokeAsync(GetInstanceProfileArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceProfileResult>("aws:iam/getInstanceProfile:getInstanceProfile", args ?? new GetInstanceProfileArgs(), options.WithVersion());

        /// <summary>
        /// This data source can be used to fetch information about a specific
        /// IAM instance profile. By using this data source, you can reference IAM
        /// instance profile properties without having to hard code ARNs as input.
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
        ///         var example = Output.Create(Aws.Iam.GetInstanceProfile.InvokeAsync(new Aws.Iam.GetInstanceProfileArgs
        ///         {
        ///             Name = "an_example_instance_profile_name",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstanceProfileResult> Invoke(GetInstanceProfileInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceProfileResult>("aws:iam/getInstanceProfile:getInstanceProfile", args ?? new GetInstanceProfileInvokeArgs(), options.WithVersion());
    }


    public sealed class GetInstanceProfileArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The friendly IAM instance profile name to match.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetInstanceProfileArgs()
        {
        }
    }

    public sealed class GetInstanceProfileInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The friendly IAM instance profile name to match.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetInstanceProfileInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceProfileResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) specifying the instance profile.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The string representation of the date the instance profile
        /// was created.
        /// </summary>
        public readonly string CreateDate;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The path to the instance profile.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// The role arn associated with this instance profile.
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// The role id associated with this instance profile.
        /// </summary>
        public readonly string RoleId;
        /// <summary>
        /// The role name associated with this instance profile.
        /// </summary>
        public readonly string RoleName;

        [OutputConstructor]
        private GetInstanceProfileResult(
            string arn,

            string createDate,

            string id,

            string name,

            string path,

            string roleArn,

            string roleId,

            string roleName)
        {
            Arn = arn;
            CreateDate = createDate;
            Id = id;
            Name = name;
            Path = path;
            RoleArn = roleArn;
            RoleId = roleId;
            RoleName = roleName;
        }
    }
}
