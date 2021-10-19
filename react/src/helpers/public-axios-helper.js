import axios from 'axios'
import { configuration } from './axios-helper'
import { ERRORS } from 'src/constants/errors'

const client = axios.create(configuration)

const handleError = (error) => {
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
}

client.interceptors.response.use(response => response, handleError)

const { get, post, patch, delete : destroy } = client
export { get, post, patch, destroy }