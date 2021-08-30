/* Base64. */
import base64 from 'react-native-base64'
/* Helpers. */
import axios from 'src/helpers/axios-helper'
import { setToken, decodeToken } from 'src/helpers/auth-helper'
/* Constants. */
import { LOGIN } from 'src/constants'

export {
    login
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