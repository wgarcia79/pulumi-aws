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

__all__ = ['DocumentationPartArgs', 'DocumentationPart']

@pulumi.input_type
class DocumentationPartArgs:
    def __init__(__self__, *,
                 location: pulumi.Input['DocumentationPartLocationArgs'],
                 properties: pulumi.Input[str],
                 rest_api_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a DocumentationPart resource.
        :param pulumi.Input['DocumentationPartLocationArgs'] location: The location of the targeted API entity of the to-be-created documentation part. See below.
        :param pulumi.Input[str] properties: A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
        :param pulumi.Input[str] rest_api_id: The ID of the associated Rest API
        """
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "properties", properties)
        pulumi.set(__self__, "rest_api_id", rest_api_id)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input['DocumentationPartLocationArgs']:
        """
        The location of the targeted API entity of the to-be-created documentation part. See below.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input['DocumentationPartLocationArgs']):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Input[str]:
        """
        A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: pulumi.Input[str]):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter(name="restApiId")
    def rest_api_id(self) -> pulumi.Input[str]:
        """
        The ID of the associated Rest API
        """
        return pulumi.get(self, "rest_api_id")

    @rest_api_id.setter
    def rest_api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "rest_api_id", value)


@pulumi.input_type
class _DocumentationPartState:
    def __init__(__self__, *,
                 location: Optional[pulumi.Input['DocumentationPartLocationArgs']] = None,
                 properties: Optional[pulumi.Input[str]] = None,
                 rest_api_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DocumentationPart resources.
        :param pulumi.Input['DocumentationPartLocationArgs'] location: The location of the targeted API entity of the to-be-created documentation part. See below.
        :param pulumi.Input[str] properties: A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
        :param pulumi.Input[str] rest_api_id: The ID of the associated Rest API
        """
        if location is not None:
            pulumi.set(__self__, "location", location)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)
        if rest_api_id is not None:
            pulumi.set(__self__, "rest_api_id", rest_api_id)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input['DocumentationPartLocationArgs']]:
        """
        The location of the targeted API entity of the to-be-created documentation part. See below.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input['DocumentationPartLocationArgs']]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input[str]]:
        """
        A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter(name="restApiId")
    def rest_api_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the associated Rest API
        """
        return pulumi.get(self, "rest_api_id")

    @rest_api_id.setter
    def rest_api_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rest_api_id", value)


class DocumentationPart(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[pulumi.InputType['DocumentationPartLocationArgs']]] = None,
                 properties: Optional[pulumi.Input[str]] = None,
                 rest_api_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a settings of an API Gateway Documentation Part.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_rest_api = aws.apigateway.RestApi("exampleRestApi")
        example_documentation_part = aws.apigateway.DocumentationPart("exampleDocumentationPart",
            location=aws.apigateway.DocumentationPartLocationArgs(
                type="METHOD",
                method="GET",
                path="/example",
            ),
            properties="{\"description\":\"Example description\"}",
            rest_api_id=example_rest_api.id)
        ```

        ## Import

        API Gateway documentation_parts can be imported using `REST-API-ID/DOC-PART-ID`, e.g.

        ```sh
         $ pulumi import aws:apigateway/documentationPart:DocumentationPart example 5i4e1ko720/3oyy3t
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DocumentationPartLocationArgs']] location: The location of the targeted API entity of the to-be-created documentation part. See below.
        :param pulumi.Input[str] properties: A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
        :param pulumi.Input[str] rest_api_id: The ID of the associated Rest API
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DocumentationPartArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a settings of an API Gateway Documentation Part.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_rest_api = aws.apigateway.RestApi("exampleRestApi")
        example_documentation_part = aws.apigateway.DocumentationPart("exampleDocumentationPart",
            location=aws.apigateway.DocumentationPartLocationArgs(
                type="METHOD",
                method="GET",
                path="/example",
            ),
            properties="{\"description\":\"Example description\"}",
            rest_api_id=example_rest_api.id)
        ```

        ## Import

        API Gateway documentation_parts can be imported using `REST-API-ID/DOC-PART-ID`, e.g.

        ```sh
         $ pulumi import aws:apigateway/documentationPart:DocumentationPart example 5i4e1ko720/3oyy3t
        ```

        :param str resource_name: The name of the resource.
        :param DocumentationPartArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DocumentationPartArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[pulumi.InputType['DocumentationPartLocationArgs']]] = None,
                 properties: Optional[pulumi.Input[str]] = None,
                 rest_api_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = DocumentationPartArgs.__new__(DocumentationPartArgs)

            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            if properties is None and not opts.urn:
                raise TypeError("Missing required property 'properties'")
            __props__.__dict__["properties"] = properties
            if rest_api_id is None and not opts.urn:
                raise TypeError("Missing required property 'rest_api_id'")
            __props__.__dict__["rest_api_id"] = rest_api_id
        super(DocumentationPart, __self__).__init__(
            'aws:apigateway/documentationPart:DocumentationPart',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            location: Optional[pulumi.Input[pulumi.InputType['DocumentationPartLocationArgs']]] = None,
            properties: Optional[pulumi.Input[str]] = None,
            rest_api_id: Optional[pulumi.Input[str]] = None) -> 'DocumentationPart':
        """
        Get an existing DocumentationPart resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DocumentationPartLocationArgs']] location: The location of the targeted API entity of the to-be-created documentation part. See below.
        :param pulumi.Input[str] properties: A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
        :param pulumi.Input[str] rest_api_id: The ID of the associated Rest API
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DocumentationPartState.__new__(_DocumentationPartState)

        __props__.__dict__["location"] = location
        __props__.__dict__["properties"] = properties
        __props__.__dict__["rest_api_id"] = rest_api_id
        return DocumentationPart(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output['outputs.DocumentationPartLocation']:
        """
        The location of the targeted API entity of the to-be-created documentation part. See below.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output[str]:
        """
        A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter(name="restApiId")
    def rest_api_id(self) -> pulumi.Output[str]:
        """
        The ID of the associated Rest API
        """
        return pulumi.get(self, "rest_api_id")

