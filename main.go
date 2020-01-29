package main

import (
	"context"
	"fmt"
	pb "github.com/keitakn/golang-grpc-server/pb"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

// TODO 後で構造化する
func main() {
	// TODO gRPCサーバーの情報は環境変数等にする
	// TODO gRPCサーバーの生成処理は分離する
	connection, err := grpc.Dial("grpc-server:9998", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("did not connect: %s", err)
	}
	defer connection.Close()

	client := pb.NewCatClient(connection)

	context, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.FindCuteCat(context, &pb.FindCuteCatMessage{CatId: "mop"})
	if err != nil {
		log.Println(err)
	}

	msg := response.GetName() + " is " + response.GetKind() + " 🐱"
	http.Handle("/", String(msg))
	http.ListenAndServe(":9999", nil)
}
