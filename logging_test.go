package paymentsapi

import (
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockNextService struct {
	called bool
}

func (m *mockNextService) GetPayment(s string) (output Payment, err error) {
	m.called = true
	return Payment{}, nil
}

func (m *mockNextService) CreatePayment(p Payment) (output CreatePaymentResponse, err error) {
	m.called = true
	return CreatePaymentResponse{}, nil
}

func (m *mockNextService) UpdatePayment(p UpdatePaymentRequest) (output UpdatePaymentResponse, err error) {
	m.called = true
	return UpdatePaymentResponse{}, nil
}

func (m *mockNextService) DeletePayment(id uuid.UUID) (t *time.Time, err error) {
	m.called = true
	return nil, nil
}

func (m *mockNextService) GetListPayments() (output []Payment, err error) {
	m.called = true
	return nil, nil
}

func TestLogGetPayment(t *testing.T) {
	m := &mockNextService{}
	s := NewLogging(log.NewNopLogger(), m)
	assert.False(t, m.called)
	_, err := s.GetPayment("")
	assert.Nil(t, err)
	assert.True(t, m.called)
	p := Payment{}
	p.Type = "Payment"
	mockService := &MockPaymentService{}
	mockService.On("GetPayment", mock.Anything).Return(p, nil)
	s1 := NewLogging(log.NewNopLogger(), mockService)
	//assert.False(t, mockService.called)
	_, err = s1.GetPayment("abcd")
	assert.Nil(t, err)
	//assert.True(t, mockService.called)
}

func TestLogCreatePayment(t *testing.T) {
	m := &mockNextService{}
	s := NewLogging(log.NewNopLogger(), m)
	assert.False(t, m.called)
	p := Payment{}
	_, err := s.CreatePayment(p)
	assert.Nil(t, err)
	assert.True(t, m.called)
}

func TestLogUpdatePayment(t *testing.T) {
	m := &mockNextService{}
	s := NewLogging(log.NewNopLogger(), m)
	assert.False(t, m.called)
	p := UpdatePaymentRequest{}
	_, err := s.UpdatePayment(p)
	assert.Nil(t, err)
	assert.True(t, m.called)
	// up := UpdatePaymentRequest{}
	// up.Payment = map[string]interface{}{
	// 	"foo": make(chan int),
	// }
	// mockService := &MockPaymentService{}
	// mockService.On("UpdateListPayments", up).Return(slice, nil)
	// s1 := NewLogging(log.NewNopLogger(), mockService)
	// assert.False(t, m.called)
	// _, err = s1.GetListPayments()
	// assert.Nil(t, err)
	// assert.True(t, m.called)

}

func TestDeleteUpdatePayment(t *testing.T) {
	m := &mockNextService{}
	s := NewLogging(log.NewNopLogger(), m)
	assert.False(t, m.called)
	uuid, _ := uuid.NewV4()
	_, err := s.DeletePayment(uuid)
	assert.Nil(t, err)
	assert.True(t, m.called)
}

func TestGetListPayment(t *testing.T) {
	m := &mockNextService{}
	s := NewLogging(log.NewNopLogger(), m)
	assert.False(t, m.called)
	_, err := s.GetListPayments()
	assert.Nil(t, err)
	assert.True(t, m.called)
	p1 := Payment{}
	p2 := Payment{}
	slice := []Payment{
		p1,
		p2,
	}
	mockService := &MockPaymentService{}
	mockService.On("GetListPayments", mock.Anything).Return(slice, nil)
	s1 := NewLogging(log.NewNopLogger(), mockService)
	//assert.False(t, mockService.called)
	output, err := s1.GetListPayments()
	assert.Nil(t, err)
	assert.NotNil(t, output)
	//assert.True(t, mockService.called)
}
