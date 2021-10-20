import React from 'react'
import { View } from 'react-native'
import { Formik } from 'formik'
import { Button, Input, PasswordInput } from 'src/components'
import { styles } from './LoginForm.styles'
import { useUser } from 'src/hooks/useUser'

const LoginForm = () => {

    const { loading, handleLogin } = useUser()

    const initialValues = { 
        email: 'martinez-angel@uadec.edu.mx', 
        password: 'college-app' 
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
            onSubmit={handleLogin}
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
