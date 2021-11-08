# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TagOptionResourceAssociationArgs', 'TagOptionResourceAssociation']

@pulumi.input_type
class TagOptionResourceAssociationArgs:
    def __init__(__self__, *,
                 resource_id: pulumi.Input[str],
                 tag_option_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a TagOptionResourceAssociation resource.
        :param pulumi.Input[str] resource_id: Resource identifier.
        :param pulumi.Input[str] tag_option_id: Tag Option identifier.
        """
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "tag_option_id", tag_option_id)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        Resource identifier.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="tagOptionId")
    def tag_option_id(self) -> pulumi.Input[str]:
        """
        Tag Option identifier.
        """
        return pulumi.get(self, "tag_option_id")

    @tag_option_id.setter
    def tag_option_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tag_option_id", value)


@pulumi.input_type
class _TagOptionResourceAssociationState:
    def __init__(__self__, *,
                 resource_arn: Optional[pulumi.Input[str]] = None,
                 resource_created_time: Optional[pulumi.Input[str]] = None,
                 resource_description: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_name: Optional[pulumi.Input[str]] = None,
                 tag_option_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TagOptionResourceAssociation resources.
        :param pulumi.Input[str] resource_arn: ARN of the resource.
        :param pulumi.Input[str] resource_created_time: Creation time of the resource.
        :param pulumi.Input[str] resource_description: Description of the resource.
        :param pulumi.Input[str] resource_id: Resource identifier.
        :param pulumi.Input[str] resource_name: Description of the resource.
        :param pulumi.Input[str] tag_option_id: Tag Option identifier.
        """
        if resource_arn is not None:
            pulumi.set(__self__, "resource_arn", resource_arn)
        if resource_created_time is not None:
            pulumi.set(__self__, "resource_created_time", resource_created_time)
        if resource_description is not None:
            pulumi.set(__self__, "resource_description", resource_description)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_name is not None:
            pulumi.set(__self__, "resource_name", resource_name)
        if tag_option_id is not None:
            pulumi.set(__self__, "tag_option_id", tag_option_id)

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the resource.
        """
        return pulumi.get(self, "resource_arn")

    @resource_arn.setter
    def resource_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_arn", value)

    @property
    @pulumi.getter(name="resourceCreatedTime")
    def resource_created_time(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time of the resource.
        """
        return pulumi.get(self, "resource_created_time")

    @resource_created_time.setter
    def resource_created_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_created_time", value)

    @property
    @pulumi.getter(name="resourceDescription")
    def resource_description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the resource.
        """
        return pulumi.get(self, "resource_description")

    @resource_description.setter
    def resource_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_description", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource identifier.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the resource.
        """
        return pulumi.get(self, "resource_name")

    @resource_name.setter
    def resource_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_name", value)

    @property
    @pulumi.getter(name="tagOptionId")
    def tag_option_id(self) -> Optional[pulumi.Input[str]]:
        """
        Tag Option identifier.
        """
        return pulumi.get(self, "tag_option_id")

    @tag_option_id.setter
    def tag_option_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag_option_id", value)


class TagOptionResourceAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 tag_option_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Service Catalog Tag Option Resource Association.

        > **Tip:** A "resource" is either a Service Catalog portfolio or product.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.servicecatalog.TagOptionResourceAssociation("example",
            resource_id="prod-dnigbtea24ste",
            tag_option_id="tag-pjtvyakdlyo3m")
        ```

        ## Import

        `aws_servicecatalog_tag_option_resource_association` can be imported using the tag option ID and resource ID, e.g.,

        ```sh
         $ pulumi import aws:servicecatalog/tagOptionResourceAssociation:TagOptionResourceAssociation example tag-pjtvyakdlyo3m:prod-dnigbtea24ste
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] resource_id: Resource identifier.
        :param pulumi.Input[str] tag_option_id: Tag Option identifier.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TagOptionResourceAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Service Catalog Tag Option Resource Association.

        > **Tip:** A "resource" is either a Service Catalog portfolio or product.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.servicecatalog.TagOptionResourceAssociation("example",
            resource_id="prod-dnigbtea24ste",
            tag_option_id="tag-pjtvyakdlyo3m")
        ```

        ## Import

        `aws_servicecatalog_tag_option_resource_association` can be imported using the tag option ID and resource ID, e.g.,

        ```sh
         $ pulumi import aws:servicecatalog/tagOptionResourceAssociation:TagOptionResourceAssociation example tag-pjtvyakdlyo3m:prod-dnigbtea24ste
        ```

        :param str resource_name: The name of the resource.
        :param TagOptionResourceAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TagOptionResourceAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 tag_option_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = TagOptionResourceAssociationArgs.__new__(TagOptionResourceAssociationArgs)

            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
            if tag_option_id is None and not opts.urn:
                raise TypeError("Missing required property 'tag_option_id'")
            __props__.__dict__["tag_option_id"] = tag_option_id
            __props__.__dict__["resource_arn"] = None
            __props__.__dict__["resource_created_time"] = None
            __props__.__dict__["resource_description"] = None
            __props__.__dict__["resource_name"] = None
        super(TagOptionResourceAssociation, __self__).__init__(
            'aws:servicecatalog/tagOptionResourceAssociation:TagOptionResourceAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            resource_arn: Optional[pulumi.Input[str]] = None,
            resource_created_time: Optional[pulumi.Input[str]] = None,
            resource_description: Optional[pulumi.Input[str]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            resource_name_: Optional[pulumi.Input[str]] = None,
            tag_option_id: Optional[pulumi.Input[str]] = None) -> 'TagOptionResourceAssociation':
        """
        Get an existing TagOptionResourceAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] resource_arn: ARN of the resource.
        :param pulumi.Input[str] resource_created_time: Creation time of the resource.
        :param pulumi.Input[str] resource_description: Description of the resource.
        :param pulumi.Input[str] resource_id: Resource identifier.
        :param pulumi.Input[str] resource_name_: Description of the resource.
        :param pulumi.Input[str] tag_option_id: Tag Option identifier.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TagOptionResourceAssociationState.__new__(_TagOptionResourceAssociationState)

        __props__.__dict__["resource_arn"] = resource_arn
        __props__.__dict__["resource_created_time"] = resource_created_time
        __props__.__dict__["resource_description"] = resource_description
        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["resource_name"] = resource_name_
        __props__.__dict__["tag_option_id"] = tag_option_id
        return TagOptionResourceAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> pulumi.Output[str]:
        """
        ARN of the resource.
        """
        return pulumi.get(self, "resource_arn")

    @property
    @pulumi.getter(name="resourceCreatedTime")
    def resource_created_time(self) -> pulumi.Output[str]:
        """
        Creation time of the resource.
        """
        return pulumi.get(self, "resource_created_time")

    @property
    @pulumi.getter(name="resourceDescription")
    def resource_description(self) -> pulumi.Output[str]:
        """
        Description of the resource.
        """
        return pulumi.get(self, "resource_description")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        Resource identifier.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> pulumi.Output[str]:
        """
        Description of the resource.
        """
        return pulumi.get(self, "resource_name")

    @property
    @pulumi.getter(name="tagOptionId")
    def tag_option_id(self) -> pulumi.Output[str]:
        """
        Tag Option identifier.
        """
        return pulumi.get(self, "tag_option_id")

