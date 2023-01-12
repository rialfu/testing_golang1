import { useEffect, useRef, useState } from "react"
import { useStore } from "../Context/context"
import  '../assets/custom.css'
import '../assets/login.css'
import { useNavigate } from "react-router-dom"
import axios from 'axios'
export default function Login(){
    const {state, dispatch} = useStore()
    const navigate = useNavigate()
    
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    
    const loadRef = useRef(false)
    const [load, setLoad] = useState(false)
    useEffect(()=>{
        if(state.user.username !== ""){
            navigate('/')
        }
        setLoad(false)
        // userRef.current.focus()
    }, [load])
    const login = async ()=>{
        
        if(loadRef.current === false){
            loadRef.current = true
            
            if(username === "" || password ===""){
                
                loadRef.current = false
                return
            }
            setLoad(true)
            try{
                const res =  await axios.post("http://localhost:8080/login",JSON.stringify({username, password}))
                const {token, data} = res.data
                dispatch({type:"set", payload:data})
                localStorage.setItem("account", JSON.stringify(data))
                setTimeout(()=>{
                    navigate('/',{replace:true} )
                }, 500)
                loadRef.current = false
                return
            }catch(e){
                if(e.response.status === 401){
                    alert("Username atau password salah")
                }else{
                    alert("Server sedang bermasalah")
                }
                console.log()
            }
            setLoad(false)
            loadRef.current = false
            return
            
        }
        
    }
    const eventFocus = (id)=>{
        document.getElementById(id).style.top="-34%"
    }
    const eventBlur = (id)=>{
        if((id==='p1' && username==="") || (id==='p2' && password === ''))
            document.getElementById(id).style.top="10%"
    }
    const eventKeyDown = (e)=>{
        if(e.key==='Enter'){
            login()
        }
    }  
    // let loadHTML=load?(<div style={{width:"100%", height:"100vh", position:"fixed",background:"rgba(204, 66, 24, 0.5)"}}>
    //         {/* <h1>Res</h1> */}
    //         <div className="d-flex justify-content-center align-items-center h-100">
    //         <div className="lds-circle"><div></div></div>
    //         </div>
            
    //     </div>):<></>
    
    
    
    return(
        <div className="row">
            <div className="col-6 d-none d-md-block left-side" style={{backgroundColor:"#fa6b4b"}}>
                <div className="d-flex align-items-center" style={{height:"100vh"}}>
                <img src={window.location.origin+"/pinjaman.png"} alt="" />
                </div>
                
            </div>
            <div className="col-6 right-side">
                <div className="login-content">
                    <img className="mb-4" src={window.location.origin+"/logo.png"} style={{maxHeight:"50px"}} alt="" />
                    <h2 className="mb-1">Hello Madafak</h2>
                    <p className="mb-4">Welcome to My Life</p>
                    <div className="input-group1">
                        <label htmlFor="userid" className="placeholder" id="p1">Username</label>
                        <input type="text" name="username" className="in" 
                            onFocus={()=>eventFocus('p1')} 
                            onBlur={()=>eventBlur('p1')} 
                            onChange={(e)=>{setUsername(e.target.value)}}
                            onKeyDown={eventKeyDown}
                        />
                        
                    </div>
                    <div className="input-group1">
                        <label htmlFor="password" className="placeholder" id="p2">Password</label>
                        <input type="password" name="password" className="in" 
                            onFocus={()=>eventFocus('p2')} onBlur={()=>eventBlur('p2')} 
                            onChange={(e)=>{setPassword(e.target.value)}}
                            onKeyDown={eventKeyDown}
                        />
                    </div>
                    <button className="btn btn-primary w-100 mb-4" disabled={load} onClick={login}>Login </button>
                </div>
            </div>
        </div>
    //     <div style={{width:"100%", height:"100vh", background:"#cc3333", display:"flex", alignItems:"center",justifyContent:"center"}}>
    //         {loadHTML}
    //         <div style={{width:"500px", background:"white", padding:"20px 30px",}} className="rounded">
    //             <div className="d-flex justify-content-center h-100">
    //             <img src={window.location.origin+"/logo.png"} alt="" style={{maxHeight:"10vh"}} />
    //             </div>
    //             <h1 className="text-center" style={{textShadow:`-1px -1px 0 red,  
    // 1px -1px 0 red,
    // -1px 1px 0 red,
    //  1px 1px 0 red`, color:"white" }}>Login</h1>
                
    //             <div className="mb-3">
    //                 <label htmlFor="exampleFormControlTextarea1" className="form-label">Username</label>
    //                 <input className="form-control" id="exampleFormControlTextarea1" ref={userRef} value={username} onChange={(e)=>{setUsername(e.target.value)}}/>
    //             </div>
    //             <div className="mb-3">
    //                 <label htmlFor="exampleFormControlTextarea1" className="form-label">Password</label>
    //                 <input type="password" className="form-control" id="exampleFormControlTextarea1" value={password} onChange={(e)=>{setPassword(e.target.value)}}/>
    //             </div>
    //             <button className="btn btn-primary w-100" onClick={login}>Login</button>
                
    //         </div>
    //     </div>
       
    )
}