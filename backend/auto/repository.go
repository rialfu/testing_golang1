package auto

import (
	"fmt"
	"math"
	"regexp"
	"rema/kredit/model"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type AutoRepository interface {
	GenerateValidatePengajuanKredit()
	GenerateSkalaAngsuran() 
}
type repository struct {
	db *gorm.DB
}
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) GenerateValidatePengajuanKredit()  {
	
	type Result struct {
		model.StagingCustomer
		jumlah_ppk int
		jumlah_company int
		jumlah_branch int
		jumlah_chasis int
		jumlah_engine int
		CompanyCode string
	}
	var StagingCustomer []Result
	
	// currentTime := time.Now()
	// r.db.Raw(`
	// select sc.*, jumlah_ppk, jumlah_company, jumlah_branch, jumlah_engine, jumlah_chasis
	// from staging_customer as sc 
	// join(
	// 	select sc.id,count(cdt.ppk) as jumlah_ppk from staging_customer sc 
	// 	left join customer_data_tab as cdt on 
	// 	sc.customer_ppk=cdt.ppk
	// 	where sc_flag='0'
	// 	group by sc.id
	// ) as child1 on child1.id=sc.id
	// join(
	// 	select sc.id,count(mct.company_short_name) as jumlah_company from staging_customer sc 
	// 	left join mst_company_tab as mct on 
	// 	sc.sc_company=mct.company_short_name
	// 	where sc_flag='0'
	// 	group by sc.id
	// ) as child2 on child2.id=sc.id


	// join(
	// 	select sc.id,count(bt.code) as jumlah_branch from staging_customer sc 
	// 	left join branch_tab as bt on 
	// 	sc.sc_branch_code=bt.code
	// 	where sc_flag='0'
	// 	group by sc.id
	// ) as child3 on child3.id=sc.id
	// join(
	// 	select sc.id,count(vdt.engine_no) as jumlah_engine, count(vdt1.chasis_no) as jumlah_chasis from staging_customer sc 
	// 	left join vehicle_data_tab as vdt on 
	// 	sc.vehicle_engine_no=vdt.engine_no
	// 	left join vehicle_data_tab as vdt1 on 
	// 	sc.vehicle_chasis_no=vdt1.chasis_no
	// 	where sc_flag='0'
	// 	group by sc.id
	// ) as child4 on child4.id=sc.id
	// `).
	
	// r.db.Raw(`
	// select 
	// 	sc.*, jumlah_ppk, 
	// 	jumlah_company, 
	// 	jumlah_branch, 
	// 	jumlah_engine, 
	// 	jumlah_chasis,
	// 	company_code
	// from staging_customer as sc 
	// join(
	// 	select 
	// 		sc.id, 
	// 		count(cdt.ppk) as jumlah_ppk ,
	// 		count(mct.company_short_name) as jumlah_company,
	// 		count(bt.code) as jumlah_branch,
	// 		count(vdt.engine_no) as jumlah_engine, 
	// 		count(vdt1.chasis_no) as jumlah_chasis
	// 	from staging_customer sc 
	// 	left join customer_data_tab as cdt on 
	// 	sc.customer_ppk=cdt.ppk
	// 	left join mst_company_tab as mct on 
	// 	sc.sc_company=mct.company_short_name
	// 	left join branch_tab as bt on 
	// 	sc.sc_branch_code=bt.code
	// 	left join vehicle_data_tab as vdt on 
	// 	sc.vehicle_engine_no=vdt.engine_no
	// 	left join vehicle_data_tab as vdt1 on 
	// 	sc.vehicle_chasis_no=vdt1.chasis_no
	// 	where sc_flag='0' and sc_create_date = ?
	// 	group by sc.id
	// ) as child on child.id=sc.id
	// left join mst_company_tab as mct on 
	// sc.sc_company=mct.company_short_name
	// `, time.Now().Format("2006-01-02")).
	r.db.Raw(`
	select 
		sc.*, jumlah_ppk, 
		jumlah_company, 
		jumlah_branch, 
		jumlah_engine, 
		jumlah_chasis,
		company_code
	from (
		select 
			sc.id, 
			count(cdt.ppk) as jumlah_ppk ,
			count(mct.company_short_name) as jumlah_company,
			count(bt.code) as jumlah_branch,
			count(vdt.engine_no) as jumlah_engine, 
			count(vdt1.chasis_no) as jumlah_chasis
		from staging_customer sc 
		left join customer_data_tab as cdt on 
			sc.customer_ppk=cdt.ppk
		left join mst_company_tab as mct on 
			sc.sc_company=mct.company_short_name
		left join branch_tab as bt on 
			sc.sc_branch_code=bt.code
		left join vehicle_data_tab as vdt on 
			sc.vehicle_engine_no=vdt.engine_no
		left join vehicle_data_tab as vdt1 on 
			sc.vehicle_chasis_no=vdt1.chasis_no
		where sc_flag='0' and sc_create_date = ?
		group by sc.id
	) as child join staging_customer as sc on child.id=sc.id
	left join mst_company_tab as mct on 
	sc.sc_company=mct.company_short_name
	`, time.Now().Format("2006-01-02")).
	Find(&StagingCustomer)
	fmt.Println("run")
	for _, el := range StagingCustomer{
		var reason []string
		if el.jumlah_ppk > 0{
			reason = append(reason, "Terjadi duplikasi PPK")
		}
		if el.jumlah_branch == 0 {
			reason = append(reason, "Branch tidak terdaftar")
		}
		if el.jumlah_company == 0 {
			reason = append(reason, "Company tidak terdaftar")
		}
		TglPk, err := time.Parse("2006-01-02", el.LoanTglPk)
		if(err != nil){
			fmt.Println("format tgl PK tidak sesuai")
		}

		monthTglPk := int(TglPk.Month())
		yearTglPk := TglPk.Year()
		if int(time.Now().Month()) != monthTglPk || time.Now().Year() !=  yearTglPk {
			reason = append(reason, "Tgl PK berbeda dari bulan, tahun ini")
		}
		if el.CustomerIDType == "1" && strings.TrimSpace(el.CustomerIDNumber) == ""{
			reason = append(reason, "ID Number harus terisi")
		}
		regex := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]`)
		if regex.MatchString(strings.TrimSpace(el.CustomerName)) {
			reason = append(reason, "Nama debitur mengandung special character")
		}
		if strings.TrimSpace(el.VehicleBpkb) == "" {
			reason = append(reason, "BPKB tidak terisi")
		}
		if strings.TrimSpace(el.VehicleStnk) == "" {
			reason = append(reason, "STNK tidak terisi")
		}
		if strings.TrimSpace(el.VehicleEngineNo) == "" {
			reason = append(reason, "Nomor mesin tidak terisi")
		}
		if strings.TrimSpace(el.VehicleChasisNo) == "" {
			reason = append(reason, "Nomor chasis tidak terisi")
		}
		if el.jumlah_chasis > 0 {
			reason = append(reason, "Nomor chasis terduplikasi")
		}
		if el.jumlah_engine > 0 {
			reason = append(reason, "Nomor engine terduplikasi")
		}
		
		var sc model.StagingCustomer
		if len(reason) > 0{
			se :=model.StagingError{
				SeReff: el.ScReff,
				SeCreateDate: time.Now().Format(time.RFC1123Z),
				BranchCode: el.ScBranchCode,
				Company: el.ScCompany,
				Ppk: el.CustomerPpk,
				Name: el.CustomerName,
				ErrorDesc: strings.Join(reason[:], ", "),
			}
			tx := r.db.Begin()
			res := tx.Model(&sc).Where("id", el.ID).Update("sc_flag","8")
			fmt.Println("sukses?")
			if res.Error != nil{
				tx.Rollback()
				continue
			}
			
			res = tx.Create(&se)
			if res.Error != nil{
				tx.Rollback()
				continue
			}
			tx.Commit()
			
		}else{
			BirthDate, err := time.Parse("2006-01-02 15:04:05", el.CustomerBirthDate)
			if(err != nil){
				fmt.Println("birth date form")
				continue
			}

			idType, err := strconv.ParseInt(el.CustomerIDType, 10, 8)
			if(err != nil){
				continue
			}
			otr, err := strconv.ParseFloat(el.LoanOtr, 64)
			if(err != nil){
				continue
			}
			DrawdownDate, err := time.Parse("2006-01-02", el.LoanTglPk)
			if(err != nil){
				continue
			}
			tglPkChanneling, err := time.Parse("2006-01-02", el.LoanTglPkChanneling)
			if(err != nil){
				continue
			}
			downPayment, err := strconv.ParseFloat(el.LoanDownPayment, 64)
			if(err != nil){
				continue
			}
			amount, err := strconv.ParseFloat(el.LoanLoanAmountChanneling, 64)
			EffectivePaymentType, err := strconv.ParseInt(el.LoanEffectivePaymentType, 10, 8)
			interestFlat, err := strconv.ParseFloat(el.LoanInterestFlatChanneling, 32)
			InterestEffective, err := strconv.ParseFloat(el.LoanInterestEffectiveChanneling, 32)
			MonthlyPayment, err := strconv.ParseFloat(el.LoanMonthlyPaymentChanneling, 64)
			VehicleStatus, err := strconv.ParseInt(el.LoanEffectivePaymentType, 10, 8)
			VehicleTglStnk, err := time.Parse("2006-01-02 15:04:05", el.VehicleTglStnk)
			VehicleTglBkpb, err := time.Parse("2006-01-02 15:04:05", el.VehicleTglBpkb)
			var VehicleDealerID int64 = 0
			if el.VehicleDealerID != ""{
				VehicleDealerID, err = strconv.ParseInt(el.VehicleDealerID, 10, 8)
			}
			
			// id := custom.RandStringBytes(8)
			now := time.Now()
			var AppCustCode string = "006";
			var NewCustCode string
			getDate := fmt.Sprintf("%s", now.Format("200601"))
			NewCustCode = AppCustCode  +el.CompanyCode+ getDate
			customer:= model.CustomerDataTab{
				Custcode: NewCustCode,
				Ppk: strings.TrimSpace(el.CustomerPpk),
				Name: strings.TrimSpace(el.CustomerName),
				Address1: el.CustomerAddress1,
				Address2: el.CustomerAddress2,
				City: el.CustomerCity,
				Zip: el.CustomerZip,
				BirthPlace: el.CustomerBirthPlace,
				BirthDate: BirthDate,
				IDType: int8(idType),
				IDNumber: el.CustomerIDNumber,
				MobileNo: el.CustomerMobileNo,
				DrawdownDate: DrawdownDate,
				TglPkChannelling: tglPkChanneling,
				MotherMaidenName: el.CustomerMotherMaidenName,
				ChannelingCompany: el.ScCompany,
				ApprovalStatus: "9",
			}
			
			loan:= model.LoanDataTab{
				Custcode: NewCustCode,
				Branch: el.ScBranchCode,
				Otr: otr,
				DownPayment: downPayment,
				LoanAmount: amount,
				LoanPeriod: el.LoanLoanPeriodChanneling,
				// InterestType: 0,
				InterestFlat: float32(interestFlat),
				InterestEffective: float32(InterestEffective),
				EffectivePaymentType:int8(EffectivePaymentType),
				// AdminFee: 0,
				MonthlyPayment: MonthlyPayment,
				InputDate: now,
				Inputdate: now,
				Inputby: "System",
			}
			vehicle := model.VehicleDataTab{
				Custcode: NewCustCode,
				Brand: el.VehicleBrand,
				Type:strings.TrimSpace(el.VehicleType),
				Year:el.VehicleYear,
				// Golongan: el.vehicleg,
				Jenis: el.VehicleJenis,
				Status: int8(VehicleStatus),
				Color:el.VehicleColor,
				PoliceNo:el.VehiclePoliceNo,
				EngineNo:el.VehicleEngineNo,
				ChasisNo:el.VehicleChasisNo,
				Bpkb: el.VehicleBpkb,
				// RegisterNo: el.vehiclre,
				Stnk: el.VehicleStnk,
				// StnkAddress1: el.VehicleAddressDealer1,
				// StnkAddress2: el.VehicleAddressDealer2,
				// StnkCity: el.VehicleCityDealer,
				DealerID: int8(VehicleDealerID),
				Inputdate: now,
				Inputby: "System",
				TglStnk: VehicleTglStnk,
				TglBpkb: VehicleTglBkpb,
				// TglPolis: ,
				PolisNo: el.VehiclePoliceNo,
				// CollateralID: el.CollateralTypeID,
				// Ketagunan: ,
				// AgunanLbu: e,
				Dealer: el.VehicleDealer,
				AddressDealer1: el.VehicleAddressDealer1,
				AddressDealer2: el.VehicleAddressDealer2,
				CityDealer: el.VehicleCityDealer,
			}
			type SequenceDataFormat struct{
				Id int64
			}
			var sequenceData SequenceDataFormat	
			tx := r.db.Begin()
			tx.Table("sequence_data").Last(&sequenceData)
			seq :=  "0000000000"
			seqString := strconv.FormatInt(sequenceData.Id, 10)
			appCustomerCodeSeq := seq+seqString
			appCustomerCodeSeq = appCustomerCodeSeq[len(appCustomerCodeSeq)-len(seq):]
			
			NewCustCode = NewCustCode +appCustomerCodeSeq
			dataLoan := &loan
			dataLoan.Custcode = NewCustCode
			dataCustomer := &customer
			dataCustomer.Custcode = NewCustCode
			dataVehicle := &vehicle
			dataVehicle.Custcode = NewCustCode
			

			res := tx.Model(&sc).Where("id", el.ID).Update("sc_flag","1")
			if res.Error != nil{
				tx.Rollback()
				continue
			}
			res = tx.Table("sequence_data").Where("id", sequenceData.Id).Update("id", (sequenceData.Id+1))
			if res.Error != nil{
				tx.Rollback()
				continue
			}
			res = tx.Create(&customer)
			if res.Error != nil{
				tx.Rollback()
				continue
			}
			res = tx.Create(&loan)
			if res.Error != nil{
				tx.Rollback()
				continue
			}
			res = tx.Create(&vehicle)
			if res.Error != nil{
				tx.Rollback()
				continue
			}
			fmt.Println("finish")
			tx.Commit()

		}
		
	}
	// fmt.Println(StagingCustomer)
}
func (r *repository) GenerateSkalaAngsuran()  {
	type Result struct{
		model.LoanDataTab
	}
	var result []Result
	r.db.Model(&model.LoanDataTab{}).Select("loan_data_tab.*").
	Joins("join customer_data_tab  as cdt on loan_data_tab.custcode=cdt.custcode").
	Where("approval_status=?","0").Scan(&result)
	fmt.Println(result)
	for _, el :=range result{
		LoanPeriod, err := strconv.ParseInt(el.LoanPeriod, 10, 8)
		if(err != nil){
			continue
		}
		total := el.LoanAmount
		
		effRate := el.InterestEffective
		dataSkalaRental  := make([]model.SkalaRentalTab, LoanPeriod+1)
		now := time.Now()
		

		for i := range dataSkalaRental{
			if i==0{
				data := dataSkalaRental[i]
				data.Custcode = el.Custcode
				data.Counter = int8(i)
				data.Osbalance = total
				data.EndBalance = total
				data.DueDate = now
				data.EffRate = effRate
				data.Rental = el.MonthlyPayment
				data.Principle = 0
				data.Interest = 0
				data.Inputdate = now
				dataSkalaRental[i] = data
				continue
			}
			var new time.Time
			new = now.AddDate(0, i, 0)
			if new.Day() != now.Day(){
				new = new.AddDate(0, 0, -new.Day())
			}
			
			interest :=  math.Floor(total *float64(effRate) *30/36000)
			principle := el.MonthlyPayment - interest
			end_balance := total - principle

			data := dataSkalaRental[i]
			data.Custcode = el.Custcode
			data.Counter = int8(i)
			data.Osbalance = total
			data.EndBalance =end_balance
			data.DueDate = new
			data.EffRate = effRate
			data.Rental = el.MonthlyPayment
			data.Principle = principle
			data.Interest = interest
			data.Inputdate = now
			if(end_balance < 0){
				data.EndBalance = 0
				data.Interest = data.Interest + end_balance
			}

			total = end_balance
			dataSkalaRental[i] = data
			
		}
		// fmt.Printf("%v", dataSkalaRental)
		// fmt.Println(dataSkalaRental)
		tx := r.db.Begin()
		res := tx.Create(&dataSkalaRental)
		if res.Error != nil{
			tx.Rollback()
			continue
		}
		res = tx.Model(&model.CustomerDataTab{}).Where("custcode",el.Custcode ).Update("approval_status", "1")
		if res.Error != nil{
			tx.Rollback()
			continue
		}
		tx.Commit()
		
	}
	
}