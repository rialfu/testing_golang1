package model

import (
	"time"
)

type CustomerDataTab struct {
	Custcode           string    `gorm:"column:custcode;type:varchar(25);NOT NULL"`
	Ppk                string    `gorm:"column:ppk;type:varchar(20)"`
	Name               string    `gorm:"column:name;type:varchar(100)"`
	Address1           string    `gorm:"column:address1;type:varchar(40)"`
	Address2           string    `gorm:"column:address2;type:varchar(40)"`
	City               string    `gorm:"column:city;type:varchar(100)"`
	Zip                string    `gorm:"column:zip;type:varchar(6)"`
	BirthPlace         string    `gorm:"column:birth_place;type:varchar(20)"`
	BirthDate          time.Time `gorm:"column:birth_date;type:date"`
	IDType             int8       `gorm:"column:id_type;type:smallint"`
	IDNumber           string    `gorm:"column:id_number;type:varchar(30)"`
	MobileNo           string    `gorm:"column:mobile_no;type:varchar(20)"`
	DrawdownDate       time.Time `gorm:"column:drawdown_date;type:date"`
	TglPkChannelling   time.Time `gorm:"column:tgl_pk_channelling;type:date"`
	MotherMaidenName   string    `gorm:"column:mother_maiden_name;type:varchar(100)"`
	ChannelingCompany string    `gorm:"column:channeling_company;type:varchar(100)"`
	ApprovalStatus     string    `gorm:"column:approval_status;type:varchar(2)"`
}
func (m *CustomerDataTab) TableName() string {
	return "customer_data_tab"
}