import React, { createContext, useEffect, useState } from 'react'
import StorageService from '../services/storage.service';
import { useNavigate } from 'react-router-dom';
import { toast } from 'react-toastify';
import { JwtHelper } from '../helpers/jwtHelper';
import AuthService from '../services/auth.service';


export const AuthContext = createContext({
    currentUserId: null as number | null,
    isAuthenticated: false,
    isTokenChecked: false,
    isAdmin: false,
    login: (dto: { email: string, password: string }) => { },
    logout: () => { },
})


export const AuthProvider = ({ children }: any) => {
    const [currentUserId, setCurrentUserId] = useState<number | null>(null);
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [isTokenChecked, setIsTokenChecked] = useState(false);
    const [isAdmin, setIsAdmin] = useState(false);


    useEffect(() => {
        const checkToken = async () => {
            const token = StorageService.getAccessToken();
            if (token) {
                setIsAuthenticated(true);
                setIsTokenChecked(true);
                setCurrentUserId(JwtHelper.getUserId(token));
                return;
            }
            setIsAuthenticated(false);
            setIsTokenChecked(true);
            setCurrentUserId(null);
            return;
        }

        checkToken();
    }, [isAuthenticated])

    useEffect(() => {
        const token = StorageService.getAccessToken();
        if (token && isAuthenticated && isTokenChecked) {
            const exp = JwtHelper.getTokenInfos(token).exp;

            const remainingTime = exp * 1000 - new Date().getTime();
            console.log(remainingTime / 1000)
            if (remainingTime <= 0) {
                logout();
                toast.info("Your session has expired. Please sign in again.");
                return;
            } else {
                setTimeout(() => {
                    logout();
                    toast.info("Your session has expired. Please sign in again.");
                }, remainingTime)
            }
        }
    }, [isAuthenticated, isTokenChecked])


    const login = async (dto: { email: string, password: string }) => {
        await AuthService.login(dto).then(res => {
            StorageService.setAccessToken(res.data.data);
            setIsAuthenticated(true);
            setIsTokenChecked(true);
        }).catch(err => {
            toast.error(err.response.data.message);
            setIsAuthenticated(false);
            setIsTokenChecked(true);
        })
    }

    const logout = () => {
        setIsAuthenticated(false);
        setIsAdmin(false);
        setIsTokenChecked(false);
        setCurrentUserId(null);
        StorageService.clearAccessToken();
    }



    return (
        <AuthContext.Provider value={{
            currentUserId,
            isAuthenticated,
            isTokenChecked,
            isAdmin,
            login,
            logout,
        }}>
            {children}
        </AuthContext.Provider>
    )
}