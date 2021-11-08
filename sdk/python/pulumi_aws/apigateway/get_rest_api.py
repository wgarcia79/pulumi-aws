# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetRestApiResult',
    'AwaitableGetRestApiResult',
    'get_rest_api',
    'get_rest_api_output',
]

@pulumi.output_type
class GetRestApiResult:
    """
    A collection of values returned by getRestApi.
    """
    def __init__(__self__, api_key_source=None, arn=None, binary_media_types=None, description=None, endpoint_configurations=None, execution_arn=None, id=None, minimum_compression_size=None, name=None, policy=None, root_resource_id=None, tags=None):
        if api_key_source and not isinstance(api_key_source, str):
            raise TypeError("Expected argument 'api_key_source' to be a str")
        pulumi.set(__self__, "api_key_source", api_key_source)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if binary_media_types and not isinstance(binary_media_types, list):
            raise TypeError("Expected argument 'binary_media_types' to be a list")
        pulumi.set(__self__, "binary_media_types", binary_media_types)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if endpoint_configurations and not isinstance(endpoint_configurations, list):
            raise TypeError("Expected argument 'endpoint_configurations' to be a list")
        pulumi.set(__self__, "endpoint_configurations", endpoint_configurations)
        if execution_arn and not isinstance(execution_arn, str):
            raise TypeError("Expected argument 'execution_arn' to be a str")
        pulumi.set(__self__, "execution_arn", execution_arn)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if minimum_compression_size and not isinstance(minimum_compression_size, int):
            raise TypeError("Expected argument 'minimum_compression_size' to be a int")
        pulumi.set(__self__, "minimum_compression_size", minimum_compression_size)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if policy and not isinstance(policy, str):
            raise TypeError("Expected argument 'policy' to be a str")
        pulumi.set(__self__, "policy", policy)
        if root_resource_id and not isinstance(root_resource_id, str):
            raise TypeError("Expected argument 'root_resource_id' to be a str")
        pulumi.set(__self__, "root_resource_id", root_resource_id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="apiKeySource")
    def api_key_source(self) -> str:
        """
        The source of the API key for requests.
        """
        return pulumi.get(self, "api_key_source")

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The ARN of the REST API.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="binaryMediaTypes")
    def binary_media_types(self) -> Sequence[str]:
        """
        The list of binary media types supported by the REST API.
        """
        return pulumi.get(self, "binary_media_types")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the REST API.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endpointConfigurations")
    def endpoint_configurations(self) -> Sequence['outputs.GetRestApiEndpointConfigurationResult']:
        """
        The endpoint configuration of this RestApi showing the endpoint types of the API.
        """
        return pulumi.get(self, "endpoint_configurations")

    @property
    @pulumi.getter(name="executionArn")
    def execution_arn(self) -> str:
        """
        The execution ARN part to be used in `lambda_permission`'s `source_arn` when allowing API Gateway to invoke a Lambda function, e.g., `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j`, which can be concatenated with allowed stage, method and resource path.
        """
        return pulumi.get(self, "execution_arn")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="minimumCompressionSize")
    def minimum_compression_size(self) -> int:
        """
        Minimum response size to compress for the REST API.
        """
        return pulumi.get(self, "minimum_compression_size")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policy(self) -> str:
        """
        JSON formatted policy document that controls access to the API Gateway.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="rootResourceId")
    def root_resource_id(self) -> str:
        """
        Set to the ID of the API Gateway Resource on the found REST API where the route matches '/'.
        """
        return pulumi.get(self, "root_resource_id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Key-value map of resource tags.
        """
        return pulumi.get(self, "tags")


class AwaitableGetRestApiResult(GetRestApiResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRestApiResult(
            api_key_source=self.api_key_source,
            arn=self.arn,
            binary_media_types=self.binary_media_types,
            description=self.description,
            endpoint_configurations=self.endpoint_configurations,
            execution_arn=self.execution_arn,
            id=self.id,
            minimum_compression_size=self.minimum_compression_size,
            name=self.name,
            policy=self.policy,
            root_resource_id=self.root_resource_id,
            tags=self.tags)


def get_rest_api(name: Optional[str] = None,
                 tags: Optional[Mapping[str, str]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRestApiResult:
    """
    Use this data source to get the id and root_resource_id of a REST API in
    API Gateway. To fetch the REST API you must provide a name to match against.
    As there is no unique name constraint on REST APIs this data source will
    error if there is more than one match.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    my_rest_api = aws.apigateway.get_rest_api(name="my-rest-api")
    ```


    :param str name: The name of the REST API to look up. If no REST API is found with this name, an error will be returned. If multiple REST APIs are found with this name, an error will be returned.
    :param Mapping[str, str] tags: Key-value map of resource tags.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:apigateway/getRestApi:getRestApi', __args__, opts=opts, typ=GetRestApiResult).value

    return AwaitableGetRestApiResult(
        api_key_source=__ret__.api_key_source,
        arn=__ret__.arn,
        binary_media_types=__ret__.binary_media_types,
        description=__ret__.description,
        endpoint_configurations=__ret__.endpoint_configurations,
        execution_arn=__ret__.execution_arn,
        id=__ret__.id,
        minimum_compression_size=__ret__.minimum_compression_size,
        name=__ret__.name,
        policy=__ret__.policy,
        root_resource_id=__ret__.root_resource_id,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_rest_api)
def get_rest_api_output(name: Optional[pulumi.Input[str]] = None,
                        tags: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRestApiResult]:
    """
    Use this data source to get the id and root_resource_id of a REST API in
    API Gateway. To fetch the REST API you must provide a name to match against.
    As there is no unique name constraint on REST APIs this data source will
    error if there is more than one match.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    my_rest_api = aws.apigateway.get_rest_api(name="my-rest-api")
    ```


    :param str name: The name of the REST API to look up. If no REST API is found with this name, an error will be returned. If multiple REST APIs are found with this name, an error will be returned.
    :param Mapping[str, str] tags: Key-value map of resource tags.
    """
    ...
