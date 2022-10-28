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
