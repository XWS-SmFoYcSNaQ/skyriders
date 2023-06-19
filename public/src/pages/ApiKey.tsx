import { Box, Button, Checkbox, Container, FormControlLabel, Tooltip, Typography } from "@mui/material";
import { LocalizationProvider, DateTimePicker } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import axios from "axios";
import { useEffect, useState } from "react";
import dayjs, { Dayjs } from "dayjs";
import ContentCopyIcon from '@mui/icons-material/ContentCopy';

interface APIKey {
  keyString: string;
  expiration?: Date;
}

const ApiKey = () => {
  const [apiKey, setApiKey] = useState<APIKey | null>(null);
  const [duration, setDuration] = useState<Dayjs | null>(null);
  const [expires, setExpires] = useState(false);

  const registerAPI = async () => {
    try {
      const response = await axios.post('user/apikey', duration)
      setApiKey(response.data);
    } catch (error) {
      console.log('Error:', error);
    }
  };

  const revokeAPI = async () => {
    try {
      await axios.delete('user/apikey')
      setApiKey(null);
    } catch (error) {
      console.log('Error:', error);
    }
  };

  const fetchApiKey = async () => {
    try {
      const response = await axios.get('user/apikey');
      setApiKey(response.data);
    } catch (error) {
      console.log('Error:', error);
    }
  };

  useEffect(() => {
    fetchApiKey();
  }, []);

  const formattedDate = apiKey?.expiration
    ? dayjs(apiKey.expiration).format('YYYY-MM-DD HH:mm')
    : '';

  return (
    <LocalizationProvider dateAdapter={AdapterDayjs}>
      <Container>
        {apiKey ? (
          <div>
            <h2> Your API Key </h2>
            <Box sx={{ display: 'flex', alignItems: 'center' }}>
              <Typography sx={{ marginRight: '0.5rem' }}>Key:</Typography>
              <Typography>{apiKey.keyString}</Typography>
              <Tooltip title="Copy to clipboard" placement="top">
                <Button startIcon={<ContentCopyIcon />} onClick={() => {
                  navigator.clipboard.writeText(apiKey.keyString);
                }}>
                </Button>
              </Tooltip>
            </Box>
            {apiKey?.expiration && (
              <span style={{ opacity: 0.6 }}> (Valid until {formattedDate})</span>
            )}
            <div style={{ marginTop: '30px' }}>
              <Button variant="contained" onClick={revokeAPI}> Revoke API Key </Button>
            </div>
          </div>
        ) : (
          <div>
            <h2> Register for a new API Key </h2>
            <p> Registering for an API Key will allow another users to book a flight on your behalf. </p>
            <div style={{ minHeight: '60px' }}>
              <FormControlLabel
                control={
                  <Checkbox checked={expires} onChange={(event) => {
                    setExpires(event.target.checked)
                    if (!event.target.checked) {
                      setDuration(null);
                    }
                  }
                  } />
                }
                label="Expires"
              />
              {expires && (
                <DateTimePicker
                  label="Duration"
                  value={duration}
                  minDateTime={dayjs()}
                  ampm={false}
                  onChange={(date) => setDuration(date)}
                />
              )}
            </div>
            <div style={{ marginTop: '20px' }}>
              <Button variant="contained" onClick={registerAPI}> Generate </Button>
            </div>
          </div>
        )}
      </Container>
    </LocalizationProvider>
  )
}

export default ApiKey;