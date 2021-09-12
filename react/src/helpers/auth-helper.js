/* Async storage. */
import AsyncStorage from '@react-native-async-storage/async-storage'
/* Base64. */
import { Base64 } from 'js-base64'

const getToken = async () => {
    try{
        const token = await AsyncStorage.getItem('token')
        return token
    } catch(error) {
        throw error
    }
}

const setToken = async (token) => {
    try{
        await AsyncStorage.setItem('token', token)
    }catch(error){
        throw error
    }
}

const deleteToken = async () => {
    try{
        await AsyncStorage.removeItem('token')
    }catch(error){
        throw error
    }
}

const decodeToken = async () => {
    try{
        const bearerToken = await getToken()
        const encodedToken = bearerToken.split('Bearer ')[1]
        const userToken = encodedToken.split('.')[1]
        const user = Base64.decode(userToken)
        const parsedUser = JSON.parse(user)
        return parsedUser
    }catch(error){
        throw error
    }
}

export { 
    getToken,
    setToken,
    deleteToken,
    decodeToken
}