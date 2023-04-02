import { CircularProgress, Container } from "@mui/material";
import CustomerTicket from "../model/customerTicket";
import { useCallback, useEffect, useState } from "react";
import axios, { HttpStatusCode } from "axios";
import TicketList from "../components/ticket-list/TicketList";

const MyTickets = () => {
  const [myTickets, setMyTickets] = useState<CustomerTicket[]>();
  const [error, setError] = useState<any>(null);
  const [isPending, setIsPending] = useState(true);
  const ticketsUrl = `${process.env.REACT_APP_API}/tickets`;

  const fetchMyTickets = useCallback(async() => {
    const userId = '6421dbe6ef986c1e2cbbd5bd'; // TODO: Remove after login is implemented
    setIsPending(true);
    setError(null);
    try {
      const res = await axios.get<CustomerTicket[]>(`${ticketsUrl}?userId=${userId}`, {
        "headers": {
          "Accept": "application/json"
        }
      });
      if (res.status === HttpStatusCode.Ok) {
        setMyTickets(res.data)
        setIsPending(false);
      } else {
        console.log("Unknown error")
        setIsPending(false);
        setError("unknown error occured");
      }
    }
    catch (err: any) {
      setIsPending(false);
      setError("Error fetching tickets");
      console.log(err);
    }
  }, [ticketsUrl]);

  useEffect(() => {
    fetchMyTickets()
  }, [fetchMyTickets])

  if (isPending) return ( 
    <Container  sx={{ textAlign: 'center', py: 5}} >
      <CircularProgress /> 
    </Container>
  );
  if (error) return ( 
    <Container  sx={{ textAlign: 'center', py: 5}} >
      <p className="error-block">{error}</p> 
    </Container>
  );

  return (
    <Container>
      <TicketList tickets={myTickets}/>
    </Container>
  );
}

export default MyTickets;