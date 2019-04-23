package paymentsapi

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMakeGetListPaymentsEndpoint(t *testing.T) {
	response := []Payment{}
	tError := errors.New("error in test")

	tests := []struct {
		name        string
		Service     func() PaymentService
		isError     bool
		ExpError    string
		ExpResponse []Payment
	}{
		{
			name: "MakeGetListPaymentsEndpoint successful for all payments",
			Service: func() PaymentService {
				mockSvc := &MockPaymentService{}
				mockSvc.On("GetListPayments", mock.Anything).Return(response, nil)
				return mockSvc
			},
			isError:     false,
			ExpResponse: response,
			ExpError:    "",
		},
		{
			name: "MakeGetListPaymentsEndpoint failed for all payments",
			Service: func() PaymentService {
				mockSvc := &MockPaymentService{}
				mockSvc.On("GetListPayments", mock.Anything).Return(response, tError)
				return mockSvc
			},
			isError:     true,
			ExpResponse: nil,
			ExpError:    "error in test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ep := MakeGetListPaymentsEndpoint(tt.Service())
			rt := GetListPaymentRequest{}
			lp, err := ep(nil, rt)

			if tt.isError {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tt.ExpError)
				result, ok := lp.([]Payment)
				assert.Equal(t, true, ok)
				assert.Equal(t, []Payment{}, result)
				return
			}
			assert.Nil(t, err)
			result, ok := lp.([]Payment)
			assert.Equal(t, true, ok)
			assert.NotNil(t, result)

		})
	}
}

func TestMakeGetPaymentEndpoint(t *testing.T) {
	response := Payment{}
	tError := errors.New(" test")

	tests := []struct {
		name        string
		Service     func() PaymentService
		isError     bool
		ExpError    string
		ExpResponse Payment
	}{
		{
			name: "MakeGetPaymentEndpoint successfully GETs a payment",
			Service: func() PaymentService {
				mockSvc := &MockPaymentService{}
				mockSvc.On("GetPayment", mock.Anything).Return(response, nil)
				return mockSvc
			},
			isError:     false,
			ExpResponse: response,
			ExpError:    "",
		},
		{
			name: "MakeGetPaymentEndpoint failed to GET a payment",
			Service: func() PaymentService {
				mockSvc := &MockPaymentService{}
				mockSvc.On("GetPayment", mock.Anything).Return(response, tError)
				return mockSvc
			},
			isError:     true,
			ExpResponse: Payment{},
			ExpError:    "err: Could not GET payment test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ep := MakeGetPaymentEndpoint(tt.Service())
			rt := GetPaymentRequest{}
			lp, err := ep(nil, rt)
			if tt.isError {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tt.ExpError)
				assert.Nil(t, lp)
				return
			}
			assert.Nil(t, err)
			result, ok := lp.(Payment)
			assert.Equal(t, true, ok)
			assert.NotNil(t, result)

		})
	}
}

func TestMakeCreatePaymentEndpoint(t *testing.T) {
	response := CreatePaymentResponse{}
	tError := errors.New(" test")

	tests := []struct {
		name        string
		Service     func() PaymentService
		isError     bool
		ExpError    string
		ExpResponse CreatePaymentResponse
	}{
		{
			name: "MakeCreatePaymentEndpoint successfully POSTs a payment",
			Service: func() PaymentService {
				mockSvc := &MockPaymentService{}
				mockSvc.On("CreatePayment", mock.Anything).Return(response, nil)
				return mockSvc
			},
			isError:     false,
			ExpResponse: response,
			ExpError:    "",
		},
		{
			name: "MakeCreatePaymentEndpoint failed to POST a payment",
			Service: func() PaymentService {
				mockSvc := &MockPaymentService{}
				mockSvc.On("CreatePayment", mock.Anything).Return(response, tError)
				return mockSvc
			},
			isError:     true,
			ExpResponse: CreatePaymentResponse{},
			ExpError:    "err: Could not Create(POST) payment test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ep := MakeCreatePaymentEndpoint(tt.Service())
			rt := CreatePaymentRequest{}
			lp, err := ep(nil, rt)
			if tt.isError {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tt.ExpError)
				assert.Nil(t, lp)
				return
			}
			assert.Nil(t, err)
			result, ok := lp.(CreatePaymentResponse)
			assert.Equal(t, true, ok)
			assert.NotNil(t, result)

		})
	}
}

func TestMakeDeletePaymentEndpoint(t *testing.T) {
	response := new(time.Time)
	tError := errors.New(" test")

	tests := []struct {
		name        string
		Service     func() PaymentService
		isError     bool
		ExpError    string
		ExpResponse *time.Time
	}{
		{
			name: "MakeDeletePaymentEndpoint successfully DELETEs a payment",
			Service: func() PaymentService {
				mockSvc := &MockPaymentService{}
				mockSvc.On("DeletePayment", mock.Anything).Return(response, nil)
				return mockSvc
			},
			isError:     false,
			ExpResponse: response,
			ExpError:    "",
		},
		{
			name: "MakeDeletePaymentEndpoint failed to DELETE a payment",
			Service: func() PaymentService {
				mockSvc := &MockPaymentService{}
				mockSvc.On("DeletePayment", mock.Anything).Return(response, tError)
				return mockSvc
			},
			isError:     true,
			ExpResponse: response,
			ExpError:    "err: Could not DELETE payment test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ep := MakeDeletePaymentEndpoint(tt.Service())
			rt := DeletePaymentRequest{}
			lp, err := ep(nil, rt)
			result, ok := lp.(DeletePaymentResponse)
			dp := DeletePaymentResponse{}
			if tt.isError {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tt.ExpError)
				assert.Equal(t, result, dp)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, true, ok)
			assert.NotNil(t, result)

		})
	}
}

func TestMakeUpdatePaymentEndpoint(t *testing.T) {
	response := UpdatePaymentResponse{}
	tError := errors.New(" test")

	tests := []struct {
		name        string
		Service     func() PaymentService
		isError     bool
		ExpError    string
		ExpResponse UpdatePaymentResponse
	}{
		{
			name: "MakeUpdatePaymentEndpoint successfully PUTs a payment",
			Service: func() PaymentService {
				mockSvc := &MockPaymentService{}
				mockSvc.On("UpdatePayment", mock.Anything).Return(response, nil)
				return mockSvc
			},
			isError:     false,
			ExpResponse: response,
			ExpError:    "",
		},
		{
			name: "MakeUpdatePaymentEndpoint failed to PUT a payment",
			Service: func() PaymentService {
				mockSvc := &MockPaymentService{}
				mockSvc.On("UpdatePayment", mock.Anything).Return(response, tError)
				return mockSvc
			},
			isError:     true,
			ExpResponse: response,
			ExpError:    "err: Could not Update(PUT) payment test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ep := MakeUpdatePaymentEndpoint(tt.Service())
			rt := UpdatePaymentRequest{}
			lp, err := ep(nil, rt)
			result, ok := lp.(UpdatePaymentResponse)
			if tt.isError {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tt.ExpError)
				assert.Equal(t, result, response)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, true, ok)
			assert.NotNil(t, result)

		})
	}
}
