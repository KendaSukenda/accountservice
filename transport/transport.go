package transport

import (
	"accountservice/model"
	"accountservice/service"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

// Endpoint for the Account service.
func MakeCreateCustomerEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCustomerRequest)
		msg, err := s.CreateCustomer(ctx, req.customer)
		return CreateCustomerResponse{Message: msg, Err: err}, nil
	}
}

func MakeGetCustomerByIdEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCustomerByIdRequest)
		customerDetails, err := s.GetCustomerById(ctx, req.Id)
		if err != nil {
			return GetCustomerByIdResponse{Customer: customerDetails, Err: "Id not found"}, nil
		}
		return GetCustomerByIdResponse{Customer: customerDetails, Err: ""}, nil
	}
}

func MakeGetAllCustomersEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		customerDetails, err := s.GetAllCustomers(ctx)
		if err != nil {
			return GetAllCustomersResponse{Customer: customerDetails, Err: "no data found"}, nil
		}
		return GetAllCustomersResponse{Customer: customerDetails, Err: ""}, nil
	}
}

func MakeDeleteCustomerEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCustomerRequest)
		msg, err := s.DeleteCustomer(ctx, req.CustomerId)
		if err != nil {
			return DeleteCustomerResponse{Msg: msg, Err: err}, nil
		}
		return DeleteCustomerResponse{Msg: msg, Err: nil}, nil
	}
}

func MakeUpdateCustomerEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateCustomerRequest)
		msg, err := s.UpdateCustomer(ctx, req.customer)
		return msg, err
	}
}

func DecodeCreateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateCustomerRequest
	fmt.Println("-------->>>>into Decoding")
	if err := json.NewDecoder(r.Body).Decode(&req.customer); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeGetCustomerByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetCustomerByIdRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = GetCustomerByIdRequest{
		Id: vars["customerId"],
	}
	return req, nil
}

func DecodeGetAllCustomersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into GETALL Decoding")
	var req GetAllCustomersRequest
	return req, nil
}

func DecodeDeleteCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Delete Decoding")
	var req DeleteCustomerRequest
	vars := mux.Vars(r)
	req = DeleteCustomerRequest{
		CustomerId: vars["customerId"],
	}
	return req, nil
}

func DecodeUpdateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Update Decoding")
	var req UpdateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req.customer); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("into Encoding <<<<<<----------------")
	return json.NewEncoder(w).Encode(response)
}

type (
	CreateCustomerRequest struct {
		customer model.Customer
	}

	CreateCustomerResponse struct {
		Message string `json:"message"`
		Err     error  `json:"error,omitempty"`
	}

	GetCustomerByIdRequest struct {
		Id string `json:"customerId"`
	}

	GetCustomerByIdResponse struct {
		Customer interface{} `json:"customer,omitempty"`
		Err      string      `json:"error,omitempty"`
	}

	GetAllCustomersRequest struct{}

	GetAllCustomersResponse struct {
		Customer interface{} `json:"customer,omitempty"`
		Err      string      `json:"error,omitempty"`
	}

	DeleteCustomerRequest struct {
		CustomerId string `json:"customerId"`
	}

	DeleteCustomerResponse struct {
		Msg string `json:"response"`
		Err error  `json:"error,omitempty"`
	}

	UpdateCustomerRequest struct {
		customer model.Customer
	}
)
