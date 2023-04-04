import { useLocation, Navigate, Outlet } from 'react-router-dom';
import useAuth from '../../hooks/useAuth';
import { useEffect, useState } from 'react';
import Loader from '../loader/Loader';
import axios from 'axios';

const RequireAuth = ({ allowedRoles }: any) => {
  const { auth, setAuth } = useAuth();
  const location = useLocation();
  const [isLoading, setIsLoading] = useState(true)

  const checkAuthStatus = async () => {
    try {
      const response = await axios.get('auth/check');
      if (response.status === 200) {
        const user = true
        const roles = [response.data['roles']]
        const isAuthenticated = true;
        setAuth({ user, roles, isAuthenticated });
        setIsLoading(false);
      } else {
        const user = null
        const roles = null
        const isAuthenticated = false
        setAuth({ user, roles, isAuthenticated })
      }
    } catch (error) {
      console.log(error);
      setIsLoading(false);
    }
  };

  useEffect(() => {
    checkAuthStatus();
  }, [])

  console.log(auth, location)

  return (
    <>
      {!isLoading && (
        auth?.roles?.some((role: any) => allowedRoles?.includes(role))
          ? <Outlet />
          : auth?.user
            ? <Navigate to="/unathorized" state={{ from: location }} replace />
            : <Navigate to="/login" state={{ from: location }} replace />
      )}
      {isLoading && (
        <Loader />
      )}
    </>
  );
}

export default RequireAuth;