import { Container } from '@mui/material';
import CustomForm, { FormField } from '../components/CustomForm';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import { AuthContext } from '../contexts/AuthContext';
import { useContext } from 'react';

interface LoginFormData {
    email: string;
    password: string;
}

const Login = () => {
    const authContext=useContext(AuthContext);
    const fields: FormField[] = [
        {
            id: 'email',
            label: 'Email',
            type: 'email',
            validation: {
                required: 'Email is required',
                pattern: {
                    value: /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$/,
                    message: 'Invalid email'
                }
            },
            grid: { xs: 12 }
        },
        {
            id: 'password',
            label: 'Password',
            type: 'password',
            validation: {
                required: 'Password is required',
                
            },
            grid: { xs: 12 }
        }
    ];

    const onSubmit = (data: LoginFormData) => {
        authContext.login(data);
    };

    return (
        <>

            <CustomForm<LoginFormData> title="Login" icon={<LockOutlinedIcon fontSize="large" />} fields={fields}
                defaultValues={{ email: 'emreaknci@github.com', password: '123456' }} onSubmit={onSubmit} />
        </>
    )
};

export default Login;