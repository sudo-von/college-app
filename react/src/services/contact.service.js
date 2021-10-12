import { get, post, patch } from 'src/helpers/protected-axios-helper'

const getContactByUserID = async (userID) => {
    try{
        const response = await get(`/contacts/users/${userID}`)
        const user = await response.data
        return user
    }catch(error){
        throw error
    }
}

const createContactByUserID = async (userID, contact) => {
    try{
        await post(`/contacts/users/${userID}`, contact)
        return 'El contacto ha sido guardado con éxito.'
    }catch(error){
        throw error
    }
}

const updateContactByID = async (contactID, contact) => {
    try{
        await patch(`/contacts/${contactID}`, contact)
        return 'El contacto ha sido actualizado con éxito.'
    }catch(error){
        throw error
    }
}

export {
    getContactByUserID,
    createContactByUserID,
    updateContactByID
}