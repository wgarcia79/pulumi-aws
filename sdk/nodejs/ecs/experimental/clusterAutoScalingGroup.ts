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

import { AutoScalingGroupArgs, Cluster, DefaultAutoScalingGroupArgs, } from "./cluster";

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