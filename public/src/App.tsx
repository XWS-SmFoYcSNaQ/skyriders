import React from 'react';
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import './App.css';
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

import Error from './pages/Error';
import Home from './pages/Home';
import Mainlayout from './components/layout/Main';
import Flights from './pages/Flights';
import Login from './pages/Login';
import Registration from './pages/Registration';
import MyTickets from './pages/MyTickets';
import Unauthorized from './pages/Unauthorized';
import RequireAuth from './components/require-auth/RequireAuth';
import { AuthProvider } from './context/AuthProvider';
import RequireUnAuth from './components/require-unauth/RequireUnauth';


const router = createBrowserRouter([
  {
    path: "/",
    element: <Mainlayout />,
    errorElement: <Error />,
    children: [
      {
        element: <RequireUnAuth />,
        children: [
          {
            index: true,
            path: "register",
            element: <Registration />
          },
          {
            index: true,
            path: "login",
            element: <Login />
          }
        ]
      },
      {
        index: true,
        element: <Home />,
      },
      {
        index: true,
        path: "unathorized",
        element: <Unauthorized />
      },
      {
        element: <RequireAuth allowedRoles={[0]} />,
        children: [
          {
            index: true,
            path: "myTickets",
            element: <MyTickets />
          }
        ]
      },
      {
        element: <RequireAuth allowedRoles={[0, 1]} />,
        children: [
          {
            index: true,
            path: "flights",
            element: <Flights />
          }
        ]
      }
    ]
  },
]);

function App() {
  return (
    <div className="App">
      <AuthProvider>
        <RouterProvider router={router} />
      </AuthProvider>
    </div>
  );
}

export default App;
