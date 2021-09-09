/* Base64. */
import base64 from 'react-native-base64'
/* Helpers. */
import { publicAxios, protectedAxios } from 'src/helpers/axios-helper'
import { setToken, decodeToken } from 'src/helpers/auth-helper'
/* Constants. */
import { LOGIN, SIGNUP, USERS } from 'src/constants/endpoints'
import { USER_ERRORS } from 'src/constants/errors'

export {
    login,
    signup,
    getUserByID,
    updateUserByID
}

/* Handles basic auth. */
const login = async ({ email, password }) => {
    try{
        const response = await publicAxios.post(LOGIN,{},{
            "headers": { 
                'Authorization': "Basic " + base64.encode(`${email}:${password}`)
            }
        })
        await setToken(response.headers.authorization)
        const user = await decodeToken(response.headers.authorization)
        return user
    }catch(error){
        throw new Error(USER_ERRORS[error.response.data.code]['esp'])
    }
}

/* Adds a new user. */
const signup = async (user) => {
    try{
        await protectedAxios.post(SIGNUP, user)
        return "¡Registro éxitoso!"
    }catch(error){
        throw new Error(USER_ERRORS[error.response.data.code]['esp'])
    }
}

const getUserByID = async (userID) => {
    try{
        const response = await protectedAxios.get(`${USERS}/${userID}`)
        const user = await response.data
        return user
    }catch(error){
        throw new Error(USER_ERRORS[error.response.data.code]['esp'])
    }
}

const updateUserByID = async (userID, user) => {
    try{
        await protectedAxios.patch(`${USERS}/${userID}`, user)
    }catch(error){
        throw new Error(USER_ERRORS[error.response.data.code]['esp'])
    }
}