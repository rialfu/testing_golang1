package model

type MstCompanyTab struct {
	ID               int64   `json:"id" gorm:"column:id;type:bigint;AUTO_INCREMENT;NOT NULL"`
	CompanyCode      string  `json:"company_code" gorm:"column:company_code;type:varchar(5)"`
	CompanyShortName string  `json:"company_short_name" gorm:"column:company_short_name;type:varchar(50)"`
	CompanyName      string  `json:"company_name" gorm:"column:company_name;type:varchar(200)"`
	CompanyAddress1  string  `json:"address1" gorm:"column:company_address1;type:varchar(200)"`
	CompanyAddress2  string  `json:"address2" gorm:"column:company_address2;type:varchar(200)"`
	CompanyCity      string  `json:"city" gorm:"column:company_city;type:varchar(100)"`
	CompanyPhone     string  `json:"phone" gorm:"column:company_phone;type:varchar(50)"`
	CompanyFax       string  `json:"fax" gorm:"column:company_fax;type:varchar(50)"`
	BungaEffMin      float32 `json:"bunga_eff_min" gorm:"column:bunga_eff_min;type:real"`
	BungaEffMax      float32 `json:"bunga_eff_max" gorm:"column:bunga_eff_max;type:real"`
	BungaFlatMin     float32 `json:"bunga_flat_min" gorm:"column:bunga_flat_min;type:real"`
	BungaFlatMax     float32 `json:"bunga_flat_max" gorm:"column:bunga_flat_max;type:real"`
	LaMin            float64 `json:"la_min" gorm:"column:la_min; type:decimal(16,2);"`
	LaMax            float64 `json:"la_max" gorm:"column:la_max; type:decimal(16,2);"`
	PeriodeMin       string  `json:"periode_min" gorm:"column:periode_min;type:varchar(10)"`
	PeriodeMax       string  `json:"periode_max" gorm:"column:periode_max;type:varchar(10)"`
}

func (m *MstCompanyTab) TableName() string {
	return "mst_company_tab"
}