const configuration = {
    baseURL: process.env.NODE_ENV === 'development' ? 'http://192.168.0.4:4000' : 'https://www.website.com:port/',
    timeout: 1000,
    headers: {
        'Content-Type': 'application/json'
    }
}

export { configuration }