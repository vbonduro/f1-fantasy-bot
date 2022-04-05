# Configuration

In order to run this app, you will need:

1. F1 Fantasy user
1. F1 Fantasy password
1. F1 Fantasy league id
1. Slack Bot OAuth Token
1. Slack Verification Token

Run `make configure`, you will be prompted to provide all of the above information. It will be added to your [aws ssm parameter store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html).
The [template.yml](./template.yml) file will add these parameters as environment variables for the AWS Lambda. Note that Environment Variables are encrypted at rest.

# Deployment

This app is built to run in an [AWS Lambda](https://docs.aws.amazon.com/lambda/latest/dg/welcome.html).

In order to deploy this app, you must first create an aws account at https://aws.amazon.com/.

After you have an aws account, follow the aws cli configuration steps [here](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html).

Once that is all done, simply `make deploy`!

> :warning: Using aws services may incur charges! While aws lambdas do have a free tier, be aware of aws' pricing model [here](https://aws.amazon.com/pricing/).
