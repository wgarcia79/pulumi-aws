# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetRulesPackagesResult',
    'AwaitableGetRulesPackagesResult',
    'get_rules_packages',
]

@pulumi.output_type
class GetRulesPackagesResult:
    """
    A collection of values returned by getRulesPackages.
    """
    def __init__(__self__, arns=None, id=None):
        if arns and not isinstance(arns, list):
            raise TypeError("Expected argument 'arns' to be a list")
        pulumi.set(__self__, "arns", arns)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def arns(self) -> Sequence[str]:
        """
        A list of the AWS Inspector Rules Packages arns available in the AWS region.
        """
        return pulumi.get(self, "arns")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetRulesPackagesResult(GetRulesPackagesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRulesPackagesResult(
            arns=self.arns,
            id=self.id)


def get_rules_packages(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRulesPackagesResult:
    """
    The AWS Inspector Rules Packages data source allows access to the list of AWS
    Inspector Rules Packages which can be used by AWS Inspector within the region
    configured in the provider.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    rules = aws.inspector.get_rules_packages()
    # e.g., Use in aws_inspector_assessment_template
    group = aws.inspector.ResourceGroup("group", tags={
        "test": "test",
    })
    assessment_assessment_target = aws.inspector.AssessmentTarget("assessmentAssessmentTarget", resource_group_arn=group.arn)
    assessment_assessment_template = aws.inspector.AssessmentTemplate("assessmentAssessmentTemplate",
        target_arn=assessment_assessment_target.arn,
        duration=60,
        rules_package_arns=rules.arns)
    ```
    """
    __args__ = dict()
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:inspector/getRulesPackages:getRulesPackages', __args__, opts=opts, typ=GetRulesPackagesResult).value

    return AwaitableGetRulesPackagesResult(
        arns=__ret__.arns,
        id=__ret__.id)
