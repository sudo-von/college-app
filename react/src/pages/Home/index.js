import React from 'react'
/* Custom components. */
import Container from 'src/components/Container'
/* React native paper. */
import { Title } from 'react-native-paper'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'

const Home = () => {
    const { authState } = useAuth()
    console.log(authState.isLoggedIn)
    return (
        <Container>
            <Title>¡Hola</Title>
        </Container>
    )
}

export default Home
