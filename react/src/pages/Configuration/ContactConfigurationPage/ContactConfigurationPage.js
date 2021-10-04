import React, { useState, useEffect } from 'react'
import { Alert, View, SafeAreaView, ScrollView } from 'react-native'
import { Container, Small, Bold } from 'src/components'
import { Title } from 'react-native-paper'
import { styles } from './ContactConfigurationPage.styles'
import { getContactByUserID } from 'src/services/contact.service'
import { useAuth } from 'src/providers/auth.provider'
import CreateContactConfigurationForm from '../CreateContactConfigurationForm/CreateContactConfigurationForm'
import UpdateContactConfigurationForm from '../UpdateContactConfigurationForm/UpdateContactConfigurationForm'

const ContactConfigurationPage = () => {

    const { authState } = useAuth()
    const { user } = authState
    const { user_id } = user

    const [ contact, setContact ] = useState({ id: '', contact_name: '', contact_number: '', message: '' })
    const [ loading, setLoading ] = useState(true)
    
    useEffect(() => {
        const searchContact = async (userID) => {
            try{
                const response = await getContactByUserID(userID)
                setContact(response)
            }catch(error){
                Alert.alert('¡No ha sido posible cargar la información de tu contacto!', error.message)
            }finally{
                setLoading(false)
            }
        }
        searchContact(user_id)
    },  [])

    return(
        <SafeAreaView>
            <ScrollView>
                <Container justifyContent='flex-start'>
                    <View style={styles.view}>
                        <Title><Bold>¡Es importante mantener informados a tus seres queridos!</Bold></Title>
                        <Small>Modifica tu información de contacto para alertar a tus seres queridos en cualquier momento.</Small>
                    </View>
                    {   Object.values(contact).some(v => v) ? 
                        <UpdateContactConfigurationForm
                            contact={contact}
                            isLoading={loading}
                            setContact={setContact}
                        /> : 
                        <CreateContactConfigurationForm
                            contact={contact}
                            isLoading={loading}
                            userID={user_id}
                            setContact={setContact}
                        />  
                    }
                </Container>
            </ScrollView>
        </SafeAreaView>
    )
}

export default ContactConfigurationPage
