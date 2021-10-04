/* Helpers. */
import { protectedAxios } from 'src/helpers/axios-helper'
/* Constants. */
import { CONTACT, USERS } from 'src/constants/endpoints'

export {
    getContactByUserID,
    createContactByUserID,
    updateContactByID
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

const createContactByUserID = async (userID, contact) => {
    try{
        await protectedAxios.post(`${CONTACT}${USERS}/${userID}`, contact)
    }catch(error){
        throw new Error(error.message)
    }
}

const updateContactByID = async (contactID, contact) => {
    try{
        await protectedAxios.patch(`${CONTACT}/${contactID}`, contact)
    }catch(error){
        throw new Error(error.message)
    }
}