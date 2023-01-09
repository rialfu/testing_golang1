import React, { memo, useState } from 'react'

 function ChangePassword() {
    const[form, setForm] = useState({"oldpass":"", "newpass":""})
    const changeValue = (e)=>{
        if(e.target.id === "inputNewPassword"){
            setForm({...form, newpass:e.target.value})
        }else if(e.target.id === "inputPassword"){
            setForm({...form, oldpass:e.target.value})
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
                        <button className='btn btn-success'>Update</button>
                    </div>
                </div>
            </div>
        </>
    )
}
export default memo(ChangePassword)
