/* Helpers. */
import { protectedAxios } from 'src/helpers/axios-helper'
/* Constants. */
import { CONTACT, USERS } from 'src/constants/endpoints'

export {
    getContactByUserID,
    sendContactByUserID,
    updateContactByUserID
}

const getContactByUserID = async (userID) => {
    try{
        const response = await protectedAxios.get(`${CONTACT}${USERS}/${userID}`)
        const user = await response.data
        return user
    }catch(error){
        throw new Error(error.message)
    }
}

const sendContactByUserID = async (userID, user) => {
    try{
        await protectedAxios.post(`${CONTACT}${USERS}/${userID}`, user)
    }catch(error){
        throw new Error(error.message)
    }
}

const updateContactByUserID = async (userID, user) => {
    try{
        await protectedAxios.patch(`${CONTACT}${USERS}/${userID}`, user)
    }catch(error){
        throw new Error(error.message)
    }
}