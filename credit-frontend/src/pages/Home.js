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
            if(i===index){
                if(!el.check){
                    isChecked =  true
                }
                return {...el,check:!el.check}
            }
            return el
        })
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
            console.log(err)
        }
        
    }
    const paginate = ()=>{
        let next =[]
        if(page !== 1){
            next.push(
                <li key={1} className="page-item"><a className="page-link" href="#" onClick={()=>changePage(1)}>1</a></li>
            )
        }
        let min = (page -2 < 1)? 1 : page -2
        if(min-2 <= 1){
            next.push(
                <li key={1000} className="page-item  disabled"><a className="page-link" href="#">...</a></li>
            )
        }
        for(let i=min;i<page; i++){
            next.push(
                <li  key={i} className="page-item"><a className="page-link" href="#" onClick={()=>changePage(i)}>{i}</a></li>
            )
        }
        next.push(<li className="page-item active"><a className="page-link" href="#">{page}</a></li>)
        let max = (page+2 > totalPage)?totalPage : (page+2)
        for(let i=page+1;i < max+1;i++){
            next.push(
                <li key={i} className="page-item"><a className="page-link" href="#" onClick={()=>changePage(i)}>{i}</a></li>
            )
        }
        if(max+1 !== 10){
            next.push(
                <li key={1001} className="page-item  disabled"><a className="page-link" href="#">...</a></li>
            )
        }
        if(page !== totalPage){
            next.push(
                <li key={totalPage} className="page-item"><a className="page-link" href="#" onClick={()=>changePage(totalPage)}>{totalPage}</a></li>
            )
        }
        return next
    }
    
    return(
        <div className="mt-3">
            <div className="p-3 bg-white">
                
                {
                    dataReport.length === 0 ?
                    <>
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
                                        <th width={70}>No</th>
                                        <th width={170}>PPK</th>
                                        <th width={190}>Name</th>
                                        <th width={130}>Channeling Company</th>
                                        <th width={190}>Drawdown Date</th>
                                        <th width={190}>Loan Amount</th>
                                        <th width={190}>Loan Period</th>
                                        <th width={190}>Interest Eff</th>
                                        <th width={190}>action</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {
                                        dataReport.map((data, i)=>{
                                            let no =(page - 1) *limit.current + (i +1)
                                            // console.log((page.current - 1) * limit)
                                            return(
                                                <tr key={i}>
                                                    <td>{no}</td>
                                                    <td>{data.ppk}</td>
                                                    <td>{data.name}</td>
                                                    <td>{data.channeling_company}</td>
                                                    <td>{data.drawdown_date.substring(0,10)}</td>
                                                    <td>{data.loan_amount}</td>
                                                    <td>{data.loan_period}</td>
                                                    <td>{data.interest_effective}</td>
                                                    <td> 
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