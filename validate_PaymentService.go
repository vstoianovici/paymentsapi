package paymentsapi

import (
	"time"

	uuid "github.com/satori/go.uuid"
	valid "gopkg.in/go-playground/validator.v9"
)

var tags = map[string]string{
	"required": "is_required",
	"exists":   "should_exist",
}

// Validator needs to be exported as it is the return type of NewValidator who is also exported
type Validator struct {
	next PaymentService
}

// NewValidator returns a new instance of PaymentService with a model validation layer
func NewValidator(svc PaymentService) (PaymentService, error) {
	return Validator{next: svc}, nil
}

// GetPayment needs to be exported to be accessed outside of the paymentsapi package
func (v Validator) GetPayment(id string) (Payment, error) {
	if err := validatePaymentID(id); err != nil {
		return Payment{}, err
	}
	return v.next.GetPayment(id)
}

// GetListPayments needs to be exported to be accessed outside of the paymentsapi package
func (v Validator) GetListPayments() ([]Payment, error) {
	return v.next.GetListPayments()
}

// CreatePayment needs to be exported to be accessed outside of the paymentsapi package
func (v Validator) CreatePayment(p Payment) (CreatePaymentResponse, error) {
	err := validatePayload(p)
	if err != nil {
		return CreatePaymentResponse{}, err
	}
	return v.next.CreatePayment(p)
}

// UpdatePayment needs to be exported to be accessed outside of the paymentsapi package
func (v Validator) UpdatePayment(req UpdatePaymentRequest) (UpdatePaymentResponse, error) {
	if err := validatePaymentID(req.PaymentID); err != nil {
		return UpdatePaymentResponse{}, err
	}
	err := validatePayload(req.Payment)
	if err != nil {
		return UpdatePaymentResponse{}, err
	}
	return v.next.UpdatePayment(req)
}

// DeletePayment needs to be exported to be accessed outside of the paymentsapi package
func (v Validator) DeletePayment(id uuid.UUID) (*time.Time, error) {
	if err := validatePaymentID(id.String()); err != nil {
		return nil, err
	}
	return v.next.DeletePayment(id)
}

func validatePaymentID(id string) error {
	_, err := uuid.FromString(id)
	if err != nil {
		return err
	}
	return nil
}

func validatePayload(p Payment) error {
	err := valid.New().Struct(p)
	if err != nil {
		return err
	}
	return nil
}
