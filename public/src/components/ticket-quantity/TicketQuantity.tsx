import { Box, Button, Modal, Stack } from "@mui/material";
import { useState } from "react";

interface TicketQuantityProps {
  isOpen: boolean;
  showModal: (show: boolean) => void;
  buyTickets: (quantity: number | undefined) => void;
}

const style = {
  position: 'absolute' as 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  border: '2px solid var(--primary-color)',
  boxShadow: 24,
  p: 4,
};

const TicketQuantity = ({ isOpen, showModal, buyTickets }: TicketQuantityProps) => {
  const handleClose = () => showModal(false);
  const [ticketQuantity, setTicketQuantity] = useState<any>('');
  const buttonDisabled = !ticketQuantity || ticketQuantity === 0;

  return (
    <div>
      <Modal
        open={isOpen}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <Stack direction="row" spacing={3} sx={{ mb: 3}}>
            <span>Quantity:</span>
            <input type="number" min="0" max="500" value={ticketQuantity} onChange={ (e) => setTicketQuantity(e.target.value ? parseInt(e.target.value) : '')}/>
          </Stack>
          <Button color="primary" variant="outlined" disabled={buttonDisabled} onClick={() => { buyTickets(ticketQuantity); handleClose(); } }>
            Buy
          </Button>
        </Box>
      </Modal>
    </div>
  );
};

export default TicketQuantity;