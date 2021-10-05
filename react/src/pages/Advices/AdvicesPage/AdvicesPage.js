import React, { useState, useEffect } from 'react'
import { View, Alert } from 'react-native'
import { Loader, Container } from 'src/components'
import { IconButton, useTheme } from 'react-native-paper'
import { Tabs, TabScreen } from 'react-native-paper-tabs'
import { useAuth } from 'src/providers/auth.provider'
import { getAdvices } from 'src/services/advice.service'
import { useNavigation } from '@react-navigation/native'
import { styles } from './AdvicesPage.styles'
import AdvicesList from '../AdvicesList/AdvicesList'

const AdvicesPage = () => {

    const navigation = useNavigation()
    const { colors } = useTheme()
    const { authState } = useAuth()
    const { user } = authState
    const { user_id } = user

    const [ loading, setLoading ] = useState(false)
    const [ advices, setAdvices ] = useState([])
    const attendedAdvices = advices.filter(advice => advice.students_will_attend.some(student_id => student_id == user_id))
    const teachedAdvices = advices.filter(advice => advice.user.id == user_id)

    useEffect(() => {
        const searchAdvices = async () => {
            try{
                setLoading(true)
                const response = await getAdvices()
                setAdvices(response)
            }catch(error){
                Alert.alert('¡Ha ocurrido un error!', error.message)
            }finally{
                setLoading(false)
            }
        }
        const willFocusSubscription = navigation.addListener('focus', () => {
            searchAdvices()
        })
        return willFocusSubscription
    }, [])
    
    return (
        <Container style={styles(colors).container}>
            <View style={styles(colors).view}>
                { loading ? 
                    <Loader loadingMessage='Cargando asesorías'/> :
                    <Tabs>
                        <TabScreen label='Todas'>
                            <AdvicesList
                                advices={advices}
                                userID={user_id}
                                setAdvices={setAdvices}    
                            />
                        </TabScreen>
                        <TabScreen label='Por asistir'>
                            <AdvicesList
                                advices={attendedAdvices}
                                userID={user_id}
                                setAdvices={setAdvices}    
                            />
                        </TabScreen>
                        <TabScreen label='Por dar'>
                            <AdvicesList 
                                advices={teachedAdvices}
                                userID={user_id}
                                setAdvices={setAdvices}    
                            />
                        </TabScreen>
                    </Tabs>
                }
            </View>
            <View style={styles(colors).bottomView}>
                <IconButton
                    icon='plus'
                    color='white'
                    onPress={() => navigation.navigate('/create-advice')}
                />
            </View>
        </Container>
    )
}

export default AdvicesPage