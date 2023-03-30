import { Box, Button, TextField, Select, MenuItem} from "@mui/material"
import { useState } from "react"
import { User } from "../../model/user"
import classes from './RegistrationForm.module.css';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { DateTimePicker, DatePicker, LocalizationProvider } from '@mui/x-date-pickers';
import React from "react"

interface Props {
  onSubmit: (user: User) => void
}

const RegistrationForm = ({onSubmit}: Props) => {
  const [data, setData] = useState<User>({})
  const [gender] = React.useState<number>(0);
  const handleDateChange = (newDate: Date | null) => {
    data.dateOfBirth = newDate?.toISOString()
    console.log(data.dateOfBirth)
  };


  const handleSubmit = (e: any) => {
    e.preventDefault();
    onSubmit(data)
  }

  return (
    <form onSubmit={handleSubmit}>
      <h2>Register now</h2>
      <div className={classes.grid}>
        <TextField
          required
          label="First name"
          onChange={(e) => setData({...data, firstname: e.target.value})}
          />
        <TextField
          required
          label="Last name"
          onChange={(e) => setData({...data, lastname: e.target.value})}
        />
        <Select
            value={gender}
            label="Gender"
            onChange={(e) => setData({...data, gender: gender})}>
                <MenuItem value={0}>Male</MenuItem>
                <MenuItem value={1}>Female</MenuItem>
        </Select>
        <LocalizationProvider
          dateAdapter={AdapterDateFns}>
          <DatePicker
            label="Date of birth"
            onChange={handleDateChange}
            maxDate={new Date()}
          />
        </LocalizationProvider>
        <TextField
          required
          label="Phone number"
          onChange={(e) => setData({...data, phone: e.target.value})}
        />
        <TextField
          required
          label="Nationality"
          onChange={(e) => setData({...data, nationality: e.target.value})}
        />
        <TextField
          required
          label="Email"
          onChange={(e) => setData({...data, email: e.target.value})}
        />
        <TextField
          required
          label="Password"
          type="password"
          onChange={(e) => setData({...data, password: e.target.value})}
        />
      </div>
      <Box marginTop="20px">
        <Button size="large" variant="contained" type="submit" >Register</Button>
      </Box>
    </form>
  )
}

export default RegistrationForm