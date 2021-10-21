import { useState, useEffect } from 'react'
import { Alert } from 'react-native'
import { getDepartments } from 'src/services/department.service'
import { useNavigation } from '@react-navigation/native'

export const useDepartments = () => {
    
    const navigation = useNavigation()
    
    const [ loading, setLoading ] = useState(false)
    const [ departments, setDepartments ] = useState([])
  
    const handleDepartments = async () => {
        try{
            setLoading(true)
            const response = await getDepartments()
            setDepartments(response.results)
        }catch(error){
            Alert.alert('¡Ha ocurrido un error!', error.message)
        }finally{
            setLoading(false)
        }
    }
    
    useEffect(() => {
        handleDepartments()
        const willFocusSubscription = navigation.addListener('focus', () => {
            handleDepartments()
        })
        return willFocusSubscription
    }, [])
  
    return {loading, departments, setDepartments }
} 
