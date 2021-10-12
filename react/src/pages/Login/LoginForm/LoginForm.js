import React, { useState } from 'react'
import { View, Alert } from 'react-native'
import { Formik } from 'formik'
import { Button, Input, PasswordInput } from 'src/components'
import { useAuth } from 'src/providers/auth.provider'
import { login } from 'src/services/auth.service'
import { styles } from './LoginForm.styles'

const LoginForm = () => {

    const [ loading, setLoading ] = useState(false)
    const { authDispatch } = useAuth()
    const initialValues = { email: 'martinez-angel@uadec.edu.mx', password: 'college-app' }

    const onSubmit = async (form) => {
        try{
            setLoading(true)
            const user = await login(form)
            setLoading(false)
            authDispatch({type: 'login', user})
        }catch(error){
            setLoading(false)
            Alert.alert('', error.message)
        }
    }

    const handleValidation = ({ email, password }) => {
        let errors = {}
        if(!email){
            errors.email = 'Correo requerido'
        }
        if(!password){
            errors.password = 'Contraseña requerida'
        }
        return errors
    }

    return (
        <Formik 
            initialValues={initialValues} 
            validate={handleValidation}
            validateOnChange={false}
            validateOnBlur={false} 
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

export default LoginForm
