package paymentsapi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	uuid "github.com/satori/go.uuid"
)

// The logging middleware amends the wallet service with a logger

// loggingMiddleware is the type of the wrapper around the core service and any other functionality layers
type loggingMiddleware struct {
	logger log.Logger
	next   PaymentService
}

// NewLogging is how the logging middleware (loggingMiddleware struct) is constructed (the function is exported so it can be used from outside the package)
func NewLogging(logger log.Logger, next PaymentService) PaymentService {
	return &loggingMiddleware{
		logger: logger,
		next:   next,
	}
}

// GetPayment function is implemented for logging layer as the request traverses through the logging layer down to the next layer
func (mw loggingMiddleware) GetPayment(s string) (output Payment, err error) {
	// Log everything that the function sees in the provided format
	defer func(begin time.Time) {
		status := func(in Payment) string {
			if in.Type == "Payment" {
				return "success"
			}
			return "payment was not found"
		}(output)
		_ = mw.logger.Log(
			"method", "getPayment",
			"input", s,
			"output", status,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	// The function calls the next layer down
	output, err = mw.next.GetPayment(s)
	return
}

// CreatePayment function is implemented for the logging layer as the request traverses through the logging layer down to the next layer
func (mw loggingMiddleware) CreatePayment(p Payment) (output CreatePaymentResponse, err error) {
	// Log everything that the function sees in the provided format
	defer func(begin time.Time) {
		e, err := json.Marshal(p)
		if err != nil {
			return
		}
		fmt.Println(string(e))
		_ = mw.logger.Log(
			"method", "createPayment",
			"input", "Input "+string(e),
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	// The function calls the next layer down
	output, err = mw.next.CreatePayment(p)
	return
}

// UpdatePayment function is implemented for the logging layer as the request traverses through the logging layer down to the next layer
func (mw loggingMiddleware) UpdatePayment(p UpdatePaymentRequest) (output UpdatePaymentResponse, err error) {
	// Log everything that the function sees in the provided format
	defer func(begin time.Time) {
		e, err := json.Marshal(p.Payment)
		if err != nil {
			return
		}
		output := "Updated" + p.PaymentID
		_ = mw.logger.Log(
			"method", "updatePayment",
			"input", "Input "+string(e)+"for id:"+p.PaymentID,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	// The function calls the next layer down
	output, err = mw.next.UpdatePayment(p)
	return
}

// DeletePayment function is implemented for the logging layer as the request traverses through the logging layer down to the next layer
func (mw loggingMiddleware) DeletePayment(id uuid.UUID) (t *time.Time, err error) {
	// Log everything that the function sees in the provided format
	defer func(begin time.Time) {
		output := "Deleted" + id.String()
		_ = mw.logger.Log(
			"method", "deletePayment",
			"input", "Delete id:"+id.String(),
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	// The function calls the next layer down
	t, err = mw.next.DeletePayment(id)
	return
}

// GetListPaymentsfunction is implemented for the logging layer as the request traverses through the logging layer down to the next layer
func (mw loggingMiddleware) GetListPayments() (output []Payment, err error) {

	defer func(begin time.Time) {
		status := func(in []Payment) string {
			if len(in) != 0 {
				return "success"
			}
			return "no payments"
		}(output)
		_ = mw.logger.Log(
			"method", "getListPayments",
			"input", "List Payments",
			"output", status,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	// The function calls the next layer down
	output, err = mw.next.GetListPayments()
	return
}
