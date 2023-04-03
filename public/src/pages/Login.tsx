import { Box } from "@mui/system"
import LoginForm from '../components/login-form/LoginForm'
import axios, { HttpStatusCode } from "axios"
import { toast, ToastContainer } from "react-toastify"
import 'react-toastify/dist/ReactToastify.css';
import useAuth from "../hooks/useAuth";
import { useNavigate, useLocation } from 'react-router-dom';

const Login = () => {
  const {setAuth} = useAuth();
  const navigate = useNavigate();
  const location = useLocation();
  const from = location.state?.from?.pathname || "/";

  const login = async (data: any) => {
    try {
      const response = await axios.post('auth/login', data, { withCredentials: true })
      if (response?.status === HttpStatusCode.Ok) {
        const respData = response?.data
        axios.defaults.headers.common['Authorization'] = `Bearer ${respData['access_token']}`
        const roles = [respData['roles']]
        const user = data.email
        setAuth({ user, roles })
        navigate(from, { replace: true });
        toast.success('Successfully logged in', { position: toast.POSITION.BOTTOM_CENTER });
      } else {
        if (response?.request.response.includes('invalid email or password')) {
          toast.error('Invalid email or password, please try again', { position: toast.POSITION.BOTTOM_CENTER });
        }
      }

    } catch (err) {
      toast.error('Sorry, we\'re experiencing some technical difficulties. Please try again later.', { position: toast.POSITION.BOTTOM_CENTER });
    }
  }

  return (
    <div>
      <ToastContainer />
      <Box sx={{ padding: "30px" }}>
        <LoginForm onSubmit={login} />
      </Box>
    </div>
  );
}

export default Login