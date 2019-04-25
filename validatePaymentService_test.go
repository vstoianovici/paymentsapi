package paymentsapi

import (
	"errors"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidateGetPayment(t *testing.T) {
	var ErrAcc = errors.New("uuid: incorrect UUID length: 1")
	type args struct {
		id string
	}
	type serviceResult struct {
		p   Payment
		err error
	}
	tests := []struct {
		name              string
		args              args
		mockServiceResult *serviceResult
		want              Payment
		wantErr           error
	}{
		{
			name: "Should return error invalid payment id when it is not a valid uuid",
			args: args{
				id: "1",
			},
			wantErr: ErrAcc,
		},
		{
			name: "Should return a successful get response",
			args: args{
				id: "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3",
			},
			mockServiceResult: &serviceResult{
				p:   Payment{},
				err: nil,
			},
			want: Payment{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &MockPaymentService{}
			if tt.mockServiceResult != nil {
				mockService.On("GetPayment", tt.args.id).Return(tt.mockServiceResult.p, tt.mockServiceResult.err)
			}
			s, _ := NewValidator(mockService)
			got, err := s.GetPayment(tt.args.id)
			if err != nil {
				assert.Equal(t, err, tt.wantErr)
			} else {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestValidatUpdatePayment(t *testing.T) {
	var ErrAcc = errors.New("uuid: incorrect UUID length: 1")
	var ErrPay = errors.New("Payload could not be validated")
	sUUID := "b50a0337-4bfe-4af7-a02e-3d7126a5101d"
	uuid, _ := uuid.FromString(sUUID)

	type args struct {
		req UpdatePaymentRequest
	}
	type serviceResult struct {
		res UpdatePaymentResponse
		err error
	}
	tests := []struct {
		name              string
		args              args
		mockServiceResult *serviceResult
		want              UpdatePaymentResponse
		wantErr           error
	}{
		{
			name: "Should return error invalid payment id when it is not a valid uuid",
			args: args{
				req: UpdatePaymentRequest{
					PaymentID: "1",
					Payment:   mockPayment(sUUID),
				},
			},
			wantErr: ErrAcc,
		},
		{
			name: "Should return invalid payload when required fields are missing",
			args: args{
				req: UpdatePaymentRequest{
					PaymentID: sUUID,
					Payment:   mockCorruptedPayment(sUUID),
				},
			},
			wantErr: ErrPay,
		},
		{
			name: "Should return a successful update response",
			args: args{
				req: UpdatePaymentRequest{
					PaymentID: sUUID,
					Payment:   mockPayment(sUUID),
				},
			},
			mockServiceResult: &serviceResult{
				res: UpdatePaymentResponse{PaymentID: uuid},
				err: nil,
			},
			want: UpdatePaymentResponse{PaymentID: uuid},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &MockPaymentService{}
			if tt.mockServiceResult != nil {
				mockService.On("UpdatePayment", tt.args.req).Return(tt.mockServiceResult.res, tt.mockServiceResult.err)
			}
			s, _ := NewValidator(mockService)
			got, err := s.UpdatePayment(tt.args.req)
			if err != nil {
				assert.Equal(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.IsType(t, tt.want, got)
			}
		})
	}
}

func TestValidateDeletePayment(t *testing.T) {
	var ErrAcc = errors.New("uuid: incorrect UUID length: 1")
	sUUID := "b50a0337-4bfe-4af7-a02e-3d7126a5101d"
	time0 := new(time.Time)
	gUUID, _ := uuid.FromString(sUUID)
	wrongUUID := "1"
	wUUID, _ := uuid.FromString(wrongUUID)
	type args struct {
		id uuid.UUID
	}
	type serviceResult struct {
		deleteTime *time.Time
		err        error
	}
	tests := []struct {
		name              string
		args              args
		mockServiceResult *serviceResult
		want              *time.Time
		wantErr           error
	}{
		{
			name: "Should return delete payment response",
			args: args{
				id: gUUID,
			},
			mockServiceResult: &serviceResult{
				deleteTime: time0,
				err:        nil,
			},
			want: time0,
		},
		{
			name: "Should return error invalid payment id when it is not a valid uuid",
			args: args{
				id: wUUID,
			},
			wantErr: ErrAcc,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &MockPaymentService{}
			if tt.mockServiceResult != nil {
				mockService.On("DeletePayment", tt.args.id).Return(tt.mockServiceResult.deleteTime, tt.mockServiceResult.err)
			}
			s, _ := NewValidator(mockService)
			got, err := s.DeletePayment(tt.args.id)
			if err != nil {
				assert.Equal(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.IsType(t, tt.want, got)
			}
		})
	}
}

func TestValidatCreatePayment(t *testing.T) {
	var ErrPay = errors.New("Payload could not be validated")
	sUUID := "b50a0337-4bfe-4af7-a02e-3d7126a5101d"
	uuid, _ := uuid.FromString(sUUID)

	type args struct {
		p Payment
	}
	type serviceResult struct {
		p   CreatePaymentResponse
		err error
	}
	tests := []struct {
		name              string
		args              args
		mockServiceResult *serviceResult
		want              CreatePaymentResponse
		wantErr           error
	}{
		{
			name: "Should return invalid payload when required fields are missing",
			args: args{
				p: mockCorruptedPayment(sUUID),
			},
			wantErr: ErrPay,
		},
		{
			name: "Should return a successful create response",
			args: args{
				p: mockPayment(sUUID),
			},
			mockServiceResult: &serviceResult{
				p:   CreatePaymentResponse{PaymentID: uuid},
				err: nil,
			},
			want: CreatePaymentResponse{PaymentID: uuid},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &MockPaymentService{}
			if tt.mockServiceResult != nil {
				mockService.On("CreatePayment", tt.args.p).Return(tt.mockServiceResult.p, tt.mockServiceResult.err)
			}
			s, _ := NewValidator(mockService)
			got, err := s.CreatePayment(tt.args.p)
			if err != nil {
				assert.Equal(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.IsType(t, tt.want, got)
			}
		})
	}
}

func TestValidatePaymentID(t *testing.T) {
	id := "b50a0337-4bfe-4af7-a02e-3d7126a5101d"
	idIncorrect := "b50a033"
	err := validatePaymentID(id)
	assert.NoError(t, err)
	err = validatePaymentID(idIncorrect)
	assert.Error(t, err)
}

func TestValidatePayload(t *testing.T) {
	p := mockPayment("b50a0337-4bfe-4af7-a02e-3d7126a5101d")
	err := validatePayload(p)
	assert.NoError(t, err)
	pIncorrect := mockCorruptedPayment("b50a0337-4bfe-4af7-a02e-3d7126a5101d")
	err = validatePayload(pIncorrect)
	assert.Error(t, err)
}

func mockPayment(id string) Payment {
	pID, _ := uuid.FromString(id)
	uuid, _ := uuid.NewV4()
	p := Payment{
		Type:           "Payment",
		Version:        0,
		ID:             pID,
		OrganisationID: uuid,
		Attributes: Attributes{
			Amount: "100.21",
			BeneficiaryParty: BeneficiaryParty{
				AccountType: 0,
				DebtorParty: DebtorParty{
					AccountName:       "aqssbb",
					AccountNumberCode: "IBAN",
					Address:           "34 frfrf ded",
					Name:              "ING Dfh",
					SponsorParty: SponsorParty{
						AccountNumber: "5678923",
						BankID:        "134667",
						BankIDCode:    "GSDFE",
					},
				},
			},
			ChargesInformation: ChargesInformation{
				BearerCode: "SHAR",
				SenderCharges: []Charge{
					Charge{Amount: "5.00", Currency: "GBP"},
					Charge{Amount: "10.00", Currency: "USD"},
				},
				ReceiverChargesAmount:   "1.00",
				ReceiverChargesCurrency: "USD",
			},
			Currency: "GBP",
			DebtorParty: DebtorParty{
				AccountName:       "deded",
				AccountNumberCode: "IBAN",
				Address:           "1 dhhde ded",
				Name:              "alspnfh",
				SponsorParty: SponsorParty{
					AccountNumber: "5678923",
					BankID:        "134667",
					BankIDCode:    "GSDFE",
				},
			},
			EndToEndReference: "Wil def ee",
			Forex: Forex{
				ContractReference: "FX123",
				ExchangeRate:      "2.0000",
				OriginalAmount:    "200.42",
				OriginalCurrency:  "USD",
			},
			NumericReference:     "10223453",
			PayID:                "123344556790",
			PaymentPurpose:       "course",
			PaymentScheme:        "FPS",
			PaymentType:          "Credit",
			ProcessingDate:       "2017-01-18",
			Reference:            "PAYmen",
			SchemePaymentSubType: "InternetBanking",
			SchemePaymentType:    "Immediate Pay",
			SponsorParty: SponsorParty{
				AccountNumber: "5678923",
				BankID:        "134667",
				BankIDCode:    "GSDFE",
			},
		},
	}

	return p
}

func mockCorruptedPayment(id string) Payment {
	pID, _ := uuid.FromString(id)
	uuid, _ := uuid.NewV4()
	// same definition as in mockPayment but some fields are missing (commented out)
	p := Payment{
		Type:           "Payment",
		Version:        0,
		ID:             pID,
		OrganisationID: uuid,
		Attributes: Attributes{
			Amount: "100.21",
			/*BeneficiaryParty: BeneficiaryParty{
				AccountType: 0,
				DebtorParty: DebtorParty{
					AccountName:       "aqssbb",
					AccountNumberCode: "IBAN",
					Address:           "34 frfrf ded",
					Name:              "ING Dfh",
					SponsorParty: SponsorParty{
						AccountNumber: "5678923",
						BankID:        "134667",
						BankIDCode:    "GSDFE",
					},
				},
			},*/
			ChargesInformation: ChargesInformation{
				BearerCode: "SHAR",
				SenderCharges: []Charge{
					Charge{Amount: "5.00", Currency: "GBP"},
					Charge{Amount: "10.00", Currency: "USD"},
				},
				ReceiverChargesAmount:   "1.00",
				ReceiverChargesCurrency: "USD",
			},
			Currency: "GBP",
			DebtorParty: DebtorParty{
				//AccountName:       "deded",
				AccountNumberCode: "IBAN",
				Address:           "1 dhhde ded",
				Name:              "alspnfh",
				SponsorParty: SponsorParty{
					AccountNumber: "5678923",
					BankID:        "134667",
					BankIDCode:    "GSDFE",
				},
			},
			EndToEndReference: "Wil def ee",
			Forex: Forex{
				ContractReference: "FX123",
				//ExchangeRate:      "2.0000",
				OriginalAmount:   "200.42",
				OriginalCurrency: "USD",
			},
			NumericReference: "10223453",
			PayID:            "123344556790",
			PaymentPurpose:   "course",
			PaymentScheme:    "FPS",
			PaymentType:      "Credit",
			//ProcessingDate:       "2017-01-18",
			Reference:            "PAYmen",
			SchemePaymentSubType: "InternetBanking",
			SchemePaymentType:    "Immediate Pay",
			SponsorParty: SponsorParty{
				AccountNumber: "5678923",
				BankID:        "134667",
				BankIDCode:    "GSDFE",
			},
		},
	}

	return p
}
