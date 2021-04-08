# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AssessmentTargetArgs', 'AssessmentTarget']

@pulumi.input_type
class AssessmentTargetArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_arn: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AssessmentTarget resource.
        :param pulumi.Input[str] name: The name of the assessment target.
        :param pulumi.Input[str] resource_group_arn: Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_arn is not None:
            pulumi.set(__self__, "resource_group_arn", resource_group_arn)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the assessment target.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupArn")
    def resource_group_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
        """
        return pulumi.get(self, "resource_group_arn")

    @resource_group_arn.setter
    def resource_group_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_arn", value)


@pulumi.input_type
class _AssessmentTargetState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_arn: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AssessmentTarget resources.
        :param pulumi.Input[str] arn: The target assessment ARN.
        :param pulumi.Input[str] name: The name of the assessment target.
        :param pulumi.Input[str] resource_group_arn: Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_arn is not None:
            pulumi.set(__self__, "resource_group_arn", resource_group_arn)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The target assessment ARN.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the assessment target.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupArn")
    def resource_group_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
        """
        return pulumi.get(self, "resource_group_arn")

    @resource_group_arn.setter
    def resource_group_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_arn", value)


class AssessmentTarget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Inspector assessment target

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        bar = aws.inspector.ResourceGroup("bar", tags={
            "Name": "foo",
            "Env": "bar",
        })
        foo = aws.inspector.AssessmentTarget("foo", resource_group_arn=bar.arn)
        ```

        ## Import

        Inspector Assessment Targets can be imported via their Amazon Resource Name (ARN), e.g.

        ```sh
         $ pulumi import aws:inspector/assessmentTarget:AssessmentTarget example arn:aws:inspector:us-east-1:123456789012:target/0-xxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the assessment target.
        :param pulumi.Input[str] resource_group_arn: Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AssessmentTargetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Inspector assessment target

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        bar = aws.inspector.ResourceGroup("bar", tags={
            "Name": "foo",
            "Env": "bar",
        })
        foo = aws.inspector.AssessmentTarget("foo", resource_group_arn=bar.arn)
        ```

        ## Import

        Inspector Assessment Targets can be imported via their Amazon Resource Name (ARN), e.g.

        ```sh
         $ pulumi import aws:inspector/assessmentTarget:AssessmentTarget example arn:aws:inspector:us-east-1:123456789012:target/0-xxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param AssessmentTargetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AssessmentTargetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_arn: Optional[pulumi.Input[str]] = None,
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
            __props__ = AssessmentTargetArgs.__new__(AssessmentTargetArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["resource_group_arn"] = resource_group_arn
            __props__.__dict__["arn"] = None
        super(AssessmentTarget, __self__).__init__(
            'aws:inspector/assessmentTarget:AssessmentTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_arn: Optional[pulumi.Input[str]] = None) -> 'AssessmentTarget':
        """
        Get an existing AssessmentTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The target assessment ARN.
        :param pulumi.Input[str] name: The name of the assessment target.
        :param pulumi.Input[str] resource_group_arn: Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AssessmentTargetState.__new__(_AssessmentTargetState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_arn"] = resource_group_arn
        return AssessmentTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The target assessment ARN.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the assessment target.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupArn")
    def resource_group_arn(self) -> pulumi.Output[Optional[str]]:
        """
        Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
        """
        return pulumi.get(self, "resource_group_arn")

