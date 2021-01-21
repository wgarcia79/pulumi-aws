// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws:iam/accessKey:AccessKey":
		r, err = NewAccessKey(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/accountAlias:AccountAlias":
		r, err = NewAccountAlias(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/accountPasswordPolicy:AccountPasswordPolicy":
		r, err = NewAccountPasswordPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/group:Group":
		r, err = NewGroup(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/groupMembership:GroupMembership":
		r, err = NewGroupMembership(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/groupPolicy:GroupPolicy":
		r, err = NewGroupPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/groupPolicyAttachment:GroupPolicyAttachment":
		r, err = NewGroupPolicyAttachment(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/instanceProfile:InstanceProfile":
		r, err = NewInstanceProfile(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/openIdConnectProvider:OpenIdConnectProvider":
		r, err = NewOpenIdConnectProvider(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/policy:Policy":
		r, err = NewPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/policyAttachment:PolicyAttachment":
		r, err = NewPolicyAttachment(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/role:Role":
		r, err = NewRole(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/rolePolicy:RolePolicy":
		r, err = NewRolePolicy(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/rolePolicyAttachment:RolePolicyAttachment":
		r, err = NewRolePolicyAttachment(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/samlProvider:SamlProvider":
		r, err = NewSamlProvider(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/serverCertificate:ServerCertificate":
		r, err = NewServerCertificate(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/serviceLinkedRole:ServiceLinkedRole":
		r, err = NewServiceLinkedRole(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/sshKey:SshKey":
		r, err = NewSshKey(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/user:User":
		r, err = NewUser(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/userGroupMembership:UserGroupMembership":
		r, err = NewUserGroupMembership(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/userLoginProfile:UserLoginProfile":
		r, err = NewUserLoginProfile(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/userPolicy:UserPolicy":
		r, err = NewUserPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "aws:iam/userPolicyAttachment:UserPolicyAttachment":
		r, err = NewUserPolicyAttachment(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"iam/accessKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/accountAlias",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/accountPasswordPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/groupMembership",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/groupPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/groupPolicyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/instanceProfile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/openIdConnectProvider",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/policy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/policyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/role",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/rolePolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/rolePolicyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/samlProvider",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/serverCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/serviceLinkedRole",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/sshKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/userGroupMembership",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/userLoginProfile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/userPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"iam/userPolicyAttachment",
		&module{version},
	)
}
