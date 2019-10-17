package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/p-weisk/grpc-example/api"
	"github.com/p-weisk/grpc-example/server/invoiceService"
	"github.com/p-weisk/grpc-example/server/productService"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {gr
	// create database handler
	db, dsnerr := sql.Open("mysql", "dev:dev@tcp(db:3306)/grpc")
	if dsnerr != nil {
		log.Fatalf("DSN seems invalid: %+v", dsnerr)
	}
	defer db.Close()

	// listen on TCP port 7777
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("can't listen on 7777: %v", err)
	}

	// create service servers
	is := invoiceService.Server{Database:db}
	ps := productService.Server{Database:db}

	// create gRPC server
	grpcServer := grpc.NewServer()

	// attach services
	api.RegisterInvoiceServiceServer(grpcServer, &is)
	api.RegisterProductServiceServer(grpcServer, &ps)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
