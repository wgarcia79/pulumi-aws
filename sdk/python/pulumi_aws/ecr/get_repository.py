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
    'GetRepositoryResult',
    'AwaitableGetRepositoryResult',
    'get_repository',
]

@pulumi.output_type
class GetRepositoryResult:
    """
    A collection of values returned by getRepository.
    """
    def __init__(__self__, arn=None, encryption_configurations=None, id=None, image_scanning_configurations=None, image_tag_mutability=None, name=None, registry_id=None, repository_url=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if encryption_configurations and not isinstance(encryption_configurations, list):
            raise TypeError("Expected argument 'encryption_configurations' to be a list")
        pulumi.set(__self__, "encryption_configurations", encryption_configurations)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_scanning_configurations and not isinstance(image_scanning_configurations, list):
            raise TypeError("Expected argument 'image_scanning_configurations' to be a list")
        pulumi.set(__self__, "image_scanning_configurations", image_scanning_configurations)
        if image_tag_mutability and not isinstance(image_tag_mutability, str):
            raise TypeError("Expected argument 'image_tag_mutability' to be a str")
        pulumi.set(__self__, "image_tag_mutability", image_tag_mutability)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if registry_id and not isinstance(registry_id, str):
            raise TypeError("Expected argument 'registry_id' to be a str")
        pulumi.set(__self__, "registry_id", registry_id)
        if repository_url and not isinstance(repository_url, str):
            raise TypeError("Expected argument 'repository_url' to be a str")
        pulumi.set(__self__, "repository_url", repository_url)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        Full ARN of the repository.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="encryptionConfigurations")
    def encryption_configurations(self) -> Sequence['outputs.GetRepositoryEncryptionConfigurationResult']:
        """
        Encryption configuration for the repository. See Encryption Configuration below.
        """
        return pulumi.get(self, "encryption_configurations")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageScanningConfigurations")
    def image_scanning_configurations(self) -> Sequence['outputs.GetRepositoryImageScanningConfigurationResult']:
        """
        Configuration block that defines image scanning configuration for the repository. See Image Scanning Configuration below.
        """
        return pulumi.get(self, "image_scanning_configurations")

    @property
    @pulumi.getter(name="imageTagMutability")
    def image_tag_mutability(self) -> str:
        """
        The tag mutability setting for the repository.
        """
        return pulumi.get(self, "image_tag_mutability")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> str:
        return pulumi.get(self, "registry_id")

    @property
    @pulumi.getter(name="repositoryUrl")
    def repository_url(self) -> str:
        """
        The URL of the repository (in the form `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName`).
        """
        return pulumi.get(self, "repository_url")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A map of tags assigned to the resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetRepositoryResult(GetRepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryResult(
            arn=self.arn,
            encryption_configurations=self.encryption_configurations,
            id=self.id,
            image_scanning_configurations=self.image_scanning_configurations,
            image_tag_mutability=self.image_tag_mutability,
            name=self.name,
            registry_id=self.registry_id,
            repository_url=self.repository_url,
            tags=self.tags)


def get_repository(name: Optional[str] = None,
                   registry_id: Optional[str] = None,
                   tags: Optional[Mapping[str, str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryResult:
    """
    The ECR Repository data source allows the ARN, Repository URI and Registry ID to be retrieved for an ECR repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    service = aws.ecr.get_repository(name="ecr-repository")
    ```


    :param str name: The name of the ECR Repository.
    :param str registry_id: The registry ID where the repository was created.
    :param Mapping[str, str] tags: A map of tags assigned to the resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['registryId'] = registry_id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ecr/getRepository:getRepository', __args__, opts=opts, typ=GetRepositoryResult).value

    return AwaitableGetRepositoryResult(
        arn=__ret__.arn,
        encryption_configurations=__ret__.encryption_configurations,
        id=__ret__.id,
        image_scanning_configurations=__ret__.image_scanning_configurations,
        image_tag_mutability=__ret__.image_tag_mutability,
        name=__ret__.name,
        registry_id=__ret__.registry_id,
        repository_url=__ret__.repository_url,
        tags=__ret__.tags)
