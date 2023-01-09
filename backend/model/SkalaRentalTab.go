package model

import (
	"time"
)

type SkalaRentalTab struct {
	Custcode      string    `gorm:"column:custcode;type:varchar(25);NOT NULL"`
	Counter       int8       `gorm:"column:counter;type:smallint"`
	Osbalance     float64    `gorm:"column:osbalance;type:decimal(12,2)"`
	EndBalance    float64    `gorm:"column:end_balance;type:decimal(12,2)"`
	DueDate 		time.Time	`gorm:"column:due_date;type:timestamp"`
	EffRate       float32   `gorm:"column:eff_rate;type:float"`
	Rental        float64    `gorm:"column:rental;type:decimal(12,2)"`
	Principle     float64    `gorm:"column:principle;type:decimal(12,2)"`
	Interest      float64    `gorm:"column:interest;type:decimal(12,2)"`
	Inputdate     time.Time `gorm:"column:inputdate; type:date"`
	Inputby       string    `gorm:"column:inputby;type:varchar(50)"`
	Lastmodified  time.Time `gorm:"column:lastmodified; type:timestamp"`
	Modifiedby    string    `gorm:"column:modifiedby;type:varchar(50)"`
	PaymentDate   time.Time `gorm:"column:payment_date;type:timestamp"`
	Penalty       float64    `gorm:"column:penalty;type:decimal(12,2)"`
	PaymentAmount float64    `gorm:"column:payment_amount;type:decimal(12,2)"`
	PaymentType   int8       `gorm:"column:payment_type;type:smallint"`
}

func (m *SkalaRentalTab) TableName() string {
	return "skala_rental_tab"
}