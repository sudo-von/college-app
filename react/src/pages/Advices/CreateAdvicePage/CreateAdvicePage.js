import React from 'react'
import { View, SafeAreaView, ScrollView } from 'react-native'
import { Bold, Small, Container } from 'src/components'
import { Title } from 'react-native-paper'
import { styles } from './CreateAdvicePage.styles'
import CreateAdviceForm from '../CreateAdviceForm/CreateAdviceForm'

const CreateAdvicePage = () =>
    <SafeAreaView>
        <ScrollView>
            <Container justifyContent='flex-start'>
                <View style={styles.view}>
                    <Title><Bold>Apreciamos que quieras ayudar a tus compañeros!</Bold></Title>
                    <Small>Al brindar una asesoría podrás apoyar a tus compañeros logrando que todos crezcan juntos como profesionistas.</Small>
                </View>
                <CreateAdviceForm/>
            </Container>
        </ScrollView>
    </SafeAreaView>

export default CreateAdvicePage
