import {
  TableContainer,
  Paper,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Button,
} from "@mui/material";
import { Flight } from "../../model/flight";
import { Stack } from "@mui/system";
import { useState } from "react";
import TicketQuantity from "../ticket-quantity/TicketQuantity";

interface Props {
  data: Flight[];
  onDelete: (flight: Flight) => void;
  buyTickets: (quantity: number | undefined) => void;
  setActiveFlightId: (id: string) => void;
  user?: any;
}

const FlightList = ({ data, onDelete, buyTickets, setActiveFlightId, user }: Props) => {
  const [showQuantityModal, setShowQuantityModal] = useState(false);

  if (!data || data.length === 0) {
    return (
      <>
        <h3>Flights</h3>
        <div>Empty</div>
      </>
    );
  }

  const showModal = (show: boolean) => {
    setShowQuantityModal(show)
  }


  return (
    <>
      <h3>Flights</h3>
      <TicketQuantity isOpen={showQuantityModal} showModal={showModal} buyTickets={buyTickets} />
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell>Place Destination</TableCell>
              <TableCell align="right">Place Source</TableCell>
              <TableCell align="right">Ticket Price</TableCell>
              <TableCell align="right">Total Price</TableCell>
              <TableCell align="right">Date Source</TableCell>
              <TableCell align="right">Date Destination</TableCell>
              <TableCell align="right">Total Tickets</TableCell>
              <TableCell align="right">Bought Tickets</TableCell>
              <TableCell align="center">Action</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {data.map((d: any, index: number) => (
              <TableRow
                key={index}
                sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
              >
                <TableCell component="th" scope="row">
                  {d.placeDestination}
                </TableCell>
                <TableCell align="right">{d.placeSource}</TableCell>
                <TableCell align="right">{d.ticketPrice}€</TableCell>
                <TableCell align="right">
                  {d.ticketPrice * d.boughtTickets}€
                </TableCell>
                <TableCell align="right">
                  {new Date(d.dateSource).toLocaleString()}
                </TableCell>
                <TableCell align="right">
                  {new Date(d.dateDestination).toLocaleString()}
                </TableCell>
                <TableCell align="right">{d.totalTickets}</TableCell>
                <TableCell align="right">{d.boughtTickets}</TableCell>
                <TableCell align="right">
                  <Stack direction="row">
                    {user?.includes(1) && (
                      <Button
                        variant="outlined"
                        color="error"
                        sx={{ mr: 3 }}
                        onClick={() => onDelete(d)}
                      >
                        Delete
                      </Button>
                    )}
                    <Button
                      variant="outlined"
                      color="primary"
                      sx={{ whiteSpace: 'nowrap' }}
                      onClick={() => { setActiveFlightId(d.id); showModal(true) }}
                    >
                      <span>Buy ticket</span>
                    </Button>
                  </Stack>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
};

export default FlightList;
