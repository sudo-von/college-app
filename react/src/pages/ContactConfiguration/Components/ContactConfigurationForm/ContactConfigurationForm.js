import React, { useEffect, useState } from 'react'
import { View, Alert, StyleSheet } from 'react-native'
/* Formik. */
import { Formik } from 'formik'
/* Custom components. */
import Input from 'src/components/Input'
import Button from 'src/components/Button'
/* Services. */
import { getContactByUserID, sendContactByUserID, updateContactByID } from 'src/services/contact.service'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'

const ContactConfigurationForm = () => {

    /* Destructuring user properties. */
    const { authState } = useAuth()
    const { user_id } = authState.user
    const [ contact, setContact ] = useState({ contact_name: '', contact_number: '', message: '' })
    const [ loading, setLoading ] = useState(false)

    /* Fetchs contact data and then reinitialize form values. */
    useEffect(() => {
        const searchContact = async (userID) => {
            try{
                const response = await getContactByUserID(userID)
                setContact(response)
            }catch(error){
                Alert.alert('¡Parece que aún no has registrado un contacto!', error.message)
            }
        }
        searchContact(user_id)
    },[])

    const updateContact = async (data) => {
        console.log(contact.id)
        try{
            setLoading(true)
            const response = await updateContactByID(contact.id, data)
            Alert.alert(response, 'El usuario ha sido actualizado con éxito.')
        }catch(error){
            Alert.alert('¡Ha ocurrido un error!', error.message)
        }finally{
            setLoading(false)
        }
    }

    const sendContact = async (data) => {
        try{
            setLoading(true)
            const response = await sendContactByUserID(user_id, data)
            Alert.alert(response, 'El contacto ha sido guardado con éxito.')
        }catch(error){
            Alert.alert('¡Ha ocurrido un error!', error.message)
        }finally{
            setLoading(false)
        }
    }

    return (
        <Formik 
            initialValues={{ 
                contact_name: contact.contact_name,
                contact_number: contact.contact_number,
                message : contact.message
            }}
            validate={values => {
                const { contact_name, contact_number, message } = values
                const errors = {}
                if(!contact_name){
                    errors.contact_name = 'Nombre del contacto requerido'
                }
                if(!contact_number){
                    errors.contact_number = 'Número de contacto requerido'
                }
                if(!message){
                    errors.message = 'Mensaje de auxilio requerido'
                }
                return errors
            }}
            onSubmit={!Object.values(contact).some(v => v) ? sendContact : updateContact }
            enableReinitialize
        >
            {({ handleChange, handleBlur, handleSubmit, errors, values }) => (
                <View>
                    <Input
                        label='Ingresa el nombre de tu contacto'
                        onChangeText={handleChange('contact_name')}
                        onBlur={handleBlur('contact_name')}
                        value={values.contact_name}
                        error={errors.contact_name}
                    />
                    <Input
                        label='Ingresa el teléfono de tu contacto'
                        onChangeText={handleChange('contact_number')}
                        onBlur={handleBlur('contact_number')}
                        value={values.contact_number}
                        error={errors.contact_number}
                        keyboardType='phone-pad'
                    />
                    <Input
                        label='Ingresa tu mensaje de auxilio'
                        onChangeText={handleChange('message')}
                        onBlur={handleBlur('message')}
                        value={values.message}
                        error={errors.message}
                    />
                    <Button 
                        loading={loading}
                        onPress={handleSubmit}
                        loadingMessage='Actualizando información...' 
                        style={styles.button}
                    >
                        Guardar cambios
                    </Button>
                </View>
            )}
        </Formik>
    )
}

const styles = StyleSheet.create({
    button: {
        marginVertical: 40
    }
})

export default ContactConfigurationForm
