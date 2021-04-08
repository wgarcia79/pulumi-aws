# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PrivateDnsNamespaceArgs', 'PrivateDnsNamespace']

@pulumi.input_type
class PrivateDnsNamespaceArgs:
    def __init__(__self__, *,
                 vpc: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a PrivateDnsNamespace resource.
        :param pulumi.Input[str] vpc: The ID of VPC that you want to associate the namespace with.
        :param pulumi.Input[str] description: The description that you specify for the namespace when you create it.
        :param pulumi.Input[str] name: The name of the namespace.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the namespace.
        """
        pulumi.set(__self__, "vpc", vpc)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def vpc(self) -> pulumi.Input[str]:
        """
        The ID of VPC that you want to associate the namespace with.
        """
        return pulumi.get(self, "vpc")

    @vpc.setter
    def vpc(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description that you specify for the namespace when you create it.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the namespace.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the namespace.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _PrivateDnsNamespaceState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 hosted_zone: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PrivateDnsNamespace resources.
        :param pulumi.Input[str] arn: The ARN that Amazon Route 53 assigns to the namespace when you create it.
        :param pulumi.Input[str] description: The description that you specify for the namespace when you create it.
        :param pulumi.Input[str] hosted_zone: The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
        :param pulumi.Input[str] name: The name of the namespace.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the namespace.
        :param pulumi.Input[str] vpc: The ID of VPC that you want to associate the namespace with.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if hosted_zone is not None:
            pulumi.set(__self__, "hosted_zone", hosted_zone)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc is not None:
            pulumi.set(__self__, "vpc", vpc)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN that Amazon Route 53 assigns to the namespace when you create it.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description that you specify for the namespace when you create it.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="hostedZone")
    def hosted_zone(self) -> Optional[pulumi.Input[str]]:
        """
        The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
        """
        return pulumi.get(self, "hosted_zone")

    @hosted_zone.setter
    def hosted_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hosted_zone", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the namespace.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the namespace.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def vpc(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of VPC that you want to associate the namespace with.
        """
        return pulumi.get(self, "vpc")

    @vpc.setter
    def vpc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc", value)


class PrivateDnsNamespace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Service Discovery Private DNS Namespace resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.0.0.0/16")
        example_private_dns_namespace = aws.servicediscovery.PrivateDnsNamespace("examplePrivateDnsNamespace",
            description="example",
            vpc=example_vpc.id)
        ```

        ## Import

        Service Discovery Private DNS Namespace can be imported using the namespace ID and VPC ID, e.g.

        ```sh
         $ pulumi import aws:servicediscovery/privateDnsNamespace:PrivateDnsNamespace example 0123456789:vpc-123345
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description that you specify for the namespace when you create it.
        :param pulumi.Input[str] name: The name of the namespace.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the namespace.
        :param pulumi.Input[str] vpc: The ID of VPC that you want to associate the namespace with.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateDnsNamespaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Service Discovery Private DNS Namespace resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.0.0.0/16")
        example_private_dns_namespace = aws.servicediscovery.PrivateDnsNamespace("examplePrivateDnsNamespace",
            description="example",
            vpc=example_vpc.id)
        ```

        ## Import

        Service Discovery Private DNS Namespace can be imported using the namespace ID and VPC ID, e.g.

        ```sh
         $ pulumi import aws:servicediscovery/privateDnsNamespace:PrivateDnsNamespace example 0123456789:vpc-123345
        ```

        :param str resource_name: The name of the resource.
        :param PrivateDnsNamespaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateDnsNamespaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            __props__ = PrivateDnsNamespaceArgs.__new__(PrivateDnsNamespaceArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            if vpc is None and not opts.urn:
                raise TypeError("Missing required property 'vpc'")
            __props__.__dict__["vpc"] = vpc
            __props__.__dict__["arn"] = None
            __props__.__dict__["hosted_zone"] = None
        super(PrivateDnsNamespace, __self__).__init__(
            'aws:servicediscovery/privateDnsNamespace:PrivateDnsNamespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            hosted_zone: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc: Optional[pulumi.Input[str]] = None) -> 'PrivateDnsNamespace':
        """
        Get an existing PrivateDnsNamespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN that Amazon Route 53 assigns to the namespace when you create it.
        :param pulumi.Input[str] description: The description that you specify for the namespace when you create it.
        :param pulumi.Input[str] hosted_zone: The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
        :param pulumi.Input[str] name: The name of the namespace.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the namespace.
        :param pulumi.Input[str] vpc: The ID of VPC that you want to associate the namespace with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivateDnsNamespaceState.__new__(_PrivateDnsNamespaceState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["hosted_zone"] = hosted_zone
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vpc"] = vpc
        return PrivateDnsNamespace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN that Amazon Route 53 assigns to the namespace when you create it.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description that you specify for the namespace when you create it.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="hostedZone")
    def hosted_zone(self) -> pulumi.Output[str]:
        """
        The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
        """
        return pulumi.get(self, "hosted_zone")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the namespace.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the namespace.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def vpc(self) -> pulumi.Output[str]:
        """
        The ID of VPC that you want to associate the namespace with.
        """
        return pulumi.get(self, "vpc")

