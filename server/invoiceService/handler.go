package invoiceService

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/p-weisk/grpc-example/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"log"
	"strconv"
	"strings"
)

// gRPC server
type Server struct {
	Database *sql.DB
}

const MySQLErrorForeignKeyConstraint = 145

const UnknownProductErrorMsg = "UnknownProductException"
const UnknownClientErrorMsg = "UnknownClientException"
const EmptyParameterErrorMsg = "Parameters must not be empty."
const NotIntegerParameterErrorMsg = "Invoice Number must be convertable to Integer"

const FindInvoiceQuery = "SELECT ClientId, P, Number FROM grpc.invoice WHERE Number = ?;"
const CreateInvoiceQuery = "INSERT INTO grpc.invoice(ClientId, P, Number) VALUES (?,?,?);"

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
	if dberr, ok := dberr.(*mysql.MySQLError); ok {
		// mySQL returns an error code
		if dberr.Number == 145 { //foreign key constraint failed
			// product does not exist in products table
			if strings.Contains(dberr.Message, "product") {
				return nil, status.Error(codes.NotFound, UnknownProductErrorMsg)
			}
			//client does not exist in clients table
			if strings.Contains(dberr.Message, "client") {
				return nil, status.Error(codes.NotFound, UnknownClientErrorMsg)
			}
		}
	}
	// unspecific error
	if dberr != nil {
		return nil, status.Error(codes.Unknown, dberr.Error())
	}

	res.P = &api.Product{ProductId:pid}
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
		if dberr.Number == MySQLErrorForeignKeyConstraint { //foreign key constraint failed
			// product does not exist in products table
			if strings.Contains(dberr.Message, "product") {
				return &api.Null{}, status.Error(codes.NotFound, UnknownProductErrorMsg)
			}
			//client does not exist in clients table
			if strings.Contains(dberr.Message, "client") {
				return &api.Null{}, status.Error(codes.NotFound, UnknownClientErrorMsg)
			}
		}
	}
	// unspecific error
	if dberr != nil {
		return &api.Null{}, status.Error(codes.Unknown, dberr.Error())
	}
	return &api.Null{}, nil
}
