import { useEffect, useRef, useState } from "react"
import { useStore } from "../Context/context"
import  '../assets/custom.css'
import { useNavigate } from "react-router-dom"
export default function Login(){
    const {state, dispatch} = useStore()
    const navigate = useNavigate()
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const userRef = useRef()
    const loadRef = useRef(false)
    const [load, setLoad] = useState(false)
    useEffect(()=>{
        if(state.user.username !== ""){
            navigate('/')
        }
        setLoad(false)
        // userRef.current.focus()
    }, [])
    const login = ()=>{
        if(loadRef.current === false){
            loadRef.current = true
            if(username === "" || password ===""){
                
                loadRef.current = false
                return
            }
            localStorage.setItem("account", JSON.stringify({username:"adimahesa", name:"Adi Mahesa"}))
            dispatch({type:"set", payload:{username:"adimahesa", name:"Adi Mahesa"}})
            navigate('/')
            loadRef.current = false
            return
            
        }
        
    }
    console.log(state)
    let loadHTML=load?(<div style={{width:"100%", height:"100vh", position:"fixed",background:"rgba(204, 66, 24, 0.5)"}}>
            {/* <h1>Res</h1> */}
            <div className="d-flex justify-content-center align-items-center h-100">
            <div className="lds-circle"><div></div></div>
            </div>
            
        </div>):<></>
    
    
    
    return(
        <div style={{width:"100%", height:"100vh", background:"#cc3333", display:"flex", alignItems:"center",justifyContent:"center"}}>
            {loadHTML}
            <div style={{width:"500px", background:"white", padding:"20px 30px",}} className="rounded">
                <h1 className="text-center">Login</h1>
                <div className="mb-3">
                    <label htmlFor="exampleFormControlTextarea1" className="form-label">Username</label>
                    <input className="form-control" id="exampleFormControlTextarea1" ref={userRef} value={username} onChange={(e)=>{setUsername(e.target.value)}}/>
                </div>
                <div className="mb-3">
                    <label htmlFor="exampleFormControlTextarea1" className="form-label">Password</label>
                    <input className="form-control" id="exampleFormControlTextarea1" value={password} onChange={(e)=>{setPassword(e.target.value)}}/>
                </div>
                <button className="btn btn-primary w-100" onClick={login}>Login</button>
                
            </div>
        </div>
       
    )
}