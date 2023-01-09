package model

import (
	"time"
)

type StagingCustomer struct {
	ID                              int64     `gorm:"column:id;NOT NULL;type: bigint"`
	ScReff                          string    `gorm:"column:sc_reff; type: varchar(80)"`
	ScCreateDate                    time.Time `gorm:"column:sc_create_date; type: timestamp"`
	ScBranchCode                    string    `gorm:"column:sc_branch_code;type: varchar(80)"`
	ScCompany                       string    `gorm:"column:sc_company; type: varchar(80)"`
	ScFlag                          string    `gorm:"column:sc_flag; type:char;size:1"`
	CustomerPpk                     string    `gorm:"column:customer_ppk; type: varchar(80)"`
	CustomerName                    string    `gorm:"column:customer_name; type: varchar(80)"`
	CustomerAddress1                string    `gorm:"column:customer_address1; type: varchar(80)"`
	CustomerAddress2                string    `gorm:"column:customer_address2; type: varchar(80)"`
	CustomerCity                    string    `gorm:"column:customer_city; type: varchar(80)"`
	CustomerZip                     string    `gorm:"column:customer_zip; type: varchar(80)"`
	CustomerBirthPlace              string    `gorm:"column:customer_birth_place; type: varchar(80)"`
	CustomerBirthDate               string    `gorm:"column:customer_birth_date; type: varchar(80)"`
	CustomerIDType                  string    `gorm:"column:customer_id_type; type: varchar(80)"`
	CustomerIDNumber                string    `gorm:"column:customer_id_number; type: varchar(80)"`
	CustomerMobileNo                string    `gorm:"column:customer_mobile_no; type: varchar(80)"`
	CustomerMotherMaidenName        string    `gorm:"column:customer_mother_maiden_name;type: varchar(80)"`
	LoanOtr                         string    `gorm:"column:loan_otr;type: varchar(80)"`
	LoanDownPayment                 string    `gorm:"column:loan_down_payment;type: varchar(80)"`
	LoanLoanAmountChanneling        string    `gorm:"column:loan_loan_amount_channeling;type: varchar(80)"`
	LoanLoanPeriodChanneling        string    `gorm:"column:loan_loan_period_channeling;type: varchar(80)"`
	LoanInterestFlatChanneling      string    `gorm:"column:loan_interest_flat_channeling;type: varchar(80)"`
	LoanInterestEffectiveChanneling string    `gorm:"column:loan_interest_effective_channeling;type: varchar(80)"`
	LoanEffectivePaymentType        string    `gorm:"column:loan_effective_payment_type;type: varchar(80)"`
	LoanMonthlyPaymentChanneling    string    `gorm:"column:loan_monthly_payment_channeling;type: varchar(80)"`
	LoanTglPk                       string    `gorm:"column:loan_tgl_pk;type: varchar(80)"`
	LoanTglPkChanneling             string    `gorm:"column:loan_tgl_pk_channeling;type: varchar(80)"`
	CollateralTypeID                string    `gorm:"column:collateral_type_id;type: varchar(80)"`
	VehicleJenisProduk              string    `gorm:"column:vehicle_jenis_produk;type: varchar(80)"`
	VehicleBrand                    string    `gorm:"column:vehicle_brand;type: varchar(80)"`
	VehicleType                     string    `gorm:"column:vehicle_type;type: varchar(150)"`
	VehicleYear                     string    `gorm:"column:vehicle_year;type: varchar(80)"`
	VehicleJenis                    string    `gorm:"column:vehicle_jenis;type: varchar(80)"`
	VehicleStatus                   string    `gorm:"column:vehicle_status;type: varchar(80)"`
	VehicleColor                    string    `gorm:"column:vehicle_color;type: varchar(80)"`
	VehiclePoliceNo                 string    `gorm:"column:vehicle_police_no;type: varchar(80)"`
	VehicleEngineNo                 string    `gorm:"column:vehicle_engine_no;type: varchar(80)"`
	VehicleChasisNo                 string    `gorm:"column:vehicle_chasis_no;type: varchar(80)"`
	VehicleBpkb                     string    `gorm:"column:vehicle_bpkb;type: varchar(80)"`
	VehicleStnk                     string    `gorm:"column:vehicle_stnk;type: varchar(80)"`
	VehicleDealer                   string    `gorm:"column:vehicle_dealer;type: varchar(80)"`
	VehicleAddressDealer1           string    `gorm:"column:vehicle_address_dealer1;type: varchar(80)"`
	VehicleAddressDealer2           string    `gorm:"column:vehicle_address_dealer2;type: varchar(80)"`
	VehicleCityDealer               string    `gorm:"column:vehicle_city_dealer;type: varchar(80)"`
	VehicleTglStnk                  string    `gorm:"column:vehicle_tgl_stnk;type: varchar(80)"`
	VehicleTglBpkb                  string    `gorm:"column:vehicle_tgl_bpkb;type: varchar(80)"`
	VehicleUtilizationPurpose       string    `gorm:"column:vehicle_utilization_purpose;type: varchar(80)"`
	LoanDrawdownKolektibilitas      string    `gorm:"column:loan_drawdown_kolektibilitas;type: varchar(2)"`
	VehicleDealerID                 string    `gorm:"column:vehicle_dealer_id;type: varchar(10)"`
	
}
func (m *StagingCustomer) TableName() string {
	return "staging_customer"
}