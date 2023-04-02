import { Box } from "@mui/system"
import LoginForm from '../components/login-form/LoginForm'
import axios from "axios"
import { useState } from "react";
import { Navigate } from "react-router-dom";

const Login = () => {
  const [navigate, setNavigate] = useState(false);

  const login = async (data: any) => {
    try {
      const response = await axios.post('auth/login', data, { withCredentials: true })
      const respData = response.data
      axios.defaults.headers.common['Authorization'] = `Bearer ${respData['access_token']}`
      setNavigate(true)
    } catch (err) {
      console.log(err);
    }
  }

  if (navigate) {
    return <Navigate to="/" />;
  }

  return (
    <div>
      <Box sx={{ padding: "30px" }}>
        <LoginForm onSubmit={login} />
      </Box>
    </div>
  );
}

export default Login