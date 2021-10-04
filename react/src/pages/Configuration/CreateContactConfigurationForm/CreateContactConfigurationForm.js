import React, { useState } from 'react'
import { View, Alert } from 'react-native'
import { Formik } from 'formik'
import { Input, Button } from 'src/components'
import { getContactByUserID, createContactByUserID } from 'src/services/contact.service'
import { styles } from './CreateContactConfigurationForm.styles'
import withLoading from 'src/hocs/withLoading'

const CreateContactConfigurationForm = ({ contact, userID, setContact }) => {

    const [ loading, setLoading ] = useState(false)

    /* Formik configuration. */
    const initialValues = { 
        contact_name: contact.contact_name,
        contact_number: contact.contact_number,
        message : contact.message
    }

    const handleValidation = ({ contact_name, contact_number, message }) => {
        let errors = {}
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
    }

    const handleContact = async (data) => {
        try{
            setLoading(true)
            const createResponse = await createContactByUserID(userID, data)
            const contactResponse = await getContactByUserID(userID)
            setLoading(false)
            setContact(contactResponse)
            Alert.alert(createResponse, 'El contacto ha sido guardado con éxito.')
        }catch(error){
            setLoading(false)
            Alert.alert('¡Ha ocurrido un error al registrar tu contacto!', error.message)
        }
    }

    return (
        <Formik 
            initialValues={initialValues}
            validateOnChange={false}
            validateOnBlur={false} 
            validate={handleValidation}
            onSubmit={handleContact}
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
                        loadingMessage='Guardando el contacto...' 
                        style={styles.button}
                    >
                        Crear contacto
                    </Button>
                </View>
            )}
        </Formik>
    )
}

export default withLoading(CreateContactConfigurationForm, 'Cargando contacto...')
