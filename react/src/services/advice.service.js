/* Helpers. */
import { protectedAxios } from 'src/helpers/axios-helper'
/* Constants. */
import { ADVICE } from 'src/constants/endpoints'

export {
    getAdviceByID,
    getAdvices,
    createAdvice,
    deleteAdviceByID,
    updateAdviceByID,
    updateStudentsNumber
}

const getAdviceByID = async (adviceID) => {
    try{
        const advice = await protectedAxios.get(`${ADVICE}/${adviceID}`)
        return advice.data
    }catch(error){
        throw error
    }
}


const getAdvices = async () => {
    try{
        const advices = await protectedAxios.get(`${ADVICE}`)
        return advices.data.results
    }catch(error){
        throw error
    }
}

const createAdvice = async (advicePayload) => {
    try{
        await protectedAxios.post(`${ADVICE}`, advicePayload)
    }catch(error){
        throw error
    }
}


const deleteAdviceByID = async (adviceID) => {
    try{
        await protectedAxios.delete(`${ADVICE}/${adviceID}`)
    }catch(error){
        throw error
    }
}

const updateAdviceByID = async (adviceID, advice) => {
    try{
        await protectedAxios.patch(`${ADVICE}/${adviceID}`, advice)
    }catch(error){
        throw error
    }
}

const updateStudentsNumber = async (adviceID) => {
    try{
        await protectedAxios.patch(`${ADVICE}/${adviceID}/students-number`, {})
    }catch(error){
        throw error
    }
}
