import { post } from 'src/helpers/protected-axios-helper'

const sendSuggestion = async (suggestion) => {
    try{
        await post('/suggestions', suggestion)
        return 'Tu sugerencia se ha registrado con éxito'
    }catch(error){
        throw error
    }
}

export { sendSuggestion }