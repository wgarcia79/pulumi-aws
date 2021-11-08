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
    'PipelineArtifactStore',
    'PipelineArtifactStoreEncryptionKey',
    'PipelineStage',
    'PipelineStageAction',
    'WebhookAuthenticationConfiguration',
    'WebhookFilter',
]

@pulumi.output_type
class PipelineArtifactStore(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "encryptionKey":
            suggest = "encryption_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PipelineArtifactStore. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PipelineArtifactStore.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PipelineArtifactStore.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 location: str,
                 type: str,
                 encryption_key: Optional['outputs.PipelineArtifactStoreEncryptionKey'] = None,
                 region: Optional[str] = None):
        """
        :param str location: The location where AWS CodePipeline stores artifacts for a pipeline; currently only `S3` is supported.
        :param str type: The type of the artifact store, such as Amazon S3
        :param 'PipelineArtifactStoreEncryptionKeyArgs' encryption_key: The encryption key block AWS CodePipeline uses to encrypt the data in the artifact store, such as an AWS Key Management Service (AWS KMS) key. If you don't specify a key, AWS CodePipeline uses the default key for Amazon Simple Storage Service (Amazon S3). An `encryption_key` block is documented below.
        :param str region: The region where the artifact store is located. Required for a cross-region CodePipeline, do not provide for a single-region CodePipeline.
        """
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "type", type)
        if encryption_key is not None:
            pulumi.set(__self__, "encryption_key", encryption_key)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location where AWS CodePipeline stores artifacts for a pipeline; currently only `S3` is supported.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the artifact store, such as Amazon S3
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="encryptionKey")
    def encryption_key(self) -> Optional['outputs.PipelineArtifactStoreEncryptionKey']:
        """
        The encryption key block AWS CodePipeline uses to encrypt the data in the artifact store, such as an AWS Key Management Service (AWS KMS) key. If you don't specify a key, AWS CodePipeline uses the default key for Amazon Simple Storage Service (Amazon S3). An `encryption_key` block is documented below.
        """
        return pulumi.get(self, "encryption_key")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        The region where the artifact store is located. Required for a cross-region CodePipeline, do not provide for a single-region CodePipeline.
        """
        return pulumi.get(self, "region")


@pulumi.output_type
class PipelineArtifactStoreEncryptionKey(dict):
    def __init__(__self__, *,
                 id: str,
                 type: str):
        """
        :param str id: The KMS key ARN or ID
        :param str type: The type of key; currently only `KMS` is supported
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The KMS key ARN or ID
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of key; currently only `KMS` is supported
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class PipelineStage(dict):
    def __init__(__self__, *,
                 actions: Sequence['outputs.PipelineStageAction'],
                 name: str):
        """
        :param Sequence['PipelineStageActionArgs'] actions: The action(s) to include in the stage. Defined as an `action` block below
        :param str name: The name of the stage.
        """
        pulumi.set(__self__, "actions", actions)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def actions(self) -> Sequence['outputs.PipelineStageAction']:
        """
        The action(s) to include in the stage. Defined as an `action` block below
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the stage.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PipelineStageAction(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "inputArtifacts":
            suggest = "input_artifacts"
        elif key == "outputArtifacts":
            suggest = "output_artifacts"
        elif key == "roleArn":
            suggest = "role_arn"
        elif key == "runOrder":
            suggest = "run_order"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PipelineStageAction. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PipelineStageAction.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PipelineStageAction.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 category: str,
                 name: str,
                 owner: str,
                 provider: str,
                 version: str,
                 configuration: Optional[Mapping[str, str]] = None,
                 input_artifacts: Optional[Sequence[str]] = None,
                 namespace: Optional[str] = None,
                 output_artifacts: Optional[Sequence[str]] = None,
                 region: Optional[str] = None,
                 role_arn: Optional[str] = None,
                 run_order: Optional[int] = None):
        """
        :param str category: A category defines what kind of action can be taken in the stage, and constrains the provider type for the action. Possible values are `Approval`, `Build`, `Deploy`, `Invoke`, `Source` and `Test`.
        :param str name: The action declaration's name.
        :param str owner: The creator of the action being called. Possible values are `AWS`, `Custom` and `ThirdParty`.
        :param str provider: The provider of the service being called by the action. Valid providers are determined by the action category. Provider names are listed in the [Action Structure Reference](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference.html) documentation.
        :param str version: A string that identifies the action type.
        :param Mapping[str, str] configuration: A map of the action declaration's configuration. Configurations options for action types and providers can be found in the [Pipeline Structure Reference](http://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements) and [Action Structure Reference](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference.html) documentation.
        :param Sequence[str] input_artifacts: A list of artifact names to be worked on.
        :param str namespace: The namespace all output variables will be accessed from.
        :param Sequence[str] output_artifacts: A list of artifact names to output. Output artifact names must be unique within a pipeline.
        :param str region: The region in which to run the action.
        :param str role_arn: The ARN of the IAM service role that will perform the declared action. This is assumed through the roleArn for the pipeline.
        :param int run_order: The order in which actions are run.
        """
        pulumi.set(__self__, "category", category)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "owner", owner)
        pulumi.set(__self__, "provider", provider)
        pulumi.set(__self__, "version", version)
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if input_artifacts is not None:
            pulumi.set(__self__, "input_artifacts", input_artifacts)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if output_artifacts is not None:
            pulumi.set(__self__, "output_artifacts", output_artifacts)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if run_order is not None:
            pulumi.set(__self__, "run_order", run_order)

    @property
    @pulumi.getter
    def category(self) -> str:
        """
        A category defines what kind of action can be taken in the stage, and constrains the provider type for the action. Possible values are `Approval`, `Build`, `Deploy`, `Invoke`, `Source` and `Test`.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The action declaration's name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        The creator of the action being called. Possible values are `AWS`, `Custom` and `ThirdParty`.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def provider(self) -> str:
        """
        The provider of the service being called by the action. Valid providers are determined by the action category. Provider names are listed in the [Action Structure Reference](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference.html) documentation.
        """
        return pulumi.get(self, "provider")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        A string that identifies the action type.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter
    def configuration(self) -> Optional[Mapping[str, str]]:
        """
        A map of the action declaration's configuration. Configurations options for action types and providers can be found in the [Pipeline Structure Reference](http://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements) and [Action Structure Reference](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference.html) documentation.
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter(name="inputArtifacts")
    def input_artifacts(self) -> Optional[Sequence[str]]:
        """
        A list of artifact names to be worked on.
        """
        return pulumi.get(self, "input_artifacts")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        """
        The namespace all output variables will be accessed from.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="outputArtifacts")
    def output_artifacts(self) -> Optional[Sequence[str]]:
        """
        A list of artifact names to output. Output artifact names must be unique within a pipeline.
        """
        return pulumi.get(self, "output_artifacts")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        The region in which to run the action.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        """
        The ARN of the IAM service role that will perform the declared action. This is assumed through the roleArn for the pipeline.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="runOrder")
    def run_order(self) -> Optional[int]:
        """
        The order in which actions are run.
        """
        return pulumi.get(self, "run_order")


@pulumi.output_type
class WebhookAuthenticationConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowedIpRange":
            suggest = "allowed_ip_range"
        elif key == "secretToken":
            suggest = "secret_token"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in WebhookAuthenticationConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        WebhookAuthenticationConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        WebhookAuthenticationConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allowed_ip_range: Optional[str] = None,
                 secret_token: Optional[str] = None):
        """
        :param str allowed_ip_range: A valid CIDR block for `IP` filtering. Required for `IP`.
        :param str secret_token: The shared secret for the GitHub repository webhook. Set this as `secret` in your `github_repository_webhook`'s `configuration` block. Required for `GITHUB_HMAC`.
        """
        if allowed_ip_range is not None:
            pulumi.set(__self__, "allowed_ip_range", allowed_ip_range)
        if secret_token is not None:
            pulumi.set(__self__, "secret_token", secret_token)

    @property
    @pulumi.getter(name="allowedIpRange")
    def allowed_ip_range(self) -> Optional[str]:
        """
        A valid CIDR block for `IP` filtering. Required for `IP`.
        """
        return pulumi.get(self, "allowed_ip_range")

    @property
    @pulumi.getter(name="secretToken")
    def secret_token(self) -> Optional[str]:
        """
        The shared secret for the GitHub repository webhook. Set this as `secret` in your `github_repository_webhook`'s `configuration` block. Required for `GITHUB_HMAC`.
        """
        return pulumi.get(self, "secret_token")


@pulumi.output_type
class WebhookFilter(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "jsonPath":
            suggest = "json_path"
        elif key == "matchEquals":
            suggest = "match_equals"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in WebhookFilter. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        WebhookFilter.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        WebhookFilter.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 json_path: str,
                 match_equals: str):
        """
        :param str json_path: The [JSON path](https://github.com/json-path/JsonPath) to filter on.
        :param str match_equals: The value to match on (e.g., `refs/heads/{Branch}`). See [AWS docs](https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_WebhookFilterRule.html) for details.
        """
        pulumi.set(__self__, "json_path", json_path)
        pulumi.set(__self__, "match_equals", match_equals)

    @property
    @pulumi.getter(name="jsonPath")
    def json_path(self) -> str:
        """
        The [JSON path](https://github.com/json-path/JsonPath) to filter on.
        """
        return pulumi.get(self, "json_path")

    @property
    @pulumi.getter(name="matchEquals")
    def match_equals(self) -> str:
        """
        The value to match on (e.g., `refs/heads/{Branch}`). See [AWS docs](https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_WebhookFilterRule.html) for details.
        """
        return pulumi.get(self, "match_equals")


