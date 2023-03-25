import { Container } from "@mui/material";

const Error = () => {
  return (
    <Container sx={{
      textAlign: 'center',
      py: 6
    }}>
      <h1 style={{ color: 'crimson'}}>Error, page doesn't exist.</h1>
    </Container>
  )
};

export default Error;