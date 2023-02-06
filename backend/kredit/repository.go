package kredit

import (
	"fmt"
	"rema/kredit/model"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	getDataChecklistPencairan(page int, length int, fields []string, values []interface{}) ([]model.RequestKreditData, error)
	getTotalDataChecklistPencairan(fields []string, values []interface{}) (int, error)
	UpdateChecklistPencairan(fields []string, values []interface{}) error
	getDataReport(page int, length int, fields []string, values []interface{})([]model.RequestKreditData, error)
	getTotalDataReport( fields []string, values []interface{})(int, error)
	GetCompany() ([]model.MstCompanyTab, error)
	GetBranch() ([]model.BranchTab, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) GetCompany() ([]model.MstCompanyTab, error){
	var data []model.MstCompanyTab
	res := r.db.Find(&data)
	if res.Error != nil{
		return []model.MstCompanyTab{}, res.Error
	}
	return data, nil
}
func (r *repository) GetBranch() ([]model.BranchTab, error){
	var data []model.BranchTab
	res := r.db.Order("code").Find(&data)
	if res.Error != nil{
		return []model.BranchTab{}, res.Error
	}
	return data, nil
}
func(r *repository) getDataReport(page int, length int, fields []string, values []interface{})([]model.RequestKreditData, error){
	
	var rd []model.RequestKreditData
	offset := (page-1)* length
	tx := r.db.Table("customer_data_tab").Limit(length).Offset(offset).
	Select("loan_data_tab.custcode, ppk, name, channeling_company, drawdown_date, loan_amount, loan_period,interest_effective, approval_status").
	Joins("join loan_data_tab on loan_data_tab.custcode=customer_data_tab.custcode").
	Joins("join mst_company_tab as a on customer_data_tab.channeling_company = a.company_short_name").
	Joins("join branch_tab on branch_tab.code = loan_data_tab.branch").
	Where(strings.Join(fields, " AND "), values...).
	Order("drawdown_date desc").
	Find(&rd)
	if tx.Error != nil {
		fmt.Println("error get data")
		return []model.RequestKreditData{}, tx.Error
	}
	return rd, nil
}
func (r *repository) getTotalDataReport( fields []string, values []interface{})(int, error){
	type resultFormat struct{
		Total int
	}
	var result resultFormat
	tx := r.db.Table("customer_data_tab").
	Select("count(loan_data_tab.custcode) as total").
	Joins("join loan_data_tab on loan_data_tab.custcode=customer_data_tab.custcode").
	Joins("join mst_company_tab as a on customer_data_tab.channeling_company = a.company_short_name").
	Joins("join branch_tab on branch_tab.code = loan_data_tab.branch").
	Where(strings.Join(fields, " AND "), values...).
	Scan(&result)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return result.Total, nil
}
func (r *repository) getDataChecklistPencairan(page int, length int, fields []string, values []interface{}) ([]model.RequestKreditData, error){
	var rd []model.RequestKreditData
	offset := (page-1)* length
	tx := r.db.Table("customer_data_tab").Limit(length).Offset(offset).
	Select("loan_data_tab.custcode, ppk, name, channeling_company, drawdown_date, loan_amount, loan_period,interest_effective, approval_status").
	Joins("join loan_data_tab on loan_data_tab.custcode=customer_data_tab.custcode").
	Where(strings.Join(fields, " AND "), values...).
	Order("drawdown_date desc").
	Find(&rd)
	if tx.Error != nil {
		fmt.Println("error get data")
		return []model.RequestKreditData{}, tx.Error
	}
	return rd, nil
}

func (r *repository) getTotalDataChecklistPencairan(fields []string, values []interface{}) (int, error){
	type resultFormat struct{
		Total int
	}
	var result resultFormat
	tx := r.db.Table("customer_data_tab").
	Select("count(loan_data_tab.custcode) as total").
	Joins("join loan_data_tab on loan_data_tab.custcode=customer_data_tab.custcode").
	Where(strings.Join(fields, " AND "), values...).
	
	Scan(&result)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return result.Total, nil
}
func (r *repository) UpdateChecklistPencairan(fields []string, values []interface{}) error{
	res := r.db.Model(&model.CustomerDataTab{}).Where(strings.Join(fields, " AND "), values...).
	Update("approval_status","0")
	if res.Error != nil{
		return res.Error
	}
	// sql :=r.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 	return tx.Model(&model.CustomerDataTab{}).Where(strings.Join(fields, " AND "), values...).Update("approval_status","0")
	// }) 
	// fmt.Println(sql)
	return nil
}