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

__all__ = ['VoiceConnectorOrganizationArgs', 'VoiceConnectorOrganization']

@pulumi.input_type
class VoiceConnectorOrganizationArgs:
    def __init__(__self__, *,
                 routes: pulumi.Input[Sequence[pulumi.Input['VoiceConnectorOrganizationRouteArgs']]],
                 voice_connector_id: pulumi.Input[str],
                 disabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a VoiceConnectorOrganization resource.
        :param pulumi.Input[Sequence[pulumi.Input['VoiceConnectorOrganizationRouteArgs']]] routes: Set of call distribution properties defined for your SIP hosts. See route below for more details. Minimum of 1. Maximum of 20.
        :param pulumi.Input[str] voice_connector_id: The Amazon Chime Voice Connector ID.
        :param pulumi.Input[bool] disabled: When origination settings are disabled, inbound calls are not enabled for your Amazon Chime Voice Connector.
        """
        pulumi.set(__self__, "routes", routes)
        pulumi.set(__self__, "voice_connector_id", voice_connector_id)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)

    @property
    @pulumi.getter
    def routes(self) -> pulumi.Input[Sequence[pulumi.Input['VoiceConnectorOrganizationRouteArgs']]]:
        """
        Set of call distribution properties defined for your SIP hosts. See route below for more details. Minimum of 1. Maximum of 20.
        """
        return pulumi.get(self, "routes")

    @routes.setter
    def routes(self, value: pulumi.Input[Sequence[pulumi.Input['VoiceConnectorOrganizationRouteArgs']]]):
        pulumi.set(self, "routes", value)

    @property
    @pulumi.getter(name="voiceConnectorId")
    def voice_connector_id(self) -> pulumi.Input[str]:
        """
        The Amazon Chime Voice Connector ID.
        """
        return pulumi.get(self, "voice_connector_id")

    @voice_connector_id.setter
    def voice_connector_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "voice_connector_id", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When origination settings are disabled, inbound calls are not enabled for your Amazon Chime Voice Connector.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)


@pulumi.input_type
class _VoiceConnectorOrganizationState:
    def __init__(__self__, *,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 routes: Optional[pulumi.Input[Sequence[pulumi.Input['VoiceConnectorOrganizationRouteArgs']]]] = None,
                 voice_connector_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VoiceConnectorOrganization resources.
        :param pulumi.Input[bool] disabled: When origination settings are disabled, inbound calls are not enabled for your Amazon Chime Voice Connector.
        :param pulumi.Input[Sequence[pulumi.Input['VoiceConnectorOrganizationRouteArgs']]] routes: Set of call distribution properties defined for your SIP hosts. See route below for more details. Minimum of 1. Maximum of 20.
        :param pulumi.Input[str] voice_connector_id: The Amazon Chime Voice Connector ID.
        """
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if routes is not None:
            pulumi.set(__self__, "routes", routes)
        if voice_connector_id is not None:
            pulumi.set(__self__, "voice_connector_id", voice_connector_id)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When origination settings are disabled, inbound calls are not enabled for your Amazon Chime Voice Connector.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def routes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VoiceConnectorOrganizationRouteArgs']]]]:
        """
        Set of call distribution properties defined for your SIP hosts. See route below for more details. Minimum of 1. Maximum of 20.
        """
        return pulumi.get(self, "routes")

    @routes.setter
    def routes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VoiceConnectorOrganizationRouteArgs']]]]):
        pulumi.set(self, "routes", value)

    @property
    @pulumi.getter(name="voiceConnectorId")
    def voice_connector_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Chime Voice Connector ID.
        """
        return pulumi.get(self, "voice_connector_id")

    @voice_connector_id.setter
    def voice_connector_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "voice_connector_id", value)


class VoiceConnectorOrganization(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VoiceConnectorOrganizationRouteArgs']]]]] = None,
                 voice_connector_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Enable origination settings to control inbound calling to your SIP infrastructure.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        default_voice_connector = aws.chime.VoiceConnector("defaultVoiceConnector", require_encryption=True)
        default_voice_connector_organization = aws.chime.VoiceConnectorOrganization("defaultVoiceConnectorOrganization",
            disabled=False,
            voice_connector_id=default_voice_connector.id,
            routes=[
                aws.chime.VoiceConnectorOrganizationRouteArgs(
                    host="127.0.0.1",
                    port=8081,
                    protocol="TCP",
                    priority=1,
                    weight=1,
                ),
                aws.chime.VoiceConnectorOrganizationRouteArgs(
                    host="127.0.0.2",
                    port=8082,
                    protocol="TCP",
                    priority=2,
                    weight=10,
                ),
            ])
        ```

        ## Import

        Chime Voice Connector Origination can be imported using the `voice_connector_id`, e.g.

        ```sh
         $ pulumi import aws:chime/voiceConnectorOrganization:VoiceConnectorOrganization default abcdef1ghij2klmno3pqr4
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disabled: When origination settings are disabled, inbound calls are not enabled for your Amazon Chime Voice Connector.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VoiceConnectorOrganizationRouteArgs']]]] routes: Set of call distribution properties defined for your SIP hosts. See route below for more details. Minimum of 1. Maximum of 20.
        :param pulumi.Input[str] voice_connector_id: The Amazon Chime Voice Connector ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VoiceConnectorOrganizationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Enable origination settings to control inbound calling to your SIP infrastructure.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        default_voice_connector = aws.chime.VoiceConnector("defaultVoiceConnector", require_encryption=True)
        default_voice_connector_organization = aws.chime.VoiceConnectorOrganization("defaultVoiceConnectorOrganization",
            disabled=False,
            voice_connector_id=default_voice_connector.id,
            routes=[
                aws.chime.VoiceConnectorOrganizationRouteArgs(
                    host="127.0.0.1",
                    port=8081,
                    protocol="TCP",
                    priority=1,
                    weight=1,
                ),
                aws.chime.VoiceConnectorOrganizationRouteArgs(
                    host="127.0.0.2",
                    port=8082,
                    protocol="TCP",
                    priority=2,
                    weight=10,
                ),
            ])
        ```

        ## Import

        Chime Voice Connector Origination can be imported using the `voice_connector_id`, e.g.

        ```sh
         $ pulumi import aws:chime/voiceConnectorOrganization:VoiceConnectorOrganization default abcdef1ghij2klmno3pqr4
        ```

        :param str resource_name: The name of the resource.
        :param VoiceConnectorOrganizationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VoiceConnectorOrganizationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VoiceConnectorOrganizationRouteArgs']]]]] = None,
                 voice_connector_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = VoiceConnectorOrganizationArgs.__new__(VoiceConnectorOrganizationArgs)

            __props__.__dict__["disabled"] = disabled
            if routes is None and not opts.urn:
                raise TypeError("Missing required property 'routes'")
            __props__.__dict__["routes"] = routes
            if voice_connector_id is None and not opts.urn:
                raise TypeError("Missing required property 'voice_connector_id'")
            __props__.__dict__["voice_connector_id"] = voice_connector_id
        super(VoiceConnectorOrganization, __self__).__init__(
            'aws:chime/voiceConnectorOrganization:VoiceConnectorOrganization',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VoiceConnectorOrganizationRouteArgs']]]]] = None,
            voice_connector_id: Optional[pulumi.Input[str]] = None) -> 'VoiceConnectorOrganization':
        """
        Get an existing VoiceConnectorOrganization resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disabled: When origination settings are disabled, inbound calls are not enabled for your Amazon Chime Voice Connector.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VoiceConnectorOrganizationRouteArgs']]]] routes: Set of call distribution properties defined for your SIP hosts. See route below for more details. Minimum of 1. Maximum of 20.
        :param pulumi.Input[str] voice_connector_id: The Amazon Chime Voice Connector ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VoiceConnectorOrganizationState.__new__(_VoiceConnectorOrganizationState)

        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["routes"] = routes
        __props__.__dict__["voice_connector_id"] = voice_connector_id
        return VoiceConnectorOrganization(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        When origination settings are disabled, inbound calls are not enabled for your Amazon Chime Voice Connector.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def routes(self) -> pulumi.Output[Sequence['outputs.VoiceConnectorOrganizationRoute']]:
        """
        Set of call distribution properties defined for your SIP hosts. See route below for more details. Minimum of 1. Maximum of 20.
        """
        return pulumi.get(self, "routes")

    @property
    @pulumi.getter(name="voiceConnectorId")
    def voice_connector_id(self) -> pulumi.Output[str]:
        """
        The Amazon Chime Voice Connector ID.
        """
        return pulumi.get(self, "voice_connector_id")

