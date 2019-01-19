# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetVpcResult(object):
    """
    A collection of values returned by getVpc.
    """
    def __init__(__self__, arn=None, cidr_block=None, cidr_block_associations=None, default=None, dhcp_options_id=None, enable_dns_hostnames=None, enable_dns_support=None, id=None, instance_tenancy=None, ipv6_association_id=None, ipv6_cidr_block=None, main_route_table_id=None, owner_id=None, state=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError('Expected argument arn to be a str')
        __self__.arn = arn
        """
        Amazon Resource Name (ARN) of VPC
        """
        if cidr_block and not isinstance(cidr_block, str):
            raise TypeError('Expected argument cidr_block to be a str')
        __self__.cidr_block = cidr_block
        """
        The CIDR block for the association.
        """
        if cidr_block_associations and not isinstance(cidr_block_associations, list):
            raise TypeError('Expected argument cidr_block_associations to be a list')
        __self__.cidr_block_associations = cidr_block_associations
        if default and not isinstance(default, bool):
            raise TypeError('Expected argument default to be a bool')
        __self__.default = default
        if dhcp_options_id and not isinstance(dhcp_options_id, str):
            raise TypeError('Expected argument dhcp_options_id to be a str')
        __self__.dhcp_options_id = dhcp_options_id
        if enable_dns_hostnames and not isinstance(enable_dns_hostnames, bool):
            raise TypeError('Expected argument enable_dns_hostnames to be a bool')
        __self__.enable_dns_hostnames = enable_dns_hostnames
        """
        Whether or not the VPC has DNS hostname support
        """
        if enable_dns_support and not isinstance(enable_dns_support, bool):
            raise TypeError('Expected argument enable_dns_support to be a bool')
        __self__.enable_dns_support = enable_dns_support
        """
        Whether or not the VPC has DNS support
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        if instance_tenancy and not isinstance(instance_tenancy, str):
            raise TypeError('Expected argument instance_tenancy to be a str')
        __self__.instance_tenancy = instance_tenancy
        """
        The allowed tenancy of instances launched into the
        selected VPC. May be any of `"default"`, `"dedicated"`, or `"host"`.
        """
        if ipv6_association_id and not isinstance(ipv6_association_id, str):
            raise TypeError('Expected argument ipv6_association_id to be a str')
        __self__.ipv6_association_id = ipv6_association_id
        """
        The association ID for the IPv6 CIDR block.
        """
        if ipv6_cidr_block and not isinstance(ipv6_cidr_block, str):
            raise TypeError('Expected argument ipv6_cidr_block to be a str')
        __self__.ipv6_cidr_block = ipv6_cidr_block
        """
        The IPv6 CIDR block.
        """
        if main_route_table_id and not isinstance(main_route_table_id, str):
            raise TypeError('Expected argument main_route_table_id to be a str')
        __self__.main_route_table_id = main_route_table_id
        """
        The ID of the main route table associated with this VPC.
        """
        if owner_id and not isinstance(owner_id, str):
            raise TypeError('Expected argument owner_id to be a str')
        __self__.owner_id = owner_id
        """
        The ID of the AWS account that owns the VPC.
        """
        if state and not isinstance(state, str):
            raise TypeError('Expected argument state to be a str')
        __self__.state = state
        """
        The State of the association.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError('Expected argument tags to be a dict')
        __self__.tags = tags

async def get_vpc(cidr_block=None, default=None, dhcp_options_id=None, filters=None, id=None, state=None, tags=None):
    """
    `aws_vpc` provides details about a specific VPC.
    
    This resource can prove useful when a module accepts a vpc id as
    an input variable and needs to, for example, determine the CIDR block of that
    VPC.
    """
    __args__ = dict()

    __args__['cidrBlock'] = cidr_block
    __args__['default'] = default
    __args__['dhcpOptionsId'] = dhcp_options_id
    __args__['filters'] = filters
    __args__['id'] = id
    __args__['state'] = state
    __args__['tags'] = tags
    __ret__ = await pulumi.runtime.invoke('aws:ec2/getVpc:getVpc', __args__)

    return GetVpcResult(
        arn=__ret__.get('arn'),
        cidr_block=__ret__.get('cidrBlock'),
        cidr_block_associations=__ret__.get('cidrBlockAssociations'),
        default=__ret__.get('default'),
        dhcp_options_id=__ret__.get('dhcpOptionsId'),
        enable_dns_hostnames=__ret__.get('enableDnsHostnames'),
        enable_dns_support=__ret__.get('enableDnsSupport'),
        id=__ret__.get('id'),
        instance_tenancy=__ret__.get('instanceTenancy'),
        ipv6_association_id=__ret__.get('ipv6AssociationId'),
        ipv6_cidr_block=__ret__.get('ipv6CidrBlock'),
        main_route_table_id=__ret__.get('mainRouteTableId'),
        owner_id=__ret__.get('ownerId'),
        state=__ret__.get('state'),
        tags=__ret__.get('tags'))
