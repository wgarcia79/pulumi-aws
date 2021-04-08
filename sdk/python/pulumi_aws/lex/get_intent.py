# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetIntentResult',
    'AwaitableGetIntentResult',
    'get_intent',
]

@pulumi.output_type
class GetIntentResult:
    """
    A collection of values returned by getIntent.
    """
    def __init__(__self__, arn=None, checksum=None, created_date=None, description=None, id=None, last_updated_date=None, name=None, parent_intent_signature=None, version=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if checksum and not isinstance(checksum, str):
            raise TypeError("Expected argument 'checksum' to be a str")
        pulumi.set(__self__, "checksum", checksum)
        if created_date and not isinstance(created_date, str):
            raise TypeError("Expected argument 'created_date' to be a str")
        pulumi.set(__self__, "created_date", created_date)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_updated_date and not isinstance(last_updated_date, str):
            raise TypeError("Expected argument 'last_updated_date' to be a str")
        pulumi.set(__self__, "last_updated_date", last_updated_date)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parent_intent_signature and not isinstance(parent_intent_signature, str):
            raise TypeError("Expected argument 'parent_intent_signature' to be a str")
        pulumi.set(__self__, "parent_intent_signature", parent_intent_signature)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The ARN of the Lex intent.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def checksum(self) -> str:
        """
        Checksum identifying the version of the intent that was created. The checksum is not
        included as an argument because the resource will add it automatically when updating the intent.
        """
        return pulumi.get(self, "checksum")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> str:
        """
        The date when the intent version was created.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description of the intent.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastUpdatedDate")
    def last_updated_date(self) -> str:
        """
        The date when the $LATEST version of this intent was updated.
        """
        return pulumi.get(self, "last_updated_date")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the intent, not case sensitive.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentIntentSignature")
    def parent_intent_signature(self) -> str:
        """
        A unique identifier for the built-in intent to base this
        intent on. To find the signature for an intent, see
        [Standard Built-in Intents](https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents)
        in the Alexa Skills Kit.
        """
        return pulumi.get(self, "parent_intent_signature")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        The version of the bot.
        """
        return pulumi.get(self, "version")


class AwaitableGetIntentResult(GetIntentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIntentResult(
            arn=self.arn,
            checksum=self.checksum,
            created_date=self.created_date,
            description=self.description,
            id=self.id,
            last_updated_date=self.last_updated_date,
            name=self.name,
            parent_intent_signature=self.parent_intent_signature,
            version=self.version)


def get_intent(name: Optional[str] = None,
               version: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIntentResult:
    """
    Provides details about a specific Amazon Lex Intent.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    order_flowers = aws.lex.get_intent(name="OrderFlowers",
        version="$LATEST")
    ```


    :param str name: The name of the intent. The name is case sensitive.
    :param str version: The version of the intent.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['version'] = version
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:lex/getIntent:getIntent', __args__, opts=opts, typ=GetIntentResult).value

    return AwaitableGetIntentResult(
        arn=__ret__.arn,
        checksum=__ret__.checksum,
        created_date=__ret__.created_date,
        description=__ret__.description,
        id=__ret__.id,
        last_updated_date=__ret__.last_updated_date,
        name=__ret__.name,
        parent_intent_signature=__ret__.parent_intent_signature,
        version=__ret__.version)
