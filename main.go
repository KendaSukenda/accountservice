package main

import (
	"accountservice/repository"
	"accountservice/service"
	"accountservice/transport"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	// "github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	// kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	// stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	db := repository.GetDBConnection()

	newRouter := mux.NewRouter()

	var svc service.AccountService
	repo, err := repository.NewRepo(db, logger)
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}
	svc = service.NewService(repo, logger)

	// svc = loggingMiddleware{logger, svc}
	// svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	CreateAccountHandler := httptransport.NewServer(
		transport.MakeCreateCustomerEndpoint(svc),
		transport.DecodeCreateCustomerRequest,
		transport.EncodeResponse,
	)

	GetByCustomerIdHandler := httptransport.NewServer(
		transport.MakeGetCustomerByIdEndpoint(svc),
		transport.DecodeGetCustomerByIdRequest,
		transport.EncodeResponse,
	)

	GetAllCustomersHandler := httptransport.NewServer(
		transport.MakeGetAllCustomersEndpoint(svc),
		transport.DecodeGetAllCustomersRequest,
		transport.EncodeResponse,
	)

	DeleteCustomerHandler := httptransport.NewServer(
		transport.MakeDeleteCustomerEndpoint(svc),
		transport.DecodeDeleteCustomerRequest,
		transport.EncodeResponse,
	)

	UpdateCustomerHandler := httptransport.NewServer(
		transport.MakeUpdateCustomerEndpoint(svc),
		transport.DecodeUpdateCustomerRequest,
		transport.EncodeResponse,
	)

	http.Handle("/", newRouter)
	http.Handle("/account", CreateAccountHandler)
	http.Handle("/account/update", UpdateCustomerHandler)

	newRouter.Handle("/account/getAll", GetAllCustomersHandler).Methods("GET")
	newRouter.Handle("/account/{customerId}", GetByCustomerIdHandler).Methods("GET")
	newRouter.Handle("/account/{customerId}", DeleteCustomerHandler).Methods("DELETE")

	// http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8000")
	logger.Log("err", http.ListenAndServe(":8000", nil))
}
