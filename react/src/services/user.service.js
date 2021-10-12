import { post, get, patch } from 'src/helpers/protected-axios-helper'

const signup = async (user) => {
    try{
        await post('/users', user)
    }catch(error){
        throw error
    }
}

const getUserByID = async (userID) => {
    try{
        const response = await get(`/users/${userID}`)
        const user = await response.data
        return user
    }catch(error){
        throw error
    }
}

const updateUserByID = async (userID, user) => {
    try{
        await patch(`/users/${userID}`, user)
        return '¡El usuario ha sido actualizado con éxito!'
    }catch(error){
        throw error
    }
}

export {
    signup,
    getUserByID,
    updateUserByID
}