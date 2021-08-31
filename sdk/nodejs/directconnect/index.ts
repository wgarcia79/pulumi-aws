// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./bgpPeer";
export * from "./connection";
export * from "./connectionAssociation";
export * from "./gateway";
export * from "./gatewayAssociation";
export * from "./gatewayAssociationProposal";
export * from "./getConnection";
export * from "./getGateway";
export * from "./getLocation";
export * from "./getLocations";
export * from "./hostedPrivateVirtualInterface";
export * from "./hostedPrivateVirtualInterfaceAccepter";
export * from "./hostedPublicVirtualInterface";
export * from "./hostedPublicVirtualInterfaceAccepter";
export * from "./hostedTransitVirtualInterface";
export * from "./hostedTransitVirtualInterfaceAcceptor";
export * from "./linkAggregationGroup";
export * from "./privateVirtualInterface";
export * from "./publicVirtualInterface";
export * from "./transitVirtualInterface";

// Import resources to register:
import { BgpPeer } from "./bgpPeer";
import { Connection } from "./connection";
import { ConnectionAssociation } from "./connectionAssociation";
import { Gateway } from "./gateway";
import { GatewayAssociation } from "./gatewayAssociation";
import { GatewayAssociationProposal } from "./gatewayAssociationProposal";
import { HostedPrivateVirtualInterface } from "./hostedPrivateVirtualInterface";
import { HostedPrivateVirtualInterfaceAccepter } from "./hostedPrivateVirtualInterfaceAccepter";
import { HostedPublicVirtualInterface } from "./hostedPublicVirtualInterface";
import { HostedPublicVirtualInterfaceAccepter } from "./hostedPublicVirtualInterfaceAccepter";
import { HostedTransitVirtualInterface } from "./hostedTransitVirtualInterface";
import { HostedTransitVirtualInterfaceAcceptor } from "./hostedTransitVirtualInterfaceAcceptor";
import { LinkAggregationGroup } from "./linkAggregationGroup";
import { PrivateVirtualInterface } from "./privateVirtualInterface";
import { PublicVirtualInterface } from "./publicVirtualInterface";
import { TransitVirtualInterface } from "./transitVirtualInterface";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:directconnect/bgpPeer:BgpPeer":
                return new BgpPeer(name, <any>undefined, { urn })
            case "aws:directconnect/connection:Connection":
                return new Connection(name, <any>undefined, { urn })
            case "aws:directconnect/connectionAssociation:ConnectionAssociation":
                return new ConnectionAssociation(name, <any>undefined, { urn })
            case "aws:directconnect/gateway:Gateway":
                return new Gateway(name, <any>undefined, { urn })
            case "aws:directconnect/gatewayAssociation:GatewayAssociation":
                return new GatewayAssociation(name, <any>undefined, { urn })
            case "aws:directconnect/gatewayAssociationProposal:GatewayAssociationProposal":
                return new GatewayAssociationProposal(name, <any>undefined, { urn })
            case "aws:directconnect/hostedPrivateVirtualInterface:HostedPrivateVirtualInterface":
                return new HostedPrivateVirtualInterface(name, <any>undefined, { urn })
            case "aws:directconnect/hostedPrivateVirtualInterfaceAccepter:HostedPrivateVirtualInterfaceAccepter":
                return new HostedPrivateVirtualInterfaceAccepter(name, <any>undefined, { urn })
            case "aws:directconnect/hostedPublicVirtualInterface:HostedPublicVirtualInterface":
                return new HostedPublicVirtualInterface(name, <any>undefined, { urn })
            case "aws:directconnect/hostedPublicVirtualInterfaceAccepter:HostedPublicVirtualInterfaceAccepter":
                return new HostedPublicVirtualInterfaceAccepter(name, <any>undefined, { urn })
            case "aws:directconnect/hostedTransitVirtualInterface:HostedTransitVirtualInterface":
                return new HostedTransitVirtualInterface(name, <any>undefined, { urn })
            case "aws:directconnect/hostedTransitVirtualInterfaceAcceptor:HostedTransitVirtualInterfaceAcceptor":
                return new HostedTransitVirtualInterfaceAcceptor(name, <any>undefined, { urn })
            case "aws:directconnect/linkAggregationGroup:LinkAggregationGroup":
                return new LinkAggregationGroup(name, <any>undefined, { urn })
            case "aws:directconnect/privateVirtualInterface:PrivateVirtualInterface":
                return new PrivateVirtualInterface(name, <any>undefined, { urn })
            case "aws:directconnect/publicVirtualInterface:PublicVirtualInterface":
                return new PublicVirtualInterface(name, <any>undefined, { urn })
            case "aws:directconnect/transitVirtualInterface:TransitVirtualInterface":
                return new TransitVirtualInterface(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "directconnect/bgpPeer", _module)
pulumi.runtime.registerResourceModule("aws", "directconnect/connection", _module)
pulumi.runtime.registerResourceModule("aws", "directconnect/connectionAssociation", _module)
pulumi.runtime.registerResourceModule("aws", "directconnect/gateway", _module)
pulumi.runtime.registerResourceModule("aws", "directconnect/gatewayAssociation", _module)
pulumi.runtime.registerResourceModule("aws", "directconnect/gatewayAssociationProposal", _module)
pulumi.runtime.registerResourceModule("aws", "directconnect/hostedPrivateVirtualInterface", _module)
pulumi.runtime.registerResourceModule("aws", "directconnect/hostedPrivateVirtualInterfaceAccepter", _module)
pulumi.runtime.registerResourceModule("aws", "directconnect/hostedPublicVirtualInterface", _module)
pulumi.runtime.registerResourceModule("aws", "directconnect/hostedPublicVirtualInterfaceAccepter", _module)
pulumi.runtime.registerResourceModule("aws", "directconnect/hostedTransitVirtualInterface", _module)
pulumi.runtime.registerResourceModule("aws", "directconnect/hostedTransitVirtualInterfaceAcceptor", _module)
pulumi.runtime.registerResourceModule("aws", "directconnect/linkAggregationGroup", _module)
pulumi.runtime.registerResourceModule("aws", "directconnect/privateVirtualInterface", _module)
pulumi.runtime.registerResourceModule("aws", "directconnect/publicVirtualInterface", _module)
pulumi.runtime.registerResourceModule("aws", "directconnect/transitVirtualInterface", _module)
