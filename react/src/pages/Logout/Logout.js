import React, { useEffect } from 'react'
/* Helpers. */
import { deleteToken } from 'src/helpers/auth-helper'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'
/* React native paper. */
import { ActivityIndicator } from 'react-native-paper'
/* Custom components. */
import Container from 'src/components/Container'
import Center from 'src/components/Center'
import Small from 'src/components/Small'

const Logout = () => {

    /* Deletes token from the storage and deletes user from the context. */
    const { authDispatch } = useAuth()
    useEffect(() => {
        const doLogout = async () => {
            try{
                await deleteToken()
                authDispatch({ type: 'logout'})
            }catch(error){
            }
        }
        doLogout()
    }, [])
    
    return (
        <Container>
            <ActivityIndicator
                size={80}
            />
            <Center>
                <Small>Cerrando sesión...</Small>
            </Center>
        </Container>
    )
}

export default Logout
