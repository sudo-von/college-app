/* Helpers. */
import axios from 'src/helpers/axios-helper'
import { authorizationRequestInterceptor, authorizationerrorResponseInterceptor } from 'src/helpers/interceptor-helper'
/* Constants. */
import { CONTACT, USERS } from 'src/constants/endpoints'

/* Add interceptors to the axios instance. */
axios.interceptors.request.use(config => authorizationRequestInterceptor(config), error => Promise.reject(error))
axios.interceptors.response.use(response => response, error => authorizationerrorResponseInterceptor(error))

export {
    getContactByUserID,
    updateContactByUserID
}

const getContactByUserID = async (userID) => {
    try{
        const response = await axios.get(`${CONTACT}${USERS}/${userID}`)
        const user = await response.data
        console.log(user)
        return user
    }catch(error){
        console.log(error)
        throw new Error("Ha ocurrido un error, intenta de nuevo más tarde")
    }
}

const sendContactByUserID = async (userID, user) => {
    try{
        await axios.post(`${CONTACT}${USERS}/${userID}`, user)
    }catch(error){
        console.log(error)
        if(error.response?.data.message){
            throw new Error(error.response.data.message)
        }
        throw new Error("Ha ocurrido un error, intenta de nuevo más tarde")
    }
}

const updateContactByUserID = async (userID, user) => {
    try{
        await axios.patch(`${CONTACT}${USERS}/${userID}`, user)
    }catch(error){
        console.log(error)
        if(error.response?.data.message){
            throw new Error(error.response.data.message)
        }
        throw new Error("Ha ocurrido un error, intenta de nuevo más tarde")
    }
}