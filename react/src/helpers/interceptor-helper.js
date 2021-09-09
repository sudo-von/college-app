
import { getToken, deleteToken } from './auth-helper'
/* Refs. */
import { navigate } from 'src/refs/navigation.ref'

const authorizationRequestInterceptor = async (config) => {
    try{
        const token = await getToken()
        if (!config.headers.Authorization && token) {
            config.headers.Authorization = token
        }
        return config
    }catch(error){
        throw error
    }
}

const authorizationerrorResponseInterceptor = async (error) => {
    if (error.response?.status === 401) {
        await deleteToken()
        navigate('/logout')
    }
    return Promise.reject(error)
}

export {
    authorizationRequestInterceptor,
    authorizationerrorResponseInterceptor
}