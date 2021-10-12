import { get, post } from 'src/helpers/protected-axios-helper'

const getMood = async (userID) => {
    try{
        const mood = await get(`/users-mood/users/${userID}`)
        return mood.data
    }catch(error){
        throw error
    }
}

const sendMood = async (userID, mood) => {
    try{
        await post(`/users-mood/users/${userID}`, mood)
    }catch(error){
        throw error
    }
}

export {
    getMood,
    sendMood
}