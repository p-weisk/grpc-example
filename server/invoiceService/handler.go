package invoiceService

import (
	"database/sql"
	"github.com/p-weisk/grpc-example/api"
	"golang.org/x/net/context"
	"log"
	"strconv"
)

// gRPC server
type Server struct {
	Database *sql.DB
}

const FindInvoiceQuery = "SELECT ClientId, P, Number FROM grpc.invoice WHERE Number = ?;"

// Handlers for InvoiceService
func (s *Server) FindInvoiceById(ctx context.Context, in *api.InvoiceNumber) (*api.Invoice, error) {
	log.Printf("Receive message for InvoiceService (FindVoiceById), arg: %s", in.Number)
	i, err := strconv.Atoi(in.Number)
	if err != nil {
		return nil, err
	}
	dbres, dberr := s.Database.Query(FindInvoiceQuery, i)
	if dberr != nil {
		return nil, dberr
	}

	res := api.Invoice{}

	if !dbres.Next() {
		return &res, nil
	}

	var nr int
	var pid string
	scerr := dbres.Scan(&res.ClientId, &pid, &nr)
	if scerr != nil {
		return nil, scerr
	}

	res.P = &api.Product{ProductId:pid}
	res.Number = int32(nr)

	return &res, nil
}

func (s *Server) CreateInvoice(ctx context.Context, in *api.Invoice) (*api.Null, error) {
	log.Printf("Receive message for InvoiceSerivce(CreateInvoice), arg: %+v", *in)
	return &api.Null{}, nil
}
