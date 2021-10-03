import React, { useEffect } from 'react'
import { deleteToken } from 'src/helpers/auth-helper'
import { useAuth } from 'src/providers/auth.provider'
import { ActivityIndicator } from 'react-native-paper'
import { Container, Center, Small } from 'src/components'
import { Alert } from 'react-native'

const LogoutPage = () => {

    /* Deletes token from the storage and then deletes user from the context. */
    const { authDispatch } = useAuth()
    useEffect(() => {
        const doLogout = async () => {
            try{
                await deleteToken()
                authDispatch({ type: 'logout' })
            }catch(error){
                Alert.alert('¡Ha ocurrido un error!', error.message)
            }
        }
        doLogout()
    }, [])
    
    return (
        <Container>
            <ActivityIndicator size={80}/>
            <Center>
                <Small>Cerrando sesión...</Small>
            </Center>
        </Container>
    )
}

export default LogoutPage
