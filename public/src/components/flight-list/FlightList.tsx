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

interface Props {
  data: Flight[];
  onDelete: (flight: Flight) => void;
}

const FlightList = ({ data, onDelete }: Props) => {
  if (!data || data.length === 0) {
    return (
      <>
        <h3>Flights</h3>
        <div>Empty</div>
      </>
    );
  }

  return (
    <>
      <h3>Flights</h3>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell>Place Destination</TableCell>
              <TableCell align="right">Place Source</TableCell>
              <TableCell align="right">Ticket Price</TableCell>
              <TableCell align="right">Date Source</TableCell>
              <TableCell align="right">Date Destination</TableCell>
              <TableCell align="right">Total Tickets</TableCell>
              <TableCell align="right">Bought Tickets</TableCell>
              <TableCell align="right">Action</TableCell>
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
                <TableCell align="right">{d.ticketPrice}</TableCell>
                <TableCell align="right">
                  {new Date(d.dateSource).toLocaleString()}
                </TableCell>
                <TableCell align="right">
                  {new Date(d.dateDestination).toLocaleString()}
                </TableCell>
                <TableCell align="right">{d.totalTickets}</TableCell>
                <TableCell align="right">{d.boughtTickets}</TableCell>
                <TableCell align="right">
                  <Button
                    variant="outlined"
                    color="error"
                    onClick={() => onDelete(d)}
                  >
                    Delete
                  </Button>
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
