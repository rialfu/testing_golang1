package model

import (
	"time"
)

type LoanDataTab struct {
	Custcode             string    `gorm:"column:custcode;type:varchar(25);NOT NULL"`
	Branch               string    `gorm:"column:branch;type:varchar(50)"`
	Otr                  float64   `gorm:"column:otr;type:decimal(16,2)"`
	DownPayment          float64    `gorm:"column:down_payment;type:decimal(16,2)"`
	LoanAmount           float64    `gorm:"column:loan_amount;type:decimal(16,2)"`
	LoanPeriod           string    `gorm:"column:loan_period;type:varchar(6)"`
	InterestType         int8       `gorm:"column:interest_type;type:smallint"`
	InterestFlat         float32   `gorm:"column:interest_flat;type:real"`
	InterestEffective    float32   `gorm:"column:interest_effective;type:real"`
	EffectivePaymentType int8       `gorm:"column:effective_payment_type;type:smallint"`
	AdminFee             float64    `gorm:"column:admin_fee;type:decimal(16,2)"`
	MonthlyPayment       float64    `gorm:"column:monthly_payment;type:decimal(16,2)"`
	InputDate            time.Time `gorm:"column:input_date;type:timestamp"`
	LastModified         time.Time `gorm:"column:last_modified;type:timestamp"`
	ModifiedBy           string    `gorm:"column:modified_by;type:varchar(20)"`
	Inputdate            time.Time `gorm:"column:inputdate;type:timestamp"`
	Inputby              string    `gorm:"column:inputby;type:varchar(50)"`
	Lastmodified         time.Time `gorm:"column:lastmodified;type:timestamp"`
	Modifiedby           string    `gorm:"column:modifiedby;type:varchar(50)"`
}

func (m *LoanDataTab) TableName() string {
	return "loan_data_tab"
}