import React from 'react'
/* Custom components. */
import Container from 'src/components/Container'
import Logo from 'src/components/Logo'
import Header from 'src/components/Header'
import Link from 'src/components/Link'
import Small from 'src/components/Small'
import Center from 'src/components/Center'
import LoginForm from './Components/LoginForm'

const Login = () => {
    const logo = require('../../../assets/gopher.png')
    return (
        <Container>
            <Center>
                <Logo 
                    source={logo}
                    size={160}
                />
            </Center>
            <Header 
                title='Universidad Autónoma de Gojs'
                subtitle='"Siempre parece imposible, hasta que se hace".'
            />
            <LoginForm/>
            <Center>
                <Small>¿No tienes una cuenta?</Small>
                <Link url='/signup'>¡Regístrate aquí!</Link>
            </Center>
        </Container>
    )
}

export default Login
