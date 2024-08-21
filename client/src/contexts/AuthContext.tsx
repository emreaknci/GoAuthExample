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
        const getNewAccessToken = async () => {
            const refreshToken = StorageService.getRefreshToken();
            if (refreshToken) {
                await AuthService.refreshToken({ refresh_token: refreshToken }).then(res => {
                    StorageService.setAccessToken(res.data.data.access_token);
                    StorageService.setRefreshToken(res.data.data.refresh_token);
                }).catch(err => {
                    toast.error(err.response.data.message);
                    logout();
                })
            }
        }
        const token = StorageService.getAccessToken();
        if (token && isAuthenticated && isTokenChecked) {
            const exp = JwtHelper.getTokenInfos(token).exp;

            const remainingTime = exp * 1000 - new Date().getTime(); 

            if (remainingTime < 15*60*1000) { // token will expire in 15 minutes
                getNewAccessToken();
            }

            const interval = setInterval(() => {
                getNewAccessToken();
            }, 15*60*1000) // check every 15 minutes

            return () => clearInterval(interval);
        }

    }, [isAuthenticated, isTokenChecked])


    const login = async (dto: { email: string, password: string }) => {
        await AuthService.login(dto).then(res => {
            StorageService.setAccessToken(res.data.data.access_token);
            StorageService.setRefreshToken(res.data.data.refresh_token);
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
        StorageService.clearTokens();
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