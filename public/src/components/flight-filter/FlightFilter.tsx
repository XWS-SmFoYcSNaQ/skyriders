import { TextField, Box, Button } from "@mui/material";
import React from "react";
import { useState } from "react";

interface Props {
  onSubmit: (filters: any) => void;
}

const FlightFilter = ({ onSubmit }: Props) => {
  const [filters, setFilters] = useState<any>({});

  const handleSubmit = (e: any) => {
    e.preventDefault();
    onSubmit({
      ...filters,
      dateSource: filters.dateSource
        ? +new Date(filters.dateSource)
        : undefined,
      dateDestination: filters.dateDestination
        ? +new Date(filters.dateDestination)
        : undefined,
      boughtTickets:
        filters.boughtTickets !== "" ? filters.boughtTickets : undefined,
      totalTickets:
        filters.totalTickets !== "" ? filters.totalTickets : undefined,
    });
  };

  const handleClear = () => {
    setFilters({});
    onSubmit({});
  };

  return (
    <form onSubmit={handleSubmit}>
      <h3>Filters</h3>
      <div className="grid">
        <TextField
          label="Place Destination"
          value={filters.placeDestination ?? ""}
          onChange={(e) =>
            setFilters({ ...filters, placeDestination: e.target.value })
          }
        />
        <TextField
          label="Place Source"
          value={filters.placeSource ?? ""}
          onChange={(e) =>
            setFilters({ ...filters, placeSource: e.target.value })
          }
        />
        <TextField
          label="Date Source"
          type="datetime-local"
          value={filters.dateSource ?? ""}
          onChange={(e) =>
            setFilters({ ...filters, dateSource: e.target.value })
          }
        />
        <TextField
          label="Date Destination"
          type="datetime-local"
          value={filters.dateDestination ?? ""}
          onChange={(e) =>
            setFilters({ ...filters, dateDestination: e.target.value })
          }
        />
        <TextField
          label="Total Tickets"
          type="number"
          value={filters.totalTickets ?? ""}
          onChange={(e) =>
            setFilters({
              ...filters,
              totalTickets:
                e.target.value === "" ? "" : parseInt(e.target.value),
            })
          }
        />
        <TextField
          label="Bought Tickets"
          type="number"
          value={filters.boughtTickets ?? ""}
          onChange={(e) =>
            setFilters({
              ...filters,
              boughtTickets:
                e.target.value === "" ? "" : parseInt(e.target.value),
            })
          }
        />
      </div>
      <Box
        marginTop="20px"
        sx={{ display: "flex", justifyContent: "right", gap: "15px" }}
      >
        <Button
          size="large"
          variant="outlined"
          color="error"
          onClick={handleClear}
        >
          Clear
        </Button>
        <Button size="large" variant="outlined" type="submit">
          Search
        </Button>
      </Box>
    </form>
  );
};

export default FlightFilter;
