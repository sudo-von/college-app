import { get, post, patch, destroy } from 'src/helpers/protected-axios-helper'

const getAdviceByID = async (adviceID) => {
    try{
        const advice = await get(`/advices/${adviceID}`)
        return advice.data
    }catch(error){
        throw error
    }
}

const getAdvices = async () => {
    try{
        const advices = await get('/advices')
        return advices.data
    }catch(error){
        throw error
    }
}

const createAdvice = async (advicePayload) => {
    try{
        await post('/advices', advicePayload)
    }catch(error){
        throw error
    }
}

const deleteAdviceByID = async (adviceID) => {
    try{
        await destroy(`/advices/${adviceID}`)
        return '¡Has elimiando la asesoría con éxito!'
    }catch(error){
        throw error
    }
}

const updateAdviceByID = async (adviceID, advice) => {
    try{
        await patch(`/advices/${adviceID}`, advice)
    }catch(error){
        throw error
    }
}

const updateStudentsNumber = async (adviceID) => {
    try{
        await patch(`/advices/${adviceID}/students-number`, {})
    }catch(error){
        throw error
    }
}

export {
    getAdviceByID,
    getAdvices,
    createAdvice,
    deleteAdviceByID,
    updateAdviceByID,
    updateStudentsNumber
}