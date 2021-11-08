// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.RedShift
{
    /// <summary>
    /// ## Example Usage
    /// ### Pause Cluster Action
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleRole = new Aws.Iam.Role("exampleRole", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = @"{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {
    ///       ""Action"": ""sts:AssumeRole"",
    ///       ""Principal"": {
    ///         ""Service"": [
    ///           ""scheduler.redshift.amazonaws.com""
    ///         ]
    ///       },
    ///       ""Effect"": ""Allow"",
    ///       ""Sid"": """"
    ///     }
    ///   ]
    /// }
    /// ",
    ///         });
    ///         var examplePolicy = new Aws.Iam.Policy("examplePolicy", new Aws.Iam.PolicyArgs
    ///         {
    ///             Policy = @"{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///       {
    ///           ""Sid"": ""VisualEditor0"",
    ///           ""Effect"": ""Allow"",
    ///           ""Action"": [
    ///               ""redshift:PauseCluster"",
    ///               ""redshift:ResumeCluster"",
    ///               ""redshift:ResizeCluster""
    ///           ],
    ///           ""Resource"": ""*""
    ///       }
    ///   ]
    /// }
    /// ",
    ///         });
    ///         var exampleRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("exampleRolePolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
    ///         {
    ///             PolicyArn = examplePolicy.Arn,
    ///             Role = exampleRole.Name,
    ///         });
    ///         var exampleScheduledAction = new Aws.RedShift.ScheduledAction("exampleScheduledAction", new Aws.RedShift.ScheduledActionArgs
    ///         {
    ///             Schedule = "cron(00 23 * * ? *)",
    ///             IamRole = exampleRole.Arn,
    ///             TargetAction = new Aws.RedShift.Inputs.ScheduledActionTargetActionArgs
    ///             {
    ///                 PauseCluster = new Aws.RedShift.Inputs.ScheduledActionTargetActionPauseClusterArgs
    ///                 {
    ///                     ClusterIdentifier = "tf-redshift001",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Resize Cluster Action
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.RedShift.ScheduledAction("example", new Aws.RedShift.ScheduledActionArgs
    ///         {
    ///             Schedule = "cron(00 23 * * ? *)",
    ///             IamRole = aws_iam_role.Example.Arn,
    ///             TargetAction = new Aws.RedShift.Inputs.ScheduledActionTargetActionArgs
    ///             {
    ///                 ResizeCluster = new Aws.RedShift.Inputs.ScheduledActionTargetActionResizeClusterArgs
    ///                 {
    ///                     ClusterIdentifier = "tf-redshift001",
    ///                     ClusterType = "multi-node",
    ///                     NodeType = "dc1.large",
    ///                     NumberOfNodes = 2,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Redshift Scheduled Action can be imported using the `name`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:redshift/scheduledAction:ScheduledAction example tf-redshift-scheduled-action
    /// ```
    /// </summary>
    [AwsResourceType("aws:redshift/scheduledAction:ScheduledAction")]
    public partial class ScheduledAction : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the scheduled action.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether to enable the scheduled action. Default is `true` .
        /// </summary>
        [Output("enable")]
        public Output<bool?> Enable { get; private set; } = null!;

        /// <summary>
        /// The end time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
        /// </summary>
        [Output("endTime")]
        public Output<string?> EndTime { get; private set; } = null!;

        /// <summary>
        /// The IAM role to assume to run the scheduled action.
        /// </summary>
        [Output("iamRole")]
        public Output<string> IamRole { get; private set; } = null!;

        /// <summary>
        /// The scheduled action name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The schedule of action. The schedule is defined format of "at expression" or "cron expression", for example `at(2016-03-04T17:27:00)` or `cron(0 10 ? * MON *)`. See [Scheduled Action](https://docs.aws.amazon.com/redshift/latest/APIReference/API_ScheduledAction.html) for more information.
        /// </summary>
        [Output("schedule")]
        public Output<string> Schedule { get; private set; } = null!;

        /// <summary>
        /// The start time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
        /// </summary>
        [Output("startTime")]
        public Output<string?> StartTime { get; private set; } = null!;

        /// <summary>
        /// Target action. Documented below.
        /// </summary>
        [Output("targetAction")]
        public Output<Outputs.ScheduledActionTargetAction> TargetAction { get; private set; } = null!;


        /// <summary>
        /// Create a ScheduledAction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScheduledAction(string name, ScheduledActionArgs args, CustomResourceOptions? options = null)
            : base("aws:redshift/scheduledAction:ScheduledAction", name, args ?? new ScheduledActionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScheduledAction(string name, Input<string> id, ScheduledActionState? state = null, CustomResourceOptions? options = null)
            : base("aws:redshift/scheduledAction:ScheduledAction", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ScheduledAction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScheduledAction Get(string name, Input<string> id, ScheduledActionState? state = null, CustomResourceOptions? options = null)
        {
            return new ScheduledAction(name, id, state, options);
        }
    }

    public sealed class ScheduledActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the scheduled action.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to enable the scheduled action. Default is `true` .
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// The end time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// The IAM role to assume to run the scheduled action.
        /// </summary>
        [Input("iamRole", required: true)]
        public Input<string> IamRole { get; set; } = null!;

        /// <summary>
        /// The scheduled action name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The schedule of action. The schedule is defined format of "at expression" or "cron expression", for example `at(2016-03-04T17:27:00)` or `cron(0 10 ? * MON *)`. See [Scheduled Action](https://docs.aws.amazon.com/redshift/latest/APIReference/API_ScheduledAction.html) for more information.
        /// </summary>
        [Input("schedule", required: true)]
        public Input<string> Schedule { get; set; } = null!;

        /// <summary>
        /// The start time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// Target action. Documented below.
        /// </summary>
        [Input("targetAction", required: true)]
        public Input<Inputs.ScheduledActionTargetActionArgs> TargetAction { get; set; } = null!;

        public ScheduledActionArgs()
        {
        }
    }

    public sealed class ScheduledActionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the scheduled action.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to enable the scheduled action. Default is `true` .
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// The end time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// The IAM role to assume to run the scheduled action.
        /// </summary>
        [Input("iamRole")]
        public Input<string>? IamRole { get; set; }

        /// <summary>
        /// The scheduled action name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The schedule of action. The schedule is defined format of "at expression" or "cron expression", for example `at(2016-03-04T17:27:00)` or `cron(0 10 ? * MON *)`. See [Scheduled Action](https://docs.aws.amazon.com/redshift/latest/APIReference/API_ScheduledAction.html) for more information.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// The start time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// Target action. Documented below.
        /// </summary>
        [Input("targetAction")]
        public Input<Inputs.ScheduledActionTargetActionGetArgs>? TargetAction { get; set; }

        public ScheduledActionState()
        {
        }
    }
}
