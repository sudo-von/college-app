import { useCallback, useState } from 'react'
import { Alert } from 'react-native'
import { useAuth } from 'src/providers/auth.provider'
import { login } from 'src/services/auth.service'
import { deleteToken } from 'src/services/token.service'

export const useUser = () => {

    const { authState: { user }, authDispatch } = useAuth()
    const [ loading, setLoading] = useState(false)

    const handleLogin = useCallback(async ({ email, password }) => {
        try {
            setLoading(true)
            const response = await login({ email, password })
            setLoading(false)
            authDispatch({ type: 'login', user: response })
        }catch(error){
            setLoading(false)
            Alert.alert('Ha ocurrido un error', error.message)
        }
    }, [])


    const handleLogout = useCallback(async () => {
        try{
            await deleteToken()
            authDispatch({ type: 'logout' })
        }catch(error){
            Alert.alert('¡Ha ocurrido un error!', error.message)
        }
    }, [])

    return {
        user,
        loading,
        handleLogin,
        handleLogout
    }
} 

