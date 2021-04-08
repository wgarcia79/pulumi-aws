# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['OpenIdConnectProviderArgs', 'OpenIdConnectProvider']

@pulumi.input_type
class OpenIdConnectProviderArgs:
    def __init__(__self__, *,
                 client_id_lists: pulumi.Input[Sequence[pulumi.Input[str]]],
                 thumbprint_lists: pulumi.Input[Sequence[pulumi.Input[str]]],
                 url: pulumi.Input[str],
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a OpenIdConnectProvider resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] client_id_lists: A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the client_id parameter on OAuth requests.)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] thumbprint_lists: A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
        :param pulumi.Input[str] url: The URL of the identity provider. Corresponds to the _iss_ claim.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of resource tags for the IAM OIDC provider.
        """
        pulumi.set(__self__, "client_id_lists", client_id_lists)
        pulumi.set(__self__, "thumbprint_lists", thumbprint_lists)
        pulumi.set(__self__, "url", url)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="clientIdLists")
    def client_id_lists(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the client_id parameter on OAuth requests.)
        """
        return pulumi.get(self, "client_id_lists")

    @client_id_lists.setter
    def client_id_lists(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "client_id_lists", value)

    @property
    @pulumi.getter(name="thumbprintLists")
    def thumbprint_lists(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
        """
        return pulumi.get(self, "thumbprint_lists")

    @thumbprint_lists.setter
    def thumbprint_lists(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "thumbprint_lists", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        """
        The URL of the identity provider. Corresponds to the _iss_ claim.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of resource tags for the IAM OIDC provider.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _OpenIdConnectProviderState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 client_id_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 thumbprint_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OpenIdConnectProvider resources.
        :param pulumi.Input[str] arn: The ARN assigned by AWS for this provider.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] client_id_lists: A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the client_id parameter on OAuth requests.)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of resource tags for the IAM OIDC provider.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] thumbprint_lists: A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
        :param pulumi.Input[str] url: The URL of the identity provider. Corresponds to the _iss_ claim.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if client_id_lists is not None:
            pulumi.set(__self__, "client_id_lists", client_id_lists)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if thumbprint_lists is not None:
            pulumi.set(__self__, "thumbprint_lists", thumbprint_lists)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN assigned by AWS for this provider.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="clientIdLists")
    def client_id_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the client_id parameter on OAuth requests.)
        """
        return pulumi.get(self, "client_id_lists")

    @client_id_lists.setter
    def client_id_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "client_id_lists", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of resource tags for the IAM OIDC provider.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="thumbprintLists")
    def thumbprint_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
        """
        return pulumi.get(self, "thumbprint_lists")

    @thumbprint_lists.setter
    def thumbprint_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "thumbprint_lists", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the identity provider. Corresponds to the _iss_ claim.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class OpenIdConnectProvider(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 thumbprint_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an IAM OpenID Connect provider.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.iam.OpenIdConnectProvider("default",
            client_id_lists=["266362248691-342342xasdasdasda-apps.googleusercontent.com"],
            thumbprint_lists=[],
            url="https://accounts.google.com")
        ```

        ## Import

        IAM OpenID Connect Providers can be imported using the `arn`, e.g.

        ```sh
         $ pulumi import aws:iam/openIdConnectProvider:OpenIdConnectProvider default arn:aws:iam::123456789012:oidc-provider/accounts.google.com
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] client_id_lists: A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the client_id parameter on OAuth requests.)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of resource tags for the IAM OIDC provider.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] thumbprint_lists: A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
        :param pulumi.Input[str] url: The URL of the identity provider. Corresponds to the _iss_ claim.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OpenIdConnectProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an IAM OpenID Connect provider.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.iam.OpenIdConnectProvider("default",
            client_id_lists=["266362248691-342342xasdasdasda-apps.googleusercontent.com"],
            thumbprint_lists=[],
            url="https://accounts.google.com")
        ```

        ## Import

        IAM OpenID Connect Providers can be imported using the `arn`, e.g.

        ```sh
         $ pulumi import aws:iam/openIdConnectProvider:OpenIdConnectProvider default arn:aws:iam::123456789012:oidc-provider/accounts.google.com
        ```

        :param str resource_name: The name of the resource.
        :param OpenIdConnectProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OpenIdConnectProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 thumbprint_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 url: Optional[pulumi.Input[str]] = None,
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
            __props__ = OpenIdConnectProviderArgs.__new__(OpenIdConnectProviderArgs)

            if client_id_lists is None and not opts.urn:
                raise TypeError("Missing required property 'client_id_lists'")
            __props__.__dict__["client_id_lists"] = client_id_lists
            __props__.__dict__["tags"] = tags
            if thumbprint_lists is None and not opts.urn:
                raise TypeError("Missing required property 'thumbprint_lists'")
            __props__.__dict__["thumbprint_lists"] = thumbprint_lists
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            __props__.__dict__["arn"] = None
        super(OpenIdConnectProvider, __self__).__init__(
            'aws:iam/openIdConnectProvider:OpenIdConnectProvider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            client_id_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            thumbprint_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'OpenIdConnectProvider':
        """
        Get an existing OpenIdConnectProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN assigned by AWS for this provider.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] client_id_lists: A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the client_id parameter on OAuth requests.)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of resource tags for the IAM OIDC provider.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] thumbprint_lists: A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
        :param pulumi.Input[str] url: The URL of the identity provider. Corresponds to the _iss_ claim.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OpenIdConnectProviderState.__new__(_OpenIdConnectProviderState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["client_id_lists"] = client_id_lists
        __props__.__dict__["tags"] = tags
        __props__.__dict__["thumbprint_lists"] = thumbprint_lists
        __props__.__dict__["url"] = url
        return OpenIdConnectProvider(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN assigned by AWS for this provider.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="clientIdLists")
    def client_id_lists(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the client_id parameter on OAuth requests.)
        """
        return pulumi.get(self, "client_id_lists")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map of resource tags for the IAM OIDC provider.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="thumbprintLists")
    def thumbprint_lists(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
        """
        return pulumi.get(self, "thumbprint_lists")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL of the identity provider. Corresponds to the _iss_ claim.
        """
        return pulumi.get(self, "url")

