package invoiceService

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/p-weisk/grpc-example/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gRPC server with Database Handle
type Server struct {
	Database *sql.DB
}

const MySQLErrorForeignKeyConstraint = 1216

const UnknownProductErrorMsg = "UnknownProductException"
const UnknownClientErrorMsg = "UnknownClientException"
const EmptyParameterErrorMsg = "Parameters must not be empty."
const NotIntegerParameterErrorMsg = "Invoice Number must be convertable to Integer"

const FindInvoiceQuery = "SELECT ClientId, P, Number FROM grpc.invoice WHERE Number = ?;"
const CreateInvoiceQuery = "INSERT INTO grpc.invoice(ClientId, P, Number) VALUES (?,?,?);"
const CountProductQuery = "SELECT COUNT(*) FROM grpc.product WHERE Id = ?;"

// Handlers for InvoiceService
func (s *Server) FindInvoiceById(ctx context.Context, in *api.InvoiceNumber) (*api.Invoice, error) {
	log.Printf("Receive message for InvoiceService (FindVoiceById), arg: %s", in.Number)
	i, err := strconv.Atoi(in.Number)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, NotIntegerParameterErrorMsg)
	}

	res := api.Invoice{}
	var nr int
	var pid string
	dberr := s.Database.QueryRow(FindInvoiceQuery, i).Scan(&res.ClientId, &pid, &nr)
	if dberr == sql.ErrNoRows {
		return &res, nil
	}
	// unspecific error
	if dberr != nil {
		return nil, status.Error(codes.Unknown, dberr.Error())
	}

	res.P = &api.Product{ProductId: pid}
	res.Number = int32(nr)

	return &res, nil
}

func (s *Server) CreateInvoice(ctx context.Context, in *api.Invoice) (*api.Null, error) {
	log.Printf("Receive message for InvoiceSerivce(CreateInvoice), arg: %+v", *in)
	if in.ClientId == "" || in.P == nil || in.P.ProductId == "" {
		return &api.Null{}, status.Error(codes.InvalidArgument, EmptyParameterErrorMsg)
	}

	_, dberr := s.Database.Exec(CreateInvoiceQuery, in.ClientId, in.P.ProductId, in.Number)
	// mySQL returns an error code
	if dberr, ok := dberr.(*mysql.MySQLError); ok {
		log.Printf("MySQL Error %d , message: %s", dberr.Number, dberr.Error())
		if dberr.Number == MySQLErrorForeignKeyConstraint {
			// foreign key constraint failed - either productor client does not exist

			var pcount int
			s.Database.QueryRow(CountProductQuery, in.P.ProductId).Scan(&pcount)
			log.Printf("product count: %d", pcount)
			// product does exist in products table
			if pcount != 0 {
				return &api.Null{}, status.Error(codes.NotFound, UnknownClientErrorMsg)
			}
			// product does not exist in products table
			return &api.Null{}, status.Error(codes.NotFound, UnknownProductErrorMsg)
		}
	}
	// unspecific error
	if dberr != nil {
		return &api.Null{}, status.Error(codes.Unknown, dberr.Error())
	}
	return &api.Null{}, nil
}
