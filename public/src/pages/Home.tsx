import { Button, CardMedia } from '@mui/material';
import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import AirplaneImage from '../images/airplane.jpg';
import styles from './Home.module.css';
import AirplaneTicketIcon from '@mui/icons-material/AirplaneTicket';
import { NavLink } from 'react-router-dom';
import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import CardActions from '@mui/material/CardActions';
import PilotImage from '../images/pilot.png';
import ConnectingAirportsIcon from '@mui/icons-material/ConnectingAirports';

const Home = () => {
  return (
    <Box sx={{ 
      height: '500px', 
      backgroundImage: `url(${AirplaneImage})`, 
      backgroundRepeat: 'no-repeat', 
      backgroundPosition: '0px -90px', 
      backgroundSize: 'cover',
      position: 'relative'
      }}>
        <Grid container sx={{
          justifyContent: 'center',
          height: '100%',
          alignItems: 'center'
        }} >
          <Grid item xs={4} textAlign='center'>
            <Paper className={styles['main-content']}>
              <h1>Fly Safe!</h1>
              <p>Buy your ticket now.</p>
              <NavLink to='/flights'>
                <Button variant='contained' startIcon={<AirplaneTicketIcon/>} sx={{ 
                  backgroundColor: 'var(--primary-light)',
                  my: 1,
                  p: '11px 21px'
                }}>Buy ticket</Button>
              </NavLink>
            </Paper>
          </Grid>
        </Grid>
        <Grid container columnSpacing={25} sx={{
          py: 9,
          px: 7,
          backgroundColor: '#bcbec9'
        }}>
          <Grid item xs={12} md={6} lg={4} sx={{ justifyContent: 'center'}}>
            <Card variant='outlined' sx={{ 
              backgroundColor: 'var(--primary-color)',
              color: 'var(--primary-text)',
              maxHeight: '280px'}}>
              <CardHeader title='Flights'/>
              <CardContent>
                <Typography variant='body1'>
                  Take a look at our wide flight list and pick one that suits you best
                </Typography>
              </CardContent>
              <CardActions>
                <NavLink to='/flights'>
                  <Button variant='contained' startIcon={<ConnectingAirportsIcon/>} sx={{
                     my: 2, 
                     backgroundColor: 'var(--primary-dark)', 
                     color: 'var(--primary-text)'}}>Flights</Button>
                </NavLink>
              </CardActions>
            </Card>
          </Grid>
          <Grid item xs={12} md={6} lg={4} >
            <Box sx={{ position: 'relative', height: '100%',  maxHeight: '280px' }}>
              <img src={PilotImage} alt="Pilot" className={styles.pilot}/>
            </Box>
          </Grid>
          <Grid item xs={12} md={6} lg={4}>
            <Card variant='outlined' sx={{ 
              backgroundColor: 'var(--primary-color)',
              color: 'var(--primary-text)',
              maxHeight: '280px'}}>
              <CardHeader title='Tickets'/>
              <CardContent>
                <Typography variant='body1'>
                  You want to recall memories?<br/>
                  View all your tickets 
                </Typography>
              </CardContent>
              <CardActions>
                <NavLink to='/myTickets'>
                  <Button variant='contained' startIcon={<AirplaneTicketIcon/>} sx={{ 
                    my: 2, 
                    backgroundColor: 'var(--primary-dark)', 
                    color: 'var(--primary-text)'}}>My Tickets</Button>
                </NavLink>
              </CardActions>
            </Card>
          </Grid>
        </Grid>
    </Box>
  );
};

export default Home;