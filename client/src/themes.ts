import { createTheme } from "@mui/material";

export const lightTheme = createTheme({
  palette: {
    mode: 'light',
    primary: {
      main: '#19857b',
      dark: '#b0272f',
      light: '#ff784e',
    },
    secondary: {
      main: '#19857b',
      dark: '#004c40',
      light: '#4fb3bf',
    },
    background: {
      default: '#ffffff',
      paper: '#f4f4f5',

    },
  },
});

export const darkTheme = createTheme({
  palette: {
    mode: 'dark',
    primary: {
      main: '#ffffff',
      dark: '#00c89F',
      light: '#ff784e',
    },
    secondary: {
      main: '#4B6BFB',
      dark: '#004c40',
      light: '#4fb3bf',
    },
    background: {
      default: '#181a2a',
      paper: '#141624',
    },
    divider: '#242535',
  },
});