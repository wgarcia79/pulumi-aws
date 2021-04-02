# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['QueryDefinition']


class QueryDefinition(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 log_group_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_string: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a CloudWatch Logs query definition resource.

        ## Import

        CloudWatch query definitions can be imported using the query definition ARN. The ARN can be found on the "Edit Query" page for the query in the AWS Console.

        ```sh
         $ pulumi import aws:cloudwatch/queryDefinition:QueryDefinition example arn:aws:logs:us-west-2:123456789012:query-definition:269951d7-6f75-496d-9d7b-6b7a5486bdbd
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] log_group_names: Specific log groups to use with the query.
        :param pulumi.Input[str] name: The name of the query.
        :param pulumi.Input[str] query_string: The query to save. You can read more about CloudWatch Logs Query Syntax in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['log_group_names'] = log_group_names
            __props__['name'] = name
            if query_string is None and not opts.urn:
                raise TypeError("Missing required property 'query_string'")
            __props__['query_string'] = query_string
            __props__['query_definition_id'] = None
        super(QueryDefinition, __self__).__init__(
            'aws:cloudwatch/queryDefinition:QueryDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            log_group_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            query_definition_id: Optional[pulumi.Input[str]] = None,
            query_string: Optional[pulumi.Input[str]] = None) -> 'QueryDefinition':
        """
        Get an existing QueryDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] log_group_names: Specific log groups to use with the query.
        :param pulumi.Input[str] name: The name of the query.
        :param pulumi.Input[str] query_definition_id: The query definition ID.
        :param pulumi.Input[str] query_string: The query to save. You can read more about CloudWatch Logs Query Syntax in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["log_group_names"] = log_group_names
        __props__["name"] = name
        __props__["query_definition_id"] = query_definition_id
        __props__["query_string"] = query_string
        return QueryDefinition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="logGroupNames")
    def log_group_names(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Specific log groups to use with the query.
        """
        return pulumi.get(self, "log_group_names")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the query.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="queryDefinitionId")
    def query_definition_id(self) -> pulumi.Output[str]:
        """
        The query definition ID.
        """
        return pulumi.get(self, "query_definition_id")

    @property
    @pulumi.getter(name="queryString")
    def query_string(self) -> pulumi.Output[str]:
        """
        The query to save. You can read more about CloudWatch Logs Query Syntax in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
        """
        return pulumi.get(self, "query_string")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

