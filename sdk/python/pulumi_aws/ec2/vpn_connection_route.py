# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['VpnConnectionRoute']


class VpnConnectionRoute(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 vpn_connection_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a static route between a VPN connection and a customer gateway.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        vpc = aws.ec2.Vpc("vpc", cidr_block="10.0.0.0/16")
        vpn_gateway = aws.ec2.VpnGateway("vpnGateway", vpc_id=vpc.id)
        customer_gateway = aws.ec2.CustomerGateway("customerGateway",
            bgp_asn="65000",
            ip_address="172.0.0.1",
            type="ipsec.1")
        main = aws.ec2.VpnConnection("main",
            customer_gateway_id=customer_gateway.id,
            static_routes_only=True,
            type="ipsec.1",
            vpn_gateway_id=vpn_gateway.id)
        office = aws.ec2.VpnConnectionRoute("office",
            destination_cidr_block="192.168.10.0/24",
            vpn_connection_id=main.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidr_block: The CIDR block associated with the local subnet of the customer network.
        :param pulumi.Input[str] vpn_connection_id: The ID of the VPN connection.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if destination_cidr_block is None:
                raise TypeError("Missing required property 'destination_cidr_block'")
            __props__['destination_cidr_block'] = destination_cidr_block
            if vpn_connection_id is None:
                raise TypeError("Missing required property 'vpn_connection_id'")
            __props__['vpn_connection_id'] = vpn_connection_id
        super(VpnConnectionRoute, __self__).__init__(
            'aws:ec2/vpnConnectionRoute:VpnConnectionRoute',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            destination_cidr_block: Optional[pulumi.Input[str]] = None,
            vpn_connection_id: Optional[pulumi.Input[str]] = None) -> 'VpnConnectionRoute':
        """
        Get an existing VpnConnectionRoute resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidr_block: The CIDR block associated with the local subnet of the customer network.
        :param pulumi.Input[str] vpn_connection_id: The ID of the VPN connection.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["destination_cidr_block"] = destination_cidr_block
        __props__["vpn_connection_id"] = vpn_connection_id
        return VpnConnectionRoute(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> str:
        """
        The CIDR block associated with the local subnet of the customer network.
        """
        return pulumi.get(self, "destination_cidr_block")

    @property
    @pulumi.getter(name="vpnConnectionId")
    def vpn_connection_id(self) -> str:
        """
        The ID of the VPN connection.
        """
        return pulumi.get(self, "vpn_connection_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

