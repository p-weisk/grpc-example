package productService

import (
	"database/sql"
	"github.com/p-weisk/avg-grpc/api"
	"golang.org/x/net/context"
	"log"
)

const SalesVolumeQuery = "SELECT COUNT(*) FROM grpc.invoice WHERE P = ?;"

// gRPC server
type Server struct {
	Database *sql.DB
}

// Handler for ProductService
func (s *Server) GetVolumeOfSales(ctx context.Context, in *api.Product) (*api.SalesVolume, error) {
	log.Printf("Receive message for ProductService (GetVolumeOfSales), arg: %+v", *in)

	dbres, dberr := s.Database.Query(SalesVolumeQuery, in.ProductId)
	if dberr != nil {
		return nil, dberr
	}
	if !dbres.Next() {
		return nil, nil
	}

	var res int

	scerr := dbres.Scan(&res)
	if scerr != nil {
		return nil, scerr
	}

	return &api.SalesVolume{
		Volume: float32(res),
	}, nil
}