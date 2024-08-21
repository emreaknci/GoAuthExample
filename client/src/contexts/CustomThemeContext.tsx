import { createContext, useState } from "react";


export const CustomThemeContext = createContext({
    theme: false,
    toggleTheme: () => { }
});


export const CustomThemeProvider = ({ children }: any) => {
    const [theme, setTheme] = useState(localStorage.getItem('darkMode') === 'true' ? true : false);


    const toggleTheme = () => {
        setTheme(!theme);
        localStorage.setItem('darkMode', (!theme).toString());
    }

    return (
        <CustomThemeContext.Provider value={{ theme, toggleTheme }}>
            {children}
        </CustomThemeContext.Provider>
    )
}