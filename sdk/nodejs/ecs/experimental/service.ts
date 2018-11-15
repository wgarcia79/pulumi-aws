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

import * as pulumi from "@pulumi/pulumi";
import * as iam from "../../iam";
import * as ecs from "..";

import { sha1hash, Overwrite } from "../../utils";

import { Cluster } from "./cluster";

/**
 * Parameteres to control a task definition.  See
 * https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html
 * for more details.
 */
export type TaskDefinitionArgs = Overwrite<ecs.TaskDefinitionArgs, {
    /**
     * Not used.  Provide [executionRole] instead.
     */
    executionRoleArn?: never;

    /**
     * When you register a task definition, you can provide a task execution role that allows the
     * containers in the task to pull container images and publish container logs to CloudWatch on
     * your behalf. For more information, see
     * https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_execution_IAM_role.html
     */
    executionRole?: iam.Role;
    
    /**
     * Not used.  Provide [taskRole] instead.
     */
    taskRoleArn?: never;

    /**
     * When you register a task definition, you can provide a task role for an IAM role that allows
     * the containers in the task permission to call the AWS APIs that are specified in its
     * associated policies on your behalf. For more information, see IAM Roles for Tasks. 
     */
    taskRole?: iam.Role;

    containerDefinitions: ecs.ContainerDefinition[];

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

    // todo: add support for this.
    volumes?: never;
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

   // requiresCompatibilities: 
}>;

export type ServiceArgs = Overwrite<ecs.ServiceArgs, {
    cluster: Cluster;

    /**
     * 
     */
    taskDefinition: TaskDefinitionArgs;

    // loadBalancers?: pulumi.Input<pulumi.Input<ServiceLoadBalancer>[]>;
    // serviceRegistries?: pulumi.Input<pulumi.Input<ServiceRegistry>[]>;

    /**
     * The launch type on which to run your service.
     */
    launchType: pulumi.Input<"EC2" | "FARGATE">,

    /**
     * The upper limit (as a percentage of the service's desiredCount) of the number of running
     * tasks that can be running in a service during a deployment. Not valid when using the `DAEMON`
     * scheduling strategy.
     * 
     * Defaults to 200 if not specified.
     */
    deploymentMaximumPercent?: pulumi.Input<number>;

    /**
     * The lower limit (as a percentage of the service's desiredCount) of the number of running
     * tasks that must remain running and healthy in a service during a deployment.
     * 
     * Default to 50 if not specified.
     */
    deploymentMinimumHealthyPercent?: pulumi.Input<number>;
}>;

const executionAndTaskRolePolicy = {
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "sts:AssumeRole",
            "Principal": {
                "Service": "ecs-tasks.amazonaws.com",
            },
            "Effect": "Allow",
            "Sid": "",
        },
    ],
};

export class TaskDefinition extends pulumi.ComponentResource {
    public readonly executionRole: iam.Role;
    public readonly taskRole: iam.Role;
    public readonly resource: ecs.TaskDefinition;

    public constructor(name: string, args: TaskDefinitionArgs, opts?: pulumi.ComponentResourceOptions) {
        super("aws.ec2.x.TaskDefinition", name, args, opts);

        this.executionRole = args.executionRole || new iam.Role(name, {
                assumeRolePolicy: JSON.stringify(executionAndTaskRolePolicy),
            }, { parent: this });

        this.taskRole = args.taskRole || new iam.Role(name, {
            assumeRolePolicy: JSON.stringify(executionAndTaskRolePolicy),
        }, { parent: this });

        const argsCopy = {
            ...args,
            containerDefinitions: convertContainerDefinitions(args.containerDefinitions),
            executionRoleArn: this.executionRole.arn,
            taskRoleArn: this.taskRole.arn,
            cpu: "" + args.cpu,
            memory: "" + args.memory,
        };

        delete argsCopy.executionRole;
        delete argsCopy.taskRoleArn;

        this.resource = new ecs.TaskDefinition(name, argsCopy, { parent: this });

        this.registerOutputs({
            executionRole: this.executionRole,
            taskRole: this.taskRole,
            resource: this.resource,
        });
    }
}

function convertContainerDefinitions(definitions: ecs.ContainerDefinition[]) {
    return JSON.stringify(definitions);
}

export class Service extends pulumi.ComponentResource {
    public readonly taskDefinition: TaskDefinition;
    public readonly resource: ecs.Service;

    public constructor(type: string, name: string, args: ServiceArgs, opts?: pulumi.ComponentResourceOptions) {
        super(type, name, args, opts);

        this.taskDefinition = new TaskDefinition(name, args.taskDefinition, { parent: this });

        const argsCopy = {
            ...args,
            cluster: args.cluster.resource.arn,
            taskDefinition: this.taskDefinition.resource.arn,
            deploymentMaximumPercent: args.deploymentMaximumPercent !== undefined ? args.deploymentMaximumPercent : 200,
            deploymentMinimumHealthyPercent: args.deploymentMinimumHealthyPercent !== undefined ? args.deploymentMinimumHealthyPercent : 50,
        };

        this.resource = new ecs.Service(name, argsCopy, { parent: this });

        this.registerOutputs({
            taskDefinition: this.taskDefinition,
            resource: this.resource,
        })
    }
}