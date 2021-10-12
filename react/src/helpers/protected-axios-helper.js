import axios from 'axios'
import { configuration } from './axios-helper'
import { getToken, deleteToken } from './auth-helper'
import { navigate } from 'src/refs/navigation.ref'
import { ERRORS } from 'src/constants/errors'

const client = axios.create(configuration)

const handleHeader = async (config) => {
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

const handleError = async error => {
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
}

client.interceptors.request.use(handleHeader, error => Promise.reject(error))
client.interceptors.response.use(response => response, handleError)

const { get, post, patch, delete : destroy } = client
export { get, post, patch, destroy }