package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/p-weisk/grpc-example/api"
	"github.com/p-weisk/grpc-example/server/invoiceService"
	"github.com/p-weisk/grpc-example/server/productService"
	"google.golang.org/grpc"
)

func main() {
	// create database handle
	db, dsnerr := sql.Open("mysql", "dev:dev@tcp(db:3306)/grpc")
	if dsnerr != nil {
		log.Fatalf("DSN seems invalid: %+v", dsnerr)
	}
	defer db.Close()

	// listen on TCP 8888
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 8888))
	if err != nil {
		log.Fatalf("cannot listen on 8888: %v", err)
	}

	// create service servers
	is := invoiceService.Server{Database: db}
	ps := productService.Server{Database: db}

	// create gRPC server
	grpcServer := grpc.NewServer()

	// register services
	api.RegisterInvoiceServiceServer(grpcServer, &is)
	api.RegisterProductServiceServer(grpcServer, &ps)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
