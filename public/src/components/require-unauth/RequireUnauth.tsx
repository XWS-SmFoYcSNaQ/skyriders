import { Navigate, useLocation } from 'react-router-dom';
import useAuth from '../../hooks/useAuth';

const RequireUnAuth = ({ children }: any) => {
  const { auth } = useAuth();
  const location = useLocation();

  return auth?.isAuthenticated 
  ? <Navigate to="/" state={{ from: location }} replace />
  : <>{children}</>;
};

export default RequireUnAuth;