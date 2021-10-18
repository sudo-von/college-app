import React, { useEffect } from 'react'
import { Linking, Alert } from 'react-native'
import { useAuth } from 'src/providers/auth.provider'
import { ActivityIndicator } from 'react-native-paper'
import { Container, Center, Small } from 'src/components'
import { getContactByUserID } from 'src/services/contact.service'
import { useNavigation } from '@react-navigation/native'
import * as Location from 'expo-location'

const PanicButtonPage = () => {
    
    const navigation = useNavigation()
    const { authState } = useAuth()
    const { user } = authState
    const { user_id } = user

    /* Send the custom contact's message to the user's contact. */
    useEffect(() => {
        const sendMessage = async (userID) => {
            try{
                /* Get contact info. */
                const { contact_name, contact_number, message } = await getContactByUserID(userID)
                /* Get current coords. */
                await checkPermissions()
                const { latitude, longitude } = await getCoords()
                /* Formats the message. */
                const formated_message = encodeURI(`¡Hola ${contact_name}! ${message}`)
                const google_maps_url = `https://www.google.com/maps?q=${latitude},${longitude}`
                const whapp_web_url = `https://wa.me/${contact_number}?text=${formated_message}, esta es mi ubicación: ${google_maps_url}`
                /* Sends the message. */
                const supported = await Linking.canOpenURL(whapp_web_url)
                if (supported) {
                    await Linking.openURL(whapp_web_url)
                } else {
                    throw new Error('No ha sido posible acceder a la URL.')
                }
            }catch(error){
                Alert.alert('¡Ha ocurrido un error!', error.message)
            }finally{
                navigation.navigate('/')
            }
        }
        sendMessage(user_id)
    },[])

    const checkPermissions = async () => {
        try{
            let { status } = await Location.requestForegroundPermissionsAsync()
            if (status !== 'granted') {
                throw new Error('Considera habilitar el permiso de ubicación, sin este permiso el botón de pánico no podrá acceder a tu ubicación.')
            }
        }catch(error){
            throw error
        }
    }

    const getCoords = async () => {
        try{
            let { coords } = await Location.getCurrentPositionAsync({ accuracy: 1 })
            return coords
        }catch(error){
            throw new Error('No ha sido posible acceder a tu ubicación.')
        }
    }


    return (
        <Container>
            <ActivityIndicator size={80}/>
            <Center>
                <Small>Abriendo whatsapp...</Small>
            </Center>
        </Container>
    )
}

export default PanicButtonPage
