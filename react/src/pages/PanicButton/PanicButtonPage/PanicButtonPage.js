import React from 'react'
import { ActivityIndicator } from 'react-native-paper'
import { Container, Center, Small } from 'src/components'
import { useUser } from 'src/hooks/useUser'
import { useLocation } from 'src/hooks/useLocation'

const PanicButtonPage = () => {
    
    const { user } = useUser()
    const { user_id } = user
    const { message } = useLocation(user_id)

    return (
        <Container>
            <ActivityIndicator size={80}/>
            <Center>
                <Small>{message}</Small>
            </Center>
        </Container>
    )
}

export default PanicButtonPage
