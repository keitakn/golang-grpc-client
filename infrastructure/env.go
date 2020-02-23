package infrastructure

import (
	"os"
)

type RPCServerEnv struct {
	Host      string
	Port      string
	AuthToken string
}

func GetRPCServerEnv() RPCServerEnv {
	host := os.Getenv("GRPC_HOST")
	if host == "" {
		host = "grpc-server"
	}

	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "9998"
	}

	authToken := os.Getenv("AUTH_TOKEN")
	if authToken == "" {
		// https://github.com/keitakn/golang-grpc-server に書いてあるテストデータをデフォルト値として利用
		// 本来このようなデータはハードコードすべきではない
		authToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.rTCH8cLoGxAm_xw68z-zXVKi9ie6xJn9tnVWjd_9ftE"
	}

	r := RPCServerEnv{}
	r.Host = host
	r.Port = port
	r.AuthToken = authToken

	return r
}
