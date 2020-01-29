# golang-grpc-client
GoでgRPCを実装する時のサンプルコード

https://github.com/keitakn/golang-grpc-server と連携を行います。

## 起動

https://github.com/keitakn/golang-grpc-server と同じDockerネットワークに所属するようになっています。

先に https://github.com/keitakn/golang-grpc-server を起動して下さい。

### 通常の起動方法

```
./docker-compose-up.sh
```

### デバッガーによるデバッグを行う場合

```
./docker-compose-up-debug.sh
```

## ソースコードのフォーマット

プロジェクトルートで以下を実行して下さい。

`gofmt -l -s -w .`
