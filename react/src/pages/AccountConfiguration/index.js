import React, { useState, useEffect } from 'react'
import { StyleSheet, Alert,  View } from 'react-native'
/* Custom components. */
import Container from 'src/components/Container'
import Small from 'src/components/Small'
import Bold from 'src/components/Bold'
import AccountConfigurationForm from './Components/AccountConfigurationForm'
/* React native paper. */
import { Title } from 'react-native-paper'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'
/* Services. */
import { getUserByID } from 'src/services/user.service'

const AccountConfiguration = () => {

    /* Destructuring user properties. */
    const { authState } = useAuth()
    const { user: { user_id } } = authState
    const [ user, setUser ] = useState({ name: '', birth_date: '', email: '', registration_number: ''})
    const [ loading, setLoading ] = useState(true)

    /* Fetchs user data. */
    useEffect(() => {
        const searchUser = async (userID) => {
            try{
                const response = await getUserByID(userID)
                setUser(response)
            }catch(error){
                Alert.alert('¡Ha ocurrido un error!', error.message)
            }finally{
                setLoading(false)
            }
        }
        searchUser(user_id)
    },[])

    return (
        <Container justifyContent='flex-start'>
            <View style={styles.view}>
                <Title><Bold>¡Manten actualizados{'\n'}tus datos en todo momento!</Bold></Title>
                <Small>Modifica tu información personal para estar al día.</Small>
            </View>
            <AccountConfigurationForm 
                user={user} 
                isLoading={loading}
            />
        </Container>
    )
}

const styles = StyleSheet.create({
    view: {
        marginVertical: 40,
    }
})

export default AccountConfiguration
