import { FormControlLabel, styled, Switch, useTheme } from '@mui/material'
import React, { useContext } from 'react'
import Brightness7Icon from '@mui/icons-material/Brightness7';
import Brightness4Icon from '@mui/icons-material/Brightness4';
import { CustomThemeContext } from '../contexts/CustomThemeContext';


const CustomSwitch = styled(Switch)(({ theme }) => ({
  '& .MuiSwitch-switchBase': {
    color: 'black',
    '&.Mui-checked': {
      color: 'white',
      '& + .MuiSwitch-track': {
        backgroundColor: '#0476D0',
      },
    },
    '&.Mui-disabled': {
      color: 'gray',
    },
  },
  '& .MuiSwitch-track': {
    backgroundColor: 'gray',
  },
  '& .MuiSwitch-thumb': {
    color: 'white',
  },
}));

const ThemeSwitcher = () => {
    const themeContext = useContext(CustomThemeContext);

  return (
    <FormControlLabel
      label={undefined}
      control={
        <CustomSwitch
          checked={themeContext.theme}
          onChange={themeContext.toggleTheme}
          name="themeSwitch"
          aria-label="toggle dark mode"
          icon={<Brightness4Icon />}
          checkedIcon={<Brightness7Icon />}
        />
      }
    />
  );
};


export default ThemeSwitcher