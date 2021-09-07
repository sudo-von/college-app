/* Helpers. */
import axios from 'src/helpers/axios-helper'
/* Constants. */
import { SUGGESTION } from 'src/constants/endpoints'

export {
    sendSuggestion,
}

const sendSuggestion = async (suggestion) => {
    try{
        await axios.post(SUGGESTION, suggestion)
    }catch(error){
        throw new Error(error.response.data.message)
    }
}