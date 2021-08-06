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

__all__ = ['EventConnectionArgs', 'EventConnection']

@pulumi.input_type
class EventConnectionArgs:
    def __init__(__self__, *,
                 auth_parameters: pulumi.Input['EventConnectionAuthParametersArgs'],
                 authorization_type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EventConnection resource.
        :param pulumi.Input['EventConnectionAuthParametersArgs'] auth_parameters: Parameters used for authorization. A maximum of 1 are allowed. Documented below.
        :param pulumi.Input[str] authorization_type: Choose the type of authorization to use for the connection. One of `API_KEY`,`BASIC`,`OAUTH_CLIENT_CREDENTIALS`.
        :param pulumi.Input[str] description: Enter a description for the connection. Maximum of 512 characters.
        :param pulumi.Input[str] name: The name of the new connection. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
        """
        pulumi.set(__self__, "auth_parameters", auth_parameters)
        pulumi.set(__self__, "authorization_type", authorization_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="authParameters")
    def auth_parameters(self) -> pulumi.Input['EventConnectionAuthParametersArgs']:
        """
        Parameters used for authorization. A maximum of 1 are allowed. Documented below.
        """
        return pulumi.get(self, "auth_parameters")

    @auth_parameters.setter
    def auth_parameters(self, value: pulumi.Input['EventConnectionAuthParametersArgs']):
        pulumi.set(self, "auth_parameters", value)

    @property
    @pulumi.getter(name="authorizationType")
    def authorization_type(self) -> pulumi.Input[str]:
        """
        Choose the type of authorization to use for the connection. One of `API_KEY`,`BASIC`,`OAUTH_CLIENT_CREDENTIALS`.
        """
        return pulumi.get(self, "authorization_type")

    @authorization_type.setter
    def authorization_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "authorization_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Enter a description for the connection. Maximum of 512 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the new connection. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _EventConnectionState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 auth_parameters: Optional[pulumi.Input['EventConnectionAuthParametersArgs']] = None,
                 authorization_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret_arn: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EventConnection resources.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the connection.
        :param pulumi.Input['EventConnectionAuthParametersArgs'] auth_parameters: Parameters used for authorization. A maximum of 1 are allowed. Documented below.
        :param pulumi.Input[str] authorization_type: Choose the type of authorization to use for the connection. One of `API_KEY`,`BASIC`,`OAUTH_CLIENT_CREDENTIALS`.
        :param pulumi.Input[str] description: Enter a description for the connection. Maximum of 512 characters.
        :param pulumi.Input[str] name: The name of the new connection. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
        :param pulumi.Input[str] secret_arn: The Amazon Resource Name (ARN) of the secret created from the authorization parameters specified for the connection.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if auth_parameters is not None:
            pulumi.set(__self__, "auth_parameters", auth_parameters)
        if authorization_type is not None:
            pulumi.set(__self__, "authorization_type", authorization_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if secret_arn is not None:
            pulumi.set(__self__, "secret_arn", secret_arn)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the connection.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="authParameters")
    def auth_parameters(self) -> Optional[pulumi.Input['EventConnectionAuthParametersArgs']]:
        """
        Parameters used for authorization. A maximum of 1 are allowed. Documented below.
        """
        return pulumi.get(self, "auth_parameters")

    @auth_parameters.setter
    def auth_parameters(self, value: Optional[pulumi.Input['EventConnectionAuthParametersArgs']]):
        pulumi.set(self, "auth_parameters", value)

    @property
    @pulumi.getter(name="authorizationType")
    def authorization_type(self) -> Optional[pulumi.Input[str]]:
        """
        Choose the type of authorization to use for the connection. One of `API_KEY`,`BASIC`,`OAUTH_CLIENT_CREDENTIALS`.
        """
        return pulumi.get(self, "authorization_type")

    @authorization_type.setter
    def authorization_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Enter a description for the connection. Maximum of 512 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the new connection. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="secretArn")
    def secret_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the secret created from the authorization parameters specified for the connection.
        """
        return pulumi.get(self, "secret_arn")

    @secret_arn.setter
    def secret_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_arn", value)


class EventConnection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_parameters: Optional[pulumi.Input[pulumi.InputType['EventConnectionAuthParametersArgs']]] = None,
                 authorization_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an EventBridge connection resource.

        > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.cloudwatch.EventConnection("test",
            auth_parameters=aws.cloudwatch.EventConnectionAuthParametersArgs(
                api_key=aws.cloudwatch.EventConnectionAuthParametersApiKeyArgs(
                    key="x-signature",
                    value="1234",
                ),
            ),
            authorization_type="API_KEY",
            description="A connection description")
        ```
        ### Basic Authorization

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.cloudwatch.EventConnection("test",
            auth_parameters=aws.cloudwatch.EventConnectionAuthParametersArgs(
                basic=aws.cloudwatch.EventConnectionAuthParametersBasicArgs(
                    password="Pass1234!",
                    username="user",
                ),
            ),
            authorization_type="BASIC",
            description="A connection description")
        ```
        ### OAuth Authorization

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.cloudwatch.EventConnection("test",
            auth_parameters=aws.cloudwatch.EventConnectionAuthParametersArgs(
                oauth=aws.cloudwatch.EventConnectionAuthParametersOauthArgs(
                    authorization_endpoint="https://auth.url.com/endpoint",
                    client_parameters=aws.cloudwatch.EventConnectionAuthParametersOauthClientParametersArgs(
                        client_id="1234567890",
                        client_secret="Pass1234!",
                    ),
                    http_method="GET",
                    oauth_http_parameters=aws.cloudwatch.EventConnectionAuthParametersOauthOauthHttpParametersArgs(
                        body=[{
                            "isValueSecret": False,
                            "key": "body-parameter-key",
                            "value": "body-parameter-value",
                        }],
                        header=[{
                            "isValueSecret": False,
                            "key": "header-parameter-key",
                            "value": "header-parameter-value",
                        }],
                        query_string=[{
                            "isValueSecret": False,
                            "key": "query-string-parameter-key",
                            "value": "query-string-parameter-value",
                        }],
                    ),
                ),
            ),
            authorization_type="OAUTH_CLIENT_CREDENTIALS",
            description="A connection description")
        ```
        ### Invocation Http Parameters

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.cloudwatch.EventConnection("test",
            auth_parameters=aws.cloudwatch.EventConnectionAuthParametersArgs(
                basic=aws.cloudwatch.EventConnectionAuthParametersBasicArgs(
                    password="Pass1234!",
                    username="user",
                ),
                invocation_http_parameters=aws.cloudwatch.EventConnectionAuthParametersInvocationHttpParametersArgs(
                    body=[
                        {
                            "isValueSecret": False,
                            "key": "body-parameter-key",
                            "value": "body-parameter-value",
                        },
                        {
                            "isValueSecret": True,
                            "key": "body-parameter-key2",
                            "value": "body-parameter-value2",
                        },
                    ],
                    header=[{
                        "isValueSecret": False,
                        "key": "header-parameter-key",
                        "value": "header-parameter-value",
                    }],
                    query_string=[{
                        "isValueSecret": False,
                        "key": "query-string-parameter-key",
                        "value": "query-string-parameter-value",
                    }],
                ),
            ),
            authorization_type="BASIC",
            description="A connection description")
        ```

        ## Import

        EventBridge Connection can be imported using the `name`, e.g. console

        ```sh
         $ pulumi import aws:cloudwatch/eventConnection:EventConnection test ngrok-connection
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['EventConnectionAuthParametersArgs']] auth_parameters: Parameters used for authorization. A maximum of 1 are allowed. Documented below.
        :param pulumi.Input[str] authorization_type: Choose the type of authorization to use for the connection. One of `API_KEY`,`BASIC`,`OAUTH_CLIENT_CREDENTIALS`.
        :param pulumi.Input[str] description: Enter a description for the connection. Maximum of 512 characters.
        :param pulumi.Input[str] name: The name of the new connection. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EventConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an EventBridge connection resource.

        > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.cloudwatch.EventConnection("test",
            auth_parameters=aws.cloudwatch.EventConnectionAuthParametersArgs(
                api_key=aws.cloudwatch.EventConnectionAuthParametersApiKeyArgs(
                    key="x-signature",
                    value="1234",
                ),
            ),
            authorization_type="API_KEY",
            description="A connection description")
        ```
        ### Basic Authorization

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.cloudwatch.EventConnection("test",
            auth_parameters=aws.cloudwatch.EventConnectionAuthParametersArgs(
                basic=aws.cloudwatch.EventConnectionAuthParametersBasicArgs(
                    password="Pass1234!",
                    username="user",
                ),
            ),
            authorization_type="BASIC",
            description="A connection description")
        ```
        ### OAuth Authorization

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.cloudwatch.EventConnection("test",
            auth_parameters=aws.cloudwatch.EventConnectionAuthParametersArgs(
                oauth=aws.cloudwatch.EventConnectionAuthParametersOauthArgs(
                    authorization_endpoint="https://auth.url.com/endpoint",
                    client_parameters=aws.cloudwatch.EventConnectionAuthParametersOauthClientParametersArgs(
                        client_id="1234567890",
                        client_secret="Pass1234!",
                    ),
                    http_method="GET",
                    oauth_http_parameters=aws.cloudwatch.EventConnectionAuthParametersOauthOauthHttpParametersArgs(
                        body=[{
                            "isValueSecret": False,
                            "key": "body-parameter-key",
                            "value": "body-parameter-value",
                        }],
                        header=[{
                            "isValueSecret": False,
                            "key": "header-parameter-key",
                            "value": "header-parameter-value",
                        }],
                        query_string=[{
                            "isValueSecret": False,
                            "key": "query-string-parameter-key",
                            "value": "query-string-parameter-value",
                        }],
                    ),
                ),
            ),
            authorization_type="OAUTH_CLIENT_CREDENTIALS",
            description="A connection description")
        ```
        ### Invocation Http Parameters

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.cloudwatch.EventConnection("test",
            auth_parameters=aws.cloudwatch.EventConnectionAuthParametersArgs(
                basic=aws.cloudwatch.EventConnectionAuthParametersBasicArgs(
                    password="Pass1234!",
                    username="user",
                ),
                invocation_http_parameters=aws.cloudwatch.EventConnectionAuthParametersInvocationHttpParametersArgs(
                    body=[
                        {
                            "isValueSecret": False,
                            "key": "body-parameter-key",
                            "value": "body-parameter-value",
                        },
                        {
                            "isValueSecret": True,
                            "key": "body-parameter-key2",
                            "value": "body-parameter-value2",
                        },
                    ],
                    header=[{
                        "isValueSecret": False,
                        "key": "header-parameter-key",
                        "value": "header-parameter-value",
                    }],
                    query_string=[{
                        "isValueSecret": False,
                        "key": "query-string-parameter-key",
                        "value": "query-string-parameter-value",
                    }],
                ),
            ),
            authorization_type="BASIC",
            description="A connection description")
        ```

        ## Import

        EventBridge Connection can be imported using the `name`, e.g. console

        ```sh
         $ pulumi import aws:cloudwatch/eventConnection:EventConnection test ngrok-connection
        ```

        :param str resource_name: The name of the resource.
        :param EventConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EventConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_parameters: Optional[pulumi.Input[pulumi.InputType['EventConnectionAuthParametersArgs']]] = None,
                 authorization_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
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
            __props__ = EventConnectionArgs.__new__(EventConnectionArgs)

            if auth_parameters is None and not opts.urn:
                raise TypeError("Missing required property 'auth_parameters'")
            __props__.__dict__["auth_parameters"] = auth_parameters
            if authorization_type is None and not opts.urn:
                raise TypeError("Missing required property 'authorization_type'")
            __props__.__dict__["authorization_type"] = authorization_type
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["arn"] = None
            __props__.__dict__["secret_arn"] = None
        super(EventConnection, __self__).__init__(
            'aws:cloudwatch/eventConnection:EventConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            auth_parameters: Optional[pulumi.Input[pulumi.InputType['EventConnectionAuthParametersArgs']]] = None,
            authorization_type: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            secret_arn: Optional[pulumi.Input[str]] = None) -> 'EventConnection':
        """
        Get an existing EventConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the connection.
        :param pulumi.Input[pulumi.InputType['EventConnectionAuthParametersArgs']] auth_parameters: Parameters used for authorization. A maximum of 1 are allowed. Documented below.
        :param pulumi.Input[str] authorization_type: Choose the type of authorization to use for the connection. One of `API_KEY`,`BASIC`,`OAUTH_CLIENT_CREDENTIALS`.
        :param pulumi.Input[str] description: Enter a description for the connection. Maximum of 512 characters.
        :param pulumi.Input[str] name: The name of the new connection. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
        :param pulumi.Input[str] secret_arn: The Amazon Resource Name (ARN) of the secret created from the authorization parameters specified for the connection.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EventConnectionState.__new__(_EventConnectionState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["auth_parameters"] = auth_parameters
        __props__.__dict__["authorization_type"] = authorization_type
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["secret_arn"] = secret_arn
        return EventConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the connection.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="authParameters")
    def auth_parameters(self) -> pulumi.Output['outputs.EventConnectionAuthParameters']:
        """
        Parameters used for authorization. A maximum of 1 are allowed. Documented below.
        """
        return pulumi.get(self, "auth_parameters")

    @property
    @pulumi.getter(name="authorizationType")
    def authorization_type(self) -> pulumi.Output[str]:
        """
        Choose the type of authorization to use for the connection. One of `API_KEY`,`BASIC`,`OAUTH_CLIENT_CREDENTIALS`.
        """
        return pulumi.get(self, "authorization_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Enter a description for the connection. Maximum of 512 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the new connection. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="secretArn")
    def secret_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the secret created from the authorization parameters specified for the connection.
        """
        return pulumi.get(self, "secret_arn")

