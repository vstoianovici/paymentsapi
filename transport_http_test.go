package paymentsapi

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_NewHTTPTransport(t *testing.T) {
	svc := &MockPaymentService{}
	h := NewHTTPTransport(svc)
	assert.NotNil(t, h)
}

func TestDecodeGetListPaymentsRequest(t *testing.T) {
	type empty struct{}
	expected := empty{}
	r := httptest.NewRequest("GET", "/v1/payments", bytes.NewBufferString("{}"))
	o, err := DecodeGetListPaymentsRequest(context.Background(), r)
	assert.Nil(t, err)
	assert.EqualValues(t, expected, o)
}

func TestDecodeGetPaymentRequest(t *testing.T) {
	id := "400a75b8-a0aa-4aad-9366-5c609ae390a7"
	expected := GetPaymentRequest{PaymentID: id}
	httpRequest, _ := http.NewRequest("GET", "/v1/payments/400a75b8-a0aa-4aad-9366-5c609ae390a7/", nil)
	httpRequest = mux.SetURLVars(httpRequest, map[string]string{"id": "400a75b8-a0aa-4aad-9366-5c609ae390a7"})
	req, err := DecodeGetPaymentRequest(context.Background(), httpRequest)
	assert.NoError(t, err)
	assert.Equal(t, expected, req)
}

func TestDecodeCreatePaymentRequest(t *testing.T) {
	p := Payment{}
	expected := CreatePaymentRequest{p}
	r := httptest.NewRequest("POST", "/v1/payments", bytes.NewBufferString("{}"))
	req, err := DecodeCreatePaymentRequest(context.Background(), r)
	assert.Nil(t, err)
	assert.Equal(t, expected, req.(CreatePaymentRequest))
}

func TestDecodeUpdatePayementRequest(t *testing.T) {
	id := "400a75b8-a0aa-4aad-9366-5c609ae390a7"
	expectedResult := UpdatePaymentRequest{PaymentID: id}
	httpRequest, err := http.NewRequest("PUT", "/v1/payments/400a75b8-a0aa-4aad-9366-5c609ae390a7/", bytes.NewBufferString("{}"))
	httpRequest = mux.SetURLVars(httpRequest, map[string]string{"id": "400a75b8-a0aa-4aad-9366-5c609ae390a7"})
	req, err := DecodeUpdatePayementRequest(context.Background(), httpRequest)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, req)
}

func TestDecodeDeletePayementRequest(t *testing.T) {
	id := "400a75b8-a0aa-4aad-9366-5c609ae390a7"
	uuid, _ := uuid.FromString(id)
	expectedResult := DeletePaymentRequest{PaymentID: uuid}
	httpRequest, err := http.NewRequest("DELETE", "/v1/payments/400a75b8-a0aa-4aad-9366-5c609ae390a7/", nil)
	httpRequest = mux.SetURLVars(httpRequest, map[string]string{"id": "400a75b8-a0aa-4aad-9366-5c609ae390a7"})
	req, err := DecodeDeletePayementRequest(context.Background(), httpRequest)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, req)
}

func TestEncodeBasicResponse(t *testing.T) {
	rec := httptest.NewRecorder()
	err := EncodeBasicResponse(context.Background(), rec, GetPaymentRequest{})
	assert.Equal(t, rec.Header().Get("Content-Type"), "")
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.Nil(t, err)
}

func TestEncodeCreationResponse(t *testing.T) {
	rec := httptest.NewRecorder()
	err := EncodeCreationResponse(context.Background(), rec, CreatePaymentRequest{})
	assert.Equal(t, rec.Code, http.StatusCreated)
	assert.Nil(t, err)
}

func TestTreatErr(t *testing.T) {
	var err = errors.New("test error")
	var s = "on top of"
	newErr := treatErr(err, s)
	c := s + err.Error()
	assert.Equal(t, c, newErr.Error())
}
