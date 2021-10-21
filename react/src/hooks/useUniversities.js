import { useState, useEffect } from 'react'
import { Alert } from 'react-native'
import { getUniversities } from 'src/services/university.service'

export const useUniversities = () => {

    const [ loading, setLoading] = useState(false)
    const [ universities, setUniversities ] = useState([])

    const handleUniversities = async () => {
        try{
            setLoading(true)
            const response = await getUniversities()
            setUniversities(response.results)
        }catch(error){
            Alert.alert('¡Ha ocurrido un error!', error.message)
        }finally{
            setLoading(false)
        }
    }

    useEffect(() => {
        handleUniversities()
    }, [])

    return {
        universities,
        loading
    }
} 

