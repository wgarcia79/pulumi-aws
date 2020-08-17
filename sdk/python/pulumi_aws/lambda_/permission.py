# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Permission']


class Permission(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 event_source_token: Optional[pulumi.Input[str]] = None,
                 function: Optional[pulumi.Input[str]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 qualifier: Optional[pulumi.Input[str]] = None,
                 source_account: Optional[pulumi.Input[str]] = None,
                 source_arn: Optional[pulumi.Input[str]] = None,
                 statement_id: Optional[pulumi.Input[str]] = None,
                 statement_id_prefix: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Gives an external source (like a CloudWatch Event Rule, SNS, or S3) permission to access the Lambda function.

        ## Example Usage
        ### Basic Example

        ```python
        import pulumi
        import pulumi_aws as aws

        iam_for_lambda = aws.iam.Role("iamForLambda", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": "sts:AssumeRole",
              "Principal": {
                "Service": "lambda.amazonaws.com"
              },
              "Effect": "Allow",
              "Sid": ""
            }
          ]
        }

        \"\"\")
        test_lambda = aws.lambda_.Function("testLambda",
            code=pulumi.FileArchive("lambdatest.zip"),
            handler="exports.handler",
            role=iam_for_lambda.arn,
            runtime="nodejs8.10")
        test_alias = aws.lambda_.Alias("testAlias",
            description="a sample description",
            function_name=test_lambda.name,
            function_version="$LATEST")
        allow_cloudwatch = aws.lambda_.Permission("allowCloudwatch",
            action="lambda:InvokeFunction",
            function=test_lambda.name,
            principal="events.amazonaws.com",
            qualifier=test_alias.name,
            source_arn="arn:aws:events:eu-west-1:111122223333:rule/RunDaily")
        ```
        ### Usage with SNS

        ```python
        import pulumi
        import pulumi_aws as aws

        default_topic = aws.sns.Topic("defaultTopic")
        default_role = aws.iam.Role("defaultRole", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": "sts:AssumeRole",
              "Principal": {
                "Service": "lambda.amazonaws.com"
              },
              "Effect": "Allow",
              "Sid": ""
            }
          ]
        }

        \"\"\")
        func = aws.lambda_.Function("func",
            code=pulumi.FileArchive("lambdatest.zip"),
            handler="exports.handler",
            role=default_role.arn,
            runtime="python2.7")
        with_sns = aws.lambda_.Permission("withSns",
            action="lambda:InvokeFunction",
            function=func.name,
            principal="sns.amazonaws.com",
            source_arn=default_topic.arn)
        lambda_ = aws.sns.TopicSubscription("lambda",
            endpoint=func.arn,
            protocol="lambda",
            topic=default_topic.arn)
        ```
        ### Specify Lambda permissions for API Gateway REST API

        ```python
        import pulumi
        import pulumi_aws as aws

        my_demo_api = aws.apigateway.RestApi("myDemoAPI", description="This is my API for demonstration purposes")
        lambda_permission = aws.lambda_.Permission("lambdaPermission",
            action="lambda:InvokeFunction",
            function="MyDemoFunction",
            principal="apigateway.amazonaws.com",
            source_arn=my_demo_api.execution_arn.apply(lambda execution_arn: f"{execution_arn}/*/*/*"))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The AWS Lambda action you want to allow in this statement. (e.g. `lambda:InvokeFunction`)
        :param pulumi.Input[str] event_source_token: The Event Source Token to validate.  Used with [Alexa Skills](https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-an-aws-lambda-function.html#use-aws-cli).
        :param pulumi.Input[str] function: Name of the Lambda function whose resource policy you are updating
        :param pulumi.Input[str] principal: The principal who is getting this permission.
               e.g. `s3.amazonaws.com`, an AWS account ID, or any valid AWS service principal
               such as `events.amazonaws.com` or `sns.amazonaws.com`.
        :param pulumi.Input[str] qualifier: Query parameter to specify function version or alias name.
               The permission will then apply to the specific qualified ARN.
               e.g. `arn:aws:lambda:aws-region:acct-id:function:function-name:2`
        :param pulumi.Input[str] source_account: This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.
        :param pulumi.Input[str] source_arn: When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to.
               Without this, any resource from `principal` will be granted permission – even if that resource is from another account.
               For S3, this should be the ARN of the S3 Bucket.
               For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule.
               For API Gateway, this should be the ARN of the API, as described [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
        :param pulumi.Input[str] statement_id: A unique statement identifier. By default generated by this provider.
        :param pulumi.Input[str] statement_id_prefix: A statement identifier prefix. This provider will generate a unique suffix. Conflicts with `statement_id`.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if action is None:
                raise TypeError("Missing required property 'action'")
            __props__['action'] = action
            __props__['event_source_token'] = event_source_token
            if function is None:
                raise TypeError("Missing required property 'function'")
            __props__['function'] = function
            if principal is None:
                raise TypeError("Missing required property 'principal'")
            __props__['principal'] = principal
            __props__['qualifier'] = qualifier
            __props__['source_account'] = source_account
            __props__['source_arn'] = source_arn
            __props__['statement_id'] = statement_id
            __props__['statement_id_prefix'] = statement_id_prefix
        super(Permission, __self__).__init__(
            'aws:lambda/permission:Permission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            event_source_token: Optional[pulumi.Input[str]] = None,
            function: Optional[pulumi.Input[str]] = None,
            principal: Optional[pulumi.Input[str]] = None,
            qualifier: Optional[pulumi.Input[str]] = None,
            source_account: Optional[pulumi.Input[str]] = None,
            source_arn: Optional[pulumi.Input[str]] = None,
            statement_id: Optional[pulumi.Input[str]] = None,
            statement_id_prefix: Optional[pulumi.Input[str]] = None) -> 'Permission':
        """
        Get an existing Permission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The AWS Lambda action you want to allow in this statement. (e.g. `lambda:InvokeFunction`)
        :param pulumi.Input[str] event_source_token: The Event Source Token to validate.  Used with [Alexa Skills](https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-an-aws-lambda-function.html#use-aws-cli).
        :param pulumi.Input[str] function: Name of the Lambda function whose resource policy you are updating
        :param pulumi.Input[str] principal: The principal who is getting this permission.
               e.g. `s3.amazonaws.com`, an AWS account ID, or any valid AWS service principal
               such as `events.amazonaws.com` or `sns.amazonaws.com`.
        :param pulumi.Input[str] qualifier: Query parameter to specify function version or alias name.
               The permission will then apply to the specific qualified ARN.
               e.g. `arn:aws:lambda:aws-region:acct-id:function:function-name:2`
        :param pulumi.Input[str] source_account: This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.
        :param pulumi.Input[str] source_arn: When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to.
               Without this, any resource from `principal` will be granted permission – even if that resource is from another account.
               For S3, this should be the ARN of the S3 Bucket.
               For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule.
               For API Gateway, this should be the ARN of the API, as described [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
        :param pulumi.Input[str] statement_id: A unique statement identifier. By default generated by this provider.
        :param pulumi.Input[str] statement_id_prefix: A statement identifier prefix. This provider will generate a unique suffix. Conflicts with `statement_id`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["action"] = action
        __props__["event_source_token"] = event_source_token
        __props__["function"] = function
        __props__["principal"] = principal
        __props__["qualifier"] = qualifier
        __props__["source_account"] = source_account
        __props__["source_arn"] = source_arn
        __props__["statement_id"] = statement_id
        __props__["statement_id_prefix"] = statement_id_prefix
        return Permission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> str:
        """
        The AWS Lambda action you want to allow in this statement. (e.g. `lambda:InvokeFunction`)
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="eventSourceToken")
    def event_source_token(self) -> Optional[str]:
        """
        The Event Source Token to validate.  Used with [Alexa Skills](https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-an-aws-lambda-function.html#use-aws-cli).
        """
        return pulumi.get(self, "event_source_token")

    @property
    @pulumi.getter
    def function(self) -> str:
        """
        Name of the Lambda function whose resource policy you are updating
        """
        return pulumi.get(self, "function")

    @property
    @pulumi.getter
    def principal(self) -> str:
        """
        The principal who is getting this permission.
        e.g. `s3.amazonaws.com`, an AWS account ID, or any valid AWS service principal
        such as `events.amazonaws.com` or `sns.amazonaws.com`.
        """
        return pulumi.get(self, "principal")

    @property
    @pulumi.getter
    def qualifier(self) -> Optional[str]:
        """
        Query parameter to specify function version or alias name.
        The permission will then apply to the specific qualified ARN.
        e.g. `arn:aws:lambda:aws-region:acct-id:function:function-name:2`
        """
        return pulumi.get(self, "qualifier")

    @property
    @pulumi.getter(name="sourceAccount")
    def source_account(self) -> Optional[str]:
        """
        This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.
        """
        return pulumi.get(self, "source_account")

    @property
    @pulumi.getter(name="sourceArn")
    def source_arn(self) -> Optional[str]:
        """
        When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to.
        Without this, any resource from `principal` will be granted permission – even if that resource is from another account.
        For S3, this should be the ARN of the S3 Bucket.
        For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule.
        For API Gateway, this should be the ARN of the API, as described [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
        """
        return pulumi.get(self, "source_arn")

    @property
    @pulumi.getter(name="statementId")
    def statement_id(self) -> str:
        """
        A unique statement identifier. By default generated by this provider.
        """
        return pulumi.get(self, "statement_id")

    @property
    @pulumi.getter(name="statementIdPrefix")
    def statement_id_prefix(self) -> Optional[str]:
        """
        A statement identifier prefix. This provider will generate a unique suffix. Conflicts with `statement_id`.
        """
        return pulumi.get(self, "statement_id_prefix")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

