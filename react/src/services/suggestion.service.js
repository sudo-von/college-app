/* Helpers. */
import { protectedAxios } from 'src/helpers/axios-helper'

export {
    sendSuggestion,
}

const sendSuggestion = async (suggestion) => {
    try{
        await protectedAxios.post(SUGGESTION, suggestion)
    }catch(error){
        throw new Error(error.message)
    }
}