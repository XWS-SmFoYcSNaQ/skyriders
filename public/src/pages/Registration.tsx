
import { Box } from '@mui/system';
import axios, { AxiosError } from 'axios';
import RegistrationForm from '../components/registration-form/RegistrationForm';
import { User } from '../model/user';
import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const Registration = () => {

  interface ErrorResponse {
    error: string;
  }
  const onRegistrationSubmit = async (user: User) => {
    try {
      await axios.post('auth/register', user)
      toast.success('Customer successfully registered', {position: toast.POSITION.BOTTOM_CENTER});
    } catch (err: unknown) {
      if (axios.isAxiosError(err)) {
        const axiosError = err as AxiosError<ErrorResponse>;
        if (axiosError.response?.status === 400) {        
          toast.error(axiosError.response.data.error, {position: toast.POSITION.BOTTOM_CENTER});
        }
        else {
          console.log(err)
        }
      }
    }
  }

  return (
    <div>
      <ToastContainer/>
      <Box sx={{padding: "30px"}}>
          <RegistrationForm onSubmit={onRegistrationSubmit}></RegistrationForm>
      </Box>
    </div>
  );
}

export default Registration