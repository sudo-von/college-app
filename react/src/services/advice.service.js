/* Helpers. */
import { protectedAxios } from 'src/helpers/axios-helper'
/* Constants. */
import { ADVICE } from 'src/constants/endpoints'

export {
    getAdvices,
    deleteAdviceByID,
    updateStudentsNumber
}

const getAdvices = async () => {
    try{
        const advices = await protectedAxios.get(`${ADVICE}`)
        return advices.data.results
    }catch(error){
        throw new Error(error.message)
    }
}

const deleteAdviceByID = async (adviceID) => {
    try{
        await protectedAxios.delete(`${ADVICE}/${adviceID}`)
    }catch(error){
        throw new Error(error.message)
    }
}

const updateStudentsNumber = async (adviceID) => {
    try{
        await protectedAxios.patch(`${ADVICE}/${adviceID}/students-number`, {})
    }catch(error){
        throw new Error(error.message)
    }
}
