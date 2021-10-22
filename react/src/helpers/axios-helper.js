export const configuration = {
    baseURL: process.env.NODE_ENV === 'development' ? 'http://192.168.1.78:4000' : 'https://www.website.com:port/',
    timeout: 1000,
    headers: {
        'Content-Type': 'application/json'
    }
}