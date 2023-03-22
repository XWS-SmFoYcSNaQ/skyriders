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
import MailIcon from '@mui/icons-material/Mail';
import HowToRegIcon from '@mui/icons-material/HowToReg';
import LoginIcon from '@mui/icons-material/Login';
import HomeIcon from '@mui/icons-material/Home';
import { Outlet, NavLink  } from 'react-router-dom';

const drawerWidth = 240;

interface NavItem {
  route: string;
  text: string;
}

const MainLayout = () => {
  const upperNavItems: NavItem[] = [
    {
      route: '/',
      text: 'Home'
    },
    {
      route: '/flights',
      text: 'Flights'
    },
  ];

  const lowerNavItems: NavItem[] = [
    {
      route: '/login',
      text: 'Login'
    },
    {
      route: '/register',
      text: 'Register'
    }
  ];

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
            <NavLink to={navItem.route}>
              <ListItem key={index} disablePadding>
                <ListItemButton>
                  <ListItemIcon>
                    {index % 2 === 0 ? <HomeIcon /> : <MailIcon />}
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
          {lowerNavItems.map((navItem, index) => (
            <NavLink to={navItem.route} >
              <ListItem key={index} disablePadding>
                <ListItemButton>
                  <ListItemIcon>
                    {index % 2 === 0 ? <LoginIcon /> : <HowToRegIcon />}
                  </ListItemIcon>
                  {navItem.text}
                </ListItemButton>
              </ListItem>
            </NavLink>
          ))}
        </List>
      </Drawer>
      <Box
        component="main"
        sx={{ flexGrow: 1, bgcolor: 'background.default', py: 8, overflowX: 'visible'}}
      >
        <Outlet/>
      </Box>
    </Box>
  );
}

export default MainLayout;