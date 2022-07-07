package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	position_gw "github.com/cory-evans/Ichnaea/pkg/proto/position"
	users_gw "github.com/cory-evans/Ichnaea/pkg/proto/users"
)

var (
	listenPort                  = os.Getenv("LISTEN_PORT")
	positionGrpcServiceEndpoint = os.Getenv("POSITION_SERVICE")
	userGrpcServiceEndpoint     = os.Getenv("USER_SERVICE")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	positionMux := runtime.NewServeMux()
	err := position_gw.RegisterPositionServiceHandlerFromEndpoint(ctx, positionMux, positionGrpcServiceEndpoint, opts)
	if err != nil {
		return err
	}

	userMux := runtime.NewServeMux()
	err = users_gw.RegisterUserServiceHandlerFromEndpoint(ctx, userMux, userGrpcServiceEndpoint, opts)
	if err != nil {
		return err
	}

	gwMux := http.NewServeMux()
	gwMux.Handle("/position/", http.StripPrefix("/position", positionMux))
	gwMux.Handle("/users/", http.StripPrefix("/users", positionMux))

	srv := &http.Server{
		Addr:    ":" + listenPort,
		Handler: cors(gwMux),
	}

	return srv.ListenAndServe()
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
