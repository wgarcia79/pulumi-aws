# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ConnectionArgs', 'Connection']

@pulumi.input_type
class ConnectionArgs:
    def __init__(__self__, *,
                 bandwidth: pulumi.Input[str],
                 location: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Connection resource.
        :param pulumi.Input[str] bandwidth: The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
        :param pulumi.Input[str] location: The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        :param pulumi.Input[str] name: The name of the connection.
        :param pulumi.Input[str] provider_name: The name of the service provider associated with the connection.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        pulumi.set(__self__, "bandwidth", bandwidth)
        pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if provider_name is not None:
            pulumi.set(__self__, "provider_name", provider_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Input[str]:
        """
        The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: pulumi.Input[str]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the connection.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the service provider associated with the connection.
        """
        return pulumi.get(self, "provider_name")

    @provider_name.setter
    def provider_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


@pulumi.input_type
class _ConnectionState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 aws_device: Optional[pulumi.Input[str]] = None,
                 bandwidth: Optional[pulumi.Input[str]] = None,
                 has_logical_redundancy: Optional[pulumi.Input[str]] = None,
                 jumbo_frame_capable: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_account_id: Optional[pulumi.Input[str]] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Connection resources.
        :param pulumi.Input[str] arn: The ARN of the connection.
        :param pulumi.Input[str] aws_device: The Direct Connect endpoint on which the physical connection terminates.
        :param pulumi.Input[str] bandwidth: The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
        :param pulumi.Input[str] has_logical_redundancy: Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
        :param pulumi.Input[bool] jumbo_frame_capable: Boolean value representing if jumbo frames have been enabled for this connection.
        :param pulumi.Input[str] location: The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        :param pulumi.Input[str] name: The name of the connection.
        :param pulumi.Input[str] owner_account_id: The ID of the AWS account that owns the connection.
        :param pulumi.Input[str] provider_name: The name of the service provider associated with the connection.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if aws_device is not None:
            pulumi.set(__self__, "aws_device", aws_device)
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if has_logical_redundancy is not None:
            pulumi.set(__self__, "has_logical_redundancy", has_logical_redundancy)
        if jumbo_frame_capable is not None:
            pulumi.set(__self__, "jumbo_frame_capable", jumbo_frame_capable)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner_account_id is not None:
            pulumi.set(__self__, "owner_account_id", owner_account_id)
        if provider_name is not None:
            pulumi.set(__self__, "provider_name", provider_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the connection.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="awsDevice")
    def aws_device(self) -> Optional[pulumi.Input[str]]:
        """
        The Direct Connect endpoint on which the physical connection terminates.
        """
        return pulumi.get(self, "aws_device")

    @aws_device.setter
    def aws_device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_device", value)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[str]]:
        """
        The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="hasLogicalRedundancy")
    def has_logical_redundancy(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
        """
        return pulumi.get(self, "has_logical_redundancy")

    @has_logical_redundancy.setter
    def has_logical_redundancy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "has_logical_redundancy", value)

    @property
    @pulumi.getter(name="jumboFrameCapable")
    def jumbo_frame_capable(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean value representing if jumbo frames have been enabled for this connection.
        """
        return pulumi.get(self, "jumbo_frame_capable")

    @jumbo_frame_capable.setter
    def jumbo_frame_capable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "jumbo_frame_capable", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the connection.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ownerAccountId")
    def owner_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the AWS account that owns the connection.
        """
        return pulumi.get(self, "owner_account_id")

    @owner_account_id.setter
    def owner_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_account_id", value)

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the service provider associated with the connection.
        """
        return pulumi.get(self, "provider_name")

    @provider_name.setter
    def provider_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class Connection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a Connection of Direct Connect.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        hoge = aws.directconnect.Connection("hoge",
            bandwidth="1Gbps",
            location="EqDC2")
        ```

        ## Import

        Direct Connect connections can be imported using the `connection id`, e.g.

        ```sh
         $ pulumi import aws:directconnect/connection:Connection test_connection dxcon-ffre0ec3
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bandwidth: The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
        :param pulumi.Input[str] location: The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        :param pulumi.Input[str] name: The name of the connection.
        :param pulumi.Input[str] provider_name: The name of the service provider associated with the connection.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Connection of Direct Connect.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        hoge = aws.directconnect.Connection("hoge",
            bandwidth="1Gbps",
            location="EqDC2")
        ```

        ## Import

        Direct Connect connections can be imported using the `connection id`, e.g.

        ```sh
         $ pulumi import aws:directconnect/connection:Connection test_connection dxcon-ffre0ec3
        ```

        :param str resource_name: The name of the resource.
        :param ConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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
            __props__ = ConnectionArgs.__new__(ConnectionArgs)

            if bandwidth is None and not opts.urn:
                raise TypeError("Missing required property 'bandwidth'")
            __props__.__dict__["bandwidth"] = bandwidth
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["provider_name"] = provider_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tags_all"] = tags_all
            __props__.__dict__["arn"] = None
            __props__.__dict__["aws_device"] = None
            __props__.__dict__["has_logical_redundancy"] = None
            __props__.__dict__["jumbo_frame_capable"] = None
            __props__.__dict__["owner_account_id"] = None
        super(Connection, __self__).__init__(
            'aws:directconnect/connection:Connection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            aws_device: Optional[pulumi.Input[str]] = None,
            bandwidth: Optional[pulumi.Input[str]] = None,
            has_logical_redundancy: Optional[pulumi.Input[str]] = None,
            jumbo_frame_capable: Optional[pulumi.Input[bool]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            owner_account_id: Optional[pulumi.Input[str]] = None,
            provider_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Connection':
        """
        Get an existing Connection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the connection.
        :param pulumi.Input[str] aws_device: The Direct Connect endpoint on which the physical connection terminates.
        :param pulumi.Input[str] bandwidth: The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
        :param pulumi.Input[str] has_logical_redundancy: Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
        :param pulumi.Input[bool] jumbo_frame_capable: Boolean value representing if jumbo frames have been enabled for this connection.
        :param pulumi.Input[str] location: The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        :param pulumi.Input[str] name: The name of the connection.
        :param pulumi.Input[str] owner_account_id: The ID of the AWS account that owns the connection.
        :param pulumi.Input[str] provider_name: The name of the service provider associated with the connection.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConnectionState.__new__(_ConnectionState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["aws_device"] = aws_device
        __props__.__dict__["bandwidth"] = bandwidth
        __props__.__dict__["has_logical_redundancy"] = has_logical_redundancy
        __props__.__dict__["jumbo_frame_capable"] = jumbo_frame_capable
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["owner_account_id"] = owner_account_id
        __props__.__dict__["provider_name"] = provider_name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return Connection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the connection.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsDevice")
    def aws_device(self) -> pulumi.Output[str]:
        """
        The Direct Connect endpoint on which the physical connection terminates.
        """
        return pulumi.get(self, "aws_device")

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Output[str]:
        """
        The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="hasLogicalRedundancy")
    def has_logical_redundancy(self) -> pulumi.Output[str]:
        """
        Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
        """
        return pulumi.get(self, "has_logical_redundancy")

    @property
    @pulumi.getter(name="jumboFrameCapable")
    def jumbo_frame_capable(self) -> pulumi.Output[bool]:
        """
        Boolean value representing if jumbo frames have been enabled for this connection.
        """
        return pulumi.get(self, "jumbo_frame_capable")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the connection.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerAccountId")
    def owner_account_id(self) -> pulumi.Output[str]:
        """
        The ID of the AWS account that owns the connection.
        """
        return pulumi.get(self, "owner_account_id")

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> pulumi.Output[str]:
        """
        The name of the service provider associated with the connection.
        """
        return pulumi.get(self, "provider_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

