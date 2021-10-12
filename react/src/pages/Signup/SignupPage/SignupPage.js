import React from 'react'
import { SafeAreaView, ScrollView } from 'react-native'
import { Container, Logo, Small, Center, Bold, Link } from 'src/components'
import { Title } from 'react-native-paper'
import SignupForm from '../SignupForm/SignupForm'

const SignupPage = () =>
    <SafeAreaView>
        <ScrollView>
            <Container>
                <Center>
                    <Logo size={60} source={require('assets/logo.png')}/>
                    <Title><Bold>Bienvenido</Bold></Title>
                    <Small>Crea tu cuenta completando el formulario.</Small>
                </Center>
                <SignupForm/>
                <Center>
                    <Small>¿Ya tienes una cuenta?</Small>
                    <Link url='/login'>¡Inicia sesión aquí!</Link>
                </Center>
            </Container>
        </ScrollView>
    </SafeAreaView>

export default SignupPage