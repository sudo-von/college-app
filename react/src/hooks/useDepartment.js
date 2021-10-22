import { useState, useCallback } from 'react'
import { Alert } from 'react-native'
import { createDepartment, updateDepartmentByID  } from 'src/services/department.service'
import { useNavigation } from '@react-navigation/native'

export const useDepartment = () => {

    const navigation = useNavigation()

    const [ loading, setLoading ] = useState(false)

    const handleCreateDepartment = useCallback(async (form, resetForm ) => {
        try{
            setLoading(true)
            await createDepartment({ ...form, cost: parseFloat(form.cost) })
            setLoading(false)
            Alert.alert('¡Felicidades!', '¡Has registrado el departamento con éxito!')
            resetForm()
            navigation.goBack()
        }catch(error){
            setLoading(false)
            Alert.alert('¡Ha ocurrido un error!', error.message)
        }
    }, [])

    const handleUpdateDepartment = useCallback(async (id, form, resetForm) => {
        try{
            console.log(form)
            setLoading(true)
            await updateDepartmentByID(id, { ...form, cost: parseFloat(form.cost) })
            setLoading(false)
            Alert.alert('¡Felicidades!', '¡Has actualizado el departamento con éxito!')
            resetForm()
            navigation.goBack()
        }catch(error){
            setLoading(false)
            Alert.alert('¡Ha ocurrido un error!', error.message)
        }
    }, [])

    return {
        loading,
        handleCreateDepartment,
        handleUpdateDepartment
    }
} 