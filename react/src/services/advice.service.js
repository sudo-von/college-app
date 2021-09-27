/* Helpers. */
import { protectedAxios } from 'src/helpers/axios-helper'
/* Constants. */
import { ADVICE } from 'src/constants/endpoints'

export {
    getAdvices
}

const getAdvices = async () => {
    try{
        const advices = await protectedAxios.get(`${ADVICE}`)
        return advices.data.results
    }catch(error){
        throw new Error(error.message)
    }
}