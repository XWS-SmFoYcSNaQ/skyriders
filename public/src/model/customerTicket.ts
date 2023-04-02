export default interface CustomerTicket {
  flightId: string;
  flightDateSource: string;
  flightDateDestination: string;
  flightPlaceSource: string;
  flightPlaceDestination: string;
  flightTicketPrice: number;
  quantity: number;
  fullName: string;
}