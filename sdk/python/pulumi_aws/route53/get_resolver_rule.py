# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetResolverRuleResult',
    'AwaitableGetResolverRuleResult',
    'get_resolver_rule',
]



@pulumi.output_type
class GetResolverRuleResult:
    """
    A collection of values returned by getResolverRule.
    """
    def __init__(__self__, arn=None, domain_name=None, id=None, name=None, owner_id=None, resolver_endpoint_id=None, resolver_rule_id=None, rule_type=None, share_status=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if domain_name and not isinstance(domain_name, str):
            raise TypeError("Expected argument 'domain_name' to be a str")
        pulumi.set(__self__, "domain_name", domain_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        pulumi.set(__self__, "owner_id", owner_id)
        if resolver_endpoint_id and not isinstance(resolver_endpoint_id, str):
            raise TypeError("Expected argument 'resolver_endpoint_id' to be a str")
        pulumi.set(__self__, "resolver_endpoint_id", resolver_endpoint_id)
        if resolver_rule_id and not isinstance(resolver_rule_id, str):
            raise TypeError("Expected argument 'resolver_rule_id' to be a str")
        pulumi.set(__self__, "resolver_rule_id", resolver_rule_id)
        if rule_type and not isinstance(rule_type, str):
            raise TypeError("Expected argument 'rule_type' to be a str")
        pulumi.set(__self__, "rule_type", rule_type)
        if share_status and not isinstance(share_status, str):
            raise TypeError("Expected argument 'share_status' to be a str")
        pulumi.set(__self__, "share_status", share_status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The ARN (Amazon Resource Name) for the resolver rule.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> str:
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> str:
        """
        When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="resolverEndpointId")
    def resolver_endpoint_id(self) -> str:
        return pulumi.get(self, "resolver_endpoint_id")

    @property
    @pulumi.getter(name="resolverRuleId")
    def resolver_rule_id(self) -> str:
        return pulumi.get(self, "resolver_rule_id")

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> str:
        return pulumi.get(self, "rule_type")

    @property
    @pulumi.getter(name="shareStatus")
    def share_status(self) -> str:
        """
        Whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account.
        Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
        """
        return pulumi.get(self, "share_status")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A map of tags assigned to the resolver rule.
        """
        return pulumi.get(self, "tags")



class AwaitableGetResolverRuleResult(GetResolverRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResolverRuleResult(
            arn=self.arn,
            domain_name=self.domain_name,
            id=self.id,
            name=self.name,
            owner_id=self.owner_id,
            resolver_endpoint_id=self.resolver_endpoint_id,
            resolver_rule_id=self.resolver_rule_id,
            rule_type=self.rule_type,
            share_status=self.share_status,
            tags=self.tags)


def get_resolver_rule(domain_name: Optional[str] = None,
                      name: Optional[str] = None,
                      resolver_endpoint_id: Optional[str] = None,
                      resolver_rule_id: Optional[str] = None,
                      rule_type: Optional[str] = None,
                      tags: Optional[Mapping[str, str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResolverRuleResult:
    """
    `route53.ResolverRule` provides details about a specific Route53 Resolver rule.

    ## Example Usage

    The following example shows how to get a Route53 Resolver rule based on its associated domain name and rule type.

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.route53.get_resolver_rule(domain_name="subdomain.example.com",
        rule_type="SYSTEM")
    ```


    :param str domain_name: The domain name the desired resolver rule forwards DNS queries for. Conflicts with `resolver_rule_id`.
    :param str name: The friendly name of the desired resolver rule. Conflicts with `resolver_rule_id`.
    :param str resolver_endpoint_id: The ID of the outbound resolver endpoint of the desired resolver rule. Conflicts with `resolver_rule_id`.
    :param str resolver_rule_id: The ID of the desired resolver rule. Conflicts with `domain_name`, `name`, `resolver_endpoint_id` and `rule_type`.
    :param str rule_type: The rule type of the desired resolver rule. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`. Conflicts with `resolver_rule_id`.
    :param Mapping[str, str] tags: A map of tags assigned to the resolver rule.
    """
    __args__ = dict()
    __args__['domainName'] = domain_name
    __args__['name'] = name
    __args__['resolverEndpointId'] = resolver_endpoint_id
    __args__['resolverRuleId'] = resolver_rule_id
    __args__['ruleType'] = rule_type
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:route53/getResolverRule:getResolverRule', __args__, opts=opts, typ=GetResolverRuleResult).value

    return AwaitableGetResolverRuleResult(
        arn=__ret__.arn,
        domain_name=__ret__.domain_name,
        id=__ret__.id,
        name=__ret__.name,
        owner_id=__ret__.owner_id,
        resolver_endpoint_id=__ret__.resolver_endpoint_id,
        resolver_rule_id=__ret__.resolver_rule_id,
        rule_type=__ret__.rule_type,
        share_status=__ret__.share_status,
        tags=__ret__.tags)
