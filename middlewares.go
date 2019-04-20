package paymentsapi

import (
	"context"
	"errors"
	"time"

	"github.com/go-kit/kit/endpoint"
	uuid "github.com/satori/go.uuid"
)

// GetPaymentRequest is the request type used to retrieve a specific payment
type GetPaymentRequest struct {
	PaymentID string
}

// GetPaymentResponse is the request type used to retrieve a specific payment
type GetPaymentResponse struct {
	Payment
}

// GetListPaymentRequest is the request type used to list all payments
type GetListPaymentRequest struct{}

// CreatePaymentRequest is the request type used to insert a new payment
type CreatePaymentRequest struct {
	Payment
}

// CreatePaymentResponse is the response returned after creating a new payment containing the Payment ID
type CreatePaymentResponse struct {
	PaymentID uuid.UUID `json:"created_id"`
}

// UpdatePaymentResponse is the response returned after updating an already existing payment based on its ID
type UpdatePaymentResponse struct {
	PaymentID uuid.UUID `json:"updated_id"`
}

// UpdatePaymentRequest is the request passed when updating a payment based on the ID and the new payment information.
type UpdatePaymentRequest struct {
	PaymentID string
	Payment   Payment
}

// DeletePaymentRequest represents the type needed when requesting to delete a payment
type DeletePaymentRequest struct {
	PaymentID uuid.UUID `json:"id"`
}

// DeletePaymentResponse represents the type needed as a response to a payment deletion
type DeletePaymentResponse struct {
	DeletedAt *time.Time `json:"DeletedAt"`
}

// MakeGetListPaymentsEndpoint is an endpoint constructor that takes a service and constructs individual endpoints for the GetListPayments method
func MakeGetListPaymentsEndpoint(svc PaymentService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		return svc.GetListPayments()
	}
}

// MakeGetPaymentEndpoint is an endpoint constructor that takes a service and constructs individual endpoints for the GetPayment method
func MakeGetPaymentEndpoint(svc PaymentService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPaymentRequest)
		v, err := svc.GetPayment(req.PaymentID)
		if err != nil {
			var ErrAcc = errors.New("err: Could not GET payment")
			cErr := errors.New(ErrAcc.Error() + err.Error())
			return nil, cErr
		}
		return v, nil
	}
}

// MakeCreatePaymentEndpoint is an endpoint constructor that takes a service and constructs individual endpoints for the CreatePayment method
func MakeCreatePaymentEndpoint(svc PaymentService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePaymentRequest)
		v, err := svc.CreatePayment(req.Payment)
		if err != nil {
			var ErrAcc = errors.New("err: Could not Create(POST) payment")
			cErr := errors.New(ErrAcc.Error() + err.Error())
			return nil, cErr
		}
		return v, nil
	}
}

// MakeUpdatePaymentEndpoint is an endpoint constructor that takes a service and constructs individual endpoints for the UpdatePayment method
func MakeUpdatePaymentEndpoint(svc PaymentService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdatePaymentRequest)
		v, err := svc.UpdatePayment(req)
		if err != nil {
			var ErrAcc = errors.New("err: Could not Update(PUT) payment")
			cErr := errors.New(ErrAcc.Error() + err.Error())
			return UpdatePaymentRequest{}, cErr
		}
		return v, nil
	}
}

// MakeDeletePaymentEndpoint is an endpoint constructor that takes a service and constructs individual endpoints for the DeletePayment method
func MakeDeletePaymentEndpoint(svc PaymentService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(DeletePaymentRequest)
		t, err := svc.DeletePayment(req.PaymentID)
		if err != nil {
			var ErrAcc = errors.New("err: Could not DELETE payment")
			cErr := errors.New(ErrAcc.Error() + err.Error())
			return DeletePaymentRequest{}, cErr
		}
		return DeletePaymentResponse{DeletedAt: t}, nil
	}
}
