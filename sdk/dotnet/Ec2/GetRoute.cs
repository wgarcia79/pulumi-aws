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
    public static class GetRoute
    {
        /// <summary>
        /// `aws.ec2.Route` provides details about a specific Route.
        /// 
        /// This resource can prove useful when finding the resource associated with a CIDR. For example, finding the peering connection associated with a CIDR value.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// The following example shows how one might use a CIDR value to find a network interface id and use this to create a data source of that network interface.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var config = new Config();
        ///         var subnetId = config.RequireObject&lt;dynamic&gt;("subnetId");
        ///         var selected = Output.Create(Aws.Ec2.GetRouteTable.InvokeAsync(new Aws.Ec2.GetRouteTableArgs
        ///         {
        ///             SubnetId = subnetId,
        ///         }));
        ///         var route = Output.Create(Aws.Ec2.GetRoute.InvokeAsync(new Aws.Ec2.GetRouteArgs
        ///         {
        ///             RouteTableId = aws_route_table.Selected.Id,
        ///             DestinationCidrBlock = "10.0.1.0/24",
        ///         }));
        ///         var @interface = route.Apply(route =&gt; Output.Create(Aws.Ec2.GetNetworkInterface.InvokeAsync(new Aws.Ec2.GetNetworkInterfaceArgs
        ///         {
        ///             Id = route.NetworkInterfaceId,
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRouteResult> InvokeAsync(GetRouteArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRouteResult>("aws:ec2/getRoute:getRoute", args ?? new GetRouteArgs(), options.WithVersion());

        /// <summary>
        /// `aws.ec2.Route` provides details about a specific Route.
        /// 
        /// This resource can prove useful when finding the resource associated with a CIDR. For example, finding the peering connection associated with a CIDR value.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// The following example shows how one might use a CIDR value to find a network interface id and use this to create a data source of that network interface.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var config = new Config();
        ///         var subnetId = config.RequireObject&lt;dynamic&gt;("subnetId");
        ///         var selected = Output.Create(Aws.Ec2.GetRouteTable.InvokeAsync(new Aws.Ec2.GetRouteTableArgs
        ///         {
        ///             SubnetId = subnetId,
        ///         }));
        ///         var route = Output.Create(Aws.Ec2.GetRoute.InvokeAsync(new Aws.Ec2.GetRouteArgs
        ///         {
        ///             RouteTableId = aws_route_table.Selected.Id,
        ///             DestinationCidrBlock = "10.0.1.0/24",
        ///         }));
        ///         var @interface = route.Apply(route =&gt; Output.Create(Aws.Ec2.GetNetworkInterface.InvokeAsync(new Aws.Ec2.GetNetworkInterfaceArgs
        ///         {
        ///             Id = route.NetworkInterfaceId,
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRouteResult> Invoke(GetRouteInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRouteResult>("aws:ec2/getRoute:getRoute", args ?? new GetRouteInvokeArgs(), options.WithVersion());
    }


    public sealed class GetRouteArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// EC2 Carrier Gateway ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("carrierGatewayId")]
        public string? CarrierGatewayId { get; set; }

        /// <summary>
        /// CIDR block of the Route belonging to the Route Table.
        /// </summary>
        [Input("destinationCidrBlock")]
        public string? DestinationCidrBlock { get; set; }

        /// <summary>
        /// IPv6 CIDR block of the Route belonging to the Route Table.
        /// </summary>
        [Input("destinationIpv6CidrBlock")]
        public string? DestinationIpv6CidrBlock { get; set; }

        /// <summary>
        /// The ID of a managed prefix list destination of the Route belonging to the Route Table.
        /// </summary>
        [Input("destinationPrefixListId")]
        public string? DestinationPrefixListId { get; set; }

        /// <summary>
        /// Egress Only Gateway ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("egressOnlyGatewayId")]
        public string? EgressOnlyGatewayId { get; set; }

        /// <summary>
        /// Gateway ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("gatewayId")]
        public string? GatewayId { get; set; }

        /// <summary>
        /// Instance ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// Local Gateway ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("localGatewayId")]
        public string? LocalGatewayId { get; set; }

        /// <summary>
        /// NAT Gateway ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("natGatewayId")]
        public string? NatGatewayId { get; set; }

        /// <summary>
        /// Network Interface ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("networkInterfaceId")]
        public string? NetworkInterfaceId { get; set; }

        /// <summary>
        /// The ID of the specific Route Table containing the Route entry.
        /// </summary>
        [Input("routeTableId", required: true)]
        public string RouteTableId { get; set; } = null!;

        /// <summary>
        /// EC2 Transit Gateway ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("transitGatewayId")]
        public string? TransitGatewayId { get; set; }

        /// <summary>
        /// VPC Peering Connection ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("vpcPeeringConnectionId")]
        public string? VpcPeeringConnectionId { get; set; }

        public GetRouteArgs()
        {
        }
    }

    public sealed class GetRouteInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// EC2 Carrier Gateway ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("carrierGatewayId")]
        public Input<string>? CarrierGatewayId { get; set; }

        /// <summary>
        /// CIDR block of the Route belonging to the Route Table.
        /// </summary>
        [Input("destinationCidrBlock")]
        public Input<string>? DestinationCidrBlock { get; set; }

        /// <summary>
        /// IPv6 CIDR block of the Route belonging to the Route Table.
        /// </summary>
        [Input("destinationIpv6CidrBlock")]
        public Input<string>? DestinationIpv6CidrBlock { get; set; }

        /// <summary>
        /// The ID of a managed prefix list destination of the Route belonging to the Route Table.
        /// </summary>
        [Input("destinationPrefixListId")]
        public Input<string>? DestinationPrefixListId { get; set; }

        /// <summary>
        /// Egress Only Gateway ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("egressOnlyGatewayId")]
        public Input<string>? EgressOnlyGatewayId { get; set; }

        /// <summary>
        /// Gateway ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// Instance ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Local Gateway ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("localGatewayId")]
        public Input<string>? LocalGatewayId { get; set; }

        /// <summary>
        /// NAT Gateway ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("natGatewayId")]
        public Input<string>? NatGatewayId { get; set; }

        /// <summary>
        /// Network Interface ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// The ID of the specific Route Table containing the Route entry.
        /// </summary>
        [Input("routeTableId", required: true)]
        public Input<string> RouteTableId { get; set; } = null!;

        /// <summary>
        /// EC2 Transit Gateway ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("transitGatewayId")]
        public Input<string>? TransitGatewayId { get; set; }

        /// <summary>
        /// VPC Peering Connection ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("vpcPeeringConnectionId")]
        public Input<string>? VpcPeeringConnectionId { get; set; }

        public GetRouteInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRouteResult
    {
        public readonly string CarrierGatewayId;
        public readonly string DestinationCidrBlock;
        public readonly string DestinationIpv6CidrBlock;
        public readonly string DestinationPrefixListId;
        public readonly string EgressOnlyGatewayId;
        public readonly string GatewayId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string LocalGatewayId;
        public readonly string NatGatewayId;
        public readonly string NetworkInterfaceId;
        public readonly string RouteTableId;
        public readonly string TransitGatewayId;
        public readonly string VpcPeeringConnectionId;

        [OutputConstructor]
        private GetRouteResult(
            string carrierGatewayId,

            string destinationCidrBlock,

            string destinationIpv6CidrBlock,

            string destinationPrefixListId,

            string egressOnlyGatewayId,

            string gatewayId,

            string id,

            string instanceId,

            string localGatewayId,

            string natGatewayId,

            string networkInterfaceId,

            string routeTableId,

            string transitGatewayId,

            string vpcPeeringConnectionId)
        {
            CarrierGatewayId = carrierGatewayId;
            DestinationCidrBlock = destinationCidrBlock;
            DestinationIpv6CidrBlock = destinationIpv6CidrBlock;
            DestinationPrefixListId = destinationPrefixListId;
            EgressOnlyGatewayId = egressOnlyGatewayId;
            GatewayId = gatewayId;
            Id = id;
            InstanceId = instanceId;
            LocalGatewayId = localGatewayId;
            NatGatewayId = natGatewayId;
            NetworkInterfaceId = networkInterfaceId;
            RouteTableId = routeTableId;
            TransitGatewayId = transitGatewayId;
            VpcPeeringConnectionId = vpcPeeringConnectionId;
        }
    }
}
