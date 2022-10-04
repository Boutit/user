package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"

	"github.com/Boutit/user/api"
	"github.com/Boutit/user/internal/config"
	"github.com/Boutit/user/internal/server"
	"github.com/Boutit/user/internal/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	cfg := config.GetConfig(env)
	
	// use net package's Listen method to create a server
	h := cfg.AppConfig.Host
	appPort := cfg.AppConfig.Port
	appValues := []interface{}{h, appPort}
	appConnStr := fmt.Sprintf("%s:%d", appValues...)

	gatewayPort := cfg.AppConfig.GatewayPort
	gatewayValues := []interface{}{h, gatewayPort}
	gatewayConnStr := fmt.Sprintf("%s:%d", gatewayValues...)


	lis, err := net.Listen("tcp", appConnStr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	log.Printf("listening on %s", appConnStr)
	 
	// create a gRPC server
	g := grpc.NewServer()
	reflection.Register(g)

	// Create store
	userStore, err := store.CreatePostgresStore(cfg)
	if err != nil {
		log.Fatalf("failed to connect to user user store: %v", err)
	}
	
	// create a service server
	u := server.NewUserServiceServer(userStore)
	
	// register the service server with the gRPC server
	api.RegisterUserServiceServer(g, u)

	go func() {
		log.Fatalln(g.Serve(lis))
	}()


	// Create client connection to GRPC server that was started
	// Proxy requests using grpc-gateway
	conn, err := grpc.DialContext(
		context.Background(),
		appConnStr,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	err = api.RegisterUserServiceHandler(context.Background(), gwmux, conn)
	
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    gatewayConnStr,
		Handler: gwmux,
	}

	log.Printf("Serving gRPC-Gateway on http:%s", gatewayConnStr)

	log.Fatalln(gwServer.ListenAndServe())

}