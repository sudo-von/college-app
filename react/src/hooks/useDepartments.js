import { useState, useEffect } from 'react'
import { getDepartments } from 'src/services/department.service'
import { useNavigation } from '@react-navigation/native'

export const useDepartments = () => {

    const [loading, setLoading] = useState(false)
    const [departments, setDepartments] = useState([])
    const navigation = useNavigation()
  
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
