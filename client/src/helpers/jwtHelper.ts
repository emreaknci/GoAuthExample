import { JwtPayload, jwtDecode } from "jwt-decode";

export interface JwtInfo {
    user: string;
    exp: number;
}

export const JwtHelper = {
    decode: (token: string): JwtPayload => {
        return jwtDecode(token);
    },

    decodedTokenToClaims: (decoded: any): JwtInfo => {
        return {
            user: decoded['user'],
            exp: decoded['exp']
        }
    },

    getTokenInfos: (token: string): JwtInfo => {
        const decoded = JwtHelper.decode(token);
        return JwtHelper.decodedTokenToClaims(decoded);
    },

    getUserId: (token: string): number => {
        const claims = JwtHelper.getTokenInfos(token);
        return parseInt(claims.user);
    }
}