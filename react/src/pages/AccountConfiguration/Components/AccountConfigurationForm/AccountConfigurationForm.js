import React, { useEffect, useState } from 'react'
import { View, Alert, StyleSheet } from 'react-native'
/* Formik. */
import { Formik, Field } from 'formik'
/* Custom components. */
import Input from 'src/components/Input'
import Datepicker from 'src/components/Datepicker'
import Button from 'src/components/Button'
/* Services. */
import { getUserByID, updateUserByID } from 'src/services/user.service'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'

const AccountConfigurationForm = () => {

    /* Destructuring user properties. */
    const { authState, authDispatch } = useAuth()
    const { user : { user_id } } = authState
    const [ user, setUser ] = useState({ name: '', birth_date: '', email: '', registration_number: ''})
    const [ loading, setLoading ] = useState(false)

    /* Fetchs user data and then reinitialize form values. */
    useEffect(() => {
        const searchUser = async (userID) => {
            try{
                const response = await getUserByID(userID)
                setUser(response)
            }catch(error){
                Alert.alert('¡Ha ocurrido un error!', error.message)
            }
        }
        searchUser(user_id)
    },[])

    const onSubmit = async (data) => {
        try{
            setLoading(true)
            const response = await updateUserByID(user_id, data)
            setLoading(false)
            Alert.alert(response, 'El usuario ha sido actualizado con éxito.')
            authDispatch({ type: 'update', user: { name : data.name }})
        }catch(error){
            Alert.alert('¡Ha ocurrido un error!', error.message)
            setLoading(false)
        }
    }

    return (
        <Formik 
            initialValues={{ 
                name: user.name,
                birth_date: user.birth_date,
                email : user.email, 
                registration_number: user.registration_number
            }}
            validate={values => {
                const { name, birth_date, email, registration_number } = values
                const errors = {}
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
            }}
            onSubmit={onSubmit}
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

export default AccountConfigurationForm
