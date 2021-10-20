import { useCallback, useState } from 'react'
import { Alert } from 'react-native'
import { sendSuggestion } from 'src/services/suggestion.service'

export const useSuggestion = () => {

    const [ loading, setLoading] = useState(false)

    const handleSuggestion = useCallback(async (form, { resetForm }) => {
        try{
            setLoading(true)
            const response = await sendSuggestion(form)
            Alert.alert('¡Felicidades!', response)
            resetForm()
        }catch(error){
            Alert.alert('¡Ha ocurrido un error!', error.message)
        }finally{
            setLoading(false)
        }
    }, [])

    return {
        loading,
        handleSuggestion
    }
} 