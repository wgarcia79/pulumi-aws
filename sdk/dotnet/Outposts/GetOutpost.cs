// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Outposts
{
    public static class GetOutpost
    {
        /// <summary>
        /// Provides details about an Outposts Outpost.
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
        ///         var example = Output.Create(Aws.Outposts.GetOutpost.InvokeAsync(new Aws.Outposts.GetOutpostArgs
        ///         {
        ///             Name = "example",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOutpostResult> InvokeAsync(GetOutpostArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOutpostResult>("aws:outposts/getOutpost:getOutpost", args ?? new GetOutpostArgs(), options.WithVersion());
    }


    public sealed class GetOutpostArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN).
        /// </summary>
        [Input("arn")]
        public string? Arn { get; set; }

        /// <summary>
        /// Identifier of the Outpost.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// Name of the Outpost.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetOutpostArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOutpostResult
    {
        public readonly string Arn;
        /// <summary>
        /// Availability Zone name.
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// Availability Zone identifier.
        /// </summary>
        public readonly string AvailabilityZoneId;
        /// <summary>
        /// Description.
        /// </summary>
        public readonly string Description;
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// AWS Account identifier of the Outpost owner.
        /// </summary>
        public readonly string OwnerId;
        /// <summary>
        /// Site identifier.
        /// </summary>
        public readonly string SiteId;

        [OutputConstructor]
        private GetOutpostResult(
            string arn,

            string availabilityZone,

            string availabilityZoneId,

            string description,

            string id,

            string name,

            string ownerId,

            string siteId)
        {
            Arn = arn;
            AvailabilityZone = availabilityZone;
            AvailabilityZoneId = availabilityZoneId;
            Description = description;
            Id = id;
            Name = name;
            OwnerId = ownerId;
            SiteId = siteId;
        }
    }
}
