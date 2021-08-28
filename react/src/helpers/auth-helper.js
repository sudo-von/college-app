import AsyncStorage from '@react-native-async-storage/async-storage'

const setToken = async (token) => {
    try{
        await AsyncStorage.setItem('token', token)
    }catch(error){
        return error
    }
}

export { 
    setToken
}