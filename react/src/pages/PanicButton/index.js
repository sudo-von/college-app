import React, { useState, useEffect } from 'react'
import { Linking} from 'react-native'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'
/* React native paper. */
import { ActivityIndicator } from 'react-native-paper'
/* Custom components. */
import Container from 'src/components/Container'
import Center from 'src/components/Center'
import Small from 'src/components/Small'
/* Services. */
import { getContactByUserID } from 'src/services/contact.service'
/* React navigation. */
import { useNavigation } from '@react-navigation/native'

const PanicButton = () => {

    /* Destructuring user properties. */
    const { authState } = useAuth()
    const { user_id } = authState.user
    const navigation = useNavigation()

    /* Fetchs contact data and open whatsapp with. */
    useEffect(() => {
        const searchContact = async (userID) => {
            try{
                const response = await getContactByUserID(userID)
                const url = `https://wa.me/${response.contact_number}?text=${encodeURI(response.message)}, esta es mi ubicación:`
                const supported = await Linking.canOpenURL(url)
                if (supported) {
                    console.log(url)
                    await Linking.openURL(url)
                    navigation.navigate('/')
                } else {
                    Alert.alert('¡Ha ocurrido un error', 'Algo ha salido mal al abrir whatsapp...')
                }
            }catch(error){
                Alert.alert('¡Parece que aún no has registrado un contacto!', error.message)
            }
        }
        searchContact(user_id)
    },[])


    return (
        <Container>
            <ActivityIndicator
                size={80}
            />
            <Center>
                <Small>Abriendo whatsapp...</Small>
            </Center>
        </Container>
    )
}

export default PanicButton
