package service

import (
	"accountservice/model"
	"context"
)

// Service describes the Account service.
type AccountService interface {
	CreateCustomer(ctx context.Context, customer model.Customer) (string, error)
	GetCustomerById(ctx context.Context, id string) (interface{}, error)
	GetAllCustomers(ctx context.Context) (interface{}, error)
	UpdateCustomer(ctx context.Context, customer model.Customer) (string, error)
	DeleteCustomer(ctx context.Context, id string) (string, error)
}
