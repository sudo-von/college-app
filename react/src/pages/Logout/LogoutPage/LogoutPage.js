import React, { useEffect } from 'react'
import { ActivityIndicator } from 'react-native-paper'
import { Container, Center, Small } from 'src/components'
import { useUser } from 'src/hooks/useUser'

const LogoutPage = () => {

    const { handleLogout } = useUser()
    
    useEffect(() => {
        handleLogout()
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