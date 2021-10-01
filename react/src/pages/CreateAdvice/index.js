import React from 'react'
import { View, StyleSheet } from 'react-native'
/* Custom components. */
import CreateAdviceForm from './Components/CreateAdviceForm'
import Bold from 'src/components/Bold'
import Small from 'src/components/Small'
import Container from 'src/components/Container'
/* React native paper. */
import { Title } from 'react-native-paper'

const CreateAdvice = () => {
    return (
        <Container justifyContent='flex-start'>
            <View style={styles.view}>
                <Title><Bold>Apreciamos que quieras ayudar a tus compañeros!</Bold></Title>
                <Small>Al brindar una asesoría podrás apoyar a tus compañeros logrando que todos crezcan juntos como profesionistas.</Small>
            </View>
            <CreateAdviceForm/>
        </Container>
    )
}

const styles = StyleSheet.create({
    view: {
        marginVertical: 40
    }
})

export default CreateAdvice
