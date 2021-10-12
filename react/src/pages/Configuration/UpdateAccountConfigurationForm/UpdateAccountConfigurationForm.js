import React, { useState } from 'react'
import { View, Alert } from 'react-native'
import { Formik, Field } from 'formik'
import { Input, Datepicker, Button } from 'src/components'
import { updateUserByID } from 'src/services/user.service'
import { useAuth } from 'src/providers/auth.provider'
import { styles } from './UpdateAccountConfigurationForm.styles'
import withLoading from 'src/hocs/withLoading'

const UpdateAccountConfigurationForm = ({ account }) => {

    const { authDispatch } = useAuth()
    const [ loading, setLoading ] = useState(false)

    /* Formik configuration. */
    const initialValues = { 
        name: account.name,
        birth_date: account.birth_date,
        email : account.email, 
        registration_number: account.registration_number
    }

    const handleValidation = ({ name, birth_date, email, registration_number }) => {
        let errors = {}
        if(!name){
            errors.name = 'Nombre requerido'
        }
        if(!birth_date){
            errors.birth_date = 'Fecha de nacimiento requerida'
        }
        if(!email){
            errors.email = 'Correo requerido'
        }
        if(!registration_number){
            errors.registration_number = 'Matrícula requerida'
        }
        return errors
    }

    const onHandleSubmit = async (data) => {
        try{
            setLoading(true)
            const response = await updateUserByID(account.id, data)
            Alert.alert('¡Felicidades!', response)
            authDispatch({ type: 'update', user : { name : data.name }})
        }catch(error){
            Alert.alert('¡Ha ocurrido un error!', error.message)
        }finally{
            setLoading(false)
        }
    }

    return (
        <Formik 
            initialValues={initialValues}
            validate={handleValidation}
            onSubmit={onHandleSubmit}
            validateOnChange={false}
            validateOnBlur={false} 
            enableReinitialize
        >
            {({ handleChange, handleBlur, handleSubmit, errors, values }) => (
                <View>
                    <Input
                        label='Ingresa tu nombre'
                        onChangeText={handleChange('name')}
                        onBlur={handleBlur('name')}
                        value={values.name}
                        error={errors.name}
                    />
                    <Field name='birth_date'>
                        {({ field, form, meta }) => (
                            <Datepicker 
                                label='Selecciona tu fecha de nacimiento'
                                field={field} 
                                form={form} 
                                meta={meta}
                            />
                        )}
                    </Field>
                    <Input
                        label='Ingresa tu correo institucional'
                        onChangeText={handleChange('email')}
                        onBlur={handleBlur('email')}
                        value={values.email}
                        error={errors.email}
                    />
                    <Input
                        label='Ingresa tu matrícula'
                        onChangeText={handleChange('registration_number')}
                        onBlur={handleBlur('registration_number')}
                        value={values.registration_number}
                        error={errors.registration_number}
                        keyboardType='phone-pad'
                    />
                    <Button 
                        loading={loading}
                        onPress={handleSubmit}
                        loadingMessage='Actualizando cuenta...' 
                        style={styles.button}
                    >
                        Actualizar cuenta
                    </Button>
                </View>
            )}
        </Formik>
    )
}


export default withLoading(UpdateAccountConfigurationForm, 'Cargando información...')