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

__all__ = ['RemediationConfigurationArgs', 'RemediationConfiguration']

@pulumi.input_type
class RemediationConfigurationArgs:
    def __init__(__self__, *,
                 config_rule_name: pulumi.Input[str],
                 target_id: pulumi.Input[str],
                 target_type: pulumi.Input[str],
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input['RemediationConfigurationParameterArgs']]]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 target_version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RemediationConfiguration resource.
        :param pulumi.Input[str] config_rule_name: The name of the AWS Config rule
        :param pulumi.Input[str] target_id: Target ID is the name of the public document
        :param pulumi.Input[str] target_type: The type of the target. Target executes remediation. For example, SSM document
        :param pulumi.Input[Sequence[pulumi.Input['RemediationConfigurationParameterArgs']]] parameters: Can be specified multiple times for each
               parameter. Each parameter block supports fields documented below.
        :param pulumi.Input[str] resource_type: The type of a resource
        :param pulumi.Input[str] target_version: Version of the target. For example, version of the SSM document
        """
        pulumi.set(__self__, "config_rule_name", config_rule_name)
        pulumi.set(__self__, "target_id", target_id)
        pulumi.set(__self__, "target_type", target_type)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if target_version is not None:
            pulumi.set(__self__, "target_version", target_version)

    @property
    @pulumi.getter(name="configRuleName")
    def config_rule_name(self) -> pulumi.Input[str]:
        """
        The name of the AWS Config rule
        """
        return pulumi.get(self, "config_rule_name")

    @config_rule_name.setter
    def config_rule_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "config_rule_name", value)

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> pulumi.Input[str]:
        """
        Target ID is the name of the public document
        """
        return pulumi.get(self, "target_id")

    @target_id.setter
    def target_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_id", value)

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> pulumi.Input[str]:
        """
        The type of the target. Target executes remediation. For example, SSM document
        """
        return pulumi.get(self, "target_type")

    @target_type.setter
    def target_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_type", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RemediationConfigurationParameterArgs']]]]:
        """
        Can be specified multiple times for each
        parameter. Each parameter block supports fields documented below.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RemediationConfigurationParameterArgs']]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of a resource
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter(name="targetVersion")
    def target_version(self) -> Optional[pulumi.Input[str]]:
        """
        Version of the target. For example, version of the SSM document
        """
        return pulumi.get(self, "target_version")

    @target_version.setter
    def target_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_version", value)


@pulumi.input_type
class _RemediationConfigurationState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 config_rule_name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input['RemediationConfigurationParameterArgs']]]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 target_id: Optional[pulumi.Input[str]] = None,
                 target_type: Optional[pulumi.Input[str]] = None,
                 target_version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RemediationConfiguration resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Config Remediation Configuration.
        :param pulumi.Input[str] config_rule_name: The name of the AWS Config rule
        :param pulumi.Input[Sequence[pulumi.Input['RemediationConfigurationParameterArgs']]] parameters: Can be specified multiple times for each
               parameter. Each parameter block supports fields documented below.
        :param pulumi.Input[str] resource_type: The type of a resource
        :param pulumi.Input[str] target_id: Target ID is the name of the public document
        :param pulumi.Input[str] target_type: The type of the target. Target executes remediation. For example, SSM document
        :param pulumi.Input[str] target_version: Version of the target. For example, version of the SSM document
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if config_rule_name is not None:
            pulumi.set(__self__, "config_rule_name", config_rule_name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if target_id is not None:
            pulumi.set(__self__, "target_id", target_id)
        if target_type is not None:
            pulumi.set(__self__, "target_type", target_type)
        if target_version is not None:
            pulumi.set(__self__, "target_version", target_version)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the Config Remediation Configuration.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="configRuleName")
    def config_rule_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the AWS Config rule
        """
        return pulumi.get(self, "config_rule_name")

    @config_rule_name.setter
    def config_rule_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_rule_name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RemediationConfigurationParameterArgs']]]]:
        """
        Can be specified multiple times for each
        parameter. Each parameter block supports fields documented below.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RemediationConfigurationParameterArgs']]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of a resource
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> Optional[pulumi.Input[str]]:
        """
        Target ID is the name of the public document
        """
        return pulumi.get(self, "target_id")

    @target_id.setter
    def target_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_id", value)

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the target. Target executes remediation. For example, SSM document
        """
        return pulumi.get(self, "target_type")

    @target_type.setter
    def target_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_type", value)

    @property
    @pulumi.getter(name="targetVersion")
    def target_version(self) -> Optional[pulumi.Input[str]]:
        """
        Version of the target. For example, version of the SSM document
        """
        return pulumi.get(self, "target_version")

    @target_version.setter
    def target_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_version", value)


class RemediationConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_rule_name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RemediationConfigurationParameterArgs']]]]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 target_id: Optional[pulumi.Input[str]] = None,
                 target_type: Optional[pulumi.Input[str]] = None,
                 target_version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an AWS Config Remediation Configuration.

        > **Note:** Config Remediation Configuration requires an existing Config Rule to be present.

        ## Example Usage

        AWS managed rules can be used by setting the source owner to `AWS` and the source identifier to the name of the managed rule. More information about AWS managed rules can be found in the [AWS Config Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html).

        ```python
        import pulumi
        import pulumi_aws as aws

        this_rule = aws.cfg.Rule("thisRule", source=aws.cfg.RuleSourceArgs(
            owner="AWS",
            source_identifier="S3_BUCKET_VERSIONING_ENABLED",
        ))
        this_remediation_configuration = aws.cfg.RemediationConfiguration("thisRemediationConfiguration",
            config_rule_name=this_rule.name,
            resource_type="AWS::S3::Bucket",
            target_type="SSM_DOCUMENT",
            target_id="AWS-EnableS3BucketEncryption",
            target_version="1",
            parameters=[
                aws.cfg.RemediationConfigurationParameterArgs(
                    name="AutomationAssumeRole",
                    static_value="arn:aws:iam::875924563244:role/security_config",
                ),
                aws.cfg.RemediationConfigurationParameterArgs(
                    name="BucketName",
                    resource_value="RESOURCE_ID",
                ),
                aws.cfg.RemediationConfigurationParameterArgs(
                    name="SSEAlgorithm",
                    static_value="AES256",
                ),
            ])
        ```

        ## Import

        Remediation Configurations can be imported using the name config_rule_name, e.g.,

        ```sh
         $ pulumi import aws:cfg/remediationConfiguration:RemediationConfiguration this example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_rule_name: The name of the AWS Config rule
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RemediationConfigurationParameterArgs']]]] parameters: Can be specified multiple times for each
               parameter. Each parameter block supports fields documented below.
        :param pulumi.Input[str] resource_type: The type of a resource
        :param pulumi.Input[str] target_id: Target ID is the name of the public document
        :param pulumi.Input[str] target_type: The type of the target. Target executes remediation. For example, SSM document
        :param pulumi.Input[str] target_version: Version of the target. For example, version of the SSM document
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RemediationConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AWS Config Remediation Configuration.

        > **Note:** Config Remediation Configuration requires an existing Config Rule to be present.

        ## Example Usage

        AWS managed rules can be used by setting the source owner to `AWS` and the source identifier to the name of the managed rule. More information about AWS managed rules can be found in the [AWS Config Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html).

        ```python
        import pulumi
        import pulumi_aws as aws

        this_rule = aws.cfg.Rule("thisRule", source=aws.cfg.RuleSourceArgs(
            owner="AWS",
            source_identifier="S3_BUCKET_VERSIONING_ENABLED",
        ))
        this_remediation_configuration = aws.cfg.RemediationConfiguration("thisRemediationConfiguration",
            config_rule_name=this_rule.name,
            resource_type="AWS::S3::Bucket",
            target_type="SSM_DOCUMENT",
            target_id="AWS-EnableS3BucketEncryption",
            target_version="1",
            parameters=[
                aws.cfg.RemediationConfigurationParameterArgs(
                    name="AutomationAssumeRole",
                    static_value="arn:aws:iam::875924563244:role/security_config",
                ),
                aws.cfg.RemediationConfigurationParameterArgs(
                    name="BucketName",
                    resource_value="RESOURCE_ID",
                ),
                aws.cfg.RemediationConfigurationParameterArgs(
                    name="SSEAlgorithm",
                    static_value="AES256",
                ),
            ])
        ```

        ## Import

        Remediation Configurations can be imported using the name config_rule_name, e.g.,

        ```sh
         $ pulumi import aws:cfg/remediationConfiguration:RemediationConfiguration this example
        ```

        :param str resource_name: The name of the resource.
        :param RemediationConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RemediationConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_rule_name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RemediationConfigurationParameterArgs']]]]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 target_id: Optional[pulumi.Input[str]] = None,
                 target_type: Optional[pulumi.Input[str]] = None,
                 target_version: Optional[pulumi.Input[str]] = None,
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
            __props__ = RemediationConfigurationArgs.__new__(RemediationConfigurationArgs)

            if config_rule_name is None and not opts.urn:
                raise TypeError("Missing required property 'config_rule_name'")
            __props__.__dict__["config_rule_name"] = config_rule_name
            __props__.__dict__["parameters"] = parameters
            __props__.__dict__["resource_type"] = resource_type
            if target_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_id'")
            __props__.__dict__["target_id"] = target_id
            if target_type is None and not opts.urn:
                raise TypeError("Missing required property 'target_type'")
            __props__.__dict__["target_type"] = target_type
            __props__.__dict__["target_version"] = target_version
            __props__.__dict__["arn"] = None
        super(RemediationConfiguration, __self__).__init__(
            'aws:cfg/remediationConfiguration:RemediationConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            config_rule_name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RemediationConfigurationParameterArgs']]]]] = None,
            resource_type: Optional[pulumi.Input[str]] = None,
            target_id: Optional[pulumi.Input[str]] = None,
            target_type: Optional[pulumi.Input[str]] = None,
            target_version: Optional[pulumi.Input[str]] = None) -> 'RemediationConfiguration':
        """
        Get an existing RemediationConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Config Remediation Configuration.
        :param pulumi.Input[str] config_rule_name: The name of the AWS Config rule
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RemediationConfigurationParameterArgs']]]] parameters: Can be specified multiple times for each
               parameter. Each parameter block supports fields documented below.
        :param pulumi.Input[str] resource_type: The type of a resource
        :param pulumi.Input[str] target_id: Target ID is the name of the public document
        :param pulumi.Input[str] target_type: The type of the target. Target executes remediation. For example, SSM document
        :param pulumi.Input[str] target_version: Version of the target. For example, version of the SSM document
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RemediationConfigurationState.__new__(_RemediationConfigurationState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["config_rule_name"] = config_rule_name
        __props__.__dict__["parameters"] = parameters
        __props__.__dict__["resource_type"] = resource_type
        __props__.__dict__["target_id"] = target_id
        __props__.__dict__["target_type"] = target_type
        __props__.__dict__["target_version"] = target_version
        return RemediationConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the Config Remediation Configuration.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="configRuleName")
    def config_rule_name(self) -> pulumi.Output[str]:
        """
        The name of the AWS Config rule
        """
        return pulumi.get(self, "config_rule_name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Sequence['outputs.RemediationConfigurationParameter']]]:
        """
        Can be specified multiple times for each
        parameter. Each parameter block supports fields documented below.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of a resource
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> pulumi.Output[str]:
        """
        Target ID is the name of the public document
        """
        return pulumi.get(self, "target_id")

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> pulumi.Output[str]:
        """
        The type of the target. Target executes remediation. For example, SSM document
        """
        return pulumi.get(self, "target_type")

    @property
    @pulumi.getter(name="targetVersion")
    def target_version(self) -> pulumi.Output[Optional[str]]:
        """
        Version of the target. For example, version of the SSM document
        """
        return pulumi.get(self, "target_version")

