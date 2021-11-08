// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker
{
    /// <summary>
    /// Provides a Sagemaker Notebook Instance resource.
    /// 
    /// ## Example Usage
    /// ### Basic usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var ni = new Aws.Sagemaker.NotebookInstance("ni", new Aws.Sagemaker.NotebookInstanceArgs
    ///         {
    ///             RoleArn = aws_iam_role.Role.Arn,
    ///             InstanceType = "ml.t2.medium",
    ///             Tags = 
    ///             {
    ///                 { "Name", "foo" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Code repository usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Sagemaker.CodeRepository("example", new Aws.Sagemaker.CodeRepositoryArgs
    ///         {
    ///             CodeRepositoryName = "my-notebook-instance-code-repo",
    ///             GitConfig = new Aws.Sagemaker.Inputs.CodeRepositoryGitConfigArgs
    ///             {
    ///                 RepositoryUrl = "https://github.com/hashicorp/terraform-provider-aws.git",
    ///             },
    ///         });
    ///         var ni = new Aws.Sagemaker.NotebookInstance("ni", new Aws.Sagemaker.NotebookInstanceArgs
    ///         {
    ///             RoleArn = aws_iam_role.Role.Arn,
    ///             InstanceType = "ml.t2.medium",
    ///             DefaultCodeRepository = example.CodeRepositoryName,
    ///             Tags = 
    ///             {
    ///                 { "Name", "foo" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Sagemaker Notebook Instances can be imported using the `name`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:sagemaker/notebookInstance:NotebookInstance test_notebook_instance my-notebook-instance
    /// ```
    /// </summary>
    [AwsResourceType("aws:sagemaker/notebookInstance:NotebookInstance")]
    public partial class NotebookInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// An array of up to three Git repositories to associate with the notebook instance.
        /// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
        /// </summary>
        [Output("additionalCodeRepositories")]
        public Output<ImmutableArray<string>> AdditionalCodeRepositories { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) assigned by AWS to this notebook instance.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository.
        /// </summary>
        [Output("defaultCodeRepository")]
        public Output<string?> DefaultCodeRepository { get; private set; } = null!;

        /// <summary>
        /// Set to `Disabled` to disable internet access to notebook. Requires `security_groups` and `subnet_id` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
        /// </summary>
        [Output("directInternetAccess")]
        public Output<string?> DirectInternetAccess { get; private set; } = null!;

        /// <summary>
        /// The name of ML compute instance type.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string?> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// The name of a lifecycle configuration to associate with the notebook instance.
        /// </summary>
        [Output("lifecycleConfigName")]
        public Output<string?> LifecycleConfigName { get; private set; } = null!;

        /// <summary>
        /// The name of the notebook instance (must be unique).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The network interface ID that Amazon SageMaker created at the time of creating the instance. Only available when setting `subnet_id`.
        /// </summary>
        [Output("networkInterfaceId")]
        public Output<string> NetworkInterfaceId { get; private set; } = null!;

        /// <summary>
        /// The platform identifier of the notebook instance runtime environment. This value can be either `notebook-al1-v1` or `notebook-al2-v1`, depending on which version of Amazon Linux you require.
        /// </summary>
        [Output("platformIdentifier")]
        public Output<string> PlatformIdentifier { get; private set; } = null!;

        /// <summary>
        /// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
        /// </summary>
        [Output("rootAccess")]
        public Output<string?> RootAccess { get; private set; } = null!;

        /// <summary>
        /// The associated security groups.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// The VPC subnet ID.
        /// </summary>
        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The URL that you use to connect to the Jupyter notebook that is running in your notebook instance.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
        /// </summary>
        [Output("volumeSize")]
        public Output<int?> VolumeSize { get; private set; } = null!;


        /// <summary>
        /// Create a NotebookInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NotebookInstance(string name, NotebookInstanceArgs args, CustomResourceOptions? options = null)
            : base("aws:sagemaker/notebookInstance:NotebookInstance", name, args ?? new NotebookInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NotebookInstance(string name, Input<string> id, NotebookInstanceState? state = null, CustomResourceOptions? options = null)
            : base("aws:sagemaker/notebookInstance:NotebookInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NotebookInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NotebookInstance Get(string name, Input<string> id, NotebookInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new NotebookInstance(name, id, state, options);
        }
    }

    public sealed class NotebookInstanceArgs : Pulumi.ResourceArgs
    {
        [Input("additionalCodeRepositories")]
        private InputList<string>? _additionalCodeRepositories;

        /// <summary>
        /// An array of up to three Git repositories to associate with the notebook instance.
        /// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
        /// </summary>
        public InputList<string> AdditionalCodeRepositories
        {
            get => _additionalCodeRepositories ?? (_additionalCodeRepositories = new InputList<string>());
            set => _additionalCodeRepositories = value;
        }

        /// <summary>
        /// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository.
        /// </summary>
        [Input("defaultCodeRepository")]
        public Input<string>? DefaultCodeRepository { get; set; }

        /// <summary>
        /// Set to `Disabled` to disable internet access to notebook. Requires `security_groups` and `subnet_id` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
        /// </summary>
        [Input("directInternetAccess")]
        public Input<string>? DirectInternetAccess { get; set; }

        /// <summary>
        /// The name of ML compute instance type.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// The name of a lifecycle configuration to associate with the notebook instance.
        /// </summary>
        [Input("lifecycleConfigName")]
        public Input<string>? LifecycleConfigName { get; set; }

        /// <summary>
        /// The name of the notebook instance (must be unique).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The platform identifier of the notebook instance runtime environment. This value can be either `notebook-al1-v1` or `notebook-al2-v1`, depending on which version of Amazon Linux you require.
        /// </summary>
        [Input("platformIdentifier")]
        public Input<string>? PlatformIdentifier { get; set; }

        /// <summary>
        /// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
        /// </summary>
        [Input("rootAccess")]
        public Input<string>? RootAccess { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// The associated security groups.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// The VPC subnet ID.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
        /// </summary>
        [Input("volumeSize")]
        public Input<int>? VolumeSize { get; set; }

        public NotebookInstanceArgs()
        {
        }
    }

    public sealed class NotebookInstanceState : Pulumi.ResourceArgs
    {
        [Input("additionalCodeRepositories")]
        private InputList<string>? _additionalCodeRepositories;

        /// <summary>
        /// An array of up to three Git repositories to associate with the notebook instance.
        /// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
        /// </summary>
        public InputList<string> AdditionalCodeRepositories
        {
            get => _additionalCodeRepositories ?? (_additionalCodeRepositories = new InputList<string>());
            set => _additionalCodeRepositories = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) assigned by AWS to this notebook instance.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository.
        /// </summary>
        [Input("defaultCodeRepository")]
        public Input<string>? DefaultCodeRepository { get; set; }

        /// <summary>
        /// Set to `Disabled` to disable internet access to notebook. Requires `security_groups` and `subnet_id` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
        /// </summary>
        [Input("directInternetAccess")]
        public Input<string>? DirectInternetAccess { get; set; }

        /// <summary>
        /// The name of ML compute instance type.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// The name of a lifecycle configuration to associate with the notebook instance.
        /// </summary>
        [Input("lifecycleConfigName")]
        public Input<string>? LifecycleConfigName { get; set; }

        /// <summary>
        /// The name of the notebook instance (must be unique).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network interface ID that Amazon SageMaker created at the time of creating the instance. Only available when setting `subnet_id`.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// The platform identifier of the notebook instance runtime environment. This value can be either `notebook-al1-v1` or `notebook-al2-v1`, depending on which version of Amazon Linux you require.
        /// </summary>
        [Input("platformIdentifier")]
        public Input<string>? PlatformIdentifier { get; set; }

        /// <summary>
        /// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
        /// </summary>
        [Input("rootAccess")]
        public Input<string>? RootAccess { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// The associated security groups.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// The VPC subnet ID.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        /// <summary>
        /// The URL that you use to connect to the Jupyter notebook that is running in your notebook instance.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
        /// </summary>
        [Input("volumeSize")]
        public Input<int>? VolumeSize { get; set; }

        public NotebookInstanceState()
        {
        }
    }
}
