package paymentsapi

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	config "github.com/vstoianovici/paymentsapi/config"

	_ "github.com/lib/pq" //pq imports the postgres driver
	uuid "github.com/satori/go.uuid"
)

// PaymentService is an interface that implements a simple RESTful API for Payment Service (CRUD functionality against a postgresql DB).
// PaymentService can retrieve a list of all submitted Payments (GetListPayment), get a payment based on a payment ID (GetPayement), create a payment based on a json file and return its ID,
// update a payment based on the original payment ID and a new payment json file and delete a payment (softdelete - DeletedAt will have a timestamp but the entry will still be available)
type PaymentService interface {
	GetPayment(id string) (Payment, error)
	GetListPayments() ([]Payment, error)
	CreatePayment(p Payment) (CreatePaymentResponse, error)
	UpdatePayment(p UpdatePaymentRequest) (UpdatePaymentResponse, error)
	DeletePayment(id uuid.UUID) (*time.Time, error)
}

type paymentService struct {
	db *gorm.DB
}

const cnnctnString = "host=%s port=%d dbname=%s user=%s password=%s sslmode=%s connect_timeout=%d"

// NewPaymentService is the payment API contructor function
func NewPaymentService(db *gorm.DB) PaymentService {
	return &paymentService{
		db: db,
	}
}

// NewDBConnection implements the connection to the Postgresql DB
func NewDBConnection(file string) (*gorm.DB, error) {
	dbConfig, err := config.GetDbConfig(file)
	cnnction := parseCnctionParams(dbConfig.Host, dbConfig.Port, dbConfig.DBName, dbConfig.User, dbConfig.Password, dbConfig.Sslmode, dbConfig.Timeout)
	db, err := gorm.Open(dbConfig.Driver, cnnction)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// define the necessary format for the open connection operation
func parseCnctionParams(host string, port int, name string, user string, password string, sslmode string, timeout int) string {
	return fmt.Sprintf(cnnctnString, host, port, name, user, password, sslmode, timeout)
}

// MigrateDB initializes db schema with needed tables
func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&Payment{}, &Attributes{}, &BeneficiaryParty{}, &DebtorParty{}, &SponsorParty{}, &ChargesInformation{}, &Charge{}, &Forex{})
}

// CloseDB closes the connection to the database
func CloseDB(db *gorm.DB) {
	if db != nil {
		db.Close()
	}
}

// GetPayment retrieves (GET) and displays a payment based on a provided ID
func (r *paymentService) GetPayment(id string) (Payment, error) {
	p := Payment{}
	err := r.db.Model(&p).Where("id = ?", id).Preload("Attributes.BeneficiaryParty").Preload("Attributes.ChargesInformation.SenderCharges").Preload("Attributes.DebtorParty").Preload("Attributes.Forex").Preload("Attributes.SponsorParty").Find(&p).Error
	//err := r.db.Debug().Model(&p).Where("id = ?", id).Preload("Attributes.BeneficiaryParty").Preload("Attributes.ChargesInformation.SenderCharges").Preload("Attributes.DebtorParty").Preload("Attributes.Forex").Preload("Attributes.SponsorParty").Find(&p).Error
	if err != nil {
		return p, err
	}
	return p, nil
}

// CreatePayment creates a payment (POST) based on a provided payment json file that has all the right information
func (r *paymentService) CreatePayment(p Payment) (CreatePaymentResponse, error) {
	paymentID, err := uuid.NewV4()
	if err != nil {
		e := CreatePaymentResponse{}
		return e, err
	}
	p.ID = paymentID
	err = r.db.Save(&p).Error
	//err = r.db.Debug().Save(&p).Error
	if err != nil {
		e := CreatePaymentResponse{}
		return e, err
	}
	c := CreatePaymentResponse{PaymentID: p.ID}
	return c, nil
}

// UpdatePayment updates (PUT) an already existing payment based on the original payment's ID and and a provided payment json file
func (r *paymentService) UpdatePayment(req UpdatePaymentRequest) (UpdatePaymentResponse, error) {
	pa := &Payment{}
	id, err := uuid.FromString(req.PaymentID)
	if err != nil {
		e := UpdatePaymentResponse{}
		var ErrAcc = errors.New("err: Could not parse UUID to Update")
		cErr := errors.New(ErrAcc.Error() + err.Error())
		return e, cErr
	}

	p := req.Payment
	p.ID = id
	//if err := r.db.Debug().Model(&p).Where("id = ?", id).Find(&pa).Error; err != nil {
	if err := r.db.Model(&p).Where("id = ?", id).Find(&pa).Error; err != nil {
		e := UpdatePaymentResponse{}
		return e, err
	}

	err = r.db.Model(&p).Save(&p).Error
	//err = r.db.Debug().Model(&p).Save(&p).Error
	if err != nil {
		e := UpdatePaymentResponse{}
		return e, err
	}
	c := UpdatePaymentResponse{PaymentID: id}
	return c, nil
}

// DeletePayment soft deletes (DELETE) an existing payment entry based on a provided payment ID.
// A soft delete is the act of populating the DeletedAt field from the Payments table with a timestamp
// which tracks the time the opreation was performed and excludes the entry from other operations
func (r *paymentService) DeletePayment(id uuid.UUID) (*time.Time, error) {
	p := &Payment{}
	delTime := new(time.Time)
	//if err := r.db.Debug().Model(p).Where("id = ?", id).Find(p).Error; err != nil {
	if err := r.db.Model(p).Where("id = ?", id).Find(p).Error; err != nil {
		return delTime, err
	}
	// Delete payment by ID `Soft Delete`
	//if err := r.db.Debug().Model(p).Where("id = ?", id).Delete(p).Error; err != nil {
	if err := r.db.Model(p).Where("id = ?", id).Delete(p).Error; err != nil {
		return delTime, err
	}
	//if err := r.db.Debug().Unscoped().Where("id = ?", id).Find(p).Error; err != nil {
	if err := r.db.Unscoped().Where("id = ?", id).Find(p).Error; err != nil {
		return delTime, err
	}
	delTime = p.DeletedAt
	return delTime, nil
}

// GetListOfPayments retrieves a list of all the commited payments
func (r *paymentService) GetListPayments() ([]Payment, error) {
	var payments []Payment
	err := r.db.Find(&payments).Error
	//err := r.db.Debug().Find(&payments).Error
	if err != nil {
		return nil, err
	}
	for i, p := range payments {
		payments[i], _ = r.GetPayment(p.ID.String())
	}
	return payments, nil
}
