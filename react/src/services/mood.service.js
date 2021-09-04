/* Helpers. */
import axios from 'src/helpers/axios-helper'
/* Constants. */
import { MOOD } from 'src/constants/endpoints'

export {
    getMood
}

/* Handles basic auth. */
const getMood = async (userID) => {
    try{
        const mood = await axios.get(`${MOOD}/users/${userID}`)
        return mood
    }catch(error){
        throw new Error(error.response.data.message)
    }
}