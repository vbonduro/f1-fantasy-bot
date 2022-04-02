#!/bin/sh

STORE_PARAMETER="aws ssm put-parameter --no-cli-pager --overwrite"

echo -n "F1 Fantasy User: "
read F1_USER
F1_PASSWORD=`systemd-ask-password "F1 Fantasy Password: "`
echo -n "F1 Fantasy League ID: "
read LEAGUE_ID
SLACK_OAUTH=`systemd-ask-password "Slack OAuth Token: "`
SLACK_VERIFICATION_TOKEN=`systemd-ask-password "Slack Verification Token: "`
echo "Sending parameters to aws..."
${STORE_PARAMETER} --name "f1_user" --type "String" --value "${F1_USER}" > /dev/null
${STORE_PARAMETER} --name "f1_password" --type "String" --value "${F1_PASSWORD}" > /dev/null
${STORE_PARAMETER} --name "f1_league" --type "String" --value "${LEAGUE_ID}" > /dev/null
${STORE_PARAMETER} --name "slack_oauth" --type "String" --value "${SLACK_OAUTH}" > /dev/null
${STORE_PARAMETER} --name "slack_verification_token" --type "String" --value "${SLACK_VERIFICATION_TOKEN}" > /dev/null
echo "✅ Done!"
