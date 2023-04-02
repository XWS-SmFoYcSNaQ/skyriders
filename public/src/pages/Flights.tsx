import { useCallback, useEffect, useState } from "react";
import axios from "axios";
import FlightForm from "../components/flight-form/FlightForm";
import FlightList from "../components/flight-list/FlightList";
import { Box } from "@mui/system";
import { Flight } from "../model/flight";
import { Container } from "@mui/material";
import FlightFilter from "../components/flight-filter/FlightFilter";
import { createQueryObject } from "../utils/utils";

const Flights = () => {
  const [data, setData] = useState([]);
  const [filters, setFilters] = useState<any>({});

  const fetchData = useCallback(async () => {
    const query = createQueryObject(filters);
    try {
      const res = await axios.get('flight', {
        params: query,
      });
      setData(res.data);
    } catch (err) {
      console.log(err);
    }
  }, [filters]);

  useEffect(() => {
    fetchData();
  }, [fetchData]);

  const onFlightSubmit = async (flight: Flight) => {
    try {
      await axios.post('flight', flight);
      fetchData();
    } catch (err) {
      console.log(err);
    }
  };

  const onFlightDelete = async (flight: Flight) => {
    try {
      await axios.delete(`flight/${flight.id}`);
      fetchData();
    } catch (err) {
      console.log(err);
    }
  };

  const onFilterChanged = (filters: any) => {
    setFilters(filters);
  };

  return (
    <div>
      <Container maxWidth="lg">
        <Box sx={{ padding: "30px 30px 130px 30px" }}>
          <FlightFilter onSubmit={onFilterChanged} />
          <Box sx={{ paddingTop: "30px" }}>
            <FlightList onDelete={onFlightDelete} data={data} />
          </Box>
          <Box sx={{ paddingTop: "30px" }}>
            <FlightForm onSubmit={onFlightSubmit} />
          </Box>
        </Box>
      </Container>
    </div>
  );
};

export default Flights;
