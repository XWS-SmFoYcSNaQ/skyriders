import { Box } from "@mui/system"
import LoginForm from '../components/login-form/LoginForm'
import axios from "axios"

const Login = () => {

  const login = async (data: any) => {
    try {
      const response = await axios.post('auth/login', data, { withCredentials: true })
      const respData = response.data
      axios.defaults.headers.common['Authorization'] = `Bearer ${respData['access_token']}`
      window.location.replace('/');
    } catch (err) {
      console.log(err);
    }
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