import './App.css'
import Login from './pages/Login'
import { Container, CssBaseline, ThemeProvider } from '@mui/material';
import { darkTheme, lightTheme } from './themes'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import Register from './pages/Register';
import Navbar from './components/Navbar';
import { useContext, useEffect } from 'react';
import { AuthContext } from './contexts/AuthContext';
import { CustomThemeContext } from './contexts/CustomThemeContext';


function App() {
  const { isAuthenticated, isTokenChecked } = useContext(AuthContext);
  const { theme } = useContext(CustomThemeContext);

  return (
    <ThemeProvider theme={theme ? darkTheme : lightTheme}>
      <CssBaseline />
      <Container
        sx={{
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        <Router>
          <Navbar />
          {isTokenChecked && (
            <Routes>
              <>
                <Route path="/" element={<Home />} />
                {!isAuthenticated && (
                  <>
                    <Route path="/login" element={<Login />} />
                    <Route path="/register" element={<Register />} />
                  </>
                )}
                <Route path="*" element={<Home />} />
              </>
            </Routes>
          )}
        </Router>
      </Container>
    </ThemeProvider>
  )
}

export default App
