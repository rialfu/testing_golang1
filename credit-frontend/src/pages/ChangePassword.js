import React, { memo, useState } from 'react'
import {useStore} from '../Context/context'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'
 function ChangePassword() {
    const {state, dispatch} = useStore()
    const navigate = useNavigate()
    const[form, setForm] = useState({"oldpass":"", "newpass":""})
    const changeValue = (e)=>{
        if(e.target.id === "inputNewPassword"){
            setForm({...form, newpass:e.target.value})
        }else if(e.target.id === "inputPassword"){
            setForm({...form, oldpass:e.target.value})
        }
    }
    const send = async()=>{
        const data = {username:state.user.username, old:form.oldpass, new:form.newpass}
        try{
            await axios.post("http://localhost:8080/update-password", JSON.stringify(data))
            localStorage.removeItem("account")
            dispatch({type:"logout"}); 
            alert("Berhasil update")
            setTimeout(()=>{
                navigate('/',{replace:true} )
            }, 500)
            
        }catch(err){
            if(err.response.status===401){
                alert("gagal update")
            }else{
                alert("server sedang bermasalah")
            }
            
        }
        
    }
    return (
        <>
            <div className="row mt-4">
                <div className="col-12">
                    <div className="card-style mb-5">
                        <h5>Update Password</h5>
                        
                        <div className="mb-3 row">
                            <label htmlFor="inputPassword" className="col-sm-2 col-form-label">Old Password</label>
                            <div className="col-sm-10">
                                <input type="password" className="form-control" id="inputPassword" onChange={changeValue}/>
                            </div>
                        </div>
                        <div className="mb-3 row">
                            <label htmlFor="inputNewPassword" className="col-sm-2 col-form-label">New Password</label>
                            <div className="col-sm-10">
                                <input type="password" className="form-control" id="inputNewPassword"  onChange={changeValue}/>
                            </div>
                        </div>
                        <button className='btn btn-success' onClick={send}>Update</button>
                    </div>
                </div>
            </div>
        </>
    )
}
export default memo(ChangePassword)
