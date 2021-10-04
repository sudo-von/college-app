import React, { useState, useEffect } from 'react'
import { Alert, View } from 'react-native'
import { Container, Small, Bold } from 'src/components'
import { Title } from 'react-native-paper'
import { useAuth } from 'src/providers/auth.provider'
import { getUserByID } from 'src/services/user.service'
import { styles } from './UpdateAccountConfigurationPage.styles'
import UpdateAccountConfigurationForm from '../UpdateAccountConfigurationForm/UpdateAccountConfigurationForm'

const UpdateAccountConfigurationPage = () => {

    const { authState } = useAuth()
    const { user } = authState
    const { user_id } = user

    const [ account, setAccount ] = useState({ id: '', name: '', birth_date: '', email: '', registration_number: '' })
    const [ loading, setLoading ] = useState(true)

    useEffect(() => {
        const searchAccount = async (userID) => {
            try{
                const response = await getUserByID(userID)
                setAccount(response)
            }catch(error){
                Alert.alert('¡Ha ocurrido un error!', error.message)
            }finally{
                setLoading(false)
            }
        }
        searchAccount(user_id)
    },  [])

    return (
        <Container justifyContent='flex-start'>
            <View style={styles.view}>
                <Title><Bold>¡Manten actualizados{'\n'}tus datos en todo momento!</Bold></Title>
                <Small>Modifica tu información personal para estar al día.</Small>
            </View>
            <UpdateAccountConfigurationForm 
                account={account} 
                isLoading={loading}
            />
        </Container>
    )
}

export default UpdateAccountConfigurationPage
