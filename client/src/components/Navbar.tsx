import { AppBar, Toolbar, Typography, Button, useTheme, Box, IconButton, InputBase, Switch, Drawer, List, ListItem, ListItemText, useMediaQuery, Container, Menu, MenuItem } from '@mui/material';
import React, { useContext } from 'react';
import MenuIcon from '@mui/icons-material/Menu';
import { Link, useNavigate } from 'react-router-dom';
import ThemeSwitcher from './ThemeSwitcher';
import { AccountCircle } from '@mui/icons-material';
import { AuthContext } from '../contexts/AuthContext';





const Navbar = () => {
  const { isAuthenticated, isTokenChecked, logout } = useContext(AuthContext);


  const navigate = useNavigate();

  const theme = useTheme();
  const isMedium = useMediaQuery(theme.breakpoints.down('md'));

  const [anchorEl, setAnchorEl] = React.useState(null);

  const handleMenu = (event: any) => setAnchorEl(event.currentTarget);
  const handleClose = () => setAnchorEl(null);


  return (
    <>
      <AppBar variant='outlined' elevation={0} position="static" sx={{
        bgcolor: theme.palette.background.default,
        color: theme.palette.text.primary,
        border: 'none',
      }}>
        <Container>
          <Toolbar style={{ padding: '0' }}>
            <Typography variant="h6" component="div" sx={{ flexGrow: 1, display: 'flex', alignItems: 'center', justifyContent: 'flex-start', cursor: "pointer" }} onClick={() => { navigate("/") }}>
              GoReactAuthExample
            </Typography>

            <Box sx={{ display: 'flex', alignItems: 'center', ml: 10 }}>
              <ThemeSwitcher />
              <IconButton
                edge="end" aria-label="account of current user"
                aria-controls="menu-appbar" aria-haspopup="true"
                onClick={handleMenu} color="inherit"
              >
                <AccountCircle fontSize="large" />
              </IconButton>
              <Menu
                id="menu-appbar" anchorEl={anchorEl}
                open={Boolean(anchorEl)}
                onClose={handleClose}
              >
                {isAuthenticated && isTokenChecked ? [
                  <MenuItem key="logout" onClick={() => logout()}>Logout</MenuItem>
                ] : [
                  <MenuItem key="login" onClick={() => { navigate("/login") }}>Login</MenuItem>,
                  <MenuItem key="register" onClick={() => { navigate("/register") }}>Register</MenuItem>
                ]}
              </Menu>
            </Box>
          </Toolbar>
        </Container>
      </AppBar>

    </>
  );
};

export default Navbar;
