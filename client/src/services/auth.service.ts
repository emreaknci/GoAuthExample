import BaseService from './_base.service';

const AuthService = {
    login: async (dto: { email: string; password: string; }) => await BaseService.post('/auth/login', dto),
    register: async (dto: { email: string; password: string; }) => await BaseService.post('/auth/register', dto),

}

export default AuthService;