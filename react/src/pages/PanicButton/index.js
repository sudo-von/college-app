import React, { useState, useEffect } from 'react'
import { Linking, Alert } from 'react-native'
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
/* Geolocation. */
import * as Location from 'expo-location'

const PanicButton = () => {

    const { authState } = useAuth()
    const { user_id } = authState.user
    const navigation = useNavigation()

    useEffect(() => {
        const sendMessage = async (userID) => {
            try{
                /* Gets contact info. */
                const { contact_number, message } = await getContactByUserID(userID)
                /* Gets the formatted google maps url. */
                const { coords } = await getLocation()
                const { latitude, longitude } = coords
                const google_maps_url = `https://www.google.com/maps?q=${latitude},${longitude}`
                const whapp_web_url = `https://wa.me/${contact_number}?text=${encodeURI(message)}, esta es mi ubicación: ${google_maps_url}`
                /* Tries to send the message. */
                const supported = await Linking.canOpenURL(whapp_web_url)
                if (supported) {
                    await Linking.openURL(whapp_web_url)
                    navigation.navigate('/')
                } else {
                    throw new Error('No ha sido posible acceder a la URL.')
                }
            }catch(error){
                Alert.alert('¡Ha ocurrido un error!', error.message)
                navigation.navigate('/')
            }
        }
        sendMessage(user_id)
    },[])


    const getLocation = async () => {
        try{
            let { status } = await Location.requestForegroundPermissionsAsync()
            if (status !== 'granted') {
                throw new Error('Considera habilitar el permiso de ubicación, sin este permiso el botón de pánico no podrá mandar tu ubicación.')
            }
            let location = await Location.getCurrentPositionAsync({ accuracy: 1 })
            return location
        }catch(error){
            throw new Error(error.message)
        }
    }


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
