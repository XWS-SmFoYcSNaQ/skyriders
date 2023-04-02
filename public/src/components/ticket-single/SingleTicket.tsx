import { Card, CardContent, Grid, Typography } from "@mui/material"
import CustomerTicket from "../../model/customerTicket"
import QRImage from "../../images/qr.png"
import { Stack } from "@mui/system"

interface SingleTicketProps {
  ticket: CustomerTicket
}

const SingleTicket = ({ ticket } : SingleTicketProps ) => {
  return (
    <Card sx={{ minWidth: 275, backgroundColor:"var(--primary-color)", color: "var(--primary-text)", fontSize: "1.2rem" }}>
      <CardContent>
        <Typography sx={{ fontSize: 20, fontWeight:"bold", color:"#222", letterSpacing:"1.1px", textAlign:"center" }}  gutterBottom>
          Ticket
        </Typography>
        <Grid container>
          <Grid container item xs={8}>
            <Grid item xs={4}>
              From
            </Grid>
            <Grid item xs={8}>
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
              {ticket.flightTicketPrice}€
            </Grid>
            <Grid item xs={4} >
              Quantity
            </Grid>
            <Grid item xs={8} >
              {ticket.quantity}
            </Grid>
          </Grid>
          <Grid item xs={4}>
            <Stack sx={{ alignItems: "center"}}>
              <img src={QRImage} width="120px" height="120px" alt="qr"/>
              <p>{ticket.fullName}</p>
            </Stack>
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  )
}

export default SingleTicket;