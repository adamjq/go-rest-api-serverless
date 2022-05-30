#!/usr/bin/env node
import { App } from "aws-cdk-lib";
import "source-map-support/register";
import { ApiStack } from "../lib/api-stack";

const app = new App();
new ApiStack(app, "GoRestApiServerlessStack", {});
