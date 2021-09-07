import React from 'react'
import { SafeAreaView, ScrollView, StyleSheet } from 'react-native'
/* Custom components. */
import Container from 'src/components/Container'
import Small from 'src/components/Small'
import Center from 'src/components/Center'
import Bold from 'src/components/Bold'
import Link from 'src/components/Link'
import SignupForm from './Components/SignupForm'
/* React native paper. */
import { Title } from 'react-native-paper'

const Signup = () =>
    <SafeAreaView>
        <ScrollView>
            <Container>
                <Center style={styles.center}>
                    <Title><Bold>Crea una nueva cuenta</Bold></Title>
                    <Small style={{textAlign: 'center'}}>¡Únete a la comunidad estudiantil más increíble!</Small>
                </Center>
                <SignupForm/>
                <Center>
                    <Small>¿Ya tienes una cuenta?</Small>
                    <Link url='/login'>¡Inicia sesión aquí!</Link>
                </Center>
            </Container>
        </ScrollView>
    </SafeAreaView>

const styles = StyleSheet.create({
    center: {
        marginVertical: 30
    }
})

export default Signup