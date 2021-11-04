// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Aws.Ec2
{
    public static class GetNetworkInterfaces
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// The following shows outputing all network interface ids in a region.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var exampleNetworkInterfaces = Output.Create(Aws.Ec2.GetNetworkInterfaces.InvokeAsync());
        ///         this.Example = exampleNetworkInterfaces.Apply(exampleNetworkInterfaces =&gt; exampleNetworkInterfaces.Ids);
        ///     }
        /// 
        ///     [Output("example")]
        ///     public Output&lt;string&gt; Example { get; set; }
        /// }
        /// ```
        /// 
        /// The following example retrieves a list of all network interface ids with a custom tag of `Name` set to a value of `test`.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.Ec2.GetNetworkInterfaces.InvokeAsync(new Aws.Ec2.GetNetworkInterfacesArgs
        ///         {
        ///             Tags = 
        ///             {
        ///                 { "Name", "test" },
        ///             },
        ///         }));
        ///         this.Example1 = example.Apply(example =&gt; example.Ids);
        ///     }
        /// 
        ///     [Output("example1")]
        ///     public Output&lt;string&gt; Example1 { get; set; }
        /// }
        /// ```
        /// 
        /// The following example retrieves a network interface ids which associated
        /// with specific subnet.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var exampleNetworkInterfaces = Output.Create(Aws.Ec2.GetNetworkInterfaces.InvokeAsync(new Aws.Ec2.GetNetworkInterfacesArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Aws.Ec2.Inputs.GetNetworkInterfacesFilterArgs
        ///                 {
        ///                     Name = "subnet-id",
        ///                     Values = 
        ///                     {
        ///                         aws_subnet.Test.Id,
        ///                     },
        ///                 },
        ///             },
        ///         }));
        ///         this.Example = exampleNetworkInterfaces.Apply(exampleNetworkInterfaces =&gt; exampleNetworkInterfaces.Ids);
        ///     }
        /// 
        ///     [Output("example")]
        ///     public Output&lt;string&gt; Example { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNetworkInterfacesResult> InvokeAsync(GetNetworkInterfacesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkInterfacesResult>("aws:ec2/getNetworkInterfaces:getNetworkInterfaces", args ?? new GetNetworkInterfacesArgs(), options.WithVersion());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// The following shows outputing all network interface ids in a region.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var exampleNetworkInterfaces = Output.Create(Aws.Ec2.GetNetworkInterfaces.InvokeAsync());
        ///         this.Example = exampleNetworkInterfaces.Apply(exampleNetworkInterfaces =&gt; exampleNetworkInterfaces.Ids);
        ///     }
        /// 
        ///     [Output("example")]
        ///     public Output&lt;string&gt; Example { get; set; }
        /// }
        /// ```
        /// 
        /// The following example retrieves a list of all network interface ids with a custom tag of `Name` set to a value of `test`.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.Ec2.GetNetworkInterfaces.InvokeAsync(new Aws.Ec2.GetNetworkInterfacesArgs
        ///         {
        ///             Tags = 
        ///             {
        ///                 { "Name", "test" },
        ///             },
        ///         }));
        ///         this.Example1 = example.Apply(example =&gt; example.Ids);
        ///     }
        /// 
        ///     [Output("example1")]
        ///     public Output&lt;string&gt; Example1 { get; set; }
        /// }
        /// ```
        /// 
        /// The following example retrieves a network interface ids which associated
        /// with specific subnet.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var exampleNetworkInterfaces = Output.Create(Aws.Ec2.GetNetworkInterfaces.InvokeAsync(new Aws.Ec2.GetNetworkInterfacesArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Aws.Ec2.Inputs.GetNetworkInterfacesFilterArgs
        ///                 {
        ///                     Name = "subnet-id",
        ///                     Values = 
        ///                     {
        ///                         aws_subnet.Test.Id,
        ///                     },
        ///                 },
        ///             },
        ///         }));
        ///         this.Example = exampleNetworkInterfaces.Apply(exampleNetworkInterfaces =&gt; exampleNetworkInterfaces.Ids);
        ///     }
        /// 
        ///     [Output("example")]
        ///     public Output&lt;string&gt; Example { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNetworkInterfacesResult> Invoke(GetNetworkInterfacesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNetworkInterfacesResult>("aws:ec2/getNetworkInterfaces:getNetworkInterfaces", args ?? new GetNetworkInterfacesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetNetworkInterfacesArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetNetworkInterfacesFilterArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public List<Inputs.GetNetworkInterfacesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetNetworkInterfacesFilterArgs>());
            set => _filters = value;
        }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// A map of tags, each pair of which must exactly match
        /// a pair on the desired network interfaces.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetNetworkInterfacesArgs()
        {
        }
    }

    public sealed class GetNetworkInterfacesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetNetworkInterfacesFilterInputArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public InputList<Inputs.GetNetworkInterfacesFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetNetworkInterfacesFilterInputArgs>());
            set => _filters = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags, each pair of which must exactly match
        /// a pair on the desired network interfaces.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetNetworkInterfacesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNetworkInterfacesResult
    {
        public readonly ImmutableArray<Outputs.GetNetworkInterfacesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of all the network interface ids found. This data source will fail if none are found.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetNetworkInterfacesResult(
            ImmutableArray<Outputs.GetNetworkInterfacesFilterResult> filters,

            string id,

            ImmutableArray<string> ids,

            ImmutableDictionary<string, string> tags)
        {
            Filters = filters;
            Id = id;
            Ids = ids;
            Tags = tags;
        }
    }
}
