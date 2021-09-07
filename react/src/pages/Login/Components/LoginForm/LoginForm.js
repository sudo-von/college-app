import React, { useState } from 'react'
import { StyleSheet, View, Alert } from 'react-native'
/* Formik. */
import { Formik } from 'formik'
/* Custom components. */
import Button from 'src/components/Button'
import Input from 'src/components/Input'
import PasswordInput from 'src/components/PasswordInput'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'
/* Services. */
import { login } from 'src/services/user.service'

const LoginForm = () => {

    /* States. */
    const [ loading, setLoading ] = useState(false)
    const { authDispatch } = useAuth()

    const onSubmit = async (data) => {
        try{
            setLoading(true)
            const user = await login({email: data.email, password: data.password})
            setLoading(false)
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
                        label='Ingresa tu correo'
                        onChangeText={handleChange('email')}
                        onBlur={handleBlur('email')}
                        value={values.email}
                        error={errors.email}
                    />
                    <PasswordInput
                        label='Ingresa tu contraseña'
                        onChangeText={handleChange('password')}
                        onBlur={handleBlur('password')}
                        value={values.password}
                        error={errors.password}
                    />
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
