import { Duration, Stack, StackProps } from "aws-cdk-lib";
import { LambdaIntegration } from "aws-cdk-lib/aws-apigateway";
import { Code, Function, Runtime } from "aws-cdk-lib/aws-lambda";
import { Construct } from "constructs";

const LAMBDA_RUNTIME = Runtime.GO_1_X;
const LAMBDA_DEFAULT_TIMEOUT = Duration.seconds(15);
const LAMBDA_DEFAULT_MEMORY = 512; // megabytes
const CODE_DIST = "../../dist";
const OPENAPI_SPEC = "../../openapi/api.yaml";

export class ApiStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    const backendLambda = new Function(this, "LambdaBackend", {
      runtime: LAMBDA_RUNTIME,
      functionName: "GoRestApiLambdaBackend", // this is important to reference in openapi spec
      handler: "api",
      code: Code.fromAsset(CODE_DIST),
      timeout: LAMBDA_DEFAULT_TIMEOUT,
      memorySize: LAMBDA_DEFAULT_MEMORY,
    });

    const backendLambdaIntegration = new LambdaIntegration(backendLambda);

    // const apigw = new SpecRestApi(this, "RestApi", {
    //   apiDefinition: ApiDefinition.fromAsset(OPENAPI_SPEC)
    // });

    // apigw.root.addProxy({
    //   defaultIntegration: new LambdaIntegration(backendLambda),
    // })
  }
}
