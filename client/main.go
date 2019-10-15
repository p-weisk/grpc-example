package main

import (
	"github.com/p-weisk/grpc-example/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var c *grpc.ClientConn
	c, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection failed: %s", err)
	}
	defer c.Close()

	ic := api.NewInvoiceServiceClient(c)

	// Invoice Service request #1
	response1, err := ic.FindInvoiceById(
		context.Background(),
		&api.InvoiceNumber{Number: "2"},
	)
	if err != nil {
		log.Fatalf("Error when calling InvoiceService.findInvoiceById: %s", err)
	}

	log.Printf("Response from InvoiceService.findInvoiceById: %s", response1.String())

	// Invoice Service request #2
	response2, err := ic.CreateInvoice(
		context.Background(),
		&api.Invoice{
			ClientId: "Client-0000",
			P: &api.Product{ProductId: "Product-0000"},
			Number: 10,
		},
	)
	if err != nil {
		log.Fatalf("Error when calling InvoiceService.createInvoice: %s", err)
	}

	log.Printf("Response from InvoiceService.createInvoice: %s", response2.String())

	// ProductService request
	pc := api.NewProductServiceClient(c)

	response3, err := pc.GetVolumeOfSales(context.Background(), &api.Product{ProductId: "Product-0000"})
	if err != nil {
		log.Fatalf("Error when calling ProductService.getVolumeOfSales: %s", err)
	}

	log.Printf("Response from ProductService.getVolumeOfSales: %s", response3.String())
}
