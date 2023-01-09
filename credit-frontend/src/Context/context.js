import { createContext, useReducer, useContext } from 'react'
const Store = createContext()
const useStore = () => useContext(Store)
// Store.displayName = "store"
export { Store, useStore }

function AccountReducer(state, action) {
  switch (action.type) {
    case 'set': {
      return { ...state, user: action.payload }
    }
    case 'logout':{
        return {...state, user:{username:"", name:""}}
    }
    default: {
      console.log("not found" + action.type)
      // throw new Error(`Unhandled action type: ${action.type}`)
    }
  }
}
export default function UserContext({ children }) {
  const [state, dispatch] = useReducer(AccountReducer, { user: {username:"", name:""} }, () => {
    const account = localStorage.getItem("account")
    let user = {username:"", name:""}
    try{
        if(account != null){
            user =  JSON.parse(account)
        }
        
    }catch(e){
        console.log("gagal")
    }
    return { user:user }
  })
  const value = { state, dispatch }
  return <Store.Provider value={value}>{children}</Store.Provider>
}