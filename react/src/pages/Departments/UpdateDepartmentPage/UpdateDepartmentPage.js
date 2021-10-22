import React from 'react'
import { View, SafeAreaView, ScrollView } from 'react-native'
import { Bold, Small, Container, Center, Loader } from 'src/components'
import { Title } from 'react-native-paper'
import { styles } from './UpdateDepartmentPage.styles'
import { useSingleDepartment } from 'src/hooks/useSingleDepartment'
import { useRoute } from '@react-navigation/native'
import UpdateDepartmentForm from '../UpdateDepartmentForm/UpdateDepartmentForm'

const UpdateDepartmentPage = () => {

    const { params } = useRoute()
    const { id } = params
    const { loading, department } = useSingleDepartment(id)

    return(
        <SafeAreaView>
                <ScrollView>
                    <Container justifyContent='flex-start'>
                        <View style={styles.view}>
                            <Title><Bold>Actualiza tu departamento!</Bold></Title>
                            <Small>Modifica la información de tu departamento en cualquier momento.</Small>
                        </View>
                        { loading ? 
                            <Loader loadingMessage='Cargando departamento...'/> :
                            <UpdateDepartmentForm
                                {...department}
                            />
                        }
                    </Container>
            </ScrollView>
        </SafeAreaView>
    )
}

export default UpdateDepartmentPage
