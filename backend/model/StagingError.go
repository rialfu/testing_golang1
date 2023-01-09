package model

type StagingError struct {
	ID           int64  `gorm:"column:id;type:bigserial;NOT NULL"`
	SeReff       string `gorm:"column:se_reff;type:varchar(50);"`
	SeCreateDate string `gorm:"column:se_create_date;type:timestamp"`
	BranchCode   string `gorm:"column:branch_code;type:varchar(50)"`
	Company      string `gorm:"column:company;type:varchar(50)"`
	Ppk          string `gorm:"column:ppk;type:varchar(50)"`
	Name         string `gorm:"column:name;type:varchar(50)"`
	ErrorDesc    string `gorm:"column:error_desc;type:text"`
}

func (m *StagingError) TableName() string {
	return "staging_error"
}