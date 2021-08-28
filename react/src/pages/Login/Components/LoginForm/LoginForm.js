import React, { useState } from 'react'
import { StyleSheet, View } from 'react-native'
/* React native paper. */
import { TextInput } from 'react-native-paper'
/* Formik. */
import { Formik } from 'formik'
/* Custom components. */
import Button from 'src/components/Button'
import Input from 'src/components/Input'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'
/* Services. */
import { login } from 'src/services/user.service'

const LoginForm = () => {

    const { dispatch } = useAuth()

    const [ showPassword, setShowPassword ] = useState(false)

    const initialValues = { email: 'martinez-angel@uadec.edu.mx', password: 'college-app'}
    const onSubmit = async (data) => console.log(await login({email: data.email, password: data.password}))

    return (
        <Formik initialValues={initialValues} onSubmit={onSubmit}>
            {({ handleChange, handleBlur, handleSubmit, values }) => (
                <View style={styles.container}>
                    <Input
                        label='Ingresa tu correo'
                        onChangeText={handleChange('email')}
                        onBlur={handleBlur('email')}
                        value={values.email}
                    />
                    <Input
                        label='Ingresa tu contraseña'
                        onChangeText={handleChange('password')}
                        onBlur={handleBlur('password')}
                        value={values.password}
                        secureTextEntry={showPassword ? true : false}
                        right={<TextInput.Icon onPress={() => setShowPassword(!showPassword)} color='gray' name={showPassword ? 'eye-off' : 'eye'}/>}
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
