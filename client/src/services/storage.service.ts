
const StorageService = {

    setItem(key: string, value: string) {
        localStorage.setItem(key, value);
    },

    getItem(key: string) {
        return localStorage.getItem(key);
    },

    removeItem(key: string) {
        localStorage.removeItem(key);
    },

    clear() {
        localStorage.clear();
    },

    getAccessToken() {
        return sessionStorage.getItem("access_token");
    },

    getRefreshToken() {
        return sessionStorage.getItem("refresh_token");
    },

    setAccessToken(token: string) {
        sessionStorage.setItem("access_token", token);
    },

    setRefreshToken(token: string) {
        sessionStorage.setItem("refresh_token", token);
    },

    clearTokens() {
        sessionStorage.removeItem("access_token");
        sessionStorage.removeItem("refresh_token");
    }
}

export default StorageService;