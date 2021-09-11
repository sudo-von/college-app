import axios from 'axios'
/* Constants. */
import { API_URL } from 'src/constants/endpoints'
import { ERRORS } from 'src/constants/errors'
/* Helpers. */
import { getToken, deleteToken } from './auth-helper'
/* Refs. */
import { navigate } from 'src/refs/navigation.ref'

const config = {
    baseURL: API_URL,
    timeout: 1000,
    headers: {
        'Content-Type': 'application/json'
    }
}

/* Axios instances. */
const publicAxios = axios.create(config)
const protectedAxios = axios.create(config)

/* Public interceptors. */
publicAxios.interceptors.response.use(response => response, error => {
    if (error?.response){
        if(error.response.data?.code){
            if(Object.prototype.hasOwnProperty.call(ERRORS, error.response.data.code)){
                throw new Error(ERRORS[error.response.data.code])
            }
            console.warn(`Response code: [${error.response.data.code}] has not been implemented in the client.`)
            throw new Error(ERRORS['ERROR_NOT_IMPLEMENTED'])
        }
        console.warn(`Response code: [${error.response.data}] has not been implemented in the server.`)
        throw new Error(ERRORS['ERROR_NOT_IMPLEMENTED'])
    }
    throw new Error(ERRORS['UNKNOWN_ERROR'])
})

/* Protected interceptors. */
protectedAxios.interceptors.request.use(async config => {
    try{
        const token = await getToken()
        if (!config.headers.Authorization && token) {
            config.headers.Authorization = token
        }
        return config
    }catch(error){
        throw error
    }
}, error => Promise.reject(error))

protectedAxios.interceptors.response.use(response => response, async error => {
    if (error?.response){
        if (error?.response){
            if (error.response?.status === 401){
                await deleteToken()
                navigate('/logout')
            }
            if(error.response.data?.code){
                if(Object.prototype.hasOwnProperty.call(ERRORS, error.response.data.code)){
                    throw new Error(ERRORS[error.response.data.code])
                }
                console.warn(`Response code: [${error.response.data.code}] has not been implemented in the client.`)
                throw new Error(ERRORS['ERROR_NOT_IMPLEMENTED'])
            }
            console.warn(`Response code: [${error.response.data}] has not been implemented in the server.`)
            throw new Error(ERRORS['ERROR_NOT_IMPLEMENTED'])
        }
    }
    throw new Error(ERRORS['UNKNOWN_ERROR'])
})

export {
    publicAxios,
    protectedAxios
}