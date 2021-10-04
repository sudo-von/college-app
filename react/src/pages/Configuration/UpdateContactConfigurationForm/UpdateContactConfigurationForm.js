import React, { useState } from 'react'
import { View, Alert } from 'react-native'
import { Formik } from 'formik'
import { Input, Button } from 'src/components'
import { updateContactByID } from 'src/services/contact.service'
import { styles } from './UpdateContactConfigurationForm.styles'
import withLoading from 'src/hocs/withLoading'

const UpdateContactConfigurationForm = ({ contact, setContact }) => {

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
            const response = await updateContactByID(contact.id, data)
            setContact({...contact, ...data})
            Alert.alert(response, 'El usuario ha sido actualizado con éxito.')
        }catch(error){
            Alert.alert('¡Ha ocurrido un error al actualizar tu contacto!', error.message)
        }finally{
            setLoading(false)
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
                        loadingMessage='Actualizando el contacto...' 
                        style={styles.button}
                    >
                        Actualizar contacto
                    </Button>
                </View>
            )}
        </Formik>
    )
}

export default withLoading(UpdateContactConfigurationForm, 'Cargando contacto...')
