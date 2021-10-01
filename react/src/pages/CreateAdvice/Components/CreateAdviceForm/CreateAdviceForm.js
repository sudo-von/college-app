import React, { useState } from 'react'
import { StyleSheet, View, Alert } from 'react-native'
/* Formik. */
import { Formik, Field } from 'formik'
/* Custom components. */
import Button from 'src/components/Button'
import Input from 'src/components/Input'
import Datepicker from 'src/components/Datepicker'
import SelectInput from 'src/components/SelectInput'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'
/* Services. */
import { login } from 'src/services/user.service'

const LoginForm = () => {

    /* States. */
    const [ loading, setLoading ] = useState(false)
    const { authDispatch } = useAuth()
    /* Gets user id. */
    const { authState } = useAuth()
    const { user } = authState
    const { university_id } = user

    const onSubmit = async (data) => {
        try{
            setLoading(true)
            const user = await login({email: data.email, password: data.password})
            setLoading(false)
            console.log(user)
            authDispatch({type: 'login', user})
        }catch(error){
            Alert.alert('',error.message)
            setLoading(false)
        }
    }

    return (
        <Formik 
            initialValues={{ email: 'martinez-angel@uadec.edu.mx', password: 'college-app'}} 
            validate={values => {
                const { email, password } = values
                const errors = {}
                if(!email){
                    errors.email = 'Correo requerido'
                }
                if(!password){
                    errors.password = 'Contraseña requerida'
                }
                return errors
            }}
            onSubmit={onSubmit}
        >
            {({ handleChange, handleBlur, handleSubmit, errors, values }) => (
                <View style={styles.container}>
                    <Input
                        label='Ingresa la materia'
                        onChangeText={handleChange('subject')}
                        onBlur={handleBlur('subject')}
                        value={values.subject}
                        error={errors.subject}
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
                    <Field name='university_id'>
                        {({ field, form, meta }) => (
                            <SelectInput
                                label='Selecciona tu facultad'
                                data={universityList}
                                field={field} 
                                form={form} 
                                meta={meta}
                                error={errors.university_id}
                            />
                        )}
                    </Field>
                    <Button 
                        loading={loading} 
                        loadingMessage='Iniciando sesión' 
                        style={styles.button}
                        onPress={handleSubmit}
                    >
                        Iniciar sesión
                    </Button>
                </View>
            )}
        </Formik>
    )
}

const styles = StyleSheet.create({
    container: {
        marginVertical: 30
    },
    button: {
        marginTop: 40,
    }
})

export default LoginForm
