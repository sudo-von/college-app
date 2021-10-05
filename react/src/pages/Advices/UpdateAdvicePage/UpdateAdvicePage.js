import React, { useState, useEffect } from 'react'
import { View, SafeAreaView, ScrollView } from 'react-native'
import { Bold, Small, Container } from 'src/components'
import { Title } from 'react-native-paper'
import { styles } from './UpdateAdvicePage.styles'
import { useRoute } from '@react-navigation/native'
import { getAdviceByID } from 'src/services/advice.service'
import UpdateAdviceForm from '../UpdateAdviceForm/UpdateAdviceForm'

const UpdateAdvicePage = () => {

    const [ advice, setAdvice ] = useState({})
    const [ loading, setLoading ] = useState(true)
    const { params } = useRoute()

    useEffect(() => {
        const searchAdviceByID = async (adviceID) => {
            try{
                const response = await getAdviceByID(adviceID)
                setAdvice(response)
            }catch(error){
                Alert.alert('¡Ha ocurrido un error!', error.message)
            }finally{
                setLoading(false)
            }
        }
        searchAdviceByID(params.id)
    },[])

    return(
        <SafeAreaView>
            <ScrollView>
                <Container justifyContent='flex-start'>
                    <View style={styles.view}>
                        <Title><Bold>Actualiza tu asesoría{'\n'}en cualquier momento!</Bold></Title>
                        <Small>Manten la información de tu asesoría actualizada para que todos puedan acudir a ella.</Small>
                    </View>
                    <UpdateAdviceForm
                        advice={advice}
                        isLoading={loading}
                    />
                </Container>
            </ScrollView>
        </SafeAreaView>
    )
}

export default UpdateAdvicePage
