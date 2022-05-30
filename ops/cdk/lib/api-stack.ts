import { Duration, Stack, StackProps } from "aws-cdk-lib";
import { LambdaRestApi } from "aws-cdk-lib/aws-apigateway";
import { Code, Function, Runtime } from "aws-cdk-lib/aws-lambda";
import { Construct } from "constructs";

const LAMBDA_RUNTIME = Runtime.GO_1_X;
const LAMBDA_DEFAULT_TIMEOUT = Duration.seconds(15);
const LAMBDA_DEFAULT_MEMORY = 512; // megabytes
const CODE_DIST = "../../dist";

export class ApiStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    const backend = new Function(this, "ApiLambdaBackend", {
      runtime: LAMBDA_RUNTIME,
      handler: "api",
      code: Code.fromAsset(CODE_DIST),
      timeout: LAMBDA_DEFAULT_TIMEOUT,
      memorySize: LAMBDA_DEFAULT_MEMORY,
    });

    // proxy requests to lambda by default
    new LambdaRestApi(this, "Api", {
      handler: backend,
    });
  }
}
