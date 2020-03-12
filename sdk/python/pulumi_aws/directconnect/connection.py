# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Connection(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the connection.
    """
    aws_device: pulumi.Output[str]
    """
    The Direct Connect endpoint on which the physical connection terminates.
    """
    bandwidth: pulumi.Output[str]
    """
    The bandwidth of the connection. Available values: 1Gbps, 10Gbps. Case sensitive.
    """
    has_logical_redundancy: pulumi.Output[str]
    """
    Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
    """
    jumbo_frame_capable: pulumi.Output[bool]
    """
    Boolean value representing if jumbo frames have been enabled for this connection.
    """
    location: pulumi.Output[str]
    """
    The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
    """
    name: pulumi.Output[str]
    """
    The name of the connection.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, bandwidth=None, location=None, name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Connection of Direct Connect.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dx_connection.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bandwidth: The bandwidth of the connection. Available values: 1Gbps, 10Gbps. Case sensitive.
        :param pulumi.Input[str] location: The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        :param pulumi.Input[str] name: The name of the connection.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if bandwidth is None:
                raise TypeError("Missing required property 'bandwidth'")
            __props__['bandwidth'] = bandwidth
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['aws_device'] = None
            __props__['has_logical_redundancy'] = None
            __props__['jumbo_frame_capable'] = None
        super(Connection, __self__).__init__(
            'aws:directconnect/connection:Connection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, aws_device=None, bandwidth=None, has_logical_redundancy=None, jumbo_frame_capable=None, location=None, name=None, tags=None):
        """
        Get an existing Connection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the connection.
        :param pulumi.Input[str] aws_device: The Direct Connect endpoint on which the physical connection terminates.
        :param pulumi.Input[str] bandwidth: The bandwidth of the connection. Available values: 1Gbps, 10Gbps. Case sensitive.
        :param pulumi.Input[str] has_logical_redundancy: Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
        :param pulumi.Input[bool] jumbo_frame_capable: Boolean value representing if jumbo frames have been enabled for this connection.
        :param pulumi.Input[str] location: The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        :param pulumi.Input[str] name: The name of the connection.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["aws_device"] = aws_device
        __props__["bandwidth"] = bandwidth
        __props__["has_logical_redundancy"] = has_logical_redundancy
        __props__["jumbo_frame_capable"] = jumbo_frame_capable
        __props__["location"] = location
        __props__["name"] = name
        __props__["tags"] = tags
        return Connection(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

