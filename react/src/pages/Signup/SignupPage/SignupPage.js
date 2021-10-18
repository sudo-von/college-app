import React from 'react'
import SignupForm from '../SignupForm/SignupForm'
import { View } from 'react-native'
import { SafeAreaView, ScrollView } from 'react-native'
import { Container, Small, Center, Bold, Link } from 'src/components'
import { Title } from 'react-native-paper'
import { styles } from './SignupPage.styles'

const SignupPage = () =>
    <SafeAreaView>
        <ScrollView>
            <Container>
                <View style={styles.view}>
                    <Title style={styles.title}><Bold>¡Bienvenido!</Bold></Title>
                    <Small>Crea tu cuenta completando el formulario.</Small>
                </View>
                <SignupForm/>
                <Center>
                    <Small>¿Ya tienes una cuenta?</Small>
                    <Link url='/login'>¡Inicia sesión aquí!</Link>
                </Center>
            </Container>
        </ScrollView>
    </SafeAreaView>

export default SignupPage