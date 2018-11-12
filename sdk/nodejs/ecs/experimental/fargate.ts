// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// These APIs are currently experimental and may change.

import * as fs from "fs";
import * as mime from "mime";
import * as fspath from "path";

import * as aws from "../..";
import * as lambda from "../../lambda";
import * as pulumi from "@pulumi/pulumi";
import * as autoscaling from "../../autoscaling";
import * as ec2 from "../../ec2";
import * as iam from "../../iam";

import { sha1hash, Overwrite } from "../../utils";

export * from "./cluster";
export * from "./clusterAutoScalingGroup";
export * from "./taskDefinition";
export * from "./vpc";


export type FargateTaskDefinition = Overwrite<TaskDefinition, {
    // Tasks that use the Fargate launch type do not support all of the task definition parameters
    // that are available. Some parameters are not supported at all, and others behave differently
    // for Fargate tasks.
    //
    // The following task definition parameters are not valid in Fargate tasks: 
    disableNetworking: never;
    dnsSearchDomains: never;
    dnsServers: never;
    dockerSecurityOptions: never;
    extraHosts: never;
    links: never;
    host: never;
    sourcePath: never;
    linuxParameters: never;
    placementConstraints: never;
    privileged: never;

    // Fargate task definitions require that the network mode is set to awsvpc. The awsvpc network
    // mode provides each task with its own elastic network interface
    networkMode: "awsvpc";

    /**
     * Number of cpu units used by the task.  See
     * https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html
     * for more details.
     * 
     * Defaults to 256 if unspecified.
     */
    cpu?: 256 | 512 | 1024 | 2048 | 4096 | string;
  
    /**
     * The amount of memory (in MiB) used by the task. It can be expressed as an integer using MiB,
     * for example 1024, or as a string using GB, for example 1GB or 1 GB, in a task definition.
     * When the task definition is registered, a GB value is converted to an integer indicating the
     * MiB. 
     *
     * See https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html
     * for more details.
     * 
     * Defaults to "512" if unspecified.
     */
    memory?: string;
}>;

export interface FargateServiceArgs extends ServiceArgs {
    cluster: Cluster;
  
    taskDefinition: FargateTaskDefinition;
  
    /**
     * Whether or not awsvpc's networking configuration's [assignPublicIp] is set to "ENABLED" or
     * not.  See https://docs.aws.amazon.com/AmazonECS/latest/developerguide/AWS_Fargate.html for
     * more details.
     *
     * Defaults to false if unspecified.
     */
    assignPublicIp?: boolean;
  
    /**
     * The security group to configure the awsvpc network with.
     */
    subnets?: pulumi.Input<string[]>;
  
    /**
     * The subnets to configure the awsvpc network with.
     */
    securityGroup?: pulumi.Input<string>;
}

class FargateService extends Service {
    public constructor(name: string, args: FargateServiceArgs, opts?: pulumi.ComponentResourceOptions) {
        super("aws:ecs:x:FargateService", name, args, opts);

    }
}