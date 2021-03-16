package service

import (
	"accountservice/model"
	"accountservice/repository"
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// service implements the Account Service
type accountService struct {
	repository repository.Repository
	logger     log.Logger
}

// NewService creates and returns a new Account service instance
func NewService(rep repository.Repository, logger log.Logger) AccountService {
	return &accountService{
		repository: rep,
		logger:     logger,
	}
}

func (s accountService) CreateCustomer(ctx context.Context, customer model.Customer) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	var msg = "success"

	details := model.Customer{
		CustomerID: customer.CustomerID,
		Email:      customer.Email,
		Phone:      customer.Phone,
	}

	if err := s.repository.CreateCustomer(ctx, details); err != nil {
		level.Error(logger).Log("err from repo is ", err)
		return "", err
	}
	return msg, nil
}

func (s accountService) GetCustomerById(ctx context.Context, id string) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetcustomerById")

	var customer interface{}
	var empty interface{}
	customer, err := s.repository.GetCustomerById(ctx, id)
	if err != nil {
		level.Error(logger).Log("err ", err)
		return empty, err
	}
	return customer, nil
}

func (s accountService) GetAllCustomers(ctx context.Context) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetAllcustomers")
	var customer interface{}
	var empty interface{}
	customer, err := s.repository.GetAllCustomers(ctx)
	if err != nil {
		level.Error(logger).Log("err ", err)
		return empty, err
	}
	return customer, nil
}

func (s accountService) UpdateCustomer(ctx context.Context, customer model.Customer) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	var msg = "success"
	customerDetails := model.Customer{
		CustomerID: customer.CustomerID,
		Email:      customer.Email,
		Phone:      customer.Phone,
	}
	msg, err := s.repository.UpdateCustomer(ctx, customerDetails)
	if err != nil {
		level.Error(logger).Log("err from repo is ", err)
		return "", err
	}
	return msg, nil
}

func (s accountService) DeleteCustomer(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "DeleteCustomer")
	msg, err := s.repository.DeleteCustomer(ctx, id)
	if err != nil {
		level.Error(logger).Log("err ", err)
		return "", err
	}
	return msg, nil
}
