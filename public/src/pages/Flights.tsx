import * as React from 'react';
import { useEffect, useState } from 'react';
import axios from 'axios';
import FlightForm from '../components/flight-form/FlightForm';
import FlightList from '../components/flight-list/FlightList';
import { Box } from '@mui/system';
import { Flight } from '../model/flight';

const Flights = () => {
  const [data, setData] = useState([])

  const fetchData = async () => {
    try{
      const res = await axios.get("http://localhost:9000/api/flight")
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
      await axios.post("http://localhost:9000/api/flight", flight)
      fetchData()
    }catch (err) {
      console.log(err)
    }
  }

  const onFlightDelete = async (flight: Flight) => {
    try {
      await axios.delete(`http://localhost:9000/api/flight/${flight.id}`)
      fetchData()
    }catch (err) {
      console.log(err)
    }
  }

  return (
    <div>
      <Box sx={{padding: "30px"}}>
        <FlightList onDelete={onFlightDelete} data={data}/>
        <Box sx={{paddingTop: "30px"}}>
          <FlightForm onSubmit={onFlightSubmit}/>
        </Box>
      </Box>
    </div>
  );
}

export default Flights