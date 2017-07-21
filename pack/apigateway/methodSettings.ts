// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

import {RestApi} from "./restApi";

export class MethodSettings extends lumi.NamedResource implements MethodSettingsArgs {
    public readonly methodPath: string;
    public readonly restApi: RestApi;
    public readonly settings: { cacheDataEncrypted?: boolean, cacheTtlInSeconds?: number, cachingEnabled?: boolean, dataTraceEnabled?: boolean, loggingLevel?: string, metricsEnabled?: boolean, requireAuthorizationForCacheControl?: boolean, throttlingBurstLimit?: number, throttlingRateLimit?: number, unauthorizedCacheControlHeaderStrategy?: string }[];
    public readonly stageName: string;

    public static get(id: lumi.ID): MethodSettings {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): MethodSettings[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: MethodSettingsArgs) {
        super(name);
        if (lumirt.defaultIfComputed(args.methodPath, "") === undefined) {
            throw new Error("Property argument 'methodPath' is required, but was missing");
        }
        this.methodPath = args.methodPath;
        if (lumirt.defaultIfComputed(args.restApi, "") === undefined) {
            throw new Error("Property argument 'restApi' is required, but was missing");
        }
        this.restApi = args.restApi;
        if (lumirt.defaultIfComputed(args.settings, "") === undefined) {
            throw new Error("Property argument 'settings' is required, but was missing");
        }
        this.settings = args.settings;
        if (lumirt.defaultIfComputed(args.stageName, "") === undefined) {
            throw new Error("Property argument 'stageName' is required, but was missing");
        }
        this.stageName = args.stageName;
    }
}

export interface MethodSettingsArgs {
    readonly methodPath: string;
    readonly restApi: RestApi;
    readonly settings: { cacheDataEncrypted?: boolean, cacheTtlInSeconds?: number, cachingEnabled?: boolean, dataTraceEnabled?: boolean, loggingLevel?: string, metricsEnabled?: boolean, requireAuthorizationForCacheControl?: boolean, throttlingBurstLimit?: number, throttlingRateLimit?: number, unauthorizedCacheControlHeaderStrategy?: string }[];
    readonly stageName: string;
}
