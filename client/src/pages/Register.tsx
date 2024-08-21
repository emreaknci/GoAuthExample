import { Container } from '@mui/material';
import CustomForm, { FormField } from '../components/CustomForm';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import AuthService from '../services/auth.service';

interface RegisterFormData {
    email: string;
    password: string;
}

const Register = () => {
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
                minLength: {
                    value: 6,
                    message: 'Password must be at least 6 characters'
                }
            },
            grid: { xs: 12 }
        }
    ];

    const onSubmit = (data: RegisterFormData) => {
        AuthService.register(data).then(response => {
            console.log('Form Data:', response);
        });
    };

    return (
        <>
            <CustomForm<RegisterFormData> title="Register" icon={<LockOutlinedIcon fontSize="large" />} fields={fields}
                defaultValues={{ email: 'emreaknci@github.com', password: '123456' }} onSubmit={onSubmit} />
        </>
    )
};

export default Register;