# https://docs.aws.amazon.com/cli/latest/reference/ecs/describe-tasks.html
# https://docs.aws.amazon.com/cli/latest/reference/ecs/list-tasks.html
resource "local_file" "rds_tunnel" {
  filename = "../../cmd/${local.env}/rds_tunnel.sh"
  content  = <<DOC
#!/bin/bash

set -eu

CUR_DIR=$(pwd)

TUNNEL_LOG_DIR_PATH=$CUR_DIR"/cmd/logs/rds_tunnel"
TUNNEL_LOG_FILE_PATH=$TUNNEL_LOG_DIR_PATH"/$(date +%Y_%m%d_%H%M%S).log"

TASK_ID=`aws ecs list-tasks \
  --cluster ${local.app_name} \
  --service-name ${aws_ecs_service.server.name} \
  --desired-status RUNNING \
  --profile ${local.aws_profile} \
  | jq '.taskArns[0]' \
  | sed 's/"//g' \
  | cut -f 3 -d '/'`

DESCRIBE_TASKS=`aws ecs describe-tasks \
  --cluster ${aws_ecs_cluster.this.name} \
  --task $TASK_ID \
  --profile ${local.aws_profile}`

for container in $(echo $DESCRIBE_TASKS | jq -c '.tasks[0].containers[]'); do
  CONTAINER_NAME=$(echo $container | jq -c ".name" | sed 's/"//g')
  if [ $CONTAINER_NAME = "nest" ]; then
    RUNTIME_ID=$(echo $container | jq -c ".runtimeId" | sed 's/"//g')
    break
  fi
done

mkdir -p $TUNNEL_LOG_DIR_PATH
touch $TUNNEL_LOG_FILE_PATH

echo "30分経つと接続が切れます。"

aws ssm start-session \
  --target "ecs:${aws_ecs_cluster.this.name}_"$TASK_ID"_"$RUNTIME_ID \
  --document-name AWS-StartPortForwardingSessionToRemoteHost \
  --parameters '{"host":["${aws_rds_cluster.this.endpoint}"],"portNumber":["3306"], "localPortNumber":["13306"]}' \
  --profile ${local.aws_profile} > $TUNNEL_LOG_FILE_PATH

# 1日経ったログは消す
find $TUNNEL_LOG_DIR_PATH -type f -mtime +1 | xargs rm
DOC
}

resource "local_file" "adminer_tunnel" {
  filename = "../../cmd/${local.env}/adminer_tunnel.sh"
  content  = <<DOC
#!/bin/bash

set -eu

CUR_DIR=$(pwd)

TUNNEL_LOG_DIR_PATH=$CUR_DIR"/cmd/.logs/adminer_tunnel"
TUNNEL_LOG_FILE_PATH=$TUNNEL_LOG_DIR_PATH"/$(date +%Y_%m%d_%H%M%S).log"

TASK_ARN=$(aws ecs run-task \
  --cluster ${local.app_name} \
  --network-configuration "awsvpcConfiguration={subnets=[${aws_subnet.public_1a.id}],securityGroups=[${aws_security_group.adminer.id}],assignPublicIp=ENABLED}" \
  --profile ${local.aws_profile} \
  | jq -r '.tasks[0].taskArn')

TASK_ID=$(echo $TASK_ARN | cut -f 3 -d '/')

RUNTIME_ID=$(aws ecs describe-tasks \
  --cluster ${aws_ecs_cluster.this.name} \
  --tasks $TASK_ID \
  --profile ${local.aws_profile} \
  | jq -r '.tasks[0].containers[0].runtimeId')

echo "Adminer タスクが起動されました: $TASK_ARN"

mkdir -p $TUNNEL_LOG_DIR_PATH
touch $TUNNEL_LOG_FILE_PATH

echo "30分経つと接続が切れます。"

aws ssm start-session \
  --target "ecs:${aws_ecs_cluster.this.name}_${TASK_ID}_${RUNTIME_ID}" \
  --document-name AWS-StartPortForwardingSessionToRemoteHost \
  --parameters '{"host":["localhost"],"portNumber":["8080"], "localPortNumber":["18080"]}' \
  --profile ${local.aws_profile} > $TUNNEL_LOG_FILE_PATH

# 1日経ったログは消す
find $TUNNEL_LOG_DIR_PATH -type f -mtime +1 | xargs rm
DOC
}

resource "local_file" "adminer_tunnel" {
  filename = "../../cmd/${local.env}/rds_tunnel.sh"
  content  = <<DOC
#!/bin/bash

set -eu

CUR_DIR=$(pwd)

TUNNEL_LOG_DIR_PATH=$CUR_DIR"/logs/rds_tunnel"
TUNNEL_LOG_FILE_PATH=$TUNNEL_LOG_DIR_PATH"/$(date +%Y_%m%d_%H%M%S).log"

TASK_ID=`aws ecs list-tasks \
  --cluster ${local.app_name} \
  --service-name ${aws_ecs_service.server.name} \
  --desired-status RUNNING \
  --profile ${local.aws_profile} \
  | jq '.taskArns[0]' \
  | sed 's/"//g' \
  | cut -f 3 -d '/'`

DESCRIBE_TASKS=`aws ecs describe-tasks \
  --cluster ${aws_ecs_cluster.this.name} \
  --task $TASK_ID \
  --profile ${local.aws_profile}`

for container in $(echo $DESCRIBE_TASKS | jq -c '.tasks[0].containers[]'); do
  CONTAINER_NAME=$(echo $container | jq -c ".name" | sed 's/"//g')
  if [ $CONTAINER_NAME = "nest" ]; then
    RUNTIME_ID=$(echo $container | jq -c ".runtimeId" | sed 's/"//g')
    break
  fi
done

mkdir -p $TUNNEL_LOG_DIR_PATH
touch $TUNNEL_LOG_FILE_PATH

echo "30分経つと接続が切れます。"

aws ssm start-session \
  --target "ecs:${aws_ecs_cluster.this.name}_"$TASK_ID"_"$RUNTIME_ID \
  --document-name AWS-StartPortForwardingSessionToRemoteHost \
  --parameters '{"host":["${aws_rds_cluster.this.endpoint}"],"portNumber":["3306"], "localPortNumber":["13306"]}' \
  --profile ${local.aws_profile} > $TUNNEL_LOG_FILE_PATH

# 1日経ったログは消す
find $TUNNEL_LOG_DIR_PATH -type f -mtime +1 | xargs rm
DOC
}

resource "local_file" "server_exec" {
  filename = "../../../cmd/${local.env}/server_exec.sh"
  content  = <<DOC
#!/bin/bash

TASK_ID=`aws ecs list-tasks \
  --cluster ${local.app_name} \
  --service-name ${aws_ecs_service.server.name} \
  --desired-status RUNNING \
  --profile ${local.aws_profile} \
  | jq '.taskArns[0]' \
  | sed 's/"//g' \
  | cut -f 3 -d '/'`

aws ecs execute-command \
  --cluster ${aws_ecs_service.server.name} \
  --container ${aws_ecs_task_definition.server.arn} \
  --task $TASK_ID \
  --interactive \
  --command "/bin/bash" \
  --profile ${local.aws_profile}
DOC
}
