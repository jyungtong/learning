import { createContext, useContext, useState } from 'react';

const ThemeContext = createContext();
const ThemeUpdateContext = createContext();

export function useTheme() {
  return useContext(ThemeContext);
}

export function useThemeUpdate() {
  return useContext(ThemeUpdateContext);
}

export function ThemeProvider({ children }) {
  const [isRedTheme, setRedTheme] = useState(false);

  function toggleTheme() {
    setRedTheme(prev => !prev);
  }

  return (
    <ThemeContext.Provider value={isRedTheme}>
      <ThemeUpdateContext.Provider value={toggleTheme}>
        {children}
      </ThemeUpdateContext.Provider>
    </ThemeContext.Provider>
  )
}