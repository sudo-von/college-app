/* Helpers. */
import { protectedAxios } from 'src/helpers/axios-helper'
/* Constants. */
import { MOOD } from 'src/constants/endpoints'

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
        throw new Error(error.message)
    }
}

/* Sends user's mood for the current day. */
const sendMood = async (mood) => {
    try{
        await protectedAxios.post(`${MOOD}`, mood)
    }catch(error){
        throw new Error(error.message)
    }
}