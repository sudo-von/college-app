/* Async storage. */
import AsyncStorage from '@react-native-async-storage/async-storage'
/* Base64. */
import Base64 from 'Base64'

const getToken = async () => {
    try{
        const token = await AsyncStorage.getItem('token')
        return token
    } catch(e) {
        return error
    }
}

const setToken = async (token) => {
    try{
        await AsyncStorage.setItem('token', token)
    }catch(error){
        return error
    }
}

const deleteToken = async () => {
    try{
        await AsyncStorage.removeItem('token')
    }catch(error){
        return error
    }
}

const decodeToken = async () => {
    try{
        const bearerToken = await getToken()
        const encodedToken = bearerToken.split('Bearer ')[1]
        const userToken = encodedToken.split('.')[1]
        const user = Base64.atob(userToken)
        const parsedUser = JSON.parse(user)
        return parsedUser
    }catch(error){
        return error
    }
}

export { 
    getToken,
    setToken,
    deleteToken,
    decodeToken
}