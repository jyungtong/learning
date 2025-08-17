/// <reference path="./.sst/platform/config.d.ts" />

export default $config({
  app(input) {
    return {
      name: "learn-hono",
      removal: input?.stage === "production" ? "retain" : "remove",
      protect: ["production"].includes(input?.stage),
      home: "aws",
    };
  },
  async run() {
    const router = new sst.aws.Router("app_router", {
      domain: {
        name: "jy.example.com",
        aliases: ["*.jy.example.com"]
      }
    });

    const app = new sst.aws.Function('hono_app', {
      handler: 'src/lambda.handler',
      url: {
        router: {
          instance: router,
          path: "/",
          domain: "api.jy.example.com",
        }
      }
    });
  },
});
