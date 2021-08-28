import React from 'react'
import { StyleSheet, View } from 'react-native'
/* Custom components. */
import Logo from 'src/components/Logo'
import Header from 'src/components/Header'
import LoginForm from './Components/LoginForm'
/* React native paper. */
import { Caption } from 'react-native-paper'

const Login = () => {

    const gopher = require('../../../assets/gopher.png')

    return (
        <View style={styles.container}>
            <Logo source={gopher}/>
            <Header 
                title='Universidad Autónoma de Gojs'
                subtitle='"Siempre parece imposible, hasta que se hace".'
            />
            <LoginForm/>
            <View style={styles.navigation}>
                <Caption style={styles.message}>¿No tienes una cuenta?</Caption>
                <Caption style={styles.signup}>¡Regístrate aquí!</Caption>
            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        padding: 24,
        flex: 1,
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center'
    },
    navigation: {
        marginTop: 35,
    },
    message: {
        color: '#666666',
        fontSize: 12,
        textAlign: 'center'
    },
    signup: {
        fontSize: 14,
        color: '#4C9DAF',
        textAlign: 'center',
        lineHeight: 13
    }
})

export default Login
