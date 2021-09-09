/* Helpers. */
import axios from 'src/helpers/axios-helper'
/* Constants. */
import { CONTACT, USERS } from 'src/constants/endpoints'

export {
    getContactByUserID,
    updateContactByUserID
}

const getContactByUserID = async (userID) => {
    try{
        const response = await axios.get(`${CONTACT}${USERS}/${userID}`)
        const user = await response.data
        console.log(user)
        return user
    }catch(error){
        console.log(error)
        throw new Error("Ha ocurrido un error, intenta de nuevo más tarde")
    }
}

const sendContactByUserID = async (userID, user) => {
    try{
        await axios.post(`${CONTACT}${USERS}/${userID}`, user)
    }catch(error){
        console.log(error)
        if(error.response?.data.message){
            throw new Error(error.response.data.message)
        }
        throw new Error("Ha ocurrido un error, intenta de nuevo más tarde")
    }
}

const updateContactByUserID = async (userID, user) => {
    try{
        await axios.patch(`${CONTACT}${USERS}/${userID}`, user)
    }catch(error){
        console.log(error)
        if(error.response?.data.message){
            throw new Error(error.response.data.message)
        }
        throw new Error("Ha ocurrido un error, intenta de nuevo más tarde")
    }
}