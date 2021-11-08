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

__all__ = ['SqlInjectionMatchSetArgs', 'SqlInjectionMatchSet']

@pulumi.input_type
class SqlInjectionMatchSetArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 sql_injection_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input['SqlInjectionMatchSetSqlInjectionMatchTupleArgs']]]] = None):
        """
        The set of arguments for constructing a SqlInjectionMatchSet resource.
        :param pulumi.Input[str] name: The name or description of the SizeConstraintSet.
        :param pulumi.Input[Sequence[pulumi.Input['SqlInjectionMatchSetSqlInjectionMatchTupleArgs']]] sql_injection_match_tuples: The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if sql_injection_match_tuples is not None:
            pulumi.set(__self__, "sql_injection_match_tuples", sql_injection_match_tuples)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or description of the SizeConstraintSet.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sqlInjectionMatchTuples")
    def sql_injection_match_tuples(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SqlInjectionMatchSetSqlInjectionMatchTupleArgs']]]]:
        """
        The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
        """
        return pulumi.get(self, "sql_injection_match_tuples")

    @sql_injection_match_tuples.setter
    def sql_injection_match_tuples(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SqlInjectionMatchSetSqlInjectionMatchTupleArgs']]]]):
        pulumi.set(self, "sql_injection_match_tuples", value)


@pulumi.input_type
class _SqlInjectionMatchSetState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 sql_injection_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input['SqlInjectionMatchSetSqlInjectionMatchTupleArgs']]]] = None):
        """
        Input properties used for looking up and filtering SqlInjectionMatchSet resources.
        :param pulumi.Input[str] name: The name or description of the SizeConstraintSet.
        :param pulumi.Input[Sequence[pulumi.Input['SqlInjectionMatchSetSqlInjectionMatchTupleArgs']]] sql_injection_match_tuples: The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if sql_injection_match_tuples is not None:
            pulumi.set(__self__, "sql_injection_match_tuples", sql_injection_match_tuples)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or description of the SizeConstraintSet.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sqlInjectionMatchTuples")
    def sql_injection_match_tuples(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SqlInjectionMatchSetSqlInjectionMatchTupleArgs']]]]:
        """
        The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
        """
        return pulumi.get(self, "sql_injection_match_tuples")

    @sql_injection_match_tuples.setter
    def sql_injection_match_tuples(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SqlInjectionMatchSetSqlInjectionMatchTupleArgs']]]]):
        pulumi.set(self, "sql_injection_match_tuples", value)


class SqlInjectionMatchSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sql_injection_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SqlInjectionMatchSetSqlInjectionMatchTupleArgs']]]]] = None,
                 __props__=None):
        """
        Provides a WAF Regional SQL Injection Match Set Resource for use with Application Load Balancer.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        sql_injection_match_set = aws.wafregional.SqlInjectionMatchSet("sqlInjectionMatchSet", sql_injection_match_tuples=[aws.wafregional.SqlInjectionMatchSetSqlInjectionMatchTupleArgs(
            field_to_match=aws.wafregional.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatchArgs(
                type="QUERY_STRING",
            ),
            text_transformation="URL_DECODE",
        )])
        ```

        ## Import

        WAF Regional Sql Injection Match Set can be imported using the id, e.g.,

        ```sh
         $ pulumi import aws:wafregional/sqlInjectionMatchSet:SqlInjectionMatchSet sql_injection_match_set a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name or description of the SizeConstraintSet.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SqlInjectionMatchSetSqlInjectionMatchTupleArgs']]]] sql_injection_match_tuples: The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SqlInjectionMatchSetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a WAF Regional SQL Injection Match Set Resource for use with Application Load Balancer.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        sql_injection_match_set = aws.wafregional.SqlInjectionMatchSet("sqlInjectionMatchSet", sql_injection_match_tuples=[aws.wafregional.SqlInjectionMatchSetSqlInjectionMatchTupleArgs(
            field_to_match=aws.wafregional.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatchArgs(
                type="QUERY_STRING",
            ),
            text_transformation="URL_DECODE",
        )])
        ```

        ## Import

        WAF Regional Sql Injection Match Set can be imported using the id, e.g.,

        ```sh
         $ pulumi import aws:wafregional/sqlInjectionMatchSet:SqlInjectionMatchSet sql_injection_match_set a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
        ```

        :param str resource_name: The name of the resource.
        :param SqlInjectionMatchSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SqlInjectionMatchSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sql_injection_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SqlInjectionMatchSetSqlInjectionMatchTupleArgs']]]]] = None,
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
            __props__ = SqlInjectionMatchSetArgs.__new__(SqlInjectionMatchSetArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["sql_injection_match_tuples"] = sql_injection_match_tuples
        super(SqlInjectionMatchSet, __self__).__init__(
            'aws:wafregional/sqlInjectionMatchSet:SqlInjectionMatchSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            sql_injection_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SqlInjectionMatchSetSqlInjectionMatchTupleArgs']]]]] = None) -> 'SqlInjectionMatchSet':
        """
        Get an existing SqlInjectionMatchSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name or description of the SizeConstraintSet.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SqlInjectionMatchSetSqlInjectionMatchTupleArgs']]]] sql_injection_match_tuples: The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SqlInjectionMatchSetState.__new__(_SqlInjectionMatchSetState)

        __props__.__dict__["name"] = name
        __props__.__dict__["sql_injection_match_tuples"] = sql_injection_match_tuples
        return SqlInjectionMatchSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name or description of the SizeConstraintSet.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sqlInjectionMatchTuples")
    def sql_injection_match_tuples(self) -> pulumi.Output[Optional[Sequence['outputs.SqlInjectionMatchSetSqlInjectionMatchTuple']]]:
        """
        The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
        """
        return pulumi.get(self, "sql_injection_match_tuples")

