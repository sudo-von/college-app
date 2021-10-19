import React from 'react'
import { View, SafeAreaView, ScrollView } from 'react-native'
import { Bold, Small, Container } from 'src/components'
import { Title } from 'react-native-paper'
import { styles } from './CreateDepartmentPage.styles'
import CreateDepartmentForm from '../CreateDepartmentForm/CreateDepartmentForm'

const CreateDepartmentPage = () =>
    <SafeAreaView>
        <ScrollView>
            <Container justifyContent='flex-start'>
                <View style={styles.view}>
                    <Title><Bold>Anuncia un departamento!</Bold></Title>
                    <Small>Podrás brindarle una oportunidada tus compañeros para encontrar un lugar que les permita continuar con sus estudios.</Small>
                </View>
                <CreateDepartmentForm/>
            </Container>
        </ScrollView>
    </SafeAreaView>

export default CreateDepartmentPage
