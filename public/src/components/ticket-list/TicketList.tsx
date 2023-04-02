import { Box, Grid, Typography } from "@mui/material";
import CustomerTicket from "../../model/customerTicket";
import SingleTicket from "../ticket-single/SingleTicket";

interface TicketListProps {
  tickets: CustomerTicket[] | undefined
}

const TicketList = ({ tickets } : TicketListProps) => {
  
  return (
    <Box>
      <Typography variant="h3" sx={{ pt: 2, pb: 4, textAlign:'center'}}>My Tickets</Typography>
      { tickets === undefined 
        ? <p>No tickets</p> 
        : (
        <Grid container justifyContent="start" alignItems="center" spacing={5}>
          { tickets.map(t => ( 
            <Grid item key={t.flightId} xs={12} sm={9} md={7} lg={6}>
              <SingleTicket ticket={t}/>
            </Grid>
            )) }
        </Grid>
      )}
    </Box>
  )
}

export default TicketList;