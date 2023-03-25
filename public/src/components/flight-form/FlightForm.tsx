import { Box, Button, TextField } from "@mui/material"
import { useState } from "react"
import { Flight } from "../../model/flight";
import classes from './FlightForm.module.css';

interface Props {
  onSubmit: (flight: Flight) => void
}

const FlightForm = ({onSubmit}: Props) => {
  const [data, setData] = useState<Flight>({})

  const handleSubmit = (e: any) => {
    e.preventDefault();
    onSubmit(data)
  }

  return (
    <form onSubmit={handleSubmit}>
      <h2>Add new flight</h2>
      <div className={classes.grid}>
        <TextField
          required
          label="Place Destination"
          onChange={(e) => setData({...data, placeDestination: e.target.value})}
          />
        <TextField
          required
          label="Place Source"
          onChange={(e) => setData({...data, placeSource: e.target.value})}
        />
        <TextField
          required
          label="Ticket Price"
          type="number"
          onChange={(e) => setData({...data, ticketPrice: parseInt(e.target.value)})}
        />
        <TextField
          label="Date Source"
          onChange={(e) => setData({...data, dateSource: e.target.value})}
        />
        <TextField
          label="Date Destination"
          onChange={(e) => setData({...data, dateDestination: e.target.value})}
        />
        <TextField
          label="Total Tickets"
          type="number"
          onChange={(e) => setData({...data, totalTickets: parseInt(e.target.value)})}
        />
        <TextField
          label="Bought Tickets"
          type="number"
          onChange={(e) => setData({...data, boughtTickets: parseInt(e.target.value)})}
        />
      </div>
      <Box marginTop="20px">
        <Button size="large" variant="contained" type="submit" >Create Flight</Button>
      </Box>
    </form>
  )
}

export default FlightForm