import { useCallback, useEffect, useState } from "react";
import axios, { AxiosError, HttpStatusCode } from "axios";
import FlightForm from "../components/flight-form/FlightForm";
import FlightList from "../components/flight-list/FlightList";
import { Box } from "@mui/system";
import { Flight } from "../model/flight";
import { Container } from "@mui/material";
import FlightFilter from "../components/flight-filter/FlightFilter";
import { createQueryObject } from "../utils/utils";
import BuyTicketRequest from "../model/buyTicketRequest";
import { ToastContainer, toast } from "react-toastify";

const Flights = () => {
  const [data, setData] = useState<Flight[]>([]);
  const [filters, setFilters] = useState<any>({});
  const [activeFlightId, setActiveFlightId] = useState<string>();

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

  const buyTickets = async (quantity: number | undefined) => {
    if (!quantity || quantity === 0) {
      return;
    }
    const buyTicketRequest : BuyTicketRequest = { flightId: activeFlightId, quantity};
    try { 
      const res = await axios.post(`tickets`, buyTicketRequest, {
        headers: { "Accept": "application/json" }
      })
      if (res.status === HttpStatusCode.Created) {
        setData((prevData) => {
          return prevData.map(x => x.id !== activeFlightId ? x : {...x, boughtTickets: x.boughtTickets! + quantity });
        });
        toast.success(`Thank you for buying ${quantity} tickets`, { position: toast.POSITION.TOP_CENTER });
      } else {
        toast.error(res.request?.response, { position: toast.POSITION.TOP_CENTER });
      }
    } catch (err: any) {
      console.log(err)
      toast.error(err.response.data, { position: toast.POSITION.TOP_CENTER });
    };
  }

  return (
    <div>
      <Container maxWidth="lg">
        <Box sx={{ padding: "30px 30px 130px 30px" }}>
          <FlightFilter onSubmit={onFilterChanged} />
          <Box sx={{ paddingTop: "30px" }}>
            <FlightList onDelete={onFlightDelete} data={data} buyTickets={buyTickets} setActiveFlightId={setActiveFlightId} />
          </Box>
          <Box sx={{ paddingTop: "30px" }}>
            <FlightForm onSubmit={onFlightSubmit} />
          </Box>
        </Box>
        <ToastContainer/>
      </Container>
    </div>
  );
};

export default Flights;
