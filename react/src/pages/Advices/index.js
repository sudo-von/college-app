import React, { useState, useEffect } from 'react'
import { View, StyleSheet } from 'react-native'
/* Custom components. */
import Loader from 'src/components/Loader'
import Container from 'src/components/Container'
import AdvicesList from './Compontents/AdvicesList'
/* React native paper. */
import { IconButton, useTheme } from 'react-native-paper'
import { Tabs, TabScreen } from 'react-native-paper-tabs'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'
/* Services. */
import { getAdvices } from 'src/services/advice.service'
/* React navigation. */
import { useNavigation } from '@react-navigation/native'

const Advices = () => {

    const navigation = useNavigation()
    const { colors } = useTheme()
    /* Gets user id. */
    const { authState } = useAuth()
    const { user } = authState
    const { user_id } = user
    /* Sets the advice list that will use each component. */
    const [ advices, setAdvices ] = useState([])
    const [ loading, setLoading ] = useState(false)
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
        searchAdvices()
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

const styles = (colors) => StyleSheet.create({
    container: {
        padding: 0
    },
    view: {
        backgroundColor: '#F2F2F2',
        flex: 1,
    },
    bottomView: {
        alignItems: 'center',
        backgroundColor: colors.primary
    }
})

export default Advices
