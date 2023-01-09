package model

type BranchTab struct {
	Code        string `gorm:"column:code; type:varchar(50);NOT NULL"`
	Description string `gorm:"column:description; type:varchar(50)"`
	Address1    string `gorm:"column:address1; type:varchar(50) "`
	Address2    string `gorm:"column:address2; type:varchar(50)"`
	City        string `gorm:"column:city; type:varchar(50)"`
	Zip         string `gorm:"column:zip; type:varchar(6)"`
	Phone       string `gorm:"column:phone; type:varchar(15)"`
	Fax         string `gorm:"column:fax; type:varchar(15)"`
}

func (m *BranchTab) TableName() string {
	return "branch_tab"
}