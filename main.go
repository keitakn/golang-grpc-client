package main

import (
	"context"
	"encoding/json"
	"github.com/keitakn/golang-grpc-client/infrastructure"
	"github.com/keitakn/golang-grpc-server/pkg/pb/cat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	serverEnv := infrastructure.GetRPCServerEnv()
	serverHost := serverEnv.Host + ":" + serverEnv.Port
	authorizationMetadata := "Bearer " + serverEnv.AuthToken

	conn, err := grpc.Dial(serverHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("did not connect: %s", err)
	}
	defer conn.Close()

	client := cat.NewCatClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	catId := r.URL.Query().Get("catId")
	md := metadata.Pairs("authorization", authorizationMetadata)
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.FindCuteCat(ctx, &cat.FindCuteCatMessage{CatId: catId})
	if err != nil {
		log.Println(err)
	}

	msg := res.GetName() + " is " + res.GetKind() + " 🐱"
	type payloadType struct {
		Message string `json:"message"`
	}
	payload := payloadType{Message: msg}
	createJsonResponse(w, http.StatusOK, payload)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	type payloadType struct {
		Message string `json:"message"`
	}
	payload := payloadType{Message: "golang-grpc-client healthCheck is OK🐱"}
	createJsonResponse(w, http.StatusOK, payload)
}

func createJsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	res, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(res))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health-check", healthCheckHandler)
	http.ListenAndServe(":9999", nil)
}
