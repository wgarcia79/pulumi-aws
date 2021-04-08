# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AuthorizationRuleArgs', 'AuthorizationRule']

@pulumi.input_type
class AuthorizationRuleArgs:
    def __init__(__self__, *,
                 client_vpn_endpoint_id: pulumi.Input[str],
                 target_network_cidr: pulumi.Input[str],
                 access_group_id: Optional[pulumi.Input[str]] = None,
                 authorize_all_groups: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AuthorizationRule resource.
        :param pulumi.Input[str] client_vpn_endpoint_id: The ID of the Client VPN endpoint.
        :param pulumi.Input[str] target_network_cidr: The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
        :param pulumi.Input[str] access_group_id: The ID of the group to which the authorization rule grants access. One of `access_group_id` or `authorize_all_groups` must be set.
        :param pulumi.Input[bool] authorize_all_groups: Indicates whether the authorization rule grants access to all clients. One of `access_group_id` or `authorize_all_groups` must be set.
        :param pulumi.Input[str] description: A brief description of the authorization rule.
        """
        pulumi.set(__self__, "client_vpn_endpoint_id", client_vpn_endpoint_id)
        pulumi.set(__self__, "target_network_cidr", target_network_cidr)
        if access_group_id is not None:
            pulumi.set(__self__, "access_group_id", access_group_id)
        if authorize_all_groups is not None:
            pulumi.set(__self__, "authorize_all_groups", authorize_all_groups)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="clientVpnEndpointId")
    def client_vpn_endpoint_id(self) -> pulumi.Input[str]:
        """
        The ID of the Client VPN endpoint.
        """
        return pulumi.get(self, "client_vpn_endpoint_id")

    @client_vpn_endpoint_id.setter
    def client_vpn_endpoint_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_vpn_endpoint_id", value)

    @property
    @pulumi.getter(name="targetNetworkCidr")
    def target_network_cidr(self) -> pulumi.Input[str]:
        """
        The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
        """
        return pulumi.get(self, "target_network_cidr")

    @target_network_cidr.setter
    def target_network_cidr(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_network_cidr", value)

    @property
    @pulumi.getter(name="accessGroupId")
    def access_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the group to which the authorization rule grants access. One of `access_group_id` or `authorize_all_groups` must be set.
        """
        return pulumi.get(self, "access_group_id")

    @access_group_id.setter
    def access_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_group_id", value)

    @property
    @pulumi.getter(name="authorizeAllGroups")
    def authorize_all_groups(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the authorization rule grants access to all clients. One of `access_group_id` or `authorize_all_groups` must be set.
        """
        return pulumi.get(self, "authorize_all_groups")

    @authorize_all_groups.setter
    def authorize_all_groups(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "authorize_all_groups", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A brief description of the authorization rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _AuthorizationRuleState:
    def __init__(__self__, *,
                 access_group_id: Optional[pulumi.Input[str]] = None,
                 authorize_all_groups: Optional[pulumi.Input[bool]] = None,
                 client_vpn_endpoint_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 target_network_cidr: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AuthorizationRule resources.
        :param pulumi.Input[str] access_group_id: The ID of the group to which the authorization rule grants access. One of `access_group_id` or `authorize_all_groups` must be set.
        :param pulumi.Input[bool] authorize_all_groups: Indicates whether the authorization rule grants access to all clients. One of `access_group_id` or `authorize_all_groups` must be set.
        :param pulumi.Input[str] client_vpn_endpoint_id: The ID of the Client VPN endpoint.
        :param pulumi.Input[str] description: A brief description of the authorization rule.
        :param pulumi.Input[str] target_network_cidr: The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
        """
        if access_group_id is not None:
            pulumi.set(__self__, "access_group_id", access_group_id)
        if authorize_all_groups is not None:
            pulumi.set(__self__, "authorize_all_groups", authorize_all_groups)
        if client_vpn_endpoint_id is not None:
            pulumi.set(__self__, "client_vpn_endpoint_id", client_vpn_endpoint_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if target_network_cidr is not None:
            pulumi.set(__self__, "target_network_cidr", target_network_cidr)

    @property
    @pulumi.getter(name="accessGroupId")
    def access_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the group to which the authorization rule grants access. One of `access_group_id` or `authorize_all_groups` must be set.
        """
        return pulumi.get(self, "access_group_id")

    @access_group_id.setter
    def access_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_group_id", value)

    @property
    @pulumi.getter(name="authorizeAllGroups")
    def authorize_all_groups(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the authorization rule grants access to all clients. One of `access_group_id` or `authorize_all_groups` must be set.
        """
        return pulumi.get(self, "authorize_all_groups")

    @authorize_all_groups.setter
    def authorize_all_groups(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "authorize_all_groups", value)

    @property
    @pulumi.getter(name="clientVpnEndpointId")
    def client_vpn_endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Client VPN endpoint.
        """
        return pulumi.get(self, "client_vpn_endpoint_id")

    @client_vpn_endpoint_id.setter
    def client_vpn_endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_vpn_endpoint_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A brief description of the authorization rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="targetNetworkCidr")
    def target_network_cidr(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
        """
        return pulumi.get(self, "target_network_cidr")

    @target_network_cidr.setter
    def target_network_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_network_cidr", value)


class AuthorizationRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_group_id: Optional[pulumi.Input[str]] = None,
                 authorize_all_groups: Optional[pulumi.Input[bool]] = None,
                 client_vpn_endpoint_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 target_network_cidr: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides authorization rules for AWS Client VPN endpoints. For more information on usage, please see the
        [AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2clientvpn.AuthorizationRule("example",
            client_vpn_endpoint_id=aws_ec2_client_vpn_endpoint["example"]["id"],
            target_network_cidr=aws_subnet["example"]["cidr_block"],
            authorize_all_groups=True)
        ```

        ## Import

        AWS Client VPN authorization rules can be imported using the endpoint ID and target network CIDR. If there is a specific group name that is included as well. All values are separated by a `,`.

        ```sh
         $ pulumi import aws:ec2clientvpn/authorizationRule:AuthorizationRule example cvpn-endpoint-0ac3a1abbccddd666,10.1.0.0/24
        ```

        ```sh
         $ pulumi import aws:ec2clientvpn/authorizationRule:AuthorizationRule example cvpn-endpoint-0ac3a1abbccddd666,10.1.0.0/24,team-a
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_group_id: The ID of the group to which the authorization rule grants access. One of `access_group_id` or `authorize_all_groups` must be set.
        :param pulumi.Input[bool] authorize_all_groups: Indicates whether the authorization rule grants access to all clients. One of `access_group_id` or `authorize_all_groups` must be set.
        :param pulumi.Input[str] client_vpn_endpoint_id: The ID of the Client VPN endpoint.
        :param pulumi.Input[str] description: A brief description of the authorization rule.
        :param pulumi.Input[str] target_network_cidr: The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthorizationRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides authorization rules for AWS Client VPN endpoints. For more information on usage, please see the
        [AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2clientvpn.AuthorizationRule("example",
            client_vpn_endpoint_id=aws_ec2_client_vpn_endpoint["example"]["id"],
            target_network_cidr=aws_subnet["example"]["cidr_block"],
            authorize_all_groups=True)
        ```

        ## Import

        AWS Client VPN authorization rules can be imported using the endpoint ID and target network CIDR. If there is a specific group name that is included as well. All values are separated by a `,`.

        ```sh
         $ pulumi import aws:ec2clientvpn/authorizationRule:AuthorizationRule example cvpn-endpoint-0ac3a1abbccddd666,10.1.0.0/24
        ```

        ```sh
         $ pulumi import aws:ec2clientvpn/authorizationRule:AuthorizationRule example cvpn-endpoint-0ac3a1abbccddd666,10.1.0.0/24,team-a
        ```

        :param str resource_name: The name of the resource.
        :param AuthorizationRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthorizationRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_group_id: Optional[pulumi.Input[str]] = None,
                 authorize_all_groups: Optional[pulumi.Input[bool]] = None,
                 client_vpn_endpoint_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 target_network_cidr: Optional[pulumi.Input[str]] = None,
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
            __props__ = AuthorizationRuleArgs.__new__(AuthorizationRuleArgs)

            __props__.__dict__["access_group_id"] = access_group_id
            __props__.__dict__["authorize_all_groups"] = authorize_all_groups
            if client_vpn_endpoint_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_vpn_endpoint_id'")
            __props__.__dict__["client_vpn_endpoint_id"] = client_vpn_endpoint_id
            __props__.__dict__["description"] = description
            if target_network_cidr is None and not opts.urn:
                raise TypeError("Missing required property 'target_network_cidr'")
            __props__.__dict__["target_network_cidr"] = target_network_cidr
        super(AuthorizationRule, __self__).__init__(
            'aws:ec2clientvpn/authorizationRule:AuthorizationRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_group_id: Optional[pulumi.Input[str]] = None,
            authorize_all_groups: Optional[pulumi.Input[bool]] = None,
            client_vpn_endpoint_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            target_network_cidr: Optional[pulumi.Input[str]] = None) -> 'AuthorizationRule':
        """
        Get an existing AuthorizationRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_group_id: The ID of the group to which the authorization rule grants access. One of `access_group_id` or `authorize_all_groups` must be set.
        :param pulumi.Input[bool] authorize_all_groups: Indicates whether the authorization rule grants access to all clients. One of `access_group_id` or `authorize_all_groups` must be set.
        :param pulumi.Input[str] client_vpn_endpoint_id: The ID of the Client VPN endpoint.
        :param pulumi.Input[str] description: A brief description of the authorization rule.
        :param pulumi.Input[str] target_network_cidr: The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthorizationRuleState.__new__(_AuthorizationRuleState)

        __props__.__dict__["access_group_id"] = access_group_id
        __props__.__dict__["authorize_all_groups"] = authorize_all_groups
        __props__.__dict__["client_vpn_endpoint_id"] = client_vpn_endpoint_id
        __props__.__dict__["description"] = description
        __props__.__dict__["target_network_cidr"] = target_network_cidr
        return AuthorizationRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessGroupId")
    def access_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the group to which the authorization rule grants access. One of `access_group_id` or `authorize_all_groups` must be set.
        """
        return pulumi.get(self, "access_group_id")

    @property
    @pulumi.getter(name="authorizeAllGroups")
    def authorize_all_groups(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether the authorization rule grants access to all clients. One of `access_group_id` or `authorize_all_groups` must be set.
        """
        return pulumi.get(self, "authorize_all_groups")

    @property
    @pulumi.getter(name="clientVpnEndpointId")
    def client_vpn_endpoint_id(self) -> pulumi.Output[str]:
        """
        The ID of the Client VPN endpoint.
        """
        return pulumi.get(self, "client_vpn_endpoint_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A brief description of the authorization rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="targetNetworkCidr")
    def target_network_cidr(self) -> pulumi.Output[str]:
        """
        The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
        """
        return pulumi.get(self, "target_network_cidr")

