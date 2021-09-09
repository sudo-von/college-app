/* Helpers. */
import axios from 'src/helpers/axios-helper'
import { authorizationRequestInterceptor, authorizationerrorResponseInterceptor } from 'src/helpers/interceptor-helper'
/* Constants. */
import { SUGGESTION } from 'src/constants/endpoints'

/* Add interceptors to the axios instance. */
axios.interceptors.request.use(config => authorizationRequestInterceptor(config), error => Promise.reject(error))
axios.interceptors.response.use(response => response, error => authorizationerrorResponseInterceptor(error))


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