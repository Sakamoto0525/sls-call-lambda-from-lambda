# sls-call-lambda-from-lambda

VPC①のLambdaから、VPC②のLambdaをVPCエンドポイントを通して呼び出す

呼び出しには AWS SDK for Goの[Invoke](https://docs.aws.amazon.com/sdk-for-go/api/service/lambda/#Lambda.Invoke)を使う

## 動作環境
- Serverless Framework
- go 1.18

## コマンド

```bash
# ビルド
make build

# デプロイ
make deploy

# スタック削除
make remove
```

## 参考

### Go

- https://docs.aws.amazon.com/sdk-for-go/api/service/lambda/#Lambda.Invoke
    - aws-sdk-goのInvoke関数

### Serverless Framework

- https://qiita.com/t_okkan/items/546330b5f4da720c71a7
    - functions.hello.eventsを定義することでAPI Gatewayがトリガーとして作成される
- 

### AWS

- https://www.wantedly.com/companies/toridori/post_articles/312845
    - 良記事。VPC lambdaからvpclambdaを呼ぶcfnが書いてある。ただし、呼び出す側しか書いてないのと、インデントが揃っていないので公式サイトと照らし合わせる必要があるので注意
- https://aws.amazon.com/jp/premiumsupport/knowledge-center/cloudformation-stack-cleanup-stuck/
    - Stack状態が`UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS`から変わらない、かつ削除、更新なにもコンソールからできなかったのだが、コールバック完了もしくはタイムアウトまで待たないといけなかった
- https://blog.serverworks.co.jp/aws-gateways-1
    - Gatewayって多いなと思って調べた。Internet GatewayやNat Gatewayの違いがわかりやすい記事だった
- https://dev.classmethod.jp/articles/dynamodb-vpc-endpoint-lambda/
    - VPC内のLambdaからDynamoDBを叩くときにVPCエンドポイントを作るので、その辺りが参考になる

### エラー

- https://stackoverflow.com/questions/37498124/accessdeniedexception-user-is-not-authorized-to-perform-lambdainvokefunction
    - `lambda:InvokeFunction Error StatusCode403`エラーの対処法。リソースに対して権限を付与しただけで満足していてでたエラー。アカウントにlambda:InvokeFunctionの権限を付与する必要があった。serverless.ymlに書く。
- 




