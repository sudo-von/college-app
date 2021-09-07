import React from 'react'
import { StyleSheet, View } from 'react-native'
/* React native paper. */
import { Title } from 'react-native-paper'
/* Custom components. */
import Container from 'src/components/Container'
import SuggestionForm from './Components/SuggestionForm'
import Bold from 'src/components/Bold'
import Small from 'src/components/Small'

const Suggestions = () =>
    <Container style={styles.container}>
        <View style={styles.view}>
            <Title><Bold>Tu opinión es muy{'\n'}importante para nosotros!</Bold></Title>
            <Small>¡Realiza una sugerencia para mejorar la aplicación!</Small>
        </View>
        <SuggestionForm/>
    </Container>

const styles = StyleSheet.create({
    container: {
        justifyContent: 'flex-start'
    },
    view: {
        marginVertical: 40
    }
})


export default Suggestions
