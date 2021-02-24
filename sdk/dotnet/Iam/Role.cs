// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iam
{
    /// <summary>
    /// Provides an IAM role.
    /// 
    /// &gt; **NOTE:** If policies are attached to the role via the `aws.iam.PolicyAttachment` resource and you are modifying the role `name` or `path`, the `force_detach_policies` argument must be set to `true` and applied before attempting the operation otherwise you will encounter a `DeleteConflict` error. The `aws.iam.RolePolicyAttachment` resource does not have this requirement.
    /// 
    /// &gt; **NOTE:** If you use this resource's `managed_policy_arns` argument or `inline_policy` configuration blocks, this resource will take over exclusive management of the role's respective policy types (e.g., both policy types if both arguments are used). These arguments are incompatible with other ways of managing a role's policies, such as `aws.iam.PolicyAttachment`, `aws.iam.RolePolicyAttachment`, and `aws.iam.RolePolicy`. If you attempt to manage a role's policies by multiple means, you will get resource cycling and/or errors.
    /// 
    /// ## Example Usage
    /// ### Basic Example
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var testRole = new Aws.Iam.Role("testRole", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 { "Version", "2012-10-17" },
    ///                 { "Statement", new[]
    ///                     {
    ///                         new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             { "Action", "sts:AssumeRole" },
    ///                             { "Effect", "Allow" },
    ///                             { "Sid", "" },
    ///                             { "Principal", new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 { "Service", "ec2.amazonaws.com" },
    ///                             } },
    ///                         },
    ///                     }
    ///                  },
    ///             }),
    ///             Tags = 
    ///             {
    ///                 { "tag-key", "tag-value" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Example of Using Data Source for Assume Role Policy
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var instance_assume_role_policy = Output.Create(Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
    ///         {
    ///             Statements = 
    ///             {
    ///                 new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
    ///                 {
    ///                     Actions = 
    ///                     {
    ///                         "sts:AssumeRole",
    ///                     },
    ///                     Principals = 
    ///                     {
    ///                         new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
    ///                         {
    ///                             Type = "Service",
    ///                             Identifiers = 
    ///                             {
    ///                                 "ec2.amazonaws.com",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var instance = new Aws.Iam.Role("instance", new Aws.Iam.RoleArgs
    ///         {
    ///             Path = "/system/",
    ///             AssumeRolePolicy = instance_assume_role_policy.Apply(instance_assume_role_policy =&gt; instance_assume_role_policy.Json),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Example of Exclusive Inline Policies
    /// 
    /// This example creates an IAM role with two inline IAM policies. If someone adds another inline policy out-of-band, on the next apply, the provider will remove that policy. If someone deletes these policies out-of-band, the provider will recreate them.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var inlinePolicy = Output.Create(Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
    ///         {
    ///             Statements = 
    ///             {
    ///                 new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
    ///                 {
    ///                     Actions = 
    ///                     {
    ///                         "ec2:DescribeAccountAttributes",
    ///                     },
    ///                     Resources = 
    ///                     {
    ///                         "*",
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var example = new Aws.Iam.Role("example", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = data.Aws_iam_policy_document.Instance_assume_role_policy.Json,
    ///             InlinePolicies = 
    ///             {
    ///                 new Aws.Iam.Inputs.RoleInlinePolicyArgs
    ///                 {
    ///                     Name = "my_inline_policy",
    ///                     Policy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         { "Version", "2012-10-17" },
    ///                         { "Statement", new[]
    ///                             {
    ///                                 new Dictionary&lt;string, object?&gt;
    ///                                 {
    ///                                     { "Action", new[]
    ///                                         {
    ///                                             "ec2:Describe*",
    ///                                         }
    ///                                      },
    ///                                     { "Effect", "Allow" },
    ///                                     { "Resource", "*" },
    ///                                 },
    ///                             }
    ///                          },
    ///                     }),
    ///                 },
    ///                 new Aws.Iam.Inputs.RoleInlinePolicyArgs
    ///                 {
    ///                     Name = "policy-8675309",
    ///                     Policy = inlinePolicy.Apply(inlinePolicy =&gt; inlinePolicy.Json),
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Example of Removing Inline Policies
    /// 
    /// This example creates an IAM role with what appears to be empty IAM `inline_policy` argument instead of using `inline_policy` as a configuration block. The result is that if someone were to add an inline policy out-of-band, on the next apply, the provider will remove that policy.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Iam.Role("example", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = data.Aws_iam_policy_document.Instance_assume_role_policy.Json,
    ///             InlinePolicies = 
    ///             {
    ///                 ,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Example of Exclusive Managed Policies
    /// 
    /// This example creates an IAM role and attaches two managed IAM policies. If someone attaches another managed policy out-of-band, on the next apply, the provider will detach that policy. If someone detaches these policies out-of-band, the provider will attach them again.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var policyOne = new Aws.Iam.Policy("policyOne", new Aws.Iam.PolicyArgs
    ///         {
    ///             Policy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 { "Version", "2012-10-17" },
    ///                 { "Statement", new[]
    ///                     {
    ///                         new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             { "Action", new[]
    ///                                 {
    ///                                     "ec2:Describe*",
    ///                                 }
    ///                              },
    ///                             { "Effect", "Allow" },
    ///                             { "Resource", "*" },
    ///                         },
    ///                     }
    ///                  },
    ///             }),
    ///         });
    ///         var policyTwo = new Aws.Iam.Policy("policyTwo", new Aws.Iam.PolicyArgs
    ///         {
    ///             Policy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 { "Version", "2012-10-17" },
    ///                 { "Statement", new[]
    ///                     {
    ///                         new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             { "Action", new[]
    ///                                 {
    ///                                     "s3:ListAllMyBuckets",
    ///                                     "s3:ListBucket",
    ///                                     "s3:HeadBucket",
    ///                                 }
    ///                              },
    ///                             { "Effect", "Allow" },
    ///                             { "Resource", "*" },
    ///                         },
    ///                     }
    ///                  },
    ///             }),
    ///         });
    ///         var example = new Aws.Iam.Role("example", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = data.Aws_iam_policy_document.Instance_assume_role_policy.Json,
    ///             ManagedPolicyArns = 
    ///             {
    ///                 policyOne.Arn,
    ///                 policyTwo.Arn,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Example of Removing Managed Policies
    /// 
    /// This example creates an IAM role with an empty `managed_policy_arns` argument. If someone attaches a policy out-of-band, on the next apply, the provider will detach that policy.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Iam.Role("example", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = data.Aws_iam_policy_document.Instance_assume_role_policy.Json,
    ///             ManagedPolicyArns = {},
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// IAM Roles can be imported using the `name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:iam/role:Role developer developer_name
    /// ```
    /// </summary>
    [AwsResourceType("aws:iam/role:Role")]
    public partial class Role : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) specifying the role.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Policy that grants an entity permission to assume the role.
        /// </summary>
        [Output("assumeRolePolicy")]
        public Output<string> AssumeRolePolicy { get; private set; } = null!;

        /// <summary>
        /// Creation date of the IAM role.
        /// </summary>
        [Output("createDate")]
        public Output<string> CreateDate { get; private set; } = null!;

        /// <summary>
        /// Description of the role.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
        /// </summary>
        [Output("forceDetachPolicies")]
        public Output<bool?> ForceDetachPolicies { get; private set; } = null!;

        /// <summary>
        /// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. Defined below. If no blocks are configured, the provider will ignore any managing any inline policies in this resource. Configuring one empty block (i.e., `inline_policy {}`) will cause the provider to remove _all_ inline policies.
        /// </summary>
        [Output("inlinePolicies")]
        public Output<ImmutableArray<Outputs.RoleInlinePolicy>> InlinePolicies { get; private set; } = null!;

        /// <summary>
        /// Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, the provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managed_policy_arns = []`) will cause the provider to remove _all_ managed policy attachments.
        /// </summary>
        [Output("managedPolicyArns")]
        public Output<ImmutableArray<string>> ManagedPolicyArns { get; private set; } = null!;

        /// <summary>
        /// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
        /// </summary>
        [Output("maxSessionDuration")]
        public Output<int?> MaxSessionDuration { get; private set; } = null!;

        /// <summary>
        /// Name of the role policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// ARN of the policy that is used to set the permissions boundary for the role.
        /// </summary>
        [Output("permissionsBoundary")]
        public Output<string?> PermissionsBoundary { get; private set; } = null!;

        /// <summary>
        /// Key-value mapping of tags for the IAM role
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Stable and unique string identifying the role.
        /// </summary>
        [Output("uniqueId")]
        public Output<string> UniqueId { get; private set; } = null!;


        /// <summary>
        /// Create a Role resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Role(string name, RoleArgs args, CustomResourceOptions? options = null)
            : base("aws:iam/role:Role", name, args ?? new RoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Role(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
            : base("aws:iam/role:Role", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Role resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Role Get(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
        {
            return new Role(name, id, state, options);
        }
    }

    public sealed class RoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy that grants an entity permission to assume the role.
        /// </summary>
        [Input("assumeRolePolicy", required: true)]
        public Input<string> AssumeRolePolicy { get; set; } = null!;

        /// <summary>
        /// Description of the role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
        /// </summary>
        [Input("forceDetachPolicies")]
        public Input<bool>? ForceDetachPolicies { get; set; }

        [Input("inlinePolicies")]
        private InputList<Inputs.RoleInlinePolicyArgs>? _inlinePolicies;

        /// <summary>
        /// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. Defined below. If no blocks are configured, the provider will ignore any managing any inline policies in this resource. Configuring one empty block (i.e., `inline_policy {}`) will cause the provider to remove _all_ inline policies.
        /// </summary>
        public InputList<Inputs.RoleInlinePolicyArgs> InlinePolicies
        {
            get => _inlinePolicies ?? (_inlinePolicies = new InputList<Inputs.RoleInlinePolicyArgs>());
            set => _inlinePolicies = value;
        }

        [Input("managedPolicyArns")]
        private InputList<string>? _managedPolicyArns;

        /// <summary>
        /// Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, the provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managed_policy_arns = []`) will cause the provider to remove _all_ managed policy attachments.
        /// </summary>
        public InputList<string> ManagedPolicyArns
        {
            get => _managedPolicyArns ?? (_managedPolicyArns = new InputList<string>());
            set => _managedPolicyArns = value;
        }

        /// <summary>
        /// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
        /// </summary>
        [Input("maxSessionDuration")]
        public Input<int>? MaxSessionDuration { get; set; }

        /// <summary>
        /// Name of the role policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// ARN of the policy that is used to set the permissions boundary for the role.
        /// </summary>
        [Input("permissionsBoundary")]
        public Input<string>? PermissionsBoundary { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of tags for the IAM role
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public RoleArgs()
        {
        }
    }

    public sealed class RoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) specifying the role.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Policy that grants an entity permission to assume the role.
        /// </summary>
        [Input("assumeRolePolicy")]
        public Input<string>? AssumeRolePolicy { get; set; }

        /// <summary>
        /// Creation date of the IAM role.
        /// </summary>
        [Input("createDate")]
        public Input<string>? CreateDate { get; set; }

        /// <summary>
        /// Description of the role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
        /// </summary>
        [Input("forceDetachPolicies")]
        public Input<bool>? ForceDetachPolicies { get; set; }

        [Input("inlinePolicies")]
        private InputList<Inputs.RoleInlinePolicyGetArgs>? _inlinePolicies;

        /// <summary>
        /// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. Defined below. If no blocks are configured, the provider will ignore any managing any inline policies in this resource. Configuring one empty block (i.e., `inline_policy {}`) will cause the provider to remove _all_ inline policies.
        /// </summary>
        public InputList<Inputs.RoleInlinePolicyGetArgs> InlinePolicies
        {
            get => _inlinePolicies ?? (_inlinePolicies = new InputList<Inputs.RoleInlinePolicyGetArgs>());
            set => _inlinePolicies = value;
        }

        [Input("managedPolicyArns")]
        private InputList<string>? _managedPolicyArns;

        /// <summary>
        /// Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, the provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managed_policy_arns = []`) will cause the provider to remove _all_ managed policy attachments.
        /// </summary>
        public InputList<string> ManagedPolicyArns
        {
            get => _managedPolicyArns ?? (_managedPolicyArns = new InputList<string>());
            set => _managedPolicyArns = value;
        }

        /// <summary>
        /// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
        /// </summary>
        [Input("maxSessionDuration")]
        public Input<int>? MaxSessionDuration { get; set; }

        /// <summary>
        /// Name of the role policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// ARN of the policy that is used to set the permissions boundary for the role.
        /// </summary>
        [Input("permissionsBoundary")]
        public Input<string>? PermissionsBoundary { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of tags for the IAM role
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Stable and unique string identifying the role.
        /// </summary>
        [Input("uniqueId")]
        public Input<string>? UniqueId { get; set; }

        public RoleState()
        {
        }
    }
}
