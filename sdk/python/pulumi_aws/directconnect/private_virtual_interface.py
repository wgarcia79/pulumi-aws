# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PrivateVirtualInterfaceArgs', 'PrivateVirtualInterface']

@pulumi.input_type
class PrivateVirtualInterfaceArgs:
    def __init__(__self__, *,
                 address_family: pulumi.Input[str],
                 bgp_asn: pulumi.Input[int],
                 connection_id: pulumi.Input[str],
                 vlan: pulumi.Input[int],
                 amazon_address: Optional[pulumi.Input[str]] = None,
                 bgp_auth_key: Optional[pulumi.Input[str]] = None,
                 customer_address: Optional[pulumi.Input[str]] = None,
                 dx_gateway_id: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PrivateVirtualInterface resource.
        :param pulumi.Input[str] address_family: The address family for the BGP peer. `ipv4 ` or `ipv6`.
        :param pulumi.Input[int] bgp_asn: The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
        :param pulumi.Input[str] connection_id: The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
        :param pulumi.Input[int] vlan: The VLAN ID.
        :param pulumi.Input[str] amazon_address: The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
        :param pulumi.Input[str] bgp_auth_key: The authentication key for BGP configuration.
        :param pulumi.Input[str] customer_address: The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
        :param pulumi.Input[str] dx_gateway_id: The ID of the Direct Connect gateway to which to connect the virtual interface.
        :param pulumi.Input[int] mtu: The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
               The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
        :param pulumi.Input[str] name: The name for the virtual interface.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the virtual private gateway to which to connect the virtual interface.
        """
        pulumi.set(__self__, "address_family", address_family)
        pulumi.set(__self__, "bgp_asn", bgp_asn)
        pulumi.set(__self__, "connection_id", connection_id)
        pulumi.set(__self__, "vlan", vlan)
        if amazon_address is not None:
            pulumi.set(__self__, "amazon_address", amazon_address)
        if bgp_auth_key is not None:
            pulumi.set(__self__, "bgp_auth_key", bgp_auth_key)
        if customer_address is not None:
            pulumi.set(__self__, "customer_address", customer_address)
        if dx_gateway_id is not None:
            pulumi.set(__self__, "dx_gateway_id", dx_gateway_id)
        if mtu is not None:
            pulumi.set(__self__, "mtu", mtu)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpn_gateway_id is not None:
            pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> pulumi.Input[str]:
        """
        The address family for the BGP peer. `ipv4 ` or `ipv6`.
        """
        return pulumi.get(self, "address_family")

    @address_family.setter
    def address_family(self, value: pulumi.Input[str]):
        pulumi.set(self, "address_family", value)

    @property
    @pulumi.getter(name="bgpAsn")
    def bgp_asn(self) -> pulumi.Input[int]:
        """
        The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
        """
        return pulumi.get(self, "bgp_asn")

    @bgp_asn.setter
    def bgp_asn(self, value: pulumi.Input[int]):
        pulumi.set(self, "bgp_asn", value)

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> pulumi.Input[str]:
        """
        The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
        """
        return pulumi.get(self, "connection_id")

    @connection_id.setter
    def connection_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "connection_id", value)

    @property
    @pulumi.getter
    def vlan(self) -> pulumi.Input[int]:
        """
        The VLAN ID.
        """
        return pulumi.get(self, "vlan")

    @vlan.setter
    def vlan(self, value: pulumi.Input[int]):
        pulumi.set(self, "vlan", value)

    @property
    @pulumi.getter(name="amazonAddress")
    def amazon_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
        """
        return pulumi.get(self, "amazon_address")

    @amazon_address.setter
    def amazon_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "amazon_address", value)

    @property
    @pulumi.getter(name="bgpAuthKey")
    def bgp_auth_key(self) -> Optional[pulumi.Input[str]]:
        """
        The authentication key for BGP configuration.
        """
        return pulumi.get(self, "bgp_auth_key")

    @bgp_auth_key.setter
    def bgp_auth_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_auth_key", value)

    @property
    @pulumi.getter(name="customerAddress")
    def customer_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
        """
        return pulumi.get(self, "customer_address")

    @customer_address.setter
    def customer_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "customer_address", value)

    @property
    @pulumi.getter(name="dxGatewayId")
    def dx_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Direct Connect gateway to which to connect the virtual interface.
        """
        return pulumi.get(self, "dx_gateway_id")

    @dx_gateway_id.setter
    def dx_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dx_gateway_id", value)

    @property
    @pulumi.getter
    def mtu(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
        The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
        """
        return pulumi.get(self, "mtu")

    @mtu.setter
    def mtu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mtu", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for the virtual interface.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the virtual private gateway to which to connect the virtual interface.
        """
        return pulumi.get(self, "vpn_gateway_id")

    @vpn_gateway_id.setter
    def vpn_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpn_gateway_id", value)


@pulumi.input_type
class _PrivateVirtualInterfaceState:
    def __init__(__self__, *,
                 address_family: Optional[pulumi.Input[str]] = None,
                 amazon_address: Optional[pulumi.Input[str]] = None,
                 amazon_side_asn: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 aws_device: Optional[pulumi.Input[str]] = None,
                 bgp_asn: Optional[pulumi.Input[int]] = None,
                 bgp_auth_key: Optional[pulumi.Input[str]] = None,
                 connection_id: Optional[pulumi.Input[str]] = None,
                 customer_address: Optional[pulumi.Input[str]] = None,
                 dx_gateway_id: Optional[pulumi.Input[str]] = None,
                 jumbo_frame_capable: Optional[pulumi.Input[bool]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vlan: Optional[pulumi.Input[int]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PrivateVirtualInterface resources.
        :param pulumi.Input[str] address_family: The address family for the BGP peer. `ipv4 ` or `ipv6`.
        :param pulumi.Input[str] amazon_address: The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
        :param pulumi.Input[str] arn: The ARN of the virtual interface.
        :param pulumi.Input[str] aws_device: The Direct Connect endpoint on which the virtual interface terminates.
        :param pulumi.Input[int] bgp_asn: The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
        :param pulumi.Input[str] bgp_auth_key: The authentication key for BGP configuration.
        :param pulumi.Input[str] connection_id: The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
        :param pulumi.Input[str] customer_address: The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
        :param pulumi.Input[str] dx_gateway_id: The ID of the Direct Connect gateway to which to connect the virtual interface.
        :param pulumi.Input[bool] jumbo_frame_capable: Indicates whether jumbo frames (9001 MTU) are supported.
        :param pulumi.Input[int] mtu: The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
               The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
        :param pulumi.Input[str] name: The name for the virtual interface.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[int] vlan: The VLAN ID.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the virtual private gateway to which to connect the virtual interface.
        """
        if address_family is not None:
            pulumi.set(__self__, "address_family", address_family)
        if amazon_address is not None:
            pulumi.set(__self__, "amazon_address", amazon_address)
        if amazon_side_asn is not None:
            pulumi.set(__self__, "amazon_side_asn", amazon_side_asn)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if aws_device is not None:
            pulumi.set(__self__, "aws_device", aws_device)
        if bgp_asn is not None:
            pulumi.set(__self__, "bgp_asn", bgp_asn)
        if bgp_auth_key is not None:
            pulumi.set(__self__, "bgp_auth_key", bgp_auth_key)
        if connection_id is not None:
            pulumi.set(__self__, "connection_id", connection_id)
        if customer_address is not None:
            pulumi.set(__self__, "customer_address", customer_address)
        if dx_gateway_id is not None:
            pulumi.set(__self__, "dx_gateway_id", dx_gateway_id)
        if jumbo_frame_capable is not None:
            pulumi.set(__self__, "jumbo_frame_capable", jumbo_frame_capable)
        if mtu is not None:
            pulumi.set(__self__, "mtu", mtu)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if vlan is not None:
            pulumi.set(__self__, "vlan", vlan)
        if vpn_gateway_id is not None:
            pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> Optional[pulumi.Input[str]]:
        """
        The address family for the BGP peer. `ipv4 ` or `ipv6`.
        """
        return pulumi.get(self, "address_family")

    @address_family.setter
    def address_family(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_family", value)

    @property
    @pulumi.getter(name="amazonAddress")
    def amazon_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
        """
        return pulumi.get(self, "amazon_address")

    @amazon_address.setter
    def amazon_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "amazon_address", value)

    @property
    @pulumi.getter(name="amazonSideAsn")
    def amazon_side_asn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "amazon_side_asn")

    @amazon_side_asn.setter
    def amazon_side_asn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "amazon_side_asn", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the virtual interface.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="awsDevice")
    def aws_device(self) -> Optional[pulumi.Input[str]]:
        """
        The Direct Connect endpoint on which the virtual interface terminates.
        """
        return pulumi.get(self, "aws_device")

    @aws_device.setter
    def aws_device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_device", value)

    @property
    @pulumi.getter(name="bgpAsn")
    def bgp_asn(self) -> Optional[pulumi.Input[int]]:
        """
        The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
        """
        return pulumi.get(self, "bgp_asn")

    @bgp_asn.setter
    def bgp_asn(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bgp_asn", value)

    @property
    @pulumi.getter(name="bgpAuthKey")
    def bgp_auth_key(self) -> Optional[pulumi.Input[str]]:
        """
        The authentication key for BGP configuration.
        """
        return pulumi.get(self, "bgp_auth_key")

    @bgp_auth_key.setter
    def bgp_auth_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_auth_key", value)

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
        """
        return pulumi.get(self, "connection_id")

    @connection_id.setter
    def connection_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_id", value)

    @property
    @pulumi.getter(name="customerAddress")
    def customer_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
        """
        return pulumi.get(self, "customer_address")

    @customer_address.setter
    def customer_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "customer_address", value)

    @property
    @pulumi.getter(name="dxGatewayId")
    def dx_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Direct Connect gateway to which to connect the virtual interface.
        """
        return pulumi.get(self, "dx_gateway_id")

    @dx_gateway_id.setter
    def dx_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dx_gateway_id", value)

    @property
    @pulumi.getter(name="jumboFrameCapable")
    def jumbo_frame_capable(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether jumbo frames (9001 MTU) are supported.
        """
        return pulumi.get(self, "jumbo_frame_capable")

    @jumbo_frame_capable.setter
    def jumbo_frame_capable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "jumbo_frame_capable", value)

    @property
    @pulumi.getter
    def mtu(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
        The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
        """
        return pulumi.get(self, "mtu")

    @mtu.setter
    def mtu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mtu", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for the virtual interface.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter
    def vlan(self) -> Optional[pulumi.Input[int]]:
        """
        The VLAN ID.
        """
        return pulumi.get(self, "vlan")

    @vlan.setter
    def vlan(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vlan", value)

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the virtual private gateway to which to connect the virtual interface.
        """
        return pulumi.get(self, "vpn_gateway_id")

    @vpn_gateway_id.setter
    def vpn_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpn_gateway_id", value)


class PrivateVirtualInterface(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 amazon_address: Optional[pulumi.Input[str]] = None,
                 bgp_asn: Optional[pulumi.Input[int]] = None,
                 bgp_auth_key: Optional[pulumi.Input[str]] = None,
                 connection_id: Optional[pulumi.Input[str]] = None,
                 customer_address: Optional[pulumi.Input[str]] = None,
                 dx_gateway_id: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vlan: Optional[pulumi.Input[int]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Direct Connect private virtual interface resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foo = aws.directconnect.PrivateVirtualInterface("foo",
            address_family="ipv4",
            bgp_asn=65352,
            connection_id="dxcon-zzzzzzzz",
            vlan=4094)
        ```

        ## Import

        Direct Connect private virtual interfaces can be imported using the `vif id`, e.g.,

        ```sh
         $ pulumi import aws:directconnect/privateVirtualInterface:PrivateVirtualInterface test dxvif-33cc44dd
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_family: The address family for the BGP peer. `ipv4 ` or `ipv6`.
        :param pulumi.Input[str] amazon_address: The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
        :param pulumi.Input[int] bgp_asn: The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
        :param pulumi.Input[str] bgp_auth_key: The authentication key for BGP configuration.
        :param pulumi.Input[str] connection_id: The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
        :param pulumi.Input[str] customer_address: The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
        :param pulumi.Input[str] dx_gateway_id: The ID of the Direct Connect gateway to which to connect the virtual interface.
        :param pulumi.Input[int] mtu: The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
               The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
        :param pulumi.Input[str] name: The name for the virtual interface.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[int] vlan: The VLAN ID.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the virtual private gateway to which to connect the virtual interface.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateVirtualInterfaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Direct Connect private virtual interface resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foo = aws.directconnect.PrivateVirtualInterface("foo",
            address_family="ipv4",
            bgp_asn=65352,
            connection_id="dxcon-zzzzzzzz",
            vlan=4094)
        ```

        ## Import

        Direct Connect private virtual interfaces can be imported using the `vif id`, e.g.,

        ```sh
         $ pulumi import aws:directconnect/privateVirtualInterface:PrivateVirtualInterface test dxvif-33cc44dd
        ```

        :param str resource_name: The name of the resource.
        :param PrivateVirtualInterfaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateVirtualInterfaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 amazon_address: Optional[pulumi.Input[str]] = None,
                 bgp_asn: Optional[pulumi.Input[int]] = None,
                 bgp_auth_key: Optional[pulumi.Input[str]] = None,
                 connection_id: Optional[pulumi.Input[str]] = None,
                 customer_address: Optional[pulumi.Input[str]] = None,
                 dx_gateway_id: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vlan: Optional[pulumi.Input[int]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateVirtualInterfaceArgs.__new__(PrivateVirtualInterfaceArgs)

            if address_family is None and not opts.urn:
                raise TypeError("Missing required property 'address_family'")
            __props__.__dict__["address_family"] = address_family
            __props__.__dict__["amazon_address"] = amazon_address
            if bgp_asn is None and not opts.urn:
                raise TypeError("Missing required property 'bgp_asn'")
            __props__.__dict__["bgp_asn"] = bgp_asn
            __props__.__dict__["bgp_auth_key"] = bgp_auth_key
            if connection_id is None and not opts.urn:
                raise TypeError("Missing required property 'connection_id'")
            __props__.__dict__["connection_id"] = connection_id
            __props__.__dict__["customer_address"] = customer_address
            __props__.__dict__["dx_gateway_id"] = dx_gateway_id
            __props__.__dict__["mtu"] = mtu
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            if vlan is None and not opts.urn:
                raise TypeError("Missing required property 'vlan'")
            __props__.__dict__["vlan"] = vlan
            __props__.__dict__["vpn_gateway_id"] = vpn_gateway_id
            __props__.__dict__["amazon_side_asn"] = None
            __props__.__dict__["arn"] = None
            __props__.__dict__["aws_device"] = None
            __props__.__dict__["jumbo_frame_capable"] = None
            __props__.__dict__["tags_all"] = None
        super(PrivateVirtualInterface, __self__).__init__(
            'aws:directconnect/privateVirtualInterface:PrivateVirtualInterface',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_family: Optional[pulumi.Input[str]] = None,
            amazon_address: Optional[pulumi.Input[str]] = None,
            amazon_side_asn: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            aws_device: Optional[pulumi.Input[str]] = None,
            bgp_asn: Optional[pulumi.Input[int]] = None,
            bgp_auth_key: Optional[pulumi.Input[str]] = None,
            connection_id: Optional[pulumi.Input[str]] = None,
            customer_address: Optional[pulumi.Input[str]] = None,
            dx_gateway_id: Optional[pulumi.Input[str]] = None,
            jumbo_frame_capable: Optional[pulumi.Input[bool]] = None,
            mtu: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vlan: Optional[pulumi.Input[int]] = None,
            vpn_gateway_id: Optional[pulumi.Input[str]] = None) -> 'PrivateVirtualInterface':
        """
        Get an existing PrivateVirtualInterface resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_family: The address family for the BGP peer. `ipv4 ` or `ipv6`.
        :param pulumi.Input[str] amazon_address: The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
        :param pulumi.Input[str] arn: The ARN of the virtual interface.
        :param pulumi.Input[str] aws_device: The Direct Connect endpoint on which the virtual interface terminates.
        :param pulumi.Input[int] bgp_asn: The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
        :param pulumi.Input[str] bgp_auth_key: The authentication key for BGP configuration.
        :param pulumi.Input[str] connection_id: The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
        :param pulumi.Input[str] customer_address: The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
        :param pulumi.Input[str] dx_gateway_id: The ID of the Direct Connect gateway to which to connect the virtual interface.
        :param pulumi.Input[bool] jumbo_frame_capable: Indicates whether jumbo frames (9001 MTU) are supported.
        :param pulumi.Input[int] mtu: The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
               The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
        :param pulumi.Input[str] name: The name for the virtual interface.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[int] vlan: The VLAN ID.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the virtual private gateway to which to connect the virtual interface.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivateVirtualInterfaceState.__new__(_PrivateVirtualInterfaceState)

        __props__.__dict__["address_family"] = address_family
        __props__.__dict__["amazon_address"] = amazon_address
        __props__.__dict__["amazon_side_asn"] = amazon_side_asn
        __props__.__dict__["arn"] = arn
        __props__.__dict__["aws_device"] = aws_device
        __props__.__dict__["bgp_asn"] = bgp_asn
        __props__.__dict__["bgp_auth_key"] = bgp_auth_key
        __props__.__dict__["connection_id"] = connection_id
        __props__.__dict__["customer_address"] = customer_address
        __props__.__dict__["dx_gateway_id"] = dx_gateway_id
        __props__.__dict__["jumbo_frame_capable"] = jumbo_frame_capable
        __props__.__dict__["mtu"] = mtu
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["vlan"] = vlan
        __props__.__dict__["vpn_gateway_id"] = vpn_gateway_id
        return PrivateVirtualInterface(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> pulumi.Output[str]:
        """
        The address family for the BGP peer. `ipv4 ` or `ipv6`.
        """
        return pulumi.get(self, "address_family")

    @property
    @pulumi.getter(name="amazonAddress")
    def amazon_address(self) -> pulumi.Output[str]:
        """
        The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
        """
        return pulumi.get(self, "amazon_address")

    @property
    @pulumi.getter(name="amazonSideAsn")
    def amazon_side_asn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "amazon_side_asn")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the virtual interface.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsDevice")
    def aws_device(self) -> pulumi.Output[str]:
        """
        The Direct Connect endpoint on which the virtual interface terminates.
        """
        return pulumi.get(self, "aws_device")

    @property
    @pulumi.getter(name="bgpAsn")
    def bgp_asn(self) -> pulumi.Output[int]:
        """
        The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
        """
        return pulumi.get(self, "bgp_asn")

    @property
    @pulumi.getter(name="bgpAuthKey")
    def bgp_auth_key(self) -> pulumi.Output[str]:
        """
        The authentication key for BGP configuration.
        """
        return pulumi.get(self, "bgp_auth_key")

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> pulumi.Output[str]:
        """
        The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
        """
        return pulumi.get(self, "connection_id")

    @property
    @pulumi.getter(name="customerAddress")
    def customer_address(self) -> pulumi.Output[str]:
        """
        The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
        """
        return pulumi.get(self, "customer_address")

    @property
    @pulumi.getter(name="dxGatewayId")
    def dx_gateway_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Direct Connect gateway to which to connect the virtual interface.
        """
        return pulumi.get(self, "dx_gateway_id")

    @property
    @pulumi.getter(name="jumboFrameCapable")
    def jumbo_frame_capable(self) -> pulumi.Output[bool]:
        """
        Indicates whether jumbo frames (9001 MTU) are supported.
        """
        return pulumi.get(self, "jumbo_frame_capable")

    @property
    @pulumi.getter
    def mtu(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
        The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
        """
        return pulumi.get(self, "mtu")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name for the virtual interface.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter
    def vlan(self) -> pulumi.Output[int]:
        """
        The VLAN ID.
        """
        return pulumi.get(self, "vlan")

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the virtual private gateway to which to connect the virtual interface.
        """
        return pulumi.get(self, "vpn_gateway_id")

