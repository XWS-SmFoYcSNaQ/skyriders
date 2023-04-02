import { Box, Button, TextField } from "@mui/material";
import { useState } from "react";
import { Flight } from "../../model/flight";
import { parseDateToDatetimeLocal, convertDate } from "../../utils/utils";

const getInitialData = (): object => {
  return {
    dateSource: parseDateToDatetimeLocal(new Date()),
    dateDestination: parseDateToDatetimeLocal(new Date()),
  };
};

interface Props {
  onSubmit: (flight: Flight) => void;
}

const FlightForm = ({ onSubmit }: Props) => {
  const [data, setData] = useState<Flight>(getInitialData());

  const handleSubmit = (e: any) => {
    e.preventDefault();
    onSubmit({
      ...data,
      dateSource: convertDate(data.dateSource),
      dateDestination: convertDate(data.dateDestination),
    });
    setData(getInitialData());
  };

  return (
    <form onSubmit={handleSubmit}>
      <h3>Add new flight</h3>
      <div className="grid">
        <TextField
          required
          label="Place Destination"
          value={data.placeDestination ?? ""}
          onChange={(e) =>
            setData({ ...data, placeDestination: e.target.value })
          }
        />
        <TextField
          required
          label="Place Source"
          value={data.placeSource ?? ""}
          onChange={(e) => setData({ ...data, placeSource: e.target.value })}
        />
        <TextField
          required
          label="Ticket Price"
          type="number"
          value={data.ticketPrice ?? ""}
          onChange={(e) =>
            setData({
              ...data,
              ticketPrice:
                e.target.value === "" ? undefined : parseFloat(e.target.value),
            })
          }
        />
        <TextField
          label="Date Source"
          type="datetime-local"
          value={data.dateSource ?? ""}
          onChange={(e) => setData({ ...data, dateSource: e.target.value })}
        />
        <TextField
          label="Date Destination"
          type="datetime-local"
          value={data.dateDestination ?? ""}
          onChange={(e) =>
            setData({ ...data, dateDestination: e.target.value })
          }
        />
        <TextField
          label="Total Tickets"
          type="number"
          value={data.totalTickets ?? ""}
          onChange={(e) =>
            setData({
              ...data,
              totalTickets:
                e.target.value === "" ? undefined : parseInt(e.target.value),
            })
          }
        />
        <TextField
          label="Bought Tickets"
          type="number"
          value={data.boughtTickets ?? ""}
          onChange={(e) =>
            setData({
              ...data,
              boughtTickets:
                e.target.value === "" ? undefined : parseInt(e.target.value),
            })
          }
        />
      </div>
      <Box marginTop="20px" sx={{ display: "flex", justifyContent: "right" }}>
        <Button size="large" variant="outlined" type="submit">
          Create Flight
        </Button>
      </Box>
    </form>
  );
};

export default FlightForm;
