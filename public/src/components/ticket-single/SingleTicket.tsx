import { Card, CardContent, Grid, Typography } from "@mui/material"
import CustomerTicket from "../../model/customerTicket"

interface SingleTicketProps {
  ticket: CustomerTicket
}

const SingleTicket = ({ ticket } : SingleTicketProps ) => {
  return (
    <Card sx={{ minWidth: 275, backgroundColor:"var(--primary-color)", color: "var(--primary-text)", fontSize: "1.2rem" }}>
      <CardContent>
        <Typography sx={{ fontSize: 16, fontWeight:"bold", color:"#222", letterSpacing:"1.1px" }}  gutterBottom>
          Ticket
        </Typography>
        <Grid container justifyContent="center" alignItems="center">
          <Grid item xs={4}>
            From
          </Grid>
          <Grid item xs={8} >
            {ticket.flightPlaceSource}
          </Grid>
          <Grid item xs={4} >
            Departure
          </Grid>
          <Grid item xs={8} >
            {ticket.flightDateSource}
          </Grid>
          <Grid item xs={4} >
            To
          </Grid>
          <Grid item xs={8} >
            {ticket.flightPlaceDestination}
          </Grid>
          <Grid item xs={4} >
            Arrival
          </Grid>
          <Grid item xs={8} >
            {ticket.flightDateDestination}
          </Grid>
          <Grid item xs={4} >
            Price
          </Grid>
          <Grid item xs={8} >
            {ticket.flightTicketPrice}$
          </Grid>
          <Grid item xs={4} >
            Quantity
          </Grid>
          <Grid item xs={8} >
            {ticket.quantity}
          </Grid>
        </Grid>
        <Typography variant="body2">
          
        </Typography>
      </CardContent>
    </Card>
  )
}

export default SingleTicket;