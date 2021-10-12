
import { Base64 } from 'js-base64'
import { post } from 'src/helpers/public-axios-helper'
import { setToken, decodeToken } from './token.service'

const login = async ({ email, password }) => {
    try{
        const response = await post('/auth/login',{},{
            'headers': { 
                'Authorization': 'Basic ' + Base64.encode(`${email}:${password}`)
            }
        })
        await setToken(response.headers.authorization)
        const user = await decodeToken(response.headers.authorization)
        return user
    }catch(error){
        throw error
    }
}

export { login }