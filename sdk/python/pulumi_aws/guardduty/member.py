# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MemberArgs', 'Member']

@pulumi.input_type
class MemberArgs:
    def __init__(__self__, *,
                 account_id: pulumi.Input[str],
                 detector_id: pulumi.Input[str],
                 email: pulumi.Input[str],
                 disable_email_notification: Optional[pulumi.Input[bool]] = None,
                 invitation_message: Optional[pulumi.Input[str]] = None,
                 invite: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Member resource.
        :param pulumi.Input[str] account_id: AWS account ID for member account.
        :param pulumi.Input[str] detector_id: The detector ID of the GuardDuty account where you want to create member accounts.
        :param pulumi.Input[str] email: Email address for member account.
        :param pulumi.Input[bool] disable_email_notification: Boolean whether an email notification is sent to the accounts. Defaults to `false`.
        :param pulumi.Input[str] invitation_message: Message for invitation.
        :param pulumi.Input[bool] invite: Boolean whether to invite the account to GuardDuty as a member. Defaults to `false`. To detect if an invitation needs to be (re-)sent, the this provider state value is `true` based on a `relationship_status` of `Disabled`, `Enabled`, `Invited`, or `EmailVerificationInProgress`.
        """
        pulumi.set(__self__, "account_id", account_id)
        pulumi.set(__self__, "detector_id", detector_id)
        pulumi.set(__self__, "email", email)
        if disable_email_notification is not None:
            pulumi.set(__self__, "disable_email_notification", disable_email_notification)
        if invitation_message is not None:
            pulumi.set(__self__, "invitation_message", invitation_message)
        if invite is not None:
            pulumi.set(__self__, "invite", invite)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Input[str]:
        """
        AWS account ID for member account.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="detectorId")
    def detector_id(self) -> pulumi.Input[str]:
        """
        The detector ID of the GuardDuty account where you want to create member accounts.
        """
        return pulumi.get(self, "detector_id")

    @detector_id.setter
    def detector_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "detector_id", value)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        """
        Email address for member account.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="disableEmailNotification")
    def disable_email_notification(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean whether an email notification is sent to the accounts. Defaults to `false`.
        """
        return pulumi.get(self, "disable_email_notification")

    @disable_email_notification.setter
    def disable_email_notification(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_email_notification", value)

    @property
    @pulumi.getter(name="invitationMessage")
    def invitation_message(self) -> Optional[pulumi.Input[str]]:
        """
        Message for invitation.
        """
        return pulumi.get(self, "invitation_message")

    @invitation_message.setter
    def invitation_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "invitation_message", value)

    @property
    @pulumi.getter
    def invite(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean whether to invite the account to GuardDuty as a member. Defaults to `false`. To detect if an invitation needs to be (re-)sent, the this provider state value is `true` based on a `relationship_status` of `Disabled`, `Enabled`, `Invited`, or `EmailVerificationInProgress`.
        """
        return pulumi.get(self, "invite")

    @invite.setter
    def invite(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "invite", value)


@pulumi.input_type
class _MemberState:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 detector_id: Optional[pulumi.Input[str]] = None,
                 disable_email_notification: Optional[pulumi.Input[bool]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 invitation_message: Optional[pulumi.Input[str]] = None,
                 invite: Optional[pulumi.Input[bool]] = None,
                 relationship_status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Member resources.
        :param pulumi.Input[str] account_id: AWS account ID for member account.
        :param pulumi.Input[str] detector_id: The detector ID of the GuardDuty account where you want to create member accounts.
        :param pulumi.Input[bool] disable_email_notification: Boolean whether an email notification is sent to the accounts. Defaults to `false`.
        :param pulumi.Input[str] email: Email address for member account.
        :param pulumi.Input[str] invitation_message: Message for invitation.
        :param pulumi.Input[bool] invite: Boolean whether to invite the account to GuardDuty as a member. Defaults to `false`. To detect if an invitation needs to be (re-)sent, the this provider state value is `true` based on a `relationship_status` of `Disabled`, `Enabled`, `Invited`, or `EmailVerificationInProgress`.
        :param pulumi.Input[str] relationship_status: The status of the relationship between the member account and its primary account. More information can be found in [Amazon GuardDuty API Reference](https://docs.aws.amazon.com/guardduty/latest/ug/get-members.html).
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if detector_id is not None:
            pulumi.set(__self__, "detector_id", detector_id)
        if disable_email_notification is not None:
            pulumi.set(__self__, "disable_email_notification", disable_email_notification)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if invitation_message is not None:
            pulumi.set(__self__, "invitation_message", invitation_message)
        if invite is not None:
            pulumi.set(__self__, "invite", invite)
        if relationship_status is not None:
            pulumi.set(__self__, "relationship_status", relationship_status)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        AWS account ID for member account.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="detectorId")
    def detector_id(self) -> Optional[pulumi.Input[str]]:
        """
        The detector ID of the GuardDuty account where you want to create member accounts.
        """
        return pulumi.get(self, "detector_id")

    @detector_id.setter
    def detector_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "detector_id", value)

    @property
    @pulumi.getter(name="disableEmailNotification")
    def disable_email_notification(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean whether an email notification is sent to the accounts. Defaults to `false`.
        """
        return pulumi.get(self, "disable_email_notification")

    @disable_email_notification.setter
    def disable_email_notification(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_email_notification", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        Email address for member account.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="invitationMessage")
    def invitation_message(self) -> Optional[pulumi.Input[str]]:
        """
        Message for invitation.
        """
        return pulumi.get(self, "invitation_message")

    @invitation_message.setter
    def invitation_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "invitation_message", value)

    @property
    @pulumi.getter
    def invite(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean whether to invite the account to GuardDuty as a member. Defaults to `false`. To detect if an invitation needs to be (re-)sent, the this provider state value is `true` based on a `relationship_status` of `Disabled`, `Enabled`, `Invited`, or `EmailVerificationInProgress`.
        """
        return pulumi.get(self, "invite")

    @invite.setter
    def invite(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "invite", value)

    @property
    @pulumi.getter(name="relationshipStatus")
    def relationship_status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the relationship between the member account and its primary account. More information can be found in [Amazon GuardDuty API Reference](https://docs.aws.amazon.com/guardduty/latest/ug/get-members.html).
        """
        return pulumi.get(self, "relationship_status")

    @relationship_status.setter
    def relationship_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "relationship_status", value)


class Member(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 detector_id: Optional[pulumi.Input[str]] = None,
                 disable_email_notification: Optional[pulumi.Input[bool]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 invitation_message: Optional[pulumi.Input[str]] = None,
                 invite: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Provides a resource to manage a GuardDuty member. To accept invitations in member accounts, see the `guardduty.InviteAccepter` resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        primary = aws.guardduty.Detector("primary", enable=True)
        member_detector = aws.guardduty.Detector("memberDetector", enable=True,
        opts=pulumi.ResourceOptions(provider=aws["dev"]))
        member_member = aws.guardduty.Member("memberMember",
            account_id=member_detector.account_id,
            detector_id=primary.id,
            email="required@example.com",
            invite=True,
            invitation_message="please accept guardduty invitation")
        ```

        ## Import

        GuardDuty members can be imported using the the primary GuardDuty detector ID and member AWS account ID, e.g.,

        ```sh
         $ pulumi import aws:guardduty/member:Member MyMember 00b00fd5aecc0ab60a708659477e9617:123456789012
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: AWS account ID for member account.
        :param pulumi.Input[str] detector_id: The detector ID of the GuardDuty account where you want to create member accounts.
        :param pulumi.Input[bool] disable_email_notification: Boolean whether an email notification is sent to the accounts. Defaults to `false`.
        :param pulumi.Input[str] email: Email address for member account.
        :param pulumi.Input[str] invitation_message: Message for invitation.
        :param pulumi.Input[bool] invite: Boolean whether to invite the account to GuardDuty as a member. Defaults to `false`. To detect if an invitation needs to be (re-)sent, the this provider state value is `true` based on a `relationship_status` of `Disabled`, `Enabled`, `Invited`, or `EmailVerificationInProgress`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage a GuardDuty member. To accept invitations in member accounts, see the `guardduty.InviteAccepter` resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        primary = aws.guardduty.Detector("primary", enable=True)
        member_detector = aws.guardduty.Detector("memberDetector", enable=True,
        opts=pulumi.ResourceOptions(provider=aws["dev"]))
        member_member = aws.guardduty.Member("memberMember",
            account_id=member_detector.account_id,
            detector_id=primary.id,
            email="required@example.com",
            invite=True,
            invitation_message="please accept guardduty invitation")
        ```

        ## Import

        GuardDuty members can be imported using the the primary GuardDuty detector ID and member AWS account ID, e.g.,

        ```sh
         $ pulumi import aws:guardduty/member:Member MyMember 00b00fd5aecc0ab60a708659477e9617:123456789012
        ```

        :param str resource_name: The name of the resource.
        :param MemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 detector_id: Optional[pulumi.Input[str]] = None,
                 disable_email_notification: Optional[pulumi.Input[bool]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 invitation_message: Optional[pulumi.Input[str]] = None,
                 invite: Optional[pulumi.Input[bool]] = None,
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
            __props__ = MemberArgs.__new__(MemberArgs)

            if account_id is None and not opts.urn:
                raise TypeError("Missing required property 'account_id'")
            __props__.__dict__["account_id"] = account_id
            if detector_id is None and not opts.urn:
                raise TypeError("Missing required property 'detector_id'")
            __props__.__dict__["detector_id"] = detector_id
            __props__.__dict__["disable_email_notification"] = disable_email_notification
            if email is None and not opts.urn:
                raise TypeError("Missing required property 'email'")
            __props__.__dict__["email"] = email
            __props__.__dict__["invitation_message"] = invitation_message
            __props__.__dict__["invite"] = invite
            __props__.__dict__["relationship_status"] = None
        super(Member, __self__).__init__(
            'aws:guardduty/member:Member',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            detector_id: Optional[pulumi.Input[str]] = None,
            disable_email_notification: Optional[pulumi.Input[bool]] = None,
            email: Optional[pulumi.Input[str]] = None,
            invitation_message: Optional[pulumi.Input[str]] = None,
            invite: Optional[pulumi.Input[bool]] = None,
            relationship_status: Optional[pulumi.Input[str]] = None) -> 'Member':
        """
        Get an existing Member resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: AWS account ID for member account.
        :param pulumi.Input[str] detector_id: The detector ID of the GuardDuty account where you want to create member accounts.
        :param pulumi.Input[bool] disable_email_notification: Boolean whether an email notification is sent to the accounts. Defaults to `false`.
        :param pulumi.Input[str] email: Email address for member account.
        :param pulumi.Input[str] invitation_message: Message for invitation.
        :param pulumi.Input[bool] invite: Boolean whether to invite the account to GuardDuty as a member. Defaults to `false`. To detect if an invitation needs to be (re-)sent, the this provider state value is `true` based on a `relationship_status` of `Disabled`, `Enabled`, `Invited`, or `EmailVerificationInProgress`.
        :param pulumi.Input[str] relationship_status: The status of the relationship between the member account and its primary account. More information can be found in [Amazon GuardDuty API Reference](https://docs.aws.amazon.com/guardduty/latest/ug/get-members.html).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MemberState.__new__(_MemberState)

        __props__.__dict__["account_id"] = account_id
        __props__.__dict__["detector_id"] = detector_id
        __props__.__dict__["disable_email_notification"] = disable_email_notification
        __props__.__dict__["email"] = email
        __props__.__dict__["invitation_message"] = invitation_message
        __props__.__dict__["invite"] = invite
        __props__.__dict__["relationship_status"] = relationship_status
        return Member(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        """
        AWS account ID for member account.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="detectorId")
    def detector_id(self) -> pulumi.Output[str]:
        """
        The detector ID of the GuardDuty account where you want to create member accounts.
        """
        return pulumi.get(self, "detector_id")

    @property
    @pulumi.getter(name="disableEmailNotification")
    def disable_email_notification(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean whether an email notification is sent to the accounts. Defaults to `false`.
        """
        return pulumi.get(self, "disable_email_notification")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        Email address for member account.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="invitationMessage")
    def invitation_message(self) -> pulumi.Output[Optional[str]]:
        """
        Message for invitation.
        """
        return pulumi.get(self, "invitation_message")

    @property
    @pulumi.getter
    def invite(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean whether to invite the account to GuardDuty as a member. Defaults to `false`. To detect if an invitation needs to be (re-)sent, the this provider state value is `true` based on a `relationship_status` of `Disabled`, `Enabled`, `Invited`, or `EmailVerificationInProgress`.
        """
        return pulumi.get(self, "invite")

    @property
    @pulumi.getter(name="relationshipStatus")
    def relationship_status(self) -> pulumi.Output[str]:
        """
        The status of the relationship between the member account and its primary account. More information can be found in [Amazon GuardDuty API Reference](https://docs.aws.amazon.com/guardduty/latest/ug/get-members.html).
        """
        return pulumi.get(self, "relationship_status")

