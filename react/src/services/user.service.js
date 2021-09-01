/* Base64. */
import base64 from 'react-native-base64'
/* Helpers. */
import axios from 'src/helpers/axios-helper'
import { setToken, decodeToken } from 'src/helpers/auth-helper'
/* Constants. */
import { LOGIN, SIGNUP } from 'src/constants'

export {
    login,
    signup
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
        const response = await axios.post(SIGNUP,user)
        console.log(response)
        return "¡Registro éxitoso!"
    }catch(error){
        if(error.response?.data.message){
            throw new Error(error.response.data.message)
        }
        throw new Error("Ha ocurrido un error, intenta de nuevo más tarde")
    }
}