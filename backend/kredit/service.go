package kredit

import (
	"fmt"
	"math"
	"rema/kredit/model"
	"strconv"
)

type Service interface {
	// ServiceChecklistPencairan(page string, length string) ([]model.RequestKreditData, error)
	ServiceChecklistPencairan(page string, length string) (ResultChecklistPencairan, error)
	UpdateChecklistPencairan(datas UpdateChecklistPencairan) error

	ServiceDataReport(page string, length string, start string, end string, company string, cabang string, approval string) (ResultChecklistPencairan, error)
	GetCompanyService() ([]model.MstCompanyTab, error)
	GetBranchService() ([]model.BranchTab, error)
}
type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}
// func (s *service) ServiceChecklistPencairan(page string, length string) ([]model.RequestKreditData, error) {
// 	valPage, _ := strconv.Atoi(page)
// 	valLength, _ := strconv.Atoi(length)
// 	// data, err := s.repo.getDataChecklistPencairan(valPage, valLength)
// 	return data, err
// }
func (s *service) ServiceChecklistPencairan(page string, length string) (ResultChecklistPencairan, error) {
	valPage, _ := strconv.Atoi(page)
	valLength, _ := strconv.Atoi(length)

	fields := []string{}
	values := []interface{}{}
	fields = append(fields,"approval_status=?")
	values = append(values, "9")

	data, err := s.repo.getDataChecklistPencairan(valPage, valLength, fields, values)
	fmt.Println(data)
	if err!= nil{
		return ResultChecklistPencairan{}, err
	}
	total, err:=s.repo.getTotalDataChecklistPencairan( fields, values)
	if err!= nil{
		return ResultChecklistPencairan{}, err
	}
	total_page := int(math.Ceil(float64(total)/float64(valLength)))
	return ResultChecklistPencairan{data: data, total: total,total_halaman:total_page }, err
}
func (s *service) UpdateChecklistPencairan(datas UpdateChecklistPencairan) error{
	fields := []string{}
	values := []interface{}{}
	list := []string{}
	for _,data := range datas.Data{
		list = append(list, data.Id)
	}
	// idList := strings.Trim(strings.Join(strings.Split(fmt.Sprint(list), " "), ","), "[]")
	fields = append(fields,"custcode in (?)")
	values = append(values,list)
	err := s.repo.UpdateChecklistPencairan(fields, values)
	return err
}

func (s *service) ServiceDataReport(page string, length string, start string, end string, company string, cabang string, approval string) (ResultChecklistPencairan, error) {
	valPage, _ := strconv.Atoi(page)
	valLength, _ := strconv.Atoi(length)
	
	fields := []string{}
	values := []interface{}{}
	
	if start != ""{
		fields = append(fields, "drawdown_date >= ?")
		values = append(values, start)
	}
	if end != "" {
		fields = append(fields, "drawdown_date <= ?")
		values = append(values, end)
	}
	if company != ""{
		fields = append(fields, "company_code = ?")
		values = append(values, company)
	}
	
	if cabang != ""{
		fields = append(fields, "branch = ?")
		values = append(values, cabang)
	}
	if approval == ""{
		fields = append(fields,"approval_status in (?)")
		values = append(values,[]interface{}{"1","0"})

	}else{
		fields = append(fields,"approval_status = ?")
		values = append(values, approval)
	}
	
	data, err := s.repo.getDataReport(valPage, valLength, fields, values)
	fmt.Println(data)
	if err!= nil{
		return ResultChecklistPencairan{}, err
	}
	total, err:=s.repo.getTotalDataReport( fields, values)
	if err!= nil{
		return ResultChecklistPencairan{}, err
	}
	total_page := int(math.Ceil(float64(total)/float64(valLength)))
	return ResultChecklistPencairan{data: data, total: total,total_halaman:total_page }, nil
}
func (s *service) GetCompanyService() ([]model.MstCompanyTab, error){
	data, err := s.repo.GetCompany()
	return data, err
}
func (s *service) GetBranchService() ([]model.BranchTab, error){
	data, err := s.repo.GetBranch()
	return data, err
}