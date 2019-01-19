# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class EipAssociation(pulumi.CustomResource):
    allocation_id: pulumi.Output[str]
    """
    The allocation ID. This is required for EC2-VPC.
    """
    allow_reassociation: pulumi.Output[bool]
    """
    Whether to allow an Elastic IP to
    be re-associated. Defaults to `true` in VPC.
    """
    instance_id: pulumi.Output[str]
    """
    The ID of the instance. This is required for
    EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
    network interface ID, but not both. The operation fails if you specify an
    instance ID unless exactly one network interface is attached.
    """
    network_interface_id: pulumi.Output[str]
    """
    The ID of the network interface. If the
    instance has more than one network interface, you must specify a network
    interface ID.
    """
    private_ip_address: pulumi.Output[str]
    """
    The primary or secondary private IP address
    to associate with the Elastic IP address. If no private IP address is
    specified, the Elastic IP address is associated with the primary private IP
    address.
    """
    public_ip: pulumi.Output[str]
    """
    The Elastic IP address. This is required for EC2-Classic.
    """
    def __init__(__self__, __name__, __opts__=None, allocation_id=None, allow_reassociation=None, instance_id=None, network_interface_id=None, private_ip_address=None, public_ip=None):
        """
        Provides an AWS EIP Association as a top level resource, to associate and
        disassociate Elastic IPs from AWS Instances and Network Interfaces.
        
        > **NOTE:** Do not use this resource to associate an EIP to `aws_lb` or `aws_nat_gateway` resources. Instead use the `allocation_id` available in those resources to allow AWS to manage the association, otherwise you will see `AuthFailure` errors.
        
        > **NOTE:** `aws_eip_association` is useful in scenarios where EIPs are either
        pre-existing or distributed to customers or users and therefore cannot be changed.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] allocation_id: The allocation ID. This is required for EC2-VPC.
        :param pulumi.Input[bool] allow_reassociation: Whether to allow an Elastic IP to
               be re-associated. Defaults to `true` in VPC.
        :param pulumi.Input[str] instance_id: The ID of the instance. This is required for
               EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
               network interface ID, but not both. The operation fails if you specify an
               instance ID unless exactly one network interface is attached.
        :param pulumi.Input[str] network_interface_id: The ID of the network interface. If the
               instance has more than one network interface, you must specify a network
               interface ID.
        :param pulumi.Input[str] private_ip_address: The primary or secondary private IP address
               to associate with the Elastic IP address. If no private IP address is
               specified, the Elastic IP address is associated with the primary private IP
               address.
        :param pulumi.Input[str] public_ip: The Elastic IP address. This is required for EC2-Classic.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['allocation_id'] = allocation_id

        __props__['allow_reassociation'] = allow_reassociation

        __props__['instance_id'] = instance_id

        __props__['network_interface_id'] = network_interface_id

        __props__['private_ip_address'] = private_ip_address

        __props__['public_ip'] = public_ip

        super(EipAssociation, __self__).__init__(
            'aws:ec2/eipAssociation:EipAssociation',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

