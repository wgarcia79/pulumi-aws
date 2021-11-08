// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Waf
{
    /// <summary>
    /// Provides a WAF XSS Match Set Resource
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
    ///         var xssMatchSet = new Aws.Waf.XssMatchSet("xssMatchSet", new Aws.Waf.XssMatchSetArgs
    ///         {
    ///             XssMatchTuples = 
    ///             {
    ///                 new Aws.Waf.Inputs.XssMatchSetXssMatchTupleArgs
    ///                 {
    ///                     FieldToMatch = new Aws.Waf.Inputs.XssMatchSetXssMatchTupleFieldToMatchArgs
    ///                     {
    ///                         Type = "URI",
    ///                     },
    ///                     TextTransformation = "NONE",
    ///                 },
    ///                 new Aws.Waf.Inputs.XssMatchSetXssMatchTupleArgs
    ///                 {
    ///                     FieldToMatch = new Aws.Waf.Inputs.XssMatchSetXssMatchTupleFieldToMatchArgs
    ///                     {
    ///                         Type = "QUERY_STRING",
    ///                     },
    ///                     TextTransformation = "NONE",
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
    /// WAF XSS Match Set can be imported using their ID, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:waf/xssMatchSet:XssMatchSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
    /// ```
    /// </summary>
    [AwsResourceType("aws:waf/xssMatchSet:XssMatchSet")]
    public partial class XssMatchSet : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name or description of the SizeConstraintSet.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The parts of web requests that you want to inspect for cross-site scripting attacks.
        /// </summary>
        [Output("xssMatchTuples")]
        public Output<ImmutableArray<Outputs.XssMatchSetXssMatchTuple>> XssMatchTuples { get; private set; } = null!;


        /// <summary>
        /// Create a XssMatchSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public XssMatchSet(string name, XssMatchSetArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:waf/xssMatchSet:XssMatchSet", name, args ?? new XssMatchSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private XssMatchSet(string name, Input<string> id, XssMatchSetState? state = null, CustomResourceOptions? options = null)
            : base("aws:waf/xssMatchSet:XssMatchSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing XssMatchSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static XssMatchSet Get(string name, Input<string> id, XssMatchSetState? state = null, CustomResourceOptions? options = null)
        {
            return new XssMatchSet(name, id, state, options);
        }
    }

    public sealed class XssMatchSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name or description of the SizeConstraintSet.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("xssMatchTuples")]
        private InputList<Inputs.XssMatchSetXssMatchTupleArgs>? _xssMatchTuples;

        /// <summary>
        /// The parts of web requests that you want to inspect for cross-site scripting attacks.
        /// </summary>
        public InputList<Inputs.XssMatchSetXssMatchTupleArgs> XssMatchTuples
        {
            get => _xssMatchTuples ?? (_xssMatchTuples = new InputList<Inputs.XssMatchSetXssMatchTupleArgs>());
            set => _xssMatchTuples = value;
        }

        public XssMatchSetArgs()
        {
        }
    }

    public sealed class XssMatchSetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name or description of the SizeConstraintSet.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("xssMatchTuples")]
        private InputList<Inputs.XssMatchSetXssMatchTupleGetArgs>? _xssMatchTuples;

        /// <summary>
        /// The parts of web requests that you want to inspect for cross-site scripting attacks.
        /// </summary>
        public InputList<Inputs.XssMatchSetXssMatchTupleGetArgs> XssMatchTuples
        {
            get => _xssMatchTuples ?? (_xssMatchTuples = new InputList<Inputs.XssMatchSetXssMatchTupleGetArgs>());
            set => _xssMatchTuples = value;
        }

        public XssMatchSetState()
        {
        }
    }
}
