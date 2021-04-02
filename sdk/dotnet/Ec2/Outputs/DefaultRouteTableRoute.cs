// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Outputs
{

    [OutputType]
    public sealed class DefaultRouteTableRoute
    {
        /// <summary>
        /// The CIDR block of the route.
        /// </summary>
        public readonly string? CidrBlock;
        /// <summary>
        /// The ID of a managed prefix list destination of the route.
        /// </summary>
        public readonly string? DestinationPrefixListId;
        /// <summary>
        /// Identifier of a VPC Egress Only Internet Gateway.
        /// </summary>
        public readonly string? EgressOnlyGatewayId;
        /// <summary>
        /// Identifier of a VPC internet gateway or a virtual private gateway.
        /// </summary>
        public readonly string? GatewayId;
        /// <summary>
        /// Identifier of an EC2 instance.
        /// </summary>
        public readonly string? InstanceId;
        /// <summary>
        /// The Ipv6 CIDR block of the route
        /// </summary>
        public readonly string? Ipv6CidrBlock;
        /// <summary>
        /// Identifier of a VPC NAT gateway.
        /// </summary>
        public readonly string? NatGatewayId;
        /// <summary>
        /// Identifier of an EC2 network interface.
        /// </summary>
        public readonly string? NetworkInterfaceId;
        /// <summary>
        /// Identifier of an EC2 Transit Gateway.
        /// </summary>
        public readonly string? TransitGatewayId;
        /// <summary>
        /// Identifier of a VPC Endpoint. This route must be removed prior to VPC Endpoint deletion.
        /// </summary>
        public readonly string? VpcEndpointId;
        /// <summary>
        /// Identifier of a VPC peering connection.
        /// </summary>
        public readonly string? VpcPeeringConnectionId;

        [OutputConstructor]
        private DefaultRouteTableRoute(
            string? cidrBlock,

            string? destinationPrefixListId,

            string? egressOnlyGatewayId,

            string? gatewayId,

            string? instanceId,

            string? ipv6CidrBlock,

            string? natGatewayId,

            string? networkInterfaceId,

            string? transitGatewayId,

            string? vpcEndpointId,

            string? vpcPeeringConnectionId)
        {
            CidrBlock = cidrBlock;
            DestinationPrefixListId = destinationPrefixListId;
            EgressOnlyGatewayId = egressOnlyGatewayId;
            GatewayId = gatewayId;
            InstanceId = instanceId;
            Ipv6CidrBlock = ipv6CidrBlock;
            NatGatewayId = natGatewayId;
            NetworkInterfaceId = networkInterfaceId;
            TransitGatewayId = transitGatewayId;
            VpcEndpointId = vpcEndpointId;
            VpcPeeringConnectionId = vpcPeeringConnectionId;
        }
    }
}
