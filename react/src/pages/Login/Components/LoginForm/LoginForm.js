import React, { useState, useEffect } from 'react'
import { StyleSheet, View } from 'react-native'
/* Formik. */
import { Formik } from 'formik'
/* Custom components. */
import Button from 'src/components/Button'
import Input from 'src/components/Input'
import PasswordInput from 'src/components/PasswordInput'
import Alert from 'src/components/Alert'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'
/* Services. */
import { login } from 'src/services/user.service'

const LoginForm = () => {

    /* States. */
    const [ loading, setLoading ] = useState(false)
    const [ error, setError ] = useState(null) 
    const { authDispatch } = useAuth()

    const onSubmit = async (data) => {
        try{
            setLoading(true)
            setError(null)
            const user = await login({email: data.email, password: data.password})
            authDispatch({type: 'login', user})
        }catch(error){
            setError(error.message)
        }finally{
            setLoading(false)
        }
    }

    return (
        <Formik 
            initialValues={{ email: 'martinez-angel@uadec.edu.mx', password: 'college-app'}} 
            onSubmit={onSubmit}
        >
            {({ handleChange, handleBlur, handleSubmit, values }) => (
                <View style={styles.container}>
                    <Input
                        label='Ingresa tu correo'
                        onChangeText={handleChange('email')}
                        onBlur={handleBlur('email')}
                        value={values.email}
                    />
                    <PasswordInput
                        label='Ingresa tu contraseña'
                        onChangeText={handleChange('password')}
                        onBlur={handleBlur('password')}
                        value={values.password}
                    />
                    { error && <Alert title={error} type='error'/>}
                    <Button 
                        loading={loading} 
                        loadingMessage='Iniciando sesión' 
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
        width: '100%',
        marginTop: 30
    }
})

export default LoginForm
