/* Helpers. */
import { protectedAxios } from 'src/helpers/axios-helper'
/* Constants. */
import { MOOD } from 'src/constants/endpoints'
import { USER_MOOD_ERRORS } from 'src/constants/errors'

export {
    getMood,
    sendMood
}

/* Gets the current day user's mood. */
const getMood = async (userID) => {
    try{
        const mood = await protectedAxios.get(`${MOOD}/users/${userID}`)
        return mood.data
    }catch(error){
        throw new Error(USER_MOOD_ERRORS[error.response.data.code]['esp'])
    }
}

/* Sends user's mood for the current day. */
const sendMood = async (mood) => {
    try{
        await protectedAxios.post(`${MOOD}`, mood)
    }catch(error){
        throw new Error(USER_MOOD_ERRORS[error.response.data.code]['esp'])
    }
}