import {createBrowserRouter} from 'react-router-dom'
import Login from './pages/Login'
import HomePage from './pages/Home'
import withSidebarComponent from './Components/Sidebar'
import ChangePasswordPage from './pages/ChangePassword'
import ReportPage from './pages/Report'
const Home = withSidebarComponent(HomePage)
const ChangePassword = withSidebarComponent(ChangePasswordPage)
const Report = withSidebarComponent(ReportPage)
const routes = createBrowserRouter([
    {
        path:"/login",
        element: <Login/>
    },
    {
        path:"/",
        element: <Home/>
    },
    {
        path:"/change-password",
        element: <ChangePassword/>
    },
    {
        path:"/report",
        element: <Report/>
    }
])
export {routes}