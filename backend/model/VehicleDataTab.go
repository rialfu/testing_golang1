package model

import (
	"time"
)

type VehicleDataTab struct {
	Custcode       string    `gorm:"column:custcode;NOT NULL; type: varchar(25);"`
	Brand          string       `gorm:"column:brand;type:varchar(80)"`
	Type           string    `gorm:"column:type;type: varchar(100)"`
	Year           string    `gorm:"column:year;type: varchar(4)"`
	Golongan       int8       `gorm:"column:golongan;type: smallint"`
	Jenis          string    `gorm:"column:jenis;type: varchar(200)"`
	Status         int8       `gorm:"column:status;type: smallint"`
	Color          string    `gorm:"column:color;type: varchar(20)"`
	PoliceNo       string    `gorm:"column:police_no;type: varchar(20)"`
	EngineNo       string    `gorm:"column:engine_no;type: varchar(20)"`
	ChasisNo       string    `gorm:"column:chasis_no;ype: varchar(20)"`
	Bpkb           string    `gorm:"column:bpkb;type: varchar(20)"`
	RegisterNo     string    `gorm:"column:register_no;type: varchar(50)"`
	Stnk           string    `gorm:"column:stnk;type: varchar(50)"`
	StnkAddress1   string    `gorm:"column:stnk_address1;type: varchar(40)"`
	StnkAddress2   string    `gorm:"column:stnk_address2;type: varchar(40)"`
	StnkCity       string    `gorm:"column:stnk_city;type: varchar(20)"`
	DealerID       int8       `gorm:"column:dealer_id;type:smallint"`
	Inputdate      time.Time `gorm:"column:inputdate;type: timestamp"`
	Inputby        string    `gorm:"column:inputby;type: varchar(50)"`
	Lastmodified   time.Time `gorm:"column:lastmodified;type: timestamp;"`
	Modifiedby     string    `gorm:"column:modifiedby;type: varchar(50)"`
	TglStnk        time.Time `gorm:"column:tgl_stnk;type: timestamp"`
	TglBpkb        time.Time `gorm:"column:tgl_bpkb;type: timestamp"`
	TglPolis       time.Time `gorm:"column:tgl_polis;type: timestamp"`
	PolisNo        string    `gorm:"column:polis_no;type: varchar(17)"`
	CollateralID   int64     `gorm:"column:collateral_id;type:bigint"`
	Ketagunan      string    `gorm:"column:ketagunan; type:varchar(16)"`
	AgunanLbu      string    `gorm:"column:agunan_lbu;type: varchar(10)"`
	Dealer         string    `gorm:"column:dealer;type:varchar(100)"`
	AddressDealer1 string    `gorm:"column:address_dealer1;type:varchar(100)"`
	AddressDealer2 string    `gorm:"column:address_dealer2;type:varchar(100)"`
	CityDealer     string    `gorm:"column:city_dealer;type:varchar(100)"`
}

func (m *VehicleDataTab) TableName() string {
	return "vehicle_data_tab"
}