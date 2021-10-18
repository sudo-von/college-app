import React from 'react'
import { Container, Logo, Header, Link, Small, Center } from 'src/components'
import LoginForm from '../LoginForm/LoginForm'

const LoginPage = () =>
    <Container>
        <Center>
            <Logo 
                source={require('assets/logo.png')}
                size={180}
            />
        </Center>
        <Header 
            title='Universidad Autónoma de Von'
            subtitle='"Siempre parece imposible, hasta que se hace".'
        />
        <LoginForm/>
        <Center>
            <Small>¿No tienes una cuenta?</Small>
            <Link url='/signup'>¡Regístrate aquí!</Link>
        </Center>
    </Container>

export default LoginPage
