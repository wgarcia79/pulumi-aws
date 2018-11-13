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

import { sha1hash, Overwrite } from "../../utils";

import { Cluster } from "./cluster";
import { Service, ServiceArgs, TaskDefinitionArgs, TaskDefinition } from "./service"

export type EC2TaskDefinitionArgs = Overwrite<TaskDefinitionArgs, {

}>;

export type EC2ServiceArgs = Overwrite<ServiceArgs, {
    launchType: never;
}>;

export class EC2Service extends Service {
    public constructor(name: string, args: EC2ServiceArgs, opts?: pulumi.ComponentResourceOptions) {
        super("aws.ecs.x.EC2Service", name, args, opts);
        throw new Error("NYI");
    }
}