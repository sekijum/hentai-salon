### terraform で管理していないリソース

- acm
- route53 ホストゾーン
- ecr リポジトリ

### terraform(v1.5.3)

**install** [公式ドキュメント](https://developer.hashicorp.com/terraform/downloads)
[tfenv](https://github.com/tfutils/tfenv)を使うとバージョンの切り替えが容易

### awscli(v2)

tfvars を使えば awscli は不要だが S3 に格納している tfstate を参照する時権限エラーになるため一旦クレデンシャルで代用
**install** [公式ドキュメント](https://docs.aws.amazon.com/ja_jp/cli/latest/userguide/getting-started-install.html)

**プロファイル設定**

```
# ~/.aws/credentials
[vrsales]
aws_access_key_id=<アクセスキー>
aws_secret_access_key=<シークレットキー>
region = ap-northeast-1
```

**環境変数の設定の手順**

- aws コンソールにて ssm パラメータ作成
  - ```bash
    # 命名規則
    ## 環境変数のprefixとして`/VR-SALES/STG/`を付与する
    ### ステージング環境 (例) /VR-SALES/STG/HOGE_HOGE
    ```
  - 作成後`terraform/modules/ecs/ssm_parameter.tf`に datasource を追加
  - `terraform/modules/ecs/ecs_task_definition.tf` の該当コンテナに(datasource を参照するように)追加
  - apply する

### スケジューリングされたリリース

ecs を[Application Auto Scaling](https://docs.aws.amazon.com/ja_jp/autoscaling/application/userguide/what-is-application-auto-scaling.html)でスケジューリングしているがコンソール画面がない？ので[cli](https://docs.aws.amazon.com/ja_jp/autoscaling/application/userguide/scheduled-scaling-additional-cli-commands.html)で確認するしかなさそう。

### 運用上の注意点

プロファイル名を指定しているがセットされないことがあるため以下のコマンド打つ

```
$ export AWS_PROFILE=xxx
```

リソース変更を行っていないにも関わらず `terraform plan` で差分が発生する事がある。
terraform 以外の手段でリソースが変更された場合は差分が出るようになっている為。

```
No changes. Your infrastructure matches the configuration.

Your configuration already matches the changes detected above. If you'd like to update the Terraform state to match, create and apply a refresh-only plan:
  terraform apply -refresh-only
```

このように `No changes. Your infrastructure matches the configuration.` が表示されている場合は `terraform apply -refresh-only` を実行。
