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

__all__ = ['ResolverArgs', 'Resolver']

@pulumi.input_type
class ResolverArgs:
    def __init__(__self__, *,
                 api_id: pulumi.Input[str],
                 field: pulumi.Input[str],
                 type: pulumi.Input[str],
                 caching_config: Optional[pulumi.Input['ResolverCachingConfigArgs']] = None,
                 data_source: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 pipeline_config: Optional[pulumi.Input['ResolverPipelineConfigArgs']] = None,
                 request_template: Optional[pulumi.Input[str]] = None,
                 response_template: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Resolver resource.
        :param pulumi.Input[str] api_id: The API ID for the GraphQL API.
        :param pulumi.Input[str] field: The field name from the schema defined in the GraphQL API.
        :param pulumi.Input[str] type: The type name from the schema defined in the GraphQL API.
        :param pulumi.Input['ResolverCachingConfigArgs'] caching_config: The CachingConfig.
        :param pulumi.Input[str] data_source: The DataSource name.
        :param pulumi.Input[str] kind: The resolver type. Valid values are `UNIT` and `PIPELINE`.
        :param pulumi.Input['ResolverPipelineConfigArgs'] pipeline_config: The PipelineConfig.
        :param pulumi.Input[str] request_template: The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
        :param pulumi.Input[str] response_template: The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
        """
        pulumi.set(__self__, "api_id", api_id)
        pulumi.set(__self__, "field", field)
        pulumi.set(__self__, "type", type)
        if caching_config is not None:
            pulumi.set(__self__, "caching_config", caching_config)
        if data_source is not None:
            pulumi.set(__self__, "data_source", data_source)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if pipeline_config is not None:
            pulumi.set(__self__, "pipeline_config", pipeline_config)
        if request_template is not None:
            pulumi.set(__self__, "request_template", request_template)
        if response_template is not None:
            pulumi.set(__self__, "response_template", response_template)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[str]:
        """
        The API ID for the GraphQL API.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter
    def field(self) -> pulumi.Input[str]:
        """
        The field name from the schema defined in the GraphQL API.
        """
        return pulumi.get(self, "field")

    @field.setter
    def field(self, value: pulumi.Input[str]):
        pulumi.set(self, "field", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type name from the schema defined in the GraphQL API.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="cachingConfig")
    def caching_config(self) -> Optional[pulumi.Input['ResolverCachingConfigArgs']]:
        """
        The CachingConfig.
        """
        return pulumi.get(self, "caching_config")

    @caching_config.setter
    def caching_config(self, value: Optional[pulumi.Input['ResolverCachingConfigArgs']]):
        pulumi.set(self, "caching_config", value)

    @property
    @pulumi.getter(name="dataSource")
    def data_source(self) -> Optional[pulumi.Input[str]]:
        """
        The DataSource name.
        """
        return pulumi.get(self, "data_source")

    @data_source.setter
    def data_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_source", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        The resolver type. Valid values are `UNIT` and `PIPELINE`.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="pipelineConfig")
    def pipeline_config(self) -> Optional[pulumi.Input['ResolverPipelineConfigArgs']]:
        """
        The PipelineConfig.
        """
        return pulumi.get(self, "pipeline_config")

    @pipeline_config.setter
    def pipeline_config(self, value: Optional[pulumi.Input['ResolverPipelineConfigArgs']]):
        pulumi.set(self, "pipeline_config", value)

    @property
    @pulumi.getter(name="requestTemplate")
    def request_template(self) -> Optional[pulumi.Input[str]]:
        """
        The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
        """
        return pulumi.get(self, "request_template")

    @request_template.setter
    def request_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_template", value)

    @property
    @pulumi.getter(name="responseTemplate")
    def response_template(self) -> Optional[pulumi.Input[str]]:
        """
        The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
        """
        return pulumi.get(self, "response_template")

    @response_template.setter
    def response_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "response_template", value)


@pulumi.input_type
class _ResolverState:
    def __init__(__self__, *,
                 api_id: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 caching_config: Optional[pulumi.Input['ResolverCachingConfigArgs']] = None,
                 data_source: Optional[pulumi.Input[str]] = None,
                 field: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 pipeline_config: Optional[pulumi.Input['ResolverPipelineConfigArgs']] = None,
                 request_template: Optional[pulumi.Input[str]] = None,
                 response_template: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Resolver resources.
        :param pulumi.Input[str] api_id: The API ID for the GraphQL API.
        :param pulumi.Input[str] arn: The ARN
        :param pulumi.Input['ResolverCachingConfigArgs'] caching_config: The CachingConfig.
        :param pulumi.Input[str] data_source: The DataSource name.
        :param pulumi.Input[str] field: The field name from the schema defined in the GraphQL API.
        :param pulumi.Input[str] kind: The resolver type. Valid values are `UNIT` and `PIPELINE`.
        :param pulumi.Input['ResolverPipelineConfigArgs'] pipeline_config: The PipelineConfig.
        :param pulumi.Input[str] request_template: The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
        :param pulumi.Input[str] response_template: The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
        :param pulumi.Input[str] type: The type name from the schema defined in the GraphQL API.
        """
        if api_id is not None:
            pulumi.set(__self__, "api_id", api_id)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if caching_config is not None:
            pulumi.set(__self__, "caching_config", caching_config)
        if data_source is not None:
            pulumi.set(__self__, "data_source", data_source)
        if field is not None:
            pulumi.set(__self__, "field", field)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if pipeline_config is not None:
            pulumi.set(__self__, "pipeline_config", pipeline_config)
        if request_template is not None:
            pulumi.set(__self__, "request_template", request_template)
        if response_template is not None:
            pulumi.set(__self__, "response_template", response_template)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> Optional[pulumi.Input[str]]:
        """
        The API ID for the GraphQL API.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="cachingConfig")
    def caching_config(self) -> Optional[pulumi.Input['ResolverCachingConfigArgs']]:
        """
        The CachingConfig.
        """
        return pulumi.get(self, "caching_config")

    @caching_config.setter
    def caching_config(self, value: Optional[pulumi.Input['ResolverCachingConfigArgs']]):
        pulumi.set(self, "caching_config", value)

    @property
    @pulumi.getter(name="dataSource")
    def data_source(self) -> Optional[pulumi.Input[str]]:
        """
        The DataSource name.
        """
        return pulumi.get(self, "data_source")

    @data_source.setter
    def data_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_source", value)

    @property
    @pulumi.getter
    def field(self) -> Optional[pulumi.Input[str]]:
        """
        The field name from the schema defined in the GraphQL API.
        """
        return pulumi.get(self, "field")

    @field.setter
    def field(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        The resolver type. Valid values are `UNIT` and `PIPELINE`.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="pipelineConfig")
    def pipeline_config(self) -> Optional[pulumi.Input['ResolverPipelineConfigArgs']]:
        """
        The PipelineConfig.
        """
        return pulumi.get(self, "pipeline_config")

    @pipeline_config.setter
    def pipeline_config(self, value: Optional[pulumi.Input['ResolverPipelineConfigArgs']]):
        pulumi.set(self, "pipeline_config", value)

    @property
    @pulumi.getter(name="requestTemplate")
    def request_template(self) -> Optional[pulumi.Input[str]]:
        """
        The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
        """
        return pulumi.get(self, "request_template")

    @request_template.setter
    def request_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_template", value)

    @property
    @pulumi.getter(name="responseTemplate")
    def response_template(self) -> Optional[pulumi.Input[str]]:
        """
        The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
        """
        return pulumi.get(self, "response_template")

    @response_template.setter
    def response_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "response_template", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type name from the schema defined in the GraphQL API.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Resolver(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 caching_config: Optional[pulumi.Input[pulumi.InputType['ResolverCachingConfigArgs']]] = None,
                 data_source: Optional[pulumi.Input[str]] = None,
                 field: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 pipeline_config: Optional[pulumi.Input[pulumi.InputType['ResolverPipelineConfigArgs']]] = None,
                 request_template: Optional[pulumi.Input[str]] = None,
                 response_template: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an AppSync Resolver.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test_graph_ql_api = aws.appsync.GraphQLApi("testGraphQLApi",
            authentication_type="API_KEY",
            schema=\"\"\"type Mutation {
        	putPost(id: ID!, title: String!): Post
        }

        type Post {
        	id: ID!
        	title: String!
        }

        type Query {
        	singlePost(id: ID!): Post
        }

        schema {
        	query: Query
        	mutation: Mutation
        }
        \"\"\")
        test_data_source = aws.appsync.DataSource("testDataSource",
            api_id=test_graph_ql_api.id,
            name="tf_example",
            type="HTTP",
            http_config=aws.appsync.DataSourceHttpConfigArgs(
                endpoint="http://example.com",
            ))
        # UNIT type resolver (default)
        test_resolver = aws.appsync.Resolver("testResolver",
            api_id=test_graph_ql_api.id,
            field="singlePost",
            type="Query",
            data_source=test_data_source.name,
            request_template=\"\"\"{
            "version": "2018-05-29",
            "method": "GET",
            "resourcePath": "/",
            "params":{
                "headers": $utils.http.copyheaders($ctx.request.headers)
            }
        }
        \"\"\",
            response_template=\"\"\"#if($ctx.result.statusCode == 200)
            $ctx.result.body
        #else
            $utils.appendError($ctx.result.body, $ctx.result.statusCode)
        #end
        \"\"\",
            caching_config=aws.appsync.ResolverCachingConfigArgs(
                caching_keys=[
                    "$context.identity.sub",
                    "$context.arguments.id",
                ],
                ttl=60,
            ))
        # PIPELINE type resolver
        mutation_pipeline_test = aws.appsync.Resolver("mutationPipelineTest",
            type="Mutation",
            api_id=test_graph_ql_api.id,
            field="pipelineTest",
            request_template="{}",
            response_template="$util.toJson($ctx.result)",
            kind="PIPELINE",
            pipeline_config=aws.appsync.ResolverPipelineConfigArgs(
                functions=[
                    aws_appsync_function["test1"]["function_id"],
                    aws_appsync_function["test2"]["function_id"],
                    aws_appsync_function["test3"]["function_id"],
                ],
            ))
        ```

        ## Import

        `aws_appsync_resolver` can be imported with their `api_id`, a hyphen, `type`, a hypen and `field` e.g.,

        ```sh
         $ pulumi import aws:appsync/resolver:Resolver example abcdef123456-exampleType-exampleField
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: The API ID for the GraphQL API.
        :param pulumi.Input[pulumi.InputType['ResolverCachingConfigArgs']] caching_config: The CachingConfig.
        :param pulumi.Input[str] data_source: The DataSource name.
        :param pulumi.Input[str] field: The field name from the schema defined in the GraphQL API.
        :param pulumi.Input[str] kind: The resolver type. Valid values are `UNIT` and `PIPELINE`.
        :param pulumi.Input[pulumi.InputType['ResolverPipelineConfigArgs']] pipeline_config: The PipelineConfig.
        :param pulumi.Input[str] request_template: The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
        :param pulumi.Input[str] response_template: The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
        :param pulumi.Input[str] type: The type name from the schema defined in the GraphQL API.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResolverArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AppSync Resolver.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test_graph_ql_api = aws.appsync.GraphQLApi("testGraphQLApi",
            authentication_type="API_KEY",
            schema=\"\"\"type Mutation {
        	putPost(id: ID!, title: String!): Post
        }

        type Post {
        	id: ID!
        	title: String!
        }

        type Query {
        	singlePost(id: ID!): Post
        }

        schema {
        	query: Query
        	mutation: Mutation
        }
        \"\"\")
        test_data_source = aws.appsync.DataSource("testDataSource",
            api_id=test_graph_ql_api.id,
            name="tf_example",
            type="HTTP",
            http_config=aws.appsync.DataSourceHttpConfigArgs(
                endpoint="http://example.com",
            ))
        # UNIT type resolver (default)
        test_resolver = aws.appsync.Resolver("testResolver",
            api_id=test_graph_ql_api.id,
            field="singlePost",
            type="Query",
            data_source=test_data_source.name,
            request_template=\"\"\"{
            "version": "2018-05-29",
            "method": "GET",
            "resourcePath": "/",
            "params":{
                "headers": $utils.http.copyheaders($ctx.request.headers)
            }
        }
        \"\"\",
            response_template=\"\"\"#if($ctx.result.statusCode == 200)
            $ctx.result.body
        #else
            $utils.appendError($ctx.result.body, $ctx.result.statusCode)
        #end
        \"\"\",
            caching_config=aws.appsync.ResolverCachingConfigArgs(
                caching_keys=[
                    "$context.identity.sub",
                    "$context.arguments.id",
                ],
                ttl=60,
            ))
        # PIPELINE type resolver
        mutation_pipeline_test = aws.appsync.Resolver("mutationPipelineTest",
            type="Mutation",
            api_id=test_graph_ql_api.id,
            field="pipelineTest",
            request_template="{}",
            response_template="$util.toJson($ctx.result)",
            kind="PIPELINE",
            pipeline_config=aws.appsync.ResolverPipelineConfigArgs(
                functions=[
                    aws_appsync_function["test1"]["function_id"],
                    aws_appsync_function["test2"]["function_id"],
                    aws_appsync_function["test3"]["function_id"],
                ],
            ))
        ```

        ## Import

        `aws_appsync_resolver` can be imported with their `api_id`, a hyphen, `type`, a hypen and `field` e.g.,

        ```sh
         $ pulumi import aws:appsync/resolver:Resolver example abcdef123456-exampleType-exampleField
        ```

        :param str resource_name: The name of the resource.
        :param ResolverArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResolverArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 caching_config: Optional[pulumi.Input[pulumi.InputType['ResolverCachingConfigArgs']]] = None,
                 data_source: Optional[pulumi.Input[str]] = None,
                 field: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 pipeline_config: Optional[pulumi.Input[pulumi.InputType['ResolverPipelineConfigArgs']]] = None,
                 request_template: Optional[pulumi.Input[str]] = None,
                 response_template: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            __props__ = ResolverArgs.__new__(ResolverArgs)

            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            __props__.__dict__["caching_config"] = caching_config
            __props__.__dict__["data_source"] = data_source
            if field is None and not opts.urn:
                raise TypeError("Missing required property 'field'")
            __props__.__dict__["field"] = field
            __props__.__dict__["kind"] = kind
            __props__.__dict__["pipeline_config"] = pipeline_config
            __props__.__dict__["request_template"] = request_template
            __props__.__dict__["response_template"] = response_template
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["arn"] = None
        super(Resolver, __self__).__init__(
            'aws:appsync/resolver:Resolver',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_id: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            caching_config: Optional[pulumi.Input[pulumi.InputType['ResolverCachingConfigArgs']]] = None,
            data_source: Optional[pulumi.Input[str]] = None,
            field: Optional[pulumi.Input[str]] = None,
            kind: Optional[pulumi.Input[str]] = None,
            pipeline_config: Optional[pulumi.Input[pulumi.InputType['ResolverPipelineConfigArgs']]] = None,
            request_template: Optional[pulumi.Input[str]] = None,
            response_template: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Resolver':
        """
        Get an existing Resolver resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: The API ID for the GraphQL API.
        :param pulumi.Input[str] arn: The ARN
        :param pulumi.Input[pulumi.InputType['ResolverCachingConfigArgs']] caching_config: The CachingConfig.
        :param pulumi.Input[str] data_source: The DataSource name.
        :param pulumi.Input[str] field: The field name from the schema defined in the GraphQL API.
        :param pulumi.Input[str] kind: The resolver type. Valid values are `UNIT` and `PIPELINE`.
        :param pulumi.Input[pulumi.InputType['ResolverPipelineConfigArgs']] pipeline_config: The PipelineConfig.
        :param pulumi.Input[str] request_template: The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
        :param pulumi.Input[str] response_template: The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
        :param pulumi.Input[str] type: The type name from the schema defined in the GraphQL API.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ResolverState.__new__(_ResolverState)

        __props__.__dict__["api_id"] = api_id
        __props__.__dict__["arn"] = arn
        __props__.__dict__["caching_config"] = caching_config
        __props__.__dict__["data_source"] = data_source
        __props__.__dict__["field"] = field
        __props__.__dict__["kind"] = kind
        __props__.__dict__["pipeline_config"] = pipeline_config
        __props__.__dict__["request_template"] = request_template
        __props__.__dict__["response_template"] = response_template
        __props__.__dict__["type"] = type
        return Resolver(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        """
        The API ID for the GraphQL API.
        """
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="cachingConfig")
    def caching_config(self) -> pulumi.Output[Optional['outputs.ResolverCachingConfig']]:
        """
        The CachingConfig.
        """
        return pulumi.get(self, "caching_config")

    @property
    @pulumi.getter(name="dataSource")
    def data_source(self) -> pulumi.Output[Optional[str]]:
        """
        The DataSource name.
        """
        return pulumi.get(self, "data_source")

    @property
    @pulumi.getter
    def field(self) -> pulumi.Output[str]:
        """
        The field name from the schema defined in the GraphQL API.
        """
        return pulumi.get(self, "field")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[str]]:
        """
        The resolver type. Valid values are `UNIT` and `PIPELINE`.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="pipelineConfig")
    def pipeline_config(self) -> pulumi.Output[Optional['outputs.ResolverPipelineConfig']]:
        """
        The PipelineConfig.
        """
        return pulumi.get(self, "pipeline_config")

    @property
    @pulumi.getter(name="requestTemplate")
    def request_template(self) -> pulumi.Output[Optional[str]]:
        """
        The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
        """
        return pulumi.get(self, "request_template")

    @property
    @pulumi.getter(name="responseTemplate")
    def response_template(self) -> pulumi.Output[Optional[str]]:
        """
        The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
        """
        return pulumi.get(self, "response_template")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type name from the schema defined in the GraphQL API.
        """
        return pulumi.get(self, "type")

