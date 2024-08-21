import BaseService from './_base.service';

const AuthService = {
    login: async (dto: { email: string; password: string; }) => await BaseService.post('/auth/login', dto),
    register: async (dto: { email: string; password: string; }) => await BaseService.post('/auth/register', dto),
    refreshToken: async (dto: { refresh_token: string; }) => await BaseService.post('/auth/refresh', dto),
}

export default AuthService;