/* Base64. */
import base64 from 'react-native-base64'
/* Helpers. */
import axios from 'src/helpers/axios-helper'
import { authorizationRequestInterceptor, authorizationerrorResponseInterceptor } from 'src/helpers/interceptor-helper'
import { setToken, decodeToken } from 'src/helpers/auth-helper'
/* Constants. */
import { LOGIN, SIGNUP, USERS } from 'src/constants/endpoints'

/* Add interceptors to the axios instance. */
axios.interceptors.request.use(config => authorizationRequestInterceptor(config), error => Promise.reject(error))
axios.interceptors.response.use(response => response, error => authorizationerrorResponseInterceptor(error))


export {
    login,
    signup,
    getUserByID,
    updateUserByID
}

/* Handles basic auth. */
const login = async ({ email, password }) => {
    try{
        const response = await axios.post(LOGIN,{},{
            "headers": { 
                'Authorization': "Basic " + base64.encode(`${email}:${password}`)
            }
        })
        await setToken(response.headers.authorization)
        const user = await decodeToken(response.headers.authorization)
        return user
    }catch(error){
        if(error.response?.status == 401){
            throw new Error("Credenciales incorrectas")
        }
        throw new Error("Ha ocurrido un error, intenta de nuevo más tarde")
    }
}

/* Adds a new user. */
const signup = async (user) => {
    try{
        await axios.post(SIGNUP, user)
        return "¡Registro éxitoso!"
    }catch(error){
        if(error.response?.data.message){
            throw new Error(error.response.data.message)
        }
        throw new Error("Ha ocurrido un error, intenta de nuevo más tarde")
    }
}

const getUserByID = async (userID) => {
    try{
        const response = await axios.get(`${USERS}/${userID}`)
        const user = await response.data
        return user
    }catch(error){
        throw new Error("Ha ocurrido un error, intenta de nuevo más tarde")
    }
}

const updateUserByID = async (userID, user) => {
    try{
        await axios.patch(`${USERS}/${userID}`, user)
    }catch(error){
        if(error.response?.data.message){
            throw new Error(error.response.data.message)
        }
        throw new Error("Ha ocurrido un error, intenta de nuevo más tarde")
    }
}