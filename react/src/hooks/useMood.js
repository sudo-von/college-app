import { useState, useEffect } from 'react'
import { Alert } from 'react-native'
import { getMood, sendMood } from 'src/services/mood.service'

export const useMood = (initialMoodValue, userID) => {

    const [ mood, setMood ] = useState(initialMoodValue)
    const [ show, setShow ] = useState(false)

    const handleMood = async (userID) => {
        try{
            await getMood(userID)
        }catch(error){
            setShow(true)
        }
    }

    const handleUpdateMood = async () => {
        try{
            await sendMood(userID, { mood })
        }catch(error){
            Alert.alert('¡Ha ocurrido un error!', error.message)
        }finally{
            setShow(false)
        }
    }

    const handleShow = (show) => setShow(show)

    useEffect(() => {
        handleMood(userID)
    }, [])

    return {
        mood,
        show,
        setMood,
        handleShow,
        handleUpdateMood
    }
} 