// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws:ec2/ami:Ami":
		r = &Ami{}
	case "aws:ec2/amiCopy:AmiCopy":
		r = &AmiCopy{}
	case "aws:ec2/amiFromInstance:AmiFromInstance":
		r = &AmiFromInstance{}
	case "aws:ec2/amiLaunchPermission:AmiLaunchPermission":
		r = &AmiLaunchPermission{}
	case "aws:ec2/availabilityZoneGroup:AvailabilityZoneGroup":
		r = &AvailabilityZoneGroup{}
	case "aws:ec2/capacityReservation:CapacityReservation":
		r = &CapacityReservation{}
	case "aws:ec2/carrierGateway:CarrierGateway":
		r = &CarrierGateway{}
	case "aws:ec2/customerGateway:CustomerGateway":
		r = &CustomerGateway{}
	case "aws:ec2/dedicatedHost:DedicatedHost":
		r = &DedicatedHost{}
	case "aws:ec2/defaultNetworkAcl:DefaultNetworkAcl":
		r = &DefaultNetworkAcl{}
	case "aws:ec2/defaultRouteTable:DefaultRouteTable":
		r = &DefaultRouteTable{}
	case "aws:ec2/defaultSecurityGroup:DefaultSecurityGroup":
		r = &DefaultSecurityGroup{}
	case "aws:ec2/defaultSubnet:DefaultSubnet":
		r = &DefaultSubnet{}
	case "aws:ec2/defaultVpc:DefaultVpc":
		r = &DefaultVpc{}
	case "aws:ec2/defaultVpcDhcpOptions:DefaultVpcDhcpOptions":
		r = &DefaultVpcDhcpOptions{}
	case "aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway":
		r = &EgressOnlyInternetGateway{}
	case "aws:ec2/eip:Eip":
		r = &Eip{}
	case "aws:ec2/eipAssociation:EipAssociation":
		r = &EipAssociation{}
	case "aws:ec2/fleet:Fleet":
		r = &Fleet{}
	case "aws:ec2/flowLog:FlowLog":
		r = &FlowLog{}
	case "aws:ec2/instance:Instance":
		r = &Instance{}
	case "aws:ec2/internetGateway:InternetGateway":
		r = &InternetGateway{}
	case "aws:ec2/keyPair:KeyPair":
		r = &KeyPair{}
	case "aws:ec2/launchConfiguration:LaunchConfiguration":
		r = &LaunchConfiguration{}
	case "aws:ec2/launchTemplate:LaunchTemplate":
		r = &LaunchTemplate{}
	case "aws:ec2/localGatewayRoute:LocalGatewayRoute":
		r = &LocalGatewayRoute{}
	case "aws:ec2/localGatewayRouteTableVpcAssociation:LocalGatewayRouteTableVpcAssociation":
		r = &LocalGatewayRouteTableVpcAssociation{}
	case "aws:ec2/mainRouteTableAssociation:MainRouteTableAssociation":
		r = &MainRouteTableAssociation{}
	case "aws:ec2/managedPrefixList:ManagedPrefixList":
		r = &ManagedPrefixList{}
	case "aws:ec2/natGateway:NatGateway":
		r = &NatGateway{}
	case "aws:ec2/networkAcl:NetworkAcl":
		r = &NetworkAcl{}
	case "aws:ec2/networkAclRule:NetworkAclRule":
		r = &NetworkAclRule{}
	case "aws:ec2/networkInterface:NetworkInterface":
		r = &NetworkInterface{}
	case "aws:ec2/networkInterfaceAttachment:NetworkInterfaceAttachment":
		r = &NetworkInterfaceAttachment{}
	case "aws:ec2/networkInterfaceSecurityGroupAttachment:NetworkInterfaceSecurityGroupAttachment":
		r = &NetworkInterfaceSecurityGroupAttachment{}
	case "aws:ec2/peeringConnectionOptions:PeeringConnectionOptions":
		r = &PeeringConnectionOptions{}
	case "aws:ec2/placementGroup:PlacementGroup":
		r = &PlacementGroup{}
	case "aws:ec2/proxyProtocolPolicy:ProxyProtocolPolicy":
		r = &ProxyProtocolPolicy{}
	case "aws:ec2/route:Route":
		r = &Route{}
	case "aws:ec2/routeTable:RouteTable":
		r = &RouteTable{}
	case "aws:ec2/routeTableAssociation:RouteTableAssociation":
		r = &RouteTableAssociation{}
	case "aws:ec2/securityGroup:SecurityGroup":
		r = &SecurityGroup{}
	case "aws:ec2/securityGroupRule:SecurityGroupRule":
		r = &SecurityGroupRule{}
	case "aws:ec2/snapshotCreateVolumePermission:SnapshotCreateVolumePermission":
		r = &SnapshotCreateVolumePermission{}
	case "aws:ec2/spotDatafeedSubscription:SpotDatafeedSubscription":
		r = &SpotDatafeedSubscription{}
	case "aws:ec2/spotFleetRequest:SpotFleetRequest":
		r = &SpotFleetRequest{}
	case "aws:ec2/spotInstanceRequest:SpotInstanceRequest":
		r = &SpotInstanceRequest{}
	case "aws:ec2/subnet:Subnet":
		r = &Subnet{}
	case "aws:ec2/tag:Tag":
		r = &Tag{}
	case "aws:ec2/trafficMirrorFilter:TrafficMirrorFilter":
		r = &TrafficMirrorFilter{}
	case "aws:ec2/trafficMirrorFilterRule:TrafficMirrorFilterRule":
		r = &TrafficMirrorFilterRule{}
	case "aws:ec2/trafficMirrorSession:TrafficMirrorSession":
		r = &TrafficMirrorSession{}
	case "aws:ec2/trafficMirrorTarget:TrafficMirrorTarget":
		r = &TrafficMirrorTarget{}
	case "aws:ec2/transitGatewayPeeringAttachmentAccepter:TransitGatewayPeeringAttachmentAccepter":
		r = &TransitGatewayPeeringAttachmentAccepter{}
	case "aws:ec2/volumeAttachment:VolumeAttachment":
		r = &VolumeAttachment{}
	case "aws:ec2/vpc:Vpc":
		r = &Vpc{}
	case "aws:ec2/vpcDhcpOptions:VpcDhcpOptions":
		r = &VpcDhcpOptions{}
	case "aws:ec2/vpcDhcpOptionsAssociation:VpcDhcpOptionsAssociation":
		r = &VpcDhcpOptionsAssociation{}
	case "aws:ec2/vpcEndpoint:VpcEndpoint":
		r = &VpcEndpoint{}
	case "aws:ec2/vpcEndpointConnectionNotification:VpcEndpointConnectionNotification":
		r = &VpcEndpointConnectionNotification{}
	case "aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation":
		r = &VpcEndpointRouteTableAssociation{}
	case "aws:ec2/vpcEndpointService:VpcEndpointService":
		r = &VpcEndpointService{}
	case "aws:ec2/vpcEndpointServiceAllowedPrinciple:VpcEndpointServiceAllowedPrinciple":
		r = &VpcEndpointServiceAllowedPrinciple{}
	case "aws:ec2/vpcEndpointSubnetAssociation:VpcEndpointSubnetAssociation":
		r = &VpcEndpointSubnetAssociation{}
	case "aws:ec2/vpcIpv4CidrBlockAssociation:VpcIpv4CidrBlockAssociation":
		r = &VpcIpv4CidrBlockAssociation{}
	case "aws:ec2/vpcPeeringConnection:VpcPeeringConnection":
		r = &VpcPeeringConnection{}
	case "aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter":
		r = &VpcPeeringConnectionAccepter{}
	case "aws:ec2/vpnConnection:VpnConnection":
		r = &VpnConnection{}
	case "aws:ec2/vpnConnectionRoute:VpnConnectionRoute":
		r = &VpnConnectionRoute{}
	case "aws:ec2/vpnGateway:VpnGateway":
		r = &VpnGateway{}
	case "aws:ec2/vpnGatewayAttachment:VpnGatewayAttachment":
		r = &VpnGatewayAttachment{}
	case "aws:ec2/vpnGatewayRoutePropagation:VpnGatewayRoutePropagation":
		r = &VpnGatewayRoutePropagation{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/ami",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/amiCopy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/amiFromInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/amiLaunchPermission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/availabilityZoneGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/capacityReservation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/carrierGateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/customerGateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/dedicatedHost",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/defaultNetworkAcl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/defaultRouteTable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/defaultSecurityGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/defaultSubnet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/defaultVpc",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/defaultVpcDhcpOptions",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/egressOnlyInternetGateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/eip",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/eipAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/fleet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/flowLog",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/internetGateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/keyPair",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/launchConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/launchTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/localGatewayRoute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/localGatewayRouteTableVpcAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/mainRouteTableAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/managedPrefixList",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/natGateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/networkAcl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/networkAclRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/networkInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/networkInterfaceAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/networkInterfaceSecurityGroupAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/peeringConnectionOptions",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/placementGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/proxyProtocolPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/route",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/routeTable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/routeTableAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/securityGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/securityGroupRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/snapshotCreateVolumePermission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/spotDatafeedSubscription",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/spotFleetRequest",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/spotInstanceRequest",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/subnet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/tag",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/trafficMirrorFilter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/trafficMirrorFilterRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/trafficMirrorSession",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/trafficMirrorTarget",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/transitGatewayPeeringAttachmentAccepter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/volumeAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpc",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpcDhcpOptions",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpcDhcpOptionsAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpcEndpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpcEndpointConnectionNotification",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpcEndpointRouteTableAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpcEndpointService",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpcEndpointServiceAllowedPrinciple",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpcEndpointSubnetAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpcIpv4CidrBlockAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpcPeeringConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpcPeeringConnectionAccepter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpnConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpnConnectionRoute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpnGateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpnGatewayAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2/vpnGatewayRoutePropagation",
		&module{version},
	)
}
