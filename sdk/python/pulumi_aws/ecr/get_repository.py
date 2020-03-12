# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetRepositoryResult:
    """
    A collection of values returned by getRepository.
    """
    def __init__(__self__, arn=None, id=None, name=None, registry_id=None, repository_url=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        Full ARN of the repository.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if registry_id and not isinstance(registry_id, str):
            raise TypeError("Expected argument 'registry_id' to be a str")
        __self__.registry_id = registry_id
        """
        The registry ID where the repository was created.
        """
        if repository_url and not isinstance(repository_url, str):
            raise TypeError("Expected argument 'repository_url' to be a str")
        __self__.repository_url = repository_url
        """
        The URL of the repository (in the form `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName`).
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags assigned to the resource.
        """
class AwaitableGetRepositoryResult(GetRepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryResult(
            arn=self.arn,
            id=self.id,
            name=self.name,
            registry_id=self.registry_id,
            repository_url=self.repository_url,
            tags=self.tags)

def get_repository(name=None,tags=None,opts=None):
    """
    The ECR Repository data source allows the ARN, Repository URI and Registry ID to be retrieved for an ECR repository.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ecr_repository.html.markdown.


    :param str name: The name of the ECR Repository.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ecr/getRepository:getRepository', __args__, opts=opts).value

    return AwaitableGetRepositoryResult(
        arn=__ret__.get('arn'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        registry_id=__ret__.get('registryId'),
        repository_url=__ret__.get('repositoryUrl'),
        tags=__ret__.get('tags'))
