package paymentsapi

import (
	"log"
	"testing"
	"time"

	mocket "github.com/Selvatico/go-mocket"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

// to set up tests, you need to register the driver and override the DB instance used across the code base.
func setupTests() *gorm.DB {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	// GORM
	db, err := gorm.Open(mocket.DriverName, "cnnctnString")
	if err != nil {
		log.Fatalf("error mocking gorm: %s", err)
		return &gorm.DB{}
	}
	// Log mode shows the query gorm uses, so we can replicate and mock it
	//db.LogMode(true)

	return db
}

func mockNewPaymentResponse(id string) []map[string]interface{} {
	pid := id
	uuid, _ := uuid.NewV4()
	oid := uuid.String()
	tT := new(time.Time)
	mockResponse := []map[string]interface{}{{
		"created_at":      tT,
		"updated_at":      tT,
		"deleted_at":      tT,
		"type":            "Payment",
		"version":         0,
		"id":              pid,
		"organisation_id": oid,
		"attributes": Attributes{
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
	}}

	return mockResponse
}

func mockNewPaymentListResponse() []map[string]interface{} {
	uuid1, _ := uuid.NewV4()
	pid1 := uuid1.String()

	uuid2, _ := uuid.NewV4()
	pid2 := uuid2.String()

	uuid, _ := uuid.NewV4()
	oid := uuid.String()
	tT := new(time.Time)

	mockResponse := []map[string]interface{}{{
		"created_at":      tT,
		"updated_at":      tT,
		"deleted_at":      tT,
		"type":            "Payment",
		"version":         0,
		"id":              pid1,
		"organisation_id": oid,
		"attributes": Attributes{
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
	},
		{
			"created_at":      tT,
			"updated_at":      tT,
			"deleted_at":      tT,
			"type":            "Payment",
			"version":         0,
			"id":              pid2,
			"organisation_id": oid,
			"attributes": Attributes{
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
		},
	}

	return mockResponse
}

func TestSetupTests(t *testing.T) {
	tstDB := setupTests()
	mockDB := gorm.DB{}
	assert.IsType(t, tstDB, &mockDB)
	assert.NotEqual(t, tstDB, &mockDB)
}

func TestNewCnnctn(t *testing.T) {
	// create string
	cnnctnStr := "host=localhost port=5432 dbname=postgres user=postgres password=password sslmode=disable connect_timeout=5"
	// create connection
	cnnctn := parseCnctionParams("localhost", 5432, "postgres", "postgres", "password", "disable", 5)
	// Assert
	assert.Equal(t, cnnctnStr, cnnctn)
}

func TestGetPayment(t *testing.T) {
	id := "400a75b8-a0aa-4aad-9366-5c609ae390a7"

	db := setupTests()
	defer db.Close()

	mockResponse := mockNewPaymentResponse(id)

	mocket.Catcher.Attach([]*mocket.FakeResponse{
		{
			Pattern:  "SELECT * FROM \"payments\"",
			Response: mockResponse,
		},
	})

	s := NewPaymentService(db)
	p, err := s.GetPayment(id)

	assert.NotNil(t, p)
	assert.NoError(t, err)
}

func TestCreatePayment(t *testing.T) {
	id := "400a75b8-a0aa-4aad-9366-5c609ae390a7"
	uuid, _ := uuid.FromString(id)
	p := Payment{
		ID: uuid,
	}

	db := setupTests()
	defer db.Close()

	mocket.Catcher.Reset().NewMock().WithQuery("INSERT INTO \"payments\"")
	s := NewPaymentService(db)
	rid, err := s.CreatePayment(p)

	assert.NoError(t, err)
	assert.NotEmpty(t, rid)
}

func TestUpdatePayment(t *testing.T) {
	id := "400a75b8-a0aa-4aad-9366-5c609ae390a7"
	uuid, _ := uuid.FromString(id)
	mockResponse := []map[string]interface{}{{}}
	p := Payment{
		ID: uuid,
	}
	r := UpdatePaymentRequest{
		PaymentID: id,
		Payment:   p,
	}

	db := setupTests()
	defer db.Close()

	mocket.Catcher.Attach([]*mocket.FakeResponse{
		{
			Pattern:  "SELECT * FROM \"payments\"",
			Response: mockResponse,
		},
	})
	mocket.Catcher.Attach([]*mocket.FakeResponse{
		{
			Pattern:  "UPDATE \"payments\"",
			Response: mockResponse,
		},
	})

	s := NewPaymentService(db)
	_, err := s.UpdatePayment(r)

	assert.NoError(t, err)
}

func TestDeletePayment(t *testing.T) {
	id := "400a75b8-a0aa-4aad-9366-5c609ae390a7"
	uuid, _ := uuid.FromString(id)
	mockResponse := []map[string]interface{}{{}}

	db := setupTests()
	defer db.Close()

	mocket.Catcher.Attach([]*mocket.FakeResponse{
		{
			Pattern:  "SELECT * FROM \"payments\"",
			Response: mockResponse,
		},
	})
	mocket.Catcher.Attach([]*mocket.FakeResponse{
		{
			Pattern:  "DELETE \"payments\"",
			Response: mockResponse,
		},
	})

	s := NewPaymentService(db)
	_, err := s.DeletePayment(uuid)

	assert.NoError(t, err)
}

func TestGetListPayments(t *testing.T) {
	mockResponse := mockNewPaymentListResponse()

	db := setupTests()
	defer db.Close()

	mocket.Catcher.Attach([]*mocket.FakeResponse{
		{
			Pattern:  "SELECT * FROM \"payments\"",
			Response: mockResponse,
		},
	})

	s := NewPaymentService(db)
	p, err := s.GetListPayments()

	assert.NotNil(t, p)
	assert.NoError(t, err)
}
