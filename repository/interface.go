package repository

import (
	"accountservice/model"
	"context"
)

type Repository interface {
	CreateCustomer(ctx context.Context, customer model.Customer) error
	GetCustomerById(ctx context.Context, id string) (interface{}, error)
	GetAllCustomers(ctx context.Context) (interface{}, error)
	UpdateCustomer(ctx context.Context, customer model.Customer) (string, error)
	DeleteCustomer(ctx context.Context, id string) (string, error)
}
