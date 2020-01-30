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

## 動作確認方法

以下で動作確認が出来ます。

`curl -v http://localhost:9999\?catId\=moko`

下記のようなレスポンスが返ってくれば正常動作しています。

```
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 9999 (#0)
> GET /?catId=moko HTTP/1.1
> Host: localhost:9999
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Thu, 30 Jan 2020 02:29:11 GMT
< Content-Length: 30
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
Moko is Chinchilla silver 🐱* Closing connection 0
```

catIdに関しては [こちら](https://github.com/keitakn/golang-grpc-server/blob/5df29d624f4046e2a3dac636fad83ff7503376dd/infrastructure/gRPCService.go#L12) に定義されている値であれば何でも構いません。
