import React from 'react'
import { StyleSheet, View } from 'react-native'
/* Custom components. */
import Container from 'src/components/Container'
import Small from 'src/components/Small'
import Bold from 'src/components/Bold'
import AccountConfigurationForm from './Components/AccountConfigurationForm'
/* React native paper. */
import { Title } from 'react-native-paper'

const AccountConfiguration = () =>
    <Container style={styles.container}>
        <View style={styles.view}>
            <Title><Bold>¡Manten actualizados{'\n'}tus datos en todo momento!</Bold></Title>
            <Small>Modifica tu información personal para estar al día.</Small>
        </View>
        <AccountConfigurationForm/>
    </Container>

const styles = StyleSheet.create({
    container: {
        justifyContent: 'flex-start'
    },
    view: {
        marginVertical: 40,
    }
})

export default AccountConfiguration
