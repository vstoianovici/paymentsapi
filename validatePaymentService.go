package paymentsapi

import (
	"errors"
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
		var ErrPay = errors.New("Payload could not be validated")
		return CreatePaymentResponse{}, ErrPay
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
		var ErrPay = errors.New("Payload could not be validated")
		return UpdatePaymentResponse{}, ErrPay
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
	rUUID, err := uuid.FromString(id)
	zero := "0"
	zeroUUID, _ := uuid.FromString(zero)
	if err != nil {
		return err
	}
	// because casting a string that does not match the UUID format to a uuid.UUID type results in a uuid with
	// value "00000000-0000-0000-0000-000000000000", we test for it and output an error
	if zeroUUID == rUUID {
		var ErrAcc = errors.New("uuid: incorrect UUID length: 1")
		return ErrAcc
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
