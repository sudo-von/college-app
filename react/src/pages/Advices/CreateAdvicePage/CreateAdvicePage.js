import React from 'react'
import { View } from 'react-native'
import { Bold, Small, Container } from 'src/components'
import { Title } from 'react-native-paper'
import { styles } from './CreateAdvicePage.styles'
import CreateAdviceForm from '../CreateAdvicePageForm/CreateAdvicePageForm'

const CreateAdvicePage = () =>
    <Container justifyContent='flex-start'>
        <View style={styles.view}>
            <Title><Bold>Apreciamos que quieras ayudar a tus compañeros!</Bold></Title>
            <Small>Al brindar una asesoría podrás apoyar a tus compañeros logrando que todos crezcan juntos como profesionistas.</Small>
        </View>
        <CreateAdviceForm/>
    </Container>

export default CreateAdvicePage
