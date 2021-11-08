// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Transfer
{
    /// <summary>
    /// Provides a AWS Transfer Server resource.
    /// 
    /// &gt; **NOTE on AWS IAM permissions:** If the `endpoint_type` is set to `VPC`, the `ec2:DescribeVpcEndpoints` and `ec2:ModifyVpcEndpoint` [actions](https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonec2.html#amazonec2-actions-as-permissions) are used.
    /// 
    /// ## Example Usage
    /// ### Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Transfer.Server("example", new Aws.Transfer.ServerArgs
    ///         {
    ///             Tags = 
    ///             {
    ///                 { "Name", "Example" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Security Policy Name
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Transfer.Server("example", new Aws.Transfer.ServerArgs
    ///         {
    ///             SecurityPolicyName = "TransferSecurityPolicy-2020-06",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### VPC Endpoint
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Transfer.Server("example", new Aws.Transfer.ServerArgs
    ///         {
    ///             EndpointType = "VPC",
    ///             EndpointDetails = new Aws.Transfer.Inputs.ServerEndpointDetailsArgs
    ///             {
    ///                 AddressAllocationIds = 
    ///                 {
    ///                     aws_eip.Example.Id,
    ///                 },
    ///                 SubnetIds = 
    ///                 {
    ///                     aws_subnet.Example.Id,
    ///                 },
    ///                 VpcId = aws_vpc.Example.Id,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### AWS Directory authentication
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Transfer.Server("example", new Aws.Transfer.ServerArgs
    ///         {
    ///             IdentityProviderType = "AWS_DIRECTORY_SERVICE",
    ///             DirectoryId = aws_directory_service_directory.Example.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Protocols
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Transfer.Server("example", new Aws.Transfer.ServerArgs
    ///         {
    ///             EndpointType = "VPC",
    ///             EndpointDetails = new Aws.Transfer.Inputs.ServerEndpointDetailsArgs
    ///             {
    ///                 SubnetIds = 
    ///                 {
    ///                     aws_subnet.Example.Id,
    ///                 },
    ///                 VpcId = aws_vpc.Example.Id,
    ///             },
    ///             Protocols = 
    ///             {
    ///                 "FTP",
    ///                 "FTPS",
    ///             },
    ///             Certificate = aws_acm_certificate.Example.Arn,
    ///             IdentityProviderType = "API_GATEWAY",
    ///             Url = $"{aws_api_gateway_deployment.Example.Invoke_url}{aws_api_gateway_resource.Example.Path}",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Transfer Servers can be imported using the `server id`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:transfer/server:Server example s-12345678
    /// ```
    /// 
    ///  Certain resource arguments, such as `host_key`, cannot be read via the API and imported into the provider. This provider will display a difference for these arguments the first run after import if declared in the provider configuration for an imported resource.
    /// </summary>
    [AwsResourceType("aws:transfer/server:Server")]
    public partial class Server : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of Transfer Server
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate. This is required when `protocols` is set to `FTPS`
        /// </summary>
        [Output("certificate")]
        public Output<string?> Certificate { get; private set; } = null!;

        /// <summary>
        /// The directory service ID of the directory service you want to connect to with an `identity_provider_type` of `AWS_DIRECTORY_SERVICE`.
        /// </summary>
        [Output("directoryId")]
        public Output<string?> DirectoryId { get; private set; } = null!;

        /// <summary>
        /// The domain of the storage system that is used for file transfers. Valid values are: `S3` and `EFS`. The default value is `S3`.
        /// </summary>
        [Output("domain")]
        public Output<string?> Domain { get; private set; } = null!;

        /// <summary>
        /// The endpoint of the Transfer Server (e.g., `s-12345678.server.transfer.REGION.amazonaws.com`)
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
        /// </summary>
        [Output("endpointDetails")]
        public Output<Outputs.ServerEndpointDetails?> EndpointDetails { get; private set; } = null!;

        /// <summary>
        /// The type of endpoint that you want your SFTP server connect to. If you connect to a `VPC` (or `VPC_ENDPOINT`), your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set `PUBLIC`.  Defaults to `PUBLIC`.
        /// </summary>
        [Output("endpointType")]
        public Output<string?> EndpointType { get; private set; } = null!;

        /// <summary>
        /// A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is `false`. This option only applies to servers configured with a `SERVICE_MANAGED` `identity_provider_type`.
        /// </summary>
        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        /// <summary>
        /// RSA private key (e.g., as generated by the `ssh-keygen -N "" -m PEM -f my-new-server-key` command).
        /// </summary>
        [Output("hostKey")]
        public Output<string?> HostKey { get; private set; } = null!;

        /// <summary>
        /// This value contains the message-digest algorithm (MD5) hash of the server's host key. This value is equivalent to the output of the `ssh-keygen -l -E md5 -f my-new-server-key` command.
        /// </summary>
        [Output("hostKeyFingerprint")]
        public Output<string> HostKeyFingerprint { get; private set; } = null!;

        /// <summary>
        /// The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice. Using `AWS_DIRECTORY_SERVICE` will allow for authentication against AWS Managed Active Directory or Microsoft Active Directory in your on-premises environment, or in AWS using AD Connectors.
        /// </summary>
        [Output("identityProviderType")]
        public Output<string?> IdentityProviderType { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identity_provider_type` of `API_GATEWAY`.
        /// </summary>
        [Output("invocationRole")]
        public Output<string?> InvocationRole { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
        /// </summary>
        [Output("loggingRole")]
        public Output<string?> LoggingRole { get; private set; } = null!;

        /// <summary>
        /// Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint. This defaults to `SFTP` . The available protocols are:
        /// * `SFTP`: File transfer over SSH
        /// * `FTPS`: File transfer with TLS encryption
        /// * `FTP`: Unencrypted file transfer
        /// </summary>
        [Output("protocols")]
        public Output<ImmutableArray<string>> Protocols { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the security policy that is attached to the server. Possible values are `TransferSecurityPolicy-2018-11`, `TransferSecurityPolicy-2020-06`, and  `TransferSecurityPolicy-FIPS-2020-06`. Default value is: `TransferSecurityPolicy-2018-11`.
        /// </summary>
        [Output("securityPolicyName")]
        public Output<string?> SecurityPolicyName { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// - URL of the service endpoint used to authenticate users with an `identity_provider_type` of `API_GATEWAY`.
        /// </summary>
        [Output("url")]
        public Output<string?> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Server resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Server(string name, ServerArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:transfer/server:Server", name, args ?? new ServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Server(string name, Input<string> id, ServerState? state = null, CustomResourceOptions? options = null)
            : base("aws:transfer/server:Server", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Server resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Server Get(string name, Input<string> id, ServerState? state = null, CustomResourceOptions? options = null)
        {
            return new Server(name, id, state, options);
        }
    }

    public sealed class ServerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate. This is required when `protocols` is set to `FTPS`
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// The directory service ID of the directory service you want to connect to with an `identity_provider_type` of `AWS_DIRECTORY_SERVICE`.
        /// </summary>
        [Input("directoryId")]
        public Input<string>? DirectoryId { get; set; }

        /// <summary>
        /// The domain of the storage system that is used for file transfers. Valid values are: `S3` and `EFS`. The default value is `S3`.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
        /// </summary>
        [Input("endpointDetails")]
        public Input<Inputs.ServerEndpointDetailsArgs>? EndpointDetails { get; set; }

        /// <summary>
        /// The type of endpoint that you want your SFTP server connect to. If you connect to a `VPC` (or `VPC_ENDPOINT`), your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set `PUBLIC`.  Defaults to `PUBLIC`.
        /// </summary>
        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        /// <summary>
        /// A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is `false`. This option only applies to servers configured with a `SERVICE_MANAGED` `identity_provider_type`.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// RSA private key (e.g., as generated by the `ssh-keygen -N "" -m PEM -f my-new-server-key` command).
        /// </summary>
        [Input("hostKey")]
        public Input<string>? HostKey { get; set; }

        /// <summary>
        /// The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice. Using `AWS_DIRECTORY_SERVICE` will allow for authentication against AWS Managed Active Directory or Microsoft Active Directory in your on-premises environment, or in AWS using AD Connectors.
        /// </summary>
        [Input("identityProviderType")]
        public Input<string>? IdentityProviderType { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identity_provider_type` of `API_GATEWAY`.
        /// </summary>
        [Input("invocationRole")]
        public Input<string>? InvocationRole { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
        /// </summary>
        [Input("loggingRole")]
        public Input<string>? LoggingRole { get; set; }

        [Input("protocols")]
        private InputList<string>? _protocols;

        /// <summary>
        /// Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint. This defaults to `SFTP` . The available protocols are:
        /// * `SFTP`: File transfer over SSH
        /// * `FTPS`: File transfer with TLS encryption
        /// * `FTP`: Unencrypted file transfer
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        /// <summary>
        /// Specifies the name of the security policy that is attached to the server. Possible values are `TransferSecurityPolicy-2018-11`, `TransferSecurityPolicy-2020-06`, and  `TransferSecurityPolicy-FIPS-2020-06`. Default value is: `TransferSecurityPolicy-2018-11`.
        /// </summary>
        [Input("securityPolicyName")]
        public Input<string>? SecurityPolicyName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// - URL of the service endpoint used to authenticate users with an `identity_provider_type` of `API_GATEWAY`.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ServerArgs()
        {
        }
    }

    public sealed class ServerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of Transfer Server
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate. This is required when `protocols` is set to `FTPS`
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// The directory service ID of the directory service you want to connect to with an `identity_provider_type` of `AWS_DIRECTORY_SERVICE`.
        /// </summary>
        [Input("directoryId")]
        public Input<string>? DirectoryId { get; set; }

        /// <summary>
        /// The domain of the storage system that is used for file transfers. Valid values are: `S3` and `EFS`. The default value is `S3`.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The endpoint of the Transfer Server (e.g., `s-12345678.server.transfer.REGION.amazonaws.com`)
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
        /// </summary>
        [Input("endpointDetails")]
        public Input<Inputs.ServerEndpointDetailsGetArgs>? EndpointDetails { get; set; }

        /// <summary>
        /// The type of endpoint that you want your SFTP server connect to. If you connect to a `VPC` (or `VPC_ENDPOINT`), your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set `PUBLIC`.  Defaults to `PUBLIC`.
        /// </summary>
        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        /// <summary>
        /// A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is `false`. This option only applies to servers configured with a `SERVICE_MANAGED` `identity_provider_type`.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// RSA private key (e.g., as generated by the `ssh-keygen -N "" -m PEM -f my-new-server-key` command).
        /// </summary>
        [Input("hostKey")]
        public Input<string>? HostKey { get; set; }

        /// <summary>
        /// This value contains the message-digest algorithm (MD5) hash of the server's host key. This value is equivalent to the output of the `ssh-keygen -l -E md5 -f my-new-server-key` command.
        /// </summary>
        [Input("hostKeyFingerprint")]
        public Input<string>? HostKeyFingerprint { get; set; }

        /// <summary>
        /// The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice. Using `AWS_DIRECTORY_SERVICE` will allow for authentication against AWS Managed Active Directory or Microsoft Active Directory in your on-premises environment, or in AWS using AD Connectors.
        /// </summary>
        [Input("identityProviderType")]
        public Input<string>? IdentityProviderType { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identity_provider_type` of `API_GATEWAY`.
        /// </summary>
        [Input("invocationRole")]
        public Input<string>? InvocationRole { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
        /// </summary>
        [Input("loggingRole")]
        public Input<string>? LoggingRole { get; set; }

        [Input("protocols")]
        private InputList<string>? _protocols;

        /// <summary>
        /// Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint. This defaults to `SFTP` . The available protocols are:
        /// * `SFTP`: File transfer over SSH
        /// * `FTPS`: File transfer with TLS encryption
        /// * `FTP`: Unencrypted file transfer
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        /// <summary>
        /// Specifies the name of the security policy that is attached to the server. Possible values are `TransferSecurityPolicy-2018-11`, `TransferSecurityPolicy-2020-06`, and  `TransferSecurityPolicy-FIPS-2020-06`. Default value is: `TransferSecurityPolicy-2018-11`.
        /// </summary>
        [Input("securityPolicyName")]
        public Input<string>? SecurityPolicyName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// - URL of the service endpoint used to authenticate users with an `identity_provider_type` of `API_GATEWAY`.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ServerState()
        {
        }
    }
}
