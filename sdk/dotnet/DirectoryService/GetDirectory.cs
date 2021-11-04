// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Aws.DirectoryService
{
    public static class GetDirectory
    {
        /// <summary>
        /// Get attributes of AWS Directory Service directory (SimpleAD, Managed AD, AD Connector). It's especially useful to refer AWS Managed AD or on-premise AD in AD Connector configuration.
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
        ///         var example = Output.Create(Aws.DirectoryService.GetDirectory.InvokeAsync(new Aws.DirectoryService.GetDirectoryArgs
        ///         {
        ///             DirectoryId = aws_directory_service_directory.Main.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDirectoryResult> InvokeAsync(GetDirectoryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDirectoryResult>("aws:directoryservice/getDirectory:getDirectory", args ?? new GetDirectoryArgs(), options.WithVersion());

        /// <summary>
        /// Get attributes of AWS Directory Service directory (SimpleAD, Managed AD, AD Connector). It's especially useful to refer AWS Managed AD or on-premise AD in AD Connector configuration.
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
        ///         var example = Output.Create(Aws.DirectoryService.GetDirectory.InvokeAsync(new Aws.DirectoryService.GetDirectoryArgs
        ///         {
        ///             DirectoryId = aws_directory_service_directory.Main.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDirectoryResult> Invoke(GetDirectoryInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDirectoryResult>("aws:directoryservice/getDirectory:getDirectory", args ?? new GetDirectoryInvokeArgs(), options.WithVersion());
    }


    public sealed class GetDirectoryArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the directory.
        /// </summary>
        [Input("directoryId", required: true)]
        public string DirectoryId { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// A map of tags assigned to the directory/connector.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetDirectoryArgs()
        {
        }
    }

    public sealed class GetDirectoryInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the directory.
        /// </summary>
        [Input("directoryId", required: true)]
        public Input<string> DirectoryId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags assigned to the directory/connector.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetDirectoryInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDirectoryResult
    {
        /// <summary>
        /// The access URL for the directory/connector, such as http://alias.awsapps.com.
        /// </summary>
        public readonly string AccessUrl;
        /// <summary>
        /// The alias for the directory/connector, such as `d-991708b282.awsapps.com`.
        /// </summary>
        public readonly string Alias;
        public readonly ImmutableArray<Outputs.GetDirectoryConnectSettingResult> ConnectSettings;
        /// <summary>
        /// A textual description for the directory/connector.
        /// </summary>
        public readonly string Description;
        public readonly string DirectoryId;
        /// <summary>
        /// A list of IP addresses of the DNS servers for the directory/connector.
        /// </summary>
        public readonly ImmutableArray<string> DnsIpAddresses;
        /// <summary>
        /// (for `MicrosoftAD`) The Microsoft AD edition (`Standard` or `Enterprise`).
        /// </summary>
        public readonly string Edition;
        /// <summary>
        /// The directory/connector single-sign on status.
        /// </summary>
        public readonly bool EnableSso;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The fully qualified name for the directory/connector.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID of the security group created by the directory/connector.
        /// </summary>
        public readonly string SecurityGroupId;
        /// <summary>
        /// The short name of the directory/connector, such as `CORP`.
        /// </summary>
        public readonly string ShortName;
        /// <summary>
        /// (for `SimpleAD` and `ADConnector`) The size of the directory/connector (`Small` or `Large`).
        /// </summary>
        public readonly string Size;
        /// <summary>
        /// A map of tags assigned to the directory/connector.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD`).
        /// </summary>
        public readonly string Type;
        public readonly ImmutableArray<Outputs.GetDirectoryVpcSettingResult> VpcSettings;

        [OutputConstructor]
        private GetDirectoryResult(
            string accessUrl,

            string alias,

            ImmutableArray<Outputs.GetDirectoryConnectSettingResult> connectSettings,

            string description,

            string directoryId,

            ImmutableArray<string> dnsIpAddresses,

            string edition,

            bool enableSso,

            string id,

            string name,

            string securityGroupId,

            string shortName,

            string size,

            ImmutableDictionary<string, string>? tags,

            string type,

            ImmutableArray<Outputs.GetDirectoryVpcSettingResult> vpcSettings)
        {
            AccessUrl = accessUrl;
            Alias = alias;
            ConnectSettings = connectSettings;
            Description = description;
            DirectoryId = directoryId;
            DnsIpAddresses = dnsIpAddresses;
            Edition = edition;
            EnableSso = enableSso;
            Id = id;
            Name = name;
            SecurityGroupId = securityGroupId;
            ShortName = shortName;
            Size = size;
            Tags = tags;
            Type = type;
            VpcSettings = vpcSettings;
        }
    }
}
