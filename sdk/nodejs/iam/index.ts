// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./accessKey";
export * from "./accountAlias";
export * from "./accountPasswordPolicy";
export * from "./documents";
export * from "./getAccountAlias";
export * from "./getGroup";
export * from "./getInstanceProfile";
export * from "./getPolicy";
export * from "./getPolicyDocument";
export * from "./getRole";
export * from "./getServerCertificate";
export * from "./getSessionContext";
export * from "./getUser";
export * from "./group";
export * from "./groupMembership";
export * from "./groupPolicy";
export * from "./groupPolicyAttachment";
export * from "./instanceProfile";
export * from "./managedPolicies";
export * from "./openIdConnectProvider";
export * from "./policy";
export * from "./policyAttachment";
export * from "./principals";
export * from "./role";
export * from "./rolePolicy";
export * from "./rolePolicyAttachment";
export * from "./samlProvider";
export * from "./serverCertificate";
export * from "./serviceLinkedRole";
export * from "./sshKey";
export * from "./user";
export * from "./userGroupMembership";
export * from "./userLoginProfile";
export * from "./userPolicy";
export * from "./userPolicyAttachment";

// Export enums:
export * from "../types/enums/iam";

// Import resources to register:
import { AccessKey } from "./accessKey";
import { AccountAlias } from "./accountAlias";
import { AccountPasswordPolicy } from "./accountPasswordPolicy";
import { Group } from "./group";
import { GroupMembership } from "./groupMembership";
import { GroupPolicy } from "./groupPolicy";
import { GroupPolicyAttachment } from "./groupPolicyAttachment";
import { InstanceProfile } from "./instanceProfile";
import { OpenIdConnectProvider } from "./openIdConnectProvider";
import { Policy } from "./policy";
import { PolicyAttachment } from "./policyAttachment";
import { Role } from "./role";
import { RolePolicy } from "./rolePolicy";
import { RolePolicyAttachment } from "./rolePolicyAttachment";
import { SamlProvider } from "./samlProvider";
import { ServerCertificate } from "./serverCertificate";
import { ServiceLinkedRole } from "./serviceLinkedRole";
import { SshKey } from "./sshKey";
import { User } from "./user";
import { UserGroupMembership } from "./userGroupMembership";
import { UserLoginProfile } from "./userLoginProfile";
import { UserPolicy } from "./userPolicy";
import { UserPolicyAttachment } from "./userPolicyAttachment";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:iam/accessKey:AccessKey":
                return new AccessKey(name, <any>undefined, { urn })
            case "aws:iam/accountAlias:AccountAlias":
                return new AccountAlias(name, <any>undefined, { urn })
            case "aws:iam/accountPasswordPolicy:AccountPasswordPolicy":
                return new AccountPasswordPolicy(name, <any>undefined, { urn })
            case "aws:iam/group:Group":
                return new Group(name, <any>undefined, { urn })
            case "aws:iam/groupMembership:GroupMembership":
                return new GroupMembership(name, <any>undefined, { urn })
            case "aws:iam/groupPolicy:GroupPolicy":
                return new GroupPolicy(name, <any>undefined, { urn })
            case "aws:iam/groupPolicyAttachment:GroupPolicyAttachment":
                return new GroupPolicyAttachment(name, <any>undefined, { urn })
            case "aws:iam/instanceProfile:InstanceProfile":
                return new InstanceProfile(name, <any>undefined, { urn })
            case "aws:iam/openIdConnectProvider:OpenIdConnectProvider":
                return new OpenIdConnectProvider(name, <any>undefined, { urn })
            case "aws:iam/policy:Policy":
                return new Policy(name, <any>undefined, { urn })
            case "aws:iam/policyAttachment:PolicyAttachment":
                return new PolicyAttachment(name, <any>undefined, { urn })
            case "aws:iam/role:Role":
                return new Role(name, <any>undefined, { urn })
            case "aws:iam/rolePolicy:RolePolicy":
                return new RolePolicy(name, <any>undefined, { urn })
            case "aws:iam/rolePolicyAttachment:RolePolicyAttachment":
                return new RolePolicyAttachment(name, <any>undefined, { urn })
            case "aws:iam/samlProvider:SamlProvider":
                return new SamlProvider(name, <any>undefined, { urn })
            case "aws:iam/serverCertificate:ServerCertificate":
                return new ServerCertificate(name, <any>undefined, { urn })
            case "aws:iam/serviceLinkedRole:ServiceLinkedRole":
                return new ServiceLinkedRole(name, <any>undefined, { urn })
            case "aws:iam/sshKey:SshKey":
                return new SshKey(name, <any>undefined, { urn })
            case "aws:iam/user:User":
                return new User(name, <any>undefined, { urn })
            case "aws:iam/userGroupMembership:UserGroupMembership":
                return new UserGroupMembership(name, <any>undefined, { urn })
            case "aws:iam/userLoginProfile:UserLoginProfile":
                return new UserLoginProfile(name, <any>undefined, { urn })
            case "aws:iam/userPolicy:UserPolicy":
                return new UserPolicy(name, <any>undefined, { urn })
            case "aws:iam/userPolicyAttachment:UserPolicyAttachment":
                return new UserPolicyAttachment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "iam/accessKey", _module)
pulumi.runtime.registerResourceModule("aws", "iam/accountAlias", _module)
pulumi.runtime.registerResourceModule("aws", "iam/accountPasswordPolicy", _module)
pulumi.runtime.registerResourceModule("aws", "iam/group", _module)
pulumi.runtime.registerResourceModule("aws", "iam/groupMembership", _module)
pulumi.runtime.registerResourceModule("aws", "iam/groupPolicy", _module)
pulumi.runtime.registerResourceModule("aws", "iam/groupPolicyAttachment", _module)
pulumi.runtime.registerResourceModule("aws", "iam/instanceProfile", _module)
pulumi.runtime.registerResourceModule("aws", "iam/openIdConnectProvider", _module)
pulumi.runtime.registerResourceModule("aws", "iam/policy", _module)
pulumi.runtime.registerResourceModule("aws", "iam/policyAttachment", _module)
pulumi.runtime.registerResourceModule("aws", "iam/role", _module)
pulumi.runtime.registerResourceModule("aws", "iam/rolePolicy", _module)
pulumi.runtime.registerResourceModule("aws", "iam/rolePolicyAttachment", _module)
pulumi.runtime.registerResourceModule("aws", "iam/samlProvider", _module)
pulumi.runtime.registerResourceModule("aws", "iam/serverCertificate", _module)
pulumi.runtime.registerResourceModule("aws", "iam/serviceLinkedRole", _module)
pulumi.runtime.registerResourceModule("aws", "iam/sshKey", _module)
pulumi.runtime.registerResourceModule("aws", "iam/user", _module)
pulumi.runtime.registerResourceModule("aws", "iam/userGroupMembership", _module)
pulumi.runtime.registerResourceModule("aws", "iam/userLoginProfile", _module)
pulumi.runtime.registerResourceModule("aws", "iam/userPolicy", _module)
pulumi.runtime.registerResourceModule("aws", "iam/userPolicyAttachment", _module)
