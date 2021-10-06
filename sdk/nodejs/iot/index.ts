// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./authorizer";
export * from "./certificate";
export * from "./getEndpoint";
export * from "./policy";
export * from "./policyAttachment";
export * from "./roleAlias";
export * from "./thing";
export * from "./thingPrincipalAttachment";
export * from "./thingType";
export * from "./topicRule";

// Import resources to register:
import { Authorizer } from "./authorizer";
import { Certificate } from "./certificate";
import { Policy } from "./policy";
import { PolicyAttachment } from "./policyAttachment";
import { RoleAlias } from "./roleAlias";
import { Thing } from "./thing";
import { ThingPrincipalAttachment } from "./thingPrincipalAttachment";
import { ThingType } from "./thingType";
import { TopicRule } from "./topicRule";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:iot/authorizer:Authorizer":
                return new Authorizer(name, <any>undefined, { urn })
            case "aws:iot/certificate:Certificate":
                return new Certificate(name, <any>undefined, { urn })
            case "aws:iot/policy:Policy":
                return new Policy(name, <any>undefined, { urn })
            case "aws:iot/policyAttachment:PolicyAttachment":
                return new PolicyAttachment(name, <any>undefined, { urn })
            case "aws:iot/roleAlias:RoleAlias":
                return new RoleAlias(name, <any>undefined, { urn })
            case "aws:iot/thing:Thing":
                return new Thing(name, <any>undefined, { urn })
            case "aws:iot/thingPrincipalAttachment:ThingPrincipalAttachment":
                return new ThingPrincipalAttachment(name, <any>undefined, { urn })
            case "aws:iot/thingType:ThingType":
                return new ThingType(name, <any>undefined, { urn })
            case "aws:iot/topicRule:TopicRule":
                return new TopicRule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "iot/authorizer", _module)
pulumi.runtime.registerResourceModule("aws", "iot/certificate", _module)
pulumi.runtime.registerResourceModule("aws", "iot/policy", _module)
pulumi.runtime.registerResourceModule("aws", "iot/policyAttachment", _module)
pulumi.runtime.registerResourceModule("aws", "iot/roleAlias", _module)
pulumi.runtime.registerResourceModule("aws", "iot/thing", _module)
pulumi.runtime.registerResourceModule("aws", "iot/thingPrincipalAttachment", _module)
pulumi.runtime.registerResourceModule("aws", "iot/thingType", _module)
pulumi.runtime.registerResourceModule("aws", "iot/topicRule", _module)
