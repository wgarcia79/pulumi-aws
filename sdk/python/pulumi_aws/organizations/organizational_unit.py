# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['OrganizationalUnitArgs', 'OrganizationalUnit']

@pulumi.input_type
class OrganizationalUnitArgs:
    def __init__(__self__, *,
                 parent_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a OrganizationalUnit resource.
        :param pulumi.Input[str] parent_id: ID of the parent organizational unit, which may be the root
        :param pulumi.Input[str] name: The name for the organizational unit
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "parent_id", parent_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> pulumi.Input[str]:
        """
        ID of the parent organizational unit, which may be the root
        """
        return pulumi.get(self, "parent_id")

    @parent_id.setter
    def parent_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "parent_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for the organizational unit
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _OrganizationalUnitState:
    def __init__(__self__, *,
                 accounts: Optional[pulumi.Input[Sequence[pulumi.Input['OrganizationalUnitAccountArgs']]]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering OrganizationalUnit resources.
        :param pulumi.Input[Sequence[pulumi.Input['OrganizationalUnitAccountArgs']]] accounts: List of child accounts for this Organizational Unit. Does not return account information for child Organizational Units. All elements have these attributes:
        :param pulumi.Input[str] arn: ARN of the organizational unit
        :param pulumi.Input[str] name: The name for the organizational unit
        :param pulumi.Input[str] parent_id: ID of the parent organizational unit, which may be the root
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        if accounts is not None:
            pulumi.set(__self__, "accounts", accounts)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_id is not None:
            pulumi.set(__self__, "parent_id", parent_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OrganizationalUnitAccountArgs']]]]:
        """
        List of child accounts for this Organizational Unit. Does not return account information for child Organizational Units. All elements have these attributes:
        """
        return pulumi.get(self, "accounts")

    @accounts.setter
    def accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OrganizationalUnitAccountArgs']]]]):
        pulumi.set(self, "accounts", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the organizational unit
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for the organizational unit
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the parent organizational unit, which may be the root
        """
        return pulumi.get(self, "parent_id")

    @parent_id.setter
    def parent_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class OrganizationalUnit(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to create an organizational unit.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.organizations.OrganizationalUnit("example", parent_id=aws_organizations_organization["example"]["roots"][0]["id"])
        ```

        ## Import

        AWS Organizations Organizational Units can be imported by using the `id`, e.g.,

        ```sh
         $ pulumi import aws:organizations/organizationalUnit:OrganizationalUnit example ou-1234567
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name for the organizational unit
        :param pulumi.Input[str] parent_id: ID of the parent organizational unit, which may be the root
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationalUnitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create an organizational unit.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.organizations.OrganizationalUnit("example", parent_id=aws_organizations_organization["example"]["roots"][0]["id"])
        ```

        ## Import

        AWS Organizations Organizational Units can be imported by using the `id`, e.g.,

        ```sh
         $ pulumi import aws:organizations/organizationalUnit:OrganizationalUnit example ou-1234567
        ```

        :param str resource_name: The name of the resource.
        :param OrganizationalUnitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationalUnitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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
            __props__ = OrganizationalUnitArgs.__new__(OrganizationalUnitArgs)

            __props__.__dict__["name"] = name
            if parent_id is None and not opts.urn:
                raise TypeError("Missing required property 'parent_id'")
            __props__.__dict__["parent_id"] = parent_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["accounts"] = None
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        super(OrganizationalUnit, __self__).__init__(
            'aws:organizations/organizationalUnit:OrganizationalUnit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrganizationalUnitAccountArgs']]]]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'OrganizationalUnit':
        """
        Get an existing OrganizationalUnit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrganizationalUnitAccountArgs']]]] accounts: List of child accounts for this Organizational Unit. Does not return account information for child Organizational Units. All elements have these attributes:
        :param pulumi.Input[str] arn: ARN of the organizational unit
        :param pulumi.Input[str] name: The name for the organizational unit
        :param pulumi.Input[str] parent_id: ID of the parent organizational unit, which may be the root
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrganizationalUnitState.__new__(_OrganizationalUnitState)

        __props__.__dict__["accounts"] = accounts
        __props__.__dict__["arn"] = arn
        __props__.__dict__["name"] = name
        __props__.__dict__["parent_id"] = parent_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return OrganizationalUnit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accounts(self) -> pulumi.Output[Sequence['outputs.OrganizationalUnitAccount']]:
        """
        List of child accounts for this Organizational Unit. Does not return account information for child Organizational Units. All elements have these attributes:
        """
        return pulumi.get(self, "accounts")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of the organizational unit
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name for the organizational unit
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> pulumi.Output[str]:
        """
        ID of the parent organizational unit, which may be the root
        """
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

