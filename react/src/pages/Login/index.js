import React from 'react'
import { StyleSheet, View } from 'react-native'
/* Custom components. */
import Logo from 'src/components/Logo'
import Header from 'src/components/Header'
import Link from 'src/components/Link'
import Small from 'src/components/Small'
import LoginForm from './Components/LoginForm'

const Login = () => {
    const logo = require('../../../assets/gopher.png')
    return (
        <View style={styles.container}>
            <Logo source={logo}/>
            <Header 
                title='Universidad Autónoma de Gojs'
                subtitle='"Siempre parece imposible, hasta que se hace".'
            />
            <LoginForm/>
            <Small>¿No tienes una cuenta?</Small>
            <Link url='/signup'>¡Regístrate aquí!</Link>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        padding: 25,
        flex: 1,
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center'
    }
})

export default Login
