import * as React from 'react';
import Box from '@mui/material/Box';
import Drawer from '@mui/material/Drawer';
import CssBaseline from '@mui/material/CssBaseline';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import List from '@mui/material/List';
import Typography from '@mui/material/Typography';
import Divider from '@mui/material/Divider';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import HowToRegIcon from '@mui/icons-material/HowToReg';
import LoginIcon from '@mui/icons-material/Login';
import HomeIcon from '@mui/icons-material/Home';
import { Outlet, NavLink, useNavigate } from 'react-router-dom';
import { Button } from '@mui/material';
import axios from 'axios';
import { useEffect, useState } from 'react';
import AirplaneTicketIcon from '@mui/icons-material/AirplaneTicket';
import FlightTakeoffIcon from '@mui/icons-material/FlightTakeoff';
import useAuth from "../../hooks/useAuth";

const drawerWidth = 240;

interface NavItem {
  route: string;
  text: string;
  icon: JSX.Element;
}

const MainLayout = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const { auth, setAuth } = useAuth();

  let navigate = useNavigate();

  const checkAuthStatus = async () => {
    try {
      const response = await axios.get('auth/check');
      if (response.status === 200) {
        setIsAuthenticated(true);
        const user = 'logged'
        const roles = [response.data['roles']]
        setAuth({ user, roles })
      } else {
        setIsAuthenticated(false);
      }
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    checkAuthStatus();
  }, []);

  const logout = async () => {
    const response = await axios.get('auth/logout', { withCredentials: true })
    if (response.status === 200) {
      setIsAuthenticated(false)
      axios.defaults.headers.common['Authorization'] = ""
      navigate('/login');
    }
  };

  let upperNavItems: NavItem[] = [
    {
      route: '/',
      text: 'Home',
      icon: <HomeIcon />
    }
  ];

  const adminNavItems: NavItem[] = [
    {
      route: '/flights',
      text: 'Flights',
      icon: <FlightTakeoffIcon />
    }
  ];

  const customerNavItems: NavItem[] = [
    {
      route: '/flights',
      text: 'Flights',
      icon: <FlightTakeoffIcon />
    },
    {
      route: "/myTickets",
      text: "My Tickets",
      icon: <AirplaneTicketIcon />
    }
  ];

  try {
    if (auth?.roles.includes(0)) {
      upperNavItems.push(...customerNavItems);
    } else if (auth?.roles.includes(1)) {
      upperNavItems.push(...adminNavItems);
    }
  } catch { }

  const lowerNavItems: NavItem[] = [
    {
      route: '/register',
      text: 'Register',
      icon: <HowToRegIcon />
    },
    {
      route: '/login',
      text: 'Login',
      icon: <LoginIcon />
    }
  ];

  const filteredLowerNavItems = isAuthenticated
    ? lowerNavItems.filter(item => item.route !== '/login' && item.route !== '/register')
    : lowerNavItems;

  return (
    <Box sx={{ display: 'flex' }}>
      <CssBaseline />
      <AppBar
        position="fixed"
        sx={{ width: `calc(100% - ${drawerWidth}px)`, ml: `${drawerWidth}px` }}
      >
        <Toolbar>
          <Typography variant="h5" noWrap component="div">
            Skyriders
          </Typography>
        </Toolbar>
      </AppBar>
      <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: drawerWidth,
            boxSizing: 'border-box',
          },
        }}
        variant="permanent"
        anchor="left"
      >
        <Toolbar />
        <Divider />
        {/* Upper nav items */}
        <List>
          {upperNavItems.map((navItem, index) => (
            <NavLink to={navItem.route} key={navItem.route}>
              <ListItem disablePadding>
                <ListItemButton>
                  <ListItemIcon>
                    {navItem.icon}
                  </ListItemIcon>
                  {navItem.text}
                </ListItemButton>
              </ListItem>
            </NavLink>
          ))}
        </List>
        <Divider />
        {/* Lower nav items */}
        <List>
          {filteredLowerNavItems.map((navItem, index) => (
            <NavLink to={navItem.route} key={navItem.route} >
              <ListItem disablePadding>
                <ListItemButton>
                  <ListItemIcon>
                    {navItem.icon}
                  </ListItemIcon>
                  {navItem.text}
                </ListItemButton>
              </ListItem>
            </NavLink>
          ))}
        </List>
        {isAuthenticated && <Button onClick={logout}>Logout</Button>}
      </Drawer>
      <Box
        component="main"
        sx={{ flexGrow: 1, bgcolor: 'background.default', py: 8, overflowX: 'visible' }}
      >
        <Outlet />
      </Box>
    </Box>
  );
}

export default MainLayout;