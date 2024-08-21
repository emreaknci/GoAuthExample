import { createRoot } from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer } from 'react-toastify';
import { AuthProvider } from './contexts/AuthContext.tsx';
import { CustomThemeProvider } from './contexts/CustomThemeContext.tsx';

createRoot(document.getElementById('root')!).render(
  <>
    <ToastContainer
      position="top-left"
      autoClose={3000}
      hideProgressBar={false}
      newestOnTop={false}
      closeOnClick
      rtl={false}
      pauseOnFocusLoss
      draggable
      pauseOnHover
    />
    <AuthProvider>
      <CustomThemeProvider>
          <App />
      </CustomThemeProvider>
    </AuthProvider>
  </>,
)
