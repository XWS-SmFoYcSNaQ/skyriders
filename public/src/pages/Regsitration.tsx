
import { Box } from '@mui/system';
import RegistrationForm from '../components/registration-form/RegistrationForm';
import { User } from '../model/user';

const Registration = () => {

  const onRegistrationSubmit = async (user: User) => {
  }

  return (
    <div>
      <Box sx={{padding: "30px"}}>
          <RegistrationForm onSubmit={onRegistrationSubmit}></RegistrationForm>
      </Box>
    </div>
  );
}

export default Registration