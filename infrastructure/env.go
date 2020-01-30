package infrastructure

import (
	"os"
)

type RPCServerEnv struct {
	Host string
	Port string
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

	r := RPCServerEnv{}
	r.Host = host
	r.Port = port

	return r
}
