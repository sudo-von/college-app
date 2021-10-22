import { useState, useCallback, useEffect } from 'react'
import { Alert } from 'react-native'
import { getDepartmentByID } from 'src/services/department.service'

export const useSingleDepartment = (id) => {

    const [ loading, setLoading ] = useState(false)
    const [ department, setDepartment ] = useState({})

    const handleDepartmentByID = useCallback(async (departmentID) => {
        try{
            setLoading(true)
            const department = await getDepartmentByID(departmentID)
            setDepartment(department)
        }catch(error){
            Alert.alert('¡Ha ocurrido un error!', error.message)
        }finally{
            setLoading(false)
        }
    }, [])

    useEffect(() => {
        handleDepartmentByID(id)
    }, [])

    return {
        loading,
        department,
        setDepartment
    }
} 