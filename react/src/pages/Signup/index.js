import React from 'react'
import { SafeAreaView, ScrollView, StyleSheet } from 'react-native'
/* Custom components. */
import Container from 'src/components/Container'
import Logo from 'src/components/Logo'
import Small from 'src/components/Small'
import Center from 'src/components/Center'
import Bold from 'src/components/Bold'
import Link from 'src/components/Link'
import SignupForm from './Components/SignupForm'
/* React native paper. */
import { Title } from 'react-native-paper'

const Signup = () => {
    const logo = require('../../../assets/logo.png')
    return(
        <SafeAreaView>
            <ScrollView>
                <Container>
                    <Center style={styles.center}>
                        <Logo size={60} source={logo}/>
                        <Title><Bold>Bienvenido</Bold></Title>
                        <Small>Crea una cuenta completando el formulario.</Small>
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

const styles = StyleSheet.create({
    center: {
        marginVertical: 15
    }
})

export default Signup