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

__all__ = ['XssMatchSetArgs', 'XssMatchSet']

@pulumi.input_type
class XssMatchSetArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 xss_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input['XssMatchSetXssMatchTupleArgs']]]] = None):
        """
        The set of arguments for constructing a XssMatchSet resource.
        :param pulumi.Input[str] name: The name of the set
        :param pulumi.Input[Sequence[pulumi.Input['XssMatchSetXssMatchTupleArgs']]] xss_match_tuples: The parts of web requests that you want to inspect for cross-site scripting attacks.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if xss_match_tuples is not None:
            pulumi.set(__self__, "xss_match_tuples", xss_match_tuples)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the set
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="xssMatchTuples")
    def xss_match_tuples(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['XssMatchSetXssMatchTupleArgs']]]]:
        """
        The parts of web requests that you want to inspect for cross-site scripting attacks.
        """
        return pulumi.get(self, "xss_match_tuples")

    @xss_match_tuples.setter
    def xss_match_tuples(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['XssMatchSetXssMatchTupleArgs']]]]):
        pulumi.set(self, "xss_match_tuples", value)


@pulumi.input_type
class _XssMatchSetState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 xss_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input['XssMatchSetXssMatchTupleArgs']]]] = None):
        """
        Input properties used for looking up and filtering XssMatchSet resources.
        :param pulumi.Input[str] name: The name of the set
        :param pulumi.Input[Sequence[pulumi.Input['XssMatchSetXssMatchTupleArgs']]] xss_match_tuples: The parts of web requests that you want to inspect for cross-site scripting attacks.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if xss_match_tuples is not None:
            pulumi.set(__self__, "xss_match_tuples", xss_match_tuples)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the set
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="xssMatchTuples")
    def xss_match_tuples(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['XssMatchSetXssMatchTupleArgs']]]]:
        """
        The parts of web requests that you want to inspect for cross-site scripting attacks.
        """
        return pulumi.get(self, "xss_match_tuples")

    @xss_match_tuples.setter
    def xss_match_tuples(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['XssMatchSetXssMatchTupleArgs']]]]):
        pulumi.set(self, "xss_match_tuples", value)


class XssMatchSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 xss_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['XssMatchSetXssMatchTupleArgs']]]]] = None,
                 __props__=None):
        """
        Provides a WAF Regional XSS Match Set Resource for use with Application Load Balancer.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        xss_match_set = aws.wafregional.XssMatchSet("xssMatchSet", xss_match_tuples=[
            aws.wafregional.XssMatchSetXssMatchTupleArgs(
                field_to_match=aws.wafregional.XssMatchSetXssMatchTupleFieldToMatchArgs(
                    type="URI",
                ),
                text_transformation="NONE",
            ),
            aws.wafregional.XssMatchSetXssMatchTupleArgs(
                field_to_match=aws.wafregional.XssMatchSetXssMatchTupleFieldToMatchArgs(
                    type="QUERY_STRING",
                ),
                text_transformation="NONE",
            ),
        ])
        ```

        ## Import

        AWS WAF Regional XSS Match can be imported using the `id`, e.g.,

        ```sh
         $ pulumi import aws:wafregional/xssMatchSet:XssMatchSet example 12345abcde
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the set
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['XssMatchSetXssMatchTupleArgs']]]] xss_match_tuples: The parts of web requests that you want to inspect for cross-site scripting attacks.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[XssMatchSetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a WAF Regional XSS Match Set Resource for use with Application Load Balancer.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        xss_match_set = aws.wafregional.XssMatchSet("xssMatchSet", xss_match_tuples=[
            aws.wafregional.XssMatchSetXssMatchTupleArgs(
                field_to_match=aws.wafregional.XssMatchSetXssMatchTupleFieldToMatchArgs(
                    type="URI",
                ),
                text_transformation="NONE",
            ),
            aws.wafregional.XssMatchSetXssMatchTupleArgs(
                field_to_match=aws.wafregional.XssMatchSetXssMatchTupleFieldToMatchArgs(
                    type="QUERY_STRING",
                ),
                text_transformation="NONE",
            ),
        ])
        ```

        ## Import

        AWS WAF Regional XSS Match can be imported using the `id`, e.g.,

        ```sh
         $ pulumi import aws:wafregional/xssMatchSet:XssMatchSet example 12345abcde
        ```

        :param str resource_name: The name of the resource.
        :param XssMatchSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(XssMatchSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 xss_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['XssMatchSetXssMatchTupleArgs']]]]] = None,
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
            __props__ = XssMatchSetArgs.__new__(XssMatchSetArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["xss_match_tuples"] = xss_match_tuples
        super(XssMatchSet, __self__).__init__(
            'aws:wafregional/xssMatchSet:XssMatchSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            xss_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['XssMatchSetXssMatchTupleArgs']]]]] = None) -> 'XssMatchSet':
        """
        Get an existing XssMatchSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the set
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['XssMatchSetXssMatchTupleArgs']]]] xss_match_tuples: The parts of web requests that you want to inspect for cross-site scripting attacks.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _XssMatchSetState.__new__(_XssMatchSetState)

        __props__.__dict__["name"] = name
        __props__.__dict__["xss_match_tuples"] = xss_match_tuples
        return XssMatchSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the set
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="xssMatchTuples")
    def xss_match_tuples(self) -> pulumi.Output[Optional[Sequence['outputs.XssMatchSetXssMatchTuple']]]:
        """
        The parts of web requests that you want to inspect for cross-site scripting attacks.
        """
        return pulumi.get(self, "xss_match_tuples")

