import { Button, Container, ThemeProvider, createTheme } from "@mui/material";
import { useNavigate } from "react-router-dom";

const theme = createTheme();

const Unauthorized = () => {
    const navigate = useNavigate();

    const goBack = () => navigate(-1);

    return (
        <ThemeProvider theme={theme}>
            <Container component="main" maxWidth="xs">
                <h1>Unauthorized</h1>
                <br />
                <p>You do not have access to the requested page.</p>
                <Button onClick={goBack}>Go Back</Button>
            </Container>
        </ThemeProvider >
    )
}

export default Unauthorized