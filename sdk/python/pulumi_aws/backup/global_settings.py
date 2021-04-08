# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GlobalSettingsArgs', 'GlobalSettings']

@pulumi.input_type
class GlobalSettingsArgs:
    def __init__(__self__, *,
                 global_settings: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        """
        The set of arguments for constructing a GlobalSettings resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] global_settings: A list of resources along with the opt-in preferences for the account.
        """
        pulumi.set(__self__, "global_settings", global_settings)

    @property
    @pulumi.getter(name="globalSettings")
    def global_settings(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        """
        A list of resources along with the opt-in preferences for the account.
        """
        return pulumi.get(self, "global_settings")

    @global_settings.setter
    def global_settings(self, value: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        pulumi.set(self, "global_settings", value)


@pulumi.input_type
class _GlobalSettingsState:
    def __init__(__self__, *,
                 global_settings: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering GlobalSettings resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] global_settings: A list of resources along with the opt-in preferences for the account.
        """
        if global_settings is not None:
            pulumi.set(__self__, "global_settings", global_settings)

    @property
    @pulumi.getter(name="globalSettings")
    def global_settings(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A list of resources along with the opt-in preferences for the account.
        """
        return pulumi.get(self, "global_settings")

    @global_settings.setter
    def global_settings(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "global_settings", value)


class GlobalSettings(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 global_settings: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an AWS Backup Global Settings resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.backup.GlobalSettings("test", global_settings={
            "isCrossAccountBackupEnabled": "true",
        })
        ```

        ## Import

        Backup Global Settings can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:backup/globalSettings:GlobalSettings example 123456789012
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] global_settings: A list of resources along with the opt-in preferences for the account.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GlobalSettingsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AWS Backup Global Settings resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.backup.GlobalSettings("test", global_settings={
            "isCrossAccountBackupEnabled": "true",
        })
        ```

        ## Import

        Backup Global Settings can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:backup/globalSettings:GlobalSettings example 123456789012
        ```

        :param str resource_name: The name of the resource.
        :param GlobalSettingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GlobalSettingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 global_settings: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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
            __props__ = GlobalSettingsArgs.__new__(GlobalSettingsArgs)

            if global_settings is None and not opts.urn:
                raise TypeError("Missing required property 'global_settings'")
            __props__.__dict__["global_settings"] = global_settings
        super(GlobalSettings, __self__).__init__(
            'aws:backup/globalSettings:GlobalSettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            global_settings: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'GlobalSettings':
        """
        Get an existing GlobalSettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] global_settings: A list of resources along with the opt-in preferences for the account.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GlobalSettingsState.__new__(_GlobalSettingsState)

        __props__.__dict__["global_settings"] = global_settings
        return GlobalSettings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="globalSettings")
    def global_settings(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A list of resources along with the opt-in preferences for the account.
        """
        return pulumi.get(self, "global_settings")

