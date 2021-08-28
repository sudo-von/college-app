import AsyncStorage from '@react-native-async-storage/async-storage'

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

const deleteToken = () => {
    try{
        await AsyncStorage.removeItem('token')
    }catch(error){
        return error
    }
}

export { 
    getToken,
    setToken,
    deleteToken
}