package paymentsapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"

	httptransport "github.com/go-kit/kit/transport/http"
)

// NewHTTPTransport creates a new JSON over HTTP transport
func NewHTTPTransport(svc PaymentService) http.Handler {
	// define a way to service a request for the getListPaymentstHandler endpoint
	getListPaymentstHandler := httptransport.NewServer(
		MakeGetListPaymentsEndpoint(svc),
		DecodeGetListPaymentsRequest,
		EncodeBasicResponse,
	)

	// define a way to service a request for the getPaymentHandler endpoint
	getPaymentHandler := httptransport.NewServer(
		MakeGetPaymentEndpoint(svc),
		DecodeGetPaymentRequest,
		EncodeBasicResponse,
	)
	// define a way to service a request for the createPaymentHandler endpoint
	createPaymentHandler := httptransport.NewServer(
		MakeCreatePaymentEndpoint(svc),
		DecodeCreatePaymentRequest,
		EncodeCreationResponse,
	)
	// define a way to service a request for the updatePaymentHandler endpoint
	updatePaymentHandler := httptransport.NewServer(
		MakeUpdatePaymentEndpoint(svc),
		DecodeUpdatePayementRequest,
		EncodeCreationResponse,
	)
	// define a way to service a request for the deletePaymentHandler endpoint
	deletePaymentHandler := httptransport.NewServer(
		MakeDeletePaymentEndpoint(svc),
		DecodeDeletePayementRequest,
		EncodeBasicResponse,
	)

	// Define a new router that will handle API endpoints for all the above defined handlers
	router := mux.NewRouter()
	router.Handle("/v1/payments/", getListPaymentstHandler).Methods("GET")
	router.Handle("/v1/payments/{id}", getPaymentHandler).Methods("GET")
	router.Handle("/v1/payments/", createPaymentHandler).Methods("POST")
	router.Handle("/v1/payments/{id}", updatePaymentHandler).Methods("PUT")
	router.Handle("/v1/payments/{id}", deletePaymentHandler).Methods("DELETE")
	return router
}

// DecodeGetListPaymentsRequest exported to be accessible from outside the package (from main)
func DecodeGetListPaymentsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	type empty struct{}
	return empty{}, nil
}

// DecodeGetPaymentRequest exported to be accessible from outside the package (from main)
func DecodeGetPaymentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	return GetPaymentRequest{PaymentID: id}, nil
}

// DecodeCreatePaymentRequest exported to be accessible from outside the package (from main)
func DecodeCreatePaymentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreatePaymentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	newErr := treatErr(err, "err: Could not read 'create payment' body")
	if newErr != nil {
		return nil, newErr
	}
	return req, nil
}

// DecodeUpdatePayementRequest exported to be accessible from outside the package (from main)
func DecodeUpdatePayementRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	var req UpdatePaymentRequest
	err := json.NewDecoder(r.Body).Decode(&req.Payment)
	newErr := treatErr(err, "err: Could not read 'update payment' body")
	if newErr != nil {
		return nil, newErr
	}
	req.PaymentID = id
	return req, nil
}

// DecodeDeletePayementRequest exported to be accessible from outside the package (from main)
func DecodeDeletePayementRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	newErr := treatErr(err, "err: Could not read 'delete payment' body")
	if newErr != nil {
		return nil, newErr
	}
	return DeletePaymentRequest{PaymentID: id}, nil
}

func treatErr(err error, s string) error {
	if err != nil {
		var ErrAcc = errors.New(s)
		cErr := errors.New(ErrAcc.Error() + err.Error())
		return cErr
	}
	return nil
}

// EncodeBasicResponse exported to be accessible from outside the package (from main)
func EncodeBasicResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

// EncodeCreationResponse exported to be accessible from outside the package (from main)
func EncodeCreationResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(response)
}
