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

import * as aws from "../..";
import * as pulumi from "@pulumi/pulumi";
import * as ec2 from "../../ec2";

import { Vpc } from "./vpc";

export interface ClusterArgs {
    /**
     * VPC for the cluster.
     */
    vpc: Vpc;
}

export interface DefaultAutoScalingGroupArgs {
    /** Types of instances for the AutoScalingGroup */
    instanceType: ec2.InstanceType;

    /** Number of instances to place in the cluster. Defaults to 1 if unspecified. */
    instanceCount?: number;
}

export interface AutoScalingGroupArgs {
    /** The size of the instance to launch */
    instanceType: pulumi.Input<ec2.InstanceType>

    /**
     * The name of the ECS-optimzed AMI to use for the Container Instances in this cluster, e.g.
     * "amzn-ami-2017.09.l-amazon-ecs-optimized". Defaults to using the latest recommended ECS Optimized AMI, which may
     * change over time and cause recreation of EC2 instances when new versions are release. To control when these
     * changes are adopted, set this parameter explicitly to the version you would like to use.
     *
     * See http://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html for valid values.
     */
    machineImageId?: pulumi.Input<string>; 

    /** 
     * How to update instances in this auto scaling group.  See
     * https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html
     * for more details.
     *
     * "willReplace": 
     * Specifies whether an Auto Scaling group and the instances it contains are
     * replaced during an update. During replacement, AWS CloudFormation retains the old group until
     * it finishes creating the new one. If the update fails, AWS CloudFormation can roll back to
     * the old Auto Scaling group and delete the new Auto Scaling group. While AWS CloudFormation
     * creates the new group, it doesn't detach or attach any instances. After successfully creating
     * the new Auto Scaling group, AWS CloudFormation deletes the old Auto Scaling group during the
     * cleanup process. 
     * 
     * "rollingUpdate": 
     * Rolling updates enable you to specify whether AWS CloudFormation updates
     * instances that are in an Auto Scaling group in batches or all at once. 
     *  
     * Defaults to [none] if not specified
     */
    updateType?: "none" | "willReplace" | "rollingUpdate";
    
    /** 
     * The minimum number of instances that can run in your auto scale group. If your scale down
     * CloudWatch alarm is triggered, your auto scale group will never terminate instances below
     * this number.
     * 
     * Defaults to 1 if unspecified.
     */
    minSize?: number;

    /** 
     * The maximum number of instances that you can run in your auto scale group. If your scale up
     * CloudWatch alarm stays triggered, your auto scale group will never create instances more than
     * the maximum amount specified.
     * 
     * Defaults to 1 if unspecified.
     */
    maxSize?: number;


    /**
     * The desired amount of capcity in the group. If you trip a CloudWatch alarm for a scale up
     * event, then it will notify the auto scaler to change it's desired to a specified higher
     * amount and the auto scaler will start an instance/s to meet that number. If you trip a
     * CloudWatch alarm to scale down, then it will change the auto scaler desired to a specified
     * lower number and the auto scaler will terminate instance/s to get to that number. 
     * 
     * Defaults to 1 if unspecified.
     */
    desiredCapacity?: number;
}

export class Cluster extends pulumi.ComponentResource {
    public readonly resource: aws.ecs.Cluster;
    public readonly vpc: Vpc;

    public constructor(name: string, args: ClusterArgs, opts?: pulumi.ComponentResourceOptions) {
        super("aws:ecs:x:Cluster", name, args, opts);
        this.resource = new aws.ecs.Cluster(name, {}, { parent: this });

        this.registerOutputs({
            resource: this.resource,
        })
    }

    public createDefaultAutoScalingGroup(name: string, args: DefaultAutoScalingGroupArgs): ClusterAutoScalingGroup {
        return ClusterAutoScalingGroup.createDefault(name,  {
            ...args,
            cluster: this,
        }, { parent: this });
    }

    public createAutoScalingGroup(name: string, args: AutoScalingGroupArgs): ClusterAutoScalingGroup {
        return new ClusterAutoScalingGroup(name, {
            ...args,
            cluster: this,
        }, { parent: this });
    }
}

export interface DefaultClusterAutoScalingGroupArgs extends DefaultAutoScalingGroupArgs {
    cluster: Cluster;
}

export interface ClusterAutoScalingGroupArgs extends AutoScalingGroupArgs {
    cluster: Cluster,
}

export class ClusterAutoScalingGroup extends pulumi.ComponentResource {
    public constructor(name: string, args: ClusterAutoScalingGroupArgs, opts?: pulumi.CustomResourceOptions) {
        super("aws:ecs:x:ClusterAutoScalingGroup", name, args, opts)

        throw new Error("NYI");
    }

    public static createDefault(name: string, args: DefaultClusterAutoScalingGroupArgs, opts?: pulumi.CustomResourceOptions) {
        return new ClusterAutoScalingGroup(name, {
            cluster: args.cluster,
            instanceType: args.instanceType,
            updateType: "willReplace",
            minSize: 0,
            maxSize: args.instanceCount !== undefined ? args.instanceCount : 1,
            desiredCapacity: args.instanceCount !== undefined ? args.instanceCount : 1,
        }, opts);
    }
}