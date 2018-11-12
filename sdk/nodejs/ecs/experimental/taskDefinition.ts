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
export * from "./vpc";


/**
 * Parameteres to control a task definition.  See
 * https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html
 * for more details.
 */
export interface TaskDefinition {
    /**
     * When you register a task definition, you give it a family, which is similar to a name for
     * multiple versions of the task definition, specified with a revision number. The first task
     * definition that is registered into a particular family is given a revision of 1, and any task
     * definitions registered after that are given a sequential revision number. 
     */
    family?: string;
    
    /**
     * When you register a task definition, you can provide a task execution role that allows the
     * containers in the task to pull container images and publish container logs to CloudWatch on
     * your behalf. For more information, see
     * https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_execution_IAM_role.html
     */
    executionRole?: iam.Role;
    
    /**
     * When you register a task definition, you can provide a task role for an IAM role that allows
     * the containers in the task permission to call the AWS APIs that are specified in its
     * associated policies on your behalf. For more information, see IAM Roles for Tasks. 
     */
    taskRole?: iam.Role;

    /**
     * Number of cpu units used by the task.  See
     * https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html
     * for more details.
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
     */
    memoryMiB?: string;

    /**
     * The Docker networking mode to use for the containers in the task. The valid values are none,
     * bridge, awsvpc, and host. 
     * 
     * If the network mode is set to none, the task's containers do not have external connectivity
     * and port mappings can't be specified in the container definition. 
     * 
     * If the network mode is bridge, the task utilizes Docker's built-in virtual network which runs
     * inside each container instance. 
     * 
     * If the network mode is host, the task bypasses Docker's built-in virtual network and maps
     * container ports directly to the EC2 instance's network interface directly. In this mode, you
     * can't run multiple instantiations of the same task on a single container instance when port
     * mappings are used. 
     * 
     * If the network mode is awsvpc, the task is allocated an elastic network interface.
     * 
     * See https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#network_mode
     * for more details.
     * 
     * The default Docker network mode is bridge.
     */
    networkMode?: "none" | "bridge" | "awsvpc" | "host";

    /*
    disableNetworking
    dnsSearchDomains
    dnsServers
    dockerSecurityOptions
    extraHosts
    links
    host and sourcePath
    linuxParameters
    placementConstraints
    privileged
    */
}