/* Helpers. */
import { protectedAxios } from 'src/helpers/axios-helper'
/* Constants. */
import { CONTACT, USERS } from 'src/constants/endpoints'
import { CONTACT_ERRORS } from 'src/constants/errors'

export {
    getContactByUserID,
    updateContactByUserID
}

const getContactByUserID = async (userID) => {
    try{
        const response = await protectedAxios.get(`${CONTACT}${USERS}/${userID}`)
        const user = await response.data
        console.log(user)
        return user
    }catch(error){
        throw new Error(CONTACT_ERRORS[error.response.data.code]['esp'])
    }
}

const sendContactByUserID = async (userID, user) => {
    try{
        await protectedAxios.post(`${CONTACT}${USERS}/${userID}`, user)
    }catch(error){
        console.log(error)
        if(error.response?.data.message){
            throw new Error(error.response.data.message)
        }
        throw new Error(CONTACT_ERRORS[error.response.data.code]['esp'])
    }
}

const updateContactByUserID = async (userID, user) => {
    try{
        await protectedAxios.patch(`${CONTACT}${USERS}/${userID}`, user)
    }catch(error){
        console.log(error)
        if(error.response?.data.message){
            throw new Error(error.response.data.message)
        }
        throw new Error(CONTACT_ERRORS[error.response.data.code]['esp'])
    }
}