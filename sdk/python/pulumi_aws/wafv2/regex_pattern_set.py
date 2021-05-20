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

__all__ = ['RegexPatternSetArgs', 'RegexPatternSet']

@pulumi.input_type
class RegexPatternSetArgs:
    def __init__(__self__, *,
                 scope: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 regular_expressions: Optional[pulumi.Input[Sequence[pulumi.Input['RegexPatternSetRegularExpressionArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a RegexPatternSet resource.
        :param pulumi.Input[str] scope: Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        :param pulumi.Input[str] description: A friendly description of the regular expression pattern set.
        :param pulumi.Input[str] name: A friendly name of the regular expression pattern set.
        :param pulumi.Input[Sequence[pulumi.Input['RegexPatternSetRegularExpressionArgs']]] regular_expressions: One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: An array of key:value pairs to associate with the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        pulumi.set(__self__, "scope", scope)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if regular_expressions is not None:
            pulumi.set(__self__, "regular_expressions", regular_expressions)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Input[str]:
        """
        Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: pulumi.Input[str]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A friendly description of the regular expression pattern set.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A friendly name of the regular expression pattern set.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="regularExpressions")
    def regular_expressions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RegexPatternSetRegularExpressionArgs']]]]:
        """
        One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
        """
        return pulumi.get(self, "regular_expressions")

    @regular_expressions.setter
    def regular_expressions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RegexPatternSetRegularExpressionArgs']]]]):
        pulumi.set(self, "regular_expressions", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        An array of key:value pairs to associate with the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


@pulumi.input_type
class _RegexPatternSetState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 lock_token: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 regular_expressions: Optional[pulumi.Input[Sequence[pulumi.Input['RegexPatternSetRegularExpressionArgs']]]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering RegexPatternSet resources.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) that identifies the cluster.
        :param pulumi.Input[str] description: A friendly description of the regular expression pattern set.
        :param pulumi.Input[str] name: A friendly name of the regular expression pattern set.
        :param pulumi.Input[Sequence[pulumi.Input['RegexPatternSetRegularExpressionArgs']]] regular_expressions: One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
        :param pulumi.Input[str] scope: Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: An array of key:value pairs to associate with the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if lock_token is not None:
            pulumi.set(__self__, "lock_token", lock_token)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if regular_expressions is not None:
            pulumi.set(__self__, "regular_expressions", regular_expressions)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) that identifies the cluster.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A friendly description of the regular expression pattern set.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="lockToken")
    def lock_token(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "lock_token")

    @lock_token.setter
    def lock_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lock_token", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A friendly name of the regular expression pattern set.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="regularExpressions")
    def regular_expressions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RegexPatternSetRegularExpressionArgs']]]]:
        """
        One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
        """
        return pulumi.get(self, "regular_expressions")

    @regular_expressions.setter
    def regular_expressions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RegexPatternSetRegularExpressionArgs']]]]):
        pulumi.set(self, "regular_expressions", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        An array of key:value pairs to associate with the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class RegexPatternSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 regular_expressions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegexPatternSetRegularExpressionArgs']]]]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides an AWS WAFv2 Regex Pattern Set Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.wafv2.RegexPatternSet("example",
            description="Example regex pattern set",
            regular_expressions=[
                aws.wafv2.RegexPatternSetRegularExpressionArgs(
                    regex_string="one",
                ),
                aws.wafv2.RegexPatternSetRegularExpressionArgs(
                    regex_string="two",
                ),
            ],
            scope="REGIONAL",
            tags={
                "Tag1": "Value1",
                "Tag2": "Value2",
            })
        ```

        ## Import

        WAFv2 Regex Pattern Sets can be imported using `ID/name/scope` e.g.

        ```sh
         $ pulumi import aws:wafv2/regexPatternSet:RegexPatternSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc/example/REGIONAL
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A friendly description of the regular expression pattern set.
        :param pulumi.Input[str] name: A friendly name of the regular expression pattern set.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegexPatternSetRegularExpressionArgs']]]] regular_expressions: One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
        :param pulumi.Input[str] scope: Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: An array of key:value pairs to associate with the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegexPatternSetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AWS WAFv2 Regex Pattern Set Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.wafv2.RegexPatternSet("example",
            description="Example regex pattern set",
            regular_expressions=[
                aws.wafv2.RegexPatternSetRegularExpressionArgs(
                    regex_string="one",
                ),
                aws.wafv2.RegexPatternSetRegularExpressionArgs(
                    regex_string="two",
                ),
            ],
            scope="REGIONAL",
            tags={
                "Tag1": "Value1",
                "Tag2": "Value2",
            })
        ```

        ## Import

        WAFv2 Regex Pattern Sets can be imported using `ID/name/scope` e.g.

        ```sh
         $ pulumi import aws:wafv2/regexPatternSet:RegexPatternSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc/example/REGIONAL
        ```

        :param str resource_name: The name of the resource.
        :param RegexPatternSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegexPatternSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 regular_expressions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegexPatternSetRegularExpressionArgs']]]]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
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
            __props__ = RegexPatternSetArgs.__new__(RegexPatternSetArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["regular_expressions"] = regular_expressions
            if scope is None and not opts.urn:
                raise TypeError("Missing required property 'scope'")
            __props__.__dict__["scope"] = scope
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tags_all"] = tags_all
            __props__.__dict__["arn"] = None
            __props__.__dict__["lock_token"] = None
        super(RegexPatternSet, __self__).__init__(
            'aws:wafv2/regexPatternSet:RegexPatternSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            lock_token: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            regular_expressions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegexPatternSetRegularExpressionArgs']]]]] = None,
            scope: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'RegexPatternSet':
        """
        Get an existing RegexPatternSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) that identifies the cluster.
        :param pulumi.Input[str] description: A friendly description of the regular expression pattern set.
        :param pulumi.Input[str] name: A friendly name of the regular expression pattern set.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegexPatternSetRegularExpressionArgs']]]] regular_expressions: One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
        :param pulumi.Input[str] scope: Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: An array of key:value pairs to associate with the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RegexPatternSetState.__new__(_RegexPatternSetState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["lock_token"] = lock_token
        __props__.__dict__["name"] = name
        __props__.__dict__["regular_expressions"] = regular_expressions
        __props__.__dict__["scope"] = scope
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return RegexPatternSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) that identifies the cluster.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A friendly description of the regular expression pattern set.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lockToken")
    def lock_token(self) -> pulumi.Output[str]:
        return pulumi.get(self, "lock_token")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A friendly name of the regular expression pattern set.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="regularExpressions")
    def regular_expressions(self) -> pulumi.Output[Optional[Sequence['outputs.RegexPatternSetRegularExpression']]]:
        """
        One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
        """
        return pulumi.get(self, "regular_expressions")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        An array of key:value pairs to associate with the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

