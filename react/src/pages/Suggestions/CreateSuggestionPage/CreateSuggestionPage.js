import React from 'react'
import { View } from 'react-native'
import { Title } from 'react-native-paper'
import { Container, Bold, Small } from 'src/components'
import { styles } from './CreateSuggestionPage.styles'
import SuggestionForm from '../SuggestionForm/SuggestionForm'

const CreateSuggestionPage = () =>
    <Container justifyContent='flex-start'>
        <View style={styles.view}>
            <Title><Bold>Tu opinión es muy{'\n'}importante para nosotros!</Bold></Title>
            <Small>¡Realiza una sugerencia para mejorar la aplicación! </Small>
        </View>
        <SuggestionForm/>
    </Container>

export default CreateSuggestionPage
