package productService

import (
	"database/sql"
	"github.com/p-weisk/grpc-example/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"log"
)

const ProductIdInvalidErrorMsg = "Product Id must not be null"
const UnknownProductErrorMsg = "UnknownProductException"

const SalesVolumeQuery = "SELECT Count(P) -1 + COUNT(Id) from (select Number, ClientId, P, null as Id from invoice where P=? UNION ALL select null, null, null, Id from product where Id=?) as t;"

// gRPC server with Database Handle
type Server struct {
	Database *sql.DB
}

// Handler for ProductService
func (s *Server) GetVolumeOfSales(ctx context.Context, in *api.Product) (*api.SalesVolume, error) {
	log.Printf("Receive message for ProductService (GetVolumeOfSales), arg: %+v", *in)
	if in == nil || in.ProductId == "" {
		return nil, status.Error(codes.InvalidArgument, ProductIdInvalidErrorMsg)
	}

	// number of sales
	var res int

	// the query returns the number of sales, or -1 if the product does not exist
	dberr := s.Database.QueryRow(SalesVolumeQuery, in.ProductId, in.ProductId).Scan(&res)
	if dberr != nil {
		return nil, status.Error(codes.Unknown, dberr.Error())
	}

	// return an error if the product does not exist
	if res < 0 {
		return &api.SalesVolume{
			Volume: 0,
		}, status.Error(codes.NotFound, UnknownProductErrorMsg)
	}

	return &api.SalesVolume{
		Volume: float32(res),
	}, nil
}