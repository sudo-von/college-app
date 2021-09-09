import React from 'react'
import { StyleSheet, View } from 'react-native'
/* Custom components. */
import Container from 'src/components/Container'
import Small from 'src/components/Small'
import Bold from 'src/components/Bold'
import ContactConfigurationForm from './Components/ContactConfigurationForm'
/* React native paper. */
import { Title } from 'react-native-paper'

const ContactConfiguration = () =>
    <Container style={styles.container}>
        <View style={styles.view}>
            <Title><Bold>¡Es importante mantener informados a tus seres queridos!</Bold></Title>
            <Small>Modifica tu información de contacto para alertar a tus seres queridos en cualquier momento.</Small>
        </View>
        <ContactConfigurationForm/>
    </Container>

const styles = StyleSheet.create({
    container: {
        justifyContent: 'flex-start'
    },
    view: {
        marginVertical: 40,
    }
})

export default ContactConfiguration
