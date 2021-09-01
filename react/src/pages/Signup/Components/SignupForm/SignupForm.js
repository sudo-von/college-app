import React, { useState } from 'react'
import { View, Alert } from 'react-native'
/* Formik. */
import { Formik } from 'formik'
/* Custom components. */
import Button from 'src/components/Button'
import Input from 'src/components/Input'
import PasswordInput from 'src/components/PasswordInput'
/* Services. */
import { signup } from 'src/services/user.service'
/* React navigation. */
import { useNavigation } from '@react-navigation/native'


const SignupForm = () => {

    const navigation = useNavigation()
    
    const [ loading, setLoading ] = useState(false)

    const onSubmit = async (form) => {
        try{
            setLoading(true)
            const response = await signup(form)
            if(response){
                Alert.alert(
                    response,
                    "Ya puedes iniciar sesión."
                )
                navigation.navigate('/')
            }
        }catch(error){
            Alert.alert(
                "¡Ha ocurrido un error!",
                error.message
            )
        }finally{
            setLoading(false)
        }
    }

    return (
        <Formik 
            initialValues={{ 
                name: '',
                birth_date: '',
                email : '', 
                registration_number: '',
                university_id: '', 
                password: ''
            }}
            validate={values => {
                const { name, birth_date, email, registration_number, university_id, password } = values
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
                if(!university_id){
                    errors.university_id = 'Universidad requerida'
                }
                if(!password){
                    errors.password = 'Contraseña requerida'
                }
                return errors
            }}
            onSubmit={onSubmit}
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
                    <Input
                        label='Selecciona tu fecha de nacimiento'
                        onChangeText={handleChange('birth_date')}
                        onBlur={handleBlur('birth_date')}
                        value={values.birth_date}
                        error={errors.birth_date}
                    />
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
                    />
                    <Input
                        label='Selecciona tu facultad'
                        onChangeText={handleChange('university_id')}
                        onBlur={handleBlur('university_id')}
                        value={values.university_id}
                        error={errors.university_id}
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
                        loadingMessage='Registrando usuario...' 
                        onPress={handleSubmit}
                    >
                        Registrarse
                    </Button>
                </View>
            )}
        </Formik>
    )
}

export default SignupForm
