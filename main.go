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

func handler(w http.ResponseWriter, r *http.Request) {
	catId := r.URL.Query().Get("catId")

	// TODO gRPCサーバーの情報は環境変数等にする
	conn, err := grpc.Dial("grpc-server:9998", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewCatClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.FindCuteCat(ctx, &pb.FindCuteCatMessage{CatId: catId})
	if err != nil {
		log.Println(err)
	}

	msg := res.GetName() + " is " + res.GetKind() + " 🐱"
	fmt.Fprint(w, msg)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9999", nil)
}
