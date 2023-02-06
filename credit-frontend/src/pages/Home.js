import { memo, useEffect, useRef, useState } from "react"
import axios from 'axios'
import Pagination, { bootstrap5PaginationPreset } from 'react-responsive-pagination';
function Home(){
    const[dataReport, setDataReport] = useState([])
    const [totalPage, setTotalPage] = useState(0)
    const [page, setPage] = useState(1)
    const limit =useRef(5)
    const load = useRef(false)
    const [isChecked, setIsChecked] = useState(false)
    useEffect(()=>{
    },[])
    useEffect(()=>{
        getData()
    },[page])
    const changePage = (i)=>{
        setPage(i)
    }
    const updateData = async()=>{
        if (!isChecked){
            return 
        }
        // console.log(dataReport)
        // return
        if(load.current === false){
            load.current = true
            try{
                let data = [...dataReport]
                data = data.filter((el)=>el.check).map((el)=>{return {"id":el.custcode}})

                const res= await axios.post(`http://localhost:8080/checklist-pencairan`,JSON.stringify({"data":data}))
                console.log(res)
                await getData()
                alert("Sukses")
            }catch(err){
                alert("gagal")
                console.log(err)
            }
            load.current = false
            
            
        }
    }
    const checked = (i)=>{
        let data = [...dataReport]
        let isChecked=false
        data = data.map((el,index)=>{
            return {...el, check:i===index?!el.check:el.check}
        })
        isChecked = data.some(el=>el.check)
        setIsChecked(isChecked)
        setDataReport(data)
    }
    const getData = async ()=>{
        console.log("run")
        try{
            const res = await axios.get(`http://localhost:8080/checklist-pencairan?limit=${limit.current}&page=${page}`)
            console.log(res.data)
            const {data, total_page} = res.data.data
            const up = data.map((el,i)=>{return {...el, 'check':false}})
            setDataReport(up)
            setIsChecked(false)
            setTotalPage(total_page)
        }catch(err){
            alert("Gagal mendapatkan data")
            console.log(err)
        }
        
    }
    
    
    return(
        <div className="mt-3">
            <div className="p-3 bg-white">
                
                {
                    dataReport.length === 0 ?
                    <>
                        <button className="btn btn-primary" onClick={()=>setPage(1)}>Refresh</button>
                        <h3 className="text-center">Data Kosong</h3>
                    </>:
                    <>
                        <div style={{float:"right"}}>
                            <button className={"btn btn-primary btn-disabled "} disabled={!isChecked} onClick={updateData}>Approve</button>
                        </div>
                        <div className="table-responsive-xl table-wrapper mb-4">
                            <table className=" table">
                                <thead>
                                    <tr>
                                        <th style={{fontSize:"1vw"}} width={70}>No</th>
                                        <th style={{fontSize:"1vw"}} width={170}>PPK</th>
                                        <th style={{fontSize:"1vw"}} width={190}>Name</th>
                                        <th style={{fontSize:"1vw"}} width={130}>Channeling Company</th>
                                        <th style={{fontSize:"1vw"}} width={190}>Drawdown Date</th>
                                        <th style={{fontSize:"1vw"}} width={190}>Loan Amount</th>
                                        <th style={{fontSize:"1vw"}} width={190}>Loan Period</th>
                                        <th style={{fontSize:"1vw"}} width={190}>Interest Eff</th>
                                        <th style={{fontSize:"1vw"}} width={190}>action</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {
                                        dataReport.map((data, i)=>{
                                            let no =(page - 1) *limit.current + (i +1)
                                            // console.log((page.current - 1) * limit)
                                            return(
                                                <tr key={i}>
                                                    <td style={{fontSize:"1vw"}} >{no}</td>
                                                    <td style={{fontSize:"1vw"}} >{data.ppk}</td>
                                                    <td style={{fontSize:"1vw"}} >{data.name}</td>
                                                    <td style={{fontSize:"1vw"}} >{data.channeling_company}</td>
                                                    <td style={{fontSize:"1vw"}} >{data.drawdown_date.substring(0,10)}</td>
                                                    <td style={{fontSize:"1vw"}} >{data.loan_amount}</td>
                                                    <td style={{fontSize:"1vw"}} >{data.loan_period}</td>
                                                    <td style={{fontSize:"1vw"}} >{data.interest_effective}</td>
                                                    <td style={{fontSize:"1vw"}} > 
                                                        <div className="form-check">
                                                            <input className="form-check-input" type="checkbox" value="" id="flexCheckDefault" onChange={()=>checked(i)} checked={data.check}/>
                                                                <label className="form-check-label"  htmlFor="flexCheckDefault">
                                                                    Pilih
                                                                </label>
                                                            </div>
                                                    </td>
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
                }
                
            </div>
        </div>
        
    )
}
export default memo(Home)