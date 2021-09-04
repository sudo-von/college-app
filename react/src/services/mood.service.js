/* Helpers. */
import axios from 'src/helpers/axios-helper'
/* Constants. */
import { MOOD } from 'src/constants/endpoints'

export {
    getMood,
    sendMood
}

/* Gets the current day user's mood. */
const getMood = async (userID) => {
    try{
        const mood = await axios.get(`${MOOD}/users/${userID}`)
        return mood.data
    }catch(error){
        throw new Error(error.response.data.message)
    }
}

/* Sends user's mood for the current day. */
const sendMood = async (mood) => {
    try{
        await axios.post(`${MOOD}`, mood)
    }catch(error){
        throw new Error(error.response.data.message)
    }
}