// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecr
{
    /// <summary>
    /// Provides an Elastic Container Registry Repository.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Aws.Ecr.Repository("foo", new Aws.Ecr.RepositoryArgs
    ///         {
    ///             ImageScanningConfiguration = new Aws.Ecr.Inputs.RepositoryImageScanningConfigurationArgs
    ///             {
    ///                 ScanOnPush = true,
    ///             },
    ///             ImageTagMutability = "MUTABLE",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECR Repositories can be imported using the `name`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:ecr/repository:Repository service test-service
    /// ```
    /// </summary>
    [AwsResourceType("aws:ecr/repository:Repository")]
    public partial class Repository : Pulumi.CustomResource
    {
        /// <summary>
        /// Full ARN of the repository.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Encryption configuration for the repository. See below for schema.
        /// </summary>
        [Output("encryptionConfigurations")]
        public Output<ImmutableArray<Outputs.RepositoryEncryptionConfiguration>> EncryptionConfigurations { get; private set; } = null!;

        /// <summary>
        /// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
        /// </summary>
        [Output("imageScanningConfiguration")]
        public Output<Outputs.RepositoryImageScanningConfiguration?> ImageScanningConfiguration { get; private set; } = null!;

        /// <summary>
        /// The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
        /// </summary>
        [Output("imageTagMutability")]
        public Output<string?> ImageTagMutability { get; private set; } = null!;

        /// <summary>
        /// Name of the repository.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The registry ID where the repository was created.
        /// </summary>
        [Output("registryId")]
        public Output<string> RegistryId { get; private set; } = null!;

        /// <summary>
        /// The URL of the repository (in the form `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName`).
        /// </summary>
        [Output("repositoryUrl")]
        public Output<string> RepositoryUrl { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a Repository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Repository(string name, RepositoryArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:ecr/repository:Repository", name, args ?? new RepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Repository(string name, Input<string> id, RepositoryState? state = null, CustomResourceOptions? options = null)
            : base("aws:ecr/repository:Repository", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Repository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Repository Get(string name, Input<string> id, RepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new Repository(name, id, state, options);
        }
    }

    public sealed class RepositoryArgs : Pulumi.ResourceArgs
    {
        [Input("encryptionConfigurations")]
        private InputList<Inputs.RepositoryEncryptionConfigurationArgs>? _encryptionConfigurations;

        /// <summary>
        /// Encryption configuration for the repository. See below for schema.
        /// </summary>
        public InputList<Inputs.RepositoryEncryptionConfigurationArgs> EncryptionConfigurations
        {
            get => _encryptionConfigurations ?? (_encryptionConfigurations = new InputList<Inputs.RepositoryEncryptionConfigurationArgs>());
            set => _encryptionConfigurations = value;
        }

        /// <summary>
        /// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
        /// </summary>
        [Input("imageScanningConfiguration")]
        public Input<Inputs.RepositoryImageScanningConfigurationArgs>? ImageScanningConfiguration { get; set; }

        /// <summary>
        /// The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
        /// </summary>
        [Input("imageTagMutability")]
        public Input<string>? ImageTagMutability { get; set; }

        /// <summary>
        /// Name of the repository.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public RepositoryArgs()
        {
        }
    }

    public sealed class RepositoryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Full ARN of the repository.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("encryptionConfigurations")]
        private InputList<Inputs.RepositoryEncryptionConfigurationGetArgs>? _encryptionConfigurations;

        /// <summary>
        /// Encryption configuration for the repository. See below for schema.
        /// </summary>
        public InputList<Inputs.RepositoryEncryptionConfigurationGetArgs> EncryptionConfigurations
        {
            get => _encryptionConfigurations ?? (_encryptionConfigurations = new InputList<Inputs.RepositoryEncryptionConfigurationGetArgs>());
            set => _encryptionConfigurations = value;
        }

        /// <summary>
        /// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
        /// </summary>
        [Input("imageScanningConfiguration")]
        public Input<Inputs.RepositoryImageScanningConfigurationGetArgs>? ImageScanningConfiguration { get; set; }

        /// <summary>
        /// The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
        /// </summary>
        [Input("imageTagMutability")]
        public Input<string>? ImageTagMutability { get; set; }

        /// <summary>
        /// Name of the repository.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The registry ID where the repository was created.
        /// </summary>
        [Input("registryId")]
        public Input<string>? RegistryId { get; set; }

        /// <summary>
        /// The URL of the repository (in the form `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName`).
        /// </summary>
        [Input("repositoryUrl")]
        public Input<string>? RepositoryUrl { get; set; }

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

        public RepositoryState()
        {
        }
    }
}
