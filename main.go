package main

import (
	"context"
	"fmt"
	"github.com/keitakn/golang-grpc-client/infrastructure"
	pb "github.com/keitakn/golang-grpc-server/pb"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	serverEnv := infrastructure.GetRPCServerEnv()
	serverHost := serverEnv.Host + ":" + serverEnv.Port

	conn, err := grpc.Dial(serverHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewCatClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	catId := r.URL.Query().Get("catId")
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
