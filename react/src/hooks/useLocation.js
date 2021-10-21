import { useState, useEffect } from 'react'
import { Linking, Alert } from 'react-native'
import { getContactByUserID } from 'src/services/contact.service'
import { checkPermissions, getCoords } from 'src/services/location.service'
import { useNavigation } from '@react-navigation/native'

export const useLocation = (userID) => {

    const navigation = useNavigation()
    
    const [ message, setMessage ] = useState('')

    const handleMessage = async(userID) => {
        try{           
            setMessage('Revisando permisos...')
            await checkPermissions()
            setMessage('Obteniendo contacto...')
            const { contact_name, contact_number, message } = await getContactByUserID(userID)
            setMessage('Obteniendo localización...')
            const { latitude, longitude } = await getCoords()
            setMessage('Abriendo whatsapp...')
            const formated_message = encodeURI(`¡Hola ${contact_name}! ${message}`)
            const google_maps_url = `https://www.google.com/maps?q=${latitude},${longitude}`
            const whapp_web_url = `https://wa.me/${contact_number}?text=${formated_message}, esta es mi ubicación: ${google_maps_url}`
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

    useEffect(() => {
        handleMessage(userID)
    },  [])

    return {
        message
    }
} 

