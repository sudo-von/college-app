import React from 'react'
import { Text, StyleSheet, View } from 'react-native'
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

    const { state, dispatch } = useAuth()
    console.log(state)
    
    const onSubmit = async (data) => {
        try{
            const user = await login({email: data.email, password: data.password})
            dispatch({type: 'login', user})
        }catch(error){
            console.log(error)
            alert(error)
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
                    <Button onPress={handleSubmit}>Iniciar sesión</Button>
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
