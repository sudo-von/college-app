import React from 'react'
import { SafeAreaView, ScrollView } from 'react-native'
/* Custom components. */
import Container from 'src/components/Container'
import Small from 'src/components/Small'
import Center from 'src/components/Center'
import Bold from 'src/components/Bold'
import Link from 'src/components/Link'
import SignupForm from './Components/SignupForm'
/* React native paper. */
import { Title } from 'react-native-paper'

const Signup = () => {
    return (
        <SafeAreaView>
            <ScrollView>
                <Container>
                    <Center>
                        <Title><Bold>Registro</Bold></Title>
                        <Small>Cuéntanos sobre ti</Small>
                    </Center>
                    <SignupForm/>
                    <Center>
                        <Small>¿Ya tienes una cuenta?</Small>
                        <Link url='/login'>¡Inicia sesión aquí!</Link>
                    </Center>
                </Container>
            </ScrollView>
        </SafeAreaView>
    )
}

export default Signup