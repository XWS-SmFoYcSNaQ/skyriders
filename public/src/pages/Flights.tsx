import * as React from 'react';
import { useEffect, useState } from 'react';
import axios from 'axios';
import FlightForm from '../components/flight-form/FlightForm';
import FlightList from '../components/flight-list/FlightList';
import { Box } from '@mui/system';
import { Flight } from '../model/flight';
import { Container } from '@mui/material';

const Flights = () => {
  const [data, setData] = useState([])

  const fetchData = async () => {
    try{
      const res = await axios.get('flight')
      setData(res.data)
    }catch(err) {
      console.log(err)
    }
  }

  useEffect(() => {
    fetchData()
  }, [])

  const onFlightSubmit = async (flight: Flight) => {
    try {
      await axios.post('flight', flight);
      fetchData();
    } catch (err) {
      console.log(err);
    }
  }

  const onFlightDelete = async (flight: Flight) => {
    try {
      await axios.delete(`flight/${flight.id}`);
      fetchData();
    } catch (err) {
      console.log(err);
    }
  }

  return (
    <div>
      <Container maxWidth="lg">
        <Box sx={{padding: "30px"}}>
          <FlightList onDelete={onFlightDelete} data={data}/>
            <Box sx={{paddingTop: "30px"}}>
              <FlightForm onSubmit={onFlightSubmit}/>
            </Box>
        </Box>
      </Container>
    </div>
  );
}

export default Flights