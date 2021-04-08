# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LogDestinationPolicyArgs', 'LogDestinationPolicy']

@pulumi.input_type
class LogDestinationPolicyArgs:
    def __init__(__self__, *,
                 access_policy: pulumi.Input[str],
                 destination_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a LogDestinationPolicy resource.
        :param pulumi.Input[str] access_policy: The policy document. This is a JSON formatted string.
        :param pulumi.Input[str] destination_name: A name for the subscription filter
        """
        pulumi.set(__self__, "access_policy", access_policy)
        pulumi.set(__self__, "destination_name", destination_name)

    @property
    @pulumi.getter(name="accessPolicy")
    def access_policy(self) -> pulumi.Input[str]:
        """
        The policy document. This is a JSON formatted string.
        """
        return pulumi.get(self, "access_policy")

    @access_policy.setter
    def access_policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_policy", value)

    @property
    @pulumi.getter(name="destinationName")
    def destination_name(self) -> pulumi.Input[str]:
        """
        A name for the subscription filter
        """
        return pulumi.get(self, "destination_name")

    @destination_name.setter
    def destination_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_name", value)


@pulumi.input_type
class _LogDestinationPolicyState:
    def __init__(__self__, *,
                 access_policy: Optional[pulumi.Input[str]] = None,
                 destination_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LogDestinationPolicy resources.
        :param pulumi.Input[str] access_policy: The policy document. This is a JSON formatted string.
        :param pulumi.Input[str] destination_name: A name for the subscription filter
        """
        if access_policy is not None:
            pulumi.set(__self__, "access_policy", access_policy)
        if destination_name is not None:
            pulumi.set(__self__, "destination_name", destination_name)

    @property
    @pulumi.getter(name="accessPolicy")
    def access_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The policy document. This is a JSON formatted string.
        """
        return pulumi.get(self, "access_policy")

    @access_policy.setter
    def access_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_policy", value)

    @property
    @pulumi.getter(name="destinationName")
    def destination_name(self) -> Optional[pulumi.Input[str]]:
        """
        A name for the subscription filter
        """
        return pulumi.get(self, "destination_name")

    @destination_name.setter
    def destination_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_name", value)


class LogDestinationPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_policy: Optional[pulumi.Input[str]] = None,
                 destination_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a CloudWatch Logs destination policy resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test_destination = aws.cloudwatch.LogDestination("testDestination",
            role_arn=aws_iam_role["iam_for_cloudwatch"]["arn"],
            target_arn=aws_kinesis_stream["kinesis_for_cloudwatch"]["arn"])
        test_destination_policy_policy_document = test_destination.arn.apply(lambda arn: aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            effect="Allow",
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                type="AWS",
                identifiers=["123456789012"],
            )],
            actions=["logs:PutSubscriptionFilter"],
            resources=[arn],
        )]))
        test_destination_policy_log_destination_policy = aws.cloudwatch.LogDestinationPolicy("testDestinationPolicyLogDestinationPolicy",
            destination_name=test_destination.name,
            access_policy=test_destination_policy_policy_document.json)
        ```

        ## Import

        CloudWatch Logs destination policies can be imported using the `destination_name`, e.g.

        ```sh
         $ pulumi import aws:cloudwatch/logDestinationPolicy:LogDestinationPolicy test_destination_policy test_destination
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_policy: The policy document. This is a JSON formatted string.
        :param pulumi.Input[str] destination_name: A name for the subscription filter
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LogDestinationPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a CloudWatch Logs destination policy resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test_destination = aws.cloudwatch.LogDestination("testDestination",
            role_arn=aws_iam_role["iam_for_cloudwatch"]["arn"],
            target_arn=aws_kinesis_stream["kinesis_for_cloudwatch"]["arn"])
        test_destination_policy_policy_document = test_destination.arn.apply(lambda arn: aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            effect="Allow",
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                type="AWS",
                identifiers=["123456789012"],
            )],
            actions=["logs:PutSubscriptionFilter"],
            resources=[arn],
        )]))
        test_destination_policy_log_destination_policy = aws.cloudwatch.LogDestinationPolicy("testDestinationPolicyLogDestinationPolicy",
            destination_name=test_destination.name,
            access_policy=test_destination_policy_policy_document.json)
        ```

        ## Import

        CloudWatch Logs destination policies can be imported using the `destination_name`, e.g.

        ```sh
         $ pulumi import aws:cloudwatch/logDestinationPolicy:LogDestinationPolicy test_destination_policy test_destination
        ```

        :param str resource_name: The name of the resource.
        :param LogDestinationPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogDestinationPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_policy: Optional[pulumi.Input[str]] = None,
                 destination_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = LogDestinationPolicyArgs.__new__(LogDestinationPolicyArgs)

            if access_policy is None and not opts.urn:
                raise TypeError("Missing required property 'access_policy'")
            __props__.__dict__["access_policy"] = access_policy
            if destination_name is None and not opts.urn:
                raise TypeError("Missing required property 'destination_name'")
            __props__.__dict__["destination_name"] = destination_name
        super(LogDestinationPolicy, __self__).__init__(
            'aws:cloudwatch/logDestinationPolicy:LogDestinationPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_policy: Optional[pulumi.Input[str]] = None,
            destination_name: Optional[pulumi.Input[str]] = None) -> 'LogDestinationPolicy':
        """
        Get an existing LogDestinationPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_policy: The policy document. This is a JSON formatted string.
        :param pulumi.Input[str] destination_name: A name for the subscription filter
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogDestinationPolicyState.__new__(_LogDestinationPolicyState)

        __props__.__dict__["access_policy"] = access_policy
        __props__.__dict__["destination_name"] = destination_name
        return LogDestinationPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessPolicy")
    def access_policy(self) -> pulumi.Output[str]:
        """
        The policy document. This is a JSON formatted string.
        """
        return pulumi.get(self, "access_policy")

    @property
    @pulumi.getter(name="destinationName")
    def destination_name(self) -> pulumi.Output[str]:
        """
        A name for the subscription filter
        """
        return pulumi.get(self, "destination_name")

