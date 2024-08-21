import React, { useState } from 'react';
import { useForm, FieldValues, Path, DefaultValues } from 'react-hook-form';
import { Button, TextField, Box, Typography, Container, Grid, Paper, IconButton, InputAdornment } from '@mui/material';
import { SxProps } from '@mui/system';
import { VisibilityOff, Visibility } from '@mui/icons-material';

export interface FormField {
    id: string;
    label: string;
    type: string;
    defaultValue?: string;
    validation: Record<string, any>;
    grid?: { xs?: number; sm?: number; md?: number; lg?: number; xl?: number };
    sx?: SxProps;
}

interface FormProps<T extends FieldValues> {
    title: string;
    icon?: any;
    fields: FormField[];
    onSubmit: (data: T) => void;
    defaultValues?: DefaultValues<T>;
}

const headerStyles: SxProps = {
    my: 3,
    fontWeight: 500,
};



const CustomForm = <T extends FieldValues>({ icon, title, fields, onSubmit, defaultValues }: FormProps<T>) => {
    const [showPassword, setShowPassword] = useState(false);

    const toggleShowPassword = () => {
        setShowPassword(!showPassword);
    };

    const { register, handleSubmit, formState: { errors } } = useForm<T>({ defaultValues });

    return (
        <>
            {icon && icon}
            <Typography component="h1" variant="h5" sx={headerStyles}>
                {title}
            </Typography>
            <form onSubmit={handleSubmit(onSubmit)} noValidate>
                <Grid container spacing={3}>
                    {fields.map((field) => (
                        <Grid item key={field.id} xs={12} {...field.grid} sx={field.sx}>
                            <TextField
                                fullWidth
                                id={field.id}
                                label={field.label}
                                type={showPassword ? 'text' : field.type}
                                autoComplete='off'
                                {...register(field.id as Path<T>, field.validation)}
                                error={!!errors[field.id]}
                                helperText={errors[field.id]?.message as React.ReactNode}
                                margin="normal"
                                variant="standard"
                                InputProps={{
                                    endAdornment: (
                                        <>
                                            {(field.type === 'password') &&
                                                <InputAdornment position="end">
                                                    <IconButton onMouseDown={toggleShowPassword} onMouseUp={toggleShowPassword} edge="end">
                                                        {showPassword ? <VisibilityOff /> : <Visibility />}
                                                    </IconButton>
                                                </InputAdornment>}

                                        </>
                                    ),
                                }}

                            />
                        </Grid>
                    ))}
                </Grid>
                <Button type="submit" fullWidth variant="outlined" color="primary" sx={{ my: 3 }} >
                    Submit
                </Button>
            </form>
        </>
    );
};

export default CustomForm;
