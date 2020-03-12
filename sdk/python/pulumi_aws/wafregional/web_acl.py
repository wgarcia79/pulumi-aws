# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class WebAcl(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the WAF Regional WebACL.
    """
    default_action: pulumi.Output[dict]
    """
    The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.

      * `type` (`str`) - Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
    """
    logging_configuration: pulumi.Output[dict]
    """
    Configuration block to enable WAF logging. Detailed below.

      * `log_destination` (`str`) - Amazon Resource Name (ARN) of Kinesis Firehose Delivery Stream
      * `redactedFields` (`dict`) - Configuration block containing parts of the request that you want redacted from the logs. Detailed below.
        * `fieldToMatches` (`list`) - Set of configuration blocks for fields to redact. Detailed below.
          * `data` (`str`) - When the value of `type` is `HEADER`, enter the name of the header that you want the WAF to search, for example, `User-Agent` or `Referer`. If the value of `type` is any other value, omit `data`.
          * `type` (`str`) - Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
    """
    metric_name: pulumi.Output[str]
    """
    The name or description for the Amazon CloudWatch metric of this web ACL.
    """
    name: pulumi.Output[str]
    """
    The name or description of the web ACL.
    """
    rules: pulumi.Output[list]
    """
    Set of configuration blocks containing rules for the web ACL. Detailed below.

      * `action` (`dict`) - Configuration block of the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Not used if `type` is `GROUP`. Detailed below.
        * `type` (`str`) - Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`

      * `overrideAction` (`dict`) - Configuration block of the override the action that a group requests CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Only used if `type` is `GROUP`. Detailed below.
        * `type` (`str`) - Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`

      * `priority` (`float`) - Specifies the order in which the rules in a WebACL are evaluated.
        Rules with a lower value are evaluated before rules with a higher value.
      * `rule_id` (`str`) - ID of the associated WAF (Regional) rule (e.g. [`wafregional.Rule`](https://www.terraform.io/docs/providers/aws/r/wafregional_rule.html)). WAF (Global) rules cannot be used.
      * `type` (`str`) - Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
    """
    tags: pulumi.Output[dict]
    """
    Key-value mapping of resource tags
    """
    def __init__(__self__, resource_name, opts=None, default_action=None, logging_configuration=None, metric_name=None, name=None, rules=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a WAF Regional Web ACL Resource for use with Application Load Balancer.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/wafregional_web_acl.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] default_action: The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
        :param pulumi.Input[dict] logging_configuration: Configuration block to enable WAF logging. Detailed below.
        :param pulumi.Input[str] metric_name: The name or description for the Amazon CloudWatch metric of this web ACL.
        :param pulumi.Input[str] name: The name or description of the web ACL.
        :param pulumi.Input[list] rules: Set of configuration blocks containing rules for the web ACL. Detailed below.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags

        The **default_action** object supports the following:

          * `type` (`pulumi.Input[str]`) - Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`

        The **logging_configuration** object supports the following:

          * `log_destination` (`pulumi.Input[str]`) - Amazon Resource Name (ARN) of Kinesis Firehose Delivery Stream
          * `redactedFields` (`pulumi.Input[dict]`) - Configuration block containing parts of the request that you want redacted from the logs. Detailed below.
            * `fieldToMatches` (`pulumi.Input[list]`) - Set of configuration blocks for fields to redact. Detailed below.
              * `data` (`pulumi.Input[str]`) - When the value of `type` is `HEADER`, enter the name of the header that you want the WAF to search, for example, `User-Agent` or `Referer`. If the value of `type` is any other value, omit `data`.
              * `type` (`pulumi.Input[str]`) - Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`

        The **rules** object supports the following:

          * `action` (`pulumi.Input[dict]`) - Configuration block of the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Not used if `type` is `GROUP`. Detailed below.
            * `type` (`pulumi.Input[str]`) - Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`

          * `overrideAction` (`pulumi.Input[dict]`) - Configuration block of the override the action that a group requests CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Only used if `type` is `GROUP`. Detailed below.
            * `type` (`pulumi.Input[str]`) - Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`

          * `priority` (`pulumi.Input[float]`) - Specifies the order in which the rules in a WebACL are evaluated.
            Rules with a lower value are evaluated before rules with a higher value.
          * `rule_id` (`pulumi.Input[str]`) - ID of the associated WAF (Regional) rule (e.g. [`wafregional.Rule`](https://www.terraform.io/docs/providers/aws/r/wafregional_rule.html)). WAF (Global) rules cannot be used.
          * `type` (`pulumi.Input[str]`) - Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
        """
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if default_action is None:
                raise TypeError("Missing required property 'default_action'")
            __props__['default_action'] = default_action
            __props__['logging_configuration'] = logging_configuration
            if metric_name is None:
                raise TypeError("Missing required property 'metric_name'")
            __props__['metric_name'] = metric_name
            __props__['name'] = name
            __props__['rules'] = rules
            __props__['tags'] = tags
            __props__['arn'] = None
        super(WebAcl, __self__).__init__(
            'aws:wafregional/webAcl:WebAcl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, default_action=None, logging_configuration=None, metric_name=None, name=None, rules=None, tags=None):
        """
        Get an existing WebAcl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the WAF Regional WebACL.
        :param pulumi.Input[dict] default_action: The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
        :param pulumi.Input[dict] logging_configuration: Configuration block to enable WAF logging. Detailed below.
        :param pulumi.Input[str] metric_name: The name or description for the Amazon CloudWatch metric of this web ACL.
        :param pulumi.Input[str] name: The name or description of the web ACL.
        :param pulumi.Input[list] rules: Set of configuration blocks containing rules for the web ACL. Detailed below.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags

        The **default_action** object supports the following:

          * `type` (`pulumi.Input[str]`) - Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`

        The **logging_configuration** object supports the following:

          * `log_destination` (`pulumi.Input[str]`) - Amazon Resource Name (ARN) of Kinesis Firehose Delivery Stream
          * `redactedFields` (`pulumi.Input[dict]`) - Configuration block containing parts of the request that you want redacted from the logs. Detailed below.
            * `fieldToMatches` (`pulumi.Input[list]`) - Set of configuration blocks for fields to redact. Detailed below.
              * `data` (`pulumi.Input[str]`) - When the value of `type` is `HEADER`, enter the name of the header that you want the WAF to search, for example, `User-Agent` or `Referer`. If the value of `type` is any other value, omit `data`.
              * `type` (`pulumi.Input[str]`) - Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`

        The **rules** object supports the following:

          * `action` (`pulumi.Input[dict]`) - Configuration block of the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Not used if `type` is `GROUP`. Detailed below.
            * `type` (`pulumi.Input[str]`) - Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`

          * `overrideAction` (`pulumi.Input[dict]`) - Configuration block of the override the action that a group requests CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Only used if `type` is `GROUP`. Detailed below.
            * `type` (`pulumi.Input[str]`) - Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`

          * `priority` (`pulumi.Input[float]`) - Specifies the order in which the rules in a WebACL are evaluated.
            Rules with a lower value are evaluated before rules with a higher value.
          * `rule_id` (`pulumi.Input[str]`) - ID of the associated WAF (Regional) rule (e.g. [`wafregional.Rule`](https://www.terraform.io/docs/providers/aws/r/wafregional_rule.html)). WAF (Global) rules cannot be used.
          * `type` (`pulumi.Input[str]`) - Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["default_action"] = default_action
        __props__["logging_configuration"] = logging_configuration
        __props__["metric_name"] = metric_name
        __props__["name"] = name
        __props__["rules"] = rules
        __props__["tags"] = tags
        return WebAcl(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

