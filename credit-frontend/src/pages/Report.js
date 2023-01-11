import { memo, useEffect, useRef, useState } from "react"
import axios from 'axios'
import Pagination, { bootstrap5PaginationPreset } from 'react-responsive-pagination';
import { useLocation } from "react-router-dom";
function Report(props){
    const [dataReport, setDataReport] = useState([])
    const [companyData, setCompanyData] = useState([])
    const [branchData, setBranchData] = useState([])
    const [totalPage, setTotalPage] = useState(0)
    const [page, setPage] = useState(1)
    const limit =useRef(5)
    const startDate = useRef('')
    const endDate = useRef('')
    const company = useRef('')
    const branch = useRef('')
    const load = useRef(true)
    
    useEffect(()=>{
        getCompany()
    },[])
    useEffect(()=>{
        getData()
    },[page])
    const changePage = (i)=>{
        setPage(i)
    }
    const getCompany = async()=>{
        const res = await axios.get(`http://localhost:8080/get-search`)
        const {dataBranch, dataCompany } = res.data
        setCompanyData(dataCompany)
        setBranchData(dataBranch)
        console.log(dataBranch)    
        // console.log(res.data)
        // setCompanyData(res.data.data)
    }
    
    
    const getData = async ()=>{
        let url = `http://localhost:8080/report-pencairan?limit=${limit.current}&page=${page}`
        url += `&start=${startDate.current}&end=${endDate.current}&company=`
        url += `&branch=${branch.current}`
        try{
            const res = await axios.get(url)
            console.log(res.data)
            const {data, total_page} = res.data.data
            const up = data.map((el,i)=>{return {...el, 'check':false}})
            setDataReport(up)
            setTotalPage(total_page)
        }catch(err){
            console.log(err)
        }
        load.current = false
        
    }
    
    return(
        <div className="mt-3">
            <div className="p-3 bg-white " style={{minHeight:"300px", }}>
                <div className="row mb-3">
                    <div className="col-md-4">
                        <label htmlFor="">Start Date</label>
                        <input type="date" className="form-control" onChange={(e)=>startDate.current=e.target.value}  />
                    </div>
                    <div className="col-md-4">
                        <label htmlFor="">End Date</label>
                        <input type="date" className="form-control" onChange={(e)=>endDate.current=e.target.value}  />
                    </div>
                    <div className="col-md-4">
                        <label htmlFor="">Company</label>
                        <select id="company" name="Company" defaultValue={''} className="form-control" 
                            onChange={(e)=>company.current=e.target.value}>
                        <option  value="">Pilih Company</option>
                            {
                                companyData.map((el, id)=>{
                                    return(
                                        <option key={id} value={el.company_code}>{el.company_name}</option>
                                    )
                                })
                            }
                        </select>
                    </div>
                    <div className="col-md-4">
                        <label htmlFor="">Branch</label>
                        <select id="branch" name="branch" defaultValue={''} className="form-control" 
                            onChange={(e)=>branch.current=e.target.value}>
                        <option  value="">Pilih Branch</option>
                            {
                                branchData.map((el, id)=>{
                                    return(
                                        <option key={id} value={el.Code}>{el.Description}</option>
                                    )
                                })
                            }
                        </select>
                    </div>
                    
                    
                </div>
                <div className="row">
                    <div className="col-md-4">
                        <button className="btn btn-primary" onClick={getData}>Search</button>
                    </div>
                </div>
                {
                   dataReport.length === 0?
                   (
                    <>
                        <h3 className="text-center">Data Kosong</h3>
                    </>
                    
                   ) :
                   (
                    <>
                        <div className="table-responsive-xl table-wrapper mb-4">
                    
                    <table className=" table">
                        <thead>
                            <tr>
                                <th style={{fontSize:"1vw"}} >No</th>
                                <th style={{fontSize:"1vw"}} >PPK</th>
                                <th style={{fontSize:"1vw"}} >Name</th>
                                <th style={{fontSize:"1vw"}} >Channeling Company</th>
                                <th style={{fontSize:"1vw"}} >Drawdown Date</th>
                                <th style={{fontSize:"1vw"}} >Loan Amount</th>
                                <th style={{fontSize:"1vw"}} >Loan Period</th>
                                <th style={{fontSize:"1vw"}} >Interest Eff</th>
                            </tr>
                        </thead>
                        <tbody>
                            {
                                dataReport.map((data, i)=>{
                                    let no =(page - 1) *limit.current + (i +1)
                                    return(
                                        <tr key={i}>
                                            <td style={{fontSize:"1vw"}}>{no}</td>
                                            <td style={{fontSize:"1vw"}}>{data.ppk}</td>
                                            <td style={{fontSize:"1vw"}}>{data.name}</td>
                                            <td style={{fontSize:"1vw"}}>{data.channeling_company}</td>
                                            <td style={{fontSize:"1vw"}}>{data.drawdown_date.substring(0,10)}</td>
                                            <td style={{fontSize:"1vw"}}>{data.loan_amount}</td>
                                            <td style={{fontSize:"1vw"}}>{data.loan_period}</td>
                                            <td >{data.interest_effective}</td>
                                            
                                        </tr>
                                    )
                                })
                            }
                            
                        </tbody>
                        
                        
                    </table>
                    
                </div>
                <Pagination
                    {...bootstrap5PaginationPreset}
                    current={page}
                    total={totalPage}
                    onPageChange={(e)=>{
                        if(e !== page){
                            changePage(e)
                        }
                        
                    }}
                    previousLabel="<"
                    nextLabel=">"
                    maxWidth={300}
                />
                    </>
                   )
                }
                
            </div>
        </div>
        
    )
}
export default memo(Report)